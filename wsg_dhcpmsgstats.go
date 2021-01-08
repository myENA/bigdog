package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// WSGDHCPMessageStatsDhcpMsgStats
//
// Definition: dhcpMsgStats_dhcpMsgStats
//
// DHCP Message Statistic Per AP
type WSGDHCPMessageStatsDhcpMsgStats struct {
	ApMac *string `json:"apMac,omitempty"`

	// DhcpMsgRecvdStats
	// DHCP Message Received Statistic
	DhcpMsgRecvdStats *WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgRecvdStatsType `json:"dhcpMsgRecvdStats,omitempty"`

	// DhcpMsgSentStats
	// DHCP Message Sent Statistic
	DhcpMsgSentStats *WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgSentStatsType `json:"dhcpMsgSentStats,omitempty"`

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

type WSGDHCPMessageStatsDhcpMsgStatsAPIResponse struct {
	*RawAPIResponse
	Data *WSGDHCPMessageStatsDhcpMsgStats
}

func newWSGDHCPMessageStatsDhcpMsgStatsAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDHCPMessageStatsDhcpMsgStatsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDHCPMessageStatsDhcpMsgStatsAPIResponse) Hydrate() error {
	r.Data = new(WSGDHCPMessageStatsDhcpMsgStats)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDHCPMessageStatsDhcpMsgStats() *WSGDHCPMessageStatsDhcpMsgStats {
	m := new(WSGDHCPMessageStatsDhcpMsgStats)
	return m
}

// WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgRecvdStatsType
//
// Definition: dhcpMsgStats_dhcpMsgStatsDhcpMsgRecvdStatsType
//
// DHCP Message Received Statistic
type WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgRecvdStatsType struct {
	DhcpDecline *int `json:"dhcpDecline,omitempty"`

	DhcpDiscover *int `json:"dhcpDiscover,omitempty"`

	DhcpInform *int `json:"dhcpInform,omitempty"`

	DhcpRelease *int `json:"dhcpRelease,omitempty"`

	DhcpRequest *int `json:"dhcpRequest,omitempty"`
}

func NewWSGDHCPMessageStatsDhcpMsgStatsDhcpMsgRecvdStatsType() *WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgRecvdStatsType {
	m := new(WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgRecvdStatsType)
	return m
}

// WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgSentStatsType
//
// Definition: dhcpMsgStats_dhcpMsgStatsDhcpMsgSentStatsType
//
// DHCP Message Sent Statistic
type WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgSentStatsType struct {
	DhcpAck *int `json:"dhcpAck,omitempty"`

	DhcpNak *int `json:"dhcpNak,omitempty"`

	DhcpOffer *int `json:"dhcpOffer,omitempty"`
}

func NewWSGDHCPMessageStatsDhcpMsgStatsDhcpMsgSentStatsType() *WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgSentStatsType {
	m := new(WSGDHCPMessageStatsDhcpMsgStatsDhcpMsgSentStatsType)
	return m
}
