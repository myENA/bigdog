package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGBonjourGatewayPoliciesService struct {
    client *Client
}

func NewWSGBonjourGatewayPoliciesService (client *Client) *WSGBonjourGatewayPoliciesService {
    s := new(WSGBonjourGatewayPoliciesService)
    s.client = client
    return s
}

func (ss *WSGService) WSGBonjourGatewayPoliciesService () *WSGBonjourGatewayPoliciesService {
    serv := new(WSGBonjourGatewayPoliciesService)
    serv.client = ss.client
    return serv
}

