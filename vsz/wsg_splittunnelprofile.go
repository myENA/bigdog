package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGSplitTunnelProfileService struct {
	apiClient *APIClient
}

func NewWSGSplitTunnelProfileService(c *APIClient) *WSGSplitTunnelProfileService {
	s := new(WSGSplitTunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSplitTunnelProfileService() *WSGSplitTunnelProfileService {
	return NewWSGSplitTunnelProfileService(ss.apiClient)
}

// AddRkszonesSplitTunnelProfilesByZoneId
//
// Create a split tunnel profile.
//
// Request Body:
//	 - body *WSGSplitTunnelCreateSplitTunnelProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) AddRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, zoneId string) (*WSGCommonCreateResult, error) {
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesSplitTunnelProfilesByZoneId, true)
}

// DeleteRkszonesSplitTunnelProfilesById
//
// Use this API command to delete a split tunnel profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) DeleteRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesSplitTunnelProfilesById, true)
}

// FindRkszonesSplitTunnelProfilesById
//
// Get a split tunnel profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string) (*WSGSplitTunnelProfile, error) {
	var (
		resp *WSGSplitTunnelProfile
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesSplitTunnelProfilesById, true)
}

// FindRkszonesSplitTunnelProfilesByQueryCriteria
//
// Use this API command to retrieve a list of split tunnel profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSplitTunnelProfileQuery, error) {
	var (
		resp *WSGSplitTunnelProfileQuery
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesSplitTunnelProfilesByQueryCriteria, true)
}

// FindRkszonesSplitTunnelProfilesByZoneId
//
// Get a ID list of split tunnel profile in this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, zoneId string) (*WSGSplitTunnelProfileList, error) {
	var (
		resp *WSGSplitTunnelProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesSplitTunnelProfilesByZoneId, true)
}

// PartialUpdateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify a split tunnel profile.
//
// Request Body:
//	 - body *WSGSplitTunnelModifySplitTunnelProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) PartialUpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelModifySplitTunnelProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesSplitTunnelProfilesById, true)
}

// UpdateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify entire information of a split tunnel profile.
//
// Request Body:
//	 - body *WSGSplitTunnelCreateSplitTunnelProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) UpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesSplitTunnelProfilesById, true)
}
