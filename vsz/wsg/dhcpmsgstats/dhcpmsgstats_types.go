package dhcpmsgstats

// API Version: v8_0

type DHCPMsgStats struct {
	ApMac             *string            `json:"apMac,omitempty"`
	DHCPMsgRecvdStats *DHCPMsgRecvdStats `json:"dhcpMsgRecvdStats,omitempty"`
	DHCPMsgSentStats  *DHCPMsgSentStats  `json:"dhcpMsgSentStats,omitempty"`
	DomainID          *string            `json:"domainId,omitempty"`
	ID                *string            `json:"id,omitempty"`
	TenantID          *string            `json:"tenantId,omitempty"`
}

type DHCPMsgStatsDHCPMsgRecvdStatsType struct {
	DHCPDecline  *int `json:"dhcpDecline,omitempty"`
	DHCPDiscover *int `json:"dhcpDiscover,omitempty"`
	DHCPInform   *int `json:"dhcpInform,omitempty"`
	DHCPRelease  *int `json:"dhcpRelease,omitempty"`
	DHCPRequest  *int `json:"dhcpRequest,omitempty"`
}

type DHCPMsgStatsDHCPMsgSentStatsType struct {
	DHCPAck   *int `json:"dhcpAck,omitempty"`
	DHCPNak   *int `json:"dhcpNak,omitempty"`
	DHCPOffer *int `json:"dhcpOffer,omitempty"`
}
