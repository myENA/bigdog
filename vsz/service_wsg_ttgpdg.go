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

type WSGTTGPDGService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGTTGPDGService(c *APIClient) *WSGTTGPDGService {
	s := new(WSGTTGPDGService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGTTGPDGService() *WSGTTGPDGService {
	return NewWSGTTGPDGService(ss.apiClient)
}

// AddProfilesTtgpdg
//
// Use this API command to create TTG+PDG profile.
//
// Request Body:
//	 - body *profile.CreateTtgpdgProfile
func (s *WSGTTGPDGService) AddProfilesTtgpdg(ctx context.Context, body *profile.CreateTtgpdgProfile) (*common.CreateResult, error) {
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

// DeleteProfilesTtgpdg
//
// Use this API command to delete multiple TTG PDG profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGTTGPDGService) DeleteProfilesTtgpdg(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
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

// DeleteProfilesTtgpdgApnRealmsById
//
// Use this API command to disable the APN realm of TTG PDG profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgApnRealmsById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTtgpdgById
//
// Use this API command to delete TTG PDG profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesTtgpdgDhcpRelayById
//
// Use this API command to disable the DHCP relay of TTG PDG profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgDhcpRelayById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTtgpdg
//
// Use this API command to retrieve a list of TTG+PDG profile.
func (s *WSGTTGPDGService) FindProfilesTtgpdg(ctx context.Context) (*profile.ProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTtgpdgById
//
// Use this API command to retrieve TTG+PDG profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) FindProfilesTtgpdgById(ctx context.Context, pId string) (*profile.TtgpdgProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesTtgpdgByQueryCriteria
//
// Use this API command to query a list of TTG+PDG profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGTTGPDGService) FindProfilesTtgpdgByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.TtgpdgProfileList, error) {
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

// PartialUpdateProfilesTtgpdgById
//
// Use this API command to modify the basic information of TTG+PDG profile.
//
// Request Body:
//	 - body *profile.TtgpdgProfileConfiguration
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) PartialUpdateProfilesTtgpdgById(ctx context.Context, body *profile.TtgpdgProfileConfiguration, pId string) (*common.EmptyResult, error) {
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
