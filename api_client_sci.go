package bigdog

import (
	"context"
	"log"
	"net/http"
)

type SCIClientConfig struct {
	// Address [required]
	//
	// Full address of SCI, including scheme and port
	Address string

	// PathPrefix [optional]
	//
	// Custom path prefix to prepend to all api calls.  Default is to leave this blank.
	PathPrefix string

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool

	// ServiceTicketProvider [required]
	//
	// ServiceTicketProvider to use to handle request auth session
	AccessTokenProvider SCIAccessTokenProvider

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger

	// HTTPClient [optional]
	HTTPClient *http.Client
}

type SCIClient struct {
	*baseClient
	atp SCIAccessTokenProvider
}

func NewSCIClient(config *SCIClientConfig) *SCIClient {
	c := new(SCIClient)
	c.baseClient = newBaseClient(config.Logger, config.Debug, config.Address, config.PathPrefix, config.HTTPClient)
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
