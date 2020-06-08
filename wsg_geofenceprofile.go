package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGGeofenceprofileService struct {
	apiClient *VSZClient
}

func NewWSGGeofenceprofileService(c *VSZClient) *WSGGeofenceprofileService {
	s := new(WSGGeofenceprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGeofenceprofileService() *WSGGeofenceprofileService {
	return NewWSGGeofenceprofileService(ss.apiClient)
}

type WSGGeofenceprofile struct {
	// Description
	// Geofence Profile's description
	Description *string `json:"description,omitempty"`

	// LocationList
	// Geofence Profile's location list
	// Constraints:
	//    - required
	LocationList []*WSGGeofenceprofileLocationData `json:"locationList"`

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

func NewWSGGeofenceprofile() *WSGGeofenceprofile {
	m := new(WSGGeofenceprofile)
	return m
}

type WSGGeofenceprofileLocationData struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Location
	// Geofence Profile's location Latitude and longitude
	Location *string `json:"location,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGGeofenceprofileLocationData() *WSGGeofenceprofileLocationData {
	m := new(WSGGeofenceprofileLocationData)
	return m
}

type WSGGeofenceprofileGetGeofenceProfile struct {
	// Description
	// Geofence Profile's description
	Description *string `json:"description,omitempty"`

	// Id
	// Geofence Profile's id
	Id *string `json:"id,omitempty"`

	// LocationList
	// Geofence Profile's location list
	LocationList []*WSGGeofenceprofileLocationData `json:"locationList,omitempty"`

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

func NewWSGGeofenceprofileGetGeofenceProfile() *WSGGeofenceprofileGetGeofenceProfile {
	m := new(WSGGeofenceprofileGetGeofenceProfile)
	return m
}

type WSGGeofenceprofileGetGeofenceProfileProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGGeofenceprofileGetGeofenceProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGGeofenceprofileGetGeofenceProfileProfileList() *WSGGeofenceprofileGetGeofenceProfileProfileList {
	m := new(WSGGeofenceprofileGetGeofenceProfileProfileList)
	return m
}

// AddRkszonesGeofenceProfilesByZoneId
//
// Use this API command to create a Geofence Profile.
//
// Request Body:
//	 - body *WSGGeofenceprofile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceprofileService) AddRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGGeofenceprofile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
func (s *WSGGeofenceprofileService) DeleteRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
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
func (s *WSGGeofenceprofileService) DeleteRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
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
func (s *WSGGeofenceprofileService) FindRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string) (*WSGGeofenceprofileGetGeofenceProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGGeofenceprofileGetGeofenceProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGGeofenceprofileGetGeofenceProfile()
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
func (s *WSGGeofenceprofileService) FindRkszonesGeofenceProfilesByZoneId(ctx context.Context, zoneId string) (*WSGGeofenceprofileGetGeofenceProfileProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGGeofenceprofileGetGeofenceProfileProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGGeofenceprofileGetGeofenceProfileProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesGeofenceProfilesById
//
// Use this API command to update a Geofence Profile.
//
// Request Body:
//	 - body *WSGGeofenceprofile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceprofileService) UpdateRkszonesGeofenceProfilesById(ctx context.Context, body *WSGGeofenceprofile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
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
