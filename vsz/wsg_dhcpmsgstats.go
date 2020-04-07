package vsz

// API Version: v9_0

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

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`
}

func NewWSGDHCPMsgStats() *WSGDHCPMsgStats {
	m := new(WSGDHCPMsgStats)
	return m
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

func NewWSGDHCPMsgStatsDhcpMsgRecvdStatsType() *WSGDHCPMsgStatsDhcpMsgRecvdStatsType {
	m := new(WSGDHCPMsgStatsDhcpMsgRecvdStatsType)
	return m
}

// WSGDHCPMsgStatsDhcpMsgSentStatsType
//
// DHCP Message Sent Statistic
type WSGDHCPMsgStatsDhcpMsgSentStatsType struct {
	DhcpAck *int `json:"dhcpAck,omitempty"`

	DhcpNak *int `json:"dhcpNak,omitempty"`

	DhcpOffer *int `json:"dhcpOffer,omitempty"`
}

func NewWSGDHCPMsgStatsDhcpMsgSentStatsType() *WSGDHCPMsgStatsDhcpMsgSentStatsType {
	m := new(WSGDHCPMsgStatsDhcpMsgSentStatsType)
	return m
}
