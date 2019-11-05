package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBlockClientService struct {
	apiClient *APIClient
}

func NewWSGBlockClientService(c *APIClient) *WSGBlockClientService {
	s := new(WSGBlockClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBlockClientService() *WSGBlockClientService {
	serv := new(WSGBlockClientService)
	serv.apiClient = ss.apiClient
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
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *profile.BlockClient, pApMac string) (*common.CreateResult, error) {
}

// DeleteBlockClient
//
// Delete Block Client List.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGBlockClientService) DeleteBlockClient(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteBlockClientById
//
// Delete a Block Client.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) DeleteBlockClientById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindBlockClientById
//
// Retrieve a Block Client.
//
// Path Parameters:
// - pId string
//		- required
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
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, pZoneId string) (*profile.BlockClientList, error) {
}

// PartialUpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *profile.ModifyBlockClient
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) PartialUpdateBlockClientById(ctx context.Context, body *profile.ModifyBlockClient, pId string) (*common.EmptyResult, error) {
}

// UpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *profile.ModifyBlockClient
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBlockClientService) UpdateBlockClientById(ctx context.Context, body *profile.ModifyBlockClient, pId string) (*common.EmptyResult, error) {
}
