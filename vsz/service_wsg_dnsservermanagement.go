package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGDNSServerManagementService struct {
    client *Client
}

func NewWSGDNSServerManagementService (client *Client) *WSGDNSServerManagementService {
    s := new(WSGDNSServerManagementService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDNSServerManagementService () *WSGDNSServerManagementService {
    serv := new(WSGDNSServerManagementService)
    serv.client = ss.client
    return serv
}

func (s *WSGDNSServerManagementService) AddProfilesDnsserverCloneById (ctx context.Context, id string) (*profile.ProfileCloneResponse, error) {
}

func (s *WSGDNSServerManagementService) FindProfilesDnsserver (ctx context.Context) (*profile.DNSServerProfileList, error) {
}

func (s *WSGDNSServerManagementService) FindProfilesDnsserverById (ctx context.Context, id string) (*profile.DNSServerProfile, error) {
}

func (s *WSGDNSServerManagementService) FindProfilesDnsserverByQueryCriteria (ctx context.Context) (*profile.DNSServerProfileList, error) {
}

