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

func NewDhcpMsgStats() *DhcpMsgStats {
	dhcpMsgStatsType := new(DhcpMsgStats)
	return dhcpMsgStatsType
}

func NewDefaultDhcpMsgStats() *DhcpMsgStats {
	dhcpMsgStatsType := new(DhcpMsgStats)
	return dhcpMsgStatsType
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

func NewDhcpMsgStatsDhcpMsgRecvdStatsType() *DhcpMsgStatsDhcpMsgRecvdStatsType {
	dhcpMsgStatsDhcpMsgRecvdStatsTypeType := new(DhcpMsgStatsDhcpMsgRecvdStatsType)
	return dhcpMsgStatsDhcpMsgRecvdStatsTypeType
}

func NewDefaultDhcpMsgStatsDhcpMsgRecvdStatsType() *DhcpMsgStatsDhcpMsgRecvdStatsType {
	dhcpMsgStatsDhcpMsgRecvdStatsTypeType := new(DhcpMsgStatsDhcpMsgRecvdStatsType)
	return dhcpMsgStatsDhcpMsgRecvdStatsTypeType
}

// DhcpMsgStatsDhcpMsgSentStatsType
//
// DHCP Message Sent Statistic
type DhcpMsgStatsDhcpMsgSentStatsType struct {
	DhcpAck *int `json:"dhcpAck,omitempty"`

	DhcpNak *int `json:"dhcpNak,omitempty"`

	DhcpOffer *int `json:"dhcpOffer,omitempty"`
}

func NewDhcpMsgStatsDhcpMsgSentStatsType() *DhcpMsgStatsDhcpMsgSentStatsType {
	dhcpMsgStatsDhcpMsgSentStatsTypeType := new(DhcpMsgStatsDhcpMsgSentStatsType)
	return dhcpMsgStatsDhcpMsgSentStatsTypeType
}

func NewDefaultDhcpMsgStatsDhcpMsgSentStatsType() *DhcpMsgStatsDhcpMsgSentStatsType {
	dhcpMsgStatsDhcpMsgSentStatsTypeType := new(DhcpMsgStatsDhcpMsgSentStatsType)
	return dhcpMsgStatsDhcpMsgSentStatsTypeType
}
