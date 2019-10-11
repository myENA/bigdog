package client

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type DeAuthClient struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *common.Mac `json:"apMac" validate:"required"`

	// Mac
	// Constraints:
	//    - required
	Mac *common.Mac `json:"mac" validate:"required"`
}

func NewDeAuthClient() *DeAuthClient {
	deAuthClientType := new(DeAuthClient)
	return deAuthClientType
}

func NewDefaultDeAuthClient() *DeAuthClient {
	deAuthClientType := new(DeAuthClient)
	return deAuthClientType
}

type DeAuthClientList struct {
	ClientList []*DeAuthClient `json:"clientList,omitempty"`
}

func NewDeAuthClientList() *DeAuthClientList {
	deAuthClientListType := new(DeAuthClientList)
	return deAuthClientListType
}

func NewDefaultDeAuthClientList() *DeAuthClientList {
	deAuthClientListType := new(DeAuthClientList)
	return deAuthClientListType
}

type DisconnectClient struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *common.Mac `json:"apMac" validate:"required"`

	// Mac
	// Constraints:
	//    - required
	Mac *common.Mac `json:"mac" validate:"required"`
}

func NewDisconnectClient() *DisconnectClient {
	disconnectClientType := new(DisconnectClient)
	return disconnectClientType
}

func NewDefaultDisconnectClient() *DisconnectClient {
	disconnectClientType := new(DisconnectClient)
	return disconnectClientType
}

type DisconnectClientList struct {
	ClientList []*DisconnectClient `json:"clientList,omitempty"`
}

func NewDisconnectClientList() *DisconnectClientList {
	disconnectClientListType := new(DisconnectClientList)
	return disconnectClientListType
}

func NewDefaultDisconnectClientList() *DisconnectClientList {
	disconnectClientListType := new(DisconnectClientList)
	return disconnectClientListType
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

	// IpAddress
	// Client IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Ipv6Address
	// Client IPv6 address
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	// MvnoName
	// MVNO name of the client
	MvnoName *string `json:"mvnoName,omitempty"`

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

func NewHistoricalClient() *HistoricalClient {
	historicalClientType := new(HistoricalClient)
	return historicalClientType
}

func NewDefaultHistoricalClient() *HistoricalClient {
	historicalClientType := new(HistoricalClient)
	return historicalClientType
}

type HistoricalClientList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*HistoricalClient `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewHistoricalClientList() *HistoricalClientList {
	historicalClientListType := new(HistoricalClientList)
	return historicalClientListType
}

func NewDefaultHistoricalClientList() *HistoricalClientList {
	historicalClientListType := new(HistoricalClientList)
	return historicalClientListType
}
