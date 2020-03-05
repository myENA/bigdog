package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
func (s *WSGSplitTunnelProfileService) AddRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddRkszonesSplitTunnelProfilesByZoneId), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGSplitTunnelProfileService) DeleteRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteRkszonesSplitTunnelProfilesById), true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesById(ctx context.Context, id string, zoneId string) (*WSGSplitTunnelProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSplitTunnelProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindRkszonesSplitTunnelProfilesById), true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSplitTunnelProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindRkszonesSplitTunnelProfilesByQueryCriteria
//
// Use this API command to retrieve a list of split tunnel profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSplitTunnelProfileQuery, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSplitTunnelProfileQuery
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindRkszonesSplitTunnelProfilesByQueryCriteria), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSplitTunnelProfileQuery()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindRkszonesSplitTunnelProfilesByZoneId
//
// Get a ID list of split tunnel profile in this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, zoneId string) (*WSGSplitTunnelProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSplitTunnelProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindRkszonesSplitTunnelProfilesByZoneId), true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSplitTunnelProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGSplitTunnelProfileService) PartialUpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelModifySplitTunnelProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateRkszonesSplitTunnelProfilesById), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGSplitTunnelProfileService) UpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *WSGSplitTunnelCreateSplitTunnelProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGUpdateRkszonesSplitTunnelProfilesById), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
