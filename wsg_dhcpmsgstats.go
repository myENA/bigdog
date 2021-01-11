package bigdog

// API Version: v9_1

import (
	"errors"
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

func newWSGDHCPMessageStatsDhcpMsgStatsAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDHCPMessageStatsDhcpMsgStatsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDHCPMessageStatsDhcpMsgStatsAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDHCPMessageStatsDhcpMsgStats)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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
