package dhcpmsgstats

// API Version: v8_1

// DHCPMsgStats
//
// DHCP Message Statistic Per AP
type DHCPMsgStats struct {
	ApMac *string `json:"apMac,omitempty"`

	// DHCPMsgRecvdStats
	// DHCP Message Received Statistic
	DHCPMsgRecvdStats *DHCPMsgStatsDHCPMsgRecvdStatsType `json:"dhcpMsgRecvdStats,omitempty"`

	// DHCPMsgSentStats
	// DHCP Message Sent Statistic
	DHCPMsgSentStats *DHCPMsgStatsDHCPMsgSentStatsType `json:"dhcpMsgSentStats,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	ID *string `json:"id,omitempty"`

	TenantID *string `json:"tenantId,omitempty"`
}

// DHCPMsgStatsDHCPMsgRecvdStatsType
//
// DHCP Message Received Statistic
type DHCPMsgStatsDHCPMsgRecvdStatsType struct {
	DHCPDecline *int `json:"dhcpDecline,omitempty"`

	DHCPDiscover *int `json:"dhcpDiscover,omitempty"`

	DHCPInform *int `json:"dhcpInform,omitempty"`

	DHCPRelease *int `json:"dhcpRelease,omitempty"`

	DHCPRequest *int `json:"dhcpRequest,omitempty"`
}

// DHCPMsgStatsDHCPMsgSentStatsType
//
// DHCP Message Sent Statistic
type DHCPMsgStatsDHCPMsgSentStatsType struct {
	DHCPAck *int `json:"dhcpAck,omitempty"`

	DHCPNak *int `json:"dhcpNak,omitempty"`

	DHCPOffer *int `json:"dhcpOffer,omitempty"`
}
