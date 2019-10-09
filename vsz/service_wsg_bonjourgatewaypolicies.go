package vsz

// API Version: v8_1

import (
	"context"
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

func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesById (ctx context.Context, id string, zoneId string) (*zone.BonjourGatewayPolicyConfiguration, error) {
}

func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesByZoneId (ctx context.Context, zoneId string) (*zone.BonjourGatewayPolicyList, error) {
}

