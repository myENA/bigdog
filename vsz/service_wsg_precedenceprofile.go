package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGPrecedenceProfileService struct {
	apiClient *APIClient
}

func NewWSGPrecedenceProfileService(c *APIClient) *WSGPrecedenceProfileService {
	s := new(WSGPrecedenceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGPrecedenceProfileService() *WSGPrecedenceProfileService {
	serv := new(WSGPrecedenceProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddPrecedence
//
// Use this API command to create Precedence Profile.
//
// Request Body:
//	 - body *profile.CreatePrecedenceProfile
func (s *WSGPrecedenceProfileService) AddPrecedence(ctx context.Context, body *profile.CreatePrecedenceProfile) (*common.CreateResult, error) {
}

// DeletePrecedence
//
// Use this API command to Bulk Delete Precedence Profile.
//
// Request Body:
//	 - body *profile.DeleteBulkPrecedenceProfile
func (s *WSGPrecedenceProfileService) DeletePrecedence(ctx context.Context, body *profile.DeleteBulkPrecedenceProfile) (*common.EmptyResult, error) {
}

// DeletePrecedenceById
//
// Use this API command to Delete Precedence Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGPrecedenceProfileService) DeletePrecedenceById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindPrecedence
//
// Use this API command to Get Precedence Profile list.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGPrecedenceProfileService) FindPrecedence(ctx context.Context, qIndex string, qListSize string) (*profile.PrecedenceList, error) {
}

// FindPrecedenceById
//
// Use this API command to Get Precedence Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGPrecedenceProfileService) FindPrecedenceById(ctx context.Context, pId string) (*profile.CreatePrecedenceProfile, error) {
}

// FindPrecedenceByQueryCriteria
//
// Use this API command to query Precedence Profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGPrecedenceProfileService) FindPrecedenceByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.PrecedenceList, error) {
}

// PartialUpdatePrecedenceById
//
// Use this API command to Modify Precedence Profile by profile's ID.
//
// Request Body:
//	 - body *profile.UpdatePrecedenceProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGPrecedenceProfileService) PartialUpdatePrecedenceById(ctx context.Context, body *profile.UpdatePrecedenceProfile, pId string) (*common.EmptyResult, error) {
}
