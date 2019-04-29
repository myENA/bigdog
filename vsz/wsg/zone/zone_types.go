package zone

// API Version: v8_0

type ApFirmware struct {
	FirmwareVersion           *string                      `json:"firmwareVersion,omitempty"`
	Supported                 *bool                        `json:"supported,omitempty"`
	UnsupportedApModelSummary []*UnsupportedApModelSummary `json:"unsupportedApModelSummary,omitempty"`
}

type ApFirmwareList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type ApLogin struct {
	ApLoginName     *string `json:"apLoginName,omitempty"`
	ApLoginPassword *string `json:"apLoginPassword,omitempty"`
}

type ApSNMPOptions struct {
	ApSNMPEnabled *bool          `json:"apSnmpEnabled,omitempty"`
	SNMPV2Agent   []*SNMPV2Agent `json:"snmpV2Agent,omitempty"`
	SNMPV3Agent   []*SNMPV3Agent `json:"snmpV3Agent,omitempty"`
}

type AvailableTunnelProfile struct {
	AaaAffinityEnabled *bool   `json:"aaaAffinityEnabled,omitempty"`
	ID                 *string `json:"id,omitempty"`
	IPMode             *string `json:"ipMode,omitempty"`
	Name               *string `json:"name,omitempty"`
	TunnelType         *string `json:"tunnelType,omitempty"`
}

type AvailableTunnelProfileList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type BackgroundScanning struct {
	FrequencyInSec *int `json:"frequencyInSec,omitempty"`
}

type BandBalancing struct {
	Mode             *string `json:"mode,omitempty"`
	Wifi24Percentage *int    `json:"wifi24Percentage,omitempty"`
}

type BonjourGatewayPolicyConfiguration struct {
	BonjourPolicyRuleList []*BonjourPolicyRuleList `json:"bonjourPolicyRuleList,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	Name                  *string                  `json:"name,omitempty"`
}

type BonjourGatewayPolicyList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type BonjourGatewayPolicySummary struct {
	Description    *string `json:"description,omitempty"`
	ID             *string `json:"id,omitempty"`
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	LastModifiedOn *string `json:"lastModifiedOn,omitempty"`
	Name           *string `json:"name,omitempty"`
}

type BonjourPolicyRule struct {
	BridgeService *string `json:"bridgeService,omitempty"`
	FromVlan      *int    `json:"fromVlan,omitempty"`
	Notes         *string `json:"notes,omitempty"`
	Protocol      *string `json:"protocol,omitempty"`
	ToVlan        *int    `json:"toVlan,omitempty"`
}

type BonjourPolicyRuleConfiguration struct {
	BridgeService *string `json:"bridgeService,omitempty"`
	FromVlan      *int    `json:"fromVlan,omitempty"`
	Notes         *string `json:"notes,omitempty"`
	Priority      *string `json:"priority,omitempty"`
	Protocol      *string `json:"protocol,omitempty"`
	ToVlan        *int    `json:"toVlan,omitempty"`
}

type ClientLoadBalancing struct {
	AdjacentRadioThreshold *int `json:"adjacentRadioThreshold,omitempty"`
}

type CreateBonjourGatewayPolicy struct {
	BonjourPolicyRuleList []*BonjourPolicyRuleList `json:"bonjourPolicyRuleList,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	Name                  *string                  `json:"name,omitempty"`
}

type CreateDiffServProfile struct {
	Description      *string           `json:"description,omitempty"`
	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`
	Name             *string           `json:"name,omitempty"`
	PreservedList    []string          `json:"preservedList,omitempty"`
	UplinkDiffServ   *UplinkDiffServ   `json:"uplinkDiffServ,omitempty"`
}

type CreateZone struct {
	Altitude                                   *common.Altitude               `json:"altitude,omitempty"`
	ApLatencyInterval                          *common.ApLatencyInterval      `json:"apLatencyInterval,omitempty"`
	ApMgmtVlan                                 *common.ApManagementVlan       `json:"apMgmtVlan,omitempty"`
	ApRebootTimeout                            *common.ApRebootTimeout        `json:"apRebootTimeout,omitempty"`
	AutoChannelSelection24                     *common.AutoChannelSelection   `json:"autoChannelSelection24,omitempty"`
	AutoChannelSelection50                     *common.AutoChannelSelection   `json:"autoChannelSelection50,omitempty"`
	BackgroundScanning24                       *BackgroundScanning24          `json:"backgroundScanning24,omitempty"`
	BackgroundScanning50                       *BackgroundScanning50          `json:"backgroundScanning50,omitempty"`
	BandBalancing                              *BandBalancing                 `json:"bandBalancing,omitempty"`
	BonjourFencingPolicy                       *common.GenericRef             `json:"bonjourFencingPolicy,omitempty"`
	BonjourFencingPolicyEnabled                *bool                          `json:"bonjourFencingPolicyEnabled,omitempty"`
	CbandChannelEnabled                        *bool                          `json:"cbandChannelEnabled,omitempty"`
	CbandChannelLicenseEnabled                 *bool                          `json:"cbandChannelLicenseEnabled,omitempty"`
	ChannelEvaluationInterval                  *int                           `json:"channelEvaluationInterval,omitempty"`
	ChannelModeEnabled                         *bool                          `json:"channelModeEnabled,omitempty"`
	ClientAdmissionControl24                   *common.ClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`
	ClientAdmissionControl50                   *common.ClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`
	ClientLoadBalancing24                      *ClientLoadBalancing24         `json:"clientLoadBalancing24,omitempty"`
	ClientLoadBalancing50                      *ClientLoadBalancing50         `json:"clientLoadBalancing50,omitempty"`
	ClusterRedundancyEnabled                   *bool                          `json:"clusterRedundancyEnabled,omitempty"`
	CountryCode                                *string                        `json:"countryCode,omitempty"`
	Description                                *string                        `json:"description,omitempty"`
	DfsChannelEnabled                          *bool                          `json:"dfsChannelEnabled,omitempty"`
	DHCPSiteConfig                             *common.DHCPSiteConfigRef      `json:"dhcpSiteConfig,omitempty"`
	DirectedMulticastFromNetworkEnabled        *bool                          `json:"directedMulticastFromNetworkEnabled,omitempty"`
	DirectedMulticastFromWiredClientEnabled    *bool                          `json:"directedMulticastFromWiredClientEnabled,omitempty"`
	DirectedMulticastFromWirelessClientEnabled *bool                          `json:"directedMulticastFromWirelessClientEnabled,omitempty"`
	DomainID                                   *string                        `json:"domainId,omitempty"`
	DosBarringCheckPeriod                      *int                           `json:"dosBarringCheckPeriod,omitempty"`
	DosBarringEnable                           *int                           `json:"dosBarringEnable,omitempty"`
	DosBarringPeriod                           *int                           `json:"dosBarringPeriod,omitempty"`
	DosBarringThreshold                        *int                           `json:"dosBarringThreshold,omitempty"`
	HealthCheckSites                           []string                       `json:"healthCheckSites,omitempty"`
	HealthCheckSitesEnabled                    *bool                          `json:"healthCheckSitesEnabled,omitempty"`
	IpsecProfile                               *common.GenericRef             `json:"ipsecProfile,omitempty"`
	IpsecProfiles                              []*IpsecProfiles               `json:"ipsecProfiles,omitempty"`
	Ipv6TrafficFilterEnabled                   *int                           `json:"ipv6TrafficFilterEnabled,omitempty"`
	Latitude                                   *float64                       `json:"latitude,omitempty"`
	LoadBalancingMethod                        *string                        `json:"loadBalancingMethod,omitempty"`
	Location                                   *string                        `json:"location,omitempty"`
	LocationAdditionalInfo                     *string                        `json:"locationAdditionalInfo,omitempty"`
	LocationBasedService                       *common.GenericRef             `json:"locationBasedService,omitempty"`
	Login                                      *Login                         `json:"login,omitempty"`
	Longitude                                  *float64                       `json:"longitude,omitempty"`
	Mesh                                       *Mesh                          `json:"mesh,omitempty"`
	Name                                       *string                        `json:"name,omitempty"`
	NodeAffinityProfile                        *common.GenericRef             `json:"nodeAffinityProfile,omitempty"`
	ProtectionMode24                           *string                        `json:"protectionMode24,omitempty"`
	RecoverySsid                               *common.RecoverySsid           `json:"recoverySsid,omitempty"`
	Rogue                                      *Rogue                         `json:"rogue,omitempty"`
	RogueApAggressivenessMode                  *int                           `json:"rogueApAggressivenessMode,omitempty"`
	RogueApJammingDetection                    *bool                          `json:"rogueApJammingDetection,omitempty"`
	RogueApJammingThreshold                    *int                           `json:"rogueApJammingThreshold,omitempty"`
	RogueApReportThreshold                     *int                           `json:"rogueApReportThreshold,omitempty"`
	RuckusGreTunnelProfile                     *common.GenericRef             `json:"ruckusGreTunnelProfile,omitempty"`
	SmartMonitor                               *common.SmartMonitor           `json:"smartMonitor,omitempty"`
	SNMPAgent                                  *SNMPAgent                     `json:"snmpAgent,omitempty"`
	SoftGreTunnelProflies                      []*SoftGreTunnelProflies       `json:"softGreTunnelProflies,omitempty"`
	SSHTunnelEncryption                        *string                        `json:"sshTunnelEncryption,omitempty"`
	Syslog                                     *Syslog                        `json:"syslog,omitempty"`
	Timezone                                   *Timezone                      `json:"timezone,omitempty"`
	TunnelProfile                              *common.GenericRef             `json:"tunnelProfile,omitempty"`
	TunnelType                                 *string                        `json:"tunnelType,omitempty"`
	VenueProfile                               *common.GenericRef             `json:"venueProfile,omitempty"`
	Version                                    *string                        `json:"version,omitempty"`
	VlanOverlappingEnabled                     *bool                          `json:"vlanOverlappingEnabled,omitempty"`
	Wifi24                                     *common.Radio24                `json:"wifi24,omitempty"`
	Wifi50                                     *common.Radio50                `json:"wifi50,omitempty"`
}

type CustomizedTimeZone struct {
	Abbreviation    *string `json:"abbreviation,omitempty"`
	End             *End    `json:"end,omitempty"`
	GmtOffset       *int    `json:"gmtOffset,omitempty"`
	GmtOffsetMinute *int    `json:"gmtOffsetMinute,omitempty"`
	Start           *Start  `json:"start,omitempty"`
}

type DaylightSavingTime struct {
	Day   *int `json:"day,omitempty"`
	Hour  *int `json:"hour,omitempty"`
	Month *int `json:"month,omitempty"`
	Week  *int `json:"week,omitempty"`
}

type DHCPSiteConfigList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type DiffServConfiguration struct {
	Description      *string           `json:"description,omitempty"`
	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`
	ID               *string           `json:"id,omitempty"`
	Name             *string           `json:"name,omitempty"`
	PreservedList    []string          `json:"preservedList,omitempty"`
	UplinkDiffServ   *UplinkDiffServ   `json:"uplinkDiffServ,omitempty"`
}

type DiffServList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type DiffServSummary struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type DownlinkDiffServ struct {
	Downlink       *string `json:"downlink,omitempty"`
	DownlinkEnable *bool   `json:"downlinkEnable,omitempty"`
}

type MeshConfiguration struct {
	MeshRadioIdx    *string `json:"meshRadioIdx,omitempty"`
	Passphrase      *string `json:"passphrase,omitempty"`
	Ssid            *string `json:"ssid,omitempty"`
	ZeroTouchStatus *bool   `json:"zeroTouchStatus,omitempty"`
}

type ModfiyApFirmware struct {
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
}

type ModifyBonjourGatewayPolicy struct {
	BonjourPolicyRuleList []*BonjourPolicyRuleList `json:"bonjourPolicyRuleList,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	Name                  *string                  `json:"name,omitempty"`
}

type ModifyDiffServProfile struct {
	Description      *string           `json:"description,omitempty"`
	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`
	Name             *string           `json:"name,omitempty"`
	PreservedList    []string          `json:"preservedList,omitempty"`
	UplinkDiffServ   *UplinkDiffServ   `json:"uplinkDiffServ,omitempty"`
}

type ModifyZone struct {
	AaaAffinityEnabled                         *bool                          `json:"aaaAffinityEnabled,omitempty"`
	Altitude                                   *common.Altitude               `json:"altitude,omitempty"`
	ApLatencyInterval                          *common.ApLatencyInterval      `json:"apLatencyInterval,omitempty"`
	ApMgmtVlan                                 *common.ApManagementVlan       `json:"apMgmtVlan,omitempty"`
	ApRebootTimeout                            *common.ApRebootTimeout        `json:"apRebootTimeout,omitempty"`
	AutoChannelSelection24                     *common.AutoChannelSelection   `json:"autoChannelSelection24,omitempty"`
	AutoChannelSelection50                     *common.AutoChannelSelection   `json:"autoChannelSelection50,omitempty"`
	AwsVenue                                   *string                        `json:"awsVenue,omitempty"`
	BackgroundScanning24                       *BackgroundScanning24          `json:"backgroundScanning24,omitempty"`
	BackgroundScanning50                       *BackgroundScanning50          `json:"backgroundScanning50,omitempty"`
	BandBalancing                              *BandBalancing                 `json:"bandBalancing,omitempty"`
	BonjourFencingPolicy                       *common.GenericRef             `json:"bonjourFencingPolicy,omitempty"`
	BonjourFencingPolicyEnabled                *bool                          `json:"bonjourFencingPolicyEnabled,omitempty"`
	CbandChannelEnabled                        *bool                          `json:"cbandChannelEnabled,omitempty"`
	CbandChannelLicenseEnabled                 *bool                          `json:"cbandChannelLicenseEnabled,omitempty"`
	ChannelEvaluationInterval                  *int                           `json:"channelEvaluationInterval,omitempty"`
	ChannelModeEnabled                         *bool                          `json:"channelModeEnabled,omitempty"`
	ClientAdmissionControl24                   *common.ClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`
	ClientAdmissionControl50                   *common.ClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`
	ClientLoadBalancing24                      *ClientLoadBalancing24         `json:"clientLoadBalancing24,omitempty"`
	ClientLoadBalancing50                      *ClientLoadBalancing50         `json:"clientLoadBalancing50,omitempty"`
	ClusterRedundancyEnabled                   *bool                          `json:"clusterRedundancyEnabled,omitempty"`
	CountryCode                                *string                        `json:"countryCode,omitempty"`
	Description                                *string                        `json:"description,omitempty"`
	DfsChannelEnabled                          *bool                          `json:"dfsChannelEnabled,omitempty"`
	DHCPSiteConfig                             *common.DHCPSiteConfigRef      `json:"dhcpSiteConfig,omitempty"`
	DirectedMulticastFromNetworkEnabled        *bool                          `json:"directedMulticastFromNetworkEnabled,omitempty"`
	DirectedMulticastFromWiredClientEnabled    *bool                          `json:"directedMulticastFromWiredClientEnabled,omitempty"`
	DirectedMulticastFromWirelessClientEnabled *bool                          `json:"directedMulticastFromWirelessClientEnabled,omitempty"`
	DomainID                                   *string                        `json:"domainId,omitempty"`
	DosBarringCheckPeriod                      *int                           `json:"dosBarringCheckPeriod,omitempty"`
	DosBarringEnable                           *int                           `json:"dosBarringEnable,omitempty"`
	DosBarringPeriod                           *int                           `json:"dosBarringPeriod,omitempty"`
	DosBarringThreshold                        *int                           `json:"dosBarringThreshold,omitempty"`
	HealthCheckSites                           []string                       `json:"healthCheckSites,omitempty"`
	HealthCheckSitesEnabled                    *bool                          `json:"healthCheckSitesEnabled,omitempty"`
	IpsecProfile                               *common.GenericRef             `json:"ipsecProfile,omitempty"`
	IpsecProfiles                              []*IpsecProfiles               `json:"ipsecProfiles,omitempty"`
	Ipv6TrafficFilterEnabled                   *int                           `json:"ipv6TrafficFilterEnabled,omitempty"`
	Latitude                                   *float64                       `json:"latitude,omitempty"`
	LoadBalancingMethod                        *string                        `json:"loadBalancingMethod,omitempty"`
	Location                                   *string                        `json:"location,omitempty"`
	LocationAdditionalInfo                     *string                        `json:"locationAdditionalInfo,omitempty"`
	LocationBasedService                       *common.GenericRef             `json:"locationBasedService,omitempty"`
	Login                                      *Login                         `json:"login,omitempty"`
	Longitude                                  *float64                       `json:"longitude,omitempty"`
	Mesh                                       *Mesh                          `json:"mesh,omitempty"`
	Name                                       *string                        `json:"name,omitempty"`
	NodeAffinityProfile                        *common.GenericRef             `json:"nodeAffinityProfile,omitempty"`
	ProtectionMode24                           *string                        `json:"protectionMode24,omitempty"`
	RecoverySsid                               *common.RecoverySsid           `json:"recoverySsid,omitempty"`
	Rogue                                      *Rogue                         `json:"rogue,omitempty"`
	RogueApAggressivenessMode                  *int                           `json:"rogueApAggressivenessMode,omitempty"`
	RogueApJammingDetection                    *bool                          `json:"rogueApJammingDetection,omitempty"`
	RogueApJammingThreshold                    *int                           `json:"rogueApJammingThreshold,omitempty"`
	RogueApReportThreshold                     *int                           `json:"rogueApReportThreshold,omitempty"`
	RuckusGreTunnelProfile                     *common.GenericRef             `json:"ruckusGreTunnelProfile,omitempty"`
	SmartMonitor                               *common.SmartMonitor           `json:"smartMonitor,omitempty"`
	SNMPAgent                                  *SNMPAgent                     `json:"snmpAgent,omitempty"`
	SoftGreTunnelProflies                      []*SoftGreTunnelProflies       `json:"softGreTunnelProflies,omitempty"`
	SSHTunnelEncryption                        *string                        `json:"sshTunnelEncryption,omitempty"`
	Syslog                                     *Syslog                        `json:"syslog,omitempty"`
	Timezone                                   *Timezone                      `json:"timezone,omitempty"`
	TunnelProfile                              *common.GenericRef             `json:"tunnelProfile,omitempty"`
	TunnelType                                 *string                        `json:"tunnelType,omitempty"`
	VenueProfile                               *common.GenericRef             `json:"venueProfile,omitempty"`
	VlanOverlappingEnabled                     *bool                          `json:"vlanOverlappingEnabled,omitempty"`
	Wifi24                                     *common.Radio24                `json:"wifi24,omitempty"`
	Wifi50                                     *common.Radio50                `json:"wifi50,omitempty"`
	ZoneAffinityProfileID                      *string                        `json:"zoneAffinityProfileId,omitempty"`
}

type QueryCriteria struct {
	Attributes      []string           `json:"attributes,omitempty"`
	Criteria        *string            `json:"criteria,omitempty"`
	ExpandDomains   *bool              `json:"expandDomains,omitempty"`
	ExtraFilters    []*ExtraFilters    `json:"extraFilters,omitempty"`
	ExtraNotFilters []*ExtraNotFilters `json:"extraNotFilters,omitempty"`
	ExtraTimeRange  *ExtraTimeRange    `json:"extraTimeRange,omitempty"`
	Filters         []*Filters         `json:"filters,omitempty"`
	FullTextSearch  *FullTextSearch    `json:"fullTextSearch,omitempty"`
	Limit           *int               `json:"limit,omitempty"`
	Options         *Options           `json:"options,omitempty"`
	Page            *int               `json:"page,omitempty"`
	Query           *string            `json:"query,omitempty"`
	SortInfo        *SortInfo          `json:"sortInfo,omitempty"`
}

type QueryCriteriaExtraFiltersType struct {
	Operator *string `json:"operator,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type QueryCriteriaFiltersType struct {
	Operator *string `json:"operator,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type QueryCriteriaOptionsType struct {
	IncludeSharedResources *bool   `json:"includeSharedResources,omitempty"`
	ZoneIpmode             *string `json:"zone_ipmode,omitempty"`
}

type Rogue struct {
	MaliciousTypes    []string `json:"maliciousTypes,omitempty"`
	ProtectionEnabled *bool    `json:"protectionEnabled,omitempty"`
	ReportType        *string  `json:"reportType,omitempty"`
}

type SNMPUser struct {
	AuthPassword        *string               `json:"authPassword,omitempty"`
	AuthProtocol        *string               `json:"authProtocol,omitempty"`
	NotificationEnabled *bool                 `json:"notificationEnabled,omitempty"`
	NotificationTarget  []*NotificationTarget `json:"notificationTarget,omitempty"`
	NotificationType    *string               `json:"notificationType,omitempty"`
	PrivPassword        *string               `json:"privPassword,omitempty"`
	PrivProtocol        *string               `json:"privProtocol,omitempty"`
	ReadEnabled         *bool                 `json:"readEnabled,omitempty"`
	UserName            *string               `json:"userName,omitempty"`
	WriteEnabled        *bool                 `json:"writeEnabled,omitempty"`
}

type SoftGreRef struct {
	AaaAffinityEnabled *bool   `json:"aaaAffinityEnabled,omitempty"`
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
}

type Syslog struct {
	Address           *string `json:"address,omitempty"`
	Facility          *string `json:"facility,omitempty"`
	FlowLevel         *string `json:"flowLevel,omitempty"`
	Port              *int    `json:"port,omitempty"`
	Priority          *string `json:"priority,omitempty"`
	Protocol          *string `json:"protocol,omitempty"`
	SecondaryAddress  *string `json:"secondaryAddress,omitempty"`
	SecondaryPort     *int    `json:"secondaryPort,omitempty"`
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty"`
}

type TimezoneSetting struct {
	CustomizedTimezone *CustomizedTimezone `json:"customizedTimezone,omitempty"`
	SystemTimezone     *string             `json:"systemTimezone,omitempty"`
}

type UnsupportedApModel struct {
	Amount *int    `json:"amount,omitempty"`
	Model  *string `json:"model,omitempty"`
}

type UplinkDiffServ struct {
	Uplink       *string `json:"uplink,omitempty"`
	UplinkEnable *bool   `json:"uplinkEnable,omitempty"`
}

type ZoneConfiguration struct {
	AaaAffinityEnabled                         *bool                          `json:"aaaAffinityEnabled,omitempty"`
	Altitude                                   *common.Altitude               `json:"altitude,omitempty"`
	ApLatencyInterval                          *common.ApLatencyInterval      `json:"apLatencyInterval,omitempty"`
	ApMgmtVlan                                 *common.ApManagementVlan       `json:"apMgmtVlan,omitempty"`
	ApRebootTimeout                            *common.ApRebootTimeout        `json:"apRebootTimeout,omitempty"`
	AutoChannelSelection24                     *common.AutoChannelSelection   `json:"autoChannelSelection24,omitempty"`
	AutoChannelSelection50                     *common.AutoChannelSelection   `json:"autoChannelSelection50,omitempty"`
	AwsVenue                                   *string                        `json:"awsVenue,omitempty"`
	BackgroundScanning24                       *BackgroundScanning24          `json:"backgroundScanning24,omitempty"`
	BackgroundScanning50                       *BackgroundScanning50          `json:"backgroundScanning50,omitempty"`
	BandBalancing                              *BandBalancing                 `json:"bandBalancing,omitempty"`
	BonjourFencingPolicy                       *common.GenericRef             `json:"bonjourFencingPolicy,omitempty"`
	BonjourFencingPolicyEnabled                *bool                          `json:"bonjourFencingPolicyEnabled,omitempty"`
	CbandChannelEnabled                        *bool                          `json:"cbandChannelEnabled,omitempty"`
	CbandChannelLicenseEnabled                 *bool                          `json:"cbandChannelLicenseEnabled,omitempty"`
	ChannelEvaluationInterval                  *int                           `json:"channelEvaluationInterval,omitempty"`
	ChannelModeEnabled                         *bool                          `json:"channelModeEnabled,omitempty"`
	ClientAdmissionControl24                   *common.ClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`
	ClientAdmissionControl50                   *common.ClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`
	ClientLoadBalancing24                      *ClientLoadBalancing24         `json:"clientLoadBalancing24,omitempty"`
	ClientLoadBalancing50                      *ClientLoadBalancing50         `json:"clientLoadBalancing50,omitempty"`
	ClusterRedundancyEnabled                   *bool                          `json:"clusterRedundancyEnabled,omitempty"`
	CountryCode                                *string                        `json:"countryCode,omitempty"`
	Description                                *string                        `json:"description,omitempty"`
	DfsChannelEnabled                          *bool                          `json:"dfsChannelEnabled,omitempty"`
	DHCPSiteConfig                             *common.DHCPSiteConfigRef      `json:"dhcpSiteConfig,omitempty"`
	DirectedMulticastFromNetworkEnabled        *bool                          `json:"directedMulticastFromNetworkEnabled,omitempty"`
	DirectedMulticastFromWiredClientEnabled    *bool                          `json:"directedMulticastFromWiredClientEnabled,omitempty"`
	DirectedMulticastFromWirelessClientEnabled *bool                          `json:"directedMulticastFromWirelessClientEnabled,omitempty"`
	DomainID                                   *string                        `json:"domainId,omitempty"`
	DosBarringCheckPeriod                      *int                           `json:"dosBarringCheckPeriod,omitempty"`
	DosBarringEnable                           *int                           `json:"dosBarringEnable,omitempty"`
	DosBarringPeriod                           *int                           `json:"dosBarringPeriod,omitempty"`
	DosBarringThreshold                        *int                           `json:"dosBarringThreshold,omitempty"`
	HealthCheckSites                           []string                       `json:"healthCheckSites,omitempty"`
	HealthCheckSitesEnabled                    *bool                          `json:"healthCheckSitesEnabled,omitempty"`
	ID                                         *string                        `json:"id,omitempty"`
	IPMode                                     *string                        `json:"ipMode,omitempty"`
	IpsecProfile                               *common.GenericRef             `json:"ipsecProfile,omitempty"`
	IpsecProfiles                              []*IpsecProfiles               `json:"ipsecProfiles,omitempty"`
	Ipv6TrafficFilterEnabled                   *int                           `json:"ipv6TrafficFilterEnabled,omitempty"`
	Latitude                                   *float64                       `json:"latitude,omitempty"`
	LoadBalancingMethod                        *string                        `json:"loadBalancingMethod,omitempty"`
	Location                                   *string                        `json:"location,omitempty"`
	LocationAdditionalInfo                     *string                        `json:"locationAdditionalInfo,omitempty"`
	LocationBasedService                       *common.GenericRef             `json:"locationBasedService,omitempty"`
	Login                                      *Login                         `json:"login,omitempty"`
	Longitude                                  *float64                       `json:"longitude,omitempty"`
	Mesh                                       *Mesh                          `json:"mesh,omitempty"`
	Name                                       *string                        `json:"name,omitempty"`
	NodeAffinityProfile                        *common.GenericRef             `json:"nodeAffinityProfile,omitempty"`
	ProtectionMode24                           *string                        `json:"protectionMode24,omitempty"`
	RecoverySsid                               *common.RecoverySsid           `json:"recoverySsid,omitempty"`
	Rogue                                      *Rogue                         `json:"rogue,omitempty"`
	RogueApAggressivenessMode                  *int                           `json:"rogueApAggressivenessMode,omitempty"`
	RogueApJammingDetection                    *bool                          `json:"rogueApJammingDetection,omitempty"`
	RogueApJammingThreshold                    *int                           `json:"rogueApJammingThreshold,omitempty"`
	RogueApReportThreshold                     *int                           `json:"rogueApReportThreshold,omitempty"`
	RuckusGreTunnelProfile                     *common.GenericRef             `json:"ruckusGreTunnelProfile,omitempty"`
	SmartMonitor                               *common.SmartMonitor           `json:"smartMonitor,omitempty"`
	SNMPAgent                                  *SNMPAgent                     `json:"snmpAgent,omitempty"`
	SoftGreTunnelProflies                      []*SoftGreTunnelProflies       `json:"softGreTunnelProflies,omitempty"`
	SSHTunnelEncryption                        *string                        `json:"sshTunnelEncryption,omitempty"`
	Syslog                                     *Syslog                        `json:"syslog,omitempty"`
	Timezone                                   *Timezone                      `json:"timezone,omitempty"`
	TunnelProfile                              *common.GenericRef             `json:"tunnelProfile,omitempty"`
	TunnelType                                 *string                        `json:"tunnelType,omitempty"`
	VenueProfile                               *common.GenericRef             `json:"venueProfile,omitempty"`
	Version                                    *string                        `json:"version,omitempty"`
	VlanOverlappingEnabled                     *bool                          `json:"vlanOverlappingEnabled,omitempty"`
	Wifi24                                     *common.Radio24SuperSet        `json:"wifi24,omitempty"`
	Wifi50                                     *common.Radio50SuperSet        `json:"wifi50,omitempty"`
	ZoneAffinityProfileID                      *string                        `json:"zoneAffinityProfileId,omitempty"`
}

type ZoneList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type ZoneSummary struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
}
