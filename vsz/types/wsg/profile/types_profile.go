package profile

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AccountingProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.DescriptionTo128 `json:"description,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*AcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

func NewAccountingProfile() *AccountingProfile {
	accountingProfileType := new(AccountingProfile)
	return accountingProfileType
}

func NewDefaultAccountingProfile() *AccountingProfile {
	accountingProfileType := new(AccountingProfile)
	return accountingProfileType
}

type AccountingProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AccountingProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAccountingProfileList() *AccountingProfileList {
	accountingProfileListType := new(AccountingProfileList)
	return accountingProfileListType
}

func NewDefaultAccountingProfileList() *AccountingProfileList {
	accountingProfileListType := new(AccountingProfileList)
	return accountingProfileListType
}

// AcctServiceRealmMapping
//
// Accounting service per realm
type AcctServiceRealmMapping struct {
	// Id
	// Accounting service UUID
	Id *string `json:"id,omitempty"`

	// Name
	// Accounting service name
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`

	// ServiceType
	// Accounting service type, NA is NA-Request Rejected
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,CGF]
	ServiceType *string `json:"serviceType" validate:"required,oneof=NA RADIUS CGF"`
}

func NewAcctServiceRealmMapping() *AcctServiceRealmMapping {
	acctServiceRealmMappingType := new(AcctServiceRealmMapping)
	return acctServiceRealmMappingType
}

func NewDefaultAcctServiceRealmMapping() *AcctServiceRealmMapping {
	acctServiceRealmMappingType := new(AcctServiceRealmMapping)
	return acctServiceRealmMappingType
}

// AdvancedOptionContent
//
// advanced option Content
type AdvancedOptionContent struct {
	// DhcpOpt43Subcode
	// dhcpOpt43Subcode of the ipsec profile
	DhcpOpt43Subcode *float64 `json:"dhcpOpt43Subcode,omitempty"`

	// DpdDelay
	// dpdDelay of the ipsec profile
	DpdDelay *float64 `json:"dpdDelay,omitempty"`

	// EnforceNatt
	// enforceNatt Enable of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[Disabled,Enabled]
	EnforceNatt *string `json:"enforceNatt,omitempty" validate:"omitempty,oneof=Disabled Enabled"`

	// FailoverMode
	// mode of the failover
	// Constraints:
	//    - nullable
	//    - oneof:[Non_Revertive,Revertive]
	FailoverMode *string `json:"failoverMode,omitempty" validate:"omitempty,oneof=Non_Revertive Revertive"`

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
	//    - nullable
	//    - oneof:[Disabled,Enabled]
	IpcompEnable *string `json:"ipcompEnable,omitempty" validate:"omitempty,oneof=Disabled Enabled"`

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

func NewAdvancedOptionContent() *AdvancedOptionContent {
	advancedOptionContentType := new(AdvancedOptionContent)
	return advancedOptionContentType
}

func NewDefaultAdvancedOptionContent() *AdvancedOptionContent {
	advancedOptionContentType := new(AdvancedOptionContent)
	return advancedOptionContentType
}

type ApnRealm struct {
	// DefaultAPN
	// name of the apnForwardingPolicys.
	DefaultAPN *string `json:"defaultAPN,omitempty"`

	// Realm
	// name of the apnRealm.
	Realm *string `json:"realm,omitempty"`
}

func NewApnRealm() *ApnRealm {
	apnRealmType := new(ApnRealm)
	return apnRealmType
}

func NewDefaultApnRealm() *ApnRealm {
	apnRealmType := new(ApnRealm)
	return apnRealmType
}

type AuthenticationProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.DescriptionTo128 `json:"description,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*RealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewAuthenticationProfile() *AuthenticationProfile {
	authenticationProfileType := new(AuthenticationProfile)
	return authenticationProfileType
}

func NewDefaultAuthenticationProfile() *AuthenticationProfile {
	authenticationProfileType := new(AuthenticationProfile)
	return authenticationProfileType
}

type AuthenticationProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AuthenticationProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAuthenticationProfileList() *AuthenticationProfileList {
	authenticationProfileListType := new(AuthenticationProfileList)
	return authenticationProfileListType
}

func NewDefaultAuthenticationProfileList() *AuthenticationProfileList {
	authenticationProfileListType := new(AuthenticationProfileList)
	return authenticationProfileListType
}

type BaseServiceInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.BaseServiceInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewBaseServiceInfoList() *BaseServiceInfoList {
	baseServiceInfoListType := new(BaseServiceInfoList)
	return baseServiceInfoListType
}

func NewDefaultBaseServiceInfoList() *BaseServiceInfoList {
	baseServiceInfoListType := new(BaseServiceInfoList)
	return baseServiceInfoListType
}

type BlockClient struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *common.Mac `json:"mac" validate:"required"`

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

func NewBlockClient() *BlockClient {
	blockClientType := new(BlockClient)
	return blockClientType
}

func NewDefaultBlockClient() *BlockClient {
	blockClientType := new(BlockClient)
	return blockClientType
}

type BlockClientList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BlockClientListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewBlockClientList() *BlockClientList {
	blockClientListType := new(BlockClientList)
	return blockClientListType
}

func NewDefaultBlockClientList() *BlockClientList {
	blockClientListType := new(BlockClientList)
	return blockClientListType
}

type BlockClientListType struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

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

func NewBlockClientListType() *BlockClientListType {
	blockClientListTypeType := new(BlockClientListType)
	return blockClientListTypeType
}

func NewDefaultBlockClientListType() *BlockClientListType {
	blockClientListTypeType := new(BlockClientListType)
	return blockClientListTypeType
}

type BonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	// Constraints:
	//    - required
	BonjourFencingRuleList []*BonjourFencingRule `json:"bonjourFencingRuleList" validate:"required"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*BonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

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
	Name *common.NormalName `json:"name" validate:"required"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewBonjourFencingPolicy() *BonjourFencingPolicy {
	bonjourFencingPolicyType := new(BonjourFencingPolicy)
	return bonjourFencingPolicyType
}

func NewDefaultBonjourFencingPolicy() *BonjourFencingPolicy {
	bonjourFencingPolicyType := new(BonjourFencingPolicy)
	return bonjourFencingPolicyType
}

type BonjourFencingPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BonjourFencingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewBonjourFencingPolicyList() *BonjourFencingPolicyList {
	bonjourFencingPolicyListType := new(BonjourFencingPolicyList)
	return bonjourFencingPolicyListType
}

func NewDefaultBonjourFencingPolicyList() *BonjourFencingPolicyList {
	bonjourFencingPolicyListType := new(BonjourFencingPolicyList)
	return bonjourFencingPolicyListType
}

type BonjourFencingRule struct {
	ClosestAp *common.Mac `json:"closestAp,omitempty"`

	CustomServiceName *common.NormalName `json:"customServiceName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DeviceMacList
	// Specify the device list providing Bonjour Service
	DeviceMacList []*BonjourFencingRuleDeviceMac `json:"deviceMacList,omitempty"`

	// DeviceType
	// Name of the Bonjour Fencing Rule
	// Constraints:
	//    - required
	//    - oneof:[Wired,Wireless]
	DeviceType *string `json:"deviceType" validate:"required,oneof=Wired Wireless"`

	// FencingRange
	// The range of AP can take Bonjour work
	// Constraints:
	//    - required
	//    - oneof:[SameAp,OneHopAp]
	FencingRange *string `json:"fencingRange" validate:"required,oneof=SameAp OneHopAp"`

	// ServiceType
	// Constraints:
	//    - required
	ServiceType *BridgeService `json:"serviceType" validate:"required"`
}

func NewBonjourFencingRule() *BonjourFencingRule {
	bonjourFencingRuleType := new(BonjourFencingRule)
	return bonjourFencingRuleType
}

func NewDefaultBonjourFencingRule() *BonjourFencingRule {
	bonjourFencingRuleType := new(BonjourFencingRule)
	return bonjourFencingRuleType
}

type BonjourFencingRuleDeviceMac struct {
	Mac *common.Mac `json:"mac,omitempty"`
}

func NewBonjourFencingRuleDeviceMac() *BonjourFencingRuleDeviceMac {
	bonjourFencingRuleDeviceMacType := new(BonjourFencingRuleDeviceMac)
	return bonjourFencingRuleDeviceMacType
}

func NewDefaultBonjourFencingRuleDeviceMac() *BonjourFencingRuleDeviceMac {
	bonjourFencingRuleDeviceMacType := new(BonjourFencingRuleDeviceMac)
	return bonjourFencingRuleDeviceMacType
}

type BonjourFencingRuleMapping struct {
	CustomServiceName *common.NormalName `json:"customServiceName,omitempty"`

	// CustomStringList
	// The array of mdns string
	CustomStringList []string `json:"customStringList,omitempty"`

	ServiceType *BridgeService `json:"serviceType,omitempty"`
}

func NewBonjourFencingRuleMapping() *BonjourFencingRuleMapping {
	bonjourFencingRuleMappingType := new(BonjourFencingRuleMapping)
	return bonjourFencingRuleMappingType
}

func NewDefaultBonjourFencingRuleMapping() *BonjourFencingRuleMapping {
	bonjourFencingRuleMappingType := new(BonjourFencingRuleMapping)
	return bonjourFencingRuleMappingType
}

type BonjourFencingService struct {
	NeighborApMac *string `json:"neighborApMac,omitempty"`

	NeighborApName *string `json:"neighborApName,omitempty"`

	ServiceType *BridgeService `json:"serviceType,omitempty"`

	// SourceType
	// Constraints:
	//    - nullable
	//    - oneof:[UNKNOWN,DIRECT,NEIGHBOR]
	SourceType *string `json:"sourceType,omitempty" validate:"omitempty,oneof=UNKNOWN DIRECT NEIGHBOR"`
}

func NewBonjourFencingService() *BonjourFencingService {
	bonjourFencingServiceType := new(BonjourFencingService)
	return bonjourFencingServiceType
}

func NewDefaultBonjourFencingService() *BonjourFencingService {
	bonjourFencingServiceType := new(BonjourFencingService)
	return bonjourFencingServiceType
}

type BonjourFencingStatistic struct {
	ApMac *string `json:"apMac,omitempty"`

	DroppedPacketsDueToNeighbor *int `json:"droppedPacketsDueToNeighbor,omitempty"`

	DroppedPacketsDueToServiceType *int `json:"droppedPacketsDueToServiceType,omitempty"`

	ForwardedPackets *int `json:"forwardedPackets,omitempty"`

	ServiceList []*BonjourFencingService `json:"serviceList,omitempty"`
}

func NewBonjourFencingStatistic() *BonjourFencingStatistic {
	bonjourFencingStatisticType := new(BonjourFencingStatistic)
	return bonjourFencingStatisticType
}

func NewDefaultBonjourFencingStatistic() *BonjourFencingStatistic {
	bonjourFencingStatisticType := new(BonjourFencingStatistic)
	return bonjourFencingStatisticType
}

type BridgeProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`
}

func NewBridgeProfile() *BridgeProfile {
	bridgeProfileType := new(BridgeProfile)
	return bridgeProfileType
}

func NewDefaultBridgeProfile() *BridgeProfile {
	bridgeProfileType := new(BridgeProfile)
	return bridgeProfileType
}

type BridgeProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BridgeProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewBridgeProfileList() *BridgeProfileList {
	bridgeProfileListType := new(BridgeProfileList)
	return bridgeProfileListType
}

func NewDefaultBridgeProfileList() *BridgeProfileList {
	bridgeProfileListType := new(BridgeProfileList)
	return bridgeProfileListType
}

type BridgeService string

func NewBridgeService() *BridgeService {
	bridgeServiceType := new(BridgeService)
	return bridgeServiceType
}

func NewDefaultBridgeService() *BridgeService {
	bridgeServiceType := new(BridgeService)
	return bridgeServiceType
}

type BulkBlockClient struct {
	BlockClientList []*BulkBlockClientBlockClientListType `json:"blockClientList,omitempty"`

	Description *common.Description `json:"description,omitempty"`
}

func NewBulkBlockClient() *BulkBlockClient {
	bulkBlockClientType := new(BulkBlockClient)
	return bulkBlockClientType
}

func NewDefaultBulkBlockClient() *BulkBlockClient {
	bulkBlockClientType := new(BulkBlockClient)
	return bulkBlockClientType
}

type BulkBlockClientBlockClientListType struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *common.Mac `json:"apMac" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *common.Mac `json:"mac" validate:"required"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewBulkBlockClientBlockClientListType() *BulkBlockClientBlockClientListType {
	bulkBlockClientBlockClientListTypeType := new(BulkBlockClientBlockClientListType)
	return bulkBlockClientBlockClientListTypeType
}

func NewDefaultBulkBlockClientBlockClientListType() *BulkBlockClientBlockClientListType {
	bulkBlockClientBlockClientListTypeType := new(BulkBlockClientBlockClientListType)
	return bulkBlockClientBlockClientListTypeType
}

type ClientIsolationEntry struct {
	Description *common.Description `json:"description,omitempty"`

	// Ip
	// Client Entry ip
	Ip *string `json:"ip,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *common.Mac `json:"mac" validate:"required"`
}

func NewClientIsolationEntry() *ClientIsolationEntry {
	clientIsolationEntryType := new(ClientIsolationEntry)
	return clientIsolationEntryType
}

func NewDefaultClientIsolationEntry() *ClientIsolationEntry {
	clientIsolationEntryType := new(ClientIsolationEntry)
	return clientIsolationEntryType
}

type ClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	// Constraints:
	//    - required
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled" validate:"required"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

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
	Name *common.NormalName `json:"name" validate:"required"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - required
	Whitelist []*ClientIsolationEntry `json:"whitelist" validate:"required"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewClientIsolationWhitelist() *ClientIsolationWhitelist {
	clientIsolationWhitelistType := new(ClientIsolationWhitelist)
	return clientIsolationWhitelistType
}

func NewDefaultClientIsolationWhitelist() *ClientIsolationWhitelist {
	clientIsolationWhitelistType := new(ClientIsolationWhitelist)
	return clientIsolationWhitelistType
}

type ClientIsolationWhitelistArray struct {
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

	List []*ClientIsolationWhitelist `json:"list,omitempty"`

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

func NewClientIsolationWhitelistArray() *ClientIsolationWhitelistArray {
	clientIsolationWhitelistArrayType := new(ClientIsolationWhitelistArray)
	return clientIsolationWhitelistArrayType
}

func NewDefaultClientIsolationWhitelistArray() *ClientIsolationWhitelistArray {
	clientIsolationWhitelistArrayType := new(ClientIsolationWhitelistArray)
	return clientIsolationWhitelistArrayType
}

// CmProtocolOptionContent
//
// Certificate Management Protocol Option
type CmProtocolOptionContent struct {
	// CmpDhcpOpt43Subcode
	// Certificate Management Protocol dhcpOpt43Subcode
	// Constraints:
	//    - required
	CmpDhcpOpt43Subcode *float64 `json:"cmpDhcpOpt43Subcode" validate:"required"`

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

func NewCmProtocolOptionContent() *CmProtocolOptionContent {
	cmProtocolOptionContentType := new(CmProtocolOptionContent)
	return cmProtocolOptionContentType
}

func NewDefaultCmProtocolOptionContent() *CmProtocolOptionContent {
	cmProtocolOptionContentType := new(CmProtocolOptionContent)
	return cmProtocolOptionContentType
}

type CoreNetworkGateway struct {
	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:1
	//    - max:32767
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty" validate:"omitempty,gte=1,lte=32767"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - nullable
	//    - default:3
	//    - min:1
	//    - max:255
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty" validate:"omitempty,gte=1,lte=255"`

	// PrimaryGateway
	// Primary Gateway
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary Gateway
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMTU
	// Gateway path MTU
	// Constraints:
	//    - nullable
	//    - oneof:[AUTO,MANUAL]
	TunnelMTU *string `json:"tunnelMTU,omitempty" validate:"omitempty,oneof=AUTO MANUAL"`

	// TunnelMTUSize
	// Manual setting value of Gateway path MTU
	// Constraints:
	//    - nullable
	//    - default:1500
	//    - min:850
	//    - max:1500
	TunnelMTUSize *int `json:"tunnelMTUSize,omitempty" validate:"omitempty,gte=850,lte=1500"`
}

func NewCoreNetworkGateway() *CoreNetworkGateway {
	coreNetworkGatewayType := new(CoreNetworkGateway)
	return coreNetworkGatewayType
}

func NewDefaultCoreNetworkGateway() *CoreNetworkGateway {
	coreNetworkGatewayType := new(CoreNetworkGateway)
	keepAlivePeriodField := 10
	coreNetworkGatewayType.KeepAlivePeriod = &keepAlivePeriodField
	keepAliveRetryField := 3
	coreNetworkGatewayType.KeepAliveRetry = &keepAliveRetryField
	tunnelMTUSizeField := 1500
	coreNetworkGatewayType.TunnelMTUSize = &tunnelMTUSizeField
	return coreNetworkGatewayType
}

type CreateAccountingProfile struct {
	Description *common.DescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*AcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

func NewCreateAccountingProfile() *CreateAccountingProfile {
	createAccountingProfileType := new(CreateAccountingProfile)
	return createAccountingProfileType
}

func NewDefaultCreateAccountingProfile() *CreateAccountingProfile {
	createAccountingProfileType := new(CreateAccountingProfile)
	return createAccountingProfileType
}

type CreateAuthenticationProfile struct {
	Description *common.DescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	// Constraints:
	//    - required
	//    - default:false
	GppSuppportEnabled *bool `json:"gppSuppportEnabled" validate:"required"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*RealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewCreateAuthenticationProfile() *CreateAuthenticationProfile {
	createAuthenticationProfileType := new(CreateAuthenticationProfile)
	return createAuthenticationProfileType
}

func NewDefaultCreateAuthenticationProfile() *CreateAuthenticationProfile {
	createAuthenticationProfileType := new(CreateAuthenticationProfile)
	gppSuppportEnabledField := false
	createAuthenticationProfileType.GppSuppportEnabled = &gppSuppportEnabledField
	return createAuthenticationProfileType
}

type CreateBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	// Constraints:
	//    - required
	BonjourFencingRuleList []*BonjourFencingRule `json:"bonjourFencingRuleList" validate:"required"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*BonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewCreateBonjourFencingPolicy() *CreateBonjourFencingPolicy {
	createBonjourFencingPolicyType := new(CreateBonjourFencingPolicy)
	return createBonjourFencingPolicyType
}

func NewDefaultCreateBonjourFencingPolicy() *CreateBonjourFencingPolicy {
	createBonjourFencingPolicyType := new(CreateBonjourFencingPolicy)
	return createBonjourFencingPolicyType
}

type CreateBridgeProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewCreateBridgeProfile() *CreateBridgeProfile {
	createBridgeProfileType := new(CreateBridgeProfile)
	return createBridgeProfileType
}

func NewDefaultCreateBridgeProfile() *CreateBridgeProfile {
	createBridgeProfileType := new(CreateBridgeProfile)
	return createBridgeProfileType
}

type CreateClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	// Constraints:
	//    - required
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - required
	Whitelist []*ClientIsolationEntry `json:"whitelist" validate:"required"`
}

func NewCreateClientIsolationWhitelist() *CreateClientIsolationWhitelist {
	createClientIsolationWhitelistType := new(CreateClientIsolationWhitelist)
	return createClientIsolationWhitelistType
}

func NewDefaultCreateClientIsolationWhitelist() *CreateClientIsolationWhitelist {
	createClientIsolationWhitelistType := new(CreateClientIsolationWhitelist)
	return createClientIsolationWhitelistType
}

type CreateDhcpProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// LeaseTimeHours
	// Lease time in hours of the DHCP Profile
	// Constraints:
	//    - required
	//    - min:0
	//    - max:24
	LeaseTimeHours *int `json:"leaseTimeHours" validate:"required,gte=0,lte=24"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	// Constraints:
	//    - required
	//    - min:0
	//    - max:59
	LeaseTimeMinutes *int `json:"leaseTimeMinutes" validate:"required,gte=0,lte=59"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PoolEndIp
	// Constraints:
	//    - required
	PoolEndIp *common.IpAddress `json:"poolEndIp" validate:"required"`

	// PoolStartIp
	// Constraints:
	//    - required
	PoolStartIp *common.IpAddress `json:"poolStartIp" validate:"required"`

	PrimaryDnsIp *common.IpAddress `json:"primaryDnsIp,omitempty"`

	SecondaryDnsIp *common.IpAddress `json:"secondaryDnsIp,omitempty"`

	// SubnetMask
	// Constraints:
	//    - required
	SubnetMask *common.IpAddress `json:"subnetMask" validate:"required"`

	// SubnetNetworkIp
	// Constraints:
	//    - required
	SubnetNetworkIp *common.IpAddress `json:"subnetNetworkIp" validate:"required"`

	// VlanId
	// VLAN ID of the DHCP Profile
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	VlanId *int `json:"vlanId" validate:"required,gte=1,lte=4094"`
}

func NewCreateDhcpProfile() *CreateDhcpProfile {
	createDhcpProfileType := new(CreateDhcpProfile)
	return createDhcpProfileType
}

func NewDefaultCreateDhcpProfile() *CreateDhcpProfile {
	createDhcpProfileType := new(CreateDhcpProfile)
	return createDhcpProfileType
}

type CreateDnsServerProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PrimaryIp
	// Primary ip of DNS server service
	// Constraints:
	//    - required
	PrimaryIp *string `json:"primaryIp" validate:"required"`

	// SecondaryIp
	// Secondary ip of DNS server service
	SecondaryIp *string `json:"secondaryIp,omitempty"`

	// TertiaryIp
	// Tertiary ip of DNS server service
	TertiaryIp *string `json:"tertiaryIp,omitempty"`
}

func NewCreateDnsServerProfile() *CreateDnsServerProfile {
	createDnsServerProfileType := new(CreateDnsServerProfile)
	return createDnsServerProfileType
}

func NewDefaultCreateDnsServerProfile() *CreateDnsServerProfile {
	createDnsServerProfileType := new(CreateDnsServerProfile)
	return createDnsServerProfileType
}

type CreateIpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"omitempty,oneof=PresharedKey Certificate"`

	CmProtocolOption *CmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	// Constraints:
	//    - required
	EspRekeyTime *float64 `json:"espRekeyTime" validate:"required"`

	EspRekeyTimeUnit *common.TimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *EspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	// Constraints:
	//    - required
	IkeRekeyTime *float64 `json:"ikeRekeyTime" validate:"required"`

	IkeRekeyTimeUnit *common.TimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *IkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	// IpMode
	// Constraints:
	//    - required
	IpMode *IpMode `json:"ipMode" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`

	// TunnelMode
	// Tunnel mode of IPsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[SOFT_GRE,RUCKUS_GRE]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"omitempty,oneof=SOFT_GRE RUCKUS_GRE"`
}

func NewCreateIpsecProfile() *CreateIpsecProfile {
	createIpsecProfileType := new(CreateIpsecProfile)
	return createIpsecProfileType
}

func NewDefaultCreateIpsecProfile() *CreateIpsecProfile {
	createIpsecProfileType := new(CreateIpsecProfile)
	return createIpsecProfileType
}

type CreateL2oGREProfile struct {
	// CoreNetworkGateway
	// Constraints:
	//    - required
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewCreateL2oGREProfile() *CreateL2oGREProfile {
	createL2oGREProfileType := new(CreateL2oGREProfile)
	return createL2oGREProfileType
}

func NewDefaultCreateL2oGREProfile() *CreateL2oGREProfile {
	createL2oGREProfileType := new(CreateL2oGREProfile)
	return createL2oGREProfileType
}

type CreatePrecedenceProfile struct {
	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*RateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*VlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

func NewCreatePrecedenceProfile() *CreatePrecedenceProfile {
	createPrecedenceProfileType := new(CreatePrecedenceProfile)
	return createPrecedenceProfileType
}

func NewDefaultCreatePrecedenceProfile() *CreatePrecedenceProfile {
	createPrecedenceProfileType := new(CreatePrecedenceProfile)
	return createPrecedenceProfileType
}

type CreateResultList []*common.CreateResult

func NewCreateResultList() *CreateResultList {
	createResultListType := make(CreateResultList, 0)
	return &createResultListType
}

func NewDefaultCreateResultList() *CreateResultList {
	createResultListType := make(CreateResultList, 0)
	return &createResultListType
}

type CreateRogueApPolicy struct {
	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Rules
	// Constraints:
	//    - required
	Rules []*RogueApRuleList `json:"rules" validate:"required"`
}

func NewCreateRogueApPolicy() *CreateRogueApPolicy {
	createRogueApPolicyType := new(CreateRogueApPolicy)
	return createRogueApPolicyType
}

func NewDefaultCreateRogueApPolicy() *CreateRogueApPolicy {
	createRogueApPolicyType := new(CreateRogueApPolicy)
	return createRogueApPolicyType
}

type CreateRtlsProfile struct {
	// EkahauEnabled
	// Eekahau Location Service Enabled
	// Constraints:
	//    - required
	EkahauEnabled *bool `json:"ekahauEnabled" validate:"required"`

	EkahauIp *common.IpAddress `json:"ekahauIp,omitempty"`

	// EkahauPort
	// Eekahau Location Server Port
	EkahauPort *int `json:"ekahauPort,omitempty"`

	Id *string `json:"id,omitempty"`

	// StanleyEnabled
	// Stanley Location Service Enabled
	// Constraints:
	//    - required
	StanleyEnabled *bool `json:"stanleyEnabled" validate:"required"`
}

func NewCreateRtlsProfile() *CreateRtlsProfile {
	createRtlsProfileType := new(CreateRtlsProfile)
	return createRtlsProfileType
}

func NewDefaultCreateRtlsProfile() *CreateRtlsProfile {
	createRtlsProfileType := new(CreateRtlsProfile)
	return createRtlsProfileType
}

type CreateRuckusGREProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// TunnelEncryption
	// Tunnel Encryption of the RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,AES128,AES256]
	TunnelEncryption *string `json:"tunnelEncryption,omitempty" validate:"omitempty,oneof=DISABLE AES128 AES256"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"omitempty,oneof=GRE GREUDP"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - required
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled" validate:"required,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"omitempty,gte=850,lte=9018"`
}

func NewCreateRuckusGREProfile() *CreateRuckusGREProfile {
	createRuckusGREProfileType := new(CreateRuckusGREProfile)
	return createRuckusGREProfileType
}

func NewDefaultCreateRuckusGREProfile() *CreateRuckusGREProfile {
	createRuckusGREProfileType := new(CreateRuckusGREProfile)
	tunnelMtuSizeField := 1500
	createRuckusGREProfileType.TunnelMtuSize = &tunnelMtuSizeField
	return createRuckusGREProfileType
}

type CreateSoftGREProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	IpMode *IpMode `json:"ipMode,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	// Constraints:
	//    - required
	//    - min:1
	//    - max:180
	KeepAlivePeriod *int `json:"keepAlivePeriod" validate:"required,gte=1,lte=180"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - required
	//    - min:2
	//    - max:20
	KeepAliveRetry *int `json:"keepAliveRetry" validate:"required,gte=2,lte=20"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	// Constraints:
	//    - required
	PrimaryGateway *string `json:"primaryGateway" validate:"required"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	// Constraints:
	//    - required
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled" validate:"required,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile. IPV4:850-1500, IPV6:1384-1500. Default 1500.
	// Constraints:
	//    - nullable
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"omitempty,gte=850,lte=9018"`
}

func NewCreateSoftGREProfile() *CreateSoftGREProfile {
	createSoftGREProfileType := new(CreateSoftGREProfile)
	return createSoftGREProfileType
}

func NewDefaultCreateSoftGREProfile() *CreateSoftGREProfile {
	createSoftGREProfileType := new(CreateSoftGREProfile)
	tunnelMtuSizeField := 1500
	createSoftGREProfileType.TunnelMtuSize = &tunnelMtuSizeField
	return createSoftGREProfileType
}

type CreateTrafficClassProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName2to64 `json:"name" validate:"required"`

	// TrafficClasses
	// Constraints:
	//    - required
	TrafficClasses []*common.TrafficClassRef `json:"trafficClasses" validate:"required"`
}

func NewCreateTrafficClassProfile() *CreateTrafficClassProfile {
	createTrafficClassProfileType := new(CreateTrafficClassProfile)
	return createTrafficClassProfileType
}

func NewDefaultCreateTrafficClassProfile() *CreateTrafficClassProfile {
	createTrafficClassProfileType := new(CreateTrafficClassProfile)
	return createTrafficClassProfileType
}

type CreateTtgpdgProfile struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	// Constraints:
	//    - required
	ApnForwardingRealms []*TtgpdgApnForwardingRealm `json:"apnForwardingRealms" validate:"required"`

	// ApnRealms
	// List of the Default APN
	ApnRealms []*ApnRealm `json:"apnRealms,omitempty"`

	// CommonSetting
	// Constraints:
	//    - required
	CommonSetting *TtgpdgCommonSetting `json:"commonSetting" validate:"required"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	// Constraints:
	//    - required
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN" validate:"required"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	// Constraints:
	//    - required
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewCreateTtgpdgProfile() *CreateTtgpdgProfile {
	createTtgpdgProfileType := new(CreateTtgpdgProfile)
	return createTtgpdgProfileType
}

func NewDefaultCreateTtgpdgProfile() *CreateTtgpdgProfile {
	createTtgpdgProfileType := new(CreateTtgpdgProfile)
	return createTtgpdgProfileType
}

type CreateUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=BLOCK ALLOW"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// IpAclRules
	// Traffic access control list
	IpAclRules []*IpAclRules `json:"ipAclRules,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// QmAppPolicyId
	// Application Policy UUID
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *UplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewCreateUserTrafficProfile() *CreateUserTrafficProfile {
	createUserTrafficProfileType := new(CreateUserTrafficProfile)
	return createUserTrafficProfileType
}

func NewDefaultCreateUserTrafficProfile() *CreateUserTrafficProfile {
	createUserTrafficProfileType := new(CreateUserTrafficProfile)
	defaultActionField := `ALLOW`
	createUserTrafficProfileType.DefaultAction = &defaultActionField
	return createUserTrafficProfileType
}

type CreateZoneAffinityProfile struct {
	// Description
	// The description of the profile
	// Constraints:
	//    - nullable
	//    - max:64
	Description *string `json:"description,omitempty" validate:"omitempty,max=64"`

	// Name
	// Zone affinity profile name
	// Constraints:
	//    - required
	//    - max:64
	//    - min:1
	Name *string `json:"name" validate:"required,max=64,min=1"`

	// ZoneAffinityList
	// Constraints:
	//    - required
	ZoneAffinityList []string `json:"zoneAffinityList" validate:"required"`
}

func NewCreateZoneAffinityProfile() *CreateZoneAffinityProfile {
	createZoneAffinityProfileType := new(CreateZoneAffinityProfile)
	return createZoneAffinityProfileType
}

func NewDefaultCreateZoneAffinityProfile() *CreateZoneAffinityProfile {
	createZoneAffinityProfileType := new(CreateZoneAffinityProfile)
	return createZoneAffinityProfileType
}

type DataPlaneL3RoamingData struct {
	// Activated
	// Show if this DP is included in the L3 roaming feature or not, 0 means excluded and 1 means included
	// Constraints:
	//    - required
	Activated *int `json:"activated" validate:"required"`

	// FirmwareVersion
	// DP firmware version
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Key
	// Data plane key
	// Constraints:
	//    - required
	Key *string `json:"key" validate:"required"`

	// Name
	// DP name
	Name *string `json:"name,omitempty"`

	// SubCriteriaType
	// Constraints:
	//    - nullable
	//    - oneof:[VLAN,SUBNET]
	SubCriteriaType *string `json:"subCriteriaType,omitempty" validate:"omitempty,oneof=VLAN SUBNET"`

	// Value
	// A list of L3 roaming configuration for this DP
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`
}

func NewDataPlaneL3RoamingData() *DataPlaneL3RoamingData {
	dataPlaneL3RoamingDataType := new(DataPlaneL3RoamingData)
	return dataPlaneL3RoamingDataType
}

func NewDefaultDataPlaneL3RoamingData() *DataPlaneL3RoamingData {
	dataPlaneL3RoamingDataType := new(DataPlaneL3RoamingData)
	return dataPlaneL3RoamingDataType
}

type DeleteBulkAccountingProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulkAccountingProfile() *DeleteBulkAccountingProfile {
	deleteBulkAccountingProfileType := new(DeleteBulkAccountingProfile)
	return deleteBulkAccountingProfileType
}

func NewDefaultDeleteBulkAccountingProfile() *DeleteBulkAccountingProfile {
	deleteBulkAccountingProfileType := new(DeleteBulkAccountingProfile)
	return deleteBulkAccountingProfileType
}

type DeleteBulkAuthenticationProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulkAuthenticationProfile() *DeleteBulkAuthenticationProfile {
	deleteBulkAuthenticationProfileType := new(DeleteBulkAuthenticationProfile)
	return deleteBulkAuthenticationProfileType
}

func NewDefaultDeleteBulkAuthenticationProfile() *DeleteBulkAuthenticationProfile {
	deleteBulkAuthenticationProfileType := new(DeleteBulkAuthenticationProfile)
	return deleteBulkAuthenticationProfileType
}

type DeleteBulkPrecedenceProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulkPrecedenceProfile() *DeleteBulkPrecedenceProfile {
	deleteBulkPrecedenceProfileType := new(DeleteBulkPrecedenceProfile)
	return deleteBulkPrecedenceProfileType
}

func NewDefaultDeleteBulkPrecedenceProfile() *DeleteBulkPrecedenceProfile {
	deleteBulkPrecedenceProfileType := new(DeleteBulkPrecedenceProfile)
	return deleteBulkPrecedenceProfileType
}

type DeleteBulkUserTrafficProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulkUserTrafficProfile() *DeleteBulkUserTrafficProfile {
	deleteBulkUserTrafficProfileType := new(DeleteBulkUserTrafficProfile)
	return deleteBulkUserTrafficProfileType
}

func NewDefaultDeleteBulkUserTrafficProfile() *DeleteBulkUserTrafficProfile {
	deleteBulkUserTrafficProfileType := new(DeleteBulkUserTrafficProfile)
	return deleteBulkUserTrafficProfileType
}

type DhcpOption82 struct {
	// DhcpOption82Enabled
	// Enable DHCP Option 82
	DhcpOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// Subopt1Enabled
	// Enable subopt-1
	Subopt1Enabled *bool `json:"subopt1Enabled,omitempty"`

	// Subopt1Format
	// Subopt-1 format
	// Constraints:
	//    - nullable
	//    - oneof:[AP_INFO,AP_MAC_hex,AP_MAC_hex_ESSID,AP_INFO_LOCATION]
	Subopt1Format *string `json:"subopt1Format,omitempty" validate:"omitempty,oneof=AP_INFO AP_MAC_hex AP_MAC_hex_ESSID AP_INFO_LOCATION"`

	// Subopt2Enabled
	// Enable subopt-2
	Subopt2Enabled *bool `json:"subopt2Enabled,omitempty"`

	// Subopt2Format
	// Subopt-2 format
	// Constraints:
	//    - nullable
	//    - oneof:[CLIENT_MAC_hex,CLIENT_MAC_hex_ESSID,AP_MAC_hex,AP_MAC__hex_ESSID]
	Subopt2Format *string `json:"subopt2Format,omitempty" validate:"omitempty,oneof=CLIENT_MAC_hex CLIENT_MAC_hex_ESSID AP_MAC_hex AP_MAC__hex_ESSID"`

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
	//    - nullable
	//    - oneof:[AREA_NAME,ESSID]
	Subopt151Format *string `json:"subopt151Format,omitempty" validate:"omitempty,oneof=AREA_NAME ESSID"`
}

func NewDhcpOption82() *DhcpOption82 {
	dhcpOption82Type := new(DhcpOption82)
	return dhcpOption82Type
}

func NewDefaultDhcpOption82() *DhcpOption82 {
	dhcpOption82Type := new(DhcpOption82)
	return dhcpOption82Type
}

type DhcpProfileList struct {
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

	List []*common.DhcpProfileRef `json:"list,omitempty"`

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

func NewDhcpProfileList() *DhcpProfileList {
	dhcpProfileListType := new(DhcpProfileList)
	return dhcpProfileListType
}

func NewDefaultDhcpProfileList() *DhcpProfileList {
	dhcpProfileListType := new(DhcpProfileList)
	return dhcpProfileListType
}

type DhcpRelayNoRelayTunnel struct {
	DhcpOption82 *DhcpOption82 `json:"dhcpOption82,omitempty"`

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

func NewDhcpRelayNoRelayTunnel() *DhcpRelayNoRelayTunnel {
	dhcpRelayNoRelayTunnelType := new(DhcpRelayNoRelayTunnel)
	return dhcpRelayNoRelayTunnelType
}

func NewDefaultDhcpRelayNoRelayTunnel() *DhcpRelayNoRelayTunnel {
	dhcpRelayNoRelayTunnelType := new(DhcpRelayNoRelayTunnel)
	return dhcpRelayNoRelayTunnelType
}

type DnsServerProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

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

func NewDnsServerProfile() *DnsServerProfile {
	dnsServerProfileType := new(DnsServerProfile)
	return dnsServerProfileType
}

func NewDefaultDnsServerProfile() *DnsServerProfile {
	dnsServerProfileType := new(DnsServerProfile)
	return dnsServerProfileType
}

type DnsServerProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DnsServerProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDnsServerProfileList() *DnsServerProfileList {
	dnsServerProfileListType := new(DnsServerProfileList)
	return dnsServerProfileListType
}

func NewDefaultDnsServerProfileList() *DnsServerProfileList {
	dnsServerProfileListType := new(DnsServerProfileList)
	return dnsServerProfileListType
}

type DownlinkRateLimiting struct {
	// DownlinkRateLimitingBps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	DownlinkRateLimitingBps *string `json:"downlinkRateLimitingBps,omitempty"`

	// DownlinkRateLimitingEnabled
	// Downlink rate limiting enabled or disabled
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`
}

func NewDownlinkRateLimiting() *DownlinkRateLimiting {
	downlinkRateLimitingType := new(DownlinkRateLimiting)
	return downlinkRateLimitingType
}

func NewDefaultDownlinkRateLimiting() *DownlinkRateLimiting {
	downlinkRateLimitingType := new(DownlinkRateLimiting)
	downlinkRateLimitingEnabledField := false
	downlinkRateLimitingType.DownlinkRateLimitingEnabled = &downlinkRateLimitingEnabledField
	return downlinkRateLimitingType
}

// EspProposal
//
// EspProposal based ipsec service mappings
type EspProposal struct {
	// AuthAlg
	// authAlg of espProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[MD5,SHA1,AESXCBC,SHA256,SHA384,SHA512]
	AuthAlg *string `json:"authAlg" validate:"required,oneof=MD5 SHA1 AESXCBC SHA256 SHA384 SHA512"`

	// DhGroup
	// dhGroup of espProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[None,Modp768,Modp1024,Modp1536,Modp2048,Modp3072,Modp4096,Modp6144,Modp8192,Ecp384]
	DhGroup *string `json:"dhGroup" validate:"required,oneof=None Modp768 Modp1024 Modp1536 Modp2048 Modp3072 Modp4096 Modp6144 Modp8192 Ecp384"`

	// EncAlg
	// encAlg of espProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[None,ThreeDES,AES128,AES192,AES256]
	EncAlg *string `json:"encAlg" validate:"required,oneof=None ThreeDES AES128 AES192 AES256"`
}

func NewEspProposal() *EspProposal {
	espProposalType := new(EspProposal)
	return espProposalType
}

func NewDefaultEspProposal() *EspProposal {
	espProposalType := new(EspProposal)
	return espProposalType
}

// EspSecurityAssociationContent
//
// espProposal Security Association Content
type EspSecurityAssociationContent struct {
	// EspProposals
	// espProposal list of the ipsec profile
	EspProposals []*EspProposal `json:"espProposals,omitempty"`

	// EspProposalType
	// espProposal Type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[Default,Specific]
	EspProposalType *string `json:"espProposalType,omitempty" validate:"omitempty,oneof=Default Specific"`
}

func NewEspSecurityAssociationContent() *EspSecurityAssociationContent {
	espSecurityAssociationContentType := new(EspSecurityAssociationContent)
	return espSecurityAssociationContentType
}

func NewDefaultEspSecurityAssociationContent() *EspSecurityAssociationContent {
	espSecurityAssociationContentType := new(EspSecurityAssociationContent)
	return espSecurityAssociationContentType
}

type FlexiVpnProfile struct {
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

func NewFlexiVpnProfile() *FlexiVpnProfile {
	flexiVpnProfileType := new(FlexiVpnProfile)
	return flexiVpnProfileType
}

func NewDefaultFlexiVpnProfile() *FlexiVpnProfile {
	flexiVpnProfileType := new(FlexiVpnProfile)
	return flexiVpnProfileType
}

type FlexiVpnProfileList struct {
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

	List []*FlexiVpnProfile `json:"list,omitempty"`

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

func NewFlexiVpnProfileList() *FlexiVpnProfileList {
	flexiVpnProfileListType := new(FlexiVpnProfileList)
	return flexiVpnProfileListType
}

func NewDefaultFlexiVpnProfileList() *FlexiVpnProfileList {
	flexiVpnProfileListType := new(FlexiVpnProfileList)
	return flexiVpnProfileListType
}

type GetL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*DataPlaneL3RoamingData `json:"dataPlanes,omitempty"`
}

func NewGetL3RoamingConfig() *GetL3RoamingConfig {
	getL3RoamingConfigType := new(GetL3RoamingConfig)
	return getL3RoamingConfigType
}

func NewDefaultGetL3RoamingConfig() *GetL3RoamingConfig {
	getL3RoamingConfigType := new(GetL3RoamingConfig)
	return getL3RoamingConfigType
}

type Hs20FriendlyName struct {
	// Language
	// Constraints:
	//    - required
	Language *common.LanguageName `json:"language" validate:"required"`

	// Name
	// Name of the friendly name
	// Constraints:
	//    - required
	//    - max:32
	//    - min:2
	Name *string `json:"name" validate:"required,max=32,min=2"`
}

func NewHs20FriendlyName() *Hs20FriendlyName {
	hs20FriendlyNameType := new(Hs20FriendlyName)
	return hs20FriendlyNameType
}

func NewDefaultHs20FriendlyName() *Hs20FriendlyName {
	hs20FriendlyNameType := new(Hs20FriendlyName)
	return hs20FriendlyNameType
}

type Hs20Operator struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	// Constraints:
	//    - required
	DomainNames []common.WildFQDN `json:"domainNames" validate:"required"`

	// FriendlyNames
	// Friendly names
	// Constraints:
	//    - required
	FriendlyNames []*Hs20FriendlyName `json:"friendlyNames" validate:"required"`

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
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewHs20Operator() *Hs20Operator {
	hs20OperatorType := new(Hs20Operator)
	return hs20OperatorType
}

func NewDefaultHs20Operator() *Hs20Operator {
	hs20OperatorType := new(Hs20Operator)
	return hs20OperatorType
}

type Hs20OperatorList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Hs20Operator `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewHs20OperatorList() *Hs20OperatorList {
	hs20OperatorListType := new(Hs20OperatorList)
	return hs20OperatorListType
}

func NewDefaultHs20OperatorList() *Hs20OperatorList {
	hs20OperatorListType := new(Hs20OperatorList)
	return hs20OperatorListType
}

type Hs20Provider struct {
	// Accountings
	// Accountings
	Accountings []*ProviderAccounting `json:"accountings,omitempty"`

	// Authentications
	// Authentications
	Authentications []*ProviderAuthentication `json:"authentications,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// HomeOis
	// Home OIs
	HomeOis []*ProviderHomeOIs `json:"homeOis,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	Osu *ProviderOnlineSignup `json:"osu,omitempty"`

	// Plmns
	// PLMNs
	Plmns []*ProviderPLMN `json:"plmns,omitempty"`

	// Realms
	// Realms
	Realms []*ProviderRealm `json:"realms,omitempty"`
}

func NewHs20Provider() *Hs20Provider {
	hs20ProviderType := new(Hs20Provider)
	return hs20ProviderType
}

func NewDefaultHs20Provider() *Hs20Provider {
	hs20ProviderType := new(Hs20Provider)
	return hs20ProviderType
}

type Hs20ProviderList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Hs20Provider `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewHs20ProviderList() *Hs20ProviderList {
	hs20ProviderListType := new(Hs20ProviderList)
	return hs20ProviderListType
}

func NewDefaultHs20ProviderList() *Hs20ProviderList {
	hs20ProviderListType := new(Hs20ProviderList)
	return hs20ProviderListType
}

// IkeProposal
//
// IkeProposal based ipsec service mappings
type IkeProposal struct {
	// AuthAlg
	// authAlg of ikeProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[MD5,SHA1,AESXCBC,SHA256,SHA384,SHA512]
	AuthAlg *string `json:"authAlg" validate:"required,oneof=MD5 SHA1 AESXCBC SHA256 SHA384 SHA512"`

	// DhGroup
	// dhGroup of ikeProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[Modp768,Modp1024,Modp1536,Modp2048,Modp3072,Modp4096,Modp6144,Modp8192,Ecp384]
	DhGroup *string `json:"dhGroup" validate:"required,oneof=Modp768 Modp1024 Modp1536 Modp2048 Modp3072 Modp4096 Modp6144 Modp8192 Ecp384"`

	// EncAlg
	// encAlg of ikeProposal Specific
	// Constraints:
	//    - required
	//    - oneof:[ThreeDES,AES128,AES192,AES256]
	EncAlg *string `json:"encAlg" validate:"required,oneof=ThreeDES AES128 AES192 AES256"`

	// PrfAlg
	// prfAlg of ikeProposal Specific
	// Constraints:
	//    - nullable
	//    - oneof:[UseIntegrityALG,PRF_MD5,PRF_SHA1,PRF_AES_CBC,PRF_AES_MAC,PRF_SHA256,PRF_SHA384,PRF_SHA512]
	PrfAlg *string `json:"prfAlg,omitempty" validate:"omitempty,oneof=UseIntegrityALG PRF_MD5 PRF_SHA1 PRF_AES_CBC PRF_AES_MAC PRF_SHA256 PRF_SHA384 PRF_SHA512"`
}

func NewIkeProposal() *IkeProposal {
	ikeProposalType := new(IkeProposal)
	return ikeProposalType
}

func NewDefaultIkeProposal() *IkeProposal {
	ikeProposalType := new(IkeProposal)
	return ikeProposalType
}

// IkeSecurityAssociationContent
//
// ikeProposal Security Association Content
type IkeSecurityAssociationContent struct {
	// IkeProposals
	// ikeProposal list of the ipsec profile
	IkeProposals []*IkeProposal `json:"ikeProposals,omitempty"`

	// IkeProposalType
	// ikeProposal Type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[Default,Specific]
	IkeProposalType *string `json:"ikeProposalType,omitempty" validate:"omitempty,oneof=Default Specific"`
}

func NewIkeSecurityAssociationContent() *IkeSecurityAssociationContent {
	ikeSecurityAssociationContentType := new(IkeSecurityAssociationContent)
	return ikeSecurityAssociationContentType
}

func NewDefaultIkeSecurityAssociationContent() *IkeSecurityAssociationContent {
	ikeSecurityAssociationContentType := new(IkeSecurityAssociationContent)
	return ikeSecurityAssociationContentType
}

type IpAclRules struct {
	// Action
	// The access of traffic access control.
	// Constraints:
	//    - nullable
	//    - default:'ALLOW'
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty" validate:"omitempty,oneof=ALLOW BLOCK"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty" validate:"omitempty,gte=1,lte=255"`

	Description *common.Description `json:"description,omitempty"`

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
	//    - nullable
	//    - default:'UPSTREAM'
	//    - oneof:[UPSTREAM]
	Direction *string `json:"direction,omitempty" validate:"omitempty,oneof=UPSTREAM"`

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
	//    - nullable
	//    - default:'IPv4'
	//    - oneof:[IPv4,IPv6]
	IpType *string `json:"ipType,omitempty" validate:"omitempty,oneof=IPv4 IPv6"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP UDPLITE ICMP_ICMPV4 ICMPV6 IGMP ESP AH SCTP CUSTOM"`

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

func NewIpAclRules() *IpAclRules {
	ipAclRulesType := new(IpAclRules)
	return ipAclRulesType
}

func NewDefaultIpAclRules() *IpAclRules {
	ipAclRulesType := new(IpAclRules)
	actionField := `ALLOW`
	ipAclRulesType.Action = &actionField
	directionField := `UPSTREAM`
	ipAclRulesType.Direction = &directionField
	ipTypeField := `IPv4`
	ipAclRulesType.IpType = &ipTypeField
	return ipAclRulesType
}

type IpMode string

func NewIpMode() *IpMode {
	ipModeType := new(IpMode)
	return ipModeType
}

func NewDefaultIpMode() *IpMode {
	ipModeType := new(IpMode)
	return ipModeType
}

type IpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"omitempty,oneof=PresharedKey Certificate"`

	CmProtocolOption *CmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	EspRekeyTimeUnit *common.TimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *EspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	IkeRekeyTimeUnit *common.TimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *IkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	IpMode *IpMode `json:"ipMode,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`

	// TunnelMode
	// Tunnel mode of IPsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[SOFT_GRE,RUCKUS_GRE]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"omitempty,oneof=SOFT_GRE RUCKUS_GRE"`
}

func NewIpsecProfile() *IpsecProfile {
	ipsecProfileType := new(IpsecProfile)
	return ipsecProfileType
}

func NewDefaultIpsecProfile() *IpsecProfile {
	ipsecProfileType := new(IpsecProfile)
	return ipsecProfileType
}

type IpsecProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*IpsecProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewIpsecProfileList() *IpsecProfileList {
	ipsecProfileListType := new(IpsecProfileList)
	return ipsecProfileListType
}

func NewDefaultIpsecProfileList() *IpsecProfileList {
	ipsecProfileListType := new(IpsecProfileList)
	return ipsecProfileListType
}

type L2oGREProfile struct {
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`
}

func NewL2oGREProfile() *L2oGREProfile {
	l2oGREProfileType := new(L2oGREProfile)
	return l2oGREProfileType
}

func NewDefaultL2oGREProfile() *L2oGREProfile {
	l2oGREProfileType := new(L2oGREProfile)
	return l2oGREProfileType
}

type L2oGREProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*L2oGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewL2oGREProfileList() *L2oGREProfileList {
	l2oGREProfileListType := new(L2oGREProfileList)
	return l2oGREProfileListType
}

func NewDefaultL2oGREProfileList() *L2oGREProfileList {
	l2oGREProfileListType := new(L2oGREProfileList)
	return l2oGREProfileListType
}

type LbsProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

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

func NewLbsProfile() *LbsProfile {
	lbsProfileType := new(LbsProfile)
	return lbsProfileType
}

func NewDefaultLbsProfile() *LbsProfile {
	lbsProfileType := new(LbsProfile)
	return lbsProfileType
}

type LbsProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LbsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewLbsProfileList() *LbsProfileList {
	lbsProfileListType := new(LbsProfileList)
	return lbsProfileListType
}

func NewDefaultLbsProfileList() *LbsProfileList {
	lbsProfileListType := new(LbsProfileList)
	return lbsProfileListType
}

type ModifyAccountingProfile struct {
	Description *common.DescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// ID of Accounting Profile
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*AcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

func NewModifyAccountingProfile() *ModifyAccountingProfile {
	modifyAccountingProfileType := new(ModifyAccountingProfile)
	return modifyAccountingProfileType
}

func NewDefaultModifyAccountingProfile() *ModifyAccountingProfile {
	modifyAccountingProfileType := new(ModifyAccountingProfile)
	return modifyAccountingProfileType
}

type ModifyAuthenticationProfile struct {
	Description *common.DescriptionTo128 `json:"description,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*RealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewModifyAuthenticationProfile() *ModifyAuthenticationProfile {
	modifyAuthenticationProfileType := new(ModifyAuthenticationProfile)
	return modifyAuthenticationProfileType
}

func NewDefaultModifyAuthenticationProfile() *ModifyAuthenticationProfile {
	modifyAuthenticationProfileType := new(ModifyAuthenticationProfile)
	gppSuppportEnabledField := false
	modifyAuthenticationProfileType.GppSuppportEnabled = &gppSuppportEnabledField
	return modifyAuthenticationProfileType
}

type ModifyBlockClient struct {
	Description *common.Description `json:"description,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`
}

func NewModifyBlockClient() *ModifyBlockClient {
	modifyBlockClientType := new(ModifyBlockClient)
	return modifyBlockClientType
}

func NewDefaultModifyBlockClient() *ModifyBlockClient {
	modifyBlockClientType := new(ModifyBlockClient)
	return modifyBlockClientType
}

type ModifyBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	BonjourFencingRuleList []*BonjourFencingRule `json:"bonjourFencingRuleList,omitempty"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*BonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewModifyBonjourFencingPolicy() *ModifyBonjourFencingPolicy {
	modifyBonjourFencingPolicyType := new(ModifyBonjourFencingPolicy)
	return modifyBonjourFencingPolicyType
}

func NewDefaultModifyBonjourFencingPolicy() *ModifyBonjourFencingPolicy {
	modifyBonjourFencingPolicyType := new(ModifyBonjourFencingPolicy)
	return modifyBonjourFencingPolicyType
}

type ModifyBridgeProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewModifyBridgeProfile() *ModifyBridgeProfile {
	modifyBridgeProfileType := new(ModifyBridgeProfile)
	return modifyBridgeProfileType
}

func NewDefaultModifyBridgeProfile() *ModifyBridgeProfile {
	modifyBridgeProfileType := new(ModifyBridgeProfile)
	return modifyBridgeProfileType
}

type ModifyClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	Whitelist []*ClientIsolationEntry `json:"whitelist,omitempty"`
}

func NewModifyClientIsolationWhitelist() *ModifyClientIsolationWhitelist {
	modifyClientIsolationWhitelistType := new(ModifyClientIsolationWhitelist)
	return modifyClientIsolationWhitelistType
}

func NewDefaultModifyClientIsolationWhitelist() *ModifyClientIsolationWhitelist {
	modifyClientIsolationWhitelistType := new(ModifyClientIsolationWhitelist)
	return modifyClientIsolationWhitelistType
}

type ModifyDnsServerProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

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

func NewModifyDnsServerProfile() *ModifyDnsServerProfile {
	modifyDnsServerProfileType := new(ModifyDnsServerProfile)
	return modifyDnsServerProfileType
}

func NewDefaultModifyDnsServerProfile() *ModifyDnsServerProfile {
	modifyDnsServerProfileType := new(ModifyDnsServerProfile)
	return modifyDnsServerProfileType
}

type ModifyHS20Operator struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	DomainNames []common.WildFQDN `json:"domainNames,omitempty"`

	// FriendlyNames
	// Friendly names
	FriendlyNames []*Hs20FriendlyName `json:"friendlyNames,omitempty"`

	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewModifyHS20Operator() *ModifyHS20Operator {
	modifyHS20OperatorType := new(ModifyHS20Operator)
	return modifyHS20OperatorType
}

func NewDefaultModifyHS20Operator() *ModifyHS20Operator {
	modifyHS20OperatorType := new(ModifyHS20Operator)
	return modifyHS20OperatorType
}

// ModifyIpAclRules
//
// Traffic access control list
type ModifyIpAclRules struct {
	// Action
	// The access of traffic access control.
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action" validate:"required,oneof=ALLOW BLOCK"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty" validate:"omitempty,gte=1,lte=255"`

	Description *common.Description `json:"description,omitempty"`

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
	Direction *string `json:"direction" validate:"required,oneof=UPSTREAM"`

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
	//    - nullable
	//    - default:'IPv4'
	//    - oneof:[IPv4,IPv6]
	IpType *string `json:"ipType,omitempty" validate:"omitempty,oneof=IPv4 IPv6"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP UDPLITE ICMP_ICMPV4 ICMPV6 IGMP ESP AH SCTP CUSTOM"`

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

func NewModifyIpAclRules() *ModifyIpAclRules {
	modifyIpAclRulesType := new(ModifyIpAclRules)
	return modifyIpAclRulesType
}

func NewDefaultModifyIpAclRules() *ModifyIpAclRules {
	modifyIpAclRulesType := new(ModifyIpAclRules)
	actionField := `ALLOW`
	modifyIpAclRulesType.Action = &actionField
	directionField := `UPSTREAM`
	modifyIpAclRulesType.Direction = &directionField
	ipTypeField := `IPv4`
	modifyIpAclRulesType.IpType = &ipTypeField
	return modifyIpAclRulesType
}

type ModifyIpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"omitempty,oneof=PresharedKey Certificate"`

	CmProtocolOption *CmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	EspRekeyTimeUnit *common.TimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *EspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	IkeRekeyTimeUnit *common.TimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *IkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	IpMode *IpMode `json:"ipMode,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`
}

func NewModifyIpsecProfile() *ModifyIpsecProfile {
	modifyIpsecProfileType := new(ModifyIpsecProfile)
	return modifyIpsecProfileType
}

func NewDefaultModifyIpsecProfile() *ModifyIpsecProfile {
	modifyIpsecProfileType := new(ModifyIpsecProfile)
	return modifyIpsecProfileType
}

type ModifyL2oGREProfile struct {
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewModifyL2oGREProfile() *ModifyL2oGREProfile {
	modifyL2oGREProfileType := new(ModifyL2oGREProfile)
	return modifyL2oGREProfileType
}

func NewDefaultModifyL2oGREProfile() *ModifyL2oGREProfile {
	modifyL2oGREProfileType := new(ModifyL2oGREProfile)
	return modifyL2oGREProfileType
}

type ModifyRuckusGREProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// TunnelEncryption
	// Tunnel Encryption of the RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,AES128,AES256]
	TunnelEncryption *string `json:"tunnelEncryption,omitempty" validate:"omitempty,oneof=DISABLE AES128 AES256"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"omitempty,oneof=GRE GREUDP"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"omitempty,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"omitempty,gte=850,lte=9018"`
}

func NewModifyRuckusGREProfile() *ModifyRuckusGREProfile {
	modifyRuckusGREProfileType := new(ModifyRuckusGREProfile)
	return modifyRuckusGREProfileType
}

func NewDefaultModifyRuckusGREProfile() *ModifyRuckusGREProfile {
	modifyRuckusGREProfileType := new(ModifyRuckusGREProfile)
	tunnelMtuSizeField := 1500
	modifyRuckusGREProfileType.TunnelMtuSize = &tunnelMtuSizeField
	return modifyRuckusGREProfileType
}

type ModifySoftGREProfile struct {
	Description *common.Description `json:"description,omitempty"`

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
	//    - nullable
	//    - default:10
	//    - min:1
	//    - max:180
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty" validate:"omitempty,gte=1,lte=180"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - nullable
	//    - default:5
	//    - min:2
	//    - max:20
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty" validate:"omitempty,gte=2,lte=20"`

	Name *common.NormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"omitempty,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile. IPV4:850-1500, IPV6:1384-1500. Default 1500.
	// Constraints:
	//    - nullable
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"omitempty,gte=850,lte=9018"`
}

func NewModifySoftGREProfile() *ModifySoftGREProfile {
	modifySoftGREProfileType := new(ModifySoftGREProfile)
	return modifySoftGREProfileType
}

func NewDefaultModifySoftGREProfile() *ModifySoftGREProfile {
	modifySoftGREProfileType := new(ModifySoftGREProfile)
	keepAlivePeriodField := 10
	modifySoftGREProfileType.KeepAlivePeriod = &keepAlivePeriodField
	keepAliveRetryField := 5
	modifySoftGREProfileType.KeepAliveRetry = &keepAliveRetryField
	tunnelMtuSizeField := 1500
	modifySoftGREProfileType.TunnelMtuSize = &tunnelMtuSizeField
	return modifySoftGREProfileType
}

type ModifyUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - nullable
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"omitempty,oneof=BLOCK ALLOW"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// Id
	// Identifier of the user traffic profile
	Id *string `json:"id,omitempty"`

	// IpAclRules
	// Traffic access control list
	IpAclRules []*ModifyIpAclRules `json:"ipAclRules,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// QmAppPolicyId
	// Application Policy UUID
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *UplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewModifyUserTrafficProfile() *ModifyUserTrafficProfile {
	modifyUserTrafficProfileType := new(ModifyUserTrafficProfile)
	return modifyUserTrafficProfileType
}

func NewDefaultModifyUserTrafficProfile() *ModifyUserTrafficProfile {
	modifyUserTrafficProfileType := new(ModifyUserTrafficProfile)
	defaultActionField := `ALLOW`
	modifyUserTrafficProfileType.DefaultAction = &defaultActionField
	return modifyUserTrafficProfileType
}

type ModifyZoneAffinityProfile struct {
	// Description
	// The description of the profile
	// Constraints:
	//    - nullable
	//    - max:64
	Description *string `json:"description,omitempty" validate:"omitempty,max=64"`

	// Name
	// Zone affinity profile name
	// Constraints:
	//    - nullable
	//    - max:64
	//    - min:1
	Name *string `json:"name,omitempty" validate:"omitempty,max=64,min=1"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
}

func NewModifyZoneAffinityProfile() *ModifyZoneAffinityProfile {
	modifyZoneAffinityProfileType := new(ModifyZoneAffinityProfile)
	return modifyZoneAffinityProfileType
}

func NewDefaultModifyZoneAffinityProfile() *ModifyZoneAffinityProfile {
	modifyZoneAffinityProfileType := new(ModifyZoneAffinityProfile)
	return modifyZoneAffinityProfileType
}

type PrecedenceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PrecedenceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewPrecedenceList() *PrecedenceList {
	precedenceListType := new(PrecedenceList)
	return precedenceListType
}

func NewDefaultPrecedenceList() *PrecedenceList {
	precedenceListType := new(PrecedenceList)
	return precedenceListType
}

type PrecedenceListType struct {
	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*RateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*VlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

func NewPrecedenceListType() *PrecedenceListType {
	precedenceListTypeType := new(PrecedenceListType)
	return precedenceListTypeType
}

func NewDefaultPrecedenceListType() *PrecedenceListType {
	precedenceListTypeType := new(PrecedenceListType)
	return precedenceListTypeType
}

type ProfileCloneRequest struct {
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

func NewProfileCloneRequest() *ProfileCloneRequest {
	profileCloneRequestType := new(ProfileCloneRequest)
	return profileCloneRequestType
}

func NewDefaultProfileCloneRequest() *ProfileCloneRequest {
	profileCloneRequestType := new(ProfileCloneRequest)
	return profileCloneRequestType
}

type ProfileCloneResponse struct {
	*ProfileCloneRequest
}

type ProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ProfileListType `json:"list,omitempty"`

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

func NewProfileList() *ProfileList {
	profileListType := new(ProfileList)
	return profileListType
}

func NewDefaultProfileList() *ProfileList {
	profileListType := new(ProfileList)
	return profileListType
}

type ProfileListType struct {
	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewProfileListType() *ProfileListType {
	profileListTypeType := new(ProfileListType)
	return profileListTypeType
}

func NewDefaultProfileListType() *ProfileListType {
	profileListTypeType := new(ProfileListType)
	return profileListTypeType
}

type ProviderAccounting struct {
	// Id
	// Accounting id
	Id *string `json:"id,omitempty"`

	// Name
	// Accounting name
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`

	// ServiceType
	// Accounting service type
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,CGF]
	ServiceType *string `json:"serviceType" validate:"required,oneof=NA RADIUS CGF"`
}

func NewProviderAccounting() *ProviderAccounting {
	providerAccountingType := new(ProviderAccounting)
	return providerAccountingType
}

func NewDefaultProviderAccounting() *ProviderAccounting {
	providerAccountingType := new(ProviderAccounting)
	return providerAccountingType
}

type ProviderAuthentication struct {
	// Id
	// Authentication id
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication name
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`

	// ServiceType
	// Authentication service type
	// Constraints:
	//    - required
	//    - oneof:[NA,LOCAL_DB,RADIUS,GUEST]
	ServiceType *string `json:"serviceType" validate:"required,oneof=NA LOCAL_DB RADIUS GUEST"`

	// VlanId
	// Dynamic vlan ID
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	VlanId *int `json:"vlanId,omitempty" validate:"omitempty,gte=1,lte=4094"`
}

func NewProviderAuthentication() *ProviderAuthentication {
	providerAuthenticationType := new(ProviderAuthentication)
	return providerAuthenticationType
}

func NewDefaultProviderAuthentication() *ProviderAuthentication {
	providerAuthenticationType := new(ProviderAuthentication)
	return providerAuthenticationType
}

type ProviderEAPAuthSetting struct {
	// Info
	// EAP auth info
	// Constraints:
	//    - required
	//    - oneof:[Expanded,Non,Inner,Expanded_Inner,Credential,Tunneled]
	Info *string `json:"info" validate:"required,oneof=Expanded Non Inner Expanded_Inner Credential Tunneled"`

	// Type
	// EAP auth type
	Type *string `json:"type,omitempty"`

	// VendorId
	// EAP auth vendor ID
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:16777215
	VendorId *int `json:"vendorId,omitempty" validate:"omitempty,gte=0,lte=16777215"`

	// VendorType
	// EAP auth vendor type
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:4294967295
	VendorType *int `json:"vendorType,omitempty" validate:"omitempty,gte=0,lte=4294967295"`
}

func NewProviderEAPAuthSetting() *ProviderEAPAuthSetting {
	providerEAPAuthSettingType := new(ProviderEAPAuthSetting)
	return providerEAPAuthSettingType
}

func NewDefaultProviderEAPAuthSetting() *ProviderEAPAuthSetting {
	providerEAPAuthSettingType := new(ProviderEAPAuthSetting)
	return providerEAPAuthSettingType
}

type ProviderEAPMethod struct {
	// AuthSettings
	// EAP method auth settings
	AuthSettings []*ProviderEAPAuthSetting `json:"authSettings,omitempty"`

	// Type
	// EAP method type
	// Constraints:
	//    - required
	//    - oneof:[NA,MD5,EAP_TLS,EAP_Cisco,EAP_SIM,EAP_TTLS,EAP_AKA,PEAP,EAP_MSCHAP_V2,EAP_AKAs,Reserved]
	Type *string `json:"type" validate:"required,oneof=NA MD5 EAP_TLS EAP_Cisco EAP_SIM EAP_TTLS EAP_AKA PEAP EAP_MSCHAP_V2 EAP_AKAs Reserved"`
}

func NewProviderEAPMethod() *ProviderEAPMethod {
	providerEAPMethodType := new(ProviderEAPMethod)
	return providerEAPMethodType
}

func NewDefaultProviderEAPMethod() *ProviderEAPMethod {
	providerEAPMethodType := new(ProviderEAPMethod)
	return providerEAPMethodType
}

type ProviderExternalOSU struct {
	// CommonLanguageIcon
	// The base64 encoded data of icon.
	// Constraints:
	//    - required
	CommonLanguageIcon *string `json:"commonLanguageIcon" validate:"required"`

	// OsuNaiRealm
	// Online signup NAI realm, it should be one of realm as defined in Hotspot 2.0 identity provider
	// Constraints:
	//    - required
	OsuNaiRealm *string `json:"osuNaiRealm" validate:"required"`

	// OsuServiceUrl
	// Constraints:
	//    - required
	OsuServiceUrl *common.HTTPS `json:"osuServiceUrl" validate:"required"`

	// ProvisioningProtocals
	// Provisioning protocal
	// Constraints:
	//    - required
	ProvisioningProtocals []ProviderProvisionProtocal `json:"provisioningProtocals" validate:"required"`

	// SubscriptionDescriptions
	// Subscription descriptions
	// Constraints:
	//    - required
	SubscriptionDescriptions []*ProviderSubscriptionDescription `json:"subscriptionDescriptions" validate:"required"`

	// WhitelistedDomains
	// Whitelisted domains
	WhitelistedDomains []common.WildFQDN `json:"whitelistedDomains,omitempty"`
}

func NewProviderExternalOSU() *ProviderExternalOSU {
	providerExternalOSUType := new(ProviderExternalOSU)
	return providerExternalOSUType
}

func NewDefaultProviderExternalOSU() *ProviderExternalOSU {
	providerExternalOSUType := new(ProviderExternalOSU)
	return providerExternalOSUType
}

type ProviderHomeOIs struct {
	// Name
	// Name of the home OI
	// Constraints:
	//    - required
	//    - max:255
	Name *string `json:"name" validate:"required,max=255"`

	// Oi
	// Orgnization ID(3Hex or 5Hex)
	// Constraints:
	//    - required
	Oi *string `json:"oi" validate:"required"`
}

func NewProviderHomeOIs() *ProviderHomeOIs {
	providerHomeOIsType := new(ProviderHomeOIs)
	return providerHomeOIsType
}

func NewDefaultProviderHomeOIs() *ProviderHomeOIs {
	providerHomeOIsType := new(ProviderHomeOIs)
	return providerHomeOIsType
}

type ProviderInternalOSU struct {
	// Certificate
	// Constraints:
	//    - required
	Certificate *common.GenericRef `json:"certificate" validate:"required"`

	// CommonLanguageIcon
	// The base64 encoded data of icon.
	// Constraints:
	//    - required
	CommonLanguageIcon *string `json:"commonLanguageIcon" validate:"required"`

	// OsuAuthServices
	// Online signup authentication services
	// Constraints:
	//    - required
	OsuAuthServices []*ProviderInternalOSUOsuAuthServicesType `json:"osuAuthServices" validate:"required"`

	// OsuPortal
	// Constraints:
	//    - required
	OsuPortal *ProviderInternalOSUOsuPortalType `json:"osuPortal" validate:"required"`

	// ProvisioningFormat
	// Provisioning format
	// Constraints:
	//    - required
	//    - oneof:[R1_R2_ZEROIT,R2_ZEROIT]
	ProvisioningFormat *string `json:"provisioningFormat" validate:"required,oneof=R1_R2_ZEROIT R2_ZEROIT"`

	// ProvisioningProtocals
	// Provisioning protocal
	// Constraints:
	//    - required
	ProvisioningProtocals []ProviderProvisionProtocal `json:"provisioningProtocals" validate:"required"`

	// ProvisioningUpdateType
	// Provisioning update at
	// Constraints:
	//    - required
	//    - oneof:[ALWAYS,KNOWN_ROAM_PARTNERS,NEVER]
	ProvisioningUpdateType *string `json:"provisioningUpdateType" validate:"required,oneof=ALWAYS KNOWN_ROAM_PARTNERS NEVER"`

	// SubscriptionDescriptions
	// Subscription descriptions
	// Constraints:
	//    - required
	SubscriptionDescriptions []*ProviderSubscriptionDescription `json:"subscriptionDescriptions" validate:"required"`

	// WhitelistedDomains
	// whitelisted domains
	WhitelistedDomains []common.WildFQDN `json:"whitelistedDomains,omitempty"`
}

func NewProviderInternalOSU() *ProviderInternalOSU {
	providerInternalOSUType := new(ProviderInternalOSU)
	return providerInternalOSUType
}

func NewDefaultProviderInternalOSU() *ProviderInternalOSU {
	providerInternalOSUType := new(ProviderInternalOSU)
	return providerInternalOSUType
}

type ProviderInternalOSUOsuAuthServicesType struct {
	// CredentialType
	// Authentication credential type
	// Constraints:
	//    - required
	//    - oneof:[LOCAL,REMOTE]
	CredentialType *string `json:"credentialType" validate:"required,oneof=LOCAL REMOTE"`

	// Expiration
	// Expiration hour. null mean never expire
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:175200
	Expiration *int `json:"expiration,omitempty" validate:"omitempty,gte=1,lte=175200"`

	// Id
	// Identifier of authentication service
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication service name
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`
}

func NewProviderInternalOSUOsuAuthServicesType() *ProviderInternalOSUOsuAuthServicesType {
	providerInternalOSUOsuAuthServicesTypeType := new(ProviderInternalOSUOsuAuthServicesType)
	return providerInternalOSUOsuAuthServicesTypeType
}

func NewDefaultProviderInternalOSUOsuAuthServicesType() *ProviderInternalOSUOsuAuthServicesType {
	providerInternalOSUOsuAuthServicesTypeType := new(ProviderInternalOSUOsuAuthServicesType)
	return providerInternalOSUOsuAuthServicesTypeType
}

type ProviderInternalOSUOsuPortalType struct {
	ExternalUrl *common.HTTPS `json:"externalUrl,omitempty"`

	InternalOSUPortal *common.GenericRef `json:"internalOSUPortal,omitempty"`

	// Type
	// Portal type
	// Constraints:
	//    - required
	//    - oneof:[Internal,External]
	Type *string `json:"type" validate:"required,oneof=Internal External"`
}

func NewProviderInternalOSUOsuPortalType() *ProviderInternalOSUOsuPortalType {
	providerInternalOSUOsuPortalTypeType := new(ProviderInternalOSUOsuPortalType)
	return providerInternalOSUOsuPortalTypeType
}

func NewDefaultProviderInternalOSUOsuPortalType() *ProviderInternalOSUOsuPortalType {
	providerInternalOSUOsuPortalTypeType := new(ProviderInternalOSUOsuPortalType)
	return providerInternalOSUOsuPortalTypeType
}

type ProviderOnlineSignup struct {
	ExternalOSU *ProviderExternalOSU `json:"externalOSU,omitempty"`

	InternalOSU *ProviderInternalOSU `json:"internalOSU,omitempty"`

	// Type
	// Online singup type
	// Constraints:
	//    - required
	//    - oneof:[Internal,External]
	Type *string `json:"type" validate:"required,oneof=Internal External"`
}

func NewProviderOnlineSignup() *ProviderOnlineSignup {
	providerOnlineSignupType := new(ProviderOnlineSignup)
	return providerOnlineSignupType
}

func NewDefaultProviderOnlineSignup() *ProviderOnlineSignup {
	providerOnlineSignupType := new(ProviderOnlineSignup)
	return providerOnlineSignupType
}

type ProviderPLMN struct {
	// Mcc
	// MCC
	// Constraints:
	//    - required
	Mcc *string `json:"mcc" validate:"required"`

	// Mnc
	// MNC
	// Constraints:
	//    - required
	Mnc *string `json:"mnc" validate:"required"`
}

func NewProviderPLMN() *ProviderPLMN {
	providerPLMNType := new(ProviderPLMN)
	return providerPLMNType
}

func NewDefaultProviderPLMN() *ProviderPLMN {
	providerPLMNType := new(ProviderPLMN)
	return providerPLMNType
}

type ProviderProvisionProtocal string

func NewProviderProvisionProtocal() *ProviderProvisionProtocal {
	providerProvisionProtocalType := new(ProviderProvisionProtocal)
	return providerProvisionProtocalType
}

func NewDefaultProviderProvisionProtocal() *ProviderProvisionProtocal {
	providerProvisionProtocalType := new(ProviderProvisionProtocal)
	return providerProvisionProtocalType
}

type ProviderRealm struct {
	// EapMethods
	// EAP methods
	// Constraints:
	//    - required
	EapMethods []*ProviderEAPMethod `json:"eapMethods" validate:"required"`

	// Encoding
	// Encoding
	// Constraints:
	//    - required
	//    - oneof:[RFC4282,UTF8]
	Encoding *string `json:"encoding" validate:"required,oneof=RFC4282 UTF8"`

	// Name
	// Name of realm
	// Constraints:
	//    - required
	//    - max:243
	//    - min:2
	Name *string `json:"name" validate:"required,max=243,min=2"`
}

func NewProviderRealm() *ProviderRealm {
	providerRealmType := new(ProviderRealm)
	return providerRealmType
}

func NewDefaultProviderRealm() *ProviderRealm {
	providerRealmType := new(ProviderRealm)
	return providerRealmType
}

type ProviderSubscriptionDescription struct {
	// Description
	// Description of the friendly name
	// Constraints:
	//    - nullable
	//    - max:64
	Description *string `json:"description,omitempty" validate:"omitempty,max=64"`

	// Icon
	// The binary data of icon, maximum size 65536
	Icon *string `json:"icon,omitempty"`

	// Language
	// Constraints:
	//    - required
	Language *common.LanguageName `json:"language" validate:"required"`

	// Name
	// Name of the friendly name
	// Constraints:
	//    - required
	//    - max:252
	//    - min:2
	Name *string `json:"name" validate:"required,max=252,min=2"`
}

func NewProviderSubscriptionDescription() *ProviderSubscriptionDescription {
	providerSubscriptionDescriptionType := new(ProviderSubscriptionDescription)
	return providerSubscriptionDescriptionType
}

func NewDefaultProviderSubscriptionDescription() *ProviderSubscriptionDescription {
	providerSubscriptionDescriptionType := new(ProviderSubscriptionDescription)
	return providerSubscriptionDescriptionType
}

// RateLimitingPrecedenceItem
//
// Rate limiting precedence item
type RateLimitingPrecedenceItem struct {
	// Name
	// Name of rate limiting precedence item
	// Constraints:
	//    - nullable
	//    - oneof:[AAA,DEVICE,WLANUTP]
	Name *string `json:"name,omitempty" validate:"omitempty,oneof=AAA DEVICE WLANUTP"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

func NewRateLimitingPrecedenceItem() *RateLimitingPrecedenceItem {
	rateLimitingPrecedenceItemType := new(RateLimitingPrecedenceItem)
	return rateLimitingPrecedenceItemType
}

func NewDefaultRateLimitingPrecedenceItem() *RateLimitingPrecedenceItem {
	rateLimitingPrecedenceItemType := new(RateLimitingPrecedenceItem)
	return rateLimitingPrecedenceItemType
}

// RealmAuthServiceMapping
//
// Realm based authentication service mappings
type RealmAuthServiceMapping struct {
	// AuthorizationMethod
	// Authorization method
	// Constraints:
	//    - required
	//    - oneof:[NonGPPCallFlow,GPPCallFlow,UpdateGPRSLocation,RestoreData,NoAutz]
	AuthorizationMethod *string `json:"authorizationMethod" validate:"required,oneof=NonGPPCallFlow GPPCallFlow UpdateGPRSLocation RestoreData NoAutz"`

	// DynamicVlanId
	// Dynamic VLAN ID
	// Constraints:
	//    - nullable
	//    - min:2
	//    - max:4094
	DynamicVlanId *int `json:"dynamicVlanId,omitempty" validate:"omitempty,gte=2,lte=4094"`

	HostedAaaEnabled *bool `json:"hostedAaaEnabled,omitempty"`

	// Id
	// Authentication service UUID
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication service name
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`

	// ServiceType
	// Authentication service type, NA is NA-Request Rejected
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,LOCAL_DB,HLR,AD,LDAP]
	ServiceType *string `json:"serviceType" validate:"required,oneof=NA RADIUS LOCAL_DB HLR AD LDAP"`
}

func NewRealmAuthServiceMapping() *RealmAuthServiceMapping {
	realmAuthServiceMappingType := new(RealmAuthServiceMapping)
	return realmAuthServiceMappingType
}

func NewDefaultRealmAuthServiceMapping() *RealmAuthServiceMapping {
	realmAuthServiceMappingType := new(RealmAuthServiceMapping)
	return realmAuthServiceMappingType
}

type ReturnZoneAffinityProfile struct {
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
	//    - nullable
	//    - max:64
	//    - min:1
	Name *string `json:"name,omitempty" validate:"omitempty,max=64,min=1"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`

	ZoneAffinityListWithPriority []*ReturnZoneAffinityProfileZoneAffinityListWithPriorityType `json:"zoneAffinityListWithPriority,omitempty"`
}

func NewReturnZoneAffinityProfile() *ReturnZoneAffinityProfile {
	returnZoneAffinityProfileType := new(ReturnZoneAffinityProfile)
	return returnZoneAffinityProfileType
}

func NewDefaultReturnZoneAffinityProfile() *ReturnZoneAffinityProfile {
	returnZoneAffinityProfileType := new(ReturnZoneAffinityProfile)
	return returnZoneAffinityProfileType
}

type ReturnZoneAffinityProfileZoneAffinityListWithPriorityType struct {
	// DpId
	// DP ID
	DpId *string `json:"dpId,omitempty"`

	// Priority
	// The priority of DP in zone affinity
	Priority *float64 `json:"priority,omitempty"`
}

func NewReturnZoneAffinityProfileZoneAffinityListWithPriorityType() *ReturnZoneAffinityProfileZoneAffinityListWithPriorityType {
	returnZoneAffinityProfileZoneAffinityListWithPriorityTypeType := new(ReturnZoneAffinityProfileZoneAffinityListWithPriorityType)
	return returnZoneAffinityProfileZoneAffinityListWithPriorityTypeType
}

func NewDefaultReturnZoneAffinityProfileZoneAffinityListWithPriorityType() *ReturnZoneAffinityProfileZoneAffinityListWithPriorityType {
	returnZoneAffinityProfileZoneAffinityListWithPriorityTypeType := new(ReturnZoneAffinityProfileZoneAffinityListWithPriorityType)
	return returnZoneAffinityProfileZoneAffinityListWithPriorityTypeType
}

type RogueApPolicy struct {
	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Id *string `json:"id,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Rules []*RogueApRuleList `json:"rules,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewRogueApPolicy() *RogueApPolicy {
	rogueApPolicyType := new(RogueApPolicy)
	return rogueApPolicyType
}

func NewDefaultRogueApPolicy() *RogueApPolicy {
	rogueApPolicyType := new(RogueApPolicy)
	return rogueApPolicyType
}

type RogueApPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RogueApPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRogueApPolicyList() *RogueApPolicyList {
	rogueApPolicyListType := new(RogueApPolicyList)
	return rogueApPolicyListType
}

func NewDefaultRogueApPolicyList() *RogueApPolicyList {
	rogueApPolicyListType := new(RogueApPolicyList)
	return rogueApPolicyListType
}

type RogueApRuleList struct {
	// Classification
	// Constraints:
	//    - required
	//    - oneof:[Ignore,Known,Rogue,Malicious]
	Classification *string `json:"classification" validate:"required,oneof=Ignore Known Rogue Malicious"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalNameAllowBlank `json:"name" validate:"required"`

	// Priority
	// Constraints:
	//    - required
	Priority *int `json:"priority" validate:"required"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[AdhocRule,SsidSpoofingRule,MacSpoofingRule,SameNetworkRule,CustomSsidRule,CustomRssiRule,CustomMacOuiRule]
	Type *string `json:"type" validate:"required,oneof=AdhocRule SsidSpoofingRule MacSpoofingRule SameNetworkRule CustomSsidRule CustomRssiRule CustomMacOuiRule"`

	Value interface{} `json:"value,omitempty"`
}

func NewRogueApRuleList() *RogueApRuleList {
	rogueApRuleListType := new(RogueApRuleList)
	return rogueApRuleListType
}

func NewDefaultRogueApRuleList() *RogueApRuleList {
	rogueApRuleListType := new(RogueApRuleList)
	return rogueApRuleListType
}

type RtlsProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateRtlsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRtlsProfileList() *RtlsProfileList {
	rtlsProfileListType := new(RtlsProfileList)
	return rtlsProfileListType
}

func NewDefaultRtlsProfileList() *RtlsProfileList {
	rtlsProfileListType := new(RtlsProfileList)
	return rtlsProfileListType
}

type RuckusGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// TunnelEncryption
	// Tunnel Encryption of the RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,AES128,AES256]
	TunnelEncryption *string `json:"tunnelEncryption,omitempty" validate:"omitempty,oneof=DISABLE AES128 AES256"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"omitempty,oneof=GRE GREUDP"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"omitempty,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewRuckusGREProfile() *RuckusGREProfile {
	ruckusGREProfileType := new(RuckusGREProfile)
	return ruckusGREProfileType
}

func NewDefaultRuckusGREProfile() *RuckusGREProfile {
	ruckusGREProfileType := new(RuckusGREProfile)
	return ruckusGREProfileType
}

type RuckusGREProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RuckusGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRuckusGREProfileList() *RuckusGREProfileList {
	ruckusGREProfileListType := new(RuckusGREProfileList)
	return ruckusGREProfileListType
}

func NewDefaultRuckusGREProfileList() *RuckusGREProfileList {
	ruckusGREProfileListType := new(RuckusGREProfileList)
	return ruckusGREProfileListType
}

type SoftGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	IpMode *IpMode `json:"ipMode,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"omitempty,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewSoftGREProfile() *SoftGREProfile {
	softGREProfileType := new(SoftGREProfile)
	return softGREProfileType
}

func NewDefaultSoftGREProfile() *SoftGREProfile {
	softGREProfileType := new(SoftGREProfile)
	return softGREProfileType
}

type SoftGREProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SoftGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSoftGREProfileList() *SoftGREProfileList {
	softGREProfileListType := new(SoftGREProfileList)
	return softGREProfileListType
}

func NewDefaultSoftGREProfileList() *SoftGREProfileList {
	softGREProfileListType := new(SoftGREProfileList)
	return softGREProfileListType
}

type TrafficClassProfileList struct {
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

	List []*common.TrafficClassProfileRef `json:"list,omitempty"`

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

func NewTrafficClassProfileList() *TrafficClassProfileList {
	trafficClassProfileListType := new(TrafficClassProfileList)
	return trafficClassProfileListType
}

func NewDefaultTrafficClassProfileList() *TrafficClassProfileList {
	trafficClassProfileListType := new(TrafficClassProfileList)
	return trafficClassProfileListType
}

// TtgCommonSetting
//
// Hosted AAA server RADIUS settings & PLMN ID settings
type TtgCommonSetting struct {
	// MobileCountryCode
	// Mobile country code
	// Constraints:
	//    - nullable
	//    - max:3
	//    - min:3
	MobileCountryCode *string `json:"mobileCountryCode,omitempty" validate:"omitempty,max=3,min=3"`

	// MobileNetworkCode
	// Mobile network code
	// Constraints:
	//    - nullable
	//    - max:3
	//    - min:2
	MobileNetworkCode *string `json:"mobileNetworkCode,omitempty" validate:"omitempty,max=3,min=2"`
}

func NewTtgCommonSetting() *TtgCommonSetting {
	ttgCommonSettingType := new(TtgCommonSetting)
	return ttgCommonSettingType
}

func NewDefaultTtgCommonSetting() *TtgCommonSetting {
	ttgCommonSettingType := new(TtgCommonSetting)
	return ttgCommonSettingType
}

type TtgpdgApnForwardingRealm struct {
	// Apn
	// the forwarding policy APN, if apnType is NIOI, APN Example : internet-v4.mnc111.mcc222.gprs
	Apn *string `json:"apn,omitempty"`

	// ApnType
	// type of the forwarding policy APN.
	// Constraints:
	//    - nullable
	//    - oneof:[NI,NIOI]
	ApnType *string `json:"apnType,omitempty" validate:"omitempty,oneof=NI NIOI"`

	// RouteType
	// routeType of the forwarding policy APN.
	// Constraints:
	//    - nullable
	//    - oneof:[GTPv1,GTPv2,PDG]
	RouteType *string `json:"routeType,omitempty" validate:"omitempty,oneof=GTPv1 GTPv2 PDG"`
}

func NewTtgpdgApnForwardingRealm() *TtgpdgApnForwardingRealm {
	ttgpdgApnForwardingRealmType := new(TtgpdgApnForwardingRealm)
	return ttgpdgApnForwardingRealmType
}

func NewDefaultTtgpdgApnForwardingRealm() *TtgpdgApnForwardingRealm {
	ttgpdgApnForwardingRealmType := new(TtgpdgApnForwardingRealm)
	return ttgpdgApnForwardingRealmType
}

type TtgpdgCommonSetting struct {
	// AcctRetry
	// Accounting retry of TTG PDG common setting
	AcctRetry *int `json:"acctRetry,omitempty"`

	// AcctRetryTimeout
	// Accounting retry timeout(secs) of TTG PDG common setting
	AcctRetryTimeout *int `json:"acctRetryTimeout,omitempty"`

	// ApnFormat2GGSN
	// APN format to GGSN of TTG PDG common setting
	// Constraints:
	//    - nullable
	//    - oneof:[DNS,String]
	ApnFormat2GGSN *string `json:"apnFormat2GGSN,omitempty" validate:"omitempty,oneof=DNS String"`

	// ApnOIInUse
	// APN-OI of TTG PDG common setting
	ApnOIInUse *bool `json:"apnOIInUse,omitempty"`

	// PdgUeIdleTimeout
	// PDG UE session idle timeout(secs) of TTG PDG common setting
	PdgUeIdleTimeout *int `json:"pdgUeIdleTimeout,omitempty"`
}

func NewTtgpdgCommonSetting() *TtgpdgCommonSetting {
	ttgpdgCommonSettingType := new(TtgpdgCommonSetting)
	return ttgpdgCommonSettingType
}

func NewDefaultTtgpdgCommonSetting() *TtgpdgCommonSetting {
	ttgpdgCommonSettingType := new(TtgpdgCommonSetting)
	return ttgpdgCommonSettingType
}

type TtgpdgProfile struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	ApnForwardingRealms []*TtgpdgApnForwardingRealm `json:"apnForwardingRealms,omitempty"`

	// ApnRealms
	// List of the Default APN
	ApnRealms []*ApnRealm `json:"apnRealms,omitempty"`

	CommonSetting *TtgpdgCommonSetting `json:"commonSetting,omitempty"`

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

	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`
}

func NewTtgpdgProfile() *TtgpdgProfile {
	ttgpdgProfileType := new(TtgpdgProfile)
	return ttgpdgProfileType
}

func NewDefaultTtgpdgProfile() *TtgpdgProfile {
	ttgpdgProfileType := new(TtgpdgProfile)
	return ttgpdgProfileType
}

type TtgpdgProfileConfiguration struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	ApnForwardingRealms []*TtgpdgApnForwardingRealm `json:"apnForwardingRealms,omitempty"`

	// ApnRealms
	// List of the Default APN
	ApnRealms []*ApnRealm `json:"apnRealms,omitempty"`

	CommonSetting *TtgpdgCommonSetting `json:"commonSetting,omitempty"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN,omitempty"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DhcpRelay *DhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewTtgpdgProfileConfiguration() *TtgpdgProfileConfiguration {
	ttgpdgProfileConfigurationType := new(TtgpdgProfileConfiguration)
	return ttgpdgProfileConfigurationType
}

func NewDefaultTtgpdgProfileConfiguration() *TtgpdgProfileConfiguration {
	ttgpdgProfileConfigurationType := new(TtgpdgProfileConfiguration)
	return ttgpdgProfileConfigurationType
}

type TtgpdgProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*TtgpdgProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewTtgpdgProfileList() *TtgpdgProfileList {
	ttgpdgProfileListType := new(TtgpdgProfileList)
	return ttgpdgProfileListType
}

func NewDefaultTtgpdgProfileList() *TtgpdgProfileList {
	ttgpdgProfileListType := new(TtgpdgProfileList)
	return ttgpdgProfileListType
}

type UpdateL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*DataPlaneL3RoamingData `json:"dataPlanes,omitempty"`
}

func NewUpdateL3RoamingConfig() *UpdateL3RoamingConfig {
	updateL3RoamingConfigType := new(UpdateL3RoamingConfig)
	return updateL3RoamingConfigType
}

func NewDefaultUpdateL3RoamingConfig() *UpdateL3RoamingConfig {
	updateL3RoamingConfigType := new(UpdateL3RoamingConfig)
	return updateL3RoamingConfigType
}

type UpdatePrecedenceProfile struct {
	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*RateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*VlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

func NewUpdatePrecedenceProfile() *UpdatePrecedenceProfile {
	updatePrecedenceProfileType := new(UpdatePrecedenceProfile)
	return updatePrecedenceProfileType
}

func NewDefaultUpdatePrecedenceProfile() *UpdatePrecedenceProfile {
	updatePrecedenceProfileType := new(UpdatePrecedenceProfile)
	return updatePrecedenceProfileType
}

type UpdateRogueApPolicy struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Rules []*RogueApRuleList `json:"rules,omitempty"`
}

func NewUpdateRogueApPolicy() *UpdateRogueApPolicy {
	updateRogueApPolicyType := new(UpdateRogueApPolicy)
	return updateRogueApPolicyType
}

func NewDefaultUpdateRogueApPolicy() *UpdateRogueApPolicy {
	updateRogueApPolicyType := new(UpdateRogueApPolicy)
	return updateRogueApPolicyType
}

type UpdateRtlsProfile struct {
	EkahauEnabled *bool `json:"ekahauEnabled,omitempty"`

	EkahauIp *common.IpAddress `json:"ekahauIp,omitempty"`

	EkahauPort *int `json:"ekahauPort,omitempty"`

	Id *string `json:"id,omitempty"`

	StanleyEnabled *bool `json:"stanleyEnabled,omitempty"`
}

func NewUpdateRtlsProfile() *UpdateRtlsProfile {
	updateRtlsProfileType := new(UpdateRtlsProfile)
	return updateRtlsProfileType
}

func NewDefaultUpdateRtlsProfile() *UpdateRtlsProfile {
	updateRtlsProfileType := new(UpdateRtlsProfile)
	return updateRtlsProfileType
}

type UplinkRateLimiting struct {
	// UplinkRateLimitingBps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	UplinkRateLimitingBps *string `json:"uplinkRateLimitingBps,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled or disabled
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`
}

func NewUplinkRateLimiting() *UplinkRateLimiting {
	uplinkRateLimitingType := new(UplinkRateLimiting)
	return uplinkRateLimitingType
}

func NewDefaultUplinkRateLimiting() *UplinkRateLimiting {
	uplinkRateLimitingType := new(UplinkRateLimiting)
	uplinkRateLimitingEnabledField := false
	uplinkRateLimitingType.UplinkRateLimitingEnabled = &uplinkRateLimitingEnabledField
	return uplinkRateLimitingType
}

type UserTrafficProfile struct {
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
	//    - nullable
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"omitempty,oneof=BLOCK ALLOW"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// Id
	// Identifier of the user traffic profile
	Id *string `json:"id,omitempty"`

	// IpAclRules
	// Traffic access control list
	IpAclRules []*IpAclRules `json:"ipAclRules,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// QmAppPolicyId
	// Application Policy UUID
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *UplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewUserTrafficProfile() *UserTrafficProfile {
	userTrafficProfileType := new(UserTrafficProfile)
	return userTrafficProfileType
}

func NewDefaultUserTrafficProfile() *UserTrafficProfile {
	userTrafficProfileType := new(UserTrafficProfile)
	defaultActionField := `ALLOW`
	userTrafficProfileType.DefaultAction = &defaultActionField
	return userTrafficProfileType
}

type UserTrafficProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UserTrafficProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewUserTrafficProfileList() *UserTrafficProfileList {
	userTrafficProfileListType := new(UserTrafficProfileList)
	return userTrafficProfileListType
}

func NewDefaultUserTrafficProfileList() *UserTrafficProfileList {
	userTrafficProfileListType := new(UserTrafficProfileList)
	return userTrafficProfileListType
}

type VdpProfile struct {
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

func NewVdpProfile() *VdpProfile {
	vdpProfileType := new(VdpProfile)
	return vdpProfileType
}

func NewDefaultVdpProfile() *VdpProfile {
	vdpProfileType := new(VdpProfile)
	return vdpProfileType
}

// VlanPrecedenceItem
//
// Vlan precedence item
type VlanPrecedenceItem struct {
	// Name
	// Name of the Vlan precedence item
	// Constraints:
	//    - nullable
	//    - oneof:[AAA,DEVICE,WLAN]
	Name *string `json:"name,omitempty" validate:"omitempty,oneof=AAA DEVICE WLAN"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

func NewVlanPrecedenceItem() *VlanPrecedenceItem {
	vlanPrecedenceItemType := new(VlanPrecedenceItem)
	return vlanPrecedenceItemType
}

func NewDefaultVlanPrecedenceItem() *VlanPrecedenceItem {
	vlanPrecedenceItemType := new(VlanPrecedenceItem)
	return vlanPrecedenceItemType
}

type ZoneAffinityProfileList struct {
	List []*ReturnZoneAffinityProfile `json:"list,omitempty"`
}

func NewZoneAffinityProfileList() *ZoneAffinityProfileList {
	zoneAffinityProfileListType := new(ZoneAffinityProfileList)
	return zoneAffinityProfileListType
}

func NewDefaultZoneAffinityProfileList() *ZoneAffinityProfileList {
	zoneAffinityProfileListType := new(ZoneAffinityProfileList)
	return zoneAffinityProfileListType
}
