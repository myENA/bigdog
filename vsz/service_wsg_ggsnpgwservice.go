package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGGGSNPGWServiceService struct {
    client *Client
}

func NewWSGGGSNPGWServiceService (client *Client) *WSGGGSNPGWServiceService {
    s := new(WSGGGSNPGWServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGGGSNPGWServiceService () *WSGGGSNPGWServiceService {
    serv := new(WSGGGSNPGWServiceService)
    serv.client = ss.client
    return serv
}

func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnDnsServerList (ctx context.Context) error {
}

func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnGgsnList (ctx context.Context) error {
}

func (s *WSGGGSNPGWServiceService) FindServicesGgsn (ctx context.Context) (service.GgsnConfig, error) {
}

func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGtpSettings (ctx context.Context) error {
}
