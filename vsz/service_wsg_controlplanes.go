package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGControlPlanesService struct {
    client *Client
}

func NewWSGControlPlanesService (client *Client) *WSGControlPlanesService {
    s := new(WSGControlPlanesService)
    s.client = client
    return s
}

func (ss *WSGService) WSGControlPlanesService () *WSGControlPlanesService {
    serv := new(WSGControlPlanesService)
    serv.client = ss.client
    return serv
}

func (s *WSGControlPlanesService) FindControlPlanes (ctx context.Context) (system.ControlPlaneList, error) {
}

func (s *WSGControlPlanesService) FindControlPlanesByBladeUUID (ctx context.Context, bladeUUID string) (system.ControlPlaneConfiguration, error) {
}

func (s *WSGControlPlanesService) FindControlPlanesStaticRoutesByBladeUUID (ctx context.Context, bladeUUID string) (system.StaticRouteList, error) {
}

func (s *WSGControlPlanesService) PartialUpdateControlPlanesIpSupport (ctx context.Context) error {
}

