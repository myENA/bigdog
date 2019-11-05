package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/staticroute"
)

type SwitchMSwitchStaticRouteService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchStaticRouteService(c *APIClient) *SwitchMSwitchStaticRouteService {
	s := new(SwitchMSwitchStaticRouteService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchStaticRouteService() *SwitchMSwitchStaticRouteService {
	serv := new(SwitchMSwitchStaticRouteService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddStaticRoutes
//
// Use this API command to Create Static Route.
//
// Request Body:
//	 - body *staticroute.CreateStaticRoute
func (s *SwitchMSwitchStaticRouteService) AddStaticRoutes(ctx context.Context, body *staticroute.CreateStaticRoute) (*common.CreateResult, error) {
}

// DeleteStaticRoutes
//
// Use this API command to Delete Static Route by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchStaticRouteService) DeleteStaticRoutes(ctx context.Context, body *common.BulkDeleteRequest) error {
}

// DeleteStaticRoutesById
//
// Use this API command to Delete Static Route.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchStaticRouteService) DeleteStaticRoutesById(ctx context.Context, pId string) error {
}

// FindStaticRoutesById
//
// Use this API command to Retrieve Static Route.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchStaticRouteService) FindStaticRoutesById(ctx context.Context, pId string) (*staticroute.StaticRoute, error) {
}

// FindStaticRoutesByQueryCriteria
//
// Use this API command to Retrieve Static Route list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchStaticRouteService) FindStaticRoutesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*staticroute.StaticRoutesQueryResult, error) {
}

// UpdateStaticRoutesById
//
// Use this API command to Update Static Route.
//
// Request Body:
//	 - body *staticroute.UpdateStaticRoute
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchStaticRouteService) UpdateStaticRoutesById(ctx context.Context, body *staticroute.UpdateStaticRoute, pId string) (*staticroute.EmptyResult, error) {
}
