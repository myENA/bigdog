package vsz

// API Version: v9_0

type WSGEthernetPortCreateEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	// Constraints:
	//    - nullable
	//    - default:false
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// DynamicVlanEnabled
	// Constraints:
	//    - nullable
	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	// Constraints:
	//    - nullable
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// GuestVlan
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// IpsecProfile
	// Constraints:
	//    - nullable
	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Qinq
	// Constraints:
	//    - nullable
	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	// Constraints:
	//    - nullable
	//    - default:false
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	// TunnelProfile
	// Constraints:
	//    - nullable
	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	// Constraints:
	//    - required
	//    - default:'TrunkPort'
	//    - oneof:[AccessPort,TrunkPort,GeneralPort]
	Type *string `json:"type" validate:"required,oneof=AccessPort TrunkPort GeneralPort"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - default:1
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// UserSidePortEnabled
	// User side port enabled.
	// Constraints:
	//    - nullable
	//    - default:false
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - nullable
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty" validate:"omitempty,gte=1,lte=32"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	VlanMembers *string `json:"vlanMembers,omitempty"`

	// X8021X
	// Constraints:
	//    - required
	X8021X *WSGAPModelLanPort8021X `json:"_8021X" validate:"required"`
}

func NewWSGEthernetPortCreateEthernetPortProfile() *WSGEthernetPortCreateEthernetPortProfile {
	m := new(WSGEthernetPortCreateEthernetPortProfile)
	return m
}

type WSGEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	// Constraints:
	//    - nullable
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// DynamicVlanEnabled
	// Constraints:
	//    - nullable
	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	// Constraints:
	//    - nullable
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// GuestVlan
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// Id
	// identifier of the ethernet port profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpsecProfile
	// Constraints:
	//    - nullable
	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Qinq
	// Constraints:
	//    - nullable
	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	// Constraints:
	//    - nullable
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	// TunnelProfile
	// Constraints:
	//    - nullable
	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - oneof:[AccessPort,TrunkPort,GeneralPort]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AccessPort TrunkPort GeneralPort"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// UserSidePortEnabled
	// User side port enabled.
	// Constraints:
	//    - nullable
	//    - default:false
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - nullable
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty" validate:"omitempty,gte=1,lte=32"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	VlanMembers *string `json:"vlanMembers,omitempty"`

	// X8021X
	// Constraints:
	//    - nullable
	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

func NewWSGEthernetPortProfile() *WSGEthernetPortProfile {
	m := new(WSGEthernetPortProfile)
	return m
}

type WSGEthernetPortModifyEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	// Constraints:
	//    - nullable
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - default:15
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// DynamicVlanEnabled
	// Constraints:
	//    - nullable
	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	// Constraints:
	//    - nullable
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// GuestVlan
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// IpsecProfile
	// Constraints:
	//    - nullable
	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Qinq
	// Constraints:
	//    - nullable
	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	// Constraints:
	//    - nullable
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	// TunnelProfile
	// Constraints:
	//    - nullable
	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// UserSidePortEnabled
	// User side port enabled.
	// Constraints:
	//    - nullable
	//    - default:false
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - nullable
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty" validate:"omitempty,gte=1,lte=32"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	VlanMembers *string `json:"vlanMembers,omitempty"`

	// X8021X
	// Constraints:
	//    - nullable
	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

func NewWSGEthernetPortModifyEthernetPortProfile() *WSGEthernetPortModifyEthernetPortProfile {
	m := new(WSGEthernetPortModifyEthernetPortProfile)
	return m
}

type WSGEthernetPortProfileList struct {
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
	List []*WSGEthernetPortProfileListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGEthernetPortProfileList() *WSGEthernetPortProfileList {
	m := new(WSGEthernetPortProfileList)
	return m
}

type WSGEthernetPortProfileListType struct {
	// Id
	// Identifier of the service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGEthernetPortProfileListType() *WSGEthernetPortProfileListType {
	m := new(WSGEthernetPortProfileListType)
	return m
}
