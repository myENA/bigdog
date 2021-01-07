package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addFirewallProfiles
//
// Create a Firewall Profile.
//
// Request Body:
//	 - body *WSGProfileCreateFirewallProfile
func (s *WSGFirewallProfileService) AddFirewallProfiles(ctx context.Context, body *WSGProfileCreateFirewallProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddFirewallProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteFirewallProfiles
//
// Operation ID: deleteFirewallProfiles
//
// Use this API command to delete Bulk Firewall Profiles.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGFirewallProfileService) DeleteFirewallProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteFirewallProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteFirewallProfilesById
//
// Operation ID: deleteFirewallProfilesById
//
// Delete a Firewall Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) DeleteFirewallProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteFirewallProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindFirewallProfiles
//
// Operation ID: findFirewallProfiles
//
// Retrieve Firewall Profile list.
//
// Optional Parameters:
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
		resp     APIResponse
		err      error

		respFn = newWSGProfileIdListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileIdListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindFirewallProfiles, true)
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
	return resp.(*WSGProfileIdListAPIResponse), err
}

// FindFirewallProfilesById
//
// Operation ID: findFirewallProfilesById
//
// Retrieve a Firewall Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileFirewallProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileFirewallProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileFirewallProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindFirewallProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileFirewallProfileAPIResponse), err
}

// FindFirewallProfilesByQueryCriteria
//
// Operation ID: findFirewallProfilesByQueryCriteria
//
// Retrieve a list of Firewall Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFirewallProfileService) FindFirewallProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileFirewallProfileArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileFirewallProfileArrayAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileFirewallProfileArrayAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindFirewallProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileFirewallProfileArrayAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileFirewallProfileArrayAPIResponse), err
}

// FindFirewallProfilesEthernetPortProfilesById
//
// Operation ID: findFirewallProfilesEthernetPortProfilesById
//
// Retrieve a EthernetPort Profile list of Firewall Profile is used by
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesEthernetPortProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGEthernetPortProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGEthernetPortProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGEthernetPortProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindFirewallProfilesEthernetPortProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGEthernetPortProfileListAPIResponse), err
}

// FindFirewallProfilesWlansById
//
// Operation ID: findFirewallProfilesWlansById
//
// Retrieve a WLAN list of Firewall Profile is used by
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesWlansById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGWLANQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindFirewallProfilesWlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANQueryListAPIResponse), err
}

// UpdateFirewallProfilesById
//
// Operation ID: updateFirewallProfilesById
//
// Modify a Firewall Profile.
//
// Request Body:
//	 - body *WSGProfileModifyFirewallProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) UpdateFirewallProfilesById(ctx context.Context, body *WSGProfileModifyFirewallProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateFirewallProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
