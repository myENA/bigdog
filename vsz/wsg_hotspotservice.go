package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotExternal, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddRkszonesPortalsHotspotExternalByZoneId), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotInternal, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddRkszonesPortalsHotspotInternalByZoneId), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotSmartClientOnly, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddRkszonesPortalsHotspotSmartClientOnlyByZoneId), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGHotspotServiceService) DeleteRkszonesPortalsHotspotById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteRkszonesPortalsHotspotById), true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceHotspot, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalServiceHotspot
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindRkszonesPortalsHotspotById), true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceHotspot()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindRkszonesPortalsHotspotByZoneId
//
// Use this API command to retrieve a list of Hotspot (WISPr) of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId(ctx context.Context, zoneId string) (*WSGPortalServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindRkszonesPortalsHotspotByZoneId), true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindServicesHotspotByQueryCriteria
//
// Query Hotspot Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspotServiceService) FindServicesHotspotByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindServicesHotspotByQueryCriteria), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesPortalsHotspotById
//
// Use this API command to modify the configuration on Hotspot (WISPr) of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceModifyHotspot
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGHotspotServiceService) PartialUpdateRkszonesPortalsHotspotById(ctx context.Context, body *WSGPortalServiceModifyHotspot, id string, zoneId string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateRkszonesPortalsHotspotById), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
