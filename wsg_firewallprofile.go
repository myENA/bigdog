package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGFirewallProfileService struct {
	apiClient *VSZClient
}

func NewWSGFirewallProfileService(c *VSZClient) *WSGFirewallProfileService {
	s := new(WSGFirewallProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFirewallProfileService() *WSGFirewallProfileService {
	return NewWSGFirewallProfileService(ss.apiClient)
}

// AddFirewallProfiles
//
// Create a Firewall Profile.
//
// Operation ID: addFirewallProfiles
// Operation path: /firewallProfiles
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateFirewallProfile
func (s *WSGFirewallProfileService) AddFirewallProfiles(ctx context.Context, body *WSGProfileCreateFirewallProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddFirewallProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteFirewallProfiles
//
// Use this API command to delete Bulk Firewall Profiles.
//
// Operation ID: deleteFirewallProfiles
// Operation path: /firewallProfiles
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGFirewallProfileService) DeleteFirewallProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteFirewallProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteFirewallProfilesById
//
// Delete a Firewall Profile.
//
// Operation ID: deleteFirewallProfilesById
// Operation path: /firewallProfiles/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) DeleteFirewallProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteFirewallProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindFirewallProfiles
//
// Retrieve Firewall Profile list.
//
// Operation ID: findFirewallProfiles
// Operation path: /firewallProfiles
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGFirewallProfileService) FindFirewallProfiles(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileIdListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileIdListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindFirewallProfiles, true)
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
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileIdListAPIResponse), err
}

// FindFirewallProfilesById
//
// Retrieve a Firewall Profile.
//
// Operation ID: findFirewallProfilesById
// Operation path: /firewallProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileFirewallProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileFirewallProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindFirewallProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileFirewallProfileAPIResponse), err
}

// FindFirewallProfilesByQueryCriteria
//
// Retrieve a list of Firewall Profile.
//
// Operation ID: findFirewallProfilesByQueryCriteria
// Operation path: /firewallProfiles/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFirewallProfileService) FindFirewallProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileFirewallProfileArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileFirewallProfileArrayAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindFirewallProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileFirewallProfileArrayAPIResponse), err
}

// FindFirewallProfilesEthernetPortProfilesById
//
// Retrieve a EthernetPort Profile list of Firewall Profile is used by
//
// Operation ID: findFirewallProfilesEthernetPortProfilesById
// Operation path: /firewallProfiles/{id}/ethernetPortProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesEthernetPortProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGEthernetPortProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGEthernetPortProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindFirewallProfilesEthernetPortProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGEthernetPortProfileListAPIResponse), err
}

// FindFirewallProfilesWlansById
//
// Retrieve a WLAN list of Firewall Profile is used by
//
// Operation ID: findFirewallProfilesWlansById
// Operation path: /firewallProfiles/{id}/wlans
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesWlansById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGWLANQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGWLANQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindFirewallProfilesWlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGWLANQueryListAPIResponse), err
}

// UpdateFirewallProfilesById
//
// Modify a Firewall Profile.
//
// Operation ID: updateFirewallProfilesById
// Operation path: /firewallProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyFirewallProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) UpdateFirewallProfilesById(ctx context.Context, body *WSGProfileModifyFirewallProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateFirewallProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
