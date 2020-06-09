package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGMarkRogueService struct {
	apiClient *VSZClient
}

func NewWSGMarkRogueService(c *VSZClient) *WSGMarkRogueService {
	s := new(WSGMarkRogueService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGMarkRogueService() *WSGMarkRogueService {
	return NewWSGMarkRogueService(ss.apiClient)
}

// AddRogueMarkIgnore
//
// Mark a rogue AP as ignore.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkIgnore(ctx context.Context, body *WSGAPModifyRogueType) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkIgnore, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddRogueMarkKnown
//
// Mark a rogue AP as know.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkKnown(ctx context.Context, body *WSGAPModifyRogueType) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkKnown, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddRogueMarkMalicious
//
// Mark a rogue AP as malicious.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkMalicious(ctx context.Context, body *WSGAPModifyRogueType) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkMalicious, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddRogueMarkRogue
//
// Mark a rogue AP as rogue.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkRogue(ctx context.Context, body *WSGAPModifyRogueType) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkRogue, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddRogueUnMark
//
// Use this API command to remove the manual admin classification marking.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueUnMark(ctx context.Context, body *WSGAPModifyRogueType) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueUnMark, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRogueMarkKnown
//
// Get Known Rogue AP list.
func (s *WSGMarkRogueService) FindRogueMarkKnown(ctx context.Context) (*WSGAPModifyRogueType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPModifyRogueType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRogueMarkKnown, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPModifyRogueType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
