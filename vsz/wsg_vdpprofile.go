package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGVDPProfileService struct {
	apiClient *APIClient
}

func NewWSGVDPProfileService(c *APIClient) *WSGVDPProfileService {
	s := new(WSGVDPProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVDPProfileService() *WSGVDPProfileService {
	return NewWSGVDPProfileService(ss.apiClient)
}

// DeleteProfilesVdpById
//
// Use this API command to delete an vdp.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVDPProfileService) DeleteProfilesVdpById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesVdpById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindProfilesVdp
//
// Use this API command to retrieve a list of vdp.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGVDPProfileService) FindProfilesVdp(ctx context.Context, optionalParams map[string][]string) (*WSGProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesVdp, true)
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

// FindProfilesVdpById
//
// Use this API command to retrieve an vdp.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVDPProfileService) FindProfilesVdpById(ctx context.Context, id string) (*WSGProfileVdpProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileVdpProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesVdpById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileVdpProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateProfilesVdpApproveById
//
// Use this API command to approve vdp.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVDPProfileService) UpdateProfilesVdpApproveById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateProfilesVdpApproveById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}
