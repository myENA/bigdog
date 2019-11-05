package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDNSServerManagementService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGDNSServerManagementService(c *APIClient) *WSGDNSServerManagementService {
	s := new(WSGDNSServerManagementService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGDNSServerManagementService() *WSGDNSServerManagementService {
	return NewWSGDNSServerManagementService(ss.apiClient)
}

// AddProfilesDnsserver
//
// Use this API command to create DNS server profile.
//
// Request Body:
//	 - body *profile.CreateDnsServerProfile
func (s *WSGDNSServerManagementService) AddProfilesDnsserver(ctx context.Context, body *profile.CreateDnsServerProfile) (*common.CreateResult, error) {
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

// AddProfilesDnsserverCloneById
//
// Use this API command to clone an DNS server profile.
//
// Request Body:
//	 - body *profile.ProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) AddProfilesDnsserverCloneById(ctx context.Context, body *profile.ProfileCloneRequest, pId string) (*profile.ProfileCloneResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteProfilesDnsserver
//
// Use this API command to delete a list of DNS server profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserver(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteProfilesDnsserverById
//
// Use this API command to delete DNS server profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserverById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesDnsserver
//
// Use this API command to retrieve a list of DNS server profile.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGDNSServerManagementService) FindProfilesDnsserver(ctx context.Context, qIndex string, qListSize string) (*profile.DnsServerProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesDnsserverById
//
// Use this API command to retrieve DNS server profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) FindProfilesDnsserverById(ctx context.Context, pId string) (*profile.DnsServerProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesDnsserverByQueryCriteria
//
// Use this API command to retrieve a list of DNS server profile  by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGDNSServerManagementService) FindProfilesDnsserverByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.DnsServerProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// PartialUpdateProfilesDnsserverById
//
// Use this API command to modify the basic information of DNS server profile.
//
// Request Body:
//	 - body *profile.ModifyDnsServerProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) PartialUpdateProfilesDnsserverById(ctx context.Context, body *profile.ModifyDnsServerProfile, pId string) (*common.EmptyResult, error) {
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
