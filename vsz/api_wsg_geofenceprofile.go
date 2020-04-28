package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGGeofenceProfileService struct {
	apiClient *APIClient
}

func NewWSGGeofenceProfileService(c *APIClient) *WSGGeofenceProfileService {
	s := new(WSGGeofenceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGeofenceProfileService() *WSGGeofenceProfileService {
	return NewWSGGeofenceProfileService(ss.apiClient)
}

type WSGGeofenceProfile struct {
	// Description
	// Geofence Profile's description
	Description *string `json:"description,omitempty"`

	// LocationList
	// Geofence Profile's location list
	// Constraints:
	//    - required
	LocationList []*WSGGeofenceProfileLocationData `json:"locationList"`

	// Name
	// Geofence Profile's name
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// RadiusMeter
	// Geofence Profile's radius (1 - 100) meter
	// Constraints:
	//    - required
	RadiusMeter *int `json:"radiusMeter"`
}

func NewWSGGeofenceProfile() *WSGGeofenceProfile {
	m := new(WSGGeofenceProfile)
	return m
}

type WSGGeofenceProfileLocationData struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Location
	// Geofence Profile's location Latitude and longitude
	Location *string `json:"location,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGGeofenceProfileLocationData() *WSGGeofenceProfileLocationData {
	m := new(WSGGeofenceProfileLocationData)
	return m
}

type WSGGeofenceProfileGetGeofenceProfile struct {
	// Description
	// Geofence Profile's description
	Description *string `json:"description,omitempty"`

	// Id
	// Geofence Profile's id
	Id *string `json:"id,omitempty"`

	// LocationList
	// Geofence Profile's location list
	LocationList []*WSGGeofenceProfileLocationData `json:"locationList,omitempty"`

	// Name
	// Geofence Profile's name
	Name *string `json:"name,omitempty"`

	// RadiusMeter
	// Geofence Profile's radius
	RadiusMeter *int `json:"radiusMeter,omitempty"`

	// ZoneId
	// The zone which Geofence Profile belong to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGGeofenceProfileGetGeofenceProfile() *WSGGeofenceProfileGetGeofenceProfile {
	m := new(WSGGeofenceProfileGetGeofenceProfile)
	return m
}

type WSGGeofenceProfileGetGeofenceProfileProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGGeofenceProfileGetGeofenceProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGGeofenceProfileGetGeofenceProfileProfileList() *WSGGeofenceProfileGetGeofenceProfileProfileList {
	m := new(WSGGeofenceProfileGetGeofenceProfileProfileList)
	return m
}

// AddRkszonesGeofenceProfilesByZoneId
//
// Use this API command to create a Geofence Profile.
//
// Request Body:
//	 - body *WSGGeofenceProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) AddRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGGeofenceProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesGeofenceProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesGeofenceProfilesById
//
// Use this API command to delete a Geofence Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) DeleteRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesGeofenceProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesGeofenceProfilesByZoneId
//
// Use this API command to delete a list of Geofence Profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) DeleteRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesGeofenceProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesGeofenceProfilesById
//
// Use this API command to retrieve a Geofence Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) FindRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string) (*WSGGeofenceProfileGetGeofenceProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGGeofenceProfileGetGeofenceProfile
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGGeofenceProfileGetGeofenceProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesGeofenceProfilesByZoneId
//
// Query Geofence Profile with specified filters.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) FindRkszonesGeofenceProfilesByZoneId(ctx context.Context, zoneId string) (*WSGGeofenceProfileGetGeofenceProfileProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGGeofenceProfileGetGeofenceProfileProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGGeofenceProfileGetGeofenceProfileProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesGeofenceProfilesById
//
// Use this API command to update a Geofence Profile.
//
// Request Body:
//	 - body *WSGGeofenceProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) UpdateRkszonesGeofenceProfilesById(ctx context.Context, body *WSGGeofenceProfile, id string, zoneId string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesGeofenceProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
