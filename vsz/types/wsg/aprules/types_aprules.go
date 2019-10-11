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
	// Constraints:
	//    - nullable
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

func NewApRuleConfiguration() *ApRuleConfiguration {
	apRuleConfigurationType := new(ApRuleConfiguration)
	return apRuleConfigurationType
}

func NewDefaultApRuleConfiguration() *ApRuleConfiguration {
	apRuleConfigurationType := new(ApRuleConfiguration)
	return apRuleConfigurationType
}

type ApRuleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApRuleListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApRuleList() *ApRuleList {
	apRuleListType := new(ApRuleList)
	return apRuleListType
}

func NewDefaultApRuleList() *ApRuleList {
	apRuleListType := new(ApRuleList)
	return apRuleListType
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

func NewApRuleListType() *ApRuleListType {
	apRuleListTypeType := new(ApRuleListType)
	return apRuleListTypeType
}

func NewDefaultApRuleListType() *ApRuleListType {
	apRuleListTypeType := new(ApRuleListType)
	return apRuleListTypeType
}

type CreateApRule struct {
	Description *common.Description `json:"description,omitempty"`

	GpsCoordinates *GpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *IpAddressRange `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - required
	MobilityZone *common.GenericRef `json:"mobilityZone" validate:"required"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *Subnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - required
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type" validate:"required,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

func NewCreateApRule() *CreateApRule {
	createApRuleType := new(CreateApRule)
	return createApRuleType
}

func NewDefaultCreateApRule() *CreateApRule {
	createApRuleType := new(CreateApRule)
	return createApRuleType
}

type GpsCoordinates struct {
	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`
}

func NewGpsCoordinates() *GpsCoordinates {
	gpsCoordinatesType := new(GpsCoordinates)
	return gpsCoordinatesType
}

func NewDefaultGpsCoordinates() *GpsCoordinates {
	gpsCoordinatesType := new(GpsCoordinates)
	return gpsCoordinatesType
}

type IpAddressRange struct {
	FromIp *common.IpAddress `json:"fromIp,omitempty"`

	ToIp *common.IpAddress `json:"toIp,omitempty"`
}

func NewIpAddressRange() *IpAddressRange {
	ipAddressRangeType := new(IpAddressRange)
	return ipAddressRangeType
}

func NewDefaultIpAddressRange() *IpAddressRange {
	ipAddressRangeType := new(IpAddressRange)
	return ipAddressRangeType
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
	// Constraints:
	//    - nullable
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

func NewModifyApRule() *ModifyApRule {
	modifyApRuleType := new(ModifyApRule)
	return modifyApRuleType
}

func NewDefaultModifyApRule() *ModifyApRule {
	modifyApRuleType := new(ModifyApRule)
	return modifyApRuleType
}

type Subnet struct {
	NetworkAddress *common.IpAddress `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewSubnet() *Subnet {
	subnetType := new(Subnet)
	return subnetType
}

func NewDefaultSubnet() *Subnet {
	subnetType := new(Subnet)
	return subnetType
}
