package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGBlockClientService struct {
	apiClient *VSZClient
}

func NewWSGBlockClientService(c *VSZClient) *WSGBlockClientService {
	s := new(WSGBlockClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBlockClientService() *WSGBlockClientService {
	return NewWSGBlockClientService(ss.apiClient)
}

// AddBlockClient
//
// Create new Block Clients by list.
//
// Operation ID: addBlockClient
// Operation path: /blockClient
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileBulkBlockClient
func (s *WSGBlockClientService) AddBlockClient(ctx context.Context, body *WSGProfileBulkBlockClient, mutators ...RequestMutator) (*WSGProfileCreateResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileCreateResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddBlockClient, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileCreateResultListAPIResponse), err
}

// AddBlockClientByApMacByApMac
//
// Create a new Block Client by AP MAC.
//
// Operation ID: addBlockClientByApMacByApMac
// Operation path: /blockClient/byApMac/{apMac}
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileBlockClient
//
// Required parameters:
// - apMac string
//		- required
func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *WSGProfileBlockClient, apMac string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddBlockClientByApMacByApMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteBlockClient
//
// Delete Block Client List.
//
// Operation ID: deleteBlockClient
// Operation path: /blockClient
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBlockClientService) DeleteBlockClient(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteBlockClient, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteBlockClientById
//
// Delete a Block Client.
//
// Operation ID: deleteBlockClientById
// Operation path: /blockClient/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGBlockClientService) DeleteBlockClientById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteBlockClientById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindBlockClientById
//
// Retrieve a Block Client.
//
// Operation ID: findBlockClientById
// Operation path: /blockClient/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGBlockClientService) FindBlockClientById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileBlockClientAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileBlockClientAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindBlockClientById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBlockClientAPIResponse), err
}

// FindBlockClientByQueryCriteria
//
// Retrieve a list of Block Client.
//
// Operation ID: findBlockClientByQueryCriteria
// Operation path: /blockClient/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBlockClientService) FindBlockClientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileBlockClientListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileBlockClientListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindBlockClientByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBlockClientListAPIResponse), err
}

// FindBlockClientByZoneByZoneId
//
// Retrieve a list of Block Client.
//
// Operation ID: findBlockClientByZoneByZoneId
// Operation path: /blockClient/byZone/{zoneId}
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileBlockClientListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileBlockClientListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindBlockClientByZoneByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileBlockClientListAPIResponse), err
}

// PartialUpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Operation ID: partialUpdateBlockClientById
// Operation path: /blockClient/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyBlockClient
//
// Required parameters:
// - id string
//		- required
func (s *WSGBlockClientService) PartialUpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateBlockClientById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Operation ID: updateBlockClientById
// Operation path: /blockClient/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyBlockClient
//
// Required parameters:
// - id string
//		- required
func (s *WSGBlockClientService) UpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateBlockClientById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
