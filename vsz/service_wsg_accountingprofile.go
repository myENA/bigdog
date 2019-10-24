package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAccountingProfileService struct {
	client *Client
}

func NewWSGAccountingProfileService(client *Client) *WSGAccountingProfileService {
	s := new(WSGAccountingProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAccountingProfileService() *WSGAccountingProfileService {
	serv := new(WSGAccountingProfileService)
	serv.client = ss.client
	return serv
}

// AddProfilesAcctCloneById
//
// Use this API command to clone an accounting profile.
//
// Request Body:
//	 - body *profile.ProfileCloneRequest
func (s *WSGAccountingProfileService) AddProfilesAcctCloneById(ctx context.Context, body *profile.ProfileCloneRequest, pId string) (*profile.ProfileCloneResponse, error) {
}

// FindProfilesAcct
//
// Use this API command to retrieve a list of accounting profiles.
func (s *WSGAccountingProfileService) FindProfilesAcct(ctx context.Context) (*profile.AccountingProfileList, error) {
}

// FindProfilesAcctById
//
// Use this API command to retrieve an accounting profile.
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
