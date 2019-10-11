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

func NewBatchGenUnbound() *BatchGenUnbound {
	batchGenUnboundType := new(BatchGenUnbound)
	return batchGenUnboundType
}

func NewDefaultBatchGenUnbound() *BatchGenUnbound {
	batchGenUnboundType := new(BatchGenUnbound)
	return batchGenUnboundType
}

type DeleteDpskResult struct {
	ResultCount *int `json:"resultCount,omitempty"`
}

func NewDeleteDpskResult() *DeleteDpskResult {
	deleteDpskResultType := new(DeleteDpskResult)
	return deleteDpskResultType
}

func NewDefaultDeleteDpskResult() *DeleteDpskResult {
	deleteDpskResultType := new(DeleteDpskResult)
	return deleteDpskResultType
}

type DeleteDPSKs struct {
	IdList []string `json:"idList,omitempty"`
}

func NewDeleteDPSKs() *DeleteDPSKs {
	deleteDPSKsType := new(DeleteDPSKs)
	return deleteDPSKsType
}

func NewDefaultDeleteDPSKs() *DeleteDPSKs {
	deleteDPSKsType := new(DeleteDPSKs)
	return deleteDPSKsType
}

type DeleteExpiredDpskConfig struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	// Constraints:
	//    - nullable
	//    - oneof:[Never,AfterOneDay,AfterSixMonths]
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty" validate:"omitempty,oneof=Never AfterOneDay AfterSixMonths"`
}

func NewDeleteExpiredDpskConfig() *DeleteExpiredDpskConfig {
	deleteExpiredDpskConfigType := new(DeleteExpiredDpskConfig)
	return deleteExpiredDpskConfigType
}

func NewDefaultDeleteExpiredDpskConfig() *DeleteExpiredDpskConfig {
	deleteExpiredDpskConfigType := new(DeleteExpiredDpskConfig)
	return deleteExpiredDpskConfigType
}

// DpskInfo
//
// Information list of DPSK
type DpskInfo []*DpskInfoType

func NewDpskInfo() *DpskInfo {
	dpskInfoType := make(DpskInfo, 0)
	return &dpskInfoType
}

func NewDefaultDpskInfo() *DpskInfo {
	dpskInfoType := make(DpskInfo, 0)
	return &dpskInfoType
}

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

func NewDpskInfoType() *DpskInfoType {
	dpskInfoTypeType := new(DpskInfoType)
	return dpskInfoTypeType
}

func NewDefaultDpskInfoType() *DpskInfoType {
	dpskInfoTypeType := new(DpskInfoType)
	return dpskInfoTypeType
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

func NewDpskQueryList() *DpskQueryList {
	dpskQueryListType := new(DpskQueryList)
	return dpskQueryListType
}

func NewDefaultDpskQueryList() *DpskQueryList {
	dpskQueryListType := new(DpskQueryList)
	return dpskQueryListType
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

func NewDpskQueryListType() *DpskQueryListType {
	dpskQueryListTypeType := new(DpskQueryListType)
	return dpskQueryListTypeType
}

func NewDefaultDpskQueryListType() *DpskQueryListType {
	dpskQueryListTypeType := new(DpskQueryListType)
	return dpskQueryListTypeType
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

func NewGetDpskEnabledWlans() *GetDpskEnabledWlans {
	getDpskEnabledWlansType := new(GetDpskEnabledWlans)
	return getDpskEnabledWlansType
}

func NewDefaultGetDpskEnabledWlans() *GetDpskEnabledWlans {
	getDpskEnabledWlansType := new(GetDpskEnabledWlans)
	return getDpskEnabledWlansType
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

func NewGetDpskEnabledWlansListType() *GetDpskEnabledWlansListType {
	getDpskEnabledWlansListTypeType := new(GetDpskEnabledWlansListType)
	return getDpskEnabledWlansListTypeType
}

func NewDefaultGetDpskEnabledWlansListType() *GetDpskEnabledWlansListType {
	getDpskEnabledWlansListTypeType := new(GetDpskEnabledWlansListType)
	return getDpskEnabledWlansListTypeType
}

type GetDpskInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List DpskInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewGetDpskInfoList() *GetDpskInfoList {
	getDpskInfoListType := new(GetDpskInfoList)
	return getDpskInfoListType
}

func NewDefaultGetDpskInfoList() *GetDpskInfoList {
	getDpskInfoListType := new(GetDpskInfoList)
	return getDpskInfoListType
}

type GetDpskResult struct {
	DpskInfoList DpskInfo `json:"dpskInfoList,omitempty"`

	ResultCount *int `json:"resultCount,omitempty"`
}

func NewGetDpskResult() *GetDpskResult {
	getDpskResultType := new(GetDpskResult)
	return getDpskResultType
}

func NewDefaultGetDpskResult() *GetDpskResult {
	getDpskResultType := new(GetDpskResult)
	return getDpskResultType
}

type ModifyDeleteExpiredDpsk struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	// Constraints:
	//    - nullable
	//    - oneof:[Never,AfterOneDay,AfterSixMonths]
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty" validate:"omitempty,oneof=Never AfterOneDay AfterSixMonths"`
}

func NewModifyDeleteExpiredDpsk() *ModifyDeleteExpiredDpsk {
	modifyDeleteExpiredDpskType := new(ModifyDeleteExpiredDpsk)
	return modifyDeleteExpiredDpskType
}

func NewDefaultModifyDeleteExpiredDpsk() *ModifyDeleteExpiredDpsk {
	modifyDeleteExpiredDpskType := new(ModifyDeleteExpiredDpsk)
	return modifyDeleteExpiredDpskType
}

type UpdateDpsk struct {
	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`
}

func NewUpdateDpsk() *UpdateDpsk {
	updateDpskType := new(UpdateDpsk)
	return updateDpskType
}

func NewDefaultUpdateDpsk() *UpdateDpsk {
	updateDpskType := new(UpdateDpsk)
	return updateDpskType
}

type WlanDpskSetting struct {
	// DpskEnabled
	// DPSK enabled
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
	Length *int `json:"length,omitempty"`
}

func NewWlanDpskSetting() *WlanDpskSetting {
	wlanDpskSettingType := new(WlanDpskSetting)
	return wlanDpskSettingType
}

func NewDefaultWlanDpskSetting() *WlanDpskSetting {
	wlanDpskSettingType := new(WlanDpskSetting)
	dpskEnabledField := false
	wlanDpskSettingType.DpskEnabled = &dpskEnabledField
	dpskFromTypeField := `CreateTime`
	wlanDpskSettingType.DpskFromType = &dpskFromTypeField
	dpskTypeField := `Secure`
	wlanDpskSettingType.DpskType = &dpskTypeField
	expirationField := `Unlimited`
	wlanDpskSettingType.Expiration = &expirationField
	lengthField := 62
	wlanDpskSettingType.Length = &lengthField
	return wlanDpskSettingType
}

type WlanExternalDpsk struct {
	// AuthService
	// Autentication of the WLAN relate to external DPSK
	AuthService *WlanExternalDpskAuthServiceType `json:"authService,omitempty"`

	// Enabled
	// Enable External DPSK
	// Constraints:
	//    - required
	//    - default:false
	Enabled *bool `json:"enabled" validate:"required"`

	// Encryption
	// Encryption of the WLAN relate to external DPSK
	Encryption *WlanExternalDpskEncryptionType `json:"encryption,omitempty"`
}

func NewWlanExternalDpsk() *WlanExternalDpsk {
	wlanExternalDpskType := new(WlanExternalDpsk)
	return wlanExternalDpskType
}

func NewDefaultWlanExternalDpsk() *WlanExternalDpsk {
	wlanExternalDpskType := new(WlanExternalDpsk)
	enabledField := false
	wlanExternalDpskType.Enabled = &enabledField
	return wlanExternalDpskType
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

func NewWlanExternalDpskAuthServiceType() *WlanExternalDpskAuthServiceType {
	wlanExternalDpskAuthServiceTypeType := new(WlanExternalDpskAuthServiceType)
	return wlanExternalDpskAuthServiceTypeType
}

func NewDefaultWlanExternalDpskAuthServiceType() *WlanExternalDpskAuthServiceType {
	wlanExternalDpskAuthServiceTypeType := new(WlanExternalDpskAuthServiceType)
	return wlanExternalDpskAuthServiceTypeType
}

// WlanExternalDpskEncryptionType
//
// Encryption of the WLAN relate to external DPSK
type WlanExternalDpskEncryptionType struct {
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
}

func NewWlanExternalDpskEncryptionType() *WlanExternalDpskEncryptionType {
	wlanExternalDpskEncryptionTypeType := new(WlanExternalDpskEncryptionType)
	return wlanExternalDpskEncryptionTypeType
}

func NewDefaultWlanExternalDpskEncryptionType() *WlanExternalDpskEncryptionType {
	wlanExternalDpskEncryptionTypeType := new(WlanExternalDpskEncryptionType)
	return wlanExternalDpskEncryptionTypeType
}
