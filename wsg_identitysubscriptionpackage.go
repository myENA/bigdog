package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addIdentityPackageList
//
// Use this API command to retrieve a list of subscription package.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackageList(ctx context.Context, body *WSGIdentityQueryCriteria, mutators ...RequestMutator) (*WSGIdentitySubscriptionPackageListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentitySubscriptionPackageListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentitySubscriptionPackageListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityPackageList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGIdentitySubscriptionPackageListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentitySubscriptionPackageListAPIResponse), err
}

// AddIdentityPackages
//
// Operation ID: addIdentityPackages
//
// Use this API command to create subscription package.
//
// Request Body:
//	 - body *WSGIdentityCreateSubscriptionPackage
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackages(ctx context.Context, body *WSGIdentityCreateSubscriptionPackage, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityPackages, true)
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

// DeleteIdentityPackages
//
// Operation ID: deleteIdentityPackages
//
// Use this API command to delete multiple subscription packages.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackages(ctx context.Context, body *WSGIdentityDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityPackages, true)
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

// DeleteIdentityPackagesById
//
// Operation ID: deleteIdentityPackagesById
//
// Use this API command to delete subscription package.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackagesById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityPackagesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindIdentityPackages
//
// Operation ID: findIdentityPackages
//
// Use this API command to retrieve a list of subscription package.
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackages(ctx context.Context, mutators ...RequestMutator) (*WSGIdentitySubscriptionPackageListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentitySubscriptionPackageListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentitySubscriptionPackageListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityPackages, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentitySubscriptionPackageListAPIResponse), err
}

// FindIdentityPackagesById
//
// Operation ID: findIdentityPackagesById
//
// Use this API command to retrieve subscription package.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackagesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGIdentitySubscriptionPackageAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentitySubscriptionPackageAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentitySubscriptionPackageAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityPackagesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentitySubscriptionPackageAPIResponse), err
}

// PartialUpdateIdentityPackagesById
//
// Operation ID: partialUpdateIdentityPackagesById
//
// Use this API command to modify the configuration of subscription package.
//
// Request Body:
//	 - body *WSGIdentityModifySubscriptionPackage
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitySubscriptionPackageService) PartialUpdateIdentityPackagesById(ctx context.Context, body *WSGIdentityModifySubscriptionPackage, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateIdentityPackagesById, true)
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
