package vsz

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	ServiceTicketCAS time.Duration

	ServiceTicketProvider interface {
		// Current must do one of two things:
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
		Current() (ServiceTicketCAS, string, error)

		// Refresh must do one of two things:
		//
		// If the provided CAS value is current, it must assume that its current state is no longer valid and try to do
		// what is needed to get back to a state that Current is able to do what it needs to do.
		//
		// If the provided CAS value is NOT equal to the current state, it must assume that a refresh attempt has
		// already been made in another process, and just return the current CAS value with no error.
		//
		// The client will only attempt a maximum of 2 times per execution:
		//
		// 1. If the FIRST Current call fails
		// 2. If initial Current did not fail but VSZ returned a 401, causing an Invalidate -> Refresh -> Current loop
		// that will execute exactly 1 times.
		//
		// Arguments:
		//
		// 0. the context provided to the original API call
		// 1. the current API client
		// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or
		// Invalidate)
		Refresh(context.Context, *APIClient, ServiceTicketCAS) (ServiceTicketCAS, error)

		// Invalidate will only be called if a 401 is seen after a refresh has been attempted, and should indicate to
		// the implementor that whatever decoration the authenticator is current using is no longer considered valid by
		// the VSZ being queried
		//
		// Arguments:
		//
		// 0. the context provided to the original API call
		// 1. the current API client
		// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or
		// Refresh)
		Invalidate(context.Context, *APIClient, ServiceTicketCAS) (ServiceTicketCAS, error)
	}
)

// UsernamePasswordServiceTicketProvider is a simple example implementation of an ServiceTicketProvider that will
// attempt to automatically  handle the login / logout cycle for api requests using a username and password.
type UsernamePasswordServiceTicketProvider struct {
	username string
	password string

	mu sync.RWMutex

	cas           ServiceTicketCAS
	expires       time.Time
	sessionTTL    time.Duration
	serviceTicket string
}

func NewUsernamePasswordServiceTicketProvider(username, password string, sessionTTL time.Duration) *UsernamePasswordServiceTicketProvider {
	a := new(UsernamePasswordServiceTicketProvider)
	a.username = username
	a.password = password
	a.sessionTTL = sessionTTL
	return a
}

func (a *UsernamePasswordServiceTicketProvider) Username() string {
	return a.username
}

func (a *UsernamePasswordServiceTicketProvider) Password() string {
	return a.password
}

// Current returns the current service ticket, if there is one.
func (a *UsernamePasswordServiceTicketProvider) Current() (ServiceTicketCAS, string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.serviceTicket != "" && time.Now().Before(a.expires) {
		return a.cas, a.serviceTicket, nil
	}

	return a.cas, "", ErrServiceTicketRequiresRefresh
}

// Refresh will attempt to fetch a new serviceTicket from the VSZ for use in subsequent requests
func (a *UsernamePasswordServiceTicketProvider) Refresh(ctx context.Context, client *APIClient, cas ServiceTicketCAS) (ServiceTicketCAS, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var (
		loginRequest  *WSGServiceTicketLoginRequest
		loginResponse *WSGServiceTicketLoginResponse
		rm            *APIResponseMeta
		err           error

		username = a.username
		password = a.password
	)

	if client == nil {
		return a.cas, NewServiceTicketProviderError(newErrAPIResponseMeta(), ErrServiceTicketClientNil)
	}

	// if the passed cas value is greater than the internal CAS, assume weirdness and return current CAS and an error
	if a.cas < cas {
		return a.cas, NewServiceTicketProviderError(newErrAPIResponseMeta(), fmt.Errorf("%w: provided cas value is greater than possible", ErrServiceTicketCASInvalid))
	}
	// if the passed in CAS value is less than the currently stored one, assume another routine called either Refresh
	// or Invalidate and just return current cas
	if a.cas > cas {
		// todo: is this ok...?
		return a.cas, nil
	}

	// if cas matches internal...

	// try to execute logon
	loginRequest = &WSGServiceTicketLoginRequest{Username: &username, Password: &password}
	if loginResponse, rm, err = client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest); err != nil {
		a.cas = a.iterateCAS()
		a.serviceTicket = ""
		a.expires = time.Time{}
		return a.cas, NewServiceTicketProviderError(rm, err)
	}

	// this case is extremely unlikely, but keep it here just in case.  don't want no npe.
	if loginResponse == nil || loginResponse.ServiceTicket == nil || *loginResponse.ServiceTicket == "" {
		a.cas = a.iterateCAS()
		a.serviceTicket = ""
		a.expires = time.Time{}
		return a.cas, NewServiceTicketProviderError(rm, ErrServiceTicketResponseEmpty)
	}

	// if reached, assume login successful
	a.serviceTicket = *loginResponse.ServiceTicket
	a.cas = a.iterateCAS()
	a.expires = time.Now().Add(a.sessionTTL)

	return a.cas, nil
}

// Invalidate will mark the current session as invalid.
func (a *UsernamePasswordServiceTicketProvider) Invalidate(ctx context.Context, client *APIClient, cas ServiceTicketCAS) (ServiceTicketCAS, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var (
		rm  *APIResponseMeta
		err error
	)

	// if current cas is less than provided, assume insanity.
	if a.cas < cas {
		return a.cas, NewServiceTicketProviderError(newErrAPIResponseMeta(), fmt.Errorf("%w: provided cas value greater than possible", ErrServiceTicketCASInvalid))
	}

	// if current cas is greater than provided, assume Refresh or Invalidate has already been called.
	if a.cas > cas {
		// todo: is this ok?
		return a.cas, nil
	}

	// if we have a service ticket stored, attempt to invalidate it at the VSZ
	if a.serviceTicket != "" {
		if _, rm, err = client.WSG().WSGServiceTicketService().DeleteServiceTicket(ctx, a.serviceTicket); err != nil {
			err = NewServiceTicketProviderError(rm, err)
		}

		// update internal state
		a.serviceTicket = ""
		a.cas = a.iterateCAS()
		a.expires = time.Time{}
	}

	return a.cas, err
}

func (a *UsernamePasswordServiceTicketProvider) iterateCAS() ServiceTicketCAS {
	return ServiceTicketCAS(time.Now().UnixNano())
}
