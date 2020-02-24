package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGIdentityUserService struct {
	apiClient *APIClient
}

func NewWSGIdentityUserService(c *APIClient) *WSGIdentityUserService {
	s := new(WSGIdentityUserService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityUserService() *WSGIdentityUserService {
	return NewWSGIdentityUserService(ss.apiClient)
}

// AddIdentityUserList
//
// Use this API command to retrieve a list of identity user.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityUserList, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUserList, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddIdentityUsers
//
// Use this API command to create identity user.
//
// Request Body:
//	 - body *WSGIdentityCreateUser
func (s *WSGIdentityUserService) AddIdentityUsers(ctx context.Context, body *WSGIdentityCreateUser) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUsers, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteIdentityUsers
//
// Use this API command to delete multiple identity users.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserService) DeleteIdentityUsers(ctx context.Context, body *WSGIdentityDeleteBulk) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUsers, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
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
func (s *WSGIdentityUserService) DeleteIdentityUsersById(ctx context.Context, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUsersById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, optionalParams map[string][]string) (*WSGIdentityUserList, *APIResponseMeta, error) {
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
	if v, ok := optionalParams["createdOnFrom"]; ok {
		req.AddQueryParameter("createdOnFrom", v)
	}
	if v, ok := optionalParams["createdOnTo"]; ok {
		req.AddQueryParameter("createdOnTo", v)
	}
	if v, ok := optionalParams["displayName"]; ok {
		req.AddQueryParameter("displayName", v)
	}
	if v, ok := optionalParams["email"]; ok {
		req.AddQueryParameter("email", v)
	}
	if v, ok := optionalParams["firstName"]; ok {
		req.AddQueryParameter("firstName", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["isDisabled"]; ok {
		req.AddQueryParameter("isDisabled", v)
	}
	if v, ok := optionalParams["lastName"]; ok {
		req.AddQueryParameter("lastName", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["phone"]; ok {
		req.AddQueryParameter("phone", v)
	}
	if v, ok := optionalParams["timeZone"]; ok {
		req.AddQueryParameter("timeZone", v)
	}
	if v, ok := optionalParams["userName"]; ok {
		req.AddQueryParameter("userName", v)
	}
	if v, ok := optionalParams["userSource"]; ok {
		req.AddQueryParameter("userSource", v)
	}
	if v, ok := optionalParams["userType"]; ok {
		req.AddQueryParameter("userType", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindIdentityUsersAaaserver
//
// Use this API command to retrieve a list of aaa server.
func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context) (*WSGIdentityAaaServerList, *APIResponseMeta, error) {
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindIdentityUsersById
//
// Use this API command to retrieve identity user.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, id string) (*WSGIdentityUserConfiguration, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityUserConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindIdentityUsersCountries
//
// Use this API command to retrieve a list of countries.
func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context) (*WSGIdentityCountryList, *APIResponseMeta, error) {
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindIdentityUsersPackages
//
// Use this API command to retrieve a list of packages.
func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context) (*WSGIdentityPackageList, *APIResponseMeta, error) {
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGIdentityUserService) PartialUpdateIdentityUsersById(ctx context.Context, body *WSGIdentityModifyUser, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateIdentityUsersById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
