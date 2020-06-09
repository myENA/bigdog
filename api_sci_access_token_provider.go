package bigdog

import (
	"context"
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
	username string
	password string

	mu sync.Mutex

	cas         SCIAccessTokenCAS
	expires     time.Time
	sessionTTL  time.Duration
	accessToken string
}
