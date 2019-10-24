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

func (s *WSGL2oGREService) FindProfilesL2ogre(ctx context.Context) (*profile.ProfileList, error) {
}

func (s *WSGL2oGREService) FindProfilesL2ogreById(ctx context.Context, pId string) (*profile.L2oGREProfile, error) {
}

func (s *WSGL2oGREService) FindProfilesL2ogreByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.L2oGREProfileList, error) {
}
