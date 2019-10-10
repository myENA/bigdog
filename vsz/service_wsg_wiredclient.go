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

func (s *WSGWiredClientService) AddWiredClientsBulkDeauth(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGWiredClientService) AddWiredClientsDeauth(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGWiredClientService) FindWiredclientByQueryCriteria(ctx context.Context) (*wiredclientquery.ClientQueryList, error) {
}
