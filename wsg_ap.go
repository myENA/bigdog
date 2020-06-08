package ruckus

// API Version: v9_0

import (
	"encoding/json"
)

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

type WSGAPConfiguration struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - default:'Unlocked'
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty"`

	Altitude **WSGAPConfiguration `json:"altitude,omitempty"`

	ApGroupId *string `json:"apGroupId,omitempty"`

	ApMgmtVlan **WSGAPConfiguration `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 **WSGAPConfiguration `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 **WSGAPConfiguration `json:"autoChannelSelection50,omitempty"`

	AwsVenue **WSGAPConfiguration `json:"awsVenue,omitempty"`

	BonjourGateway **WSGAPConfiguration `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 **WSGAPConfiguration `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 **WSGAPConfiguration `json:"clientAdmissionControl50,omitempty"`

	Description **WSGAPConfiguration `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource **WSGAPConfiguration `json:"gpsSource,omitempty"`

	Latitude **WSGAPConfiguration `json:"latitude,omitempty"`

	Location **WSGAPConfiguration `json:"location,omitempty"`

	LocationAdditionalInfo **WSGAPConfiguration `json:"locationAdditionalInfo,omitempty"`

	Login **WSGAPConfiguration `json:"login,omitempty"`

	Longitude **WSGAPConfiguration `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []**WSGAPConfiguration `json:"lteBandLockChannels,omitempty"`

	Mac **WSGAPConfiguration `json:"mac,omitempty"`

	MeshOptions **WSGAPConfiguration `json:"meshOptions,omitempty"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name **WSGAPConfiguration `json:"name,omitempty"`

	Network **WSGAPConfiguration `json:"network,omitempty"`

	NetworkIpv6 **WSGAPConfiguration `json:"networkIpv6,omitempty"`

	ProtectionMode24 **WSGAPConfiguration `json:"protectionMode24,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	RecoverySsid **WSGAPConfiguration `json:"recoverySsid,omitempty"`

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

	SmartMonitor **WSGAPConfiguration `json:"smartMonitor,omitempty"`

	Specific **WSGAPConfiguration `json:"specific,omitempty"`

	SwapInMac *string `json:"swapInMac,omitempty"`

	SwapOutMac *string `json:"swapOutMac,omitempty"`

	Syslog **WSGAPConfiguration `json:"syslog,omitempty"`

	VenueProfile **WSGAPConfiguration `json:"venueProfile,omitempty"`

	Wifi24 **WSGAPConfiguration `json:"wifi24,omitempty"`

	Wifi50 **WSGAPConfiguration `json:"wifi50,omitempty"`

	WlanGroup24 **WSGAPConfiguration `json:"wlanGroup24,omitempty"`

	WlanGroup50 **WSGAPConfiguration `json:"wlanGroup50,omitempty"`

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

func NewWSGAPConfiguration() *WSGAPConfiguration {
	m := new(WSGAPConfiguration)
	return m
}

type WSGAPLinemanSummary struct {
	// FirstIndex
	// Index of the first AP returned out of the complete AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more APs after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGAPLinemanSummary `json:"list,omitempty"`

	// TotalCount
	// Total AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPLinemanSummary() *WSGAPLinemanSummary {
	m := new(WSGAPLinemanSummary)
	return m
}

type WSGAPLinemanSummaryListType struct {
	Alarms **WSGAPLinemanSummaryListType `json:"alarms,omitempty"`

	// ConfigState
	// State of the AP configuration
	// Constraints:
	//    - oneof:[newConfig,fwApplied,fwDownloaded,fwFailed,configApplied,completed,configFailed]
	ConfigState *string `json:"configState,omitempty"`

	Latitude **WSGAPLinemanSummaryListType `json:"latitude,omitempty"`

	Location **WSGAPLinemanSummaryListType `json:"location,omitempty"`

	Longitude **WSGAPLinemanSummaryListType `json:"longitude,omitempty"`

	Mac **WSGAPLinemanSummaryListType `json:"mac,omitempty"`

	// Name
	// Name of the AP
	Name *string `json:"name,omitempty"`
}

func NewWSGAPLinemanSummaryListType() *WSGAPLinemanSummaryListType {
	m := new(WSGAPLinemanSummaryListType)
	return m
}

type WSGAPListEntry struct {
	// FirstIndex
	// Index of the first AP returned out of the complete AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more APs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGAPListEntry `json:"list,omitempty"`

	// TotalCount
	// Total AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPListEntry() *WSGAPListEntry {
	m := new(WSGAPListEntry)
	return m
}

type WSGAPListEntryListType struct {
	// ApGroupId
	// Identifier of the AP group to which the AP belongs
	ApGroupId *string `json:"apGroupId,omitempty"`

	Mac **WSGAPListEntryListType `json:"mac,omitempty"`

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

type WSGAPName string

func NewWSGAPName() *WSGAPName {
	m := new(WSGAPName)
	return m
}

type WSGAPOperationalSummary struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty"`

	Altitude **WSGAPOperationalSummary `json:"altitude,omitempty"`

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

	Description **WSGAPOperationalSummary `json:"description,omitempty"`

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

	Latitude **WSGAPOperationalSummary `json:"latitude,omitempty"`

	Location **WSGAPOperationalSummary `json:"location,omitempty"`

	LocationAdditionalInfo **WSGAPOperationalSummary `json:"locationAdditionalInfo,omitempty"`

	Longitude **WSGAPOperationalSummary `json:"longitude,omitempty"`

	Mac **WSGAPOperationalSummary `json:"mac,omitempty"`

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

	Name **WSGAPOperationalSummary `json:"name,omitempty"`

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

	Version **WSGAPOperationalSummary `json:"version,omitempty"`

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

func NewWSGAPOperationalSummary() *WSGAPOperationalSummary {
	m := new(WSGAPOperationalSummary)
	return m
}

type WSGAPRadioConfiguration struct {
	Data **WSGAPRadioConfiguration `json:"data,omitempty"`

	Error **WSGAPRadioConfiguration `json:"error,omitempty"`

	Extra **WSGAPRadioConfiguration `json:"extra,omitempty"`

	MetaData **WSGAPRadioConfiguration `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewWSGAPRadioConfiguration() *WSGAPRadioConfiguration {
	m := new(WSGAPRadioConfiguration)
	return m
}

type WSGAPRadioConfigurationDataType struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGAPRadioConfigurationDataType `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPRadioConfigurationDataType() *WSGAPRadioConfigurationDataType {
	m := new(WSGAPRadioConfigurationDataType)
	return m
}

type WSGAPRadioConfigurationDataTypeListType struct {
	ApMac *string `json:"apMac,omitempty"`

	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	BackgroundScan *bool `json:"backgroundScan,omitempty"`

	Busy *int `json:"busy,omitempty"`

	Channel *string `json:"channel,omitempty"`

	ChannelBlacklist *string `json:"channelBlacklist,omitempty"`

	ChannelList *string `json:"channelList,omitempty"`

	ChannelWidth *int `json:"channelWidth,omitempty"`

	ClientCount *int `json:"clientCount,omitempty"`

	Columns **WSGAPRadioConfigurationDataTypeListType `json:"columns,omitempty"`

	DeployedWlanNum *int `json:"deployedWlanNum,omitempty"`

	Drop *int `json:"drop,omitempty"`

	Key *string `json:"key,omitempty"`

	MaxWlanNum *int `json:"maxWlanNum,omitempty"`

	Mode *string `json:"mode,omitempty"`

	NoiseFloor *int `json:"noiseFloor,omitempty"`

	NullValueColumnNames []**WSGAPRadioConfigurationDataTypeListType `json:"nullValueColumnNames,omitempty"`

	NumOfAuthorizedClients *int `json:"numOfAuthorizedClients,omitempty"`

	PhyError *int `json:"phyError,omitempty"`

	RadioId *string `json:"radioId,omitempty"`

	Retry *int `json:"retry,omitempty"`

	Rx *int `json:"rx,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	RxFrames *int `json:"rxFrames,omitempty"`

	RxMulticast *int `json:"rxMulticast,omitempty"`

	RxRadioFrames *int `json:"rxRadioFrames,omitempty"`

	SecondaryChannel *string `json:"secondaryChannel,omitempty"`

	Total *int `json:"total,omitempty"`

	Tx *int `json:"tx,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	TxFrames *int `json:"txFrames,omitempty"`

	TxMulticast *int `json:"txMulticast,omitempty"`

	TxPower *string `json:"txPower,omitempty"`

	TxRadioFrames *int `json:"txRadioFrames,omitempty"`

	WlanGroupId *string `json:"wlanGroupId,omitempty"`

	WlanGroupName *string `json:"wlanGroupName,omitempty"`

	WlanGroupWlanNum *int `json:"wlanGroupWlanNum,omitempty"`
}

func NewWSGAPRadioConfigurationDataTypeListType() *WSGAPRadioConfigurationDataTypeListType {
	m := new(WSGAPRadioConfigurationDataTypeListType)
	return m
}

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

type WSGAPRadioConfigurationDataTypeListTypeNullValueColumnNamesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAPRadioConfigurationDataTypeListTypeNullValueColumnNamesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAPRadioConfigurationDataTypeListTypeNullValueColumnNamesType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAPRadioConfigurationDataTypeListTypeNullValueColumnNamesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAPRadioConfigurationDataTypeListTypeNullValueColumnNamesType() *WSGAPRadioConfigurationDataTypeListTypeNullValueColumnNamesType {
	m := new(WSGAPRadioConfigurationDataTypeListTypeNullValueColumnNamesType)
	return m
}

type WSGAPRadioConfigurationErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAPRadioConfigurationErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAPRadioConfigurationErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAPRadioConfigurationErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAPRadioConfigurationErrorType() *WSGAPRadioConfigurationErrorType {
	m := new(WSGAPRadioConfigurationErrorType)
	return m
}

type WSGAPRadioConfigurationExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAPRadioConfigurationExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAPRadioConfigurationExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAPRadioConfigurationExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAPRadioConfigurationExtraType() *WSGAPRadioConfigurationExtraType {
	m := new(WSGAPRadioConfigurationExtraType)
	return m
}

type WSGAPRadioConfigurationMetaDataType struct {
	Fields []**WSGAPRadioConfigurationMetaDataType `json:"fields,omitempty"`

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

type WSGAPRadioConfigurationMetaDataTypeFieldsType struct {
	Name *string `json:"name,omitempty"`
}

func NewWSGAPRadioConfigurationMetaDataTypeFieldsType() *WSGAPRadioConfigurationMetaDataTypeFieldsType {
	m := new(WSGAPRadioConfigurationMetaDataTypeFieldsType)
	return m
}

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

	AwsVenue **WSGAPCreateAP `json:"awsVenue,omitempty"`

	Description **WSGAPCreateAP `json:"description,omitempty"`

	GpsSource **WSGAPCreateAP `json:"gpsSource,omitempty"`

	Latitude **WSGAPCreateAP `json:"latitude,omitempty"`

	Location **WSGAPCreateAP `json:"location,omitempty"`

	Longitude **WSGAPCreateAP `json:"longitude,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac **WSGAPCreateAP `json:"mac"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name **WSGAPCreateAP `json:"name,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

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

type WSGAPLogin struct {
	// ApLoginName
	// Constraints:
	//    - required
	ApLoginName **WSGAPLogin `json:"apLoginName"`

	// ApLoginPassword
	// Constraints:
	//    - required
	ApLoginPassword **WSGAPLogin `json:"apLoginPassword"`
}

func NewWSGAPLogin() *WSGAPLogin {
	m := new(WSGAPLogin)
	return m
}

type WSGAPMesh struct {
	// MeshMode
	// mesh mode
	// Constraints:
	//    - oneof:[AUTO,ROOT_AP,MESH_AP,DISABLE]
	MeshMode *string `json:"meshMode,omitempty"`

	// MeshUplinkEntryList
	// MAC address of the neighbor AP
	MeshUplinkEntryList []*WSGAPMesh `json:"meshUplinkEntryList,omitempty"`

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

type WSGAPModifyAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty"`

	Altitude **WSGAPModifyAP `json:"altitude,omitempty"`

	// ApGroupId
	// Identifier of the AP group to which the AP belongs
	ApGroupId *string `json:"apGroupId,omitempty"`

	ApMgmtVlan **WSGAPModifyAP `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 **WSGAPModifyAP `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 **WSGAPModifyAP `json:"autoChannelSelection50,omitempty"`

	AwsVenue **WSGAPModifyAP `json:"awsVenue,omitempty"`

	BonjourGateway **WSGAPModifyAP `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	// Constraints:
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	ClientAdmissionControl24 **WSGAPModifyAP `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 **WSGAPModifyAP `json:"clientAdmissionControl50,omitempty"`

	Description **WSGAPModifyAP `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource **WSGAPModifyAP `json:"gpsSource,omitempty"`

	Latitude **WSGAPModifyAP `json:"latitude,omitempty"`

	Location **WSGAPModifyAP `json:"location,omitempty"`

	LocationAdditionalInfo **WSGAPModifyAP `json:"locationAdditionalInfo,omitempty"`

	Login **WSGAPModifyAP `json:"login,omitempty"`

	Longitude **WSGAPModifyAP `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []**WSGAPModifyAP `json:"lteBandLockChannels,omitempty"`

	MeshOptions **WSGAPModifyAP `json:"meshOptions,omitempty"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name **WSGAPModifyAP `json:"name,omitempty"`

	Network **WSGAPModifyAP `json:"network,omitempty"`

	NetworkIpv6 **WSGAPModifyAP `json:"networkIpv6,omitempty"`

	ProtectionMode24 **WSGAPModifyAP `json:"protectionMode24,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	RecoverySsid **WSGAPModifyAP `json:"recoverySsid,omitempty"`

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

	SmartMonitor **WSGAPModifyAP `json:"smartMonitor,omitempty"`

	Syslog **WSGAPModifyAP `json:"syslog,omitempty"`

	VenueProfile **WSGAPModifyAP `json:"venueProfile,omitempty"`

	Wifi24 **WSGAPModifyAP `json:"wifi24,omitempty"`

	Wifi50 **WSGAPModifyAP `json:"wifi50,omitempty"`

	WlanGroup24 **WSGAPModifyAP `json:"wlanGroup24,omitempty"`

	WlanGroup50 **WSGAPModifyAP `json:"wlanGroup50,omitempty"`

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

type WSGAPModifyRogueType struct {
	// RogueMacList
	// rogue mac list
	RogueMacList []*WSGAPModifyRogueType `json:"rogueMacList,omitempty"`
}

func NewWSGAPModifyRogueType() *WSGAPModifyRogueType {
	m := new(WSGAPModifyRogueType)
	return m
}

type WSGAPNeighborAPList struct {
	// FirstIndex
	// Index of the first Mesh Neighbor AP returned out of the complete Mesh Neighbor AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Mesh Neighbor APs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGAPNeighborAPList `json:"list,omitempty"`

	// TotalCount
	// Total mesh neighbor APs count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPNeighborAPList() *WSGAPNeighborAPList {
	m := new(WSGAPNeighborAPList)
	return m
}

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

	Mac **WSGAPNeighborAPListType `json:"mac,omitempty"`

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

type WSGAPNetwork struct {
	Gateway **WSGAPNetwork `json:"gateway,omitempty"`

	Ip **WSGAPNetwork `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty"`

	Netmask **WSGAPNetwork `json:"netmask,omitempty"`

	PrimaryDns **WSGAPNetwork `json:"primaryDns,omitempty"`

	SecondaryDns **WSGAPNetwork `json:"secondaryDns,omitempty"`
}

func NewWSGAPNetwork() *WSGAPNetwork {
	m := new(WSGAPNetwork)
	return m
}

type WSGAPNetworkIpv6 struct {
	Gateway **WSGAPNetworkIpv6 `json:"gateway,omitempty"`

	Ip **WSGAPNetworkIpv6 `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty"`

	PrimaryDns **WSGAPNetworkIpv6 `json:"primaryDns,omitempty"`

	SecondaryDns **WSGAPNetworkIpv6 `json:"secondaryDns,omitempty"`
}

func NewWSGAPNetworkIpv6() *WSGAPNetworkIpv6 {
	m := new(WSGAPNetworkIpv6)
	return m
}

type WSGAPSwapApConfigure struct {
	// SwapInMac
	// Constraints:
	//    - required
	SwapInMac **WSGAPSwapApConfigure `json:"swapInMac"`

	// SwapOutMac
	// Constraints:
	//    - required
	SwapOutMac **WSGAPSwapApConfigure `json:"swapOutMac"`
}

func NewWSGAPSwapApConfigure() *WSGAPSwapApConfigure {
	m := new(WSGAPSwapApConfigure)
	return m
}

type WSGAPSwitchoverAP struct {
	// ApMacList
	// AP MAC address list
	ApMacList []*WSGAPSwitchoverAP `json:"apMacList,omitempty"`

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

type WSGAPSyslog struct {
	Address **WSGAPSyslog `json:"address,omitempty"`

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

	SecondaryAddress **WSGAPSyslog `json:"secondaryAddress,omitempty"`

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
