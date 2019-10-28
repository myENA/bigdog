package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGControlPlanesService struct {
	client *Client
}

func NewWSGControlPlanesService(client *Client) *WSGControlPlanesService {
	s := new(WSGControlPlanesService)
	s.client = client
	return s
}

func (ss *WSGService) WSGControlPlanesService() *WSGControlPlanesService {
	serv := new(WSGControlPlanesService)
	serv.client = ss.client
	return serv
}

// FindControlPlanes
//
// Use this API command to retrieve the list of control plane.
func (s *WSGControlPlanesService) FindControlPlanes(ctx context.Context) (*system.ControlPlaneList, error) {
}

// FindControlPlanesByBladeUUID
//
// Use this API command to retrieve control plane.
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesByBladeUUID(ctx context.Context, pBladeUUID string) (*system.ControlPlaneConfiguration, error) {
}

// FindControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to retrieve static route of control plane.
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesStaticRoutesByBladeUUID(ctx context.Context, pBladeUUID string) (*system.StaticRouteList, error) {
}

// FindControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to retrieve user defined interface of control plane.
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, pBladeUUID string) (*system.UserDefinedInterfaceList, error) {
}

// PartialUpdateControlPlanesIpSupport
//
// Use this API command to modify ip support of control plane.
//
// Request Body:
//	 - body *system.ModifyIpSupportType
func (s *WSGControlPlanesService) PartialUpdateControlPlanesIpSupport(ctx context.Context, body *system.ModifyIpSupportType) (*common.EmptyResult, error) {
}
