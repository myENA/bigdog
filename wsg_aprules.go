package bigdog

// API Version: v9_0

// WSGAPRulesApRuleConfiguration
//
// Definition: aprules_apRuleConfiguration
type WSGAPRulesApRuleConfiguration struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *WSGCommonGenericRef `json:"mobilityZone,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty"`
}

func NewWSGAPRulesApRuleConfiguration() *WSGAPRulesApRuleConfiguration {
	m := new(WSGAPRulesApRuleConfiguration)
	return m
}

// WSGAPRulesApRuleList
//
// Definition: aprules_apRuleList
type WSGAPRulesApRuleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPRulesApRuleListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPRulesApRuleList() *WSGAPRulesApRuleList {
	m := new(WSGAPRulesApRuleList)
	return m
}

// WSGAPRulesApRuleListType
//
// Definition: aprules_apRuleListType
type WSGAPRulesApRuleListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`
}

func NewWSGAPRulesApRuleListType() *WSGAPRulesApRuleListType {
	m := new(WSGAPRulesApRuleListType)
	return m
}

// WSGAPRulesCreateApRule
//
// Definition: aprules_createApRule
type WSGAPRulesCreateApRule struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - required
	MobilityZone *WSGCommonGenericRef `json:"mobilityZone"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - required
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type"`
}

func NewWSGAPRulesCreateApRule() *WSGAPRulesCreateApRule {
	m := new(WSGAPRulesCreateApRule)
	return m
}

// WSGAPRulesGpsCoordinates
//
// Definition: aprules_gpsCoordinates
type WSGAPRulesGpsCoordinates struct {
	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`
}

func NewWSGAPRulesGpsCoordinates() *WSGAPRulesGpsCoordinates {
	m := new(WSGAPRulesGpsCoordinates)
	return m
}

// WSGAPRulesIpAddressRange
//
// Definition: aprules_ipAddressRange
type WSGAPRulesIpAddressRange struct {
	FromIp *WSGCommonIpAddress `json:"fromIp,omitempty"`

	ToIp *WSGCommonIpAddress `json:"toIp,omitempty"`
}

func NewWSGAPRulesIpAddressRange() *WSGAPRulesIpAddressRange {
	m := new(WSGAPRulesIpAddressRange)
	return m
}

// WSGAPRulesModifyApRule
//
// Definition: aprules_modifyApRule
type WSGAPRulesModifyApRule struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *WSGCommonGenericRef `json:"mobilityZone,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty"`
}

func NewWSGAPRulesModifyApRule() *WSGAPRulesModifyApRule {
	m := new(WSGAPRulesModifyApRule)
	return m
}

// WSGAPRulesSubnet
//
// Definition: aprules_subnet
type WSGAPRulesSubnet struct {
	NetworkAddress *WSGCommonIpAddress `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGAPRulesSubnet() *WSGAPRulesSubnet {
	m := new(WSGAPRulesSubnet)
	return m
}
