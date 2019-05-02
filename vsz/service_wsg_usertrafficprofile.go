package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGUserTrafficProfileService struct {
    c *Client
}

func NewWSGUserTrafficProfileService (c *Client) *WSGUserTrafficProfileService {
    s := new(WSGUserTrafficProfileService)
    s.c = c
    return s
}

