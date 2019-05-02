package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/accountsecurityprofile"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAccountSecurityService struct {
    c *Client
}

func NewWSGAccountSecurityService (c *Client) *WSGAccountSecurityService {
    s := new(WSGAccountSecurityService)
    s.c = c
    return s
}

