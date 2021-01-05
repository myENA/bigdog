package bigdog

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	VSZServiceTicketCAS time.Duration

	VSZServiceTicketProvider interface {
		Current() (VSZServiceTicketCAS, string, error)

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
		// 1. the current VSZ client
		// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or
		// Invalidate)
		Refresh(context.Context, *VSZClient, VSZServiceTicketCAS) (VSZServiceTicketCAS, error)

		// Invalidate will only be called if a 401 is seen after a refresh has been attempted, and should indicate to
		// the implementor that whatever decoration the authenticator is current using is no longer considered valid by
		// the VSZ being queried
		//
		// Arguments:
		//
		// 0. the context provided to the original API call
		// 1. the current VSZ client
		// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or
		// Refresh)
		Invalidate(context.Context, *VSZClient, VSZServiceTicketCAS) (VSZServiceTicketCAS, error)
	}
)

// UsernamePasswordVSZSServiceTicketProvider is a simple example implementation of an VSZServiceTicketProvider that will
// attempt to automatically  handle the login / logout cycle for api requests using a username and password.
type UsernamePasswordVSZSServiceTicketProvider struct {
	mu sync.RWMutex

	username   string
	password   string
	retries    uint
	retryWait  time.Duration
	sessionTTL time.Duration

	cas           VSZServiceTicketCAS
	expires       time.Time
	serviceTicket string
}

func NewUsernamePasswordServiceTicketProvider(username, password string, retries uint, retryWait, sessionTTL time.Duration) *UsernamePasswordVSZSServiceTicketProvider {
	stp := new(UsernamePasswordVSZSServiceTicketProvider)
	stp.username = username
	stp.password = password
	stp.retries = retries
	stp.retryWait = retryWait
	stp.sessionTTL = sessionTTL
	return stp
}

func (stp *UsernamePasswordVSZSServiceTicketProvider) Username() string {
	return stp.username
}

func (stp *UsernamePasswordVSZSServiceTicketProvider) Password() string {
	return stp.password
}

func (stp *UsernamePasswordVSZSServiceTicketProvider) SessionTTL() time.Duration {
	return stp.sessionTTL
}

// Current returns the current service ticket, if there is one.
func (stp *UsernamePasswordVSZSServiceTicketProvider) Current() (VSZServiceTicketCAS, string, error) {
	stp.mu.RLock()
	defer stp.mu.RUnlock()

	if stp.serviceTicket != "" && time.Now().Before(stp.expires) {
		return stp.cas, stp.serviceTicket, nil
	}

	return stp.cas, "", NewVSZServiceTicketProviderError(NewErrAPIResponseMeta(), ErrServiceTicketRequiresRefresh)
}

// Refresh will attempt to fetch a new serviceTicket from the VSZ for use in subsequent requests
func (stp *UsernamePasswordVSZSServiceTicketProvider) Refresh(ctx context.Context, client *VSZClient, cas VSZServiceTicketCAS) (VSZServiceTicketCAS, error) {
	stp.mu.Lock()
	defer stp.mu.Unlock()

	var (
		loginRequest  *WSGServiceTicketLoginRequest
		loginResponse *WSGServiceTicketLoginResponse
		rm            *APIResponseMeta
		err           error

		username = stp.username
		password = stp.password
	)

	if client == nil {
		return stp.cas, NewVSZServiceTicketProviderError(NewErrAPIResponseMeta(), ErrServiceTicketClientNil)
	}

	// if the passed cas value is greater than the internal CAS, assume weirdness and return current CAS and an error
	if stp.cas < cas {
		return stp.cas, NewVSZServiceTicketProviderError(NewErrAPIResponseMeta(), fmt.Errorf("%w: provided cas value is greater than possible", ErrServiceTicketCASInvalid))
	}
	// if the passed in CAS value is less than the currently stored one, assume another routine called either Refresh
	// or Invalidate and just return current cas
	if stp.cas > cas {
		// todo: is this ok...?
		return stp.cas, nil
	}

	// if cas matches internal...

	// try to execute logon
	loginRequest = &WSGServiceTicketLoginRequest{Username: &username, Password: &password}

	// retry logic...
	loginResponse, rm, err = client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	for i := uint(1); err != nil && i <= stp.retries; i++ {
		time.Sleep(stp.retryWait)
		loginResponse, rm, err = client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	}

	if err != nil {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return stp.cas, NewVSZServiceTicketProviderError(rm, err)
	}

	// this case is extremely unlikely, but keep it here just in case.  don't want no npe.
	if loginResponse == nil || loginResponse.ServiceTicket == nil || *loginResponse.ServiceTicket == "" {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return stp.cas, NewVSZServiceTicketProviderError(rm, ErrServiceTicketResponseEmpty)
	}

	// if reached, assume login successful
	stp.cas = stp.iterateCAS()
	stp.serviceTicket = *loginResponse.ServiceTicket
	stp.expires = time.Now().Add(stp.sessionTTL)

	return stp.cas, nil
}

// Invalidate will mark the current session as invalid.
func (stp *UsernamePasswordVSZSServiceTicketProvider) Invalidate(ctx context.Context, client *VSZClient, cas VSZServiceTicketCAS) (VSZServiceTicketCAS, error) {
	stp.mu.Lock()
	defer stp.mu.Unlock()

	var (
		rm  *APIResponseMeta
		err error
	)

	// if current cas is less than provided, assume insanity.
	if stp.cas < cas {
		return stp.cas, NewVSZServiceTicketProviderError(NewErrAPIResponseMeta(), fmt.Errorf("%w: provided cas value greater than possible", ErrServiceTicketCASInvalid))
	}

	// if current cas is greater than provided, assume Refresh or Invalidate has already been called.
	if stp.cas > cas {
		return stp.cas, nil
	}

	// if we have a service ticket stored, attempt to invalidate it at the VSZ
	if stp.serviceTicket != "" {
		if _, rm, err = client.WSG().WSGServiceTicketService().DeleteServiceTicket(ctx, stp.serviceTicket); err != nil {
			err = NewVSZServiceTicketProviderError(rm, err)
		}

		// update internal state
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
	}

	return stp.cas, err
}

func (*UsernamePasswordVSZSServiceTicketProvider) iterateCAS() VSZServiceTicketCAS {
	return VSZServiceTicketCAS(time.Now().UnixNano())
}
