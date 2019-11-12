package vsz

// API Version: v8_1

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
	Type *string `json:"type,omitempty" validate:"oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

type WSGAPRulesApRuleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPRulesApRuleListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAPRulesApRuleListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`
}

type WSGAPRulesCreateApRule struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - required
	MobilityZone *WSGCommonGenericRef `json:"mobilityZone" validate:"required"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - required
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type" validate:"required,oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

type WSGAPRulesGpsCoordinates struct {
	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`
}

type WSGAPRulesIpAddressRange struct {
	FromIp *WSGCommonIpAddress `json:"fromIp,omitempty"`

	ToIp *WSGCommonIpAddress `json:"toIp,omitempty"`
}

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
	Type *string `json:"type,omitempty" validate:"oneof=IPAddressRange Subnet GPSCoordinates ProvisionTag"`
}

type WSGAPRulesSubnet struct {
	NetworkAddress *WSGCommonIpAddress `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}
