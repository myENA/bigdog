package vsz

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type Config struct {
	// Address [required]
	//
	// Address or hostname of your VSZ instance
	Address string `json:"hostname"`

	// Authenticator [required]
	//
	// Authenticator to use to handle request auth session
	Authenticator Authenticator `json:"-"`

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger `json:"-"`

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool `json:"debug"`

	// HTTPClient [optional]
	HTTPClient *http.Client `json:"-"`
}

type Client struct {
	log   *log.Logger
	debug bool

	addr string
	auth Authenticator

	client *http.Client
}

func NewClient(conf *Config) (*Client, error) {
	if conf.Address == "" {
		return nil, errors.New("address must be defined")
	}
	if conf.Authenticator == nil {
		return nil, errors.New("authenticator must be defined")
	}

	c := new(Client)
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

	return c, nil
}

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) WSG() *WSGService {
	return NewWSGService(c)
}

func (c *Client) SwitchM() *SwitchMService {
	return NewSwitchMService(c)
}
