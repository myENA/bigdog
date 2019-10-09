package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGLBSprofileService struct {
    client *Client
}

func NewWSGLBSprofileService (client *Client) *WSGLBSprofileService {
    s := new(WSGLBSprofileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGLBSprofileService () *WSGLBSprofileService {
    serv := new(WSGLBSprofileService)
    serv.client = ss.client
    return serv
}

func (s *WSGLBSprofileService) AddProfilesLbs (ctx context.Context) (*common.CreateResult, error) {
}

func (s *WSGLBSprofileService) FindProfilesLbsById (ctx context.Context, id string) (*profile.LbsProfile, error) {
}

func (s *WSGLBSprofileService) FindProfilesLbsByQueryCriteria (ctx context.Context) (*profile.LbsProfileList, error) {
}

