package bigdog

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type VSZServiceTicketCAS time.Duration

type VSZServiceTicketResponse interface {
	// CAS must return the CAS value as it stands at the time of this response
	CAS() VSZServiceTicketCAS
	// Ticket must return the current service ticket if it is still valid and no error occurred
	Ticket() string
}

type BaseVSZServiceTicketResponse struct {
	cas    VSZServiceTicketCAS
	ticket string
}

func NewBaseVSZServiceTicketResponse(cas VSZServiceTicketCAS, ticket string) BaseVSZServiceTicketResponse {
	str := BaseVSZServiceTicketResponse{
		cas:    cas,
		ticket: ticket,
	}
	return str
}

func (str BaseVSZServiceTicketResponse) CAS() VSZServiceTicketCAS {
	return str.cas
}

func (str BaseVSZServiceTicketResponse) Ticket() string {
	return str.ticket
}

type UsernamePasswordVSZTicketResponse struct {
	BaseVSZServiceTicketResponse
	meta APIResponseMeta
}

func newUsernamePasswordVSZTicketProviderResponse(cas VSZServiceTicketCAS, ticket string, meta APIResponseMeta) UsernamePasswordVSZTicketResponse {
	str := UsernamePasswordVSZTicketResponse{
		BaseVSZServiceTicketResponse: NewBaseVSZServiceTicketResponse(cas, ticket),
		meta:                         meta,
	}
	return str
}

func (str UsernamePasswordVSZTicketResponse) ResponseMeta() APIResponseMeta {
	return str.meta
}

type VSZServiceTicketProvider interface {
	// Current must return the current cas and ticket value, if and only if it is still valid.
	Current() (VSZServiceTicketResponse, error)

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
	// 2. If initial Current did not fail but VSZ returned a 401, causing an Invalidate -> Refresh -> Current chain
	// that will execute exactly 1 times.
	//
	// Arguments:
	//
	// 0. the context provided to the original API call
	// 1. the current VSZ client
	// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or
	// Invalidate)
	Refresh(context.Context, *VSZClient, VSZServiceTicketCAS) (VSZServiceTicketResponse, error)

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
	Invalidate(context.Context, *VSZClient, VSZServiceTicketCAS) (VSZServiceTicketResponse, error)
}

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
func (stp *UsernamePasswordVSZSServiceTicketProvider) Current() (VSZServiceTicketResponse, error) {
	stp.mu.RLock()
	defer stp.mu.RUnlock()

	if stp.serviceTicket != "" && time.Now().Before(stp.expires) {
		return newUsernamePasswordVSZTicketProviderResponse(stp.cas, stp.serviceTicket, APIResponseMeta{}), nil
	}

	return newUsernamePasswordVSZTicketProviderResponse(stp.cas, "", APIResponseMeta{}), ErrServiceTicketRequiresRefresh
}

// Refresh will attempt to fetch a new serviceTicket from the VSZ for use in subsequent requests
func (stp *UsernamePasswordVSZSServiceTicketProvider) Refresh(ctx context.Context, client *VSZClient, cas VSZServiceTicketCAS) (VSZServiceTicketResponse, error) {
	stp.mu.Lock()
	defer stp.mu.Unlock()

	var (
		loginRequest  *WSGServiceTicketLoginRequest
		loginResponse *WSGServiceTicketLoginResponseAPIResponse
		loginMeta     APIResponseMeta
		rm            *APIResponseMeta
		err           error

		username = stp.username
		password = stp.password
	)

	if client == nil {
		return newUsernamePasswordVSZTicketProviderResponse(stp.cas, "", loginMeta), ErrServiceTicketClientNil
	}

	// if the passed cas value is greater than the internal CAS, assume weirdness and return current CAS and an error
	if stp.cas < cas {
		return newUsernamePasswordVSZTicketProviderResponse(stp.cas, "", loginMeta), fmt.Errorf("%w: provided cas value is greater than possible", ErrServiceTicketCASInvalid)
	}
	// if the passed in CAS value is less than the currently stored one, assume another routine called either Refresh
	// or Invalidate and just return current cas
	if stp.cas > cas {
		// todo: is this ok...?
		return newUsernamePasswordVSZTicketProviderResponse(stp.cas, stp.serviceTicket, loginMeta), nil
	}

	// if cas matches internal...

	// try to execute logon
	loginRequest = &WSGServiceTicketLoginRequest{Username: &username, Password: &password}

	// retry logic...
	loginResponse, err = client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	for i := uint(1); err != nil && i <= stp.retries; i++ {
		time.Sleep(stp.retryWait)
		loginResponse, err = client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	}

	if loginResponse != nil {
		loginMeta = loginResponse.ResponseMeta()
	}

	if err != nil {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return newUsernamePasswordVSZTicketProviderResponse(stp.cas, "", loginMeta), err
	}

	if loginResponse == nil {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return newUsernamePasswordVSZTicketProviderResponse(stp.cas, "", loginMeta), ErrServiceTicketResponseEmpty
	}

	// test if we need to try to hydrate the response model
	if loginResponse.Data == nil {
		// attempt model hydration
		if err = loginResponse.Hydrate(); err != nil {
			stp.cas = stp.iterateCAS()
			stp.serviceTicket = ""
			stp.expires = time.Time{}
			return newUsernamePasswordVSZTicketProviderResponse(stp.cas, "", loginMeta), fmt.Errorf("%w: error unmarshalling response body: %v", ErrServiceTicketResponseEmpty, err)
		}
	}

	// this case is extremely unlikely, but keep it here just in case.  don't want no npe.
	if loginResponse.Data.ServiceTicket == nil || *loginResponse.Data.ServiceTicket == "" {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return newUsernamePasswordVSZTicketProviderResponse(stp.cas, "", loginMeta), ErrServiceTicketResponseEmpty
	}

	// if reached, assume login successful
	stp.cas = stp.iterateCAS()
	stp.serviceTicket = *loginResponse.Data.ServiceTicket
	stp.expires = time.Now().Add(stp.sessionTTL)

	return newUsernamePasswordVSZTicketProviderResponse(stp.cas, stp.serviceTicket, loginMeta), nil
}

// Invalidate will mark the current session as invalid.
func (stp *UsernamePasswordVSZSServiceTicketProvider) Invalidate(ctx context.Context, client *VSZClient, cas VSZServiceTicketCAS) (VSZServiceTicketResponse, error) {
	stp.mu.Lock()
	defer stp.mu.Unlock()

	var (
		rm  *APIResponseMeta
		err error
	)

	// if current cas is less than provided, assume insanity.
	if stp.cas < cas {
		return stp.cas, NewVSZServiceTicketProviderError(newErrAPIResponseMeta(), fmt.Errorf("%w: provided cas value greater than possible", ErrServiceTicketCASInvalid))
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
