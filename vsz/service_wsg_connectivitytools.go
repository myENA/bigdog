package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/tool"
)

type WSGConnectivityToolsService struct {
	client *Client
}

func NewWSGConnectivityToolsService(client *Client) *WSGConnectivityToolsService {
	s := new(WSGConnectivityToolsService)
	s.client = client
	return s
}

func (ss *WSGService) WSGConnectivityToolsService() *WSGConnectivityToolsService {
	serv := new(WSGConnectivityToolsService)
	serv.client = ss.client
	return serv
}

func (s *WSGConnectivityToolsService) AddToolSpeedflex(ctx context.Context, body *tool.SpeedFlex) (*tool.TestResult, error) {
}

func (s *WSGConnectivityToolsService) FindToolPing(ctx context.Context, qApMac string, qTargetIP string) error {
}

func (s *WSGConnectivityToolsService) FindToolSpeedflexByWcid(ctx context.Context, pWcid string) (*tool.TestResult, error) {
}

func (s *WSGConnectivityToolsService) FindToolTraceRoute(ctx context.Context, qApMac string, qTargetIP string, qTimeoutInSec string) error {
}
