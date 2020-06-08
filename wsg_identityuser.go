package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGIdentityuserService struct {
	apiClient *VSZClient
}

func NewWSGIdentityuserService(c *VSZClient) *WSGIdentityuserService {
	s := new(WSGIdentityuserService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityuserService() *WSGIdentityuserService {
	return NewWSGIdentityuserService(ss.apiClient)
}

// AddIdentityUserList
//
// Use this API command to retrieve a list of identity user.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityuserService) AddIdentityUserList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityUserList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityUserList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUserList, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddIdentityUsers
//
// Use this API command to create identity user.
//
// Request Body:
//	 - body *WSGIdentityCreateUser
func (s *WSGIdentityuserService) AddIdentityUsers(ctx context.Context, body *WSGIdentityCreateUser) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUsers, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteIdentityUsers
//
// Use this API command to delete multiple identity users.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityuserService) DeleteIdentityUsers(ctx context.Context, body *WSGIdentityDeleteBulk) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUsers, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteIdentityUsersById
//
// Use this API command to delete identity user.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityuserService) DeleteIdentityUsersById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUsersById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindIdentityUsers
//
// Use this API command to retrieve a list of identity user.
//
// Optional Parameters:
// - createdOnFrom string
//		- nullable
// - createdOnTo string
//		- nullable
// - displayName string
//		- nullable
// - email string
//		- nullable
// - firstName string
//		- nullable
// - index string
//		- nullable
// - isDisabled string
//		- nullable
// - lastName string
//		- nullable
// - listSize string
//		- nullable
// - phone string
//		- nullable
// - timeZone string
//		- nullable
// - userName string
//		- nullable
// - userSource string
//		- nullable
// - userType string
//		- nullable
func (s *WSGIdentityuserService) FindIdentityUsers(ctx context.Context, optionalParams map[string][]string) (*WSGIdentityUserList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityUserList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsers, true)
	if v, ok := optionalParams["createdOnFrom"]; ok && len(v) > 0 {
		req.SetQueryParameter("createdOnFrom", v)
	}
	if v, ok := optionalParams["createdOnTo"]; ok && len(v) > 0 {
		req.SetQueryParameter("createdOnTo", v)
	}
	if v, ok := optionalParams["displayName"]; ok && len(v) > 0 {
		req.SetQueryParameter("displayName", v)
	}
	if v, ok := optionalParams["email"]; ok && len(v) > 0 {
		req.SetQueryParameter("email", v)
	}
	if v, ok := optionalParams["firstName"]; ok && len(v) > 0 {
		req.SetQueryParameter("firstName", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["isDisabled"]; ok && len(v) > 0 {
		req.SetQueryParameter("isDisabled", v)
	}
	if v, ok := optionalParams["lastName"]; ok && len(v) > 0 {
		req.SetQueryParameter("lastName", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["phone"]; ok && len(v) > 0 {
		req.SetQueryParameter("phone", v)
	}
	if v, ok := optionalParams["timeZone"]; ok && len(v) > 0 {
		req.SetQueryParameter("timeZone", v)
	}
	if v, ok := optionalParams["userName"]; ok && len(v) > 0 {
		req.SetQueryParameter("userName", v)
	}
	if v, ok := optionalParams["userSource"]; ok && len(v) > 0 {
		req.SetQueryParameter("userSource", v)
	}
	if v, ok := optionalParams["userType"]; ok && len(v) > 0 {
		req.SetQueryParameter("userType", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIdentityUsersAaaserver
//
// Use this API command to retrieve a list of aaa server.
func (s *WSGIdentityuserService) FindIdentityUsersAaaserver(ctx context.Context) (*WSGIdentityAaaServerList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityAaaServerList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersAaaserver, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityAaaServerList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIdentityUsersById
//
// Use this API command to retrieve identity user.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityuserService) FindIdentityUsersById(ctx context.Context, id string) (*WSGIdentityUserConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityUserConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIdentityUsersCountries
//
// Use this API command to retrieve a list of countries.
func (s *WSGIdentityuserService) FindIdentityUsersCountries(ctx context.Context) (*WSGIdentityCountryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityCountryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersCountries, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityCountryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIdentityUsersPackages
//
// Use this API command to retrieve a list of packages.
func (s *WSGIdentityuserService) FindIdentityUsersPackages(ctx context.Context) (*WSGIdentityPackageList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityPackageList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersPackages, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityPackageList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateIdentityUsersById
//
// Use this API command to modify the configuration of identity user.
//
// Request Body:
//	 - body *WSGIdentityModifyUser
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityuserService) PartialUpdateIdentityUsersById(ctx context.Context, body *WSGIdentityModifyUser, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateIdentityUsersById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
