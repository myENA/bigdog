package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAccountingProfileService struct {
	apiClient *APIClient
}

func NewWSGAccountingProfileService(c *APIClient) *WSGAccountingProfileService {
	s := new(WSGAccountingProfileService)
	s.apiClient = c
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
//	 - body *WSGProfileCreateAccountingProfile
func (s *WSGAccountingProfileService) AddProfilesAcct(ctx context.Context, body *WSGProfileCreateAccountingProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddProfilesAcctCloneById
//
// Use this API command to clone an accounting profile.
//
// Request Body:
//	 - body *WSGProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) AddProfilesAcctCloneById(ctx context.Context, body *WSGProfileCloneRequest, pId string) (*WSGProfileCloneResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesAcct
//
// Use this API command to delete a list of accounting profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkAccountingProfile
func (s *WSGAccountingProfileService) DeleteProfilesAcct(ctx context.Context, body *WSGProfileDeleteBulkAccountingProfile) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesAcctById
//
// Use this API command to delete an accounting profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) DeleteProfilesAcctById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesAcct
//
// Use this API command to retrieve a list of accounting profiles.
func (s *WSGAccountingProfileService) FindProfilesAcct(ctx context.Context) (*WSGProfileAccountingProfileList, error) {
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
func (s *WSGAccountingProfileService) FindProfilesAcctById(ctx context.Context, pId string) (*WSGProfileAccountingProfile, error) {
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
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingProfileService) FindProfilesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileAccountingProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesAcctById
//
// Use this API command to modify the basic information of an accounting profile.
//
// Request Body:
//	 - body *WSGProfileModifyAccountingProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) PartialUpdateProfilesAcctById(ctx context.Context, body *WSGProfileModifyAccountingProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
