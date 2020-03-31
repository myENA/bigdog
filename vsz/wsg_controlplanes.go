package vsz

// API Version: v9_0

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
func (s *WSGControlPlanesService) DeleteControlPlanesStaticRoutesByBladeUUID(ctx context.Context, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteControlPlanesStaticRoutesByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to delete the user defined interface of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) DeleteControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteControlPlanesUserDefinedInterfaceByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindControlPlanes
//
// Use this API command to retrieve the list of control plane.
func (s *WSGControlPlanesService) FindControlPlanes(ctx context.Context) (*WSGSystemControlPlaneList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemControlPlaneList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanes, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemControlPlaneList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindControlPlanesByBladeUUID
//
// Use this API command to retrieve control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemControlPlaneConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemControlPlaneConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanesByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemControlPlaneConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindControlPlanesInterfaces
//
// Use this API command to retrieve Control Plane Interface list.
//
// Optional Parameters:
// - bladeUUID string
//		- nullable
func (s *WSGControlPlanesService) FindControlPlanesInterfaces(ctx context.Context, optionalParams map[string][]string) (*WSGSystemControlPlaneInterfaceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemControlPlaneInterfaceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanesInterfaces, true)
	if v, ok := optionalParams["bladeUUID"]; ok {
		req.AddQueryParameter("bladeUUID", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemControlPlaneInterfaceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to retrieve static route of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesStaticRoutesByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemStaticRouteList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemStaticRouteList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanesStaticRoutesByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemStaticRouteList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to retrieve user defined interface of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemUserDefinedInterfaceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemUserDefinedInterfaceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControlPlanesUserDefinedInterfaceByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemUserDefinedInterfaceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateControlPlanesByBladeUUID
//
// Use this API command to modify the configuration of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyControlPlane
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyControlPlane, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateControlPlanesByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateControlPlanesIpSupport
//
// Use this API command to modify ip support of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyIpSupportType
func (s *WSGControlPlanesService) PartialUpdateControlPlanesIpSupport(ctx context.Context, body *WSGSystemModifyIpSupportType) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateControlPlanesIpSupport, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to modify the static route of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyCPStaticRoute
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesStaticRoutesByBladeUUID(ctx context.Context, body *WSGSystemModifyCPStaticRoute, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateControlPlanesStaticRoutesByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to modify user defined interface of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyCPUserDefinedInterface
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, body *WSGSystemModifyCPUserDefinedInterface, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
