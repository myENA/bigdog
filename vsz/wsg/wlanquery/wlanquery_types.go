package wlanquery

// API Version: v8_0

type ApWLANBssid struct {
	ApMac      *string      `json:"apMac,omitempty"`
	DeviceName *string      `json:"deviceName,omitempty"`
	WLANBssids []*WLANBssid `json:"wlanBssids,omitempty"`
}

type ApWLANBssidQueryList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*ApWLANBssid       `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type CreateWLANQuery struct {
	Alerts                *int    `json:"alerts,omitempty"`
	ApplicationVisibility *string `json:"applicationVisibility,omitempty"`
	AuthMethod            *string `json:"authMethod,omitempty"`
	AuthType              *string `json:"authType,omitempty"`
	Clients               *int    `json:"clients,omitempty"`
	Description           *string `json:"description,omitempty"`
	DomainName            *string `json:"domainName,omitempty"`
	Enability1K           *string `json:"enability11k,omitempty"`
	Enability1R           *string `json:"enability11r,omitempty"`
	EncryptionMethod      *string `json:"encryptionMethod,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Ssid                  *string `json:"ssid,omitempty"`
	Status                *string `json:"status,omitempty"`
	TenantDomainName      *string `json:"tenantDomainName,omitempty"`
	TenantID              *string `json:"tenantId,omitempty"`
	Traffic               *int    `json:"traffic,omitempty"`
	TrafficDownlink       *int    `json:"trafficDownlink,omitempty"`
	TrafficUplink         *int    `json:"trafficUplink,omitempty"`
	Tunneled              *string `json:"tunneled,omitempty"`
	Utp                   *string `json:"utp,omitempty"`
	Vlan                  *int    `json:"vlan,omitempty"`
	WepEncryptionStrength *int    `json:"wepEncryptionStrength,omitempty"`
	WLANID                *string `json:"wlanId,omitempty"`
	WpaVersion            *string `json:"wpaVersion,omitempty"`
	ZeroITEnabled         *string `json:"zeroITEnabled,omitempty"`
	ZeroITOnboard         *string `json:"zeroITOnboard,omitempty"`
	ZoneName              *string `json:"zoneName,omitempty"`
}

type WLANBssid struct {
	Bssid    *string `json:"bssid,omitempty"`
	RadioID  *int    `json:"radioId,omitempty"`
	WLANID   *int    `json:"wlanId,omitempty"`
	WLANName *string `json:"wlanName,omitempty"`
}

type WLANQueryList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*CreateWLANQuery   `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}
