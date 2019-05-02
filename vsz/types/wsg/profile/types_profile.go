package profile

// API Version: v8_0

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

type AccountingProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AccountingProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

	Realm *common.Realm `json:"realm,omitempty"`

	// ServiceType
	// Accounting service type, NA is NA-Request Rejected
	ServiceType *string `json:"serviceType,omitempty"`
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
	EnforceNatt *string `json:"enforceNatt,omitempty"`

	// FailoverMode
	// mode of the failover
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

type ApnRealm struct {
	// DefaultAPN
	// name of the apnForwardingPolicys.
	DefaultAPN *string `json:"defaultAPN,omitempty"`

	// Realm
	// name of the apnRealm.
	Realm *string `json:"realm,omitempty"`
}

type AuthenticationProfile struct {
	// AaaSuppportEnabled
	// Hosted AAA support enabled or disabled
	AaaSuppportEnabled *bool `json:"aaaSuppportEnabled,omitempty"`

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

type AuthenticationProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AuthenticationProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type BaseServiceInfoList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.BaseServiceInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

	Mac *common.Mac `json:"mac,omitempty"`

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

type BlockClientList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BlockClientListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

type BonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	BonjourFencingRuleList []*BonjourFencingRule `json:"bonjourFencingRuleList,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
}

type BonjourFencingPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BonjourFencingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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
	DeviceType *string `json:"deviceType,omitempty"`

	// FencingRange
	// The range of AP can take Bonjour work
	FencingRange *string `json:"fencingRange,omitempty"`

	ServiceType *BridgeService `json:"serviceType,omitempty"`
}

type BonjourFencingRuleDeviceMac struct {
	Mac *common.Mac `json:"mac,omitempty"`
}

type BonjourFencingRuleMapping struct {
	CustomServiceName *common.NormalName `json:"customServiceName,omitempty"`

	// CustomStringList
	// The array of mdns string
	CustomStringList []string `json:"customStringList,omitempty"`

	ServiceType *BridgeService `json:"serviceType,omitempty"`
}

type BonjourFencingService struct {
	NeighborApMac *string `json:"neighborApMac,omitempty"`

	ServiceType *BridgeService `json:"serviceType,omitempty"`

	SourceType *string `json:"sourceType,omitempty"`
}

type BonjourFencingStatistic struct {
	ApMac *string `json:"apMac,omitempty"`

	DroppedPacketsDueToNeighbor *int `json:"droppedPacketsDueToNeighbor,omitempty"`

	DroppedPacketsDueToServiceType *int `json:"droppedPacketsDueToServiceType,omitempty"`

	ForwardedPackets *int `json:"forwardedPackets,omitempty"`

	ServiceList []*BonjourFencingService `json:"serviceList,omitempty"`
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

type BridgeProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BridgeProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type BridgeService string

type BulkBlockClient struct {
	BlockClientList []*BulkBlockClientBlockClientListType `json:"blockClientList,omitempty"`

	Description *common.Description `json:"description,omitempty"`
}

type BulkBlockClientBlockClientListType struct {
	ApMac *common.Mac `json:"apMac,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

type ClientIsolationEntry struct {
	Description *common.Description `json:"description,omitempty"`

	// Ip
	// Client Entry ip
	Ip *string `json:"ip,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`
}

type ClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	Whitelist []*ClientIsolationEntry `json:"whitelist,omitempty"`

	// ZoneId
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneId *string `json:"zoneId,omitempty"`
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

// CmProtocolOptionContent
//
// Certificate Management Protocol Option
type CmProtocolOptionContent struct {
	// CmpDhcpOpt43Subcode
	// Certificate Management Protocol dhcpOpt43Subcode
	CmpDhcpOpt43Subcode *float64 `json:"cmpDhcpOpt43Subcode,omitempty"`

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

type CoreNetworkGateway struct {
	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	// PrimaryGateway
	// Primary Gateway
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary Gateway
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMTU
	// Gateway path MTU
	TunnelMTU *string `json:"tunnelMTU,omitempty"`

	// TunnelMTUSize
	// Manual setting value of Gateway path MTU
	TunnelMTUSize *int `json:"tunnelMTUSize,omitempty"`
}

type CreateAccountingProfile struct {
	Description *common.DescriptionTo128 `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*AcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

type CreateAuthenticationProfile struct {
	// AaaSuppportEnabled
	// Hosted AAA support enabled or disabled
	AaaSuppportEnabled *bool `json:"aaaSuppportEnabled,omitempty"`

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

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*RealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

type CreateBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	BonjourFencingRuleList []*BonjourFencingRule `json:"bonjourFencingRuleList,omitempty"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*BonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
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

	Name *common.NormalName `json:"name,omitempty"`
}

type CreateClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	Whitelist []*ClientIsolationEntry `json:"whitelist,omitempty"`
}

type CreateDhcpProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// LeaseTimeHours
	// Lease time in hours of the DHCP Profile
	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PoolEndIp *common.IpAddress `json:"poolEndIp,omitempty"`

	PoolStartIp *common.IpAddress `json:"poolStartIp,omitempty"`

	PrimaryDnsIp *common.IpAddress `json:"primaryDnsIp,omitempty"`

	SecondaryDnsIp *common.IpAddress `json:"secondaryDnsIp,omitempty"`

	SubnetMask *common.IpAddress `json:"subnetMask,omitempty"`

	SubnetNetworkIp *common.IpAddress `json:"subnetNetworkIp,omitempty"`

	// VlanId
	// VLAN ID of the DHCP Profile
	VlanId *int `json:"vlanId,omitempty"`
}

type CreateDnsServerProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain UUID
	DomainId *string `json:"domainId,omitempty"`

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

type CreateIpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	AuthType *string `json:"authType,omitempty"`

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

type CreateLOGREProfile struct {
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

type CreateResultList []*common.CreateResult

type CreateRogueApPolicy struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Rules []*RogueApRuleList `json:"rules,omitempty"`
}

type CreateRtlsProfile struct {
	// EkahauEnabled
	// Eekahau Location Service Enabled
	EkahauEnabled *bool `json:"ekahauEnabled,omitempty"`

	EkahauIp *common.IpAddress `json:"ekahauIp,omitempty"`

	// EkahauPort
	// Eekahau Location Server Port
	EkahauPort *int `json:"ekahauPort,omitempty"`

	Id *string `json:"id,omitempty"`

	// StanleyEnabled
	// Stanley Location Service Enabled
	StanleyEnabled *bool `json:"stanleyEnabled,omitempty"`
}

type CreateRuckusGREProfile struct {
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
	TunnelEncryption *string `json:"tunnelEncryption,omitempty"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	TunnelMode *string `json:"tunnelMode,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
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
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile. IPV4:850-1500, IPV6:1384-1500. Default 1500.
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

type CreateTrafficClassProfile struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalNameTo64 `json:"name,omitempty"`

	TrafficClasses []*common.TrafficClassRef `json:"trafficClasses,omitempty"`
}

type CreateTtgpdgProfile struct {
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

	// Id
	// Profile Id
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type CreateUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	DefaultAction *string `json:"defaultAction,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`

	// QmAppPolicyId
	// Application Policy UUID
	QmAppPolicyId *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *UplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// UrlFilteringPolicyId
	// URL Filtering Policy UUID
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`
}

type CreateZoneAffinityProfile struct {
	// Description
	// The description of the profile
	Description *string `json:"description,omitempty"`

	// Name
	// Zone affinity profile name
	Name *string `json:"name,omitempty"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
}

type DataPlaneL3RoamingData struct {
	// Activated
	// Show if this DP is included in the L3 roaming feature or not, 0 means excluded and 1 means included
	Activated *int `json:"activated,omitempty"`

	// FirmwareVersion
	// DP firmware version
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Key
	// Data plane key
	Key *string `json:"key,omitempty"`

	// Name
	// DP name
	Name *string `json:"name,omitempty"`

	SubCriteriaType *string `json:"subCriteriaType,omitempty"`

	// Value
	// A list of L3 roaming configuration for this DP
	Value *string `json:"value,omitempty"`
}

type DeleteBulkAccountingProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type DeleteBulkAuthenticationProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type DeleteBulkPrecedenceProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type DeleteBulkUserTrafficProfile struct {
	IdList common.IdList `json:"idList,omitempty"`
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
	Subopt1Format *string `json:"subopt1Format,omitempty"`

	// Subopt2Enabled
	// Enable subopt-2
	Subopt2Enabled *bool `json:"subopt2Enabled,omitempty"`

	// Subopt2Format
	// Subopt-2 format
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
	Subopt151Format *string `json:"subopt151Format,omitempty"`
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

type DnsServerProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DnsServerProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DownlinkRateLimiting struct {
	// DownlinkRateLimitingBps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	DownlinkRateLimitingBps *string `json:"downlinkRateLimitingBps,omitempty"`

	// DownlinkRateLimitingEnabled
	// Downlink rate limiting enabled or disabled
	DownlinkRateLimitingEnabled *bool `json:"downlinkRateLimitingEnabled,omitempty"`
}

// EspProposal
//
// EspProposal based ipsec service mappings
type EspProposal struct {
	// AuthAlg
	// authAlg of espProposal Specific
	AuthAlg *string `json:"authAlg,omitempty"`

	// DhGroup
	// dhGroup of espProposal Specific
	DhGroup *string `json:"dhGroup,omitempty"`

	// EncAlg
	// encAlg of espProposal Specific
	EncAlg *string `json:"encAlg,omitempty"`
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
	EspProposalType *string `json:"espProposalType,omitempty"`
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

type GetL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*DataPlaneL3RoamingData `json:"dataPlanes,omitempty"`

	// FeatureEnabled
	// Show if L3 roaming feature is enabled or not
	FeatureEnabled *int `json:"featureEnabled,omitempty"`
}

type Hs20FriendlyName struct {
	Language *common.LanguageName `json:"language,omitempty"`

	// Name
	// Name of the friendly name
	Name *string `json:"name,omitempty"`
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
	DomainNames []common.WildFQDN `json:"domainNames,omitempty"`

	// FriendlyNames
	// Friendly names
	FriendlyNames []*Hs20FriendlyName `json:"friendlyNames,omitempty"`

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

	Name *common.NormalName `json:"name,omitempty"`
}

type Hs20OperatorList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Hs20Operator `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

type Hs20ProviderList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Hs20Provider `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// IkeProposal
//
// IkeProposal based ipsec service mappings
type IkeProposal struct {
	// AuthAlg
	// authAlg of ikeProposal Specific
	AuthAlg *string `json:"authAlg,omitempty"`

	// DhGroup
	// dhGroup of ikeProposal Specific
	DhGroup *string `json:"dhGroup,omitempty"`

	// EncAlg
	// encAlg of ikeProposal Specific
	EncAlg *string `json:"encAlg,omitempty"`

	// PrfAlg
	// prfAlg of ikeProposal Specific
	PrfAlg *string `json:"prfAlg,omitempty"`
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
	IkeProposalType *string `json:"ikeProposalType,omitempty"`
}

type IpAclRules struct {
	// Action
	// The access of traffic access control.
	Action *string `json:"action,omitempty"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	CustomProtocol *int `json:"customProtocol,omitempty"`

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
	IpType *string `json:"ipType,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
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

type IpMode string

type IpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	AuthType *string `json:"authType,omitempty"`

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
}

type IpsecProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*IpsecProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type LOGREProfile struct {
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

type LOGREProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LOGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

type LbsProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LbsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

type ModifyAuthenticationProfile struct {
	// AaaSuppportEnabled
	// Hosted AAA support enabled or disabled
	AaaSuppportEnabled *bool `json:"aaaSuppportEnabled,omitempty"`

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

type ModifyBlockClient struct {
	Description *common.Description `json:"description,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`
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

// ModifyIpAclRules
//
// Traffic access control list
type ModifyIpAclRules struct {
	// Action
	// The access of traffic access control.
	Action *string `json:"action,omitempty"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	CustomProtocol *int `json:"customProtocol,omitempty"`

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
	// IP Type(IPv4 or IPv6)
	IpType *string `json:"ipType,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
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

type ModifyIpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	AuthType *string `json:"authType,omitempty"`

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

type ModifyLOGREProfile struct {
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
	TunnelEncryption *string `json:"tunnelEncryption,omitempty"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	TunnelMode *string `json:"tunnelMode,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
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
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PrimaryGateway
	// Primary gateway address of the SoftGRE profile
	PrimaryGateway *string `json:"primaryGateway,omitempty"`

	// SecondaryGateway
	// Secondary gateway address of the SoftGRE profile
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the SoftGRE profile
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile. IPV4:850-1500, IPV6:1384-1500. Default 1500.
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

type ModifyUserTrafficProfile struct {
	// AppPolicyId
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyId *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	DefaultAction *string `json:"defaultAction,omitempty"`

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

type ModifyZoneAffinityProfile struct {
	// Description
	// The description of the profile
	Description *string `json:"description,omitempty"`

	// Name
	// Zone affinity profile name
	Name *string `json:"name,omitempty"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
}

type PrecedenceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PrecedenceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

type ProfileListType struct {
	// Id
	// Identifier of the profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type ProviderAccounting struct {
	// Id
	// Accounting id
	Id *string `json:"id,omitempty"`

	// Name
	// Accounting name
	Name *string `json:"name,omitempty"`

	Realm *common.Realm `json:"realm,omitempty"`

	// ServiceType
	// Accounting service type
	ServiceType *string `json:"serviceType,omitempty"`
}

type ProviderAuthentication struct {
	// Id
	// Authentication id
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication name
	Name *string `json:"name,omitempty"`

	Realm *common.Realm `json:"realm,omitempty"`

	// ServiceType
	// Authentication service type
	ServiceType *string `json:"serviceType,omitempty"`

	// VlanId
	// Dynamic vlan ID
	VlanId *int `json:"vlanId,omitempty"`
}

type ProviderEAPAuthSetting struct {
	// Info
	// EAP auth info
	Info *string `json:"info,omitempty"`

	// Type
	// EAP auth type
	Type *string `json:"type,omitempty"`

	// VendorId
	// EAP auth vendor ID
	VendorId *int `json:"vendorId,omitempty"`

	// VendorType
	// EAP auth vendor type
	VendorType *int `json:"vendorType,omitempty"`
}

type ProviderEAPMethod struct {
	// AuthSettings
	// EAP method auth settings
	AuthSettings []*ProviderEAPAuthSetting `json:"authSettings,omitempty"`

	// Type
	// EAP method type
	Type *string `json:"type,omitempty"`
}

type ProviderExternalOSU struct {
	// CommonLanguageIcon
	// The base64 encoded data of icon.
	CommonLanguageIcon *string `json:"commonLanguageIcon,omitempty"`

	// OsuNaiRealm
	// Online signup NAI realm, it should be one of realm as defined in Hotspot 2.0 identity provider
	OsuNaiRealm *string `json:"osuNaiRealm,omitempty"`

	OsuServiceUrl *common.HTTPS `json:"osuServiceUrl,omitempty"`

	// ProvisioningProtocals
	// Provisioning protocal
	ProvisioningProtocals []ProviderProvisionProtocal `json:"provisioningProtocals,omitempty"`

	// SubscriptionDescriptions
	// Subscription descriptions
	SubscriptionDescriptions []*ProviderSubscriptionDescription `json:"subscriptionDescriptions,omitempty"`

	// WhitelistedDomains
	// Whitelisted domains
	WhitelistedDomains []common.WildFQDN `json:"whitelistedDomains,omitempty"`
}

type ProviderHomeOIs struct {
	// Name
	// Name of the home OI
	Name *string `json:"name,omitempty"`

	// Oi
	// Orgnization ID(3Hex or 5Hex)
	Oi *string `json:"oi,omitempty"`
}

type ProviderInternalOSU struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`

	// CommonLanguageIcon
	// The base64 encoded data of icon.
	CommonLanguageIcon *string `json:"commonLanguageIcon,omitempty"`

	// OsuAuthServices
	// Online signup authentication services
	OsuAuthServices []*ProviderInternalOSUOsuAuthServicesType `json:"osuAuthServices,omitempty"`

	OsuPortal *ProviderInternalOSUOsuPortalType `json:"osuPortal,omitempty"`

	// ProvisioningFormat
	// Provisioning format
	ProvisioningFormat *string `json:"provisioningFormat,omitempty"`

	// ProvisioningProtocals
	// Provisioning protocal
	ProvisioningProtocals []ProviderProvisionProtocal `json:"provisioningProtocals,omitempty"`

	// ProvisioningUpdateType
	// Provisioning update at
	ProvisioningUpdateType *string `json:"provisioningUpdateType,omitempty"`

	// SubscriptionDescriptions
	// Subscription descriptions
	SubscriptionDescriptions []*ProviderSubscriptionDescription `json:"subscriptionDescriptions,omitempty"`

	// WhitelistedDomains
	// whitelisted domains
	WhitelistedDomains []common.WildFQDN `json:"whitelistedDomains,omitempty"`
}

type ProviderInternalOSUOsuAuthServicesType struct {
	// CredentialType
	// Authentication credential type
	CredentialType *string `json:"credentialType,omitempty"`

	// Expiration
	// Expiration hour. null mean never expire
	Expiration *int `json:"expiration,omitempty"`

	// Id
	// Identifier of authentication service
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication service name
	Name *string `json:"name,omitempty"`

	Realm *common.Realm `json:"realm,omitempty"`
}

type ProviderInternalOSUOsuPortalType struct {
	ExternalUrl *common.HTTPS `json:"externalUrl,omitempty"`

	InternalOSUPortal *common.GenericRef `json:"internalOSUPortal,omitempty"`

	// Type
	// Portal type
	Type *string `json:"type,omitempty"`
}

type ProviderOnlineSignup struct {
	ExternalOSU *ProviderExternalOSU `json:"externalOSU,omitempty"`

	InternalOSU *ProviderInternalOSU `json:"internalOSU,omitempty"`

	// Type
	// Online singup type
	Type *string `json:"type,omitempty"`
}

type ProviderPLMN struct {
	// Mcc
	// MCC
	Mcc *string `json:"mcc,omitempty"`

	// Mnc
	// MNC
	Mnc *string `json:"mnc,omitempty"`
}

type ProviderProvisionProtocal string

type ProviderRealm struct {
	// EapMethods
	// EAP methods
	EapMethods []*ProviderEAPMethod `json:"eapMethods,omitempty"`

	// Encoding
	// Encoding
	Encoding *string `json:"encoding,omitempty"`

	// Name
	// Name of realm
	Name *string `json:"name,omitempty"`
}

type ProviderSubscriptionDescription struct {
	// Description
	// Description of the friendly name
	Description *string `json:"description,omitempty"`

	// Icon
	// The binary data of icon, maximum size 65536
	Icon *string `json:"icon,omitempty"`

	Language *common.LanguageName `json:"language,omitempty"`

	// Name
	// Name of the friendly name
	Name *string `json:"name,omitempty"`
}

// RateLimitingPrecedenceItem
//
// Rate limiting precedence item
type RateLimitingPrecedenceItem struct {
	// Name
	// Name of rate limiting precedence item
	Name *string `json:"name,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

// RealmAuthServiceMapping
//
// Realm based authentication service mappings
type RealmAuthServiceMapping struct {
	// AuthorizationMethod
	// Authorization method
	AuthorizationMethod *string `json:"authorizationMethod,omitempty"`

	// DynamicVlanId
	// Dynamic VLAN ID
	DynamicVlanId *int `json:"dynamicVlanId,omitempty"`

	HostedAaaEnabled *bool `json:"hostedAaaEnabled,omitempty"`

	// Id
	// Authentication service UUID
	Id *string `json:"id,omitempty"`

	// Name
	// Authentication service name
	Name *string `json:"name,omitempty"`

	Realm *common.Realm `json:"realm,omitempty"`

	// ServiceType
	// Authentication service type, NA is NA-Request Rejected
	ServiceType *string `json:"serviceType,omitempty"`
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
	Name *string `json:"name,omitempty"`

	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
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

type RogueApPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RogueApPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type RogueApRuleList struct {
	Classification *string `json:"classification,omitempty"`

	Name *common.NormalNameAllowBlank `json:"name,omitempty"`

	Priority *int `json:"priority,omitempty"`

	Type *string `json:"type,omitempty"`

	Value interface{} `json:"value,omitempty"`
}

type RtlsProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateRtlsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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
	TunnelEncryption *string `json:"tunnelEncryption,omitempty"`

	// TunnelMode
	// Ruckus Tunnel Mode of RuckusGRE profile
	TunnelMode *string `json:"tunnelMode,omitempty"`

	// TunnelMtuAutoEnabled
	// WAN Interface MTU of the RuckusGRE profile
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of RuckusGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

type RuckusGREProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RuckusGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`

	// TunnelMtuSize
	// Tunnel MTU size of SoftGRE profile
	TunnelMtuSize *int `json:"tunnelMtuSize,omitempty"`
}

type SoftGREProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SoftGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

// TtgCommonSetting
//
// Hosted AAA server RADIUS settings & PLMN ID settings
type TtgCommonSetting struct {
	// InterimAcctInterval
	// Interim accounting interval (value should be 0 or 600~65536, unit: seconds)
	InterimAcctInterval *int `json:"interimAcctInterval,omitempty"`

	// MobileCountryCode
	// Mobile country code
	MobileCountryCode *string `json:"mobileCountryCode,omitempty"`

	// MobileNetworkCode
	// Mobile network code
	MobileNetworkCode *string `json:"mobileNetworkCode,omitempty"`

	// SessionIdleTimeout
	// Session idle timeout (unit: seconds)
	SessionIdleTimeout *int `json:"sessionIdleTimeout,omitempty"`

	// SessionTimeout
	// Session timeout (unit: seconds)
	SessionTimeout *int `json:"sessionTimeout,omitempty"`
}

type TtgpdgApnForwardingRealm struct {
	// Apn
	// the forwarding policy APN, if apnType is NIOI, APN Example : internet-v4.mnc111.mcc222.gprs
	Apn *string `json:"apn,omitempty"`

	// ApnType
	// type of the forwarding policy APN.
	ApnType *string `json:"apnType,omitempty"`

	// RouteType
	// routeType of the forwarding policy APN.
	RouteType *string `json:"routeType,omitempty"`
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
	ApnFormat2GGSN *string `json:"apnFormat2GGSN,omitempty"`

	// ApnOIInUse
	// APN-OI of TTG PDG common setting
	ApnOIInUse *bool `json:"apnOIInUse,omitempty"`

	// PdgUeIdleTimeout
	// PDG UE session idle timeout(secs) of TTG PDG common setting
	PdgUeIdleTimeout *int `json:"pdgUeIdleTimeout,omitempty"`
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

type TtgpdgProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*TtgpdgProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type UpdateL3RoamingConfig struct {
	// DataPlanes
	// L3 roaming configuration for DPs
	DataPlanes []*DataPlaneL3RoamingData `json:"dataPlanes,omitempty"`

	// FeatureEnabled
	// Show if L3 roaming feature is enabled or not
	FeatureEnabled *int `json:"featureEnabled,omitempty"`
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

type UpdateRogueApPolicy struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Rules []*RogueApRuleList `json:"rules,omitempty"`
}

type UpdateRtlsProfile struct {
	EkahauEnabled *bool `json:"ekahauEnabled,omitempty"`

	EkahauIp *common.IpAddress `json:"ekahauIp,omitempty"`

	EkahauPort *int `json:"ekahauPort,omitempty"`

	Id *string `json:"id,omitempty"`

	StanleyEnabled *bool `json:"stanleyEnabled,omitempty"`
}

type UplinkRateLimiting struct {
	// UplinkRateLimitingBps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	UplinkRateLimitingBps *string `json:"uplinkRateLimitingBps,omitempty"`

	// UplinkRateLimitingEnabled
	// Uplink rate limiting enabled or disabled
	UplinkRateLimitingEnabled *bool `json:"uplinkRateLimitingEnabled,omitempty"`
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
	DefaultAction *string `json:"defaultAction,omitempty"`

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

type UserTrafficProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UserTrafficProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
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

// VlanPrecedenceItem
//
// Vlan precedence item
type VlanPrecedenceItem struct {
	// Name
	// Name of the Vlan precedence item
	Name *string `json:"name,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

type ZoneAffinityProfileList struct {
	List []*ReturnZoneAffinityProfile `json:"list,omitempty"`
}