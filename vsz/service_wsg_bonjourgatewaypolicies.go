package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGBonjourGatewayPoliciesService struct {
    c *Client
}

func NewWSGBonjourGatewayPoliciesService (c *Client) *WSGBonjourGatewayPoliciesService {
    s := new(WSGBonjourGatewayPoliciesService)
    s.c = c
    return s
}

