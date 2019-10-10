package wlanquery

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApWLANBssid struct {
	ApMac *string `json:"apMac,omitempty"`

	DeviceName *string `json:"deviceName,omitempty"`

	WLANBssids []*WLANBssid `json:"wlanBssids,omitempty"`
}

type ApWLANBssidQueryList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApWLANBssid `json:"list,omitempty"`

	// TotalCount
	// Total matched AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

type CreateWLANQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	ApplicationVisibility *string `json:"applicationVisibility,omitempty" validate:"oneof=Enabled Disabled"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthType *string `json:"authType,omitempty"`

	Clients *int `json:"clients,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	Enability11K *string `json:"enability11k,omitempty" validate:"oneof=Enabled Disabled"`

	Enability11R *string `json:"enability11r,omitempty" validate:"oneof=Enabled Disabled"`

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	Name *string `json:"name,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Status *string `json:"status,omitempty" validate:"oneof=Online Flagged Offline"`

	TenantDomainName *string `json:"tenantDomainName,omitempty"`

	TenantID *string `json:"tenantId,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TrafficDownlink *int `json:"trafficDownlink,omitempty"`

	TrafficUplink *int `json:"trafficUplink,omitempty"`

	Tunneled *string `json:"tunneled,omitempty" validate:"oneof=Tunneled APBridged"`

	Utp *string `json:"utp,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WepEncryptionStrength *int `json:"wepEncryptionStrength,omitempty"`

	WLANID *string `json:"wlanId,omitempty"`

	WpaVersion *string `json:"wpaVersion,omitempty"`

	ZeroITEnabled *string `json:"zeroITEnabled,omitempty" validate:"oneof=Enabled Disabled"`

	ZeroITOnboard *string `json:"zeroITOnboard,omitempty" validate:"oneof=Enabled Disabled"`

	ZoneName *string `json:"zoneName,omitempty"`
}

type WLANBssid struct {
	Bssid *string `json:"bssid,omitempty"`

	RadioID *int `json:"radioId,omitempty"`

	WLANID *int `json:"wlanId,omitempty"`

	WLANName *string `json:"wlanName,omitempty"`
}

type WLANQueryList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateWLANQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
