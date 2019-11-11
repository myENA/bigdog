package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDNSServerManagementService struct {
	apiClient *APIClient
}

func NewWSGDNSServerManagementService(c *APIClient) *WSGDNSServerManagementService {
	s := new(WSGDNSServerManagementService)
	s.apiClient = c
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
//	 - body *WSGProfileCreateDnsServerProfile
func (s *WSGDNSServerManagementService) AddProfilesDnsserver(ctx context.Context, body *WSGProfileCreateDnsServerProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddProfilesDnsserverCloneById
//
// Use this API command to clone an DNS server profile.
//
// Request Body:
//	 - body *WSGProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) AddProfilesDnsserverCloneById(ctx context.Context, body *WSGProfileCloneRequest, pId string) (*WSGProfileCloneResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesDnsserver
//
// Use this API command to delete a list of DNS server profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserver(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesDnsserverById
//
// Use this API command to delete DNS server profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserverById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
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
func (s *WSGDNSServerManagementService) FindProfilesDnsserver(ctx context.Context, qIndex string, qListSize string) (*WSGProfileDnsServerProfileList, error) {
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
func (s *WSGDNSServerManagementService) FindProfilesDnsserverById(ctx context.Context, pId string) (*WSGProfileDnsServerProfile, error) {
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
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDNSServerManagementService) FindProfilesDnsserverByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileDnsServerProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesDnsserverById
//
// Use this API command to modify the basic information of DNS server profile.
//
// Request Body:
//	 - body *WSGProfileModifyDnsServerProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) PartialUpdateProfilesDnsserverById(ctx context.Context, body *WSGProfileModifyDnsServerProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
