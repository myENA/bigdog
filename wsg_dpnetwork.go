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
// Operation ID: deletePlanesStaticRouteByBladeUUID
//
// Use this API command to delete static route.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) DeletePlanesStaticRouteByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeletePlanesStaticRouteByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindPlanes
//
// Operation ID: findPlanes
//
// Use this API command to retrieve a list of data planes.
func (s *WSGDPNetworkService) FindPlanes(ctx context.Context, mutators ...RequestMutator) (*WSGSystemDataPlaneListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemDataPlaneListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemDataPlaneListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindPlanes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemDataPlaneListAPIResponse), err
}

// FindPlanesByBladeUUID
//
// Operation ID: findPlanesByBladeUUID
//
// Use this API command to retrieve data plane by id.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemDataPlaneConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindPlanesByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemDataPlaneConfigurationAPIResponse), err
}

// FindPlanesDpTunnelSetting
//
// Operation ID: findPlanesDpTunnelSetting
//
// Use this API command to get DP mesh tunnel setting.
//
// Optional Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemGetDataPlaneMeshTunnelSettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindPlanesDpTunnelSetting, true)
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
// Operation ID: partialUpdatePlanesByBladeUUID
//
// Use this API command to modify the configuration of data plane.
//
// Request Body:
//	 - body *WSGSystemModifyDataPlane
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlane, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdatePlanesByBladeUUID, true)
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

// PartialUpdatePlaneStatesByBladeUUID
//
// Operation ID: partialUpdatePlaneStatesByBladeUUID
//
// Use this API command to update DP state profile.
//
// Request Body:
//	 - body *WSGSystemModifyDataPlaneState
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) PartialUpdatePlaneStatesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlaneState, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdatePlaneStatesByBladeUUID, true)
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

// UpdatePlanesDpTunnelSetting
//
// Operation ID: updatePlanesDpTunnelSetting
//
// Use this API command to update DP mesh tunnel setting.
//
// Request Body:
//	 - body *WSGSystemUpdateDpMeshTunnelSetting
func (s *WSGDPNetworkService) UpdatePlanesDpTunnelSetting(ctx context.Context, body *WSGSystemUpdateDpMeshTunnelSetting, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdatePlanesDpTunnelSetting, true)
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
