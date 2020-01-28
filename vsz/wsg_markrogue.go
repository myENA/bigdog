package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGMarkRogueService struct {
	apiClient *APIClient
}

func NewWSGMarkRogueService(c *APIClient) *WSGMarkRogueService {
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
func (s *WSGMarkRogueService) AddRogueMarkIgnore(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkIgnore, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddRogueMarkKnown
//
// Mark a rogue AP as know.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkKnown(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkKnown, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddRogueMarkMalicious
//
// Mark a rogue AP as malicious.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkMalicious(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkMalicious, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddRogueMarkRogue
//
// Mark a rogue AP as rogue.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkRogue(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkRogue, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddRogueUnMark
//
// Use this API command to remove the manual admin classification marking.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueUnMark(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRogueUnMark, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindRogueMarkKnown
//
// Get Known Rogue AP list.
func (s *WSGMarkRogueService) FindRogueMarkKnown(ctx context.Context) (*WSGAPModifyRogueType, error) {
	var (
		req      *APIRequest
		resp     *WSGAPModifyRogueType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRogueMarkKnown, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPModifyRogueType()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
