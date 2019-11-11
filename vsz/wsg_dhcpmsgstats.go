package vsz

// API Version: v8_1

// WSGDHCPMsgStats
//
// DHCP Message Statistic Per AP
type WSGDHCPMsgStats struct {
	ApMac *string `json:"apMac,omitempty"`

	// DhcpMsgRecvdStats
	// DHCP Message Received Statistic
	DhcpMsgRecvdStats *WSGDHCPMsgStatsDhcpMsgRecvdStatsType `json:"dhcpMsgRecvdStats,omitempty"`

	// DhcpMsgSentStats
	// DHCP Message Sent Statistic
	DhcpMsgSentStats *WSGDHCPMsgStatsDhcpMsgSentStatsType `json:"dhcpMsgSentStats,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`
}

// WSGDHCPMsgStatsDhcpMsgRecvdStatsType
//
// DHCP Message Received Statistic
type WSGDHCPMsgStatsDhcpMsgRecvdStatsType struct {
	DhcpDecline *int `json:"dhcpDecline,omitempty"`

	DhcpDiscover *int `json:"dhcpDiscover,omitempty"`

	DhcpInform *int `json:"dhcpInform,omitempty"`

	DhcpRelease *int `json:"dhcpRelease,omitempty"`

	DhcpRequest *int `json:"dhcpRequest,omitempty"`
}

// WSGDHCPMsgStatsDhcpMsgSentStatsType
//
// DHCP Message Sent Statistic
type WSGDHCPMsgStatsDhcpMsgSentStatsType struct {
	DhcpAck *int `json:"dhcpAck,omitempty"`

	DhcpNak *int `json:"dhcpNak,omitempty"`

	DhcpOffer *int `json:"dhcpOffer,omitempty"`
}
