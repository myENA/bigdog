package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
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
	serv := new(WSGAccountingProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesAcct
//
// Use this API command to create a new accounting profile.
//
// Request Body:
//	 - body *profile.CreateAccountingProfile
func (s *WSGAccountingProfileService) AddProfilesAcct(ctx context.Context, body *profile.CreateAccountingProfile) (*common.CreateResult, error) {
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
}

// DeleteProfilesAcct
//
// Use this API command to delete a list of accounting profile.
//
// Request Body:
//	 - body *profile.DeleteBulkAccountingProfile
func (s *WSGAccountingProfileService) DeleteProfilesAcct(ctx context.Context, body *profile.DeleteBulkAccountingProfile) (*common.EmptyResult, error) {
}

// DeleteProfilesAcctById
//
// Use this API command to delete an accounting profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) DeleteProfilesAcctById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindProfilesAcct
//
// Use this API command to retrieve a list of accounting profiles.
func (s *WSGAccountingProfileService) FindProfilesAcct(ctx context.Context) (*profile.AccountingProfileList, error) {
}

// FindProfilesAcctById
//
// Use this API command to retrieve an accounting profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingProfileService) FindProfilesAcctById(ctx context.Context, pId string) (*profile.AccountingProfile, error) {
}

// FindProfilesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting profiles by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingProfileService) FindProfilesAcctByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.AccountingProfileList, error) {
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
}
