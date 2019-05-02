package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityGuestPassService struct {
    c *Client
}

func NewWSGIdentityGuestPassService (c *Client) *WSGIdentityGuestPassService {
    s := new(WSGIdentityGuestPassService)
    s.c = c
    return s
}

