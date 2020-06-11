package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDHCPService struct {
	apiClient *VSZClient
}

func NewWSGDHCPService(c *VSZClient) *WSGDHCPService {
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
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGProfileCreateDhcpProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpProfileByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Request Body:
//	 - body *WSGCommonDoAssignIp
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *WSGCommonDoAssignIp, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRef, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpSiteConfigListRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonDhcpSiteConfigListRef()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindDhcpDataDhcpMsgStatsByApMac
//
// Use this API command to get AP DHCP Message Statistic.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpMsgStatsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPMessageStatsDhcpMsgStats, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPMessageStatsDhcpMsgStats
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpMsgStatsByApMac, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDHCPMessageStatsDhcpMsgStats()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDhcpDataDhcpPoolsByApMac
//
// Use this API command to get AP DHCP Pools Usage.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPPools, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPPools
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByApMac, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDHCPPools()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByPoolIndex(ctx context.Context, apMac string, poolIndex string, mutators ...RequestMutator) (*WSGDHCPPoolsDhcpPoolInfo, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPPoolsDhcpPoolInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByPoolIndex, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	req.SetPathParameter("poolIndex", poolIndex)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDHCPPoolsDhcpPoolInfo()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpProfileRef, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpProfileRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonDhcpProfileRef()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to get DHCP Pool list.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileDhcpProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileDhcpProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileDhcpProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Use this API command to get DHCP Configuration.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRef, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpSiteConfigListRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpSiteConfigByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonDhcpSiteConfigListRef()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Request Body:
//	 - body *WSGZoneQueryCriteria
func (s *WSGDHCPService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *WSGZoneQueryCriteria, mutators ...RequestMutator) (*WSGZoneDhcpSiteConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGZoneDhcpSiteConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesServicesDhcpSiteConfigByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGZoneDhcpSiteConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesDhcpProfileByQueryCriteria
//
// Query DHCP Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDHCPService) FindServicesDhcpProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileDhcpProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileDhcpProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesDhcpProfileByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileDhcpProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGDHCPService) PartialUpdateRkszonesDhcpSiteDhcpProfileById(ctx context.Context, body *WSGProfileCreateDhcpProfile, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesDhcpSiteDhcpProfileById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
