package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGL3AccessControlPolicyService struct {
	apiClient *APIClient
}

func NewWSGL3AccessControlPolicyService(c *APIClient) *WSGL3AccessControlPolicyService {
	s := new(WSGL3AccessControlPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL3AccessControlPolicyService() *WSGL3AccessControlPolicyService {
	return NewWSGL3AccessControlPolicyService(ss.apiClient)
}

// AddL3AccessControlPolicies
//
// Create a L3 Access Control Policy.
//
// Request Body:
//	 - body *WSGProfileCreateL3AccessControlPolicy
func (s *WSGL3AccessControlPolicyService) AddL3AccessControlPolicies(ctx context.Context, body *WSGProfileCreateL3AccessControlPolicy) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddL3AccessControlPolicies, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteL3AccessControlPolicies
//
// Use this API command to delete Bulk L3 Access Control Policies.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL3AccessControlPolicyService) DeleteL3AccessControlPolicies(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteL3AccessControlPolicies, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteL3AccessControlPoliciesById
//
// Delete a L3 Access Control Policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL3AccessControlPolicyService) DeleteL3AccessControlPoliciesById(ctx context.Context, id string) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteL3AccessControlPoliciesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindL3AccessControlPolicies
//
// Retrieve L3 Access Control Policy list.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGL3AccessControlPolicyService) FindL3AccessControlPolicies(ctx context.Context, optionalParams map[string][]string) (*WSGProfileIdList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileIdList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindL3AccessControlPolicies, true)
	if v, ok := optionalParams["domainId"]; ok {
		req.AddQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIdList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindL3AccessControlPoliciesById
//
// Retrieve a L3 Access Control Policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL3AccessControlPolicyService) FindL3AccessControlPoliciesById(ctx context.Context, id string) (*WSGProfileL3AccessControlPolicy, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileL3AccessControlPolicy
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindL3AccessControlPoliciesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileL3AccessControlPolicy()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindL3AccessControlPoliciesByQueryCriteria
//
// Retrieve a list of L3 Access Control Policy.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL3AccessControlPolicyService) FindL3AccessControlPoliciesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileL3AccessControlPolicyArray, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileL3AccessControlPolicyArray
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindL3AccessControlPoliciesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileL3AccessControlPolicyArray()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateL3AccessControlPoliciesById
//
// Modify a L3 Access Control Policy.
//
// Request Body:
//	 - body *WSGProfileModifyL3AccessControlPolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL3AccessControlPolicyService) UpdateL3AccessControlPoliciesById(ctx context.Context, body *WSGProfileModifyL3AccessControlPolicy, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateL3AccessControlPoliciesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
