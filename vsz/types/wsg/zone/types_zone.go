package zone

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

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

func NewApFirmware() *ApFirmware {
	apFirmwareType := new(ApFirmware)
	return apFirmwareType
}

func NewApFirmwareWithDefaults() *ApFirmware {
	apFirmwareType := new(ApFirmware)
	return apFirmwareType
}

type ApFirmwareList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApFirmware `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApFirmwareList() *ApFirmwareList {
	apFirmwareListType := new(ApFirmwareList)
	return apFirmwareListType
}

func NewApFirmwareListWithDefaults() *ApFirmwareList {
	apFirmwareListType := new(ApFirmwareList)
	return apFirmwareListType
}

type ApLogin struct {
	// ApLoginName
	// Constraints:
	//    - required
	ApLoginName *common.ApLoginName `json:"apLoginName" validate:"required"`

	// ApLoginPassword
	// Constraints:
	//    - required
	ApLoginPassword *common.ApLoginPassword `json:"apLoginPassword" validate:"required"`
}

func NewApLogin() *ApLogin {
	apLoginType := new(ApLogin)
	return apLoginType
}

func NewApLoginWithDefaults() *ApLogin {
	apLoginType := new(ApLogin)
	return apLoginType
}

type ApSnmpOptions struct {
	// ApSnmpEnabled
	// Enable AP SNMP
	ApSnmpEnabled *bool `json:"apSnmpEnabled,omitempty"`

	// SnmpV2Agent
	// Community List of the SNMP V2 Agent.
	SnmpV2Agent []*common.SnmpCommunity `json:"snmpV2Agent,omitempty"`

	// SnmpV3Agent
	// User List of the SNMP V3 Agent.
	SnmpV3Agent []*SnmpUser `json:"snmpV3Agent,omitempty"`
}

func NewApSnmpOptions() *ApSnmpOptions {
	apSnmpOptionsType := new(ApSnmpOptions)
	return apSnmpOptionsType
}

func NewApSnmpOptionsWithDefaults() *ApSnmpOptions {
	apSnmpOptionsType := new(ApSnmpOptions)
	return apSnmpOptionsType
}

type AvailableTunnelProfile struct {
	// AaaAffinityEnabled
	// Enable AAA affinity (Soft GRE only)
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	// Id
	// Tunnel Profile ID
	Id *string `json:"id,omitempty"`

	IpMode *common.IpMode `json:"ipMode,omitempty"`

	// Name
	// Tunnel Profile Name
	Name *string `json:"name,omitempty"`

	// TunnelType
	// Tunnel Profile Type ("RuckusGRE", "SoftGRE",or "Ipsec")
	// Constraints:
	//    - nullable
	//    - oneof:[RuckusGRE,SoftGRE,Ipsec]
	TunnelType *string `json:"tunnelType,omitempty" validate:"omitempty,oneof=RuckusGRE SoftGRE Ipsec"`
}

func NewAvailableTunnelProfile() *AvailableTunnelProfile {
	availableTunnelProfileType := new(AvailableTunnelProfile)
	return availableTunnelProfileType
}

func NewAvailableTunnelProfileWithDefaults() *AvailableTunnelProfile {
	availableTunnelProfileType := new(AvailableTunnelProfile)
	return availableTunnelProfileType
}

type AvailableTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AvailableTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAvailableTunnelProfileList() *AvailableTunnelProfileList {
	availableTunnelProfileListType := new(AvailableTunnelProfileList)
	return availableTunnelProfileListType
}

func NewAvailableTunnelProfileListWithDefaults() *AvailableTunnelProfileList {
	availableTunnelProfileListType := new(AvailableTunnelProfileList)
	return availableTunnelProfileListType
}

type BackgroundScanning struct {
	// FrequencyInSec
	// Frequency in second
	// Constraints:
	//    - nullable
	//    - default:20
	//    - min:1
	//    - max:65535
	FrequencyInSec *int `json:"frequencyInSec,omitempty" validate:"omitempty,gte=1,lte=65535"`
}

func NewBackgroundScanning() *BackgroundScanning {
	backgroundScanningType := new(BackgroundScanning)
	return backgroundScanningType
}

func NewBackgroundScanningWithDefaults() *BackgroundScanning {
	backgroundScanningType := new(BackgroundScanning)
	frequencyInSecField := 20
	backgroundScanningType.FrequencyInSec = &frequencyInSecField
	return backgroundScanningType
}

type BandBalancing struct {
	// Mode
	// Band Balancing Mode: BASIC-Withholds probe and authentication responses at connetcion time in heavily loaded band to balance clients to the other band, PROACTIVE-Uses BASIC functionality and actively rebalances clients via 802.11v BTM, STRICT-Uses PROACTIVE functionality and forcefully rebalances clients via 802.11v BTM
	// Constraints:
	//    - nullable
	//    - default:'BASIC'
	//    - oneof:[BASIC,PROACTIVE,STRICT]
	Mode *string `json:"mode,omitempty" validate:"omitempty,oneof=BASIC PROACTIVE STRICT"`

	// Wifi24Percentage
	// Percentage of client load on 2.4GHz radio band
	// Constraints:
	//    - nullable
	//    - default:25
	//    - min:0
	//    - max:100
	Wifi24Percentage *int `json:"wifi24Percentage,omitempty" validate:"omitempty,gte=0,lte=100"`
}

func NewBandBalancing() *BandBalancing {
	bandBalancingType := new(BandBalancing)
	return bandBalancingType
}

func NewBandBalancingWithDefaults() *BandBalancing {
	bandBalancingType := new(BandBalancing)
	modeField := `BASIC`
	bandBalancingType.Mode = &modeField
	wifi24PercentageField := 25
	bandBalancingType.Wifi24Percentage = &wifi24PercentageField
	return bandBalancingType
}

type BonjourGatewayPolicyConfiguration struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*BonjourPolicyRuleConfiguration `json:"bonjourPolicyRuleList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewBonjourGatewayPolicyConfiguration() *BonjourGatewayPolicyConfiguration {
	bonjourGatewayPolicyConfigurationType := new(BonjourGatewayPolicyConfiguration)
	return bonjourGatewayPolicyConfigurationType
}

func NewBonjourGatewayPolicyConfigurationWithDefaults() *BonjourGatewayPolicyConfiguration {
	bonjourGatewayPolicyConfigurationType := new(BonjourGatewayPolicyConfiguration)
	return bonjourGatewayPolicyConfigurationType
}

type BonjourGatewayPolicyList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BonjourGatewayPolicySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewBonjourGatewayPolicyList() *BonjourGatewayPolicyList {
	bonjourGatewayPolicyListType := new(BonjourGatewayPolicyList)
	return bonjourGatewayPolicyListType
}

func NewBonjourGatewayPolicyListWithDefaults() *BonjourGatewayPolicyList {
	bonjourGatewayPolicyListType := new(BonjourGatewayPolicyList)
	return bonjourGatewayPolicyListType
}

type BonjourGatewayPolicySummary struct {
	// Description
	// Description of the bonjour gateway policy
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the bonjour gateway policy
	Id *string `json:"id,omitempty"`

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

func NewBonjourGatewayPolicySummary() *BonjourGatewayPolicySummary {
	bonjourGatewayPolicySummaryType := new(BonjourGatewayPolicySummary)
	return bonjourGatewayPolicySummaryType
}

func NewBonjourGatewayPolicySummaryWithDefaults() *BonjourGatewayPolicySummary {
	bonjourGatewayPolicySummaryType := new(BonjourGatewayPolicySummary)
	return bonjourGatewayPolicySummaryType
}

// BonjourPolicyRule
//
// Bonjour policy rule
type BonjourPolicyRule struct {
	// BridgeService
	// Bridge service
	// Constraints:
	//    - required
	//    - oneof:[AIRDISK,AIRPLAY,AIRPORT_MANAGEMENT,AIRPRINT,AIRTUNES,APPLE_FILE_SHARING,APPLE_MOBILE_DEVICES,APPLETV,ICLOUD_SYNC,ITUNES_REMOTE,ITUNES_SHARING,OPEN_DIRECTORY_MASTER,OPTICAL_DISK_SHARING,SCREEN_SHARING,SECURE_FILE_SHARING,SECURE_SHELL,WWW_HTTP,WWW_HTTPS,WORKGROUP_MANAGER,XGRID,OTHER]
	BridgeService *string `json:"bridgeService" validate:"required,oneof=AIRDISK AIRPLAY AIRPORT_MANAGEMENT AIRPRINT AIRTUNES APPLE_FILE_SHARING APPLE_MOBILE_DEVICES APPLETV ICLOUD_SYNC ITUNES_REMOTE ITUNES_SHARING OPEN_DIRECTORY_MASTER OPTICAL_DISK_SHARING SCREEN_SHARING SECURE_FILE_SHARING SECURE_SHELL WWW_HTTP WWW_HTTPS WORKGROUP_MANAGER XGRID OTHER"`

	// FromVlan
	// From VLAN
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	FromVlan *int `json:"fromVlan" validate:"required,gte=1,lte=4094"`

	// Notes
	// Notes
	Notes *string `json:"notes,omitempty"`

	// Protocol
	// protocol. This is only available when bridgeService is OTHER
	Protocol *string `json:"protocol,omitempty"`

	// ToVlan
	// To VLAN
	// Constraints:
	//    - required
	//    - min:1
	//    - max:4094
	ToVlan *int `json:"toVlan" validate:"required,gte=1,lte=4094"`
}

func NewBonjourPolicyRule() *BonjourPolicyRule {
	bonjourPolicyRuleType := new(BonjourPolicyRule)
	return bonjourPolicyRuleType
}

func NewBonjourPolicyRuleWithDefaults() *BonjourPolicyRule {
	bonjourPolicyRuleType := new(BonjourPolicyRule)
	return bonjourPolicyRuleType
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

func NewBonjourPolicyRuleConfiguration() *BonjourPolicyRuleConfiguration {
	bonjourPolicyRuleConfigurationType := new(BonjourPolicyRuleConfiguration)
	return bonjourPolicyRuleConfigurationType
}

func NewBonjourPolicyRuleConfigurationWithDefaults() *BonjourPolicyRuleConfiguration {
	bonjourPolicyRuleConfigurationType := new(BonjourPolicyRuleConfiguration)
	return bonjourPolicyRuleConfigurationType
}

type ClientLoadBalancing struct {
	// AdjacentRadioThreshold
	// Adjacent radio threshold
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AdjacentRadioThreshold *int `json:"adjacentRadioThreshold,omitempty" validate:"omitempty,gte=1,lte=100"`
}

func NewClientLoadBalancing() *ClientLoadBalancing {
	clientLoadBalancingType := new(ClientLoadBalancing)
	return clientLoadBalancingType
}

func NewClientLoadBalancingWithDefaults() *ClientLoadBalancing {
	clientLoadBalancingType := new(ClientLoadBalancing)
	return clientLoadBalancingType
}

type CreateBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*BonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewCreateBonjourGatewayPolicy() *CreateBonjourGatewayPolicy {
	createBonjourGatewayPolicyType := new(CreateBonjourGatewayPolicy)
	return createBonjourGatewayPolicyType
}

func NewCreateBonjourGatewayPolicyWithDefaults() *CreateBonjourGatewayPolicy {
	createBonjourGatewayPolicyType := new(CreateBonjourGatewayPolicy)
	return createBonjourGatewayPolicyType
}

type CreateDiffServProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *UplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewCreateDiffServProfile() *CreateDiffServProfile {
	createDiffServProfileType := new(CreateDiffServProfile)
	return createDiffServProfileType
}

func NewCreateDiffServProfileWithDefaults() *CreateDiffServProfile {
	createDiffServProfileType := new(CreateDiffServProfile)
	return createDiffServProfileType
}

type CreateZone struct {
	Altitude *common.Altitude `json:"altitude,omitempty"`

	// ApHccdEnabled
	// Historical Connection Failures allows the AP to report historical client connection failures for this zone.
	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	// ApHccdPersist
	// Allow Historical Connection Failures to be persisted.
	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

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

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
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

	Description *common.Description `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code.
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DhcpSiteConfig *common.DhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	DomainId *string `json:"domainId,omitempty"`

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

	// EnforcePriorityZoneAffinityEnable
	// Enforced the priority of Affinity Profile.
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

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

	// IpsecTunnelMode
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty" validate:"omitempty,oneof=DISABLE SOFT_GRE RUCKUS_GRE"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	// LoadBalancingMethod
	// Constraints:
	//    - nullable
	//    - default:'BASED_ON_CLIENT_COUNT'
	//    - oneof:[BASED_ON_CLIENT_COUNT,BASED_ON_CAPACITY,OFF]
	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty" validate:"omitempty,oneof=BASED_ON_CLIENT_COUNT BASED_ON_CAPACITY OFF"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.GenericRef `json:"locationBasedService,omitempty"`

	// Login
	// Constraints:
	//    - required
	Login *ApLogin `json:"login" validate:"required"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*common.LteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mesh *MeshConfiguration `json:"mesh,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	NodeAffinityProfile *common.GenericRef `json:"nodeAffinityProfile,omitempty"`

	ProtectionMode24 *common.ProtectionMode `json:"protectionMode24,omitempty"`

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

	SnmpAgent *ApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*SoftGreRef `json:"softGreTunnelProflies,omitempty"`

	// SshTunnelEncryption
	// Constraints:
	//    - nullable
	//    - default:'AES128'
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty" validate:"omitempty,oneof=AES128 AES256"`

	Syslog *Syslog `json:"syslog,omitempty"`

	Timezone *TimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *common.ZoneTunnelType `json:"tunnelType,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Version *common.FirmwareVersion `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *common.Radio24 `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50 `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// Zone affinity profile of the zone
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewCreateZone() *CreateZone {
	createZoneType := new(CreateZone)
	return createZoneType
}

func NewCreateZoneWithDefaults() *CreateZone {
	createZoneType := new(CreateZone)
	apHccdEnabledField := false
	createZoneType.ApHccdEnabled = &apHccdEnabledField
	apHccdPersistField := false
	createZoneType.ApHccdPersist = &apHccdPersistField
	cbandChannelEnabledField := false
	createZoneType.CbandChannelEnabled = &cbandChannelEnabledField
	cbandChannelLicenseEnabledField := false
	createZoneType.CbandChannelLicenseEnabled = &cbandChannelLicenseEnabledField
	channel144EnabledField := false
	createZoneType.Channel144Enabled = &channel144EnabledField
	channelEvaluationIntervalField := 600
	createZoneType.ChannelEvaluationInterval = &channelEvaluationIntervalField
	channelModeEnabledField := false
	createZoneType.ChannelModeEnabled = &channelModeEnabledField
	clusterRedundancyEnabledField := false
	createZoneType.ClusterRedundancyEnabled = &clusterRedundancyEnabledField
	dfsChannelEnabledField := false
	createZoneType.DfsChannelEnabled = &dfsChannelEnabledField
	directedMulticastFromNetworkEnabledField := false
	createZoneType.DirectedMulticastFromNetworkEnabled = &directedMulticastFromNetworkEnabledField
	directedMulticastFromWiredClientEnabledField := false
	createZoneType.DirectedMulticastFromWiredClientEnabled = &directedMulticastFromWiredClientEnabledField
	directedMulticastFromWirelessClientEnabledField := false
	createZoneType.DirectedMulticastFromWirelessClientEnabled = &directedMulticastFromWirelessClientEnabledField
	dosBarringCheckPeriodField := 30
	createZoneType.DosBarringCheckPeriod = &dosBarringCheckPeriodField
	dosBarringPeriodField := 60
	createZoneType.DosBarringPeriod = &dosBarringPeriodField
	dosBarringThresholdField := 5
	createZoneType.DosBarringThreshold = &dosBarringThresholdField
	healthCheckSitesEnabledField := false
	createZoneType.HealthCheckSitesEnabled = &healthCheckSitesEnabledField
	loadBalancingMethodField := `BASED_ON_CLIENT_COUNT`
	createZoneType.LoadBalancingMethod = &loadBalancingMethodField
	rogueApJammingDetectionField := false
	createZoneType.RogueApJammingDetection = &rogueApJammingDetectionField
	sshTunnelEncryptionField := `AES128`
	createZoneType.SshTunnelEncryption = &sshTunnelEncryptionField
	return createZoneType
}

type CustomizedTimeZone struct {
	// Abbreviation
	// Time zone abbreviation
	// Constraints:
	//    - required
	Abbreviation *string `json:"abbreviation" validate:"required"`

	End *DaylightSavingTime `json:"end,omitempty"`

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

	Start *DaylightSavingTime `json:"start,omitempty"`
}

func NewCustomizedTimeZone() *CustomizedTimeZone {
	customizedTimeZoneType := new(CustomizedTimeZone)
	return customizedTimeZoneType
}

func NewCustomizedTimeZoneWithDefaults() *CustomizedTimeZone {
	customizedTimeZoneType := new(CustomizedTimeZone)
	return customizedTimeZoneType
}

type DaylightSavingTime struct {
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

func NewDaylightSavingTime() *DaylightSavingTime {
	daylightSavingTimeType := new(DaylightSavingTime)
	return daylightSavingTimeType
}

func NewDaylightSavingTimeWithDefaults() *DaylightSavingTime {
	daylightSavingTimeType := new(DaylightSavingTime)
	return daylightSavingTimeType
}

type DhcpSiteConfigList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*common.DhcpSiteConfigListRef `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDhcpSiteConfigList() *DhcpSiteConfigList {
	dhcpSiteConfigListType := new(DhcpSiteConfigList)
	return dhcpSiteConfigListType
}

func NewDhcpSiteConfigListWithDefaults() *DhcpSiteConfigList {
	dhcpSiteConfigListType := new(DhcpSiteConfigList)
	return dhcpSiteConfigListType
}

type DiffServConfiguration struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Id
	// Identifier of the zone
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *UplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewDiffServConfiguration() *DiffServConfiguration {
	diffServConfigurationType := new(DiffServConfiguration)
	return diffServConfigurationType
}

func NewDiffServConfigurationWithDefaults() *DiffServConfiguration {
	diffServConfigurationType := new(DiffServConfiguration)
	return diffServConfigurationType
}

type DiffServList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DiffServSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDiffServList() *DiffServList {
	diffServListType := new(DiffServList)
	return diffServListType
}

func NewDiffServListWithDefaults() *DiffServList {
	diffServListType := new(DiffServList)
	return diffServListType
}

type DiffServSummary struct {
	// Id
	// Identifier of the diff serv
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the diff serv
	Name *string `json:"name,omitempty"`
}

func NewDiffServSummary() *DiffServSummary {
	diffServSummaryType := new(DiffServSummary)
	return diffServSummaryType
}

func NewDiffServSummaryWithDefaults() *DiffServSummary {
	diffServSummaryType := new(DiffServSummary)
	return diffServSummaryType
}

type DownlinkDiffServ struct {
	// Downlink
	// Downlink
	Downlink *string `json:"downlink,omitempty"`

	// DownlinkEnable
	// Downlink enable
	DownlinkEnable *bool `json:"downlinkEnable,omitempty"`
}

func NewDownlinkDiffServ() *DownlinkDiffServ {
	downlinkDiffServType := new(DownlinkDiffServ)
	return downlinkDiffServType
}

func NewDownlinkDiffServWithDefaults() *DownlinkDiffServ {
	downlinkDiffServType := new(DownlinkDiffServ)
	return downlinkDiffServType
}

type MeshConfiguration struct {
	// MeshRadioIdx
	// Mesh radio index
	// Constraints:
	//    - nullable
	//    - default:'Radio5G'
	//    - oneof:[Radio24G,Radio5G]
	MeshRadioIdx *string `json:"meshRadioIdx,omitempty" validate:"omitempty,oneof=Radio24G Radio5G"`

	// Passphrase
	// Passphrase for the mesh network
	Passphrase *string `json:"passphrase,omitempty"`

	// Ssid
	// SSID of the mesh network
	Ssid *string `json:"ssid,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

func NewMeshConfiguration() *MeshConfiguration {
	meshConfigurationType := new(MeshConfiguration)
	return meshConfigurationType
}

func NewMeshConfigurationWithDefaults() *MeshConfiguration {
	meshConfigurationType := new(MeshConfiguration)
	meshRadioIdxField := `Radio5G`
	meshConfigurationType.MeshRadioIdx = &meshRadioIdxField
	return meshConfigurationType
}

type ModfiyApFirmware struct {
	// FirmwareVersion
	// new version of the AP firmare
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
}

func NewModfiyApFirmware() *ModfiyApFirmware {
	modfiyApFirmwareType := new(ModfiyApFirmware)
	return modfiyApFirmwareType
}

func NewModfiyApFirmwareWithDefaults() *ModfiyApFirmware {
	modfiyApFirmwareType := new(ModfiyApFirmware)
	return modfiyApFirmwareType
}

type ModifyBonjourGatewayPolicy struct {
	// BonjourPolicyRuleList
	// Bonjour policy rule list
	BonjourPolicyRuleList []*BonjourPolicyRule `json:"bonjourPolicyRuleList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewModifyBonjourGatewayPolicy() *ModifyBonjourGatewayPolicy {
	modifyBonjourGatewayPolicyType := new(ModifyBonjourGatewayPolicy)
	return modifyBonjourGatewayPolicyType
}

func NewModifyBonjourGatewayPolicyWithDefaults() *ModifyBonjourGatewayPolicy {
	modifyBonjourGatewayPolicyType := new(ModifyBonjourGatewayPolicy)
	return modifyBonjourGatewayPolicyType
}

type ModifyDiffServProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkDiffServ *DownlinkDiffServ `json:"downlinkDiffServ,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PreservedList
	// Preserved list
	PreservedList []string `json:"preservedList,omitempty"`

	UplinkDiffServ *UplinkDiffServ `json:"uplinkDiffServ,omitempty"`
}

func NewModifyDiffServProfile() *ModifyDiffServProfile {
	modifyDiffServProfileType := new(ModifyDiffServProfile)
	return modifyDiffServProfileType
}

func NewModifyDiffServProfileWithDefaults() *ModifyDiffServProfile {
	modifyDiffServProfileType := new(ModifyDiffServProfile)
	return modifyDiffServProfileType
}

type ModifyZone struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

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

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
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

	Description *common.Description `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DhcpSiteConfig *common.DhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	DomainId *string `json:"domainId,omitempty"`

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

	// EnforcePriorityZoneAffinityEnable
	// Enforce the priority of zone affinity
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

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

	// IpsecTunnelMode
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty" validate:"omitempty,oneof=DISABLE SOFT_GRE RUCKUS_GRE"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	// LoadBalancingMethod
	// Constraints:
	//    - nullable
	//    - default:'BASED_ON_CLIENT_COUNT'
	//    - oneof:[BASED_ON_CLIENT_COUNT,BASED_ON_CAPACITY,OFF]
	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty" validate:"omitempty,oneof=BASED_ON_CLIENT_COUNT BASED_ON_CAPACITY OFF"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.GenericRef `json:"locationBasedService,omitempty"`

	Login *ApLogin `json:"login,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*common.LteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mesh *MeshConfiguration `json:"mesh,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	NodeAffinityProfile *common.GenericRef `json:"nodeAffinityProfile,omitempty"`

	ProtectionMode24 *common.ProtectionMode `json:"protectionMode24,omitempty"`

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

	SnmpAgent *ApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*SoftGreRef `json:"softGreTunnelProflies,omitempty"`

	// SshTunnelEncryption
	// Constraints:
	//    - nullable
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty" validate:"omitempty,oneof=AES128 AES256"`

	Syslog *Syslog `json:"syslog,omitempty"`

	Timezone *TimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *common.ZoneTunnelType `json:"tunnelType,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *common.Radio24 `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50 `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// zone affinity profile Id
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewModifyZone() *ModifyZone {
	modifyZoneType := new(ModifyZone)
	return modifyZoneType
}

func NewModifyZoneWithDefaults() *ModifyZone {
	modifyZoneType := new(ModifyZone)
	apHccdEnabledField := false
	modifyZoneType.ApHccdEnabled = &apHccdEnabledField
	apHccdPersistField := false
	modifyZoneType.ApHccdPersist = &apHccdPersistField
	channelEvaluationIntervalField := 600
	modifyZoneType.ChannelEvaluationInterval = &channelEvaluationIntervalField
	healthCheckSitesEnabledField := false
	modifyZoneType.HealthCheckSitesEnabled = &healthCheckSitesEnabledField
	loadBalancingMethodField := `BASED_ON_CLIENT_COUNT`
	modifyZoneType.LoadBalancingMethod = &loadBalancingMethodField
	return modifyZoneType
}

type QueryCriteria struct{}

func NewQueryCriteria() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
}

func NewQueryCriteriaWithDefaults() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
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
	// Constraints:
	//    - nullable
	//    - oneof:[All,Malicious]
	ReportType *string `json:"reportType,omitempty" validate:"omitempty,oneof=All Malicious"`
}

func NewRogue() *Rogue {
	rogueType := new(Rogue)
	return rogueType
}

func NewRogueWithDefaults() *Rogue {
	rogueType := new(Rogue)
	return rogueType
}

type SnmpUser struct {
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
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP User
	NotificationTarget []*common.TargetConfig `json:"notificationTarget,omitempty"`

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
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// UserName
	// name of the SNMP User.
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`

	// WriteEnabled
	// write privilege of the SNMP User
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

func NewSnmpUser() *SnmpUser {
	snmpUserType := new(SnmpUser)
	return snmpUserType
}

func NewSnmpUserWithDefaults() *SnmpUser {
	snmpUserType := new(SnmpUser)
	return snmpUserType
}

type SoftGreRef struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSoftGreRef() *SoftGreRef {
	softGreRefType := new(SoftGreRef)
	return softGreRefType
}

func NewSoftGreRefWithDefaults() *SoftGreRef {
	softGreRefType := new(SoftGreRef)
	return softGreRefType
}

type Syslog struct {
	Address *common.IpAddress `json:"address,omitempty"`

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

	SecondaryAddress *common.IpAddress `json:"secondaryAddress,omitempty"`

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

func NewSyslog() *Syslog {
	syslogType := new(Syslog)
	return syslogType
}

func NewSyslogWithDefaults() *Syslog {
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

type TimezoneSetting struct {
	CustomizedTimezone *CustomizedTimeZone `json:"customizedTimezone,omitempty"`

	// SystemTimezone
	// System defined time zone, please refer to the “Overview > Time Zone” list
	SystemTimezone *string `json:"systemTimezone,omitempty"`
}

func NewTimezoneSetting() *TimezoneSetting {
	timezoneSettingType := new(TimezoneSetting)
	return timezoneSettingType
}

func NewTimezoneSettingWithDefaults() *TimezoneSetting {
	timezoneSettingType := new(TimezoneSetting)
	return timezoneSettingType
}

type UnsupportedApModel struct {
	// Amount
	// amount of the AP Model
	Amount *int `json:"amount,omitempty"`

	// Model
	// name of the AP Model
	Model *string `json:"model,omitempty"`
}

func NewUnsupportedApModel() *UnsupportedApModel {
	unsupportedApModelType := new(UnsupportedApModel)
	return unsupportedApModelType
}

func NewUnsupportedApModelWithDefaults() *UnsupportedApModel {
	unsupportedApModelType := new(UnsupportedApModel)
	return unsupportedApModelType
}

type UplinkDiffServ struct {
	// Uplink
	// Uplink
	Uplink *string `json:"uplink,omitempty"`

	// UplinkEnable
	// Uplink enable
	UplinkEnable *bool `json:"uplinkEnable,omitempty"`
}

func NewUplinkDiffServ() *UplinkDiffServ {
	uplinkDiffServType := new(UplinkDiffServ)
	return uplinkDiffServType
}

func NewUplinkDiffServWithDefaults() *UplinkDiffServ {
	uplinkDiffServType := new(UplinkDiffServ)
	return uplinkDiffServType
}

type ZoneConfiguration struct {
	AaaAffinityEnabled *bool `json:"aaaAffinityEnabled,omitempty"`

	Altitude *common.Altitude `json:"altitude,omitempty"`

	ApHccdEnabled *bool `json:"apHccdEnabled,omitempty"`

	ApHccdPersist *bool `json:"apHccdPersist,omitempty"`

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

	// Channel144Enabled
	// Channel 144 enabled configuration of the zone.
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

	Description *common.Description `json:"description,omitempty"`

	// DfsChannelEnabled
	// DFS Channel enabled configuration of the zone, only for the US country code .
	DfsChannelEnabled *bool `json:"dfsChannelEnabled,omitempty"`

	DhcpSiteConfig *common.DhcpSiteConfigRef `json:"dhcpSiteConfig,omitempty"`

	// DirectedMulticastFromNetworkEnabled
	// Directed multicast from network.
	DirectedMulticastFromNetworkEnabled *bool `json:"directedMulticastFromNetworkEnabled,omitempty"`

	// DirectedMulticastFromWiredClientEnabled
	// Directed multicast from wired client.
	DirectedMulticastFromWiredClientEnabled *bool `json:"directedMulticastFromWiredClientEnabled,omitempty"`

	// DirectedMulticastFromWirelessClientEnabled
	// Directed multicast from wireless client.
	DirectedMulticastFromWirelessClientEnabled *bool `json:"directedMulticastFromWirelessClientEnabled,omitempty"`

	// DomainId
	// Identifier of the management domain to which the zone belongs
	DomainId *string `json:"domainId,omitempty"`

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

	// EnforcePriorityZoneAffinityEnable
	// Enforce the priority of zone affinity
	EnforcePriorityZoneAffinityEnable *bool `json:"enforcePriorityZoneAffinityEnable,omitempty"`

	// HealthCheckSites
	// Health Check Sites.
	HealthCheckSites []string `json:"healthCheckSites,omitempty"`

	// HealthCheckSitesEnabled
	// Enabled Health Check Sites.
	HealthCheckSitesEnabled *bool `json:"healthCheckSitesEnabled,omitempty"`

	// Id
	// Identifier of the zone
	Id *string `json:"id,omitempty"`

	IpMode *common.IpMode `json:"ipMode,omitempty"`

	IpsecProfile *common.GenericRef `json:"ipsecProfile,omitempty"`

	// IpsecProfiles
	// Ipsec profile for Multiple Tunnel (Start from SZ 5.0)
	IpsecProfiles []*common.GenericRef `json:"ipsecProfiles,omitempty"`

	// IpsecTunnelMode
	// Constraints:
	//    - nullable
	//    - oneof:[DISABLE,SOFT_GRE,RUCKUS_GRE]
	IpsecTunnelMode *string `json:"ipsecTunnelMode,omitempty" validate:"omitempty,oneof=DISABLE SOFT_GRE RUCKUS_GRE"`

	// Ipv6TrafficFilterEnabled
	// IPv6 Traffic filtering on the AP
	Ipv6TrafficFilterEnabled *int `json:"ipv6TrafficFilterEnabled,omitempty"`

	Latitude *common.Latitude `json:"latitude,omitempty"`

	// LoadBalancingMethod
	// Constraints:
	//    - nullable
	//    - default:'BASED_ON_CLIENT_COUNT'
	//    - oneof:[BASED_ON_CLIENT_COUNT,BASED_ON_CAPACITY,OFF]
	LoadBalancingMethod *string `json:"loadBalancingMethod,omitempty" validate:"omitempty,oneof=BASED_ON_CLIENT_COUNT BASED_ON_CAPACITY OFF"`

	Location *common.Location `json:"location,omitempty"`

	LocationAdditionalInfo *common.LocationAdditionalInfo `json:"locationAdditionalInfo,omitempty"`

	LocationBasedService *common.GenericRef `json:"locationBasedService,omitempty"`

	Login *ApLogin `json:"login,omitempty"`

	Longitude *common.Longitude `json:"longitude,omitempty"`

	// LteBandLockChannels
	// LTE band lock channels options
	LteBandLockChannels []*common.LteBandLockChannel `json:"lteBandLockChannels,omitempty"`

	Mesh *MeshConfiguration `json:"mesh,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	NodeAffinityProfile *common.GenericRef `json:"nodeAffinityProfile,omitempty"`

	ProtectionMode24 *common.ProtectionMode `json:"protectionMode24,omitempty"`

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

	SnmpAgent *ApSnmpOptions `json:"snmpAgent,omitempty"`

	// SoftGreTunnelProflies
	// SoftGRE Profiles for Multiple Tunnel (Start from SZ 5.0)
	SoftGreTunnelProflies []*SoftGreRef `json:"softGreTunnelProflies,omitempty"`

	// SshTunnelEncryption
	// Constraints:
	//    - nullable
	//    - oneof:[AES128,AES256]
	SshTunnelEncryption *string `json:"sshTunnelEncryption,omitempty" validate:"omitempty,oneof=AES128 AES256"`

	Syslog *Syslog `json:"syslog,omitempty"`

	Timezone *TimezoneSetting `json:"timezone,omitempty"`

	TunnelProfile *common.GenericRef `json:"tunnelProfile,omitempty"`

	TunnelType *common.ZoneTunnelType `json:"tunnelType,omitempty"`

	VenueProfile *common.GenericRef `json:"venueProfile,omitempty"`

	Version *common.FirmwareVersion `json:"version,omitempty"`

	// VlanOverlappingEnabled
	// VLAN pooling overlapping of the zone
	VlanOverlappingEnabled *bool `json:"vlanOverlappingEnabled,omitempty"`

	Wifi24 *common.Radio24SuperSet `json:"wifi24,omitempty"`

	Wifi50 *common.Radio50SuperSet `json:"wifi50,omitempty"`

	// ZoneAffinityProfileId
	// zone affinity profile Id
	ZoneAffinityProfileId *string `json:"zoneAffinityProfileId,omitempty"`
}

func NewZoneConfiguration() *ZoneConfiguration {
	zoneConfigurationType := new(ZoneConfiguration)
	return zoneConfigurationType
}

func NewZoneConfigurationWithDefaults() *ZoneConfiguration {
	zoneConfigurationType := new(ZoneConfiguration)
	apHccdEnabledField := false
	zoneConfigurationType.ApHccdEnabled = &apHccdEnabledField
	apHccdPersistField := false
	zoneConfigurationType.ApHccdPersist = &apHccdPersistField
	channelEvaluationIntervalField := 600
	zoneConfigurationType.ChannelEvaluationInterval = &channelEvaluationIntervalField
	healthCheckSitesEnabledField := false
	zoneConfigurationType.HealthCheckSitesEnabled = &healthCheckSitesEnabledField
	loadBalancingMethodField := `BASED_ON_CLIENT_COUNT`
	zoneConfigurationType.LoadBalancingMethod = &loadBalancingMethodField
	return zoneConfigurationType
}

type ZoneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ZoneSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewZoneList() *ZoneList {
	zoneListType := new(ZoneList)
	return zoneListType
}

func NewZoneListWithDefaults() *ZoneList {
	zoneListType := new(ZoneList)
	return zoneListType
}

type ZoneSummary struct {
	// Id
	// Identifier of the zone
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the zone
	Name *string `json:"name,omitempty"`

	// ServiceName
	// Name of the zone
	ServiceName *string `json:"serviceName,omitempty"`
}

func NewZoneSummary() *ZoneSummary {
	zoneSummaryType := new(ZoneSummary)
	return zoneSummaryType
}

func NewZoneSummaryWithDefaults() *ZoneSummary {
	zoneSummaryType := new(ZoneSummary)
	return zoneSummaryType
}
