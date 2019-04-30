package apquery

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/types/common"
)

type ApQueryList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateApQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CreateApQuery struct {
	AdministrativeState *string `json:"administrativeState,omitempty"`

	Airtime *int `json:"airtime,omitempty"`

	Airtime5G *int `json:"airtime5G,omitempty"`

	Airtime24G *int `json:"airtime24G,omitempty"`

	Alerts *int `json:"alerts,omitempty"`

	ApGroupID *string `json:"apGroupId,omitempty"`

	ApGroupName *string `json:"apGroupName,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	CableModemResetSupported *bool `json:"cableModemResetSupported,omitempty"`

	CableModemSupported *bool `json:"cableModemSupported,omitempty"`

	Capacity *int `json:"capacity,omitempty"`

	Capacity24G *int `json:"capacity24G,omitempty"`

	Capacity50G *int `json:"capacity50G,omitempty"`

	Cellular3G4GChannel *int `json:"cellular3G4GChannel,omitempty"`

	CellularActiveSim *string `json:"cellularActiveSim,omitempty"`

	CellularConnectionStatus *string `json:"cellularConnectionStatus,omitempty"`

	CellularCountry *string `json:"cellularCountry,omitempty"`

	CellularDefaultGateway *string `json:"cellularDefaultGateway,omitempty"`

	CellularICCIDSIM0 *string `json:"cellularICCIDSIM0,omitempty"`

	CellularICCIDSIM1 *string `json:"cellularICCIDSIM1,omitempty"`

	CellularIMSISIM0 *string `json:"cellularIMSISIM0,omitempty"`

	CellularIMSISIM1 *string `json:"cellularIMSISIM1,omitempty"`

	CellularIPaddress *string `json:"cellularIPaddress,omitempty"`

	CellularIsSIM0Present *string `json:"cellularIsSIM0Present,omitempty"`

	CellularIsSIM1Present *string `json:"cellularIsSIM1Present,omitempty"`

	CellularOperator *string `json:"cellularOperator,omitempty"`

	CellularRadioUptime *int `json:"cellularRadioUptime,omitempty"`

	CellularRxBytesSIM0 *int `json:"cellularRxBytesSIM0,omitempty"`

	CellularRxBytesSIM1 *int `json:"cellularRxBytesSIM1,omitempty"`

	CellularSignalStrength *string `json:"cellularSignalStrength,omitempty"`

	CellularSubnetMask *string `json:"cellularSubnetMask,omitempty"`

	CellularTxBytesSIM0 *int `json:"cellularTxBytesSIM0,omitempty"`

	CellularTxBytesSIM1 *int `json:"cellularTxBytesSIM1,omitempty"`

	CellularWanInterface *string `json:"cellularWanInterface,omitempty"`

	Channel5G *string `json:"channel5G,omitempty"`

	Channel24G *string `json:"channel24G,omitempty"`

	Channel2GValue *int `json:"channel24gValue,omitempty"`

	Channel5GValue *int `json:"channel50gValue,omitempty"`

	ConfigOverride *bool `json:"configOverride,omitempty"`

	ConfigurationStatus *string `json:"configurationStatus,omitempty"`

	ConnectionFailure *float64 `json:"connectionFailure,omitempty"`

	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	ControlBladeID *string `json:"controlBladeId,omitempty"`

	ControlBladeName *string `json:"controlBladeName,omitempty"`

	CrashDump *int `json:"crashDump,omitempty"`

	DataBladeName *string `json:"dataBladeName,omitempty"`

	Description *string `json:"description,omitempty"`

	DeviceGps *string `json:"deviceGps,omitempty"`

	DeviceName *string `json:"deviceName,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	DpIP *string `json:"dpIp,omitempty"`

	Eirp24G *int `json:"eirp24G,omitempty"`

	Eirp50G *int `json:"eirp50G,omitempty"`

	EnabledBonjourGateway *bool `json:"enabledBonjourGateway,omitempty"`

	ExtIP *string `json:"extIp,omitempty"`

	ExtPort *string `json:"extPort,omitempty"`

	FipsEnabled *bool `json:"fipsEnabled,omitempty"`

	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	IndoorMapID *string `json:"indoorMapId,omitempty"`

	IndoorMapLocation *string `json:"indoorMapLocation,omitempty"`

	IndoorMapName *string `json:"indoorMapName,omitempty"`

	IndoorMapXy *CreateApQueryIndoorMapXyType `json:"indoorMapXy,omitempty"`

	IP *string `json:"ip,omitempty"`

	IpsecRxBytes *int `json:"ipsecRxBytes,omitempty"`

	IpsecRxDropPkts *int `json:"ipsecRxDropPkts,omitempty"`

	IpsecRxIdleTime *int `json:"ipsecRxIdleTime,omitempty"`

	IpsecRxPkts *int `json:"ipsecRxPkts,omitempty"`

	IpsecSessionTime *int `json:"ipsecSessionTime,omitempty"`

	IpsecTxBytes *int `json:"ipsecTxBytes,omitempty"`

	IpsecTxDropPkts *int `json:"ipsecTxDropPkts,omitempty"`

	IpsecTxIdleTime *int `json:"ipsecTxIdleTime,omitempty"`

	IpsecTxPkts *int `json:"ipsecTxPkts,omitempty"`

	IPType *string `json:"ipType,omitempty"`

	Ipv6Address *string `json:"ipv6Address,omitempty"`

	IsAirtimeUtilization24GFlagged *bool `json:"isAirtimeUtilization24GFlagged,omitempty"`

	IsAirtimeUtilization50GFlagged *bool `json:"isAirtimeUtilization50GFlagged,omitempty"`

	IsCapacity24GFlagged *bool `json:"isCapacity24GFlagged,omitempty"`

	IsCapacity50GFlagged *bool `json:"isCapacity50GFlagged,omitempty"`

	IsConnectionFailure24GFlagged *bool `json:"isConnectionFailure24GFlagged,omitempty"`

	IsConnectionFailure50GFlagged *bool `json:"isConnectionFailure50GFlagged,omitempty"`

	IsConnectionFailureFlagged *bool `json:"isConnectionFailureFlagged,omitempty"`

	IsConnectionTotalCountFlagged *bool `json:"isConnectionTotalCountFlagged,omitempty"`

	IsCriticalAp *bool `json:"isCriticalAp,omitempty"`

	IsLatency24GFlagged *bool `json:"isLatency24GFlagged,omitempty"`

	IsLatency50GFlagged *bool `json:"isLatency50GFlagged,omitempty"`

	IsOverallHealthStatusFlagged *bool `json:"isOverallHealthStatusFlagged,omitempty"`

	LastSeen *int `json:"lastSeen,omitempty"`

	Latency24G *int `json:"latency24G,omitempty"`

	Latency50G *int `json:"latency50G,omitempty"`

	LbsStatus *string `json:"lbsStatus,omitempty"`

	Location *string `json:"location,omitempty"`

	ManagementVlan *int `json:"managementVlan,omitempty"`

	MeshMode *string `json:"meshMode,omitempty"`

	MeshRole *string `json:"meshRole,omitempty"`

	Model *string `json:"model,omitempty"`

	Noise5G *int `json:"noise5G,omitempty"`

	Noise24G *int `json:"noise24G,omitempty"`

	NumClients *int `json:"numClients,omitempty"`

	NumClients5G *int `json:"numClients5G,omitempty"`

	NumClients24G *int `json:"numClients24G,omitempty"`

	PacketCaptureState *string `json:"packetCaptureState,omitempty"`

	ProvisionMethod *string `json:"provisionMethod,omitempty"`

	ProvisionStage *string `json:"provisionStage,omitempty"`

	RegistrationState *string `json:"registrationState,omitempty"`

	RegistrationTime *int `json:"registrationTime,omitempty"`

	Retry5G *int `json:"retry5G,omitempty"`

	Retry24G *int `json:"retry24G,omitempty"`

	Rx *int `json:"rx,omitempty"`

	Serial *string `json:"serial,omitempty"`

	Status *string `json:"status,omitempty"`

	SupportFips *bool `json:"supportFips,omitempty"`

	SwapInMac *string `json:"swapInMac,omitempty"`

	SwapOutMac *string `json:"swapOutMac,omitempty"`

	Tx *int `json:"tx,omitempty"`

	TxRx *int `json:"txRx,omitempty"`

	Uptime *int `json:"uptime,omitempty"`

	WLANGroup24ID *string `json:"wlanGroup24Id,omitempty"`

	WLANGroup24Name *string `json:"wlanGroup24Name,omitempty"`

	WLANGroup50ID *string `json:"wlanGroup50Id,omitempty"`

	WLANGroup50Name *string `json:"wlanGroup50Name,omitempty"`

	ZoneAffinityProfileName *string `json:"zoneAffinityProfileName,omitempty"`

	ZoneFirmwareVersion *string `json:"zoneFirmwareVersion,omitempty"`

	ZoneID *string `json:"zoneId,omitempty"`

	ZoneName *string `json:"zoneName,omitempty"`
}

type CreateApQueryIndoorMapXyType struct {
	X *float64 `json:"x,omitempty"`

	Y *float64 `json:"y,omitempty"`
}
