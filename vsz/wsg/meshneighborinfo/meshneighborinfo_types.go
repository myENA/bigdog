package meshneighborinfo

// API Version: v8_0

type MeshNeighborInfo struct {
	ApFirmware        *string `json:"apFirmware,omitempty"`
	ApMac             *string `json:"apMac,omitempty"`
	ApModel           *string `json:"apModel,omitempty"`
	ApName            *string `json:"apName,omitempty"`
	Channel           *string `json:"channel,omitempty"`
	ConnectionStatus  *string `json:"connectionStatus,omitempty"`
	ExternalIPAddress *string `json:"externalIPAddress,omitempty"`
	IPAddress         *string `json:"ipAddress,omitempty"`
	Rssi              *int    `json:"rssi,omitempty"`
	ZoneName          *string `json:"zoneName,omitempty"`
}

type MeshNeighborInfoList struct {
	Extra             interface{}         `json:"extra,omitempty"`
	FirstIndex        *int                `json:"firstIndex,omitempty"`
	HasMore           *bool               `json:"hasMore,omitempty"`
	List              []*MeshNeighborInfo `json:"list,omitempty"`
	RawDataTotalCount *int                `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int                `json:"totalCount,omitempty"`
}
