package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGIndoorMapService struct {
	apiClient *VSZClient
}

func NewWSGIndoorMapService(c *VSZClient) *WSGIndoorMapService {
	s := new(WSGIndoorMapService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIndoorMapService() *WSGIndoorMapService {
	return NewWSGIndoorMapService(ss.apiClient)
}

// WSGIndoorMapAccessPointList
//
// Definition: indoorMap_accessPointList
type WSGIndoorMapAccessPointList []*WSGIndoorMapAp

func MakeWSGIndoorMapAccessPointList() WSGIndoorMapAccessPointList {
	m := make(WSGIndoorMapAccessPointList, 0)
	return m
}

// WSGIndoorMapBasicIndoorMap
//
// Definition: indoorMap_basicIndoorMap
type WSGIndoorMapBasicIndoorMap struct {
	Address *string `json:"address,omitempty"`

	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *WSGCommonDescription `json:"description"`

	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType"`

	Id *string `json:"id,omitempty"`

	ImageFileName *string `json:"imageFileName,omitempty"`

	Latitude *float64 `json:"latitude,omitempty"`

	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Orientation
	// Constraints:
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty"`

	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGIndoorMapBasicIndoorMap() *WSGIndoorMapBasicIndoorMap {
	m := new(WSGIndoorMapBasicIndoorMap)
	return m
}

// WSGIndoorMapIndooMapAuditId
//
// Definition: indoorMap_indooMapAuditId
type WSGIndoorMapIndooMapAuditId struct {
	// Id
	// the identifier of the indoor map
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the indoor map
	Name *string `json:"name,omitempty"`
}

type WSGIndoorMapIndooMapAuditIdAPIResponse struct {
	*RawAPIResponse
	Data *WSGIndoorMapIndooMapAuditId
}

func newWSGIndoorMapIndooMapAuditIdAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGIndoorMapIndooMapAuditIdAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGIndoorMapIndooMapAuditIdAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGIndoorMapIndooMapAuditId)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGIndoorMapIndooMapAuditId() *WSGIndoorMapIndooMapAuditId {
	m := new(WSGIndoorMapIndooMapAuditId)
	return m
}

// WSGIndoorMap
//
// Definition: indoorMap_indoorMap
type WSGIndoorMap struct {
	// Address
	// address
	Address *string `json:"address,omitempty"`

	// ApGroupId
	// apGroupId
	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *WSGCommonDescription `json:"description"`

	// DomainId
	// domainId
	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType"`

	// Id
	// id
	Id *string `json:"id,omitempty"`

	// ImageData
	// imageData
	ImageData *string `json:"imageData,omitempty"`

	// ImageFileName
	// imageFileName
	ImageFileName *string `json:"imageFileName,omitempty"`

	// Latitude
	// latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude
	// longitude
	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Orientation
	// orientation
	// Constraints:
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty"`

	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	// TenantId
	// tenantId
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// zoneId
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGIndoorMapAPIResponse struct {
	*RawAPIResponse
	Data *WSGIndoorMap
}

func newWSGIndoorMapAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGIndoorMapAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGIndoorMapAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGIndoorMap)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGIndoorMap() *WSGIndoorMap {
	m := new(WSGIndoorMap)
	return m
}

// WSGIndoorMapAp
//
// Definition: indoorMap_indoorMapAp
type WSGIndoorMapAp struct {
	IndoorMapXy *WSGIndoorMapXy `json:"indoorMapXy,omitempty"`

	// Mac
	// the identifier of the create object
	Mac *string `json:"mac,omitempty"`
}

func NewWSGIndoorMapAp() *WSGIndoorMapAp {
	m := new(WSGIndoorMapAp)
	return m
}

// WSGIndoorMapList
//
// Definition: indoorMap_indoorMapList
type WSGIndoorMapList struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first indoorMapList returned out of the complete indoor maps list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more indoor maps after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIndoorMapBasicIndoorMap `json:"list,omitempty"`

	// TotalCount
	// Total indoor maps count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGIndoorMapListAPIResponse struct {
	*RawAPIResponse
	Data *WSGIndoorMapList
}

func newWSGIndoorMapListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGIndoorMapListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGIndoorMapListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGIndoorMapList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGIndoorMapList() *WSGIndoorMapList {
	m := new(WSGIndoorMapList)
	return m
}

// WSGIndoorMapSummary
//
// Definition: indoorMap_indoorMapSummary
type WSGIndoorMapSummary struct {
	// Address
	// address
	Address *string `json:"address,omitempty"`

	// ApCount
	// AP count in this indoor map
	ApCount *float64 `json:"apCount,omitempty"`

	// ApGroupId
	// apGroupId
	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *WSGCommonDescription `json:"description"`

	// DomainId
	// domainId
	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType"`

	// Id
	// id
	Id *string `json:"id,omitempty"`

	// ImageFileName
	// imageFileName
	ImageFileName *string `json:"imageFileName,omitempty"`

	// Key
	// id
	Key *string `json:"key,omitempty"`

	// Latitude
	// latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude
	// longitude
	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	// TenantId
	// tenantId
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// zoneId
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGIndoorMapSummary() *WSGIndoorMapSummary {
	m := new(WSGIndoorMapSummary)
	return m
}

// WSGIndoorMapSummaryList
//
// Definition: indoorMap_indoorMapSummaryList
type WSGIndoorMapSummaryList struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first indoorMapList returned out of the complete indoor maps list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more indoor maps after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGIndoorMapSummary `json:"list,omitempty"`

	// TotalCount
	// Total indoor maps count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGIndoorMapSummaryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGIndoorMapSummaryList
}

func newWSGIndoorMapSummaryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGIndoorMapSummaryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGIndoorMapSummaryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGIndoorMapSummaryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGIndoorMapSummaryList() *WSGIndoorMapSummaryList {
	m := new(WSGIndoorMapSummaryList)
	return m
}

// WSGIndoorMapXy
//
// Definition: indoorMap_indoorMapXy
type WSGIndoorMapXy struct {
	// X
	// x
	X *float64 `json:"x,omitempty"`

	// Y
	// y
	Y *float64 `json:"y,omitempty"`
}

func NewWSGIndoorMapXy() *WSGIndoorMapXy {
	m := new(WSGIndoorMapXy)
	return m
}

// WSGIndoorMapScale
//
// Definition: indoorMap_scale
type WSGIndoorMapScale struct {
	A *WSGIndoorMapXy `json:"a,omitempty"`

	B *WSGIndoorMapXy `json:"b,omitempty"`

	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	// Unit
	// unit
	// Constraints:
	//    - oneof:[MM,CM,M,Foot,Yard]
	Unit *string `json:"unit,omitempty"`
}

func NewWSGIndoorMapScale() *WSGIndoorMapScale {
	m := new(WSGIndoorMapScale)
	return m
}

// AddMaps
//
// Use this API command to create indoorMap.
//
// Operation ID: addMaps
// Operation path: /maps
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGIndoorMap
func (s *WSGIndoorMapService) AddMaps(ctx context.Context, body *WSGIndoorMap, mutators ...RequestMutator) (*WSGIndoorMapIndooMapAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIndoorMapIndooMapAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddMaps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIndoorMapIndooMapAuditIdAPIResponse), err
}

// DeleteMapsByIndoorMapId
//
// Use this API command to delete indoor map.
//
// Operation ID: deleteMapsByIndoorMapId
// Operation path: /maps/{indoorMapId}
// Success code: 204 (No Content)
//
// Required parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) DeleteMapsByIndoorMapId(ctx context.Context, indoorMapId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteMapsByIndoorMapId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("indoorMapId", indoorMapId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindMaps
//
// Use this API command to get indoor map list.
//
// Operation ID: findMaps
// Operation path: /maps
// Success code: 200 (OK)
//
// Required parameters:
// - groupId string
//		- required
// - groupType string
//		- required
func (s *WSGIndoorMapService) FindMaps(ctx context.Context, groupId string, groupType string, mutators ...RequestMutator) (*WSGIndoorMapListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIndoorMapListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindMaps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("groupId", groupId)
	req.QueryParams.Set("groupType", groupType)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIndoorMapListAPIResponse), err
}

// FindMapsByIndoorMapId
//
// Use this API command to get indoor maps.
//
// Operation ID: findMapsByIndoorMapId
// Operation path: /maps/{indoorMapId}
// Success code: 200 (OK)
//
// Required parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) FindMapsByIndoorMapId(ctx context.Context, indoorMapId string, mutators ...RequestMutator) (*WSGIndoorMapAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIndoorMapAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindMapsByIndoorMapId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("indoorMapId", indoorMapId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIndoorMapAPIResponse), err
}

// FindMapsByQueryCriteria
//
// Use this API command to query indoorMap.
//
// Operation ID: findMapsByQueryCriteria
// Operation path: /maps/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIndoorMapService) FindMapsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGIndoorMapListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIndoorMapListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindMapsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIndoorMapListAPIResponse), err
}

// FindMapsByRogueMac
//
// Use this API command to query the location of a rogue device.
//
// Operation ID: findMapsByRogueMac
// Operation path: /maps/rogues/{rogueMAC}/{apMac}
// Success code: 200 (OK)
//
// Required parameters:
// - rogueMac string
//		- required
// - apMac string
//		- required
func (s *WSGIndoorMapService) FindMapsByRogueMac(ctx context.Context, rogueMac string, apMac string, mutators ...RequestMutator) (*WSGIndoorMapAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGIndoorMapListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindMapsByRogueMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("rogueMac", rogueMac)
	req.QueryParams.Set("apMac", apMac)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGIndoorMapAPIResponse), err
}

// PartialUpdateMapsByIndoorMapId
//
// Use this API command to update specific indoor map.
//
// Operation ID: partialUpdateMapsByIndoorMapId
// Operation path: /maps/{indoorMapId}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGIndoorMap
//
// Required parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) PartialUpdateMapsByIndoorMapId(ctx context.Context, body *WSGIndoorMap, indoorMapId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateMapsByIndoorMapId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.SetBody(body)
	req.PathParams.Set("indoorMapId", indoorMapId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateMapsApsByIndoorMapId
//
// Use this API command to put Aps in indoor map.
//
// Operation ID: updateMapsApsByIndoorMapId
// Operation path: /maps/{indoorMapId}/aps
// Success code: 204 (No Content)
//
// Request body:
//	 - body WSGIndoorMapAccessPointList
//
// Required parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) UpdateMapsApsByIndoorMapId(ctx context.Context, body WSGIndoorMapAccessPointList, indoorMapId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateMapsApsByIndoorMapId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.SetBody(body)
	req.PathParams.Set("indoorMapId", indoorMapId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
