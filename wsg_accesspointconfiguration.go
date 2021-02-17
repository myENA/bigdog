package bigdog

// API Version: v9_1

import (
	"context"
	"io"
	"net/http"
	"time"
)

type WSGAccessPointConfigurationService struct {
	apiClient *VSZClient
}

func NewWSGAccessPointConfigurationService(c *VSZClient) *WSGAccessPointConfigurationService {
	s := new(WSGAccessPointConfigurationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointConfigurationService() *WSGAccessPointConfigurationService {
	return NewWSGAccessPointConfigurationService(ss.apiClient)
}

// AddAps
//
// Use this API command to create a new access point.
//
// Operation ID: addAps
// Operation path: /aps
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAPCreateAP
func (s *WSGAccessPointConfigurationService) AddAps(ctx context.Context, body *WSGAPCreateAP, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddAps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// AddApsPictureByApMac
//
// Use this API command to upload a new AP picture.
//
// Operation ID: addApsPictureByApMac
// Operation path: /aps/{apMac}/picture
// Success code: 204 (No Content)
//
// Form data parameters:
// - uploadFile io.Reader
//		- required
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) AddApsPictureByApMac(ctx context.Context, filename string, uploadFile io.Reader, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApsPictureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// AddSwapAps
//
// Use this API command to swap in specific AP
//
// Operation ID: addSwapAps
// Operation path: /swapAps
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAPSwapApConfigure
func (s *WSGAccessPointConfigurationService) AddSwapAps(ctx context.Context, body *WSGAPSwapApConfigure, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddSwapAps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsAltitudeByApMac
//
// Use this API command to disable AP level override of altitude. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsAltitudeByApMac
// Operation path: /aps/{apMac}/altitude
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsAltitudeByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsAltitudeByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsApMgmtVlanByApMac
//
// Disable AP Management Vlan Override of an AP.
//
// Operation ID: deleteApsApMgmtVlanByApMac
// Operation path: /aps/{apMac}/apMgmtVlan
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsApMgmtVlanByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsApMgmtVlanByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsAutoChannelSelection24ByApMac
//
// Use this API command to disable the AP level override of auto channel selection on the 2.4GHz radio. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsAutoChannelSelection24ByApMac
// Operation path: /aps/{apMac}/autoChannelSelection24
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsAutoChannelSelection24ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsAutoChannelSelection24ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsAutoChannelSelection50ByApMac
//
// Use this API command to disable the AP level override of auto channel selection on the 5GHz radio. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsAutoChannelSelection50ByApMac
// Operation path: /aps/{apMac}/autoChannelSelection50
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsAutoChannelSelection50ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsAutoChannelSelection50ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsBonjourGatewayByApMac
//
// Use this API command to disable AP level override of bonjour gateway. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsBonjourGatewayByApMac
// Operation path: /aps/{apMac}/bonjourGateway
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsBonjourGatewayByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsBonjourGatewayByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsByApMac
//
// Use this API command to delete an access point.
//
// Operation ID: deleteApsByApMac
// Operation path: /aps/{apMac}
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsChannelEvaluationIntervalByApMac
//
// Disable AP lChannel Evaluation Interval. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsChannelEvaluationIntervalByApMac
// Operation path: /aps/{apMac}/channelEvaluationInterval
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsChannelEvaluationIntervalByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsChannelEvaluationIntervalByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsClientAdmissionControl24ByApMac
//
// Use this API command to disable AP level override of client admission control 2.4GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsClientAdmissionControl24ByApMac
// Operation path: /aps/{apMac}/clientAdmissionControl24
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsClientAdmissionControl24ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsClientAdmissionControl24ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsClientAdmissionControl50ByApMac
//
// Use this API command to disable AP level override of client admission control 5GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsClientAdmissionControl50ByApMac
// Operation path: /aps/{apMac}/clientAdmissionControl50
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsClientAdmissionControl50ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsClientAdmissionControl50ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsDirectedMulticastFromNetworkEnabledByApMac
//
// Use this API command to disable Directed Multicast from network to wired/wireless client configuration override.The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsDirectedMulticastFromNetworkEnabledByApMac
// Operation path: /aps/{apMac}/directedMulticastFromNetworkEnabled
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromNetworkEnabledByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsDirectedMulticastFromNetworkEnabledByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsDirectedMulticastFromWiredClientEnabledByApMac
//
// Use this API command to disable Island SSID Broadcast enabled configuration override.The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsDirectedMulticastFromWiredClientEnabledByApMac
// Operation path: /aps/{apMac}/directedMulticastFromWiredClientEnabled
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromWiredClientEnabledByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsDirectedMulticastFromWiredClientEnabledByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsDirectedMulticastFromWirelessClientEnabledByApMac
//
// Use this API command to disable Island SSID Broadcast enabled configuration override.The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsDirectedMulticastFromWirelessClientEnabledByApMac
// Operation path: /aps/{apMac}/directedMulticastFromWirelessClientEnabled
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromWirelessClientEnabledByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsDirectedMulticastFromWirelessClientEnabledByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsGpsCoordinatesByApMac
//
// Disable AP Management GPS Cooordinates of an AP.
//
// Operation ID: deleteApsGpsCoordinatesByApMac
// Operation path: /aps/{apMac}/gpsCoordinates
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsGpsCoordinatesByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsGpsCoordinatesByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsLocationAdditionalInfoByApMac
//
// Use this API command to disable AP level override of location additionalInfo. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsLocationAdditionalInfoByApMac
// Operation path: /aps/{apMac}/locationAdditionalInfo
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLocationAdditionalInfoByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsLocationAdditionalInfoByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsLocationByApMac
//
// Use this API command to disable AP level override of location. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsLocationByApMac
// Operation path: /aps/{apMac}/location
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLocationByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsLocationByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsLoginByApMac
//
// Use this API command to disable the AP-level logon override. The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsLoginByApMac
// Operation path: /aps/{apMac}/login
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLoginByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsLoginByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsLteBandLockChannelsByApMac
//
// Use this API command to disable LTE band lock channel override. The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsLteBandLockChannelsByApMac
// Operation path: /aps/{apMac}/lteBandLockChannels
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLteBandLockChannelsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsLteBandLockChannelsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsMeshOptionsByApMac
//
// Use this API command to disable mesh options.
//
// Operation ID: deleteApsMeshOptionsByApMac
// Operation path: /aps/{apMac}/meshOptions
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsMeshOptionsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsMeshOptionsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsPictureByApMac
//
// Use this API command to delete an AP picture.
//
// Operation ID: deleteApsPictureByApMac
// Operation path: /aps/{apMac}/picture
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsPictureByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsPictureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsProtectionMode24ByApMac
//
// Use this API command to disable 2.4GHz radio protection mode configuration override.The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsProtectionMode24ByApMac
// Operation path: /aps/{apMac}/protectionMode24
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsProtectionMode24ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsProtectionMode24ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsRecoverySsidByApMac
//
// Use this API command to disable Recovery SSID configuration override.The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsRecoverySsidByApMac
// Operation path: /aps/{apMac}/recoverySsid
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRecoverySsidByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsRecoverySsidByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsRksGreForwardBroadcastByApMac
//
// Use this API command to disable Ruckus GRE Broadcast packet forwarding override. The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsRksGreForwardBroadcastByApMac
// Operation path: /aps/{apMac}/rksGreForwardBroadcast
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRksGreForwardBroadcastByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsRksGreForwardBroadcastByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsRogueApAggressivenessModeByApMac
//
// Use this API command to disable rogue AP aggressiveness mode override. The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsRogueApAggressivenessModeByApMac
// Operation path: /aps/{apMac}/rogueApAggressivenessMode
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRogueApAggressivenessModeByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsRogueApAggressivenessModeByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsRogueApJammingThresholdByApMac
//
// Use this API command to disable rogue AP jamming threshold override. The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsRogueApJammingThresholdByApMac
// Operation path: /aps/{apMac}/rogueApJammingThreshold
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRogueApJammingThresholdByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsRogueApJammingThresholdByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsRogueApReportThresholdByApMac
//
// Use this API command to disable rogue AP report threshold override. The AP will apply its group's or zone's configuration.
//
// Operation ID: deleteApsRogueApReportThresholdByApMac
// Operation path: /aps/{apMac}/rogueApReportThreshold
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRogueApReportThresholdByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsRogueApReportThresholdByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsSmartMonitorByApMac
//
// Use this API command to disable AP level override of smart monitor. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsSmartMonitorByApMac
// Operation path: /aps/{apMac}/smartMonitor
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsSmartMonitorByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsSmartMonitorByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsSpecificByApMac
//
// Use this API command to disable specific configuration override from AP group or zone.
//
// Operation ID: deleteApsSpecificByApMac
// Operation path: /aps/{apMac}/specific
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsSpecificByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsSpecificByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsSyslogByApMac
//
// Use this API command to disable the AP level syslog override. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsSyslogByApMac
// Operation path: /aps/{apMac}/syslog
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsSyslogByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsSyslogByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsVenueProfileByApMac
//
// Use this API command to disable AP level override of venue profile. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsVenueProfileByApMac
// Operation path: /aps/{apMac}/venueProfile
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsVenueProfileByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsVenueProfileByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi24ByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsWifi24ByApMac
// Operation path: /aps/{apMac}/wifi24
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi24ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi24ChannelByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio channel. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsWifi24ChannelByApMac
// Operation path: /aps/{apMac}/wifi24/channel
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi24ChannelByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi24ChannelRangeByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio channelRange. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsWifi24ChannelRangeByApMac
// Operation path: /aps/{apMac}/wifi24/channelRange
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelRangeByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi24ChannelRangeByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi24ChannelWidthByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio channelWidth. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsWifi24ChannelWidthByApMac
// Operation path: /aps/{apMac}/wifi24/channelWidth
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelWidthByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi24ChannelWidthByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi24TxPowerByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio txPower. The access point will take its group's configuration or zone's configuration.
//
// Operation ID: deleteApsWifi24TxPowerByApMac
// Operation path: /aps/{apMac}/wifi24/txPower
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24TxPowerByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi24TxPowerByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi50ByApMac
//
// Use this API command to disable the AP level override of 5GHz radio configuration. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsWifi50ByApMac
// Operation path: /aps/{apMac}/wifi50
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi50ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi50ChannelByApMac
//
// Use this API command to disable the AP level override of 5GHz radio channel. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsWifi50ChannelByApMac
// Operation path: /aps/{apMac}/wifi50/channel
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi50ChannelByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi50ChannelRangeByApMac
//
// Use this API command to disable the AP level override of 5GHz radio channelRange. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsWifi50ChannelRangeByApMac
// Operation path: /aps/{apMac}/wifi50/channelRange
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelRangeByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi50ChannelRangeByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi50ChannelWidthByApMac
//
// Use this API command to disable the AP level override of 5GHz radio channelWidth. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsWifi50ChannelWidthByApMac
// Operation path: /aps/{apMac}/wifi50/channelWidth
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelWidthByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi50ChannelWidthByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWifi50TxPowerByApMac
//
// Use this API command to disable the AP level override of 5GHz radio txPower. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsWifi50TxPowerByApMac
// Operation path: /aps/{apMac}/wifi50/txPower
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50TxPowerByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWifi50TxPowerByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWlanGroup24ByApMac
//
// Use this API command to disable the AP level override of WLAN group configuration on 2.4GHz radio. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsWlanGroup24ByApMac
// Operation path: /aps/{apMac}/wlanGroup24
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWlanGroup24ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWlanGroup24ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteApsWlanGroup50ByApMac
//
// Use this API command to disable the AP level override of WLAN group on the 5GHz radio. The access point will take its group's or zone's configuration.
//
// Operation ID: deleteApsWlanGroup50ByApMac
// Operation path: /aps/{apMac}/wlanGroup50
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWlanGroup50ByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteApsWlanGroup50ByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindApRadiosByApMac
//
// Fetch radio details about an access point by MAC address
//
// Operation ID: findApRadiosByApMac
// Operation path: /aps/{apMac}/radios
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) FindApRadiosByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAPRadioConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPRadioConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApRadiosByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPRadioConfigurationAPIResponse), err
}

// FindAps
//
// Use this API command to retrieve the list of APs that belong to a zone or a domain.
//
// Operation ID: findAps
// Operation path: /aps
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointConfigurationService) FindAps(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAPListEntryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPListEntryAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindAps, true)
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
	if v, ok := optionalParams["zoneId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("zoneId", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPListEntryAPIResponse), err
}

// FindApsByApMac
//
// Use this API command to retrieve the configuration of an AP.
//
// Operation ID: findApsByApMac
// Operation path: /aps/{apMac}
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) FindApsByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*WSGAPConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPConfigurationAPIResponse), err
}

// FindApsPictureByApMac
//
// Use this API command to retrieve the current AP picture.
//
// Operation ID: findApsPictureByApMac
// Operation path: /aps/{apMac}/picture
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) FindApsPictureByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsPictureByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}

// FindApsSupportLogByApMac
//
// Use this API command to download AP support log.
//
// Operation ID: findApsSupportLogByApMac
// Operation path: /aps/{apMac}/supportLog
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) FindApsSupportLogByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsSupportLogByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}

// FindMeshZeroTouch
//
// Use this API command to retrieve a list of unapproved AP.
//
// Operation ID: findMeshZeroTouch
// Operation path: /mesh/zeroTouch
// Success code: 200 (OK)
func (s *WSGAccessPointConfigurationService) FindMeshZeroTouch(ctx context.Context, mutators ...RequestMutator) (*WSGMeshNodeInfoListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGMeshNodeInfoListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindMeshZeroTouch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGMeshNodeInfoListAPIResponse), err
}

// PartialUpdateApsByApMac
//
// Use this API command to modify the configuration of an AP.
//
// Operation ID: partialUpdateApsByApMac
// Operation path: /aps/{apMac}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAPModifyAP
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) PartialUpdateApsByApMac(ctx context.Context, body *WSGAPModifyAP, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateApsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateApsByApMac
//
// Use this API command to modify the entire information of an AP.
//
// Operation ID: updateApsByApMac
// Operation path: /aps/{apMac}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAPModifyAP
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) UpdateApsByApMac(ctx context.Context, body *WSGAPModifyAP, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateApsByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateApsRebootByApMac
//
// reboot an access point.
//
// Operation ID: updateApsRebootByApMac
// Operation path: /aps/{apMac}/reboot
// Success code: 204 (No Content)
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) UpdateApsRebootByApMac(ctx context.Context, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateApsRebootByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateApsSpecificByApMac
//
// Use this API command to modify specific configuration.
//
// Operation ID: updateApsSpecificByApMac
// Operation path: /aps/{apMac}/specific
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAPModel
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) UpdateApsSpecificByApMac(ctx context.Context, body *WSGAPModel, apMac string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateApsSpecificByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateMeshZeroTouch
//
// Use this API command to approve/reject unapproved AP. Recommend to deploy 20 island APs to join per batch at the same time.
//
// Operation ID: updateMeshZeroTouch
// Operation path: /mesh/zeroTouch
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGMeshNodeInfoUpdateAPZeroTouch
func (s *WSGAccessPointConfigurationService) UpdateMeshZeroTouch(ctx context.Context, body *WSGMeshNodeInfoUpdateAPZeroTouch, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateMeshZeroTouch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}
