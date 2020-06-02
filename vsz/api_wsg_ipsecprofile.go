package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGIpsecprofileService struct {
	apiClient *APIClient
}

func NewWSGIpsecprofileService(c *APIClient) *WSGIpsecprofileService {
	s := new(WSGIpsecprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIpsecprofileService() *WSGIpsecprofileService {
	return NewWSGIpsecprofileService(ss.apiClient)
}

// AddProfilesTunnelIpsec
//
// Create a new ipsec.
//
// Request Body:
//	 - body *WSGProfileCreateIpsecProfile
func (s *WSGIpsecprofileService) AddProfilesTunnelIpsec(ctx context.Context, body *WSGProfileCreateIpsecProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesTunnelIpsec, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteProfilesTunnelIpsec
//
// Delete multiple ipsec.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGIpsecprofileService) DeleteProfilesTunnelIpsec(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsec, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesTunnelIpsecById
//
// Delete a ipsec.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIpsecprofileService) DeleteProfilesTunnelIpsecById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsecById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesTunnelIpsec
//
// Retrieve a list of IPSEC.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGIpsecprofileService) FindProfilesTunnelIpsec(ctx context.Context, optionalParams map[string][]string) (*WSGProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelIpsec, true)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesTunnelIpsecById
//
// Retrieve a IPSEC.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIpsecprofileService) FindProfilesTunnelIpsecById(ctx context.Context, id string) (*WSGProfileIpsecProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileIpsecProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelIpsecById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIpsecProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesTunnelIpsecByQueryCriteria
//
// Query a list of IPSEC.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIpsecprofileService) FindProfilesTunnelIpsecByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileIpsecProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileIpsecProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTunnelIpsecByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIpsecProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesTunnelIpsecById
//
// Modify a specific ipsec basic.
//
// Request Body:
//	 - body *WSGProfileModifyIpsecProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIpsecprofileService) PartialUpdateProfilesTunnelIpsecById(ctx context.Context, body *WSGProfileModifyIpsecProfile, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelIpsecById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
