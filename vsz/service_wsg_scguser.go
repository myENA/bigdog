package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserService struct {
	apiClient *APIClient
}

func NewWSGSCGUserService(c *APIClient) *WSGSCGUserService {
	s := new(WSGSCGUserService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCGUserService() *WSGSCGUserService {
	serv := new(WSGSCGUserService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddUsers
//
// Add SCG user.
//
// Request Body:
//	 - body *scguser.CreateScgUser
func (s *WSGSCGUserService) AddUsers(ctx context.Context, body *scguser.CreateScgUser) (*scguser.ScgUserAuditId, error) {
}

// DeleteUsers
//
// Delete multiple SCG user.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGSCGUserService) DeleteUsers(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteUsersByUserId
//
// Delete SCG user.
//
// Path Parameters:
// - pUserId string
//		- required
func (s *WSGSCGUserService) DeleteUsersByUserId(ctx context.Context, pUserId string) (*scguser.ScgUserAuditId, error) {
}

// FindUsersByQueryCriteria
//
// Query SCG users.
//
// Request Body:
//	 - body *scguser.QueryCriteria
func (s *WSGSCGUserService) FindUsersByQueryCriteria(ctx context.Context, body *scguser.QueryCriteria) (*scguser.ScgUserList, error) {
}

// FindUsersByUserId
//
// Get SCG user.
//
// Path Parameters:
// - pUserId string
//		- required
func (s *WSGSCGUserService) FindUsersByUserId(ctx context.Context, pUserId string) (*scguser.GetScgUser, error) {
}

// PartialUpdateUsersByUserId
//
// Update SCG user.
//
// Request Body:
//	 - body *scguser.ModifyScgUser
//
// Path Parameters:
// - pUserId string
//		- required
func (s *WSGSCGUserService) PartialUpdateUsersByUserId(ctx context.Context, body *scguser.ModifyScgUser, pUserId string) (*scguser.ScgUserAuditId, error) {
}
