package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGFirewallProfileService struct {
	apiClient *APIClient
}

func NewWSGFirewallProfileService(c *APIClient) *WSGFirewallProfileService {
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
func (s *WSGFirewallProfileService) AddFirewallProfiles(ctx context.Context, body *WSGProfileCreateFirewallProfile) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddFirewallProfiles, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteFirewallProfiles
//
// Use this API command to delete Bulk Firewall Profiles.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGFirewallProfileService) DeleteFirewallProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFirewallProfiles, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteFirewallProfilesById
//
// Delete a Firewall Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) DeleteFirewallProfilesById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFirewallProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
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
func (s *WSGFirewallProfileService) FindFirewallProfiles(ctx context.Context, optionalParams map[string][]string) (*WSGProfileIdList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileIdList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfiles, true)
	if v, ok := optionalParams["domainId"]; ok {
		req.AddQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIdList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindFirewallProfilesById
//
// Retrieve a Firewall Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesById(ctx context.Context, id string) (*WSGProfileFirewallProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileFirewallProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileFirewallProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindFirewallProfilesByQueryCriteria
//
// Retrieve a list of Firewall Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFirewallProfileService) FindFirewallProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileFirewallProfileArray, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileFirewallProfileArray
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindFirewallProfilesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileFirewallProfileArray()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindFirewallProfilesEthernetPortProfilesById
//
// Retrieve a EthernetPort Profile list of Firewall Profile is used by
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesEthernetPortProfilesById(ctx context.Context, id string) (*WSGEthernetPortProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGEthernetPortProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfilesEthernetPortProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEthernetPortProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindFirewallProfilesWlansById
//
// Retrieve a WLAN list of Firewall Profile is used by
//
// Required Parameters:
// - id string
//		- required
func (s *WSGFirewallProfileService) FindFirewallProfilesWlansById(ctx context.Context, id string) (*WSGWLANQueryList, error) {
	var (
		req      *APIRequest
		resp     *WSGWLANQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFirewallProfilesWlansById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWLANQueryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGFirewallProfileService) UpdateFirewallProfilesById(ctx context.Context, body *WSGProfileModifyFirewallProfile, id string) (*WSGCommonEmptyResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateFirewallProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
