package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/accountsecurityprofile"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAccountSecurityService struct {
	apiClient *APIClient
}

func NewWSGAccountSecurityService(c *APIClient) *WSGAccountSecurityService {
	s := new(WSGAccountSecurityService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountSecurityService() *WSGAccountSecurityService {
	serv := new(WSGAccountSecurityService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddAccountSecurity
//
// Use this API command to create the account security proile.
//
// Request Body:
//	 - body *accountsecurityprofile.Create
func (s *WSGAccountSecurityService) AddAccountSecurity(ctx context.Context, body *accountsecurityprofile.Create) (*common.CreateResultIdName, error) {
}

// DeleteAccountSecurity
//
// Use this API command to selete the account security profile.
//
// Request Body:
//	 - body *accountsecurityprofile.DeleteList
func (s *WSGAccountSecurityService) DeleteAccountSecurity(ctx context.Context, body *accountsecurityprofile.DeleteList) error {
}

// DeleteAccountSecurityById
//
// Use this API command to delete the account security profile by id.
//
// Request Body:
//	 - body *accountsecurityprofile.Delete
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) DeleteAccountSecurityById(ctx context.Context, body *accountsecurityprofile.Delete, pId string) (*common.CreateResultIdName, error) {
}

// FindAccountSecurity
//
// Use this API command to get account security profiles.
func (s *WSGAccountSecurityService) FindAccountSecurity(ctx context.Context) (*accountsecurityprofile.ProfileListResult, error) {
}

// FindAccountSecurityById
//
// Use this API command to retrieve the specific account security profile.
//
// Request Body:
//	 - body *accountsecurityprofile.GetById
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) FindAccountSecurityById(ctx context.Context, body *accountsecurityprofile.GetById, pId string) (*accountsecurityprofile.GetByIdResult, error) {
}

// PartialUpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *accountsecurityprofile.Update
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) PartialUpdateAccountSecurityById(ctx context.Context, body *accountsecurityprofile.Update, pId string) (*common.CreateResultIdName, error) {
}

// UpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *accountsecurityprofile.Update
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountSecurityService) UpdateAccountSecurityById(ctx context.Context, body *accountsecurityprofile.Update, pId string) (*common.CreateResultIdName, error) {
}
