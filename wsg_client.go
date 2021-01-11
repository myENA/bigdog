package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGClientDeAuthClient
//
// Definition: client_deAuthClient
type WSGClientDeAuthClient struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *WSGCommonMac `json:"apMac"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac"`
}

func NewWSGClientDeAuthClient() *WSGClientDeAuthClient {
	m := new(WSGClientDeAuthClient)
	return m
}

// WSGClientDeAuthClientList
//
// Definition: client_deAuthClientList
type WSGClientDeAuthClientList struct {
	ClientList []*WSGClientDeAuthClient `json:"clientList,omitempty"`
}

func NewWSGClientDeAuthClientList() *WSGClientDeAuthClientList {
	m := new(WSGClientDeAuthClientList)
	return m
}

// WSGClientDisconnectClient
//
// Definition: client_disconnectClient
type WSGClientDisconnectClient struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *WSGCommonMac `json:"apMac"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac"`
}

func NewWSGClientDisconnectClient() *WSGClientDisconnectClient {
	m := new(WSGClientDisconnectClient)
	return m
}

// WSGClientDisconnectClientList
//
// Definition: client_disconnectClientList
type WSGClientDisconnectClientList struct {
	ClientList []*WSGClientDisconnectClient `json:"clientList,omitempty"`
}

func NewWSGClientDisconnectClientList() *WSGClientDisconnectClientList {
	m := new(WSGClientDisconnectClientList)
	return m
}

// WSGClientHistoricalClient
//
// Definition: client_historicalClient
type WSGClientHistoricalClient struct {
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

	// ModelName
	// Model Name of the client
	ModelName *string `json:"modelName,omitempty"`

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

func NewWSGClientHistoricalClient() *WSGClientHistoricalClient {
	m := new(WSGClientHistoricalClient)
	return m
}

// WSGClientHistoricalClientList
//
// Definition: client_historicalClientList
type WSGClientHistoricalClientList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGClientHistoricalClient `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGClientHistoricalClientListAPIResponse struct {
	*RawAPIResponse
	Data *WSGClientHistoricalClientList
}

func newWSGClientHistoricalClientListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGClientHistoricalClientListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGClientHistoricalClientListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGClientHistoricalClientList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGClientHistoricalClientList() *WSGClientHistoricalClientList {
	m := new(WSGClientHistoricalClientList)
	return m
}
