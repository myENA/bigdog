package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/urlfiltering"
)

type WSGURLFilteringPolicyService struct {
    c *Client
}

func NewWSGURLFilteringPolicyService (c *Client) *WSGURLFilteringPolicyService {
    s := new(WSGURLFilteringPolicyService)
    s.c = c
    return s
}

