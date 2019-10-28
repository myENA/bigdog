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

// AddClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *client.DeAuthClientList
func (s *WSGWirelessClientService) AddClientsBulkDeauth(ctx context.Context, body *client.DeAuthClientList) (*common.EmptyResult, error) {
}

// AddClientsBulkDisconnect
//
// Use this API command to bulk disconnect client.
//
// Request Body:
//	 - body *client.DisconnectClientList
func (s *WSGWirelessClientService) AddClientsBulkDisconnect(ctx context.Context, body *client.DisconnectClientList) (*common.EmptyResult, error) {
}

// AddClientsByWlanNameByWlanname
//
// Use this API command to query client by wlan name.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
//
// Path Parameters:
// - pWlanname string
//		- required
func (s *WSGWirelessClientService) AddClientsByWlanNameByWlanname(ctx context.Context, body *common.QueryCriteriaSuperSet, pWlanname string) (*clientquery.ClientQueryList, error) {
}

// AddClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *client.DeAuthClient
func (s *WSGWirelessClientService) AddClientsDeauth(ctx context.Context, body *client.DeAuthClient) (*common.EmptyResult, error) {
}

// AddClientsDisconnect
//
// Use this API command to disconnect client.
//
// Request Body:
//	 - body *client.DisconnectClient
func (s *WSGWirelessClientService) AddClientsDisconnect(ctx context.Context, body *client.DisconnectClient) (*common.EmptyResult, error) {
}

// FindApsOperationalClientByApMac
//
// Use this API command to retrieve the client list per AP.
//
// Path Parameters:
// - pApMac string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGWirelessClientService) FindApsOperationalClientByApMac(ctx context.Context, pApMac string, qIndex string, qListSize string) (*ap.ClientList, error) {
}

// FindApsOperationalClientTotalCountByApMac
//
// Use this API command to retrieve the total client count per AP.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGWirelessClientService) FindApsOperationalClientTotalCountByApMac(ctx context.Context, pApMac string) error {
}

// FindHistoricalclientByQueryCriteria
//
// Use this API command to retrive historical client.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGWirelessClientService) FindHistoricalclientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*client.HistoricalClientList, error) {
}
