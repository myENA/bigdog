package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
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
	serv := new(WSGL2oGREService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesL2ogre
//
// Use this API command to create L2oGRE profile.
//
// Request Body:
//	 - body *profile.CreateL2oGREProfile
func (s *WSGL2oGREService) AddProfilesL2ogre(ctx context.Context, body *profile.CreateL2oGREProfile) (*common.CreateResult, error) {
}

// DeleteProfilesL2ogre
//
// Use this API command to delete multiple L2oGRE profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGL2oGREService) DeleteProfilesL2ogre(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteProfilesL2ogreById
//
// Use this API command to delete L2oGRE profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGL2oGREService) DeleteProfilesL2ogreById(ctx context.Context, pId string) error {
}

// FindProfilesL2ogre
//
// Use this API command to retrieve a list of L2oGRE profile.
func (s *WSGL2oGREService) FindProfilesL2ogre(ctx context.Context) (*profile.ProfileList, error) {
}

// FindProfilesL2ogreById
//
// Use this API command to retrieve L2oGRE profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGL2oGREService) FindProfilesL2ogreById(ctx context.Context, pId string) (*profile.L2oGREProfile, error) {
}

// FindProfilesL2ogreByQueryCriteria
//
// Use this API command to query a list of L2oGRE profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGL2oGREService) FindProfilesL2ogreByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.L2oGREProfileList, error) {
}

// PartialUpdateProfilesL2ogreById
//
// Use this API command to modify the basic information of L2oGRE profile.
//
// Request Body:
//	 - body *profile.ModifyL2oGREProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGL2oGREService) PartialUpdateProfilesL2ogreById(ctx context.Context, body *profile.ModifyL2oGREProfile, pId string) (*common.EmptyResult, error) {
}
