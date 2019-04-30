package aprules

// API Version: v8_0

type ApRuleConfiguration struct {
	Description *string `json:"description,omitempty"`

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
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the AP Registration Rules
	ID *string `json:"id,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`
}

type CreateApRule struct {
	Description *string `json:"description,omitempty"`

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

	Latitude *float64 `json:"latitude,omitempty"`

	Longitude *float64 `json:"longitude,omitempty"`
}

type IPAddressRange struct {
	FromIP *string `json:"fromIp,omitempty"`

	ToIP *string `json:"toIp,omitempty"`
}

type ModifyApRule struct {
	Description *string `json:"description,omitempty"`

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
	NetworkAddress *string `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}
