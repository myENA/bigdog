package vsz

// API Version: v9_0

type WSGClientDeAuthClient struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *WSGCommonMac `json:"apMac" validate:"required"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`
}

func NewWSGClientDeAuthClient() *WSGClientDeAuthClient {
	m := new(WSGClientDeAuthClient)
	return m
}

type WSGClientDeAuthClientList struct {
	// ClientList
	// Constraints:
	//    - nullable
	ClientList []*WSGClientDeAuthClient `json:"clientList,omitempty" validate:"omitempty,dive"`
}

func NewWSGClientDeAuthClientList() *WSGClientDeAuthClientList {
	m := new(WSGClientDeAuthClientList)
	return m
}

type WSGClientDisconnectClient struct {
	// ApMac
	// Constraints:
	//    - required
	ApMac *WSGCommonMac `json:"apMac" validate:"required"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`
}

func NewWSGClientDisconnectClient() *WSGClientDisconnectClient {
	m := new(WSGClientDisconnectClient)
	return m
}

type WSGClientDisconnectClientList struct {
	// ClientList
	// Constraints:
	//    - nullable
	ClientList []*WSGClientDisconnectClient `json:"clientList,omitempty" validate:"omitempty,dive"`
}

func NewWSGClientDisconnectClientList() *WSGClientDisconnectClientList {
	m := new(WSGClientDisconnectClientList)
	return m
}

type WSGClientHistoricalClient struct {
	// ApMac
	// Client connected AP's MAC address
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// ClientMac
	// Client MAC address
	// Constraints:
	//    - nullable
	ClientMac *string `json:"clientMac,omitempty"`

	// CoreNetworkType
	// Core network type of the client
	// Constraints:
	//    - nullable
	CoreNetworkType *string `json:"coreNetworkType,omitempty"`

	// Hostname
	// Hostname of the client
	// Constraints:
	//    - nullable
	Hostname *string `json:"hostname,omitempty"`

	// IpAddress
	// Client IP address
	// Constraints:
	//    - nullable
	IpAddress *string `json:"ipAddress,omitempty"`

	// Ipv6Address
	// Client IPv6 address
	// Constraints:
	//    - nullable
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	// ModelName
	// Model Name of the client
	// Constraints:
	//    - nullable
	ModelName *string `json:"modelName,omitempty"`

	// MvnoName
	// MVNO name of the client
	// Constraints:
	//    - nullable
	MvnoName *string `json:"mvnoName,omitempty"`

	// OsType
	// OS type of the client
	// Constraints:
	//    - nullable
	OsType *string `json:"osType,omitempty"`

	// RxBytes
	// Bytes to client
	// Constraints:
	//    - nullable
	RxBytes *int `json:"rxBytes,omitempty"`

	// RxDrops
	// Dropped packets to client
	// Constraints:
	//    - nullable
	RxDrops *int `json:"rxDrops,omitempty"`

	// RxFrames
	// Bytes to client
	// Constraints:
	//    - nullable
	RxFrames *int `json:"rxFrames,omitempty"`

	// SessionEndTime
	// Session end time of the client
	// Constraints:
	//    - nullable
	SessionEndTime *int `json:"sessionEndTime,omitempty"`

	// SessionStartTime
	// Session start time of the client
	// Constraints:
	//    - nullable
	SessionStartTime *int `json:"sessionStartTime,omitempty"`

	// Ssid
	// Client connected SSID name
	// Constraints:
	//    - nullable
	Ssid *string `json:"ssid,omitempty"`

	// TxBytes
	// Bytes from client
	// Constraints:
	//    - nullable
	TxBytes *int `json:"txBytes,omitempty"`

	// TxDrops
	// Dropped packets from client
	// Constraints:
	//    - nullable
	TxDrops *int `json:"txDrops,omitempty"`

	// TxFrames
	// Bytes from client
	// Constraints:
	//    - nullable
	TxFrames *int `json:"txFrames,omitempty"`
}

func NewWSGClientHistoricalClient() *WSGClientHistoricalClient {
	m := new(WSGClientHistoricalClient)
	return m
}

type WSGClientHistoricalClientList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGClientHistoricalClient `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGClientHistoricalClientList() *WSGClientHistoricalClientList {
	m := new(WSGClientHistoricalClientList)
	return m
}
