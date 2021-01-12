package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
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

func newWSGGeofenceProfileGetGeofenceProfileAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGGeofenceProfileGetGeofenceProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGGeofenceProfileGetGeofenceProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGGeofenceProfileGetGeofenceProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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

func newWSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGGeofenceProfileGetGeofenceProfileProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGGeofenceProfileGetGeofenceProfileProfileList() *WSGGeofenceProfileGetGeofenceProfileProfileList {
	m := new(WSGGeofenceProfileGetGeofenceProfileProfileList)
	return m
}

// AddRkszonesGeofenceProfilesByZoneId
//
// Use this API command to create a Geofence Profile.
//
// Operation ID: addRkszonesGeofenceProfilesByZoneId
// Operation path: /rkszones/{zoneId}/geofenceProfiles
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGGeofenceProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) AddRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGGeofenceProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesGeofenceProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesGeofenceProfilesById
//
// Use this API command to delete a Geofence Profile.
//
// Operation ID: deleteRkszonesGeofenceProfilesById
// Operation path: /rkszones/{zoneId}/geofenceProfiles/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) DeleteRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesGeofenceProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesGeofenceProfilesByZoneId
//
// Use this API command to delete a list of Geofence Profile.
//
// Operation ID: deleteRkszonesGeofenceProfilesByZoneId
// Operation path: /rkszones/{zoneId}/geofenceProfiles
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) DeleteRkszonesGeofenceProfilesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesGeofenceProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesGeofenceProfilesById
//
// Use this API command to retrieve a Geofence Profile.
//
// Operation ID: findRkszonesGeofenceProfilesById
// Operation path: /rkszones/{zoneId}/geofenceProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) FindRkszonesGeofenceProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGGeofenceProfileGetGeofenceProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGGeofenceProfileGetGeofenceProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGGeofenceProfileGetGeofenceProfileAPIResponse), err
}

// FindRkszonesGeofenceProfilesByZoneId
//
// Query Geofence Profile with specified filters.
//
// Operation ID: findRkszonesGeofenceProfilesByZoneId
// Operation path: /rkszones/{zoneId}/geofenceProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) FindRkszonesGeofenceProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesGeofenceProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGGeofenceProfileGetGeofenceProfileProfileListAPIResponse), err
}

// UpdateRkszonesGeofenceProfilesById
//
// Use this API command to update a Geofence Profile.
//
// Operation ID: updateRkszonesGeofenceProfilesById
// Operation path: /rkszones/{zoneId}/geofenceProfiles/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGGeofenceProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGeofenceProfileService) UpdateRkszonesGeofenceProfilesById(ctx context.Context, body *WSGGeofenceProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesGeofenceProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
