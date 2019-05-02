package dpprofile

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type BulkDelete struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type DpDhcpProfileBasicBO struct {
	// DefaultLeaseTime
	// defaultLeaseTime
	DefaultLeaseTime *int `json:"defaultLeaseTime,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DomainName
	// domainName
	DomainName *string `json:"domainName,omitempty"`

	// PrimaryDnsServer
	// primaryDnsServer
	PrimaryDnsServer *string `json:"primaryDnsServer,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// ProfileName
	// profileName
	ProfileName *string `json:"profileName,omitempty"`

	// SecondaryDnsServer
	// secondaryDnsServer
	SecondaryDnsServer *string `json:"secondaryDnsServer,omitempty"`
}

type DpDhcpProfileBasicBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfileBasicBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DpDhcpProfileHostBO struct {
	// BroadcastAddress
	// broadcastAddress
	BroadcastAddress *string `json:"broadcastAddress,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DnsServers
	// dnsServers
	DnsServers []string `json:"dnsServers,omitempty"`

	// DomainName
	// domainName
	DomainName *string `json:"domainName,omitempty"`

	// FixedAddress
	// fixedAddress
	FixedAddress *string `json:"fixedAddress,omitempty"`

	// HardwareEthernet
	// hardwareEthernet
	HardwareEthernet *string `json:"hardwareEthernet,omitempty"`

	// HostId
	// hostId
	HostId *string `json:"hostId,omitempty"`

	// HostName
	// hostName
	HostName *string `json:"hostName,omitempty"`

	// LeaseTime
	// leaseTime
	LeaseTime *int `json:"leaseTime,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// Routers
	// routers
	Routers []string `json:"routers,omitempty"`
}

type DpDhcpProfileHostBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfileHostBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DpDhcpProfileOptionBO struct {
	// CodeNumber
	// codeNumber
	CodeNumber *int `json:"codeNumber,omitempty"`

	// FunctionName
	// functionName
	FunctionName *string `json:"functionName,omitempty"`

	// Type
	// type
	Type *string `json:"type,omitempty"`

	// Value
	// value
	Value *string `json:"value,omitempty"`
}

type DpDhcpProfileOptionInstance struct {
	FunctionName *string `json:"functionName,omitempty"`

	Value *string `json:"value,omitempty"`
}

type DpDhcpProfileOptionSpaceApplyToBO struct {
	// AppliedPoolNames
	// appliedPoolNames
	AppliedPoolNames []string `json:"appliedPoolNames,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// Options
	// options
	Options []*DpDhcpProfileOptionBO `json:"options,omitempty"`

	SpaceId *string `json:"spaceId,omitempty"`
}

type DpDhcpProfileOptionSpaceApplyToBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfileOptionSpaceApplyToBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DpDhcpProfileOptionSpaceBO struct {
	// Description
	// description
	Description *string `json:"description,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// Options
	// options
	Options []*DpDhcpProfileOptionBO `json:"options,omitempty"`

	// SpaceId
	// spaceId
	SpaceId *string `json:"spaceId,omitempty"`
}

type DpDhcpProfileOptionSpaceInstance struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	Options []*DpDhcpProfileOptionInstance `json:"options,omitempty"`

	SpaceId *string `json:"spaceId,omitempty"`
}

type DpDhcpProfilePoolBO struct {
	// BroadcastAddress
	// broadcastAddress
	BroadcastAddress *string `json:"broadcastAddress,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DomainName
	// domainName
	DomainName *string `json:"domainName,omitempty"`

	// ExcludeAddressRange
	// excludeAddressRange
	ExcludeAddressRange *string `json:"excludeAddressRange,omitempty"`

	// HostName
	// hostName
	HostName *string `json:"hostName,omitempty"`

	// Ip
	// ip
	Ip *string `json:"ip,omitempty"`

	// IpRange
	// ipRange
	IpRange *string `json:"ipRange,omitempty"`

	// LeaseTime
	// leaseTime
	LeaseTime *int `json:"leaseTime,omitempty"`

	// NetMask
	// netMask
	NetMask *string `json:"netMask,omitempty"`

	// PoolId
	// poolId
	PoolId *string `json:"poolId,omitempty"`

	// PoolName
	// poolName
	PoolName *string `json:"poolName,omitempty"`

	// PrimaryDnsServer
	// primaryDnsServer
	PrimaryDnsServer *string `json:"primaryDnsServer,omitempty"`

	// PrimaryRouter
	// primaryRouter
	PrimaryRouter *string `json:"primaryRouter,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// QinqVlanRanges
	// qinqVlanRanges
	QinqVlanRanges []*DpDhcpProfileQinqVlanRangeBO `json:"qinqVlanRanges,omitempty"`

	// SecondaryDnsServer
	// secondaryDnsServer
	SecondaryDnsServer *string `json:"secondaryDnsServer,omitempty"`

	// SecondaryRouter
	// secondaryRouter
	SecondaryRouter *string `json:"secondaryRouter,omitempty"`

	// SubOptionSpaces
	// subOptionSpaces
	SubOptionSpaces []*DpDhcpProfileOptionSpaceInstance `json:"subOptionSpaces,omitempty"`

	// VlanRange
	// vlanRange
	VlanRange *string `json:"vlanRange,omitempty"`
}

type DpDhcpProfilePoolBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfilePoolBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DpDhcpProfileQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	Svlan *string `json:"svlan,omitempty"`
}

type DpNatProfileBasicBO struct {
	// AppliedDpKey
	// The applied DP keys
	AppliedDpKey *string `json:"appliedDpKey,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	NatPublicSubnetId *DpNatProfilePublicSubnetIdBO `json:"natPublicSubnetId,omitempty"`

	// NatPublicVlanId
	// natPublicVlanId
	NatPublicVlanId *int `json:"natPublicVlanId,omitempty"`

	// PrimaryNatDefaultRouteGateway
	// primaryNatDefaultRouteGateway
	PrimaryNatDefaultRouteGateway *string `json:"primaryNatDefaultRouteGateway,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// ProfileName
	// profileName
	ProfileName *string `json:"profileName,omitempty"`

	// SecondaryNatDefaultRouteGateway
	// secondaryNatDefaultRouteGateway
	SecondaryNatDefaultRouteGateway *string `json:"secondaryNatDefaultRouteGateway,omitempty"`
}

type DpNatProfileBasicBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpNatProfileBasicBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DpNatProfilePoolBO struct {
	// Description
	// description
	Description *string `json:"description,omitempty"`

	// NatPortRange
	// natPortRange
	NatPortRange *string `json:"natPortRange,omitempty"`

	// PoolId
	// poolId
	PoolId *string `json:"poolId,omitempty"`

	// PoolName
	// poolName
	PoolName *string `json:"poolName,omitempty"`

	// PrivateQinqVlanRange
	// privateQinqVlanRange
	PrivateQinqVlanRange []*DpNatProfilePrivateQinqVlanRangeBO `json:"privateQinqVlanRange,omitempty"`

	// PrivateVlanRange
	// privateVlanRange
	PrivateVlanRange []string `json:"privateVlanRange,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// PublicAddressRange
	// publicAddressRange
	PublicAddressRange []string `json:"publicAddressRange,omitempty"`

	// PublicPrefix
	// publicPrefix
	PublicPrefix *int `json:"publicPrefix,omitempty"`

	// PublicVlan
	// publicVlan
	PublicVlan *int `json:"publicVlan,omitempty"`
}

type DpNatProfilePoolBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpNatProfilePoolBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DpNatProfilePrivateQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	Svlan *string `json:"svlan,omitempty"`
}

type DpNatProfilePublicSubnetIdBO struct {
	// Ip
	// ip
	Ip *string `json:"ip,omitempty"`

	// PrefixLength
	// prefixLength
	PrefixLength *int `json:"prefixLength,omitempty"`
}

type DpProfileSettingBO struct {
	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DhcpProfileId
	// dhcpProfileId
	DhcpProfileId *string `json:"dhcpProfileId,omitempty"`

	// DhcpProfileName
	// dhcpProfileName
	DhcpProfileName *string `json:"dhcpProfileName,omitempty"`

	// DpKey
	// dpKey
	DpKey *string `json:"dpKey,omitempty"`

	// DpName
	// dpName
	DpName *string `json:"dpName,omitempty"`

	// DpVersion
	// dpVersion
	DpVersion *string `json:"dpVersion,omitempty"`

	// NatProfileId
	// natProfileId
	NatProfileId *string `json:"natProfileId,omitempty"`

	// NatProfileName
	// natProfileName
	NatProfileName *string `json:"natProfileName,omitempty"`
}

type DpProfileSettingBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpProfileSettingBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
