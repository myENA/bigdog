package vsz

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	DefaultAPIPort           = 8443
	DefaultWSGPathPrefix     = "wsg/api"
	DefaultSwitchMPathPrefix = "switchm/api"
)

type APIConfig struct {
	// Address [required]
	//
	// Address or hostname of your VSZ instance
	Address string `json:"address"`

	// WSGPathPrefix [optional]
	//
	// Path prefix to access Wireless Security Gateway API endpoints
	WSGPathPrefix string `json:"wsgPathPrefix"`

	// SwitchMPathPrefix [optional]
	//
	// Path prefix to access Switch Management API endpoints
	SwitchMPathPrefix string `json:"switchMPathPrefix"`

	// APIPort [optional]
	//
	// APIPort to use when executing API calls.
	APIPort int `json:"port"`

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool `json:"debug"`

	// Authenticator [required]
	//
	// Authenticator to use to handle request auth session
	Authenticator Authenticator `json:"-"`

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger `json:"-"`

	// HTTPClient [optional]
	HTTPClient *http.Client `json:"-"`
}

type APIClient struct {
	log   *log.Logger
	debug bool

	addr        string
	wsgPath     string
	switchMPath string
	port        int
	auth        Authenticator

	client *http.Client
}

func NewAPIClient(conf *APIConfig) (*APIClient, error) {
	if conf.Address == "" {
		return nil, errors.New("address must be defined")
	}
	if conf.Authenticator == nil {
		return nil, errors.New("authenticator must be defined")
	}

	c := new(APIClient)
	c.addr = conf.Address
	c.auth = conf.Authenticator
	c.debug = conf.Debug

	if conf.Logger != nil {
		c.log = conf.Logger
	} else {
		c.log = log.New(ioutil.Discard, "", 0)
	}

	if conf.HTTPClient != nil {
		c.client = conf.HTTPClient
	} else {
		// shamelessly borrowed from https://github.com/hashicorp/go-cleanhttp/blob/master/cleanhttp.go
		c.client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     true,
				MaxIdleConnsPerHost:   -1,
			},
		}
	}

	if conf.APIPort != 0 {
		c.port = conf.APIPort
	} else {
		c.port = DefaultAPIPort
	}
	if conf.WSGPathPrefix != "" {
		c.wsgPath = conf.WSGPathPrefix
	} else {
		c.wsgPath = DefaultWSGPathPrefix
	}
	if conf.SwitchMPathPrefix != "" {
		c.switchMPath = conf.SwitchMPathPrefix
	} else {
		c.switchMPath = DefaultSwitchMPathPrefix
	}

	return c, nil
}

func (c *APIClient) Debug() bool {
	return c.debug
}

func (c *APIClient) WSG() *WSGService {
	return NewWSGService(c)
}

func (c *APIClient) SwitchM() *SwitchMService {
	return NewSwitchMService(c)
}

func (c *APIClient) Do(ctx context.Context, request *Request) (*http.Response, error) {
	_, httpResponse, err := c.do(ctx, request)
	return httpResponse, err
}

func (c *APIClient) tryDo(ctx context.Context, request *Request) (AuthCAS, *http.Response, error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		cas          AuthCAS
		err          error
	)
	if httpRequest, err = request.toHTTP(ctx, c.addr, c.port, c.debug); err != nil {
		return cas, nil, err
	}
	if request.authenticated {
		if cas, err = c.auth.Decorate(ctx, httpRequest); err != nil {
			if cas, err = c.auth.Refresh(ctx, c, cas); err != nil {
				return cas, nil, err
			} else if cas, err = c.auth.Decorate(ctx, httpRequest); err != nil {
				return cas, nil, err
			}
		}
	}
	httpResponse, err = c.client.Do(httpRequest)
	return cas, httpResponse, err
}

func (c *APIClient) do(ctx context.Context, request *Request) (AuthCAS, *http.Response, error) {
	if ctx == nil {
		return 0, nil, errors.New("ctx must not be nil")
	}
	return c.tryDo(ctx, request)
}
