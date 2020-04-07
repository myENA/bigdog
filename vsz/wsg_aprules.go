package vsz

// API Version: v9_0

type WSGAPRulesApRuleConfiguration struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// GpsCoordinates
	// Constraints:
	//    - nullable
	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpAddressRange
	// Constraints:
	//    - nullable
	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - nullable
	MobilityZone *WSGCommonGenericRef `json:"mobilityZone,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	// Constraints:
	//    - nullable
	ProvisionTag *string `json:"provisionTag,omitempty"`

	// Subnet
	// Constraints:
	//    - nullable
	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - nullable
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

func NewWSGAPRulesApRuleConfiguration() *WSGAPRulesApRuleConfiguration {
	m := new(WSGAPRulesApRuleConfiguration)
	return m
}

type WSGAPRulesApRuleList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAPRulesApRuleListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPRulesApRuleList() *WSGAPRulesApRuleList {
	m := new(WSGAPRulesApRuleList)
	return m
}

type WSGAPRulesApRuleListType struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`
}

func NewWSGAPRulesApRuleListType() *WSGAPRulesApRuleListType {
	m := new(WSGAPRulesApRuleListType)
	return m
}

type WSGAPRulesCreateApRule struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// GpsCoordinates
	// Constraints:
	//    - nullable
	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	// IpAddressRange
	// Constraints:
	//    - nullable
	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - required
	MobilityZone *WSGCommonGenericRef `json:"mobilityZone" validate:"required"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	// Constraints:
	//    - nullable
	ProvisionTag *string `json:"provisionTag,omitempty"`

	// Subnet
	// Constraints:
	//    - nullable
	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - required
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type" validate:"required,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

func NewWSGAPRulesCreateApRule() *WSGAPRulesCreateApRule {
	m := new(WSGAPRulesCreateApRule)
	return m
}

type WSGAPRulesGpsCoordinates struct {
	// Distance
	// distance
	// Constraints:
	//    - nullable
	Distance *float64 `json:"distance,omitempty"`

	// Latitude
	// Constraints:
	//    - nullable
	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	// Longitude
	// Constraints:
	//    - nullable
	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`
}

func NewWSGAPRulesGpsCoordinates() *WSGAPRulesGpsCoordinates {
	m := new(WSGAPRulesGpsCoordinates)
	return m
}

type WSGAPRulesIpAddressRange struct {
	// FromIp
	// Constraints:
	//    - nullable
	FromIp *WSGCommonIpAddress `json:"fromIp,omitempty"`

	// ToIp
	// Constraints:
	//    - nullable
	ToIp *WSGCommonIpAddress `json:"toIp,omitempty"`
}

func NewWSGAPRulesIpAddressRange() *WSGAPRulesIpAddressRange {
	m := new(WSGAPRulesIpAddressRange)
	return m
}

type WSGAPRulesModifyApRule struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// GpsCoordinates
	// Constraints:
	//    - nullable
	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	// IpAddressRange
	// Constraints:
	//    - nullable
	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - nullable
	MobilityZone *WSGCommonGenericRef `json:"mobilityZone,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	// Constraints:
	//    - nullable
	ProvisionTag *string `json:"provisionTag,omitempty"`

	// Subnet
	// Constraints:
	//    - nullable
	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - nullable
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

func NewWSGAPRulesModifyApRule() *WSGAPRulesModifyApRule {
	m := new(WSGAPRulesModifyApRule)
	return m
}

type WSGAPRulesSubnet struct {
	// NetworkAddress
	// Constraints:
	//    - nullable
	NetworkAddress *WSGCommonIpAddress `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGAPRulesSubnet() *WSGAPRulesSubnet {
	m := new(WSGAPRulesSubnet)
	return m
}
