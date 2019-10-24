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

// FindSmsGateway
//
// Get SMS gateway.
func (s *WSGSMSGatewayService) FindSmsGateway(ctx context.Context, qDomainId string) (*system.Sms, error) {
}

// FindSmsGatewayByQueryCriteria
//
// Query SMS gateway.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGSMSGatewayService) FindSmsGatewayByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*system.SmsList, error) {
}
