package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGIdentitySubscriptionPackageService struct {
	apiClient *VSZClient
}

func NewWSGIdentitySubscriptionPackageService(c *VSZClient) *WSGIdentitySubscriptionPackageService {
	s := new(WSGIdentitySubscriptionPackageService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentitySubscriptionPackageService() *WSGIdentitySubscriptionPackageService {
	return NewWSGIdentitySubscriptionPackageService(ss.apiClient)
}

// AddIdentityPackageList
//
// Use this API command to retrieve a list of subscription package.
//
// Operation ID: addIdentityPackageList
// Operation path: /identity/packageList
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackageList(ctx context.Context, body *WSGIdentityQueryCriteria, mutators ...RequestMutator) (*WSGIdentitySubscriptionPackageListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentitySubscriptionPackageListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddIdentityPackageList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentitySubscriptionPackageListAPIResponse), err
}

// AddIdentityPackages
//
// Use this API command to create subscription package.
//
// Operation ID: addIdentityPackages
// Operation path: /identity/packages
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGIdentityCreateSubscriptionPackage
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackages(ctx context.Context, body *WSGIdentityCreateSubscriptionPackage, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddIdentityPackages, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteIdentityPackages
//
// Use this API command to delete multiple subscription packages.
//
// Operation ID: deleteIdentityPackages
// Operation path: /identity/packages
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackages(ctx context.Context, body *WSGIdentityDeleteBulk, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteIdentityPackages, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteIdentityPackagesById
//
// Use this API command to delete subscription package.
//
// Operation ID: deleteIdentityPackagesById
// Operation path: /identity/packages/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackagesById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteIdentityPackagesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindIdentityPackages
//
// Use this API command to retrieve a list of subscription package.
//
// Operation ID: findIdentityPackages
// Operation path: /identity/packages
// Success code: 200 (OK)
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackages(ctx context.Context, mutators ...RequestMutator) (*WSGIdentitySubscriptionPackageListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentitySubscriptionPackageListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityPackages, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentitySubscriptionPackageListAPIResponse), err
}

// FindIdentityPackagesById
//
// Use this API command to retrieve subscription package.
//
// Operation ID: findIdentityPackagesById
// Operation path: /identity/packages/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackagesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGIdentitySubscriptionPackageAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIdentitySubscriptionPackageAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindIdentityPackagesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIdentitySubscriptionPackageAPIResponse), err
}

// PartialUpdateIdentityPackagesById
//
// Use this API command to modify the configuration of subscription package.
//
// Operation ID: partialUpdateIdentityPackagesById
// Operation path: /identity/packages/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGIdentityModifySubscriptionPackage
//
// Required parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) PartialUpdateIdentityPackagesById(ctx context.Context, body *WSGIdentityModifySubscriptionPackage, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateIdentityPackagesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
