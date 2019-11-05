package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
	"gopkg.in/go-playground/validator.v9"
)

type WSGBonjourGatewayPoliciesService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGBonjourGatewayPoliciesService(c *APIClient) *WSGBonjourGatewayPoliciesService {
	s := new(WSGBonjourGatewayPoliciesService)
	s.apiClient = c
	s.validate = validator.New()
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
//	 - body *zone.CreateBonjourGatewayPolicy
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, body *zone.CreateBonjourGatewayPolicy, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
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
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesById(ctx context.Context, pId string, pZoneId string) (*zone.BonjourGatewayPolicyConfiguration, error) {
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
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, pZoneId string) (*zone.BonjourGatewayPolicyList, error) {
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
//	 - body *zone.ModifyBonjourGatewayPolicy
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) PartialUpdateRkszonesBonjourGatewayPoliciesById(ctx context.Context, body *zone.ModifyBonjourGatewayPolicy, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
