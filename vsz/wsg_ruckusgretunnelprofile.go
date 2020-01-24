package vsz

// API Version: v8_1

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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesTunnelRuckusgre, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesTunnelRuckusgre
//
// Use this API command to delete multiple RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRuckusGRETunnelProfileService) DeleteProfilesTunnelRuckusgre(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelRuckusgre, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
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
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelRuckusgreById, true)
}

// FindProfilesTunnelRuckusgre
//
// Use this API command to retrieve a list of RuckusGRE tunnel profile.
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgre(ctx context.Context) (*WSGProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelRuckusgre, true)
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
		req  *APIRequest
		resp *WSGProfileRuckusGREProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelRuckusgreById, true)
}

// FindProfilesTunnelRuckusgreByQueryCriteria
//
// Use this API command to query a list of RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRuckusGRETunnelProfileService) FindProfilesTunnelRuckusgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileRuckusGREProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileRuckusGREProfileList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTunnelRuckusgreByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateProfilesTunnelRuckusgreById
//
// Use this API command to modify the basic information of RuckusGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileModifyRuckusGREProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusGRETunnelProfileService) PartialUpdateProfilesTunnelRuckusgreById(ctx context.Context, body *WSGProfileModifyRuckusGREProfile, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelRuckusgreById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
