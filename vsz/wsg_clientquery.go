package vsz

// API Version: v9_0

type WSGClientQueryList struct {
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
	List []*WSGClientQueryCreateClientQuery `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGClientQueryList() *WSGClientQueryList {
	m := new(WSGClientQueryList)
	return m
}

type WSGClientQueryCreateClientQuery struct {
	// Alerts
	// Constraints:
	//    - nullable
	Alerts *int `json:"alerts,omitempty"`

	// ApMac
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// ApName
	// Constraints:
	//    - nullable
	ApName *string `json:"apName,omitempty"`

	// AuthMethod
	// Constraints:
	//    - nullable
	AuthMethod *string `json:"authMethod,omitempty"`

	// AuthStatus
	// Constraints:
	//    - nullable
	AuthStatus *string `json:"authStatus,omitempty"`

	// Bssid
	// Constraints:
	//    - nullable
	Bssid *string `json:"bssid,omitempty"`

	// Channel
	// Constraints:
	//    - nullable
	Channel *int `json:"channel,omitempty"`

	// ClientMac
	// Constraints:
	//    - nullable
	ClientMac *string `json:"clientMac,omitempty"`

	// ControlPlaneName
	// Constraints:
	//    - nullable
	ControlPlaneName *string `json:"controlPlaneName,omitempty"`

	// CpeMac
	// Constraints:
	//    - nullable
	CpeMac *string `json:"cpeMac,omitempty"`

	// DataPlaneName
	// Constraints:
	//    - nullable
	DataPlaneName *string `json:"dataPlaneName,omitempty"`

	// DeviceType
	// Constraints:
	//    - nullable
	DeviceType *string `json:"deviceType,omitempty"`

	// Downlink
	// Constraints:
	//    - nullable
	Downlink *int `json:"downlink,omitempty"`

	// DownlinkRate
	// Constraints:
	//    - nullable
	DownlinkRate *int `json:"downlinkRate,omitempty"`

	// EncryptionMethod
	// Constraints:
	//    - nullable
	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

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

	// MedianRxMCSRate
	// Constraints:
	//    - nullable
	MedianRxMCSRate *int `json:"medianRxMCSRate,omitempty"`

	// MedianTxMCSRate
	// Constraints:
	//    - nullable
	MedianTxMCSRate *int `json:"medianTxMCSRate,omitempty"`

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

	// RadioType
	// Constraints:
	//    - nullable
	RadioType *string `json:"radioType,omitempty"`

	// Role
	// Constraints:
	//    - nullable
	Role *string `json:"role,omitempty"`

	// Rssi
	// Constraints:
	//    - nullable
	Rssi *int `json:"rssi,omitempty"`

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

	// Snr
	// Constraints:
	//    - nullable
	Snr *int `json:"snr,omitempty"`

	// Speedflex
	// Constraints:
	//    - nullable
	Speedflex *int `json:"speedflex,omitempty"`

	// Ssid
	// Constraints:
	//    - nullable
	Ssid *string `json:"ssid,omitempty"`

	// Status
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// TcWithQuotaList
	// Constraints:
	//    - nullable
	TcWithQuotaList []*WSGClientQueryTcWithQuota `json:"tcWithQuotaList,omitempty" validate:"omitempty,dive"`

	// Traffic
	// Constraints:
	//    - nullable
	Traffic *int `json:"traffic,omitempty"`

	// TxBytes
	// Constraints:
	//    - nullable
	TxBytes *int `json:"txBytes,omitempty"`

	// TxDropDataFrames
	// Constraints:
	//    - nullable
	TxDropDataFrames *int `json:"txDropDataFrames,omitempty"`

	// TxFrames
	// Constraints:
	//    - nullable
	TxFrames *int `json:"txFrames,omitempty"`

	// TxRatebps
	// Constraints:
	//    - nullable
	TxRatebps *int `json:"txRatebps,omitempty"`

	// TxRxBytes
	// Constraints:
	//    - nullable
	TxRxBytes *int `json:"txRxBytes,omitempty"`

	// Uplink
	// Constraints:
	//    - nullable
	Uplink *int `json:"uplink,omitempty"`

	// UplinkRate
	// Constraints:
	//    - nullable
	UplinkRate *int `json:"uplinkRate,omitempty"`

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

	// ZoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`

	// ZoneVersion
	// Constraints:
	//    - nullable
	ZoneVersion *string `json:"zoneVersion,omitempty"`
}

func NewWSGClientQueryCreateClientQuery() *WSGClientQueryCreateClientQuery {
	m := new(WSGClientQueryCreateClientQuery)
	return m
}

type WSGClientQueryTcWithQuota struct {
	// TcMaxQuota
	// Constraints:
	//    - nullable
	TcMaxQuota *string `json:"tcMaxQuota,omitempty"`

	// TcName
	// Constraints:
	//    - nullable
	TcName *string `json:"tcName,omitempty"`

	// TcRemainingQuota
	// Constraints:
	//    - nullable
	TcRemainingQuota *string `json:"tcRemainingQuota,omitempty"`
}

func NewWSGClientQueryTcWithQuota() *WSGClientQueryTcWithQuota {
	m := new(WSGClientQueryTcWithQuota)
	return m
}
