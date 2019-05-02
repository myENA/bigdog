package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGDPNetworkService struct {
    client *Client
}

func NewWSGDPNetworkService (client *Client) *WSGDPNetworkService {
    s := new(WSGDPNetworkService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDPNetworkService () *WSGDPNetworkService {
    serv := new(WSGDPNetworkService)
    serv.client = ss.client
    return serv
}

func (s *WSGDPNetworkService) DeletePlanesStaticRouteByBladeUUID (ctx context.Context, bladeUUID string) error {
}

func (s *WSGDPNetworkService) FindPlanes (ctx context.Context) (system.DataPlaneList, error) {
}

func (s *WSGDPNetworkService) FindPlanesByBladeUUID (ctx context.Context, bladeUUID string) (system.DataPlaneConfiguration, error) {
}

func (s *WSGDPNetworkService) FindPlanesDpTunnelSetting (ctx context.Context) (system.GetDataPlaneMeshTunnelSetting, error) {
}

func (s *WSGDPNetworkService) PartialUpdatePlaneStatesByBladeUUID (ctx context.Context, bladeUUID string) error {
}

