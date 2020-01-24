package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGControlPlanesService struct {
	apiClient *APIClient
}

func NewWSGControlPlanesService(c *APIClient) *WSGControlPlanesService {
	s := new(WSGControlPlanesService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGControlPlanesService() *WSGControlPlanesService {
	return NewWSGControlPlanesService(ss.apiClient)
}

// DeleteControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to delete the static route of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) DeleteControlPlanesStaticRoutesByBladeUUID(ctx context.Context, bladeUUID string) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteControlPlanesStaticRoutesByBladeUUID, true)
}

// DeleteControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to delete the user defined interface of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) DeleteControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, bladeUUID string) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteControlPlanesUserDefinedInterfaceByBladeUUID, true)
}

// FindControlPlanes
//
// Use this API command to retrieve the list of control plane.
func (s *WSGControlPlanesService) FindControlPlanes(ctx context.Context) (*WSGSystemControlPlaneList, error) {
	var (
		req  *APIRequest
		resp *WSGSystemControlPlaneList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanes, true)
}

// FindControlPlanesByBladeUUID
//
// Use this API command to retrieve control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemControlPlaneConfiguration, error) {
	var (
		req  *APIRequest
		resp *WSGSystemControlPlaneConfiguration
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanesByBladeUUID, true)
}

// FindControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to retrieve static route of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesStaticRoutesByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemStaticRouteList, error) {
	var (
		req  *APIRequest
		resp *WSGSystemStaticRouteList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanesStaticRoutesByBladeUUID, true)
}

// FindControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to retrieve user defined interface of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemUserDefinedInterfaceList, error) {
	var (
		req  *APIRequest
		resp *WSGSystemUserDefinedInterfaceList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanesUserDefinedInterfaceByBladeUUID, true)
}

// PartialUpdateControlPlanesByBladeUUID
//
// Use this API command to modify the basic information of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyControlPlane
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyControlPlane, bladeUUID string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateControlPlanesByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateControlPlanesIpSupport
//
// Use this API command to modify ip support of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyIpSupportType
func (s *WSGControlPlanesService) PartialUpdateControlPlanesIpSupport(ctx context.Context, body *WSGSystemModifyIpSupportType) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateControlPlanesIpSupport, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
