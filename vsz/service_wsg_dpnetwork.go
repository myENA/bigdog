package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGDPNetworkService struct {
	apiClient *APIClient
}

func NewWSGDPNetworkService(c *APIClient) *WSGDPNetworkService {
	s := new(WSGDPNetworkService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPNetworkService() *WSGDPNetworkService {
	serv := new(WSGDPNetworkService)
	serv.apiClient = ss.apiClient
	return serv
}

// DeletePlanesStaticRouteByBladeUUID
//
// Use this API command to delete static route.
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGDPNetworkService) DeletePlanesStaticRouteByBladeUUID(ctx context.Context, pBladeUUID string) (*common.EmptyResult, error) {
}

// FindPlanes
//
// Use this API command to retrieve a list of data planes.
func (s *WSGDPNetworkService) FindPlanes(ctx context.Context) (*system.DataPlaneList, error) {
}

// FindPlanesByBladeUUID
//
// Use this API command to retrieve data plane by id.
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGDPNetworkService) FindPlanesByBladeUUID(ctx context.Context, pBladeUUID string) (*system.DataPlaneConfiguration, error) {
}

// FindPlanesDpTunnelSetting
//
// Use this API command to get DP mesh tunnel setting.
func (s *WSGDPNetworkService) FindPlanesDpTunnelSetting(ctx context.Context) (*system.GetDataPlaneMeshTunnelSetting, error) {
}

// PartialUpdatePlanesByBladeUUID
//
// Use this API command to modify the basic information of data plane.
//
// Request Body:
//	 - body *system.ModifyDataPlane
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlanesByBladeUUID(ctx context.Context, body *system.ModifyDataPlane, pBladeUUID string) (*common.EmptyResult, error) {
}

// PartialUpdatePlaneStatesByBladeUUID
//
// Use this API command to update DP state profile.
//
// Request Body:
//	 - body *system.ModifyDataPlaneState
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlaneStatesByBladeUUID(ctx context.Context, body *system.ModifyDataPlaneState, pBladeUUID string) (*common.EmptyResult, error) {
}

// UpdatePlanesDpTunnelSetting
//
// Use this API command to update DP mesh tunnel setting.
//
// Request Body:
//	 - body *system.UpdateDpMeshTunnelSetting
func (s *WSGDPNetworkService) UpdatePlanesDpTunnelSetting(ctx context.Context, body *system.UpdateDpMeshTunnelSetting) (*common.EmptyResult, error) {
}
