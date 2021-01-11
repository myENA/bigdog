package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGProfileAccountingProfile
//
// Definition: profile_accountingProfile
type WSGProfileAccountingProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the accounting profile
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*WSGProfileAcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

type WSGProfileAccountingProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileAccountingProfile
}

func newWSGProfileAccountingProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileAccountingProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileAccountingProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileAccountingProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileAccountingProfile() *WSGProfileAccountingProfile {
	m := new(WSGProfileAccountingProfile)
	return m
}

// WSGProfileAccountingProfileList
//
// Definition: profile_accountingProfileList
type WSGProfileAccountingProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileAccountingProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileAccountingProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileAccountingProfileList
}

func newWSGProfileAccountingProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileAccountingProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileAccountingProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileAccountingProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileAccountingProfileList() *WSGProfileAccountingProfileList {
	m := new(WSGProfileAccountingProfileList)
	return m
}

// WSGProfileAcctServiceRealmMapping
//
// Definition: profile_acctServiceRealmMapping
//
// Accounting service per realm
type WSGProfileAcctServiceRealmMapping struct {
	// Id
	// Accounting service UUID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Accounting service name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// ServiceType
	// Accounting service type, NA is NA-Request Rejected
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,CGF]
	ServiceType *string `json:"serviceType"`
}

func NewWSGProfileAcctServiceRealmMapping() *WSGProfileAcctServiceRealmMapping {
	m := new(WSGProfileAcctServiceRealmMapping)
	return m
}

// WSGProfileAdvancedOptionContent
//
// Definition: profile_advancedOptionContent
//
// advanced option Content
type WSGProfileAdvancedOptionContent struct {
	// DhcpOpt43Subcode
	// dhcpOpt43Subcode of the ipsec profile
	DhcpOpt43Subcode *float64 `json:"dhcpOpt43Subcode,omitempty"`

	// DpdDelay
	// dpdDelay of the ipsec profile
	DpdDelay *float64 `json:"dpdDelay,omitempty"`

	// EnforceNatt
	// enforceNatt Enable of the ipsec profile
	// Constraints:
	//    - oneof:[Disabled,Enabled]
	EnforceNatt *string `json:"enforceNatt,omitempty"`

	// FailoverMode
	// mode of the failover
	// Constraints:
	//    - oneof:[Non_Revertive,Revertive]
	FailoverMode *string `json:"failoverMode,omitempty"`

	// FailoverPrimaryCheckInterval
	// Primary Check Interval of the failover
	FailoverPrimaryCheckInterval *float64 `json:"failoverPrimaryCheckInterval,omitempty"`

	// FailoverRetryInterval
	// Retry Interval of the failover
	FailoverRetryInterval *float64 `json:"failoverRetryInterval,omitempty"`

	// FailoverRetryPeriod
	// Retry Period of the failover
	FailoverRetryPeriod *float64 `json:"failoverRetryPeriod,omitempty"`

	// IpcompEnable
	// ipcomp Enable of the ipsec profile
	// Constraints:
	//    - oneof:[Disabled,Enabled]
	IpcompEnable *string `json:"ipcompEnable,omitempty"`

	// KeepAliveIntval
	// keepAliveIntval of the ipsec profile
	KeepAliveIntval *float64 `json:"keepAliveIntval,omitempty"`

	// ReplayWindow
	// replayWindow of the ipsec profile
	ReplayWindow *float64 `json:"replayWindow,omitempty"`

	// RetryLimit
	// retryLimit of the ipsec profile
	RetryLimit *float64 `json:"retryLimit,omitempty"`
}

func NewWSGProfileAdvancedOptionContent() *WSGProfileAdvancedOptionContent {
	m := new(WSGProfileAdvancedOptionContent)
	return m
}

// WSGProfileApnRealm
//
// Definition: profile_apnRealm
type WSGProfileApnRealm struct {
	// DefaultAPN
	// name of the apnForwardingPolicys.
	DefaultAPN *string `json:"defaultAPN,omitempty"`

	// Realm
	// name of the apnRealm.
	Realm *string `json:"realm,omitempty"`
}

func NewWSGProfileApnRealm() *WSGProfileApnRealm {
	m := new(WSGProfileApnRealm)
	return m
}

// WSGProfileAuthenticationProfile
//
// Definition: profile_authenticationProfile
type WSGProfileAuthenticationProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	GppSuppportEnabled *bool `json:"gppSuppportEnabled,omitempty"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// Id
	// Identifier of the authentication profile
	Id *string `json:"id,omitempty"`

	// IsContainDirectoryService
	// Realm based authentication service mappings contains LDAP or AD service type
	IsContainDirectoryService *bool `json:"isContainDirectoryService,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*WSGProfileRealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *WSGProfileTtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

type WSGProfileAuthenticationProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileAuthenticationProfile
}

func newWSGProfileAuthenticationProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileAuthenticationProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileAuthenticationProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileAuthenticationProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileAuthenticationProfile() *WSGProfileAuthenticationProfile {
	m := new(WSGProfileAuthenticationProfile)
	return m
}

// WSGProfileAuthenticationProfileList
//
// Definition: profile_authenticationProfileList
type WSGProfileAuthenticationProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileAuthenticationProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileAuthenticationProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileAuthenticationProfileList
}

func newWSGProfileAuthenticationProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileAuthenticationProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileAuthenticationProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileAuthenticationProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileAuthenticationProfileList() *WSGProfileAuthenticationProfileList {
	m := new(WSGProfileAuthenticationProfileList)
	return m
}

// WSGProfileBaseServiceInfoList
//
// Definition: profile_baseServiceInfoList
type WSGProfileBaseServiceInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCommonBaseServiceInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileBaseServiceInfoListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBaseServiceInfoList
}

func newWSGProfileBaseServiceInfoListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBaseServiceInfoListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBaseServiceInfoListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBaseServiceInfoList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBaseServiceInfoList() *WSGProfileBaseServiceInfoList {
	m := new(WSGProfileBaseServiceInfoList)
	return m
}

// WSGProfileBlockClient
//
// Definition: profile_blockClient
type WSGProfileBlockClient struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// ZoneId
	// Zone Id of the Block Client for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGProfileBlockClientAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBlockClient
}

func newWSGProfileBlockClientAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBlockClientAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBlockClientAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBlockClient)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBlockClient() *WSGProfileBlockClient {
	m := new(WSGProfileBlockClient)
	return m
}

// WSGProfileBlockClientList
//
// Definition: profile_blockClientList
type WSGProfileBlockClientList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileBlockClientListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileBlockClientListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBlockClientList
}

func newWSGProfileBlockClientListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBlockClientListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBlockClientListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBlockClientList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBlockClientList() *WSGProfileBlockClientList {
	m := new(WSGProfileBlockClientList)
	return m
}

// WSGProfileBlockClientListType
//
// Definition: profile_blockClientListType
type WSGProfileBlockClientListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

	// ModifiedDateTime
	// Date blocked of the Block Client
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Modifier blocked of the Block Client
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// ZoneId
	// Zone Id of the Block Client for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBlockClientListType() *WSGProfileBlockClientListType {
	m := new(WSGProfileBlockClientListType)
	return m
}

// WSGProfileBlockedPort
//
// Definition: profile_blockedPort
type WSGProfileBlockedPort struct {
	// Port
	// port of the Blocked Port
	Port *string `json:"port,omitempty"`

	// Protocol
	// Protocol of the Blocked Port
	// Constraints:
	//    - oneof:[BOTH,TCP,UDP]
	Protocol *string `json:"protocol,omitempty"`
}

func NewWSGProfileBlockedPort() *WSGProfileBlockedPort {
	m := new(WSGProfileBlockedPort)
	return m
}

// WSGProfileBonjourFencingPolicy
//
// Definition: profile_bonjourFencingPolicy
type WSGProfileBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	// Constraints:
	//    - required
	BonjourFencingRuleList []*WSGProfileBonjourFencingRule `json:"bonjourFencingRuleList"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*WSGProfileBonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Bonjour Fencing Policy id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGProfileBonjourFencingPolicyAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBonjourFencingPolicy
}

func newWSGProfileBonjourFencingPolicyAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBonjourFencingPolicyAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBonjourFencingPolicyAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBonjourFencingPolicy)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBonjourFencingPolicy() *WSGProfileBonjourFencingPolicy {
	m := new(WSGProfileBonjourFencingPolicy)
	return m
}

// WSGProfileBonjourFencingPolicyList
//
// Definition: profile_bonjourFencingPolicyList
type WSGProfileBonjourFencingPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileBonjourFencingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileBonjourFencingPolicyListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBonjourFencingPolicyList
}

func newWSGProfileBonjourFencingPolicyListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBonjourFencingPolicyListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBonjourFencingPolicyListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBonjourFencingPolicyList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBonjourFencingPolicyList() *WSGProfileBonjourFencingPolicyList {
	m := new(WSGProfileBonjourFencingPolicyList)
	return m
}

// WSGProfileBonjourFencingRule
//
// Definition: profile_bonjourFencingRule
type WSGProfileBonjourFencingRule struct {
	ClosestAp *WSGCommonMac `json:"closestAp,omitempty"`

	CustomServiceName *WSGCommonNormalName `json:"customServiceName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DeviceMacList
	// Specify the device list providing Bonjour Service
	DeviceMacList []*WSGProfileBonjourFencingRuleDeviceMac `json:"deviceMacList,omitempty"`

	// DeviceType
	// Name of the Bonjour Fencing Rule
	// Constraints:
	//    - required
	//    - oneof:[Wired,Wireless]
	DeviceType *string `json:"deviceType"`

	// FencingRange
	// The range of AP can take Bonjour work
	// Constraints:
	//    - required
	//    - oneof:[SameAp,OneHopAp]
	FencingRange *string `json:"fencingRange"`

	// ServiceType
	// Constraints:
	//    - required
	ServiceType *WSGProfileBridgeService `json:"serviceType"`
}

func NewWSGProfileBonjourFencingRule() *WSGProfileBonjourFencingRule {
	m := new(WSGProfileBonjourFencingRule)
	return m
}

// WSGProfileBonjourFencingRuleDeviceMac
//
// Definition: profile_bonjourFencingRuleDeviceMac
type WSGProfileBonjourFencingRuleDeviceMac struct {
	Mac *WSGCommonMac `json:"mac,omitempty"`
}

func NewWSGProfileBonjourFencingRuleDeviceMac() *WSGProfileBonjourFencingRuleDeviceMac {
	m := new(WSGProfileBonjourFencingRuleDeviceMac)
	return m
}

// WSGProfileBonjourFencingRuleMapping
//
// Definition: profile_bonjourFencingRuleMapping
type WSGProfileBonjourFencingRuleMapping struct {
	CustomServiceName *WSGCommonNormalName `json:"customServiceName,omitempty"`

	// CustomStringList
	// The array of mdns string
	CustomStringList []string `json:"customStringList,omitempty"`

	ServiceType *WSGProfileBridgeService `json:"serviceType,omitempty"`
}

func NewWSGProfileBonjourFencingRuleMapping() *WSGProfileBonjourFencingRuleMapping {
	m := new(WSGProfileBonjourFencingRuleMapping)
	return m
}

// WSGProfileBonjourFencingService
//
// Definition: profile_bonjourFencingService
type WSGProfileBonjourFencingService struct {
	NeighborApMac *string `json:"neighborApMac,omitempty"`

	NeighborApName *string `json:"neighborApName,omitempty"`

	ServiceType *WSGProfileBridgeService `json:"serviceType,omitempty"`

	// SourceType
	// Constraints:
	//    - oneof:[UNKNOWN,DIRECT,NEIGHBOR]
	SourceType *string `json:"sourceType,omitempty"`
}

func NewWSGProfileBonjourFencingService() *WSGProfileBonjourFencingService {
	m := new(WSGProfileBonjourFencingService)
	return m
}

// WSGProfileBonjourFencingStatistic
//
// Definition: profile_bonjourFencingStatistic
type WSGProfileBonjourFencingStatistic struct {
	ApMac *string `json:"apMac,omitempty"`

	DroppedPacketsDueToNeighbor *int `json:"droppedPacketsDueToNeighbor,omitempty"`

	DroppedPacketsDueToServiceType *int `json:"droppedPacketsDueToServiceType,omitempty"`

	ForwardedPackets *int `json:"forwardedPackets,omitempty"`

	ServiceList []*WSGProfileBonjourFencingService `json:"serviceList,omitempty"`
}

type WSGProfileBonjourFencingStatisticAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBonjourFencingStatistic
}

func newWSGProfileBonjourFencingStatisticAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBonjourFencingStatisticAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBonjourFencingStatisticAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBonjourFencingStatistic)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBonjourFencingStatistic() *WSGProfileBonjourFencingStatistic {
	m := new(WSGProfileBonjourFencingStatistic)
	return m
}

// WSGProfileBridgeProfile
//
// Definition: profile_bridgeProfile
type WSGProfileBridgeProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

type WSGProfileBridgeProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBridgeProfile
}

func newWSGProfileBridgeProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBridgeProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBridgeProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBridgeProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBridgeProfile() *WSGProfileBridgeProfile {
	m := new(WSGProfileBridgeProfile)
	return m
}

// WSGProfileBridgeProfileList
//
// Definition: profile_bridgeProfileList
type WSGProfileBridgeProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileBridgeProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileBridgeProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileBridgeProfileList
}

func newWSGProfileBridgeProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileBridgeProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileBridgeProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileBridgeProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileBridgeProfileList() *WSGProfileBridgeProfileList {
	m := new(WSGProfileBridgeProfileList)
	return m
}

// WSGProfileBridgeService
//
// Definition: profile_bridgeService
//
// Bonjour Service Type
// Constraints:
//    - oneof:[AIRDISK,AIRPLAY,AIRPORT_MANAGEMENT,AIRPRINT,AIRTUNES,APPLE_FILE_SHARING,APPLE_MOBILE_DEVICES,APPLETV,ICLOUD_SYNC,ITUNES_REMOTE,ITUNES_SHARING,OPEN_DIRECTORY_MASTER,OPTICAL_DISK_SHARING,SCREEN_SHARING,SECURE_FILE_SHARING,SECURE_SHELL,WWW_HTTP,WWW_HTTPS,WORKGROUP_MANAGER,XGRID,GOOGLE_CHROMECAST,OTHER]
type WSGProfileBridgeService string

func NewWSGProfileBridgeService() *WSGProfileBridgeService {
	m := new(WSGProfileBridgeService)
	return m
}

// WSGProfileBulkBlockClient
//
// Definition: profile_bulkBlockClient
type WSGProfileBulkBlockClient struct {
	BlockClientList []*WSGProfileBulkBlockClientBlockClientListType `json:"blockClientList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`
}

func NewWSGProfileBulkBlockClient() *WSGProfileBulkBlockClient {
	m := new(WSGProfileBulkBlockClient)
	return m
}

// WSGProfileBulkBlockClientBlockClientListType
//
// Definition: profile_bulkBlockClientBlockClientListType
type WSGProfileBulkBlockClientBlockClientListType struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *WSGCommonMac `json:"apMac"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBulkBlockClientBlockClientListType() *WSGProfileBulkBlockClientBlockClientListType {
	m := new(WSGProfileBulkBlockClientBlockClientListType)
	return m
}

// WSGProfileClientIsolationEntry
//
// Definition: profile_clientIsolationEntry
type WSGProfileClientIsolationEntry struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Ip
	// Client Entry ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac"`
}

func NewWSGProfileClientIsolationEntry() *WSGProfileClientIsolationEntry {
	m := new(WSGProfileClientIsolationEntry)
	return m
}

// WSGProfileClientIsolationWhitelist
//
// Definition: profile_clientIsolationWhitelist
type WSGProfileClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	// Constraints:
	//    - required
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Client Isolation Whitelist id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - required
	Whitelist []*WSGProfileClientIsolationEntry `json:"whitelist"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGProfileClientIsolationWhitelistAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileClientIsolationWhitelist
}

func newWSGProfileClientIsolationWhitelistAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileClientIsolationWhitelistAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileClientIsolationWhitelistAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileClientIsolationWhitelist)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileClientIsolationWhitelist() *WSGProfileClientIsolationWhitelist {
	m := new(WSGProfileClientIsolationWhitelist)
	return m
}

// WSGProfileClientIsolationWhitelistArray
//
// Definition: profile_clientIsolationWhitelistArray
type WSGProfileClientIsolationWhitelistArray struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileClientIsolationWhitelist `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileClientIsolationWhitelistArrayAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileClientIsolationWhitelistArray
}

func newWSGProfileClientIsolationWhitelistArrayAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileClientIsolationWhitelistArrayAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileClientIsolationWhitelistArrayAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileClientIsolationWhitelistArray)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileClientIsolationWhitelistArray() *WSGProfileClientIsolationWhitelistArray {
	m := new(WSGProfileClientIsolationWhitelistArray)
	return m
}

// WSGProfileCmProtocolOptionContent
//
// Definition: profile_cmProtocolOptionContent
//
// Certificate Management Protocol Option
type WSGProfileCmProtocolOptionContent struct {
	// CmpDhcpOpt43Subcode
	// Certificate Management Protocol dhcpOpt43Subcode
	// Constraints:
	//    - required
	CmpDhcpOpt43Subcode *float64 `json:"cmpDhcpOpt43Subcode"`

	// CmpDhcpOpt43SubcodeRecipient
	// Certificate Management Protocol dhcpOpt43SubcodeRecipient
	CmpDhcpOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`

	// CmpRecipient
	// Certificate Management Protocol Recipient
	CmpRecipient *string `json:"cmpRecipient,omitempty"`

	// CmpServerAddr
	// Certificate Management Protocol Server addr
	CmpServerAddr *string `json:"cmpServerAddr,omitempty"`

	// CmpServerPath
	// Certificate Management Protocol Server Path
	CmpServerPath *string `json:"cmpServerPath,omitempty"`
}

func NewWSGProfileCmProtocolOptionContent() *WSGProfileCmProtocolOptionContent {
	m := new(WSGProfileCmProtocolOptionContent)
	return m
}

// WSGProfileCoreNetworkGateway
//
// Definition: profile_coreNetworkGateway
type WSGProfileCoreNetworkGateway struct {
	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:32767
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - default:3
	//    - min:1
	//    - max:255
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	// PrimaryGateway
	// Primary Gateway
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary Gateway
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMTU
	// Gateway path MTU
	// Constraints:
	//    - oneof:[AUTO,MANUAL]
	TunnelMTU *string `json:"tunnelMTU,omitempty"`

	// TunnelMTUSize
	// Manual setting value of Gateway path MTU
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:1500
	TunnelMTUSize *int `json:"tunnelMTUSize,omitempty"`
}

func NewWSGProfileCoreNetworkGateway() *WSGProfileCoreNetworkGateway {
	m := new(WSGProfileCoreNetworkGateway)
	return m
}

// WSGProfileCreateAccountingProfile
//
// Definition: profile_createAccountingProfile
type WSGProfileCreateAccountingProfile struct {
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*WSGProfileAcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

func NewWSGProfileCreateAccountingProfile() *WSGProfileCreateAccountingProfile {
	m := new(WSGProfileCreateAccountingProfile)
	return m
}

// WSGProfileCreateAuthenticationProfile
//
// Definition: profile_createAuthenticationProfile
type WSGProfileCreateAuthenticationProfile struct {
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	// Constraints:
	//    - required
	//    - default:false
	GppSuppportEnabled *bool `json:"gppSuppportEnabled"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*WSGProfileRealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *WSGProfileTtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewWSGProfileCreateAuthenticationProfile() *WSGProfileCreateAuthenticationProfile {
	m := new(WSGProfileCreateAuthenticationProfile)
	return m
}

// WSGProfileCreateBonjourFencingPolicy
//
// Definition: profile_createBonjourFencingPolicy
type WSGProfileCreateBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	// Constraints:
	//    - required
	BonjourFencingRuleList []*WSGProfileBonjourFencingRule `json:"bonjourFencingRuleList"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*WSGProfileBonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileCreateBonjourFencingPolicy() *WSGProfileCreateBonjourFencingPolicy {
	m := new(WSGProfileCreateBonjourFencingPolicy)
	return m
}

// WSGProfileCreateBridgeProfile
//
// Definition: profile_createBridgeProfile
type WSGProfileCreateBridgeProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileCreateBridgeProfile() *WSGProfileCreateBridgeProfile {
	m := new(WSGProfileCreateBridgeProfile)
	return m
}

// WSGProfileCreateClientIsolationWhitelist
//
// Definition: profile_createClientIsolationWhitelist
type WSGProfileCreateClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	// Constraints:
	//    - required
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - required
	Whitelist []*WSGProfileClientIsolationEntry `json:"whitelist"`
}

func NewWSGProfileCreateClientIsolationWhitelist() *WSGProfileCreateClientIsolationWhitelist {
	m := new(WSGProfileCreateClientIsolationWhitelist)
	return m
}

// WSGProfileCreateDhcpProfile
//
// Definition: profile_createDhcpProfile
type WSGProfileCreateDhcpProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// LeaseTimeHours
	// Lease time in hours of the DHCP Profile
	// Constraints:
	//    - required
	//    - min:0
	//    - max:24
	LeaseTimeHours *int `json:"leaseTimeHours"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	// Constraints:
	//    - required
	//    - min:0
	//    - max:59
	LeaseTimeMinutes *int `json:"leaseTimeMinutes"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PoolEndIp
	// Constraints:
	//    - required
	PoolEndIp *WSGCommonIpAddress `json:"poolEndIp"`

	// PoolStartIp
	// Constraints:
	//    - required
	PoolStartIp *WSGCommonIpAddress `json:"poolStartIp"`

	PrimaryDnsIp *WSGCommonIpAddress `json:"primaryDnsIp,omitempty"`

	SecondaryDnsIp *WSGCommonIpAddress `json:"secondaryDnsIp,omitempty"`

	// SubnetMask
	// Constraints:
	//    - required
	SubnetMask *WSGCommonIpAddress `json:"subnetMask"`

	// SubnetNetworkIp
	// Constraints:
	//    - required
	SubnetNetworkIp *WSGCommonIpAddress `json:"subnetNetworkIp"`

	// VlanId
	// VLAN ID of the DHCP Profile
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	VlanId *int `json:"vlanId"`
}

func NewWSGProfileCreateDhcpProfile() *WSGProfileCreateDhcpProfile {
	m := new(WSGProfileCreateDhcpProfile)
	return m
}

// WSGProfileCreateDnsServerProfile
//
// Definition: profile_createDnsServerProfile
type WSGProfileCreateDnsServerProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PrimaryIp
	// Primary ip of DNS server service
	// Constraints:
	//    - required
	PrimaryIp *string `json:"primaryIp"`

	// SecondaryIp
	// Secondary ip of DNS server service
	SecondaryIp *string `json:"secondaryIp,omitempty"`

	// TertiaryIp
	// Tertiary ip of DNS server service
	TertiaryIp *string `json:"tertiaryIp,omitempty"`
}

func NewWSGProfileCreateDnsServerProfile() *WSGProfileCreateDnsServerProfile {
	m := new(WSGProfileCreateDnsServerProfile)
	return m
}

// WSGProfileCreateFirewallProfile
//
// Definition: profile_createFirewallProfile
type WSGProfileCreateFirewallProfile struct {
	// AppPolicyId
	// Application Policy
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DevicePolicyId
	// Device Policy
	DevicePolicyId *string `json:"devicePolicyId,omitempty"`

	// DomainId
	// Domain Id of The Firewall Profile
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	// L3AccessControlPolicyId
	// L3 Access Control Policy
	L3AccessControlPolicyId *string `json:"l3AccessControlPolicyId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`

	// UrlFilteringPolicyId
	// Url Filtering Policy
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileCreateFirewallProfile() *WSGProfileCreateFirewallProfile {
	m := new(WSGProfileCreateFirewallProfile)
	return m
}

// WSGProfileCreateIpsecProfile
//
// Definition: profile_createIpsecProfile
type WSGProfileCreateIpsecProfile struct {
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty"`

	CmProtocolOption *WSGProfileCmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	// Constraints:
	//    - required
	EspRekeyTime *float64 `json:"espRekeyTime"`

	EspRekeyTimeUnit *WSGCommonTimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *WSGProfileEspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	// Constraints:
	//    - required
	IkeRekeyTime *float64 `json:"ikeRekeyTime"`

	IkeRekeyTimeUnit *WSGCommonTimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *WSGProfileIkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	// IpMode
	// Constraints:
	//    - required
	IpMode *WSGProfileIpMode `json:"ipMode"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`

	// TunnelMode
	// Tunnel mode of IPsec profile
	// Constraints:
	//    - oneof:[SOFT_GRE,RUCKUS_GRE]
	TunnelMode *string `json:"tunnelMode,omitempty"`
}

func NewWSGProfileCreateIpsecProfile() *WSGProfileCreateIpsecProfile {
	m := new(WSGProfileCreateIpsecProfile)
	return m
}

// WSGProfileCreateL2oGREProfile
//
// Definition: profile_createL2oGREProfile
type WSGProfileCreateL2oGREProfile struct {
	// CoreNetworkGateway
	// Constraints:
	//    - required
	CoreNetworkGateway *WSGProfileCoreNetworkGateway `json:"coreNetworkGateway"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileCreateL2oGREProfile() *WSGProfileCreateL2oGREProfile {
	m := new(WSGProfileCreateL2oGREProfile)
	return m
}

// WSGProfileCreateL3AccessControlPolicy
//
// Definition: profile_createL3AccessControlPolicy
type WSGProfileCreateL3AccessControlPolicy struct {
	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// L3AclRuleList
	// L3 access control list
	L3AclRuleList []*WSGProfileL3AclRuleList `json:"l3AclRuleList,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileCreateL3AccessControlPolicy() *WSGProfileCreateL3AccessControlPolicy {
	m := new(WSGProfileCreateL3AccessControlPolicy)
	return m
}

// WSGProfileCreatePrecedenceProfile
//
// Definition: profile_createPrecedenceProfile
type WSGProfileCreatePrecedenceProfile struct {
	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*WSGProfileRateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*WSGProfileVlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

type WSGProfileCreatePrecedenceProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileCreatePrecedenceProfile
}

func newWSGProfileCreatePrecedenceProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileCreatePrecedenceProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileCreatePrecedenceProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileCreatePrecedenceProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileCreatePrecedenceProfile() *WSGProfileCreatePrecedenceProfile {
	m := new(WSGProfileCreatePrecedenceProfile)
	return m
}

// WSGProfileCreateRestrictedApAccessProfile
//
// Definition: profile_createRestrictedApAccessProfile
type WSGProfileCreateRestrictedApAccessProfile struct {
	// BlockedPortList
	// Blocked Port List
	BlockedPortList []*WSGProfileBlockedPort `json:"blockedPortList,omitempty"`

	// BlockWellKnownPort
	// Block well known ports
	BlockWellKnownPort *bool `json:"blockWellKnownPort,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IpAddressWhitelist
	// IP Address Whitelist
	IpAddressWhitelist []*WSGCommonIpAddress `json:"ipAddressWhitelist,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileCreateRestrictedApAccessProfile() *WSGProfileCreateRestrictedApAccessProfile {
	m := new(WSGProfileCreateRestrictedApAccessProfile)
	return m
}

// WSGProfileCreateResultList
//
// Definition: profile_createResultList
type WSGProfileCreateResultList []*WSGCommonCreateResult

type WSGProfileCreateResultListAPIResponse struct {
	*RawAPIResponse
	Data WSGProfileCreateResultList
}

func newWSGProfileCreateResultListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileCreateResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileCreateResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(WSGProfileCreateResultList, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeWSGProfileCreateResultList() WSGProfileCreateResultList {
	m := make(WSGProfileCreateResultList, 0)
	return m
}

// WSGProfileCreateRogueApPolicy
//
// Definition: profile_createRogueApPolicy
type WSGProfileCreateRogueApPolicy struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Rules
	// Constraints:
	//    - required
	Rules []*WSGProfileRogueApRuleList `json:"rules"`
}

func NewWSGProfileCreateRogueApPolicy() *WSGProfileCreateRogueApPolicy {
	m := new(WSGProfileCreateRogueApPolicy)
	return m
}

// WSGProfileCreateRtlsProfile
//
// Definition: profile_createRtlsProfile
type WSGProfileCreateRtlsProfile struct {
	// EkahauEnabled
	// Eekahau Location Service Enabled
	// Constraints:
	//    - required
	EkahauEnabled *bool `json:"ekahauEnabled"`

	EkahauIp *WSGCommonIpAddress `json:"ekahauIp,omitempty"`

	// EkahauPort
	// Eekahau Location Server Port
	EkahauPort *int `json:"ekahauPort,omitempty"`

	Id *string `json:"id,omitempty"`

	// StanleyEnabled
	// Stanley Location Service Enabled
	// Constraints:
	//    - required
	StanleyEnabled *bool `json:"stanleyEnabled"`
}

type WSGProfileCreateRtlsProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileCreateRtlsProfile
}

func newWSGProfileCreateRtlsProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileCreateRtlsProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileCreateRtlsProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileCreateRtlsProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileCreateRtlsProfile() *WSGProfileCreateRtlsProfile {
	m := new(WSGProfileCreateRtlsProfile)
	return m
}

// WSGProfileCreateRuckusGREProfile
//
// Definition: profile_createRuckusGREProfile
type WSGProfileCreateRuckusGREProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// TunnelEncryption
	// Tunnel Encryption of the RuckusGRE profile
	// Constraints:
	//    - oneof:[DISABLE,AES128,AES256]
	TunnelEncryption *string `json:"tunnelEncryption,omitempty"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - required
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileCreateRuckusGREProfile() *WSGProfileCreateRuckusGREProfile {
	m := new(WSGProfileCreateRuckusGREProfile)
	return m
}

// WSGProfileCreateSoftGREProfile
//
// Definition: profile_createSoftGREProfile
type WSGProfileCreateSoftGREProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	IpMode *WSGProfileIpMode `json:"ipMode,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	// Constraints:
	//    - required
	//    - min:1
	//    - max:180
	KeepAlivePeriod *int `json:"keepAlivePeriod"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - required
	//    - min:2
	//    - max:20
	KeepAliveRetry *int `json:"keepAliveRetry"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	// Constraints:
	//    - required
	PrimaryGateway *string `json:"primaryGateway"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	// Constraints:
	//    - required
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile. IPV4:850-1500, IPV6:1384-1500. Default 1500.
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileCreateSoftGREProfile() *WSGProfileCreateSoftGREProfile {
	m := new(WSGProfileCreateSoftGREProfile)
	return m
}

// WSGProfileCreateTrafficClassProfile
//
// Definition: profile_createTrafficClassProfile
type WSGProfileCreateTrafficClassProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName2to64 `json:"name"`

	// TrafficClasses
	// Constraints:
	//    - required
	TrafficClasses []*WSGCommonTrafficClassRef `json:"trafficClasses"`
}

func NewWSGProfileCreateTrafficClassProfile() *WSGProfileCreateTrafficClassProfile {
	m := new(WSGProfileCreateTrafficClassProfile)
	return m
}

// WSGProfileCreateTtgpdgProfile
//
// Definition: profile_createTtgpdgProfile
type WSGProfileCreateTtgpdgProfile struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	// Constraints:
	//    - required
	ApnForwardingRealms []*WSGProfileTtgpdgApnForwardingRealm `json:"apnForwardingRealms"`

	// ApnRealms
	// List of the Default APN
	ApnRealms []*WSGProfileApnRealm `json:"apnRealms,omitempty"`

	// CommonSetting
	// Constraints:
	//    - required
	CommonSetting *WSGProfileTtgpdgCommonSetting `json:"commonSetting"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	// Constraints:
	//    - required
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	// Constraints:
	//    - required
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileCreateTtgpdgProfile() *WSGProfileCreateTtgpdgProfile {
	m := new(WSGProfileCreateTtgpdgProfile)
	return m
}

// WSGProfileCreateUserTrafficProfile
//
// Definition: profile_createUserTrafficProfile
type WSGProfileCreateUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *WSGProfileDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// IpAclRules
	// Traffic access control list
	IpAclRules []*WSGProfileIpAclRules `json:"ipAclRules,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// QmAppPolicyId
	// Application Policy UUID
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *WSGProfileUplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileCreateUserTrafficProfile() *WSGProfileCreateUserTrafficProfile {
	m := new(WSGProfileCreateUserTrafficProfile)
	return m
}

// WSGProfileCreateZoneAffinityProfile
//
// Definition: profile_createZoneAffinityProfile
type WSGProfileCreateZoneAffinityProfile struct {
	// Description
	// The description of the profile
	// Constraints:
	//    - max:64
	Description *string `json:"description,omitempty"`

	// Name
	// Zone affinity profile name
	// Constraints:
	//    - required
	//    - max:64
	//    - min:1
	Name *string `json:"name"`

	// ZoneAffinityList
	// Constraints:
	//    - required
	ZoneAffinityList []string `json:"zoneAffinityList"`
}

func NewWSGProfileCreateZoneAffinityProfile() *WSGProfileCreateZoneAffinityProfile {
	m := new(WSGProfileCreateZoneAffinityProfile)
	return m
}

// WSGProfileDataPlaneL3RoamingData
//
// Definition: profile_dataPlaneL3RoamingData
type WSGProfileDataPlaneL3RoamingData struct {
	// Activated
	// Show if this DP is included in the L3 roaming feature or not, 0 means excluded and 1 means included
	// Constraints:
	//    - required
	Activated *int `json:"activated"`

	// FirmwareVersion
	// DP firmware version
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Key
	// Data plane key
	// Constraints:
	//    - required
	Key *string `json:"key"`

	// Name
	// DP name
	Name *string `json:"name,omitempty"`

	// SubCriteriaType
	// Constraints:
	//    - oneof:[VLAN,SUBNET]
	SubCriteriaType *string `json:"subCriteriaType,omitempty"`

	// Value
	// A list of L3 roaming configuration for this DP
	// Constraints:
	//    - required
	Value *string `json:"value"`
}

func NewWSGProfileDataPlaneL3RoamingData() *WSGProfileDataPlaneL3RoamingData {
	m := new(WSGProfileDataPlaneL3RoamingData)
	return m
}

// WSGProfileDeleteBulkAccountingProfile
//
// Definition: profile_deleteBulkAccountingProfile
type WSGProfileDeleteBulkAccountingProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkAccountingProfile() *WSGProfileDeleteBulkAccountingProfile {
	m := new(WSGProfileDeleteBulkAccountingProfile)
	return m
}

// WSGProfileDeleteBulkAuthenticationProfile
//
// Definition: profile_deleteBulkAuthenticationProfile
type WSGProfileDeleteBulkAuthenticationProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkAuthenticationProfile() *WSGProfileDeleteBulkAuthenticationProfile {
	m := new(WSGProfileDeleteBulkAuthenticationProfile)
	return m
}

// WSGProfileDeleteBulkPrecedenceProfile
//
// Definition: profile_deleteBulkPrecedenceProfile
type WSGProfileDeleteBulkPrecedenceProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkPrecedenceProfile() *WSGProfileDeleteBulkPrecedenceProfile {
	m := new(WSGProfileDeleteBulkPrecedenceProfile)
	return m
}

// WSGProfileDeleteBulkUserTrafficProfile
//
// Definition: profile_deleteBulkUserTrafficProfile
type WSGProfileDeleteBulkUserTrafficProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkUserTrafficProfile() *WSGProfileDeleteBulkUserTrafficProfile {
	m := new(WSGProfileDeleteBulkUserTrafficProfile)
	return m
}

// WSGProfileDhcpOption82
//
// Definition: profile_dhcpOption82
type WSGProfileDhcpOption82 struct {
	// DhcpOption82Enabled
	// Enable DHCP Option 82
	DhcpOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// Subopt1Enabled
	// Enable subopt-1
	Subopt1Enabled *bool `json:"subopt1Enabled,omitempty"`

	// Subopt1Format
	// Subopt-1 format
	// Constraints:
	//    - oneof:[AP_INFO,AP_MAC_hex,AP_MAC_hex_ESSID,AP_INFO_LOCATION]
	Subopt1Format *string `json:"subopt1Format,omitempty"`

	// Subopt2Enabled
	// Enable subopt-2
	Subopt2Enabled *bool `json:"subopt2Enabled,omitempty"`

	// Subopt2Format
	// Subopt-2 format
	// Constraints:
	//    - oneof:[CLIENT_MAC_hex,CLIENT_MAC_hex_ESSID,AP_MAC_hex,AP_MAC__hex_ESSID]
	Subopt2Format *string `json:"subopt2Format,omitempty"`

	// Subopt150Enabled
	// Subopt-150 with VLAN
	Subopt150Enabled *bool `json:"subopt150Enabled,omitempty"`

	// Subopt151AreaName
	// Subopt-151 Area Name value
	Subopt151AreaName *string `json:"subopt151AreaName,omitempty"`

	// Subopt151Enabled
	// Enable subopt-151
	Subopt151Enabled *bool `json:"subopt151Enabled,omitempty"`

	// Subopt151Format
	// Subopt-151 format
	// Constraints:
	//    - oneof:[AREA_NAME,ESSID]
	Subopt151Format *string `json:"subopt151Format,omitempty"`
}

func NewWSGProfileDhcpOption82() *WSGProfileDhcpOption82 {
	m := new(WSGProfileDhcpOption82)
	return m
}

// WSGProfileDhcpProfileList
//
// Definition: profile_dhcpProfileList
type WSGProfileDhcpProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCommonDhcpProfileRef `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileDhcpProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileDhcpProfileList
}

func newWSGProfileDhcpProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileDhcpProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileDhcpProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileDhcpProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileDhcpProfileList() *WSGProfileDhcpProfileList {
	m := new(WSGProfileDhcpProfileList)
	return m
}

// WSGProfileDhcpRelayNoRelayTunnel
//
// Definition: profile_dhcpRelayNoRelayTunnel
type WSGProfileDhcpRelayNoRelayTunnel struct {
	DhcpOption82 *WSGProfileDhcpOption82 `json:"dhcpOption82,omitempty"`

	// DhcpRelayEnabled
	// Enable DHCP Relay
	DhcpRelayEnabled *bool `json:"dhcpRelayEnabled,omitempty"`

	// DhcpServer1
	// DHCP Server 1
	DhcpServer1 *string `json:"dhcpServer1,omitempty"`

	// DhcpServer2
	// DHCP Server 2
	DhcpServer2 *string `json:"dhcpServer2,omitempty"`

	// RelayBothEnabled
	// Send DHCP requests to both servers simultaneously.
	RelayBothEnabled *bool `json:"relayBothEnabled,omitempty"`
}

func NewWSGProfileDhcpRelayNoRelayTunnel() *WSGProfileDhcpRelayNoRelayTunnel {
	m := new(WSGProfileDhcpRelayNoRelayTunnel)
	return m
}

// WSGProfileDnsServerProfile
//
// Definition: profile_dnsServerProfile
type WSGProfileDnsServerProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryIp
	// Primary ip of DNS server service
	PrimaryIp *string `json:"primaryIp,omitempty"`

	// SecondaryIp
	// Secondary ip of DNS server service
	SecondaryIp *string `json:"secondaryIp,omitempty"`

	// TertiaryIp
	// Tertiary ip of DNS server service
	TertiaryIp *string `json:"tertiaryIp,omitempty"`
}

type WSGProfileDnsServerProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileDnsServerProfile
}

func newWSGProfileDnsServerProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileDnsServerProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileDnsServerProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileDnsServerProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileDnsServerProfile() *WSGProfileDnsServerProfile {
	m := new(WSGProfileDnsServerProfile)
	return m
}

// WSGProfileDnsServerProfileList
//
// Definition: profile_dnsServerProfileList
type WSGProfileDnsServerProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileDnsServerProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileDnsServerProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileDnsServerProfileList
}

func newWSGProfileDnsServerProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileDnsServerProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileDnsServerProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileDnsServerProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileDnsServerProfileList() *WSGProfileDnsServerProfileList {
	m := new(WSGProfileDnsServerProfileList)
	return m
}

// WSGProfileDownlinkRateLimiting
//
// Definition: profile_downlinkRateLimiting
type WSGProfileDownlinkRateLimiting struct {
	// DownlinkRateLimitingBps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	DownlinkRateLimitingBps *string `json:"downlinkRateLimitingBps,omitempty"`

	// DownlinkRateLimitingEnabled
	// Downlink rate limiting enabled or disabled
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`
}

func NewWSGProfileDownlinkRateLimiting() *WSGProfileDownlinkRateLimiting {
	m := new(WSGProfileDownlinkRateLimiting)
	return m
}

// WSGProfileEspProposal
//
// Definition: profile_espProposal
//
// EspProposal based ipsec service mappings
type WSGProfileEspProposal struct {
	// AuthAlg
	// authAlg of espProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[MD5,SHA1,AESXCBC,SHA256,SHA384,SHA512]
	AuthAlg *string `json:"authAlg"`

	// DhGroup
	// dhGroup of espProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[None,Modp768,Modp1024,Modp1536,Modp2048,Modp3072,Modp4096,Modp6144,Modp8192,Ecp384]
	DhGroup *string `json:"dhGroup"`

	// EncAlg
	// encAlg of espProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[None,ThreeDES,AES128,AES192,AES256]
	EncAlg *string `json:"encAlg"`
}

func NewWSGProfileEspProposal() *WSGProfileEspProposal {
	m := new(WSGProfileEspProposal)
	return m
}

// WSGProfileEspSecurityAssociationContent
//
// Definition: profile_espSecurityAssociationContent
//
// espProposal Security Association Content
type WSGProfileEspSecurityAssociationContent struct {
	// EspProposals
	// espProposal list of the ipsec profile
	EspProposals []*WSGProfileEspProposal `json:"espProposals,omitempty"`

	// EspProposalType
	// espProposal Type of the ipsec profile
	// Constraints:
	//    - oneof:[Default,Specific]
	EspProposalType *string `json:"espProposalType,omitempty"`
}

func NewWSGProfileEspSecurityAssociationContent() *WSGProfileEspSecurityAssociationContent {
	m := new(WSGProfileEspSecurityAssociationContent)
	return m
}

// WSGProfileFirewallProfile
//
// Definition: profile_firewallProfile
type WSGProfileFirewallProfile struct {
	// AppPolicyId
	// Application Policy
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DevicePolicyId
	// Device Policy
	DevicePolicyId *string `json:"devicePolicyId,omitempty"`

	// DomainId
	// Domain Id of The Firewall Profile
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// FactoryDefault
	// Whether the proFirewall Profile is factory default or not
	FactoryDefault *bool `json:"factoryDefault,omitempty"`

	// Id
	// Firewall Profile id
	Id *string `json:"id,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	// L3AccessControlPolicyId
	// L3 Access Control Policy
	L3AccessControlPolicyId *string `json:"l3AccessControlPolicyId,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`

	// UrlFilteringPolicyId
	// Url Filtering Policy
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

type WSGProfileFirewallProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileFirewallProfile
}

func newWSGProfileFirewallProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileFirewallProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileFirewallProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileFirewallProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileFirewallProfile() *WSGProfileFirewallProfile {
	m := new(WSGProfileFirewallProfile)
	return m
}

// WSGProfileFirewallProfileArray
//
// Definition: profile_firewallProfileArray
type WSGProfileFirewallProfileArray struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileFirewallProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileFirewallProfileArrayAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileFirewallProfileArray
}

func newWSGProfileFirewallProfileArrayAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileFirewallProfileArrayAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileFirewallProfileArrayAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileFirewallProfileArray)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileFirewallProfileArray() *WSGProfileFirewallProfileArray {
	m := new(WSGProfileFirewallProfileArray)
	return m
}

// WSGProfileFlexiVpnProfile
//
// Definition: profile_flexiVpnProfile
type WSGProfileFlexiVpnProfile struct {
	// DestinationZoneAffinityId
	// Flexi-VPN Profile ID (Destination)
	DestinationZoneAffinityId *string `json:"destinationZoneAffinityId,omitempty"`

	// DestinationZoneAffinityName
	// Flexi-VPN Profile (Destination)
	DestinationZoneAffinityName *string `json:"destinationZoneAffinityName,omitempty"`

	// DomainId
	// Domain ID
	DomainId *string `json:"domainId,omitempty"`

	// SourceZoneAffinityId
	// Zone Affinity Profile ID (Source)
	SourceZoneAffinityId *string `json:"sourceZoneAffinityId,omitempty"`

	// SourceZoneAffinityName
	// Zone Affinity Profile (Source)
	SourceZoneAffinityName *string `json:"sourceZoneAffinityName,omitempty"`

	// WlanId
	// Wlan ID
	WlanId *string `json:"wlanId,omitempty"`

	// WlanName
	// Wlan name
	WlanName *string `json:"wlanName,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`

	// ZoneName
	// Zone name
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGProfileFlexiVpnProfile() *WSGProfileFlexiVpnProfile {
	m := new(WSGProfileFlexiVpnProfile)
	return m
}

// WSGProfileFlexiVpnProfileList
//
// Definition: profile_flexiVpnProfileList
type WSGProfileFlexiVpnProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileFlexiVpnProfile `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileFlexiVpnProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileFlexiVpnProfileList
}

func newWSGProfileFlexiVpnProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileFlexiVpnProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileFlexiVpnProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileFlexiVpnProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileFlexiVpnProfileList() *WSGProfileFlexiVpnProfileList {
	m := new(WSGProfileFlexiVpnProfileList)
	return m
}

// WSGProfileGetL3RoamingConfig
//
// Definition: profile_getL3RoamingConfig
type WSGProfileGetL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*WSGProfileDataPlaneL3RoamingData `json:"dataPlanes,omitempty"`
}

type WSGProfileGetL3RoamingConfigAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileGetL3RoamingConfig
}

func newWSGProfileGetL3RoamingConfigAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileGetL3RoamingConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileGetL3RoamingConfigAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileGetL3RoamingConfig)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileGetL3RoamingConfig() *WSGProfileGetL3RoamingConfig {
	m := new(WSGProfileGetL3RoamingConfig)
	return m
}

// WSGProfileHs20FriendlyName
//
// Definition: profile_hs20FriendlyName
type WSGProfileHs20FriendlyName struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonLanguageName `json:"language"`

	// Name
	// Name of the friendly name
	// Constraints:
	//    - required
	//    - max:32
	//    - min:2
	Name *string `json:"name"`
}

func NewWSGProfileHs20FriendlyName() *WSGProfileHs20FriendlyName {
	m := new(WSGProfileHs20FriendlyName)
	return m
}

// WSGProfileHs20Operator
//
// Definition: profile_hs20Operator
type WSGProfileHs20Operator struct {
	Certificate *WSGCommonGenericRef `json:"certificate,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	// Constraints:
	//    - required
	DomainNames []*WSGCommonWildFQDN `json:"domainNames"`

	// FriendlyNames
	// Friendly names
	// Constraints:
	//    - required
	FriendlyNames []*WSGProfileHs20FriendlyName `json:"friendlyNames"`

	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

type WSGProfileHs20OperatorAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileHs20Operator
}

func newWSGProfileHs20OperatorAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileHs20OperatorAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileHs20OperatorAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileHs20Operator)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileHs20Operator() *WSGProfileHs20Operator {
	m := new(WSGProfileHs20Operator)
	return m
}

// WSGProfileHs20OperatorList
//
// Definition: profile_hs20OperatorList
type WSGProfileHs20OperatorList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileHs20Operator `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileHs20OperatorListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileHs20OperatorList
}

func newWSGProfileHs20OperatorListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileHs20OperatorListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileHs20OperatorListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileHs20OperatorList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileHs20OperatorList() *WSGProfileHs20OperatorList {
	m := new(WSGProfileHs20OperatorList)
	return m
}

// WSGProfileHs20Provider
//
// Definition: profile_hs20Provider
type WSGProfileHs20Provider struct {
	// Accountings
	// Accountings
	Accountings []*WSGProfileProviderAccounting `json:"accountings,omitempty"`

	// Authentications
	// Authentications
	Authentications []*WSGProfileProviderAuthentication `json:"authentications,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// HomeOis
	// Home OIs
	HomeOis []*WSGProfileProviderHomeOIs `json:"homeOis,omitempty"`

	// Id
	// Identifier of the Hotspot 2.0 identity provider profile
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Osu *WSGProfileProviderOnlineSignup `json:"osu,omitempty"`

	// Plmns
	// PLMNs
	Plmns []*WSGProfileProviderPLMN `json:"plmns,omitempty"`

	// Realms
	// Realms
	Realms []*WSGProfileProviderRealm `json:"realms,omitempty"`
}

type WSGProfileHs20ProviderAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileHs20Provider
}

func newWSGProfileHs20ProviderAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileHs20ProviderAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileHs20ProviderAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileHs20Provider)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileHs20Provider() *WSGProfileHs20Provider {
	m := new(WSGProfileHs20Provider)
	return m
}

// WSGProfileHs20ProviderList
//
// Definition: profile_hs20ProviderList
type WSGProfileHs20ProviderList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileHs20Provider `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileHs20ProviderListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileHs20ProviderList
}

func newWSGProfileHs20ProviderListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileHs20ProviderListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileHs20ProviderListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileHs20ProviderList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileHs20ProviderList() *WSGProfileHs20ProviderList {
	m := new(WSGProfileHs20ProviderList)
	return m
}

// WSGProfileIkeProposal
//
// Definition: profile_ikeProposal
//
// IkeProposal based ipsec service mappings
type WSGProfileIkeProposal struct {
	// AuthAlg
	// authAlg of ikeProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[MD5,SHA1,AESXCBC,SHA256,SHA384,SHA512]
	AuthAlg *string `json:"authAlg"`

	// DhGroup
	// dhGroup of ikeProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[Modp768,Modp1024,Modp1536,Modp2048,Modp3072,Modp4096,Modp6144,Modp8192,Ecp384]
	DhGroup *string `json:"dhGroup"`

	// EncAlg
	// encAlg of ikeProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[ThreeDES,AES128,AES192,AES256]
	EncAlg *string `json:"encAlg"`

	// PrfAlg
	// prfAlg of ikeProposal Specific
	// Constraints:
	//    - oneof:[UseIntegrityALG,PRF_MD5,PRF_SHA1,PRF_AES_CBC,PRF_AES_MAC,PRF_SHA256,PRF_SHA384,PRF_SHA512]
	PrfAlg *string `json:"prfAlg,omitempty"`
}

func NewWSGProfileIkeProposal() *WSGProfileIkeProposal {
	m := new(WSGProfileIkeProposal)
	return m
}

// WSGProfileIkeSecurityAssociationContent
//
// Definition: profile_ikeSecurityAssociationContent
//
// ikeProposal Security Association Content
type WSGProfileIkeSecurityAssociationContent struct {
	// IkeProposals
	// ikeProposal list of the ipsec profile
	IkeProposals []*WSGProfileIkeProposal `json:"ikeProposals,omitempty"`

	// IkeProposalType
	// ikeProposal Type of the ipsec profile
	// Constraints:
	//    - oneof:[Default,Specific]
	IkeProposalType *string `json:"ikeProposalType,omitempty"`
}

func NewWSGProfileIkeSecurityAssociationContent() *WSGProfileIkeSecurityAssociationContent {
	m := new(WSGProfileIkeSecurityAssociationContent)
	return m
}

// WSGProfileIpAclRules
//
// Definition: profile_ipAclRules
type WSGProfileIpAclRules struct {
	// Action
	// The access of traffic access control.
	// Constraints:
	//    - default:'ALLOW'
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	// Constraints:
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestinationIp
	// Subnet network address or ip address of destination IP.
	DestinationIp *string `json:"destinationIp,omitempty"`

	// DestinationIpMask
	// Subnet mask of destination IP
	DestinationIpMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIpV6
	// Destination IPv6 Address.
	DestinationIpV6 *string `json:"destinationIpV6,omitempty"`

	// DestinationMaxPort
	// The maxinum port of destination port range.
	DestinationMaxPort *int `json:"destinationMaxPort,omitempty"`

	// DestinationMinPort
	// The mininum port of destination port range.
	DestinationMinPort *int `json:"destinationMinPort,omitempty"`

	// Direction
	// The direction of traffic access control.
	// Constraints:
	//    - default:'UPSTREAM'
	//    - oneof:[UPSTREAM]
	Direction *string `json:"direction,omitempty"`

	// DownlinkRateLimitingEnabled
	// Downlink rate limiting enabled
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// EnableDestinationIpSubnet
	// Destination IP subnet enabled or disabled
	EnableDestinationIpSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIpSubnet
	// Source IP subnet enabled or disabled
	EnableSourceIpSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	EnableSourceV6Prefix *bool `json:"enableSourceV6Prefix,omitempty"`

	// IpType
	// IP Type(IPv4 or IPv6).
	// Constraints:
	//    - default:'IPv4'
	//    - oneof:[IPv4,IPv6]
	IpType *string `json:"ipType,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty"`

	// SourceIp
	// Subnet network address or ip address of source IP.
	SourceIp *string `json:"sourceIp,omitempty"`

	// SourceIpMask
	// Subnet mask of source IP
	SourceIpMask *string `json:"sourceIpMask,omitempty"`

	// SourceIpV6
	// Source IPv6 Address.
	SourceIpV6 *string `json:"sourceIpV6,omitempty"`

	// SourceMaxPort
	// The maxinum port of source port range.
	SourceMaxPort *int `json:"sourceMaxPort,omitempty"`

	// SourceMinPort
	// The minunum port of source port range.
	SourceMinPort *int `json:"sourceMinPort,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`
}

func NewWSGProfileIpAclRules() *WSGProfileIpAclRules {
	m := new(WSGProfileIpAclRules)
	return m
}

// WSGProfileIpMode
//
// Definition: profile_ipMode
//
// Constraints:
//    - oneof:[IPV4,IPV6]
type WSGProfileIpMode string

func NewWSGProfileIpMode() *WSGProfileIpMode {
	m := new(WSGProfileIpMode)
	return m
}

// WSGProfileIpsecProfile
//
// Definition: profile_ipsecProfile
type WSGProfileIpsecProfile struct {
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty"`

	CmProtocolOption *WSGProfileCmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	EspRekeyTimeUnit *WSGCommonTimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *WSGProfileEspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	IkeRekeyTimeUnit *WSGCommonTimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *WSGProfileIkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	IpMode *WSGProfileIpMode `json:"ipMode,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`

	// TunnelMode
	// Tunnel mode of IPsec profile
	// Constraints:
	//    - oneof:[SOFT_GRE,RUCKUS_GRE]
	TunnelMode *string `json:"tunnelMode,omitempty"`
}

type WSGProfileIpsecProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileIpsecProfile
}

func newWSGProfileIpsecProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileIpsecProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileIpsecProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileIpsecProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileIpsecProfile() *WSGProfileIpsecProfile {
	m := new(WSGProfileIpsecProfile)
	return m
}

// WSGProfileIpsecProfileList
//
// Definition: profile_ipsecProfileList
type WSGProfileIpsecProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileIpsecProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileIpsecProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileIpsecProfileList
}

func newWSGProfileIpsecProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileIpsecProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileIpsecProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileIpsecProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileIpsecProfileList() *WSGProfileIpsecProfileList {
	m := new(WSGProfileIpsecProfileList)
	return m
}

// WSGProfileL2oGREProfile
//
// Definition: profile_l2oGREProfile
type WSGProfileL2oGREProfile struct {
	CoreNetworkGateway *WSGProfileCoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

type WSGProfileL2oGREProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileL2oGREProfile
}

func newWSGProfileL2oGREProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileL2oGREProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileL2oGREProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileL2oGREProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileL2oGREProfile() *WSGProfileL2oGREProfile {
	m := new(WSGProfileL2oGREProfile)
	return m
}

// WSGProfileL2oGREProfileList
//
// Definition: profile_l2oGREProfileList
type WSGProfileL2oGREProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileL2oGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileL2oGREProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileL2oGREProfileList
}

func newWSGProfileL2oGREProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileL2oGREProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileL2oGREProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileL2oGREProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileL2oGREProfileList() *WSGProfileL2oGREProfileList {
	m := new(WSGProfileL2oGREProfileList)
	return m
}

// WSGProfileL3AccessControlPolicy
//
// Definition: profile_l3AccessControlPolicy
type WSGProfileL3AccessControlPolicy struct {
	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// ID of the L3 Access Control Policy
	Id *string `json:"id,omitempty"`

	// L3AclRuleList
	// L3 access control list
	L3AclRuleList []*WSGProfileL3AclRuleList `json:"l3AclRuleList,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

type WSGProfileL3AccessControlPolicyAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileL3AccessControlPolicy
}

func newWSGProfileL3AccessControlPolicyAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileL3AccessControlPolicyAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileL3AccessControlPolicyAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileL3AccessControlPolicy)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileL3AccessControlPolicy() *WSGProfileL3AccessControlPolicy {
	m := new(WSGProfileL3AccessControlPolicy)
	return m
}

// WSGProfileL3AccessControlPolicyArray
//
// Definition: profile_l3AccessControlPolicyArray
type WSGProfileL3AccessControlPolicyArray struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileL3AccessControlPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileL3AccessControlPolicyArrayAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileL3AccessControlPolicyArray
}

func newWSGProfileL3AccessControlPolicyArrayAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileL3AccessControlPolicyArrayAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileL3AccessControlPolicyArrayAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileL3AccessControlPolicyArray)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileL3AccessControlPolicyArray() *WSGProfileL3AccessControlPolicyArray {
	m := new(WSGProfileL3AccessControlPolicyArray)
	return m
}

// WSGProfileL3AclRuleList
//
// Definition: profile_l3AclRuleList
type WSGProfileL3AclRuleList struct {
	// Action
	// The access of traffic access control.
	// Constraints:
	//    - default:'ALLOW'
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	// Constraints:
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestinationIp
	// Subnet network address or ip address of destination IP.
	DestinationIp *string `json:"destinationIp,omitempty"`

	// DestinationIpMask
	// Subnet mask of destination IP
	DestinationIpMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIpV6
	// Destination IPv6 Address.
	DestinationIpV6 *string `json:"destinationIpV6,omitempty"`

	// DestinationMaxPort
	// The maxinum port of destination port range.
	DestinationMaxPort *int `json:"destinationMaxPort,omitempty"`

	// DestinationMinPort
	// The mininum port of destination port range.
	DestinationMinPort *int `json:"destinationMinPort,omitempty"`

	// Direction
	// The direction of traffic access control.
	// Constraints:
	//    - default:'INBOUND'
	//    - oneof:[INBOUND,OUTBOUND,DUAL]
	Direction *string `json:"direction,omitempty"`

	// EnableDestinationIpSubnet
	// Destination IP subnet enabled or disabled
	EnableDestinationIpSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIpSubnet
	// Source IP subnet enabled or disabled
	EnableSourceIpSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	EnableSourceV6Prefix *bool `json:"enableSourceV6Prefix,omitempty"`

	// IpType
	// IP Type(IPv4 or IPv6).
	// Constraints:
	//    - default:'IPv4'
	//    - oneof:[IPv4,IPv6]
	IpType *string `json:"ipType,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty"`

	// SourceIp
	// Subnet network address or ip address of source IP.
	SourceIp *string `json:"sourceIp,omitempty"`

	// SourceIpMask
	// Subnet mask of source IP
	SourceIpMask *string `json:"sourceIpMask,omitempty"`

	// SourceIpV6
	// Source IPv6 Address.
	SourceIpV6 *string `json:"sourceIpV6,omitempty"`

	// SourceMaxPort
	// The maxinum port of source port range.
	SourceMaxPort *int `json:"sourceMaxPort,omitempty"`

	// SourceMinPort
	// The minunum port of source port range.
	SourceMinPort *int `json:"sourceMinPort,omitempty"`
}

func NewWSGProfileL3AclRuleList() *WSGProfileL3AclRuleList {
	m := new(WSGProfileL3AclRuleList)
	return m
}

// WSGProfileLbsProfile
//
// Definition: profile_lbsProfile
type WSGProfileLbsProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Password
	// Password
	Password *string `json:"password,omitempty"`

	// Port
	// LBS port
	Port *int `json:"port,omitempty"`

	// Url
	// LBS url
	Url *string `json:"url,omitempty"`

	// Venue
	// Venue
	Venue *string `json:"venue,omitempty"`
}

type WSGProfileLbsProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileLbsProfile
}

func newWSGProfileLbsProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileLbsProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileLbsProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileLbsProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileLbsProfile() *WSGProfileLbsProfile {
	m := new(WSGProfileLbsProfile)
	return m
}

// WSGProfileLbsProfileList
//
// Definition: profile_lbsProfileList
type WSGProfileLbsProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileLbsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileLbsProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileLbsProfileList
}

func newWSGProfileLbsProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileLbsProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileLbsProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileLbsProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileLbsProfileList() *WSGProfileLbsProfileList {
	m := new(WSGProfileLbsProfileList)
	return m
}

// WSGProfileModifyAccountingProfile
//
// Definition: profile_modifyAccountingProfile
type WSGProfileModifyAccountingProfile struct {
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// ID of Accounting Profile
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*WSGProfileAcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

func NewWSGProfileModifyAccountingProfile() *WSGProfileModifyAccountingProfile {
	m := new(WSGProfileModifyAccountingProfile)
	return m
}

// WSGProfileModifyAuthenticationProfile
//
// Definition: profile_modifyAuthenticationProfile
type WSGProfileModifyAuthenticationProfile struct {
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	GppSuppportEnabled *bool `json:"gppSuppportEnabled,omitempty"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// Id
	// ID of Accounting Profile
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*WSGProfileRealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *WSGProfileTtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewWSGProfileModifyAuthenticationProfile() *WSGProfileModifyAuthenticationProfile {
	m := new(WSGProfileModifyAuthenticationProfile)
	return m
}

// WSGProfileModifyBlockClient
//
// Definition: profile_modifyBlockClient
type WSGProfileModifyBlockClient struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`
}

func NewWSGProfileModifyBlockClient() *WSGProfileModifyBlockClient {
	m := new(WSGProfileModifyBlockClient)
	return m
}

// WSGProfileModifyBonjourFencingPolicy
//
// Definition: profile_modifyBonjourFencingPolicy
type WSGProfileModifyBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	BonjourFencingRuleList []*WSGProfileBonjourFencingRule `json:"bonjourFencingRuleList,omitempty"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*WSGProfileBonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyBonjourFencingPolicy() *WSGProfileModifyBonjourFencingPolicy {
	m := new(WSGProfileModifyBonjourFencingPolicy)
	return m
}

// WSGProfileModifyBridgeProfile
//
// Definition: profile_modifyBridgeProfile
type WSGProfileModifyBridgeProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyBridgeProfile() *WSGProfileModifyBridgeProfile {
	m := new(WSGProfileModifyBridgeProfile)
	return m
}

// WSGProfileModifyClientIsolationWhitelist
//
// Definition: profile_modifyClientIsolationWhitelist
type WSGProfileModifyClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	Whitelist []*WSGProfileClientIsolationEntry `json:"whitelist,omitempty"`
}

func NewWSGProfileModifyClientIsolationWhitelist() *WSGProfileModifyClientIsolationWhitelist {
	m := new(WSGProfileModifyClientIsolationWhitelist)
	return m
}

// WSGProfileModifyDnsServerProfile
//
// Definition: profile_modifyDnsServerProfile
type WSGProfileModifyDnsServerProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryIp
	// Primary ip of DNS server service
	PrimaryIp *string `json:"primaryIp,omitempty"`

	// SecondaryIp
	// Secondary ip of DNS server service
	SecondaryIp *string `json:"secondaryIp,omitempty"`

	// TertiaryIp
	// Tertiary ip of DNS server service
	TertiaryIp *string `json:"tertiaryIp,omitempty"`
}

func NewWSGProfileModifyDnsServerProfile() *WSGProfileModifyDnsServerProfile {
	m := new(WSGProfileModifyDnsServerProfile)
	return m
}

// WSGProfileModifyFirewallProfile
//
// Definition: profile_modifyFirewallProfile
type WSGProfileModifyFirewallProfile struct {
	// AppPolicyId
	// Application Policy
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DevicePolicyId
	// Device Policy
	DevicePolicyId *string `json:"devicePolicyId,omitempty"`

	// DomainId
	// Domain Id of The Firewall Profile
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	// L3AccessControlPolicyId
	// L3 Access Control Policy
	L3AccessControlPolicyId *string `json:"l3AccessControlPolicyId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`

	// UrlFilteringPolicyId
	// Url Filtering Policy
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileModifyFirewallProfile() *WSGProfileModifyFirewallProfile {
	m := new(WSGProfileModifyFirewallProfile)
	return m
}

// WSGProfileModifyHS20Operator
//
// Definition: profile_modifyHS20Operator
type WSGProfileModifyHS20Operator struct {
	Certificate *WSGCommonGenericRef `json:"certificate,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	DomainNames []*WSGCommonWildFQDN `json:"domainNames,omitempty"`

	// FriendlyNames
	// Friendly names
	FriendlyNames []*WSGProfileHs20FriendlyName `json:"friendlyNames,omitempty"`

	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyHS20Operator() *WSGProfileModifyHS20Operator {
	m := new(WSGProfileModifyHS20Operator)
	return m
}

// WSGProfileModifyIpAclRules
//
// Definition: profile_modifyIpAclRules
//
// Traffic access control list
type WSGProfileModifyIpAclRules struct {
	// Action
	// The access of traffic access control.
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	// Constraints:
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestinationIp
	// Subnet network address or ip address of destination IP.
	DestinationIp *string `json:"destinationIp,omitempty"`

	// DestinationIpMask
	// Subnet mask of destination IP
	DestinationIpMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIpV6
	// Destination IPv6 Address.
	DestinationIpV6 *string `json:"destinationIpV6,omitempty"`

	// DestinationMaxPort
	// The maxinum port of destination port range.
	DestinationMaxPort *int `json:"destinationMaxPort,omitempty"`

	// DestinationMinPort
	// The mininum port of destination port range.
	DestinationMinPort *int `json:"destinationMinPort,omitempty"`

	// Direction
	// The direction of traffic access control.
	// Constraints:
	//    - required
	//    - default:'UPSTREAM'
	//    - oneof:[UPSTREAM]
	Direction *string `json:"direction"`

	// DownlinkRateLimitingEnabled
	// Downlink rate limiting enabled
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// EnableDestinationIpSubnet
	// Destination IP subnet enabled or disabled
	EnableDestinationIpSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIpSubnet
	// Source IP subnet enabled or disabled
	EnableSourceIpSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	EnableSourceV6Prefix *bool `json:"enableSourceV6Prefix,omitempty"`

	// IpType
	// IP Type(IPv4 or IPv6)
	// Constraints:
	//    - default:'IPv4'
	//    - oneof:[IPv4,IPv6]
	IpType *string `json:"ipType,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty"`

	// SourceIp
	// Subnet network address or ip address of source IP.
	SourceIp *string `json:"sourceIp,omitempty"`

	// SourceIpMask
	// Subnet mask of source IP
	SourceIpMask *string `json:"sourceIpMask,omitempty"`

	// SourceIpV6
	// Source IPv6 Address.
	SourceIpV6 *string `json:"sourceIpV6,omitempty"`

	// SourceMaxPort
	// The maxinum port of source port range.
	SourceMaxPort *int `json:"sourceMaxPort,omitempty"`

	// SourceMinPort
	// The minunum port of source port range.
	SourceMinPort *int `json:"sourceMinPort,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`
}

func NewWSGProfileModifyIpAclRules() *WSGProfileModifyIpAclRules {
	m := new(WSGProfileModifyIpAclRules)
	return m
}

// WSGProfileModifyIpsecProfile
//
// Definition: profile_modifyIpsecProfile
type WSGProfileModifyIpsecProfile struct {
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty"`

	CmProtocolOption *WSGProfileCmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	EspRekeyTimeUnit *WSGCommonTimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *WSGProfileEspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	IkeRekeyTimeUnit *WSGCommonTimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *WSGProfileIkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	IpMode *WSGProfileIpMode `json:"ipMode,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`
}

func NewWSGProfileModifyIpsecProfile() *WSGProfileModifyIpsecProfile {
	m := new(WSGProfileModifyIpsecProfile)
	return m
}

// WSGProfileModifyL2oGREProfile
//
// Definition: profile_modifyL2oGREProfile
type WSGProfileModifyL2oGREProfile struct {
	CoreNetworkGateway *WSGProfileCoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyL2oGREProfile() *WSGProfileModifyL2oGREProfile {
	m := new(WSGProfileModifyL2oGREProfile)
	return m
}

// WSGProfileModifyL3AccessControlPolicy
//
// Definition: profile_modifyL3AccessControlPolicy
type WSGProfileModifyL3AccessControlPolicy struct {
	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// L3AclRuleList
	// L3 access control list
	L3AclRuleList []*WSGProfileL3AclRuleList `json:"l3AclRuleList,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileModifyL3AccessControlPolicy() *WSGProfileModifyL3AccessControlPolicy {
	m := new(WSGProfileModifyL3AccessControlPolicy)
	return m
}

// WSGProfileModifyRestrictedApAccessProfile
//
// Definition: profile_modifyRestrictedApAccessProfile
type WSGProfileModifyRestrictedApAccessProfile struct {
	// BlockedPortList
	// Blocked Port List
	BlockedPortList []*WSGProfileBlockedPort `json:"blockedPortList,omitempty"`

	// BlockWellKnownPort
	// Block well known ports
	BlockWellKnownPort *bool `json:"blockWellKnownPort,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IpAddressWhitelist
	// IP Address Whitelist
	IpAddressWhitelist []*WSGCommonIpAddress `json:"ipAddressWhitelist,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGProfileModifyRestrictedApAccessProfile() *WSGProfileModifyRestrictedApAccessProfile {
	m := new(WSGProfileModifyRestrictedApAccessProfile)
	return m
}

// WSGProfileModifyRuckusGREProfile
//
// Definition: profile_modifyRuckusGREProfile
type WSGProfileModifyRuckusGREProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// TunnelEncryption
	// Tunnel Encryption of the RuckusGRE profile
	// Constraints:
	//    - oneof:[DISABLE,AES128,AES256]
	TunnelEncryption *string `json:"tunnelEncryption,omitempty"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileModifyRuckusGREProfile() *WSGProfileModifyRuckusGREProfile {
	m := new(WSGProfileModifyRuckusGREProfile)
	return m
}

// WSGProfileModifySoftGREProfile
//
// Definition: profile_modifySoftGREProfile
type WSGProfileModifySoftGREProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:180
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - default:5
	//    - min:2
	//    - max:20
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	// Constraints:
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile. IPV4:850-1500, IPV6:1384-1500. Default 1500.
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileModifySoftGREProfile() *WSGProfileModifySoftGREProfile {
	m := new(WSGProfileModifySoftGREProfile)
	return m
}

// WSGProfileModifyUserTrafficProfile
//
// Definition: profile_modifyUserTrafficProfile
type WSGProfileModifyUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *WSGProfileDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// Id
	// Identifier of the user traffic profile
	Id *string `json:"id,omitempty"`

	// IpAclRules
	// Traffic access control list
	IpAclRules []*WSGProfileModifyIpAclRules `json:"ipAclRules,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// QmAppPolicyId
	// Application Policy UUID
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *WSGProfileUplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileModifyUserTrafficProfile() *WSGProfileModifyUserTrafficProfile {
	m := new(WSGProfileModifyUserTrafficProfile)
	return m
}

// WSGProfileModifyZoneAffinityProfile
//
// Definition: profile_modifyZoneAffinityProfile
type WSGProfileModifyZoneAffinityProfile struct {
	// Description
	// The description of the profile
	// Constraints:
	//    - max:64
	Description *string `json:"description,omitempty"`

	// Name
	// Zone affinity profile name
	// Constraints:
	//    - max:64
	//    - min:1
	Name *string `json:"name,omitempty"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
}

func NewWSGProfileModifyZoneAffinityProfile() *WSGProfileModifyZoneAffinityProfile {
	m := new(WSGProfileModifyZoneAffinityProfile)
	return m
}

// WSGProfilePrecedenceList
//
// Definition: profile_precedenceList
type WSGProfilePrecedenceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfilePrecedenceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfilePrecedenceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfilePrecedenceList
}

func newWSGProfilePrecedenceListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfilePrecedenceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfilePrecedenceListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfilePrecedenceList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfilePrecedenceList() *WSGProfilePrecedenceList {
	m := new(WSGProfilePrecedenceList)
	return m
}

// WSGProfilePrecedenceListType
//
// Definition: profile_precedenceListType
type WSGProfilePrecedenceListType struct {
	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*WSGProfileRateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*WSGProfileVlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

func NewWSGProfilePrecedenceListType() *WSGProfilePrecedenceListType {
	m := new(WSGProfilePrecedenceListType)
	return m
}

// WSGProfileClone
//
// Definition: profile_profileClone
type WSGProfileClone struct {
	// NewId
	// name for new profile
	NewId *string `json:"newId,omitempty"`

	// NewName
	// Id for new profile
	NewName *string `json:"newName,omitempty"`

	// OldId
	// original name
	OldId *string `json:"oldId,omitempty"`

	// OldName
	// original name
	OldName *string `json:"oldName,omitempty"`
}

type WSGProfileCloneAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileClone
}

func newWSGProfileCloneAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileCloneAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileCloneAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileClone)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileClone() *WSGProfileClone {
	m := new(WSGProfileClone)
	return m
}

// WSGProfileIdList
//
// Definition: profile_profileIdList
type WSGProfileIdList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCommonGenericRef `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileIdListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileIdList
}

func newWSGProfileIdListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileIdListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileIdListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileIdList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileIdList() *WSGProfileIdList {
	m := new(WSGProfileIdList)
	return m
}

// WSGProfileList
//
// Definition: profile_profileList
type WSGProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileListType `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileList
}

func newWSGProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileList() *WSGProfileList {
	m := new(WSGProfileList)
	return m
}

// WSGProfileListType
//
// Definition: profile_profileListType
type WSGProfileListType struct {
	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileListType() *WSGProfileListType {
	m := new(WSGProfileListType)
	return m
}

// WSGProfileProviderAccounting
//
// Definition: profile_providerAccounting
type WSGProfileProviderAccounting struct {
	// Id
	// Accounting id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Accounting name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// ServiceType
	// Accounting service type
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,CGF]
	ServiceType *string `json:"serviceType"`
}

func NewWSGProfileProviderAccounting() *WSGProfileProviderAccounting {
	m := new(WSGProfileProviderAccounting)
	return m
}

// WSGProfileProviderAuthentication
//
// Definition: profile_providerAuthentication
type WSGProfileProviderAuthentication struct {
	// Id
	// Authentication id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// ServiceType
	// Authentication service type
	// Constraints:
	//    - required
	//    - oneof:[NA,LOCAL_DB,RADIUS,GUEST]
	ServiceType *string `json:"serviceType"`

	// VlanId
	// Dynamic vlan ID
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	VlanId *int `json:"vlanId,omitempty"`
}

func NewWSGProfileProviderAuthentication() *WSGProfileProviderAuthentication {
	m := new(WSGProfileProviderAuthentication)
	return m
}

// WSGProfileProviderEAPAuthSetting
//
// Definition: profile_providerEAPAuthSetting
type WSGProfileProviderEAPAuthSetting struct {
	// Info
	// EAP auth info
	// Constraints:
	//    - required
	//    - oneof:[Expanded,Non,Inner,Expanded_Inner,Credential,Tunneled]
	Info *string `json:"info"`

	// Type
	// EAP auth type
	Type *string `json:"type,omitempty"`

	// VendorId
	// EAP auth vendor ID
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:16777215
	VendorId *int `json:"vendorId,omitempty"`

	// VendorType
	// EAP auth vendor type
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:4294967295
	VendorType *int `json:"vendorType,omitempty"`
}

func NewWSGProfileProviderEAPAuthSetting() *WSGProfileProviderEAPAuthSetting {
	m := new(WSGProfileProviderEAPAuthSetting)
	return m
}

// WSGProfileProviderEAPMethod
//
// Definition: profile_providerEAPMethod
type WSGProfileProviderEAPMethod struct {
	// AuthSettings
	// EAP method auth settings
	// Constraints:
	//    - nullable
	AuthSettings []*WSGProfileProviderEAPAuthSetting `json:"authSettings,omitempty"`

	// Type
	// EAP method type
	// Constraints:
	//    - required
	//    - oneof:[NA,MD5,EAP_TLS,EAP_Cisco,EAP_SIM,EAP_TTLS,EAP_AKA,PEAP,EAP_MSCHAP_V2,EAP_AKAs,Reserved]
	Type *string `json:"type"`
}

func NewWSGProfileProviderEAPMethod() *WSGProfileProviderEAPMethod {
	m := new(WSGProfileProviderEAPMethod)
	return m
}

// WSGProfileProviderExternalOSU
//
// Definition: profile_providerExternalOSU
type WSGProfileProviderExternalOSU struct {
	// CommonLanguageIcon
	// The base64 encoded data of icon.
	// Constraints:
	//    - required
	CommonLanguageIcon *string `json:"commonLanguageIcon"`

	// OsuNaiRealm
	// Online signup NAI realm, it should be one of realm as defined in Hotspot 2.0 identity provider
	// Constraints:
	//    - required
	OsuNaiRealm *string `json:"osuNaiRealm"`

	// OsuServiceUrl
	// Constraints:
	//    - required
	OsuServiceUrl *WSGCommonHTTPS `json:"osuServiceUrl"`

	// ProvisioningProtocals
	// Provisioning protocal
	// Constraints:
	//    - required
	ProvisioningProtocals []*WSGProfileProviderProvisionProtocal `json:"provisioningProtocals"`

	// SubscriptionDescriptions
	// Subscription descriptions
	// Constraints:
	//    - required
	SubscriptionDescriptions []*WSGProfileProviderSubscriptionDescription `json:"subscriptionDescriptions"`

	// WhitelistedDomains
	// Whitelisted domains
	WhitelistedDomains []*WSGCommonWildFQDN `json:"whitelistedDomains,omitempty"`
}

func NewWSGProfileProviderExternalOSU() *WSGProfileProviderExternalOSU {
	m := new(WSGProfileProviderExternalOSU)
	return m
}

// WSGProfileProviderHomeOIs
//
// Definition: profile_providerHomeOIs
type WSGProfileProviderHomeOIs struct {
	// Name
	// Name of the home OI
	// Constraints:
	//    - required
	//    - max:255
	Name *string `json:"name"`

	// Oi
	// Orgnization ID(3Hex or 5Hex)
	// Constraints:
	//    - required
	Oi *string `json:"oi"`
}

func NewWSGProfileProviderHomeOIs() *WSGProfileProviderHomeOIs {
	m := new(WSGProfileProviderHomeOIs)
	return m
}

// WSGProfileProviderOnlineSignup
//
// Definition: profile_providerOnlineSignup
type WSGProfileProviderOnlineSignup struct {
	ExternalOSU *WSGProfileProviderExternalOSU `json:"externalOSU,omitempty"`
}

func NewWSGProfileProviderOnlineSignup() *WSGProfileProviderOnlineSignup {
	m := new(WSGProfileProviderOnlineSignup)
	return m
}

// WSGProfileProviderPLMN
//
// Definition: profile_providerPLMN
type WSGProfileProviderPLMN struct {
	// Mcc
	// MCC
	// Constraints:
	//    - required
	Mcc *string `json:"mcc"`

	// Mnc
	// MNC
	// Constraints:
	//    - required
	Mnc *string `json:"mnc"`
}

func NewWSGProfileProviderPLMN() *WSGProfileProviderPLMN {
	m := new(WSGProfileProviderPLMN)
	return m
}

// WSGProfileProviderProvisionProtocal
//
// Definition: profile_providerProvisionProtocal
//
// Constraints:
//    - oneof:[SOAP_XML,OMA_DM]
type WSGProfileProviderProvisionProtocal string

func NewWSGProfileProviderProvisionProtocal() *WSGProfileProviderProvisionProtocal {
	m := new(WSGProfileProviderProvisionProtocal)
	return m
}

// WSGProfileProviderRealm
//
// Definition: profile_providerRealm
type WSGProfileProviderRealm struct {
	// EapMethods
	// EAP methods
	// Constraints:
	//    - required
	EapMethods []*WSGProfileProviderEAPMethod `json:"eapMethods"`

	// Encoding
	// Encoding
	// Constraints:
	//    - required
	//    - oneof:[RFC4282,UTF8]
	Encoding *string `json:"encoding"`

	// Name
	// Name of realm
	// Constraints:
	//    - required
	//    - max:243
	//    - min:2
	Name *string `json:"name"`
}

func NewWSGProfileProviderRealm() *WSGProfileProviderRealm {
	m := new(WSGProfileProviderRealm)
	return m
}

// WSGProfileProviderSubscriptionDescription
//
// Definition: profile_providerSubscriptionDescription
type WSGProfileProviderSubscriptionDescription struct {
	// Description
	// Description of the friendly name
	// Constraints:
	//    - nullable
	//    - max:64
	Description *string `json:"description,omitempty"`

	// Icon
	// The binary data of icon, maximum size 65536
	Icon *string `json:"icon,omitempty"`

	// Language
	// Constraints:
	//    - required
	Language *WSGCommonLanguageName `json:"language"`

	// Name
	// Name of the friendly name
	// Constraints:
	//    - required
	//    - max:252
	//    - min:2
	Name *string `json:"name"`
}

func NewWSGProfileProviderSubscriptionDescription() *WSGProfileProviderSubscriptionDescription {
	m := new(WSGProfileProviderSubscriptionDescription)
	return m
}

// WSGProfileRateLimitingPrecedenceItem
//
// Definition: profile_RateLimitingPrecedenceItem
//
// Rate limiting precedence item
type WSGProfileRateLimitingPrecedenceItem struct {
	// Name
	// Name of rate limiting precedence item
	// Constraints:
	//    - oneof:[AAA,DEVICE,WLANUTP]
	Name *string `json:"name,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

func NewWSGProfileRateLimitingPrecedenceItem() *WSGProfileRateLimitingPrecedenceItem {
	m := new(WSGProfileRateLimitingPrecedenceItem)
	return m
}

// WSGProfileRealmAuthServiceMapping
//
// Definition: profile_realmAuthServiceMapping
//
// Realm based authentication service mappings
type WSGProfileRealmAuthServiceMapping struct {
	// AuthorizationMethod
	// Authorization method
	// Constraints:
	//    - required
	//    - oneof:[NonGPPCallFlow,GPPCallFlow,UpdateGPRSLocation,RestoreData,NoAutz]
	AuthorizationMethod *string `json:"authorizationMethod"`

	// DynamicVlanId
	// Dynamic VLAN ID
	// Constraints:
	//    - nullable
	//    - min:2
	//    - max:4094
	DynamicVlanId *int `json:"dynamicVlanId,omitempty"`

	HostedAaaEnabled *bool `json:"hostedAaaEnabled,omitempty"`

	// Id
	// Authentication service UUID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication service name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// ServiceType
	// Authentication service type, NA is NA-Request Rejected
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,LOCAL_DB,HLR,AD,LDAP]
	ServiceType *string `json:"serviceType"`
}

func NewWSGProfileRealmAuthServiceMapping() *WSGProfileRealmAuthServiceMapping {
	m := new(WSGProfileRealmAuthServiceMapping)
	return m
}

// WSGProfileRestrictedApAccessProfile
//
// Definition: profile_restrictedApAccessProfile
type WSGProfileRestrictedApAccessProfile struct {
	// BlockedPortList
	// Blocked Port List
	BlockedPortList []*WSGProfileBlockedPort `json:"blockedPortList,omitempty"`

	// BlockWellKnownPort
	// Block well known ports
	BlockWellKnownPort *bool `json:"blockWellKnownPort,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Restricted AP Access Profile id
	Id *string `json:"id,omitempty"`

	// IpAddressWhitelist
	// IP Address Whitelist
	IpAddressWhitelist []*WSGCommonIpAddress `json:"ipAddressWhitelist,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// ZoneId
	// Zone Id of The Restricted AP Access Profile for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGProfileRestrictedApAccessProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileRestrictedApAccessProfile
}

func newWSGProfileRestrictedApAccessProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileRestrictedApAccessProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileRestrictedApAccessProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileRestrictedApAccessProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileRestrictedApAccessProfile() *WSGProfileRestrictedApAccessProfile {
	m := new(WSGProfileRestrictedApAccessProfile)
	return m
}

// WSGProfileRestrictedApAccessProfileArray
//
// Definition: profile_restrictedApAccessProfileArray
type WSGProfileRestrictedApAccessProfileArray struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileRestrictedApAccessProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileRestrictedApAccessProfileArrayAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileRestrictedApAccessProfileArray
}

func newWSGProfileRestrictedApAccessProfileArrayAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileRestrictedApAccessProfileArrayAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileRestrictedApAccessProfileArrayAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileRestrictedApAccessProfileArray)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileRestrictedApAccessProfileArray() *WSGProfileRestrictedApAccessProfileArray {
	m := new(WSGProfileRestrictedApAccessProfileArray)
	return m
}

// WSGProfileReturnZoneAffinityProfile
//
// Definition: profile_returnZoneAffinityProfile
type WSGProfileReturnZoneAffinityProfile struct {
	// BaseDpVersion
	// The lowest DP version in an Zone Affinity Profile
	BaseDpVersion *string `json:"baseDpVersion,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// The description of the profile
	Description *string `json:"description,omitempty"`

	// Id
	// Zone affinity profile key
	Id *string `json:"id,omitempty"`

	// IsDpVersionConsistent
	// True if all DPs are the same version
	IsDpVersionConsistent *bool `json:"isDpVersionConsistent,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Zone affinity profile name
	// Constraints:
	//    - max:64
	//    - min:1
	Name *string `json:"name,omitempty"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`

	ZoneAffinityListWithPriority []*WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType `json:"zoneAffinityListWithPriority,omitempty"`
}

type WSGProfileReturnZoneAffinityProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileReturnZoneAffinityProfile
}

func newWSGProfileReturnZoneAffinityProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileReturnZoneAffinityProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileReturnZoneAffinityProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileReturnZoneAffinityProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileReturnZoneAffinityProfile() *WSGProfileReturnZoneAffinityProfile {
	m := new(WSGProfileReturnZoneAffinityProfile)
	return m
}

// WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType
//
// Definition: profile_returnZoneAffinityProfileZoneAffinityListWithPriorityType
type WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType struct {
	// DpId
	// DP ID
	DpId *string `json:"dpId,omitempty"`

	// Priority
	// The priority of DP in zone affinity
	Priority *float64 `json:"priority,omitempty"`
}

func NewWSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType() *WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType {
	m := new(WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType)
	return m
}

// WSGProfileRogueApPolicy
//
// Definition: profile_rogueApPolicy
type WSGProfileRogueApPolicy struct {
	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Id *string `json:"id,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Rules []*WSGProfileRogueApRuleList `json:"rules,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGProfileRogueApPolicyAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileRogueApPolicy
}

func newWSGProfileRogueApPolicyAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileRogueApPolicyAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileRogueApPolicyAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileRogueApPolicy)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileRogueApPolicy() *WSGProfileRogueApPolicy {
	m := new(WSGProfileRogueApPolicy)
	return m
}

// WSGProfileRogueApPolicyList
//
// Definition: profile_rogueApPolicyList
type WSGProfileRogueApPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileRogueApPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileRogueApPolicyListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileRogueApPolicyList
}

func newWSGProfileRogueApPolicyListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileRogueApPolicyListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileRogueApPolicyListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileRogueApPolicyList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileRogueApPolicyList() *WSGProfileRogueApPolicyList {
	m := new(WSGProfileRogueApPolicyList)
	return m
}

// WSGProfileRogueApRuleList
//
// Definition: profile_rogueApRuleList
type WSGProfileRogueApRuleList struct {
	// Classification
	// Constraints:
	//    - required
	//    - oneof:[Ignore,Known,Rogue,Malicious]
	Classification *string `json:"classification"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalNameAllowBlank `json:"name"`

	// Priority
	// Constraints:
	//    - required
	Priority *int `json:"priority"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[AdhocRule,SsidSpoofingRule,MacSpoofingRule,SameNetworkRule,CTSAbuseRule,RTSAbuseRule,DeauthFloodRule,DisassocFloodRule,ExcessivePowerRule,NullSSIDRule,CustomSsidRule,CustomRssiRule,CustomMacOuiRule]
	Type *string `json:"type"`

	Value *string `json:"value,omitempty"`
}

func NewWSGProfileRogueApRuleList() *WSGProfileRogueApRuleList {
	m := new(WSGProfileRogueApRuleList)
	return m
}

// WSGProfileRtlsProfileList
//
// Definition: profile_rtlsProfileList
type WSGProfileRtlsProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileCreateRtlsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileRtlsProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileRtlsProfileList
}

func newWSGProfileRtlsProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileRtlsProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileRtlsProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileRtlsProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileRtlsProfileList() *WSGProfileRtlsProfileList {
	m := new(WSGProfileRtlsProfileList)
	return m
}

// WSGProfileRuckusGREProfile
//
// Definition: profile_ruckusGREProfile
type WSGProfileRuckusGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// TunnelEncryption
	// Tunnel Encryption of the RuckusGRE profile
	// Constraints:
	//    - oneof:[DISABLE,AES128,AES256]
	TunnelEncryption *string `json:"tunnelEncryption,omitempty"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

type WSGProfileRuckusGREProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileRuckusGREProfile
}

func newWSGProfileRuckusGREProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileRuckusGREProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileRuckusGREProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileRuckusGREProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileRuckusGREProfile() *WSGProfileRuckusGREProfile {
	m := new(WSGProfileRuckusGREProfile)
	return m
}

// WSGProfileRuckusGREProfileList
//
// Definition: profile_ruckusGREProfileList
type WSGProfileRuckusGREProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileRuckusGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileRuckusGREProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileRuckusGREProfileList
}

func newWSGProfileRuckusGREProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileRuckusGREProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileRuckusGREProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileRuckusGREProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileRuckusGREProfileList() *WSGProfileRuckusGREProfileList {
	m := new(WSGProfileRuckusGREProfileList)
	return m
}

// WSGProfileSoftGREProfile
//
// Definition: profile_SoftGREProfile
type WSGProfileSoftGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	IpMode *WSGProfileIpMode `json:"ipMode,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	// Constraints:
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

type WSGProfileSoftGREProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileSoftGREProfile
}

func newWSGProfileSoftGREProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileSoftGREProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileSoftGREProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileSoftGREProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileSoftGREProfile() *WSGProfileSoftGREProfile {
	m := new(WSGProfileSoftGREProfile)
	return m
}

// WSGProfileSoftGREProfileList
//
// Definition: profile_softGREProfileList
type WSGProfileSoftGREProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileSoftGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileSoftGREProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileSoftGREProfileList
}

func newWSGProfileSoftGREProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileSoftGREProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileSoftGREProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileSoftGREProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileSoftGREProfileList() *WSGProfileSoftGREProfileList {
	m := new(WSGProfileSoftGREProfileList)
	return m
}

// WSGProfileTrafficClassProfileList
//
// Definition: profile_trafficClassProfileList
type WSGProfileTrafficClassProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCommonTrafficClassProfileRef `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileTrafficClassProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileTrafficClassProfileList
}

func newWSGProfileTrafficClassProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileTrafficClassProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileTrafficClassProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileTrafficClassProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileTrafficClassProfileList() *WSGProfileTrafficClassProfileList {
	m := new(WSGProfileTrafficClassProfileList)
	return m
}

// WSGProfileTtgCommonSetting
//
// Definition: profile_ttgCommonSetting
//
// Hosted AAA server RADIUS settings & PLMN ID settings
type WSGProfileTtgCommonSetting struct {
	// MobileCountryCode
	// Mobile country code
	// Constraints:
	//    - max:3
	//    - min:3
	MobileCountryCode *string `json:"mobileCountryCode,omitempty"`

	// MobileNetworkCode
	// Mobile network code
	// Constraints:
	//    - max:3
	//    - min:2
	MobileNetworkCode *string `json:"mobileNetworkCode,omitempty"`
}

func NewWSGProfileTtgCommonSetting() *WSGProfileTtgCommonSetting {
	m := new(WSGProfileTtgCommonSetting)
	return m
}

// WSGProfileTtgpdgApnForwardingRealm
//
// Definition: profile_ttgpdgApnForwardingRealm
type WSGProfileTtgpdgApnForwardingRealm struct {
	// Apn
	// the forwarding policy APN, if apnType is NIOI, APN Example : internet-v4.mnc111.mcc222.gprs
	Apn *string `json:"apn,omitempty"`

	// ApnType
	// type of the forwarding policy APN.
	// Constraints:
	//    - oneof:[NI,NIOI]
	ApnType *string `json:"apnType,omitempty"`

	// RouteType
	// routeType of the forwarding policy APN.
	// Constraints:
	//    - oneof:[GTPv1,GTPv2,PDG]
	RouteType *string `json:"routeType,omitempty"`
}

func NewWSGProfileTtgpdgApnForwardingRealm() *WSGProfileTtgpdgApnForwardingRealm {
	m := new(WSGProfileTtgpdgApnForwardingRealm)
	return m
}

// WSGProfileTtgpdgCommonSetting
//
// Definition: profile_ttgpdgCommonSetting
type WSGProfileTtgpdgCommonSetting struct {
	// AcctRetry
	// Accounting retry of TTG PDG common setting
	AcctRetry *int `json:"acctRetry,omitempty"`

	// AcctRetryTimeout
	// Accounting retry timeout(secs) of TTG PDG common setting
	AcctRetryTimeout *int `json:"acctRetryTimeout,omitempty"`

	// ApnFormat2GGSN
	// APN format to GGSN of TTG PDG common setting
	// Constraints:
	//    - oneof:[DNS,String]
	ApnFormat2GGSN *string `json:"apnFormat2GGSN,omitempty"`

	// ApnOIInUse
	// APN-OI of TTG PDG common setting
	ApnOIInUse *bool `json:"apnOIInUse,omitempty"`

	// PdgUeIdleTimeout
	// PDG UE session idle timeout(secs) of TTG PDG common setting
	PdgUeIdleTimeout *int `json:"pdgUeIdleTimeout,omitempty"`
}

func NewWSGProfileTtgpdgCommonSetting() *WSGProfileTtgpdgCommonSetting {
	m := new(WSGProfileTtgpdgCommonSetting)
	return m
}

// WSGProfileTtgpdgProfile
//
// Definition: profile_ttgpdgProfile
type WSGProfileTtgpdgProfile struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	ApnForwardingRealms []*WSGProfileTtgpdgApnForwardingRealm `json:"apnForwardingRealms,omitempty"`

	// ApnRealms
	// List of the Default APN
	ApnRealms []*WSGProfileApnRealm `json:"apnRealms,omitempty"`

	CommonSetting *WSGProfileTtgpdgCommonSetting `json:"commonSetting,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN,omitempty"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

type WSGProfileTtgpdgProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileTtgpdgProfile
}

func newWSGProfileTtgpdgProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileTtgpdgProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileTtgpdgProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileTtgpdgProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileTtgpdgProfile() *WSGProfileTtgpdgProfile {
	m := new(WSGProfileTtgpdgProfile)
	return m
}

// WSGProfileTtgpdgProfileConfiguration
//
// Definition: profile_ttgpdgProfileConfiguration
type WSGProfileTtgpdgProfileConfiguration struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	ApnForwardingRealms []*WSGProfileTtgpdgApnForwardingRealm `json:"apnForwardingRealms,omitempty"`

	// ApnRealms
	// List of the Default APN
	ApnRealms []*WSGProfileApnRealm `json:"apnRealms,omitempty"`

	CommonSetting *WSGProfileTtgpdgCommonSetting `json:"commonSetting,omitempty"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN,omitempty"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileTtgpdgProfileConfiguration() *WSGProfileTtgpdgProfileConfiguration {
	m := new(WSGProfileTtgpdgProfileConfiguration)
	return m
}

// WSGProfileTtgpdgProfileList
//
// Definition: profile_ttgpdgProfileList
type WSGProfileTtgpdgProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileTtgpdgProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileTtgpdgProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileTtgpdgProfileList
}

func newWSGProfileTtgpdgProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileTtgpdgProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileTtgpdgProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileTtgpdgProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileTtgpdgProfileList() *WSGProfileTtgpdgProfileList {
	m := new(WSGProfileTtgpdgProfileList)
	return m
}

// WSGProfileUpdateL3RoamingConfig
//
// Definition: profile_updateL3RoamingConfig
type WSGProfileUpdateL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*WSGProfileDataPlaneL3RoamingData `json:"dataPlanes,omitempty"`
}

func NewWSGProfileUpdateL3RoamingConfig() *WSGProfileUpdateL3RoamingConfig {
	m := new(WSGProfileUpdateL3RoamingConfig)
	return m
}

// WSGProfileUpdatePrecedenceProfile
//
// Definition: profile_updatePrecedenceProfile
type WSGProfileUpdatePrecedenceProfile struct {
	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*WSGProfileRateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*WSGProfileVlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

func NewWSGProfileUpdatePrecedenceProfile() *WSGProfileUpdatePrecedenceProfile {
	m := new(WSGProfileUpdatePrecedenceProfile)
	return m
}

// WSGProfileUpdateRogueApPolicy
//
// Definition: profile_updateRogueApPolicy
type WSGProfileUpdateRogueApPolicy struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Rules []*WSGProfileRogueApRuleList `json:"rules,omitempty"`
}

func NewWSGProfileUpdateRogueApPolicy() *WSGProfileUpdateRogueApPolicy {
	m := new(WSGProfileUpdateRogueApPolicy)
	return m
}

// WSGProfileUpdateRtlsProfile
//
// Definition: profile_updateRtlsProfile
type WSGProfileUpdateRtlsProfile struct {
	EkahauEnabled *bool `json:"ekahauEnabled,omitempty"`

	EkahauIp *WSGCommonIpAddress `json:"ekahauIp,omitempty"`

	EkahauPort *int `json:"ekahauPort,omitempty"`

	Id *string `json:"id,omitempty"`

	StanleyEnabled *bool `json:"stanleyEnabled,omitempty"`
}

func NewWSGProfileUpdateRtlsProfile() *WSGProfileUpdateRtlsProfile {
	m := new(WSGProfileUpdateRtlsProfile)
	return m
}

// WSGProfileUplinkRateLimiting
//
// Definition: profile_uplinkRateLimiting
type WSGProfileUplinkRateLimiting struct {
	// UplinkRateLimitingBps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	UplinkRateLimitingBps *string `json:"uplinkRateLimitingBps,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled or disabled
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`
}

func NewWSGProfileUplinkRateLimiting() *WSGProfileUplinkRateLimiting {
	m := new(WSGProfileUplinkRateLimiting)
	return m
}

// WSGProfileUserTrafficProfile
//
// Definition: profile_userTrafficProfile
type WSGProfileUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *WSGProfileDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// Id
	// Identifier of the user traffic profile
	Id *string `json:"id,omitempty"`

	// IpAclRules
	// Traffic access control list
	IpAclRules []*WSGProfileIpAclRules `json:"ipAclRules,omitempty"`

	// IsFactoryDefault
	// Whether the UTP is factory default or not
	IsFactoryDefault *bool `json:"isFactoryDefault,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// QmAppPolicyId
	// Application Policy UUID
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *WSGProfileUplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

type WSGProfileUserTrafficProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileUserTrafficProfile
}

func newWSGProfileUserTrafficProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileUserTrafficProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileUserTrafficProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileUserTrafficProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileUserTrafficProfile() *WSGProfileUserTrafficProfile {
	m := new(WSGProfileUserTrafficProfile)
	return m
}

// WSGProfileUserTrafficProfileList
//
// Definition: profile_userTrafficProfileList
type WSGProfileUserTrafficProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileUserTrafficProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGProfileUserTrafficProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileUserTrafficProfileList
}

func newWSGProfileUserTrafficProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileUserTrafficProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileUserTrafficProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileUserTrafficProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileUserTrafficProfileList() *WSGProfileUserTrafficProfileList {
	m := new(WSGProfileUserTrafficProfileList)
	return m
}

// WSGProfileVdpProfile
//
// Definition: profile_vdpProfile
type WSGProfileVdpProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DataVlan
	// data vlan
	DataVlan *int `json:"dataVlan,omitempty"`

	// ExtIp
	// external ip
	ExtIp *string `json:"extIp,omitempty"`

	// FwVersion
	// Firmware version
	FwVersion *string `json:"fwVersion,omitempty"`

	// Ip
	// IP
	Ip *string `json:"ip,omitempty"`

	// Ipv6
	// IPv6
	Ipv6 *string `json:"ipv6,omitempty"`

	// IsSupport
	// is support vdp
	IsSupport *bool `json:"isSupport,omitempty"`

	// LastSeenOn
	// last seen
	LastSeenOn *string `json:"lastSeenOn,omitempty"`

	// Mac
	// mac
	Mac *string `json:"mac,omitempty"`

	// ManagedBy
	// managed by
	ManagedBy *string `json:"managedBy,omitempty"`

	// MgmtExtIp
	// management external ip
	MgmtExtIp *string `json:"mgmtExtIp,omitempty"`

	// MgmtIp
	// management ip
	MgmtIp *string `json:"mgmtIp,omitempty"`

	// MgmtVlan
	// management vlan
	MgmtVlan *int `json:"mgmtVlan,omitempty"`

	// Model
	// model
	Model *string `json:"model,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// name of vdp  profile
	Name *string `json:"name,omitempty"`

	// RegistrationState
	// registrationState
	RegistrationState *string `json:"registrationState,omitempty"`

	// SerialNumber
	// serialNumber
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Status
	// status
	Status *string `json:"status,omitempty"`

	// Uptime
	// uptime
	Uptime *string `json:"uptime,omitempty"`
}

type WSGProfileVdpProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileVdpProfile
}

func newWSGProfileVdpProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileVdpProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileVdpProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileVdpProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileVdpProfile() *WSGProfileVdpProfile {
	m := new(WSGProfileVdpProfile)
	return m
}

// WSGProfileVlanPrecedenceItem
//
// Definition: profile_VlanPrecedenceItem
//
// Vlan precedence item
type WSGProfileVlanPrecedenceItem struct {
	// Name
	// Name of the Vlan precedence item
	// Constraints:
	//    - oneof:[AAA,DEVICE,WLAN]
	Name *string `json:"name,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

func NewWSGProfileVlanPrecedenceItem() *WSGProfileVlanPrecedenceItem {
	m := new(WSGProfileVlanPrecedenceItem)
	return m
}

// WSGProfileZoneAffinityProfileList
//
// Definition: profile_zoneAffinityProfileList
type WSGProfileZoneAffinityProfileList struct {
	List []*WSGProfileReturnZoneAffinityProfile `json:"list,omitempty"`
}

type WSGProfileZoneAffinityProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGProfileZoneAffinityProfileList
}

func newWSGProfileZoneAffinityProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGProfileZoneAffinityProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGProfileZoneAffinityProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGProfileZoneAffinityProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGProfileZoneAffinityProfileList() *WSGProfileZoneAffinityProfileList {
	m := new(WSGProfileZoneAffinityProfileList)
	return m
}
