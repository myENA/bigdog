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

type WSGAccountingProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGAccountingProfileService(c *APIClient) *WSGAccountingProfileService {
	s := new(WSGAccountingProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGAccountingProfileService() *WSGAccountingProfileService {
	return NewWSGAccountingProfileService(ss.apiClient)
}

// AddProfilesAcct
//
// Use this API command to create a new accounting profile.
//
// Request Body:
//	 - body *profile.CreateAccountingProfile
func (s *WSGAccountingProfileService) AddProfilesAcct(ctx context.Context, body *profile.CreateAccountingProfile) (*common.CreateResult, error) {
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

// AddProfilesAcctCloneById
//
// Use this API command to clone an accounting profile.
//
// Request Body:
//	 - body *profile.ProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) AddProfilesAcctCloneById(ctx context.Context, body *profile.ProfileCloneRequest, pId string) (*profile.ProfileCloneResponse, error) {
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

// DeleteProfilesAcct
//
// Use this API command to delete a list of accounting profile.
//
// Request Body:
//	 - body *profile.DeleteBulkAccountingProfile
func (s *WSGAccountingProfileService) DeleteProfilesAcct(ctx context.Context, body *profile.DeleteBulkAccountingProfile) (*common.EmptyResult, error) {
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

// DeleteProfilesAcctById
//
// Use this API command to delete an accounting profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) DeleteProfilesAcctById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesAcct
//
// Use this API command to retrieve a list of accounting profiles.
func (s *WSGAccountingProfileService) FindProfilesAcct(ctx context.Context) (*profile.AccountingProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesAcctById
//
// Use this API command to retrieve an accounting profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) FindProfilesAcctById(ctx context.Context, pId string) (*profile.AccountingProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting profiles by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingProfileService) FindProfilesAcctByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.AccountingProfileList, error) {
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

// PartialUpdateProfilesAcctById
//
// Use this API command to modify the basic information of an accounting profile.
//
// Request Body:
//	 - body *profile.ModifyAccountingProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) PartialUpdateProfilesAcctById(ctx context.Context, body *profile.ModifyAccountingProfile, pId string) (*common.EmptyResult, error) {
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
