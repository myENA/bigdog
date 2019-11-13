package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	return NewWSGBonjourGatewayPoliciesService(ss.apiClient)
}

// AddRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to create bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneCreateBonjourGatewayPolicy
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, body *WSGZoneCreateBonjourGatewayPolicy, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesById(ctx context.Context, pId string, pZoneId string) (*WSGZoneBonjourGatewayPolicyConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to retrieve a list of bonjour gateway policies.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, pZoneId string) (*WSGZoneBonjourGatewayPolicyList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesBonjourGatewayPoliciesById
//
// Use this API command to modify the basic information of bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneModifyBonjourGatewayPolicy
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) PartialUpdateRkszonesBonjourGatewayPoliciesById(ctx context.Context, body *WSGZoneModifyBonjourGatewayPolicy, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
