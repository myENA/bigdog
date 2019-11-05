package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGDNSServerManagementService struct {
	apiClient *APIClient
}

func NewWSGDNSServerManagementService(c *APIClient) *WSGDNSServerManagementService {
	s := new(WSGDNSServerManagementService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDNSServerManagementService() *WSGDNSServerManagementService {
	serv := new(WSGDNSServerManagementService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesDnsserver
//
// Use this API command to create DNS server profile.
//
// Request Body:
//	 - body *profile.CreateDnsServerProfile
func (s *WSGDNSServerManagementService) AddProfilesDnsserver(ctx context.Context, body *profile.CreateDnsServerProfile) (*common.CreateResult, error) {
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

// DeleteProfilesDnsserver
//
// Use this API command to delete a list of DNS server profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserver(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteProfilesDnsserverById
//
// Use this API command to delete DNS server profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserverById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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

// PartialUpdateProfilesDnsserverById
//
// Use this API command to modify the basic information of DNS server profile.
//
// Request Body:
//	 - body *profile.ModifyDnsServerProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDNSServerManagementService) PartialUpdateProfilesDnsserverById(ctx context.Context, body *profile.ModifyDnsServerProfile, pId string) (*common.EmptyResult, error) {
}
