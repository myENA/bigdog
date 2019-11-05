package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGGGSNPGWServiceService struct {
	apiClient *APIClient
}

func NewWSGGGSNPGWServiceService(c *APIClient) *WSGGGSNPGWServiceService {
	s := new(WSGGGSNPGWServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGGSNPGWServiceService() *WSGGGSNPGWServiceService {
	serv := new(WSGGGSNPGWServiceService)
	serv.apiClient = ss.apiClient
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

// PartialUpdateServicesGgsn
//
// Use this API command to modify the setting of GGSN/PGW.
//
// Request Body:
//	 - body *service.GgsnConfig
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsn(ctx context.Context, body *service.GgsnConfig) (*common.EmptyResult, error) {
}

// PartialUpdateServicesGgsnDnsServerList
//
// Use this API command to modify the dns server list of GGSN/PGW.
//
// Request Body:
//	 - body service.DnsServerList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnDnsServerList(ctx context.Context, body service.DnsServerList) (*common.EmptyResult, error) {
}

// PartialUpdateServicesGgsnGgsnList
//
// Use this API command to modify the ggsn server list of GGSN/PGW.
//
// Request Body:
//	 - body service.GgsnList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGgsnList(ctx context.Context, body service.GgsnList) (*common.EmptyResult, error) {
}

// PartialUpdateServicesGgsnGtpSettings
//
// Use this API command to modify the gtp setting of GGSN/PGW.
//
// Request Body:
//	 - body *service.GtpSettings
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGtpSettings(ctx context.Context, body *service.GtpSettings) (*common.EmptyResult, error) {
}
