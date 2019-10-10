package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGSMSGatewayService struct {
	client *Client
}

func NewWSGSMSGatewayService(client *Client) *WSGSMSGatewayService {
	s := new(WSGSMSGatewayService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSMSGatewayService() *WSGSMSGatewayService {
	serv := new(WSGSMSGatewayService)
	serv.client = ss.client
	return serv
}

func (s *WSGSMSGatewayService) FindSmsGateway(ctx context.Context) (*system.Sms, error) {
}

func (s *WSGSMSGatewayService) FindSmsGatewayByQueryCriteria(ctx context.Context) (*system.SmsList, error) {
}
