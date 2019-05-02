package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGTrafficClassProfileService struct {
    c *Client
}

func NewWSGTrafficClassProfileService (c *Client) *WSGTrafficClassProfileService {
    s := new(WSGTrafficClassProfileService)
    s.c = c
    return s
}

