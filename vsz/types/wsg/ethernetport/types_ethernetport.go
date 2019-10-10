package ethernetport

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	Description *common.Description `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	GuestVlan *int `json:"guestVlan,omitempty" validate:"gte=1,lte=4094"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	Type *string `json:"type,omitempty" validate:"required,oneof=AccessPort TrunkPort GeneralPort"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	UntagId *int `json:"untagId,omitempty" validate:"gte=1,lte=4094"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *apmodel.LanPort8021X `json:"_8021X,omitempty" validate:"required"`
}

type EthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	Description *common.Description `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	GuestVlan *int `json:"guestVlan,omitempty" validate:"gte=1,lte=4094"`

	// Id
	// identifier of the ethernet port profile
	Id *string `json:"id,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	Type *string `json:"type,omitempty" validate:"oneof=AccessPort TrunkPort GeneralPort"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	UntagId *int `json:"untagId,omitempty" validate:"gte=1,lte=4094"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *apmodel.LanPort8021X `json:"_8021X,omitempty"`
}

type ModifyEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	Description *common.Description `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	GuestVlan *int `json:"guestVlan,omitempty" validate:"gte=1,lte=4094"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	UntagId *int `json:"untagId,omitempty" validate:"gte=1,lte=4094"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *apmodel.LanPort8021X `json:"_8021X,omitempty"`
}

type ProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ProfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}
