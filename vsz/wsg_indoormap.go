package vsz

// API Version: v8_1

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
	Address *string `json:"address,omitempty"`

	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *WSGCommonDescription `json:"description" validate:"required,omitempty,max=64"`

	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

	Id *string `json:"id,omitempty"`

	ImageFileName *string `json:"imageFileName,omitempty"`

	Latitude *float64 `json:"latitude,omitempty"`

	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Orientation
	// Constraints:
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty" validate:"oneof=HORIZONTAL VERTICAL"`

	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGIndoorMapBasicIndoorMap() *WSGIndoorMapBasicIndoorMap {
	m := new(WSGIndoorMapBasicIndoorMap)
	return m
}

type WSGIndoorMapIndooMapAuditId struct {
	// Id
	// the identifier of the indoor map
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the indoor map
	Name *string `json:"name,omitempty"`
}

func NewWSGIndoorMapIndooMapAuditId() *WSGIndoorMapIndooMapAuditId {
	m := new(WSGIndoorMapIndooMapAuditId)
	return m
}

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
	Description *WSGCommonDescription `json:"description" validate:"required,omitempty,max=64"`

	// DomainId
	// domainId
	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Orientation
	// orientation
	// Constraints:
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty" validate:"oneof=HORIZONTAL VERTICAL"`

	Scale *WSGIndoorMapScale `json:"scale,omitempty"`

	// TenantId
	// tenantId
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// zoneId
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGIndoorMap() *WSGIndoorMap {
	m := new(WSGIndoorMap)
	return m
}

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

type WSGIndoorMapList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

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

func NewWSGIndoorMapList() *WSGIndoorMapList {
	m := new(WSGIndoorMapList)
	return m
}

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
	Description *WSGCommonDescription `json:"description" validate:"required,omitempty,max=64"`

	// DomainId
	// domainId
	DomainId *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP]
	GroupType *string `json:"groupType" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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

type WSGIndoorMapSummaryList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

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

func NewWSGIndoorMapSummaryList() *WSGIndoorMapSummaryList {
	m := new(WSGIndoorMapSummaryList)
	return m
}

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
	Unit *string `json:"unit,omitempty" validate:"oneof=MM CM M Foot Yard"`
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
func (s *WSGIndoorMapService) AddMaps(ctx context.Context, body *WSGIndoorMap) (*WSGIndoorMapIndooMapAuditId, error) {
	var (
		req  *APIRequest
		resp *WSGIndoorMapIndooMapAuditId
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddMaps, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteMapsByIndoorMapId
//
// Use this API command to delete indoor map.
//
// Required Parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) DeleteMapsByIndoorMapId(ctx context.Context, indoorMapId string) (*WSGIndoorMapIndooMapAuditId, error) {
	var (
		req  *APIRequest
		resp *WSGIndoorMapIndooMapAuditId
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteMapsByIndoorMapId, true)
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
func (s *WSGIndoorMapService) FindMaps(ctx context.Context, groupId string, groupType string) (*WSGIndoorMapList, error) {
	var (
		req  *APIRequest
		resp *WSGIndoorMapList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, groupType, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindMaps, true)
}

// FindMapsByIndoorMapId
//
// Use this API command to get indoor maps.
//
// Required Parameters:
// - indoorMapId string
//		- required
func (s *WSGIndoorMapService) FindMapsByIndoorMapId(ctx context.Context, indoorMapId string) (*WSGIndoorMap, error) {
	var (
		req  *APIRequest
		resp *WSGIndoorMap
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindMapsByIndoorMapId, true)
}

// FindMapsByQueryCriteria
//
// Use this API command to query indoorMap.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIndoorMapService) FindMapsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGIndoorMapList, error) {
	var (
		req  *APIRequest
		resp *WSGIndoorMapList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindMapsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
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
func (s *WSGIndoorMapService) PartialUpdateMapsByIndoorMapId(ctx context.Context, body *WSGIndoorMap, indoorMapId string) (*WSGIndoorMapIndooMapAuditId, error) {
	var (
		req  *APIRequest
		resp *WSGIndoorMapIndooMapAuditId
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
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateMapsByIndoorMapId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
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
func (s *WSGIndoorMapService) UpdateMapsApsByIndoorMapId(ctx context.Context, body WSGIndoorMapAccessPointList, indoorMapId string) (*WSGIndoorMapIndooMapAuditId, error) {
	var (
		req  *APIRequest
		resp *WSGIndoorMapIndooMapAuditId
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, indoorMapId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateMapsApsByIndoorMapId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
