package vsz

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	SessionCookieToken = "JSESSIONID"
)

type (
	AuthCAS uint64

	Authenticator interface {
		// Decorate must do one of two things:
		//
		// If the internal state of this authenticator is such that it currently has whatever is needed to modify a
		// given request with the appropriate authentication cookie / token / header / etc., then it must do so,
		// returning the current CAS and a nil error
		//
		// If the internal state is such that decoration CANNOT happen, it must return the current CAS and an error,
		// describing the reason it cannot decorate the request.  If the error is not nil, Refresh will be called with
		// the CAS value returned by this method.
		//
		// In all cases, the current CAS must be returned.
		//
		// Arguments:
		//
		// 0. the context provided to original API call
		// 1. the raw http request created as part of api call chain
		Decorate(context.Context, *http.Request) (AuthCAS, error)

		// Refresh must do one of two things:
		//
		// If the provided CAS value is current, it must assume that its current state is no longer valid and try to do
		// what is needed to get back to a state that Decorate is able to do what it needs to do.
		//
		// If the provided CAS value is NOT equal to the current state, it must assume that a refresh attempt has
		// already been made in another process, and just return the current CAS value with no error.
		//
		// The client will only attempt a maximum of 2 times per execution:
		//
		// 1. If the FIRST Decorate call fails
		// 2. If initial Decorate did not fail but VSZ returned a 401, causing an Invalidate -> Refresh -> Decorate loop
		// that will execute exactly 1 times.
		//
		// Arguments:
		//
		// 0. the context provided to the original API call
		// 1. the current VSZ API client
		// 2. the CAS value seen from the calling routine's most recent action (could be from either Decorate or
		// Invalidate)
		Refresh(context.Context, *Client, AuthCAS) (AuthCAS, error)

		// Invalidate will only be called if a 401 is seen after a refresh has been attempted, and should indicate to
		// the implementor that whatever decoration the authenticator is current using is no longer considered valid by
		// the VSZ being queried
		//
		// Arguments:
		//
		// 0. the context provided to the original API call
		// 1. the CAS value seen from the calling routine's most recent action (could be from either Decorate or
		// Refresh)
		Invalidate(context.Context, AuthCAS) (AuthCAS, error)
	}
)

// PasswordAuthenticator is a simple example implementation of an Authenticator that will decorate a given request
// with a session id bearing cookie if one exists, and attempt to create one if it doesn't.
type PasswordAuthenticator struct {
	username string
	password string

	mu sync.RWMutex

	cas       uint64
	refreshed time.Time
	cookieTTL time.Duration
	cookie    *http.Cookie

	debug bool
}

func NewPasswordAuthenticator(username, password string, cookieTTL time.Duration, debug bool) *PasswordAuthenticator {
	pa := &PasswordAuthenticator{
		username:  username,
		password:  password,
		cookieTTL: cookieTTL,
		debug:     debug,
	}
	return pa
}

func (pa *PasswordAuthenticator) Username() string {
	return pa.username
}

func (pa *PasswordAuthenticator) Password() string {
	return pa.password
}

func (pa *PasswordAuthenticator) Decorate(ctx context.Context, httpRequest *http.Request) (AuthCAS, error) {
	if pa.debug {
		log.Printf("[pw-auth-%s] Decorate called for request \"%s %s\"", pa.username, httpRequest.Method, httpRequest.URL)
	}

	pa.mu.RLock()
	cas := pa.cas

	if httpRequest == nil {
		// TODO: yell a bit more if the request is nil?
		pa.mu.RUnlock()
		return AuthCAS(cas), errors.New("httpRequest cannot be nil")
	}
	// is context still valid?
	if err := ctx.Err(); err != nil {
		pa.mu.RUnlock()
		return AuthCAS(cas), err
	}

	cookie := pa.cookie
	// TODO improve efficiency of this?
	if cookie != nil && !pa.refreshed.IsZero() && pa.refreshed.Add(pa.cookieTTL).After(time.Now()) {
		httpRequest.AddCookie(cookie)
		pa.mu.RUnlock()
		return AuthCAS(cas), nil
	}
	pa.mu.RUnlock()
	return AuthCAS(cas), errors.New("cookie requires refresh")
}

func (pa *PasswordAuthenticator) Refresh(ctx context.Context, client *Client, cas AuthCAS) (AuthCAS, error) {
	if pa.debug {
		log.Printf("[pw-auth-%s] Refresh called", pa.username)
	}

	pa.mu.Lock()
	ccas := pa.cas

	if client == nil {
		pa.mu.Unlock()
		return AuthCAS(ccas), errors.New("client cannot be nil")
	}
	// is context still valid?
	if err := ctx.Err(); err != nil {
		pa.mu.Unlock()
		return AuthCAS(ccas), err
	}
	// if the passed cas value is greater than the internal CAS, assume weirdness and return current CAS and an error
	if ccas < uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value is greater than possible")
	}
	// if the passed in CAS value is less than the currently stored one, assume another routine called either Refresh
	// or Invalidate and just return current cas
	if ccas > uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), nil
	}

	// if cas matches internal...

	// try to execute logon
	username := pa.username
	password := pa.password
	resp, _, err := client.Session().LoginSessionLogonPost(ctx, &LoginSessionLogonPostRequest{
		Username: &username,
		Password: &password,
	})

	if resp != nil {
		resp.Body.Close()
	}

	if err != nil {
		pa.cookie = nil
		pa.cas++
		ncas := pa.cas
		pa.mu.Unlock()
		return AuthCAS(ncas), err
	}

	cookie := TryExtractSessionCookie(resp)
	if cookie == nil {
		pa.cookie = nil
		pa.cas++
		ncas := pa.cas
		pa.mu.Unlock()
		return AuthCAS(ncas), errors.New("unable to locate cookie in response")
	}

	pa.cookie = cookie
	pa.refreshed = time.Now()
	pa.cas++
	ncas := pa.cas
	pa.mu.Unlock()
	return AuthCAS(ncas), nil
}

func (pa *PasswordAuthenticator) Invalidate(ctx context.Context, cas AuthCAS) (AuthCAS, error) {
	if pa.debug {
		log.Printf("[pw-auth-%s] Invalidate called", pa.username)
	}

	pa.mu.Lock()
	ccas := pa.cas

	// is context still valid?
	if err := ctx.Err(); err != nil {
		pa.mu.Unlock()
		return AuthCAS(ccas), err
	}
	// if current cas is less than provided, assume insanity.
	if ccas < uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value greater than possible")
	}
	// if current cas is greater than provided, assume Refresh or Invalidate has already been called.
	if ccas > uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), nil
	}

	pa.cas++
	ncas := pa.cas
	pa.cookie = nil
	pa.refreshed = time.Now()
	pa.mu.Unlock()
	return AuthCAS(ncas), nil
}
