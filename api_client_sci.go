package bigdog

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	ErrAccessTokenRequiresRefresh = errors.New("access token requires refresh")
	ErrAccessTokenCASInvalid      = errors.New("invalid CAS provided")
	ErrAccessTokenClientNil       = errors.New("client cannot be nil")
	ErrAccessTokenResponseEmpty   = errors.New("empty response from login request")

	accessTokenErrors = []error{
		ErrAccessTokenRequiresRefresh,
		ErrAccessTokenCASInvalid,
		ErrAccessTokenClientNil,
		ErrAccessTokenResponseEmpty,
	}

	accessTokenFatalErrors = []error{
		ErrAccessTokenCASInvalid,
		ErrAccessTokenClientNil,
		ErrAccessTokenResponseEmpty,
	}
)

type SCIAPIErrorDetails struct {
	Name       string `json:"name"`
	Status     int    `json:"status"`
	Message    string `json:"message"`
	Length     int    `json:"length"`
	Severity   string `json:"severity"`
	Code       string `json:"code"`
	StatusCode int    `json:"statusCode"`
	File       string `json:"file"`
	Line       int    `json:"line"`
	Routine    string `json:"routine"`
}

type SCIAPIError struct {
	ErrorDetails *SCIAPIErrorDetails `json:"error,omitempty"`

	Err error `json:"err,omitempty"`
}

func (e *SCIAPIError) Is(err error) bool {
	return e != nil && e.Err != nil && errors.Is(err, e.Err)
}

func (e *SCIAPIError) Unwrap() error {
	if e == nil || e.Err == nil {
		return nil
	}
	return e.Err
}

func (e *SCIAPIError) Error() string {
	if e == nil {
		return ""
	}
	if e.ErrorDetails != nil {
		if e.ErrorDetails.Code == "" {
			return fmt.Sprintf(
				"name=%q; status=%q; statusCode=%q; message=%q",
				e.ErrorDetails.Name,
				e.ErrorDetails.Status,
				e.ErrorDetails.StatusCode,
				e.ErrorDetails.Message,
			)
		}
		return fmt.Sprintf(
			"name=%q; status=%q; length=%q; severity=%q; code=%q; statusCode=%q; file=%q; line=%q; routine=%q; message=%q",
			e.ErrorDetails.Name,
			e.ErrorDetails.Status,
			e.ErrorDetails.Length,
			e.ErrorDetails.Severity,
			e.ErrorDetails.Code,
			e.ErrorDetails.StatusCode,
			e.ErrorDetails.File,
			e.ErrorDetails.Line,
			e.ErrorDetails.Routine,
			e.ErrorDetails.Message,
		)
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func IsSCIAPIError(err error) bool {
	for err != nil {
		if serr, ok := err.(*SCIAPIError); ok && serr != nil {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

func IsAccessTokenProviderError(err error) bool {
	for _, aerr := range accessTokenErrors {
		if errors.Is(err, aerr) {
			return true
		}
	}
	return false
}

func IsFatalAccessTokenProviderError(err error) bool {
	for _, aerr := range accessTokenFatalErrors {
		if errors.Is(err, aerr) {
			return true
		}
	}
	return false
}

type SCIAccessTokenCAS time.Duration

type SCIAccessTokenProvider interface {
	// Current must return the current cas and token value, if and only if it is still valid.
	Current() (SCIAccessTokenCAS, string, error)

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
	// 2. If initial Current did not fail but SCI returned a 401, causing an Invalidate -> Refresh -> Current chain
	// that will execute exactly 1 times.
	//
	// Arguments:
	//
	// 0. the context provided to the original API call
	// 1. the current SCI client
	// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or Invalidate)
	Refresh(context.Context, *SCIClient, SCIAccessTokenCAS) (SCIAccessTokenCAS, error)

	// Invalidate will only be called if a 401 is seen after a refresh has been attempted, and should indicate to
	// the implementor that whatever decoration the authenticator is current using is no longer considered valid by
	// the SCI being queried
	//
	// Arguments:
	//
	// 0. the context provided to the original API call
	// 1. the current SCI client
	// 2. the CAS value seen from the calling routine's most recent action (could be from either Current or Refresh)
	Invalidate(context.Context, *SCIClient, SCIAccessTokenCAS) (SCIAccessTokenCAS, error)
}

type UsernamePasswordSCIAccessTokenProvider struct {
	mu sync.RWMutex

	username   string
	password   string
	retries    uint
	retryWait  time.Duration
	sessionTTL time.Duration

	cas         SCIAccessTokenCAS
	expires     time.Time
	accessToken string
}

func NewUsernamePasswordSCIAccessTokenProvider(username, password string, retries uint, retryWait, sessionTTL time.Duration) *UsernamePasswordSCIAccessTokenProvider {
	atp := new(UsernamePasswordSCIAccessTokenProvider)
	atp.username = username
	atp.password = password
	atp.retries = retries
	atp.retryWait = retryWait
	atp.sessionTTL = sessionTTL
	return atp
}

func (atp *UsernamePasswordSCIAccessTokenProvider) Username() string {
	return atp.username
}

func (atp *UsernamePasswordSCIAccessTokenProvider) Password() string {
	return atp.password
}

func (atp *UsernamePasswordSCIAccessTokenProvider) SessionTTL() time.Duration {
	return atp.sessionTTL
}

// Current returns the current access token, if there is one.
func (atp *UsernamePasswordSCIAccessTokenProvider) Current() (SCIAccessTokenCAS, string, error) {
	atp.mu.RLock()
	defer atp.mu.RUnlock()

	if atp.tokenValid() {
		return atp.cas, atp.accessToken, nil
	}

	return atp.cas, "", NewAPIAuthProviderError(APIResponseMeta{}, ErrAccessTokenRequiresRefresh)
}

func (atp *UsernamePasswordSCIAccessTokenProvider) Refresh(ctx context.Context, client *SCIClient, cas SCIAccessTokenCAS) (SCIAccessTokenCAS, error) {
	atp.mu.Lock()
	defer atp.mu.Unlock()

	var (
		loginRequest  *SCIUserLoginRequest
		loginResponse *SCIUserLoginResponseAPIResponse
		loginMeta     APIResponseMeta
		err           error

		username = atp.username
		password = atp.password
	)

	// if client is nil, fail immediately
	if client == nil {
		return atp.cas, NewAPIAuthProviderError(loginMeta, ErrAccessTokenClientNil)
	}

	// if incoming cas is greater than current cas, assume weirdness
	if atp.cas < cas {
		return atp.cas, NewAPIAuthProviderError(loginMeta, fmt.Errorf("%w: provided cas value is greater than possible", ErrAccessTokenCASInvalid))
	}

	// if stored cas is greater than incoming cas, assume delayed call and return OK
	if atp.cas > cas {
		return atp.cas, nil
	}

	loginRequest = &SCIUserLoginRequest{Username: &username, Password: &password}
	loginResponse, err = client.SCI().SCIUserService().UserLogin(ctx, loginRequest, nil)
	for i := uint(1); err != nil && i < atp.retries; i++ {
		time.Sleep(atp.retryWait)
		loginResponse, err = client.SCI().SCIUserService().UserLogin(ctx, loginRequest, nil)
	}

	if loginResponse != nil {
		defer cleanupReadCloser(loginResponse)
		loginMeta = loginResponse.ResponseMeta()
	}

	if err != nil {
		atp.cas = atp.iterateCAS()
		atp.accessToken = ""
		atp.expires = time.Time{}
		return atp.cas, NewAPIAuthProviderError(loginMeta, err)
	}

	if loginResponse == nil {
		atp.cas = atp.iterateCAS()
		atp.accessToken = ""
		atp.expires = time.Time{}
		return atp.cas, NewAPIAuthProviderError(loginMeta, ErrAccessTokenResponseEmpty)
	}

	// test if we need to hydrate response
	if loginResponse.Data == nil {
		if _, err = loginResponse.Hydrate(); err != nil {
			atp.cas = atp.iterateCAS()
			atp.accessToken = ""
			atp.expires = time.Time{}
			return atp.cas, NewAPIAuthProviderError(loginMeta, fmt.Errorf("%w: error unmarshalling response body: %v", ErrAccessTokenResponseEmpty, err))
		}
	}

	if loginResponse.Data.Id == nil || *loginResponse.Data.Id == "" {
		atp.cas = atp.iterateCAS()
		atp.accessToken = ""
		atp.expires = time.Time{}
		return atp.cas, NewAPIAuthProviderError(loginMeta, ErrAccessTokenResponseEmpty)
	}

	atp.cas = atp.iterateCAS()
	atp.accessToken = *loginResponse.Data.Id
	atp.expires = time.Now().Add(atp.sessionTTL)

	return atp.cas, nil
}

func (atp *UsernamePasswordSCIAccessTokenProvider) Invalidate(ctx context.Context, client *SCIClient, cas SCIAccessTokenCAS) (SCIAccessTokenCAS, error) {
	atp.mu.Lock()
	defer atp.mu.Unlock()

	var (
		logoutResp APIResponse
		logoutMeta APIResponseMeta
		err        error
	)

	if atp.cas < cas {
		return atp.cas, NewAPIAuthProviderError(logoutMeta, fmt.Errorf("%w: provided cas value is greater than possible", ErrAccessTokenCASInvalid))
	}

	if atp.cas > cas {
		return atp.cas, nil
	}

	if atp.accessToken != "" {
		logoutResp, err = client.SCI().SCIUserService().UserLogout(ctx, atp.accessToken)
		if logoutResp != nil {
			cleanupReadCloser(logoutResp)
			logoutMeta = logoutResp.ResponseMeta()
		}
		if err != nil {
			err = NewAPIAuthProviderError(logoutMeta, err)
		}
		atp.cas = atp.iterateCAS()
		atp.accessToken = ""
		atp.expires = time.Time{}
	}

	return atp.cas, err
}

func (atp *UsernamePasswordSCIAccessTokenProvider) tokenValid() bool {
	return atp.accessToken != "" && time.Now().Before(atp.expires)
}

func (*UsernamePasswordSCIAccessTokenProvider) iterateCAS() SCIAccessTokenCAS {
	return SCIAccessTokenCAS(time.Now().UnixNano())
}

type SCIClientConfig struct {
	// Address [required]
	//
	// Full address of service, including scheme and port
	Address string

	// PathPrefix [optional]
	//
	// Custom path prefix to prepend to all api calls.  Default is to leave this blank.
	PathPrefix string

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool

	// DisableAutoHydrate [bool]
	//
	// If true, response bodies will no longer automatically hydrated into the returned response models.  This enables
	// you to instead use the response models as Readers to receive the raw bytes of the response in favor of having
	// then unmarshalled if you don't need it.
	DisableAutoHydrate bool

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger

	// HTTPClient [optional]
	HTTPClient *http.Client

	// ServiceTicketProvider [required]
	//
	// ServiceTicketProvider to use to handle request auth session
	AccessTokenProvider SCIAccessTokenProvider
}

type SCIClient struct {
	*baseClient
	atp SCIAccessTokenProvider
}

func NewSCIClient(config *SCIClientConfig) *SCIClient {
	c := new(SCIClient)
	c.baseClient = newBaseClient(config.Address, config.PathPrefix, config.Debug, config.DisableAutoHydrate, config.Logger, config.HTTPClient)
	c.atp = config.AccessTokenProvider
	return c
}

func (c *SCIClient) AccessTokenProvider() SCIAccessTokenProvider {
	return c.atp
}

func (c *SCIClient) SCI() *SCIService {
	return NewSCIService(c)
}

func (c *SCIClient) Do(ctx context.Context, request *APIRequest, mutators ...RequestMutator) (*http.Response, error) {
	var (
		cas         SCIAccessTokenCAS
		accessToken string
		err         error
	)
	if request.Authenticated {
		if cas, accessToken, err = c.atp.Current(); err != nil {
			if c.debug {
				c.log.Printf("Error fetching current access token: %v", err)
			}
			if !errors.Is(err, ErrAccessTokenRequiresRefresh) {
				return nil, err
			}
			// always invalidate, just in case...
			if cas, err = c.atp.Invalidate(ctx, c, cas); err != nil {
				if c.debug {
					c.log.Printf("Error invalidating existing access token: %v", err)
				}
			}
			if cas, err = c.atp.Refresh(ctx, c, cas); err != nil {
				if c.debug {
					c.log.Printf("Error refreshing access token: %v", err)
				}
				return nil, err
			} else if cas, accessToken, err = c.atp.Current(); err != nil {
				if c.debug {
					c.log.Printf("Error fetching current access token after refresh: %v", err)
				}
				return nil, err
			}
		}
	}
	return c.do(ctx, request, SCIAccessTokenQueryParameter, accessToken, mutators...)
}
