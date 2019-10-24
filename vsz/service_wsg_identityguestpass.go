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

// AddIdentityGuestpassGenerate
//
// Use this API command to generate identity guest pass.
//
// Request Body:
//	 - body *identity.CreateIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassGenerate(ctx context.Context, body *identity.CreateIdentityGuestPass) (*common.CreateResult, error) {
}

// AddIdentityGuestpassList
//
// Use this API command to retrieve a list of identity guest pass.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassList(ctx context.Context, body *identity.QueryCriteria) (*identity.IdentityGuestPassList, error) {
}

// AddIdentityGuestpassUpload
//
// Use this API command to upload identity guest pass csv file.
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUpload(ctx context.Context) error {
}

// AddIdentityGuestpassUploadCommon
//
// Use this API command to update common identity guest pass settings.
//
// Request Body:
//	 - body *identity.ImportIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUploadCommon(ctx context.Context, body *identity.ImportIdentityGuestPass) error {
}

// DeleteIdentityGuestpassByUserId
//
// Use this API command to delete identity guest pass.
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpassByUserId(ctx context.Context, pUserId string) (*common.EmptyResult, error) {
}

// FindIdentityGuestpass
//
// Use this API command to retrieve a list of identity guest pass.
func (s *WSGIdentityGuestPassService) FindIdentityGuestpass(ctx context.Context, qDisplayName string, qExpirationFrom string, qExpirationTo string, qGeneratedTimeFrom string, qGeneratedTimeTo string, qIndex string, qListSize string, qTimeZone string, qWlan string) (*identity.IdentityGuestPassList, error) {
}
