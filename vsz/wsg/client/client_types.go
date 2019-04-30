package client

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/common"
)

type DeAuthClient struct {
	ApMac *common.Mac `json:"apMac,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`
}

type DeAuthClientList struct {
	ClientList []*DeAuthClient `json:"clientList,omitempty"`
}

type DisconnectClient struct {
	ApMac *common.Mac `json:"apMac,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`
}

type DisconnectClientList struct {
	ClientList []*DisconnectClient `json:"clientList,omitempty"`
}

type HistoricalClient struct {
	// ApMac
	// Client connected AP's MAC address
	ApMac *string `json:"apMac,omitempty"`

	// ClientMac
	// Client MAC address
	ClientMac *string `json:"clientMac,omitempty"`

	// CoreNetworkType
	// Core network type of the client
	CoreNetworkType *string `json:"coreNetworkType,omitempty"`

	// Hostname
	// Hostname of the client
	Hostname *string `json:"hostname,omitempty"`

	// IPAddress
	// Client IP address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Ipv6Address
	// Client IPv6 address
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	// MVNOName
	// MVNO name of the client
	MVNOName *string `json:"mvnoName,omitempty"`

	// OsType
	// OS type of the client
	OsType *string `json:"osType,omitempty"`

	// RxBytes
	// Bytes to client
	RxBytes *int `json:"rxBytes,omitempty"`

	// RxDrops
	// Dropped packets to client
	RxDrops *int `json:"rxDrops,omitempty"`

	// RxFrames
	// Bytes to client
	RxFrames *int `json:"rxFrames,omitempty"`

	// SessionEndTime
	// Session end time of the client
	SessionEndTime *int `json:"sessionEndTime,omitempty"`

	// SessionStartTime
	// Session start time of the client
	SessionStartTime *int `json:"sessionStartTime,omitempty"`

	// Ssid
	// Client connected SSID name
	Ssid *string `json:"ssid,omitempty"`

	// TxBytes
	// Bytes from client
	TxBytes *int `json:"txBytes,omitempty"`

	// TxDrops
	// Dropped packets from client
	TxDrops *int `json:"txDrops,omitempty"`

	// TxFrames
	// Bytes from client
	TxFrames *int `json:"txFrames,omitempty"`
}

type HistoricalClientList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*HistoricalClient `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
