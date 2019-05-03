package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAccountingServiceService struct {
    client *Client
}

func NewWSGAccountingServiceService (client *Client) *WSGAccountingServiceService {
    s := new(WSGAccountingServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccountingServiceService () *WSGAccountingServiceService {
    serv := new(WSGAccountingServiceService)
    serv.client = ss.client
    return serv
}

func (s *WSGAccountingServiceService) AddServicesAcctTestById (ctx context.Context, id string) error {
}

func (s *WSGAccountingServiceService) DeleteServicesAcct (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGAccountingServiceService) DeleteServicesAcctById (ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById (ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById (ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria (ctx context.Context) (*service.CommonAccountingServiceList, error) {
}

func (s *WSGAccountingServiceService) FindServicesAcctRadius (ctx context.Context) (*service.RadiusAccountingServiceList, error) {
}

func (s *WSGAccountingServiceService) FindServicesAcctRadiusById (ctx context.Context, id string) (*service.RadiusAccountingService, error) {
}

func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria (ctx context.Context) (*service.RadiusAccountingServiceList, error) {
}

