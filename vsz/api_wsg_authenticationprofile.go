package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAuthenticationProfileService struct {
	apiClient *APIClient
}

func NewWSGAuthenticationProfileService(c *APIClient) *WSGAuthenticationProfileService {
	s := new(WSGAuthenticationProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAuthenticationProfileService() *WSGAuthenticationProfileService {
	return NewWSGAuthenticationProfileService(ss.apiClient)
}

// AddProfilesAuth
//
// Use this API command to create a new authentication profile.
//
// Request Body:
//	 - body *WSGProfileCreateAuthenticationProfile
func (s *WSGAuthenticationProfileService) AddProfilesAuth(ctx context.Context, body *WSGProfileCreateAuthenticationProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAuth, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddProfilesAuthCloneById
//
// Use this API command to clone an authentication profile.
//
// Request Body:
//	 - body *WSGProfileClone
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) AddProfilesAuthCloneById(ctx context.Context, body *WSGProfileClone, id string) (*WSGProfileClone, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileClone
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAuthCloneById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClone()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteProfilesAuth
//
// Use this API command to delete a list of authentication profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkAuthenticationProfile
func (s *WSGAuthenticationProfileService) DeleteProfilesAuth(ctx context.Context, body *WSGProfileDeleteBulkAuthenticationProfile) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAuth, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesAuthById
//
// Use this API command to delete an authentication profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) DeleteProfilesAuthById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAuthById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesAuth
//
// Use this API command to retrieve a list of authentication profiles.
func (s *WSGAuthenticationProfileService) FindProfilesAuth(ctx context.Context) (*WSGProfileAuthenticationProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileAuthenticationProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAuth, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileAuthenticationProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesAuthAuthorizationList
//
// Use this API command to retrieve a list of authorization profiles.
//
// Required Parameters:
// - type_ string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthorizationList(ctx context.Context, type_ string) (*WSGProfileBaseServiceInfoList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBaseServiceInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAuthAuthorizationList, true)
	req.SetQueryParameter("type_", []string{type_})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBaseServiceInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesAuthAuthServiceListByQueryCriteria
//
// Use this API command to retrieve a list of authentication service.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthServiceListByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBaseServiceInfoList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBaseServiceInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesAuthAuthServiceListByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBaseServiceInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesAuthById
//
// Use this API command to retrieve an authentication profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthById(ctx context.Context, id string) (*WSGProfileAuthenticationProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileAuthenticationProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAuthById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileAuthenticationProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesAuthByQueryCriteria
//
// Use this API command to retrieve a list of authentication profiles by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileAuthenticationProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileAuthenticationProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesAuthByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileAuthenticationProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRadiusProxyStatsByQueryCriteria
//
// Use this API command to retrieve a list of Radius Proxy.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindRadiusProxyStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGRacStatsRadiusProxyList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRacStatsRadiusProxyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRadiusProxyStatsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGRacStatsRadiusProxyList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesAuthById
//
// Use this API command to modify the configuration of an authentication profile.
//
// Request Body:
//	 - body *WSGProfileModifyAuthenticationProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) PartialUpdateProfilesAuthById(ctx context.Context, body *WSGProfileModifyAuthenticationProfile, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesAuthById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
