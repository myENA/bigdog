package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGGuestAccessService struct {
    c *Client
}

func NewWSGGuestAccessService (c *Client) *WSGGuestAccessService {
    s := new(WSGGuestAccessService)
    s.c = c
    return s
}

