package vsz

// API Version: v9_0

type WSGProfileAccountingProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the accounting profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	// Constraints:
	//    - nullable
	RealmMappings []*WSGProfileAcctServiceRealmMapping `json:"realmMappings,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileAccountingProfile() *WSGProfileAccountingProfile {
	m := new(WSGProfileAccountingProfile)
	return m
}

type WSGProfileAccountingProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileAccountingProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileAccountingProfileList() *WSGProfileAccountingProfileList {
	m := new(WSGProfileAccountingProfileList)
	return m
}

// WSGProfileAcctServiceRealmMapping
//
// Accounting service per realm
// Constraints:
//    - nullable
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
	Realm *WSGCommonRealm `json:"realm" validate:"required"`

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
// Constraints:
//    - nullable
type WSGProfileAdvancedOptionContent struct {
	// DhcpOpt43Subcode
	// dhcpOpt43Subcode of the ipsec profile
	// Constraints:
	//    - nullable
	DhcpOpt43Subcode *float64 `json:"dhcpOpt43Subcode,omitempty"`

	// DpdDelay
	// dpdDelay of the ipsec profile
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	FailoverPrimaryCheckInterval *float64 `json:"failoverPrimaryCheckInterval,omitempty"`

	// FailoverRetryInterval
	// Retry Interval of the failover
	// Constraints:
	//    - nullable
	FailoverRetryInterval *float64 `json:"failoverRetryInterval,omitempty"`

	// FailoverRetryPeriod
	// Retry Period of the failover
	// Constraints:
	//    - nullable
	FailoverRetryPeriod *float64 `json:"failoverRetryPeriod,omitempty"`

	// IpcompEnable
	// ipcomp Enable of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[Disabled,Enabled]
	IpcompEnable *string `json:"ipcompEnable,omitempty" validate:"omitempty,oneof=Disabled Enabled"`

	// KeepAliveIntval
	// keepAliveIntval of the ipsec profile
	// Constraints:
	//    - nullable
	KeepAliveIntval *float64 `json:"keepAliveIntval,omitempty"`

	// ReplayWindow
	// replayWindow of the ipsec profile
	// Constraints:
	//    - nullable
	ReplayWindow *float64 `json:"replayWindow,omitempty"`

	// RetryLimit
	// retryLimit of the ipsec profile
	// Constraints:
	//    - nullable
	RetryLimit *float64 `json:"retryLimit,omitempty"`
}

func NewWSGProfileAdvancedOptionContent() *WSGProfileAdvancedOptionContent {
	m := new(WSGProfileAdvancedOptionContent)
	return m
}

type WSGProfileApnRealm struct {
	// DefaultAPN
	// name of the apnForwardingPolicys.
	// Constraints:
	//    - nullable
	DefaultAPN *string `json:"defaultAPN,omitempty"`

	// Realm
	// name of the apnRealm.
	// Constraints:
	//    - nullable
	Realm *string `json:"realm,omitempty"`
}

func NewWSGProfileApnRealm() *WSGProfileApnRealm {
	m := new(WSGProfileApnRealm)
	return m
}

type WSGProfileAuthenticationProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	// Constraints:
	//    - nullable
	GppSuppportEnabled *bool `json:"gppSuppportEnabled,omitempty"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	// Constraints:
	//    - nullable
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// Id
	// Identifier of the authentication profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IsContainDirectoryService
	// Realm based authentication service mappings contains LDAP or AD service type
	// Constraints:
	//    - nullable
	IsContainDirectoryService *bool `json:"isContainDirectoryService,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	// Constraints:
	//    - nullable
	RealmMappings []*WSGProfileRealmAuthServiceMapping `json:"realmMappings,omitempty" validate:"omitempty,dive"`

	// TtgCommonSetting
	// Constraints:
	//    - nullable
	TtgCommonSetting *WSGProfileTtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewWSGProfileAuthenticationProfile() *WSGProfileAuthenticationProfile {
	m := new(WSGProfileAuthenticationProfile)
	return m
}

type WSGProfileAuthenticationProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileAuthenticationProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileAuthenticationProfileList() *WSGProfileAuthenticationProfileList {
	m := new(WSGProfileAuthenticationProfileList)
	return m
}

type WSGProfileBaseServiceInfoList struct {
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
	List []*WSGCommonBaseServiceInfo `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBaseServiceInfoList() *WSGProfileBaseServiceInfoList {
	m := new(WSGProfileBaseServiceInfoList)
	return m
}

type WSGProfileBlockClient struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// ZoneId
	// Zone Id of the Block Client for clone in System Domain
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBlockClient() *WSGProfileBlockClient {
	m := new(WSGProfileBlockClient)
	return m
}

type WSGProfileBlockClientList struct {
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
	List []*WSGProfileBlockClientListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBlockClientList() *WSGProfileBlockClientList {
	m := new(WSGProfileBlockClientList)
	return m
}

type WSGProfileBlockClientListType struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Mac
	// Constraints:
	//    - nullable
	Mac *WSGCommonMac `json:"mac,omitempty"`

	// ModifiedDateTime
	// Date blocked of the Block Client
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Modifier blocked of the Block Client
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// ZoneId
	// Zone Id of the Block Client for clone in System Domain
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBlockClientListType() *WSGProfileBlockClientListType {
	m := new(WSGProfileBlockClientListType)
	return m
}

type WSGProfileBlockedPort struct {
	// Port
	// port of the Blocked Port
	// Constraints:
	//    - nullable
	Port *string `json:"port,omitempty"`

	// Protocol
	// Protocol of the Blocked Port
	// Constraints:
	//    - nullable
	//    - oneof:[BOTH,TCP,UDP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=BOTH TCP UDP"`
}

func NewWSGProfileBlockedPort() *WSGProfileBlockedPort {
	m := new(WSGProfileBlockedPort)
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
	// Constraints:
	//    - nullable
	BonjourFencingRuleMappingList []*WSGProfileBonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty" validate:"omitempty,dive"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Bonjour Fencing Policy id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBonjourFencingPolicy() *WSGProfileBonjourFencingPolicy {
	m := new(WSGProfileBonjourFencingPolicy)
	return m
}

type WSGProfileBonjourFencingPolicyList struct {
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
	List []*WSGProfileBonjourFencingPolicy `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBonjourFencingPolicyList() *WSGProfileBonjourFencingPolicyList {
	m := new(WSGProfileBonjourFencingPolicyList)
	return m
}

type WSGProfileBonjourFencingRule struct {
	// ClosestAp
	// Constraints:
	//    - nullable
	ClosestAp *WSGCommonMac `json:"closestAp,omitempty"`

	// CustomServiceName
	// Constraints:
	//    - nullable
	CustomServiceName *WSGCommonNormalName `json:"customServiceName,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DeviceMacList
	// Specify the device list providing Bonjour Service
	// Constraints:
	//    - nullable
	DeviceMacList []*WSGProfileBonjourFencingRuleDeviceMac `json:"deviceMacList,omitempty" validate:"omitempty,dive"`

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
	ServiceType *WSGProfileBridgeService `json:"serviceType" validate:"required"`
}

func NewWSGProfileBonjourFencingRule() *WSGProfileBonjourFencingRule {
	m := new(WSGProfileBonjourFencingRule)
	return m
}

type WSGProfileBonjourFencingRuleDeviceMac struct {
	// Mac
	// Constraints:
	//    - nullable
	Mac *WSGCommonMac `json:"mac,omitempty"`
}

func NewWSGProfileBonjourFencingRuleDeviceMac() *WSGProfileBonjourFencingRuleDeviceMac {
	m := new(WSGProfileBonjourFencingRuleDeviceMac)
	return m
}

type WSGProfileBonjourFencingRuleMapping struct {
	// CustomServiceName
	// Constraints:
	//    - nullable
	CustomServiceName *WSGCommonNormalName `json:"customServiceName,omitempty"`

	// CustomStringList
	// The array of mdns string
	// Constraints:
	//    - nullable
	CustomStringList []string `json:"customStringList,omitempty" validate:"omitempty,dive"`

	// ServiceType
	// Constraints:
	//    - nullable
	ServiceType *WSGProfileBridgeService `json:"serviceType,omitempty"`
}

func NewWSGProfileBonjourFencingRuleMapping() *WSGProfileBonjourFencingRuleMapping {
	m := new(WSGProfileBonjourFencingRuleMapping)
	return m
}

type WSGProfileBonjourFencingService struct {
	// NeighborApMac
	// Constraints:
	//    - nullable
	NeighborApMac *string `json:"neighborApMac,omitempty"`

	// NeighborApName
	// Constraints:
	//    - nullable
	NeighborApName *string `json:"neighborApName,omitempty"`

	// ServiceType
	// Constraints:
	//    - nullable
	ServiceType *WSGProfileBridgeService `json:"serviceType,omitempty"`

	// SourceType
	// Constraints:
	//    - nullable
	//    - oneof:[UNKNOWN,DIRECT,NEIGHBOR]
	SourceType *string `json:"sourceType,omitempty" validate:"omitempty,oneof=UNKNOWN DIRECT NEIGHBOR"`
}

func NewWSGProfileBonjourFencingService() *WSGProfileBonjourFencingService {
	m := new(WSGProfileBonjourFencingService)
	return m
}

type WSGProfileBonjourFencingStatistic struct {
	// ApMac
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// DroppedPacketsDueToNeighbor
	// Constraints:
	//    - nullable
	DroppedPacketsDueToNeighbor *int `json:"droppedPacketsDueToNeighbor,omitempty"`

	// DroppedPacketsDueToServiceType
	// Constraints:
	//    - nullable
	DroppedPacketsDueToServiceType *int `json:"droppedPacketsDueToServiceType,omitempty"`

	// ForwardedPackets
	// Constraints:
	//    - nullable
	ForwardedPackets *int `json:"forwardedPackets,omitempty"`

	// ServiceList
	// Constraints:
	//    - nullable
	ServiceList []*WSGProfileBonjourFencingService `json:"serviceList,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileBonjourFencingStatistic() *WSGProfileBonjourFencingStatistic {
	m := new(WSGProfileBonjourFencingStatistic)
	return m
}

type WSGProfileBridgeProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileBridgeProfile() *WSGProfileBridgeProfile {
	m := new(WSGProfileBridgeProfile)
	return m
}

type WSGProfileBridgeProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileBridgeProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileBridgeProfileList() *WSGProfileBridgeProfileList {
	m := new(WSGProfileBridgeProfileList)
	return m
}

// WSGProfileBridgeService
//
// Bonjour Service Type
// Constraints:
//    - nullable
//    - oneof:[AIRDISK,AIRPLAY,AIRPORT_MANAGEMENT,AIRPRINT,AIRTUNES,APPLE_FILE_SHARING,APPLE_MOBILE_DEVICES,APPLETV,ICLOUD_SYNC,ITUNES_REMOTE,ITUNES_SHARING,OPEN_DIRECTORY_MASTER,OPTICAL_DISK_SHARING,SCREEN_SHARING,SECURE_FILE_SHARING,SECURE_SHELL,WWW_HTTP,WWW_HTTPS,WORKGROUP_MANAGER,XGRID,GOOGLE_CHROMECAST,OTHER]
type WSGProfileBridgeService string

func NewWSGProfileBridgeService() *WSGProfileBridgeService {
	m := new(WSGProfileBridgeService)
	return m
}

type WSGProfileBulkBlockClient struct {
	// BlockClientList
	// Constraints:
	//    - nullable
	BlockClientList []*WSGProfileBulkBlockClientBlockClientListType `json:"blockClientList,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
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

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`

	// ZoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileBulkBlockClientBlockClientListType() *WSGProfileBulkBlockClientBlockClientListType {
	m := new(WSGProfileBulkBlockClientBlockClientListType)
	return m
}

type WSGProfileClientIsolationEntry struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Ip
	// Client Entry ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

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
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Client Isolation Whitelist id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - required
	Whitelist []*WSGProfileClientIsolationEntry `json:"whitelist" validate:"required,dive"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileClientIsolationWhitelist() *WSGProfileClientIsolationWhitelist {
	m := new(WSGProfileClientIsolationWhitelist)
	return m
}

type WSGProfileClientIsolationWhitelistArray struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

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
	List []*WSGProfileClientIsolationWhitelist `json:"list,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileClientIsolationWhitelistArray() *WSGProfileClientIsolationWhitelistArray {
	m := new(WSGProfileClientIsolationWhitelistArray)
	return m
}

// WSGProfileCmProtocolOptionContent
//
// Certificate Management Protocol Option
// Constraints:
//    - nullable
type WSGProfileCmProtocolOptionContent struct {
	// CmpDhcpOpt43Subcode
	// Certificate Management Protocol dhcpOpt43Subcode
	// Constraints:
	//    - required
	CmpDhcpOpt43Subcode *float64 `json:"cmpDhcpOpt43Subcode" validate:"required"`

	// CmpDhcpOpt43SubcodeRecipient
	// Certificate Management Protocol dhcpOpt43SubcodeRecipient
	// Constraints:
	//    - nullable
	CmpDhcpOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`

	// CmpRecipient
	// Certificate Management Protocol Recipient
	// Constraints:
	//    - nullable
	CmpRecipient *string `json:"cmpRecipient,omitempty"`

	// CmpServerAddr
	// Certificate Management Protocol Server addr
	// Constraints:
	//    - nullable
	CmpServerAddr *string `json:"cmpServerAddr,omitempty"`

	// CmpServerPath
	// Certificate Management Protocol Server Path
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary Gateway
	// Constraints:
	//    - nullable
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

func NewWSGProfileCoreNetworkGateway() *WSGProfileCoreNetworkGateway {
	m := new(WSGProfileCoreNetworkGateway)
	return m
}

type WSGProfileCreateAccountingProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// RealmMappings
	// Accounting service per realm
	// Constraints:
	//    - nullable
	RealmMappings []*WSGProfileAcctServiceRealmMapping `json:"realmMappings,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileCreateAccountingProfile() *WSGProfileCreateAccountingProfile {
	m := new(WSGProfileCreateAccountingProfile)
	return m
}

type WSGProfileCreateAuthenticationProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	// Constraints:
	//    - required
	//    - default:false
	GppSuppportEnabled *bool `json:"gppSuppportEnabled" validate:"required"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	// Constraints:
	//    - nullable
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// RealmMappings
	// Realm based authentication service mappings
	// Constraints:
	//    - nullable
	RealmMappings []*WSGProfileRealmAuthServiceMapping `json:"realmMappings,omitempty" validate:"omitempty,dive"`

	// TtgCommonSetting
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	BonjourFencingRuleMappingList []*WSGProfileBonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileCreateBonjourFencingPolicy() *WSGProfileCreateBonjourFencingPolicy {
	m := new(WSGProfileCreateBonjourFencingPolicy)
	return m
}

type WSGProfileCreateBridgeProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
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

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

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
	// Description
	// Constraints:
	//    - nullable
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
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PoolEndIp
	// Constraints:
	//    - required
	PoolEndIp *WSGCommonIpAddress `json:"poolEndIp" validate:"required"`

	// PoolStartIp
	// Constraints:
	//    - required
	PoolStartIp *WSGCommonIpAddress `json:"poolStartIp" validate:"required"`

	// PrimaryDnsIp
	// Constraints:
	//    - nullable
	PrimaryDnsIp *WSGCommonIpAddress `json:"primaryDnsIp,omitempty"`

	// SecondaryDnsIp
	// Constraints:
	//    - nullable
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
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PrimaryIp
	// Primary ip of DNS server service
	// Constraints:
	//    - required
	PrimaryIp *string `json:"primaryIp" validate:"required"`

	// SecondaryIp
	// Secondary ip of DNS server service
	// Constraints:
	//    - nullable
	SecondaryIp *string `json:"secondaryIp,omitempty"`

	// TertiaryIp
	// Tertiary ip of DNS server service
	// Constraints:
	//    - nullable
	TertiaryIp *string `json:"tertiaryIp,omitempty"`
}

func NewWSGProfileCreateDnsServerProfile() *WSGProfileCreateDnsServerProfile {
	m := new(WSGProfileCreateDnsServerProfile)
	return m
}

type WSGProfileCreateFirewallProfile struct {
	// AppPolicyId
	// Application Policy
	// Constraints:
	//    - nullable
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DevicePolicyId
	// Device Policy
	// Constraints:
	//    - nullable
	DevicePolicyId *string `json:"devicePolicyId,omitempty"`

	// DomainId
	// Domain Id of The Firewall Profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy
	// Constraints:
	//    - nullable
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	// L3AccessControlPolicyId
	// L3 Access Control Policy
	// Constraints:
	//    - nullable
	L3AccessControlPolicyId *string `json:"l3AccessControlPolicyId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`

	// UrlFilteringPolicyId
	// Url Filtering Policy
	// Constraints:
	//    - nullable
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileCreateFirewallProfile() *WSGProfileCreateFirewallProfile {
	m := new(WSGProfileCreateFirewallProfile)
	return m
}

type WSGProfileCreateIpsecProfile struct {
	// AdvancedOption
	// Constraints:
	//    - nullable
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"omitempty,oneof=PresharedKey Certificate"`

	// CmProtocolOption
	// Constraints:
	//    - nullable
	CmProtocolOption *WSGProfileCmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	// Constraints:
	//    - required
	EspRekeyTime *float64 `json:"espRekeyTime" validate:"required"`

	// EspRekeyTimeUnit
	// Constraints:
	//    - nullable
	EspRekeyTimeUnit *WSGCommonTimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	// EspSecurityAssociation
	// Constraints:
	//    - nullable
	EspSecurityAssociation *WSGProfileEspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	// Constraints:
	//    - required
	IkeRekeyTime *float64 `json:"ikeRekeyTime" validate:"required"`

	// IkeRekeyTimeUnit
	// Constraints:
	//    - nullable
	IkeRekeyTimeUnit *WSGCommonTimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	// IkeSecurityAssociation
	// Constraints:
	//    - nullable
	IkeSecurityAssociation *WSGProfileIkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	// IpMode
	// Constraints:
	//    - required
	IpMode *WSGProfileIpMode `json:"ipMode" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	// Constraints:
	//    - nullable
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	// Constraints:
	//    - nullable
	ServerAddr *string `json:"serverAddr,omitempty"`

	// TunnelMode
	// Tunnel mode of IPsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[SOFT_GRE,RUCKUS_GRE]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"omitempty,oneof=SOFT_GRE RUCKUS_GRE"`
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

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileCreateL2oGREProfile() *WSGProfileCreateL2oGREProfile {
	m := new(WSGProfileCreateL2oGREProfile)
	return m
}

type WSGProfileCreateL3AccessControlPolicy struct {
	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=BLOCK ALLOW"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// L3AclRuleList
	// L3 access control list
	// Constraints:
	//    - nullable
	L3AclRuleList []*WSGProfileL3AclRuleList `json:"l3AclRuleList,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileCreateL3AccessControlPolicy() *WSGProfileCreateL3AccessControlPolicy {
	m := new(WSGProfileCreateL3AccessControlPolicy)
	return m
}

type WSGProfileCreatePrecedenceProfile struct {
	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	// Constraints:
	//    - nullable
	RateLimitingPrecedence []*WSGProfileRateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty" validate:"omitempty,dive"`

	// VlanPrecedence
	// vlan precedence
	// Constraints:
	//    - nullable
	VlanPrecedence []*WSGProfileVlanPrecedenceItem `json:"vlanPrecedence,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileCreatePrecedenceProfile() *WSGProfileCreatePrecedenceProfile {
	m := new(WSGProfileCreatePrecedenceProfile)
	return m
}

type WSGProfileCreateRestrictedApAccessProfile struct {
	// BlockedPortList
	// Blocked Port List
	// Constraints:
	//    - nullable
	BlockedPortList []*WSGProfileBlockedPort `json:"blockedPortList,omitempty" validate:"omitempty,dive"`

	// BlockWellKnownPort
	// Block well known ports
	// Constraints:
	//    - nullable
	//    - default:false
	BlockWellKnownPort *bool `json:"blockWellKnownPort,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// IpAddressWhitelist
	// IP Address Whitelist
	// Constraints:
	//    - nullable
	IpAddressWhitelist []WSGCommonIpAddress `json:"ipAddressWhitelist,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileCreateRestrictedApAccessProfile() *WSGProfileCreateRestrictedApAccessProfile {
	m := new(WSGProfileCreateRestrictedApAccessProfile)
	return m
}

type WSGProfileCreateResultList []*WSGCommonCreateResult

func MakeWSGProfileCreateResultList() WSGProfileCreateResultList {
	m := make(WSGProfileCreateResultList, 0)
	return m
}

type WSGProfileCreateRogueApPolicy struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

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

	// EkahauIp
	// Constraints:
	//    - nullable
	EkahauIp *WSGCommonIpAddress `json:"ekahauIp,omitempty"`

	// EkahauPort
	// Eekahau Location Server Port
	// Constraints:
	//    - nullable
	EkahauPort *int `json:"ekahauPort,omitempty"`

	// Id
	// Constraints:
	//    - nullable
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
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

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

func NewWSGProfileCreateRuckusGREProfile() *WSGProfileCreateRuckusGREProfile {
	m := new(WSGProfileCreateRuckusGREProfile)
	return m
}

type WSGProfileCreateSoftGREProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	// Constraints:
	//    - nullable
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpMode
	// Constraints:
	//    - nullable
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
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	// Constraints:
	//    - required
	PrimaryGateway *string `json:"primaryGateway" validate:"required"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	// Constraints:
	//    - nullable
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

func NewWSGProfileCreateSoftGREProfile() *WSGProfileCreateSoftGREProfile {
	m := new(WSGProfileCreateSoftGREProfile)
	return m
}

type WSGProfileCreateTrafficClassProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName2to64 `json:"name" validate:"required"`

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
	// Constraints:
	//    - nullable
	ApnRealms []*WSGProfileApnRealm `json:"apnRealms,omitempty" validate:"omitempty,dive"`

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

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileCreateTtgpdgProfile() *WSGProfileCreateTtgpdgProfile {
	m := new(WSGProfileCreateTtgpdgProfile)
	return m
}

type WSGProfileCreateUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	// Constraints:
	//    - nullable
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=BLOCK ALLOW"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimiting
	// Constraints:
	//    - nullable
	DownlinkRateLimiting *WSGProfileDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// IpAclRules
	// Traffic access control list
	// Constraints:
	//    - nullable
	IpAclRules []*WSGProfileIpAclRules `json:"ipAclRules,omitempty" validate:"omitempty,dive"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// QmAppPolicyId
	// Application Policy UUID
	// Constraints:
	//    - nullable
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	// UplinkRateLimiting
	// Constraints:
	//    - nullable
	UplinkRateLimiting *WSGProfileUplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Key
	// Data plane key
	// Constraints:
	//    - required
	Key *string `json:"key" validate:"required"`

	// Name
	// DP name
	// Constraints:
	//    - nullable
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

func NewWSGProfileDataPlaneL3RoamingData() *WSGProfileDataPlaneL3RoamingData {
	m := new(WSGProfileDataPlaneL3RoamingData)
	return m
}

type WSGProfileDeleteBulkAccountingProfile struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkAccountingProfile() *WSGProfileDeleteBulkAccountingProfile {
	m := new(WSGProfileDeleteBulkAccountingProfile)
	return m
}

type WSGProfileDeleteBulkAuthenticationProfile struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkAuthenticationProfile() *WSGProfileDeleteBulkAuthenticationProfile {
	m := new(WSGProfileDeleteBulkAuthenticationProfile)
	return m
}

type WSGProfileDeleteBulkPrecedenceProfile struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkPrecedenceProfile() *WSGProfileDeleteBulkPrecedenceProfile {
	m := new(WSGProfileDeleteBulkPrecedenceProfile)
	return m
}

type WSGProfileDeleteBulkUserTrafficProfile struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGProfileDeleteBulkUserTrafficProfile() *WSGProfileDeleteBulkUserTrafficProfile {
	m := new(WSGProfileDeleteBulkUserTrafficProfile)
	return m
}

type WSGProfileDhcpOption82 struct {
	// DhcpOption82Enabled
	// Enable DHCP Option 82
	// Constraints:
	//    - nullable
	DhcpOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// Subopt1Enabled
	// Enable subopt-1
	// Constraints:
	//    - nullable
	Subopt1Enabled *bool `json:"subopt1Enabled,omitempty"`

	// Subopt1Format
	// Subopt-1 format
	// Constraints:
	//    - nullable
	//    - oneof:[AP_INFO,AP_MAC_hex,AP_MAC_hex_ESSID,AP_INFO_LOCATION]
	Subopt1Format *string `json:"subopt1Format,omitempty" validate:"omitempty,oneof=AP_INFO AP_MAC_hex AP_MAC_hex_ESSID AP_INFO_LOCATION"`

	// Subopt2Enabled
	// Enable subopt-2
	// Constraints:
	//    - nullable
	Subopt2Enabled *bool `json:"subopt2Enabled,omitempty"`

	// Subopt2Format
	// Subopt-2 format
	// Constraints:
	//    - nullable
	//    - oneof:[CLIENT_MAC_hex,CLIENT_MAC_hex_ESSID,AP_MAC_hex,AP_MAC__hex_ESSID]
	Subopt2Format *string `json:"subopt2Format,omitempty" validate:"omitempty,oneof=CLIENT_MAC_hex CLIENT_MAC_hex_ESSID AP_MAC_hex AP_MAC__hex_ESSID"`

	// Subopt150Enabled
	// Subopt-150 with VLAN
	// Constraints:
	//    - nullable
	Subopt150Enabled *bool `json:"subopt150Enabled,omitempty"`

	// Subopt151AreaName
	// Subopt-151 Area Name value
	// Constraints:
	//    - nullable
	Subopt151AreaName *string `json:"subopt151AreaName,omitempty"`

	// Subopt151Enabled
	// Enable subopt-151
	// Constraints:
	//    - nullable
	Subopt151Enabled *bool `json:"subopt151Enabled,omitempty"`

	// Subopt151Format
	// Subopt-151 format
	// Constraints:
	//    - nullable
	//    - oneof:[AREA_NAME,ESSID]
	Subopt151Format *string `json:"subopt151Format,omitempty" validate:"omitempty,oneof=AREA_NAME ESSID"`
}

func NewWSGProfileDhcpOption82() *WSGProfileDhcpOption82 {
	m := new(WSGProfileDhcpOption82)
	return m
}

type WSGProfileDhcpProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

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
	List []*WSGCommonDhcpProfileRef `json:"list,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileDhcpProfileList() *WSGProfileDhcpProfileList {
	m := new(WSGProfileDhcpProfileList)
	return m
}

type WSGProfileDhcpRelayNoRelayTunnel struct {
	// DhcpOption82
	// Constraints:
	//    - nullable
	DhcpOption82 *WSGProfileDhcpOption82 `json:"dhcpOption82,omitempty"`

	// DhcpRelayEnabled
	// Enable DHCP Relay
	// Constraints:
	//    - nullable
	DhcpRelayEnabled *bool `json:"dhcpRelayEnabled,omitempty"`

	// DhcpServer1
	// DHCP Server 1
	// Constraints:
	//    - nullable
	DhcpServer1 *string `json:"dhcpServer1,omitempty"`

	// DhcpServer2
	// DHCP Server 2
	// Constraints:
	//    - nullable
	DhcpServer2 *string `json:"dhcpServer2,omitempty"`

	// RelayBothEnabled
	// Send DHCP requests to both servers simultaneously.
	// Constraints:
	//    - nullable
	RelayBothEnabled *bool `json:"relayBothEnabled,omitempty"`
}

func NewWSGProfileDhcpRelayNoRelayTunnel() *WSGProfileDhcpRelayNoRelayTunnel {
	m := new(WSGProfileDhcpRelayNoRelayTunnel)
	return m
}

type WSGProfileDnsServerProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryIp
	// Primary ip of DNS server service
	// Constraints:
	//    - nullable
	PrimaryIp *string `json:"primaryIp,omitempty"`

	// SecondaryIp
	// Secondary ip of DNS server service
	// Constraints:
	//    - nullable
	SecondaryIp *string `json:"secondaryIp,omitempty"`

	// TertiaryIp
	// Tertiary ip of DNS server service
	// Constraints:
	//    - nullable
	TertiaryIp *string `json:"tertiaryIp,omitempty"`
}

func NewWSGProfileDnsServerProfile() *WSGProfileDnsServerProfile {
	m := new(WSGProfileDnsServerProfile)
	return m
}

type WSGProfileDnsServerProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileDnsServerProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileDnsServerProfileList() *WSGProfileDnsServerProfileList {
	m := new(WSGProfileDnsServerProfileList)
	return m
}

type WSGProfileDownlinkRateLimiting struct {
	// DownlinkRateLimitingBps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	DownlinkRateLimitingBps *string `json:"downlinkRateLimitingBps,omitempty"`

	// DownlinkRateLimitingEnabled
	// Downlink rate limiting enabled or disabled
	// Constraints:
	//    - nullable
	//    - default:false
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`
}

func NewWSGProfileDownlinkRateLimiting() *WSGProfileDownlinkRateLimiting {
	m := new(WSGProfileDownlinkRateLimiting)
	return m
}

// WSGProfileEspProposal
//
// EspProposal based ipsec service mappings
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
type WSGProfileEspSecurityAssociationContent struct {
	// EspProposals
	// espProposal list of the ipsec profile
	// Constraints:
	//    - nullable
	EspProposals []*WSGProfileEspProposal `json:"espProposals,omitempty" validate:"omitempty,dive"`

	// EspProposalType
	// espProposal Type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[Default,Specific]
	EspProposalType *string `json:"espProposalType,omitempty" validate:"omitempty,oneof=Default Specific"`
}

func NewWSGProfileEspSecurityAssociationContent() *WSGProfileEspSecurityAssociationContent {
	m := new(WSGProfileEspSecurityAssociationContent)
	return m
}

type WSGProfileFirewallProfile struct {
	// AppPolicyId
	// Application Policy
	// Constraints:
	//    - nullable
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DevicePolicyId
	// Device Policy
	// Constraints:
	//    - nullable
	DevicePolicyId *string `json:"devicePolicyId,omitempty"`

	// DomainId
	// Domain Id of The Firewall Profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// FactoryDefault
	// Whether the proFirewall Profile is factory default or not
	// Constraints:
	//    - nullable
	FactoryDefault *bool `json:"factoryDefault,omitempty"`

	// Id
	// Firewall Profile id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy
	// Constraints:
	//    - nullable
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	// L3AccessControlPolicyId
	// L3 Access Control Policy
	// Constraints:
	//    - nullable
	L3AccessControlPolicyId *string `json:"l3AccessControlPolicyId,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`

	// UrlFilteringPolicyId
	// Url Filtering Policy
	// Constraints:
	//    - nullable
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileFirewallProfile() *WSGProfileFirewallProfile {
	m := new(WSGProfileFirewallProfile)
	return m
}

type WSGProfileFirewallProfileArray struct {
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
	List []*WSGProfileFirewallProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileFirewallProfileArray() *WSGProfileFirewallProfileArray {
	m := new(WSGProfileFirewallProfileArray)
	return m
}

type WSGProfileFlexiVpnProfile struct {
	// DestinationZoneAffinityId
	// Flexi-VPN Profile ID (Destination)
	// Constraints:
	//    - nullable
	DestinationZoneAffinityId *string `json:"destinationZoneAffinityId,omitempty"`

	// DestinationZoneAffinityName
	// Flexi-VPN Profile (Destination)
	// Constraints:
	//    - nullable
	DestinationZoneAffinityName *string `json:"destinationZoneAffinityName,omitempty"`

	// DomainId
	// Domain ID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// SourceZoneAffinityId
	// Zone Affinity Profile ID (Source)
	// Constraints:
	//    - nullable
	SourceZoneAffinityId *string `json:"sourceZoneAffinityId,omitempty"`

	// SourceZoneAffinityName
	// Zone Affinity Profile (Source)
	// Constraints:
	//    - nullable
	SourceZoneAffinityName *string `json:"sourceZoneAffinityName,omitempty"`

	// WlanId
	// Wlan ID
	// Constraints:
	//    - nullable
	WlanId *string `json:"wlanId,omitempty"`

	// WlanName
	// Wlan name
	// Constraints:
	//    - nullable
	WlanName *string `json:"wlanName,omitempty"`

	// ZoneId
	// Zone ID
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`

	// ZoneName
	// Zone name
	// Constraints:
	//    - nullable
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGProfileFlexiVpnProfile() *WSGProfileFlexiVpnProfile {
	m := new(WSGProfileFlexiVpnProfile)
	return m
}

type WSGProfileFlexiVpnProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

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
	List []*WSGProfileFlexiVpnProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileFlexiVpnProfileList() *WSGProfileFlexiVpnProfileList {
	m := new(WSGProfileFlexiVpnProfileList)
	return m
}

type WSGProfileGetL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	// Constraints:
	//    - nullable
	DataPlanes []*WSGProfileDataPlaneL3RoamingData `json:"dataPlanes,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileGetL3RoamingConfig() *WSGProfileGetL3RoamingConfig {
	m := new(WSGProfileGetL3RoamingConfig)
	return m
}

type WSGProfileHs20FriendlyName struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonLanguageName `json:"language" validate:"required"`

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
	// Certificate
	// Constraints:
	//    - nullable
	Certificate *WSGCommonGenericRef `json:"certificate,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileHs20Operator() *WSGProfileHs20Operator {
	m := new(WSGProfileHs20Operator)
	return m
}

type WSGProfileHs20OperatorList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileHs20Operator `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileHs20OperatorList() *WSGProfileHs20OperatorList {
	m := new(WSGProfileHs20OperatorList)
	return m
}

type WSGProfileHs20Provider struct {
	// Accountings
	// Accountings
	// Constraints:
	//    - nullable
	Accountings []*WSGProfileProviderAccounting `json:"accountings,omitempty" validate:"omitempty,dive"`

	// Authentications
	// Authentications
	// Constraints:
	//    - nullable
	Authentications []*WSGProfileProviderAuthentication `json:"authentications,omitempty" validate:"omitempty,dive"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// HomeOis
	// Home OIs
	// Constraints:
	//    - nullable
	HomeOis []*WSGProfileProviderHomeOIs `json:"homeOis,omitempty" validate:"omitempty,dive"`

	// Id
	// Identifier of the Hotspot 2.0 identity provider profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Osu
	// Constraints:
	//    - nullable
	Osu *WSGProfileProviderOnlineSignup `json:"osu,omitempty"`

	// Plmns
	// PLMNs
	// Constraints:
	//    - nullable
	Plmns []*WSGProfileProviderPLMN `json:"plmns,omitempty" validate:"omitempty,dive"`

	// Realms
	// Realms
	// Constraints:
	//    - nullable
	Realms []*WSGProfileProviderRealm `json:"realms,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileHs20Provider() *WSGProfileHs20Provider {
	m := new(WSGProfileHs20Provider)
	return m
}

type WSGProfileHs20ProviderList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileHs20Provider `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileHs20ProviderList() *WSGProfileHs20ProviderList {
	m := new(WSGProfileHs20ProviderList)
	return m
}

// WSGProfileIkeProposal
//
// IkeProposal based ipsec service mappings
// Constraints:
//    - nullable
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
	//    - nullable
	//    - oneof:[UseIntegrityALG,PRF_MD5,PRF_SHA1,PRF_AES_CBC,PRF_AES_MAC,PRF_SHA256,PRF_SHA384,PRF_SHA512]
	PrfAlg *string `json:"prfAlg,omitempty" validate:"omitempty,oneof=UseIntegrityALG PRF_MD5 PRF_SHA1 PRF_AES_CBC PRF_AES_MAC PRF_SHA256 PRF_SHA384 PRF_SHA512"`
}

func NewWSGProfileIkeProposal() *WSGProfileIkeProposal {
	m := new(WSGProfileIkeProposal)
	return m
}

// WSGProfileIkeSecurityAssociationContent
//
// ikeProposal Security Association Content
// Constraints:
//    - nullable
type WSGProfileIkeSecurityAssociationContent struct {
	// IkeProposals
	// ikeProposal list of the ipsec profile
	// Constraints:
	//    - nullable
	IkeProposals []*WSGProfileIkeProposal `json:"ikeProposals,omitempty" validate:"omitempty,dive"`

	// IkeProposalType
	// ikeProposal Type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[Default,Specific]
	IkeProposalType *string `json:"ikeProposalType,omitempty" validate:"omitempty,oneof=Default Specific"`
}

func NewWSGProfileIkeSecurityAssociationContent() *WSGProfileIkeSecurityAssociationContent {
	m := new(WSGProfileIkeSecurityAssociationContent)
	return m
}

type WSGProfileIpAclRules struct {
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

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestinationIp
	// Subnet network address or ip address of destination IP.
	// Constraints:
	//    - nullable
	DestinationIp *string `json:"destinationIp,omitempty"`

	// DestinationIpMask
	// Subnet mask of destination IP
	// Constraints:
	//    - nullable
	DestinationIpMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIpV6
	// Destination IPv6 Address.
	// Constraints:
	//    - nullable
	DestinationIpV6 *string `json:"destinationIpV6,omitempty"`

	// DestinationMaxPort
	// The maxinum port of destination port range.
	// Constraints:
	//    - nullable
	DestinationMaxPort *int `json:"destinationMaxPort,omitempty"`

	// DestinationMinPort
	// The mininum port of destination port range.
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting
	// Constraints:
	//    - nullable
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// EnableDestinationIpSubnet
	// Destination IP subnet enabled or disabled
	// Constraints:
	//    - nullable
	EnableDestinationIpSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	// Constraints:
	//    - nullable
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	// Constraints:
	//    - nullable
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIpSubnet
	// Source IP subnet enabled or disabled
	// Constraints:
	//    - nullable
	EnableSourceIpSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	// Constraints:
	//    - nullable
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP UDPLITE ICMP_ICMPV4 ICMPV6 IGMP ESP AH SCTP CUSTOM"`

	// SourceIp
	// Subnet network address or ip address of source IP.
	// Constraints:
	//    - nullable
	SourceIp *string `json:"sourceIp,omitempty"`

	// SourceIpMask
	// Subnet mask of source IP
	// Constraints:
	//    - nullable
	SourceIpMask *string `json:"sourceIpMask,omitempty"`

	// SourceIpV6
	// Source IPv6 Address.
	// Constraints:
	//    - nullable
	SourceIpV6 *string `json:"sourceIpV6,omitempty"`

	// SourceMaxPort
	// The maxinum port of source port range.
	// Constraints:
	//    - nullable
	SourceMaxPort *int `json:"sourceMaxPort,omitempty"`

	// SourceMinPort
	// The minunum port of source port range.
	// Constraints:
	//    - nullable
	SourceMinPort *int `json:"sourceMinPort,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled
	// Constraints:
	//    - nullable
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting
	// Constraints:
	//    - nullable
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
	// AdvancedOption
	// Constraints:
	//    - nullable
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"omitempty,oneof=PresharedKey Certificate"`

	// CmProtocolOption
	// Constraints:
	//    - nullable
	CmProtocolOption *WSGProfileCmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	// Constraints:
	//    - nullable
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Constraints:
	//    - nullable
	EspRekeyTimeUnit *WSGCommonTimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	// EspSecurityAssociation
	// Constraints:
	//    - nullable
	EspSecurityAssociation *WSGProfileEspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	// Constraints:
	//    - nullable
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Constraints:
	//    - nullable
	IkeRekeyTimeUnit *WSGCommonTimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	// IkeSecurityAssociation
	// Constraints:
	//    - nullable
	IkeSecurityAssociation *WSGProfileIkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	// IpMode
	// Constraints:
	//    - nullable
	IpMode *WSGProfileIpMode `json:"ipMode,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	// Constraints:
	//    - nullable
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	// Constraints:
	//    - nullable
	ServerAddr *string `json:"serverAddr,omitempty"`

	// TunnelMode
	// Tunnel mode of IPsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[SOFT_GRE,RUCKUS_GRE]
	TunnelMode *string `json:"tunnelMode,omitempty" validate:"omitempty,oneof=SOFT_GRE RUCKUS_GRE"`
}

func NewWSGProfileIpsecProfile() *WSGProfileIpsecProfile {
	m := new(WSGProfileIpsecProfile)
	return m
}

type WSGProfileIpsecProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileIpsecProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileIpsecProfileList() *WSGProfileIpsecProfileList {
	m := new(WSGProfileIpsecProfileList)
	return m
}

type WSGProfileL2oGREProfile struct {
	// CoreNetworkGateway
	// Constraints:
	//    - nullable
	CoreNetworkGateway *WSGProfileCoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileL2oGREProfile() *WSGProfileL2oGREProfile {
	m := new(WSGProfileL2oGREProfile)
	return m
}

type WSGProfileL2oGREProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileL2oGREProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileL2oGREProfileList() *WSGProfileL2oGREProfileList {
	m := new(WSGProfileL2oGREProfileList)
	return m
}

type WSGProfileL3AccessControlPolicy struct {
	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=BLOCK ALLOW"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// ID of the L3 Access Control Policy
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// L3AclRuleList
	// L3 access control list
	// Constraints:
	//    - nullable
	L3AclRuleList []*WSGProfileL3AclRuleList `json:"l3AclRuleList,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileL3AccessControlPolicy() *WSGProfileL3AccessControlPolicy {
	m := new(WSGProfileL3AccessControlPolicy)
	return m
}

type WSGProfileL3AccessControlPolicyArray struct {
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
	List []*WSGProfileL3AccessControlPolicy `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileL3AccessControlPolicyArray() *WSGProfileL3AccessControlPolicyArray {
	m := new(WSGProfileL3AccessControlPolicyArray)
	return m
}

type WSGProfileL3AclRuleList struct {
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

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestinationIp
	// Subnet network address or ip address of destination IP.
	// Constraints:
	//    - nullable
	DestinationIp *string `json:"destinationIp,omitempty"`

	// DestinationIpMask
	// Subnet mask of destination IP
	// Constraints:
	//    - nullable
	DestinationIpMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIpV6
	// Destination IPv6 Address.
	// Constraints:
	//    - nullable
	DestinationIpV6 *string `json:"destinationIpV6,omitempty"`

	// DestinationMaxPort
	// The maxinum port of destination port range.
	// Constraints:
	//    - nullable
	DestinationMaxPort *int `json:"destinationMaxPort,omitempty"`

	// DestinationMinPort
	// The mininum port of destination port range.
	// Constraints:
	//    - nullable
	DestinationMinPort *int `json:"destinationMinPort,omitempty"`

	// Direction
	// The direction of traffic access control.
	// Constraints:
	//    - nullable
	//    - default:'INBOUND'
	//    - oneof:[INBOUND,OUTBOUND,DUAL]
	Direction *string `json:"direction,omitempty" validate:"omitempty,oneof=INBOUND OUTBOUND DUAL"`

	// EnableDestinationIpSubnet
	// Destination IP subnet enabled or disabled
	// Constraints:
	//    - nullable
	EnableDestinationIpSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	// Constraints:
	//    - nullable
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	// Constraints:
	//    - nullable
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIpSubnet
	// Source IP subnet enabled or disabled
	// Constraints:
	//    - nullable
	EnableSourceIpSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	// Constraints:
	//    - nullable
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP UDPLITE ICMP_ICMPV4 ICMPV6 IGMP ESP AH SCTP CUSTOM"`

	// SourceIp
	// Subnet network address or ip address of source IP.
	// Constraints:
	//    - nullable
	SourceIp *string `json:"sourceIp,omitempty"`

	// SourceIpMask
	// Subnet mask of source IP
	// Constraints:
	//    - nullable
	SourceIpMask *string `json:"sourceIpMask,omitempty"`

	// SourceIpV6
	// Source IPv6 Address.
	// Constraints:
	//    - nullable
	SourceIpV6 *string `json:"sourceIpV6,omitempty"`

	// SourceMaxPort
	// The maxinum port of source port range.
	// Constraints:
	//    - nullable
	SourceMaxPort *int `json:"sourceMaxPort,omitempty"`

	// SourceMinPort
	// The minunum port of source port range.
	// Constraints:
	//    - nullable
	SourceMinPort *int `json:"sourceMinPort,omitempty"`
}

func NewWSGProfileL3AclRuleList() *WSGProfileL3AclRuleList {
	m := new(WSGProfileL3AclRuleList)
	return m
}

type WSGProfileLbsProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Password
	// Password
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// Port
	// LBS port
	// Constraints:
	//    - nullable
	Port *int `json:"port,omitempty"`

	// Url
	// LBS url
	// Constraints:
	//    - nullable
	Url *string `json:"url,omitempty"`

	// Venue
	// Venue
	// Constraints:
	//    - nullable
	Venue *string `json:"venue,omitempty"`
}

func NewWSGProfileLbsProfile() *WSGProfileLbsProfile {
	m := new(WSGProfileLbsProfile)
	return m
}

type WSGProfileLbsProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileLbsProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileLbsProfileList() *WSGProfileLbsProfileList {
	m := new(WSGProfileLbsProfileList)
	return m
}

type WSGProfileModifyAccountingProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// ID of Accounting Profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	// Constraints:
	//    - nullable
	RealmMappings []*WSGProfileAcctServiceRealmMapping `json:"realmMappings,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileModifyAccountingProfile() *WSGProfileModifyAccountingProfile {
	m := new(WSGProfileModifyAccountingProfile)
	return m
}

type WSGProfileModifyAuthenticationProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	// Constraints:
	//    - nullable
	//    - default:false
	GppSuppportEnabled *bool `json:"gppSuppportEnabled,omitempty"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	// Constraints:
	//    - nullable
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// Id
	// ID of Accounting Profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	// Constraints:
	//    - nullable
	RealmMappings []*WSGProfileRealmAuthServiceMapping `json:"realmMappings,omitempty" validate:"omitempty,dive"`

	// TtgCommonSetting
	// Constraints:
	//    - nullable
	TtgCommonSetting *WSGProfileTtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

func NewWSGProfileModifyAuthenticationProfile() *WSGProfileModifyAuthenticationProfile {
	m := new(WSGProfileModifyAuthenticationProfile)
	return m
}

type WSGProfileModifyBlockClient struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mac
	// Constraints:
	//    - nullable
	Mac *WSGCommonMac `json:"mac,omitempty"`
}

func NewWSGProfileModifyBlockClient() *WSGProfileModifyBlockClient {
	m := new(WSGProfileModifyBlockClient)
	return m
}

type WSGProfileModifyBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	// Constraints:
	//    - nullable
	BonjourFencingRuleList []*WSGProfileBonjourFencingRule `json:"bonjourFencingRuleList,omitempty" validate:"omitempty,dive"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	// Constraints:
	//    - nullable
	BonjourFencingRuleMappingList []*WSGProfileBonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyBonjourFencingPolicy() *WSGProfileModifyBonjourFencingPolicy {
	m := new(WSGProfileModifyBonjourFencingPolicy)
	return m
}

type WSGProfileModifyBridgeProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyBridgeProfile() *WSGProfileModifyBridgeProfile {
	m := new(WSGProfileModifyBridgeProfile)
	return m
}

type WSGProfileModifyClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	// Constraints:
	//    - nullable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	// Constraints:
	//    - nullable
	Whitelist []*WSGProfileClientIsolationEntry `json:"whitelist,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileModifyClientIsolationWhitelist() *WSGProfileModifyClientIsolationWhitelist {
	m := new(WSGProfileModifyClientIsolationWhitelist)
	return m
}

type WSGProfileModifyDnsServerProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryIp
	// Primary ip of DNS server service
	// Constraints:
	//    - nullable
	PrimaryIp *string `json:"primaryIp,omitempty"`

	// SecondaryIp
	// Secondary ip of DNS server service
	// Constraints:
	//    - nullable
	SecondaryIp *string `json:"secondaryIp,omitempty"`

	// TertiaryIp
	// Tertiary ip of DNS server service
	// Constraints:
	//    - nullable
	TertiaryIp *string `json:"tertiaryIp,omitempty"`
}

func NewWSGProfileModifyDnsServerProfile() *WSGProfileModifyDnsServerProfile {
	m := new(WSGProfileModifyDnsServerProfile)
	return m
}

type WSGProfileModifyFirewallProfile struct {
	// AppPolicyId
	// Application Policy
	// Constraints:
	//    - nullable
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DevicePolicyId
	// Device Policy
	// Constraints:
	//    - nullable
	DevicePolicyId *string `json:"devicePolicyId,omitempty"`

	// DomainId
	// Domain Id of The Firewall Profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy
	// Constraints:
	//    - nullable
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	// L3AccessControlPolicyId
	// L3 Access Control Policy
	// Constraints:
	//    - nullable
	L3AccessControlPolicyId *string `json:"l3AccessControlPolicyId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`

	// UrlFilteringPolicyId
	// Url Filtering Policy
	// Constraints:
	//    - nullable
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileModifyFirewallProfile() *WSGProfileModifyFirewallProfile {
	m := new(WSGProfileModifyFirewallProfile)
	return m
}

type WSGProfileModifyHS20Operator struct {
	// Certificate
	// Constraints:
	//    - nullable
	Certificate *WSGCommonGenericRef `json:"certificate,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	// Constraints:
	//    - nullable
	DomainNames []WSGCommonWildFQDN `json:"domainNames,omitempty" validate:"omitempty,dive"`

	// FriendlyNames
	// Friendly names
	// Constraints:
	//    - nullable
	FriendlyNames []*WSGProfileHs20FriendlyName `json:"friendlyNames,omitempty" validate:"omitempty,dive"`

	// Id
	// Identifier of the profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyHS20Operator() *WSGProfileModifyHS20Operator {
	m := new(WSGProfileModifyHS20Operator)
	return m
}

// WSGProfileModifyIpAclRules
//
// Traffic access control list
// Constraints:
//    - nullable
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
	//    - nullable
	//    - min:1
	//    - max:255
	CustomProtocol *int `json:"customProtocol,omitempty" validate:"omitempty,gte=1,lte=255"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestinationIp
	// Subnet network address or ip address of destination IP.
	// Constraints:
	//    - nullable
	DestinationIp *string `json:"destinationIp,omitempty"`

	// DestinationIpMask
	// Subnet mask of destination IP
	// Constraints:
	//    - nullable
	DestinationIpMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIpV6
	// Destination IPv6 Address.
	// Constraints:
	//    - nullable
	DestinationIpV6 *string `json:"destinationIpV6,omitempty"`

	// DestinationMaxPort
	// The maxinum port of destination port range.
	// Constraints:
	//    - nullable
	DestinationMaxPort *int `json:"destinationMaxPort,omitempty"`

	// DestinationMinPort
	// The mininum port of destination port range.
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`

	// DownlinkRateLimitingMbps
	// Downlink rate limiting
	// Constraints:
	//    - nullable
	DownlinkRateLimitingMbps *float64 `json:"downlinkRateLimitingMbps,omitempty"`

	// EnableDestinationIpSubnet
	// Destination IP subnet enabled or disabled
	// Constraints:
	//    - nullable
	EnableDestinationIpSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	// Constraints:
	//    - nullable
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	// Constraints:
	//    - nullable
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIpSubnet
	// Source IP subnet enabled or disabled
	// Constraints:
	//    - nullable
	EnableSourceIpSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	// Constraints:
	//    - nullable
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP,UDPLITE,ICMP_ICMPV4,ICMPV6,IGMP,ESP,AH,SCTP,CUSTOM]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP UDPLITE ICMP_ICMPV4 ICMPV6 IGMP ESP AH SCTP CUSTOM"`

	// SourceIp
	// Subnet network address or ip address of source IP.
	// Constraints:
	//    - nullable
	SourceIp *string `json:"sourceIp,omitempty"`

	// SourceIpMask
	// Subnet mask of source IP
	// Constraints:
	//    - nullable
	SourceIpMask *string `json:"sourceIpMask,omitempty"`

	// SourceIpV6
	// Source IPv6 Address.
	// Constraints:
	//    - nullable
	SourceIpV6 *string `json:"sourceIpV6,omitempty"`

	// SourceMaxPort
	// The maxinum port of source port range.
	// Constraints:
	//    - nullable
	SourceMaxPort *int `json:"sourceMaxPort,omitempty"`

	// SourceMinPort
	// The minunum port of source port range.
	// Constraints:
	//    - nullable
	SourceMinPort *int `json:"sourceMinPort,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled
	// Constraints:
	//    - nullable
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`

	// UplinkRateLimitingMbps
	// Uplink rate limiting
	// Constraints:
	//    - nullable
	UplinkRateLimitingMbps *float64 `json:"uplinkRateLimitingMbps,omitempty"`
}

func NewWSGProfileModifyIpAclRules() *WSGProfileModifyIpAclRules {
	m := new(WSGProfileModifyIpAclRules)
	return m
}

type WSGProfileModifyIpsecProfile struct {
	// AdvancedOption
	// Constraints:
	//    - nullable
	AdvancedOption *WSGProfileAdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	// Constraints:
	//    - nullable
	//    - oneof:[PresharedKey,Certificate]
	AuthType *string `json:"authType,omitempty" validate:"omitempty,oneof=PresharedKey Certificate"`

	// CmProtocolOption
	// Constraints:
	//    - nullable
	CmProtocolOption *WSGProfileCmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the IPSec profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	// Constraints:
	//    - nullable
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Constraints:
	//    - nullable
	EspRekeyTimeUnit *WSGCommonTimeUnitStore `json:"espRekeyTimeUnit,omitempty"`

	// EspSecurityAssociation
	// Constraints:
	//    - nullable
	EspSecurityAssociation *WSGProfileEspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// Id
	// identifier of the ipsec profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	// Constraints:
	//    - nullable
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Constraints:
	//    - nullable
	IkeRekeyTimeUnit *WSGCommonTimeUnitStore `json:"ikeRekeyTimeUnit,omitempty"`

	// IkeSecurityAssociation
	// Constraints:
	//    - nullable
	IkeSecurityAssociation *WSGProfileIkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	// IpMode
	// Constraints:
	//    - nullable
	IpMode *WSGProfileIpMode `json:"ipMode,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	// Constraints:
	//    - nullable
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	// Constraints:
	//    - nullable
	ServerAddr *string `json:"serverAddr,omitempty"`
}

func NewWSGProfileModifyIpsecProfile() *WSGProfileModifyIpsecProfile {
	m := new(WSGProfileModifyIpsecProfile)
	return m
}

type WSGProfileModifyL2oGREProfile struct {
	// CoreNetworkGateway
	// Constraints:
	//    - nullable
	CoreNetworkGateway *WSGProfileCoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileModifyL2oGREProfile() *WSGProfileModifyL2oGREProfile {
	m := new(WSGProfileModifyL2oGREProfile)
	return m
}

type WSGProfileModifyL3AccessControlPolicy struct {
	// DefaultAction
	// Default action
	// Constraints:
	//    - required
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=BLOCK ALLOW"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// L3AclRuleList
	// L3 access control list
	// Constraints:
	//    - nullable
	L3AclRuleList []*WSGProfileL3AclRuleList `json:"l3AclRuleList,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileModifyL3AccessControlPolicy() *WSGProfileModifyL3AccessControlPolicy {
	m := new(WSGProfileModifyL3AccessControlPolicy)
	return m
}

type WSGProfileModifyRestrictedApAccessProfile struct {
	// BlockedPortList
	// Blocked Port List
	// Constraints:
	//    - nullable
	BlockedPortList []*WSGProfileBlockedPort `json:"blockedPortList,omitempty" validate:"omitempty,dive"`

	// BlockWellKnownPort
	// Block well known ports
	// Constraints:
	//    - nullable
	//    - default:false
	BlockWellKnownPort *bool `json:"blockWellKnownPort,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// IpAddressWhitelist
	// IP Address Whitelist
	// Constraints:
	//    - nullable
	IpAddressWhitelist []WSGCommonIpAddress `json:"ipAddressWhitelist,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGProfileModifyRestrictedApAccessProfile() *WSGProfileModifyRestrictedApAccessProfile {
	m := new(WSGProfileModifyRestrictedApAccessProfile)
	return m
}

type WSGProfileModifyRuckusGREProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

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

func NewWSGProfileModifyRuckusGREProfile() *WSGProfileModifyRuckusGREProfile {
	m := new(WSGProfileModifyRuckusGREProfile)
	return m
}

type WSGProfileModifySoftGREProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	// Constraints:
	//    - nullable
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
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

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	// Constraints:
	//    - nullable
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	// Constraints:
	//    - nullable
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

func NewWSGProfileModifySoftGREProfile() *WSGProfileModifySoftGREProfile {
	m := new(WSGProfileModifySoftGREProfile)
	return m
}

type WSGProfileModifyUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	// Constraints:
	//    - nullable
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - nullable
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"omitempty,oneof=BLOCK ALLOW"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimiting
	// Constraints:
	//    - nullable
	DownlinkRateLimiting *WSGProfileDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// Id
	// Identifier of the user traffic profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpAclRules
	// Traffic access control list
	// Constraints:
	//    - nullable
	IpAclRules []*WSGProfileModifyIpAclRules `json:"ipAclRules,omitempty" validate:"omitempty,dive"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// QmAppPolicyId
	// Application Policy UUID
	// Constraints:
	//    - nullable
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	// UplinkRateLimiting
	// Constraints:
	//    - nullable
	UplinkRateLimiting *WSGProfileUplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	// Constraints:
	//    - nullable
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

	// ZoneAffinityList
	// Constraints:
	//    - nullable
	ZoneAffinityList []string `json:"zoneAffinityList,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileModifyZoneAffinityProfile() *WSGProfileModifyZoneAffinityProfile {
	m := new(WSGProfileModifyZoneAffinityProfile)
	return m
}

type WSGProfilePrecedenceList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfilePrecedenceListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfilePrecedenceList() *WSGProfilePrecedenceList {
	m := new(WSGProfilePrecedenceList)
	return m
}

type WSGProfilePrecedenceListType struct {
	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	// Constraints:
	//    - nullable
	RateLimitingPrecedence []*WSGProfileRateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty" validate:"omitempty,dive"`

	// VlanPrecedence
	// vlan precedence
	// Constraints:
	//    - nullable
	VlanPrecedence []*WSGProfileVlanPrecedenceItem `json:"vlanPrecedence,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfilePrecedenceListType() *WSGProfilePrecedenceListType {
	m := new(WSGProfilePrecedenceListType)
	return m
}

type WSGProfileClone struct {
	// NewId
	// name for new profile
	// Constraints:
	//    - nullable
	NewId *string `json:"newId,omitempty"`

	// NewName
	// Id for new profile
	// Constraints:
	//    - nullable
	NewName *string `json:"newName,omitempty"`

	// OldId
	// original name
	// Constraints:
	//    - nullable
	OldId *string `json:"oldId,omitempty"`

	// OldName
	// original name
	// Constraints:
	//    - nullable
	OldName *string `json:"oldName,omitempty"`
}

func NewWSGProfileClone() *WSGProfileClone {
	m := new(WSGProfileClone)
	return m
}

type WSGProfileIdList struct {
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
	List []*WSGCommonGenericRef `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileIdList() *WSGProfileIdList {
	m := new(WSGProfileIdList)
	return m
}

type WSGProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

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
	// Constraints:
	//    - nullable
	List []*WSGProfileListType `json:"list,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileList() *WSGProfileList {
	m := new(WSGProfileList)
	return m
}

type WSGProfileListType struct {
	// Id
	// Identifier of the profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
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
	Id *string `json:"id,omitempty"`

	// Name
	// Accounting name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required"`

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
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required"`

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
	// Constraints:
	//    - nullable
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
	OsuServiceUrl *WSGCommonHTTPS `json:"osuServiceUrl" validate:"required"`

	// ProvisioningProtocals
	// Provisioning protocal
	// Constraints:
	//    - required
	ProvisioningProtocals []WSGProfileProviderProvisionProtocal `json:"provisioningProtocals" validate:"required,dive"`

	// SubscriptionDescriptions
	// Subscription descriptions
	// Constraints:
	//    - required
	SubscriptionDescriptions []*WSGProfileProviderSubscriptionDescription `json:"subscriptionDescriptions" validate:"required,dive"`

	// WhitelistedDomains
	// Whitelisted domains
	// Constraints:
	//    - nullable
	WhitelistedDomains []WSGCommonWildFQDN `json:"whitelistedDomains,omitempty" validate:"omitempty,dive"`
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
	ProvisioningProtocals []WSGProfileProviderProvisionProtocal `json:"provisioningProtocals" validate:"required,dive"`

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
	// Constraints:
	//    - nullable
	WhitelistedDomains []WSGCommonWildFQDN `json:"whitelistedDomains,omitempty" validate:"omitempty,dive"`
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
	Realm *WSGCommonRealm `json:"realm" validate:"required"`
}

func NewWSGProfileProviderInternalOSUOsuAuthServicesType() *WSGProfileProviderInternalOSUOsuAuthServicesType {
	m := new(WSGProfileProviderInternalOSUOsuAuthServicesType)
	return m
}

type WSGProfileProviderInternalOSUOsuPortalType struct {
	// ExternalUrl
	// Constraints:
	//    - nullable
	ExternalUrl *WSGCommonHTTPS `json:"externalUrl,omitempty"`

	// InternalOSUPortal
	// Constraints:
	//    - nullable
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
	// ExternalOSU
	// Constraints:
	//    - nullable
	ExternalOSU *WSGProfileProviderExternalOSU `json:"externalOSU,omitempty"`

	// InternalOSU
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Icon *string `json:"icon,omitempty"`

	// Language
	// Constraints:
	//    - required
	Language *WSGCommonLanguageName `json:"language" validate:"required"`

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
// Constraints:
//    - nullable
type WSGProfileRateLimitingPrecedenceItem struct {
	// Name
	// Name of rate limiting precedence item
	// Constraints:
	//    - nullable
	//    - oneof:[AAA,DEVICE,WLANUTP]
	Name *string `json:"name,omitempty" validate:"omitempty,oneof=AAA DEVICE WLANUTP"`

	// Priority
	// Priority
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`
}

func NewWSGProfileRateLimitingPrecedenceItem() *WSGProfileRateLimitingPrecedenceItem {
	m := new(WSGProfileRateLimitingPrecedenceItem)
	return m
}

// WSGProfileRealmAuthServiceMapping
//
// Realm based authentication service mappings
// Constraints:
//    - nullable
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

	// HostedAaaEnabled
	// Constraints:
	//    - nullable
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
	Realm *WSGCommonRealm `json:"realm" validate:"required"`

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

type WSGProfileRestrictedApAccessProfile struct {
	// BlockedPortList
	// Blocked Port List
	// Constraints:
	//    - nullable
	BlockedPortList []*WSGProfileBlockedPort `json:"blockedPortList,omitempty" validate:"omitempty,dive"`

	// BlockWellKnownPort
	// Block well known ports
	// Constraints:
	//    - nullable
	BlockWellKnownPort *bool `json:"blockWellKnownPort,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Restricted AP Access Profile id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpAddressWhitelist
	// IP Address Whitelist
	// Constraints:
	//    - nullable
	IpAddressWhitelist []WSGCommonIpAddress `json:"ipAddressWhitelist,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// ZoneId
	// Zone Id of The Restricted AP Access Profile for clone in System Domain
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileRestrictedApAccessProfile() *WSGProfileRestrictedApAccessProfile {
	m := new(WSGProfileRestrictedApAccessProfile)
	return m
}

type WSGProfileRestrictedApAccessProfileArray struct {
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
	List []*WSGProfileRestrictedApAccessProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileRestrictedApAccessProfileArray() *WSGProfileRestrictedApAccessProfileArray {
	m := new(WSGProfileRestrictedApAccessProfileArray)
	return m
}

type WSGProfileReturnZoneAffinityProfile struct {
	// BaseDpVersion
	// The lowest DP version in an Zone Affinity Profile
	// Constraints:
	//    - nullable
	BaseDpVersion *string `json:"baseDpVersion,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// The description of the profile
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// Id
	// Zone affinity profile key
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IsDpVersionConsistent
	// True if all DPs are the same version
	// Constraints:
	//    - nullable
	IsDpVersionConsistent *bool `json:"isDpVersionConsistent,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Zone affinity profile name
	// Constraints:
	//    - nullable
	//    - max:64
	//    - min:1
	Name *string `json:"name,omitempty" validate:"omitempty,max=64,min=1"`

	// ZoneAffinityList
	// Constraints:
	//    - nullable
	ZoneAffinityList []string `json:"zoneAffinityList,omitempty" validate:"omitempty,dive"`

	// ZoneAffinityListWithPriority
	// Constraints:
	//    - nullable
	ZoneAffinityListWithPriority []*WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType `json:"zoneAffinityListWithPriority,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileReturnZoneAffinityProfile() *WSGProfileReturnZoneAffinityProfile {
	m := new(WSGProfileReturnZoneAffinityProfile)
	return m
}

type WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType struct {
	// DpId
	// DP ID
	// Constraints:
	//    - nullable
	DpId *string `json:"dpId,omitempty"`

	// Priority
	// The priority of DP in zone affinity
	// Constraints:
	//    - nullable
	Priority *float64 `json:"priority,omitempty"`
}

func NewWSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType() *WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType {
	m := new(WSGProfileReturnZoneAffinityProfileZoneAffinityListWithPriorityType)
	return m
}

type WSGProfileRogueApPolicy struct {
	// CreateDateTime
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rules
	// Constraints:
	//    - nullable
	Rules []*WSGProfileRogueApRuleList `json:"rules,omitempty" validate:"omitempty,dive"`

	// ZoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGProfileRogueApPolicy() *WSGProfileRogueApPolicy {
	m := new(WSGProfileRogueApPolicy)
	return m
}

type WSGProfileRogueApPolicyList struct {
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
	List []*WSGProfileRogueApPolicy `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
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
	Name *WSGCommonNormalNameAllowBlank `json:"name" validate:"required"`

	// Priority
	// Constraints:
	//    - required
	Priority *int `json:"priority" validate:"required"`

	// Type
	// Constraints:
	//    - required
	//    - oneof:[AdhocRule,SsidSpoofingRule,MacSpoofingRule,SameNetworkRule,CTSAbuseRule,RTSAbuseRule,DeauthFloodRule,DisassocFloodRule,ExcessivePowerRule,NullSSIDRule,CustomSsidRule,CustomRssiRule,CustomMacOuiRule]
	Type *string `json:"type" validate:"required,oneof=AdhocRule SsidSpoofingRule MacSpoofingRule SameNetworkRule CTSAbuseRule RTSAbuseRule DeauthFloodRule DisassocFloodRule ExcessivePowerRule NullSSIDRule CustomSsidRule CustomRssiRule CustomMacOuiRule"`

	// Value
	// Constraints:
	//    - nullable
	Value interface{} `json:"value,omitempty"`
}

func NewWSGProfileRogueApRuleList() *WSGProfileRogueApRuleList {
	m := new(WSGProfileRogueApRuleList)
	return m
}

type WSGProfileRtlsProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileCreateRtlsProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileRtlsProfileList() *WSGProfileRtlsProfileList {
	m := new(WSGProfileRtlsProfileList)
	return m
}

type WSGProfileRuckusGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the RuckusGRE profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

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
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileRuckusGREProfile() *WSGProfileRuckusGREProfile {
	m := new(WSGProfileRuckusGREProfile)
	return m
}

type WSGProfileRuckusGREProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileRuckusGREProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileRuckusGREProfileList() *WSGProfileRuckusGREProfileList {
	m := new(WSGProfileRuckusGREProfileList)
	return m
}

type WSGProfileSoftGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain id of the SoftGRE profile
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	// Constraints:
	//    - nullable
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpMode
	// Constraints:
	//    - nullable
	IpMode *WSGProfileIpMode `json:"ipMode,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	// Constraints:
	//    - nullable
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	// Constraints:
	//    - nullable
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	// Constraints:
	//    - nullable
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	// Constraints:
	//    - nullable
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	// Constraints:
	//    - nullable
	//    - oneof:[AUTO,MANUAL]
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty" validate:"omitempty,oneof=AUTO MANUAL"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile
	// Constraints:
	//    - nullable
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

func NewWSGProfileSoftGREProfile() *WSGProfileSoftGREProfile {
	m := new(WSGProfileSoftGREProfile)
	return m
}

type WSGProfileSoftGREProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileSoftGREProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileSoftGREProfileList() *WSGProfileSoftGREProfileList {
	m := new(WSGProfileSoftGREProfileList)
	return m
}

type WSGProfileTrafficClassProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

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
	List []*WSGCommonTrafficClassProfileRef `json:"list,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileTrafficClassProfileList() *WSGProfileTrafficClassProfileList {
	m := new(WSGProfileTrafficClassProfileList)
	return m
}

// WSGProfileTtgCommonSetting
//
// Hosted AAA server RADIUS settings & PLMN ID settings
// Constraints:
//    - nullable
type WSGProfileTtgCommonSetting struct {
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

func NewWSGProfileTtgCommonSetting() *WSGProfileTtgCommonSetting {
	m := new(WSGProfileTtgCommonSetting)
	return m
}

type WSGProfileTtgpdgApnForwardingRealm struct {
	// Apn
	// the forwarding policy APN, if apnType is NIOI, APN Example : internet-v4.mnc111.mcc222.gprs
	// Constraints:
	//    - nullable
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

func NewWSGProfileTtgpdgApnForwardingRealm() *WSGProfileTtgpdgApnForwardingRealm {
	m := new(WSGProfileTtgpdgApnForwardingRealm)
	return m
}

type WSGProfileTtgpdgCommonSetting struct {
	// AcctRetry
	// Accounting retry of TTG PDG common setting
	// Constraints:
	//    - nullable
	AcctRetry *int `json:"acctRetry,omitempty"`

	// AcctRetryTimeout
	// Accounting retry timeout(secs) of TTG PDG common setting
	// Constraints:
	//    - nullable
	AcctRetryTimeout *int `json:"acctRetryTimeout,omitempty"`

	// ApnFormat2GGSN
	// APN format to GGSN of TTG PDG common setting
	// Constraints:
	//    - nullable
	//    - oneof:[DNS,String]
	ApnFormat2GGSN *string `json:"apnFormat2GGSN,omitempty" validate:"omitempty,oneof=DNS String"`

	// ApnOIInUse
	// APN-OI of TTG PDG common setting
	// Constraints:
	//    - nullable
	ApnOIInUse *bool `json:"apnOIInUse,omitempty"`

	// PdgUeIdleTimeout
	// PDG UE session idle timeout(secs) of TTG PDG common setting
	// Constraints:
	//    - nullable
	PdgUeIdleTimeout *int `json:"pdgUeIdleTimeout,omitempty"`
}

func NewWSGProfileTtgpdgCommonSetting() *WSGProfileTtgpdgCommonSetting {
	m := new(WSGProfileTtgpdgCommonSetting)
	return m
}

type WSGProfileTtgpdgProfile struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	// Constraints:
	//    - nullable
	ApnForwardingRealms []*WSGProfileTtgpdgApnForwardingRealm `json:"apnForwardingRealms,omitempty" validate:"omitempty,dive"`

	// ApnRealms
	// List of the Default APN
	// Constraints:
	//    - nullable
	ApnRealms []*WSGProfileApnRealm `json:"apnRealms,omitempty" validate:"omitempty,dive"`

	// CommonSetting
	// Constraints:
	//    - nullable
	CommonSetting *WSGProfileTtgpdgCommonSetting `json:"commonSetting,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	// Constraints:
	//    - nullable
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN,omitempty"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	// Constraints:
	//    - nullable
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Profile Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileTtgpdgProfile() *WSGProfileTtgpdgProfile {
	m := new(WSGProfileTtgpdgProfile)
	return m
}

type WSGProfileTtgpdgProfileConfiguration struct {
	// ApnForwardingRealms
	// List of the APN Forwarding Policy Per Realm
	// Constraints:
	//    - nullable
	ApnForwardingRealms []*WSGProfileTtgpdgApnForwardingRealm `json:"apnForwardingRealms,omitempty" validate:"omitempty,dive"`

	// ApnRealms
	// List of the Default APN
	// Constraints:
	//    - nullable
	ApnRealms []*WSGProfileApnRealm `json:"apnRealms,omitempty" validate:"omitempty,dive"`

	// CommonSetting
	// Constraints:
	//    - nullable
	CommonSetting *WSGProfileTtgpdgCommonSetting `json:"commonSetting,omitempty"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	// Constraints:
	//    - nullable
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN,omitempty"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	// Constraints:
	//    - nullable
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRelay
	// Constraints:
	//    - nullable
	DhcpRelay *WSGProfileDhcpRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGProfileTtgpdgProfileConfiguration() *WSGProfileTtgpdgProfileConfiguration {
	m := new(WSGProfileTtgpdgProfileConfiguration)
	return m
}

type WSGProfileTtgpdgProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileTtgpdgProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileTtgpdgProfileList() *WSGProfileTtgpdgProfileList {
	m := new(WSGProfileTtgpdgProfileList)
	return m
}

type WSGProfileUpdateL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	// Constraints:
	//    - nullable
	DataPlanes []*WSGProfileDataPlaneL3RoamingData `json:"dataPlanes,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileUpdateL3RoamingConfig() *WSGProfileUpdateL3RoamingConfig {
	m := new(WSGProfileUpdateL3RoamingConfig)
	return m
}

type WSGProfileUpdatePrecedenceProfile struct {
	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	// Constraints:
	//    - nullable
	RateLimitingPrecedence []*WSGProfileRateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty" validate:"omitempty,dive"`

	// VlanPrecedence
	// vlan precedence
	// Constraints:
	//    - nullable
	VlanPrecedence []*WSGProfileVlanPrecedenceItem `json:"vlanPrecedence,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileUpdatePrecedenceProfile() *WSGProfileUpdatePrecedenceProfile {
	m := new(WSGProfileUpdatePrecedenceProfile)
	return m
}

type WSGProfileUpdateRogueApPolicy struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rules
	// Constraints:
	//    - nullable
	Rules []*WSGProfileRogueApRuleList `json:"rules,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileUpdateRogueApPolicy() *WSGProfileUpdateRogueApPolicy {
	m := new(WSGProfileUpdateRogueApPolicy)
	return m
}

type WSGProfileUpdateRtlsProfile struct {
	// EkahauEnabled
	// Constraints:
	//    - nullable
	EkahauEnabled *bool `json:"ekahauEnabled,omitempty"`

	// EkahauIp
	// Constraints:
	//    - nullable
	EkahauIp *WSGCommonIpAddress `json:"ekahauIp,omitempty"`

	// EkahauPort
	// Constraints:
	//    - nullable
	EkahauPort *int `json:"ekahauPort,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// StanleyEnabled
	// Constraints:
	//    - nullable
	StanleyEnabled *bool `json:"stanleyEnabled,omitempty"`
}

func NewWSGProfileUpdateRtlsProfile() *WSGProfileUpdateRtlsProfile {
	m := new(WSGProfileUpdateRtlsProfile)
	return m
}

type WSGProfileUplinkRateLimiting struct {
	// UplinkRateLimitingBps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	// Constraints:
	//    - nullable
	UplinkRateLimitingBps *string `json:"uplinkRateLimitingBps,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled or disabled
	// Constraints:
	//    - nullable
	//    - default:false
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`
}

func NewWSGProfileUplinkRateLimiting() *WSGProfileUplinkRateLimiting {
	m := new(WSGProfileUplinkRateLimiting)
	return m
}

type WSGProfileUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	// Constraints:
	//    - nullable
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultAction
	// Default action
	// Constraints:
	//    - nullable
	//    - default:'ALLOW'
	//    - oneof:[BLOCK,ALLOW]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"omitempty,oneof=BLOCK ALLOW"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DownlinkRateLimiting
	// Constraints:
	//    - nullable
	DownlinkRateLimiting *WSGProfileDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// Id
	// Identifier of the user traffic profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpAclRules
	// Traffic access control list
	// Constraints:
	//    - nullable
	IpAclRules []*WSGProfileIpAclRules `json:"ipAclRules,omitempty" validate:"omitempty,dive"`

	// IsFactoryDefault
	// Whether the UTP is factory default or not
	// Constraints:
	//    - nullable
	IsFactoryDefault *bool `json:"isFactoryDefault,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// QmAppPolicyId
	// Application Policy UUID
	// Constraints:
	//    - nullable
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	// UplinkRateLimiting
	// Constraints:
	//    - nullable
	UplinkRateLimiting *WSGProfileUplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	// Constraints:
	//    - nullable
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

func NewWSGProfileUserTrafficProfile() *WSGProfileUserTrafficProfile {
	m := new(WSGProfileUserTrafficProfile)
	return m
}

type WSGProfileUserTrafficProfileList struct {
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
	// Constraints:
	//    - nullable
	List []*WSGProfileUserTrafficProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGProfileUserTrafficProfileList() *WSGProfileUserTrafficProfileList {
	m := new(WSGProfileUserTrafficProfileList)
	return m
}

type WSGProfileVdpProfile struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DataVlan
	// data vlan
	// Constraints:
	//    - nullable
	DataVlan *int `json:"dataVlan,omitempty"`

	// ExtIp
	// external ip
	// Constraints:
	//    - nullable
	ExtIp *string `json:"extIp,omitempty"`

	// FwVersion
	// Firmware version
	// Constraints:
	//    - nullable
	FwVersion *string `json:"fwVersion,omitempty"`

	// Ip
	// IP
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// Ipv6
	// IPv6
	// Constraints:
	//    - nullable
	Ipv6 *string `json:"ipv6,omitempty"`

	// IsSupport
	// is support vdp
	// Constraints:
	//    - nullable
	IsSupport *bool `json:"isSupport,omitempty"`

	// LastSeenOn
	// last seen
	// Constraints:
	//    - nullable
	LastSeenOn *string `json:"lastSeenOn,omitempty"`

	// Mac
	// mac
	// Constraints:
	//    - nullable
	Mac *string `json:"mac,omitempty"`

	// ManagedBy
	// managed by
	// Constraints:
	//    - nullable
	ManagedBy *string `json:"managedBy,omitempty"`

	// MgmtExtIp
	// management external ip
	// Constraints:
	//    - nullable
	MgmtExtIp *string `json:"mgmtExtIp,omitempty"`

	// MgmtIp
	// management ip
	// Constraints:
	//    - nullable
	MgmtIp *string `json:"mgmtIp,omitempty"`

	// MgmtVlan
	// management vlan
	// Constraints:
	//    - nullable
	MgmtVlan *int `json:"mgmtVlan,omitempty"`

	// Model
	// model
	// Constraints:
	//    - nullable
	Model *string `json:"model,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// name of vdp  profile
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// RegistrationState
	// registrationState
	// Constraints:
	//    - nullable
	RegistrationState *string `json:"registrationState,omitempty"`

	// SerialNumber
	// serialNumber
	// Constraints:
	//    - nullable
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Status
	// status
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// Uptime
	// uptime
	// Constraints:
	//    - nullable
	Uptime *string `json:"uptime,omitempty"`
}

func NewWSGProfileVdpProfile() *WSGProfileVdpProfile {
	m := new(WSGProfileVdpProfile)
	return m
}

// WSGProfileVlanPrecedenceItem
//
// Vlan precedence item
// Constraints:
//    - nullable
type WSGProfileVlanPrecedenceItem struct {
	// Name
	// Name of the Vlan precedence item
	// Constraints:
	//    - nullable
	//    - oneof:[AAA,DEVICE,WLAN]
	Name *string `json:"name,omitempty" validate:"omitempty,oneof=AAA DEVICE WLAN"`

	// Priority
	// Priority
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`
}

func NewWSGProfileVlanPrecedenceItem() *WSGProfileVlanPrecedenceItem {
	m := new(WSGProfileVlanPrecedenceItem)
	return m
}

type WSGProfileZoneAffinityProfileList struct {
	// List
	// Constraints:
	//    - nullable
	List []*WSGProfileReturnZoneAffinityProfile `json:"list,omitempty" validate:"omitempty,dive"`
}

func NewWSGProfileZoneAffinityProfileList() *WSGProfileZoneAffinityProfileList {
	m := new(WSGProfileZoneAffinityProfileList)
	return m
}
