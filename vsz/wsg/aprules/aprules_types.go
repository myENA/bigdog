package aprules

// API Version: v8_0

type ApRuleConfiguration struct {
	Description    *string            `json:"description,omitempty"`
	GpsCoordinates *GpsCoordinates    `json:"gpsCoordinates,omitempty"`
	ID             *string            `json:"id,omitempty"`
	IPAddressRange *IPAddressRange    `json:"ipAddressRange,omitempty"`
	MobilityZone   *common.GenericRef `json:"mobilityZone,omitempty"`
	Priority       *int               `json:"priority,omitempty"`
	ProvisionTag   *string            `json:"provisionTag,omitempty"`
	Subnet         *Subnet            `json:"subnet,omitempty"`
	Type           *string            `json:"type,omitempty"`
}

type ApRuleList struct {
	FirstIndex *int              `json:"firstIndex,omitempty"`
	HasMore    *bool             `json:"hasMore,omitempty"`
	List       []*ApRuleListType `json:"list,omitempty"`
	TotalCount *int              `json:"totalCount,omitempty"`
}

type ApRuleListType struct {
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
}

type CreateApRule struct {
	Description    *string            `json:"description,omitempty"`
	GpsCoordinates *GpsCoordinates    `json:"gpsCoordinates,omitempty"`
	IPAddressRange *IPAddressRange    `json:"ipAddressRange,omitempty"`
	MobilityZone   *common.GenericRef `json:"mobilityZone,omitempty"`
	ProvisionTag   *string            `json:"provisionTag,omitempty"`
	Subnet         *Subnet            `json:"subnet,omitempty"`
	Type           *string            `json:"type,omitempty"`
}

type GpsCoordinates struct {
	Distance  *float64 `json:"distance,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type IPAddressRange struct {
	FromIP *string `json:"fromIp,omitempty"`
	ToIP   *string `json:"toIp,omitempty"`
}

type ModifyApRule struct {
	Description    *string            `json:"description,omitempty"`
	GpsCoordinates *GpsCoordinates    `json:"gpsCoordinates,omitempty"`
	IPAddressRange *IPAddressRange    `json:"ipAddressRange,omitempty"`
	MobilityZone   *common.GenericRef `json:"mobilityZone,omitempty"`
	ProvisionTag   *string            `json:"provisionTag,omitempty"`
	Subnet         *Subnet            `json:"subnet,omitempty"`
	Type           *string            `json:"type,omitempty"`
}

type Subnet struct {
	NetworkAddress *string `json:"networkAddress,omitempty"`
	SubnetMask     *string `json:"subnetMask,omitempty"`
}
