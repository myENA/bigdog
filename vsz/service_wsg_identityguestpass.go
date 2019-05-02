package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityGuestPassService struct {
    client *Client
}

func NewWSGIdentityGuestPassService (client *Client) *WSGIdentityGuestPassService {
    s := new(WSGIdentityGuestPassService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIdentityGuestPassService () *WSGIdentityGuestPassService {
    serv := new(WSGIdentityGuestPassService)
    serv.client = ss.client
    return serv
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassGenerate (ctx context.Context) (common.CreateResult, error) {
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassList (ctx context.Context) (identity.IdentityGuestPassList, error) {
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUpload (ctx context.Context) error {
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUploadCommon (ctx context.Context) error {
}

func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpassByUserId (ctx context.Context, userId string) error {
}

func (s *WSGIdentityGuestPassService) FindIdentityGuestpass (ctx context.Context) (identity.IdentityGuestPassList, error) {
}

