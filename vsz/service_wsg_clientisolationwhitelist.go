package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGClientIsolationWhitelistService struct {
	client *Client
}

func NewWSGClientIsolationWhitelistService(client *Client) *WSGClientIsolationWhitelistService {
	s := new(WSGClientIsolationWhitelistService)
	s.client = client
	return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService() *WSGClientIsolationWhitelistService {
	serv := new(WSGClientIsolationWhitelistService)
	serv.client = ss.client
	return serv
}

// DeleteRkszonesClientIsolationWhitelist
//
// Use this API command to delete Bulk Client Isolation Whitelist.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelist(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteRkszonesClientIsolationWhitelistById
//
// Delete a Client Isolation Whitelist.
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelistById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindRkszonesClientIsolationWhitelistById
//
// Retrieve an Client Isolation Whitelist.
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistById(ctx context.Context, pId string, pZoneId string) (*profile.ClientIsolationWhitelist, error) {
}

// FindRkszonesClientIsolationWhitelistByZoneId
//
// Retrieve a list of Client Isolation Whitelist.
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, pZoneId string) (*profile.ClientIsolationWhitelistArray, error) {
}

// FindServicesClientIsolationWhitelistByQueryCriteria
//
// Retrieve a list of Client Isolation Whitelist.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGClientIsolationWhitelistService) FindServicesClientIsolationWhitelistByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.ClientIsolationWhitelistArray, error) {
}
