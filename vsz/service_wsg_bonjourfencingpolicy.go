package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBonjourFencingPolicyService struct {
    c *Client
}

func NewWSGBonjourFencingPolicyService (c *Client) *WSGBonjourFencingPolicyService {
    s := new(WSGBonjourFencingPolicyService)
    s.c = c
    return s
}

