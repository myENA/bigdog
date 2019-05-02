package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAuthenticationServiceService struct {
    c *Client
}

func NewWSGAuthenticationServiceService (c *Client) *WSGAuthenticationServiceService {
    s := new(WSGAuthenticationServiceService)
    s.c = c
    return s
}

