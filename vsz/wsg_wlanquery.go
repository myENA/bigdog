package vsz

// API Version: v9_0

type WSGWLANQueryApWlanBssid struct {
	// ApMac
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// DeviceName
	// Constraints:
	//    - nullable
	DeviceName *string `json:"deviceName,omitempty"`

	// WlanBssids
	// Constraints:
	//    - nullable
	WlanBssids []*WSGWLANQueryWlanBssid `json:"wlanBssids,omitempty" validate:"omitempty,dive"`
}

func NewWSGWLANQueryApWlanBssid() *WSGWLANQueryApWlanBssid {
	m := new(WSGWLANQueryApWlanBssid)
	return m
}

type WSGWLANQueryApWlanBssidQueryList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGWLANQueryApWlanBssid `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total matched AP count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWLANQueryApWlanBssidQueryList() *WSGWLANQueryApWlanBssidQueryList {
	m := new(WSGWLANQueryApWlanBssidQueryList)
	return m
}

type WSGWLANQueryCreateWlanQuery struct {
	// Alerts
	// Constraints:
	//    - nullable
	Alerts *int `json:"alerts,omitempty"`

	// ApplicationVisibility
	// Constraints:
	//    - nullable
	//    - oneof:[Enabled,Disabled]
	ApplicationVisibility *string `json:"applicationVisibility,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// AuthMethod
	// Constraints:
	//    - nullable
	AuthMethod *string `json:"authMethod,omitempty"`

	// AuthType
	// Constraints:
	//    - nullable
	AuthType *string `json:"authType,omitempty"`

	// Clients
	// Constraints:
	//    - nullable
	Clients *int `json:"clients,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DomainName
	// Constraints:
	//    - nullable
	DomainName *string `json:"domainName,omitempty"`

	// Enability11k
	// Constraints:
	//    - nullable
	//    - oneof:[Enabled,Disabled]
	Enability11k *string `json:"enability11k,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// Enability11r
	// Constraints:
	//    - nullable
	//    - oneof:[Enabled,Disabled]
	Enability11r *string `json:"enability11r,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// EncryptionMethod
	// Constraints:
	//    - nullable
	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	// FirewallProfileId
	// Constraints:
	//    - nullable
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Ssid
	// Constraints:
	//    - nullable
	Ssid *string `json:"ssid,omitempty"`

	// Status
	// Constraints:
	//    - nullable
	//    - oneof:[Online,Flagged,Offline]
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=Online Flagged Offline"`

	// TenantDomainName
	// Constraints:
	//    - nullable
	TenantDomainName *string `json:"tenantDomainName,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Traffic
	// Constraints:
	//    - nullable
	Traffic *int `json:"traffic,omitempty"`

	// TrafficDownlink
	// Constraints:
	//    - nullable
	TrafficDownlink *int `json:"trafficDownlink,omitempty"`

	// TrafficUplink
	// Constraints:
	//    - nullable
	TrafficUplink *int `json:"trafficUplink,omitempty"`

	// Tunneled
	// Constraints:
	//    - nullable
	//    - oneof:[Tunneled,APBridged]
	Tunneled *string `json:"tunneled,omitempty" validate:"omitempty,oneof=Tunneled APBridged"`

	// Utp
	// Constraints:
	//    - nullable
	Utp *string `json:"utp,omitempty"`

	// Vlan
	// Constraints:
	//    - nullable
	Vlan *int `json:"vlan,omitempty"`

	// WepEncryptionStrength
	// Constraints:
	//    - nullable
	WepEncryptionStrength *int `json:"wepEncryptionStrength,omitempty"`

	// WlanId
	// Constraints:
	//    - nullable
	WlanId *string `json:"wlanId,omitempty"`

	// WpaVersion
	// Constraints:
	//    - nullable
	WpaVersion *string `json:"wpaVersion,omitempty"`

	// ZeroITEnabled
	// Constraints:
	//    - nullable
	//    - oneof:[Enabled,Disabled]
	ZeroITEnabled *string `json:"zeroITEnabled,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// ZeroITOnboard
	// Constraints:
	//    - nullable
	//    - oneof:[Enabled,Disabled]
	ZeroITOnboard *string `json:"zeroITOnboard,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// ZoneName
	// Constraints:
	//    - nullable
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGWLANQueryCreateWlanQuery() *WSGWLANQueryCreateWlanQuery {
	m := new(WSGWLANQueryCreateWlanQuery)
	return m
}

type WSGWLANQueryWlanBssid struct {
	// Bssid
	// Constraints:
	//    - nullable
	Bssid *string `json:"bssid,omitempty"`

	// RadioId
	// Constraints:
	//    - nullable
	RadioId *int `json:"radioId,omitempty"`

	// WlanId
	// Constraints:
	//    - nullable
	WlanId *int `json:"wlanId,omitempty"`

	// WlanName
	// Constraints:
	//    - nullable
	WlanName *string `json:"wlanName,omitempty"`
}

func NewWSGWLANQueryWlanBssid() *WSGWLANQueryWlanBssid {
	m := new(WSGWLANQueryWlanBssid)
	return m
}

type WSGWLANQueryList struct {
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
	List []*WSGWLANQueryCreateWlanQuery `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWLANQueryList() *WSGWLANQueryList {
	m := new(WSGWLANQueryList)
	return m
}
