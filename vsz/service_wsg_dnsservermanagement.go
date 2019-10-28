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

func NewWSGDNSServerManagementService(client *Client) *WSGDNSServerManagementService {
	s := new(WSGDNSServerManagementService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDNSServerManagementService() *WSGDNSServerManagementService {
	serv := new(WSGDNSServerManagementService)
	serv.client = ss.client
	return serv
}

// AddProfilesDnsserverCloneById
//
// Use this API command to clone an DNS server profile.
//
// Request Body:
//	 - body *profile.ProfileCloneRequest
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) AddProfilesDnsserverCloneById(ctx context.Context, body *profile.ProfileCloneRequest, pId string) (*profile.ProfileCloneResponse, error) {
}

// FindProfilesDnsserver
//
// Use this API command to retrieve a list of DNS server profile.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGDNSServerManagementService) FindProfilesDnsserver(ctx context.Context, qIndex string, qListSize string) (*profile.DnsServerProfileList, error) {
}

// FindProfilesDnsserverById
//
// Use this API command to retrieve DNS server profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) FindProfilesDnsserverById(ctx context.Context, pId string) (*profile.DnsServerProfile, error) {
}

// FindProfilesDnsserverByQueryCriteria
//
// Use this API command to retrieve a list of DNS server profile  by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGDNSServerManagementService) FindProfilesDnsserverByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.DnsServerProfileList, error) {
}
