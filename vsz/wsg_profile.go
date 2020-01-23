package vsz

// API Version: v8_1

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

func NewWSGProfileAccountingProfile() *WSGProfileAccountingProfile {
	m := new(WSGProfileAccountingProfile)
	return m
}

type WSGProfileAccountingProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileAccountingProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileAccountingProfileList() *WSGProfileAccountingProfileList {
	m := new(WSGProfileAccountingProfileList)
	return m
}

// WSGProfileAcctServiceRealmMapping
//
// Accounting service per realm
type WSGProfileAcctServiceRealmMapping struct {
	// Id
	// Accounting service UUID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty" validate:"omitempty"`

	// Name
	// Accounting service name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty" validate:"omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`

	// ServiceType
	// Accounting service type, NA is NA-Request Rejected
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,CGF]
	ServiceType *string `json:"serviceType" validate:"required,oneof=NA RADIUS CGF"`
}

func NewWSGProfileAcctServiceRealmMapping() *WSGProfileAcctServiceRealmMapping {
	m := new(WSGProfileAcctServiceRealmMapping)
	return m
}

// WSGProfileAdvancedOptionContent
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
	EnforceNatt *string `json:"enforceNatt,omitempty" validate:"oneof=Disabled Enabled"`

	// FailoverMode
	// mode of the failover
	// Constraints:
	//    - oneof:[Non_Revertive,Revertive]
	FailoverMode *string `json:"failoverMode,omitempty" validate:"oneof=Non_Revertive Revertive"`

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
	IpcompEnable *string `json:"ipcompEnable,omitempty" validate:"oneof=Disabled Enabled"`

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

func NewWSGProfileAuthenticationProfile() *WSGProfileAuthenticationProfile {
	m := new(WSGProfileAuthenticationProfile)
	return m
}

type WSGProfileAuthenticationProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileAuthenticationProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileAuthenticationProfileList() *WSGProfileAuthenticationProfileList {
	m := new(WSGProfileAuthenticationProfileList)
	return m
}

type WSGProfileBaseServiceInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCommonBaseServiceInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBaseServiceInfoList() *WSGProfileBaseServiceInfoList {
	m := new(WSGProfileBaseServiceInfoList)
	return m
}

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
	Mac *WSGCommonMac `json:"mac" validate:"required"`

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

func NewWSGProfileBlockClient() *WSGProfileBlockClient {
	m := new(WSGProfileBlockClient)
	return m
}

type WSGProfileBlockClientList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileBlockClientListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBlockClientList() *WSGProfileBlockClientList {
	m := new(WSGProfileBlockClientList)
	return m
}

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

type WSGProfileBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	// Constraints:
	//    - required
	BonjourFencingRuleList []*WSGProfileBonjourFencingRule `json:"bonjourFencingRuleList" validate:"required,dive"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBonjourFencingPolicy() *WSGProfileBonjourFencingPolicy {
	m := new(WSGProfileBonjourFencingPolicy)
	return m
}

type WSGProfileBonjourFencingPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileBonjourFencingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBonjourFencingPolicyList() *WSGProfileBonjourFencingPolicyList {
	m := new(WSGProfileBonjourFencingPolicyList)
	return m
}

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
	ServiceType *WSGProfileBridgeService `json:"serviceType" validate:"required,oneof=AIRDISK AIRPLAY AIRPORT_MANAGEMENT AIRPRINT AIRTUNES APPLE_FILE_SHARING APPLE_MOBILE_DEVICES APPLETV ICLOUD_SYNC ITUNES_REMOTE ITUNES_SHARING OPEN_DIRECTORY_MASTER OPTICAL_DISK_SHARING SCREEN_SHARING SECURE_FILE_SHARING SECURE_SHELL WWW_HTTP WWW_HTTPS WORKGROUP_MANAGER XGRID GOOGLE_CHROMECAST OTHER"`
}

func NewWSGProfileBonjourFencingRule() *WSGProfileBonjourFencingRule {
	m := new(WSGProfileBonjourFencingRule)
	return m
}

type WSGProfileBonjourFencingRuleDeviceMac struct {
	Mac *WSGCommonMac `json:"mac,omitempty"`
}

func NewWSGProfileBonjourFencingRuleDeviceMac() *WSGProfileBonjourFencingRuleDeviceMac {
	m := new(WSGProfileBonjourFencingRuleDeviceMac)
	return m
}

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

type WSGProfileBonjourFencingService struct {
	NeighborApMac *string `json:"neighborApMac,omitempty"`

	NeighborApName *string `json:"neighborApName,omitempty"`

	ServiceType *WSGProfileBridgeService `json:"serviceType,omitempty"`

	// SourceType
	// Constraints:
	//    - oneof:[UNKNOWN,DIRECT,NEIGHBOR]
	SourceType *string `json:"sourceType,omitempty" validate:"oneof=UNKNOWN DIRECT NEIGHBOR"`
}

func NewWSGProfileBonjourFencingService() *WSGProfileBonjourFencingService {
	m := new(WSGProfileBonjourFencingService)
	return m
}

type WSGProfileBonjourFencingStatistic struct {
	ApMac *string `json:"apMac,omitempty"`

	DroppedPacketsDueToNeighbor *int `json:"droppedPacketsDueToNeighbor,omitempty"`

	DroppedPacketsDueToServiceType *int `json:"droppedPacketsDueToServiceType,omitempty"`

	ForwardedPackets *int `json:"forwardedPackets,omitempty"`

	ServiceList []*WSGProfileBonjourFencingService `json:"serviceList,omitempty"`
}

func NewWSGProfileBonjourFencingStatistic() *WSGProfileBonjourFencingStatistic {
	m := new(WSGProfileBonjourFencingStatistic)
	return m
}

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

func NewWSGProfileBridgeProfile() *WSGProfileBridgeProfile {
	m := new(WSGProfileBridgeProfile)
	return m
}

type WSGProfileBridgeProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileBridgeProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBridgeProfileList() *WSGProfileBridgeProfileList {
	m := new(WSGProfileBridgeProfileList)
	return m
}

type WSGProfileBridgeService string

func NewWSGProfileBridgeService() *WSGProfileBridgeService {
	m := new(WSGProfileBridgeService)
	return m
}

type WSGProfileBulkBlockClient struct {
	BlockClientList []*WSGProfileBulkBlockClientBlockClientListType `json:"blockClientList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`
}

func NewWSGProfileBulkBlockClient() *WSGProfileBulkBlockClient {
	m := new(WSGProfileBulkBlockClient)
	return m
}

type WSGProfileBulkBlockClientBlockClientListType struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *WSGCommonMac `json:"apMac" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBulkBlockClientBlockClientListType() *WSGProfileBulkBlockClientBlockClientListType {
	m := new(WSGProfileBulkBlockClientBlockClientListType)
	return m
}

type WSGProfileClientIsolationEntry struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Ip
	// Client Entry ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty" validate:"omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`
}

func NewWSGProfileClientIsolationEntry() *WSGProfileClientIsolationEntry {
	m := new(WSGProfileClientIsolationEntry)
	return m
}

type WSGProfileClientIsolationWhitelist struct {
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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - required
	Whitelist []*WSGProfileClientIsolationEntry `json:"whitelist" validate:"required,dive"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileClientIsolationWhitelist() *WSGProfileClientIsolationWhitelist {
	m := new(WSGProfileClientIsolationWhitelist)
	return m
}

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

func NewWSGProfileClientIsolationWhitelistArray() *WSGProfileClientIsolationWhitelistArray {
	m := new(WSGProfileClientIsolationWhitelistArray)
	return m
}

// WSGProfileCmProtocolOptionContent
//
// Certificate Management Protocol Option
type WSGProfileCmProtocolOptionContent struct {
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

func NewWSGProfileCmProtocolOptionContent() *WSGProfileCmProtocolOptionContent {
	m := new(WSGProfileCmProtocolOptionContent)
	return m
}

type WSGProfileCoreNetworkGateway struct {
	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:32767
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty" validate:"gte=1,lte=32767"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - default:3
	//    - min:1
	//    - max:255
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty" validate:"gte=1,lte=255"`

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
	TunnelMTU *string `json:"tunnelMTU,omitempty" validate:"oneof=AUTO MANUAL"`

	// TunnelMTUSize
	// Manual setting value of Gateway path MTU
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:1500
	TunnelMTUSize *int `json:"tunnelMTUSize,omitempty" validate:"gte=850,lte=1500"`
}

func NewWSGProfileCoreNetworkGateway() *WSGProfileCoreNetworkGateway {
	m := new(WSGProfileCoreNetworkGateway)
	return m
}

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*WSGProfileAcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

func NewWSGProfileCreateAccountingProfile() *WSGProfileCreateAccountingProfile {
	m := new(WSGProfileCreateAccountingProfile)
	return m
}

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*WSGProfileRealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *WSGProfileTtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewWSGProfileCreateAuthenticationProfile() *WSGProfileCreateAuthenticationProfile {
	m := new(WSGProfileCreateAuthenticationProfile)
	return m
}

type WSGProfileCreateBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	// Constraints:
	//    - required
	BonjourFencingRuleList []*WSGProfileBonjourFencingRule `json:"bonjourFencingRuleList" validate:"required,dive"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*WSGProfileBonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGProfileCreateBonjourFencingPolicy() *WSGProfileCreateBonjourFencingPolicy {
	m := new(WSGProfileCreateBonjourFencingPolicy)
	return m
}

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGProfileCreateBridgeProfile() *WSGProfileCreateBridgeProfile {
	m := new(WSGProfileCreateBridgeProfile)
	return m
}

type WSGProfileCreateClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	// Constraints:
	//    - required
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - required
	Whitelist []*WSGProfileClientIsolationEntry `json:"whitelist" validate:"required,dive"`
}

func NewWSGProfileCreateClientIsolationWhitelist() *WSGProfileCreateClientIsolationWhitelist {
	m := new(WSGProfileCreateClientIsolationWhitelist)
	return m
}

type WSGProfileCreateDhcpProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// PoolEndIp
	// Constraints:
	//    - required
	PoolEndIp *WSGCommonIpAddress `json:"poolEndIp" validate:"required"`

	// PoolStartIp
	// Constraints:
	//    - required
	PoolStartIp *WSGCommonIpAddress `json:"poolStartIp" validate:"required"`

	PrimaryDnsIp *WSGCommonIpAddress `json:"primaryDnsIp,omitempty"`

	SecondaryDnsIp *WSGCommonIpAddress `json:"secondaryDnsIp,omitempty"`

	// SubnetMask
	// Constraints:
	//    - required
	SubnetMask *WSGCommonIpAddress `json:"subnetMask" validate:"required"`

	// SubnetNetworkIp
	// Constraints:
	//    - required
	SubnetNetworkIp *WSGCommonIpAddress `json:"subnetNetworkIp" validate:"required"`

	// VlanId
	// VLAN ID of the DHCP Profile
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	VlanId *int `json:"vlanId" validate:"required,gte=1,lte=4094"`
}

func NewWSGProfileCreateDhcpProfile() *WSGProfileCreateDhcpProfile {
	m := new(WSGProfileCreateDhcpProfile)
	return m
}

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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

func NewWSGProfileCreateDnsServerProfile() *WSGProfileCreateDnsServerProfile {
	m := new(WSGProfileCreateDnsServerProfile)
	return m
}

type WSGProfileCreateIpsecProfile struct {
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"oneof=PresharedKey Certificate"`

	CmProtocolOption *WSGProfileCmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	// Constraints:
	//    - required
	EspRekeyTime *float64 `json:"espRekeyTime" validate:"required"`

	EspRekeyTimeUnit *WSGCommonTimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *WSGProfileEspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	// Constraints:
	//    - required
	IkeRekeyTime *float64 `json:"ikeRekeyTime" validate:"required"`

	IkeRekeyTimeUnit *WSGCommonTimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *WSGProfileIkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	// IpMode
	// Constraints:
	//    - required
	IpMode *WSGProfileIpMode `json:"ipMode" validate:"required,oneof=IPV4 IPV6"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"oneof=SOFT_GRE RUCKUS_GRE"`
}

func NewWSGProfileCreateIpsecProfile() *WSGProfileCreateIpsecProfile {
	m := new(WSGProfileCreateIpsecProfile)
	return m
}

type WSGProfileCreateL2oGREProfile struct {
	// CoreNetworkGateway
	// Constraints:
	//    - required
	CoreNetworkGateway *WSGProfileCoreNetworkGateway `json:"coreNetworkGateway" validate:"required"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGProfileCreateL2oGREProfile() *WSGProfileCreateL2oGREProfile {
	m := new(WSGProfileCreateL2oGREProfile)
	return m
}

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

func NewWSGProfileCreatePrecedenceProfile() *WSGProfileCreatePrecedenceProfile {
	m := new(WSGProfileCreatePrecedenceProfile)
	return m
}

type WSGProfileCreateResultList []*WSGCommonCreateResult

func MakeWSGProfileCreateResultList() WSGProfileCreateResultList {
	m := make(WSGProfileCreateResultList, 0)
	return m
}

type WSGProfileCreateRogueApPolicy struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Rules
	// Constraints:
	//    - required
	Rules []*WSGProfileRogueApRuleList `json:"rules" validate:"required,dive"`
}

func NewWSGProfileCreateRogueApPolicy() *WSGProfileCreateRogueApPolicy {
	m := new(WSGProfileCreateRogueApPolicy)
	return m
}

type WSGProfileCreateRtlsProfile struct {
	// EkahauEnabled
	// Eekahau Location Service Enabled
	// Constraints:
	//    - required
	EkahauEnabled *bool `json:"ekahauEnabled" validate:"required"`

	EkahauIp *WSGCommonIpAddress `json:"ekahauIp,omitempty"`

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

func NewWSGProfileCreateRtlsProfile() *WSGProfileCreateRtlsProfile {
	m := new(WSGProfileCreateRtlsProfile)
	return m
}

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// TunnelEncryption
	// Tunnel Encryption of the RuckusGRE profile
	// Constraints:
	//    - oneof:[DISABLE,AES128,AES256]
	TunnelEncryption *string `json:"tunnelEncryption,omitempty" validate:"oneof=DISABLE AES128 AES256"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"oneof=GRE GREUDP"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - required
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled" validate:"required,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"gte=850,lte=9018"`
}

func NewWSGProfileCreateRuckusGREProfile() *WSGProfileCreateRuckusGREProfile {
	m := new(WSGProfileCreateRuckusGREProfile)
	return m
}

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"gte=850,lte=9018"`
}

func NewWSGProfileCreateSoftGREProfile() *WSGProfileCreateSoftGREProfile {
	m := new(WSGProfileCreateSoftGREProfile)
	return m
}

type WSGProfileCreateTrafficClassProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName2to64 `json:"name" validate:"required,max=64,min=2"`

	// TrafficClasses
	// Constraints:
	//    - required
	TrafficClasses []*WSGCommonTrafficClassRef `json:"trafficClasses" validate:"required,dive"`
}

func NewWSGProfileCreateTrafficClassProfile() *WSGProfileCreateTrafficClassProfile {
	m := new(WSGProfileCreateTrafficClassProfile)
	return m
}

type WSGProfileCreateTtgpdgProfile struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	// Constraints:
	//    - required
	ApnForwardingRealms []*WSGProfileTtgpdgApnForwardingRealm `json:"apnForwardingRealms" validate:"required,dive"`

	// ApnRealms
	// List of the Default APN
	ApnRealms []*WSGProfileApnRealm `json:"apnRealms,omitempty"`

	// CommonSetting
	// Constraints:
	//    - required
	CommonSetting *WSGProfileTtgpdgCommonSetting `json:"commonSetting" validate:"required"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGProfileCreateTtgpdgProfile() *WSGProfileCreateTtgpdgProfile {
	m := new(WSGProfileCreateTtgpdgProfile)
	return m
}

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
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=BLOCK ALLOW"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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

type WSGProfileCreateZoneAffinityProfile struct {
	// Description
	// The description of the profile
	// Constraints:
	//    - max:64
	Description *string `json:"description,omitempty" validate:"max=64"`

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
	ZoneAffinityList []string `json:"zoneAffinityList" validate:"required,dive"`
}

func NewWSGProfileCreateZoneAffinityProfile() *WSGProfileCreateZoneAffinityProfile {
	m := new(WSGProfileCreateZoneAffinityProfile)
	return m
}

type WSGProfileDataPlaneL3RoamingData struct {
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
	//    - oneof:[VLAN,SUBNET]
	SubCriteriaType *string `json:"subCriteriaType,omitempty" validate:"oneof=VLAN SUBNET"`

	// Value
	// A list of L3 roaming configuration for this DP
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`
}

func NewWSGProfileDataPlaneL3RoamingData() *WSGProfileDataPlaneL3RoamingData {
	m := new(WSGProfileDataPlaneL3RoamingData)
	return m
}

type WSGProfileDeleteBulkAccountingProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkAccountingProfile() *WSGProfileDeleteBulkAccountingProfile {
	m := new(WSGProfileDeleteBulkAccountingProfile)
	return m
}

type WSGProfileDeleteBulkAuthenticationProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkAuthenticationProfile() *WSGProfileDeleteBulkAuthenticationProfile {
	m := new(WSGProfileDeleteBulkAuthenticationProfile)
	return m
}

type WSGProfileDeleteBulkPrecedenceProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkPrecedenceProfile() *WSGProfileDeleteBulkPrecedenceProfile {
	m := new(WSGProfileDeleteBulkPrecedenceProfile)
	return m
}

type WSGProfileDeleteBulkUserTrafficProfile struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkUserTrafficProfile() *WSGProfileDeleteBulkUserTrafficProfile {
	m := new(WSGProfileDeleteBulkUserTrafficProfile)
	return m
}

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
	Subopt1Format *string `json:"subopt1Format,omitempty" validate:"oneof=AP_INFO AP_MAC_hex AP_MAC_hex_ESSID AP_INFO_LOCATION"`

	// Subopt2Enabled
	// Enable subopt-2
	Subopt2Enabled *bool `json:"subopt2Enabled,omitempty"`

	// Subopt2Format
	// Subopt-2 format
	// Constraints:
	//    - oneof:[CLIENT_MAC_hex,CLIENT_MAC_hex_ESSID,AP_MAC_hex,AP_MAC__hex_ESSID]
	Subopt2Format *string `json:"subopt2Format,omitempty" validate:"oneof=CLIENT_MAC_hex CLIENT_MAC_hex_ESSID AP_MAC_hex AP_MAC__hex_ESSID"`

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
	Subopt151Format *string `json:"subopt151Format,omitempty" validate:"oneof=AREA_NAME ESSID"`
}

func NewWSGProfileDhcpOption82() *WSGProfileDhcpOption82 {
	m := new(WSGProfileDhcpOption82)
	return m
}

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

func NewWSGProfileDhcpProfileList() *WSGProfileDhcpProfileList {
	m := new(WSGProfileDhcpProfileList)
	return m
}

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

func NewWSGProfileDnsServerProfile() *WSGProfileDnsServerProfile {
	m := new(WSGProfileDnsServerProfile)
	return m
}

type WSGProfileDnsServerProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileDnsServerProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileDnsServerProfileList() *WSGProfileDnsServerProfileList {
	m := new(WSGProfileDnsServerProfileList)
	return m
}

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
// EspProposal based ipsec service mappings
type WSGProfileEspProposal struct {
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

func NewWSGProfileEspProposal() *WSGProfileEspProposal {
	m := new(WSGProfileEspProposal)
	return m
}

// WSGProfileEspSecurityAssociationContent
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
	EspProposalType *string `json:"espProposalType,omitempty" validate:"oneof=Default Specific"`
}

func NewWSGProfileEspSecurityAssociationContent() *WSGProfileEspSecurityAssociationContent {
	m := new(WSGProfileEspSecurityAssociationContent)
	return m
}

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

func NewWSGProfileFlexiVpnProfileList() *WSGProfileFlexiVpnProfileList {
	m := new(WSGProfileFlexiVpnProfileList)
	return m
}

type WSGProfileGetL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*WSGProfileDataPlaneL3RoamingData `json:"dataPlanes,omitempty"`
}

func NewWSGProfileGetL3RoamingConfig() *WSGProfileGetL3RoamingConfig {
	m := new(WSGProfileGetL3RoamingConfig)
	return m
}

type WSGProfileHs20FriendlyName struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonLanguageName `json:"language" validate:"required,oneof=English Chinese Czech Danish Dutch French German Japanese Spanish Korean Swedish Turkish eng chi cze dan dut fre ger jpn kor spa swe tur"`

	// Name
	// Name of the friendly name
	// Constraints:
	//    - required
	//    - max:32
	//    - min:2
	Name *string `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGProfileHs20FriendlyName() *WSGProfileHs20FriendlyName {
	m := new(WSGProfileHs20FriendlyName)
	return m
}

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
	DomainNames []WSGCommonWildFQDN `json:"domainNames" validate:"required,dive"`

	// FriendlyNames
	// Friendly names
	// Constraints:
	//    - required
	FriendlyNames []*WSGProfileHs20FriendlyName `json:"friendlyNames" validate:"required,dive"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGProfileHs20Operator() *WSGProfileHs20Operator {
	m := new(WSGProfileHs20Operator)
	return m
}

type WSGProfileHs20OperatorList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileHs20Operator `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileHs20OperatorList() *WSGProfileHs20OperatorList {
	m := new(WSGProfileHs20OperatorList)
	return m
}

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

func NewWSGProfileHs20Provider() *WSGProfileHs20Provider {
	m := new(WSGProfileHs20Provider)
	return m
}

type WSGProfileHs20ProviderList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileHs20Provider `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileHs20ProviderList() *WSGProfileHs20ProviderList {
	m := new(WSGProfileHs20ProviderList)
	return m
}

// WSGProfileIkeProposal
//
// IkeProposal based ipsec service mappings
type WSGProfileIkeProposal struct {
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
	//    - oneof:[UseIntegrityALG,PRF_MD5,PRF_SHA1,PRF_AES_CBC,PRF_AES_MAC,PRF_SHA256,PRF_SHA384,PRF_SHA512]
	PrfAlg *string `json:"prfAlg,omitempty" validate:"oneof=UseIntegrityALG PRF_MD5 PRF_SHA1 PRF_AES_CBC PRF_AES_MAC PRF_SHA256 PRF_SHA384 PRF_SHA512"`
}

func NewWSGProfileIkeProposal() *WSGProfileIkeProposal {
	m := new(WSGProfileIkeProposal)
	return m
}

// WSGProfileIkeSecurityAssociationContent
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
	IkeProposalType *string `json:"ikeProposalType,omitempty" validate:"oneof=Default Specific"`
}

func NewWSGProfileIkeSecurityAssociationContent() *WSGProfileIkeSecurityAssociationContent {
	m := new(WSGProfileIkeSecurityAssociationContent)
	return m
}

type WSGProfileIpAclRules struct {
	// Action
	// The access of traffic access control.
	// Constraints:
	//    - default:'ALLOW'
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty" validate:"oneof=ALLOW BLOCK"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	// Constraints:
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty" validate:"gte=1,lte=255"`

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
	Direction *string `json:"direction,omitempty" validate:"oneof=UPSTREAM"`

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
	IpType *string `json:"ipType,omitempty" validate:"oneof=IPv4 IPv6"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=TCP UDP UDPLITE ICMP_ICMPV4 ICMPV6 IGMP ESP AH SCTP CUSTOM"`

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

type WSGProfileIpMode string

func NewWSGProfileIpMode() *WSGProfileIpMode {
	m := new(WSGProfileIpMode)
	return m
}

type WSGProfileIpsecProfile struct {
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"oneof=PresharedKey Certificate"`

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
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"oneof=SOFT_GRE RUCKUS_GRE"`
}

func NewWSGProfileIpsecProfile() *WSGProfileIpsecProfile {
	m := new(WSGProfileIpsecProfile)
	return m
}

type WSGProfileIpsecProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileIpsecProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileIpsecProfileList() *WSGProfileIpsecProfileList {
	m := new(WSGProfileIpsecProfileList)
	return m
}

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

func NewWSGProfileL2oGREProfile() *WSGProfileL2oGREProfile {
	m := new(WSGProfileL2oGREProfile)
	return m
}

type WSGProfileL2oGREProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileL2oGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileL2oGREProfileList() *WSGProfileL2oGREProfileList {
	m := new(WSGProfileL2oGREProfileList)
	return m
}

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

func NewWSGProfileLbsProfile() *WSGProfileLbsProfile {
	m := new(WSGProfileLbsProfile)
	return m
}

type WSGProfileLbsProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileLbsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileLbsProfileList() *WSGProfileLbsProfileList {
	m := new(WSGProfileLbsProfileList)
	return m
}

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

type WSGProfileModifyBlockClient struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`
}

func NewWSGProfileModifyBlockClient() *WSGProfileModifyBlockClient {
	m := new(WSGProfileModifyBlockClient)
	return m
}

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

type WSGProfileModifyHS20Operator struct {
	Certificate *WSGCommonGenericRef `json:"certificate,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	DomainNames []WSGCommonWildFQDN `json:"domainNames,omitempty"`

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
// Traffic access control list
type WSGProfileModifyIpAclRules struct {
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
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty" validate:"gte=1,lte=255"`

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
	//    - default:'IPv4'
	//    - oneof:[IPv4,IPv6]
	IpType *string `json:"ipType,omitempty" validate:"oneof=IPv4 IPv6"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=TCP UDP UDPLITE ICMP_ICMPV4 ICMPV6 IGMP ESP AH SCTP CUSTOM"`

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

type WSGProfileModifyIpsecProfile struct {
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"oneof=PresharedKey Certificate"`

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
	TunnelEncryption *string `json:"tunnelEncryption,omitempty" validate:"oneof=DISABLE AES128 AES256"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"oneof=GRE GREUDP"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"gte=850,lte=9018"`
}

func NewWSGProfileModifyRuckusGREProfile() *WSGProfileModifyRuckusGREProfile {
	m := new(WSGProfileModifyRuckusGREProfile)
	return m
}

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
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty" validate:"gte=1,lte=180"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - default:5
	//    - min:2
	//    - max:20
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty" validate:"gte=2,lte=20"`

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
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile. IPV4:850-1500, IPV6:1384-1500. Default 1500.
	// Constraints:
	//    - default:1500
	//    - min:850
	//    - max:9018
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty" validate:"gte=850,lte=9018"`
}

func NewWSGProfileModifySoftGREProfile() *WSGProfileModifySoftGREProfile {
	m := new(WSGProfileModifySoftGREProfile)
	return m
}

type WSGProfileModifyUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"oneof=BLOCK ALLOW"`

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

type WSGProfileModifyZoneAffinityProfile struct {
	// Description
	// The description of the profile
	// Constraints:
	//    - max:64
	Description *string `json:"description,omitempty" validate:"max=64"`

	// Name
	// Zone affinity profile name
	// Constraints:
	//    - max:64
	//    - min:1
	Name *string `json:"name,omitempty" validate:"max=64,min=1"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
}

func NewWSGProfileModifyZoneAffinityProfile() *WSGProfileModifyZoneAffinityProfile {
	m := new(WSGProfileModifyZoneAffinityProfile)
	return m
}

type WSGProfilePrecedenceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfilePrecedenceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfilePrecedenceList() *WSGProfilePrecedenceList {
	m := new(WSGProfilePrecedenceList)
	return m
}

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

type WSGProfileCloneRequest struct {
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

func NewWSGProfileCloneRequest() *WSGProfileCloneRequest {
	m := new(WSGProfileCloneRequest)
	return m
}

type WSGProfileCloneResponse struct {
	*WSGProfileCloneRequest
}

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

	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

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

func NewWSGProfileList() *WSGProfileList {
	m := new(WSGProfileList)
	return m
}

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

type WSGProfileProviderAccounting struct {
	// Id
	// Accounting id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty" validate:"omitempty"`

	// Name
	// Accounting name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty" validate:"omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`

	// ServiceType
	// Accounting service type
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,CGF]
	ServiceType *string `json:"serviceType" validate:"required,oneof=NA RADIUS CGF"`
}

func NewWSGProfileProviderAccounting() *WSGProfileProviderAccounting {
	m := new(WSGProfileProviderAccounting)
	return m
}

type WSGProfileProviderAuthentication struct {
	// Id
	// Authentication id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty" validate:"omitempty"`

	// Name
	// Authentication name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty" validate:"omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`

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

func NewWSGProfileProviderAuthentication() *WSGProfileProviderAuthentication {
	m := new(WSGProfileProviderAuthentication)
	return m
}

type WSGProfileProviderEAPAuthSetting struct {
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

func NewWSGProfileProviderEAPAuthSetting() *WSGProfileProviderEAPAuthSetting {
	m := new(WSGProfileProviderEAPAuthSetting)
	return m
}

type WSGProfileProviderEAPMethod struct {
	// AuthSettings
	// EAP method auth settings
	// Constraints:
	//    - nullable
	AuthSettings []*WSGProfileProviderEAPAuthSetting `json:"authSettings,omitempty" validate:"omitempty,dive"`

	// Type
	// EAP method type
	// Constraints:
	//    - required
	//    - oneof:[NA,MD5,EAP_TLS,EAP_Cisco,EAP_SIM,EAP_TTLS,EAP_AKA,PEAP,EAP_MSCHAP_V2,EAP_AKAs,Reserved]
	Type *string `json:"type" validate:"required,oneof=NA MD5 EAP_TLS EAP_Cisco EAP_SIM EAP_TTLS EAP_AKA PEAP EAP_MSCHAP_V2 EAP_AKAs Reserved"`
}

func NewWSGProfileProviderEAPMethod() *WSGProfileProviderEAPMethod {
	m := new(WSGProfileProviderEAPMethod)
	return m
}

type WSGProfileProviderExternalOSU struct {
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
	OsuServiceUrl *WSGCommonHTTPS `json:"osuServiceUrl" validate:"required,omitempty"`

	// ProvisioningProtocals
	// Provisioning protocal
	// Constraints:
	//    - required
	ProvisioningProtocals []WSGProfileProviderProvisionProtocal `json:"provisioningProtocals" validate:"required,dive,oneof=SOAP_XML OMA_DM"`

	// SubscriptionDescriptions
	// Subscription descriptions
	// Constraints:
	//    - required
	SubscriptionDescriptions []*WSGProfileProviderSubscriptionDescription `json:"subscriptionDescriptions" validate:"required,dive"`

	// WhitelistedDomains
	// Whitelisted domains
	WhitelistedDomains []WSGCommonWildFQDN `json:"whitelistedDomains,omitempty"`
}

func NewWSGProfileProviderExternalOSU() *WSGProfileProviderExternalOSU {
	m := new(WSGProfileProviderExternalOSU)
	return m
}

type WSGProfileProviderHomeOIs struct {
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

func NewWSGProfileProviderHomeOIs() *WSGProfileProviderHomeOIs {
	m := new(WSGProfileProviderHomeOIs)
	return m
}

type WSGProfileProviderInternalOSU struct {
	// Certificate
	// Constraints:
	//    - required
	Certificate *WSGCommonGenericRef `json:"certificate" validate:"required"`

	// CommonLanguageIcon
	// The base64 encoded data of icon.
	// Constraints:
	//    - required
	CommonLanguageIcon *string `json:"commonLanguageIcon" validate:"required"`

	// OsuAuthServices
	// Online signup authentication services
	// Constraints:
	//    - required
	OsuAuthServices []*WSGProfileProviderInternalOSUOsuAuthServicesType `json:"osuAuthServices" validate:"required,dive"`

	// OsuPortal
	// Constraints:
	//    - required
	OsuPortal *WSGProfileProviderInternalOSUOsuPortalType `json:"osuPortal" validate:"required"`

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
	ProvisioningProtocals []WSGProfileProviderProvisionProtocal `json:"provisioningProtocals" validate:"required,dive,oneof=SOAP_XML OMA_DM"`

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
	SubscriptionDescriptions []*WSGProfileProviderSubscriptionDescription `json:"subscriptionDescriptions" validate:"required,dive"`

	// WhitelistedDomains
	// whitelisted domains
	WhitelistedDomains []WSGCommonWildFQDN `json:"whitelistedDomains,omitempty"`
}

func NewWSGProfileProviderInternalOSU() *WSGProfileProviderInternalOSU {
	m := new(WSGProfileProviderInternalOSU)
	return m
}

type WSGProfileProviderInternalOSUOsuAuthServicesType struct {
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
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`
}

func NewWSGProfileProviderInternalOSUOsuAuthServicesType() *WSGProfileProviderInternalOSUOsuAuthServicesType {
	m := new(WSGProfileProviderInternalOSUOsuAuthServicesType)
	return m
}

type WSGProfileProviderInternalOSUOsuPortalType struct {
	ExternalUrl *WSGCommonHTTPS `json:"externalUrl,omitempty"`

	InternalOSUPortal *WSGCommonGenericRef `json:"internalOSUPortal,omitempty"`

	// Type
	// Portal type
	// Constraints:
	//    - required
	//    - oneof:[Internal,External]
	Type *string `json:"type" validate:"required,oneof=Internal External"`
}

func NewWSGProfileProviderInternalOSUOsuPortalType() *WSGProfileProviderInternalOSUOsuPortalType {
	m := new(WSGProfileProviderInternalOSUOsuPortalType)
	return m
}

type WSGProfileProviderOnlineSignup struct {
	ExternalOSU *WSGProfileProviderExternalOSU `json:"externalOSU,omitempty"`

	InternalOSU *WSGProfileProviderInternalOSU `json:"internalOSU,omitempty"`

	// Type
	// Online singup type
	// Constraints:
	//    - required
	//    - oneof:[Internal,External]
	Type *string `json:"type" validate:"required,oneof=Internal External"`
}

func NewWSGProfileProviderOnlineSignup() *WSGProfileProviderOnlineSignup {
	m := new(WSGProfileProviderOnlineSignup)
	return m
}

type WSGProfileProviderPLMN struct {
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

func NewWSGProfileProviderPLMN() *WSGProfileProviderPLMN {
	m := new(WSGProfileProviderPLMN)
	return m
}

type WSGProfileProviderProvisionProtocal string

func NewWSGProfileProviderProvisionProtocal() *WSGProfileProviderProvisionProtocal {
	m := new(WSGProfileProviderProvisionProtocal)
	return m
}

type WSGProfileProviderRealm struct {
	// EapMethods
	// EAP methods
	// Constraints:
	//    - required
	EapMethods []*WSGProfileProviderEAPMethod `json:"eapMethods" validate:"required,dive"`

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

func NewWSGProfileProviderRealm() *WSGProfileProviderRealm {
	m := new(WSGProfileProviderRealm)
	return m
}

type WSGProfileProviderSubscriptionDescription struct {
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
	Language *WSGCommonLanguageName `json:"language" validate:"required,oneof=English Chinese Czech Danish Dutch French German Japanese Spanish Korean Swedish Turkish eng chi cze dan dut fre ger jpn kor spa swe tur"`

	// Name
	// Name of the friendly name
	// Constraints:
	//    - required
	//    - max:252
	//    - min:2
	Name *string `json:"name" validate:"required,max=252,min=2"`
}

func NewWSGProfileProviderSubscriptionDescription() *WSGProfileProviderSubscriptionDescription {
	m := new(WSGProfileProviderSubscriptionDescription)
	return m
}

// WSGProfileRateLimitingPrecedenceItem
//
// Rate limiting precedence item
type WSGProfileRateLimitingPrecedenceItem struct {
	// Name
	// Name of rate limiting precedence item
	// Constraints:
	//    - oneof:[AAA,DEVICE,WLANUTP]
	Name *string `json:"name,omitempty" validate:"oneof=AAA DEVICE WLANUTP"`

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
// Realm based authentication service mappings
type WSGProfileRealmAuthServiceMapping struct {
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
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty" validate:"omitempty"`

	// Name
	// Authentication service name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty" validate:"omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`

	// ServiceType
	// Authentication service type, NA is NA-Request Rejected
	// Constraints:
	//    - required
	//    - oneof:[NA,RADIUS,LOCAL_DB,HLR,AD,LDAP]
	ServiceType *string `json:"serviceType" validate:"required,oneof=NA RADIUS LOCAL_DB HLR AD LDAP"`
}

func NewWSGProfileRealmAuthServiceMapping() *WSGProfileRealmAuthServiceMapping {
	m := new(WSGProfileRealmAuthServiceMapping)
	return m
}

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
	Name *string `json:"name,omitempty" validate:"max=64,min=1"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`

	ZoneAffinityListWithPriority []*WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType `json:"zoneAffinityListWithPriority,omitempty"`
}

func NewWSGProfileReturnZoneAffinityProfile() *WSGProfileReturnZoneAffinityProfile {
	m := new(WSGProfileReturnZoneAffinityProfile)
	return m
}

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

func NewWSGProfileRogueApPolicy() *WSGProfileRogueApPolicy {
	m := new(WSGProfileRogueApPolicy)
	return m
}

type WSGProfileRogueApPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileRogueApPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileRogueApPolicyList() *WSGProfileRogueApPolicyList {
	m := new(WSGProfileRogueApPolicyList)
	return m
}

type WSGProfileRogueApRuleList struct {
	// Classification
	// Constraints:
	//    - required
	//    - oneof:[Ignore,Known,Rogue,Malicious]
	Classification *string `json:"classification" validate:"required,oneof=Ignore Known Rogue Malicious"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalNameAllowBlank `json:"name" validate:"required,omitempty,max=64"`

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

func NewWSGProfileRogueApRuleList() *WSGProfileRogueApRuleList {
	m := new(WSGProfileRogueApRuleList)
	return m
}

type WSGProfileRtlsProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileCreateRtlsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileRtlsProfileList() *WSGProfileRtlsProfileList {
	m := new(WSGProfileRtlsProfileList)
	return m
}

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
	TunnelEncryption *string `json:"tunnelEncryption,omitempty" validate:"oneof=DISABLE AES128 AES256"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	// Constraints:
	//    - oneof:[GRE,GREUDP]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"oneof=GRE GREUDP"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	// Constraints:
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileRuckusGREProfile() *WSGProfileRuckusGREProfile {
	m := new(WSGProfileRuckusGREProfile)
	return m
}

type WSGProfileRuckusGREProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileRuckusGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileRuckusGREProfileList() *WSGProfileRuckusGREProfileList {
	m := new(WSGProfileRuckusGREProfileList)
	return m
}

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
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileSoftGREProfile() *WSGProfileSoftGREProfile {
	m := new(WSGProfileSoftGREProfile)
	return m
}

type WSGProfileSoftGREProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileSoftGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileSoftGREProfileList() *WSGProfileSoftGREProfileList {
	m := new(WSGProfileSoftGREProfileList)
	return m
}

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

func NewWSGProfileTrafficClassProfileList() *WSGProfileTrafficClassProfileList {
	m := new(WSGProfileTrafficClassProfileList)
	return m
}

// WSGProfileTtgCommonSetting
//
// Hosted AAA server RADIUS settings & PLMN ID settings
type WSGProfileTtgCommonSetting struct {
	// MobileCountryCode
	// Mobile country code
	// Constraints:
	//    - max:3
	//    - min:3
	MobileCountryCode *string `json:"mobileCountryCode,omitempty" validate:"max=3,min=3"`

	// MobileNetworkCode
	// Mobile network code
	// Constraints:
	//    - max:3
	//    - min:2
	MobileNetworkCode *string `json:"mobileNetworkCode,omitempty" validate:"max=3,min=2"`
}

func NewWSGProfileTtgCommonSetting() *WSGProfileTtgCommonSetting {
	m := new(WSGProfileTtgCommonSetting)
	return m
}

type WSGProfileTtgpdgApnForwardingRealm struct {
	// Apn
	// the forwarding policy APN, if apnType is NIOI, APN Example : internet-v4.mnc111.mcc222.gprs
	Apn *string `json:"apn,omitempty"`

	// ApnType
	// type of the forwarding policy APN.
	// Constraints:
	//    - oneof:[NI,NIOI]
	ApnType *string `json:"apnType,omitempty" validate:"oneof=NI NIOI"`

	// RouteType
	// routeType of the forwarding policy APN.
	// Constraints:
	//    - oneof:[GTPv1,GTPv2,PDG]
	RouteType *string `json:"routeType,omitempty" validate:"oneof=GTPv1 GTPv2 PDG"`
}

func NewWSGProfileTtgpdgApnForwardingRealm() *WSGProfileTtgpdgApnForwardingRealm {
	m := new(WSGProfileTtgpdgApnForwardingRealm)
	return m
}

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
	ApnFormat2GGSN *string `json:"apnFormat2GGSN,omitempty" validate:"oneof=DNS String"`

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

func NewWSGProfileTtgpdgProfile() *WSGProfileTtgpdgProfile {
	m := new(WSGProfileTtgpdgProfile)
	return m
}

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

type WSGProfileTtgpdgProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileTtgpdgProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileTtgpdgProfileList() *WSGProfileTtgpdgProfileList {
	m := new(WSGProfileTtgpdgProfileList)
	return m
}

type WSGProfileUpdateL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*WSGProfileDataPlaneL3RoamingData `json:"dataPlanes,omitempty"`
}

func NewWSGProfileUpdateL3RoamingConfig() *WSGProfileUpdateL3RoamingConfig {
	m := new(WSGProfileUpdateL3RoamingConfig)
	return m
}

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

type WSGProfileUpdateRogueApPolicy struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Rules []*WSGProfileRogueApRuleList `json:"rules,omitempty"`
}

func NewWSGProfileUpdateRogueApPolicy() *WSGProfileUpdateRogueApPolicy {
	m := new(WSGProfileUpdateRogueApPolicy)
	return m
}

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
	DefaultAction *string `json:"defaultAction,omitempty" validate:"oneof=BLOCK ALLOW"`

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

func NewWSGProfileUserTrafficProfile() *WSGProfileUserTrafficProfile {
	m := new(WSGProfileUserTrafficProfile)
	return m
}

type WSGProfileUserTrafficProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGProfileUserTrafficProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileUserTrafficProfileList() *WSGProfileUserTrafficProfileList {
	m := new(WSGProfileUserTrafficProfileList)
	return m
}

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

func NewWSGProfileVdpProfile() *WSGProfileVdpProfile {
	m := new(WSGProfileVdpProfile)
	return m
}

// WSGProfileVlanPrecedenceItem
//
// Vlan precedence item
type WSGProfileVlanPrecedenceItem struct {
	// Name
	// Name of the Vlan precedence item
	// Constraints:
	//    - oneof:[AAA,DEVICE,WLAN]
	Name *string `json:"name,omitempty" validate:"oneof=AAA DEVICE WLAN"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

func NewWSGProfileVlanPrecedenceItem() *WSGProfileVlanPrecedenceItem {
	m := new(WSGProfileVlanPrecedenceItem)
	return m
}

type WSGProfileZoneAffinityProfileList struct {
	List []*WSGProfileReturnZoneAffinityProfile `json:"list,omitempty"`
}

func NewWSGProfileZoneAffinityProfileList() *WSGProfileZoneAffinityProfileList {
	m := new(WSGProfileZoneAffinityProfileList)
	return m
}
