package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGBridgeService struct {
	apiClient *VSZClient
}

func NewWSGBridgeService(c *VSZClient) *WSGBridgeService {
	s := new(WSGBridgeService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBridgeService() *WSGBridgeService {
	return NewWSGBridgeService(ss.apiClient)
}

// AddProfilesBridge
//
// Operation ID: addProfilesBridge
//
// Use this API command to create Bridge profile.
//
// Request Body:
//	 - body *WSGProfileCreateBridgeProfile
func (s *WSGBridgeService) AddProfilesBridge(ctx context.Context, body *WSGProfileCreateBridgeProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddProfilesBridge, true)
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

// DeleteProfilesBridge
//
// Operation ID: deleteProfilesBridge
//
// Use this API command to delete multiple bridge profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBridgeService) DeleteProfilesBridge(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesBridge, true)
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

// DeleteProfilesBridgeById
//
// Operation ID: deleteProfilesBridgeById
//
// Use this API command to delete Bridge profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBridgeService) DeleteProfilesBridgeById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesBridgeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindProfilesBridge
//
// Operation ID: findProfilesBridge
//
// Use this API command to retrieve a list of Bridge profile.
func (s *WSGBridgeService) FindProfilesBridge(ctx context.Context, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesBridge, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileListAPIResponse), err
}

// FindProfilesBridgeById
//
// Operation ID: findProfilesBridgeById
//
// Use this API command to retrieve Bridge profile by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBridgeService) FindProfilesBridgeById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileBridgeProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBridgeProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileBridgeProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesBridgeById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBridgeProfileAPIResponse), err
}

// FindProfilesBridgeByQueryCriteria
//
// Operation ID: findProfilesBridgeByQueryCriteria
//
// Use this API command to query a list of Bridge profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBridgeService) FindProfilesBridgeByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileBridgeProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileBridgeProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileBridgeProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindProfilesBridgeByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileBridgeProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBridgeProfileListAPIResponse), err
}

// PartialUpdateProfilesBridgeById
//
// Operation ID: partialUpdateProfilesBridgeById
//
// Use this API command to modify the configuration of Bridge profile.
//
// Request Body:
//	 - body *WSGProfileModifyBridgeProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBridgeService) PartialUpdateProfilesBridgeById(ctx context.Context, body *WSGProfileModifyBridgeProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesBridgeById, true)
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
