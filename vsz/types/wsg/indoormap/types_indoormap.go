package indoormap

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AccessPointList []*IndoorMapAp

func NewAccessPointList() *AccessPointList {
	accessPointListType := make(AccessPointList, 0)
	return &accessPointListType
}

func NewAccessPointListWithDefaults() *AccessPointList {
	accessPointListType := make(AccessPointList, 0)
	return &accessPointListType
}

type BasicIndoorMap struct {
	Address *string `json:"address,omitempty"`

	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *common.Description `json:"description" validate:"required"`

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
	Name *common.NormalName `json:"name" validate:"required"`

	// Orientation
	// Constraints:
	//    - nullable
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty" validate:"omitempty,oneof=HORIZONTAL VERTICAL"`

	Scale *Scale `json:"scale,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewBasicIndoorMap() *BasicIndoorMap {
	basicIndoorMapType := new(BasicIndoorMap)
	return basicIndoorMapType
}

func NewBasicIndoorMapWithDefaults() *BasicIndoorMap {
	basicIndoorMapType := new(BasicIndoorMap)
	return basicIndoorMapType
}

type IndooMapAuditId struct {
	// Id
	// the identifier of the indoor map
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the indoor map
	Name *string `json:"name,omitempty"`
}

func NewIndooMapAuditId() *IndooMapAuditId {
	indooMapAuditIdType := new(IndooMapAuditId)
	return indooMapAuditIdType
}

func NewIndooMapAuditIdWithDefaults() *IndooMapAuditId {
	indooMapAuditIdType := new(IndooMapAuditId)
	return indooMapAuditIdType
}

type IndoorMap struct {
	// Address
	// address
	Address *string `json:"address,omitempty"`

	// ApGroupId
	// apGroupId
	ApGroupId *string `json:"apGroupId,omitempty"`

	// Description
	// Constraints:
	//    - required
	Description *common.Description `json:"description" validate:"required"`

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
	Name *common.NormalName `json:"name" validate:"required"`

	// Orientation
	// orientation
	// Constraints:
	//    - nullable
	//    - oneof:[HORIZONTAL,VERTICAL]
	Orientation *string `json:"orientation,omitempty" validate:"omitempty,oneof=HORIZONTAL VERTICAL"`

	Scale *Scale `json:"scale,omitempty"`

	// TenantId
	// tenantId
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// zoneId
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewIndoorMap() *IndoorMap {
	indoorMapType := new(IndoorMap)
	return indoorMapType
}

func NewIndoorMapWithDefaults() *IndoorMap {
	indoorMapType := new(IndoorMap)
	return indoorMapType
}

type IndoorMapAp struct {
	IndoorMapXy *IndoorMapXy `json:"indoorMapXy,omitempty"`

	// Mac
	// the identifier of the create object
	Mac *string `json:"mac,omitempty"`
}

func NewIndoorMapAp() *IndoorMapAp {
	indoorMapApType := new(IndoorMapAp)
	return indoorMapApType
}

func NewIndoorMapApWithDefaults() *IndoorMapAp {
	indoorMapApType := new(IndoorMapAp)
	return indoorMapApType
}

type IndoorMapList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first indoorMapList returned out of the complete indoor maps list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more indoor maps after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*BasicIndoorMap `json:"list,omitempty"`

	// TotalCount
	// Total indoor maps count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewIndoorMapList() *IndoorMapList {
	indoorMapListType := new(IndoorMapList)
	return indoorMapListType
}

func NewIndoorMapListWithDefaults() *IndoorMapList {
	indoorMapListType := new(IndoorMapList)
	return indoorMapListType
}

type IndoorMapSummary struct {
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
	Description *common.Description `json:"description" validate:"required"`

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
	Name *common.NormalName `json:"name" validate:"required"`

	Scale *Scale `json:"scale,omitempty"`

	// TenantId
	// tenantId
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// zoneId
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewIndoorMapSummary() *IndoorMapSummary {
	indoorMapSummaryType := new(IndoorMapSummary)
	return indoorMapSummaryType
}

func NewIndoorMapSummaryWithDefaults() *IndoorMapSummary {
	indoorMapSummaryType := new(IndoorMapSummary)
	return indoorMapSummaryType
}

type IndoorMapSummaryList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first indoorMapList returned out of the complete indoor maps list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more indoor maps after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*IndoorMapSummary `json:"list,omitempty"`

	// TotalCount
	// Total indoor maps count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewIndoorMapSummaryList() *IndoorMapSummaryList {
	indoorMapSummaryListType := new(IndoorMapSummaryList)
	return indoorMapSummaryListType
}

func NewIndoorMapSummaryListWithDefaults() *IndoorMapSummaryList {
	indoorMapSummaryListType := new(IndoorMapSummaryList)
	return indoorMapSummaryListType
}

type IndoorMapXy struct {
	// X
	// x
	X *float64 `json:"x,omitempty"`

	// Y
	// y
	Y *float64 `json:"y,omitempty"`
}

func NewIndoorMapXy() *IndoorMapXy {
	indoorMapXyType := new(IndoorMapXy)
	return indoorMapXyType
}

func NewIndoorMapXyWithDefaults() *IndoorMapXy {
	indoorMapXyType := new(IndoorMapXy)
	return indoorMapXyType
}

type Scale struct {
	A *IndoorMapXy `json:"a,omitempty"`

	B *IndoorMapXy `json:"b,omitempty"`

	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	// Unit
	// unit
	// Constraints:
	//    - nullable
	//    - oneof:[MM,CM,M,Foot,Yard]
	Unit *string `json:"unit,omitempty" validate:"omitempty,oneof=MM CM M Foot Yard"`
}

func NewScale() *Scale {
	scaleType := new(Scale)
	return scaleType
}

func NewScaleWithDefaults() *Scale {
	scaleType := new(Scale)
	return scaleType
}
