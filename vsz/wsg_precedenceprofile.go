package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGPrecedenceProfileService struct {
	apiClient *APIClient
}

func NewWSGPrecedenceProfileService(c *APIClient) *WSGPrecedenceProfileService {
	s := new(WSGPrecedenceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGPrecedenceProfileService() *WSGPrecedenceProfileService {
	return NewWSGPrecedenceProfileService(ss.apiClient)
}

// AddPrecedence
//
// Use this API command to create Precedence Profile.
//
// Request Body:
//	 - body *WSGProfileCreatePrecedenceProfile
func (s *WSGPrecedenceProfileService) AddPrecedence(ctx context.Context, body *WSGProfileCreatePrecedenceProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddPrecedence), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeletePrecedence
//
// Use this API command to Bulk Delete Precedence Profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkPrecedenceProfile
func (s *WSGPrecedenceProfileService) DeletePrecedence(ctx context.Context, body *WSGProfileDeleteBulkPrecedenceProfile) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeletePrecedence), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// DeletePrecedenceById
//
// Use this API command to Delete Precedence Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) DeletePrecedenceById(ctx context.Context, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeletePrecedenceById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// FindPrecedence
//
// Use this API command to Get Precedence Profile list.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGPrecedenceProfileService) FindPrecedence(ctx context.Context, optionalParams map[string][]string) (*WSGProfilePrecedenceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfilePrecedenceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindPrecedence), true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfilePrecedenceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindPrecedenceById
//
// Use this API command to Get Precedence Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) FindPrecedenceById(ctx context.Context, id string) (*WSGProfileCreatePrecedenceProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileCreatePrecedenceProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindPrecedenceById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileCreatePrecedenceProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindPrecedenceByQueryCriteria
//
// Use this API command to query Precedence Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGPrecedenceProfileService) FindPrecedenceByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfilePrecedenceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfilePrecedenceList
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindPrecedenceByQueryCriteria), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfilePrecedenceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdatePrecedenceById
//
// Use this API command to Modify Precedence Profile by profile's ID.
//
// Request Body:
//	 - body *WSGProfileUpdatePrecedenceProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGPrecedenceProfileService) PartialUpdatePrecedenceById(ctx context.Context, body *WSGProfileUpdatePrecedenceProfile, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdatePrecedenceById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
