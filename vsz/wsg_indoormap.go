package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGIndoorMapService struct {
	apiClient *APIClient
}

func NewWSGIndoorMapService(c *APIClient) *WSGIndoorMapService {
	s := new(WSGIndoorMapService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIndoorMapService() *WSGIndoorMapService {
	return NewWSGIndoorMapService(ss.apiClient)
}

type WSGIndoorMapAccessPointList []*WSGIndoorMapAp

func MakeWSGIndoorMapAccessPointList() WSGIndoorMapAccessPointList {
	m := make(WSGIndoorMapAccessPointList, 0)
	return m
}

type WSGIndoorMapBasicIndoorMap struct {
	// Address
	// Constraints:
	//    - nullable
	Address *string `json:"address,omitempty"`

	// ApGroupId
	// Constraints:
	//    - nullable
	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *WSGCommonDescription `json:"description" validate:"required"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ImageFileName
	// Constraints:
	//    - nullable
	ImageFileName *string `json:"imageFileName,omitempty"`

	// Latitude
	// Constraints:
	//    - nullable
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude
	// Constraints:
	//    - nullable
	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Orientation
	// Constraints:
	//    - nullable
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty" validate:"omitempty,oneof=HORIZONTAL VERTICAL"`

	// Scale
	// Constraints:
	//    - nullable
	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGIndoorMapBasicIndoorMap() *WSGIndoorMapBasicIndoorMap {
	m := new(WSGIndoorMapBasicIndoorMap)
	return m
}

type WSGIndoorMapIndooMapAuditId struct {
	// Id
	// the identifier of the indoor map
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the indoor map
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGIndoorMapIndooMapAuditId() *WSGIndoorMapIndooMapAuditId {
	m := new(WSGIndoorMapIndooMapAuditId)
	return m
}

type WSGIndoorMap struct {
	// Address
	// address
	// Constraints:
	//    - nullable
	Address *string `json:"address,omitempty"`

	// ApGroupId
	// apGroupId
	// Constraints:
	//    - nullable
	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *WSGCommonDescription `json:"description" validate:"required"`

	// DomainId
	// domainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

	// Id
	// id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ImageData
	// imageData
	// Constraints:
	//    - nullable
	ImageData *string `json:"imageData,omitempty"`

	// ImageFileName
	// imageFileName
	// Constraints:
	//    - nullable
	ImageFileName *string `json:"imageFileName,omitempty"`

	// Latitude
	// latitude
	// Constraints:
	//    - nullable
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude
	// longitude
	// Constraints:
	//    - nullable
	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Orientation
	// orientation
	// Constraints:
	//    - nullable
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty" validate:"omitempty,oneof=HORIZONTAL VERTICAL"`

	// Scale
	// Constraints:
	//    - nullable
	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	// TenantId
	// tenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// zoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGIndoorMap() *WSGIndoorMap {
	m := new(WSGIndoorMap)
	return m
}

type WSGIndoorMapAp struct {
	// IndoorMapXy
	// Constraints:
	//    - nullable
	IndoorMapXy *WSGIndoorMapXy `json:"indoorMapXy,omitempty"`

	// Mac
	// the identifier of the create object
	// Constraints:
	//    - nullable
	Mac *string `json:"mac,omitempty"`
}

func NewWSGIndoorMapAp() *WSGIndoorMapAp {
	m := new(WSGIndoorMapAp)
	return m
}

type WSGIndoorMapList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first indoorMapList returned out of the complete indoor maps list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more indoor maps after the list that is currently displayed
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGIndoorMapBasicIndoorMap `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total indoor maps count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIndoorMapList() *WSGIndoorMapList {
	m := new(WSGIndoorMapList)
	return m
}

type WSGIndoorMapSummary struct {
	// Address
	// address
	// Constraints:
	//    - nullable
	Address *string `json:"address,omitempty"`

	// ApCount
	// AP count in this indoor map
	// Constraints:
	//    - nullable
	ApCount *float64 `json:"apCount,omitempty"`

	// ApGroupId
	// apGroupId
	// Constraints:
	//    - nullable
	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *WSGCommonDescription `json:"description" validate:"required"`

	// DomainId
	// domainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

	// Id
	// id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ImageFileName
	// imageFileName
	// Constraints:
	//    - nullable
	ImageFileName *string `json:"imageFileName,omitempty"`

	// Key
	// id
	// Constraints:
	//    - nullable
	Key *string `json:"key,omitempty"`

	// Latitude
	// latitude
	// Constraints:
	//    - nullable
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude
	// longitude
	// Constraints:
	//    - nullable
	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Scale
	// Constraints:
	//    - nullable
	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	// TenantId
	// tenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// zoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGIndoorMapSummary() *WSGIndoorMapSummary {
	m := new(WSGIndoorMapSummary)
	return m
}

type WSGIndoorMapSummaryList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first indoorMapList returned out of the complete indoor maps list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more indoor maps after the list that is currently displayed
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGIndoorMapSummary `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total indoor maps count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGIndoorMapSummaryList() *WSGIndoorMapSummaryList {
	m := new(WSGIndoorMapSummaryList)
	return m
}

type WSGIndoorMapXy struct {
	// X
	// x
	// Constraints:
	//    - nullable
	X *float64 `json:"x,omitempty"`

	// Y
	// y
	// Constraints:
	//    - nullable
	Y *float64 `json:"y,omitempty"`
}

func NewWSGIndoorMapXy() *WSGIndoorMapXy {
	m := new(WSGIndoorMapXy)
	return m
}

type WSGIndoorMapScale struct {
	// A
	// Constraints:
	//    - nullable
	A *WSGIndoorMapXy `json:"a,omitempty"`

	// B
	// Constraints:
	//    - nullable
	B *WSGIndoorMapXy `json:"b,omitempty"`

	// Distance
	// distance
	// Constraints:
	//    - nullable
	Distance *float64 `json:"distance,omitempty"`

	// Unit
	// unit
	// Constraints:
	//    - nullable
	//    - oneof:[MM,CM,M,Foot,Yard]
	Unit *string `json:"unit,omitempty" validate:"omitempty,oneof=MM CM M Foot Yard"`
}

func NewWSGIndoorMapScale() *WSGIndoorMapScale {
	m := new(WSGIndoorMapScale)
	return m
}

// AddMaps
//
// Use this API command to create indoorMap.
//
// Request Body:
//	 - body *WSGIndoorMap
func (s *WSGIndoorMapService) AddMaps(ctx context.Context, body *WSGIndoorMap) (*WSGIndoorMapIndooMapAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIndoorMapIndooMapAuditId
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddMaps, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIndoorMapIndooMapAuditId()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteMapsByIndoorMapId
//
// Use this API command to delete indoor map.
//
// Required Parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) DeleteMapsByIndoorMapId(ctx context.Context, indoorMapId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteMapsByIndoorMapId, true)
	req.SetPathParameter("indoorMapId", indoorMapId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindMaps
//
// Use this API command to get indoor map list.
//
// Required Parameters:
// - groupId string
//		- required
// - groupType string
//		- required
func (s *WSGIndoorMapService) FindMaps(ctx context.Context, groupId string, groupType string) (*WSGIndoorMapList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIndoorMapList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupType, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindMaps, true)
	req.SetQueryParameter("groupId", []string{groupId})
	req.SetQueryParameter("groupType", []string{groupType})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIndoorMapList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindMapsByIndoorMapId
//
// Use this API command to get indoor maps.
//
// Required Parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) FindMapsByIndoorMapId(ctx context.Context, indoorMapId string) (*WSGIndoorMap, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIndoorMap
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindMapsByIndoorMapId, true)
	req.SetPathParameter("indoorMapId", indoorMapId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIndoorMap()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindMapsByQueryCriteria
//
// Use this API command to query indoorMap.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIndoorMapService) FindMapsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGIndoorMapList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIndoorMapList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindMapsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIndoorMapList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateMapsByIndoorMapId
//
// Use this API command to update specific indoor map.
//
// Request Body:
//	 - body *WSGIndoorMap
//
// Required Parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) PartialUpdateMapsByIndoorMapId(ctx context.Context, body *WSGIndoorMap, indoorMapId string) (*APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateMapsByIndoorMapId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("indoorMapId", indoorMapId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateMapsApsByIndoorMapId
//
// Use this API command to put Aps in indoor map.
//
// Request Body:
//	 - body WSGIndoorMapAccessPointList
//
// Required Parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) UpdateMapsApsByIndoorMapId(ctx context.Context, body WSGIndoorMapAccessPointList, indoorMapId string) (*APIResponseMeta, error) {
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
	}
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateMapsApsByIndoorMapId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("indoorMapId", indoorMapId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
