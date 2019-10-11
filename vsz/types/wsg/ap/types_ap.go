package ap

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AlarmList struct {
	// FirstIndex
	// Index of the first alarm returned from the complete alarm list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more alarms after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.Alarm `json:"list,omitempty"`

	// TotalCount
	// Total alarm count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAlarmList() *AlarmList {
	alarmListType := new(AlarmList)
	return alarmListType
}

func NewDefaultAlarmList() *AlarmList {
	alarmListType := new(AlarmList)
	return alarmListType
}

type AlarmSummary struct {
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

func NewAlarmSummary() *AlarmSummary {
	alarmSummaryType := new(AlarmSummary)
	return alarmSummaryType
}

func NewDefaultAlarmSummary() *AlarmSummary {
	alarmSummaryType := new(AlarmSummary)
	return alarmSummaryType
}

type ApConfiguration struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - nullable
	//    - default:'Unlocked'
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"omitempty,oneof=Locked Unlocked"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApGroupId *string `json:"apGroupId,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	BonjourGateway *common.GenericRef `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	// Constraints:
	//    - nullable
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"omitempty,gte=60,lte=3600"`

	ClientAdmissionControl24 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource *common.ApGpsSource `json:"gpsSource,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Login *Login `json:"login,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*common.LteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

	MeshOptions *Mesh `json:"meshOptions,omitempty"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *ApName `json:"name,omitempty"`

	Network *Network `json:"network,omitempty"`

	NetworkIpv6 *NetworkIpv6 `json:"networkIpv6,omitempty"`

	ProtectionMode24 *common.ProtectionMode `json:"protectionMode24,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	RecoverySsid *common.RecoverySsid `json:"recoverySsid,omitempty"`

	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

	SmartMonitor *common.OverrideSmartMonitor `json:"smartMonitor,omitempty"`

	Specific *apmodel.ApModel `json:"specific,omitempty"`

	Syslog *Syslog `json:"syslog,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Wifi24 *common.Radio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *common.ApRadio50 `json:"wifi50,omitempty"`

	WlanGroup24 *WlanGroup `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WlanGroup `json:"wlanGroup50,omitempty"`

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

func NewApConfiguration() *ApConfiguration {
	apConfigurationType := new(ApConfiguration)
	return apConfigurationType
}

func NewDefaultApConfiguration() *ApConfiguration {
	apConfigurationType := new(ApConfiguration)
	administrativeStateField := `Unlocked`
	apConfigurationType.AdministrativeState = &administrativeStateField
	channelEvaluationIntervalField := 600
	apConfigurationType.ChannelEvaluationInterval = &channelEvaluationIntervalField
	return apConfigurationType
}

type ApLinemanSummary struct {
	// FirstIndex
	// Index of the first AP returned out of the complete AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more APs after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApLinemanSummaryListType `json:"list,omitempty"`

	// TotalCount
	// Total AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApLinemanSummary() *ApLinemanSummary {
	apLinemanSummaryType := new(ApLinemanSummary)
	return apLinemanSummaryType
}

func NewDefaultApLinemanSummary() *ApLinemanSummary {
	apLinemanSummaryType := new(ApLinemanSummary)
	return apLinemanSummaryType
}

type ApLinemanSummaryListType struct {
	Alarms *AlarmSummary `json:"alarms,omitempty"`

	// ConfigState
	// State of the AP configuration
	// Constraints:
	//    - nullable
	//    - oneof:[newConfig,fwApplied,fwDownloaded,fwFailed,configApplied,completed,configFailed]
	ConfigState *string `json:"configState,omitempty" validate:"omitempty,oneof=newConfig fwApplied fwDownloaded fwFailed configApplied completed configFailed"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

	// Name
	// Name of the AP
	Name *string `json:"name,omitempty"`
}

func NewApLinemanSummaryListType() *ApLinemanSummaryListType {
	apLinemanSummaryListTypeType := new(ApLinemanSummaryListType)
	return apLinemanSummaryListTypeType
}

func NewDefaultApLinemanSummaryListType() *ApLinemanSummaryListType {
	apLinemanSummaryListTypeType := new(ApLinemanSummaryListType)
	return apLinemanSummaryListTypeType
}

type ApListEntry struct {
	// FirstIndex
	// Index of the first AP returned out of the complete AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more APs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApListEntryListType `json:"list,omitempty"`

	// TotalCount
	// Total AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApListEntry() *ApListEntry {
	apListEntryType := new(ApListEntry)
	return apListEntryType
}

func NewDefaultApListEntry() *ApListEntry {
	apListEntryType := new(ApListEntry)
	return apListEntryType
}

type ApListEntryListType struct {
	// ApGroupId
	// Identifier of the AP group to which the AP belongs
	ApGroupId *string `json:"apGroupId,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

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

func NewApListEntryListType() *ApListEntryListType {
	apListEntryListTypeType := new(ApListEntryListType)
	return apListEntryListTypeType
}

func NewDefaultApListEntryListType() *ApListEntryListType {
	apListEntryListTypeType := new(ApListEntryListType)
	return apListEntryListTypeType
}

type ApName string

func NewApName() *ApName {
	apNameType := new(ApName)
	return apNameType
}

func NewDefaultApName() *ApName {
	apNameType := new(ApName)
	return apNameType
}

type ApOperationalSummary struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - nullable
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"omitempty,oneof=Locked Unlocked"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

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
	//    - nullable
	//    - oneof:[completed,configApplied,configFailed,fwApplied,fwDownloaded,fwFailed,newConfig]
	ConfigState *string `json:"configState,omitempty" validate:"omitempty,oneof=completed configApplied configFailed fwApplied fwDownloaded fwFailed newConfig"`

	// ConnectionState
	// Connection state of the AP (value: 'Discovery','Connect','Rebooting','Disconnect','Provisioned')
	ConnectionState *string `json:"connectionState,omitempty"`

	// CountryCode
	// Country code of the AP
	CountryCode *string `json:"countryCode,omitempty"`

	// CpId
	// Identifier of the control plane to which the AP is currently connected
	CpId *string `json:"cpId,omitempty"`

	Description *common.Description `json:"description,omitempty"`

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
	//    - nullable
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty" validate:"omitempty,oneof=Dynamic Keep Static"`

	// IsCriticalAP
	// Indicates critical APs. Critical AP are APs that were tagged by the controller based on predefined rules.
	IsCriticalAP *bool `json:"isCriticalAP,omitempty"`

	// LastSeenTime
	// Timestamp of the last successful communication with the AP
	LastSeenTime *int `json:"lastSeenTime,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

	// ManagementVlan
	// Management vlan on the AP
	ManagementVlan *int `json:"managementVlan,omitempty"`

	// MeshHop
	// Number of mesh hops of the AP. This is only applicable to mesh APs.
	MeshHop *int `json:"meshHop,omitempty"`

	// MeshRole
	// Mesh role of the AP
	// Constraints:
	//    - nullable
	//    - oneof:[Disabled,Down,Map,Root,Undefined,eMap]
	MeshRole *string `json:"meshRole,omitempty" validate:"omitempty,oneof=Disabled Down Map Root Undefined eMap"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *ApName `json:"name,omitempty"`

	// ProvisionMethod
	// Provisioning method of the AP. Discovered indicates that the AP contacted the controller using discovery and the AP did not have pre-existing record on the controller. Preprovision indicates that the AP was provisioned to the controller before AP made the first contact. Swap indicates that the AP was provisioned to be a replacement of an existing AP.
	// Constraints:
	//    - nullable
	//    - oneof:[Discovered,Preprovision,Swap]
	ProvisionMethod *string `json:"provisionMethod,omitempty" validate:"omitempty,oneof=Discovered Preprovision Swap"`

	// ProvisionStage
	// Provisioning stage of the AP. This indicates the stage at which the AP is at in the provisioning process. (value: 'Waiting for Registration','Pre-Provision AP Joined','Waiting for Swap In;Waiting for registration','Waiting for Swap In;Swap In AP Joined','Swapped In;Waiting for registration','Swapped In','Waiting for Swap Out','Swapped Out','Waiting for Swap In, the other AP was deleted','Swapped In, the other AP was deleted','Waiting for Swap Out, the other AP was deleted','Swapped Out, the other AP was deleted')
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

	Version *common.FirmwareVersion `json:"version,omitempty"`

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

func NewApOperationalSummary() *ApOperationalSummary {
	apOperationalSummaryType := new(ApOperationalSummary)
	return apOperationalSummaryType
}

func NewDefaultApOperationalSummary() *ApOperationalSummary {
	apOperationalSummaryType := new(ApOperationalSummary)
	return apOperationalSummaryType
}

type ClientList struct {
	// FirstIndex
	// Index of the first client returned out of the complete client list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more clients after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.Client `json:"list,omitempty"`

	// TotalCount
	// Total client count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewClientList() *ClientList {
	clientListType := new(ClientList)
	return clientListType
}

func NewDefaultClientList() *ClientList {
	clientListType := new(ClientList)
	return clientListType
}

type CreateAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - nullable
	//    - default:'Unlocked'
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"omitempty,oneof=Locked Unlocked"`

	// ApGroupId
	// Identifier of the AP group to which the AP belongs. If the AP belongs to the default AP group, this property is not needed.
	ApGroupId *string `json:"apGroupId,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	GpsSource *common.ApGpsSource `json:"gpsSource,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *common.Mac `json:"mac" validate:"required"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *ApName `json:"name,omitempty"`

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
	ZoneId *string `json:"zoneId" validate:"required"`
}

func NewCreateAP() *CreateAP {
	createAPType := new(CreateAP)
	return createAPType
}

func NewDefaultCreateAP() *CreateAP {
	createAPType := new(CreateAP)
	administrativeStateField := `Unlocked`
	createAPType.AdministrativeState = &administrativeStateField
	return createAPType
}

type EventSummary struct {
	// CriticalCount
	// Critical event count
	CriticalCount *int `json:"criticalCount,omitempty"`

	// DebugCount
	// Debug event count
	DebugCount *int `json:"debugCount,omitempty"`

	// InformationalCount
	// Informational event count
	InformationalCount *int `json:"informationalCount,omitempty"`

	// MajorCount
	// Major event count
	MajorCount *int `json:"majorCount,omitempty"`

	// MinorCount
	// Minor event count
	MinorCount *int `json:"minorCount,omitempty"`

	// WarningCount
	// Warning event count
	WarningCount *int `json:"warningCount,omitempty"`
}

func NewEventSummary() *EventSummary {
	eventSummaryType := new(EventSummary)
	return eventSummaryType
}

func NewDefaultEventSummary() *EventSummary {
	eventSummaryType := new(EventSummary)
	return eventSummaryType
}

type Login struct {
	// ApLoginName
	// Constraints:
	//    - required
	ApLoginName *common.ApLoginName `json:"apLoginName" validate:"required"`

	// ApLoginPassword
	// Constraints:
	//    - required
	ApLoginPassword *common.ApLoginPassword `json:"apLoginPassword" validate:"required"`
}

func NewLogin() *Login {
	loginType := new(Login)
	return loginType
}

func NewDefaultLogin() *Login {
	loginType := new(Login)
	return loginType
}

type Mesh struct {
	// MeshMode
	// mesh mode
	// Constraints:
	//    - nullable
	//    - oneof:[AUTO,ROOT_AP,MESH_AP,DISABLE]
	MeshMode *string `json:"meshMode,omitempty" validate:"omitempty,oneof=AUTO ROOT_AP MESH_AP DISABLE"`

	// MeshUplinkEntryList
	// MAC address of the neighbor AP
	MeshUplinkEntryList []common.Mac `json:"meshUplinkEntryList,omitempty"`

	// UplinkSelection
	// Uplink selection
	// Constraints:
	//    - nullable
	//    - oneof:[SMART,MANUAL]
	UplinkSelection *string `json:"uplinkSelection,omitempty" validate:"omitempty,oneof=SMART MANUAL"`
}

func NewMesh() *Mesh {
	meshType := new(Mesh)
	return meshType
}

func NewDefaultMesh() *Mesh {
	meshType := new(Mesh)
	return meshType
}

type ModifyAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - nullable
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"omitempty,oneof=Locked Unlocked"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	// ApGroupId
	// Identifier of the AP group to which the AP belongs
	ApGroupId *string `json:"apGroupId,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	BonjourGateway *common.GenericRef `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	// Constraints:
	//    - nullable
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"omitempty,gte=60,lte=3600"`

	ClientAdmissionControl24 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource *common.ApGpsSource `json:"gpsSource,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Login *Login `json:"login,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*common.LteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	MeshOptions *Mesh `json:"meshOptions,omitempty"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *ApName `json:"name,omitempty"`

	Network *Network `json:"network,omitempty"`

	NetworkIpv6 *NetworkIpv6 `json:"networkIpv6,omitempty"`

	ProtectionMode24 *common.ProtectionMode `json:"protectionMode24,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	RecoverySsid *common.RecoverySsid `json:"recoverySsid,omitempty"`

	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

	SmartMonitor *common.OverrideSmartMonitor `json:"smartMonitor,omitempty"`

	Syslog *Syslog `json:"syslog,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Wifi24 *common.Radio24 `json:"wifi24,omitempty"`

	Wifi50 *common.ApRadio50 `json:"wifi50,omitempty"`

	WlanGroup24 *WlanGroup `json:"wlanGroup24,omitempty"`

	WlanGroup50 *WlanGroup `json:"wlanGroup50,omitempty"`

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

func NewModifyAP() *ModifyAP {
	modifyAPType := new(ModifyAP)
	return modifyAPType
}

func NewDefaultModifyAP() *ModifyAP {
	modifyAPType := new(ModifyAP)
	channelEvaluationIntervalField := 600
	modifyAPType.ChannelEvaluationInterval = &channelEvaluationIntervalField
	return modifyAPType
}

type ModifyRogueType struct {
	// RogueMacList
	// rogue mac list
	RogueMacList []common.Mac `json:"rogueMacList,omitempty"`
}

func NewModifyRogueType() *ModifyRogueType {
	modifyRogueTypeType := new(ModifyRogueType)
	return modifyRogueTypeType
}

func NewDefaultModifyRogueType() *ModifyRogueType {
	modifyRogueTypeType := new(ModifyRogueType)
	return modifyRogueTypeType
}

type NeighborAPList struct {
	// FirstIndex
	// Index of the first Mesh Neighbor AP returned out of the complete Mesh Neighbor AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Mesh Neighbor APs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*NeighborAPListType `json:"list,omitempty"`

	// TotalCount
	// Total mesh neighbor APs count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewNeighborAPList() *NeighborAPList {
	neighborAPListType := new(NeighborAPList)
	return neighborAPListType
}

func NewDefaultNeighborAPList() *NeighborAPList {
	neighborAPListType := new(NeighborAPList)
	return neighborAPListType
}

type NeighborAPListType struct {
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

	Mac *common.Mac `json:"mac,omitempty"`

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

func NewNeighborAPListType() *NeighborAPListType {
	neighborAPListTypeType := new(NeighborAPListType)
	return neighborAPListTypeType
}

func NewDefaultNeighborAPListType() *NeighborAPListType {
	neighborAPListTypeType := new(NeighborAPListType)
	return neighborAPListTypeType
}

type Network struct {
	Gateway *common.IpAddress `json:"gateway,omitempty"`

	Ip *common.IpAddress `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - nullable
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty" validate:"omitempty,oneof=Dynamic Keep Static"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	PrimaryDns *common.IpAddress `json:"primaryDns,omitempty"`

	SecondaryDns *common.IpAddress `json:"secondaryDns,omitempty"`
}

func NewNetwork() *Network {
	networkType := new(Network)
	return networkType
}

func NewDefaultNetwork() *Network {
	networkType := new(Network)
	return networkType
}

type NetworkIpv6 struct {
	Gateway *common.IpAddress `json:"gateway,omitempty"`

	Ip *common.IpAddress `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - nullable
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty" validate:"omitempty,oneof=Dynamic Keep Static"`

	PrimaryDns *common.IpAddress `json:"primaryDns,omitempty"`

	SecondaryDns *common.IpAddress `json:"secondaryDns,omitempty"`
}

func NewNetworkIpv6() *NetworkIpv6 {
	networkIpv6Type := new(NetworkIpv6)
	return networkIpv6Type
}

func NewDefaultNetworkIpv6() *NetworkIpv6 {
	networkIpv6Type := new(NetworkIpv6)
	return networkIpv6Type
}

type SwitchoverAP struct {
	// ApMacList
	// AP MAC address list
	ApMacList []common.Mac `json:"apMacList,omitempty"`

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

func NewSwitchoverAP() *SwitchoverAP {
	switchoverAPType := new(SwitchoverAP)
	return switchoverAPType
}

func NewDefaultSwitchoverAP() *SwitchoverAP {
	switchoverAPType := new(SwitchoverAP)
	return switchoverAPType
}

type Syslog struct {
	Address *common.IpAddress `json:"address,omitempty"`

	// Enabled
	// Indicates whether syslog is enabled or disabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled" validate:"required"`

	// Facility
	// Facility of the syslog server
	// Constraints:
	//    - nullable
	//    - default:'Keep_Original'
	//    - oneof:[Keep_Original,Local0,Local1,Local2,Local3,Local4,Local5,Local6,Local7]
	Facility *string `json:"facility,omitempty" validate:"omitempty,oneof=Keep_Original Local0 Local1 Local2 Local3 Local4 Local5 Local6 Local7"`

	// FlowLevel
	// Flow Level of the syslog
	// Constraints:
	//    - nullable
	//    - default:'GENERAL_LOGS'
	//    - oneof:[GENERAL_LOGS,CLIENT_FLOW,ALL]
	FlowLevel *string `json:"flowLevel,omitempty" validate:"omitempty,oneof=GENERAL_LOGS CLIENT_FLOW ALL"`

	// Port
	// Port number of the syslog server
	// Constraints:
	//    - nullable
	//    - default:'514'
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// Priority
	// Priority of the log messages
	// Constraints:
	//    - nullable
	//    - default:'Error'
	//    - oneof:[Emergency,Alert,Critical,Error,Warning,Notice,Info,All]
	Priority *string `json:"priority,omitempty" validate:"omitempty,oneof=Emergency Alert Critical Error Warning Notice Info All"`

	// Protocol
	// Protocol of the syslog server
	// Constraints:
	//    - nullable
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=IPPROTO_TCP IPPROTO_UDP"`

	SecondaryAddress *common.IpAddress `json:"secondaryAddress,omitempty"`

	// SecondaryPort
	// Secondary Server Port of the syslog server
	// Constraints:
	//    - nullable
	//    - default:'514'
	//    - min:1
	//    - max:65535
	SecondaryPort *int `json:"secondaryPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// SecondaryProtocol
	// Secondary Server Protocol of the syslog server
	// Constraints:
	//    - nullable
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty" validate:"omitempty,oneof=IPPROTO_TCP IPPROTO_UDP"`
}

func NewSyslog() *Syslog {
	syslogType := new(Syslog)
	return syslogType
}

func NewDefaultSyslog() *Syslog {
	syslogType := new(Syslog)
	facilityField := `Keep_Original`
	syslogType.Facility = &facilityField
	flowLevelField := `GENERAL_LOGS`
	syslogType.FlowLevel = &flowLevelField
	portField := 514
	syslogType.Port = &portField
	priorityField := `Error`
	syslogType.Priority = &priorityField
	protocolField := `IPPROTO_TCP`
	syslogType.Protocol = &protocolField
	secondaryPortField := 514
	syslogType.SecondaryPort = &secondaryPortField
	secondaryProtocolField := `IPPROTO_TCP`
	syslogType.SecondaryProtocol = &secondaryProtocolField
	return syslogType
}

type WlanGroup struct {
	// Id
	// Identifier of the WLAN group
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the WLAN group
	Name *string `json:"name,omitempty"`
}

func NewWlanGroup() *WlanGroup {
	wlanGroupType := new(WlanGroup)
	return wlanGroupType
}

func NewDefaultWlanGroup() *WlanGroup {
	wlanGroupType := new(WlanGroup)
	return wlanGroupType
}
