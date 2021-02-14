package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Operation ID: addRkszonesDhcpSiteDhcpProfileByZoneId
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpProfile
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGProfileCreateDhcpProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Operation ID: addRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpSiteConfig/doAssignIp
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonDoAssignIp
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *WSGCommonDoAssignIp, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRefAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonDhcpSiteConfigListRefAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonDhcpSiteConfigListRefAPIResponse), err
}

// DeleteRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to delete DHCP Pool by pool's ID.
//
// Operation ID: deleteRkszonesDhcpSiteDhcpProfileById
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpProfile/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to delete multiple DHCP Pools.
//
// Operation ID: deleteRkszonesDhcpSiteDhcpProfileByZoneId
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpProfile
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindDhcpDataDhcpMsgStatsByApMac
//
// Use this API command to get AP DHCP Message Statistic.
//
// Operation ID: findDhcpDataDhcpMsgStatsByApMac
// Operation path: /dhcpData/dhcpMsgStats/{apMac}
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpMsgStatsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPMessageStatsDhcpMsgStatsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDHCPMessageStatsDhcpMsgStatsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDhcpDataDhcpMsgStatsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDHCPMessageStatsDhcpMsgStatsAPIResponse), err
}

// FindDhcpDataDhcpPoolsByApMac
//
// Use this API command to get AP DHCP Pools Usage.
//
// Operation ID: findDhcpDataDhcpPoolsByApMac
// Operation path: /dhcpData/dhcpPools/{apMac}
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPPoolsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDHCPPoolsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDHCPPoolsAPIResponse), err
}

// FindDhcpDataDhcpPoolsByPoolIndex
//
// Use this API command to get AP DHCP Pool Usage by pool's index.
//
// Operation ID: findDhcpDataDhcpPoolsByPoolIndex
// Operation path: /dhcpData/dhcpPools/{apMac}/{poolIndex}
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
// - poolIndex string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByPoolIndex(ctx context.Context, apMac string, poolIndex string, mutators ...RequestMutator) (*WSGDHCPPoolsDhcpPoolInfoAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDHCPPoolsDhcpPoolInfoAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByPoolIndex, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	req.PathParams.Set("poolIndex", poolIndex)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDHCPPoolsDhcpPoolInfoAPIResponse), err
}

// FindRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to get DHCP Pool by pool's ID.
//
// Operation ID: findRkszonesDhcpSiteDhcpProfileById
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpProfile/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpProfileRefAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonDhcpProfileRefAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonDhcpProfileRefAPIResponse), err
}

// FindRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to get DHCP Pool list.
//
// Operation ID: findRkszonesDhcpSiteDhcpProfileByZoneId
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpProfile
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileDhcpProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileDhcpProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileDhcpProfileListAPIResponse), err
}

// FindRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Use this API command to get DHCP Configuration.
//
// Operation ID: findRkszonesDhcpSiteDhcpSiteConfigByZoneId
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpSiteConfig
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRefAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonDhcpSiteConfigListRefAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpSiteConfigByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonDhcpSiteConfigListRefAPIResponse), err
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Operation ID: findRkszonesServicesDhcpSiteConfigByQueryCriteria
// Operation path: /rkszones/services/dhcpSiteConfig/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGZoneQueryCriteria
func (s *WSGDHCPService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *WSGZoneQueryCriteria, mutators ...RequestMutator) (*WSGZoneDhcpSiteConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGZoneDhcpSiteConfigListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindRkszonesServicesDhcpSiteConfigByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneDhcpSiteConfigListAPIResponse), err
}

// FindServicesDhcpProfileByQueryCriteria
//
// Query DHCP Profiles with specified filters.
//
// Operation ID: findServicesDhcpProfileByQueryCriteria
// Operation path: /query/services/dhcpProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDHCPService) FindServicesDhcpProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileDhcpProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileDhcpProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesDhcpProfileByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileDhcpProfileListAPIResponse), err
}

// PartialUpdateRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to modify DHCP Pool by pool's ID.
//
// Operation ID: partialUpdateRkszonesDhcpSiteDhcpProfileById
// Operation path: /rkszones/{zoneId}/dhcpSite/dhcpProfile/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) PartialUpdateRkszonesDhcpSiteDhcpProfileById(ctx context.Context, body *WSGProfileCreateDhcpProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
