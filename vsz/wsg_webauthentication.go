package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGWebAuthenticationService struct {
	apiClient *APIClient
}

func NewWSGWebAuthenticationService(c *APIClient) *WSGWebAuthenticationService {
	s := new(WSGWebAuthenticationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWebAuthenticationService() *WSGWebAuthenticationService {
	return NewWSGWebAuthenticationService(ss.apiClient)
}

// AddRkszonesPortalsWebauthByZoneId
//
// Use this API command to create a new web authentication of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceCreateWebAuthentication
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) AddRkszonesPortalsWebauthByZoneId(ctx context.Context, body *WSGPortalServiceCreateWebAuthentication, zoneId string) (*WSGCommonCreateResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsWebauthByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteRkszonesPortalsWebauthById
//
// Use this API command to delete an web authentication of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) DeleteRkszonesPortalsWebauthById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalsWebauthById, true)
}

// DeleteRkszonesPortalsWebauthRedirectById
//
// Use this API command to set redirect to the URL that user intends to visit on web authentication of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) DeleteRkszonesPortalsWebauthRedirectById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalsWebauthRedirectById, true)
}

// FindRkszonesPortalsWebauthById
//
// Use this API command to retrieve a web authentication of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceWebAuthentication, error) {
	var (
		req  *APIRequest
		resp *WSGPortalServiceWebAuthentication
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsWebauthById, true)
}

// FindRkszonesPortalsWebauthByZoneId
//
// Use this API command to retrieve a list of web authentication of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) FindRkszonesPortalsWebauthByZoneId(ctx context.Context, zoneId string) (*WSGPortalServiceList, error) {
	var (
		req  *APIRequest
		resp *WSGPortalServiceList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsWebauthByZoneId, true)
}

// PartialUpdateRkszonesPortalsWebauthById
//
// Use this API command to modify the basic information on web authentication of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceModifyWebAuthentication
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWebAuthenticationService) PartialUpdateRkszonesPortalsWebauthById(ctx context.Context, body *WSGPortalServiceModifyWebAuthentication, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalsWebauthById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
