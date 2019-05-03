package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBlockClientService struct {
    client *Client
}

func NewWSGBlockClientService (client *Client) *WSGBlockClientService {
    s := new(WSGBlockClientService)
    s.client = client
    return s
}

func (ss *WSGService) WSGBlockClientService () *WSGBlockClientService {
    serv := new(WSGBlockClientService)
    serv.client = ss.client
    return serv
}

func (s *WSGBlockClientService) AddBlockClient (ctx context.Context) (profile.CreateResultList, error) {
}

func (s *WSGBlockClientService) AddBlockClientByApMacByApMac (ctx context.Context, apMac string) (*common.CreateResult, error) {
}

func (s *WSGBlockClientService) FindBlockClientById (ctx context.Context, id string) (*profile.BlockClient, error) {
}

func (s *WSGBlockClientService) FindBlockClientByQueryCriteria (ctx context.Context) (*profile.BlockClientList, error) {
}

func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId (ctx context.Context, zoneId string) (*profile.BlockClientList, error) {
}

