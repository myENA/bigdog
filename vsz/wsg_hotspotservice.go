package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGHotspotServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspotServiceService(c *APIClient) *WSGHotspotServiceService {
	s := new(WSGHotspotServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspotServiceService() *WSGHotspotServiceService {
	return NewWSGHotspotServiceService(ss.apiClient)
}

// AddRkszonesPortalsHotspotExternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with external logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotExternal
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotExternal, zoneId string) (*WSGCommonCreateResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsHotspotExternalByZoneId, true)
}

// AddRkszonesPortalsHotspotInternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with internal logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotInternal
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotInternal, zoneId string) (*WSGCommonCreateResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsHotspotInternalByZoneId, true)
}

// AddRkszonesPortalsHotspotSmartClientOnlyByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with smart client only of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotSmartClientOnly
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotSmartClientOnly, zoneId string) (*WSGCommonCreateResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesPortalsHotspotSmartClientOnlyByZoneId, true)
}

// DeleteRkszonesPortalsHotspotById
//
// Use this API command to delete a Hotspot (WISPr) of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) DeleteRkszonesPortalsHotspotById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesPortalsHotspotById, true)
}

// FindRkszonesPortalsHotspotById
//
// Use this API command to retrieve a Hotspot (WISPr) of zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceHotspot, error) {
	var (
		resp *WSGPortalServiceHotspot
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsHotspotById, true)
}

// FindRkszonesPortalsHotspotByZoneId
//
// Use this API command to retrieve a list of Hotspot (WISPr) of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId(ctx context.Context, zoneId string) (*WSGPortalServiceList, error) {
	var (
		resp *WSGPortalServiceList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesPortalsHotspotByZoneId, true)
}

// PartialUpdateRkszonesPortalsHotspotById
//
// Use this API command to modify the basic information on Hotspot (WISPr) of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceModifyHotspot
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) PartialUpdateRkszonesPortalsHotspotById(ctx context.Context, body *WSGPortalServiceModifyHotspot, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalsHotspotById, true)
}
