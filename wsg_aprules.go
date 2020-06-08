package ruckus

// API Version: v9_0

type WSGAPRulesApRuleConfiguration struct {
	Description **WSGAPRulesApRuleConfiguration `json:"description,omitempty"`

	GpsCoordinates **WSGAPRulesApRuleConfiguration `json:"gpsCoordinates,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	IpAddressRange **WSGAPRulesApRuleConfiguration `json:"ipAddressRange,omitempty"`

	MobilityZone **WSGAPRulesApRuleConfiguration `json:"mobilityZone,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet **WSGAPRulesApRuleConfiguration `json:"subnet,omitempty"`

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

type WSGAPRulesApRuleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGAPRulesApRuleList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPRulesApRuleList() *WSGAPRulesApRuleList {
	m := new(WSGAPRulesApRuleList)
	return m
}

type WSGAPRulesApRuleListType struct {
	Description **WSGAPRulesApRuleListType `json:"description,omitempty"`

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

type WSGAPRulesCreateApRule struct {
	Description **WSGAPRulesCreateApRule `json:"description,omitempty"`

	GpsCoordinates **WSGAPRulesCreateApRule `json:"gpsCoordinates,omitempty"`

	IpAddressRange **WSGAPRulesCreateApRule `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - required
	MobilityZone **WSGAPRulesCreateApRule `json:"mobilityZone"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet **WSGAPRulesCreateApRule `json:"subnet,omitempty"`

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

type WSGAPRulesGpsCoordinates struct {
	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	Latitude **WSGAPRulesGpsCoordinates `json:"latitude,omitempty"`

	Longitude **WSGAPRulesGpsCoordinates `json:"longitude,omitempty"`
}

func NewWSGAPRulesGpsCoordinates() *WSGAPRulesGpsCoordinates {
	m := new(WSGAPRulesGpsCoordinates)
	return m
}

type WSGAPRulesIpAddressRange struct {
	FromIp **WSGAPRulesIpAddressRange `json:"fromIp,omitempty"`

	ToIp **WSGAPRulesIpAddressRange `json:"toIp,omitempty"`
}

func NewWSGAPRulesIpAddressRange() *WSGAPRulesIpAddressRange {
	m := new(WSGAPRulesIpAddressRange)
	return m
}

type WSGAPRulesModifyApRule struct {
	Description **WSGAPRulesModifyApRule `json:"description,omitempty"`

	GpsCoordinates **WSGAPRulesModifyApRule `json:"gpsCoordinates,omitempty"`

	IpAddressRange **WSGAPRulesModifyApRule `json:"ipAddressRange,omitempty"`

	MobilityZone **WSGAPRulesModifyApRule `json:"mobilityZone,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet **WSGAPRulesModifyApRule `json:"subnet,omitempty"`

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

type WSGAPRulesSubnet struct {
	NetworkAddress **WSGAPRulesSubnet `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGAPRulesSubnet() *WSGAPRulesSubnet {
	m := new(WSGAPRulesSubnet)
	return m
}
