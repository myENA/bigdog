package bigdog

// API Version: v9_0

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
// Create a Firewall Profile.
//
// Request Body:
//	 - body *WSGProfileCreateFirewallProfile
func (s *WSGFirewallProfileService) AddFirewallProfiles(ctx context.Context, body *WSGProfileCreateFirewallProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddFirewallProfiles, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteFirewallProfiles
//
// Use this API command to delete Bulk Firewall Profiles.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGFirewallProfileService) DeleteFirewallProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFirewallProfiles, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteFirewallProfilesById
//
// Delete a Firewall Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) DeleteFirewallProfilesById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFirewallProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindFirewallProfiles
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
func (s *WSGFirewallProfileService) FindFirewallProfiles(ctx context.Context, optionalParams map[string][]string) (*WSGProfileIdList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileIdList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfiles, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIdList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFirewallProfilesById
//
// Retrieve a Firewall Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesById(ctx context.Context, id string) (*WSGProfileFirewallProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileFirewallProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileFirewallProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFirewallProfilesByQueryCriteria
//
// Retrieve a list of Firewall Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFirewallProfileService) FindFirewallProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileFirewallProfileArray, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileFirewallProfileArray
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindFirewallProfilesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileFirewallProfileArray()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFirewallProfilesEthernetPortProfilesById
//
// Retrieve a EthernetPort Profile list of Firewall Profile is used by
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesEthernetPortProfilesById(ctx context.Context, id string) (*WSGEthernetPortProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGEthernetPortProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfilesEthernetPortProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEthernetPortProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFirewallProfilesWlansById
//
// Retrieve a WLAN list of Firewall Profile is used by
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesWlansById(ctx context.Context, id string) (*WSGWLANQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWLANQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfilesWlansById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateFirewallProfilesById
//
// Modify a Firewall Profile.
//
// Request Body:
//	 - body *WSGProfileModifyFirewallProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) UpdateFirewallProfilesById(ctx context.Context, body *WSGProfileModifyFirewallProfile, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateFirewallProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
