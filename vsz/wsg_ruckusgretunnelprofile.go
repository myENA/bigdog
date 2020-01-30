package vsz

// API Version: v9_0

import (
	"context"
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
func (s *WSGRuckusGRETunnelProfileService) AddProfilesTunnelRuckusgre(ctx context.Context, body *WSGProfileCreateRuckusGREProfile) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesTunnelRuckusgre, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteProfilesTunnelRuckusgre
//
// Use this API command to delete multiple RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgre(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelRuckusgre, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteProfilesTunnelRuckusgreById
//
// Use this API command to delete RuckusGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgreById(ctx context.Context, id string) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelRuckusgreById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindProfilesTunnelRuckusgre
//
// Use this API command to retrieve a list of RuckusGRE tunnel profile.
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgre(ctx context.Context) (*WSGProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelRuckusgre, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindProfilesTunnelRuckusgreById
//
// Use this API command to retrieve RuckusGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreById(ctx context.Context, id string) (*WSGProfileRuckusGREProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileRuckusGREProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelRuckusgreById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRuckusGREProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindProfilesTunnelRuckusgreByQueryCriteria
//
// Use this API command to query a list of RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileRuckusGREProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileRuckusGREProfileList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTunnelRuckusgreByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRuckusGREProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGRuckusGRETunnelProfileService) PartialUpdateProfilesTunnelRuckusgreById(ctx context.Context, body *WSGProfileModifyRuckusGREProfile, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelRuckusgreById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
