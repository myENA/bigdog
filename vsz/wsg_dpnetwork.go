package vsz

// API Version: v9_0

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
func (s *WSGDPNetworkService) DeletePlanesStaticRouteByBladeUUID(ctx context.Context, bladeUUID string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeletePlanesStaticRouteByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// FindPlanes
//
// Use this API command to retrieve a list of data planes.
func (s *WSGDPNetworkService) FindPlanes(ctx context.Context) (*WSGSystemDataPlaneList, *APIResponseMeta, error) {
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindPlanesByBladeUUID
//
// Use this API command to retrieve data plane by id.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGDPNetworkService) FindPlanesByBladeUUID(ctx context.Context, bladeUUID string) (*WSGSystemDataPlaneConfiguration, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindPlanesByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemDataPlaneConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindPlanesDpTunnelSetting
//
// Use this API command to get DP mesh tunnel setting.
//
// Request Body:
//	 - body []byte
func (s *WSGDPNetworkService) FindPlanesDpTunnelSetting(ctx context.Context, body []byte) (*WSGSystemGetDataPlaneMeshTunnelSetting, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindPlanesDpTunnelSetting, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemGetDataPlaneMeshTunnelSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGDPNetworkService) PartialUpdatePlanesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlane, bladeUUID string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdatePlanesByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGDPNetworkService) PartialUpdatePlaneStatesByBladeUUID(ctx context.Context, body *WSGSystemModifyDataPlaneState, bladeUUID string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdatePlaneStatesByBladeUUID, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// UpdatePlanesDpTunnelSetting
//
// Use this API command to update DP mesh tunnel setting.
//
// Request Body:
//	 - body *WSGSystemUpdateDpMeshTunnelSetting
func (s *WSGDPNetworkService) UpdatePlanesDpTunnelSetting(ctx context.Context, body *WSGSystemUpdateDpMeshTunnelSetting) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdatePlanesDpTunnelSetting, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
