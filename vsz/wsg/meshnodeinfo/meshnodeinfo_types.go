package meshnodeinfo

// API Version: v8_0

type HelperZoneInfo struct {
	HelperAPZoneID   *string `json:"helperAPZoneId,omitempty"`
	HelperAPZoneName *string `json:"helperAPZoneName,omitempty"`
}

type MeshNodeInfo struct {
	ApMac             *string           `json:"apMac,omitempty"`
	ApModel           *string           `json:"apModel,omitempty"`
	ApName            *string           `json:"apName,omitempty"`
	Channel           *string           `json:"channel,omitempty"`
	ClientCount       *int              `json:"clientCount,omitempty"`
	DownlinkRssi      *int              `json:"downlinkRssi,omitempty"`
	ExternalIPAddress *string           `json:"externalIPAddress,omitempty"`
	HasDownLink       *bool             `json:"hasDownLink,omitempty"`
	HelperZoneInfo    []*HelperZoneInfo `json:"helperZoneInfo,omitempty"`
	Hops              *int              `json:"hops,omitempty"`
	IPAddress         *int              `json:"ipAddress,omitempty"`
	MeshRole          *string           `json:"meshRole,omitempty"`
	UplinkRssi        *int              `json:"uplinkRssi,omitempty"`
}

type MeshNodeInfoList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type UpdateAPZeroTouch struct {
	ApMac        *string `json:"apMac,omitempty"`
	HelperZoneID *string `json:"helperZoneId,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	Status       *string `json:"status,omitempty"`
}
