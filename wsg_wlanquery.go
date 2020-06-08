package ruckus

// API Version: v9_0

type WSGWLANQueryApWlanBssid struct {
	ApMac *string `json:"apMac,omitempty"`

	DeviceName *string `json:"deviceName,omitempty"`

	WlanBssids []*WSGWLANQueryApWlanBssid `json:"wlanBssids,omitempty"`
}

func NewWSGWLANQueryApWlanBssid() *WSGWLANQueryApWlanBssid {
	m := new(WSGWLANQueryApWlanBssid)
	return m
}

type WSGWLANQueryApWlanBssidQueryList struct {
	Extra *WSGWLANQueryApWlanBssidQueryList `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANQueryApWlanBssidQueryList `json:"list,omitempty"`

	// TotalCount
	// Total matched AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWLANQueryApWlanBssidQueryList() *WSGWLANQueryApWlanBssidQueryList {
	m := new(WSGWLANQueryApWlanBssidQueryList)
	return m
}

type WSGWLANQueryCreateWlanQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	// ApplicationVisibility
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ApplicationVisibility *string `json:"applicationVisibility,omitempty"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthType *string `json:"authType,omitempty"`

	Clients *int `json:"clients,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	// Enability11k
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	Enability11k *string `json:"enability11k,omitempty"`

	// Enability11r
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	Enability11r *string `json:"enability11r,omitempty"`

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	Name *string `json:"name,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	// Status
	// Constraints:
	//    - oneof:[Online,Flagged,Offline]
	Status *string `json:"status,omitempty"`

	TenantDomainName *string `json:"tenantDomainName,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TrafficDownlink *int `json:"trafficDownlink,omitempty"`

	TrafficUplink *int `json:"trafficUplink,omitempty"`

	// Tunneled
	// Constraints:
	//    - oneof:[Tunneled,APBridged]
	Tunneled *string `json:"tunneled,omitempty"`

	Utp *string `json:"utp,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WepEncryptionStrength *int `json:"wepEncryptionStrength,omitempty"`

	WlanId *string `json:"wlanId,omitempty"`

	WpaVersion *string `json:"wpaVersion,omitempty"`

	// ZeroITEnabled
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ZeroITEnabled *string `json:"zeroITEnabled,omitempty"`

	// ZeroITOnboard
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ZeroITOnboard *string `json:"zeroITOnboard,omitempty"`

	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGWLANQueryCreateWlanQuery() *WSGWLANQueryCreateWlanQuery {
	m := new(WSGWLANQueryCreateWlanQuery)
	return m
}

type WSGWLANQueryWlanBssid struct {
	Bssid *string `json:"bssid,omitempty"`

	RadioId *int `json:"radioId,omitempty"`

	WlanId *int `json:"wlanId,omitempty"`

	WlanName *string `json:"wlanName,omitempty"`
}

func NewWSGWLANQueryWlanBssid() *WSGWLANQueryWlanBssid {
	m := new(WSGWLANQueryWlanBssid)
	return m
}

type WSGWLANQueryList struct {
	Extra *WSGWLANQueryList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANQueryList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWLANQueryList() *WSGWLANQueryList {
	m := new(WSGWLANQueryList)
	return m
}
