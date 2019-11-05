package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/client"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clientquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"gopkg.in/go-playground/validator.v9"
)

type WSGWirelessClientService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGWirelessClientService(c *APIClient) *WSGWirelessClientService {
	s := new(WSGWirelessClientService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGWirelessClientService() *WSGWirelessClientService {
	return NewWSGWirelessClientService(ss.apiClient)
}

// AddClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *client.DeAuthClientList
func (s *WSGWirelessClientService) AddClientsBulkDeauth(ctx context.Context, body *client.DeAuthClientList) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// AddClientsBulkDisconnect
//
// Use this API command to bulk disconnect client.
//
// Request Body:
//	 - body *client.DisconnectClientList
func (s *WSGWirelessClientService) AddClientsBulkDisconnect(ctx context.Context, body *client.DisconnectClientList) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// AddClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *client.DeAuthClient
func (s *WSGWirelessClientService) AddClientsDeauth(ctx context.Context, body *client.DeAuthClient) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddClientsDisconnect
//
// Use this API command to disconnect client.
//
// Request Body:
//	 - body *client.DisconnectClient
func (s *WSGWirelessClientService) AddClientsDisconnect(ctx context.Context, body *client.DisconnectClient) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsOperationalClientTotalCountByApMac
//
// Use this API command to retrieve the total client count per AP.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGWirelessClientService) FindApsOperationalClientTotalCountByApMac(ctx context.Context, pApMac string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindHistoricalclientByQueryCriteria
//
// Use this API command to retrive historical client.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGWirelessClientService) FindHistoricalclientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*client.HistoricalClientList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}
