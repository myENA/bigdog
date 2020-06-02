package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGWiredclientService struct {
	apiClient *APIClient
}

func NewWSGWiredclientService(c *APIClient) *WSGWiredclientService {
	s := new(WSGWiredclientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWiredclientService() *WSGWiredclientService {
	return NewWSGWiredclientService(ss.apiClient)
}

// AddWiredClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWiredclientService) AddWiredClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddWiredClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWiredclientService) AddWiredClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindWiredclientByQueryCriteria
//
// Query wired clients with specified filters
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWiredclientService) FindWiredclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGWiredClientQueryClientQueryList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGWiredClientQueryClientQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
