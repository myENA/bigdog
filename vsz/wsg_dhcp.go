package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGDHCPService struct {
	apiClient *APIClient
}

func NewWSGDHCPService(c *APIClient) *WSGDHCPService {
	s := new(WSGDHCPService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDHCPService() *WSGDHCPService {
	return NewWSGDHCPService(ss.apiClient)
}

// AddRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to create DHCP Pool.
//
// Request Body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGProfileCreateDhcpProfile, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpProfileByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to delete DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to delete multiple DHCP Pools.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindDhcpDataDhcpMsgStatsByApMac
//
// Use this API command to get AP DHCP Message Statistic.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpMsgStatsByApMac(ctx context.Context, apMac string) (*WSGDHCPMsgStats, error) {
	var (
		req      *APIRequest
		resp     *WSGDHCPMsgStats
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpMsgStatsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDHCPMsgStats()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindDhcpDataDhcpPoolsByApMac
//
// Use this API command to get AP DHCP Pools Usage.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByApMac(ctx context.Context, apMac string) (*WSGDHCPPools, error) {
	var (
		req      *APIRequest
		resp     *WSGDHCPPools
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDHCPPools()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindDhcpDataDhcpPoolsByPoolIndex
//
// Use this API command to get AP DHCP Pool Usage by pool's index.
//
// Required Parameters:
// - apMac string
//		- required
// - poolIndex string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByPoolIndex(ctx context.Context, apMac string, poolIndex string) (*WSGDHCPPoolsDhcpPoolInfo, error) {
	var (
		req      *APIRequest
		resp     *WSGDHCPPoolsDhcpPoolInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, poolIndex, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByPoolIndex, true)
	req.SetPathParameter("apMac", apMac)
	req.SetPathParameter("poolIndex", poolIndex)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDHCPPoolsDhcpPoolInfo()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to get DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string) (*WSGCommonDhcpProfileRef, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonDhcpProfileRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonDhcpProfileRef()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to get DHCP Pool list.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, zoneId string) (*WSGProfileDhcpProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileDhcpProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileDhcpProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Use this API command to get DHCP Configuration.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId(ctx context.Context, zoneId string) (*WSGCommonDhcpSiteConfigListRef, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonDhcpSiteConfigListRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpSiteConfigByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonDhcpSiteConfigListRef()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to modify DHCP Pool by pool's ID.
//
// Request Body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) PartialUpdateRkszonesDhcpSiteDhcpProfileById(ctx context.Context, body *WSGProfileCreateDhcpProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesDhcpSiteDhcpProfileById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
