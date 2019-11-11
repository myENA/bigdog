package vsz

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/myENA/ruckus-client/vsz/types/wsg/serviceticket"
)

const (
	serviceTicketQueryParameter = "serviceTicket"
)

type (
	AuthCAS uint64

	Authenticator interface {
		// Decorate must do one of two things:
		//
		// If the internal state of this authenticator is such that it currently has whatever is needed to modify a
		// given request with the appropriate authentication serviceTicket / token / header / etc., then it must do so,
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
		// 1. the  request created as part of api call chain
		Decorate(context.Context, *APIRequest) (AuthCAS, error)

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
		// 1. the current API client
		// 2. the CAS value seen from the calling routine's most recent action (could be from either Decorate or
		// Invalidate)
		Refresh(context.Context, *APIClient, AuthCAS) (AuthCAS, error)

		// Invalidate will only be called if a 401 is seen after a refresh has been attempted, and should indicate to
		// the implementor that whatever decoration the authenticator is current using is no longer considered valid by
		// the VSZ being queried
		//
		// Arguments:
		//
		// 0. the context provided to the original API call
		// 1. the current API client
		// 2. the CAS value seen from the calling routine's most recent action (could be from either Decorate or
		// Refresh)
		Invalidate(context.Context, *APIClient, AuthCAS) (AuthCAS, error)
	}
)

// PasswordAuthenticator is a simple example implementation of an Authenticator that will decorate a given request
// with a session id bearing serviceTicket if one exists, and attempt to create one if it doesn't.
type PasswordAuthenticator struct {
	username string
	password string

	mu sync.RWMutex

	cas           uint64
	refreshed     time.Time
	sessionTTL    time.Duration
	serviceTicket string

	logger *log.Logger
	debug  bool
}

func NewPasswordAuthenticator(username, password string, sessionTTL time.Duration, l *log.Logger, debug bool) *PasswordAuthenticator {
	pa := &PasswordAuthenticator{
		username:   username,
		password:   password,
		sessionTTL: sessionTTL,
		logger:     l,
		debug:      debug,
	}
	return pa
}

func (pa *PasswordAuthenticator) Username() string {
	return pa.username
}

func (pa *PasswordAuthenticator) Password() string {
	return pa.password
}

// Decorate will, if a current serviceTicket is found, add it to the provided request
func (pa *PasswordAuthenticator) Decorate(ctx context.Context, request *APIRequest) (AuthCAS, error) {
	pa.mu.RLock()
	cas := pa.cas

	if request == nil {
		// TODO: yell a bit more if request is nil?
		return AuthCAS(cas), errors.New("request cannot be nil")
	}

	pa.log(true, "Decorate called for request \"%s %s\"", request.Method(), request.CompiledURI())

	// is context still valid?
	if err := ctx.Err(); err != nil {
		pa.mu.RUnlock()
		return AuthCAS(cas), err
	}

	if pa.serviceTicket == "" && !pa.refreshed.IsZero() && pa.refreshed.Add(pa.sessionTTL).After(time.Now()) {
		request.SetQueryParameter(serviceTicketQueryParameter, pa.serviceTicket)
		pa.mu.RUnlock()
		return AuthCAS(cas), nil
	}

	pa.mu.RUnlock()

	return AuthCAS(cas), errors.New("serviceTicket requires refresh")
}

// Refresh will attempt to fetch a new serviceTicket from the VSZ for use in subsequent requests
func (pa *PasswordAuthenticator) Refresh(ctx context.Context, client *APIClient, cas AuthCAS) (AuthCAS, error) {
	pa.log(true, "Refresh called")

	pa.mu.Lock()
	ccas := pa.cas

	if client == nil {
		pa.mu.Unlock()
		return AuthCAS(ccas), errors.New("apiClient cannot be nil")
	}
	// is context still valid?
	if err := ctx.Err(); err != nil {
		pa.log(true, "Context error seen: %s", err)
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
	loginRequest := &serviceticket.LoginRequest{Username: &username, Password: &password}
	resp, err := client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	if err != nil {
		pa.log(true, "Error seen calling Login: %s", err)
		pa.serviceTicket = ""
		pa.cas++
		ncas := pa.cas
		pa.mu.Unlock()
		return AuthCAS(ncas), err
	}

	if resp == nil || resp.ServiceTicket == nil || *resp.ServiceTicket == "" {
		pa.log(true, "%q empty in response: %v", serviceTicketQueryParameter, resp)
		pa.serviceTicket = ""
		pa.cas++
		ncas := pa.cas
		pa.mu.Unlock()
		return AuthCAS(ncas), errors.New("nil response from login request seen")
	}

	pa.serviceTicket = *resp.ServiceTicket
	pa.refreshed = time.Now()
	pa.cas++
	ncas := pa.cas
	pa.mu.Unlock()
	return AuthCAS(ncas), nil
}

// Invalidate will mark the current session as invalid.
func (pa *PasswordAuthenticator) Invalidate(ctx context.Context, client *APIClient, cas AuthCAS) (AuthCAS, error) {
	pa.log(true, "Invalidate called", pa.username)

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

	// if we have a service ticket stored, attempt to invalidate it at the VSZ
	if pa.serviceTicket != "" {
		pa.log(true, "Calling Logout...")
		err := client.WSG().WSGServiceTicketService().DeleteServiceTicket(ctx, pa.serviceTicket)
		if err != nil {
			pa.log(false, "Error calling Logout: %s", err)
		} else {
			pa.log(true, "Logout called successfully")
		}
	}

	pa.cas++
	ncas := pa.cas
	pa.serviceTicket = ""
	pa.refreshed = time.Now()
	pa.mu.Unlock()
	return AuthCAS(ncas), nil
}

func (pa *PasswordAuthenticator) log(debug bool, f string, v ...interface{}) {
	if debug {
		if pa.debug {
			pa.logger.Printf(fmt.Sprintf("[pw-auth-%s] %s", pa.username, f), v...)
		}
	} else {
		pa.logger.Printf(fmt.Sprintf("[pw-auth-%s] %s", pa.username, f), v...)
	}
}
