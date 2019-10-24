package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGL2AccessControlService struct {
	client *Client
}

func NewWSGL2AccessControlService(client *Client) *WSGL2AccessControlService {
	s := new(WSGL2AccessControlService)
	s.client = client
	return s
}

func (ss *WSGService) WSGL2AccessControlService() *WSGL2AccessControlService {
	serv := new(WSGL2AccessControlService)
	serv.client = ss.client
	return serv
}

func (s *WSGL2AccessControlService) FindRkszonesL2ACLById(ctx context.Context, pId string, pZoneId string) (*portalservice.L2ACL, error) {
}

func (s *WSGL2AccessControlService) FindRkszonesL2ACLByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*portalservice.PortalServiceList, error) {
}
