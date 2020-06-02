package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGIdentitysubscriptionpackageService struct {
	apiClient *APIClient
}

func NewWSGIdentitysubscriptionpackageService(c *APIClient) *WSGIdentitysubscriptionpackageService {
	s := new(WSGIdentitysubscriptionpackageService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentitysubscriptionpackageService() *WSGIdentitysubscriptionpackageService {
	return NewWSGIdentitysubscriptionpackageService(ss.apiClient)
}

// AddIdentityPackageList
//
// Use this API command to retrieve a list of subscription package.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentitysubscriptionpackageService) AddIdentityPackageList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentitySubscriptionPackageList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentitySubscriptionPackageList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityPackageList, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentitySubscriptionPackageList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddIdentityPackages
//
// Use this API command to create subscription package.
//
// Request Body:
//	 - body *WSGIdentityCreateSubscriptionPackage
func (s *WSGIdentitysubscriptionpackageService) AddIdentityPackages(ctx context.Context, body *WSGIdentityCreateSubscriptionPackage) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityPackages, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteIdentityPackages
//
// Use this API command to delete multiple subscription packages.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentitysubscriptionpackageService) DeleteIdentityPackages(ctx context.Context, body *WSGIdentityDeleteBulk) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityPackages, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteIdentityPackagesById
//
// Use this API command to delete subscription package.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitysubscriptionpackageService) DeleteIdentityPackagesById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityPackagesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindIdentityPackages
//
// Use this API command to retrieve a list of subscription package.
func (s *WSGIdentitysubscriptionpackageService) FindIdentityPackages(ctx context.Context) (*WSGIdentitySubscriptionPackageList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentitySubscriptionPackageList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityPackages, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentitySubscriptionPackageList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindIdentityPackagesById
//
// Use this API command to retrieve subscription package.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitysubscriptionpackageService) FindIdentityPackagesById(ctx context.Context, id string) (*WSGIdentitySubscriptionPackage, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentitySubscriptionPackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityPackagesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentitySubscriptionPackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateIdentityPackagesById
//
// Use this API command to modify the configuration of subscription package.
//
// Request Body:
//	 - body *WSGIdentityModifySubscriptionPackage
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentitysubscriptionpackageService) PartialUpdateIdentityPackagesById(ctx context.Context, body *WSGIdentityModifySubscriptionPackage, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateIdentityPackagesById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
