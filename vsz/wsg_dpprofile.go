package vsz

// API Version: v9_0

type WSGDPProfileBulkDelete struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGDPProfileBulkDelete() *WSGDPProfileBulkDelete {
	m := new(WSGDPProfileBulkDelete)
	return m
}

type WSGDPProfileDpDhcpProfileBasicBO struct {
	// DefaultLeaseTime
	// defaultLeaseTime
	// Constraints:
	//    - nullable
	DefaultLeaseTime *int `json:"defaultLeaseTime,omitempty"`

	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DomainName
	// domainName
	// Constraints:
	//    - nullable
	DomainName *string `json:"domainName,omitempty"`

	// PrimaryDnsServer
	// primaryDnsServer
	// Constraints:
	//    - nullable
	PrimaryDnsServer *string `json:"primaryDnsServer,omitempty"`

	// ProfileId
	// profileId
	// Constraints:
	//    - nullable
	ProfileId *string `json:"profileId,omitempty"`

	// ProfileName
	// profileName
	// Constraints:
	//    - nullable
	ProfileName *string `json:"profileName,omitempty"`

	// SecondaryDnsServer
	// secondaryDnsServer
	// Constraints:
	//    - nullable
	SecondaryDnsServer *string `json:"secondaryDnsServer,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileBasicBO() *WSGDPProfileDpDhcpProfileBasicBO {
	m := new(WSGDPProfileDpDhcpProfileBasicBO)
	return m
}

type WSGDPProfileDpDhcpProfileBasicBOList struct {
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
	List []*WSGDPProfileDpDhcpProfileBasicBO `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileBasicBOList() *WSGDPProfileDpDhcpProfileBasicBOList {
	m := new(WSGDPProfileDpDhcpProfileBasicBOList)
	return m
}

type WSGDPProfileDpDhcpProfileHostBO struct {
	// BroadcastAddress
	// broadcastAddress
	// Constraints:
	//    - nullable
	BroadcastAddress *string `json:"broadcastAddress,omitempty"`

	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DnsServers
	// dnsServers
	// Constraints:
	//    - nullable
	DnsServers []string `json:"dnsServers,omitempty" validate:"omitempty,dive"`

	// DomainName
	// domainName
	// Constraints:
	//    - nullable
	DomainName *string `json:"domainName,omitempty"`

	// FixedAddress
	// fixedAddress
	// Constraints:
	//    - nullable
	FixedAddress *string `json:"fixedAddress,omitempty"`

	// HardwareEthernet
	// hardwareEthernet
	// Constraints:
	//    - nullable
	HardwareEthernet *string `json:"hardwareEthernet,omitempty"`

	// HostId
	// hostId
	// Constraints:
	//    - nullable
	HostId *string `json:"hostId,omitempty"`

	// HostName
	// hostName
	// Constraints:
	//    - nullable
	HostName *string `json:"hostName,omitempty"`

	// LeaseTime
	// leaseTime
	// Constraints:
	//    - nullable
	LeaseTime *int `json:"leaseTime,omitempty"`

	// Name
	// name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// ProfileId
	// profileId
	// Constraints:
	//    - nullable
	ProfileId *string `json:"profileId,omitempty"`

	// Routers
	// routers
	// Constraints:
	//    - nullable
	Routers []string `json:"routers,omitempty" validate:"omitempty,dive"`
}

func NewWSGDPProfileDpDhcpProfileHostBO() *WSGDPProfileDpDhcpProfileHostBO {
	m := new(WSGDPProfileDpDhcpProfileHostBO)
	return m
}

type WSGDPProfileDpDhcpProfileHostBOList struct {
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
	List []*WSGDPProfileDpDhcpProfileHostBO `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileHostBOList() *WSGDPProfileDpDhcpProfileHostBOList {
	m := new(WSGDPProfileDpDhcpProfileHostBOList)
	return m
}

type WSGDPProfileDpDhcpProfileOptionBO struct {
	// CodeNumber
	// codeNumber
	// Constraints:
	//    - nullable
	CodeNumber *int `json:"codeNumber,omitempty"`

	// FunctionName
	// functionName
	// Constraints:
	//    - nullable
	FunctionName *string `json:"functionName,omitempty"`

	// Type
	// type
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`

	// Value
	// value
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionBO() *WSGDPProfileDpDhcpProfileOptionBO {
	m := new(WSGDPProfileDpDhcpProfileOptionBO)
	return m
}

type WSGDPProfileDpDhcpProfileOptionInstance struct {
	// FunctionName
	// Constraints:
	//    - nullable
	FunctionName *string `json:"functionName,omitempty"`

	// Value
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionInstance() *WSGDPProfileDpDhcpProfileOptionInstance {
	m := new(WSGDPProfileDpDhcpProfileOptionInstance)
	return m
}

type WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO struct {
	// AppliedPoolNames
	// appliedPoolNames
	// Constraints:
	//    - nullable
	AppliedPoolNames []string `json:"appliedPoolNames,omitempty" validate:"omitempty,dive"`

	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// Name
	// name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Options
	// options
	// Constraints:
	//    - nullable
	Options []*WSGDPProfileDpDhcpProfileOptionBO `json:"options,omitempty" validate:"omitempty,dive"`

	// SpaceId
	// Constraints:
	//    - nullable
	SpaceId *string `json:"spaceId,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBO() *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO)
	return m
}

type WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList struct {
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
	List []*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList() *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList)
	return m
}

type WSGDPProfileDpDhcpProfileOptionSpaceBO struct {
	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// Name
	// name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Options
	// options
	// Constraints:
	//    - nullable
	Options []*WSGDPProfileDpDhcpProfileOptionBO `json:"options,omitempty" validate:"omitempty,dive"`

	// SpaceId
	// spaceId
	// Constraints:
	//    - nullable
	SpaceId *string `json:"spaceId,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionSpaceBO() *WSGDPProfileDpDhcpProfileOptionSpaceBO {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceBO)
	return m
}

type WSGDPProfileDpDhcpProfileOptionSpaceInstance struct {
	// Description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Options
	// Constraints:
	//    - nullable
	Options []*WSGDPProfileDpDhcpProfileOptionInstance `json:"options,omitempty" validate:"omitempty,dive"`

	// SpaceId
	// Constraints:
	//    - nullable
	SpaceId *string `json:"spaceId,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionSpaceInstance() *WSGDPProfileDpDhcpProfileOptionSpaceInstance {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceInstance)
	return m
}

type WSGDPProfileDpDhcpProfilePoolBO struct {
	// BroadcastAddress
	// broadcastAddress
	// Constraints:
	//    - nullable
	BroadcastAddress *string `json:"broadcastAddress,omitempty"`

	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DomainName
	// domainName
	// Constraints:
	//    - nullable
	DomainName *string `json:"domainName,omitempty"`

	// ExcludeAddressRange
	// excludeAddressRange
	// Constraints:
	//    - nullable
	ExcludeAddressRange *string `json:"excludeAddressRange,omitempty"`

	// HostName
	// hostName
	// Constraints:
	//    - nullable
	HostName *string `json:"hostName,omitempty"`

	// Ip
	// ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// IpRange
	// ipRange
	// Constraints:
	//    - nullable
	IpRange *string `json:"ipRange,omitempty"`

	// LeaseTime
	// leaseTime
	// Constraints:
	//    - nullable
	LeaseTime *int `json:"leaseTime,omitempty"`

	// NetMask
	// netMask
	// Constraints:
	//    - nullable
	NetMask *string `json:"netMask,omitempty"`

	// PoolId
	// poolId
	// Constraints:
	//    - nullable
	PoolId *string `json:"poolId,omitempty"`

	// PoolName
	// poolName
	// Constraints:
	//    - nullable
	PoolName *string `json:"poolName,omitempty"`

	// PrimaryDnsServer
	// primaryDnsServer
	// Constraints:
	//    - nullable
	PrimaryDnsServer *string `json:"primaryDnsServer,omitempty"`

	// PrimaryRouter
	// primaryRouter
	// Constraints:
	//    - nullable
	PrimaryRouter *string `json:"primaryRouter,omitempty"`

	// ProfileId
	// profileId
	// Constraints:
	//    - nullable
	ProfileId *string `json:"profileId,omitempty"`

	// QinqVlanRanges
	// qinqVlanRanges
	// Constraints:
	//    - nullable
	QinqVlanRanges []*WSGDPProfileDpDhcpProfileQinqVlanRangeBO `json:"qinqVlanRanges,omitempty" validate:"omitempty,dive"`

	// SecondaryDnsServer
	// secondaryDnsServer
	// Constraints:
	//    - nullable
	SecondaryDnsServer *string `json:"secondaryDnsServer,omitempty"`

	// SecondaryRouter
	// secondaryRouter
	// Constraints:
	//    - nullable
	SecondaryRouter *string `json:"secondaryRouter,omitempty"`

	// SubOptionSpaces
	// subOptionSpaces
	// Constraints:
	//    - nullable
	SubOptionSpaces []*WSGDPProfileDpDhcpProfileOptionSpaceInstance `json:"subOptionSpaces,omitempty" validate:"omitempty,dive"`

	// VlanRange
	// vlanRange
	// Constraints:
	//    - nullable
	VlanRange *string `json:"vlanRange,omitempty"`
}

func NewWSGDPProfileDpDhcpProfilePoolBO() *WSGDPProfileDpDhcpProfilePoolBO {
	m := new(WSGDPProfileDpDhcpProfilePoolBO)
	return m
}

type WSGDPProfileDpDhcpProfilePoolBOList struct {
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
	List []*WSGDPProfileDpDhcpProfilePoolBO `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPProfileDpDhcpProfilePoolBOList() *WSGDPProfileDpDhcpProfilePoolBOList {
	m := new(WSGDPProfileDpDhcpProfilePoolBOList)
	return m
}

type WSGDPProfileDpDhcpProfileQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	// Constraints:
	//    - nullable
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	// Constraints:
	//    - nullable
	Svlan *string `json:"svlan,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileQinqVlanRangeBO() *WSGDPProfileDpDhcpProfileQinqVlanRangeBO {
	m := new(WSGDPProfileDpDhcpProfileQinqVlanRangeBO)
	return m
}

type WSGDPProfileDpNatProfileBasicBO struct {
	// AppliedDpKey
	// The applied DP keys
	// Constraints:
	//    - nullable
	AppliedDpKey *string `json:"appliedDpKey,omitempty"`

	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// NatPublicSubnetId
	// Constraints:
	//    - nullable
	NatPublicSubnetId *WSGDPProfileDpNatProfilePublicSubnetIdBO `json:"natPublicSubnetId,omitempty"`

	// NatPublicVlanId
	// natPublicVlanId
	// Constraints:
	//    - nullable
	NatPublicVlanId *int `json:"natPublicVlanId,omitempty"`

	// PrimaryNatDefaultRouteGateway
	// primaryNatDefaultRouteGateway
	// Constraints:
	//    - nullable
	PrimaryNatDefaultRouteGateway *string `json:"primaryNatDefaultRouteGateway,omitempty"`

	// ProfileId
	// profileId
	// Constraints:
	//    - nullable
	ProfileId *string `json:"profileId,omitempty"`

	// ProfileName
	// profileName
	// Constraints:
	//    - nullable
	ProfileName *string `json:"profileName,omitempty"`

	// SecondaryNatDefaultRouteGateway
	// secondaryNatDefaultRouteGateway
	// Constraints:
	//    - nullable
	SecondaryNatDefaultRouteGateway *string `json:"secondaryNatDefaultRouteGateway,omitempty"`
}

func NewWSGDPProfileDpNatProfileBasicBO() *WSGDPProfileDpNatProfileBasicBO {
	m := new(WSGDPProfileDpNatProfileBasicBO)
	return m
}

type WSGDPProfileDpNatProfileBasicBOList struct {
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
	List []*WSGDPProfileDpNatProfileBasicBO `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPProfileDpNatProfileBasicBOList() *WSGDPProfileDpNatProfileBasicBOList {
	m := new(WSGDPProfileDpNatProfileBasicBOList)
	return m
}

type WSGDPProfileDpNatProfilePoolBO struct {
	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// NatPortRange
	// natPortRange
	// Constraints:
	//    - nullable
	NatPortRange *string `json:"natPortRange,omitempty"`

	// PoolId
	// poolId
	// Constraints:
	//    - nullable
	PoolId *string `json:"poolId,omitempty"`

	// PoolName
	// poolName
	// Constraints:
	//    - nullable
	PoolName *string `json:"poolName,omitempty"`

	// PrivateQinqVlanRange
	// privateQinqVlanRange
	// Constraints:
	//    - nullable
	PrivateQinqVlanRange []*WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO `json:"privateQinqVlanRange,omitempty" validate:"omitempty,dive"`

	// PrivateVlanRange
	// privateVlanRange
	// Constraints:
	//    - nullable
	PrivateVlanRange []string `json:"privateVlanRange,omitempty" validate:"omitempty,dive"`

	// ProfileId
	// profileId
	// Constraints:
	//    - nullable
	ProfileId *string `json:"profileId,omitempty"`

	// PublicAddressRange
	// publicAddressRange
	// Constraints:
	//    - nullable
	PublicAddressRange []string `json:"publicAddressRange,omitempty" validate:"omitempty,dive"`

	// PublicPrefix
	// publicPrefix
	// Constraints:
	//    - nullable
	PublicPrefix *int `json:"publicPrefix,omitempty"`

	// PublicVlan
	// publicVlan
	// Constraints:
	//    - nullable
	PublicVlan *int `json:"publicVlan,omitempty"`
}

func NewWSGDPProfileDpNatProfilePoolBO() *WSGDPProfileDpNatProfilePoolBO {
	m := new(WSGDPProfileDpNatProfilePoolBO)
	return m
}

type WSGDPProfileDpNatProfilePoolBOList struct {
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
	List []*WSGDPProfileDpNatProfilePoolBO `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPProfileDpNatProfilePoolBOList() *WSGDPProfileDpNatProfilePoolBOList {
	m := new(WSGDPProfileDpNatProfilePoolBOList)
	return m
}

type WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	// Constraints:
	//    - nullable
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	// Constraints:
	//    - nullable
	Svlan *string `json:"svlan,omitempty"`
}

func NewWSGDPProfileDpNatProfilePrivateQinqVlanRangeBO() *WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO {
	m := new(WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO)
	return m
}

type WSGDPProfileDpNatProfilePublicSubnetIdBO struct {
	// Ip
	// ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// PrefixLength
	// prefixLength
	// Constraints:
	//    - nullable
	PrefixLength *int `json:"prefixLength,omitempty"`
}

func NewWSGDPProfileDpNatProfilePublicSubnetIdBO() *WSGDPProfileDpNatProfilePublicSubnetIdBO {
	m := new(WSGDPProfileDpNatProfilePublicSubnetIdBO)
	return m
}

type WSGDPProfileSettingBO struct {
	// Description
	// description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DhcpProfileId
	// dhcpProfileId
	// Constraints:
	//    - nullable
	DhcpProfileId *string `json:"dhcpProfileId,omitempty"`

	// DhcpProfileName
	// dhcpProfileName
	// Constraints:
	//    - nullable
	DhcpProfileName *string `json:"dhcpProfileName,omitempty"`

	// DpKey
	// dpKey
	// Constraints:
	//    - nullable
	DpKey *string `json:"dpKey,omitempty"`

	// DpName
	// dpName
	// Constraints:
	//    - nullable
	DpName *string `json:"dpName,omitempty"`

	// DpVersion
	// dpVersion
	// Constraints:
	//    - nullable
	DpVersion *string `json:"dpVersion,omitempty"`

	// NatProfileId
	// natProfileId
	// Constraints:
	//    - nullable
	NatProfileId *string `json:"natProfileId,omitempty"`

	// NatProfileName
	// natProfileName
	// Constraints:
	//    - nullable
	NatProfileName *string `json:"natProfileName,omitempty"`
}

func NewWSGDPProfileSettingBO() *WSGDPProfileSettingBO {
	m := new(WSGDPProfileSettingBO)
	return m
}

type WSGDPProfileSettingBOList struct {
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
	List []*WSGDPProfileSettingBO `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDPProfileSettingBOList() *WSGDPProfileSettingBOList {
	m := new(WSGDPProfileSettingBOList)
	return m
}
