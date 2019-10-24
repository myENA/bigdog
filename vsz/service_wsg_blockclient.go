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

func (s *WSGBlockClientService) AddBlockClient(ctx context.Context, body *profile.BulkBlockClient) (profile.CreateResultList, error) {
}

func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *profile.BlockClient, pApMac string) (*common.CreateResult, error) {
}

func (s *WSGBlockClientService) FindBlockClientById(ctx context.Context, pId string) (*profile.BlockClient, error) {
}

func (s *WSGBlockClientService) FindBlockClientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.BlockClientList, error) {
}

func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, pZoneId string) (*profile.BlockClientList, error) {
}
