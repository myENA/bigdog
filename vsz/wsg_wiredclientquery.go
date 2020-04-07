package vsz

// API Version: v9_0

type WSGWiredClientQueryClientQueryList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGWiredClientQueryCreateClientQuery `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWiredClientQueryClientQueryList() *WSGWiredClientQueryClientQueryList {
	m := new(WSGWiredClientQueryClientQueryList)
	return m
}

type WSGWiredClientQueryCreateClientQuery struct {
	// ApEthID
	// Constraints:
	//    - nullable
	ApEthID *int `json:"apEthID,omitempty"`

	// ApMac
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// ApName
	// Constraints:
	//    - nullable
	ApName *string `json:"apName,omitempty"`

	// AuthStatus
	// Constraints:
	//    - nullable
	AuthStatus *string `json:"authStatus,omitempty"`

	// ClientMac
	// Constraints:
	//    - nullable
	ClientMac *string `json:"clientMac,omitempty"`

	// DeviceType
	// Constraints:
	//    - nullable
	DeviceType *string `json:"deviceType,omitempty"`

	// Downlink
	// Constraints:
	//    - nullable
	Downlink *int `json:"downlink,omitempty"`

	// Hostname
	// Constraints:
	//    - nullable
	Hostname *string `json:"hostname,omitempty"`

	// IpAddress
	// Constraints:
	//    - nullable
	IpAddress *string `json:"ipAddress,omitempty"`

	// Ipv6Address
	// Constraints:
	//    - nullable
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	// ModelName
	// Constraints:
	//    - nullable
	ModelName *string `json:"modelName,omitempty"`

	// OsType
	// Constraints:
	//    - nullable
	OsType *string `json:"osType,omitempty"`

	// OsVendorType
	// Constraints:
	//    - nullable
	OsVendorType *string `json:"osVendorType,omitempty"`

	// RxBytes
	// Constraints:
	//    - nullable
	RxBytes *int `json:"rxBytes,omitempty"`

	// RxFrames
	// Constraints:
	//    - nullable
	RxFrames *int `json:"rxFrames,omitempty"`

	// SessionStartTime
	// Constraints:
	//    - nullable
	SessionStartTime *int `json:"sessionStartTime,omitempty"`

	// Traffic
	// Constraints:
	//    - nullable
	Traffic *int `json:"traffic,omitempty"`

	// TxBytes
	// Constraints:
	//    - nullable
	TxBytes *int `json:"txBytes,omitempty"`

	// TxFrames
	// Constraints:
	//    - nullable
	TxFrames *int `json:"txFrames,omitempty"`

	// TxRetry
	// Constraints:
	//    - nullable
	TxRetry *int `json:"txRetry,omitempty"`

	// TxRxBytes
	// Constraints:
	//    - nullable
	TxRxBytes *int `json:"txRxBytes,omitempty"`

	// Uplink
	// Constraints:
	//    - nullable
	Uplink *int `json:"uplink,omitempty"`

	// UserName
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`

	// Vlan
	// Constraints:
	//    - nullable
	Vlan *int `json:"vlan,omitempty"`

	// WlanType
	// Constraints:
	//    - nullable
	WlanType *string `json:"wlanType,omitempty"`
}

func NewWSGWiredClientQueryCreateClientQuery() *WSGWiredClientQueryCreateClientQuery {
	m := new(WSGWiredClientQueryCreateClientQuery)
	return m
}
