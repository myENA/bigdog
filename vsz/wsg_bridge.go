package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGBridgeService struct {
	apiClient *APIClient
}

func NewWSGBridgeService(c *APIClient) *WSGBridgeService {
	s := new(WSGBridgeService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBridgeService() *WSGBridgeService {
	return NewWSGBridgeService(ss.apiClient)
}

// AddProfilesBridge
//
// Use this API command to create Bridge profile.
//
// Request Body:
//	 - body *WSGProfileCreateBridgeProfile
func (s *WSGBridgeService) AddProfilesBridge(ctx context.Context, body *WSGProfileCreateBridgeProfile) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesBridge, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesBridge
//
// Use this API command to delete multiple bridge profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBridgeService) DeleteProfilesBridge(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesBridge, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesBridgeById
//
// Use this API command to delete Bridge profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBridgeService) DeleteProfilesBridgeById(ctx context.Context, id string) error {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesBridgeById, true)
	req.SetPathParameter("id", id)
}

// FindProfilesBridge
//
// Use this API command to retrieve a list of Bridge profile.
func (s *WSGBridgeService) FindProfilesBridge(ctx context.Context) (*WSGProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesBridge, true)
}

// FindProfilesBridgeById
//
// Use this API command to retrieve Bridge profile by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBridgeService) FindProfilesBridgeById(ctx context.Context, id string) (*WSGProfileBridgeProfile, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBridgeProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesBridgeById, true)
	req.SetPathParameter("id", id)
}

// FindProfilesBridgeByQueryCriteria
//
// Use this API command to query a list of Bridge profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBridgeService) FindProfilesBridgeByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBridgeProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBridgeProfileList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesBridgeByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateProfilesBridgeById
//
// Use this API command to modify the basic information of Bridge profile.
//
// Request Body:
//	 - body *WSGProfileModifyBridgeProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBridgeService) PartialUpdateProfilesBridgeById(ctx context.Context, body *WSGProfileModifyBridgeProfile, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesBridgeById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
}
