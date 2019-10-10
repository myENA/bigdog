package aprules

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApRuleConfiguration struct {
	Description *common.Description `json:"description,omitempty"`

	GpsCoordinates *GpsCoordinates `json:"gpsCoordinates,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	IpAddressRange *IpAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *common.GenericRef `json:"mobilityZone,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *Subnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	Type *string `json:"type,omitempty" validate:"oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

type ApRuleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApRuleListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApRuleListType struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`
}

type CreateApRule struct {
	Description *common.Description `json:"description,omitempty"`

	GpsCoordinates *GpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *IpAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *common.GenericRef `json:"mobilityZone,omitempty" validate:"required"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *Subnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	Type *string `json:"type,omitempty" validate:"required,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

type GpsCoordinates struct {
	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`
}

type IpAddressRange struct {
	FromIp *common.IpAddress `json:"fromIp,omitempty"`

	ToIp *common.IpAddress `json:"toIp,omitempty"`
}

type ModifyApRule struct {
	Description *common.Description `json:"description,omitempty"`

	GpsCoordinates *GpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *IpAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *common.GenericRef `json:"mobilityZone,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *Subnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	Type *string `json:"type,omitempty" validate:"oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

type Subnet struct {
	NetworkAddress *common.IpAddress `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}
