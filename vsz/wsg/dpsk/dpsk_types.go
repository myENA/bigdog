package dpsk

// API Version: v8_0

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

	// UserRoleID
	// Identity User Role ID of DPSK
	UserRoleID *string `json:"userRoleId,omitempty"`

	// VlanID
	// VLAN ID of DPSK
	VlanID *int `json:"vlanId,omitempty"`
}

type DeleteDpskResult struct {
	ResultCount *int `json:"resultCount,omitempty"`
}

type DeleteDPSKs struct {
	IDList []string `json:"idList,omitempty"`
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

type DpskQueryList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

	// WLANID
	// Identifier of the DPSK enabled WLAN
	WLANID *string `json:"wlanId,omitempty"`

	// WLANName
	// Name of the the DPSK enabled WLAN
	WLANName *string `json:"wlanName,omitempty"`
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

type WLANDpskSetting struct {
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

type WLANExternalDpsk struct {
	// AuthService
	// Autentication of the WLAN relate to external DPSK
	AuthService *WLANExternalDpskAuthServiceType `json:"authService,omitempty"`

	// Enabled
	// Enable External DPSK
	Enabled *bool `json:"enabled,omitempty"`

	// Encryption
	// Encryption of the WLAN relate to external DPSK
	Encryption *WLANExternalDpskEncryptionType `json:"encryption,omitempty"`
}

// WLANExternalDpskAuthServiceType
//
// Autentication of the WLAN relate to external DPSK
type WLANExternalDpskAuthServiceType struct {
	// ID
	// Identifier of the authentication service. At least one ID or name is required in the request.
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the authentication service. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`
}

// WLANExternalDpskEncryptionType
//
// Encryption of the WLAN relate to external DPSK
type WLANExternalDpskEncryptionType struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	Algorithm *string `json:"algorithm,omitempty"`

	// Method
	// Encryption method
	Method *string `json:"method,omitempty"`
}
