package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGIdentityUserService struct {
	apiClient *VSZClient
}

func NewWSGIdentityUserService(c *VSZClient) *WSGIdentityUserService {
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
// Operation ID: addIdentityUserList
// Operation path: /identity/userList
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *WSGIdentityQueryCriteria, mutators ...RequestMutator) (*WSGIdentityUserListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentityUserListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddIdentityUserList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentityUserListAPIResponse), err
}

// AddIdentityUsers
//
// Use this API command to create identity user.
//
// Operation ID: addIdentityUsers
// Operation path: /identity/users
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGIdentityCreateUser
func (s *WSGIdentityUserService) AddIdentityUsers(ctx context.Context, body *WSGIdentityCreateUser, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddIdentityUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteIdentityUsers
//
// Use this API command to delete multiple identity users.
//
// Operation ID: deleteIdentityUsers
// Operation path: /identity/users
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserService) DeleteIdentityUsers(ctx context.Context, body *WSGIdentityDeleteBulk, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteIdentityUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteIdentityUsersById
//
// Use this API command to delete identity user.
//
// Operation ID: deleteIdentityUsersById
// Operation path: /identity/users/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) DeleteIdentityUsersById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteIdentityUsersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindIdentityUsers
//
// Use this API command to retrieve a list of identity user.
//
// Operation ID: findIdentityUsers
// Operation path: /identity/users
// Success code: 200 (OK)
//
// Optional parameters:
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
func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGIdentityUserListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentityUserListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["createdOnFrom"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("createdOnFrom", v)
	}
	if v, ok := optionalParams["createdOnTo"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("createdOnTo", v)
	}
	if v, ok := optionalParams["displayName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("displayName", v)
	}
	if v, ok := optionalParams["email"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("email", v)
	}
	if v, ok := optionalParams["firstName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("firstName", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["isDisabled"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("isDisabled", v)
	}
	if v, ok := optionalParams["lastName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("lastName", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["phone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("phone", v)
	}
	if v, ok := optionalParams["timeZone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timeZone", v)
	}
	if v, ok := optionalParams["userName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("userName", v)
	}
	if v, ok := optionalParams["userSource"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("userSource", v)
	}
	if v, ok := optionalParams["userType"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("userType", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentityUserListAPIResponse), err
}

// FindIdentityUsersAaaserver
//
// Use this API command to retrieve a list of aaa server.
//
// Operation ID: findIdentityUsersAaaserver
// Operation path: /identity/users/aaaserver
// Success code: 200 (OK)
func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context, mutators ...RequestMutator) (*WSGIdentityAaaServerListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentityAaaServerListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityUsersAaaserver, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentityAaaServerListAPIResponse), err
}

// FindIdentityUsersById
//
// Use this API command to retrieve identity user.
//
// Operation ID: findIdentityUsersById
// Operation path: /identity/users/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGIdentityUserConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentityUserConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityUsersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentityUserConfigurationAPIResponse), err
}

// FindIdentityUsersCountries
//
// Use this API command to retrieve a list of countries.
//
// Operation ID: findIdentityUsersCountries
// Operation path: /identity/users/countries
// Success code: 200 (OK)
func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context, mutators ...RequestMutator) (*WSGIdentityCountryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentityCountryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityUsersCountries, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentityCountryListAPIResponse), err
}

// FindIdentityUsersPackages
//
// Use this API command to retrieve a list of packages.
//
// Operation ID: findIdentityUsersPackages
// Operation path: /identity/users/packages
// Success code: 200 (OK)
func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context, mutators ...RequestMutator) (*WSGIdentityPackageListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentityPackageListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityUsersPackages, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentityPackageListAPIResponse), err
}

// PartialUpdateIdentityUsersById
//
// Use this API command to modify the configuration of identity user.
//
// Operation ID: partialUpdateIdentityUsersById
// Operation path: /identity/users/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGIdentityModifyUser
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) PartialUpdateIdentityUsersById(ctx context.Context, body *WSGIdentityModifyUser, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateIdentityUsersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
