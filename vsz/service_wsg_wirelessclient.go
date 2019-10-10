package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/client"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clientquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGWirelessClientService struct {
	client *Client
}

func NewWSGWirelessClientService(client *Client) *WSGWirelessClientService {
	s := new(WSGWirelessClientService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWirelessClientService() *WSGWirelessClientService {
	serv := new(WSGWirelessClientService)
	serv.client = ss.client
	return serv
}

func (s *WSGWirelessClientService) AddClientsBulkDeauth(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGWirelessClientService) AddClientsBulkDisconnect(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGWirelessClientService) AddClientsByWlanNameByWlanname(ctx context.Context, wlanname string) (*clientquery.ClientQueryList, error) {
}

func (s *WSGWirelessClientService) AddClientsDeauth(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGWirelessClientService) AddClientsDisconnect(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGWirelessClientService) FindApsOperationalClientByApMac(ctx context.Context, apMac string) (*ap.ClientList, error) {
}

func (s *WSGWirelessClientService) FindApsOperationalClientTotalCountByApMac(ctx context.Context, apMac string) error {
}

func (s *WSGWirelessClientService) FindHistoricalclientByQueryCriteria(ctx context.Context) (*client.HistoricalClientList, error) {
}
