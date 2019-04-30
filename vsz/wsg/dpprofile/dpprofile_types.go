package dpprofile

// API Version: v8_0

type BulkDelete struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DpDHCPProfileBasicBO struct {
	DefaultLeaseTime   *int    `json:"defaultLeaseTime,omitempty"`
	Description        *string `json:"description,omitempty"`
	DomainName         *string `json:"domainName,omitempty"`
	PrimaryDNSServer   *string `json:"primaryDnsServer,omitempty"`
	ProfileID          *string `json:"profileId,omitempty"`
	ProfileName        *string `json:"profileName,omitempty"`
	SecondaryDNSServer *string `json:"secondaryDnsServer,omitempty"`
}

type DpDHCPProfileBasicBOList struct {
	FirstIndex *int                    `json:"firstIndex,omitempty"`
	HasMore    *bool                   `json:"hasMore,omitempty"`
	List       []*DpDHCPProfileBasicBO `json:"list,omitempty"`
	TotalCount *int                    `json:"totalCount,omitempty"`
}

type DpDHCPProfileHostBO struct {
	BroadcastAddress *string  `json:"broadcastAddress,omitempty"`
	Description      *string  `json:"description,omitempty"`
	DNSServers       []string `json:"dnsServers,omitempty"`
	DomainName       *string  `json:"domainName,omitempty"`
	FixedAddress     *string  `json:"fixedAddress,omitempty"`
	HardwareEthernet *string  `json:"hardwareEthernet,omitempty"`
	HostID           *string  `json:"hostId,omitempty"`
	HostName         *string  `json:"hostName,omitempty"`
	LeaseTime        *int     `json:"leaseTime,omitempty"`
	Name             *string  `json:"name,omitempty"`
	ProfileID        *string  `json:"profileId,omitempty"`
	Routers          []string `json:"routers,omitempty"`
}

type DpDHCPProfileHostBOList struct {
	FirstIndex *int                   `json:"firstIndex,omitempty"`
	HasMore    *bool                  `json:"hasMore,omitempty"`
	List       []*DpDHCPProfileHostBO `json:"list,omitempty"`
	TotalCount *int                   `json:"totalCount,omitempty"`
}

type DpDHCPProfileOptionBO struct {
	CodeNumber   *int    `json:"codeNumber,omitempty"`
	FunctionName *string `json:"functionName,omitempty"`
	Type         *string `json:"type,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type DpDHCPProfileOptionInstance struct {
	FunctionName *string `json:"functionName,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type DpDHCPProfileOptionSpaceApplyToBO struct {
	AppliedPoolNames []string                 `json:"appliedPoolNames,omitempty"`
	Description      *string                  `json:"description,omitempty"`
	Name             *string                  `json:"name,omitempty"`
	Options          []*DpDHCPProfileOptionBO `json:"options,omitempty"`
	SpaceID          *string                  `json:"spaceId,omitempty"`
}

type DpDHCPProfileOptionSpaceApplyToBOList struct {
	FirstIndex *int                                 `json:"firstIndex,omitempty"`
	HasMore    *bool                                `json:"hasMore,omitempty"`
	List       []*DpDHCPProfileOptionSpaceApplyToBO `json:"list,omitempty"`
	TotalCount *int                                 `json:"totalCount,omitempty"`
}

type DpDHCPProfileOptionSpaceBO struct {
	Description *string                  `json:"description,omitempty"`
	Name        *string                  `json:"name,omitempty"`
	Options     []*DpDHCPProfileOptionBO `json:"options,omitempty"`
	SpaceID     *string                  `json:"spaceId,omitempty"`
}

type DpDHCPProfileOptionSpaceInstance struct {
	Description *string                        `json:"description,omitempty"`
	Name        *string                        `json:"name,omitempty"`
	Options     []*DpDHCPProfileOptionInstance `json:"options,omitempty"`
	SpaceID     *string                        `json:"spaceId,omitempty"`
}

type DpDHCPProfilePoolBO struct {
	BroadcastAddress    *string                             `json:"broadcastAddress,omitempty"`
	Description         *string                             `json:"description,omitempty"`
	DomainName          *string                             `json:"domainName,omitempty"`
	ExcludeAddressRange *string                             `json:"excludeAddressRange,omitempty"`
	HostName            *string                             `json:"hostName,omitempty"`
	IP                  *string                             `json:"ip,omitempty"`
	IPRange             *string                             `json:"ipRange,omitempty"`
	LeaseTime           *int                                `json:"leaseTime,omitempty"`
	NetMask             *string                             `json:"netMask,omitempty"`
	PoolID              *string                             `json:"poolId,omitempty"`
	PoolName            *string                             `json:"poolName,omitempty"`
	PrimaryDNSServer    *string                             `json:"primaryDnsServer,omitempty"`
	PrimaryRouter       *string                             `json:"primaryRouter,omitempty"`
	ProfileID           *string                             `json:"profileId,omitempty"`
	QinqVlanRanges      []*DpDHCPProfileQinqVlanRangeBO     `json:"qinqVlanRanges,omitempty"`
	SecondaryDNSServer  *string                             `json:"secondaryDnsServer,omitempty"`
	SecondaryRouter     *string                             `json:"secondaryRouter,omitempty"`
	SubOptionSpaces     []*DpDHCPProfileOptionSpaceInstance `json:"subOptionSpaces,omitempty"`
	VlanRange           *string                             `json:"vlanRange,omitempty"`
}

type DpDHCPProfilePoolBOList struct {
	FirstIndex *int                   `json:"firstIndex,omitempty"`
	HasMore    *bool                  `json:"hasMore,omitempty"`
	List       []*DpDHCPProfilePoolBO `json:"list,omitempty"`
	TotalCount *int                   `json:"totalCount,omitempty"`
}

type DpDHCPProfileQinqVlanRangeBO struct {
	Cvlan *string `json:"cvlan,omitempty"`
	Svlan *string `json:"svlan,omitempty"`
}

type DpNatProfileBasicBO struct {
	AppliedDpKey                    *string                       `json:"appliedDpKey,omitempty"`
	Description                     *string                       `json:"description,omitempty"`
	NatPublicSubnetID               *DpNatProfilePublicSubnetIDBO `json:"natPublicSubnetId,omitempty"`
	NatPublicVlanID                 *int                          `json:"natPublicVlanId,omitempty"`
	PrimaryNatDefaultRouteGateway   *string                       `json:"primaryNatDefaultRouteGateway,omitempty"`
	ProfileID                       *string                       `json:"profileId,omitempty"`
	ProfileName                     *string                       `json:"profileName,omitempty"`
	SecondaryNatDefaultRouteGateway *string                       `json:"secondaryNatDefaultRouteGateway,omitempty"`
}

type DpNatProfileBasicBOList struct {
	FirstIndex *int                   `json:"firstIndex,omitempty"`
	HasMore    *bool                  `json:"hasMore,omitempty"`
	List       []*DpNatProfileBasicBO `json:"list,omitempty"`
	TotalCount *int                   `json:"totalCount,omitempty"`
}

type DpNatProfilePoolBO struct {
	Description          *string                               `json:"description,omitempty"`
	NatPortRange         *string                               `json:"natPortRange,omitempty"`
	PoolID               *string                               `json:"poolId,omitempty"`
	PoolName             *string                               `json:"poolName,omitempty"`
	PrivateQinqVlanRange []*DpNatProfilePrivateQinqVlanRangeBO `json:"privateQinqVlanRange,omitempty"`
	PrivateVlanRange     []string                              `json:"privateVlanRange,omitempty"`
	ProfileID            *string                               `json:"profileId,omitempty"`
	PublicAddressRange   []string                              `json:"publicAddressRange,omitempty"`
	PublicPrefix         *int                                  `json:"publicPrefix,omitempty"`
	PublicVlan           *int                                  `json:"publicVlan,omitempty"`
}

type DpNatProfilePoolBOList struct {
	FirstIndex *int                  `json:"firstIndex,omitempty"`
	HasMore    *bool                 `json:"hasMore,omitempty"`
	List       []*DpNatProfilePoolBO `json:"list,omitempty"`
	TotalCount *int                  `json:"totalCount,omitempty"`
}

type DpNatProfilePrivateQinqVlanRangeBO struct {
	Cvlan *string `json:"cvlan,omitempty"`
	Svlan *string `json:"svlan,omitempty"`
}

type DpNatProfilePublicSubnetIDBO struct {
	IP           *string `json:"ip,omitempty"`
	PrefixLength *int    `json:"prefixLength,omitempty"`
}

type DpProfileSettingBO struct {
	Description     *string `json:"description,omitempty"`
	DHCPProfileID   *string `json:"dhcpProfileId,omitempty"`
	DHCPProfileName *string `json:"dhcpProfileName,omitempty"`
	DpKey           *string `json:"dpKey,omitempty"`
	DpName          *string `json:"dpName,omitempty"`
	DpVersion       *string `json:"dpVersion,omitempty"`
	NatProfileID    *string `json:"natProfileId,omitempty"`
	NatProfileName  *string `json:"natProfileName,omitempty"`
}

type DpProfileSettingBOList struct {
	FirstIndex *int                  `json:"firstIndex,omitempty"`
	HasMore    *bool                 `json:"hasMore,omitempty"`
	List       []*DpProfileSettingBO `json:"list,omitempty"`
	TotalCount *int                  `json:"totalCount,omitempty"`
}
