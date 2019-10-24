package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBlockClientService struct {
	client *Client
}

func NewWSGBlockClientService(client *Client) *WSGBlockClientService {
	s := new(WSGBlockClientService)
	s.client = client
	return s
}

func (ss *WSGService) WSGBlockClientService() *WSGBlockClientService {
	serv := new(WSGBlockClientService)
	serv.client = ss.client
	return serv
}

// AddBlockClient
//
// Create new Block Clients by list.
//
// Request Body:
//	 - body *profile.BulkBlockClient
func (s *WSGBlockClientService) AddBlockClient(ctx context.Context, body *profile.BulkBlockClient) (profile.CreateResultList, error) {
}

// AddBlockClientByApMacByApMac
//
// Create a new Block Client by AP MAC.
//
// Request Body:
//	 - body *profile.BlockClient
func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *profile.BlockClient, pApMac string) (*common.CreateResult, error) {
}

// FindBlockClientById
//
// Retrieve a Block Client.
func (s *WSGBlockClientService) FindBlockClientById(ctx context.Context, pId string) (*profile.BlockClient, error) {
}

// FindBlockClientByQueryCriteria
//
// Retrieve a list of Block Client.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGBlockClientService) FindBlockClientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.BlockClientList, error) {
}

// FindBlockClientByZoneByZoneId
//
// Retrieve a list of Block Client.
func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, pZoneId string) (*profile.BlockClientList, error) {
}
