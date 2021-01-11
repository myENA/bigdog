package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGIdentityUserRoleService struct {
	apiClient *VSZClient
}

func NewWSGIdentityUserRoleService(c *VSZClient) *WSGIdentityUserRoleService {
	s := new(WSGIdentityUserRoleService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityUserRoleService() *WSGIdentityUserRoleService {
	return NewWSGIdentityUserRoleService(ss.apiClient)
}

// AddIdentityUserrole
//
// Use this API command to create identity user role.
//
// Operation ID: addIdentityUserrole
// Operation path: /identity/userrole
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGIdentityCreateIdentityUserRole
func (s *WSGIdentityUserRoleService) AddIdentityUserrole(ctx context.Context, body *WSGIdentityCreateIdentityUserRole, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddIdentityUserrole, true)
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

// AddIdentityUserRoleList
//
// Use this API command to retrieve a list of identity user role.
//
// Operation ID: addIdentityUserRoleList
// Operation path: /identity/userRoleList
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserRoleService) AddIdentityUserRoleList(ctx context.Context, body *WSGIdentityQueryCriteria, mutators ...RequestMutator) (*WSGIdentityListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddIdentityUserRoleList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGIdentityListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityListAPIResponse), err
}

// DeleteIdentityUserrole
//
// Use this API command to delete multiple identity user roles.
//
// Operation ID: deleteIdentityUserrole
// Operation path: /identity/userrole
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserRoleService) DeleteIdentityUserrole(ctx context.Context, body *WSGIdentityDeleteBulk, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteIdentityUserrole, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteIdentityUserroleById
//
// Use this API command to delete identity user role.
//
// Operation ID: deleteIdentityUserroleById
// Operation path: /identity/userrole/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) DeleteIdentityUserroleById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteIdentityUserroleById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindIdentityUserrole
//
// Use this API command to retrieve a list of identity user role.
//
// Operation ID: findIdentityUserrole
// Operation path: /identity/userrole
// Success code: 200 (OK)
func (s *WSGIdentityUserRoleService) FindIdentityUserrole(ctx context.Context, mutators ...RequestMutator) (*WSGIdentityListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityUserrole, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityListAPIResponse), err
}

// FindIdentityUserroleById
//
// Use this API command to retrieve identity user role by ID.
//
// Operation ID: findIdentityUserroleById
// Operation path: /identity/userrole/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) FindIdentityUserroleById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGIdentityUserRoleAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityUserRoleAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityUserRoleAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityUserroleById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityUserRoleAPIResponse), err
}

// PartialUpdateIdentityUserroleById
//
// Use this API command to modify the configuration of identity user role.
//
// Operation ID: partialUpdateIdentityUserroleById
// Operation path: /identity/userrole/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGIdentityModifyIdentityUserRole
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentityUserRoleService) PartialUpdateIdentityUserroleById(ctx context.Context, body *WSGIdentityModifyIdentityUserRole, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateIdentityUserroleById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
