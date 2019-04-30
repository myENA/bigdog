package aprules

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/common"
)

type ApRuleConfiguration struct {
	Description *common.Description `json:"description,omitempty"`

	GpsCoordinates *GpsCoordinates `json:"gpsCoordinates,omitempty"`

	// ID
	// Identifier of the AP Registration Rules
	ID *string `json:"id,omitempty"`

	IPAddressRange *IPAddressRange `json:"ipAddressRange,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

type ApRuleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApRuleListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApRuleListType struct {
	Description *common.Description `json:"description,omitempty"`

	// ID
	// Identifier of the AP Registration Rules
	ID *string `json:"id,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`
}

type CreateApRule struct {
	Description *common.Description `json:"description,omitempty"`

	GpsCoordinates *GpsCoordinates `json:"gpsCoordinates,omitempty"`

	IPAddressRange *IPAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *common.GenericRef `json:"mobilityZone,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *Subnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	Type *string `json:"type,omitempty"`
}

type GpsCoordinates struct {
	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`
}

type IPAddressRange struct {
	FromIP *common.IPAddress `json:"fromIp,omitempty"`

	ToIP *common.IPAddress `json:"toIp,omitempty"`
}

type ModifyApRule struct {
	Description *common.Description `json:"description,omitempty"`

	GpsCoordinates *GpsCoordinates `json:"gpsCoordinates,omitempty"`

	IPAddressRange *IPAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *common.GenericRef `json:"mobilityZone,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *Subnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	Type *string `json:"type,omitempty"`
}

type Subnet struct {
	NetworkAddress *common.IPAddress `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}
