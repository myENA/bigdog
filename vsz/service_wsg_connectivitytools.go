package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/tool"
)

type WSGConnectivityToolsService struct {
	apiClient *APIClient
}

func NewWSGConnectivityToolsService(c *APIClient) *WSGConnectivityToolsService {
	s := new(WSGConnectivityToolsService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGConnectivityToolsService() *WSGConnectivityToolsService {
	serv := new(WSGConnectivityToolsService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddToolSpeedflex
//
// Use this API command to start the SpeedFlex test.
//
// Request Body:
//	 - body *tool.SpeedFlex
func (s *WSGConnectivityToolsService) AddToolSpeedflex(ctx context.Context, body *tool.SpeedFlex) (*tool.TestResult, error) {
}

// FindToolPing
//
// Use this API command to run the PING test on an AP.
//
// Query Parameters:
// - qApMac string
//		- required
// - qTargetIP string
//		- required
func (s *WSGConnectivityToolsService) FindToolPing(ctx context.Context, qApMac string, qTargetIP string) error {
}

// FindToolSpeedflexByWcid
//
// Use this API command to retrieve existing SpeedFlex test results.
//
// Path Parameters:
// - pWcid string
//		- required
func (s *WSGConnectivityToolsService) FindToolSpeedflexByWcid(ctx context.Context, pWcid string) (*tool.TestResult, error) {
}

// FindToolTraceRoute
//
// Use this API command to run the traceroute test on an AP.
//
// Query Parameters:
// - qApMac string
//		- required
// - qTargetIP string
//		- required
// - qTimeoutInSec string
func (s *WSGConnectivityToolsService) FindToolTraceRoute(ctx context.Context, qApMac string, qTargetIP string, qTimeoutInSec string) error {
}
