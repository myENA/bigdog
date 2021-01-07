package bigdog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	ErrServiceTicketRequiresRefresh = errors.New("service ticket requires refresh")
	ErrServiceTicketCASInvalid      = errors.New("invalid CAS provided")
	ErrServiceTicketClientNil       = errors.New("client cannot be nil")
	ErrServiceTicketResponseEmpty   = errors.New("empty response from login request")

	serviceTicketErrors = []error{
		ErrServiceTicketRequiresRefresh,
		ErrServiceTicketCASInvalid,
		ErrServiceTicketClientNil,
		ErrServiceTicketResponseEmpty,
	}

	serviceTicketFatalErrors = []error{
		ErrServiceTicketCASInvalid,
		ErrServiceTicketClientNil,
		ErrServiceTicketResponseEmpty,
	}
)

type VSZAPIError struct {
	Success bool `json:"success"`

	Message      string      `json:"message,omitempty"`
	NoSession    string      `json:"noSession,omitempty"`
	ErrorDetails interface{} `json:"error,omitempty"` // probably can be map[string]interface{}
	ErrorCode    int         `json:"errorCode,omitempty"`
	ErrorType    string      `json:"errorType,omitempty"`
	Extra        interface{} `json:"extra,omitempty"`
	Metadata     interface{} `json:"metadata,omitempty"`
	Data         interface{} `json:"data,omitempty"`

	Err error `json:"err,omitempty"`
}

func (e *VSZAPIError) Is(err error) bool {
	return e != nil && e.Err != nil && errors.Is(err, e.Err)
}

func (e *VSZAPIError) Unwrap() error {
	if e == nil || e.Err == nil {
		return nil
	}
	return e.Err
}

func (e *VSZAPIError) Error() string {
	if e == nil {
		return ""
	}
	if e.ErrorDetails != nil {
		return fmt.Sprintf("success=%t; errorDetails=%v", e.Success, e.ErrorDetails)
	}
	return fmt.Sprintf("success=%t; errorCode=%q; errorType=%q; message=%q; err=%v", e.Success, e.ErrorCode, e.ErrorType, e.Message, e.Err)
}

func IsVSZAPIError(err error) bool {
	for err != nil {
		if verr, ok := err.(*VSZAPIError); ok && verr != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

func IsServiceTicketProviderError(err error) bool {
	for _, sterr := range serviceTicketErrors {
		if errors.Is(err, sterr) {
			return true
		}
	}
	return false
}

func IsFatalServiceTicketProviderError(err error) bool {
	for _, sterr := range serviceTicketFatalErrors {
		if errors.Is(err, sterr) {
			return true
		}
	}
	return false
}

type VSZServiceTicketCAS time.Duration

type VSZServiceTicketProvider interface {
	// Current must return the current cas and ticket value, if and only if it is still valid.
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
	// 2. If initial Current did not fail but VSZ returned a 401, causing an Invalidate -> Refresh -> Current chain
	// that will execute exactly 1 times.
	//
	// Arguments:
	//
	// 0. the context provided to the original API call
	// 1. the current VSZ client
	// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or Invalidate)
	Refresh(context.Context, *VSZClient, VSZServiceTicketCAS) (VSZServiceTicketCAS, error)

	// Invalidate will only be called if a 401 is seen after a refresh has been attempted, and should indicate to
	// the implementor that whatever decoration the authenticator is current using is no longer considered valid by
	// the VSZ being queried
	//
	// Arguments:
	//
	// 0. the context provided to the original API call
	// 1. the current VSZ client
	// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or Refresh)
	Invalidate(context.Context, *VSZClient, VSZServiceTicketCAS) (VSZServiceTicketCAS, error)
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
func (stp *UsernamePasswordVSZSServiceTicketProvider) Current() (VSZServiceTicketCAS, string, error) {
	stp.mu.RLock()
	defer stp.mu.RUnlock()

	if stp.ticketValid() {
		return stp.cas, stp.serviceTicket, nil
	}

	return stp.cas, "", NewAPIAuthProviderError(APIResponseMeta{}, ErrServiceTicketRequiresRefresh)
}

// Refresh will attempt to fetch a new serviceTicket from the VSZ for use in subsequent requests
func (stp *UsernamePasswordVSZSServiceTicketProvider) Refresh(ctx context.Context, client *VSZClient, cas VSZServiceTicketCAS) (VSZServiceTicketCAS, error) {
	stp.mu.Lock()
	defer stp.mu.Unlock()

	var (
		loginRequest  *WSGServiceTicketLoginRequest
		loginResponse *WSGServiceTicketLoginResponseAPIResponse
		loginMeta     APIResponseMeta
		err           error

		username = stp.username
		password = stp.password
	)

	// if client is nil, fail immediately
	if client == nil {
		return stp.cas, NewAPIAuthProviderError(loginMeta, ErrServiceTicketClientNil)
	}

	// if the passed cas value is greater than the internal CAS, assume weirdness and return current CAS and an error
	if stp.cas < cas {
		return stp.cas, NewAPIAuthProviderError(loginMeta, fmt.Errorf("%w: provided cas value is greater than possible", ErrServiceTicketCASInvalid))
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
	loginResponse, err = client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	for i := uint(1); err != nil && i <= stp.retries; i++ {
		time.Sleep(stp.retryWait)
		loginResponse, err = client.WSG().WSGServiceTicketService().AddServiceTicket(ctx, loginRequest)
	}

	if loginResponse != nil {
		defer CleanupReadCloser(loginResponse)
		loginMeta = loginResponse.ResponseMeta()
	}

	if err != nil {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return stp.cas, NewAPIAuthProviderError(loginMeta, err)
	}

	if loginResponse == nil {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return stp.cas, NewAPIAuthProviderError(loginMeta, ErrServiceTicketResponseEmpty)
	}

	// test if we need to try to hydrate the response model
	if loginResponse.Data == nil {
		// attempt model hydration
		if err = loginResponse.Hydrate(); err != nil {
			stp.cas = stp.iterateCAS()
			stp.serviceTicket = ""
			stp.expires = time.Time{}
			return stp.cas, NewAPIAuthProviderError(loginMeta, fmt.Errorf("%w: error unmarshalling response body: %v", ErrServiceTicketResponseEmpty, err))
		}
	}

	// this case is extremely unlikely, but keep it here just in case.  don't want no npe.
	if loginResponse.Data.ServiceTicket == nil || *loginResponse.Data.ServiceTicket == "" {
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
		return stp.cas, NewAPIAuthProviderError(loginMeta, ErrServiceTicketResponseEmpty)
	}

	// if reached, assume login successful
	stp.cas = stp.iterateCAS()
	stp.serviceTicket = *loginResponse.Data.ServiceTicket
	stp.expires = time.Now().Add(stp.sessionTTL)

	return stp.cas, nil
}

// Invalidate will mark the current session as invalid.
func (stp *UsernamePasswordVSZSServiceTicketProvider) Invalidate(ctx context.Context, client *VSZClient, cas VSZServiceTicketCAS) (VSZServiceTicketCAS, error) {
	stp.mu.Lock()
	defer stp.mu.Unlock()

	var (
		logoutResp *RawAPIResponse
		logoutMeta APIResponseMeta
		err        error
	)

	// if current cas is less than provided, assume insanity.
	if stp.cas < cas {
		return stp.cas, NewAPIAuthProviderError(logoutMeta, fmt.Errorf("%w: provided cas value greater than possible", ErrServiceTicketCASInvalid))
	}

	// if current cas is greater than provided, assume Refresh or Invalidate has already been called.
	if stp.cas > cas {
		return stp.cas, nil
	}

	// if we have a service ticket stored, attempt to invalidate it at the VSZ
	if stp.serviceTicket != "" {
		logoutResp, err = client.WSG().WSGServiceTicketService().DeleteServiceTicket(ctx, stp.serviceTicket)
		if logoutResp != nil {
			CleanupReadCloser(logoutResp)
			logoutMeta = logoutResp.ResponseMeta()
		}
		if err != nil {
			err = NewAPIAuthProviderError(logoutMeta, err)
		}

		// update internal state
		stp.cas = stp.iterateCAS()
		stp.serviceTicket = ""
		stp.expires = time.Time{}
	}

	return stp.cas, err
}

func (stp *UsernamePasswordVSZSServiceTicketProvider) ticketValid() bool {
	return stp.serviceTicket != "" && time.Now().Before(stp.expires)
}

func (*UsernamePasswordVSZSServiceTicketProvider) iterateCAS() VSZServiceTicketCAS {
	return VSZServiceTicketCAS(time.Now().UnixNano())
}

type VSZClientConfig struct {
	baseClientConfig
	// ServiceTicketProvider [required]
	//
	// ServiceTicketProvider to use to handle request auth session
	ServiceTicketProvider VSZServiceTicketProvider
}

type VSZClient struct {
	*baseClient
	stp VSZServiceTicketProvider
}

func NewVSZClient(config *VSZClientConfig) *VSZClient {
	c := new(VSZClient)
	c.baseClient = newBaseClient(config.baseClientConfig)
	c.stp = config.ServiceTicketProvider
	return c
}

// VSZServiceTicketProvider returns the provider used by this client
func (c *VSZClient) ServiceTicketProvider() VSZServiceTicketProvider {
	return c.stp
}

func (c *VSZClient) WSG() *WSGService {
	return NewWSGService(c)
}

func (c *VSZClient) SwitchM() *SwitchMService {
	return NewSwitchMService(c)
}

func (c *VSZClient) Do(ctx context.Context, request *APIRequest, mutators ...RequestMutator) (*http.Response, error) {
	var (
		cas           VSZServiceTicketCAS
		serviceTicket string
		err           error
	)
	if request.Authenticated {
		if cas, serviceTicket, err = c.stp.Current(); err != nil {
			if c.debug {
				c.log.Printf("Error fetching current service ticket: %v", err)
			}
			if !errors.Is(err, ErrServiceTicketRequiresRefresh) {
				return nil, err
			}
			// always call invalidate prior to refresh, just in case...
			if cas, err = c.stp.Invalidate(ctx, c, cas); err != nil {
				if c.debug {
					c.log.Printf("Error invalidating existing service ticket before refresh: %v", err)
				}
			}
			if cas, err = c.stp.Refresh(ctx, c, cas); err != nil {
				if c.debug {
					c.log.Printf("Error refreshing service ticket: %v", err)
				}
				return nil, err
			} else if cas, serviceTicket, err = c.stp.Current(); err != nil {
				if c.debug {
					c.log.Printf("Error fetching current service ticket after refresh: %v", err)
				}
				return nil, err
			}
		}
	}
	return c.do(ctx, request, VSZServiceTicketQueryParameter, serviceTicket, mutators...)
}
