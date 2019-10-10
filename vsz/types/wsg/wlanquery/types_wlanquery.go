package wlanquery

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApWlanBssid struct {
	ApMac *string `json:"apMac,omitempty"`

	DeviceName *string `json:"deviceName,omitempty"`

	WlanBssids []*WlanBssid `json:"wlanBssids,omitempty"`
}

type ApWlanBssidQueryList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApWlanBssid `json:"list,omitempty"`

	// TotalCount
	// Total matched AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

type CreateWlanQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	ApplicationVisibility *string `json:"applicationVisibility,omitempty" validate:"oneof=Enabled Disabled"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthType *string `json:"authType,omitempty"`

	Clients *int `json:"clients,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	Enability11k *string `json:"enability11k,omitempty" validate:"oneof=Enabled Disabled"`

	Enability11r *string `json:"enability11r,omitempty" validate:"oneof=Enabled Disabled"`

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	Name *string `json:"name,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Status *string `json:"status,omitempty" validate:"oneof=Online Flagged Offline"`

	TenantDomainName *string `json:"tenantDomainName,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TrafficDownlink *int `json:"trafficDownlink,omitempty"`

	TrafficUplink *int `json:"trafficUplink,omitempty"`

	Tunneled *string `json:"tunneled,omitempty" validate:"oneof=Tunneled APBridged"`

	Utp *string `json:"utp,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WepEncryptionStrength *int `json:"wepEncryptionStrength,omitempty"`

	WlanId *string `json:"wlanId,omitempty"`

	WpaVersion *string `json:"wpaVersion,omitempty"`

	ZeroITEnabled *string `json:"zeroITEnabled,omitempty" validate:"oneof=Enabled Disabled"`

	ZeroITOnboard *string `json:"zeroITOnboard,omitempty" validate:"oneof=Enabled Disabled"`

	ZoneName *string `json:"zoneName,omitempty"`
}

type WlanBssid struct {
	Bssid *string `json:"bssid,omitempty"`

	RadioId *int `json:"radioId,omitempty"`

	WlanId *int `json:"wlanId,omitempty"`

	WlanName *string `json:"wlanName,omitempty"`
}

type WlanQueryList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateWlanQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
