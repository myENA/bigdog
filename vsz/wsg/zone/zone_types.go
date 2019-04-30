package zone

// API Version: v8_0

type ApFirmware struct {
	// FirmwareVersion
	// version of the AP firmare
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Supported
	// version of the AP firmare is supported for Upgrade or Downgrade.
	Supported *bool `json:"supported,omitempty"`

	// UnsupportedApModelSummary
	// summary of the AP Model is unsupported for AP firmware version.
	UnsupportedApModelSummary []*UnsupportedApModel `json:"unsupportedApModelSummary,omitempty"`
}

type ApFirmwareList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApFirmware `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApLogin struct {
	ApLoginName *string `json:"apLoginName,omitempty"`

	ApLoginPassword *string `json:"apLoginPassword,omitempty"`
}

type ApSNMPOptions struct {
	// ApSNMPEnabled
	// Enable AP SNMP
	ApSNMPEnabled *bool `json:"apSnmpEnabled,omitempty"`

	// SNMPV2Agent
	// Community List of the SNMP V2 Agent.
	SNMPV2Agent []*common.SNMPCommunity `json:"snmpV2Agent,omitempty"`

	// SNMPV3Agent
	// User List of the SNMP V3 Agent.
	SNMPV3Agent []*SNMPUser `json:"snmpV3Agent,omitempty"`
}

type AvailableTunnelProfile struct {
	// AaaAffinityEnabled
	// Enable AAA affinity (Soft GRE only)
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	// ID
	// Tunnel Profile ID
	ID *string `json:"id,omitempty"`

	IPMode *string `json:"ipMode,omitempty"`

	// Name
	// Tunnel Profile Name
	Name *string `json:"name,omitempty"`

	// TunnelType
	// Tunnel Profile Type ("RuckusGRE", "SoftGRE",or "Ipsec")
	TunnelType *string `json:"tunnelType,omitempty"`
}

type AvailableTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AvailableTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type BackgroundScanning struct {
	// FrequencyInSec
	// Frequency in second
	FrequencyInSec *int `json:"frequencyInSec,omitempty"`
}

type BandBalancing struct {
	// Mode
	// Band Balancing Mode: BASIC-Withholds probe and authentication responses at connetcion time in heavily
	// loaded band to balance clients to the other band, PROACTIVE-Uses BASIC functionality and actively
	// rebalances clients via 802.11v BTM, STRICT-Uses PROACTIVE functionality and forcefully rebalances
	// clients via 802.11v BTM
	Mode *string `json:"mode,omitempty"`

	// Wifi24Percentage
	// Percentage of client load on 2.4GHz radio band
	Wifi24Percentage *int `json:"wifi24Percentage,omitempty"`
}

type BonjourGatewayPolicyConfiguration struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*BonjourPolicyRuleConfiguration `json:"bonjourPolicyRuleList,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

type BonjourGatewayPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BonjourGatewayPolicySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type BonjourGatewayPolicySummary struct {
	// Description
	// Description of the bonjour gateway policy
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the bonjour gateway policy
	ID *string `json:"id,omitempty"`

	// LastModifiedBy
	// Last modified user of the bonjour gateway policy
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedOn
	// Last modified time of the bonjour gateway policy
	LastModifiedOn *string `json:"lastModifiedOn,omitempty"`

	// Name
	// Name of the bonjour gateway policy
	Name *string `json:"name,omitempty"`
}

// BonjourPolicyRule
//
// Bonjour policy rule
type BonjourPolicyRule struct {
	// BridgeService
	// Bridge service
	BridgeService *string `json:"bridgeService,omitempty"`

	// FromVlan
	// From VLAN
	FromVlan *int `json:"fromVlan,omitempty"`

	// Notes
	// Notes
	Notes *string `json:"notes,omitempty"`

	// Protocol
	// protocol. This is only available when bridgeService is OTHER
	Protocol *string `json:"protocol,omitempty"`

	// ToVlan
	// To VLAN
	ToVlan *int `json:"toVlan,omitempty"`
}

// BonjourPolicyRuleConfiguration
//
// Bonjour policy rule
type BonjourPolicyRuleConfiguration struct {
	// BridgeService
	// Bridge service
	BridgeService *string `json:"bridgeService,omitempty"`

	// FromVlan
	// From VLAN
	FromVlan *int `json:"fromVlan,omitempty"`

	// Notes
	// Notes
	Notes *string `json:"notes,omitempty"`

	// Priority
	// Priority
	Priority *string `json:"priority,omitempty"`

	// Protocol
	// protocol
	Protocol *string `json:"protocol,omitempty"`

	// ToVlan
	// To VLAN
	ToVlan *int `json:"toVlan,omitempty"`
}

type ClientLoadBalancing struct {
	// AdjacentRadioThreshold
	// Adjacent radio threshold
	AdjacentRadioThreshold *int `json:"adjacentRadioThreshold,omitempty"`
}

type CreateBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*BonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

type CreateDiffServProfile struct {
	Description *string `json:"description,omitempty"`

	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	Name *string `json:"name,omitempty"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *UplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

type CreateZone struct {
	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApLatencyInterval *common.ApLatencyInterval `json:"apLatencyInterval,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	ApRebootTimeout *common.ApRebootTimeout `json:"apRebootTimeout,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	BackgroundScanning24 *BackgroundScanning `json:"backgroundScanning24,omitempty"`

	BackgroundScanning50 *BackgroundScanning `json:"backgroundScanning50,omitempty"`

	BandBalancing *BandBalancing `json:"bandBalancing,omitempty"`

	BonjourFencingPolicy *common.GenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	ClientAdmissionControl24 *common.ClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *common.ClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	ClientLoadBalancing24 *ClientLoadBalancing `json:"clientLoadBalancing24,omitempty"`

	ClientLoadBalancing50 *ClientLoadBalancing `json:"clientLoadBalancing50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone.
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	CountryCode *string `json:"countryCode,omitempty"`

	Description *string `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DHCPSiteConfig *common.DHCPSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainID
	// Identifier of the management domain to which the zone belongs
	DomainID *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	HealthCheckSites []string `json:"healthCheckSites,omitempty"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	IpsecProfiles []*common.GenericRef `json:"ipsecProfiles,omitempty"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *float64 `json:"latitude,omitempty"`

	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty"`

	Location *string `json:"location,omitempty"`

	LocationAdditionalInfo *string `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.GenericRef `json:"locationBasedService,omitempty"`

	Login *ApLogin `json:"login,omitempty"`

	Longitude *float64 `json:"longitude,omitempty"`

	Mesh *MeshConfiguration `json:"mesh,omitempty"`

	Name *string `json:"name,omitempty"`

	NodeAffinityProfile *common.GenericRef `json:"nodeAffinityProfile,omitempty"`

	ProtectionMode24 *string `json:"protectionMode24,omitempty"`

	RecoverySsid *common.RecoverySsid `json:"recoverySsid,omitempty"`

	Rogue *Rogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	RuckusGreTunnelProfile *common.GenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	SmartMonitor *common.SmartMonitor `json:"smartMonitor,omitempty"`

	SNMPAgent *ApSNMPOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*SoftGreRef `json:"softGreTunnelProflies,omitempty"`

	SSHTunnelEncryption *string `json:"sshTunnelEncryption,omitempty"`

	Syslog *Syslog `json:"syslog,omitempty"`

	Timezone *TimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *string `json:"tunnelType,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Version *string `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *common.Radio24 `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50 `json:"wifi50,omitempty"`
}

type CustomizedTimeZone struct {
	// Abbreviation
	// Time zone abbreviation
	Abbreviation *string `json:"abbreviation,omitempty"`

	End *DaylightSavingTime `json:"end,omitempty"`

	// GmtOffset
	// GMT offset
	GmtOffset *int `json:"gmtOffset,omitempty"`

	// GmtOffsetMinute
	// GMT offset minute
	GmtOffsetMinute *int `json:"gmtOffsetMinute,omitempty"`

	Start *DaylightSavingTime `json:"start,omitempty"`
}

type DaylightSavingTime struct {
	// Day
	// Day of the week (0 for Sunday, 1 for Monday, 2 for Tuesday, and so on)
	Day *int `json:"day,omitempty"`

	// Hour
	// Hour of the day
	Hour *int `json:"hour,omitempty"`

	// Month
	// Month when daylight saving time begins
	Month *int `json:"month,omitempty"`

	// Week
	// Week of the month (1 for the first week, 2 for the second week, and so on)
	Week *int `json:"week,omitempty"`
}

type DHCPSiteConfigList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.DHCPSiteConfigListRef `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DiffServConfiguration struct {
	Description *string `json:"description,omitempty"`

	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// ID
	// Identifier of the zone
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *UplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

type DiffServList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DiffServSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DiffServSummary struct {
	// ID
	// Identifier of the diff serv
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the diff serv
	Name *string `json:"name,omitempty"`
}

type DownlinkDiffServ struct {
	// Downlink
	// Downlink
	Downlink *string `json:"downlink,omitempty"`

	// DownlinkEnable
	// Downlink enable
	DownlinkEnable *bool `json:"downlinkEnable,omitempty"`
}

type MeshConfiguration struct {
	// MeshRadioIdx
	// Mesh radio index
	MeshRadioIdx *string `json:"meshRadioIdx,omitempty"`

	// Passphrase
	// Passphrase for the mesh network
	Passphrase *string `json:"passphrase,omitempty"`

	// Ssid
	// SSID of the mesh network
	Ssid *string `json:"ssid,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

type ModfiyApFirmware struct {
	// FirmwareVersion
	// new version of the AP firmare
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
}

type ModifyBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*BonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

type ModifyDiffServProfile struct {
	Description *string `json:"description,omitempty"`

	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	Name *string `json:"name,omitempty"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *UplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

type ModifyZone struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApLatencyInterval *common.ApLatencyInterval `json:"apLatencyInterval,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	ApRebootTimeout *common.ApRebootTimeout `json:"apRebootTimeout,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	BackgroundScanning24 *BackgroundScanning `json:"backgroundScanning24,omitempty"`

	BackgroundScanning50 *BackgroundScanning `json:"backgroundScanning50,omitempty"`

	BandBalancing *BandBalancing `json:"bandBalancing,omitempty"`

	BonjourFencingPolicy *common.GenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	ClientAdmissionControl24 *common.ClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *common.ClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	ClientLoadBalancing24 *ClientLoadBalancing `json:"clientLoadBalancing24,omitempty"`

	ClientLoadBalancing50 *ClientLoadBalancing `json:"clientLoadBalancing50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	CountryCode *string `json:"countryCode,omitempty"`

	Description *string `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DHCPSiteConfig *common.DHCPSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainID
	// Identifier of the management domain to which the zone belongs
	DomainID *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	HealthCheckSites []string `json:"healthCheckSites,omitempty"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	IpsecProfiles []*common.GenericRef `json:"ipsecProfiles,omitempty"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *float64 `json:"latitude,omitempty"`

	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty"`

	Location *string `json:"location,omitempty"`

	LocationAdditionalInfo *string `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.GenericRef `json:"locationBasedService,omitempty"`

	Login *ApLogin `json:"login,omitempty"`

	Longitude *float64 `json:"longitude,omitempty"`

	Mesh *MeshConfiguration `json:"mesh,omitempty"`

	Name *string `json:"name,omitempty"`

	NodeAffinityProfile *common.GenericRef `json:"nodeAffinityProfile,omitempty"`

	ProtectionMode24 *string `json:"protectionMode24,omitempty"`

	RecoverySsid *common.RecoverySsid `json:"recoverySsid,omitempty"`

	Rogue *Rogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	RuckusGreTunnelProfile *common.GenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	SmartMonitor *common.SmartMonitor `json:"smartMonitor,omitempty"`

	SNMPAgent *ApSNMPOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*SoftGreRef `json:"softGreTunnelProflies,omitempty"`

	SSHTunnelEncryption *string `json:"sshTunnelEncryption,omitempty"`

	Syslog *Syslog `json:"syslog,omitempty"`

	Timezone *TimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *string `json:"tunnelType,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *common.Radio24 `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50 `json:"wifi50,omitempty"`

	// ZoneAffinityProfileID
	// zone affinity profile Id
	ZoneAffinityProfileID *string `json:"zoneAffinityProfileId,omitempty"`
}

type QueryCriteria struct {
}

type Rogue struct {
	// MaliciousTypes
	// Malicious type when reportType is Malicious
	MaliciousTypes []string `json:"maliciousTypes,omitempty"`

	// ProtectionEnabled
	// Protection enabled
	ProtectionEnabled *bool `json:"protectionEnabled,omitempty"`

	// ReportType
	// Report type
	ReportType *string `json:"reportType,omitempty"`
}

type SNMPUser struct {
	// AuthPassword
	// authPassword of the SNMP User.
	AuthPassword *string `json:"authPassword,omitempty"`

	// AuthProtocol
	// authProtocol of the SNMP User.
	AuthProtocol *string `json:"authProtocol,omitempty"`

	// NotificationEnabled
	// notification privilege of the SNMP User
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP User
	NotificationTarget []*common.TargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	NotificationType *string `json:"notificationType,omitempty"`

	// PrivPassword
	// privPassword of the SNMP User.
	PrivPassword *string `json:"privPassword,omitempty"`

	// PrivProtocol
	// privProtocol of the SNMP User.
	PrivProtocol *string `json:"privProtocol,omitempty"`

	// ReadEnabled
	// read privilege of the SNMP User
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// UserName
	// name of the SNMP User.
	UserName *string `json:"userName,omitempty"`

	// WriteEnabled
	// write privilege of the SNMP User
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

type SoftGreRef struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type Syslog struct {
	Address *string `json:"address,omitempty"`

	// Facility
	// Facility of the syslog server
	Facility *string `json:"facility,omitempty"`

	// FlowLevel
	// Flow Level of the syslog
	FlowLevel *string `json:"flowLevel,omitempty"`

	// Port
	// Port number of the syslog server
	Port *int `json:"port,omitempty"`

	// Priority
	// Priority of the log messages
	Priority *string `json:"priority,omitempty"`

	// Protocol
	// Protocol of the syslog server
	Protocol *string `json:"protocol,omitempty"`

	SecondaryAddress *string `json:"secondaryAddress,omitempty"`

	// SecondaryPort
	// Secondary Server Port of the syslog server
	SecondaryPort *int `json:"secondaryPort,omitempty"`

	// SecondaryProtocol
	// Secondary Server Protocol of the syslog server
	SecondaryProtocol *string `json:"secondaryProtocol,omitempty"`
}

type TimezoneSetting struct {
	CustomizedTimezone *CustomizedTimeZone `json:"customizedTimezone,omitempty"`

	// SystemTimezone
	// System defined time zone, please refer to the “Overview > Time Zone” list
	SystemTimezone *string `json:"systemTimezone,omitempty"`
}

type UnsupportedApModel struct {
	// Amount
	// amount of the AP Model
	Amount *int `json:"amount,omitempty"`

	// Model
	// name of the AP Model
	Model *string `json:"model,omitempty"`
}

type UplinkDiffServ struct {
	// Uplink
	// Uplink
	Uplink *string `json:"uplink,omitempty"`

	// UplinkEnable
	// Uplink enable
	UplinkEnable *bool `json:"uplinkEnable,omitempty"`
}

type ZoneConfiguration struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApLatencyInterval *common.ApLatencyInterval `json:"apLatencyInterval,omitempty"`

	ApMgmtVlan *common.ApManagementVlan `json:"apMgmtVlan,omitempty"`

	ApRebootTimeout *common.ApRebootTimeout `json:"apRebootTimeout,omitempty"`

	AutoChannelSelection24 *common.AutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	AutoChannelSelection50 *common.AutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// venue code
	AwsVenue *string `json:"awsVenue,omitempty"`

	BackgroundScanning24 *BackgroundScanning `json:"backgroundScanning24,omitempty"`

	BackgroundScanning50 *BackgroundScanning `json:"backgroundScanning50,omitempty"`

	BandBalancing *BandBalancing `json:"bandBalancing,omitempty"`

	BonjourFencingPolicy *common.GenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	ClientAdmissionControl24 *common.ClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	ClientAdmissionControl50 *common.ClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	ClientLoadBalancing24 *ClientLoadBalancing `json:"clientLoadBalancing24,omitempty"`

	ClientLoadBalancing50 *ClientLoadBalancing `json:"clientLoadBalancing50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	CountryCode *string `json:"countryCode,omitempty"`

	Description *string `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DHCPSiteConfig *common.DHCPSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainID
	// Identifier of the management domain to which the zone belongs
	DomainID *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	HealthCheckSites []string `json:"healthCheckSites,omitempty"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	// ID
	// Identifier of the zone
	ID *string `json:"id,omitempty"`

	IPMode *string `json:"ipMode,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	IpsecProfiles []*common.GenericRef `json:"ipsecProfiles,omitempty"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *float64 `json:"latitude,omitempty"`

	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty"`

	Location *string `json:"location,omitempty"`

	LocationAdditionalInfo *string `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.GenericRef `json:"locationBasedService,omitempty"`

	Login *ApLogin `json:"login,omitempty"`

	Longitude *float64 `json:"longitude,omitempty"`

	Mesh *MeshConfiguration `json:"mesh,omitempty"`

	Name *string `json:"name,omitempty"`

	NodeAffinityProfile *common.GenericRef `json:"nodeAffinityProfile,omitempty"`

	ProtectionMode24 *string `json:"protectionMode24,omitempty"`

	RecoverySsid *common.RecoverySsid `json:"recoverySsid,omitempty"`

	Rogue *Rogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	RuckusGreTunnelProfile *common.GenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	SmartMonitor *common.SmartMonitor `json:"smartMonitor,omitempty"`

	SNMPAgent *ApSNMPOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*SoftGreRef `json:"softGreTunnelProflies,omitempty"`

	SSHTunnelEncryption *string `json:"sshTunnelEncryption,omitempty"`

	Syslog *Syslog `json:"syslog,omitempty"`

	Timezone *TimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *string `json:"tunnelType,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Version *string `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *common.Radio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50SuperSet `json:"wifi50,omitempty"`

	// ZoneAffinityProfileID
	// zone affinity profile Id
	ZoneAffinityProfileID *string `json:"zoneAffinityProfileId,omitempty"`
}

type ZoneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ZoneSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ZoneSummary struct {
	// ID
	// Identifier of the zone
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the zone
	Name *string `json:"name,omitempty"`

	// ServiceName
	// Name of the zone
	ServiceName *string `json:"serviceName,omitempty"`
}
