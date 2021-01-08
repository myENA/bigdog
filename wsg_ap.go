package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// WSGAPAlarmSummary
//
// Definition: ap_alarmSummary
type WSGAPAlarmSummary struct {
	// CriticalCount
	// Critical alarm count
	CriticalCount *int `json:"criticalCount,omitempty"`

	// MajorCount
	// Major alarm count
	MajorCount *int `json:"majorCount,omitempty"`

	// MinorCount
	// Minor alarm count
	MinorCount *int `json:"minorCount,omitempty"`

	// WarningCount
	// Warning alarm count
	WarningCount *int `json:"warningCount,omitempty"`
}

func NewWSGAPAlarmSummary() *WSGAPAlarmSummary {
	m := new(WSGAPAlarmSummary)
	return m
}

// WSGAPConfiguration
//
// Definition: ap_apConfiguration
type WSGAPConfiguration struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - default:'Unlocked'
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty"`

	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	ApGroupId *string `json:"apGroupId,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	BonjourGateway *WSGCommonGenericRef `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource *WSGCommonApGpsSource `json:"gpsSource,omitempty"`

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Login *WSGAPLogin `json:"login,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

	MeshOptions *WSGAPMesh `json:"meshOptions,omitempty"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *WSGAPName `json:"name,omitempty"`

	Network *WSGAPNetwork `json:"network,omitempty"`

	NetworkIpv6 *WSGAPNetworkIpv6 `json:"networkIpv6,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

	SmartMonitor *WSGCommonOverrideSmartMonitor `json:"smartMonitor,omitempty"`

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	Specific *WSGAPModel `json:"specific,omitempty"`

	SwapInMac *string `json:"swapInMac,omitempty"`

	SwapOutMac *string `json:"swapOutMac,omitempty"`

	Syslog *WSGAPSyslog `json:"syslog,omitempty"`

	// TestSpeedEnabled
	// Test Speed is enabled or disabled
	TestSpeedEnabled *bool `json:"testSpeedEnabled,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Wifi24 *WSGCommonRadio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonApRadio50 `json:"wifi50,omitempty"`

	WlanGroup24 *WSGAPWlanGroup `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WSGAPWlanGroup `json:"wlanGroup50,omitempty"`

	// WlanService24Enabled
	// WLAN service enabled or disabled on 2.4GHz radio
	WlanService24Enabled *bool `json:"wlanService24Enabled,omitempty"`

	// WlanService50Enabled
	// WLAN service enabled or disabled on 5GHz radio
	WlanService50Enabled *bool `json:"wlanService50Enabled,omitempty"`

	// ZoneId
	// Identifier of the AP group to which the AP belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGAPConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPConfiguration
}

func newWSGAPConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPConfigurationAPIResponse) Hydrate() error {
	r.Data = new(WSGAPConfiguration)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPConfiguration() *WSGAPConfiguration {
	m := new(WSGAPConfiguration)
	return m
}

// WSGAPLinemanSummary
//
// Definition: ap_apLinemanSummary
type WSGAPLinemanSummary struct {
	// FirstIndex
	// Index of the first AP returned out of the complete AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more APs after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPLinemanSummaryListType `json:"list,omitempty"`

	// TotalCount
	// Total AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAPLinemanSummaryAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPLinemanSummary
}

func newWSGAPLinemanSummaryAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPLinemanSummaryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPLinemanSummaryAPIResponse) Hydrate() error {
	r.Data = new(WSGAPLinemanSummary)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPLinemanSummary() *WSGAPLinemanSummary {
	m := new(WSGAPLinemanSummary)
	return m
}

// WSGAPLinemanSummaryListType
//
// Definition: ap_apLinemanSummaryListType
type WSGAPLinemanSummaryListType struct {
	Alarms *WSGAPAlarmSummary `json:"alarms,omitempty"`

	// ConfigState
	// State of the AP configuration
	// Constraints:
	//    - oneof:[newConfig,fwApplied,fwDownloaded,fwFailed,configApplied,completed,configFailed]
	ConfigState *string `json:"configState,omitempty"`

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

	// Name
	// Name of the AP
	Name *string `json:"name,omitempty"`
}

func NewWSGAPLinemanSummaryListType() *WSGAPLinemanSummaryListType {
	m := new(WSGAPLinemanSummaryListType)
	return m
}

// WSGAPListEntry
//
// Definition: ap_apListEntry
type WSGAPListEntry struct {
	// FirstIndex
	// Index of the first AP returned out of the complete AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more APs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPListEntryListType `json:"list,omitempty"`

	// TotalCount
	// Total AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAPListEntryAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPListEntry
}

func newWSGAPListEntryAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPListEntryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPListEntryAPIResponse) Hydrate() error {
	r.Data = new(WSGAPListEntry)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPListEntry() *WSGAPListEntry {
	m := new(WSGAPListEntry)
	return m
}

// WSGAPListEntryListType
//
// Definition: ap_apListEntryListType
type WSGAPListEntryListType struct {
	// ApGroupId
	// Identifier of the AP group to which the AP belongs
	ApGroupId *string `json:"apGroupId,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

	// Name
	// Name of the AP
	Name *string `json:"name,omitempty"`

	// Serial
	// Serial Number
	Serial *string `json:"serial,omitempty"`

	// ZoneId
	// Identifier of the zone to which the AP belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGAPListEntryListType() *WSGAPListEntryListType {
	m := new(WSGAPListEntryListType)
	return m
}

// WSGAPName
//
// Definition: ap_apName
//
// Constraints:
//    - max:64
//    - min:2
type WSGAPName string

func NewWSGAPName() *WSGAPName {
	m := new(WSGAPName)
	return m
}

// WSGAPOperationalSummary
//
// Definition: ap_apOperationalSummary
type WSGAPOperationalSummary struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty"`

	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApGroupId
	// Identifier of the AP group to which the AP belongs
	ApGroupId *string `json:"apGroupId,omitempty"`

	// ApprovedTime
	// Timestamp when the AP was approved by the controller
	ApprovedTime *int `json:"approvedTime,omitempty"`

	// ClientCount
	// Number of clients on the AP
	ClientCount *int `json:"clientCount,omitempty"`

	// ConfigState
	// State of the AP configuration.
	// Constraints:
	//    - oneof:[newConfig,fwApplied,fwDownloaded,fwFailed,configApplied,completed,configFailed]
	ConfigState *string `json:"configState,omitempty"`

	// ConnectionState
	// Connection state of the AP (value: 'Discovery','Connect','Rebooting','Disconnect','Provisioned')
	ConnectionState *string `json:"connectionState,omitempty"`

	// CountryCode
	// Country code of the AP
	CountryCode *string `json:"countryCode,omitempty"`

	// CpId
	// Identifier of the control plane to which the AP is currently connected
	CpId *string `json:"cpId,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DpId
	// Identifier of the data plane to which the AP is currently connected
	DpId *string `json:"dpId,omitempty"`

	// ExternalIp
	// External IP address of the AP. This is only applicable when the AP is behind a NAT server.
	ExternalIp *string `json:"externalIp,omitempty"`

	// ExternalPort
	// External port number of the AP. This is only applicable when the AP is behind a NAT server.
	ExternalPort *int `json:"externalPort,omitempty"`

	// Ip
	// IP address of the AP
	Ip *string `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. The AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Static,Dynamic,Keep]
	IpType *string `json:"ipType,omitempty"`

	// Ipv6
	// IP address of the AP
	Ipv6 *string `json:"ipv6,omitempty"`

	// Ipv6Type
	// Indicates how the AP's IP address was obtained. The AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Static,Autoconfig,Keep]
	Ipv6Type *string `json:"ipv6Type,omitempty"`

	// IsCriticalAP
	// Indicates critical APs. Critical AP are APs that were tagged by the controller based on predefined rules.
	IsCriticalAP *bool `json:"isCriticalAP,omitempty"`

	// LastSeenTime
	// Timestamp of the last successful communication with the AP
	LastSeenTime *int `json:"lastSeenTime,omitempty"`

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

	// ManagementVlan
	// Management vlan on the AP
	ManagementVlan *int `json:"managementVlan,omitempty"`

	// MeshHop
	// Number of mesh hops of the AP. This is only applicable to mesh APs.
	MeshHop *int `json:"meshHop,omitempty"`

	// MeshRole
	// Mesh role of the AP
	// Constraints:
	//    - oneof:[Disabled,Root,Map,eMap,Down,Undefined]
	MeshRole *string `json:"meshRole,omitempty"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *WSGAPName `json:"name,omitempty"`

	// ProvisionMethod
	// Provisioning method of the AP. Discovered indicates that the AP contacted the controller using discovery and the AP did not have pre-existing record on the controller. Preprovision indicates that the AP was provisioned to the controller before AP made the first contact. Swap indicates that the AP was provisioned to be a replacement of an existing AP.
	// Constraints:
	//    - oneof:[Discovered,Preprovision,Swap]
	ProvisionMethod *string `json:"provisionMethod,omitempty"`

	// ProvisionStage
	// Provisioning stage of the AP. This indicates the stage at which the AP is at in the provisioning process. (value
	ProvisionStage *string `json:"provisionStage,omitempty"`

	// RegistrationState
	// Registration state of the AP
	RegistrationState *string `json:"registrationState,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

	// Uptime
	// Uptime of the AP since it was last rebooted (unit:second)
	Uptime *int `json:"uptime,omitempty"`

	Version *WSGCommonFirmwareVersion `json:"version,omitempty"`

	// Wifi24Channel
	// 2.4GHz radio channel number that the AP is currently using
	Wifi24Channel *string `json:"wifi24Channel,omitempty"`

	// Wifi50Channel
	// 5GHZ radio channel number that the AP is currently using
	Wifi50Channel *string `json:"wifi50Channel,omitempty"`

	// ZoneId
	// Identifier of the zone to which the AP belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGAPOperationalSummaryAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPOperationalSummary
}

func newWSGAPOperationalSummaryAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPOperationalSummaryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPOperationalSummaryAPIResponse) Hydrate() error {
	r.Data = new(WSGAPOperationalSummary)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPOperationalSummary() *WSGAPOperationalSummary {
	m := new(WSGAPOperationalSummary)
	return m
}

// WSGAPRadioConfiguration
//
// Definition: ap_apRadioConfiguration
type WSGAPRadioConfiguration struct {
	Data *WSGAPRadioConfigurationDataType `json:"data,omitempty"`

	Error interface{} `json:"error,omitempty"`

	Extra interface{} `json:"extra,omitempty"`

	MetaData *WSGAPRadioConfigurationMetaDataType `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAPRadioConfiguration) UnmarshalJSON(b []byte) error {
	type _WSGAPRadioConfiguration WSGAPRadioConfiguration
	tmpType := new(_WSGAPRadioConfiguration)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "data")
	delete(tmpType.XAdditionalProperties, "error")
	delete(tmpType.XAdditionalProperties, "extra")
	delete(tmpType.XAdditionalProperties, "metaData")
	delete(tmpType.XAdditionalProperties, "success")
	*t = WSGAPRadioConfiguration(*tmpType)
	return nil
}

func (t *WSGAPRadioConfiguration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Data != nil {
		tmp["data"] = t.Data
	}
	if t.Error != nil {
		tmp["error"] = t.Error
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.MetaData != nil {
		tmp["metaData"] = t.MetaData
	}
	if t.Success != nil {
		tmp["success"] = t.Success
	}
	return json.Marshal(tmp)
}

type WSGAPRadioConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPRadioConfiguration
}

func newWSGAPRadioConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPRadioConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPRadioConfigurationAPIResponse) Hydrate() error {
	r.Data = new(WSGAPRadioConfiguration)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPRadioConfiguration() *WSGAPRadioConfiguration {
	m := new(WSGAPRadioConfiguration)
	return m
}

// WSGAPRadioConfigurationDataType
//
// Definition: ap_apRadioConfigurationDataType
type WSGAPRadioConfigurationDataType struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPRadioConfigurationDataTypeListType `json:"list,omitempty"`

	RawDataTotalCount *float64 `json:"rawDataTotalCount,omitempty"`

	TotalCount *float64 `json:"totalCount,omitempty"`
}

func NewWSGAPRadioConfigurationDataType() *WSGAPRadioConfigurationDataType {
	m := new(WSGAPRadioConfigurationDataType)
	return m
}

// WSGAPRadioConfigurationDataTypeListType
//
// Definition: ap_apRadioConfigurationDataTypeListType
type WSGAPRadioConfigurationDataTypeListType struct {
	ApMac *string `json:"apMac,omitempty"`

	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	BackgroundScan *bool `json:"backgroundScan,omitempty"`

	Busy *float64 `json:"busy,omitempty"`

	Channel *string `json:"channel,omitempty"`

	ChannelBlacklist *string `json:"channelBlacklist,omitempty"`

	ChannelList *string `json:"channelList,omitempty"`

	ChannelWidth *float64 `json:"channelWidth,omitempty"`

	ClientCount *string `json:"clientCount,omitempty"`

	Columns *WSGAPRadioConfigurationDataTypeListTypeColumnsType `json:"columns,omitempty"`

	DeployedWlanNum *float64 `json:"deployedWlanNum,omitempty"`

	Drop *float64 `json:"drop,omitempty"`

	Key *string `json:"key,omitempty"`

	MaxWlanNum *float64 `json:"maxWlanNum,omitempty"`

	Mode *string `json:"mode,omitempty"`

	NoiseFloor *float64 `json:"noiseFloor,omitempty"`

	NullValueColumnNames []interface{} `json:"nullValueColumnNames,omitempty"`

	NumOfAuthorizedClients *float64 `json:"numOfAuthorizedClients,omitempty"`

	PhyError *float64 `json:"phyError,omitempty"`

	RadioId *string `json:"radioId,omitempty"`

	Retry *float64 `json:"retry,omitempty"`

	Rx *float64 `json:"rx,omitempty"`

	RxBytes *float64 `json:"rxBytes,omitempty"`

	RxFrames *float64 `json:"rxFrames,omitempty"`

	RxMulticast *float64 `json:"rxMulticast,omitempty"`

	RxRadioFrames *float64 `json:"rxRadioFrames,omitempty"`

	SecondaryChannel *string `json:"secondaryChannel,omitempty"`

	Total *float64 `json:"total,omitempty"`

	Tx *float64 `json:"tx,omitempty"`

	TxBytes *float64 `json:"txBytes,omitempty"`

	TxFrames *float64 `json:"txFrames,omitempty"`

	TxMulticast *float64 `json:"txMulticast,omitempty"`

	TxPower *string `json:"txPower,omitempty"`

	TxRadioFrames *float64 `json:"txRadioFrames,omitempty"`

	WlanGroupId *string `json:"wlanGroupId,omitempty"`

	WlanGroupName *string `json:"wlanGroupName,omitempty"`

	WlanGroupWlanNum *float64 `json:"wlanGroupWlanNum,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAPRadioConfigurationDataTypeListType) UnmarshalJSON(b []byte) error {
	type _WSGAPRadioConfigurationDataTypeListType WSGAPRadioConfigurationDataTypeListType
	tmpType := new(_WSGAPRadioConfigurationDataTypeListType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "autoCellSizing")
	delete(tmpType.XAdditionalProperties, "backgroundScan")
	delete(tmpType.XAdditionalProperties, "busy")
	delete(tmpType.XAdditionalProperties, "channel")
	delete(tmpType.XAdditionalProperties, "channelBlacklist")
	delete(tmpType.XAdditionalProperties, "channelList")
	delete(tmpType.XAdditionalProperties, "channelWidth")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "columns")
	delete(tmpType.XAdditionalProperties, "deployedWlanNum")
	delete(tmpType.XAdditionalProperties, "drop")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "maxWlanNum")
	delete(tmpType.XAdditionalProperties, "mode")
	delete(tmpType.XAdditionalProperties, "noiseFloor")
	delete(tmpType.XAdditionalProperties, "nullValueColumnNames")
	delete(tmpType.XAdditionalProperties, "numOfAuthorizedClients")
	delete(tmpType.XAdditionalProperties, "phyError")
	delete(tmpType.XAdditionalProperties, "radioId")
	delete(tmpType.XAdditionalProperties, "retry")
	delete(tmpType.XAdditionalProperties, "rx")
	delete(tmpType.XAdditionalProperties, "rxBytes")
	delete(tmpType.XAdditionalProperties, "rxFrames")
	delete(tmpType.XAdditionalProperties, "rxMulticast")
	delete(tmpType.XAdditionalProperties, "rxRadioFrames")
	delete(tmpType.XAdditionalProperties, "secondaryChannel")
	delete(tmpType.XAdditionalProperties, "total")
	delete(tmpType.XAdditionalProperties, "tx")
	delete(tmpType.XAdditionalProperties, "txBytes")
	delete(tmpType.XAdditionalProperties, "txFrames")
	delete(tmpType.XAdditionalProperties, "txMulticast")
	delete(tmpType.XAdditionalProperties, "txPower")
	delete(tmpType.XAdditionalProperties, "txRadioFrames")
	delete(tmpType.XAdditionalProperties, "wlanGroupId")
	delete(tmpType.XAdditionalProperties, "wlanGroupName")
	delete(tmpType.XAdditionalProperties, "wlanGroupWlanNum")
	*t = WSGAPRadioConfigurationDataTypeListType(*tmpType)
	return nil
}

func (t *WSGAPRadioConfigurationDataTypeListType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.AutoCellSizing != nil {
		tmp["autoCellSizing"] = t.AutoCellSizing
	}
	if t.BackgroundScan != nil {
		tmp["backgroundScan"] = t.BackgroundScan
	}
	if t.Busy != nil {
		tmp["busy"] = t.Busy
	}
	if t.Channel != nil {
		tmp["channel"] = t.Channel
	}
	if t.ChannelBlacklist != nil {
		tmp["channelBlacklist"] = t.ChannelBlacklist
	}
	if t.ChannelList != nil {
		tmp["channelList"] = t.ChannelList
	}
	if t.ChannelWidth != nil {
		tmp["channelWidth"] = t.ChannelWidth
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.Columns != nil {
		tmp["columns"] = t.Columns
	}
	if t.DeployedWlanNum != nil {
		tmp["deployedWlanNum"] = t.DeployedWlanNum
	}
	if t.Drop != nil {
		tmp["drop"] = t.Drop
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.MaxWlanNum != nil {
		tmp["maxWlanNum"] = t.MaxWlanNum
	}
	if t.Mode != nil {
		tmp["mode"] = t.Mode
	}
	if t.NoiseFloor != nil {
		tmp["noiseFloor"] = t.NoiseFloor
	}
	if t.NullValueColumnNames != nil {
		tmp["nullValueColumnNames"] = t.NullValueColumnNames
	}
	if t.NumOfAuthorizedClients != nil {
		tmp["numOfAuthorizedClients"] = t.NumOfAuthorizedClients
	}
	if t.PhyError != nil {
		tmp["phyError"] = t.PhyError
	}
	if t.RadioId != nil {
		tmp["radioId"] = t.RadioId
	}
	if t.Retry != nil {
		tmp["retry"] = t.Retry
	}
	if t.Rx != nil {
		tmp["rx"] = t.Rx
	}
	if t.RxBytes != nil {
		tmp["rxBytes"] = t.RxBytes
	}
	if t.RxFrames != nil {
		tmp["rxFrames"] = t.RxFrames
	}
	if t.RxMulticast != nil {
		tmp["rxMulticast"] = t.RxMulticast
	}
	if t.RxRadioFrames != nil {
		tmp["rxRadioFrames"] = t.RxRadioFrames
	}
	if t.SecondaryChannel != nil {
		tmp["secondaryChannel"] = t.SecondaryChannel
	}
	if t.Total != nil {
		tmp["total"] = t.Total
	}
	if t.Tx != nil {
		tmp["tx"] = t.Tx
	}
	if t.TxBytes != nil {
		tmp["txBytes"] = t.TxBytes
	}
	if t.TxFrames != nil {
		tmp["txFrames"] = t.TxFrames
	}
	if t.TxMulticast != nil {
		tmp["txMulticast"] = t.TxMulticast
	}
	if t.TxPower != nil {
		tmp["txPower"] = t.TxPower
	}
	if t.TxRadioFrames != nil {
		tmp["txRadioFrames"] = t.TxRadioFrames
	}
	if t.WlanGroupId != nil {
		tmp["wlanGroupId"] = t.WlanGroupId
	}
	if t.WlanGroupName != nil {
		tmp["wlanGroupName"] = t.WlanGroupName
	}
	if t.WlanGroupWlanNum != nil {
		tmp["wlanGroupWlanNum"] = t.WlanGroupWlanNum
	}
	return json.Marshal(tmp)
}

func NewWSGAPRadioConfigurationDataTypeListType() *WSGAPRadioConfigurationDataTypeListType {
	m := new(WSGAPRadioConfigurationDataTypeListType)
	return m
}

// WSGAPRadioConfigurationDataTypeListTypeColumnsType
//
// Definition: ap_apRadioConfigurationDataTypeListTypeColumnsType
type WSGAPRadioConfigurationDataTypeListTypeColumnsType struct {
	ApMac *string `json:"apMac,omitempty"`

	AutoCellSizing *string `json:"autoCellSizing,omitempty"`

	BackgroundScan *string `json:"backgroundScan,omitempty"`

	Busy *string `json:"busy,omitempty"`

	Channel *string `json:"channel,omitempty"`

	ChannelBlacklist *string `json:"channelBlacklist,omitempty"`

	ChannelList *string `json:"channelList,omitempty"`

	ChannelWidth *string `json:"channelWidth,omitempty"`

	Drop *string `json:"drop,omitempty"`

	Mode *string `json:"mode,omitempty"`

	NoiseFloor *string `json:"noiseFloor,omitempty"`

	NumOfAuthorizedClients *string `json:"numOfAuthorizedClients,omitempty"`

	PhyError *string `json:"phyError,omitempty"`

	RadioId *string `json:"radioId,omitempty"`

	Retry *string `json:"retry,omitempty"`

	Rx *string `json:"rx,omitempty"`

	RxBytes *string `json:"rxBytes,omitempty"`

	RxFrames *string `json:"rxFrames,omitempty"`

	RxMulticast *string `json:"rxMulticast,omitempty"`

	RxRadioFrames *string `json:"rxRadioFrames,omitempty"`

	Total *string `json:"total,omitempty"`

	Tx *string `json:"tx,omitempty"`

	TxBytes *string `json:"txBytes,omitempty"`

	TxFrames *string `json:"txFrames,omitempty"`

	TxMulticast *string `json:"txMulticast,omitempty"`

	TxPower *string `json:"txPower,omitempty"`

	TxRadioFrames *string `json:"txRadioFrames,omitempty"`

	WlanGroupId *string `json:"wlanGroupId,omitempty"`

	WlanGroupName *string `json:"wlanGroupName,omitempty"`
}

func NewWSGAPRadioConfigurationDataTypeListTypeColumnsType() *WSGAPRadioConfigurationDataTypeListTypeColumnsType {
	m := new(WSGAPRadioConfigurationDataTypeListTypeColumnsType)
	return m
}

// WSGAPRadioConfigurationMetaDataType
//
// Definition: ap_apRadioConfigurationMetaDataType
type WSGAPRadioConfigurationMetaDataType struct {
	Fields []*WSGAPRadioConfigurationMetaDataTypeFieldsType `json:"fields,omitempty"`

	IdProperty *string `json:"idProperty,omitempty"`

	MessageProperty *string `json:"messageProperty,omitempty"`

	Root *string `json:"root,omitempty"`

	SuccessProperty *string `json:"successProperty,omitempty"`

	TotalProperty *string `json:"totalProperty,omitempty"`
}

func NewWSGAPRadioConfigurationMetaDataType() *WSGAPRadioConfigurationMetaDataType {
	m := new(WSGAPRadioConfigurationMetaDataType)
	return m
}

// WSGAPRadioConfigurationMetaDataTypeFieldsType
//
// Definition: ap_apRadioConfigurationMetaDataTypeFieldsType
type WSGAPRadioConfigurationMetaDataTypeFieldsType struct {
	Name *string `json:"name,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAPRadioConfigurationMetaDataTypeFieldsType) UnmarshalJSON(b []byte) error {
	type _WSGAPRadioConfigurationMetaDataTypeFieldsType WSGAPRadioConfigurationMetaDataTypeFieldsType
	tmpType := new(_WSGAPRadioConfigurationMetaDataTypeFieldsType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "name")
	*t = WSGAPRadioConfigurationMetaDataTypeFieldsType(*tmpType)
	return nil
}

func (t *WSGAPRadioConfigurationMetaDataTypeFieldsType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	return json.Marshal(tmp)
}

func NewWSGAPRadioConfigurationMetaDataTypeFieldsType() *WSGAPRadioConfigurationMetaDataTypeFieldsType {
	m := new(WSGAPRadioConfigurationMetaDataTypeFieldsType)
	return m
}

// WSGAPCreateAP
//
// Definition: ap_createAP
type WSGAPCreateAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - default:'Unlocked'
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty"`

	// ApGroupId
	// Identifier of the AP group to which the AP belongs. If the AP belongs to the default AP group, this property is not needed.
	ApGroupId *string `json:"apGroupId,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsSource *WSGCommonApGpsSource `json:"gpsSource,omitempty"`

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *WSGAPName `json:"name,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

	// TestSpeedEnabled
	// Test Speed is enabled or disabled
	TestSpeedEnabled *bool `json:"testSpeedEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone to which the AP belongs
	// Constraints:
	//    - required
	ZoneId *string `json:"zoneId"`
}

func NewWSGAPCreateAP() *WSGAPCreateAP {
	m := new(WSGAPCreateAP)
	return m
}

// WSGAPLogin
//
// Definition: ap_login
type WSGAPLogin struct {
	// ApLoginName
	// Constraints:
	//    - required
	ApLoginName *WSGCommonApLoginName `json:"apLoginName"`

	// ApLoginPassword
	// Constraints:
	//    - required
	ApLoginPassword *WSGCommonApLoginPassword `json:"apLoginPassword"`
}

func NewWSGAPLogin() *WSGAPLogin {
	m := new(WSGAPLogin)
	return m
}

// WSGAPMesh
//
// Definition: ap_mesh
type WSGAPMesh struct {
	// MeshMode
	// mesh mode
	// Constraints:
	//    - oneof:[AUTO,ROOT_AP,MESH_AP,DISABLE]
	MeshMode *string `json:"meshMode,omitempty"`

	// MeshUplinkEntryList
	// MAC address of the neighbor AP
	MeshUplinkEntryList []*WSGCommonMac `json:"meshUplinkEntryList,omitempty"`

	// UplinkSelection
	// Uplink selection
	// Constraints:
	//    - oneof:[SMART,MANUAL]
	UplinkSelection *string `json:"uplinkSelection,omitempty"`
}

func NewWSGAPMesh() *WSGAPMesh {
	m := new(WSGAPMesh)
	return m
}

// WSGAPModifyAP
//
// Definition: ap_modifyAP
type WSGAPModifyAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty"`

	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApGroupId
	// Identifier of the AP group to which the AP belongs
	ApGroupId *string `json:"apGroupId,omitempty"`

	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	BonjourGateway *WSGCommonGenericRef `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource *WSGCommonApGpsSource `json:"gpsSource,omitempty"`

	Latitude interface{} `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Login *WSGAPLogin `json:"login,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	MeshOptions *WSGAPMesh `json:"meshOptions,omitempty"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *WSGAPName `json:"name,omitempty"`

	Network *WSGAPNetwork `json:"network,omitempty"`

	NetworkIpv6 *WSGAPNetworkIpv6 `json:"networkIpv6,omitempty"`

	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

	SmartMonitor *WSGCommonOverrideSmartMonitor `json:"smartMonitor,omitempty"`

	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	Syslog *WSGAPSyslog `json:"syslog,omitempty"`

	// TestSpeedEnabled
	// Test Speed is enabled or disabled
	TestSpeedEnabled *bool `json:"testSpeedEnabled,omitempty"`

	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	Wifi24 *WSGCommonRadio24 `json:"wifi24,omitempty"`

	Wifi50 *WSGCommonApRadio50 `json:"wifi50,omitempty"`

	WlanGroup24 *WSGAPWlanGroup `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WSGAPWlanGroup `json:"wlanGroup50,omitempty"`

	// WlanService24Enabled
	// WLAN service enabled or disabled on 2.4GHz radio
	WlanService24Enabled *bool `json:"wlanService24Enabled,omitempty"`

	// WlanService50Enabled
	// WLAN service enabled or disabled on 5GHz radio
	WlanService50Enabled *bool `json:"wlanService50Enabled,omitempty"`

	// ZoneId
	// Identifier of the zone to which the AP belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGAPModifyAP() *WSGAPModifyAP {
	m := new(WSGAPModifyAP)
	return m
}

// WSGAPModifyRogueType
//
// Definition: ap_modifyRogueType
type WSGAPModifyRogueType struct {
	// RogueMacList
	// rogue mac list
	RogueMacList []*WSGCommonMac `json:"rogueMacList,omitempty"`
}

type WSGAPModifyRogueTypeAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPModifyRogueType
}

func newWSGAPModifyRogueTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPModifyRogueTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPModifyRogueTypeAPIResponse) Hydrate() error {
	r.Data = new(WSGAPModifyRogueType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPModifyRogueType() *WSGAPModifyRogueType {
	m := new(WSGAPModifyRogueType)
	return m
}

// WSGAPNeighborAPList
//
// Definition: ap_neighborAPList
type WSGAPNeighborAPList struct {
	// FirstIndex
	// Index of the first Mesh Neighbor AP returned out of the complete Mesh Neighbor AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Mesh Neighbor APs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPNeighborAPListType `json:"list,omitempty"`

	// TotalCount
	// Total mesh neighbor APs count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAPNeighborAPListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPNeighborAPList
}

func newWSGAPNeighborAPListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPNeighborAPListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPNeighborAPListAPIResponse) Hydrate() error {
	r.Data = new(WSGAPNeighborAPList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAPNeighborAPList() *WSGAPNeighborAPList {
	m := new(WSGAPNeighborAPList)
	return m
}

// WSGAPNeighborAPListType
//
// Definition: ap_neighborAPListType
type WSGAPNeighborAPListType struct {
	// Channel
	// Channel of the mesh neighbor AP
	Channel *string `json:"channel,omitempty"`

	// ConnectionState
	// Connection state of the mesh neighbor AP
	ConnectionState *string `json:"connectionState,omitempty"`

	// ExternalIp
	// External IP of the mesh neighbor AP
	ExternalIp *string `json:"externalIp,omitempty"`

	// ExternalPort
	// External port of the mesh neighbor AP
	ExternalPort *string `json:"externalPort,omitempty"`

	// Ip
	// IP address of the mesh neighbor AP
	Ip *string `json:"ip,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

	// Model
	// Model of the mesh neighbor AP
	Model *string `json:"model,omitempty"`

	// Name
	// Name of the mesh neighbor AP
	Name *string `json:"name,omitempty"`

	// Signal
	// Signal of the mesh neighbor AP
	Signal *string `json:"signal,omitempty"`

	// Version
	// Firmware version of the mesh neighbor AP
	Version *string `json:"version,omitempty"`

	// ZoneName
	// Name of the zone which the mesh neighbor AP belongs to
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGAPNeighborAPListType() *WSGAPNeighborAPListType {
	m := new(WSGAPNeighborAPListType)
	return m
}

// WSGAPNetwork
//
// Definition: ap_network
type WSGAPNetwork struct {
	Gateway *WSGCommonIpAddress `json:"gateway,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty"`

	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

	PrimaryDns *WSGCommonIpAddress `json:"primaryDns,omitempty"`

	SecondaryDns *WSGCommonIpAddress `json:"secondaryDns,omitempty"`
}

func NewWSGAPNetwork() *WSGAPNetwork {
	m := new(WSGAPNetwork)
	return m
}

// WSGAPNetworkIpv6
//
// Definition: ap_networkIpv6
type WSGAPNetworkIpv6 struct {
	Gateway *WSGCommonIpAddress `json:"gateway,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty"`

	PrimaryDns *WSGCommonIpAddress `json:"primaryDns,omitempty"`

	SecondaryDns *WSGCommonIpAddress `json:"secondaryDns,omitempty"`
}

func NewWSGAPNetworkIpv6() *WSGAPNetworkIpv6 {
	m := new(WSGAPNetworkIpv6)
	return m
}

// WSGAPSwapApConfigure
//
// Definition: ap_swapApConfigure
type WSGAPSwapApConfigure struct {
	// SwapInMac
	// Constraints:
	//    - required
	SwapInMac *WSGCommonMac `json:"swapInMac"`

	// SwapOutMac
	// Constraints:
	//    - required
	SwapOutMac *WSGCommonMac `json:"swapOutMac"`
}

func NewWSGAPSwapApConfigure() *WSGAPSwapApConfigure {
	m := new(WSGAPSwapApConfigure)
	return m
}

// WSGAPSwitchoverAP
//
// Definition: ap_switchoverAP
type WSGAPSwitchoverAP struct {
	// ApMacList
	// AP MAC address list
	ApMacList []*WSGCommonMac `json:"apMacList,omitempty"`

	// ClusterName
	// Name of destination cluster, Notice: Once user has set ipOrFqdn, this value will be ignored.
	ClusterName *string `json:"clusterName,omitempty"`

	// DeleteRecord
	// Flag to delete AP record after switchover cluster. Default value is false.
	DeleteRecord *bool `json:"deleteRecord,omitempty"`

	// IpOrFqdn
	// IP or FQDN address of destination cluster, Notice: Once this value been set, clusterName will be ignored.
	IpOrFqdn *string `json:"ipOrFqdn,omitempty"`

	// ZoneIdList
	// Zone ID list for which APs attached to will be switchovered.
	ZoneIdList []string `json:"zoneIdList,omitempty"`
}

func NewWSGAPSwitchoverAP() *WSGAPSwitchoverAP {
	m := new(WSGAPSwitchoverAP)
	return m
}

// WSGAPSyslog
//
// Definition: ap_syslog
type WSGAPSyslog struct {
	Address *WSGCommonIpAddress `json:"address,omitempty"`

	// Enabled
	// Indicates whether syslog is enabled or disabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled"`

	// Facility
	// Facility of the syslog server
	// Constraints:
	//    - default:'Keep_Original'
	//    - oneof:[Keep_Original,Local0,Local1,Local2,Local3,Local4,Local5,Local6,Local7]
	Facility *string `json:"facility,omitempty"`

	// FlowLevel
	// Flow Level of the syslog
	// Constraints:
	//    - default:'GENERAL_LOGS'
	//    - oneof:[GENERAL_LOGS,CLIENT_FLOW,ALL]
	FlowLevel *string `json:"flowLevel,omitempty"`

	// Port
	// Port number of the syslog server
	// Constraints:
	//    - default:'514'
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// Priority
	// Priority of the log messages
	// Constraints:
	//    - default:'Error'
	//    - oneof:[Emergency,Alert,Critical,Error,Warning,Notice,Info,All]
	Priority *string `json:"priority,omitempty"`

	// Protocol
	// Protocol of the syslog server
	// Constraints:
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	Protocol *string `json:"protocol,omitempty"`

	SecondaryAddress *WSGCommonIpAddress `json:"secondaryAddress,omitempty"`

	// SecondaryPort
	// Secondary Server Port of the syslog server
	// Constraints:
	//    - default:'514'
	//    - min:1
	//    - max:65535
	SecondaryPort *int `json:"secondaryPort,omitempty"`

	// SecondaryProtocol
	// Secondary Server Protocol of the syslog server
	// Constraints:
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty"`
}

func NewWSGAPSyslog() *WSGAPSyslog {
	m := new(WSGAPSyslog)
	return m
}

// WSGAPWlanGroup
//
// Definition: ap_wlanGroup
type WSGAPWlanGroup struct {
	// Id
	// Identifier of the WLAN group
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the WLAN group
	Name *string `json:"name,omitempty"`
}

func NewWSGAPWlanGroup() *WSGAPWlanGroup {
	m := new(WSGAPWlanGroup)
	return m
}
