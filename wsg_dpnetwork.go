package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDpnetworkService struct {
	apiClient *VSZClient
}

func NewWSGDpnetworkService(c *VSZClient) *WSGDpnetworkService {
	s := new(WSGDpnetworkService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDpnetworkService() *WSGDpnetworkService {
	return NewWSGDpnetworkService(ss.apiClient)
}

// DeletePlanesStaticRouteByBladeUUID
//
// Use this API command to delete static route.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDpnetworkService) DeletePlanesStaticRouteByBladeUUID(ctx context.Context, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeletePlanesStaticRouteByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindPlanes
//
// Use this API command to retrieve a list of data planes.
func (s *WSGDpnetworkService) FindPlanes(ctx context.Context) (*WSGSystemDataPlaneList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemDataPlaneList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindPlanes, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemDataPlaneList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindPlanesByBladeUUID
//
// Use this API command to retrieve data plane by id.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDpnetworkService) FindPlanesByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemDataPlaneConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemDataPlaneConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindPlanesByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemDataPlaneConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindPlanesDpTunnelSetting
//
// Use this API command to get DP mesh tunnel setting.
//
// Request Body:
//	 - body []byte
func (s *WSGDpnetworkService) FindPlanesDpTunnelSetting(ctx context.Context, body []byte) (*WSGSystemGetDataPlaneMeshTunnelSetting, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemGetDataPlaneMeshTunnelSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindPlanesDpTunnelSetting, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemGetDataPlaneMeshTunnelSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdatePlanesByBladeUUID
//
// Use this API command to modify the configuration of data plane.
//
// Request Body:
//	 - body *WSGSystemModifyDataPlane
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDpnetworkService) PartialUpdatePlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlane, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdatePlanesByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGDpnetworkService) PartialUpdatePlaneStatesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlaneState, bladeUUID string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdatePlaneStatesByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdatePlanesDpTunnelSetting
//
// Use this API command to update DP mesh tunnel setting.
//
// Request Body:
//	 - body *WSGSystemUpdateDpMeshTunnelSetting
func (s *WSGDpnetworkService) UpdatePlanesDpTunnelSetting(ctx context.Context, body *WSGSystemUpdateDpMeshTunnelSetting) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdatePlanesDpTunnelSetting, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
