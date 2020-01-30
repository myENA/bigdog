package vsz

// API Version: v9_0

type WSGClientQueryList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGClientQueryCreateClientQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGClientQueryList() *WSGClientQueryList {
	m := new(WSGClientQueryList)
	return m
}

type WSGClientQueryCreateClientQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthStatus *string `json:"authStatus,omitempty"`

	Bssid *string `json:"bssid,omitempty"`

	Channel *int `json:"channel,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	ControlPlaneName *string `json:"controlPlaneName,omitempty"`

	CpeMac *string `json:"cpeMac,omitempty"`

	DataPlaneName *string `json:"dataPlaneName,omitempty"`

	DeviceType *string `json:"deviceType,omitempty"`

	Downlink *int `json:"downlink,omitempty"`

	DownlinkRate *int `json:"downlinkRate,omitempty"`

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	IpAddress *string `json:"ipAddress,omitempty"`

	Ipv6Address *string `json:"ipv6Address,omitempty"`

	MedianRxMCSRate *int `json:"medianRxMCSRate,omitempty"`

	MedianTxMCSRate *int `json:"medianTxMCSRate,omitempty"`

	ModelName *string `json:"modelName,omitempty"`

	OsType *string `json:"osType,omitempty"`

	OsVendorType *string `json:"osVendorType,omitempty"`

	RadioType *string `json:"radioType,omitempty"`

	Role *string `json:"role,omitempty"`

	Rssi *int `json:"rssi,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	RxFrames *int `json:"rxFrames,omitempty"`

	SessionStartTime *int `json:"sessionStartTime,omitempty"`

	Snr *int `json:"snr,omitempty"`

	Speedflex *int `json:"speedflex,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Status *string `json:"status,omitempty"`

	TcWithQuotaList []*WSGClientQueryTcWithQuota `json:"tcWithQuotaList,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	TxDropDataFrames *int `json:"txDropDataFrames,omitempty"`

	TxFrames *int `json:"txFrames,omitempty"`

	TxRatebps *int `json:"txRatebps,omitempty"`

	TxRxBytes *int `json:"txRxBytes,omitempty"`

	Uplink *int `json:"uplink,omitempty"`

	UplinkRate *int `json:"uplinkRate,omitempty"`

	UserName *string `json:"userName,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WlanType *string `json:"wlanType,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`

	ZoneVersion *string `json:"zoneVersion,omitempty"`
}

func NewWSGClientQueryCreateClientQuery() *WSGClientQueryCreateClientQuery {
	m := new(WSGClientQueryCreateClientQuery)
	return m
}

type WSGClientQueryTcWithQuota struct {
	TcMaxQuota *string `json:"tcMaxQuota,omitempty"`

	TcName *string `json:"tcName,omitempty"`

	TcRemainingQuota *string `json:"tcRemainingQuota,omitempty"`
}

func NewWSGClientQueryTcWithQuota() *WSGClientQueryTcWithQuota {
	m := new(WSGClientQueryTcWithQuota)
	return m
}
