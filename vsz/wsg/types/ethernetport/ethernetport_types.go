package ethernetport

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/types/apmodel"
	"github.com/myENA/ruckus-client/vsz/wsg/types/common"
)

type CreateEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DHCPRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DHCPRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	GuestVlan *int `json:"guestVlan,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	Type *string `json:"type,omitempty"`

	// UntagID
	// VLAN Untag ID of the ethernet port profile
	UntagID *int `json:"untagId,omitempty"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *apmodel.LanPort8021X `json:"_8021X,omitempty"`
}

type EthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DHCPRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DHCPRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	GuestVlan *int `json:"guestVlan,omitempty"`

	// ID
	// identifier of the ethernet port profile
	ID *string `json:"id,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	Type *string `json:"type,omitempty"`

	// UntagID
	// VLAN Untag ID of the ethernet port profile
	UntagID *int `json:"untagId,omitempty"`

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
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DHCPRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DHCPRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	GuestVlan *int `json:"guestVlan,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	// UntagID
	// VLAN Untag ID of the ethernet port profile
	UntagID *int `json:"untagId,omitempty"`

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
	// ID
	// Identifier of the service
	ID *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}
