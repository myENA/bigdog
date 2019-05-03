package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserService struct {
    client *Client
}

func NewWSGSCGUserService (client *Client) *WSGSCGUserService {
    s := new(WSGSCGUserService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSCGUserService () *WSGSCGUserService {
    serv := new(WSGSCGUserService)
    serv.client = ss.client
    return serv
}

func (s *WSGSCGUserService) AddUsers (ctx context.Context) (*scguser.ScgUserAuditId, error) {
}

func (s *WSGSCGUserService) FindUsersByQueryCriteria (ctx context.Context) (*scguser.ScgUserList, error) {
}

func (s *WSGSCGUserService) FindUsersByUserId (ctx context.Context, userId string) (*scguser.ScgUser, error) {
}

