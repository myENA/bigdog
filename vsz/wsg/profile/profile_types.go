package profile

// API Version: v8_0

type AccountingProfile struct {
	CreateDateTime   *int             `json:"createDateTime,omitempty"`
	CreatorID        *string          `json:"creatorId,omitempty"`
	CreatorUsername  *string          `json:"creatorUsername,omitempty"`
	Description      *string          `json:"description,omitempty"`
	DomainID         *string          `json:"domainId,omitempty"`
	ID               *string          `json:"id,omitempty"`
	ModifiedDateTime *int             `json:"modifiedDateTime,omitempty"`
	ModifierID       *string          `json:"modifierId,omitempty"`
	ModifierUsername *string          `json:"modifierUsername,omitempty"`
	MvnoID           *string          `json:"mvnoId,omitempty"`
	Name             *string          `json:"name,omitempty"`
	RealmMappings    []*RealmMappings `json:"realmMappings,omitempty"`
}

type AccountingProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type AcctServiceRealmMapping struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Realm       *string `json:"realm,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
}

type AdvancedOptionContent struct {
	DHCPOpt43Subcode             *float64 `json:"dhcpOpt43Subcode,omitempty"`
	DpdDelay                     *float64 `json:"dpdDelay,omitempty"`
	EnforceNatt                  *string  `json:"enforceNatt,omitempty"`
	FailoverMode                 *string  `json:"failoverMode,omitempty"`
	FailoverPrimaryCheckInterval *float64 `json:"failoverPrimaryCheckInterval,omitempty"`
	FailoverRetryInterval        *float64 `json:"failoverRetryInterval,omitempty"`
	FailoverRetryPeriod          *float64 `json:"failoverRetryPeriod,omitempty"`
	IpcompEnable                 *string  `json:"ipcompEnable,omitempty"`
	KeepAliveIntval              *float64 `json:"keepAliveIntval,omitempty"`
	ReplayWindow                 *float64 `json:"replayWindow,omitempty"`
	RetryLimit                   *float64 `json:"retryLimit,omitempty"`
}

type ApnRealm struct {
	DefaultAPN *string `json:"defaultAPN,omitempty"`
	Realm      *string `json:"realm,omitempty"`
}

type AuthenticationProfile struct {
	AaaSuppportEnabled        *bool             `json:"aaaSuppportEnabled,omitempty"`
	CreateDateTime            *int              `json:"createDateTime,omitempty"`
	CreatorID                 *string           `json:"creatorId,omitempty"`
	CreatorUsername           *string           `json:"creatorUsername,omitempty"`
	Description               *string           `json:"description,omitempty"`
	DomainID                  *string           `json:"domainId,omitempty"`
	GppSuppportEnabled        *bool             `json:"gppSuppportEnabled,omitempty"`
	H20SuppportEnabled        *bool             `json:"h20SuppportEnabled,omitempty"`
	ID                        *string           `json:"id,omitempty"`
	IsContainDirectoryService *bool             `json:"isContainDirectoryService,omitempty"`
	ModifiedDateTime          *int              `json:"modifiedDateTime,omitempty"`
	ModifierID                *string           `json:"modifierId,omitempty"`
	ModifierUsername          *string           `json:"modifierUsername,omitempty"`
	MvnoID                    *string           `json:"mvnoId,omitempty"`
	Name                      *string           `json:"name,omitempty"`
	RealmMappings             []*RealmMappings  `json:"realmMappings,omitempty"`
	TtgCommonSetting          *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

type AuthenticationProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type BaseServiceInfoList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type BlockClient struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	Description      *string `json:"description,omitempty"`
	Mac              *string `json:"mac,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	ZoneID           *string `json:"zoneId,omitempty"`
}

type BlockClientList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type BlockClientListType struct {
	Description      *string `json:"description,omitempty"`
	ID               *string `json:"id,omitempty"`
	Mac              *string `json:"mac,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	ZoneID           *string `json:"zoneId,omitempty"`
}

type BonjourFencingPolicy struct {
	BonjourFencingRuleList        []*BonjourFencingRuleList        `json:"bonjourFencingRuleList,omitempty"`
	BonjourFencingRuleMappingList []*BonjourFencingRuleMappingList `json:"bonjourFencingRuleMappingList,omitempty"`
	CreateDateTime                *int                             `json:"createDateTime,omitempty"`
	CreatorID                     *string                          `json:"creatorId,omitempty"`
	CreatorUsername               *string                          `json:"creatorUsername,omitempty"`
	Description                   *string                          `json:"description,omitempty"`
	ID                            *string                          `json:"id,omitempty"`
	ModifiedDateTime              *int                             `json:"modifiedDateTime,omitempty"`
	ModifierID                    *string                          `json:"modifierId,omitempty"`
	ModifierUsername              *string                          `json:"modifierUsername,omitempty"`
	Name                          *string                          `json:"name,omitempty"`
	ZoneID                        *string                          `json:"zoneId,omitempty"`
}

type BonjourFencingPolicyList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type BonjourFencingRule struct {
	ClosestAp         *string          `json:"closestAp,omitempty"`
	CustomServiceName *string          `json:"customServiceName,omitempty"`
	Description       *string          `json:"description,omitempty"`
	DeviceMacList     []*DeviceMacList `json:"deviceMacList,omitempty"`
	DeviceType        *string          `json:"deviceType,omitempty"`
	FencingRange      *string          `json:"fencingRange,omitempty"`
	ServiceType       *string          `json:"serviceType,omitempty"`
}

type BonjourFencingRuleDeviceMac struct {
	Mac *string `json:"mac,omitempty"`
}

type BonjourFencingRuleMapping struct {
	CustomServiceName *string  `json:"customServiceName,omitempty"`
	CustomStringList  []string `json:"customStringList,omitempty"`
	ServiceType       *string  `json:"serviceType,omitempty"`
}

type BonjourFencingService struct {
	NeighborApMac *string `json:"neighborApMac,omitempty"`
	ServiceType   *string `json:"serviceType,omitempty"`
	SourceType    *string `json:"sourceType,omitempty"`
}

type BonjourFencingStatistic struct {
	ApMac                          *string        `json:"apMac,omitempty"`
	DroppedPacketsDueToNeighbor    *int           `json:"droppedPacketsDueToNeighbor,omitempty"`
	DroppedPacketsDueToServiceType *int           `json:"droppedPacketsDueToServiceType,omitempty"`
	ForwardedPackets               *int           `json:"forwardedPackets,omitempty"`
	ServiceList                    []*ServiceList `json:"serviceList,omitempty"`
}

type BridgeProfile struct {
	CreateDateTime   *int       `json:"createDateTime,omitempty"`
	CreatorID        *string    `json:"creatorId,omitempty"`
	CreatorUsername  *string    `json:"creatorUsername,omitempty"`
	Description      *string    `json:"description,omitempty"`
	DHCPRelay        *DHCPRelay `json:"dhcpRelay,omitempty"`
	DomainID         *string    `json:"domainId,omitempty"`
	ID               *string    `json:"id,omitempty"`
	ModifiedDateTime *int       `json:"modifiedDateTime,omitempty"`
	ModifierID       *string    `json:"modifierId,omitempty"`
	ModifierUsername *string    `json:"modifierUsername,omitempty"`
	Name             *string    `json:"name,omitempty"`
}

type BridgeProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type BulkBlockClient struct {
	BlockClientList []*BlockClientList `json:"blockClientList,omitempty"`
	Description     *string            `json:"description,omitempty"`
}

type BulkBlockClientBlockClientListType struct {
	ApMac       *string `json:"apMac,omitempty"`
	Description *string `json:"description,omitempty"`
	Mac         *string `json:"mac,omitempty"`
	ZoneID      *string `json:"zoneId,omitempty"`
}

type ClientIsolationEntry struct {
	Description *string `json:"description,omitempty"`
	IP          *string `json:"ip,omitempty"`
	Mac         *string `json:"mac,omitempty"`
}

type ClientIsolationWhitelist struct {
	ClientIsolationAutoEnabled *bool        `json:"clientIsolationAutoEnabled,omitempty"`
	CreateDateTime             *int         `json:"createDateTime,omitempty"`
	CreatorID                  *string      `json:"creatorId,omitempty"`
	CreatorUsername            *string      `json:"creatorUsername,omitempty"`
	Description                *string      `json:"description,omitempty"`
	ID                         *string      `json:"id,omitempty"`
	ModifiedDateTime           *int         `json:"modifiedDateTime,omitempty"`
	ModifierID                 *string      `json:"modifierId,omitempty"`
	ModifierUsername           *string      `json:"modifierUsername,omitempty"`
	Name                       *string      `json:"name,omitempty"`
	Whitelist                  []*Whitelist `json:"whitelist,omitempty"`
	ZoneID                     *string      `json:"zoneId,omitempty"`
}

type ClientIsolationWhitelistArray struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	FirstIndex       *int    `json:"firstIndex,omitempty"`
	HasMore          *bool   `json:"hasMore,omitempty"`
	List             []*List `json:"list,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	TotalCount       *int    `json:"totalCount,omitempty"`
}

type CmProtocolOptionContent struct {
	CmpDHCPOpt43Subcode          *float64 `json:"cmpDhcpOpt43Subcode,omitempty"`
	CmpDHCPOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`
	CmpRecipient                 *string  `json:"cmpRecipient,omitempty"`
	CmpServerAddr                *string  `json:"cmpServerAddr,omitempty"`
	CmpServerPath                *string  `json:"cmpServerPath,omitempty"`
}

type CoreNetworkGateway struct {
	KeepAlivePeriod  *int    `json:"keepAlivePeriod,omitempty"`
	KeepAliveRetry   *int    `json:"keepAliveRetry,omitempty"`
	PrimaryGateway   *string `json:"primaryGateway,omitempty"`
	SecondaryGateway *string `json:"secondaryGateway,omitempty"`
	TunnelMTU        *string `json:"tunnelMTU,omitempty"`
	TunnelMTUSize    *int    `json:"tunnelMTUSize,omitempty"`
}

type CreateAccountingProfile struct {
	Description   *string          `json:"description,omitempty"`
	DomainID      *string          `json:"domainId,omitempty"`
	MvnoID        *string          `json:"mvnoId,omitempty"`
	Name          *string          `json:"name,omitempty"`
	RealmMappings []*RealmMappings `json:"realmMappings,omitempty"`
}

type CreateAuthenticationProfile struct {
	AaaSuppportEnabled *bool             `json:"aaaSuppportEnabled,omitempty"`
	Description        *string           `json:"description,omitempty"`
	DomainID           *string           `json:"domainId,omitempty"`
	GppSuppportEnabled *bool             `json:"gppSuppportEnabled,omitempty"`
	H20SuppportEnabled *bool             `json:"h20SuppportEnabled,omitempty"`
	MvnoID             *string           `json:"mvnoId,omitempty"`
	Name               *string           `json:"name,omitempty"`
	RealmMappings      []*RealmMappings  `json:"realmMappings,omitempty"`
	TtgCommonSetting   *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

type CreateBonjourFencingPolicy struct {
	BonjourFencingRuleList        []*BonjourFencingRuleList        `json:"bonjourFencingRuleList,omitempty"`
	BonjourFencingRuleMappingList []*BonjourFencingRuleMappingList `json:"bonjourFencingRuleMappingList,omitempty"`
	Description                   *string                          `json:"description,omitempty"`
	Name                          *string                          `json:"name,omitempty"`
}

type CreateBridgeProfile struct {
	Description *string    `json:"description,omitempty"`
	DHCPRelay   *DHCPRelay `json:"dhcpRelay,omitempty"`
	DomainID    *string    `json:"domainId,omitempty"`
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type CreateClientIsolationWhitelist struct {
	ClientIsolationAutoEnabled *bool        `json:"clientIsolationAutoEnabled,omitempty"`
	Description                *string      `json:"description,omitempty"`
	Name                       *string      `json:"name,omitempty"`
	Whitelist                  []*Whitelist `json:"whitelist,omitempty"`
}

type CreateDHCPProfile struct {
	Description      *string `json:"description,omitempty"`
	LeaseTimeHours   *int    `json:"leaseTimeHours,omitempty"`
	LeaseTimeMinutes *int    `json:"leaseTimeMinutes,omitempty"`
	Name             *string `json:"name,omitempty"`
	PoolEndIP        *string `json:"poolEndIp,omitempty"`
	PoolStartIP      *string `json:"poolStartIp,omitempty"`
	PrimaryDNSIP     *string `json:"primaryDnsIp,omitempty"`
	SecondaryDNSIP   *string `json:"secondaryDnsIp,omitempty"`
	SubnetMask       *string `json:"subnetMask,omitempty"`
	SubnetNetworkIP  *string `json:"subnetNetworkIp,omitempty"`
	VlanID           *int    `json:"vlanId,omitempty"`
}

type CreateDNSServerProfile struct {
	Description *string `json:"description,omitempty"`
	DomainID    *string `json:"domainId,omitempty"`
	MvnoID      *string `json:"mvnoId,omitempty"`
	Name        *string `json:"name,omitempty"`
	PrimaryIP   *string `json:"primaryIp,omitempty"`
	SecondaryIP *string `json:"secondaryIp,omitempty"`
	TertiaryIP  *string `json:"tertiaryIp,omitempty"`
}

type CreateIpsecProfile struct {
	AdvancedOption         *AdvancedOption         `json:"advancedOption,omitempty"`
	AuthType               *string                 `json:"authType,omitempty"`
	CmProtocolOption       *CmProtocolOption       `json:"cmProtocolOption,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	DomainID               *string                 `json:"domainId,omitempty"`
	EspRekeyTime           *float64                `json:"espRekeyTime,omitempty"`
	EspRekeyTimeUnit       *string                 `json:"espRekeyTimeUnit,omitempty"`
	EspSecurityAssociation *EspSecurityAssociation `json:"espSecurityAssociation,omitempty"`
	ID                     *string                 `json:"id,omitempty"`
	IkeRekeyTime           *float64                `json:"ikeRekeyTime,omitempty"`
	IkeRekeyTimeUnit       *string                 `json:"ikeRekeyTimeUnit,omitempty"`
	IkeSecurityAssociation *IkeSecurityAssociation `json:"ikeSecurityAssociation,omitempty"`
	IPMode                 *string                 `json:"ipMode,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
	PreSharedKey           *string                 `json:"preSharedKey,omitempty"`
	ServerAddr             *string                 `json:"serverAddr,omitempty"`
}

type CreateLOGREProfile struct {
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway,omitempty"`
	Description        *string             `json:"description,omitempty"`
	DHCPRelay          *DHCPRelay          `json:"dhcpRelay,omitempty"`
	DomainID           *string             `json:"domainId,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
}

type CreatePrecedenceProfile struct {
	DomainID               *string                   `json:"domainId,omitempty"`
	Name                   *string                   `json:"name,omitempty"`
	RateLimitingPrecedence []*RateLimitingPrecedence `json:"rateLimitingPrecedence,omitempty"`
	VlanPrecedence         []*VlanPrecedence         `json:"vlanPrecedence,omitempty"`
}

type CreateRogueApPolicy struct {
	Description *string  `json:"description,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Rules       []*Rules `json:"rules,omitempty"`
}

type CreateRtlsProfile struct {
	EkahauEnabled  *bool   `json:"ekahauEnabled,omitempty"`
	EkahauIP       *string `json:"ekahauIp,omitempty"`
	EkahauPort     *int    `json:"ekahauPort,omitempty"`
	ID             *string `json:"id,omitempty"`
	StanleyEnabled *bool   `json:"stanleyEnabled,omitempty"`
}

type CreateRuckusGREProfile struct {
	Description          *string `json:"description,omitempty"`
	DomainID             *string `json:"domainId,omitempty"`
	ID                   *string `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	TunnelEncryption     *string `json:"tunnelEncryption,omitempty"`
	TunnelMode           *string `json:"tunnelMode,omitempty"`
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`
	TunnelMtuSize        *int    `json:"tunnelMtuSize,omitempty"`
}

type CreateSoftGREProfile struct {
	Description             *string `json:"description,omitempty"`
	DomainID                *string `json:"domainId,omitempty"`
	ForceDisassociateClient *bool   `json:"forceDisassociateClient,omitempty"`
	ID                      *string `json:"id,omitempty"`
	IPMode                  *string `json:"ipMode,omitempty"`
	KeepAlivePeriod         *int    `json:"keepAlivePeriod,omitempty"`
	KeepAliveRetry          *int    `json:"keepAliveRetry,omitempty"`
	Name                    *string `json:"name,omitempty"`
	PrimaryGateway          *string `json:"primaryGateway,omitempty"`
	SecondaryGateway        *string `json:"secondaryGateway,omitempty"`
	TunnelMtuAutoEnabled    *string `json:"tunnelMtuAutoEnabled,omitempty"`
	TunnelMtuSize           *int    `json:"tunnelMtuSize,omitempty"`
}

type CreateTrafficClassProfile struct {
	Description    *string           `json:"description,omitempty"`
	Name           *string           `json:"name,omitempty"`
	TrafficClasses []*TrafficClasses `json:"trafficClasses,omitempty"`
}

type CreateTtgpdgProfile struct {
	ApnForwardingRealms  []*ApnForwardingRealms `json:"apnForwardingRealms,omitempty"`
	ApnRealms            []*ApnRealms           `json:"apnRealms,omitempty"`
	CommonSetting        *CommonSetting         `json:"commonSetting,omitempty"`
	DefaultNoMatchingAPN *string                `json:"defaultNoMatchingAPN,omitempty"`
	DefaultNoRealmAPN    *string                `json:"defaultNoRealmAPN,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DHCPRelay            *DHCPRelay             `json:"dhcpRelay,omitempty"`
	DomainID             *string                `json:"domainId,omitempty"`
	ID                   *string                `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
}

type CreateUserTrafficProfile struct {
	AppPolicyID          *string               `json:"appPolicyId,omitempty"`
	DefaultAction        *string               `json:"defaultAction,omitempty"`
	Description          *string               `json:"description,omitempty"`
	DomainID             *string               `json:"domainId,omitempty"`
	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
	IPAclRules           []*IPAclRules         `json:"ipAclRules,omitempty"`
	MvnoID               *string               `json:"mvnoId,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	QmAppPolicyID        *string               `json:"qmAppPolicyId,omitempty"`
	UplinkRateLimiting   *UplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
	URLFilteringPolicyID *string               `json:"urlFilteringPolicyId,omitempty"`
}

type CreateZoneAffinityProfile struct {
	Description      *string  `json:"description,omitempty"`
	Name             *string  `json:"name,omitempty"`
	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
}

type DataPlaneL3RoamingData struct {
	Activated       *int    `json:"activated,omitempty"`
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	Key             *string `json:"key,omitempty"`
	Name            *string `json:"name,omitempty"`
	SubCriteriaType *string `json:"subCriteriaType,omitempty"`
	Value           *string `json:"value,omitempty"`
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
	DHCPOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
	Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
	Subopt1Format       *string `json:"subopt1Format,omitempty"`
	Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
	Subopt2Format       *string `json:"subopt2Format,omitempty"`
	Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
	Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
	Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
	Subopt151Format     *string `json:"subopt151Format,omitempty"`
}

type DHCPProfileList struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	FirstIndex       *int    `json:"firstIndex,omitempty"`
	HasMore          *bool   `json:"hasMore,omitempty"`
	List             []*List `json:"list,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	TotalCount       *int    `json:"totalCount,omitempty"`
}

type DHCPRelayNoRelayTunnel struct {
	DHCPOption82     *DHCPOption82 `json:"dhcpOption82,omitempty"`
	DHCPRelayEnabled *bool         `json:"dhcpRelayEnabled,omitempty"`
	DHCPServer1      *string       `json:"dhcpServer1,omitempty"`
	DHCPServer2      *string       `json:"dhcpServer2,omitempty"`
	RelayBothEnabled *bool         `json:"relayBothEnabled,omitempty"`
}

type DNSServerProfile struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	Description      *string `json:"description,omitempty"`
	DomainID         *string `json:"domainId,omitempty"`
	ID               *string `json:"id,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	MvnoID           *string `json:"mvnoId,omitempty"`
	Name             *string `json:"name,omitempty"`
	PrimaryIP        *string `json:"primaryIp,omitempty"`
	SecondaryIP      *string `json:"secondaryIp,omitempty"`
	TertiaryIP       *string `json:"tertiaryIp,omitempty"`
}

type DNSServerProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type DownlinkRateLimiting struct {
	DownlinkRateLimitingBps     *string `json:"downlinkRateLimitingBps,omitempty"`
	DownlinkRateLimitingEnabled *bool   `json:"downlinkRateLimitingEnabled,omitempty"`
}

type EspProposal struct {
	AuthAlg *string `json:"authAlg,omitempty"`
	DhGroup *string `json:"dhGroup,omitempty"`
	EncAlg  *string `json:"encAlg,omitempty"`
}

type EspSecurityAssociationContent struct {
	EspProposals    []*EspProposals `json:"espProposals,omitempty"`
	EspProposalType *string         `json:"espProposalType,omitempty"`
}

type FlexiVpnProfile struct {
	DestinationZoneAffinityID   *string `json:"destinationZoneAffinityId,omitempty"`
	DestinationZoneAffinityName *string `json:"destinationZoneAffinityName,omitempty"`
	DomainID                    *string `json:"domainId,omitempty"`
	SourceZoneAffinityID        *string `json:"sourceZoneAffinityId,omitempty"`
	SourceZoneAffinityName      *string `json:"sourceZoneAffinityName,omitempty"`
	WLANID                      *string `json:"wlanId,omitempty"`
	WLANName                    *string `json:"wlanName,omitempty"`
	ZoneID                      *string `json:"zoneId,omitempty"`
	ZoneName                    *string `json:"zoneName,omitempty"`
}

type FlexiVpnProfileList struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	FirstIndex       *int    `json:"firstIndex,omitempty"`
	HasMore          *bool   `json:"hasMore,omitempty"`
	List             []*List `json:"list,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	TotalCount       *int    `json:"totalCount,omitempty"`
}

type GetL3RoamingConfig struct {
	DataPlanes     []*DataPlanes `json:"dataPlanes,omitempty"`
	FeatureEnabled *int          `json:"featureEnabled,omitempty"`
}

type Hs20FriendlyName struct {
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type Hs20Operator struct {
	Certificate      *common.GenericRef `json:"certificate,omitempty"`
	CreateDateTime   *int               `json:"createDateTime,omitempty"`
	CreatorID        *string            `json:"creatorId,omitempty"`
	CreatorUsername  *string            `json:"creatorUsername,omitempty"`
	Description      *string            `json:"description,omitempty"`
	DomainID         *string            `json:"domainId,omitempty"`
	DomainNames      []string           `json:"domainNames,omitempty"`
	FriendlyNames    []*FriendlyNames   `json:"friendlyNames,omitempty"`
	ID               *string            `json:"id,omitempty"`
	ModifiedDateTime *int               `json:"modifiedDateTime,omitempty"`
	ModifierID       *string            `json:"modifierId,omitempty"`
	ModifierUsername *string            `json:"modifierUsername,omitempty"`
	Name             *string            `json:"name,omitempty"`
}

type Hs20OperatorList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type Hs20Provider struct {
	Accountings      []*Accountings     `json:"accountings,omitempty"`
	Authentications  []*Authentications `json:"authentications,omitempty"`
	CreateDateTime   *int               `json:"createDateTime,omitempty"`
	CreatorID        *string            `json:"creatorId,omitempty"`
	CreatorUsername  *string            `json:"creatorUsername,omitempty"`
	Description      *string            `json:"description,omitempty"`
	DomainID         *string            `json:"domainId,omitempty"`
	HomeOis          []*HomeOis         `json:"homeOis,omitempty"`
	ID               *string            `json:"id,omitempty"`
	ModifiedDateTime *int               `json:"modifiedDateTime,omitempty"`
	ModifierID       *string            `json:"modifierId,omitempty"`
	ModifierUsername *string            `json:"modifierUsername,omitempty"`
	Name             *string            `json:"name,omitempty"`
	Osu              *Osu               `json:"osu,omitempty"`
	Plmns            []*Plmns           `json:"plmns,omitempty"`
	Realms           []*Realms          `json:"realms,omitempty"`
}

type Hs20ProviderList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type IkeProposal struct {
	AuthAlg *string `json:"authAlg,omitempty"`
	DhGroup *string `json:"dhGroup,omitempty"`
	EncAlg  *string `json:"encAlg,omitempty"`
	PrfAlg  *string `json:"prfAlg,omitempty"`
}

type IkeSecurityAssociationContent struct {
	IkeProposals    []*IkeProposals `json:"ikeProposals,omitempty"`
	IkeProposalType *string         `json:"ikeProposalType,omitempty"`
}

type IPAclRules struct {
	Action                      *string  `json:"action,omitempty"`
	CustomProtocol              *int     `json:"customProtocol,omitempty"`
	Description                 *string  `json:"description,omitempty"`
	DestinationIP               *string  `json:"destinationIp,omitempty"`
	DestinationIPMask           *string  `json:"destinationIpMask,omitempty"`
	DestinationIPV6             *string  `json:"destinationIpV6,omitempty"`
	DestinationMaxPort          *int     `json:"destinationMaxPort,omitempty"`
	DestinationMinPort          *int     `json:"destinationMinPort,omitempty"`
	Direction                   *string  `json:"direction,omitempty"`
	DownlinkRateLimitingEnabled *bool    `json:"downlinkRateLimitingEnabled,omitempty"`
	DownlinkRateLimitingMbps    *float64 `json:"downlinkRateLimitingMbps,omitempty"`
	EnableDestinationIPSubnet   *bool    `json:"enableDestinationIpSubnet,omitempty"`
	EnableDestinationPortRange  *bool    `json:"enableDestinationPortRange,omitempty"`
	EnableDestinationV6Prefix   *bool    `json:"enableDestinationV6Prefix,omitempty"`
	EnableSourceIPSubnet        *bool    `json:"enableSourceIpSubnet,omitempty"`
	EnableSourcePortRange       *bool    `json:"enableSourcePortRange,omitempty"`
	EnableSourceV6Prefix        *bool    `json:"enableSourceV6Prefix,omitempty"`
	IPType                      *string  `json:"ipType,omitempty"`
	Priority                    *int     `json:"priority,omitempty"`
	Protocol                    *string  `json:"protocol,omitempty"`
	SourceIP                    *string  `json:"sourceIp,omitempty"`
	SourceIPMask                *string  `json:"sourceIpMask,omitempty"`
	SourceIPV6                  *string  `json:"sourceIpV6,omitempty"`
	SourceMaxPort               *int     `json:"sourceMaxPort,omitempty"`
	SourceMinPort               *int     `json:"sourceMinPort,omitempty"`
	UplinkRateLimitingEnabled   *bool    `json:"uplinkRateLimitingEnabled,omitempty"`
	UplinkRateLimitingMbps      *float64 `json:"uplinkRateLimitingMbps,omitempty"`
}

type IpsecProfile struct {
	AdvancedOption         *AdvancedOption         `json:"advancedOption,omitempty"`
	AuthType               *string                 `json:"authType,omitempty"`
	CmProtocolOption       *CmProtocolOption       `json:"cmProtocolOption,omitempty"`
	CreateDateTime         *int                    `json:"createDateTime,omitempty"`
	CreatorID              *string                 `json:"creatorId,omitempty"`
	CreatorUsername        *string                 `json:"creatorUsername,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	DomainID               *string                 `json:"domainId,omitempty"`
	EspRekeyTime           *float64                `json:"espRekeyTime,omitempty"`
	EspRekeyTimeUnit       *string                 `json:"espRekeyTimeUnit,omitempty"`
	EspSecurityAssociation *EspSecurityAssociation `json:"espSecurityAssociation,omitempty"`
	ID                     *string                 `json:"id,omitempty"`
	IkeRekeyTime           *float64                `json:"ikeRekeyTime,omitempty"`
	IkeRekeyTimeUnit       *string                 `json:"ikeRekeyTimeUnit,omitempty"`
	IkeSecurityAssociation *IkeSecurityAssociation `json:"ikeSecurityAssociation,omitempty"`
	IPMode                 *string                 `json:"ipMode,omitempty"`
	ModifiedDateTime       *int                    `json:"modifiedDateTime,omitempty"`
	ModifierID             *string                 `json:"modifierId,omitempty"`
	ModifierUsername       *string                 `json:"modifierUsername,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
	PreSharedKey           *string                 `json:"preSharedKey,omitempty"`
	ServerAddr             *string                 `json:"serverAddr,omitempty"`
}

type IpsecProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type LOGREProfile struct {
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway,omitempty"`
	CreateDateTime     *int                `json:"createDateTime,omitempty"`
	CreatorID          *string             `json:"creatorId,omitempty"`
	CreatorUsername    *string             `json:"creatorUsername,omitempty"`
	Description        *string             `json:"description,omitempty"`
	DHCPRelay          *DHCPRelay          `json:"dhcpRelay,omitempty"`
	DomainID           *string             `json:"domainId,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	ModifiedDateTime   *int                `json:"modifiedDateTime,omitempty"`
	ModifierID         *string             `json:"modifierId,omitempty"`
	ModifierUsername   *string             `json:"modifierUsername,omitempty"`
	Name               *string             `json:"name,omitempty"`
}

type LOGREProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type LbsProfile struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	Description      *string `json:"description,omitempty"`
	DomainID         *string `json:"domainId,omitempty"`
	ID               *string `json:"id,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	Name             *string `json:"name,omitempty"`
	Password         *string `json:"password,omitempty"`
	Port             *int    `json:"port,omitempty"`
	URL              *string `json:"url,omitempty"`
	Venue            *string `json:"venue,omitempty"`
}

type LbsProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type ModifyAccountingProfile struct {
	Description   *string          `json:"description,omitempty"`
	DomainID      *string          `json:"domainId,omitempty"`
	ID            *string          `json:"id,omitempty"`
	MvnoID        *string          `json:"mvnoId,omitempty"`
	Name          *string          `json:"name,omitempty"`
	RealmMappings []*RealmMappings `json:"realmMappings,omitempty"`
}

type ModifyAuthenticationProfile struct {
	AaaSuppportEnabled *bool             `json:"aaaSuppportEnabled,omitempty"`
	Description        *string           `json:"description,omitempty"`
	DomainID           *string           `json:"domainId,omitempty"`
	GppSuppportEnabled *bool             `json:"gppSuppportEnabled,omitempty"`
	H20SuppportEnabled *bool             `json:"h20SuppportEnabled,omitempty"`
	ID                 *string           `json:"id,omitempty"`
	MvnoID             *string           `json:"mvnoId,omitempty"`
	Name               *string           `json:"name,omitempty"`
	RealmMappings      []*RealmMappings  `json:"realmMappings,omitempty"`
	TtgCommonSetting   *TtgCommonSetting `json:"ttgCommonSetting,omitempty"`
}

type ModifyBlockClient struct {
	Description *string `json:"description,omitempty"`
	Mac         *string `json:"mac,omitempty"`
}

type ModifyBonjourFencingPolicy struct {
	BonjourFencingRuleList        []*BonjourFencingRuleList        `json:"bonjourFencingRuleList,omitempty"`
	BonjourFencingRuleMappingList []*BonjourFencingRuleMappingList `json:"bonjourFencingRuleMappingList,omitempty"`
	Description                   *string                          `json:"description,omitempty"`
	Name                          *string                          `json:"name,omitempty"`
}

type ModifyBridgeProfile struct {
	Description *string    `json:"description,omitempty"`
	DHCPRelay   *DHCPRelay `json:"dhcpRelay,omitempty"`
	DomainID    *string    `json:"domainId,omitempty"`
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type ModifyClientIsolationWhitelist struct {
	ClientIsolationAutoEnabled *bool        `json:"clientIsolationAutoEnabled,omitempty"`
	Description                *string      `json:"description,omitempty"`
	Name                       *string      `json:"name,omitempty"`
	Whitelist                  []*Whitelist `json:"whitelist,omitempty"`
}

type ModifyDNSServerProfile struct {
	Description *string `json:"description,omitempty"`
	DomainID    *string `json:"domainId,omitempty"`
	ID          *string `json:"id,omitempty"`
	MvnoID      *string `json:"mvnoId,omitempty"`
	Name        *string `json:"name,omitempty"`
	PrimaryIP   *string `json:"primaryIp,omitempty"`
	SecondaryIP *string `json:"secondaryIp,omitempty"`
	TertiaryIP  *string `json:"tertiaryIp,omitempty"`
}

type ModifyHS20Operator struct {
	Certificate   *common.GenericRef `json:"certificate,omitempty"`
	Description   *string            `json:"description,omitempty"`
	DomainID      *string            `json:"domainId,omitempty"`
	DomainNames   []string           `json:"domainNames,omitempty"`
	FriendlyNames []*FriendlyNames   `json:"friendlyNames,omitempty"`
	ID            *string            `json:"id,omitempty"`
	Name          *string            `json:"name,omitempty"`
}

type ModifyIPAclRules struct {
	Action                      *string  `json:"action,omitempty"`
	CustomProtocol              *int     `json:"customProtocol,omitempty"`
	Description                 *string  `json:"description,omitempty"`
	DestinationIP               *string  `json:"destinationIp,omitempty"`
	DestinationIPMask           *string  `json:"destinationIpMask,omitempty"`
	DestinationIPV6             *string  `json:"destinationIpV6,omitempty"`
	DestinationMaxPort          *int     `json:"destinationMaxPort,omitempty"`
	DestinationMinPort          *int     `json:"destinationMinPort,omitempty"`
	Direction                   *string  `json:"direction,omitempty"`
	DownlinkRateLimitingEnabled *bool    `json:"downlinkRateLimitingEnabled,omitempty"`
	DownlinkRateLimitingMbps    *float64 `json:"downlinkRateLimitingMbps,omitempty"`
	EnableDestinationIPSubnet   *bool    `json:"enableDestinationIpSubnet,omitempty"`
	EnableDestinationPortRange  *bool    `json:"enableDestinationPortRange,omitempty"`
	EnableDestinationV6Prefix   *bool    `json:"enableDestinationV6Prefix,omitempty"`
	EnableSourceIPSubnet        *bool    `json:"enableSourceIpSubnet,omitempty"`
	EnableSourcePortRange       *bool    `json:"enableSourcePortRange,omitempty"`
	EnableSourceV6Prefix        *bool    `json:"enableSourceV6Prefix,omitempty"`
	IPType                      *string  `json:"ipType,omitempty"`
	Priority                    *int     `json:"priority,omitempty"`
	Protocol                    *string  `json:"protocol,omitempty"`
	SourceIP                    *string  `json:"sourceIp,omitempty"`
	SourceIPMask                *string  `json:"sourceIpMask,omitempty"`
	SourceIPV6                  *string  `json:"sourceIpV6,omitempty"`
	SourceMaxPort               *int     `json:"sourceMaxPort,omitempty"`
	SourceMinPort               *int     `json:"sourceMinPort,omitempty"`
	UplinkRateLimitingEnabled   *bool    `json:"uplinkRateLimitingEnabled,omitempty"`
	UplinkRateLimitingMbps      *float64 `json:"uplinkRateLimitingMbps,omitempty"`
}

type ModifyIpsecProfile struct {
	AdvancedOption         *AdvancedOption         `json:"advancedOption,omitempty"`
	AuthType               *string                 `json:"authType,omitempty"`
	CmProtocolOption       *CmProtocolOption       `json:"cmProtocolOption,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	DomainID               *string                 `json:"domainId,omitempty"`
	EspRekeyTime           *float64                `json:"espRekeyTime,omitempty"`
	EspRekeyTimeUnit       *string                 `json:"espRekeyTimeUnit,omitempty"`
	EspSecurityAssociation *EspSecurityAssociation `json:"espSecurityAssociation,omitempty"`
	ID                     *string                 `json:"id,omitempty"`
	IkeRekeyTime           *float64                `json:"ikeRekeyTime,omitempty"`
	IkeRekeyTimeUnit       *string                 `json:"ikeRekeyTimeUnit,omitempty"`
	IkeSecurityAssociation *IkeSecurityAssociation `json:"ikeSecurityAssociation,omitempty"`
	IPMode                 *string                 `json:"ipMode,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
	PreSharedKey           *string                 `json:"preSharedKey,omitempty"`
	ServerAddr             *string                 `json:"serverAddr,omitempty"`
}

type ModifyLOGREProfile struct {
	CoreNetworkGateway *CoreNetworkGateway `json:"coreNetworkGateway,omitempty"`
	Description        *string             `json:"description,omitempty"`
	DHCPRelay          *DHCPRelay          `json:"dhcpRelay,omitempty"`
	DomainID           *string             `json:"domainId,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
}

type ModifyRuckusGREProfile struct {
	Description          *string `json:"description,omitempty"`
	DomainID             *string `json:"domainId,omitempty"`
	ID                   *string `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	TunnelEncryption     *string `json:"tunnelEncryption,omitempty"`
	TunnelMode           *string `json:"tunnelMode,omitempty"`
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`
	TunnelMtuSize        *int    `json:"tunnelMtuSize,omitempty"`
}

type ModifySoftGREProfile struct {
	Description             *string `json:"description,omitempty"`
	DomainID                *string `json:"domainId,omitempty"`
	ForceDisassociateClient *bool   `json:"forceDisassociateClient,omitempty"`
	ID                      *string `json:"id,omitempty"`
	KeepAlivePeriod         *int    `json:"keepAlivePeriod,omitempty"`
	KeepAliveRetry          *int    `json:"keepAliveRetry,omitempty"`
	Name                    *string `json:"name,omitempty"`
	PrimaryGateway          *string `json:"primaryGateway,omitempty"`
	SecondaryGateway        *string `json:"secondaryGateway,omitempty"`
	TunnelMtuAutoEnabled    *string `json:"tunnelMtuAutoEnabled,omitempty"`
	TunnelMtuSize           *int    `json:"tunnelMtuSize,omitempty"`
}

type ModifyUserTrafficProfile struct {
	AppPolicyID          *string               `json:"appPolicyId,omitempty"`
	DefaultAction        *string               `json:"defaultAction,omitempty"`
	Description          *string               `json:"description,omitempty"`
	DomainID             *string               `json:"domainId,omitempty"`
	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
	ID                   *string               `json:"id,omitempty"`
	IPAclRules           []*IPAclRules         `json:"ipAclRules,omitempty"`
	MvnoID               *string               `json:"mvnoId,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	QmAppPolicyID        *string               `json:"qmAppPolicyId,omitempty"`
	UplinkRateLimiting   *UplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
	URLFilteringPolicyID *string               `json:"urlFilteringPolicyId,omitempty"`
}

type ModifyZoneAffinityProfile struct {
	Description      *string  `json:"description,omitempty"`
	Name             *string  `json:"name,omitempty"`
	ZoneAffinityList []string `json:"zoneAffinityList,omitempty"`
}

type PrecedenceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type PrecedenceListType struct {
	DomainID               *string                   `json:"domainId,omitempty"`
	ID                     *string                   `json:"id,omitempty"`
	Name                   *string                   `json:"name,omitempty"`
	RateLimitingPrecedence []*RateLimitingPrecedence `json:"rateLimitingPrecedence,omitempty"`
	VlanPrecedence         []*VlanPrecedence         `json:"vlanPrecedence,omitempty"`
}

type ProfileCloneRequest struct {
	NewID   *string `json:"newId,omitempty"`
	NewName *string `json:"newName,omitempty"`
	OldID   *string `json:"oldId,omitempty"`
	OldName *string `json:"oldName,omitempty"`
}

type ProfileCloneResponse struct {
	*ProfileCloneRequest
}

type ProfileList struct {
	CreateDateTime   *int                 `json:"createDateTime,omitempty"`
	CreatorID        *string              `json:"creatorId,omitempty"`
	CreatorUsername  *string              `json:"creatorUsername,omitempty"`
	Extra            *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex       *int                 `json:"firstIndex,omitempty"`
	HasMore          *bool                `json:"hasMore,omitempty"`
	List             []*List              `json:"list,omitempty"`
	ModifiedDateTime *int                 `json:"modifiedDateTime,omitempty"`
	ModifierID       *string              `json:"modifierId,omitempty"`
	ModifierUsername *string              `json:"modifierUsername,omitempty"`
	TotalCount       *int                 `json:"totalCount,omitempty"`
}

type ProfileListType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ProviderAccounting struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Realm       *string `json:"realm,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
}

type ProviderAuthentication struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Realm       *string `json:"realm,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
	VlanID      *int    `json:"vlanId,omitempty"`
}

type ProviderEAPAuthSetting struct {
	Info       *string `json:"info,omitempty"`
	Type       *string `json:"type,omitempty"`
	VendorID   *int    `json:"vendorId,omitempty"`
	VendorType *int    `json:"vendorType,omitempty"`
}

type ProviderEAPMethod struct {
	AuthSettings []*AuthSettings `json:"authSettings,omitempty"`
	Type         *string         `json:"type,omitempty"`
}

type ProviderExternalOSU struct {
	CommonLanguageIcon       *string                     `json:"commonLanguageIcon,omitempty"`
	OsuNaiRealm              *string                     `json:"osuNaiRealm,omitempty"`
	OsuServiceURL            *string                     `json:"osuServiceUrl,omitempty"`
	ProvisioningProtocals    []string                    `json:"provisioningProtocals,omitempty"`
	SubscriptionDescriptions []*SubscriptionDescriptions `json:"subscriptionDescriptions,omitempty"`
	WhitelistedDomains       []string                    `json:"whitelistedDomains,omitempty"`
}

type ProviderHomeOIs struct {
	Name *string `json:"name,omitempty"`
	Oi   *string `json:"oi,omitempty"`
}

type ProviderInternalOSU struct {
	Certificate              *common.GenericRef          `json:"certificate,omitempty"`
	CommonLanguageIcon       *string                     `json:"commonLanguageIcon,omitempty"`
	OsuAuthServices          []*OsuAuthServices          `json:"osuAuthServices,omitempty"`
	OsuPortal                *OsuPortal                  `json:"osuPortal,omitempty"`
	ProvisioningFormat       *string                     `json:"provisioningFormat,omitempty"`
	ProvisioningProtocals    []string                    `json:"provisioningProtocals,omitempty"`
	ProvisioningUpdateType   *string                     `json:"provisioningUpdateType,omitempty"`
	SubscriptionDescriptions []*SubscriptionDescriptions `json:"subscriptionDescriptions,omitempty"`
	WhitelistedDomains       []string                    `json:"whitelistedDomains,omitempty"`
}

type ProviderInternalOSUOsuAuthServicesType struct {
	CredentialType *string `json:"credentialType,omitempty"`
	Expiration     *int    `json:"expiration,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	Realm          *string `json:"realm,omitempty"`
}

type ProviderInternalOSUOsuPortalType struct {
	ExternalURL       *string            `json:"externalUrl,omitempty"`
	InternalOSUPortal *common.GenericRef `json:"internalOSUPortal,omitempty"`
	Type              *string            `json:"type,omitempty"`
}

type ProviderOnlineSignup struct {
	ExternalOSU *ExternalOSU `json:"externalOSU,omitempty"`
	InternalOSU *InternalOSU `json:"internalOSU,omitempty"`
	Type        *string      `json:"type,omitempty"`
}

type ProviderPLMN struct {
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
}

type ProviderRealm struct {
	EapMethods []*EapMethods `json:"eapMethods,omitempty"`
	Encoding   *string       `json:"encoding,omitempty"`
	Name       *string       `json:"name,omitempty"`
}

type ProviderSubscriptionDescription struct {
	Description *string `json:"description,omitempty"`
	Icon        *string `json:"icon,omitempty"`
	Language    *string `json:"language,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type RateLimitingPrecedenceItem struct {
	Name     *string `json:"name,omitempty"`
	Priority *int    `json:"priority,omitempty"`
}

type RealmAuthServiceMapping struct {
	AuthorizationMethod *string `json:"authorizationMethod,omitempty"`
	DynamicVlanID       *int    `json:"dynamicVlanId,omitempty"`
	HostedAaaEnabled    *bool   `json:"hostedAaaEnabled,omitempty"`
	ID                  *string `json:"id,omitempty"`
	Name                *string `json:"name,omitempty"`
	Realm               *string `json:"realm,omitempty"`
	ServiceType         *string `json:"serviceType,omitempty"`
}

type ReturnZoneAffinityProfile struct {
	BaseDpVersion         *string  `json:"baseDpVersion,omitempty"`
	CreateDateTime        *int     `json:"createDateTime,omitempty"`
	CreatorID             *string  `json:"creatorId,omitempty"`
	CreatorUsername       *string  `json:"creatorUsername,omitempty"`
	Description           *string  `json:"description,omitempty"`
	ID                    *string  `json:"id,omitempty"`
	IsDpVersionConsistent *bool    `json:"isDpVersionConsistent,omitempty"`
	ModifiedDateTime      *int     `json:"modifiedDateTime,omitempty"`
	ModifierID            *string  `json:"modifierId,omitempty"`
	ModifierUsername      *string  `json:"modifierUsername,omitempty"`
	Name                  *string  `json:"name,omitempty"`
	ZoneAffinityList      []string `json:"zoneAffinityList,omitempty"`
}

type RogueApPolicy struct {
	CreateDateTime   *int     `json:"createDateTime,omitempty"`
	CreatorID        *string  `json:"creatorId,omitempty"`
	CreatorUsername  *string  `json:"creatorUsername,omitempty"`
	Description      *string  `json:"description,omitempty"`
	ID               *string  `json:"id,omitempty"`
	ModifiedDateTime *int     `json:"modifiedDateTime,omitempty"`
	ModifierID       *string  `json:"modifierId,omitempty"`
	ModifierUsername *string  `json:"modifierUsername,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Rules            []*Rules `json:"rules,omitempty"`
	ZoneID           *string  `json:"zoneId,omitempty"`
}

type RogueApPolicyList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type RogueApRuleList struct {
	Classification *string     `json:"classification,omitempty"`
	Name           *string     `json:"name,omitempty"`
	Priority       *int        `json:"priority,omitempty"`
	Type           *string     `json:"type,omitempty"`
	Value          interface{} `json:"value,omitempty"`
}

type RtlsProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type RuckusGREProfile struct {
	CreateDateTime       *int    `json:"createDateTime,omitempty"`
	CreatorID            *string `json:"creatorId,omitempty"`
	CreatorUsername      *string `json:"creatorUsername,omitempty"`
	Description          *string `json:"description,omitempty"`
	DomainID             *string `json:"domainId,omitempty"`
	ID                   *string `json:"id,omitempty"`
	ModifiedDateTime     *int    `json:"modifiedDateTime,omitempty"`
	ModifierID           *string `json:"modifierId,omitempty"`
	ModifierUsername     *string `json:"modifierUsername,omitempty"`
	Name                 *string `json:"name,omitempty"`
	TunnelEncryption     *string `json:"tunnelEncryption,omitempty"`
	TunnelMode           *string `json:"tunnelMode,omitempty"`
	TunnelMtuAutoEnabled *string `json:"tunnelMtuAutoEnabled,omitempty"`
	TunnelMtuSize        *int    `json:"tunnelMtuSize,omitempty"`
}

type RuckusGREProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type SoftGREProfile struct {
	CreateDateTime          *int    `json:"createDateTime,omitempty"`
	CreatorID               *string `json:"creatorId,omitempty"`
	CreatorUsername         *string `json:"creatorUsername,omitempty"`
	Description             *string `json:"description,omitempty"`
	DomainID                *string `json:"domainId,omitempty"`
	ForceDisassociateClient *bool   `json:"forceDisassociateClient,omitempty"`
	ID                      *string `json:"id,omitempty"`
	IPMode                  *string `json:"ipMode,omitempty"`
	KeepAlivePeriod         *int    `json:"keepAlivePeriod,omitempty"`
	KeepAliveRetry          *int    `json:"keepAliveRetry,omitempty"`
	ModifiedDateTime        *int    `json:"modifiedDateTime,omitempty"`
	ModifierID              *string `json:"modifierId,omitempty"`
	ModifierUsername        *string `json:"modifierUsername,omitempty"`
	Name                    *string `json:"name,omitempty"`
	PrimaryGateway          *string `json:"primaryGateway,omitempty"`
	SecondaryGateway        *string `json:"secondaryGateway,omitempty"`
	TunnelMtuAutoEnabled    *string `json:"tunnelMtuAutoEnabled,omitempty"`
	TunnelMtuSize           *int    `json:"tunnelMtuSize,omitempty"`
}

type SoftGREProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type TrafficClassProfileList struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	FirstIndex       *int    `json:"firstIndex,omitempty"`
	HasMore          *bool   `json:"hasMore,omitempty"`
	List             []*List `json:"list,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	TotalCount       *int    `json:"totalCount,omitempty"`
}

type TtgCommonSetting struct {
	InterimAcctInterval *int    `json:"interimAcctInterval,omitempty"`
	MobileCountryCode   *string `json:"mobileCountryCode,omitempty"`
	MobileNetworkCode   *string `json:"mobileNetworkCode,omitempty"`
	SessionIdleTimeout  *int    `json:"sessionIdleTimeout,omitempty"`
	SessionTimeout      *int    `json:"sessionTimeout,omitempty"`
}

type TtgpdgApnForwardingRealm struct {
	Apn       *string `json:"apn,omitempty"`
	ApnType   *string `json:"apnType,omitempty"`
	RouteType *string `json:"routeType,omitempty"`
}

type TtgpdgCommonSetting struct {
	AcctRetry        *int    `json:"acctRetry,omitempty"`
	AcctRetryTimeout *int    `json:"acctRetryTimeout,omitempty"`
	ApnFormat2GGSN   *string `json:"apnFormat2GGSN,omitempty"`
	ApnOIInUse       *bool   `json:"apnOIInUse,omitempty"`
	PdgUeIdleTimeout *int    `json:"pdgUeIdleTimeout,omitempty"`
}

type TtgpdgProfile struct {
	ApnForwardingRealms  []*ApnForwardingRealms `json:"apnForwardingRealms,omitempty"`
	ApnRealms            []*ApnRealms           `json:"apnRealms,omitempty"`
	CommonSetting        *CommonSetting         `json:"commonSetting,omitempty"`
	CreateDateTime       *int                   `json:"createDateTime,omitempty"`
	CreatorID            *string                `json:"creatorId,omitempty"`
	CreatorUsername      *string                `json:"creatorUsername,omitempty"`
	DefaultNoMatchingAPN *string                `json:"defaultNoMatchingAPN,omitempty"`
	DefaultNoRealmAPN    *string                `json:"defaultNoRealmAPN,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DHCPRelay            *DHCPRelay             `json:"dhcpRelay,omitempty"`
	DomainID             *string                `json:"domainId,omitempty"`
	ID                   *string                `json:"id,omitempty"`
	ModifiedDateTime     *int                   `json:"modifiedDateTime,omitempty"`
	ModifierID           *string                `json:"modifierId,omitempty"`
	ModifierUsername     *string                `json:"modifierUsername,omitempty"`
	Name                 *string                `json:"name,omitempty"`
}

type TtgpdgProfileConfiguration struct {
	ApnForwardingRealms  []*ApnForwardingRealms `json:"apnForwardingRealms,omitempty"`
	ApnRealms            []*ApnRealms           `json:"apnRealms,omitempty"`
	CommonSetting        *CommonSetting         `json:"commonSetting,omitempty"`
	DefaultNoMatchingAPN *string                `json:"defaultNoMatchingAPN,omitempty"`
	DefaultNoRealmAPN    *string                `json:"defaultNoRealmAPN,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DHCPRelay            *DHCPRelay             `json:"dhcpRelay,omitempty"`
	DomainID             *string                `json:"domainId,omitempty"`
	Name                 *string                `json:"name,omitempty"`
}

type TtgpdgProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type UpdateL3RoamingConfig struct {
	DataPlanes     []*DataPlanes `json:"dataPlanes,omitempty"`
	FeatureEnabled *int          `json:"featureEnabled,omitempty"`
}

type UpdatePrecedenceProfile struct {
	DomainID               *string                   `json:"domainId,omitempty"`
	Name                   *string                   `json:"name,omitempty"`
	RateLimitingPrecedence []*RateLimitingPrecedence `json:"rateLimitingPrecedence,omitempty"`
	VlanPrecedence         []*VlanPrecedence         `json:"vlanPrecedence,omitempty"`
}

type UpdateRogueApPolicy struct {
	Description *string  `json:"description,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Rules       []*Rules `json:"rules,omitempty"`
}

type UpdateRtlsProfile struct {
	EkahauEnabled  *bool   `json:"ekahauEnabled,omitempty"`
	EkahauIP       *string `json:"ekahauIp,omitempty"`
	EkahauPort     *int    `json:"ekahauPort,omitempty"`
	ID             *string `json:"id,omitempty"`
	StanleyEnabled *bool   `json:"stanleyEnabled,omitempty"`
}

type UplinkRateLimiting struct {
	UplinkRateLimitingBps     *string `json:"uplinkRateLimitingBps,omitempty"`
	UplinkRateLimitingEnabled *bool   `json:"uplinkRateLimitingEnabled,omitempty"`
}

type UserTrafficProfile struct {
	AppPolicyID          *string               `json:"appPolicyId,omitempty"`
	CreateDateTime       *int                  `json:"createDateTime,omitempty"`
	CreatorID            *string               `json:"creatorId,omitempty"`
	CreatorUsername      *string               `json:"creatorUsername,omitempty"`
	DefaultAction        *string               `json:"defaultAction,omitempty"`
	Description          *string               `json:"description,omitempty"`
	DomainID             *string               `json:"domainId,omitempty"`
	DownlinkRateLimiting *DownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
	ID                   *string               `json:"id,omitempty"`
	IPAclRules           []*IPAclRules         `json:"ipAclRules,omitempty"`
	IsFactoryDefault     *bool                 `json:"isFactoryDefault,omitempty"`
	ModifiedDateTime     *int                  `json:"modifiedDateTime,omitempty"`
	ModifierID           *string               `json:"modifierId,omitempty"`
	ModifierUsername     *string               `json:"modifierUsername,omitempty"`
	MvnoID               *string               `json:"mvnoId,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	QmAppPolicyID        *string               `json:"qmAppPolicyId,omitempty"`
	UplinkRateLimiting   *UplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
	URLFilteringPolicyID *string               `json:"urlFilteringPolicyId,omitempty"`
}

type UserTrafficProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type VdpProfile struct {
	CreateDateTime    *int    `json:"createDateTime,omitempty"`
	CreatorID         *string `json:"creatorId,omitempty"`
	CreatorUsername   *string `json:"creatorUsername,omitempty"`
	DataVlan          *int    `json:"dataVlan,omitempty"`
	ExtIP             *string `json:"extIp,omitempty"`
	FwVersion         *string `json:"fwVersion,omitempty"`
	IP                *string `json:"ip,omitempty"`
	Ipv6              *string `json:"ipv6,omitempty"`
	IsSupport         *bool   `json:"isSupport,omitempty"`
	LastSeenOn        *string `json:"lastSeenOn,omitempty"`
	Mac               *string `json:"mac,omitempty"`
	ManagedBy         *string `json:"managedBy,omitempty"`
	MgmtExtIP         *string `json:"mgmtExtIp,omitempty"`
	MgmtIP            *string `json:"mgmtIp,omitempty"`
	MgmtVlan          *int    `json:"mgmtVlan,omitempty"`
	Model             *string `json:"model,omitempty"`
	ModifiedDateTime  *int    `json:"modifiedDateTime,omitempty"`
	ModifierID        *string `json:"modifierId,omitempty"`
	ModifierUsername  *string `json:"modifierUsername,omitempty"`
	Name              *string `json:"name,omitempty"`
	RegistrationState *string `json:"registrationState,omitempty"`
	SerialNumber      *string `json:"serialNumber,omitempty"`
	Status            *string `json:"status,omitempty"`
	Uptime            *string `json:"uptime,omitempty"`
}

type VlanPrecedenceItem struct {
	Name     *string `json:"name,omitempty"`
	Priority *int    `json:"priority,omitempty"`
}

type ZoneAffinityProfileList struct {
	List []*List `json:"list,omitempty"`
}
