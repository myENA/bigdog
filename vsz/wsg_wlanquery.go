package vsz

// API Version: v8_1

type WSGWLANQueryApWlanBssid struct {
	ApMac *string `json:"apMac,omitempty"`

	DeviceName *string `json:"deviceName,omitempty"`

	WlanBssids []*WSGWLANQueryWlanBssid `json:"wlanBssids,omitempty"`
}

type WSGWLANQueryApWlanBssidQueryList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANQueryApWlanBssid `json:"list,omitempty"`

	// TotalCount
	// Total matched AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGWLANQueryCreateWlanQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	// ApplicationVisibility
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ApplicationVisibility *string `json:"applicationVisibility,omitempty" validate:"oneof=Enabled Disabled"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthType *string `json:"authType,omitempty"`

	Clients *int `json:"clients,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	// Enability11k
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	Enability11k *string `json:"enability11k,omitempty" validate:"oneof=Enabled Disabled"`

	// Enability11r
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	Enability11r *string `json:"enability11r,omitempty" validate:"oneof=Enabled Disabled"`

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	Name *string `json:"name,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	// Status
	// Constraints:
	//    - oneof:[Online,Flagged,Offline]
	Status *string `json:"status,omitempty" validate:"oneof=Online Flagged Offline"`

	TenantDomainName *string `json:"tenantDomainName,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TrafficDownlink *int `json:"trafficDownlink,omitempty"`

	TrafficUplink *int `json:"trafficUplink,omitempty"`

	// Tunneled
	// Constraints:
	//    - oneof:[Tunneled,APBridged]
	Tunneled *string `json:"tunneled,omitempty" validate:"oneof=Tunneled APBridged"`

	Utp *string `json:"utp,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WepEncryptionStrength *int `json:"wepEncryptionStrength,omitempty"`

	WlanId *string `json:"wlanId,omitempty"`

	WpaVersion *string `json:"wpaVersion,omitempty"`

	// ZeroITEnabled
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ZeroITEnabled *string `json:"zeroITEnabled,omitempty" validate:"oneof=Enabled Disabled"`

	// ZeroITOnboard
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ZeroITOnboard *string `json:"zeroITOnboard,omitempty" validate:"oneof=Enabled Disabled"`

	ZoneName *string `json:"zoneName,omitempty"`
}

type WSGWLANQueryWlanBssid struct {
	Bssid *string `json:"bssid,omitempty"`

	RadioId *int `json:"radioId,omitempty"`

	WlanId *int `json:"wlanId,omitempty"`

	WlanName *string `json:"wlanName,omitempty"`
}

type WSGWLANQueryList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANQueryCreateWlanQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
