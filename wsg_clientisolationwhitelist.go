package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGClientIsolationWhitelistService struct {
	apiClient *VSZClient
}

func NewWSGClientIsolationWhitelistService(c *VSZClient) *WSGClientIsolationWhitelistService {
	s := new(WSGClientIsolationWhitelistService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService() *WSGClientIsolationWhitelistService {
	return NewWSGClientIsolationWhitelistService(ss.apiClient)
}

// AddRkszonesClientIsolationWhitelistByZoneId
//
// Operation ID: addRkszonesClientIsolationWhitelistByZoneId
//
// Create a new ClientIsolationWhitelist.
//
// Request Body:
//	 - body *WSGProfileCreateClientIsolationWhitelist
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) AddRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, body *WSGProfileCreateClientIsolationWhitelist, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesClientIsolationWhitelistByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesClientIsolationWhitelist
//
// Operation ID: deleteRkszonesClientIsolationWhitelist
//
// Use this API command to delete Bulk Client Isolation Whitelist.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelist(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelist, true)
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

// DeleteRkszonesClientIsolationWhitelistById
//
// Operation ID: deleteRkszonesClientIsolationWhitelistById
//
// Delete a Client Isolation Whitelist.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelistById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelistById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindRkszonesClientIsolationWhitelistById
//
// Operation ID: findRkszonesClientIsolationWhitelistById
//
// Retrieve an Client Isolation Whitelist.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileClientIsolationWhitelistAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileClientIsolationWhitelistAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileClientIsolationWhitelistAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileClientIsolationWhitelistAPIResponse), err
}

// FindRkszonesClientIsolationWhitelistByZoneId
//
// Operation ID: findRkszonesClientIsolationWhitelistByZoneId
//
// Retrieve a list of Client Isolation Whitelist.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileClientIsolationWhitelistArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileClientIsolationWhitelistArrayAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileClientIsolationWhitelistArrayAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileClientIsolationWhitelistArrayAPIResponse), err
}

// FindServicesClientIsolationWhitelistByQueryCriteria
//
// Operation ID: findServicesClientIsolationWhitelistByQueryCriteria
//
// Retrieve a list of Client Isolation Whitelist.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGClientIsolationWhitelistService) FindServicesClientIsolationWhitelistByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileClientIsolationWhitelistArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileClientIsolationWhitelistArrayAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileClientIsolationWhitelistArrayAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesClientIsolationWhitelistByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileClientIsolationWhitelistArrayAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileClientIsolationWhitelistArrayAPIResponse), err
}

// PartialUpdateRkszonesClientIsolationWhitelistById
//
// Operation ID: partialUpdateRkszonesClientIsolationWhitelistById
//
// Modify a specific Client Isolation Whitelist basic.
//
// Request Body:
//	 - body *WSGProfileModifyClientIsolationWhitelist
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) PartialUpdateRkszonesClientIsolationWhitelistById(ctx context.Context, body *WSGProfileModifyClientIsolationWhitelist, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesClientIsolationWhitelistById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
