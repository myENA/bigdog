package vsz

// API Version: v9_0

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
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

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
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"gte=60,lte=3600"`

	ClientAdmissionControl24 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource *WSGCommonApGpsSource `json:"gpsSource,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Login *WSGAPLogin `json:"login,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels"`

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

	Specific *WSGAPModel `json:"specific,omitempty"`

	SwapInMac *string `json:"swapInMac,omitempty"`

	SwapOutMac *string `json:"swapOutMac,omitempty"`

	Syslog *WSGAPSyslog `json:"syslog,omitempty"`

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

	List []*WSGAPLinemanSummaryListType `json:"list"`

	// TotalCount
	// Total AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAPLinemanSummary() *WSGAPLinemanSummary {
	m := new(WSGAPLinemanSummary)
	return m
}

type WSGAPLinemanSummaryListType struct {
	Alarms *WSGAPAlarmSummary `json:"alarms,omitempty"`

	// ConfigState
	// State of the AP configuration
	// Constraints:
	//    - oneof:[newConfig,fwApplied,fwDownloaded,fwFailed,configApplied,completed,configFailed]
	ConfigState *string `json:"configState,omitempty" validate:"oneof=newConfig fwApplied fwDownloaded fwFailed configApplied completed configFailed"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

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

	List []*WSGAPListEntryListType `json:"list"`

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
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

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
	ConfigState *string `json:"configState,omitempty" validate:"oneof=newConfig fwApplied fwDownloaded fwFailed configApplied completed configFailed"`

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
	IpType *string `json:"ipType,omitempty" validate:"oneof=Static Dynamic Keep"`

	// Ipv6
	// IP address of the AP
	Ipv6 *string `json:"ipv6,omitempty"`

	// Ipv6Type
	// Indicates how the AP's IP address was obtained. The AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Static,Autoconfig,Keep]
	Ipv6Type *string `json:"ipv6Type,omitempty" validate:"oneof=Static Autoconfig Keep"`

	// IsCriticalAP
	// Indicates critical APs. Critical AP are APs that were tagged by the controller based on predefined rules.
	IsCriticalAP *bool `json:"isCriticalAP,omitempty"`

	// LastSeenTime
	// Timestamp of the last successful communication with the AP
	LastSeenTime *int `json:"lastSeenTime,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

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
	MeshRole *string `json:"meshRole,omitempty" validate:"oneof=Disabled Root Map eMap Down Undefined"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *WSGAPName `json:"name,omitempty"`

	// ProvisionMethod
	// Provisioning method of the AP. Discovered indicates that the AP contacted the controller using discovery and the AP did not have pre-existing record on the controller. Preprovision indicates that the AP was provisioned to the controller before AP made the first contact. Swap indicates that the AP was provisioned to be a replacement of an existing AP.
	// Constraints:
	//    - oneof:[Discovered,Preprovision,Swap]
	ProvisionMethod *string `json:"provisionMethod,omitempty" validate:"oneof=Discovered Preprovision Swap"`

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

func NewWSGAPOperationalSummary() *WSGAPOperationalSummary {
	m := new(WSGAPOperationalSummary)
	return m
}

type WSGAPCreateAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	// Constraints:
	//    - default:'Unlocked'
	//    - oneof:[Locked,Unlocked]
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

	// ApGroupId
	// Identifier of the AP group to which the AP belongs. If the AP belongs to the default AP group, this property is not needed.
	ApGroupId *string `json:"apGroupId,omitempty"`

	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsSource *WSGCommonApGpsSource `json:"gpsSource,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`

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

	// ZoneId
	// Identifier of the zone to which the AP belongs
	// Constraints:
	//    - required
	ZoneId *string `json:"zoneId" validate:"required"`
}

func NewWSGAPCreateAP() *WSGAPCreateAP {
	m := new(WSGAPCreateAP)
	return m
}

type WSGAPLogin struct {
	// ApLoginName
	// Constraints:
	//    - required
	ApLoginName *WSGCommonApLoginName `json:"apLoginName" validate:"required,max=64"`

	// ApLoginPassword
	// Constraints:
	//    - required
	ApLoginPassword *WSGCommonApLoginPassword `json:"apLoginPassword" validate:"required,max=64"`
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
	MeshMode *string `json:"meshMode,omitempty" validate:"oneof=AUTO ROOT_AP MESH_AP DISABLE"`

	// MeshUplinkEntryList
	// MAC address of the neighbor AP
	MeshUplinkEntryList []WSGCommonMac `json:"meshUplinkEntryList"`

	// UplinkSelection
	// Uplink selection
	// Constraints:
	//    - oneof:[SMART,MANUAL]
	UplinkSelection *string `json:"uplinkSelection,omitempty" validate:"oneof=SMART MANUAL"`
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
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

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
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"gte=60,lte=3600"`

	ClientAdmissionControl24 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *WSGCommonOverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	GpsSource *WSGCommonApGpsSource `json:"gpsSource,omitempty"`

	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	Location *WSGCommonLocation `json:"location,omitempty"`

	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	Login *WSGAPLogin `json:"login,omitempty"`

	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels"`

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

	Syslog *WSGAPSyslog `json:"syslog,omitempty"`

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

type WSGAPModifyRogueType struct {
	// RogueMacList
	// rogue mac list
	RogueMacList []WSGCommonMac `json:"rogueMacList"`
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

	List []*WSGAPNeighborAPListType `json:"list"`

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

type WSGAPNetwork struct {
	Gateway *WSGCommonIpAddress `json:"gateway,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty" validate:"oneof=Dynamic Keep Static"`

	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

	PrimaryDns *WSGCommonIpAddress `json:"primaryDns,omitempty"`

	SecondaryDns *WSGCommonIpAddress `json:"secondaryDns,omitempty"`
}

func NewWSGAPNetwork() *WSGAPNetwork {
	m := new(WSGAPNetwork)
	return m
}

type WSGAPNetworkIpv6 struct {
	Gateway *WSGCommonIpAddress `json:"gateway,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// IpType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically assigned or kept unchanged.
	// Constraints:
	//    - oneof:[Dynamic,Keep,Static]
	IpType *string `json:"ipType,omitempty" validate:"oneof=Dynamic Keep Static"`

	PrimaryDns *WSGCommonIpAddress `json:"primaryDns,omitempty"`

	SecondaryDns *WSGCommonIpAddress `json:"secondaryDns,omitempty"`
}

func NewWSGAPNetworkIpv6() *WSGAPNetworkIpv6 {
	m := new(WSGAPNetworkIpv6)
	return m
}

type WSGAPSwapApConfigure struct {
	// SwapInMac
	// Constraints:
	//    - required
	SwapInMac *WSGCommonMac `json:"swapInMac" validate:"required"`

	// SwapOutMac
	// Constraints:
	//    - required
	SwapOutMac *WSGCommonMac `json:"swapOutMac" validate:"required"`
}

func NewWSGAPSwapApConfigure() *WSGAPSwapApConfigure {
	m := new(WSGAPSwapApConfigure)
	return m
}

type WSGAPSwitchoverAP struct {
	// ApMacList
	// AP MAC address list
	ApMacList []WSGCommonMac `json:"apMacList"`

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
	ZoneIdList []string `json:"zoneIdList"`
}

func NewWSGAPSwitchoverAP() *WSGAPSwitchoverAP {
	m := new(WSGAPSwitchoverAP)
	return m
}

type WSGAPSyslog struct {
	Address *WSGCommonIpAddress `json:"address,omitempty"`

	// Enabled
	// Indicates whether syslog is enabled or disabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled" validate:"required"`

	// Facility
	// Facility of the syslog server
	// Constraints:
	//    - default:'Keep_Original'
	//    - oneof:[Keep_Original,Local0,Local1,Local2,Local3,Local4,Local5,Local6,Local7]
	Facility *string `json:"facility,omitempty" validate:"oneof=Keep_Original Local0 Local1 Local2 Local3 Local4 Local5 Local6 Local7"`

	// FlowLevel
	// Flow Level of the syslog
	// Constraints:
	//    - default:'GENERAL_LOGS'
	//    - oneof:[GENERAL_LOGS,CLIENT_FLOW,ALL]
	FlowLevel *string `json:"flowLevel,omitempty" validate:"oneof=GENERAL_LOGS CLIENT_FLOW ALL"`

	// Port
	// Port number of the syslog server
	// Constraints:
	//    - default:'514'
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	// Priority
	// Priority of the log messages
	// Constraints:
	//    - default:'Error'
	//    - oneof:[Emergency,Alert,Critical,Error,Warning,Notice,Info,All]
	Priority *string `json:"priority,omitempty" validate:"oneof=Emergency Alert Critical Error Warning Notice Info All"`

	// Protocol
	// Protocol of the syslog server
	// Constraints:
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=IPPROTO_TCP IPPROTO_UDP"`

	SecondaryAddress *WSGCommonIpAddress `json:"secondaryAddress,omitempty"`

	// SecondaryPort
	// Secondary Server Port of the syslog server
	// Constraints:
	//    - default:'514'
	//    - min:1
	//    - max:65535
	SecondaryPort *int `json:"secondaryPort,omitempty" validate:"gte=1,lte=65535"`

	// SecondaryProtocol
	// Secondary Server Protocol of the syslog server
	// Constraints:
	//    - default:'IPPROTO_TCP'
	//    - oneof:[IPPROTO_TCP,IPPROTO_UDP]
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty" validate:"oneof=IPPROTO_TCP IPPROTO_UDP"`
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
