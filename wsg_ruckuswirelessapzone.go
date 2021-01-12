package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGRuckusWirelessAPZoneService struct {
	apiClient *VSZClient
}

func NewWSGRuckusWirelessAPZoneService(c *VSZClient) *WSGRuckusWirelessAPZoneService {
	s := new(WSGRuckusWirelessAPZoneService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRuckusWirelessAPZoneService() *WSGRuckusWirelessAPZoneService {
	return NewWSGRuckusWirelessAPZoneService(ss.apiClient)
}

// AddRkszones
//
// Use this API command to create a new Ruckus Wireless AP zone.
//
// Operation ID: addRkszones
// Operation path: /rkszones
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszones(ctx context.Context, body *WSGZoneCreateZone, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszones, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesDual
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv4/IPv6.
//
// Operation ID: addRkszonesDual
// Operation path: /rkszones/dual
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDual(ctx context.Context, body *WSGZoneCreateZone, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesDual, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesIpv6
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv6.
//
// Operation ID: addRkszonesIpv6
// Operation path: /rkszones/ipv6
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesIpv6(ctx context.Context, body *WSGZoneCreateZone, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesIpv6, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesAltitudeById
//
// Use this API command to disable altitude configuration of zone.
//
// Operation ID: deleteRkszonesAltitudeById
// Operation path: /rkszones/{id}/altitude
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesAltitudeById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAltitudeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesBackgroundScanning24ById
//
// Use this API command to disable background scanning 2.4GHz radio configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesBackgroundScanning24ById
// Operation path: /rkszones/{id}/backgroundScanning24
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning24ById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesBackgroundScanning24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesBackgroundScanning50ById
//
// Use this API command to disable background scanning 5GHz radio configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesBackgroundScanning50ById
// Operation path: /rkszones/{id}/backgroundScanning50
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning50ById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesBackgroundScanning50ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesById
//
// Use this API command to delete a zone.
//
// Operation ID: deleteRkszonesById
// Operation path: /rkszones/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesClientAdmissionControl24ById
//
// Use this API command to disable client admission control 2.4GHz radio configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesClientAdmissionControl24ById
// Operation path: /rkszones/{id}/clientAdmissionControl24
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl24ById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesClientAdmissionControl24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesClientAdmissionControl50ById
//
// Use this API command to disable client admission control 5GHz radio configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesClientAdmissionControl50ById
// Operation path: /rkszones/{id}/clientAdmissionControl50
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl50ById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesClientAdmissionControl50ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesIpsecProfilesById
//
// Use this API command to Delete IPsec profiles.
//
// Operation ID: deleteRkszonesIpsecProfilesById
// Operation path: /rkszones/{id}/ipsecProfiles
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesIpsecProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesIpsecProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesLoadBalancingBandBalancingById
//
// Use this API command to disable band balancing for APs that belong to a zone.
//
// Operation ID: deleteRkszonesLoadBalancingBandBalancingById
// Operation path: /rkszones/{id}/loadBalancing/bandBalancing
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLoadBalancingBandBalancingById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesLoadBalancingBandBalancingById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesLoadBalancingById
//
// Use this API command to disable overall load balancing configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesLoadBalancingById
// Operation path: /rkszones/{id}/loadBalancing
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLoadBalancingById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesLoadBalancingById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesLoadBalancingClientLoadBalancing24ById
//
// Use this API command to disable client load balancing 2.4GHz radio configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesLoadBalancingClientLoadBalancing24ById
// Operation path: /rkszones/{id}/loadBalancing/clientLoadBalancing24
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLoadBalancingClientLoadBalancing24ById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesLoadBalancingClientLoadBalancing24ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesLoadBalancingClientLoadBalancing50ById
//
// Use this API command to disable client load balancing 5GHz radio configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesLoadBalancingClientLoadBalancing50ById
// Operation path: /rkszones/{id}/loadBalancing/clientLoadBalancing50
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLoadBalancingClientLoadBalancing50ById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesLoadBalancingClientLoadBalancing50ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesLocationBasedServiceById
//
// Use this API command to disable location based service for APs that belong to a zone.
//
// Operation ID: deleteRkszonesLocationBasedServiceById
// Operation path: /rkszones/{id}/locationBasedService
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLocationBasedServiceById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesLocationBasedServiceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesMeshById
//
// Use this API command to disable mesh networking.
//
// Operation ID: deleteRkszonesMeshById
// Operation path: /rkszones/{id}/mesh
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesMeshById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesMeshById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesRecoverySsidById
//
// Use this API command to clear recovery ssid setting of a zone.
//
// Operation ID: deleteRkszonesRecoverySsidById
// Operation path: /rkszones/{id}/recoverySsid
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRecoverySsidById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesRecoverySsidById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesRogueById
//
// Use this API command to disable rogue AP detection for APs that belong to a zone.
//
// Operation ID: deleteRkszonesRogueById
// Operation path: /rkszones/{id}/rogue
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRogueById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesRogueById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesSmartMonitorById
//
// Use this API command to disable smart monitor for APs that belong to a zone.
//
// Operation ID: deleteRkszonesSmartMonitorById
// Operation path: /rkszones/{id}/smartMonitor
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSmartMonitorById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesSmartMonitorById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesSnmpAgentById
//
// Use this API command to clear SNMPv2 and SNMPv3 agent that belong to a zone.
//
// Operation ID: deleteRkszonesSnmpAgentById
// Operation path: /rkszones/{id}/snmpAgent
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSnmpAgentById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesSnmpAgentById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesSoftGreTunnelProfliesById
//
// Use this API command to Delete IPsec profiles.
//
// Operation ID: deleteRkszonesSoftGreTunnelProfliesById
// Operation path: /rkszones/{id}/softGreTunnelProflies
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSoftGreTunnelProfliesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesSoftGreTunnelProfliesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesSyslogById
//
// Use this API command to disable syslog configuration for APs that belong to a zone.
//
// Operation ID: deleteRkszonesSyslogById
// Operation path: /rkszones/{id}/syslog
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSyslogById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesSyslogById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesVenueProfileById
//
// Use this API command to clear Hotspot 2.0 venue profile for APs that belong to a zone.
//
// Operation ID: deleteRkszonesVenueProfileById
// Operation path: /rkszones/{id}/venueProfile
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesVenueProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesVenueProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszones
//
// Use this API command to retrieve the list of Ruckus Wireless AP zones that belong to a domain.
//
// Operation ID: findRkszones
// Operation path: /rkszones
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGRuckusWirelessAPZoneService) FindRkszones(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGZoneListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszones, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneListAPIResponse), err
}

// FindRkszonesApFirmwareByZoneId
//
// Use this API command to retrieve AP Firmware the list that belong to a zone.
//
// Operation ID: findRkszonesApFirmwareByZoneId
// Operation path: /rkszones/{zoneId}/apFirmware
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApFirmwareByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGZoneApFirmwareListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneApFirmwareListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesApFirmwareByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneApFirmwareListAPIResponse), err
}

// FindRkszonesApmodelByModel
//
// Use this API command to retrieve AP model specific configuration that belong to a zone.
//
// Operation ID: findRkszonesApmodelByModel
// Operation path: /rkszones/{zoneId}/apmodel/{model}
// Success code: 200 (OK)
//
// Required parameters:
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelByModel(ctx context.Context, model string, zoneId string, mutators ...RequestMutator) (*WSGZoneAPModelApModelAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneAPModelApModelAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesApmodelByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("model", model)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneAPModelApModelAPIResponse), err
}

// FindRkszonesApmodelCommonAttributeByModel
//
// Use this API command to retrieve AP model common attribute that belong to a zone.
//
// Operation ID: findRkszonesApmodelCommonAttributeByModel
// Operation path: /rkszones/{zoneId}/apmodel/{model}/commonAttribute
// Success code: 200 (OK)
//
// Required parameters:
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelCommonAttributeByModel(ctx context.Context, model string, zoneId string, mutators ...RequestMutator) (*WSGAPModelCommonAttributeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPModelCommonAttributeAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesApmodelCommonAttributeByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("model", model)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPModelCommonAttributeAPIResponse), err
}

// FindRkszonesAvailableIpsecProfilesByZoneId
//
// Get available IPSec tunnel profiles of this Zone.
//
// Operation ID: findRkszonesAvailableIpsecProfilesByZoneId
// Operation path: /rkszones/{zoneId}/availableIpsecProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableIpsecProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGZoneAvailableTunnelProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneAvailableTunnelProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAvailableIpsecProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneAvailableTunnelProfileListAPIResponse), err
}

// FindRkszonesAvailableTunnelProfilesByZoneId
//
// Get available GRE tunnel profiles of this Zone.
//
// Operation ID: findRkszonesAvailableTunnelProfilesByZoneId
// Operation path: /rkszones/{zoneId}/availableTunnelProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableTunnelProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGZoneAvailableTunnelProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneAvailableTunnelProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAvailableTunnelProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneAvailableTunnelProfileListAPIResponse), err
}

// FindRkszonesById
//
// Use this API command to retrieve Ruckus Wireless AP zones configuration.
//
// Operation ID: findRkszonesById
// Operation path: /rkszones/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGZoneConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneConfigurationAPIResponse), err
}

// FindRkszonesMeshById
//
// Use this API command to retrieve the mesh configuration of a zone.
//
// Operation ID: findRkszonesMeshById
// Operation path: /rkszones/{id}/mesh
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesMeshById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGZoneMeshConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGZoneMeshConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesMeshById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGZoneMeshConfigurationAPIResponse), err
}

// PartialUpdateRkszonesById
//
// Use this API command to modify the configuration of a zone.
//
// Operation ID: partialUpdateRkszonesById
// Operation path: /rkszones/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGZoneModifyZone
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) PartialUpdateRkszonesById(ctx context.Context, body *WSGZoneModifyZone, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesApFirmwareByZoneId
//
// Use this API command to change the AP Firmware that belong to a zone.
//
// Operation ID: updateRkszonesApFirmwareByZoneId
// Operation path: /rkszones/{zoneId}/apFirmware
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGZoneModfiyApFirmware
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesApFirmwareByZoneId(ctx context.Context, body *WSGZoneModfiyApFirmware, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesApFirmwareByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesApmodelByModel
//
// Use this API command to modify the AP model specific configuration that belong to a zone.
//
// Operation ID: updateRkszonesApmodelByModel
// Operation path: /rkszones/{zoneId}/apmodel/{model}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGZoneAPModelApModel
//
// Required parameters:
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesApmodelByModel(ctx context.Context, body *WSGZoneAPModelApModel, model string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesApmodelByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("model", model)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesById
//
// Use this API command to modify entire information of a zone.
//
// Operation ID: updateRkszonesById
// Operation path: /rkszones/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGZoneModifyZone
//
// Required parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesById(ctx context.Context, body *WSGZoneModifyZone, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
