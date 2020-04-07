package vsz

// API Version: v9_0

// WSGDHCPMsgStats
//
// DHCP Message Statistic Per AP
// Constraints:
//    - nullable
type WSGDHCPMsgStats struct {
	// ApMac
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// DhcpMsgRecvdStats
	// DHCP Message Received Statistic
	// Constraints:
	//    - nullable
	DhcpMsgRecvdStats *WSGDHCPMsgStatsDhcpMsgRecvdStatsType `json:"dhcpMsgRecvdStats,omitempty"`

	// DhcpMsgSentStats
	// DHCP Message Sent Statistic
	// Constraints:
	//    - nullable
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
// Constraints:
//    - nullable
type WSGDHCPMsgStatsDhcpMsgRecvdStatsType struct {
	// DhcpDecline
	// Constraints:
	//    - nullable
	DhcpDecline *int `json:"dhcpDecline,omitempty"`

	// DhcpDiscover
	// Constraints:
	//    - nullable
	DhcpDiscover *int `json:"dhcpDiscover,omitempty"`

	// DhcpInform
	// Constraints:
	//    - nullable
	DhcpInform *int `json:"dhcpInform,omitempty"`

	// DhcpRelease
	// Constraints:
	//    - nullable
	DhcpRelease *int `json:"dhcpRelease,omitempty"`

	// DhcpRequest
	// Constraints:
	//    - nullable
	DhcpRequest *int `json:"dhcpRequest,omitempty"`
}

func NewWSGDHCPMsgStatsDhcpMsgRecvdStatsType() *WSGDHCPMsgStatsDhcpMsgRecvdStatsType {
	m := new(WSGDHCPMsgStatsDhcpMsgRecvdStatsType)
	return m
}

// WSGDHCPMsgStatsDhcpMsgSentStatsType
//
// DHCP Message Sent Statistic
// Constraints:
//    - nullable
type WSGDHCPMsgStatsDhcpMsgSentStatsType struct {
	// DhcpAck
	// Constraints:
	//    - nullable
	DhcpAck *int `json:"dhcpAck,omitempty"`

	// DhcpNak
	// Constraints:
	//    - nullable
	DhcpNak *int `json:"dhcpNak,omitempty"`

	// DhcpOffer
	// Constraints:
	//    - nullable
	DhcpOffer *int `json:"dhcpOffer,omitempty"`
}

func NewWSGDHCPMsgStatsDhcpMsgSentStatsType() *WSGDHCPMsgStatsDhcpMsgSentStatsType {
	m := new(WSGDHCPMsgStatsDhcpMsgSentStatsType)
	return m
}
