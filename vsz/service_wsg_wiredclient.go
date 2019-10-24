package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/client"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wiredclientquery"
)

type WSGWiredClientService struct {
	client *Client
}

func NewWSGWiredClientService(client *Client) *WSGWiredClientService {
	s := new(WSGWiredClientService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWiredClientService() *WSGWiredClientService {
	serv := new(WSGWiredClientService)
	serv.client = ss.client
	return serv
}

// AddWiredClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *client.DeAuthClientList
func (s *WSGWiredClientService) AddWiredClientsBulkDeauth(ctx context.Context, body *client.DeAuthClientList) (*common.EmptyResult, error) {
}

// AddWiredClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *client.DeAuthClient
func (s *WSGWiredClientService) AddWiredClientsDeauth(ctx context.Context, body *client.DeAuthClient) (*common.EmptyResult, error) {
}

// FindWiredclientByQueryCriteria
//
// Query wired clients with specified filters
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGWiredClientService) FindWiredclientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*wiredclientquery.ClientQueryList, error) {
}
