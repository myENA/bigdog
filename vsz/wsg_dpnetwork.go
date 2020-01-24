package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
	return NewWSGDPNetworkService(ss.apiClient)
}

// DeletePlanesStaticRouteByBladeUUID
//
// Use this API command to delete static route.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) DeletePlanesStaticRouteByBladeUUID(ctx context.Context, bladeUUID string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeletePlanesStaticRouteByBladeUUID, true)
}

// FindPlanes
//
// Use this API command to retrieve a list of data planes.
func (s *WSGDPNetworkService) FindPlanes(ctx context.Context) (*WSGSystemDataPlaneList, error) {
	var (
		resp *WSGSystemDataPlaneList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindPlanes, true)
}

// FindPlanesByBladeUUID
//
// Use this API command to retrieve data plane by id.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) FindPlanesByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemDataPlaneConfiguration, error) {
	var (
		resp *WSGSystemDataPlaneConfiguration
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindPlanesByBladeUUID, true)
}

// FindPlanesDpTunnelSetting
//
// Use this API command to get DP mesh tunnel setting.
func (s *WSGDPNetworkService) FindPlanesDpTunnelSetting(ctx context.Context) (*WSGSystemGetDataPlaneMeshTunnelSetting, error) {
	var (
		resp *WSGSystemGetDataPlaneMeshTunnelSetting
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindPlanesDpTunnelSetting, true)
}

// PartialUpdatePlanesByBladeUUID
//
// Use this API command to modify the basic information of data plane.
//
// Request Body:
//	 - body *WSGSystemModifyDataPlane
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlane, bladeUUID string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdatePlanesByBladeUUID, true)
}

// PartialUpdatePlaneStatesByBladeUUID
//
// Use this API command to update DP state profile.
//
// Request Body:
//	 - body *WSGSystemModifyDataPlaneState
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlaneStatesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlaneState, bladeUUID string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdatePlaneStatesByBladeUUID, true)
}

// UpdatePlanesDpTunnelSetting
//
// Use this API command to update DP mesh tunnel setting.
//
// Request Body:
//	 - body *WSGSystemUpdateDpMeshTunnelSetting
func (s *WSGDPNetworkService) UpdatePlanesDpTunnelSetting(ctx context.Context, body *WSGSystemUpdateDpMeshTunnelSetting) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdatePlanesDpTunnelSetting, true)
}
