package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGDPNetworkService struct {
	apiClient *VSZClient
}

func NewWSGDPNetworkService(c *VSZClient) *WSGDPNetworkService {
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
// Operation ID: deletePlanesStaticRouteByBladeUUID
// Operation path: /planes/{bladeUUID}/staticRoute
// Success code: 204 (No Content)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) DeletePlanesStaticRouteByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeletePlanesStaticRouteByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindPlanes
//
// Use this API command to retrieve a list of data planes.
//
// Operation ID: findPlanes
// Operation path: /planes
// Success code: 200 (OK)
func (s *WSGDPNetworkService) FindPlanes(ctx context.Context, mutators ...RequestMutator) (*WSGSystemDataPlaneListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemDataPlaneListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindPlanes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemDataPlaneListAPIResponse), err
}

// FindPlanesByBladeUUID
//
// Use this API command to retrieve data plane by id.
//
// Operation ID: findPlanesByBladeUUID
// Operation path: /planes/{bladeUUID}
// Success code: 200 (OK)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) FindPlanesByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*WSGSystemDataPlaneConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemDataPlaneConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindPlanesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemDataPlaneConfigurationAPIResponse), err
}

// FindPlanesDpTunnelSetting
//
// Use this API command to get DP mesh tunnel setting.
//
// Operation ID: findPlanesDpTunnelSetting
// Operation path: /planes/dpTunnel/setting
// Success code: 200 (OK)
//
// Optional parameters:
// - useless string
//		- nullable
func (s *WSGDPNetworkService) FindPlanesDpTunnelSetting(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSystemGetDataPlaneMeshTunnelSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemGetDataPlaneMeshTunnelSettingAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindPlanesDpTunnelSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["useless"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("useless", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemGetDataPlaneMeshTunnelSettingAPIResponse), err
}

// PartialUpdatePlanesByBladeUUID
//
// Use this API command to modify the configuration of data plane.
//
// Operation ID: partialUpdatePlanesByBladeUUID
// Operation path: /planes/{bladeUUID}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyDataPlane
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlane, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdatePlanesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdatePlaneStatesByBladeUUID
//
// Use this API command to update DP state profile.
//
// Operation ID: partialUpdatePlaneStatesByBladeUUID
// Operation path: /planeStates/{bladeUUID}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifyDataPlaneState
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlaneStatesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlaneState, bladeUUID string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdatePlaneStatesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdatePlanesDpTunnelSetting
//
// Use this API command to update DP mesh tunnel setting.
//
// Operation ID: updatePlanesDpTunnelSetting
// Operation path: /planes/dpTunnel/setting
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemUpdateDpMeshTunnelSetting
func (s *WSGDPNetworkService) UpdatePlanesDpTunnelSetting(ctx context.Context, body *WSGSystemUpdateDpMeshTunnelSetting, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdatePlanesDpTunnelSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
