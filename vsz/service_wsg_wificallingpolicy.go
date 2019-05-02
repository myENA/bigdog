package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wificalling"
)

type WSGWiFiCallingPolicyService struct {
    c *Client
}

func NewWSGWiFiCallingPolicyService (c *Client) *WSGWiFiCallingPolicyService {
    s := new(WSGWiFiCallingPolicyService)
    s.c = c
    return s
}

