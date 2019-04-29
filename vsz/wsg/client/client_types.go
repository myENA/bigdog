package client

// API Version: v8_0

type DeAuthClient struct {
	ApMac *string `json:"apMac,omitempty"`
	Mac   *string `json:"mac,omitempty"`
}

type DeAuthClientList struct {
	ClientList []*ClientList `json:"clientList,omitempty"`
}

type DisconnectClient struct {
	ApMac *string `json:"apMac,omitempty"`
	Mac   *string `json:"mac,omitempty"`
}

type DisconnectClientList struct {
	ClientList []*ClientList `json:"clientList,omitempty"`
}

type HistoricalClient struct {
	ApMac            *string `json:"apMac,omitempty"`
	ClientMac        *string `json:"clientMac,omitempty"`
	CoreNetworkType  *string `json:"coreNetworkType,omitempty"`
	Hostname         *string `json:"hostname,omitempty"`
	IPAddress        *string `json:"ipAddress,omitempty"`
	Ipv6Address      *string `json:"ipv6Address,omitempty"`
	MvnoName         *string `json:"mvnoName,omitempty"`
	OsType           *string `json:"osType,omitempty"`
	RxBytes          *int    `json:"rxBytes,omitempty"`
	RxDrops          *int    `json:"rxDrops,omitempty"`
	RxFrames         *int    `json:"rxFrames,omitempty"`
	SessionEndTime   *int    `json:"sessionEndTime,omitempty"`
	SessionStartTime *int    `json:"sessionStartTime,omitempty"`
	Ssid             *string `json:"ssid,omitempty"`
	TxBytes          *int    `json:"txBytes,omitempty"`
	TxDrops          *int    `json:"txDrops,omitempty"`
	TxFrames         *int    `json:"txFrames,omitempty"`
}

type HistoricalClientList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}
