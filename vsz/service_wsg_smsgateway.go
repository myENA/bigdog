package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGSMSGatewayService struct {
	apiClient *APIClient
}

func NewWSGSMSGatewayService(c *APIClient) *WSGSMSGatewayService {
	s := new(WSGSMSGatewayService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSMSGatewayService() *WSGSMSGatewayService {
	serv := new(WSGSMSGatewayService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindSmsGateway
//
// Get SMS gateway.
//
// Query Parameters:
// - qDomainId string
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

// PartialUpdateSmsGateway
//
// Update SMS gateway.
//
// Request Body:
//	 - body *system.Sms
func (s *WSGSMSGatewayService) PartialUpdateSmsGateway(ctx context.Context, body *system.Sms) error {
}
