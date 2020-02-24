package vsz

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

const (
	serviceTicketQueryParameter    = "serviceTicket"
	serviceTicketAuthLogSlugFormat = "[ticket-auth-%s] %s"
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

// ServiceTicketAuthenticator is a simple example implementation of an Authenticator that will attempt to automatically
// handle the login / logout cycle for api requests using a username and password.
type ServiceTicketAuthenticator struct {
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

func NewServiceTicketAuthenticator(username, password string, sessionTTL time.Duration, l *log.Logger, debug bool) *ServiceTicketAuthenticator {
	pa := &ServiceTicketAuthenticator{
		username:   username,
		password:   password,
		sessionTTL: sessionTTL,
		logger:     l,
		debug:      debug,
	}
	return pa
}

func (st *ServiceTicketAuthenticator) Username() string {
	return st.username
}

func (st *ServiceTicketAuthenticator) Password() string {
	return st.password
}

// Decorate will, if a current serviceTicket is found, add it to the provided request
func (st *ServiceTicketAuthenticator) Decorate(ctx context.Context, request *APIRequest) (AuthCAS, error) {
	st.mu.RLock()
	cas := st.cas

	if request == nil {
		// TODO: yell a bit more if request is nil?
		return AuthCAS(cas), errors.New("request cannot be nil")
	}

	st.log(true, "Decorate() - called for request \"%s %s\"", request.Method(), request.CompiledURI())

	// is context still valid?
	if err := ctx.Err(); err != nil {
		st.mu.RUnlock()
		return AuthCAS(cas), err
	}

	if st.serviceTicket == "" && !st.refreshed.IsZero() && st.refreshed.Add(st.sessionTTL).After(time.Now()) {
		request.SetQueryParameter(serviceTicketQueryParameter, []string{st.serviceTicket})
		st.mu.RUnlock()
		return AuthCAS(cas), nil
	}

	st.mu.RUnlock()

	return AuthCAS(cas), errors.New("serviceTicket requires refresh")
}

// Refresh will attempt to fetch a new serviceTicket from the VSZ for use in subsequent requests
func (st *ServiceTicketAuthenticator) Refresh(ctx context.Context, client *APIClient, cas AuthCAS) (AuthCAS, error) {
	st.log(true, "Refresh() - called")

	st.mu.Lock()
	ccas := st.cas

	if client == nil {
		st.mu.Unlock()
		return AuthCAS(ccas), errors.New("apiClient cannot be nil")
	}
	// is context still valid?
	if err := ctx.Err(); err != nil {
		st.log(true, "Refresh() - Context error seen: %s", err)
		st.mu.Unlock()
		return AuthCAS(ccas), err
	}
	// if the passed cas value is greater than the internal CAS, assume weirdness and return current CAS and an error
	if ccas < uint64(cas) {
		st.mu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value is greater than possible")
	}
	// if the passed in CAS value is less than the currently stored one, assume another routine called either Refresh
	// or Invalidate and just return current cas
	if ccas > uint64(cas) {
		st.mu.Unlock()
		return AuthCAS(ccas), nil
	}

	// if cas matches internal...

	// try to execute logon
	username := st.username
	password := st.password
	loginRequest := &WSGServiceTicketLoginRequest{Username: &username, Password: &password}
	resp, _, err := client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	if err != nil {
		st.log(true, "Refresh() - Error seen calling Login: %s", err)
		st.serviceTicket = ""
		ncas := st.iterateCAS()
		st.mu.Unlock()
		return AuthCAS(ncas), err
	}

	if resp == nil || resp.ServiceTicket == nil || *resp.ServiceTicket == "" {
		st.log(true, "Refresh() - %q empty in response: %v", serviceTicketQueryParameter, resp)
		st.serviceTicket = ""
		ncas := st.iterateCAS()
		st.mu.Unlock()
		return AuthCAS(ncas), errors.New("nil response from login request seen")
	}

	st.serviceTicket = *resp.ServiceTicket
	st.refreshed = time.Now()
	ncas := st.iterateCAS()
	st.mu.Unlock()
	return AuthCAS(ncas), nil
}

// Invalidate will mark the current session as invalid.
func (st *ServiceTicketAuthenticator) Invalidate(ctx context.Context, client *APIClient, cas AuthCAS) (AuthCAS, error) {
	st.log(true, "Invalidate called", st.username)

	st.mu.Lock()
	ccas := st.cas

	// is context still valid?
	if err := ctx.Err(); err != nil {
		st.mu.Unlock()
		return AuthCAS(ccas), err
	}
	// if current cas is less than provided, assume insanity.
	if ccas < uint64(cas) {
		st.mu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value greater than possible")
	}
	// if current cas is greater than provided, assume Refresh or Invalidate has already been called.
	if ccas > uint64(cas) {
		st.mu.Unlock()
		return AuthCAS(ccas), nil
	}

	// if we have a service ticket stored, attempt to invalidate it at the VSZ
	if st.serviceTicket != "" {
		resp, _, err := client.WSG().WSGServiceTicketService().DeleteServiceTicket(ctx, st.serviceTicket)
		st.log(true, "Invalidate() - DeleteServiceTicket response: %v", resp)
		if err != nil {
			st.log(false, "Invalidate() - Error calling DeleteServiceTicket: %s", err)
		} else {
			st.log(true, "Invalidate() - DeleteServiceTicket called successfully")
		}
	}

	ncas := st.iterateCAS()
	st.serviceTicket = ""
	st.refreshed = time.Now()
	st.mu.Unlock()
	return AuthCAS(ncas), nil
}

func (st *ServiceTicketAuthenticator) iterateCAS() uint64 {
	return atomic.AddUint64(&st.cas, uint64(time.Now().UnixNano()))
}

func (st *ServiceTicketAuthenticator) log(debug bool, f string, v ...interface{}) {
	if debug {
		if st.debug {
			st.logger.Printf(fmt.Sprintf(serviceTicketAuthLogSlugFormat, st.username, f), v...)
		}
	} else {
		st.logger.Printf(fmt.Sprintf(serviceTicketAuthLogSlugFormat, st.username, f), v...)
	}
}
