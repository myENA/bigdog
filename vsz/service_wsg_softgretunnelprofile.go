package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGSoftGRETunnelProfileService struct {
	client *Client
}

func NewWSGSoftGRETunnelProfileService(client *Client) *WSGSoftGRETunnelProfileService {
	s := new(WSGSoftGRETunnelProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSoftGRETunnelProfileService() *WSGSoftGRETunnelProfileService {
	serv := new(WSGSoftGRETunnelProfileService)
	serv.client = ss.client
	return serv
}

// FindProfilesTunnelSoftgre
//
// Use this API command to retrieve a list of SoftGRE tunnel profile.
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgre(ctx context.Context) (*profile.ProfileList, error) {
}

// FindProfilesTunnelSoftgreById
//
// Use this API command to retrieve SoftGRE tunnel profile.
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreById(ctx context.Context, pId string) (*profile.SoftGREProfile, error) {
}

// FindProfilesTunnelSoftgreByQueryCriteria
//
// Use this API command to query a list of SoftGRE tunnel profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.SoftGREProfileList, error) {
}
