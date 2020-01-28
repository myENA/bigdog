package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGIPSECProfileService struct {
	apiClient *APIClient
}

func NewWSGIPSECProfileService(c *APIClient) *WSGIPSECProfileService {
	s := new(WSGIPSECProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIPSECProfileService() *WSGIPSECProfileService {
	return NewWSGIPSECProfileService(ss.apiClient)
}

// AddProfilesTunnelIpsec
//
// Create a new ipsec.
//
// Request Body:
//	 - body *WSGProfileCreateIpsecProfile
func (s *WSGIPSECProfileService) AddProfilesTunnelIpsec(ctx context.Context, body *WSGProfileCreateIpsecProfile) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesTunnelIpsec, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteProfilesTunnelIpsec
//
// Delete multiple ipsec.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsec(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsec, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteProfilesTunnelIpsecById
//
// Delete a ipsec.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsecById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsecById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsec(ctx context.Context, optionalParams map[string][]string) (*WSGProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelIpsec, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindProfilesTunnelIpsecById
//
// Retrieve a IPSEC.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecById(ctx context.Context, id string) (*WSGProfileIpsecProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileIpsecProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelIpsecById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIpsecProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindProfilesTunnelIpsecByQueryCriteria
//
// Query a list of IPSEC.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileIpsecProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileIpsecProfileList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTunnelIpsecByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIpsecProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGIPSECProfileService) PartialUpdateProfilesTunnelIpsecById(ctx context.Context, body *WSGProfileModifyIpsecProfile, id string) (*WSGCommonEmptyResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelIpsecById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
