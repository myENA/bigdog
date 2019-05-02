package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGWebAuthenticationService struct {
    c *Client
}

func NewWSGWebAuthenticationService (c *Client) *WSGWebAuthenticationService {
    s := new(WSGWebAuthenticationService)
    s.c = c
    return s
}

