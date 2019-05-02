package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAccountingProfileService struct {
    client *Client
}

func NewWSGAccountingProfileService (client *Client) *WSGAccountingProfileService {
    s := new(WSGAccountingProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccountingProfileService () *WSGAccountingProfileService {
    serv := new(WSGAccountingProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGAccountingProfileService) AddProfilesAcctCloneById (ctx context.Context, id string) (profile.ProfileCloneResponse, error) {
}

func (s *WSGAccountingProfileService) FindProfilesAcct (ctx context.Context) (profile.AccountingProfileList, error) {
}

func (s *WSGAccountingProfileService) FindProfilesAcctById (ctx context.Context, id string) (profile.AccountingProfile, error) {
}

func (s *WSGAccountingProfileService) FindProfilesAcctByQueryCriteria (ctx context.Context) (profile.AccountingProfileList, error) {
}

