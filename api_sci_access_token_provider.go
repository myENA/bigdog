package bigdog

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	SCIAccessTokenCAS time.Duration

	SCIAccessTokenProvider interface {
		Current() (SCIAccessTokenCAS, string, error)
		Refresh(context.Context, *SCIClient, SCIAccessTokenCAS) (SCIAccessTokenCAS, error)
		Invalidate(context.Context, *SCIClient, SCIAccessTokenCAS) (SCIAccessTokenCAS, error)
	}
)

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

func (atp *UsernamePasswordSCIAccessTokenProvider) Current() (SCIAccessTokenCAS, string, error) {
	atp.mu.RLock()
	defer atp.mu.RUnlock()

	if atp.accessToken != "" && time.Now().Before(atp.expires) {
		return atp.cas, atp.accessToken, nil
	}

	return atp.cas, "", NewSCIAccessTokenProviderError(newErrAPIResponseMeta(), ErrAccessTokenRequiresRefresh)
}

func (atp *UsernamePasswordSCIAccessTokenProvider) Refresh(ctx context.Context, client *SCIClient, cas SCIAccessTokenCAS) (SCIAccessTokenCAS, error) {
	atp.mu.Lock()
	defer atp.mu.Unlock()

	var (
		loginRequest  *SCIUserLoginRequest
		loginResponse *SCIUserLoginResponse
		rm            *APIResponseMeta
		err           error

		username = atp.username
		password = atp.password
	)

	if client == nil {
		return atp.cas, NewSCIAccessTokenProviderError(newErrAPIResponseMeta(), ErrAccessTokenClientNil)
	}

	if atp.cas < cas {
		return atp.cas, NewSCIAccessTokenProviderError(newErrAPIResponseMeta(), fmt.Errorf("%w: provided cas value is greater than possible", ErrAccessTokenCASInvalid))
	}

	if atp.cas > cas {
		return atp.cas, nil
	}

	loginRequest = &SCIUserLoginRequest{Username: &username, Password: &password}
	loginResponse, rm, err = client.SCI().SCIUserService().UserLogin(ctx, loginRequest, nil)
	for i := uint(1); err != nil && i < atp.retries; i++ {
		time.Sleep(atp.retryWait)
		loginResponse, rm, err = client.SCI().SCIUserService().UserLogin(ctx, loginRequest, nil)
	}

	if err != nil {
		atp.cas = atp.iterateCAS()
		atp.accessToken = ""
		atp.expires = time.Time{}
		return atp.cas, NewVSZServiceTicketProviderError(rm, err)
	}

	if loginResponse == nil || loginResponse.Id == nil || *loginResponse.Id == "" {
		atp.cas = atp.iterateCAS()
		atp.accessToken = ""
		atp.expires = time.Time{}
		return atp.cas, NewSCIAccessTokenProviderError(rm, ErrAccessTokenResponseEmpty)
	}

	atp.cas = atp.iterateCAS()
	atp.accessToken = *loginResponse.Id
	atp.expires = time.Now().Add(atp.sessionTTL)

	return atp.cas, nil
}

func (atp *UsernamePasswordSCIAccessTokenProvider) Invalidate(ctx context.Context, client *SCIClient, cas SCIAccessTokenCAS) (SCIAccessTokenCAS, error) {
	atp.mu.Lock()
	defer atp.mu.Unlock()

	var (
		rm  *APIResponseMeta
		err error
	)

	if atp.cas < cas {
		return atp.cas, NewSCIAccessTokenProviderError(newErrAPIResponseMeta(), fmt.Errorf("%w: provided cas value is greater than possible", ErrAccessTokenCASInvalid))
	}

	if atp.cas > cas {
		return atp.cas, nil
	}

	if atp.accessToken != "" {
		if rm, err = client.SCI().SCIUserService().UserLogout(ctx, atp.accessToken); err != nil {
			err = NewSCIAccessTokenProviderError(rm, err)
		}

		atp.cas = atp.iterateCAS()
		atp.accessToken = ""
		atp.expires = time.Time{}
	}

	return atp.cas, err
}

func (*UsernamePasswordSCIAccessTokenProvider) iterateCAS() SCIAccessTokenCAS {
	return SCIAccessTokenCAS(time.Now().UnixNano())
}
