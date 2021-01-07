package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGGeofenceProfileService struct {
	apiClient *VSZClient
}

func NewWSGGeofenceProfileService(c *VSZClient) *WSGGeofenceProfileService {
	s := new(WSGGeofenceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGeofenceProfileService() *WSGGeofenceProfileService {
	return NewWSGGeofenceProfileService(ss.apiClient)
}

// WSGGeofenceProfile
//
// Definition: geofenceProfile_geofenceProfile
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

// WSGGeofenceProfileLocationData
//
// Definition: geofenceProfile_geofenceProfileLocationData
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

// WSGGeofenceProfileGetGeofenceProfile
//
// Definition: geofenceProfile_getGeofenceProfile
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

type WSGGeofenceProfileGetGeofenceProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGGeofenceProfileGetGeofenceProfile
}

func newWSGGeofenceProfileGetGeofenceProfileAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGGeofenceProfileGetGeofenceProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGGeofenceProfileGetGeofenceProfileAPIResponse) Hydrate() error {
	r.Data = new(WSGGeofenceProfileGetGeofenceProfile)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGGeofenceProfileGetGeofenceProfile() *WSGGeofenceProfileGetGeofenceProfile {
	m := new(WSGGeofenceProfileGetGeofenceProfile)
	return m
}

// WSGGeofenceProfileGetGeofenceProfileProfileList
//
// Definition: geofenceProfile_getGeofenceProfileProfileList
type WSGGeofenceProfileGetGeofenceProfileProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGGeofenceProfileGetGeofenceProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGGeofenceProfileGetGeofenceProfileProfileList
}

func newWSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse) Hydrate() error {
	r.Data = new(WSGGeofenceProfileGetGeofenceProfileProfileList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGGeofenceProfileGetGeofenceProfileProfileList() *WSGGeofenceProfileGetGeofenceProfileProfileList {
	m := new(WSGGeofenceProfileGetGeofenceProfileProfileList)
	return m
}

// AddRkszonesGeofenceProfilesByZoneId
//
// Operation ID: addRkszonesGeofenceProfilesByZoneId
//
// Use this API command to create a Geofence Profile.
//
// Request Body:
//	 - body *WSGGeofenceProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) AddRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGGeofenceProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesGeofenceProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesGeofenceProfilesById
//
// Operation ID: deleteRkszonesGeofenceProfilesById
//
// Use this API command to delete a Geofence Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) DeleteRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesGeofenceProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesGeofenceProfilesByZoneId
//
// Operation ID: deleteRkszonesGeofenceProfilesByZoneId
//
// Use this API command to delete a list of Geofence Profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) DeleteRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesGeofenceProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesGeofenceProfilesById
//
// Operation ID: findRkszonesGeofenceProfilesById
//
// Use this API command to retrieve a Geofence Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) FindRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGGeofenceProfileGetGeofenceProfileAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGGeofenceProfileGetGeofenceProfile()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesGeofenceProfilesByZoneId
//
// Operation ID: findRkszonesGeofenceProfilesByZoneId
//
// Query Geofence Profile with specified filters.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) FindRkszonesGeofenceProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGGeofenceProfileGetGeofenceProfileProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesGeofenceProfilesById
//
// Operation ID: updateRkszonesGeofenceProfilesById
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
func (s *WSGGeofenceProfileService) UpdateRkszonesGeofenceProfilesById(ctx context.Context, body *WSGGeofenceProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesGeofenceProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
