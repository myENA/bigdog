package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGGGSNPGWServiceService struct {
	client *Client
}

func NewWSGGGSNPGWServiceService(client *Client) *WSGGGSNPGWServiceService {
	s := new(WSGGGSNPGWServiceService)
	s.client = client
	return s
}

func (ss *WSGService) WSGGGSNPGWServiceService() *WSGGGSNPGWServiceService {
	serv := new(WSGGGSNPGWServiceService)
	serv.client = ss.client
	return serv
}

// DeleteServicesGgsnDnsServerList
//
// Use this API command to Disable the dns server list of GGSN/PGW.
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnDnsServerList(ctx context.Context) (*common.EmptyResult, error) {
}

// DeleteServicesGgsnGgsnList
//
// Use this API command to disable the ggsn server list of GGSN/PGW.
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnGgsnList(ctx context.Context) (*common.EmptyResult, error) {
}

// FindServicesGgsn
//
// Use this API command to retrieve GGSN/PGW setting.
func (s *WSGGGSNPGWServiceService) FindServicesGgsn(ctx context.Context) (*service.GgsnConfig, error) {
}

// PartialUpdateServicesGgsnGtpSettings
//
// Use this API command to modify the gtp setting of GGSN/PGW.
//
// Request Body:
//	 - body *service.GtpSettings
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGtpSettings(ctx context.Context, body *service.GtpSettings) (*common.EmptyResult, error) {
}
