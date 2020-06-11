package bigdog

// API Version: v9_0

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
// Request Body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWiredClientService) AddWiredClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddWiredClientsBulkDeauth, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddWiredClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWiredClientService) AddWiredClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddWiredClientsDeauth, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindWiredclientByQueryCriteria
//
// Query wired clients with specified filters
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWiredClientService) FindWiredclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGWiredClientQueryClientQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGWiredClientQueryClientQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindWiredclientByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGWiredClientQueryClientQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
