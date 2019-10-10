package indoormap

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AccessPointList []*IndoorMapAp

type BasicIndoorMap struct {
	Address *string `json:"address,omitempty"`

	ApGroupID *string `json:"apGroupId,omitempty"`

	Description *common.Description `json:"description,omitempty" validate:"required"`

	DomainID *string `json:"domainId,omitempty"`

	GroupType *string `json:"groupType,omitempty" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

	ID *string `json:"id,omitempty"`

	ImageFileName *string `json:"imageFileName,omitempty"`

	Latitude *float64 `json:"latitude,omitempty"`

	Longitude *float64 `json:"longitude,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	Orientation *string `json:"orientation,omitempty" validate:"oneof=HORIZONTAL VERTICAL"`

	Scale *Scale `json:"scale,omitempty"`

	TenantID *string `json:"tenantId,omitempty"`

	ZoneID *string `json:"zoneId,omitempty"`
}

type IndooMapAuditID struct {
	// ID
	// the identifier of the indoor map
	ID *string `json:"id,omitempty"`

	// Name
	// the name of the indoor map
	Name *string `json:"name,omitempty"`
}

type IndoorMap struct {
	// Address
	// address
	Address *string `json:"address,omitempty"`

	// ApGroupID
	// apGroupId
	ApGroupID *string `json:"apGroupId,omitempty"`

	Description *common.Description `json:"description,omitempty" validate:"required"`

	// DomainID
	// domainId
	DomainID *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	GroupType *string `json:"groupType,omitempty" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

	// ID
	// id
	ID *string `json:"id,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	// Orientation
	// orientation
	Orientation *string `json:"orientation,omitempty" validate:"oneof=HORIZONTAL VERTICAL"`

	Scale *Scale `json:"scale,omitempty"`

	// TenantID
	// tenantId
	TenantID *string `json:"tenantId,omitempty"`

	// ZoneID
	// zoneId
	ZoneID *string `json:"zoneId,omitempty"`
}

type IndoorMapAp struct {
	IndoorMapXy *IndoorMapXy `json:"indoorMapXy,omitempty"`

	// Mac
	// the identifier of the create object
	Mac *string `json:"mac,omitempty"`
}

type IndoorMapList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

type IndoorMapSummary struct {
	// Address
	// address
	Address *string `json:"address,omitempty"`

	// ApCount
	// AP count in this indoor map
	ApCount *float64 `json:"apCount,omitempty"`

	// ApGroupID
	// apGroupId
	ApGroupID *string `json:"apGroupId,omitempty"`

	Description *common.Description `json:"description,omitempty" validate:"required"`

	// DomainID
	// domainId
	DomainID *string `json:"domainId,omitempty"`

	// GroupType
	// group Type
	GroupType *string `json:"groupType,omitempty" validate:"required,oneof=SYSTEM DOMAIN ZONE THIRD_PARTY_ZONE APGROUP"`

	// ID
	// id
	ID *string `json:"id,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	Scale *Scale `json:"scale,omitempty"`

	// TenantID
	// tenantId
	TenantID *string `json:"tenantId,omitempty"`

	// ZoneID
	// zoneId
	ZoneID *string `json:"zoneId,omitempty"`
}

type IndoorMapSummaryList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

type IndoorMapXy struct {
	// X
	// x
	X *float64 `json:"x,omitempty"`

	// Y
	// y
	Y *float64 `json:"y,omitempty"`
}

type Scale struct {
	A *IndoorMapXy `json:"a,omitempty"`

	B *IndoorMapXy `json:"b,omitempty"`

	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	// Unit
	// unit
	Unit *string `json:"unit,omitempty" validate:"oneof=MM CM M Foot Yard"`
}
