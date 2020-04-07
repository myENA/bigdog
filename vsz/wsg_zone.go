package vsz

// API Version: v9_0

type WSGZoneApFirmware struct {
	// FirmwareVersion
	// version of the AP firmare
	// Constraints:
	//    - nullable
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Supported
	// version of the AP firmare is supported for Upgrade or Downgrade.
	// Constraints:
	//    - nullable
	Supported *bool `json:"supported,omitempty"`

	// UnsupportedApModelSummary
	// summary of the AP Model is unsupported for AP firmware version.
	// Constraints:
	//    - nullable
	UnsupportedApModelSummary []*WSGZoneUnsupportedApModel `json:"unsupportedApModelSummary,omitempty" validate:"omitempty,dive"`
}

func NewWSGZoneApFirmware() *WSGZoneApFirmware {
	m := new(WSGZoneApFirmware)
	return m
}

type WSGZoneApFirmwareList struct {
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
	List []*WSGZoneApFirmware `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneApFirmwareList() *WSGZoneApFirmwareList {
	m := new(WSGZoneApFirmwareList)
	return m
}

type WSGZoneApLogin struct {
	// ApLoginName
	// Constraints:
	//    - required
	ApLoginName *WSGCommonApLoginName `json:"apLoginName" validate:"required"`

	// ApLoginPassword
	// Constraints:
	//    - required
	ApLoginPassword *WSGCommonApLoginPassword `json:"apLoginPassword" validate:"required"`
}

func NewWSGZoneApLogin() *WSGZoneApLogin {
	m := new(WSGZoneApLogin)
	return m
}

type WSGZoneApSnmpOptions struct {
	// ApSnmpEnabled
	// Enable AP SNMP
	// Constraints:
	//    - nullable
	ApSnmpEnabled *bool `json:"apSnmpEnabled,omitempty"`

	// SnmpV2Agent
	// Community List of the SNMP V2 Agent.
	// Constraints:
	//    - nullable
	SnmpV2Agent []*WSGCommonSnmpCommunity `json:"snmpV2Agent,omitempty" validate:"omitempty,dive"`

	// SnmpV3Agent
	// User List of the SNMP V3 Agent.
	// Constraints:
	//    - nullable
	SnmpV3Agent []*WSGZoneSnmpUser `json:"snmpV3Agent,omitempty" validate:"omitempty,dive"`
}

func NewWSGZoneApSnmpOptions() *WSGZoneApSnmpOptions {
	m := new(WSGZoneApSnmpOptions)
	return m
}

type WSGZoneAvailableTunnelProfile struct {
	// AaaAffinityEnabled
	// Enable AAA affinity (Soft GRE only)
	// Constraints:
	//    - nullable
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	// Id
	// Tunnel Profile ID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpMode
	// Constraints:
	//    - nullable
	IpMode *WSGCommonIpMode `json:"ipMode,omitempty"`

	// Name
	// Tunnel Profile Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// TunnelType
	// Tunnel Profile Type ("RuckusGRE", "SoftGRE",or "Ipsec")
	// Constraints:
	//    - nullable
	//    - oneof:[RuckusGRE,SoftGRE,Ipsec]
	TunnelType *string `json:"tunnelType,omitempty" validate:"omitempty,oneof=RuckusGRE SoftGRE Ipsec"`
}

func NewWSGZoneAvailableTunnelProfile() *WSGZoneAvailableTunnelProfile {
	m := new(WSGZoneAvailableTunnelProfile)
	return m
}

type WSGZoneAvailableTunnelProfileList struct {
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
	List []*WSGZoneAvailableTunnelProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneAvailableTunnelProfileList() *WSGZoneAvailableTunnelProfileList {
	m := new(WSGZoneAvailableTunnelProfileList)
	return m
}

type WSGZoneBackgroundScanning struct {
	// FrequencyInSec
	// Frequency in second
	// Constraints:
	//    - nullable
	//    - default:20
	//    - min:1
	//    - max:65535
	FrequencyInSec *int `json:"frequencyInSec,omitempty" validate:"omitempty,gte=1,lte=65535"`
}

func NewWSGZoneBackgroundScanning() *WSGZoneBackgroundScanning {
	m := new(WSGZoneBackgroundScanning)
	return m
}

type WSGZoneBandBalancing struct {
	// Wifi24Percentage
	// Percentage of client load on 2.4GHz radio band
	// Constraints:
	//    - nullable
	//    - default:25
	//    - min:0
	//    - max:100
	Wifi24Percentage *int `json:"wifi24Percentage,omitempty" validate:"omitempty,gte=0,lte=100"`
}

func NewWSGZoneBandBalancing() *WSGZoneBandBalancing {
	m := new(WSGZoneBandBalancing)
	return m
}

type WSGZoneBonjourGatewayPolicyConfiguration struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	// Constraints:
	//    - nullable
	BonjourPolicyRuleList []*WSGZoneBonjourPolicyRuleConfiguration `json:"bonjourPolicyRuleList,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGZoneBonjourGatewayPolicyConfiguration() *WSGZoneBonjourGatewayPolicyConfiguration {
	m := new(WSGZoneBonjourGatewayPolicyConfiguration)
	return m
}

type WSGZoneBonjourGatewayPolicyList struct {
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
	List []*WSGZoneBonjourGatewayPolicySummary `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneBonjourGatewayPolicyList() *WSGZoneBonjourGatewayPolicyList {
	m := new(WSGZoneBonjourGatewayPolicyList)
	return m
}

type WSGZoneBonjourGatewayPolicySummary struct {
	// Description
	// Description of the bonjour gateway policy
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the bonjour gateway policy
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LastModifiedBy
	// Last modified user of the bonjour gateway policy
	// Constraints:
	//    - nullable
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedOn
	// Last modified time of the bonjour gateway policy
	// Constraints:
	//    - nullable
	LastModifiedOn *string `json:"lastModifiedOn,omitempty"`

	// Name
	// Name of the bonjour gateway policy
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGZoneBonjourGatewayPolicySummary() *WSGZoneBonjourGatewayPolicySummary {
	m := new(WSGZoneBonjourGatewayPolicySummary)
	return m
}

// WSGZoneBonjourPolicyRule
//
// Bonjour policy rule
// Constraints:
//    - nullable
type WSGZoneBonjourPolicyRule struct {
	// BridgeService
	// Constraints:
	//    - required
	BridgeService *WSGProfileBridgeService `json:"bridgeService" validate:"required"`

	// FromVlan
	// From VLAN
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	FromVlan *int `json:"fromVlan" validate:"required,gte=1,lte=4094"`

	// Notes
	// Notes
	// Constraints:
	//    - nullable
	Notes *string `json:"notes,omitempty"`

	// Protocol
	// protocol. This is only available when bridgeService is OTHER
	// Constraints:
	//    - nullable
	Protocol *string `json:"protocol,omitempty"`

	// ToVlan
	// To VLAN
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	ToVlan *int `json:"toVlan" validate:"required,gte=1,lte=4094"`
}

func NewWSGZoneBonjourPolicyRule() *WSGZoneBonjourPolicyRule {
	m := new(WSGZoneBonjourPolicyRule)
	return m
}

// WSGZoneBonjourPolicyRuleConfiguration
//
// Bonjour policy rule
// Constraints:
//    - nullable
type WSGZoneBonjourPolicyRuleConfiguration struct {
	// BridgeService
	// Bridge service
	// Constraints:
	//    - nullable
	BridgeService *string `json:"bridgeService,omitempty"`

	// FromVlan
	// From VLAN
	// Constraints:
	//    - nullable
	FromVlan *int `json:"fromVlan,omitempty"`

	// Notes
	// Notes
	// Constraints:
	//    - nullable
	Notes *string `json:"notes,omitempty"`

	// Priority
	// Priority
	// Constraints:
	//    - nullable
	Priority *string `json:"priority,omitempty"`

	// Protocol
	// protocol
	// Constraints:
	//    - nullable
	Protocol *string `json:"protocol,omitempty"`

	// ToVlan
	// To VLAN
	// Constraints:
	//    - nullable
	ToVlan *int `json:"toVlan,omitempty"`
}

func NewWSGZoneBonjourPolicyRuleConfiguration() *WSGZoneBonjourPolicyRuleConfiguration {
	m := new(WSGZoneBonjourPolicyRuleConfiguration)
	return m
}

type WSGZoneClientLoadBalancing struct {
	// AdjacentRadioThreshold
	// Adjacent radio threshold
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AdjacentRadioThreshold *int `json:"adjacentRadioThreshold,omitempty" validate:"omitempty,gte=1,lte=100"`
}

func NewWSGZoneClientLoadBalancing() *WSGZoneClientLoadBalancing {
	m := new(WSGZoneClientLoadBalancing)
	return m
}

type WSGZoneCreateBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	// Constraints:
	//    - nullable
	BonjourPolicyRuleList []*WSGZoneBonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGZoneCreateBonjourGatewayPolicy() *WSGZoneCreateBonjourGatewayPolicy {
	m := new(WSGZoneCreateBonjourGatewayPolicy)
	return m
}

type WSGZoneCreateDiffServProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DownlinkDiffServ
	// Constraints:
	//    - nullable
	DownlinkDiffServ *WSGZoneDownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PreservedList
	// Preserved list
	// Constraints:
	//    - nullable
	PreservedList []string `json:"preservedList,omitempty" validate:"omitempty,dive"`

	// UplinkDiffServ
	// Constraints:
	//    - nullable
	UplinkDiffServ *WSGZoneUplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewWSGZoneCreateDiffServProfile() *WSGZoneCreateDiffServProfile {
	m := new(WSGZoneCreateDiffServProfile)
	return m
}

type WSGZoneCreateZone struct {
	// Altitude
	// Constraints:
	//    - nullable
	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApHccdEnabled
	// Historical Connection Failures allows the AP to report historical client connection failures for this zone.
	// Constraints:
	//    - nullable
	//    - default:false
	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	// ApHccdPersist
	// Allow Historical Connection Failures to be persisted.
	// Constraints:
	//    - nullable
	//    - default:true
	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

	// ApLatencyInterval
	// Constraints:
	//    - nullable
	ApLatencyInterval *WSGCommonApLatencyInterval `json:"apLatencyInterval,omitempty"`

	// ApMgmtVlan
	// Constraints:
	//    - nullable
	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	// ApRebootTimeout
	// Constraints:
	//    - nullable
	ApRebootTimeout *WSGCommonApRebootTimeout `json:"apRebootTimeout,omitempty"`

	// AutoChannelSelection24
	// Constraints:
	//    - nullable
	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	// AutoChannelSelection50
	// Constraints:
	//    - nullable
	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// BackgroundScanning24
	// Constraints:
	//    - nullable
	BackgroundScanning24 *WSGZoneBackgroundScanning `json:"backgroundScanning24,omitempty"`

	// BackgroundScanning50
	// Constraints:
	//    - nullable
	BackgroundScanning50 *WSGZoneBackgroundScanning `json:"backgroundScanning50,omitempty"`

	// BonjourFencingPolicy
	// Constraints:
	//    - nullable
	BonjourFencingPolicy *WSGCommonGenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	// Constraints:
	//    - nullable
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	// Constraints:
	//    - nullable
	//    - default:false
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	// Constraints:
	//    - nullable
	//    - default:false
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
	// Constraints:
	//    - nullable
	//    - default:false
	Channel144Enabled *bool `json:"channel144Enabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	// Constraints:
	//    - nullable
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"omitempty,gte=60,lte=3600"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	// Constraints:
	//    - nullable
	//    - default:false
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	// ClientAdmissionControl24
	// Constraints:
	//    - nullable
	ClientAdmissionControl24 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	// ClientAdmissionControl50
	// Constraints:
	//    - nullable
	ClientAdmissionControl50 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone.
	// Constraints:
	//    - nullable
	//    - default:false
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	// Constraints:
	//    - nullable
	CountryCode *string `json:"countryCode,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code.
	// Constraints:
	//    - nullable
	//    - default:false
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	// DhcpSiteConfig
	// Constraints:
	//    - nullable
	DhcpSiteConfig *WSGCommonDhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	// Constraints:
	//    - nullable
	//    - default:true
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	// Constraints:
	//    - nullable
	//    - default:true
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	// Constraints:
	//    - nullable
	//    - default:true
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	// Constraints:
	//    - nullable
	//    - default:30
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	// Constraints:
	//    - nullable
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	// Constraints:
	//    - nullable
	//    - default:60
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	// Constraints:
	//    - nullable
	//    - default:5
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// EnforcePriorityZoneAffinityEnable
	// Enforced the priority of Affinity Profile.
	// Constraints:
	//    - nullable
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	// Constraints:
	//    - nullable
	HealthCheckSites []string `json:"healthCheckSites,omitempty" validate:"omitempty,dive"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	// Constraints:
	//    - nullable
	//    - default:false
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	// IpsecProfile
	// Constraints:
	//    - nullable
	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	// Constraints:
	//    - nullable
	IpsecProfiles []*WSGCommonGenericRef `json:"ipsecProfiles,omitempty" validate:"omitempty,dive"`

	// IpsecTunnelMode
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty" validate:"omitempty,oneof=DISABLE SOFT_GRE RUCKUS_GRE"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	// Constraints:
	//    - nullable
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	// Latitude
	// Constraints:
	//    - nullable
	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	// LoadBalancing
	// Constraints:
	//    - nullable
	LoadBalancing *WSGZoneLoadBalancing `json:"loadBalancing,omitempty"`

	// Location
	// Constraints:
	//    - nullable
	Location *WSGCommonLocation `json:"location,omitempty"`

	// LocationAdditionalInfo
	// Constraints:
	//    - nullable
	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	// LocationBasedService
	// Constraints:
	//    - nullable
	LocationBasedService *WSGCommonGenericRef `json:"locationBasedService,omitempty"`

	// Login
	// Constraints:
	//    - required
	Login *WSGZoneApLogin `json:"login" validate:"required"`

	// Longitude
	// Constraints:
	//    - nullable
	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	// Constraints:
	//    - nullable
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty" validate:"omitempty,dive"`

	// Mesh
	// Constraints:
	//    - nullable
	Mesh *WSGZoneMeshConfiguration `json:"mesh,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// NodeAffinityProfile
	// Constraints:
	//    - nullable
	NodeAffinityProfile *WSGCommonGenericRef `json:"nodeAffinityProfile,omitempty"`

	// PaloAltoFirewallProfileId
	// Constraints:
	//    - nullable
	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	// ProtectionMode24
	// Constraints:
	//    - nullable
	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	// RecoverySsid
	// Constraints:
	//    - nullable
	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	// Constraints:
	//    - nullable
	//    - default:false
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// Rogue
	// Constraints:
	//    - nullable
	Rogue *WSGZoneRogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	// Constraints:
	//    - nullable
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	// Constraints:
	//    - nullable
	//    - default:false
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	// RogueApJammingThreshold
	// Constraints:
	//    - nullable
	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	// Constraints:
	//    - nullable
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// RuckusGreTunnelProfile
	// Constraints:
	//    - nullable
	RuckusGreTunnelProfile *WSGCommonGenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	// SmartMonitor
	// Constraints:
	//    - nullable
	SmartMonitor *WSGCommonSmartMonitor `json:"smartMonitor,omitempty"`

	// SnmpAgent
	// Constraints:
	//    - nullable
	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	// Constraints:
	//    - nullable
	SoftGreTunnelProflies []*WSGZoneSoftGreRef `json:"softGreTunnelProflies,omitempty" validate:"omitempty,dive"`

	// SshTunnelEncryption
	// Constraints:
	//    - nullable
	//    - default:'AES128'
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty" validate:"omitempty,oneof=AES128 AES256"`

	// Syslog
	// Constraints:
	//    - nullable
	Syslog *WSGZoneSyslog `json:"syslog,omitempty"`

	// Timezone
	// Constraints:
	//    - nullable
	Timezone *WSGZoneTimezoneSetting `json:"timezone,omitempty"`

	// TunnelProfile
	// Constraints:
	//    - nullable
	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// TunnelType
	// Constraints:
	//    - nullable
	TunnelType *WSGCommonZoneTunnelType `json:"tunnelType,omitempty"`

	// VenueProfile
	// Constraints:
	//    - nullable
	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	// Version
	// Constraints:
	//    - nullable
	Version *WSGCommonFirmwareVersion `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	// Constraints:
	//    - nullable
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	// Wifi24
	// Constraints:
	//    - nullable
	Wifi24 *WSGCommonRadio24 `json:"wifi24,omitempty"`

	// Wifi50
	// Constraints:
	//    - nullable
	Wifi50 *WSGCommonRadio50 `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// Zone affinity profile of the zone
	// Constraints:
	//    - nullable
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewWSGZoneCreateZone() *WSGZoneCreateZone {
	m := new(WSGZoneCreateZone)
	return m
}

type WSGZoneCustomizedTimeZone struct {
	// Abbreviation
	// Time zone abbreviation
	// Constraints:
	//    - required
	Abbreviation *string `json:"abbreviation" validate:"required"`

	// End
	// Constraints:
	//    - nullable
	End *WSGZoneDaylightSavingTime `json:"end,omitempty"`

	// GmtOffset
	// GMT offset
	// Constraints:
	//    - required
	//    - min:-11
	//    - max:14
	GmtOffset *int `json:"gmtOffset" validate:"required,gte=-11,lte=14"`

	// GmtOffsetMinute
	// GMT offset minute
	// Constraints:
	//    - required
	//    - min:0
	//    - max:59
	GmtOffsetMinute *int `json:"gmtOffsetMinute" validate:"required,gte=0,lte=59"`

	// Start
	// Constraints:
	//    - nullable
	Start *WSGZoneDaylightSavingTime `json:"start,omitempty"`
}

func NewWSGZoneCustomizedTimeZone() *WSGZoneCustomizedTimeZone {
	m := new(WSGZoneCustomizedTimeZone)
	return m
}

type WSGZoneDaylightSavingTime struct {
	// Day
	// Day of the week (0 for Sunday, 1 for Monday, 2 for Tuesday, and so on)
	// Constraints:
	//    - required
	//    - oneof:[0,1,2,3,4,5,6]
	Day *int `json:"day" validate:"required,oneof=0 1 2 3 4 5 6"`

	// Hour
	// Hour of the day
	// Constraints:
	//    - required
	//    - min:0
	//    - max:23
	Hour *int `json:"hour" validate:"required,gte=0,lte=23"`

	// Month
	// Month when daylight saving time begins
	// Constraints:
	//    - required
	//    - oneof:[1,2,3,4,5,6,7,8,9,10,11,12]
	Month *int `json:"month" validate:"required,oneof=1 2 3 4 5 6 7 8 9 10 11 12"`

	// Week
	// Week of the month (1 for the first week, 2 for the second week, and so on)
	// Constraints:
	//    - required
	//    - oneof:[1,2,3,4,5]
	Week *int `json:"week" validate:"required,oneof=1 2 3 4 5"`
}

func NewWSGZoneDaylightSavingTime() *WSGZoneDaylightSavingTime {
	m := new(WSGZoneDaylightSavingTime)
	return m
}

type WSGZoneDhcpSiteConfigList struct {
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
	List []*WSGCommonDhcpSiteConfigListRef `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneDhcpSiteConfigList() *WSGZoneDhcpSiteConfigList {
	m := new(WSGZoneDhcpSiteConfigList)
	return m
}

type WSGZoneDiffServConfiguration struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DownlinkDiffServ
	// Constraints:
	//    - nullable
	DownlinkDiffServ *WSGZoneDownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Id
	// Identifier of the zone
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PreservedList
	// Preserved list
	// Constraints:
	//    - nullable
	PreservedList []string `json:"preservedList,omitempty" validate:"omitempty,dive"`

	// UplinkDiffServ
	// Constraints:
	//    - nullable
	UplinkDiffServ *WSGZoneUplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewWSGZoneDiffServConfiguration() *WSGZoneDiffServConfiguration {
	m := new(WSGZoneDiffServConfiguration)
	return m
}

type WSGZoneDiffServList struct {
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
	List []*WSGZoneDiffServSummary `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneDiffServList() *WSGZoneDiffServList {
	m := new(WSGZoneDiffServList)
	return m
}

type WSGZoneDiffServSummary struct {
	// Id
	// Identifier of the diff serv
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the diff serv
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGZoneDiffServSummary() *WSGZoneDiffServSummary {
	m := new(WSGZoneDiffServSummary)
	return m
}

type WSGZoneDownlinkDiffServ struct {
	// Downlink
	// Downlink
	// Constraints:
	//    - nullable
	Downlink *string `json:"downlink,omitempty"`

	// DownlinkEnable
	// Downlink enable
	// Constraints:
	//    - nullable
	DownlinkEnable *bool `json:"downlinkEnable,omitempty"`
}

func NewWSGZoneDownlinkDiffServ() *WSGZoneDownlinkDiffServ {
	m := new(WSGZoneDownlinkDiffServ)
	return m
}

type WSGZoneLoadBalancing struct {
	// BandBalancing
	// Constraints:
	//    - nullable
	BandBalancing *WSGZoneBandBalancing `json:"bandBalancing,omitempty"`

	// ClientLoadBalancing24
	// Constraints:
	//    - nullable
	ClientLoadBalancing24 *WSGZoneClientLoadBalancing `json:"clientLoadBalancing24,omitempty"`

	// ClientLoadBalancing50
	// Constraints:
	//    - nullable
	ClientLoadBalancing50 *WSGZoneClientLoadBalancing `json:"clientLoadBalancing50,omitempty"`

	// LoadBalancingMethod
	// Constraints:
	//    - nullable
	//    - default:'BASED_ON_CLIENT_COUNT'
	//    - oneof:[BASED_ON_CLIENT_COUNT,BASED_ON_CAPACITY,OFF]
	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty" validate:"omitempty,oneof=BASED_ON_CLIENT_COUNT BASED_ON_CAPACITY OFF"`

	// SteeringMode
	// Steering Mode: BASIC-Withholds probe and authentication responses at connetcion time in heavily loaded band to balance clients to the other band, PROACTIVE-Uses BASIC functionality and actively rebalances clients via 802.11v BTM, STRICT-Uses PROACTIVE functionality and forcefully rebalances clients via 802.11v BTM
	// Constraints:
	//    - nullable
	//    - default:'BASIC'
	//    - oneof:[BASIC,PROACTIVE,STRICT]
	SteeringMode *string `json:"steeringMode,omitempty" validate:"omitempty,oneof=BASIC PROACTIVE STRICT"`
}

func NewWSGZoneLoadBalancing() *WSGZoneLoadBalancing {
	m := new(WSGZoneLoadBalancing)
	return m
}

type WSGZoneMeshConfiguration struct {
	// MeshRadioIdx
	// Mesh radio index
	// Constraints:
	//    - nullable
	//    - default:'Radio5G'
	//    - oneof:[Radio24G,Radio5G]
	MeshRadioIdx *string `json:"meshRadioIdx,omitempty" validate:"omitempty,oneof=Radio24G Radio5G"`

	// Passphrase
	// Passphrase for the mesh network
	// Constraints:
	//    - nullable
	Passphrase *string `json:"passphrase,omitempty"`

	// Ssid
	// SSID of the mesh network
	// Constraints:
	//    - nullable
	Ssid *string `json:"ssid,omitempty"`

	// ZeroTouchStatus
	// Constraints:
	//    - nullable
	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

func NewWSGZoneMeshConfiguration() *WSGZoneMeshConfiguration {
	m := new(WSGZoneMeshConfiguration)
	return m
}

type WSGZoneModfiyApFirmware struct {
	// FirmwareVersion
	// new version of the AP firmare
	// Constraints:
	//    - nullable
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
}

func NewWSGZoneModfiyApFirmware() *WSGZoneModfiyApFirmware {
	m := new(WSGZoneModfiyApFirmware)
	return m
}

type WSGZoneModifyBonjourGatewayEnable struct {
	// EnabledBonjourGateway
	// Enable Bonjour gateway on th AP
	// Constraints:
	//    - required
	EnabledBonjourGateway *bool `json:"enabledBonjourGateway" validate:"required"`
}

func NewWSGZoneModifyBonjourGatewayEnable() *WSGZoneModifyBonjourGatewayEnable {
	m := new(WSGZoneModifyBonjourGatewayEnable)
	return m
}

type WSGZoneModifyBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	// Constraints:
	//    - nullable
	BonjourPolicyRuleList []*WSGZoneBonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGZoneModifyBonjourGatewayPolicy() *WSGZoneModifyBonjourGatewayPolicy {
	m := new(WSGZoneModifyBonjourGatewayPolicy)
	return m
}

type WSGZoneModifyDiffServProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DownlinkDiffServ
	// Constraints:
	//    - nullable
	DownlinkDiffServ *WSGZoneDownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PreservedList
	// Preserved list
	// Constraints:
	//    - nullable
	PreservedList []string `json:"preservedList,omitempty" validate:"omitempty,dive"`

	// UplinkDiffServ
	// Constraints:
	//    - nullable
	UplinkDiffServ *WSGZoneUplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewWSGZoneModifyDiffServProfile() *WSGZoneModifyDiffServProfile {
	m := new(WSGZoneModifyDiffServProfile)
	return m
}

type WSGZoneModifyZone struct {
	// AaaAffinityEnabled
	// Constraints:
	//    - nullable
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	// Altitude
	// Constraints:
	//    - nullable
	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApHccdEnabled
	// Historical Connection Failures allows the AP to report historical client connection failures for this zone.
	// Constraints:
	//    - nullable
	//    - default:false
	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	// ApHccdPersist
	// Allow Historical Connection Failures to be persisted.
	// Constraints:
	//    - nullable
	//    - default:true
	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

	// ApLatencyInterval
	// Constraints:
	//    - nullable
	ApLatencyInterval *WSGCommonApLatencyInterval `json:"apLatencyInterval,omitempty"`

	// ApMgmtVlan
	// Constraints:
	//    - nullable
	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	// ApRebootTimeout
	// Constraints:
	//    - nullable
	ApRebootTimeout *WSGCommonApRebootTimeout `json:"apRebootTimeout,omitempty"`

	// AutoChannelSelection24
	// Constraints:
	//    - nullable
	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	// AutoChannelSelection50
	// Constraints:
	//    - nullable
	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Constraints:
	//    - nullable
	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	// BackgroundScanning24
	// Constraints:
	//    - nullable
	BackgroundScanning24 *WSGZoneBackgroundScanning `json:"backgroundScanning24,omitempty"`

	// BackgroundScanning50
	// Constraints:
	//    - nullable
	BackgroundScanning50 *WSGZoneBackgroundScanning `json:"backgroundScanning50,omitempty"`

	// BonjourFencingPolicy
	// Constraints:
	//    - nullable
	BonjourFencingPolicy *WSGCommonGenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	// Constraints:
	//    - nullable
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	// Constraints:
	//    - nullable
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	// Constraints:
	//    - nullable
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
	// Constraints:
	//    - nullable
	Channel144Enabled *bool `json:"channel144Enabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	// Constraints:
	//    - nullable
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"omitempty,gte=60,lte=3600"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	// Constraints:
	//    - nullable
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	// ClientAdmissionControl24
	// Constraints:
	//    - nullable
	ClientAdmissionControl24 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	// ClientAdmissionControl50
	// Constraints:
	//    - nullable
	ClientAdmissionControl50 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone
	// Constraints:
	//    - nullable
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	// Constraints:
	//    - nullable
	CountryCode *string `json:"countryCode,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	// Constraints:
	//    - nullable
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	// DhcpSiteConfig
	// Constraints:
	//    - nullable
	DhcpSiteConfig *WSGCommonDhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	// Constraints:
	//    - nullable
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	// Constraints:
	//    - nullable
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless.
	// Constraints:
	//    - nullable
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	// Constraints:
	//    - nullable
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	// Constraints:
	//    - nullable
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	// Constraints:
	//    - nullable
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	// Constraints:
	//    - nullable
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// EnforcePriorityZoneAffinityEnable
	// Enforce the priority of zone affinity
	// Constraints:
	//    - nullable
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	// Constraints:
	//    - nullable
	HealthCheckSites []string `json:"healthCheckSites,omitempty" validate:"omitempty,dive"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	// Constraints:
	//    - nullable
	//    - default:false
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	// IpsecProfile
	// Constraints:
	//    - nullable
	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	// Constraints:
	//    - nullable
	IpsecProfiles []*WSGCommonGenericRef `json:"ipsecProfiles,omitempty" validate:"omitempty,dive"`

	// IpsecTunnelMode
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty" validate:"omitempty,oneof=DISABLE SOFT_GRE RUCKUS_GRE"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	// Constraints:
	//    - nullable
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	// Latitude
	// Constraints:
	//    - nullable
	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	// LoadBalancing
	// Constraints:
	//    - nullable
	LoadBalancing *WSGZoneLoadBalancing `json:"loadBalancing,omitempty"`

	// Location
	// Constraints:
	//    - nullable
	Location *WSGCommonLocation `json:"location,omitempty"`

	// LocationAdditionalInfo
	// Constraints:
	//    - nullable
	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	// LocationBasedService
	// Constraints:
	//    - nullable
	LocationBasedService *WSGCommonGenericRef `json:"locationBasedService,omitempty"`

	// Login
	// Constraints:
	//    - nullable
	Login *WSGZoneApLogin `json:"login,omitempty"`

	// Longitude
	// Constraints:
	//    - nullable
	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	// Constraints:
	//    - nullable
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty" validate:"omitempty,dive"`

	// Mesh
	// Constraints:
	//    - nullable
	Mesh *WSGZoneMeshConfiguration `json:"mesh,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// NodeAffinityProfile
	// Constraints:
	//    - nullable
	NodeAffinityProfile *WSGCommonGenericRef `json:"nodeAffinityProfile,omitempty"`

	// PaloAltoFirewallProfileId
	// Constraints:
	//    - nullable
	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	// ProtectionMode24
	// Constraints:
	//    - nullable
	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	// RecoverySsid
	// Constraints:
	//    - nullable
	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RestrictedApAccessEnabled
	// Enable Restricted AP Access of the zone.
	// Constraints:
	//    - nullable
	RestrictedApAccessEnabled *bool `json:"restrictedApAccessEnabled,omitempty"`

	// RestrictedApAccessProfileId
	// Restricted AP Access Profile Id of the zone.
	// Constraints:
	//    - nullable
	RestrictedApAccessProfileId *string `json:"restrictedApAccessProfileId,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	// Constraints:
	//    - nullable
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// Rogue
	// Constraints:
	//    - nullable
	Rogue *WSGZoneRogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	// Constraints:
	//    - nullable
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	// Constraints:
	//    - nullable
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	// RogueApJammingThreshold
	// Constraints:
	//    - nullable
	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	// Constraints:
	//    - nullable
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// RuckusGreTunnelProfile
	// Constraints:
	//    - nullable
	RuckusGreTunnelProfile *WSGCommonGenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	// SmartMonitor
	// Constraints:
	//    - nullable
	SmartMonitor *WSGCommonSmartMonitor `json:"smartMonitor,omitempty"`

	// SnmpAgent
	// Constraints:
	//    - nullable
	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	// Constraints:
	//    - nullable
	SoftGreTunnelProflies []*WSGZoneSoftGreRef `json:"softGreTunnelProflies,omitempty" validate:"omitempty,dive"`

	// SshTunnelEncryption
	// Constraints:
	//    - nullable
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty" validate:"omitempty,oneof=AES128 AES256"`

	// Syslog
	// Constraints:
	//    - nullable
	Syslog *WSGZoneSyslog `json:"syslog,omitempty"`

	// Timezone
	// Constraints:
	//    - nullable
	Timezone *WSGZoneTimezoneSetting `json:"timezone,omitempty"`

	// TunnelProfile
	// Constraints:
	//    - nullable
	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// TunnelType
	// Constraints:
	//    - nullable
	TunnelType *WSGCommonZoneTunnelType `json:"tunnelType,omitempty"`

	// VenueProfile
	// Constraints:
	//    - nullable
	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	// Constraints:
	//    - nullable
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	// Wifi24
	// Constraints:
	//    - nullable
	Wifi24 *WSGCommonRadio24 `json:"wifi24,omitempty"`

	// Wifi50
	// Constraints:
	//    - nullable
	Wifi50 *WSGCommonRadio50 `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// Identifier of the ZoneAffinityProfile
	// Constraints:
	//    - nullable
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewWSGZoneModifyZone() *WSGZoneModifyZone {
	m := new(WSGZoneModifyZone)
	return m
}

type WSGZoneQueryCriteria struct {
	// Attributes
	// Get specific columns only
	// Constraints:
	//    - nullable
	Attributes []string `json:"attributes,omitempty" validate:"omitempty,dive"`

	// Criteria
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	// Constraints:
	//    - nullable
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraFilters []*WSGZoneQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty" validate:"omitempty,dive"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraNotFilters []*WSGCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty" validate:"omitempty,dive"`

	// ExtraTimeRange
	// Constraints:
	//    - nullable
	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	// Constraints:
	//    - nullable
	Filters []*WSGZoneQueryCriteriaFiltersType `json:"filters,omitempty" validate:"omitempty,dive"`

	// FullTextSearch
	// Constraints:
	//    - nullable
	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required information.
	// Constraints:
	//    - nullable
	Options *WSGZoneQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - nullable
	//    - min:1
	Page *int `json:"page,omitempty" validate:"omitempty,gte=1"`

	// Query
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	// Constraints:
	//    - nullable
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGZoneQueryCriteria() *WSGZoneQueryCriteria {
	m := new(WSGZoneQueryCriteria)
	return m
}

type WSGZoneQueryCriteriaExtraFiltersType struct {
	// Operator
	// Constraints:
	//    - nullable
	Operator *WSGCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attribute
	// Constraints:
	//    - required
	//    - oneof:[VERSION]
	Type *string `json:"type" validate:"required,oneof=VERSION"`

	// Value
	// Value for filtering
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`
}

func NewWSGZoneQueryCriteriaExtraFiltersType() *WSGZoneQueryCriteriaExtraFiltersType {
	m := new(WSGZoneQueryCriteriaExtraFiltersType)
	return m
}

type WSGZoneQueryCriteriaFiltersType struct {
	// Operator
	// Operator for filtering
	// Constraints:
	//    - nullable
	//    - oneof:[eq]
	Operator *string `json:"operator,omitempty" validate:"omitempty,oneof=eq"`

	// Type
	// Group type
	// Constraints:
	//    - required
	//    - oneof:[SYSTEM,CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,SWITCH_GROUP]
	Type *string `json:"type" validate:"required,oneof=SYSTEM CONTROLBLADE DATABLADE DOMAIN ZONE THIRD_PARTY_ZONE APGROUP WLANGROUP INDOORMAP AP WLAN SWITCH_GROUP"`

	// Value
	// Group ID
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`
}

func NewWSGZoneQueryCriteriaFiltersType() *WSGZoneQueryCriteriaFiltersType {
	m := new(WSGZoneQueryCriteriaFiltersType)
	return m
}

// WSGZoneQueryCriteriaOptionsType
//
// Specified feature required information.
// Constraints:
//    - nullable
type WSGZoneQueryCriteriaOptionsType struct {
	// IncludeSharedResources
	// Include the resources of parent domain as well while querying.
	// Constraints:
	//    - nullable
	IncludeSharedResources *bool `json:"includeSharedResources,omitempty"`

	// Zoneipmode
	// Constraints:
	//    - nullable
	Zoneipmode *WSGCommonIpMode `json:"zone_ipmode,omitempty"`
}

func NewWSGZoneQueryCriteriaOptionsType() *WSGZoneQueryCriteriaOptionsType {
	m := new(WSGZoneQueryCriteriaOptionsType)
	return m
}

type WSGZoneRogue struct {
	// MaliciousTypes
	// Malicious type when reportType is Malicious
	// Constraints:
	//    - nullable
	MaliciousTypes []string `json:"maliciousTypes,omitempty" validate:"omitempty,dive"`

	// ProtectionEnabled
	// Protection enabled
	// Constraints:
	//    - nullable
	ProtectionEnabled *bool `json:"protectionEnabled,omitempty"`

	// ReportType
	// Report type
	// Constraints:
	//    - nullable
	//    - oneof:[All,Malicious]
	ReportType *string `json:"reportType,omitempty" validate:"omitempty,oneof=All Malicious"`
}

func NewWSGZoneRogue() *WSGZoneRogue {
	m := new(WSGZoneRogue)
	return m
}

type WSGZoneSnmpUser struct {
	// AuthPassword
	// authPassword of the SNMP User.
	// Constraints:
	//    - nullable
	//    - min:8
	AuthPassword *string `json:"authPassword,omitempty" validate:"omitempty,min=8"`

	// AuthProtocol
	// authProtocol of the SNMP User.
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,MD5,SHA]
	AuthProtocol *string `json:"authProtocol,omitempty" validate:"omitempty,oneof=NONE MD5 SHA"`

	// NotificationEnabled
	// notification privilege of the SNMP User
	// Constraints:
	//    - nullable
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP User
	// Constraints:
	//    - nullable
	NotificationTarget []*WSGCommonTargetConfig `json:"notificationTarget,omitempty" validate:"omitempty,dive"`

	// NotificationType
	// type of the notification privilege
	// Constraints:
	//    - nullable
	//    - oneof:[TRAP,INFORM]
	NotificationType *string `json:"notificationType,omitempty" validate:"omitempty,oneof=TRAP INFORM"`

	// PrivPassword
	// privPassword of the SNMP User.
	// Constraints:
	//    - nullable
	//    - min:8
	PrivPassword *string `json:"privPassword,omitempty" validate:"omitempty,min=8"`

	// PrivProtocol
	// privProtocol of the SNMP User.
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,DES,AES]
	PrivProtocol *string `json:"privProtocol,omitempty" validate:"omitempty,oneof=NONE DES AES"`

	// ReadEnabled
	// read privilege of the SNMP User
	// Constraints:
	//    - nullable
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// UserName
	// name of the SNMP User.
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`

	// WriteEnabled
	// write privilege of the SNMP User
	// Constraints:
	//    - nullable
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

func NewWSGZoneSnmpUser() *WSGZoneSnmpUser {
	m := new(WSGZoneSnmpUser)
	return m
}

type WSGZoneSoftGreRef struct {
	// AaaAffinityEnabled
	// Constraints:
	//    - nullable
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGZoneSoftGreRef() *WSGZoneSoftGreRef {
	m := new(WSGZoneSoftGreRef)
	return m
}

type WSGZoneSyslog struct {
	// Address
	// Constraints:
	//    - nullable
	Address *WSGCommonIpAddress `json:"address,omitempty"`

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
	//    - default:514
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

	// SecondaryAddress
	// Constraints:
	//    - nullable
	SecondaryAddress *WSGCommonIpAddress `json:"secondaryAddress,omitempty"`

	// SecondaryPort
	// Secondary Server Port of the syslog server
	// Constraints:
	//    - nullable
	//    - default:514
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

func NewWSGZoneSyslog() *WSGZoneSyslog {
	m := new(WSGZoneSyslog)
	return m
}

type WSGZoneTimezoneSetting struct {
	// CustomizedTimezone
	// Constraints:
	//    - nullable
	CustomizedTimezone *WSGZoneCustomizedTimeZone `json:"customizedTimezone,omitempty"`

	// SystemTimezone
	// System defined time zone, please refer to the 'Overview > Time Zone' list
	// Constraints:
	//    - nullable
	SystemTimezone *string `json:"systemTimezone,omitempty"`
}

func NewWSGZoneTimezoneSetting() *WSGZoneTimezoneSetting {
	m := new(WSGZoneTimezoneSetting)
	return m
}

type WSGZoneUnsupportedApModel struct {
	// Amount
	// amount of the AP Model
	// Constraints:
	//    - nullable
	Amount *int `json:"amount,omitempty"`

	// Model
	// name of the AP Model
	// Constraints:
	//    - nullable
	Model *string `json:"model,omitempty"`
}

func NewWSGZoneUnsupportedApModel() *WSGZoneUnsupportedApModel {
	m := new(WSGZoneUnsupportedApModel)
	return m
}

type WSGZoneUplinkDiffServ struct {
	// Uplink
	// Uplink
	// Constraints:
	//    - nullable
	Uplink *string `json:"uplink,omitempty"`

	// UplinkEnable
	// Uplink enable
	// Constraints:
	//    - nullable
	UplinkEnable *bool `json:"uplinkEnable,omitempty"`
}

func NewWSGZoneUplinkDiffServ() *WSGZoneUplinkDiffServ {
	m := new(WSGZoneUplinkDiffServ)
	return m
}

type WSGZoneConfiguration struct {
	// AaaAffinityEnabled
	// Constraints:
	//    - nullable
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	// Altitude
	// Constraints:
	//    - nullable
	Altitude *WSGCommonAltitude `json:"altitude,omitempty"`

	// ApHccdEnabled
	// Historical Connection Failures allows the AP to report historical client connection failures for this zone.
	// Constraints:
	//    - nullable
	//    - default:false
	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	// ApHccdPersist
	// Allow Historical Connection Failures to be persisted.
	// Constraints:
	//    - nullable
	//    - default:true
	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

	// ApLatencyInterval
	// Constraints:
	//    - nullable
	ApLatencyInterval *WSGCommonApLatencyInterval `json:"apLatencyInterval,omitempty"`

	// ApMgmtVlan
	// Constraints:
	//    - nullable
	ApMgmtVlan *WSGCommonApManagementVlan `json:"apMgmtVlan,omitempty"`

	// ApRebootTimeout
	// Constraints:
	//    - nullable
	ApRebootTimeout *WSGCommonApRebootTimeout `json:"apRebootTimeout,omitempty"`

	// AutoChannelSelection24
	// Constraints:
	//    - nullable
	AutoChannelSelection24 *WSGCommonAutoChannelSelection `json:"autoChannelSelection24,omitempty"`

	// AutoChannelSelection50
	// Constraints:
	//    - nullable
	AutoChannelSelection50 *WSGCommonAutoChannelSelection `json:"autoChannelSelection50,omitempty"`

	// AwsVenue
	// Constraints:
	//    - nullable
	AwsVenue *WSGCommonAwsVenue `json:"awsVenue,omitempty"`

	// BackgroundScanning24
	// Constraints:
	//    - nullable
	BackgroundScanning24 *WSGZoneBackgroundScanning `json:"backgroundScanning24,omitempty"`

	// BackgroundScanning50
	// Constraints:
	//    - nullable
	BackgroundScanning50 *WSGZoneBackgroundScanning `json:"backgroundScanning50,omitempty"`

	// BonjourFencingPolicy
	// Constraints:
	//    - nullable
	BonjourFencingPolicy *WSGCommonGenericRef `json:"bonjourFencingPolicy,omitempty"`

	// BonjourFencingPolicyEnabled
	// Enable Bonjour Fencing Policy on the AP
	// Constraints:
	//    - nullable
	BonjourFencingPolicyEnabled *bool `json:"bonjourFencingPolicyEnabled,omitempty"`

	// CbandChannelEnabled
	// 5.8Ghz channels enabled configuration of the zone.
	// Constraints:
	//    - nullable
	CbandChannelEnabled *bool `json:"cbandChannelEnabled,omitempty"`

	// CbandChannelLicenseEnabled
	// 5.8Ghz channels license enabled configuration of the zone.
	// Constraints:
	//    - nullable
	CbandChannelLicenseEnabled *bool `json:"cbandChannelLicenseEnabled,omitempty"`

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
	// Constraints:
	//    - nullable
	Channel144Enabled *bool `json:"channel144Enabled,omitempty"`

	// ChannelEvaluationInterval
	// channel evaluation Interval of the zone
	// Constraints:
	//    - nullable
	//    - default:600
	//    - min:60
	//    - max:3600
	ChannelEvaluationInterval *int `json:"channelEvaluationInterval,omitempty" validate:"omitempty,gte=60,lte=3600"`

	// ChannelModeEnabled
	// Channel mode configuration of the zone.
	// Constraints:
	//    - nullable
	ChannelModeEnabled *bool `json:"channelModeEnabled,omitempty"`

	// ClientAdmissionControl24
	// Constraints:
	//    - nullable
	ClientAdmissionControl24 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl24,omitempty"`

	// ClientAdmissionControl50
	// Constraints:
	//    - nullable
	ClientAdmissionControl50 *WSGCommonClientAdmissionControl `json:"clientAdmissionControl50,omitempty"`

	// ClusterRedundancyEnabled
	// Enable Cluster redundancy on zone
	// Constraints:
	//    - nullable
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// CountryCode
	// Country code of the zone
	// Constraints:
	//    - nullable
	CountryCode *string `json:"countryCode,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	// Constraints:
	//    - nullable
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	// DhcpSiteConfig
	// Constraints:
	//    - nullable
	DhcpSiteConfig *WSGCommonDhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	// Constraints:
	//    - nullable
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	// Constraints:
	//    - nullable
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	// Constraints:
	//    - nullable
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DosBarringCheckPeriod
	// DoS Protection(Barring UE) check period of the zone.
	// Constraints:
	//    - nullable
	DosBarringCheckPeriod *int `json:"dosBarringCheckPeriod,omitempty"`

	// DosBarringEnable
	// Enable DoS Protection(Barring UE) of the zone.
	// Constraints:
	//    - nullable
	DosBarringEnable *int `json:"dosBarringEnable,omitempty"`

	// DosBarringPeriod
	// DoS Protection(Barring UE) blocking period of the zone.
	// Constraints:
	//    - nullable
	DosBarringPeriod *int `json:"dosBarringPeriod,omitempty"`

	// DosBarringThreshold
	// DoS Protection(Barring UE) threshold of the zone.
	// Constraints:
	//    - nullable
	DosBarringThreshold *int `json:"dosBarringThreshold,omitempty"`

	// EnforcePriorityZoneAffinityEnable
	// Constraints:
	//    - nullable
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	// Constraints:
	//    - nullable
	HealthCheckSites []string `json:"healthCheckSites,omitempty" validate:"omitempty,dive"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	// Constraints:
	//    - nullable
	//    - default:false
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	// Id
	// Identifier of the zone
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpMode
	// Constraints:
	//    - nullable
	IpMode *WSGCommonIpMode `json:"ipMode,omitempty"`

	// IpsecProfile
	// Constraints:
	//    - nullable
	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	// Constraints:
	//    - nullable
	IpsecProfiles []*WSGCommonGenericRef `json:"ipsecProfiles,omitempty" validate:"omitempty,dive"`

	// IpsecTunnelMode
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty" validate:"omitempty,oneof=DISABLE SOFT_GRE RUCKUS_GRE"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	// Constraints:
	//    - nullable
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	// Latitude
	// Constraints:
	//    - nullable
	Latitude *WSGCommonLatitude `json:"latitude,omitempty"`

	// LoadBalancing
	// Constraints:
	//    - nullable
	LoadBalancing *WSGZoneLoadBalancing `json:"loadBalancing,omitempty"`

	// Location
	// Constraints:
	//    - nullable
	Location *WSGCommonLocation `json:"location,omitempty"`

	// LocationAdditionalInfo
	// Constraints:
	//    - nullable
	LocationAdditionalInfo *WSGCommonLocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	// LocationBasedService
	// Constraints:
	//    - nullable
	LocationBasedService *WSGCommonGenericRef `json:"locationBasedService,omitempty"`

	// Login
	// Constraints:
	//    - nullable
	Login *WSGZoneApLogin `json:"login,omitempty"`

	// Longitude
	// Constraints:
	//    - nullable
	Longitude *WSGCommonLongitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	// Constraints:
	//    - nullable
	LteBandLockChannels []*WSGCommonLteBandLockChannel `json:"lteBandLockChannels,omitempty" validate:"omitempty,dive"`

	// Mesh
	// Constraints:
	//    - nullable
	Mesh *WSGZoneMeshConfiguration `json:"mesh,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// NodeAffinityProfile
	// Constraints:
	//    - nullable
	NodeAffinityProfile *WSGCommonGenericRef `json:"nodeAffinityProfile,omitempty"`

	// PaloAltoFirewallProfileId
	// Constraints:
	//    - nullable
	PaloAltoFirewallProfileId *string `json:"paloAltoFirewallProfileId,omitempty"`

	// ProtectionMode24
	// Constraints:
	//    - nullable
	ProtectionMode24 *WSGCommonProtectionMode `json:"protectionMode24,omitempty"`

	// RecoverySsid
	// Constraints:
	//    - nullable
	RecoverySsid *WSGCommonRecoverySsid `json:"recoverySsid,omitempty"`

	// RestrictedApAccessEnabled
	// Enable Restricted AP Access of the zone.
	// Constraints:
	//    - nullable
	RestrictedApAccessEnabled *bool `json:"restrictedApAccessEnabled,omitempty"`

	// RestrictedApAccessProfileId
	// Restricted AP Access Profile Id of the zone.
	// Constraints:
	//    - nullable
	RestrictedApAccessProfileId *string `json:"restrictedApAccessProfileId,omitempty"`

	// RksGreForwardBroadcast
	// Ruckus GRE tunnel broadcast packet forwarding
	// Constraints:
	//    - nullable
	//    - default:false
	RksGreForwardBroadcast *bool `json:"rksGreForwardBroadcast,omitempty"`

	// Rogue
	// Constraints:
	//    - nullable
	Rogue *WSGZoneRogue `json:"rogue,omitempty"`

	// RogueApAggressivenessMode
	// Adjust the frequency interval to de-authenticate rogue APs.
	// Constraints:
	//    - nullable
	RogueApAggressivenessMode *int `json:"rogueApAggressivenessMode,omitempty"`

	// RogueApJammingDetection
	// Enable jamming detection.
	// Constraints:
	//    - nullable
	RogueApJammingDetection *bool `json:"rogueApJammingDetection,omitempty"`

	// RogueApJammingThreshold
	// Constraints:
	//    - nullable
	RogueApJammingThreshold *int `json:"rogueApJammingThreshold,omitempty"`

	// RogueApReportThreshold
	// Rogue AP report will leave out all entries that have signal strength lower than this threshold.
	// Constraints:
	//    - nullable
	RogueApReportThreshold *int `json:"rogueApReportThreshold,omitempty"`

	// RuckusGreTunnelProfile
	// Constraints:
	//    - nullable
	RuckusGreTunnelProfile *WSGCommonGenericRef `json:"ruckusGreTunnelProfile,omitempty"`

	// SmartMonitor
	// Constraints:
	//    - nullable
	SmartMonitor *WSGCommonSmartMonitor `json:"smartMonitor,omitempty"`

	// SnmpAgent
	// Constraints:
	//    - nullable
	SnmpAgent *WSGZoneApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	// Constraints:
	//    - nullable
	SoftGreTunnelProflies []*WSGZoneSoftGreRef `json:"softGreTunnelProflies,omitempty" validate:"omitempty,dive"`

	// SshTunnelEncryption
	// Constraints:
	//    - nullable
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty" validate:"omitempty,oneof=AES128 AES256"`

	// Syslog
	// Constraints:
	//    - nullable
	Syslog *WSGZoneSyslog `json:"syslog,omitempty"`

	// Timezone
	// Constraints:
	//    - nullable
	Timezone *WSGZoneTimezoneSetting `json:"timezone,omitempty"`

	// TunnelProfile
	// Constraints:
	//    - nullable
	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// TunnelType
	// Constraints:
	//    - nullable
	TunnelType *WSGCommonZoneTunnelType `json:"tunnelType,omitempty"`

	// VenueProfile
	// Constraints:
	//    - nullable
	VenueProfile *WSGCommonGenericRef `json:"venueProfile,omitempty"`

	// Version
	// Constraints:
	//    - nullable
	Version *WSGCommonFirmwareVersion `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	// Constraints:
	//    - nullable
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	// Wifi24
	// Constraints:
	//    - nullable
	Wifi24 *WSGCommonRadio24SuperSet `json:"wifi24,omitempty"`

	// Wifi50
	// Constraints:
	//    - nullable
	Wifi50 *WSGCommonRadio50SuperSet `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// zone affinity profile Id
	// Constraints:
	//    - nullable
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewWSGZoneConfiguration() *WSGZoneConfiguration {
	m := new(WSGZoneConfiguration)
	return m
}

type WSGZoneList struct {
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
	List []*WSGZoneSummary `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGZoneList() *WSGZoneList {
	m := new(WSGZoneList)
	return m
}

type WSGZoneSummary struct {
	// Id
	// Identifier of the zone
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the zone
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// ServiceName
	// Name of the zone
	// Constraints:
	//    - nullable
	ServiceName *string `json:"serviceName,omitempty"`
}

func NewWSGZoneSummary() *WSGZoneSummary {
	m := new(WSGZoneSummary)
	return m
}
