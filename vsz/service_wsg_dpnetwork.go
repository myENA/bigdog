package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGDPNetworkService struct {
	client *Client
}

func NewWSGDPNetworkService(client *Client) *WSGDPNetworkService {
	s := new(WSGDPNetworkService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDPNetworkService() *WSGDPNetworkService {
	serv := new(WSGDPNetworkService)
	serv.client = ss.client
	return serv
}

// DeletePlanesStaticRouteByBladeUUID
//
// Use this API command to delete static route.
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
func (s *WSGDPNetworkService) FindPlanesByBladeUUID(ctx context.Context, pBladeUUID string) (*system.DataPlaneConfiguration, error) {
}

// FindPlanesDpTunnelSetting
//
// Use this API command to get DP mesh tunnel setting.
func (s *WSGDPNetworkService) FindPlanesDpTunnelSetting(ctx context.Context) (*system.GetDataPlaneMeshTunnelSetting, error) {
}

// PartialUpdatePlaneStatesByBladeUUID
//
// Use this API command to update DP state profile.
//
// Request Body:
//	 - body *system.ModifyDataPlaneState
func (s *WSGDPNetworkService) PartialUpdatePlaneStatesByBladeUUID(ctx context.Context, body *system.ModifyDataPlaneState, pBladeUUID string) (*common.EmptyResult, error) {
}
