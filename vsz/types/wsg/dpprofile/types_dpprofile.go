package dpprofile

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type BulkDelete struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewBulkDelete() *BulkDelete {
	bulkDeleteType := new(BulkDelete)
	return bulkDeleteType
}

func NewBulkDeleteWithDefaults() *BulkDelete {
	bulkDeleteType := new(BulkDelete)
	return bulkDeleteType
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

func NewDpDhcpProfileBasicBO() *DpDhcpProfileBasicBO {
	dpDhcpProfileBasicBOType := new(DpDhcpProfileBasicBO)
	return dpDhcpProfileBasicBOType
}

func NewDpDhcpProfileBasicBOWithDefaults() *DpDhcpProfileBasicBO {
	dpDhcpProfileBasicBOType := new(DpDhcpProfileBasicBO)
	return dpDhcpProfileBasicBOType
}

type DpDhcpProfileBasicBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfileBasicBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDpDhcpProfileBasicBOList() *DpDhcpProfileBasicBOList {
	dpDhcpProfileBasicBOListType := new(DpDhcpProfileBasicBOList)
	return dpDhcpProfileBasicBOListType
}

func NewDpDhcpProfileBasicBOListWithDefaults() *DpDhcpProfileBasicBOList {
	dpDhcpProfileBasicBOListType := new(DpDhcpProfileBasicBOList)
	return dpDhcpProfileBasicBOListType
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

func NewDpDhcpProfileHostBO() *DpDhcpProfileHostBO {
	dpDhcpProfileHostBOType := new(DpDhcpProfileHostBO)
	return dpDhcpProfileHostBOType
}

func NewDpDhcpProfileHostBOWithDefaults() *DpDhcpProfileHostBO {
	dpDhcpProfileHostBOType := new(DpDhcpProfileHostBO)
	return dpDhcpProfileHostBOType
}

type DpDhcpProfileHostBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfileHostBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDpDhcpProfileHostBOList() *DpDhcpProfileHostBOList {
	dpDhcpProfileHostBOListType := new(DpDhcpProfileHostBOList)
	return dpDhcpProfileHostBOListType
}

func NewDpDhcpProfileHostBOListWithDefaults() *DpDhcpProfileHostBOList {
	dpDhcpProfileHostBOListType := new(DpDhcpProfileHostBOList)
	return dpDhcpProfileHostBOListType
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

func NewDpDhcpProfileOptionBO() *DpDhcpProfileOptionBO {
	dpDhcpProfileOptionBOType := new(DpDhcpProfileOptionBO)
	return dpDhcpProfileOptionBOType
}

func NewDpDhcpProfileOptionBOWithDefaults() *DpDhcpProfileOptionBO {
	dpDhcpProfileOptionBOType := new(DpDhcpProfileOptionBO)
	return dpDhcpProfileOptionBOType
}

type DpDhcpProfileOptionInstance struct {
	FunctionName *string `json:"functionName,omitempty"`

	Value *string `json:"value,omitempty"`
}

func NewDpDhcpProfileOptionInstance() *DpDhcpProfileOptionInstance {
	dpDhcpProfileOptionInstanceType := new(DpDhcpProfileOptionInstance)
	return dpDhcpProfileOptionInstanceType
}

func NewDpDhcpProfileOptionInstanceWithDefaults() *DpDhcpProfileOptionInstance {
	dpDhcpProfileOptionInstanceType := new(DpDhcpProfileOptionInstance)
	return dpDhcpProfileOptionInstanceType
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

func NewDpDhcpProfileOptionSpaceApplyToBO() *DpDhcpProfileOptionSpaceApplyToBO {
	dpDhcpProfileOptionSpaceApplyToBOType := new(DpDhcpProfileOptionSpaceApplyToBO)
	return dpDhcpProfileOptionSpaceApplyToBOType
}

func NewDpDhcpProfileOptionSpaceApplyToBOWithDefaults() *DpDhcpProfileOptionSpaceApplyToBO {
	dpDhcpProfileOptionSpaceApplyToBOType := new(DpDhcpProfileOptionSpaceApplyToBO)
	return dpDhcpProfileOptionSpaceApplyToBOType
}

type DpDhcpProfileOptionSpaceApplyToBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfileOptionSpaceApplyToBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDpDhcpProfileOptionSpaceApplyToBOList() *DpDhcpProfileOptionSpaceApplyToBOList {
	dpDhcpProfileOptionSpaceApplyToBOListType := new(DpDhcpProfileOptionSpaceApplyToBOList)
	return dpDhcpProfileOptionSpaceApplyToBOListType
}

func NewDpDhcpProfileOptionSpaceApplyToBOListWithDefaults() *DpDhcpProfileOptionSpaceApplyToBOList {
	dpDhcpProfileOptionSpaceApplyToBOListType := new(DpDhcpProfileOptionSpaceApplyToBOList)
	return dpDhcpProfileOptionSpaceApplyToBOListType
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

func NewDpDhcpProfileOptionSpaceBO() *DpDhcpProfileOptionSpaceBO {
	dpDhcpProfileOptionSpaceBOType := new(DpDhcpProfileOptionSpaceBO)
	return dpDhcpProfileOptionSpaceBOType
}

func NewDpDhcpProfileOptionSpaceBOWithDefaults() *DpDhcpProfileOptionSpaceBO {
	dpDhcpProfileOptionSpaceBOType := new(DpDhcpProfileOptionSpaceBO)
	return dpDhcpProfileOptionSpaceBOType
}

type DpDhcpProfileOptionSpaceInstance struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	Options []*DpDhcpProfileOptionInstance `json:"options,omitempty"`

	SpaceId *string `json:"spaceId,omitempty"`
}

func NewDpDhcpProfileOptionSpaceInstance() *DpDhcpProfileOptionSpaceInstance {
	dpDhcpProfileOptionSpaceInstanceType := new(DpDhcpProfileOptionSpaceInstance)
	return dpDhcpProfileOptionSpaceInstanceType
}

func NewDpDhcpProfileOptionSpaceInstanceWithDefaults() *DpDhcpProfileOptionSpaceInstance {
	dpDhcpProfileOptionSpaceInstanceType := new(DpDhcpProfileOptionSpaceInstance)
	return dpDhcpProfileOptionSpaceInstanceType
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

func NewDpDhcpProfilePoolBO() *DpDhcpProfilePoolBO {
	dpDhcpProfilePoolBOType := new(DpDhcpProfilePoolBO)
	return dpDhcpProfilePoolBOType
}

func NewDpDhcpProfilePoolBOWithDefaults() *DpDhcpProfilePoolBO {
	dpDhcpProfilePoolBOType := new(DpDhcpProfilePoolBO)
	return dpDhcpProfilePoolBOType
}

type DpDhcpProfilePoolBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpDhcpProfilePoolBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDpDhcpProfilePoolBOList() *DpDhcpProfilePoolBOList {
	dpDhcpProfilePoolBOListType := new(DpDhcpProfilePoolBOList)
	return dpDhcpProfilePoolBOListType
}

func NewDpDhcpProfilePoolBOListWithDefaults() *DpDhcpProfilePoolBOList {
	dpDhcpProfilePoolBOListType := new(DpDhcpProfilePoolBOList)
	return dpDhcpProfilePoolBOListType
}

type DpDhcpProfileQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	Svlan *string `json:"svlan,omitempty"`
}

func NewDpDhcpProfileQinqVlanRangeBO() *DpDhcpProfileQinqVlanRangeBO {
	dpDhcpProfileQinqVlanRangeBOType := new(DpDhcpProfileQinqVlanRangeBO)
	return dpDhcpProfileQinqVlanRangeBOType
}

func NewDpDhcpProfileQinqVlanRangeBOWithDefaults() *DpDhcpProfileQinqVlanRangeBO {
	dpDhcpProfileQinqVlanRangeBOType := new(DpDhcpProfileQinqVlanRangeBO)
	return dpDhcpProfileQinqVlanRangeBOType
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

func NewDpNatProfileBasicBO() *DpNatProfileBasicBO {
	dpNatProfileBasicBOType := new(DpNatProfileBasicBO)
	return dpNatProfileBasicBOType
}

func NewDpNatProfileBasicBOWithDefaults() *DpNatProfileBasicBO {
	dpNatProfileBasicBOType := new(DpNatProfileBasicBO)
	return dpNatProfileBasicBOType
}

type DpNatProfileBasicBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpNatProfileBasicBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDpNatProfileBasicBOList() *DpNatProfileBasicBOList {
	dpNatProfileBasicBOListType := new(DpNatProfileBasicBOList)
	return dpNatProfileBasicBOListType
}

func NewDpNatProfileBasicBOListWithDefaults() *DpNatProfileBasicBOList {
	dpNatProfileBasicBOListType := new(DpNatProfileBasicBOList)
	return dpNatProfileBasicBOListType
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

func NewDpNatProfilePoolBO() *DpNatProfilePoolBO {
	dpNatProfilePoolBOType := new(DpNatProfilePoolBO)
	return dpNatProfilePoolBOType
}

func NewDpNatProfilePoolBOWithDefaults() *DpNatProfilePoolBO {
	dpNatProfilePoolBOType := new(DpNatProfilePoolBO)
	return dpNatProfilePoolBOType
}

type DpNatProfilePoolBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpNatProfilePoolBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDpNatProfilePoolBOList() *DpNatProfilePoolBOList {
	dpNatProfilePoolBOListType := new(DpNatProfilePoolBOList)
	return dpNatProfilePoolBOListType
}

func NewDpNatProfilePoolBOListWithDefaults() *DpNatProfilePoolBOList {
	dpNatProfilePoolBOListType := new(DpNatProfilePoolBOList)
	return dpNatProfilePoolBOListType
}

type DpNatProfilePrivateQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	Svlan *string `json:"svlan,omitempty"`
}

func NewDpNatProfilePrivateQinqVlanRangeBO() *DpNatProfilePrivateQinqVlanRangeBO {
	dpNatProfilePrivateQinqVlanRangeBOType := new(DpNatProfilePrivateQinqVlanRangeBO)
	return dpNatProfilePrivateQinqVlanRangeBOType
}

func NewDpNatProfilePrivateQinqVlanRangeBOWithDefaults() *DpNatProfilePrivateQinqVlanRangeBO {
	dpNatProfilePrivateQinqVlanRangeBOType := new(DpNatProfilePrivateQinqVlanRangeBO)
	return dpNatProfilePrivateQinqVlanRangeBOType
}

type DpNatProfilePublicSubnetIdBO struct {
	// Ip
	// ip
	Ip *string `json:"ip,omitempty"`

	// PrefixLength
	// prefixLength
	PrefixLength *int `json:"prefixLength,omitempty"`
}

func NewDpNatProfilePublicSubnetIdBO() *DpNatProfilePublicSubnetIdBO {
	dpNatProfilePublicSubnetIdBOType := new(DpNatProfilePublicSubnetIdBO)
	return dpNatProfilePublicSubnetIdBOType
}

func NewDpNatProfilePublicSubnetIdBOWithDefaults() *DpNatProfilePublicSubnetIdBO {
	dpNatProfilePublicSubnetIdBOType := new(DpNatProfilePublicSubnetIdBO)
	return dpNatProfilePublicSubnetIdBOType
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

func NewDpProfileSettingBO() *DpProfileSettingBO {
	dpProfileSettingBOType := new(DpProfileSettingBO)
	return dpProfileSettingBOType
}

func NewDpProfileSettingBOWithDefaults() *DpProfileSettingBO {
	dpProfileSettingBOType := new(DpProfileSettingBO)
	return dpProfileSettingBOType
}

type DpProfileSettingBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DpProfileSettingBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDpProfileSettingBOList() *DpProfileSettingBOList {
	dpProfileSettingBOListType := new(DpProfileSettingBOList)
	return dpProfileSettingBOListType
}

func NewDpProfileSettingBOListWithDefaults() *DpProfileSettingBOList {
	dpProfileSettingBOListType := new(DpProfileSettingBOList)
	return dpProfileSettingBOListType
}
