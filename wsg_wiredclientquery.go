package bigdog

// API Version: v9_0

// WSGWiredClientQueryClientQueryList
//
// Definition: wiredClientQuery_clientQueryList
type WSGWiredClientQueryClientQueryList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWiredClientQueryCreateClientQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWiredClientQueryClientQueryList() *WSGWiredClientQueryClientQueryList {
	m := new(WSGWiredClientQueryClientQueryList)
	return m
}

// WSGWiredClientQueryCreateClientQuery
//
// Definition: wiredClientQuery_createClientQuery
type WSGWiredClientQueryCreateClientQuery struct {
	ApEthID *int `json:"apEthID,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AuthStatus *string `json:"authStatus,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	DeviceType *string `json:"deviceType,omitempty"`

	Downlink *int `json:"downlink,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	IpAddress *string `json:"ipAddress,omitempty"`

	Ipv6Address *string `json:"ipv6Address,omitempty"`

	ModelName *string `json:"modelName,omitempty"`

	OsType *string `json:"osType,omitempty"`

	OsVendorType *string `json:"osVendorType,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	RxFrames *int `json:"rxFrames,omitempty"`

	SessionStartTime *int `json:"sessionStartTime,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	TxFrames *int `json:"txFrames,omitempty"`

	TxRetry *int `json:"txRetry,omitempty"`

	TxRxBytes *int `json:"txRxBytes,omitempty"`

	Uplink *int `json:"uplink,omitempty"`

	UserName *string `json:"userName,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WlanType *string `json:"wlanType,omitempty"`
}

func NewWSGWiredClientQueryCreateClientQuery() *WSGWiredClientQueryCreateClientQuery {
	m := new(WSGWiredClientQueryCreateClientQuery)
	return m
}
