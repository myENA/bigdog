package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGBonjourGatewayPoliciesService struct {
	apiClient *APIClient
}

func NewWSGBonjourGatewayPoliciesService(c *APIClient) *WSGBonjourGatewayPoliciesService {
	s := new(WSGBonjourGatewayPoliciesService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBonjourGatewayPoliciesService() *WSGBonjourGatewayPoliciesService {
	serv := new(WSGBonjourGatewayPoliciesService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to create bonjour gateway policy.
//
// Request Body:
//	 - body *zone.CreateBonjourGatewayPolicy
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, body *zone.CreateBonjourGatewayPolicy, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesBonjourGatewayPoliciesById
//
// Use this API command to delete bonjour gateway policy.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) DeleteRkszonesBonjourGatewayPoliciesById(ctx context.Context, pId string, pZoneId string) error {
}

// FindRkszonesBonjourGatewayPoliciesById
//
// Use this API command to retrieve bonjour gateway policy.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesById(ctx context.Context, pId string, pZoneId string) (*zone.BonjourGatewayPolicyConfiguration, error) {
}

// FindRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to retrieve a list of bonjour gateway policies.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, pZoneId string) (*zone.BonjourGatewayPolicyList, error) {
}

// PartialUpdateRkszonesBonjourGatewayPoliciesById
//
// Use this API command to modify the basic information of bonjour gateway policy.
//
// Request Body:
//	 - body *zone.ModifyBonjourGatewayPolicy
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) PartialUpdateRkszonesBonjourGatewayPoliciesById(ctx context.Context, body *zone.ModifyBonjourGatewayPolicy, pId string, pZoneId string) (*common.EmptyResult, error) {
}
