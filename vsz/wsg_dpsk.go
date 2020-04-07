package vsz

// API Version: v9_0

type WSGDPSKBatchGenUnbound struct {
	// Amount
	// Amount of generate unbound DPSK
	// Constraints:
	//    - nullable
	Amount *int `json:"amount,omitempty"`

	// GroupDpsk
	// Using group DPSK or not
	// Constraints:
	//    - nullable
	GroupDpsk *bool `json:"groupDpsk,omitempty"`

	// PassphraseList
	// Constraints:
	//    - nullable
	PassphraseList []string `json:"passphraseList,omitempty" validate:"omitempty,dive"`

	// UserName
	// User Name of DPSK
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`

	// UserRoleId
	// Identity User Role ID of DPSK
	// Constraints:
	//    - nullable
	UserRoleId *string `json:"userRoleId,omitempty"`

	// VlanId
	// VLAN ID of DPSK
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`
}

func NewWSGDPSKBatchGenUnbound() *WSGDPSKBatchGenUnbound {
	m := new(WSGDPSKBatchGenUnbound)
	return m
}

type WSGDPSKDeleteDpskResult struct {
	// ResultCount
	// Constraints:
	//    - nullable
	ResultCount *int `json:"resultCount,omitempty"`
}

func NewWSGDPSKDeleteDpskResult() *WSGDPSKDeleteDpskResult {
	m := new(WSGDPSKDeleteDpskResult)
	return m
}

type WSGDPSKDeleteDPSKs struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList []string `json:"idList,omitempty" validate:"omitempty,dive"`
}

func NewWSGDPSKDeleteDPSKs() *WSGDPSKDeleteDPSKs {
	m := new(WSGDPSKDeleteDPSKs)
	return m
}

type WSGDPSKDeleteExpiredDpskConfig struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	// Constraints:
	//    - nullable
	//    - oneof:[Never,AfterOneDay,AfterSixMonths]
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty" validate:"omitempty,oneof=Never AfterOneDay AfterSixMonths"`
}

func NewWSGDPSKDeleteExpiredDpskConfig() *WSGDPSKDeleteExpiredDpskConfig {
	m := new(WSGDPSKDeleteExpiredDpskConfig)
	return m
}

// WSGDPSKInfo
//
// Information list of DPSK
// Constraints:
//    - nullable
type WSGDPSKInfo []*WSGDPSKInfoType

func MakeWSGDPSKInfo() WSGDPSKInfo {
	m := make(WSGDPSKInfo, 0)
	return m
}

type WSGDPSKInfoType struct {
	// CreationDateTime
	// Creationd date/time of DPSK
	// Constraints:
	//    - nullable
	CreationDateTime *float64 `json:"creationDateTime,omitempty"`

	// ExpirationDateTime
	// Expiration date/time of DPSK
	// Constraints:
	//    - nullable
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// GroupDpsk
	// Is a Group DPSK or not
	// Constraints:
	//    - nullable
	GroupDpsk *bool `json:"groupDpsk,omitempty"`

	// Id
	// Identifier of DPSK
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// MacAddress
	// Mac address of DPSK
	// Constraints:
	//    - nullable
	MacAddress *string `json:"macAddress,omitempty"`

	// Passphrase
	// Passphrase of DPSK
	// Constraints:
	//    - nullable
	Passphrase *string `json:"passphrase,omitempty"`

	// UserName
	// User Name of DPSK
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`

	// UserRoleId
	// Identity User Role ID of DPSK
	// Constraints:
	//    - nullable
	UserRoleId *string `json:"userRoleId,omitempty"`

	// VlanId
	// VLAN ID of DPSK
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`

	// WlanId
	// WLAN ID of DPSK
	// Constraints:
	//    - nullable
	WlanId *string `json:"wlanId,omitempty"`
}

func NewWSGDPSKInfoType() *WSGDPSKInfoType {
	m := new(WSGDPSKInfoType)
	return m
}

type WSGDPSKQueryList struct {
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
	// Information list of DPSK
	// Constraints:
	//    - nullable
	List []*WSGDPSKQueryListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPSKQueryList() *WSGDPSKQueryList {
	m := new(WSGDPSKQueryList)
	return m
}

type WSGDPSKQueryListType struct {
	// CreateDateTime
	// Creationd time of DPSK
	// Constraints:
	//    - nullable
	CreateDateTime *float64 `json:"createDateTime,omitempty"`

	// DomainId
	// Domain ID of DPSK
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationStartTime
	// Expiration start time of DPSK (Unit: Seconds)
	// Constraints:
	//    - nullable
	ExpirationStartTime *float64 `json:"expirationStartTime,omitempty"`

	// ExpirationTime
	// Expiration time of DPSK (Unit: Seconds)
	// Constraints:
	//    - nullable
	ExpirationTime *float64 `json:"expirationTime,omitempty"`

	// Expired
	// DPSK is expired or not
	// Constraints:
	//    - nullable
	Expired *bool `json:"expired,omitempty"`

	// Group
	// Is a Group DPSK or not
	// Constraints:
	//    - nullable
	Group *bool `json:"group,omitempty"`

	// Key
	// Identifier of DPSK
	// Constraints:
	//    - nullable
	Key *string `json:"key,omitempty"`

	// TenantId
	// Tenant ID of DPSK
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Ttl
	// Time To Live of DPSK (Unit: Seconds)
	// Constraints:
	//    - nullable
	Ttl *float64 `json:"ttl,omitempty"`

	// UeMac
	// Mac address of DPSK
	// Constraints:
	//    - nullable
	UeMac *string `json:"ueMac,omitempty"`

	// UserName
	// User Name of DPSK
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`

	// UserRoleId
	// Identity User Role ID of DPSK
	// Constraints:
	//    - nullable
	UserRoleId *string `json:"userRoleId,omitempty"`

	// VlanId
	// VLAN ID of DPSK
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`

	// WlanId
	// WLAN ID of DPSK
	// Constraints:
	//    - nullable
	WlanId *string `json:"wlanId,omitempty"`

	// ZoneId
	// Zone ID of DPSK
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGDPSKQueryListType() *WSGDPSKQueryListType {
	m := new(WSGDPSKQueryListType)
	return m
}

type WSGDPSKGetDpskEnabledWlans struct {
	// FirstIndex
	// Index of the first DPSK enabled WLAN returned out of the complete WLAN list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more DPSK enabled WLANs after the list that is currently displayed
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGDPSKGetDpskEnabledWlansListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total DPSK enabled WLAN count of the zone
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPSKGetDpskEnabledWlans() *WSGDPSKGetDpskEnabledWlans {
	m := new(WSGDPSKGetDpskEnabledWlans)
	return m
}

type WSGDPSKGetDpskEnabledWlansListType struct {
	// Ssid
	// SSID of the DPSK enabled WLAN
	// Constraints:
	//    - nullable
	Ssid *string `json:"ssid,omitempty"`

	// WlanId
	// Identifier of the DPSK enabled WLAN
	// Constraints:
	//    - nullable
	WlanId *string `json:"wlanId,omitempty"`

	// WlanName
	// Name of the the DPSK enabled WLAN
	// Constraints:
	//    - nullable
	WlanName *string `json:"wlanName,omitempty"`
}

func NewWSGDPSKGetDpskEnabledWlansListType() *WSGDPSKGetDpskEnabledWlansListType {
	m := new(WSGDPSKGetDpskEnabledWlansListType)
	return m
}

type WSGDPSKGetDpskInfoList struct {
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
	List WSGDPSKInfo `json:"list,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPSKGetDpskInfoList() *WSGDPSKGetDpskInfoList {
	m := new(WSGDPSKGetDpskInfoList)
	return m
}

type WSGDPSKGetDpskResult struct {
	// DpskInfoList
	// Constraints:
	//    - nullable
	DpskInfoList WSGDPSKInfo `json:"dpskInfoList,omitempty"`

	// ResultCount
	// Constraints:
	//    - nullable
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
	//    - nullable
	//    - oneof:[Never,AfterOneDay,AfterSixMonths]
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty" validate:"omitempty,oneof=Never AfterOneDay AfterSixMonths"`
}

func NewWSGDPSKModifyDeleteExpiredDpsk() *WSGDPSKModifyDeleteExpiredDpsk {
	m := new(WSGDPSKModifyDeleteExpiredDpsk)
	return m
}

type WSGDPSKUpdateDpsk struct {
	// UserName
	// User Name of DPSK
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`
}

func NewWSGDPSKUpdateDpsk() *WSGDPSKUpdateDpsk {
	m := new(WSGDPSKUpdateDpsk)
	return m
}

type WSGDPSKWlanDpskSetting struct {
	// DpskEnabled
	// DPSK enabled
	// Constraints:
	//    - nullable
	//    - default:false
	DpskEnabled *bool `json:"dpskEnabled,omitempty"`

	// DpskFromType
	// Type of expiration start from
	// Constraints:
	//    - nullable
	//    - default:'CreateTime'
	//    - oneof:[CreateTime,FirstUse]
	DpskFromType *string `json:"dpskFromType,omitempty" validate:"omitempty,oneof=CreateTime FirstUse"`

	// DpskType
	// Type of DPSK key
	// Constraints:
	//    - nullable
	//    - default:'Secure'
	//    - oneof:[Secure,KeyboardFriendly,NumbersOnly]
	DpskType *string `json:"dpskType,omitempty" validate:"omitempty,oneof=Secure KeyboardFriendly NumbersOnly"`

	// Expiration
	// Expiration of DPSK key
	// Constraints:
	//    - nullable
	//    - default:'Unlimited'
	//    - oneof:[Unlimited,OneDay,TwoDays,OneWeek,TwoWeeks,OneMonth,SixMonths,OneYear,TwoYears]
	Expiration *string `json:"expiration,omitempty" validate:"omitempty,oneof=Unlimited OneDay TwoDays OneWeek TwoWeeks OneMonth SixMonths OneYear TwoYears"`

	// Length
	// Length of DPSK key
	// Constraints:
	//    - nullable
	//    - default:62
	Length *int `json:"length,omitempty"`
}

func NewWSGDPSKWlanDpskSetting() *WSGDPSKWlanDpskSetting {
	m := new(WSGDPSKWlanDpskSetting)
	return m
}

type WSGDPSKWlanExternalDpsk struct {
	// AuthService
	// Autentication of the WLAN relate to external DPSK
	// Constraints:
	//    - nullable
	AuthService *WSGDPSKWlanExternalDpskAuthServiceType `json:"authService,omitempty"`

	// Enabled
	// Enable External DPSK
	// Constraints:
	//    - required
	//    - default:false
	Enabled *bool `json:"enabled" validate:"required"`

	// Encryption
	// Encryption of the WLAN relate to external DPSK
	// Constraints:
	//    - nullable
	Encryption *WSGDPSKWlanExternalDpskEncryptionType `json:"encryption,omitempty"`
}

func NewWSGDPSKWlanExternalDpsk() *WSGDPSKWlanExternalDpsk {
	m := new(WSGDPSKWlanExternalDpsk)
	return m
}

// WSGDPSKWlanExternalDpskAuthServiceType
//
// Autentication of the WLAN relate to external DPSK
// Constraints:
//    - nullable
type WSGDPSKWlanExternalDpskAuthServiceType struct {
	// Id
	// Identifier of the authentication service. At least one ID or name is required in the request.
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the authentication service. At least one ID or name is required in the request.
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGDPSKWlanExternalDpskAuthServiceType() *WSGDPSKWlanExternalDpskAuthServiceType {
	m := new(WSGDPSKWlanExternalDpskAuthServiceType)
	return m
}

// WSGDPSKWlanExternalDpskEncryptionType
//
// Encryption of the WLAN relate to external DPSK
// Constraints:
//    - nullable
type WSGDPSKWlanExternalDpskEncryptionType struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	// Constraints:
	//    - nullable
	//    - oneof:[AES,TKIP_AES]
	Algorithm *string `json:"algorithm,omitempty" validate:"omitempty,oneof=AES TKIP_AES"`

	// Method
	// Encryption method
	// Constraints:
	//    - nullable
	//    - oneof:[WPA2,WPA_Mixed,None]
	Method *string `json:"method,omitempty" validate:"omitempty,oneof=WPA2 WPA_Mixed None"`

	// Mfp
	// Encryption mfp
	// Constraints:
	//    - nullable
	//    - oneof:[disabled,capable,required]
	Mfp *string `json:"mfp,omitempty" validate:"omitempty,oneof=disabled capable required"`

	// Passphrase
	// Passphrase. This only applies to WPA2 and WPA mixed mode.
	// Constraints:
	//    - nullable
	Passphrase *string `json:"passphrase,omitempty"`
}

func NewWSGDPSKWlanExternalDpskEncryptionType() *WSGDPSKWlanExternalDpskEncryptionType {
	m := new(WSGDPSKWlanExternalDpskEncryptionType)
	return m
}
