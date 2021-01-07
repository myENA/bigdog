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
// Operation ID: deleteControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to delete the static route of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) DeleteControlPlanesStaticRoutesByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteControlPlanesStaticRoutesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteControlPlanesUserDefinedInterfaceByBladeUUID
//
// Operation ID: deleteControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to delete the user defined interface of control plane.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) DeleteControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteControlPlanesUserDefinedInterfaceByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindControlPlanes
//
// Operation ID: findControlPlanes
//
// Use this API command to retrieve the list of control plane.
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindControlPlanes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemControlPlaneListAPIResponse), err
}

// FindControlPlanesByBladeUUID
//
// Operation ID: findControlPlanesByBladeUUID
//
// Use this API command to retrieve control plane.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindControlPlanesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemControlPlaneConfigurationAPIResponse), err
}

// FindControlPlanesInterfaces
//
// Operation ID: findControlPlanesInterfaces
//
// Use this API command to retrieve Control Plane Interface list.
//
// Optional Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindControlPlanesInterfaces, true)
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
// Operation ID: findControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to retrieve static route of control plane.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindControlPlanesStaticRoutesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemStaticRouteListAPIResponse), err
}

// FindControlPlanesUserDefinedInterfaceByBladeUUID
//
// Operation ID: findControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to retrieve user defined interface of control plane.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindControlPlanesUserDefinedInterfaceByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemUserDefinedInterfaceListAPIResponse), err
}

// PartialUpdateControlPlanesByBladeUUID
//
// Operation ID: partialUpdateControlPlanesByBladeUUID
//
// Use this API command to modify the configuration of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyControlPlane
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyControlPlane, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateControlPlanesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateControlPlanesIpSupport
//
// Operation ID: partialUpdateControlPlanesIpSupport
//
// Use this API command to modify ip support of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyIpSupportType
func (s *WSGControlPlanesService) PartialUpdateControlPlanesIpSupport(ctx context.Context, body *WSGSystemModifyIpSupportType, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateControlPlanesIpSupport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateControlPlanesStaticRoutesByBladeUUID
//
// Operation ID: partialUpdateControlPlanesStaticRoutesByBladeUUID
//
// Use this API command to modify the static route of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyCPStaticRoute
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesStaticRoutesByBladeUUID(ctx context.Context, body *WSGSystemModifyCPStaticRoute, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateControlPlanesStaticRoutesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID
//
// Operation ID: partialUpdateControlPlanesUserDefinedInterfaceByBladeUUID
//
// Use this API command to modify user defined interface of control plane.
//
// Request Body:
//	 - body *WSGSystemModifyCPUserDefinedInterface
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGControlPlanesService) PartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID(ctx context.Context, body *WSGSystemModifyCPUserDefinedInterface, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateControlPlanesUserDefinedInterfaceByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
