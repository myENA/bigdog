package meshneighborinfo

// API Version: v8_0

type MeshNeighborInfo struct {
	// ApFirmware
	// The firmware of the neighbor AP
	ApFirmware *string `json:"apFirmware,omitempty"`

	// ApMac
	// The MAC of the neighbor AP
	ApMac *string `json:"apMac,omitempty"`

	// ApModel
	// The model of the neighbor AP
	ApModel *string `json:"apModel,omitempty"`

	// ApName
	// The name of the neighbor AP
	ApName *string `json:"apName,omitempty"`

	// Channel
	// The channel of the neighbor AP
	Channel *string `json:"channel,omitempty"`

	// ConnectionStatus
	// The connection status of the neighbor AP
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	// ExternalIPAddress
	// The external IP of the neighbor AP
	ExternalIPAddress *string `json:"externalIPAddress,omitempty"`

	// IpAddress
	// The IP of the neighbor AP
	IpAddress *string `json:"ipAddress,omitempty"`

	// Rssi
	// The RSSI of the neighbor AP
	Rssi *int `json:"rssi,omitempty"`

	// ZoneName
	// The zone name of the neighbor AP
	ZoneName *string `json:"zoneName,omitempty"`
}

type MeshNeighborInfoList struct {
	// Extra
	// Any additional response data.
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first MeshNeighborInfo returned out of the complete Rogue AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more MeshNeighborInfo after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*MeshNeighborInfo `json:"list,omitempty"`

	// RawDataTotalCount
	// MeshNeighborInfos count.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// MeshNeighborInfos count in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

type MeshNeighborInfoListExtraType map[string]interface{}
