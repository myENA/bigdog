package bigdog

import (
	"context"
	"log"
	"net/http"
)

type VSZClientConfig struct {
	// Address [required]
	//
	// Full address of VSZ, including scheme and port
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
	ServiceTicketProvider VSZServiceTicketProvider

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
}

type VSZClient struct {
	*baseClient
	stp VSZServiceTicketProvider
}

func NewVSZClient(config *VSZClientConfig) *VSZClient {
	c := new(VSZClient)
	c.baseClient = newBaseClient(config.Logger, config.Debug, config.Address, config.PathPrefix, config.HTTPClient)
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
