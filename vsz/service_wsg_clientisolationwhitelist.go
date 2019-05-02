package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGClientIsolationWhitelistService struct {
    client *Client
}

func NewWSGClientIsolationWhitelistService (client *Client) *WSGClientIsolationWhitelistService {
    s := new(WSGClientIsolationWhitelistService)
    s.client = client
    return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService () *WSGClientIsolationWhitelistService {
    serv := new(WSGClientIsolationWhitelistService)
    serv.client = ss.client
    return serv
}

func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelist (ctx context.Context) error {
}

func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelistById (ctx context.Context, id string) error {
}

func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistById (ctx context.Context, id string, zoneId string) (profile.ClientIsolationWhitelist, error) {
}

func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistByZoneId (ctx context.Context, zoneId string) (profile.ClientIsolationWhitelistArray, error) {
}

func (s *WSGClientIsolationWhitelistService) FindServicesClientIsolationWhitelistByQueryCriteria (ctx context.Context) (profile.ClientIsolationWhitelistArray, error) {
}

