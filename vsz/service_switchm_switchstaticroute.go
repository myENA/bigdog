package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/staticroute"
)

type SwitchMSwitchStaticRouteService struct {
	client *Client
}

func NewSwitchMSwitchStaticRouteService(client *Client) *SwitchMSwitchStaticRouteService {
	s := new(SwitchMSwitchStaticRouteService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchStaticRouteService() *SwitchMSwitchStaticRouteService {
	serv := new(SwitchMSwitchStaticRouteService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchStaticRouteService) AddStaticRoutes(ctx context.Context, body *staticroute.CreateStaticRoute) (*common.CreateResult, error) {
}

func (s *SwitchMSwitchStaticRouteService) FindStaticRoutesById(ctx context.Context, pId string) (*staticroute.StaticRoute, error) {
}

func (s *SwitchMSwitchStaticRouteService) FindStaticRoutesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*staticroute.StaticRoutesQueryResult, error) {
}
