package bigdog

// API Version: v9_1

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
// Operation ID: addRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to create DHCP Pool.
//
// Request Body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGProfileCreateDhcpProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Operation ID: addRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Request Body:
//	 - body *WSGCommonDoAssignIp
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *WSGCommonDoAssignIp, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRefAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonDhcpSiteConfigListRefAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonDhcpSiteConfigListRefAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonDhcpSiteConfigListRefAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonDhcpSiteConfigListRefAPIResponse), err
}

// DeleteRkszonesDhcpSiteDhcpProfileById
//
// Operation ID: deleteRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to delete DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteRkszonesDhcpSiteDhcpProfileByZoneId
//
// Operation ID: deleteRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to delete multiple DHCP Pools.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) DeleteRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindDhcpDataDhcpMsgStatsByApMac
//
// Operation ID: findDhcpDataDhcpMsgStatsByApMac
//
// Use this API command to get AP DHCP Message Statistic.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpMsgStatsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPMessageStatsDhcpMsgStatsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDHCPMessageStatsDhcpMsgStatsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDHCPMessageStatsDhcpMsgStatsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDhcpDataDhcpMsgStatsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDHCPMessageStatsDhcpMsgStatsAPIResponse), err
}

// FindDhcpDataDhcpPoolsByApMac
//
// Operation ID: findDhcpDataDhcpPoolsByApMac
//
// Use this API command to get AP DHCP Pools Usage.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGDHCPPoolsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDHCPPoolsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDHCPPoolsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDHCPPoolsAPIResponse), err
}

// FindDhcpDataDhcpPoolsByPoolIndex
//
// Operation ID: findDhcpDataDhcpPoolsByPoolIndex
//
// Use this API command to get AP DHCP Pool Usage by pool's index.
//
// Required Parameters:
// - apMac string
//		- required
// - poolIndex string
//		- required
func (s *WSGDHCPService) FindDhcpDataDhcpPoolsByPoolIndex(ctx context.Context, apMac string, poolIndex string, mutators ...RequestMutator) (*WSGDHCPPoolsDhcpPoolInfoAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDHCPPoolsDhcpPoolInfoAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDHCPPoolsDhcpPoolInfoAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByPoolIndex, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	req.PathParams.Set("poolIndex", poolIndex)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDHCPPoolsDhcpPoolInfoAPIResponse), err
}

// FindRkszonesDhcpSiteDhcpProfileById
//
// Operation ID: findRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to get DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpProfileRefAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonDhcpProfileRefAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonDhcpProfileRefAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonDhcpProfileRefAPIResponse), err
}

// FindRkszonesDhcpSiteDhcpProfileByZoneId
//
// Operation ID: findRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to get DHCP Pool list.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileDhcpProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileDhcpProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileDhcpProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileDhcpProfileListAPIResponse), err
}

// FindRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Operation ID: findRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Use this API command to get DHCP Configuration.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDHCPService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGCommonDhcpSiteConfigListRefAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonDhcpSiteConfigListRefAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonDhcpSiteConfigListRefAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpSiteConfigByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonDhcpSiteConfigListRefAPIResponse), err
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Operation ID: findRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Request Body:
//	 - body *WSGZoneQueryCriteria
func (s *WSGDHCPService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *WSGZoneQueryCriteria, mutators ...RequestMutator) (*WSGZoneDhcpSiteConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneDhcpSiteConfigListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGZoneDhcpSiteConfigListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRkszonesServicesDhcpSiteConfigByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGZoneDhcpSiteConfigListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneDhcpSiteConfigListAPIResponse), err
}

// FindServicesDhcpProfileByQueryCriteria
//
// Operation ID: findServicesDhcpProfileByQueryCriteria
//
// Query DHCP Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDHCPService) FindServicesDhcpProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileDhcpProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileDhcpProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileDhcpProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesDhcpProfileByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileDhcpProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileDhcpProfileListAPIResponse), err
}

// PartialUpdateRkszonesDhcpSiteDhcpProfileById
//
// Operation ID: partialUpdateRkszonesDhcpSiteDhcpProfileById
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
func (s *WSGDHCPService) PartialUpdateRkszonesDhcpSiteDhcpProfileById(ctx context.Context, body *WSGProfileCreateDhcpProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesDhcpSiteDhcpProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
