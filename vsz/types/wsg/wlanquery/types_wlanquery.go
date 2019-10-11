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

func NewApWlanBssid() *ApWlanBssid {
	apWlanBssidType := new(ApWlanBssid)
	return apWlanBssidType
}

func NewDefaultApWlanBssid() *ApWlanBssid {
	apWlanBssidType := new(ApWlanBssid)
	return apWlanBssidType
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

func NewApWlanBssidQueryList() *ApWlanBssidQueryList {
	apWlanBssidQueryListType := new(ApWlanBssidQueryList)
	return apWlanBssidQueryListType
}

func NewDefaultApWlanBssidQueryList() *ApWlanBssidQueryList {
	apWlanBssidQueryListType := new(ApWlanBssidQueryList)
	return apWlanBssidQueryListType
}

type CreateWlanQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	// ApplicationVisibility
	// Constraints:
	//    - nullable
	//    - oneof:[Enabled,Disabled]
	ApplicationVisibility *string `json:"applicationVisibility,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthType *string `json:"authType,omitempty"`

	Clients *int `json:"clients,omitempty"`

	Description *string `json:"description,omitempty"`

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

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	Name *string `json:"name,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	// Status
	// Constraints:
	//    - nullable
	//    - oneof:[Online,Flagged,Offline]
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=Online Flagged Offline"`

	TenantDomainName *string `json:"tenantDomainName,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TrafficDownlink *int `json:"trafficDownlink,omitempty"`

	TrafficUplink *int `json:"trafficUplink,omitempty"`

	// Tunneled
	// Constraints:
	//    - nullable
	//    - oneof:[Tunneled,APBridged]
	Tunneled *string `json:"tunneled,omitempty" validate:"omitempty,oneof=Tunneled APBridged"`

	Utp *string `json:"utp,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WepEncryptionStrength *int `json:"wepEncryptionStrength,omitempty"`

	WlanId *string `json:"wlanId,omitempty"`

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

	ZoneName *string `json:"zoneName,omitempty"`
}

func NewCreateWlanQuery() *CreateWlanQuery {
	createWlanQueryType := new(CreateWlanQuery)
	return createWlanQueryType
}

func NewDefaultCreateWlanQuery() *CreateWlanQuery {
	createWlanQueryType := new(CreateWlanQuery)
	return createWlanQueryType
}

type WlanBssid struct {
	Bssid *string `json:"bssid,omitempty"`

	RadioId *int `json:"radioId,omitempty"`

	WlanId *int `json:"wlanId,omitempty"`

	WlanName *string `json:"wlanName,omitempty"`
}

func NewWlanBssid() *WlanBssid {
	wlanBssidType := new(WlanBssid)
	return wlanBssidType
}

func NewDefaultWlanBssid() *WlanBssid {
	wlanBssidType := new(WlanBssid)
	return wlanBssidType
}

type WlanQueryList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateWlanQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWlanQueryList() *WlanQueryList {
	wlanQueryListType := new(WlanQueryList)
	return wlanQueryListType
}

func NewDefaultWlanQueryList() *WlanQueryList {
	wlanQueryListType := new(WlanQueryList)
	return wlanQueryListType
}
