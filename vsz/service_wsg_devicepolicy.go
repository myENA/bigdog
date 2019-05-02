package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/devicepolicy"
)

type WSGDevicePolicyService struct {
    c *Client
}

func NewWSGDevicePolicyService (c *Client) *WSGDevicePolicyService {
    s := new(WSGDevicePolicyService)
    s.c = c
    return s
}

