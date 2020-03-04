package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGLBSprofileService struct {
	apiClient *APIClient
}

func NewWSGLBSprofileService(c *APIClient) *WSGLBSprofileService {
	s := new(WSGLBSprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGLBSprofileService() *WSGLBSprofileService {
	return NewWSGLBSprofileService(ss.apiClient)
}

// AddProfilesLbs
//
// Create LBS profile.
//
// Request Body:
//	 - body *WSGProfileLbsProfile
func (s *WSGLBSprofileService) AddProfilesLbs(ctx context.Context, body *WSGProfileLbsProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddProfilesLbs), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteProfilesLbs
//
// Delete multiple LBS profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGLBSprofileService) DeleteProfilesLbs(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteProfilesLbs), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteProfilesLbsById
//
// Delete LBS profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGLBSprofileService) DeleteProfilesLbsById(ctx context.Context, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteProfilesLbsById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindProfilesLbsById
//
// Retrieve LBS profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGLBSprofileService) FindProfilesLbsById(ctx context.Context, id string) (*WSGProfileLbsProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileLbsProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindProfilesLbsById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileLbsProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindProfilesLbsByQueryCriteria
//
// Query LBS profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGLBSprofileService) FindProfilesLbsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileLbsProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileLbsProfileList
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindProfilesLbsByQueryCriteria), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileLbsProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesLbsById
//
// Update LBS profile.
//
// Request Body:
//	 - body *WSGProfileLbsProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGLBSprofileService) PartialUpdateProfilesLbsById(ctx context.Context, body *WSGProfileLbsProfile, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateProfilesLbsById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
