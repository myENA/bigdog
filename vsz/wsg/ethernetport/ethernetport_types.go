package ethernetport

// API Version: v8_0

type CreateEthernetPortProfile struct {
	AntiSpoofingEnabled  *bool                 `json:"antiSpoofingEnabled,omitempty"`
	ArpRequestRateLimit  *int                  `json:"arpRequestRateLimit,omitempty"`
	Description          *string               `json:"description,omitempty"`
	DHCPRequestRateLimit *int                  `json:"dhcpRequestRateLimit,omitempty"`
	DynamicVlanEnabled   *bool                 `json:"dynamicVlanEnabled,omitempty"`
	GuestVlan            *int                  `json:"guestVlan,omitempty"`
	IpsecProfile         *common.GenericRef    `json:"ipsecProfile,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	TunnelEnabled        *bool                 `json:"tunnelEnabled,omitempty"`
	TunnelProfile        *common.GenericRef    `json:"tunnelProfile,omitempty"`
	Type                 *string               `json:"type,omitempty"`
	UntagID              *int                  `json:"untagId,omitempty"`
	VlanMembers          *string               `json:"vlanMembers,omitempty"`
	X8021X               *apmodel.LanPort8021X `json:"_8021X,omitempty"`
}

type EthernetPortProfile struct {
	AntiSpoofingEnabled  *bool                 `json:"antiSpoofingEnabled,omitempty"`
	ArpRequestRateLimit  *int                  `json:"arpRequestRateLimit,omitempty"`
	Description          *string               `json:"description,omitempty"`
	DHCPRequestRateLimit *int                  `json:"dhcpRequestRateLimit,omitempty"`
	DynamicVlanEnabled   *bool                 `json:"dynamicVlanEnabled,omitempty"`
	GuestVlan            *int                  `json:"guestVlan,omitempty"`
	ID                   *string               `json:"id,omitempty"`
	IpsecProfile         *common.GenericRef    `json:"ipsecProfile,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	TunnelEnabled        *bool                 `json:"tunnelEnabled,omitempty"`
	TunnelProfile        *common.GenericRef    `json:"tunnelProfile,omitempty"`
	Type                 *string               `json:"type,omitempty"`
	UntagID              *int                  `json:"untagId,omitempty"`
	VlanMembers          *string               `json:"vlanMembers,omitempty"`
	X8021X               *apmodel.LanPort8021X `json:"_8021X,omitempty"`
}

type ModifyEthernetPortProfile struct {
	AntiSpoofingEnabled  *bool                 `json:"antiSpoofingEnabled,omitempty"`
	ArpRequestRateLimit  *int                  `json:"arpRequestRateLimit,omitempty"`
	Description          *string               `json:"description,omitempty"`
	DHCPRequestRateLimit *int                  `json:"dhcpRequestRateLimit,omitempty"`
	DynamicVlanEnabled   *bool                 `json:"dynamicVlanEnabled,omitempty"`
	GuestVlan            *int                  `json:"guestVlan,omitempty"`
	IpsecProfile         *common.GenericRef    `json:"ipsecProfile,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	TunnelEnabled        *bool                 `json:"tunnelEnabled,omitempty"`
	TunnelProfile        *common.GenericRef    `json:"tunnelProfile,omitempty"`
	UntagID              *int                  `json:"untagId,omitempty"`
	VlanMembers          *string               `json:"vlanMembers,omitempty"`
	X8021X               *apmodel.LanPort8021X `json:"_8021X,omitempty"`
}

type ProfileList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type ProfileListType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
