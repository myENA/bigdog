package profile

// API Version: v8_0

type AccountingProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the accounting profile
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*AcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

type AccountingProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AccountingProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// AcctServiceRealmMapping
//
// Accounting service per realm
type AcctServiceRealmMapping struct {
	// ID
	// Accounting service UUID
	ID *string `json:"id,omitempty"`

	// Name
	// Accounting service name
	Name *string `json:"name,omitempty"`

	Realm *string `json:"realm,omitempty"`

	// ServiceType
	// Accounting service type, NA is NA-Request Rejected
	ServiceType *string `json:"serviceType,omitempty"`
}

// AdvancedOptionContent
//
// advanced option Content
type AdvancedOptionContent struct {
	// DHCPOpt43Subcode
	// dhcpOpt43Subcode of the ipsec profile
	DHCPOpt43Subcode *float64 `json:"dhcpOpt43Subcode,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	GppSuppportEnabled *bool `json:"gppSuppportEnabled,omitempty"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// ID
	// Identifier of the authentication profile
	ID *string `json:"id,omitempty"`

	// IsContainDirectoryService
	// Realm based authentication service mappings contains LDAP or AD service type
	IsContainDirectoryService *bool `json:"isContainDirectoryService,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*RealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

type AuthenticationProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	Mac *string `json:"mac,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// ZoneID
	// Zone Id of the Block Client for clone in System Domain
	ZoneID *string `json:"zoneId,omitempty"`
}

type BlockClientList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BlockClientListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type BlockClientListType struct {
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the profile
	ID *string `json:"id,omitempty"`

	Mac *string `json:"mac,omitempty"`

	// ModifiedDateTime
	// Date blocked of the Block Client
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Modifier blocked of the Block Client
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// ZoneID
	// Zone Id of the Block Client for clone in System Domain
	ZoneID *string `json:"zoneId,omitempty"`
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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// ID
	// Bonjour Fencing Policy id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// ZoneID
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneID *string `json:"zoneId,omitempty"`
}

type BonjourFencingPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BonjourFencingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type BonjourFencingRule struct {
	ClosestAp *string `json:"closestAp,omitempty"`

	CustomServiceName *string `json:"customServiceName,omitempty"`

	Description *string `json:"description,omitempty"`

	// DeviceMacList
	// Specify the device list providing Bonjour Service
	DeviceMacList []*BonjourFencingRuleDeviceMac `json:"deviceMacList,omitempty"`

	// DeviceType
	// Name of the Bonjour Fencing Rule
	DeviceType *string `json:"deviceType,omitempty"`

	// FencingRange
	// The range of AP can take Bonjour work
	FencingRange *string `json:"fencingRange,omitempty"`

	ServiceType *string `json:"serviceType,omitempty"`
}

type BonjourFencingRuleDeviceMac struct {
	Mac *string `json:"mac,omitempty"`
}

type BonjourFencingRuleMapping struct {
	CustomServiceName *string `json:"customServiceName,omitempty"`

	// CustomStringList
	// The array of mdns string
	CustomStringList []string `json:"customStringList,omitempty"`

	ServiceType *string `json:"serviceType,omitempty"`
}

type BonjourFencingService struct {
	NeighborApMac *string `json:"neighborApMac,omitempty"`

	ServiceType *string `json:"serviceType,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`
}

type BridgeProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BridgeProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type BulkBlockClient struct {
	BlockClientList []*BulkBlockClientBlockClientListType `json:"blockClientList,omitempty"`

	Description *string `json:"description,omitempty"`
}

type BulkBlockClientBlockClientListType struct {
	ApMac *string `json:"apMac,omitempty"`

	Description *string `json:"description,omitempty"`

	Mac *string `json:"mac,omitempty"`

	ZoneID *string `json:"zoneId,omitempty"`
}

type ClientIsolationEntry struct {
	Description *string `json:"description,omitempty"`

	// IP
	// Client Entry ip
	IP *string `json:"ip,omitempty"`

	Mac *string `json:"mac,omitempty"`
}

type ClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// ID
	// Client Isolation Whitelist id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	Whitelist []*ClientIsolationEntry `json:"whitelist,omitempty"`

	// ZoneID
	// Zone Id of The Bonjour Fencing Policy for clone in System Domain
	ZoneID *string `json:"zoneId,omitempty"`
}

type ClientIsolationWhitelistArray struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ClientIsolationWhitelist `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// CmProtocolOptionContent
//
// Certificate Management Protocol Option
type CmProtocolOptionContent struct {
	// CmpDHCPOpt43Subcode
	// Certificate Management Protocol dhcpOpt43Subcode
	CmpDHCPOpt43Subcode *float64 `json:"cmpDhcpOpt43Subcode,omitempty"`

	// CmpDHCPOpt43SubcodeRecipient
	// Certificate Management Protocol dhcpOpt43SubcodeRecipient
	CmpDHCPOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`

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
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*AcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

type CreateAuthenticationProfile struct {
	// AaaSuppportEnabled
	// Hosted AAA support enabled or disabled
	AaaSuppportEnabled *bool `json:"aaaSuppportEnabled,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	GppSuppportEnabled *bool `json:"gppSuppportEnabled,omitempty"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

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

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

type CreateBridgeProfile struct {
	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type CreateClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	Whitelist []*ClientIsolationEntry `json:"whitelist,omitempty"`
}

type CreateDHCPProfile struct {
	Description *string `json:"description,omitempty"`

	// LeaseTimeHours
	// Lease time in hours of the DHCP Profile
	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`

	Name *string `json:"name,omitempty"`

	PoolEndIP *string `json:"poolEndIp,omitempty"`

	PoolStartIP *string `json:"poolStartIp,omitempty"`

	PrimaryDNSIP *string `json:"primaryDnsIp,omitempty"`

	SecondaryDNSIP *string `json:"secondaryDnsIp,omitempty"`

	SubnetMask *string `json:"subnetMask,omitempty"`

	SubnetNetworkIP *string `json:"subnetNetworkIp,omitempty"`

	// VlanID
	// VLAN ID of the DHCP Profile
	VlanID *int `json:"vlanId,omitempty"`
}

type CreateDNSServerProfile struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// PrimaryIP
	// Primary ip of DNS server service
	PrimaryIP *string `json:"primaryIp,omitempty"`

	// SecondaryIP
	// Secondary ip of DNS server service
	SecondaryIP *string `json:"secondaryIp,omitempty"`

	// TertiaryIP
	// Tertiary ip of DNS server service
	TertiaryIP *string `json:"tertiaryIp,omitempty"`
}

type CreateIpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	AuthType *string `json:"authType,omitempty"`

	CmProtocolOption *CmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the IPSec profile
	DomainID *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *EspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// ID
	// identifier of the ipsec profile
	ID *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *IkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	IPMode *string `json:"ipMode,omitempty"`

	Name *string `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`
}

type CreateLOGREProfile struct {
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type CreatePrecedenceProfile struct {
	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	Name *string `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*RateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*VlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

type CreateResultList []*common.CreateResult

type CreateRogueApPolicy struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	Rules []*RogueApRuleList `json:"rules,omitempty"`
}

type CreateRtlsProfile struct {
	// EkahauEnabled
	// Eekahau Location Service Enabled
	EkahauEnabled *bool `json:"ekahauEnabled,omitempty"`

	EkahauIP *string `json:"ekahauIp,omitempty"`

	// EkahauPort
	// Eekahau Location Server Port
	EkahauPort *int `json:"ekahauPort,omitempty"`

	ID *string `json:"id,omitempty"`

	// StanleyEnabled
	// Stanley Location Service Enabled
	StanleyEnabled *bool `json:"stanleyEnabled,omitempty"`
}

type CreateRuckusGREProfile struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the RuckusGRE profile
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

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
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the SoftGRE profile
	DomainID *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	IPMode *string `json:"ipMode,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	Name *string `json:"name,omitempty"`

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
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type CreateUserTrafficProfile struct {
	// AppPolicyID
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyID *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// IPAclRules
	// Traffic access control list
	IPAclRules []*IPAclRules `json:"ipAclRules,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// QmAppPolicyID
	// Application Policy UUID
	QmAppPolicyID *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *UplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// URLFilteringPolicyID
	// URL Filtering Policy UUID
	URLFilteringPolicyID *string `json:"urlFilteringPolicyId,omitempty"`
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
	IDList common.IDList `json:"idList,omitempty"`
}

type DeleteBulkAuthenticationProfile struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DeleteBulkPrecedenceProfile struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DeleteBulkUserTrafficProfile struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DHCPOption82 struct {
	// DHCPOption82Enabled
	// Enable DHCP Option 82
	DHCPOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

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

type DHCPProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.DHCPProfileRef `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DHCPRelayNoRelayTunnel struct {
	DHCPOption82 *DHCPOption82 `json:"dhcpOption82,omitempty"`

	// DHCPRelayEnabled
	// Enable DHCP Relay
	DHCPRelayEnabled *bool `json:"dhcpRelayEnabled,omitempty"`

	// DHCPServer1
	// DHCP Server 1
	DHCPServer1 *string `json:"dhcpServer1,omitempty"`

	// DHCPServer2
	// DHCP Server 2
	DHCPServer2 *string `json:"dhcpServer2,omitempty"`

	// RelayBothEnabled
	// Send DHCP requests to both servers simultaneously.
	RelayBothEnabled *bool `json:"relayBothEnabled,omitempty"`
}

type DNSServerProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// PrimaryIP
	// Primary ip of DNS server service
	PrimaryIP *string `json:"primaryIp,omitempty"`

	// SecondaryIP
	// Secondary ip of DNS server service
	SecondaryIP *string `json:"secondaryIp,omitempty"`

	// TertiaryIP
	// Tertiary ip of DNS server service
	TertiaryIP *string `json:"tertiaryIp,omitempty"`
}

type DNSServerProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DNSServerProfile `json:"list,omitempty"`

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
	// DestinationZoneAffinityID
	// Flexi-VPN Profile ID (Destination)
	DestinationZoneAffinityID *string `json:"destinationZoneAffinityId,omitempty"`

	// DestinationZoneAffinityName
	// Flexi-VPN Profile (Destination)
	DestinationZoneAffinityName *string `json:"destinationZoneAffinityName,omitempty"`

	// DomainID
	// Domain ID
	DomainID *string `json:"domainId,omitempty"`

	// SourceZoneAffinityID
	// Zone Affinity Profile ID (Source)
	SourceZoneAffinityID *string `json:"sourceZoneAffinityId,omitempty"`

	// SourceZoneAffinityName
	// Zone Affinity Profile (Source)
	SourceZoneAffinityName *string `json:"sourceZoneAffinityName,omitempty"`

	// WLANID
	// Wlan ID
	WLANID *string `json:"wlanId,omitempty"`

	// WLANName
	// Wlan name
	WLANName *string `json:"wlanName,omitempty"`

	// ZoneID
	// Zone ID
	ZoneID *string `json:"zoneId,omitempty"`

	// ZoneName
	// Zone name
	ZoneName *string `json:"zoneName,omitempty"`
}

type FlexiVpnProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*FlexiVpnProfile `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

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
	Language *string `json:"language,omitempty"`

	// Name
	// Name of the friendly name
	Name *string `json:"name,omitempty"`
}

type Hs20Operator struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	DomainNames []string `json:"domainNames,omitempty"`

	// FriendlyNames
	// Friendly names
	FriendlyNames []*Hs20FriendlyName `json:"friendlyNames,omitempty"`

	// ID
	// Identifier of the profile
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`
}

type Hs20OperatorList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// HomeOis
	// Home OIs
	HomeOis []*ProviderHomeOIs `json:"homeOis,omitempty"`

	// ID
	// Identifier of the Hotspot 2.0 identity provider profile
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	Osu *ProviderOnlineSignup `json:"osu,omitempty"`

	// Plmns
	// PLMNs
	Plmns []*ProviderPLMN `json:"plmns,omitempty"`

	// Realms
	// Realms
	Realms []*ProviderRealm `json:"realms,omitempty"`
}

type Hs20ProviderList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

type IPAclRules struct {
	// Action
	// The access of traffic access control.
	Action *string `json:"action,omitempty"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	CustomProtocol *int `json:"customProtocol,omitempty"`

	Description *string `json:"description,omitempty"`

	// DestinationIP
	// Subnet network address or ip address of destination IP.
	DestinationIP *string `json:"destinationIp,omitempty"`

	// DestinationIPMask
	// Subnet mask of destination IP
	DestinationIPMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIPV6
	// Destination IPv6 Address.
	DestinationIPV6 *string `json:"destinationIpV6,omitempty"`

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

	// EnableDestinationIPSubnet
	// Destination IP subnet enabled or disabled
	EnableDestinationIPSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIPSubnet
	// Source IP subnet enabled or disabled
	EnableSourceIPSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	EnableSourceV6Prefix *bool `json:"enableSourceV6Prefix,omitempty"`

	// IPType
	// IP Type(IPv4 or IPv6).
	IPType *string `json:"ipType,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	Protocol *string `json:"protocol,omitempty"`

	// SourceIP
	// Subnet network address or ip address of source IP.
	SourceIP *string `json:"sourceIp,omitempty"`

	// SourceIPMask
	// Subnet mask of source IP
	SourceIPMask *string `json:"sourceIpMask,omitempty"`

	// SourceIPV6
	// Source IPv6 Address.
	SourceIPV6 *string `json:"sourceIpV6,omitempty"`

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

type IpsecProfile struct {
	AdvancedOption *AdvancedOptionContent `json:"advancedOption,omitempty"`

	// AuthType
	// authentication type of the ipsec profile
	AuthType *string `json:"authType,omitempty"`

	CmProtocolOption *CmProtocolOptionContent `json:"cmProtocolOption,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the IPSec profile
	DomainID *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *EspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// ID
	// identifier of the ipsec profile
	ID *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *IkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	IPMode *string `json:"ipMode,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`
}

type IpsecProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`
}

type LOGREProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LOGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type LbsProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// Password
	// Password
	Password *string `json:"password,omitempty"`

	// Port
	// LBS port
	Port *int `json:"port,omitempty"`

	// URL
	// LBS url
	URL *string `json:"url,omitempty"`

	// Venue
	// Venue
	Venue *string `json:"venue,omitempty"`
}

type LbsProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LbsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ModifyAccountingProfile struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// ID of Accounting Profile
	ID *string `json:"id,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// RealmMappings
	// Accounting service per realm
	RealmMappings []*AcctServiceRealmMapping `json:"realmMappings,omitempty"`
}

type ModifyAuthenticationProfile struct {
	// AaaSuppportEnabled
	// Hosted AAA support enabled or disabled
	AaaSuppportEnabled *bool `json:"aaaSuppportEnabled,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// GppSuppportEnabled
	// 3GPP support enabled or disabled
	GppSuppportEnabled *bool `json:"gppSuppportEnabled,omitempty"`

	// H20SuppportEnabled
	// Hotspot 2.0 support enabled or disabled
	H20SuppportEnabled *bool `json:"h20SuppportEnabled,omitempty"`

	// ID
	// ID of Accounting Profile
	ID *string `json:"id,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// RealmMappings
	// Realm based authentication service mappings
	RealmMappings []*RealmAuthServiceMapping `json:"realmMappings,omitempty"`

	TtgCommonSetting *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

type ModifyBlockClient struct {
	Description *string `json:"description,omitempty"`

	Mac *string `json:"mac,omitempty"`
}

type ModifyBonjourFencingPolicy struct {
	// BonjourFencingRuleList
	// Bonjour Fencing Rule List
	BonjourFencingRuleList []*BonjourFencingRule `json:"bonjourFencingRuleList,omitempty"`

	// BonjourFencingRuleMappingList
	// Bonjour Fencing Rule Mapping List
	BonjourFencingRuleMappingList []*BonjourFencingRuleMapping `json:"bonjourFencingRuleMappingList,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

type ModifyBridgeProfile struct {
	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type ModifyClientIsolationWhitelist struct {
	// ClientIsolationAutoEnabled
	// Client Isolation Auto Enable
	ClientIsolationAutoEnabled *bool `json:"clientIsolationAutoEnabled,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// Whitelist
	// Client Isolation Whitelist array
	Whitelist []*ClientIsolationEntry `json:"whitelist,omitempty"`
}

type ModifyDNSServerProfile struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// PrimaryIP
	// Primary ip of DNS server service
	PrimaryIP *string `json:"primaryIp,omitempty"`

	// SecondaryIP
	// Secondary ip of DNS server service
	SecondaryIP *string `json:"secondaryIp,omitempty"`

	// TertiaryIP
	// Tertiary ip of DNS server service
	TertiaryIP *string `json:"tertiaryIp,omitempty"`
}

type ModifyHS20Operator struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// DomainNames
	// Domain names
	DomainNames []string `json:"domainNames,omitempty"`

	// FriendlyNames
	// Friendly names
	FriendlyNames []*Hs20FriendlyName `json:"friendlyNames,omitempty"`

	// ID
	// Identifier of the profile
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

// ModifyIPAclRules
//
// Traffic access control list
type ModifyIPAclRules struct {
	// Action
	// The access of traffic access control.
	Action *string `json:"action,omitempty"`

	// CustomProtocol
	// The protocol of traffic access control. Available if the protocol is set to CUSTOM.
	CustomProtocol *int `json:"customProtocol,omitempty"`

	Description *string `json:"description,omitempty"`

	// DestinationIP
	// Subnet network address or ip address of destination IP.
	DestinationIP *string `json:"destinationIp,omitempty"`

	// DestinationIPMask
	// Subnet mask of destination IP
	DestinationIPMask *string `json:"destinationIpMask,omitempty"`

	// DestinationIPV6
	// Destination IPv6 Address.
	DestinationIPV6 *string `json:"destinationIpV6,omitempty"`

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

	// EnableDestinationIPSubnet
	// Destination IP subnet enabled or disabled
	EnableDestinationIPSubnet *bool `json:"enableDestinationIpSubnet,omitempty"`

	// EnableDestinationPortRange
	// Destincation port range enabled or disabled
	EnableDestinationPortRange *bool `json:"enableDestinationPortRange,omitempty"`

	// EnableDestinationV6Prefix
	// Enable Destination IPv6 prefix.
	EnableDestinationV6Prefix *bool `json:"enableDestinationV6Prefix,omitempty"`

	// EnableSourceIPSubnet
	// Source IP subnet enabled or disabled
	EnableSourceIPSubnet *bool `json:"enableSourceIpSubnet,omitempty"`

	// EnableSourcePortRange
	// Source port range enabled or disabled
	EnableSourcePortRange *bool `json:"enableSourcePortRange,omitempty"`

	// EnableSourceV6Prefix
	// Enable Source IPv6 prefix.
	EnableSourceV6Prefix *bool `json:"enableSourceV6Prefix,omitempty"`

	// IPType
	// IP Type(IPv4 or IPv6)
	IPType *string `json:"ipType,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`

	// Protocol
	// The protocol of traffic access control.
	Protocol *string `json:"protocol,omitempty"`

	// SourceIP
	// Subnet network address or ip address of source IP.
	SourceIP *string `json:"sourceIp,omitempty"`

	// SourceIPMask
	// Subnet mask of source IP
	SourceIPMask *string `json:"sourceIpMask,omitempty"`

	// SourceIPV6
	// Source IPv6 Address.
	SourceIPV6 *string `json:"sourceIpV6,omitempty"`

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

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the IPSec profile
	DomainID *string `json:"domainId,omitempty"`

	// EspRekeyTime
	// espRekey Time of the ipsec profile
	EspRekeyTime *float64 `json:"espRekeyTime,omitempty"`

	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	EspSecurityAssociation *EspSecurityAssociationContent `json:"espSecurityAssociation,omitempty"`

	// ID
	// identifier of the ipsec profile
	ID *string `json:"id,omitempty"`

	// IkeRekeyTime
	// ikeRekey Time of the ipsec profile
	IkeRekeyTime *float64 `json:"ikeRekeyTime,omitempty"`

	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	IkeSecurityAssociation *IkeSecurityAssociationContent `json:"ikeSecurityAssociation,omitempty"`

	IPMode *string `json:"ipMode,omitempty"`

	Name *string `json:"name,omitempty"`

	// PreSharedKey
	// authentication preShared Key of the ipsec profile
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// ServerAddr
	// server Addr of the ipsec profile
	ServerAddr *string `json:"serverAddr,omitempty"`
}

type ModifyLOGREProfile struct {
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway,omitempty"`

	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type ModifyRuckusGREProfile struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the RuckusGRE profile
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

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
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the SoftGRE profile
	DomainID *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	Name *string `json:"name,omitempty"`

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
	// AppPolicyID
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyID *string `json:"appPolicyId,omitempty"`

	// DefaultAction
	// Default action
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// ID
	// Identifier of the user traffic profile
	ID *string `json:"id,omitempty"`

	// IPAclRules
	// Traffic access control list
	IPAclRules []*ModifyIPAclRules `json:"ipAclRules,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// QmAppPolicyID
	// Application Policy UUID
	QmAppPolicyID *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *UplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// URLFilteringPolicyID
	// URL Filtering Policy UUID
	URLFilteringPolicyID *string `json:"urlFilteringPolicyId,omitempty"`
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
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PrecedenceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type PrecedenceListType struct {
	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the profile
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*RateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*VlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

type ProfileCloneRequest struct {
	// NewID
	// name for new profile
	NewID *string `json:"newId,omitempty"`

	// NewName
	// Id for new profile
	NewName *string `json:"newName,omitempty"`

	// OldID
	// original name
	OldID *string `json:"oldId,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ProfileListType `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ProfileListType struct {
	// ID
	// Identifier of the profile
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type ProviderAccounting struct {
	// ID
	// Accounting id
	ID *string `json:"id,omitempty"`

	// Name
	// Accounting name
	Name *string `json:"name,omitempty"`

	Realm *string `json:"realm,omitempty"`

	// ServiceType
	// Accounting service type
	ServiceType *string `json:"serviceType,omitempty"`
}

type ProviderAuthentication struct {
	// ID
	// Authentication id
	ID *string `json:"id,omitempty"`

	// Name
	// Authentication name
	Name *string `json:"name,omitempty"`

	Realm *string `json:"realm,omitempty"`

	// ServiceType
	// Authentication service type
	ServiceType *string `json:"serviceType,omitempty"`

	// VlanID
	// Dynamic vlan ID
	VlanID *int `json:"vlanId,omitempty"`
}

type ProviderEAPAuthSetting struct {
	// Info
	// EAP auth info
	Info *string `json:"info,omitempty"`

	// Type
	// EAP auth type
	Type *string `json:"type,omitempty"`

	// VendorID
	// EAP auth vendor ID
	VendorID *int `json:"vendorId,omitempty"`

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

	OsuServiceURL *string `json:"osuServiceUrl,omitempty"`

	// ProvisioningProtocals
	// Provisioning protocal
	ProvisioningProtocals []string `json:"provisioningProtocals,omitempty"`

	// SubscriptionDescriptions
	// Subscription descriptions
	SubscriptionDescriptions []*ProviderSubscriptionDescription `json:"subscriptionDescriptions,omitempty"`

	// WhitelistedDomains
	// Whitelisted domains
	WhitelistedDomains []string `json:"whitelistedDomains,omitempty"`
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
	ProvisioningProtocals []string `json:"provisioningProtocals,omitempty"`

	// ProvisioningUpdateType
	// Provisioning update at
	ProvisioningUpdateType *string `json:"provisioningUpdateType,omitempty"`

	// SubscriptionDescriptions
	// Subscription descriptions
	SubscriptionDescriptions []*ProviderSubscriptionDescription `json:"subscriptionDescriptions,omitempty"`

	// WhitelistedDomains
	// whitelisted domains
	WhitelistedDomains []string `json:"whitelistedDomains,omitempty"`
}

type ProviderInternalOSUOsuAuthServicesType struct {
	// CredentialType
	// Authentication credential type
	CredentialType *string `json:"credentialType,omitempty"`

	// Expiration
	// Expiration hour. null mean never expire
	Expiration *int `json:"expiration,omitempty"`

	// ID
	// Identifier of authentication service
	ID *string `json:"id,omitempty"`

	// Name
	// Authentication service name
	Name *string `json:"name,omitempty"`

	Realm *string `json:"realm,omitempty"`
}

type ProviderInternalOSUOsuPortalType struct {
	ExternalURL *string `json:"externalUrl,omitempty"`

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

	Language *string `json:"language,omitempty"`

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

	// DynamicVlanID
	// Dynamic VLAN ID
	DynamicVlanID *int `json:"dynamicVlanId,omitempty"`

	HostedAaaEnabled *bool `json:"hostedAaaEnabled,omitempty"`

	// ID
	// Authentication service UUID
	ID *string `json:"id,omitempty"`

	// Name
	// Authentication service name
	Name *string `json:"name,omitempty"`

	Realm *string `json:"realm,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// The description of the profile
	Description *string `json:"description,omitempty"`

	// ID
	// Zone affinity profile key
	ID *string `json:"id,omitempty"`

	// IsDpVersionConsistent
	// True if all DPs are the same version
	IsDpVersionConsistent *bool `json:"isDpVersionConsistent,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

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

	CreatorID *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierID *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	Rules []*RogueApRuleList `json:"rules,omitempty"`

	ZoneID *string `json:"zoneId,omitempty"`
}

type RogueApPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RogueApPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type RogueApRuleList struct {
	Classification *string `json:"classification,omitempty"`

	Name *string `json:"name,omitempty"`

	Priority *int `json:"priority,omitempty"`

	Type *string `json:"type,omitempty"`

	Value interface{} `json:"value,omitempty"`
}

type RtlsProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateRtlsProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type RuckusGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the RuckusGRE profile
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

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
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RuckusGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SoftGREProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain id of the SoftGRE profile
	DomainID *string `json:"domainId,omitempty"`

	// ForceDisassociateClient
	// Force Disassociate Client
	ForceDisassociateClient *bool `json:"forceDisassociateClient,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	IPMode *string `json:"ipMode,omitempty"`

	// KeepAlivePeriod
	// ICMP Keep-Alive Period(secs)
	KeepAlivePeriod *int `json:"keepAlivePeriod,omitempty"`

	// KeepAliveRetry
	// ICMP Keep-Alive Retry
	KeepAliveRetry *int `json:"keepAliveRetry,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

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
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SoftGREProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type TrafficClassProfileList struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.TrafficClassProfileRef `json:"list,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

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

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultNoMatchingAPN
	// Default APN of the No Matching Realm Found
	DefaultNoMatchingAPN *string `json:"defaultNoMatchingAPN,omitempty"`

	// DefaultNoRealmAPN
	// Default APN of the No Realm Specified
	DefaultNoRealmAPN *string `json:"defaultNoRealmAPN,omitempty"`

	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Profile Id
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`
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

	Description *string `json:"description,omitempty"`

	DHCPRelay *DHCPRelayNoRelayTunnel `json:"dhcpRelay,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	Name *string `json:"name,omitempty"`
}

type TtgpdgProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

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
	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	Name *string `json:"name,omitempty"`

	// RateLimitingPrecedence
	// rate limiting precedence
	RateLimitingPrecedence []*RateLimitingPrecedenceItem `json:"rateLimitingPrecedence,omitempty"`

	// VlanPrecedence
	// vlan precedence
	VlanPrecedence []*VlanPrecedenceItem `json:"vlanPrecedence,omitempty"`
}

type UpdateRogueApPolicy struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	Rules []*RogueApRuleList `json:"rules,omitempty"`
}

type UpdateRtlsProfile struct {
	EkahauEnabled *bool `json:"ekahauEnabled,omitempty"`

	EkahauIP *string `json:"ekahauIp,omitempty"`

	EkahauPort *int `json:"ekahauPort,omitempty"`

	ID *string `json:"id,omitempty"`

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
	// AppPolicyID
	// Application Policy UUID (for 5.0 and Earlier Firmware Versions)
	AppPolicyID *string `json:"appPolicyId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultAction
	// Default action
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain UUID
	DomainID *string `json:"domainId,omitempty"`

	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`

	// ID
	// Identifier of the user traffic profile
	ID *string `json:"id,omitempty"`

	// IPAclRules
	// Traffic access control list
	IPAclRules []*IPAclRules `json:"ipAclRules,omitempty"`

	// IsFactoryDefault
	// Whether the UTP is factory default or not
	IsFactoryDefault *bool `json:"isFactoryDefault,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// QmAppPolicyID
	// Application Policy UUID
	QmAppPolicyID *string `json:"qmAppPolicyId,omitempty"`

	UplinkRateLimiting *UplinkRateLimiting `json:"uplinkRateLimiting,omitempty"`

	// URLFilteringPolicyID
	// URL Filtering Policy UUID
	URLFilteringPolicyID *string `json:"urlFilteringPolicyId,omitempty"`
}

type UserTrafficProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UserTrafficProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type VdpProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DataVlan
	// data vlan
	DataVlan *int `json:"dataVlan,omitempty"`

	// ExtIP
	// external ip
	ExtIP *string `json:"extIp,omitempty"`

	// FwVersion
	// Firmware version
	FwVersion *string `json:"fwVersion,omitempty"`

	// IP
	// IP
	IP *string `json:"ip,omitempty"`

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

	// MgmtExtIP
	// management external ip
	MgmtExtIP *string `json:"mgmtExtIp,omitempty"`

	// MgmtIP
	// management ip
	MgmtIP *string `json:"mgmtIp,omitempty"`

	// MgmtVlan
	// management vlan
	MgmtVlan *int `json:"mgmtVlan,omitempty"`

	// Model
	// model
	Model *string `json:"model,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

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
