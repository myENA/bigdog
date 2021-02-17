package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGTTGPDGService struct {
	apiClient *VSZClient
}

func NewWSGTTGPDGService(c *VSZClient) *WSGTTGPDGService {
	s := new(WSGTTGPDGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTTGPDGService() *WSGTTGPDGService {
	return NewWSGTTGPDGService(ss.apiClient)
}

// AddProfilesTtgpdg
//
// Use this API command to create TTG+PDG profile.
//
// Operation ID: addProfilesTtgpdg
// Operation path: /profiles/ttgpdg
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateTtgpdgProfile
func (s *WSGTTGPDGService) AddProfilesTtgpdg(ctx context.Context, body *WSGProfileCreateTtgpdgProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesTtgpdg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesTtgpdg
//
// Use this API command to delete multiple TTG PDG profile.
//
// Operation ID: deleteProfilesTtgpdg
// Operation path: /profiles/ttgpdg
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGTTGPDGService) DeleteProfilesTtgpdg(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTtgpdg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesTtgpdgApnRealmsById
//
// Use this API command to disable the APN realm of TTG PDG profile.
//
// Operation ID: deleteProfilesTtgpdgApnRealmsById
// Operation path: /profiles/ttgpdg/{id}/apnRealms
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgApnRealmsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTtgpdgApnRealmsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesTtgpdgById
//
// Use this API command to delete TTG PDG profile.
//
// Operation ID: deleteProfilesTtgpdgById
// Operation path: /profiles/ttgpdg/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTtgpdgById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesTtgpdgDhcpRelayById
//
// Use this API command to disable the DHCP relay of TTG PDG profile.
//
// Operation ID: deleteProfilesTtgpdgDhcpRelayById
// Operation path: /profiles/ttgpdg/{id}/dhcpRelay
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgDhcpRelayById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesTtgpdgDhcpRelayById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesTtgpdg
//
// Use this API command to retrieve a list of TTG+PDG profile.
//
// Operation ID: findProfilesTtgpdg
// Operation path: /profiles/ttgpdg
// Success code: 200 (OK)
func (s *WSGTTGPDGService) FindProfilesTtgpdg(ctx context.Context, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTtgpdg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileListAPIResponse), err
}

// FindProfilesTtgpdgById
//
// Use this API command to retrieve TTG+PDG profile by ID.
//
// Operation ID: findProfilesTtgpdgById
// Operation path: /profiles/ttgpdg/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) FindProfilesTtgpdgById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileTtgpdgProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileTtgpdgProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesTtgpdgById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileTtgpdgProfileAPIResponse), err
}

// FindProfilesTtgpdgByQueryCriteria
//
// Use this API command to query a list of TTG+PDG profile.
//
// Operation ID: findProfilesTtgpdgByQueryCriteria
// Operation path: /profiles/ttgpdg/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGTTGPDGService) FindProfilesTtgpdgByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileTtgpdgProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileTtgpdgProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesTtgpdgByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileTtgpdgProfileListAPIResponse), err
}

// PartialUpdateProfilesTtgpdgById
//
// Use this API command to modify the configuration of TTG+PDG profile.
//
// Operation ID: partialUpdateProfilesTtgpdgById
// Operation path: /profiles/ttgpdg/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileTtgpdgProfileConfiguration
//
// Required parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) PartialUpdateProfilesTtgpdgById(ctx context.Context, body *WSGProfileTtgpdgProfileConfiguration, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesTtgpdgById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
