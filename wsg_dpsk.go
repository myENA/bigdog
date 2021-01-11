package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGDPSKBatchGenUnbound
//
// Definition: dpsk_batchGenUnbound
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

// WSGDPSKDeleteDpskResult
//
// Definition: dpsk_deleteDpskResult
type WSGDPSKDeleteDpskResult struct {
	ResultCount *int `json:"resultCount,omitempty"`
}

type WSGDPSKDeleteDpskResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPSKDeleteDpskResult
}

func newWSGDPSKDeleteDpskResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPSKDeleteDpskResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPSKDeleteDpskResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPSKDeleteDpskResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPSKDeleteDpskResult() *WSGDPSKDeleteDpskResult {
	m := new(WSGDPSKDeleteDpskResult)
	return m
}

// WSGDPSKDeleteDPSKs
//
// Definition: dpsk_deleteDPSKs
type WSGDPSKDeleteDPSKs struct {
	IdList []string `json:"idList,omitempty"`
}

func NewWSGDPSKDeleteDPSKs() *WSGDPSKDeleteDPSKs {
	m := new(WSGDPSKDeleteDPSKs)
	return m
}

// WSGDPSKDeleteExpiredDpskConfig
//
// Definition: dpsk_deleteExpiredDpskConfig
type WSGDPSKDeleteExpiredDpskConfig struct {
	// DeleteExpiredDpsk
	// Delete expired DPSK interval of the Zone.
	// Constraints:
	//    - oneof:[Never,AfterOneDay,AfterSixMonths]
	DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
}

type WSGDPSKDeleteExpiredDpskConfigAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPSKDeleteExpiredDpskConfig
}

func newWSGDPSKDeleteExpiredDpskConfigAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPSKDeleteExpiredDpskConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPSKDeleteExpiredDpskConfigAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPSKDeleteExpiredDpskConfig)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPSKDeleteExpiredDpskConfig() *WSGDPSKDeleteExpiredDpskConfig {
	m := new(WSGDPSKDeleteExpiredDpskConfig)
	return m
}

// WSGDPSKInfo
//
// Definition: dpsk_dpskInfo
//
// Information list of DPSK
type WSGDPSKInfo []*WSGDPSKInfoType

func MakeWSGDPSKInfo() WSGDPSKInfo {
	m := make(WSGDPSKInfo, 0)
	return m
}

// WSGDPSKInfoType
//
// Definition: dpsk_dpskInfoType
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

// WSGDPSKQueryList
//
// Definition: dpsk_dpskQueryList
type WSGDPSKQueryList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of DPSK
	List []*WSGDPSKQueryListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPSKQueryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPSKQueryList
}

func newWSGDPSKQueryListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPSKQueryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPSKQueryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPSKQueryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPSKQueryList() *WSGDPSKQueryList {
	m := new(WSGDPSKQueryList)
	return m
}

// WSGDPSKQueryListType
//
// Definition: dpsk_dpskQueryListType
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

// WSGDPSKGetDpskEnabledWlans
//
// Definition: dpsk_getDpskEnabledWlans
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

type WSGDPSKGetDpskEnabledWlansAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPSKGetDpskEnabledWlans
}

func newWSGDPSKGetDpskEnabledWlansAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPSKGetDpskEnabledWlansAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPSKGetDpskEnabledWlansAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPSKGetDpskEnabledWlans)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPSKGetDpskEnabledWlans() *WSGDPSKGetDpskEnabledWlans {
	m := new(WSGDPSKGetDpskEnabledWlans)
	return m
}

// WSGDPSKGetDpskEnabledWlansListType
//
// Definition: dpsk_getDpskEnabledWlansListType
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

// WSGDPSKGetDpskInfoList
//
// Definition: dpsk_getDpskInfoList
type WSGDPSKGetDpskInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List WSGDPSKInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPSKGetDpskInfoListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPSKGetDpskInfoList
}

func newWSGDPSKGetDpskInfoListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPSKGetDpskInfoListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPSKGetDpskInfoListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPSKGetDpskInfoList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPSKGetDpskInfoList() *WSGDPSKGetDpskInfoList {
	m := new(WSGDPSKGetDpskInfoList)
	return m
}

// WSGDPSKGetDpskResult
//
// Definition: dpsk_getDpskResult
type WSGDPSKGetDpskResult struct {
	DpskInfoList WSGDPSKInfo `json:"dpskInfoList,omitempty"`

	ResultCount *int `json:"resultCount,omitempty"`
}

type WSGDPSKGetDpskResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPSKGetDpskResult
}

func newWSGDPSKGetDpskResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPSKGetDpskResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPSKGetDpskResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPSKGetDpskResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPSKGetDpskResult() *WSGDPSKGetDpskResult {
	m := new(WSGDPSKGetDpskResult)
	return m
}

// WSGDPSKModifyDeleteExpiredDpsk
//
// Definition: dpsk_modifyDeleteExpiredDpsk
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

// WSGDPSKUpdateDpsk
//
// Definition: dpsk_updateDpsk
type WSGDPSKUpdateDpsk struct {
	// UserName
	// User Name of DPSK
	UserName *string `json:"userName,omitempty"`
}

func NewWSGDPSKUpdateDpsk() *WSGDPSKUpdateDpsk {
	m := new(WSGDPSKUpdateDpsk)
	return m
}

// WSGDPSKWlanDpskSetting
//
// Definition: dpsk_wlanDpskSetting
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

// WSGDPSKWlanExternalDpsk
//
// Definition: dpsk_wlanExternalDpsk
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
// Definition: dpsk_wlanExternalDpskAuthServiceType
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
// Definition: dpsk_wlanExternalDpskEncryptionType
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
