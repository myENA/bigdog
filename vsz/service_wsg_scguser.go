package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserService struct {
	client *Client
}

func NewWSGSCGUserService(client *Client) *WSGSCGUserService {
	s := new(WSGSCGUserService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSCGUserService() *WSGSCGUserService {
	serv := new(WSGSCGUserService)
	serv.client = ss.client
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
func (s *WSGSCGUserService) FindUsersByUserId(ctx context.Context, pUserId string) (*scguser.GetScgUser, error) {
}
