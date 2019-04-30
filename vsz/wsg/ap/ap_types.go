package ap

// API Version: v8_0

type AlarmList struct {
	FirstIndex *int            `json:"firstIndex,omitempty"`
	HasMore    *bool           `json:"hasMore,omitempty"`
	List       []*common.Alarm `json:"list,omitempty"`
	TotalCount *int            `json:"totalCount,omitempty"`
}

type AlarmSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`
	MajorCount    *int `json:"majorCount,omitempty"`
	MinorCount    *int `json:"minorCount,omitempty"`
	WarningCount  *int `json:"warningCount,omitempty"`
}

type ApConfiguration struct {
	AdministrativeState                        *string                                `json:"administrativeState,omitempty"`
	Altitude                                   *common.Altitude                       `json:"altitude,omitempty"`
	ApGroupID                                  *string                                `json:"apGroupId,omitempty"`
	ApMgmtVlan                                 *common.ApManagementVlan               `json:"apMgmtVlan,omitempty"`
	AutoChannelSelection24                     *common.AutoChannelSelection           `json:"autoChannelSelection24,omitempty"`
	AutoChannelSelection50                     *common.AutoChannelSelection           `json:"autoChannelSelection50,omitempty"`
	AwsVenue                                   *string                                `json:"awsVenue,omitempty"`
	BonjourGateway                             *common.GenericRef                     `json:"bonjourGateway,omitempty"`
	ChannelEvaluationInterval                  *int                                   `json:"channelEvaluationInterval,omitempty"`
	ClientAdmissionControl24                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`
	ClientAdmissionControl50                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`
	Description                                *string                                `json:"description,omitempty"`
	DirectedMulticastFromNetworkEnabled        *bool                                  `json:"directedMulticastFromNetworkEnabled,omitempty"`
	DirectedMulticastFromWiredClientEnabled    *bool                                  `json:"directedMulticastFromWiredClientEnabled,omitempty"`
	DirectedMulticastFromWirelessClientEnabled *bool                                  `json:"directedMulticastFromWirelessClientEnabled,omitempty"`
	GpsSource                                  *string                                `json:"gpsSource,omitempty"`
	Latitude                                   *float64                               `json:"latitude,omitempty"`
	Location                                   *string                                `json:"location,omitempty"`
	LocationAdditionalInfo                     *string                                `json:"locationAdditionalInfo,omitempty"`
	Login                                      *Login                                 `json:"login,omitempty"`
	Longitude                                  *float64                               `json:"longitude,omitempty"`
	Mac                                        *string                                `json:"mac,omitempty"`
	MeshOptions                                *Mesh                                  `json:"meshOptions,omitempty"`
	Model                                      *string                                `json:"model,omitempty"`
	Name                                       *string                                `json:"name,omitempty"`
	Network                                    *Network                               `json:"network,omitempty"`
	NetworkIpv6                                *NetworkIpv6                           `json:"networkIpv6,omitempty"`
	ProtectionMode24                           *string                                `json:"protectionMode24,omitempty"`
	ProvisionChecklist                         *string                                `json:"provisionChecklist,omitempty"`
	RecoverySsid                               *common.RecoverySsid                   `json:"recoverySsid,omitempty"`
	RogueApAggressivenessMode                  *int                                   `json:"rogueApAggressivenessMode,omitempty"`
	RogueApJammingThreshold                    *int                                   `json:"rogueApJammingThreshold,omitempty"`
	RogueApReportThreshold                     *int                                   `json:"rogueApReportThreshold,omitempty"`
	Serial                                     *string                                `json:"serial,omitempty"`
	SmartMonitor                               *common.OverrideSmartMonitor           `json:"smartMonitor,omitempty"`
	Specific                                   *apmodel.ApModel                       `json:"specific,omitempty"`
	Syslog                                     *Syslog                                `json:"syslog,omitempty"`
	VenueProfile                               *common.GenericRef                     `json:"venueProfile,omitempty"`
	Wifi24                                     *common.Radio24SuperSet                `json:"wifi24,omitempty"`
	Wifi50                                     *common.ApRadio50                      `json:"wifi50,omitempty"`
	WLANGroup24                                *WLANGroup                             `json:"wlanGroup24,omitempty"`
	WLANGroup50                                *WLANGroup                             `json:"wlanGroup50,omitempty"`
	WLANService24Enabled                       *bool                                  `json:"wlanService24Enabled,omitempty"`
	WLANService50Enabled                       *bool                                  `json:"wlanService50Enabled,omitempty"`
	ZoneID                                     *string                                `json:"zoneId,omitempty"`
}

type ApLinemanSummary struct {
	FirstIndex *int                        `json:"firstIndex,omitempty"`
	HasMore    *bool                       `json:"hasMore,omitempty"`
	List       []*ApLinemanSummaryListType `json:"list,omitempty"`
	TotalCount *int                        `json:"totalCount,omitempty"`
}

type ApLinemanSummaryListType struct {
	Alarms      *AlarmSummary `json:"alarms,omitempty"`
	ConfigState *string       `json:"configState,omitempty"`
	Latitude    *float64      `json:"latitude,omitempty"`
	Location    *string       `json:"location,omitempty"`
	Longitude   *float64      `json:"longitude,omitempty"`
	Mac         *string       `json:"mac,omitempty"`
	Name        *string       `json:"name,omitempty"`
}

type ApListEntry struct {
	FirstIndex *int                   `json:"firstIndex,omitempty"`
	HasMore    *bool                  `json:"hasMore,omitempty"`
	List       []*ApListEntryListType `json:"list,omitempty"`
	TotalCount *int                   `json:"totalCount,omitempty"`
}

type ApListEntryListType struct {
	ApGroupID *string `json:"apGroupId,omitempty"`
	Mac       *string `json:"mac,omitempty"`
	Name      *string `json:"name,omitempty"`
	Serial    *string `json:"serial,omitempty"`
	ZoneID    *string `json:"zoneId,omitempty"`
}

type ApOperationalSummary struct {
	AdministrativeState    *string          `json:"administrativeState,omitempty"`
	Altitude               *common.Altitude `json:"altitude,omitempty"`
	ApGroupID              *string          `json:"apGroupId,omitempty"`
	ApprovedTime           *int             `json:"approvedTime,omitempty"`
	ClientCount            *int             `json:"clientCount,omitempty"`
	ConfigState            *string          `json:"configState,omitempty"`
	ConnectionState        *string          `json:"connectionState,omitempty"`
	CountryCode            *string          `json:"countryCode,omitempty"`
	CpID                   *string          `json:"cpId,omitempty"`
	Description            *string          `json:"description,omitempty"`
	DpID                   *string          `json:"dpId,omitempty"`
	ExternalIP             *string          `json:"externalIp,omitempty"`
	ExternalPort           *int             `json:"externalPort,omitempty"`
	IP                     *string          `json:"ip,omitempty"`
	IPType                 *string          `json:"ipType,omitempty"`
	IsCriticalAP           *bool            `json:"isCriticalAP,omitempty"`
	LastSeenTime           *int             `json:"lastSeenTime,omitempty"`
	Latitude               *float64         `json:"latitude,omitempty"`
	Location               *string          `json:"location,omitempty"`
	LocationAdditionalInfo *string          `json:"locationAdditionalInfo,omitempty"`
	Longitude              *float64         `json:"longitude,omitempty"`
	Mac                    *string          `json:"mac,omitempty"`
	ManagementVlan         *int             `json:"managementVlan,omitempty"`
	MeshHop                *int             `json:"meshHop,omitempty"`
	MeshRole               *string          `json:"meshRole,omitempty"`
	Model                  *string          `json:"model,omitempty"`
	Name                   *string          `json:"name,omitempty"`
	ProvisionMethod        *string          `json:"provisionMethod,omitempty"`
	ProvisionStage         *string          `json:"provisionStage,omitempty"`
	RegistrationState      *string          `json:"registrationState,omitempty"`
	Serial                 *string          `json:"serial,omitempty"`
	Uptime                 *int             `json:"uptime,omitempty"`
	Version                *string          `json:"version,omitempty"`
	Wifi24Channel          *string          `json:"wifi24Channel,omitempty"`
	Wifi50Channel          *string          `json:"wifi50Channel,omitempty"`
	ZoneID                 *string          `json:"zoneId,omitempty"`
}

type ClientList struct {
	FirstIndex *int             `json:"firstIndex,omitempty"`
	HasMore    *bool            `json:"hasMore,omitempty"`
	List       []*common.Client `json:"list,omitempty"`
	TotalCount *int             `json:"totalCount,omitempty"`
}

type CreateAP struct {
	AdministrativeState *string  `json:"administrativeState,omitempty"`
	ApGroupID           *string  `json:"apGroupId,omitempty"`
	AwsVenue            *string  `json:"awsVenue,omitempty"`
	Description         *string  `json:"description,omitempty"`
	GpsSource           *string  `json:"gpsSource,omitempty"`
	Latitude            *float64 `json:"latitude,omitempty"`
	Location            *string  `json:"location,omitempty"`
	Longitude           *float64 `json:"longitude,omitempty"`
	Mac                 *string  `json:"mac,omitempty"`
	Model               *string  `json:"model,omitempty"`
	Name                *string  `json:"name,omitempty"`
	ProvisionChecklist  *string  `json:"provisionChecklist,omitempty"`
	Serial              *string  `json:"serial,omitempty"`
	ZoneID              *string  `json:"zoneId,omitempty"`
}

type EventSummary struct {
	CriticalCount      *int `json:"criticalCount,omitempty"`
	DebugCount         *int `json:"debugCount,omitempty"`
	InformationalCount *int `json:"informationalCount,omitempty"`
	MajorCount         *int `json:"majorCount,omitempty"`
	MinorCount         *int `json:"minorCount,omitempty"`
	WarningCount       *int `json:"warningCount,omitempty"`
}

type Login struct {
	ApLoginName     *string `json:"apLoginName,omitempty"`
	ApLoginPassword *string `json:"apLoginPassword,omitempty"`
}

type Mesh struct {
	MeshMode            *string  `json:"meshMode,omitempty"`
	MeshUplinkEntryList []string `json:"meshUplinkEntryList,omitempty"`
	UplinkSelection     *string  `json:"uplinkSelection,omitempty"`
}

type ModifyAP struct {
	AdministrativeState                        *string                                `json:"administrativeState,omitempty"`
	Altitude                                   *common.Altitude                       `json:"altitude,omitempty"`
	ApGroupID                                  *string                                `json:"apGroupId,omitempty"`
	ApMgmtVlan                                 *common.ApManagementVlan               `json:"apMgmtVlan,omitempty"`
	AutoChannelSelection24                     *common.AutoChannelSelection           `json:"autoChannelSelection24,omitempty"`
	AutoChannelSelection50                     *common.AutoChannelSelection           `json:"autoChannelSelection50,omitempty"`
	AwsVenue                                   *string                                `json:"awsVenue,omitempty"`
	BonjourGateway                             *common.GenericRef                     `json:"bonjourGateway,omitempty"`
	ChannelEvaluationInterval                  *int                                   `json:"channelEvaluationInterval,omitempty"`
	ClientAdmissionControl24                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`
	ClientAdmissionControl50                   *common.OverrideClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`
	Description                                *string                                `json:"description,omitempty"`
	DirectedMulticastFromNetworkEnabled        *bool                                  `json:"directedMulticastFromNetworkEnabled,omitempty"`
	DirectedMulticastFromWiredClientEnabled    *bool                                  `json:"directedMulticastFromWiredClientEnabled,omitempty"`
	DirectedMulticastFromWirelessClientEnabled *bool                                  `json:"directedMulticastFromWirelessClientEnabled,omitempty"`
	GpsSource                                  *string                                `json:"gpsSource,omitempty"`
	Latitude                                   *float64                               `json:"latitude,omitempty"`
	Location                                   *string                                `json:"location,omitempty"`
	LocationAdditionalInfo                     *string                                `json:"locationAdditionalInfo,omitempty"`
	Login                                      *Login                                 `json:"login,omitempty"`
	Longitude                                  *float64                               `json:"longitude,omitempty"`
	MeshOptions                                *Mesh                                  `json:"meshOptions,omitempty"`
	Model                                      *string                                `json:"model,omitempty"`
	Name                                       *string                                `json:"name,omitempty"`
	Network                                    *Network                               `json:"network,omitempty"`
	NetworkIpv6                                *NetworkIpv6                           `json:"networkIpv6,omitempty"`
	ProtectionMode24                           *string                                `json:"protectionMode24,omitempty"`
	ProvisionChecklist                         *string                                `json:"provisionChecklist,omitempty"`
	RecoverySsid                               *common.RecoverySsid                   `json:"recoverySsid,omitempty"`
	RogueApAggressivenessMode                  *int                                   `json:"rogueApAggressivenessMode,omitempty"`
	RogueApJammingThreshold                    *int                                   `json:"rogueApJammingThreshold,omitempty"`
	RogueApReportThreshold                     *int                                   `json:"rogueApReportThreshold,omitempty"`
	Serial                                     *string                                `json:"serial,omitempty"`
	SmartMonitor                               *common.OverrideSmartMonitor           `json:"smartMonitor,omitempty"`
	Syslog                                     *Syslog                                `json:"syslog,omitempty"`
	VenueProfile                               *common.GenericRef                     `json:"venueProfile,omitempty"`
	Wifi24                                     *common.Radio24                        `json:"wifi24,omitempty"`
	Wifi50                                     *common.ApRadio50                      `json:"wifi50,omitempty"`
	WLANGroup24                                *WLANGroup                             `json:"wlanGroup24,omitempty"`
	WLANGroup50                                *WLANGroup                             `json:"wlanGroup50,omitempty"`
	WLANService24Enabled                       *bool                                  `json:"wlanService24Enabled,omitempty"`
	WLANService50Enabled                       *bool                                  `json:"wlanService50Enabled,omitempty"`
	ZoneID                                     *string                                `json:"zoneId,omitempty"`
}

type ModifyRogueType struct {
	RogueMacList []string `json:"rogueMacList,omitempty"`
}

type NeighborAPList struct {
	FirstIndex *int                  `json:"firstIndex,omitempty"`
	HasMore    *bool                 `json:"hasMore,omitempty"`
	List       []*NeighborAPListType `json:"list,omitempty"`
	TotalCount *int                  `json:"totalCount,omitempty"`
}

type NeighborAPListType struct {
	Channel         *string `json:"channel,omitempty"`
	ConnectionState *string `json:"connectionState,omitempty"`
	ExternalIP      *string `json:"externalIp,omitempty"`
	ExternalPort    *string `json:"externalPort,omitempty"`
	IP              *string `json:"ip,omitempty"`
	Mac             *string `json:"mac,omitempty"`
	Model           *string `json:"model,omitempty"`
	Name            *string `json:"name,omitempty"`
	Signal          *string `json:"signal,omitempty"`
	Version         *string `json:"version,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
}

type Network struct {
	Gateway      *string `json:"gateway,omitempty"`
	IP           *string `json:"ip,omitempty"`
	IPType       *string `json:"ipType,omitempty"`
	Netmask      *string `json:"netmask,omitempty"`
	PrimaryDNS   *string `json:"primaryDns,omitempty"`
	SecondaryDNS *string `json:"secondaryDns,omitempty"`
}

type NetworkIpv6 struct {
	Gateway      *string `json:"gateway,omitempty"`
	IP           *string `json:"ip,omitempty"`
	IPType       *string `json:"ipType,omitempty"`
	PrimaryDNS   *string `json:"primaryDns,omitempty"`
	SecondaryDNS *string `json:"secondaryDns,omitempty"`
}

type SwitchoverAP struct {
	ApMacList    []string `json:"apMacList,omitempty"`
	ClusterName  *string  `json:"clusterName,omitempty"`
	DeleteRecord *bool    `json:"deleteRecord,omitempty"`
	IPOrFqdn     *string  `json:"ipOrFqdn,omitempty"`
	ZoneIDList   []string `json:"zoneIdList,omitempty"`
}

type Syslog struct {
	Address           *string `json:"address,omitempty"`
	Enabled           *bool   `json:"enabled,omitempty"`
	Facility          *string `json:"facility,omitempty"`
	FlowLevel         *string `json:"flowLevel,omitempty"`
	Port              *int    `json:"port,omitempty"`
	Priority          *string `json:"priority,omitempty"`
	Protocol          *string `json:"protocol,omitempty"`
	SecondaryAddress  *string `json:"secondaryAddress,omitempty"`
	SecondaryPort     *int    `json:"secondaryPort,omitempty"`
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty"`
}

type WLANGroup struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
