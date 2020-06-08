package ruckus

// API Version: v9_0

type WSGDPSKBatchGenUnbound struct {
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

func NewWSGDPSKBatchGenUnbound() *WSGDPSKBatchGenUnbound {
	m := new(WSGDPSKBatchGenUnbound)
	return m
}

type WSGDPSKDeleteDpskResult struct {
	ResultCount *int `json:"resultCount,omitempty"`
}

func NewWSGDPSKDeleteDpskResult() *WSGDPSKDeleteDpskResult {
	m := new(WSGDPSKDeleteDpskResult)
	return m
}

type WSGDPSKDeleteDPSKs struct {
	IdList []string `json:"idList,omitempty"`
}

func NewWSGDPSKDeleteDPSKs() *WSGDPSKDeleteDPSKs {
	m := new(WSGDPSKDeleteDPSKs)
	return m
}

type WSGDPSKDeleteExpiredDpskConfig struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	// Constraints:
	//    - oneof:[Never,AfterOneDay,AfterSixMonths]
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
}

func NewWSGDPSKDeleteExpiredDpskConfig() *WSGDPSKDeleteExpiredDpskConfig {
	m := new(WSGDPSKDeleteExpiredDpskConfig)
	return m
}

// WSGDPSKInfo
//
// Information list of DPSK
type WSGDPSKInfo []*WSGDPSKInfoType

func MakeWSGDPSKInfo() WSGDPSKInfo {
	m := make(WSGDPSKInfo, 0)
	return m
}

type WSGDPSKInfoType struct {
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

func NewWSGDPSKInfoType() *WSGDPSKInfoType {
	m := new(WSGDPSKInfoType)
	return m
}

type WSGDPSKQueryList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of DPSK
	List []*WSGDPSKQueryListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPSKQueryList() *WSGDPSKQueryList {
	m := new(WSGDPSKQueryList)
	return m
}

type WSGDPSKQueryListType struct {
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

func NewWSGDPSKQueryListType() *WSGDPSKQueryListType {
	m := new(WSGDPSKQueryListType)
	return m
}

type WSGDPSKGetDpskEnabledWlans struct {
	// FirstIndex
	// Index of the first DPSK enabled WLAN returned out of the complete WLAN list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more DPSK enabled WLANs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPSKGetDpskEnabledWlansListType `json:"list,omitempty"`

	// TotalCount
	// Total DPSK enabled WLAN count of the zone
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPSKGetDpskEnabledWlans() *WSGDPSKGetDpskEnabledWlans {
	m := new(WSGDPSKGetDpskEnabledWlans)
	return m
}

type WSGDPSKGetDpskEnabledWlansListType struct {
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

func NewWSGDPSKGetDpskEnabledWlansListType() *WSGDPSKGetDpskEnabledWlansListType {
	m := new(WSGDPSKGetDpskEnabledWlansListType)
	return m
}

type WSGDPSKGetDpskInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List WSGDPSKInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPSKGetDpskInfoList() *WSGDPSKGetDpskInfoList {
	m := new(WSGDPSKGetDpskInfoList)
	return m
}

type WSGDPSKGetDpskResult struct {
	DpskInfoList WSGDPSKInfo `json:"dpskInfoList,omitempty"`

	ResultCount *int `json:"resultCount,omitempty"`
}

func NewWSGDPSKGetDpskResult() *WSGDPSKGetDpskResult {
	m := new(WSGDPSKGetDpskResult)
	return m
}

type WSGDPSKModifyDeleteExpiredDpsk struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	// Constraints:
	//    - oneof:[Never,AfterOneDay,AfterSixMonths]
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
}

func NewWSGDPSKModifyDeleteExpiredDpsk() *WSGDPSKModifyDeleteExpiredDpsk {
	m := new(WSGDPSKModifyDeleteExpiredDpsk)
	return m
}

type WSGDPSKUpdateDpsk struct {
	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`
}

func NewWSGDPSKUpdateDpsk() *WSGDPSKUpdateDpsk {
	m := new(WSGDPSKUpdateDpsk)
	return m
}

type WSGDPSKWlanDpskSetting struct {
	// DpskEnabled
	// DPSK enabled
	DpskEnabled *bool `json:"dpskEnabled,omitempty"`

	// DpskFromType
	// Type of expiration start from
	// Constraints:
	//    - default:'CreateTime'
	//    - oneof:[CreateTime,FirstUse]
	DpskFromType *string `json:"dpskFromType,omitempty"`

	// DpskType
	// Type of DPSK key
	// Constraints:
	//    - default:'Secure'
	//    - oneof:[Secure,KeyboardFriendly,NumbersOnly]
	DpskType *string `json:"dpskType,omitempty"`

	// Expiration
	// Expiration of DPSK key
	// Constraints:
	//    - default:'Unlimited'
	//    - oneof:[Unlimited,OneDay,TwoDays,OneWeek,TwoWeeks,OneMonth,SixMonths,OneYear,TwoYears]
	Expiration *string `json:"expiration,omitempty"`

	// Length
	// Length of DPSK key
	Length *int `json:"length,omitempty"`
}

func NewWSGDPSKWlanDpskSetting() *WSGDPSKWlanDpskSetting {
	m := new(WSGDPSKWlanDpskSetting)
	return m
}

type WSGDPSKWlanExternalDpsk struct {
	// AuthService
	// Autentication of the WLAN relate to external DPSK
	AuthService *WSGDPSKWlanExternalDpskAuthServiceType `json:"authService,omitempty"`

	// Enabled
	// Enable External DPSK
	// Constraints:
	//    - required
	//    - default:false
	Enabled *bool `json:"enabled"`

	// Encryption
	// Encryption of the WLAN relate to external DPSK
	Encryption *WSGDPSKWlanExternalDpskEncryptionType `json:"encryption,omitempty"`
}

func NewWSGDPSKWlanExternalDpsk() *WSGDPSKWlanExternalDpsk {
	m := new(WSGDPSKWlanExternalDpsk)
	return m
}

// WSGDPSKWlanExternalDpskAuthServiceType
//
// Autentication of the WLAN relate to external DPSK
type WSGDPSKWlanExternalDpskAuthServiceType struct {
	// Id
	// Identifier of the authentication service. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the authentication service. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`
}

func NewWSGDPSKWlanExternalDpskAuthServiceType() *WSGDPSKWlanExternalDpskAuthServiceType {
	m := new(WSGDPSKWlanExternalDpskAuthServiceType)
	return m
}

// WSGDPSKWlanExternalDpskEncryptionType
//
// Encryption of the WLAN relate to external DPSK
type WSGDPSKWlanExternalDpskEncryptionType struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	// Constraints:
	//    - oneof:[AES,TKIP_AES]
	Algorithm *string `json:"algorithm,omitempty"`

	// Method
	// Encryption method
	// Constraints:
	//    - oneof:[WPA2,WPA_Mixed,None]
	Method *string `json:"method,omitempty"`

	// Mfp
	// Encryption mfp
	// Constraints:
	//    - oneof:[disabled,capable,required]
	Mfp *string `json:"mfp,omitempty"`

	// Passphrase
	// Passphrase. This only applies to WPA2 and WPA mixed mode.
	Passphrase *string `json:"passphrase,omitempty"`
}

func NewWSGDPSKWlanExternalDpskEncryptionType() *WSGDPSKWlanExternalDpskEncryptionType {
	m := new(WSGDPSKWlanExternalDpskEncryptionType)
	return m
}
