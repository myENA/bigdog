package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGIPSECProfileService struct {
	apiClient *APIClient
}

func NewWSGIPSECProfileService(c *APIClient) *WSGIPSECProfileService {
	s := new(WSGIPSECProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIPSECProfileService() *WSGIPSECProfileService {
	return NewWSGIPSECProfileService(ss.apiClient)
}

// AddProfilesTunnelIpsec
//
// Create a new ipsec.
//
// Request Body:
//	 - body *WSGProfileCreateIpsecProfile
func (s *WSGIPSECProfileService) AddProfilesTunnelIpsec(ctx context.Context, body *WSGProfileCreateIpsecProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTunnelIpsec
//
// Delete multiple ipsec.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsec(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTunnelIpsecById
//
// Delete a ipsec.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsecById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelIpsec
//
// Retrieve a list of IPSEC.
//
// Query Parameters:
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsec(ctx context.Context, qIndex string, qListSize string) (*WSGProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelIpsecById
//
// Retrieve a IPSEC.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecById(ctx context.Context, pId string) (*WSGProfileIpsecProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTunnelIpsecByQueryCriteria
//
// Query a list of IPSEC.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileIpsecProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesTunnelIpsecById
//
// Modify a specific ipsec basic.
//
// Request Body:
//	 - body *WSGProfileModifyIpsecProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIPSECProfileService) PartialUpdateProfilesTunnelIpsecById(ctx context.Context, body *WSGProfileModifyIpsecProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
