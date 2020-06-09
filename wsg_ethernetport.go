package bigdog

// API Version: v9_0

type WSGEthernetPortCreateEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	// Constraints:
	//    - required
	//    - default:'TrunkPort'
	//    - oneof:[AccessPort,TrunkPort,GeneralPort]
	Type *string `json:"type"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - default:1
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty"`

	// UserSidePortEnabled
	// User side port enabled.
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	// X8021X
	// Constraints:
	//    - required
	X8021X *WSGAPModelLanPort8021X `json:"_8021X"`
}

func NewWSGEthernetPortCreateEthernetPortProfile() *WSGEthernetPortCreateEthernetPortProfile {
	m := new(WSGEthernetPortCreateEthernetPortProfile)
	return m
}

type WSGEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty"`

	// Id
	// identifier of the ethernet port profile
	Id *string `json:"id,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	// Constraints:
	//    - oneof:[AccessPort,TrunkPort,GeneralPort]
	Type *string `json:"type,omitempty"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty"`

	// UserSidePortEnabled
	// User side port enabled.
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

func NewWSGEthernetPortProfile() *WSGEthernetPortProfile {
	m := new(WSGEthernetPortProfile)
	return m
}

type WSGEthernetPortModifyEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - default:15
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty"`

	// UserSidePortEnabled
	// User side port enabled.
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

func NewWSGEthernetPortModifyEthernetPortProfile() *WSGEthernetPortModifyEthernetPortProfile {
	m := new(WSGEthernetPortModifyEthernetPortProfile)
	return m
}

type WSGEthernetPortProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGEthernetPortProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGEthernetPortProfileList() *WSGEthernetPortProfileList {
	m := new(WSGEthernetPortProfileList)
	return m
}

type WSGEthernetPortProfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGEthernetPortProfileListType() *WSGEthernetPortProfileListType {
	m := new(WSGEthernetPortProfileListType)
	return m
}
