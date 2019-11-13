package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGL2oGREService struct {
	apiClient *APIClient
}

func NewWSGL2oGREService(c *APIClient) *WSGL2oGREService {
	s := new(WSGL2oGREService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2oGREService() *WSGL2oGREService {
	return NewWSGL2oGREService(ss.apiClient)
}

// AddProfilesL2ogre
//
// Use this API command to create L2oGRE profile.
//
// Request Body:
//	 - body *WSGProfileCreateL2oGREProfile
func (s *WSGL2oGREService) AddProfilesL2ogre(ctx context.Context, body *WSGProfileCreateL2oGREProfile) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesL2ogre
//
// Use this API command to delete multiple L2oGRE profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL2oGREService) DeleteProfilesL2ogre(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesL2ogreById
//
// Use this API command to delete L2oGRE profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGL2oGREService) DeleteProfilesL2ogreById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesL2ogre
//
// Use this API command to retrieve a list of L2oGRE profile.
func (s *WSGL2oGREService) FindProfilesL2ogre(ctx context.Context) (*WSGProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesL2ogreById
//
// Use this API command to retrieve L2oGRE profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGL2oGREService) FindProfilesL2ogreById(ctx context.Context, pId string) (*WSGProfileL2oGREProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesL2ogreByQueryCriteria
//
// Use this API command to query a list of L2oGRE profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2oGREService) FindProfilesL2ogreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileL2oGREProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesL2ogreById
//
// Use this API command to modify the basic information of L2oGRE profile.
//
// Request Body:
//	 - body *WSGProfileModifyL2oGREProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGL2oGREService) PartialUpdateProfilesL2ogreById(ctx context.Context, body *WSGProfileModifyL2oGREProfile, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
