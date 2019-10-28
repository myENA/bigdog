package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGL2oGREService struct {
	client *Client
}

func NewWSGL2oGREService(client *Client) *WSGL2oGREService {
	s := new(WSGL2oGREService)
	s.client = client
	return s
}

func (ss *WSGService) WSGL2oGREService() *WSGL2oGREService {
	serv := new(WSGL2oGREService)
	serv.client = ss.client
	return serv
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
