package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGFlexiVPNService struct {
	apiClient *VSZClient
}

func NewWSGFlexiVPNService(c *VSZClient) *WSGFlexiVPNService {
	s := new(WSGFlexiVPNService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFlexiVPNService() *WSGFlexiVPNService {
	return NewWSGFlexiVPNService(ss.apiClient)
}

// WSGFlexiVPNSetting
//
// Definition: flexiVpn_flexiVpnSetting
type WSGFlexiVPNSetting struct {
	// ZoneAffinityId
	// Zone Affinity ID
	// Constraints:
	//    - required
	ZoneAffinityId *string `json:"zoneAffinityId"`
}

func NewWSGFlexiVPNSetting() *WSGFlexiVPNSetting {
	m := new(WSGFlexiVPNSetting)
	return m
}

// DeleteRkszonesWlansFlexiVpnProfileById
//
// Use this API command to delete Flexi-VPN on WLAN
//
// Operation ID: deleteRkszonesWlansFlexiVpnProfileById
// Operation path: /rkszones/{zoneId}/wlans/{id}/flexiVpnProfile
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGFlexiVPNService) DeleteRkszonesWlansFlexiVpnProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesWlansFlexiVpnProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindServicesFlexiVpnProfileByQueryCriteria
//
// Use this API command to query Flexi-VPN profiles.
//
// Operation ID: findServicesFlexiVpnProfileByQueryCriteria
// Operation path: /query/services/flexiVpnProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFlexiVPNService) FindServicesFlexiVpnProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileFlexiVpnProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileFlexiVpnProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesFlexiVpnProfileByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileFlexiVpnProfileListAPIResponse), err
}
