package dpsk

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type BatchGenUnbound struct {
	// Amount
	// Amount of generate unbound DPSK
	Amount *int `json:"amount,omitempty"`

	// GroupDPSK
	// Using group DPSK or not
	GroupDPSK *bool `json:"groupDpsk,omitempty"`

	PassphraseList []string `json:"passphraseList,omitempty"`

	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`

	// UserRoleID
	// Identity User Role ID of DPSK
	UserRoleID *string `json:"userRoleId,omitempty"`

	// VlanID
	// VLAN ID of DPSK
	VlanID *int `json:"vlanId,omitempty"`
}

type DeleteDPSKResult struct {
	ResultCount *int `json:"resultCount,omitempty"`
}

type DeleteDPSKs struct {
	IDList []string `json:"idList,omitempty"`
}

type DeleteExpiredDPSKConfig struct {
	// DeleteExpiredDPSK
	// Delete expired DPSK interval of the Zone.
	DeleteExpiredDPSK *string `json:"deleteExpiredDpsk,omitempty" validate:"oneof=Never AfterOneDay AfterSixMonths"`
}

// DPSKInfo
//
// Information list of DPSK
type DPSKInfo []*DPSKInfoType

type DPSKInfoType struct {
	// CreationDateTime
	// Creationd date/time of DPSK
	CreationDateTime *float64 `json:"creationDateTime,omitempty"`

	// ExpirationDateTime
	// Expiration date/time of DPSK
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// GroupDPSK
	// Is a Group DPSK or not
	GroupDPSK *bool `json:"groupDpsk,omitempty"`

	// ID
	// Identifier of DPSK
	ID *string `json:"id,omitempty"`

	// MacAddress
	// Mac address of DPSK
	MacAddress *string `json:"macAddress,omitempty"`

	// Passphrase
	// Passphrase of DPSK
	Passphrase *string `json:"passphrase,omitempty"`

	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`

	// UserRoleID
	// Identity User Role ID of DPSK
	UserRoleID *string `json:"userRoleId,omitempty"`

	// VlanID
	// VLAN ID of DPSK
	VlanID *int `json:"vlanId,omitempty"`

	// WLANID
	// WLAN ID of DPSK
	WLANID *string `json:"wlanId,omitempty"`
}

type DPSKQueryList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of DPSK
	List []*DPSKQueryListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DPSKQueryListType struct {
	// CreateDateTime
	// Creationd time of DPSK
	CreateDateTime *float64 `json:"createDateTime,omitempty"`

	// DomainID
	// Domain ID of DPSK
	DomainID *string `json:"domainId,omitempty"`

	// ExpirationStartTime
	// Expiration start time of DPSK (Unit: Seconds)
	ExpirationStartTime *float64 `json:"expirationStartTime,omitempty"`

	// ExpirationTime
	// Expiration time of DPSK (Unit: Seconds)
	ExpirationTime *float64 `json:"expirationTime,omitempty"`

	// Expired
	// DPSK is expired or not
	Expired *bool `json:"expired,omitempty"`

	// Group
	// Is a Group DPSK or not
	Group *bool `json:"group,omitempty"`

	// Key
	// Identifier of DPSK
	Key *string `json:"key,omitempty"`

	// TenantID
	// Tenant ID of DPSK
	TenantID *string `json:"tenantId,omitempty"`

	// TTL
	// Time To Live of DPSK (Unit: Seconds)
	TTL *float64 `json:"ttl,omitempty"`

	// UeMac
	// Mac address of DPSK
	UeMac *string `json:"ueMac,omitempty"`

	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`

	// UserRoleID
	// Identity User Role ID of DPSK
	UserRoleID *string `json:"userRoleId,omitempty"`

	// VlanID
	// VLAN ID of DPSK
	VlanID *int `json:"vlanId,omitempty"`

	// WLANID
	// WLAN ID of DPSK
	WLANID *string `json:"wlanId,omitempty"`

	// ZoneID
	// Zone ID of DPSK
	ZoneID *string `json:"zoneId,omitempty"`
}

type GetDPSKEnabledWlans struct {
	// FirstIndex
	// Index of the first DPSK enabled WLAN returned out of the complete WLAN list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more DPSK enabled WLANs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*GetDPSKEnabledWlansListType `json:"list,omitempty"`

	// TotalCount
	// Total DPSK enabled WLAN count of the zone
	TotalCount *int `json:"totalCount,omitempty"`
}

type GetDPSKEnabledWlansListType struct {
	// Ssid
	// SSID of the DPSK enabled WLAN
	Ssid *string `json:"ssid,omitempty"`

	// WLANID
	// Identifier of the DPSK enabled WLAN
	WLANID *string `json:"wlanId,omitempty"`

	// WLANName
	// Name of the the DPSK enabled WLAN
	WLANName *string `json:"wlanName,omitempty"`
}

type GetDPSKInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List DPSKInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type GetDPSKResult struct {
	DPSKInfoList DPSKInfo `json:"dpskInfoList,omitempty"`

	ResultCount *int `json:"resultCount,omitempty"`
}

type ModifyDeleteExpiredDPSK struct {
	// DeleteExpiredDPSK
	// Delete expired DPSK interval of the Zone.
	DeleteExpiredDPSK *string `json:"deleteExpiredDpsk,omitempty" validate:"oneof=Never AfterOneDay AfterSixMonths"`
}

type UpdateDPSK struct {
	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`
}

type WLANDPSKSetting struct {
	// DPSKEnabled
	// DPSK enabled
	DPSKEnabled *bool `json:"dpskEnabled,omitempty"`

	// DPSKFromType
	// Type of expiration start from
	DPSKFromType *string `json:"dpskFromType,omitempty" validate:"oneof=CreateTime FirstUse"`

	// DPSKType
	// Type of DPSK key
	DPSKType *string `json:"dpskType,omitempty" validate:"oneof=Secure KeyboardFriendly NumbersOnly"`

	// Expiration
	// Expiration of DPSK key
	Expiration *string `json:"expiration,omitempty" validate:"oneof=Unlimited OneDay TwoDays OneWeek TwoWeeks OneMonth SixMonths OneYear TwoYears"`

	// Length
	// Length of DPSK key
	Length *int `json:"length,omitempty"`
}

type WLANExternalDPSK struct {
	// AuthService
	// Autentication of the WLAN relate to external DPSK
	AuthService *WLANExternalDPSKAuthServiceType `json:"authService,omitempty"`

	// Enabled
	// Enable External DPSK
	Enabled *bool `json:"enabled,omitempty" validate:"required"`

	// Encryption
	// Encryption of the WLAN relate to external DPSK
	Encryption *WLANExternalDPSKEncryptionType `json:"encryption,omitempty"`
}

// WLANExternalDPSKAuthServiceType
//
// Autentication of the WLAN relate to external DPSK
type WLANExternalDPSKAuthServiceType struct {
	// ID
	// Identifier of the authentication service. At least one ID or name is required in the request.
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the authentication service. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`
}

// WLANExternalDPSKEncryptionType
//
// Encryption of the WLAN relate to external DPSK
type WLANExternalDPSKEncryptionType struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	Algorithm *string `json:"algorithm,omitempty" validate:"oneof=AES TKIP_AES"`

	// Method
	// Encryption method
	Method *string `json:"method,omitempty" validate:"oneof=WPA2 WPA_Mixed None"`
}
