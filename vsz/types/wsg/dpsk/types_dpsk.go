package dpsk

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type BatchGenUnbound struct {
	// Amount
	// Amount of generate unbound DPSK
	Amount *int `json:"amount,omitempty"`

	// GroupDpsk
	// Using group DPSK or not
	GroupDpsk *bool `json:"groupDpsk,omitempty"`

	PassphraseList []string `json:"passphraseList,omitempty"`

	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`

	// UserRoleId
	// Identity User Role ID of DPSK
	UserRoleId *string `json:"userRoleId,omitempty"`

	// VlanId
	// VLAN ID of DPSK
	VlanId *int `json:"vlanId,omitempty"`
}

type DeleteDpskResult struct {
	ResultCount *int `json:"resultCount,omitempty"`
}

type DeleteDPSKs struct {
	IdList []string `json:"idList,omitempty"`
}

type DeleteExpiredDpskConfig struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
}

// DpskInfo
//
// Information list of DPSK
type DpskInfo []*DpskInfoType

type DpskInfoType struct {
	// CreationDateTime
	// Creationd date/time of DPSK
	CreationDateTime *float64 `json:"creationDateTime,omitempty"`

	// ExpirationDateTime
	// Expiration date/time of DPSK
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// GroupDpsk
	// Is a Group DPSK or not
	GroupDpsk *bool `json:"groupDpsk,omitempty"`

	// Id
	// Identifier of DPSK
	Id *string `json:"id,omitempty"`

	// MacAddress
	// Mac address of DPSK
	MacAddress *string `json:"macAddress,omitempty"`

	// Passphrase
	// Passphrase of DPSK
	Passphrase *string `json:"passphrase,omitempty"`

	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`

	// UserRoleId
	// Identity User Role ID of DPSK
	UserRoleId *string `json:"userRoleId,omitempty"`

	// VlanId
	// VLAN ID of DPSK
	VlanId *int `json:"vlanId,omitempty"`

	// WlanId
	// WLAN ID of DPSK
	WlanId *string `json:"wlanId,omitempty"`
}

type DpskQueryList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of DPSK
	List []*DpskQueryListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DpskQueryListType struct {
	// CreateDateTime
	// Creationd time of DPSK
	CreateDateTime *float64 `json:"createDateTime,omitempty"`

	// DomainId
	// Domain ID of DPSK
	DomainId *string `json:"domainId,omitempty"`

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

	// TenantId
	// Tenant ID of DPSK
	TenantId *string `json:"tenantId,omitempty"`

	// Ttl
	// Time To Live of DPSK (Unit: Seconds)
	Ttl *float64 `json:"ttl,omitempty"`

	// UeMac
	// Mac address of DPSK
	UeMac *string `json:"ueMac,omitempty"`

	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`

	// UserRoleId
	// Identity User Role ID of DPSK
	UserRoleId *string `json:"userRoleId,omitempty"`

	// VlanId
	// VLAN ID of DPSK
	VlanId *int `json:"vlanId,omitempty"`

	// WlanId
	// WLAN ID of DPSK
	WlanId *string `json:"wlanId,omitempty"`

	// ZoneId
	// Zone ID of DPSK
	ZoneId *string `json:"zoneId,omitempty"`
}

type GetDpskEnabledWlans struct {
	// FirstIndex
	// Index of the first DPSK enabled WLAN returned out of the complete WLAN list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more DPSK enabled WLANs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*GetDpskEnabledWlansListType `json:"list,omitempty"`

	// TotalCount
	// Total DPSK enabled WLAN count of the zone
	TotalCount *int `json:"totalCount,omitempty"`
}

type GetDpskEnabledWlansListType struct {
	// Ssid
	// SSID of the DPSK enabled WLAN
	Ssid *string `json:"ssid,omitempty"`

	// WlanId
	// Identifier of the DPSK enabled WLAN
	WlanId *string `json:"wlanId,omitempty"`

	// WlanName
	// Name of the the DPSK enabled WLAN
	WlanName *string `json:"wlanName,omitempty"`
}

type GetDpskInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List DpskInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type GetDpskResult struct {
	DpskInfoList DpskInfo `json:"dpskInfoList,omitempty"`

	ResultCount *int `json:"resultCount,omitempty"`
}

type ModifyDeleteExpiredDpsk struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
}

type UpdateDpsk struct {
	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`
}

type WlanDpskSetting struct {
	// DpskEnabled
	// DPSK enabled
	DpskEnabled *bool `json:"dpskEnabled,omitempty"`

	// DpskFromType
	// Type of expiration start from
	DpskFromType *string `json:"dpskFromType,omitempty"`

	// DpskType
	// Type of DPSK key
	DpskType *string `json:"dpskType,omitempty"`

	// Expiration
	// Expiration of DPSK key
	Expiration *string `json:"expiration,omitempty"`

	// Length
	// Length of DPSK key
	Length *int `json:"length,omitempty"`
}

type WlanExternalDpsk struct {
	// AuthService
	// Autentication of the WLAN relate to external DPSK
	AuthService *WlanExternalDpskAuthServiceType `json:"authService,omitempty"`

	// Enabled
	// Enable External DPSK
	Enabled *bool `json:"enabled,omitempty"`

	// Encryption
	// Encryption of the WLAN relate to external DPSK
	Encryption *WlanExternalDpskEncryptionType `json:"encryption,omitempty"`
}

// WlanExternalDpskAuthServiceType
//
// Autentication of the WLAN relate to external DPSK
type WlanExternalDpskAuthServiceType struct {
	// Id
	// Identifier of the authentication service. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the authentication service. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`
}

// WlanExternalDpskEncryptionType
//
// Encryption of the WLAN relate to external DPSK
type WlanExternalDpskEncryptionType struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	Algorithm *string `json:"algorithm,omitempty"`

	// Method
	// Encryption method
	Method *string `json:"method,omitempty"`
}
