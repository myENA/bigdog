package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGRuckusGRETunnelProfileService struct {
	apiClient *APIClient
}

func NewWSGRuckusGRETunnelProfileService(c *APIClient) *WSGRuckusGRETunnelProfileService {
	s := new(WSGRuckusGRETunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRuckusGRETunnelProfileService() *WSGRuckusGRETunnelProfileService {
	return NewWSGRuckusGRETunnelProfileService(ss.apiClient)
}

// AddProfilesTunnelRuckusgre
//
// Use this API command to create RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileCreateRuckusGREProfile
func (s *WSGRuckusGRETunnelProfileService) AddProfilesTunnelRuckusgre(ctx context.Context, body *WSGProfileCreateRuckusGREProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddProfilesTunnelRuckusgre), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteProfilesTunnelRuckusgre
//
// Use this API command to delete multiple RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgre(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGDeleteProfilesTunnelRuckusgre), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteProfilesTunnelRuckusgreById
//
// Use this API command to delete RuckusGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgreById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGDeleteProfilesTunnelRuckusgreById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesTunnelRuckusgre
//
// Use this API command to retrieve a list of RuckusGRE tunnel profile.
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgre(ctx context.Context) (*WSGProfileList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindProfilesTunnelRuckusgre), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindProfilesTunnelRuckusgreById
//
// Use this API command to retrieve RuckusGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreById(ctx context.Context, id string) (*WSGProfileRuckusGREProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileRuckusGREProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindProfilesTunnelRuckusgreById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRuckusGREProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindProfilesTunnelRuckusgreByQueryCriteria
//
// Use this API command to query a list of RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileRuckusGREProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileRuckusGREProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindProfilesTunnelRuckusgreByQueryCriteria), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRuckusGREProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesTunnelRuckusgreById
//
// Use this API command to modify the configuration of RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileModifyRuckusGREProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) PartialUpdateProfilesTunnelRuckusgreById(ctx context.Context, body *WSGProfileModifyRuckusGREProfile, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateProfilesTunnelRuckusgreById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
