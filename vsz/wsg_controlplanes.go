package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteControlPlanesStaticRoutesByBladeUUID), true)
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteControlPlanesUserDefinedInterfaceByBladeUUID), true)
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindControlPlanes), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemControlPlaneList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindControlPlanesByBladeUUID), true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemControlPlaneConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindControlPlanesInterfaces), true)
	if v, ok := optionalParams["bladeUUID"]; ok {
		req.AddQueryParameter("bladeUUID", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemControlPlaneInterfaceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindControlPlanesStaticRoutesByBladeUUID), true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemStaticRouteList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindControlPlanesUserDefinedInterfaceByBladeUUID), true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemUserDefinedInterfaceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGControlPlanesService) PartialUpdateControlPlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyControlPlane, bladeUUID string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateControlPlanesByBladeUUID), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateControlPlanesIpSupport
//
// Use this API command to modify ip support of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyIpSupportType
func (s *WSGControlPlanesService) PartialUpdateControlPlanesIpSupport(ctx context.Context, body *WSGSystemModifyIpSupportType) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateControlPlanesIpSupport), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGControlPlanesService) PartialUpdateControlPlanesStaticRoutesByBladeUUID(ctx context.Context, body *WSGSystemModifyCPStaticRoute, bladeUUID string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateControlPlanesStaticRoutesByBladeUUID), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGControlPlanesService) PartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, body *WSGSystemModifyCPUserDefinedInterface, bladeUUID string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
