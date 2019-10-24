package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityGuestPassService struct {
	client *Client
}

func NewWSGIdentityGuestPassService(client *Client) *WSGIdentityGuestPassService {
	s := new(WSGIdentityGuestPassService)
	s.client = client
	return s
}

func (ss *WSGService) WSGIdentityGuestPassService() *WSGIdentityGuestPassService {
	serv := new(WSGIdentityGuestPassService)
	serv.client = ss.client
	return serv
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassGenerate(ctx context.Context, body *identity.CreateIdentityGuestPass) (*common.CreateResult, error) {
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassList(ctx context.Context, body *identity.QueryCriteria) (*identity.IdentityGuestPassList, error) {
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUpload(ctx context.Context) error {
}

func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUploadCommon(ctx context.Context, body *identity.ImportIdentityGuestPass) error {
}

func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpassByUserId(ctx context.Context, pUserId string) (*common.EmptyResult, error) {
}

func (s *WSGIdentityGuestPassService) FindIdentityGuestpass(ctx context.Context, qDisplayName string, qExpirationFrom string, qExpirationTo string, qGeneratedTimeFrom string, qGeneratedTimeTo string, qIndex string, qListSize string, qTimeZone string, qWlan string) (*identity.IdentityGuestPassList, error) {
}
