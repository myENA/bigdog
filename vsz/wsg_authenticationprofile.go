package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGAuthenticationProfileService struct {
	apiClient *APIClient
}

func NewWSGAuthenticationProfileService(c *APIClient) *WSGAuthenticationProfileService {
	s := new(WSGAuthenticationProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAuthenticationProfileService() *WSGAuthenticationProfileService {
	return NewWSGAuthenticationProfileService(ss.apiClient)
}

// AddProfilesAuth
//
// Use this API command to create a new authentication profile.
//
// Request Body:
//	 - body *WSGProfileCreateAuthenticationProfile
func (s *WSGAuthenticationProfileService) AddProfilesAuth(ctx context.Context, body *WSGProfileCreateAuthenticationProfile) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAuth, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddProfilesAuthCloneById
//
// Use this API command to clone an authentication profile.
//
// Request Body:
//	 - body *WSGProfileCloneRequest
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) AddProfilesAuthCloneById(ctx context.Context, body *WSGProfileCloneRequest, id string) (*WSGProfileCloneResponse, error) {
	var (
		req  *APIRequest
		resp *WSGProfileCloneResponse
		err  error
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesAuthCloneById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesAuth
//
// Use this API command to delete a list of authentication profile.
//
// Request Body:
//	 - body *WSGProfileDeleteBulkAuthenticationProfile
func (s *WSGAuthenticationProfileService) DeleteProfilesAuth(ctx context.Context, body *WSGProfileDeleteBulkAuthenticationProfile) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAuth, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesAuthById
//
// Use this API command to delete an authentication profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) DeleteProfilesAuthById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesAuthById, true)
}

// FindProfilesAuth
//
// Use this API command to retrieve a list of authentication profiles.
func (s *WSGAuthenticationProfileService) FindProfilesAuth(ctx context.Context) (*WSGProfileAuthenticationProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileAuthenticationProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAuth, true)
}

// FindProfilesAuthAuthorizationList
//
// Use this API command to retrieve a list of authorization profiles.
//
// Required Parameters:
// - type_ string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthorizationList(ctx context.Context, type_ string) (*WSGProfileBaseServiceInfoList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBaseServiceInfoList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, type_, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAuthAuthorizationList, true)
}

// FindProfilesAuthAuthServiceListByQueryCriteria
//
// Use this API command to retrieve a list of authentication service.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthAuthServiceListByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBaseServiceInfoList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBaseServiceInfoList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesAuthAuthServiceListByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// FindProfilesAuthById
//
// Use this API command to retrieve an authentication profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) FindProfilesAuthById(ctx context.Context, id string) (*WSGProfileAuthenticationProfile, error) {
	var (
		req  *APIRequest
		resp *WSGProfileAuthenticationProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesAuthById, true)
}

// FindProfilesAuthByQueryCriteria
//
// Use this API command to retrieve a list of authentication profiles by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationProfileService) FindProfilesAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileAuthenticationProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileAuthenticationProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesAuthByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateProfilesAuthById
//
// Use this API command to modify the basic information of an authentication profile.
//
// Request Body:
//	 - body *WSGProfileModifyAuthenticationProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationProfileService) PartialUpdateProfilesAuthById(ctx context.Context, body *WSGProfileModifyAuthenticationProfile, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesAuthById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
