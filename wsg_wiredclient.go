package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGWiredClientService struct {
	apiClient *VSZClient
}

func NewWSGWiredClientService(c *VSZClient) *WSGWiredClientService {
	s := new(WSGWiredClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWiredClientService() *WSGWiredClientService {
	return NewWSGWiredClientService(ss.apiClient)
}

// AddWiredClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Operation ID: addWiredClientsBulkDeauth
// Operation path: /wiredClients/bulkDeauth
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWiredClientService) AddWiredClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddWiredClientsBulkDeauth, true)
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

// AddWiredClientsDeauth
//
// Use this API command to deauth client.
//
// Operation ID: addWiredClientsDeauth
// Operation path: /wiredClients/deauth
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWiredClientService) AddWiredClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddWiredClientsDeauth, true)
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

// FindWiredclientByQueryCriteria
//
// Query wired clients with specified filters
//
// Operation ID: findWiredclientByQueryCriteria
// Operation path: /query/wiredclient
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWiredClientService) FindWiredclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGWiredClientQueryClientQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWiredClientQueryClientQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWiredClientQueryClientQueryListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindWiredclientByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGWiredClientQueryClientQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWiredClientQueryClientQueryListAPIResponse), err
}
