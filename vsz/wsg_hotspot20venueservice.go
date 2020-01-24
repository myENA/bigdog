package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGHotspot20VenueServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspot20VenueServiceService(c *APIClient) *WSGHotspot20VenueServiceService {
	s := new(WSGHotspot20VenueServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20VenueServiceService() *WSGHotspot20VenueServiceService {
	return NewWSGHotspot20VenueServiceService(ss.apiClient)
}

// AddRkszonesHs20VenuesByZoneId
//
// Use this API command to create a new Hotspot 2.0 venue profile of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspot20VenueProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) AddRkszonesHs20VenuesByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspot20VenueProfile, zoneId string) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesHs20VenuesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteRkszonesHs20VenuesById
//
// Use this API command to delete Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) DeleteRkszonesHs20VenuesById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesHs20VenuesById, true)
}

// FindRkszonesHs20VenuesById
//
// Use this API command to retrieve a Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceHotspot20VeuneProfile, error) {
	var (
		req  *APIRequest
		resp *WSGPortalServiceHotspot20VeuneProfile
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesHs20VenuesById, true)
}

// FindRkszonesHs20VenuesByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesByZoneId(ctx context.Context, zoneId string) (*WSGPortalServiceList, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesHs20VenuesByZoneId, true)
}

// PartialUpdateRkszonesHs20VenuesById
//
// Use this API command to modify the basic information on Hotspot 2.0 venue profile of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceModifyHotspot20VenueProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) PartialUpdateRkszonesHs20VenuesById(ctx context.Context, body *WSGPortalServiceModifyHotspot20VenueProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesHs20VenuesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
