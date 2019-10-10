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

type ApConfiguration struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApGroupID *string `json:"apGroupId,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	BonjourGateway *common.GenericRef `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"gte=60,lte=3600"`

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
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP
	// provisioning process.
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

	WLANGroup24 *WLANGroup `json:"wlanGroup24,omitempty"`

	WLANGroup50 *WLANGroup `json:"wlanGroup50,omitempty"`

	// WLANService24Enabled
	// WLAN service enabled or disabled on 2.4GHz radio
	WLANService24Enabled *bool `json:"wlanService24Enabled,omitempty"`

	// WLANService50Enabled
	// WLAN service enabled or disabled on 5GHz radio
	WLANService50Enabled *bool `json:"wlanService50Enabled,omitempty"`

	// ZoneID
	// Identifier of the AP group to which the AP belongs
	ZoneID *string `json:"zoneId,omitempty"`
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

type ApLinemanSummaryListType struct {
	Alarms *AlarmSummary `json:"alarms,omitempty"`

	// ConfigState
	// State of the AP configuration
	ConfigState *string `json:"configState,omitempty" validate:"oneof=newConfig fwApplied fwDownloaded fwFailed configApplied completed configFailed"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

	// Name
	// Name of the AP
	Name *string `json:"name,omitempty"`
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

type ApListEntryListType struct {
	// ApGroupID
	// Identifier of the AP group to which the AP belongs
	ApGroupID *string `json:"apGroupId,omitempty"`

	Mac *common.Mac `json:"mac,omitempty"`

	// Name
	// Name of the AP
	Name *string `json:"name,omitempty"`

	// Serial
	// Serial Number
	Serial *string `json:"serial,omitempty"`

	// ZoneID
	// Identifier of the zone to which the AP belongs
	ZoneID *string `json:"zoneId,omitempty"`
}

type ApName string

type ApOperationalSummary struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	// ApGroupID
	// Identifier of the AP group to which the AP belongs
	ApGroupID *string `json:"apGroupId,omitempty"`

	// ApprovedTime
	// Timestamp when the AP was approved by the controller
	ApprovedTime *int `json:"approvedTime,omitempty"`

	// ClientCount
	// Number of clients on the AP
	ClientCount *int `json:"clientCount,omitempty"`

	// ConfigState
	// State of the AP configuration.
	ConfigState *string `json:"configState,omitempty" validate:"oneof=completed configApplied configFailed fwApplied fwDownloaded fwFailed newConfig"`

	// ConnectionState
	// Connection state of the AP (value: 'Discovery','Connect','Rebooting','Disconnect','Provisioned')
	ConnectionState *string `json:"connectionState,omitempty"`

	// CountryCode
	// Country code of the AP
	CountryCode *string `json:"countryCode,omitempty"`

	// CpID
	// Identifier of the control plane to which the AP is currently connected
	CpID *string `json:"cpId,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DpID
	// Identifier of the data plane to which the AP is currently connected
	DpID *string `json:"dpId,omitempty"`

	// ExternalIP
	// External IP address of the AP. This is only applicable when the AP is behind a NAT server.
	ExternalIP *string `json:"externalIp,omitempty"`

	// ExternalPort
	// External port number of the AP. This is only applicable when the AP is behind a NAT server.
	ExternalPort *int `json:"externalPort,omitempty"`

	// IP
	// IP address of the AP
	IP *string `json:"ip,omitempty"`

	// IPType
	// Indicates how the AP's IP address was obtained. The AP's IP address can be statically or dynamically
	// assigned or kept unchanged.
	IPType *string `json:"ipType,omitempty" validate:"oneof=Dynamic Keep Static"`

	// IsCriticalAP
	// Indicates critical APs. Critical AP are APs that were tagged by the controller based on predefined
	// rules.
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
	MeshRole *string `json:"meshRole,omitempty" validate:"oneof=Disabled Down Map Root Undefined eMap"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *ApName `json:"name,omitempty"`

	// ProvisionMethod
	// Provisioning method of the AP. Discovered indicates that the AP contacted the controller using
	// discovery and the AP did not have pre-existing record on the controller. Preprovision indicates that
	// the AP was provisioned to the controller before AP made the first contact. Swap indicates that the AP
	// was provisioned to be a replacement of an existing AP.
	ProvisionMethod *string `json:"provisionMethod,omitempty" validate:"oneof=Discovered Preprovision Swap"`

	// ProvisionStage
	// Provisioning stage of the AP. This indicates the stage at which the AP is at in the provisioning
	// process. (value: 'Waiting for Registration','Pre-Provision AP Joined','Waiting for Swap In;Waiting for
	// registration','Waiting for Swap In;Swap In AP Joined','Swapped In;Waiting for registration','Swapped
	// In','Waiting for Swap Out','Swapped Out','Waiting for Swap In, the other AP was deleted','Swapped In,
	// the other AP was deleted','Waiting for Swap Out, the other AP was deleted','Swapped Out, the other AP
	// was deleted')
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

	// ZoneID
	// Identifier of the zone to which the AP belongs
	ZoneID *string `json:"zoneId,omitempty"`
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

type CreateAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

	// ApGroupID
	// Identifier of the AP group to which the AP belongs. If the AP belongs to the default AP group, this
	// property is not needed.
	ApGroupID *string `json:"apGroupId,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	GpsSource *common.ApGpsSource `json:"gpsSource,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	Location *common.Location `json:"location,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	Mac *common.Mac `json:"mac,omitempty" validate:"required"`

	// Model
	// Model name of the AP
	Model *string `json:"model,omitempty"`

	Name *ApName `json:"name,omitempty"`

	// ProvisionChecklist
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP
	// provisioning process.
	ProvisionChecklist *string `json:"provisionChecklist,omitempty"`

	// Serial
	// Serial number of the AP
	Serial *string `json:"serial,omitempty"`

	// ZoneID
	// Identifier of the zone to which the AP belongs
	ZoneID *string `json:"zoneId,omitempty" validate:"required"`
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

type Login struct {
	ApLoginName *common.ApLoginName `json:"apLoginName,omitempty" validate:"required"`

	ApLoginPassword *common.ApLoginPassword `json:"apLoginPassword,omitempty" validate:"required"`
}

type Mesh struct {
	// MeshMode
	// mesh mode
	MeshMode *string `json:"meshMode,omitempty" validate:"oneof=AUTO ROOT_AP MESH_AP DISABLE"`

	// MeshUplinkEntryList
	// MAC address of the neighbor AP
	MeshUplinkEntryList []common.Mac `json:"meshUplinkEntryList,omitempty"`

	// UplinkSelection
	// Uplink selection
	UplinkSelection *string `json:"uplinkSelection,omitempty" validate:"oneof=SMART MANUAL"`
}

type ModifyAP struct {
	// AdministrativeState
	// Administrative state of the AP. A locked AP will not provide any WLAN services.
	AdministrativeState *string `json:"administrativeState,omitempty" validate:"oneof=Locked Unlocked"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	// ApGroupID
	// Identifier of the AP group to which the AP belongs
	ApGroupID *string `json:"apGroupId,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	BonjourGateway *common.GenericRef `json:"bonjourGateway,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the AP
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"gte=60,lte=3600"`

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
	// Provision checklist of the AP. This field indicates the steps that have been completed in the AP
	// provisioning process.
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

	WLANGroup24 *WLANGroup `json:"wlanGroup24,omitempty"`

	WLANGroup50 *WLANGroup `json:"wlanGroup50,omitempty"`

	// WLANService24Enabled
	// WLAN service enabled or disabled on 2.4GHz radio
	WLANService24Enabled *bool `json:"wlanService24Enabled,omitempty"`

	// WLANService50Enabled
	// WLAN service enabled or disabled on 5GHz radio
	WLANService50Enabled *bool `json:"wlanService50Enabled,omitempty"`

	// ZoneID
	// Identifier of the zone to which the AP belongs
	ZoneID *string `json:"zoneId,omitempty"`
}

type ModifyRogueType struct {
	// RogueMacList
	// rogue mac list
	RogueMacList []common.Mac `json:"rogueMacList,omitempty"`
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

type NeighborAPListType struct {
	// Channel
	// Channel of the mesh neighbor AP
	Channel *string `json:"channel,omitempty"`

	// ConnectionState
	// Connection state of the mesh neighbor AP
	ConnectionState *string `json:"connectionState,omitempty"`

	// ExternalIP
	// External IP of the mesh neighbor AP
	ExternalIP *string `json:"externalIp,omitempty"`

	// ExternalPort
	// External port of the mesh neighbor AP
	ExternalPort *string `json:"externalPort,omitempty"`

	// IP
	// IP address of the mesh neighbor AP
	IP *string `json:"ip,omitempty"`

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

type Network struct {
	Gateway *common.IPAddress `json:"gateway,omitempty"`

	IP *common.IPAddress `json:"ip,omitempty"`

	// IPType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically
	// assigned or kept unchanged.
	IPType *string `json:"ipType,omitempty" validate:"oneof=Dynamic Keep Static"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	PrimaryDNS *common.IPAddress `json:"primaryDns,omitempty"`

	SecondaryDNS *common.IPAddress `json:"secondaryDns,omitempty"`
}

type NetworkIpv6 struct {
	Gateway *common.IPAddress `json:"gateway,omitempty"`

	IP *common.IPAddress `json:"ip,omitempty"`

	// IPType
	// Indicates how the AP's IP address was obtained. An AP's IP address can be statically or dynamically
	// assigned or kept unchanged.
	IPType *string `json:"ipType,omitempty" validate:"oneof=Dynamic Keep Static"`

	PrimaryDNS *common.IPAddress `json:"primaryDns,omitempty"`

	SecondaryDNS *common.IPAddress `json:"secondaryDns,omitempty"`
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

	// IPOrFQDN
	// IP or FQDN address of destination cluster, Notice: Once this value been set, clusterName will be
	// ignored.
	IPOrFQDN *string `json:"ipOrFqdn,omitempty"`

	// ZoneIDList
	// Zone ID list for which APs attached to will be switchovered.
	ZoneIDList []string `json:"zoneIdList,omitempty"`
}

type Syslog struct {
	Address *common.IPAddress `json:"address,omitempty"`

	// Enabled
	// Indicates whether syslog is enabled or disabled
	Enabled *bool `json:"enabled,omitempty" validate:"required"`

	// Facility
	// Facility of the syslog server
	Facility *string `json:"facility,omitempty" validate:"oneof=Keep_Original Local0 Local1 Local2 Local3 Local4 Local5 Local6 Local7"`

	// FlowLevel
	// Flow Level of the syslog
	FlowLevel *string `json:"flowLevel,omitempty" validate:"oneof=GENERAL_LOGS CLIENT_FLOW ALL"`

	// Port
	// Port number of the syslog server
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	// Priority
	// Priority of the log messages
	Priority *string `json:"priority,omitempty" validate:"oneof=Emergency Alert Critical Error Warning Notice Info All"`

	// Protocol
	// Protocol of the syslog server
	Protocol *string `json:"protocol,omitempty" validate:"oneof=IPPROTO_TCP IPPROTO_UDP"`

	SecondaryAddress *common.IPAddress `json:"secondaryAddress,omitempty"`

	// SecondaryPort
	// Secondary Server Port of the syslog server
	SecondaryPort *int `json:"secondaryPort,omitempty" validate:"gte=1,lte=65535"`

	// SecondaryProtocol
	// Secondary Server Protocol of the syslog server
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty" validate:"oneof=IPPROTO_TCP IPPROTO_UDP"`
}

type WLANGroup struct {
	// ID
	// Identifier of the WLAN group
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the WLAN group
	Name *string `json:"name,omitempty"`
}
