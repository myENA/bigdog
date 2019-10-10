package dhcpmsgstats

// API Version: v8_1

// DhcpMsgStats
//
// DHCP Message Statistic Per AP
type DhcpMsgStats struct {
	ApMac *string `json:"apMac,omitempty"`

	// DhcpMsgRecvdStats
	// DHCP Message Received Statistic
	DhcpMsgRecvdStats *DhcpMsgStatsDhcpMsgRecvdStatsType `json:"dhcpMsgRecvdStats,omitempty"`

	// DhcpMsgSentStats
	// DHCP Message Sent Statistic
	DhcpMsgSentStats *DhcpMsgStatsDhcpMsgSentStatsType `json:"dhcpMsgSentStats,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`
}

// DhcpMsgStatsDhcpMsgRecvdStatsType
//
// DHCP Message Received Statistic
type DhcpMsgStatsDhcpMsgRecvdStatsType struct {
	DhcpDecline *int `json:"dhcpDecline,omitempty"`

	DhcpDiscover *int `json:"dhcpDiscover,omitempty"`

	DhcpInform *int `json:"dhcpInform,omitempty"`

	DhcpRelease *int `json:"dhcpRelease,omitempty"`

	DhcpRequest *int `json:"dhcpRequest,omitempty"`
}

// DhcpMsgStatsDhcpMsgSentStatsType
//
// DHCP Message Sent Statistic
type DhcpMsgStatsDhcpMsgSentStatsType struct {
	DhcpAck *int `json:"dhcpAck,omitempty"`

	DhcpNak *int `json:"dhcpNak,omitempty"`

	DhcpOffer *int `json:"dhcpOffer,omitempty"`
}
