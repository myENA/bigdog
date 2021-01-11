package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGControlPlanesService struct {
	apiClient *VSZClient
}

func NewWSGControlPlanesService(c *VSZClient) *WSGControlPlanesService {
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
// Operation ID: deleteControlPlanesStaticRoutesByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}/staticRoutes
// Success code: 204 (No Content)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) DeleteControlPlanesStaticRoutesByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteControlPlanesStaticRoutesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to delete the user defined interface of control plane.
//
// Operation ID: deleteControlPlanesUserDefinedInterfaceByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}/userDefinedInterface
// Success code: 204 (No Content)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) DeleteControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteControlPlanesUserDefinedInterfaceByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindControlPlanes
//
// Use this API command to retrieve the list of control plane.
//
// Operation ID: findControlPlanes
// Operation path: /controlPlanes
// Success code: 200 (OK)
func (s *WSGControlPlanesService) FindControlPlanes(ctx context.Context, mutators ...RequestMutator) (*WSGSystemControlPlaneListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemControlPlaneListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemControlPlaneListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindControlPlanes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemControlPlaneListAPIResponse), err
}

// FindControlPlanesByBladeUUID
//
// Use this API command to retrieve control plane.
//
// Operation ID: findControlPlanesByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}
// Success code: 200 (OK)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*WSGSystemControlPlaneConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemControlPlaneConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemControlPlaneConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindControlPlanesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemControlPlaneConfigurationAPIResponse), err
}

// FindControlPlanesInterfaces
//
// Use this API command to retrieve Control Plane Interface list.
//
// Operation ID: findControlPlanesInterfaces
// Operation path: /controlPlanes/interfaces
// Success code: 200 (OK)
//
// Optional parameters:
// - bladeUUID string
//		- nullable
func (s *WSGControlPlanesService) FindControlPlanesInterfaces(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSystemControlPlaneInterfaceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemControlPlaneInterfaceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemControlPlaneInterfaceListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindControlPlanesInterfaces, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["bladeUUID"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("bladeUUID", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemControlPlaneInterfaceListAPIResponse), err
}

// FindControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to retrieve static route of control plane.
//
// Operation ID: findControlPlanesStaticRoutesByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}/staticRoutes
// Success code: 200 (OK)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesStaticRoutesByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*WSGSystemStaticRouteListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemStaticRouteListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemStaticRouteListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindControlPlanesStaticRoutesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemStaticRouteListAPIResponse), err
}

// FindControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to retrieve user defined interface of control plane.
//
// Operation ID: findControlPlanesUserDefinedInterfaceByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}/userDefinedInterface
// Success code: 200 (OK)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) FindControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*WSGSystemUserDefinedInterfaceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemUserDefinedInterfaceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemUserDefinedInterfaceListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindControlPlanesUserDefinedInterfaceByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemUserDefinedInterfaceListAPIResponse), err
}

// PartialUpdateControlPlanesByBladeUUID
//
// Use this API command to modify the configuration of control plane.
//
// Operation ID: partialUpdateControlPlanesByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyControlPlane
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyControlPlane, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateControlPlanesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateControlPlanesIpSupport
//
// Use this API command to modify ip support of control plane.
//
// Operation ID: partialUpdateControlPlanesIpSupport
// Operation path: /controlPlanes/ipSupport
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyIpSupportType
func (s *WSGControlPlanesService) PartialUpdateControlPlanesIpSupport(ctx context.Context, body *WSGSystemModifyIpSupportType, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateControlPlanesIpSupport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to modify the static route of control plane.
//
// Operation ID: partialUpdateControlPlanesStaticRoutesByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}/staticRoutes
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyCPStaticRoute
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesStaticRoutesByBladeUUID(ctx context.Context, body *WSGSystemModifyCPStaticRoute, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateControlPlanesStaticRoutesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to modify user defined interface of control plane.
//
// Operation ID: partialUpdateControlPlanesUserDefinedInterfaceByBladeUUID
// Operation path: /controlPlanes/{bladeUUID}/userDefinedInterface
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyCPUserDefinedInterface
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, body *WSGSystemModifyCPUserDefinedInterface, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
