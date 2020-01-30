package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAccountingProfileService struct {
	apiClient *APIClient
}

func NewWSGAccountingProfileService(c *APIClient) *WSGAccountingProfileService {
	s := new(WSGAccountingProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountingProfileService() *WSGAccountingProfileService {
	return NewWSGAccountingProfileService(ss.apiClient)
}

// AddProfilesAcct
//
// Use this API command to create a new accounting profile.
//
// Request Body:
//	 - body *WSGProfileCreateAccountingProfile
func (s *WSGAccountingProfileService) AddProfilesAcct(ctx context.Context, body *WSGProfileCreateAccountingProfile) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAcct, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddProfilesAcctCloneById
//
// Use this API command to clone an accounting profile.
//
// Request Body:
//	 - body *WSGProfileClone
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) AddProfilesAcctCloneById(ctx context.Context, body *WSGProfileClone, id string) (*WSGProfileClone, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileClone
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAcctCloneById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileClone()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// DeleteProfilesAcct
//
// Use this API command to delete a list of accounting profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkAccountingProfile
func (s *WSGAccountingProfileService) DeleteProfilesAcct(ctx context.Context, body *WSGProfileDeleteBulkAccountingProfile) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAcct, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteProfilesAcctById
//
// Use this API command to delete an accounting profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) DeleteProfilesAcctById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAcctById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindProfilesAcct
//
// Use this API command to retrieve a list of accounting profiles.
func (s *WSGAccountingProfileService) FindProfilesAcct(ctx context.Context) (*WSGProfileAccountingProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileAccountingProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAcct, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileAccountingProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindProfilesAcctById
//
// Use this API command to retrieve an accounting profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) FindProfilesAcctById(ctx context.Context, id string) (*WSGProfileAccountingProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileAccountingProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAcctById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileAccountingProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindProfilesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting profiles by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingProfileService) FindProfilesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileAccountingProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileAccountingProfileList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesAcctByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileAccountingProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateProfilesAcctById
//
// Use this API command to modify the configuration of an accounting profile.
//
// Request Body:
//	 - body *WSGProfileModifyAccountingProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) PartialUpdateProfilesAcctById(ctx context.Context, body *WSGProfileModifyAccountingProfile, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesAcctById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
