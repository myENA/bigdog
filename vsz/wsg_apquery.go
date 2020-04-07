package vsz

// API Version: v9_0

type WSGAPQueryList struct {
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
	List []*WSGAPQueryCreateApQuery `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPQueryList() *WSGAPQueryList {
	m := new(WSGAPQueryList)
	return m
}

type WSGAPQueryCreateApQuery struct {
	// AdministrativeState
	// Constraints:
	//    - nullable
	AdministrativeState *string `json:"administrativeState,omitempty"`

	// Airtime
	// Constraints:
	//    - nullable
	Airtime *int `json:"airtime,omitempty"`

	// Airtime5G
	// Constraints:
	//    - nullable
	Airtime5G *int `json:"airtime5G,omitempty"`

	// Airtime24G
	// Constraints:
	//    - nullable
	Airtime24G *int `json:"airtime24G,omitempty"`

	// Alerts
	// Constraints:
	//    - nullable
	Alerts *int `json:"alerts,omitempty"`

	// ApGroupId
	// Constraints:
	//    - nullable
	ApGroupId *string `json:"apGroupId,omitempty"`

	// ApGroupName
	// Constraints:
	//    - nullable
	ApGroupName *string `json:"apGroupName,omitempty"`

	// ApMac
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// CableModemResetSupported
	// Constraints:
	//    - nullable
	CableModemResetSupported *bool `json:"cableModemResetSupported,omitempty"`

	// CableModemSupported
	// Constraints:
	//    - nullable
	CableModemSupported *bool `json:"cableModemSupported,omitempty"`

	// Capacity
	// Constraints:
	//    - nullable
	Capacity *int `json:"capacity,omitempty"`

	// Capacity24G
	// Constraints:
	//    - nullable
	Capacity24G *int `json:"capacity24G,omitempty"`

	// Capacity50G
	// Constraints:
	//    - nullable
	Capacity50G *int `json:"capacity50G,omitempty"`

	// Cellular3G4GChannel
	// Constraints:
	//    - nullable
	Cellular3G4GChannel *int `json:"cellular3G4GChannel,omitempty"`

	// CellularActiveSim
	// Constraints:
	//    - nullable
	CellularActiveSim *string `json:"cellularActiveSim,omitempty"`

	// CellularConnectionStatus
	// Constraints:
	//    - nullable
	CellularConnectionStatus *string `json:"cellularConnectionStatus,omitempty"`

	// CellularCountry
	// Constraints:
	//    - nullable
	CellularCountry *string `json:"cellularCountry,omitempty"`

	// CellularDefaultGateway
	// Constraints:
	//    - nullable
	CellularDefaultGateway *string `json:"cellularDefaultGateway,omitempty"`

	// CellularGpsHistory
	// Constraints:
	//    - nullable
	CellularGpsHistory []*WSGAPQueryCreateApQueryCellularGpsHistoryType `json:"cellularGpsHistory,omitempty" validate:"omitempty,dive"`

	// CellularICCIDSIM0
	// Constraints:
	//    - nullable
	CellularICCIDSIM0 *string `json:"cellularICCIDSIM0,omitempty"`

	// CellularICCIDSIM1
	// Constraints:
	//    - nullable
	CellularICCIDSIM1 *string `json:"cellularICCIDSIM1,omitempty"`

	// CellularIMSISIM0
	// Constraints:
	//    - nullable
	CellularIMSISIM0 *string `json:"cellularIMSISIM0,omitempty"`

	// CellularIMSISIM1
	// Constraints:
	//    - nullable
	CellularIMSISIM1 *string `json:"cellularIMSISIM1,omitempty"`

	// CellularIPaddress
	// Constraints:
	//    - nullable
	CellularIPaddress *string `json:"cellularIPaddress,omitempty"`

	// CellularIsSIM0Present
	// Constraints:
	//    - nullable
	CellularIsSIM0Present *string `json:"cellularIsSIM0Present,omitempty"`

	// CellularIsSIM1Present
	// Constraints:
	//    - nullable
	CellularIsSIM1Present *string `json:"cellularIsSIM1Present,omitempty"`

	// CellularOperator
	// Constraints:
	//    - nullable
	CellularOperator *string `json:"cellularOperator,omitempty"`

	// CellularRadioUptime
	// Constraints:
	//    - nullable
	CellularRadioUptime *int `json:"cellularRadioUptime,omitempty"`

	// CellularRxBytesSIM0
	// Constraints:
	//    - nullable
	CellularRxBytesSIM0 *int `json:"cellularRxBytesSIM0,omitempty"`

	// CellularRxBytesSIM1
	// Constraints:
	//    - nullable
	CellularRxBytesSIM1 *int `json:"cellularRxBytesSIM1,omitempty"`

	// CellularSignalStrength
	// Constraints:
	//    - nullable
	CellularSignalStrength *string `json:"cellularSignalStrength,omitempty"`

	// CellularSubnetMask
	// Constraints:
	//    - nullable
	CellularSubnetMask *string `json:"cellularSubnetMask,omitempty"`

	// CellularTxBytesSIM0
	// Constraints:
	//    - nullable
	CellularTxBytesSIM0 *int `json:"cellularTxBytesSIM0,omitempty"`

	// CellularTxBytesSIM1
	// Constraints:
	//    - nullable
	CellularTxBytesSIM1 *int `json:"cellularTxBytesSIM1,omitempty"`

	// CellularWanInterface
	// Constraints:
	//    - nullable
	CellularWanInterface *string `json:"cellularWanInterface,omitempty"`

	// Channel5G
	// Constraints:
	//    - nullable
	Channel5G *string `json:"channel5G,omitempty"`

	// Channel24G
	// Constraints:
	//    - nullable
	Channel24G *string `json:"channel24G,omitempty"`

	// Channel24gValue
	// Constraints:
	//    - nullable
	Channel24gValue *int `json:"channel24gValue,omitempty"`

	// Channel50gValue
	// Constraints:
	//    - nullable
	Channel50gValue *int `json:"channel50gValue,omitempty"`

	// ConfigOverride
	// Constraints:
	//    - nullable
	ConfigOverride *bool `json:"configOverride,omitempty"`

	// ConfigurationStatus
	// Constraints:
	//    - nullable
	ConfigurationStatus *string `json:"configurationStatus,omitempty"`

	// ConnectionFailure
	// Constraints:
	//    - nullable
	ConnectionFailure *float64 `json:"connectionFailure,omitempty"`

	// ConnectionStatus
	// Constraints:
	//    - nullable
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	// ControlBladeId
	// Constraints:
	//    - nullable
	ControlBladeId *string `json:"controlBladeId,omitempty"`

	// ControlBladeName
	// Constraints:
	//    - nullable
	ControlBladeName *string `json:"controlBladeName,omitempty"`

	// CrashDump
	// Constraints:
	//    - nullable
	CrashDump *int `json:"crashDump,omitempty"`

	// DataBladeName
	// Constraints:
	//    - nullable
	DataBladeName *string `json:"dataBladeName,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DeviceGps
	// Constraints:
	//    - nullable
	DeviceGps *string `json:"deviceGps,omitempty"`

	// DeviceName
	// Constraints:
	//    - nullable
	DeviceName *string `json:"deviceName,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DomainName
	// Constraints:
	//    - nullable
	DomainName *string `json:"domainName,omitempty"`

	// DpIp
	// Constraints:
	//    - nullable
	DpIp *string `json:"dpIp,omitempty"`

	// Eirp24G
	// Constraints:
	//    - nullable
	Eirp24G *int `json:"eirp24G,omitempty"`

	// Eirp50G
	// Constraints:
	//    - nullable
	Eirp50G *int `json:"eirp50G,omitempty"`

	// EnabledBonjourGateway
	// Constraints:
	//    - nullable
	EnabledBonjourGateway *bool `json:"enabledBonjourGateway,omitempty"`

	// ExtIp
	// Constraints:
	//    - nullable
	ExtIp *string `json:"extIp,omitempty"`

	// ExtPort
	// Constraints:
	//    - nullable
	ExtPort *string `json:"extPort,omitempty"`

	// FipsEnabled
	// Constraints:
	//    - nullable
	FipsEnabled *bool `json:"fipsEnabled,omitempty"`

	// FirmwareVersion
	// Constraints:
	//    - nullable
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// IndoorMapId
	// Constraints:
	//    - nullable
	IndoorMapId *string `json:"indoorMapId,omitempty"`

	// IndoorMapLocation
	// Constraints:
	//    - nullable
	IndoorMapLocation *string `json:"indoorMapLocation,omitempty"`

	// IndoorMapName
	// Constraints:
	//    - nullable
	IndoorMapName *string `json:"indoorMapName,omitempty"`

	// IndoorMapXy
	// Constraints:
	//    - nullable
	IndoorMapXy *WSGAPQueryCreateApQueryIndoorMapXyType `json:"indoorMapXy,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// IpsecRxBytes
	// Constraints:
	//    - nullable
	IpsecRxBytes *int `json:"ipsecRxBytes,omitempty"`

	// IpsecRxDropPkts
	// Constraints:
	//    - nullable
	IpsecRxDropPkts *int `json:"ipsecRxDropPkts,omitempty"`

	// IpsecRxIdleTime
	// Constraints:
	//    - nullable
	IpsecRxIdleTime *int `json:"ipsecRxIdleTime,omitempty"`

	// IpsecRxPkts
	// Constraints:
	//    - nullable
	IpsecRxPkts *int `json:"ipsecRxPkts,omitempty"`

	// IpsecSessionTime
	// Constraints:
	//    - nullable
	IpsecSessionTime *int `json:"ipsecSessionTime,omitempty"`

	// IpsecTxBytes
	// Constraints:
	//    - nullable
	IpsecTxBytes *int `json:"ipsecTxBytes,omitempty"`

	// IpsecTxDropPkts
	// Constraints:
	//    - nullable
	IpsecTxDropPkts *int `json:"ipsecTxDropPkts,omitempty"`

	// IpsecTxIdleTime
	// Constraints:
	//    - nullable
	IpsecTxIdleTime *int `json:"ipsecTxIdleTime,omitempty"`

	// IpsecTxPkts
	// Constraints:
	//    - nullable
	IpsecTxPkts *int `json:"ipsecTxPkts,omitempty"`

	// IpType
	// Constraints:
	//    - nullable
	IpType *string `json:"ipType,omitempty"`

	// Ipv6Address
	// Constraints:
	//    - nullable
	Ipv6Address *string `json:"ipv6Address,omitempty"`

	// Ipv6Type
	// Constraints:
	//    - nullable
	Ipv6Type *string `json:"ipv6Type,omitempty"`

	// IsAirtimeUtilization24GFlagged
	// Constraints:
	//    - nullable
	IsAirtimeUtilization24GFlagged *bool `json:"isAirtimeUtilization24GFlagged,omitempty"`

	// IsAirtimeUtilization50GFlagged
	// Constraints:
	//    - nullable
	IsAirtimeUtilization50GFlagged *bool `json:"isAirtimeUtilization50GFlagged,omitempty"`

	// IsCapacity24GFlagged
	// Constraints:
	//    - nullable
	IsCapacity24GFlagged *bool `json:"isCapacity24GFlagged,omitempty"`

	// IsCapacity50GFlagged
	// Constraints:
	//    - nullable
	IsCapacity50GFlagged *bool `json:"isCapacity50GFlagged,omitempty"`

	// IsConnectionFailure24GFlagged
	// Constraints:
	//    - nullable
	IsConnectionFailure24GFlagged *bool `json:"isConnectionFailure24GFlagged,omitempty"`

	// IsConnectionFailure50GFlagged
	// Constraints:
	//    - nullable
	IsConnectionFailure50GFlagged *bool `json:"isConnectionFailure50GFlagged,omitempty"`

	// IsConnectionFailureFlagged
	// Constraints:
	//    - nullable
	IsConnectionFailureFlagged *bool `json:"isConnectionFailureFlagged,omitempty"`

	// IsConnectionTotalCountFlagged
	// Constraints:
	//    - nullable
	IsConnectionTotalCountFlagged *bool `json:"isConnectionTotalCountFlagged,omitempty"`

	// IsCriticalAp
	// Constraints:
	//    - nullable
	IsCriticalAp *bool `json:"isCriticalAp,omitempty"`

	// IsLatency24GFlagged
	// Constraints:
	//    - nullable
	IsLatency24GFlagged *bool `json:"isLatency24GFlagged,omitempty"`

	// IsLatency50GFlagged
	// Constraints:
	//    - nullable
	IsLatency50GFlagged *bool `json:"isLatency50GFlagged,omitempty"`

	// IsOverallHealthStatusFlagged
	// Constraints:
	//    - nullable
	IsOverallHealthStatusFlagged *bool `json:"isOverallHealthStatusFlagged,omitempty"`

	// LastSeen
	// Constraints:
	//    - nullable
	LastSeen *int `json:"lastSeen,omitempty"`

	// Latency24G
	// Constraints:
	//    - nullable
	Latency24G *int `json:"latency24G,omitempty"`

	// Latency50G
	// Constraints:
	//    - nullable
	Latency50G *int `json:"latency50G,omitempty"`

	// LbsStatus
	// Constraints:
	//    - nullable
	LbsStatus *string `json:"lbsStatus,omitempty"`

	// Location
	// Constraints:
	//    - nullable
	Location *string `json:"location,omitempty"`

	// ManagementVlan
	// Constraints:
	//    - nullable
	ManagementVlan *int `json:"managementVlan,omitempty"`

	// MedianRxRadioMCSRate24G
	// Constraints:
	//    - nullable
	MedianRxRadioMCSRate24G *int `json:"medianRxRadioMCSRate24G,omitempty"`

	// MedianRxRadioMCSRate50G
	// Constraints:
	//    - nullable
	MedianRxRadioMCSRate50G *int `json:"medianRxRadioMCSRate50G,omitempty"`

	// MedianTxRadioMCSRate24G
	// Constraints:
	//    - nullable
	MedianTxRadioMCSRate24G *int `json:"medianTxRadioMCSRate24G,omitempty"`

	// MedianTxRadioMCSRate50G
	// Constraints:
	//    - nullable
	MedianTxRadioMCSRate50G *int `json:"medianTxRadioMCSRate50G,omitempty"`

	// MeshMode
	// Constraints:
	//    - nullable
	MeshMode *string `json:"meshMode,omitempty"`

	// MeshRole
	// Constraints:
	//    - nullable
	MeshRole *string `json:"meshRole,omitempty"`

	// Model
	// Constraints:
	//    - nullable
	Model *string `json:"model,omitempty"`

	// MonitoringEnabled
	// Constraints:
	//    - nullable
	MonitoringEnabled *bool `json:"monitoringEnabled,omitempty"`

	// Noise5G
	// Constraints:
	//    - nullable
	Noise5G *int `json:"noise5G,omitempty"`

	// Noise24G
	// Constraints:
	//    - nullable
	Noise24G *int `json:"noise24G,omitempty"`

	// NumClients
	// Constraints:
	//    - nullable
	NumClients *int `json:"numClients,omitempty"`

	// NumClients5G
	// Constraints:
	//    - nullable
	NumClients5G *int `json:"numClients5G,omitempty"`

	// NumClients24G
	// Constraints:
	//    - nullable
	NumClients24G *int `json:"numClients24G,omitempty"`

	// PacketCaptureState
	// Constraints:
	//    - nullable
	PacketCaptureState *string `json:"packetCaptureState,omitempty"`

	// ProvisionMethod
	// Constraints:
	//    - nullable
	ProvisionMethod *string `json:"provisionMethod,omitempty"`

	// ProvisionStage
	// Constraints:
	//    - nullable
	ProvisionStage *string `json:"provisionStage,omitempty"`

	// RegistrationState
	// Constraints:
	//    - nullable
	RegistrationState *string `json:"registrationState,omitempty"`

	// RegistrationTime
	// Constraints:
	//    - nullable
	RegistrationTime *int `json:"registrationTime,omitempty"`

	// Retry5G
	// Constraints:
	//    - nullable
	Retry5G *int `json:"retry5G,omitempty"`

	// Retry24G
	// Constraints:
	//    - nullable
	Retry24G *int `json:"retry24G,omitempty"`

	// Rx
	// Constraints:
	//    - nullable
	Rx *int `json:"rx,omitempty"`

	// Rx24G
	// Constraints:
	//    - nullable
	Rx24G *int `json:"rx24G,omitempty"`

	// Rx50G
	// Constraints:
	//    - nullable
	Rx50G *int `json:"rx50G,omitempty"`

	// Serial
	// Constraints:
	//    - nullable
	Serial *string `json:"serial,omitempty"`

	// Status
	// Constraints:
	//    - nullable
	//    - oneof:[Online,Offline,Flagged]
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=Online Offline Flagged"`

	// SupportFips
	// Constraints:
	//    - nullable
	SupportFips *bool `json:"supportFips,omitempty"`

	// SwapInMac
	// Constraints:
	//    - nullable
	SwapInMac *string `json:"swapInMac,omitempty"`

	// SwapOutMac
	// Constraints:
	//    - nullable
	SwapOutMac *string `json:"swapOutMac,omitempty"`

	// Tx
	// Constraints:
	//    - nullable
	Tx *int `json:"tx,omitempty"`

	// Tx24G
	// Constraints:
	//    - nullable
	Tx24G *int `json:"tx24G,omitempty"`

	// Tx50G
	// Constraints:
	//    - nullable
	Tx50G *int `json:"tx50G,omitempty"`

	// TxRx
	// Constraints:
	//    - nullable
	TxRx *int `json:"txRx,omitempty"`

	// TxRx24G
	// Constraints:
	//    - nullable
	TxRx24G *int `json:"txRx24G,omitempty"`

	// TxRx50G
	// Constraints:
	//    - nullable
	TxRx50G *int `json:"txRx50G,omitempty"`

	// Uptime
	// Constraints:
	//    - nullable
	Uptime *int `json:"uptime,omitempty"`

	// WlanGroup24Id
	// Constraints:
	//    - nullable
	WlanGroup24Id *string `json:"wlanGroup24Id,omitempty"`

	// WlanGroup24Name
	// Constraints:
	//    - nullable
	WlanGroup24Name *string `json:"wlanGroup24Name,omitempty"`

	// WlanGroup50Id
	// Constraints:
	//    - nullable
	WlanGroup50Id *string `json:"wlanGroup50Id,omitempty"`

	// WlanGroup50Name
	// Constraints:
	//    - nullable
	WlanGroup50Name *string `json:"wlanGroup50Name,omitempty"`

	// ZoneAffinityProfileName
	// Constraints:
	//    - nullable
	ZoneAffinityProfileName *string `json:"zoneAffinityProfileName,omitempty"`

	// ZoneFirmwareVersion
	// Constraints:
	//    - nullable
	ZoneFirmwareVersion *string `json:"zoneFirmwareVersion,omitempty"`

	// ZoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`

	// ZoneName
	// Constraints:
	//    - nullable
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGAPQueryCreateApQuery() *WSGAPQueryCreateApQuery {
	m := new(WSGAPQueryCreateApQuery)
	return m
}

type WSGAPQueryCreateApQueryCellularGpsHistoryType struct {
	// Latitude
	// Constraints:
	//    - nullable
	Latitude *string `json:"latitude,omitempty"`

	// Longitude
	// Constraints:
	//    - nullable
	Longitude *string `json:"longitude,omitempty"`

	// Timestamp
	// Constraints:
	//    - nullable
	Timestamp *int `json:"timestamp,omitempty"`
}

func NewWSGAPQueryCreateApQueryCellularGpsHistoryType() *WSGAPQueryCreateApQueryCellularGpsHistoryType {
	m := new(WSGAPQueryCreateApQueryCellularGpsHistoryType)
	return m
}

type WSGAPQueryCreateApQueryIndoorMapXyType struct {
	// X
	// Constraints:
	//    - nullable
	X *float64 `json:"x,omitempty"`

	// Y
	// Constraints:
	//    - nullable
	Y *float64 `json:"y,omitempty"`
}

func NewWSGAPQueryCreateApQueryIndoorMapXyType() *WSGAPQueryCreateApQueryIndoorMapXyType {
	m := new(WSGAPQueryCreateApQueryIndoorMapXyType)
	return m
}
