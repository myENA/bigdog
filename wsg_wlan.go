package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type WSGWLANService struct {
	apiClient *VSZClient
}

func NewWSGWLANService(c *VSZClient) *WSGWLANService {
	s := new(WSGWLANService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWLANService() *WSGWLANService {
	return NewWSGWLANService(ss.apiClient)
}

// WSGWLANCreateGuestAccessWlan
//
// Definition: wlan_createGuestAccessWlan
type WSGWLANCreateGuestAccessWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateGuestAccessWlan() *WSGWLANCreateGuestAccessWlan {
	m := new(WSGWLANCreateGuestAccessWlan)
	return m
}

// WSGWLANCreateHotspot20OpenWlan
//
// Definition: wlan_createHotspot20OpenWlan
type WSGWLANCreateHotspot20OpenWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateHotspot20OpenWlan() *WSGWLANCreateHotspot20OpenWlan {
	m := new(WSGWLANCreateHotspot20OpenWlan)
	return m
}

// WSGWLANCreateHotspot20Wlan
//
// Definition: wlan_createHotspot20Wlan
type WSGWLANCreateHotspot20Wlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	// Hotspot20Profile
	// Constraints:
	//    - required
	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateHotspot20Wlan() *WSGWLANCreateHotspot20Wlan {
	m := new(WSGWLANCreateHotspot20Wlan)
	return m
}

// WSGWLANCreateHotspotWlan
//
// Definition: wlan_createHotspotWlan
type WSGWLANCreateHotspotWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateHotspotWlan() *WSGWLANCreateHotspotWlan {
	m := new(WSGWLANCreateHotspotWlan)
	return m
}

// WSGWLANCreateStandard80211Wlan
//
// Definition: wlan_createStandard80211Wlan
type WSGWLANCreateStandard80211Wlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable  for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable  for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateStandard80211Wlan() *WSGWLANCreateStandard80211Wlan {
	m := new(WSGWLANCreateStandard80211Wlan)
	return m
}

// WSGWLANCreateStandardOpenWlan
//
// Definition: wlan_createStandardOpenWlan
type WSGWLANCreateStandardOpenWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - default:'APLBO'
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateStandardOpenWlan() *WSGWLANCreateStandardOpenWlan {
	m := new(WSGWLANCreateStandardOpenWlan)
	return m
}

// WSGWLANCreateWebAuthWlan
//
// Definition: wlan_createWebAuthWlan
type WSGWLANCreateWebAuthWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateWebAuthWlan() *WSGWLANCreateWebAuthWlan {
	m := new(WSGWLANCreateWebAuthWlan)
	return m
}

// WSGWLANCreateWechatWlan
//
// Definition: wlan_createWechatWlan
type WSGWLANCreateWechatWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWLANNameSSID `json:"name"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WSGWLANNameSSID `json:"ssid"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANCreateWechatWlan() *WSGWLANCreateWechatWlan {
	m := new(WSGWLANCreateWechatWlan)
	return m
}

// WSGWLANModifyWlan
//
// Definition: wlan_modifyWlan
type WSGWLANModifyWlan struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	Name *WSGWLANNameSSID `json:"name,omitempty"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WSGWLANNameSSID `json:"ssid,omitempty"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`
}

func NewWSGWLANModifyWlan() *WSGWLANModifyWlan {
	m := new(WSGWLANModifyWlan)
	return m
}

// WSGWLANAccounting
//
// Definition: wlan_wlanAccounting
type WSGWLANAccounting struct {
	// AccountingDelayEnabled
	// Indicates whether accounting delay time is enabled
	AccountingDelayEnabled *bool `json:"accountingDelayEnabled,omitempty"`

	// BackupAccountingId
	// Backup accounting service or profile ID. At least one backupAccountingId or backupAccountingName is required in the request when setting backup accounting service.
	BackupAccountingId *string `json:"backupAccountingId,omitempty"`

	// BackupAccountingName
	// Backup accounting service or profile name. At least one backupAccountingId or backupAccountingName is required in the request when setting backup accounting service.
	BackupAccountingName *string `json:"backupAccountingName,omitempty"`

	// Id
	// Accounting service or profile ID. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// InterimUpdateMin
	// Interval (in minutes) for sending interim updates
	// Constraints:
	//    - min:0
	//    - max:1440
	InterimUpdateMin *int `json:"interimUpdateMin,omitempty"`

	// Name
	// Accounting service or profile name. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	RealmBasedAcct *bool `json:"realmBasedAcct,omitempty"`

	// ThroughController
	// Indicates whether accounting messages were sent through the controller
	ThroughController *bool `json:"throughController,omitempty"`
}

func NewWSGWLANAccounting() *WSGWLANAccounting {
	m := new(WSGWLANAccounting)
	return m
}

// WSGWLANAdvanced
//
// Definition: wlan_wlanAdvanced
type WSGWLANAdvanced struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	// AssocRssiThr
	// Assoc RSSI threshold.
	// Constraints:
	//    - min:-90
	//    - max:-60
	AssocRssiThr *int `json:"assocRssiThr,omitempty"`

	// AuthRssiThr
	// Auth RSSI threshold.
	// Constraints:
	//    - min:-90
	//    - max:-60
	AuthRssiThr *int `json:"authRssiThr,omitempty"`

	// AvcEnabled
	// Indicator of whether AVC support is enabled or disabled
	AvcEnabled *bool `json:"avcEnabled,omitempty"`

	// BandBalancing
	// Indicates whether band balancing is enabled or disabled
	// Constraints:
	//    - default:'UseZoneSetting'
	//    - oneof:[Disabled,UseZoneSetting]
	BandBalancing *string `json:"bandBalancing,omitempty"`

	BssMinRateMbps *WSGWLANBssMinRateMbps `json:"bssMinRateMbps,omitempty"`

	// ClientFingerprintingEnabled
	// Indicates whether client fingerprinting is enabled or disabled
	ClientFingerprintingEnabled *bool `json:"clientFingerprintingEnabled,omitempty"`

	// ClientIdleTimeoutSec
	// Client idle timeout in seconds
	// Constraints:
	//    - default:120
	//    - min:60
	//    - max:86400
	ClientIdleTimeoutSec *int `json:"clientIdleTimeoutSec,omitempty"`

	// ClientIsolationAutoVrrpEnabled
	// Indicates whether Automatic support for VRRP of wireless client isolation is enabled or disabled
	ClientIsolationAutoVrrpEnabled *bool `json:"clientIsolationAutoVrrpEnabled,omitempty"`

	// ClientIsolationEnabled
	// Indicates whether wireless client isolation is enabled or disabled
	ClientIsolationEnabled *bool `json:"clientIsolationEnabled,omitempty"`

	// ClientIsolationMulticastEnabled
	// Indicates whether isolate multicast of wireless client isolation is enabled or disabled
	ClientIsolationMulticastEnabled *bool `json:"clientIsolationMulticastEnabled,omitempty"`

	// ClientIsolationUnicastEnabled
	// Indicates whether isolate unicast of wireless client isolation is enabled or disabled
	ClientIsolationUnicastEnabled *bool `json:"clientIsolationUnicastEnabled,omitempty"`

	ClientIsolationWhitelist *WSGCommonGenericRef `json:"clientIsolationWhitelist,omitempty"`

	// ClientLoadBalancingEnabled
	// Indicates whether Client Load Balancing is enabled or disabled
	ClientLoadBalancingEnabled *bool `json:"clientLoadBalancingEnabled,omitempty"`

	// DgafEnabled
	// Indicates whether dgaf is enabled or disabled
	DgafEnabled *bool `json:"dgafEnabled,omitempty"`

	// Dhcp82Format
	// DHCP Option 82 format. This variable no longer supports from v8_1 and only kept for backward compatibility.
	// Constraints:
	//    - oneof:[RUCKUS_DEFAULT,RADIUS_DHCP,RADIUS_NAT,RADIUS_DHCP_NAT,RADIUS_NAT_RUCKUS,RADIUS_NAT_SOFTGRE,SOFTGRE_CUSTOMIZED]
	Dhcp82Format *string `json:"dhcp82Format,omitempty"`

	// Dhcp82MacFormat
	// AP and Client Mac format. If dhcpOption82Enabled is true, you have to set the dhcp82MacFormat ["COLON","HYPHEN","NODELIMITER"].
	// Constraints:
	//    - oneof:[COLON,HYPHEN,NODELIMITER]
	Dhcp82MacFormat *string `json:"dhcp82MacFormat,omitempty"`

	// Dhcp82SubOpt1Format
	// Subopt-1 format
	// Constraints:
	//    - oneof:[NONE,SUBOPT1_AP_INFO_LOCATION,SUBOPT1_AP_INFO,SUBOPT1_AP_MAC_ESSID_PRIVACYTYPE,SUBOPT1_AP_MAC_hex,SUBOPT1_AP_MAC_hex_ESSID,SUBOPT1_ESSID,SUBOPT1_AP_MAC,SUBOPT1_AP_MAC_ESSID,SUBOPT1_AP_Name_ESSID]
	Dhcp82SubOpt1Format *string `json:"dhcp82SubOpt1Format,omitempty"`

	// Dhcp82SubOpt2Format
	// Subopt-2 format
	// Constraints:
	//    - oneof:[NONE,SUBOPT2_CLIENT_MAC,SUBOPT2_CLIENT_MAC_hex,SUBOPT2_CLIENT_MAC_hex_ESSID,SUBOPT2_AP_MAC,SUBOPT2_AP_MAC_hex,SUBOPT2_AP_MAC_hex_ESSID,SUBOPT2_AP_MAC_ESSID,SUBOPT2_AP_Name]
	Dhcp82SubOpt2Format *string `json:"dhcp82SubOpt2Format,omitempty"`

	// Dhcp82SubOpt150Format
	// Subopt-150 with VLAN-Id
	// Constraints:
	//    - oneof:[NONE,SUBOPT150_VLAN_ID]
	Dhcp82SubOpt150Format *string `json:"dhcp82SubOpt150Format,omitempty"`

	// Dhcp82SubOpt151AreaName
	// Subopt-151 Area Name value
	Dhcp82SubOpt151AreaName *string `json:"dhcp82SubOpt151AreaName,omitempty"`

	// Dhcp82SubOpt151Format
	// Subopt-151 format
	// Constraints:
	//    - oneof:[NONE,SUBOPT151_AREA_NAME,SUBOPT151_ESSID]
	Dhcp82SubOpt151Format *string `json:"dhcp82SubOpt151Format,omitempty"`

	// Dhcp82SubOptRadiusFormat
	// Constraints:
	//    - oneof:[NONE,RUCKUS_DEFAULT,RADIUS_DHCP,RADIUS_NAT,RADIUS_DHCP_NAT,RADIUS_NAT_RUCKUS,RADIUS_NAT_SOFTGRE,SOFTGRE_CUSTOMIZED]
	Dhcp82SubOptRadiusFormat *string `json:"dhcp82SubOptRadiusFormat,omitempty"`

	// DhcpOption82Enabled
	// Indicates whether DCHP Option 82 is enabled or disabled. This variable no longer supports from v8_1 and only kept for backward compatibility.
	DhcpOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	// DirectedThreshold
	// Directed Threshold Setting, Defines the client count at which an AP will stop converting group addressed data traffic to unicast.
	// Constraints:
	//    - default:5
	//    - min:0
	//    - max:128
	DirectedThreshold *int `json:"directedThreshold,omitempty"`

	// DnsSpoofingProfileId
	// DNS Spoofing Profile ID
	DnsSpoofingProfileId *string `json:"dnsSpoofingProfileId,omitempty"`

	// DownlinkEnabled
	// SSID Rate Limiting downlink enabled.
	DownlinkEnabled *bool `json:"downlinkEnabled,omitempty"`

	// DownlinkRate
	// SSID Rate Limiting downlink.
	DownlinkRate *float64 `json:"downlinkRate,omitempty"`

	// DropRandomProbesEnabled
	// Drop Random Probes enabled.
	DropRandomProbesEnabled *bool `json:"dropRandomProbesEnabled,omitempty"`

	// DtimInterval
	// DTIM Interval
	// Constraints:
	//    - default:1
	//    - min:1
	//    - max:255
	DtimInterval *int `json:"dtimInterval,omitempty"`

	// FlowLogEnabled
	// Flow log enabled.
	FlowLogEnabled *bool `json:"flowLogEnabled,omitempty"`

	// ForceClientDHCPTimeoutSec
	// Force DHCP disconnects the client if the client does not obtain a valid IP address within the timeout peroid. To disable force DHCP, set this value to zero (0).
	// Constraints:
	//    - default:0
	//    - oneof:[0,5,6,7,8,9,10,11,12,13,14,15]
	ForceClientDHCPTimeoutSec *int `json:"forceClientDHCPTimeoutSec,omitempty"`

	// HdOverheadOptimizeEnable
	// Airtime decongestion enabled.
	HdOverheadOptimizeEnable *bool `json:"hdOverheadOptimizeEnable,omitempty"`

	// HideSsidEnabled
	// Indicates whether the SSID is hidden or broadcast
	HideSsidEnabled *bool `json:"hideSsidEnabled,omitempty"`

	// Hs20Onboarding
	// Allow WISPr WLAN for Hotspot 2.0 Onboarding
	Hs20Onboarding *bool `json:"hs20Onboarding,omitempty"`

	// JoinAcceptTimeout
	// Join expire time.
	// Constraints:
	//    - default:300
	//    - min:1
	//    - max:300
	JoinAcceptTimeout *int `json:"joinAcceptTimeout,omitempty"`

	// JoinIgnoreThr
	// Join wait threshold.
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:50
	JoinIgnoreThr *int `json:"joinIgnoreThr,omitempty"`

	// JoinIgnoreTimeout
	// Join wait time.
	// Constraints:
	//    - default:30
	//    - min:1
	//    - max:60
	JoinIgnoreTimeout *int `json:"joinIgnoreTimeout,omitempty"`

	// MaxAllowedRA
	// Max Allowed RAs
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:1440
	MaxAllowedRA *int `json:"maxAllowedRA,omitempty"`

	// MaxClientsPerRadio
	// Maximum number of clients per radio
	// Constraints:
	//    - default:100
	//    - min:1
	//    - max:512
	MaxClientsPerRadio *int `json:"maxClientsPerRadio,omitempty"`

	MgmtTxRateMbps *WSGWLANMgmtTxRateMbps `json:"mgmtTxRateMbps,omitempty"`

	// MulticastDownlinkRateLimit
	// Multicast Rate Limiting downlink (mbps).
	// Constraints:
	//    - min:1
	//    - max:100
	MulticastDownlinkRateLimit *int `json:"multicastDownlinkRateLimit,omitempty"`

	// MulticastDownlinkRateLimitEnabled
	// Multicast Rate Limiting downlink enabled.
	MulticastDownlinkRateLimitEnabled *bool `json:"multicastDownlinkRateLimitEnabled,omitempty"`

	// MulticastFilterDrop
	// Drop the broadcast/multicast packets from associated clients.
	MulticastFilterDrop *bool `json:"multicastFilterDrop,omitempty"`

	// MulticastUplinkRateLimit
	// Multicast Rate Limiting uplink (mbps).
	// Constraints:
	//    - min:1
	//    - max:100
	MulticastUplinkRateLimit *int `json:"multicastUplinkRateLimit,omitempty"`

	// MulticastUplinkRateLimitEnabled
	// Multicast Rate Limiting uplink enabled.
	MulticastUplinkRateLimitEnabled *bool `json:"multicastUplinkRateLimitEnabled,omitempty"`

	// NdProxyEnabled
	// Indicates whether ND Proxy is enabled or disabled
	NdProxyEnabled *bool `json:"ndProxyEnabled,omitempty"`

	// OceBroadcastProbeResponseDelay
	// Broadcast probe response delay.
	// Constraints:
	//    - default:15
	//    - min:8
	//    - max:120
	OceBroadcastProbeResponseDelay *int `json:"oceBroadcastProbeResponseDelay,omitempty"`

	// OceEnabled
	// Optimized Connectivity Experience(OCE) enabled.
	OceEnabled *bool `json:"oceEnabled,omitempty"`

	// OceRssiBasedAssociationRejectionThreshold
	// RSSI-based association rejection threshold.
	// Constraints:
	//    - default:-75
	//    - min:-90
	//    - max:-60
	OceRssiBasedAssociationRejectionThreshold *int `json:"oceRssiBasedAssociationRejectionThreshold,omitempty"`

	// OfdmOnlyEnabled
	// Indicates whether OFDM only is enabled or disabled
	OfdmOnlyEnabled *bool `json:"ofdmOnlyEnabled,omitempty"`

	// OkcEnabled
	// Indicator of whether OKC support is enabled or disabled. The default value is true when the WLAN is WPA+AES non open WLAN.
	OkcEnabled *bool `json:"okcEnabled,omitempty"`

	// PmkCachingEnabled
	// Indicator of whether PKM caching support is enabled or disabled. The default value is true when the WLAN is WPA+AES non open WLAN.
	PmkCachingEnabled *bool `json:"pmkCachingEnabled,omitempty"`

	// Priority
	// Priority of the WLAN
	// Constraints:
	//    - default:'High'
	//    - oneof:[High,Low]
	Priority *string `json:"priority,omitempty"`

	// ProbeRssiThr
	// Join RSSI threshold. Value should be 0 (disabled) or between -90 and -60
	// Constraints:
	//    - nullable
	//    - default:-85
	ProbeRssiThr *int `json:"probeRssiThr,omitempty"`

	// ProxyARPEnabled
	// Indicates whether proxy ARP is enabled or disabled
	ProxyARPEnabled *bool `json:"proxyARPEnabled,omitempty"`

	// RaInterval
	// A timer that RA proxy runs and once receives unsolicited RA checks against the configured time and allow/drop RA based on next timeout
	// Constraints:
	//    - default:10
	//    - min:1
	//    - max:256
	RaInterval *int `json:"raInterval,omitempty"`

	// RaProxyEnabled
	// Indicates whether RA proxy is enabled or disabled
	RaProxyEnabled *bool `json:"raProxyEnabled,omitempty"`

	// RatePerSTADownlink
	// UE Rate Limiting downlink.
	RatePerSTADownlink *string `json:"ratePerSTADownlink,omitempty"`

	// RatePerSTAUplink
	// UE Rate Limiting uplink.
	RatePerSTAUplink *string `json:"ratePerSTAUplink,omitempty"`

	// RaThrottlingEnabled
	// Indicates whether RA Throttling is enabled or disabled
	RaThrottlingEnabled *bool `json:"raThrottlingEnabled,omitempty"`

	// RsraGuardEnabled
	// Indicates whether RS/RA Guard is enabled or disabled
	RsraGuardEnabled *bool `json:"rsraGuardEnabled,omitempty"`

	// Support80211dEnabled
	// Indicates whether support for 802.11d is enabled or disabled
	Support80211dEnabled *bool `json:"support80211dEnabled,omitempty"`

	// Support80211kEnabled
	// Indicates whether support for 802.11k is enabled or disabled
	Support80211kEnabled *bool `json:"support80211kEnabled,omitempty"`

	// SuppressNsEnabled
	// Indicates whether supperssNS is enabled or disabled
	SuppressNsEnabled *bool `json:"suppressNsEnabled,omitempty"`

	// TransientClientMgmtEnable
	// Transient Client Management enabled.
	TransientClientMgmtEnable *bool `json:"transientClientMgmtEnable,omitempty"`

	// UnauthClientStatsEnabled
	// Indicates whether to send statistics of unauthorized clients or not
	UnauthClientStatsEnabled *bool `json:"unauthClientStatsEnabled,omitempty"`

	// UplinkEnabled
	// SSID Rate Limiting uplink enabled.
	UplinkEnabled *bool `json:"uplinkEnabled,omitempty"`

	// UplinkRate
	// SSID Rate Limiting uplink.
	UplinkRate *float64 `json:"uplinkRate,omitempty"`

	// UrlFilteringPolicyEnabled
	// Indicator of whether URL Filtering is enabled or disabled
	UrlFilteringPolicyEnabled *bool `json:"urlFilteringPolicyEnabled,omitempty"`

	// UrlFilteringPolicyId
	// The URL Filtering policy ID.
	UrlFilteringPolicyId *string `json:"urlFilteringPolicyId,omitempty"`

	// Wifi6Enabled
	// Indicates whether wifi6 feature is enabled or disabled
	Wifi6Enabled *bool `json:"wifi6Enabled,omitempty"`

	// WifiCallingPolicyEnabled
	// Indicator of whether Wi-Fi Calling is enabled or disabled
	WifiCallingPolicyEnabled *bool `json:"wifiCallingPolicyEnabled,omitempty"`

	// WifiCallingPolicyIds
	// The Wi-Fi Calling policy IDs. (Maximum allowed number is 5)
	WifiCallingPolicyIds []string `json:"wifiCallingPolicyIds,omitempty"`
}

func NewWSGWLANAdvanced() *WSGWLANAdvanced {
	m := new(WSGWLANAdvanced)
	return m
}

// WSGWLANAuthentication
//
// Definition: wlan_wlanAuthentication
type WSGWLANAuthentication struct {
	// AuthenticationOption
	// Option of the authentication service or profile, At least one ID or name or authenticationOption is required in the request. This only applies to hotspot and guest WLANs.
	// Constraints:
	//    - nullable
	//    - oneof:[Local DB,Guest,Always Accept]
	AuthenticationOption *string `json:"authenticationOption,omitempty"`

	// BackupAuthenticationId
	// Identifier of the backup authentication service or profile. At least one backupAuthenticationId or backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup authentication service.
	BackupAuthenticationId *string `json:"backupAuthenticationId,omitempty"`

	// BackupAuthenticationName
	// Name of the backup authentication service or profile. At least one backupAuthenticationId or backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup authentication service. Or could input the 'Always Accept'.
	BackupAuthenticationName *string `json:"backupAuthenticationName,omitempty"`

	// BackupAuthenticationOption
	// Option of the backup authentication service or profile, At least one backupAuthenticationId or backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup authentication service. This only applies to hotspot WLANs.
	// Constraints:
	//    - nullable
	//    - oneof:[Always Accept]
	BackupAuthenticationOption *string `json:"backupAuthenticationOption,omitempty"`

	// Id
	// Identifier of the authentication service or profile. At least one ID or name or authenticationOption is required in the request.
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 location delivery support
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Name
	// Name of the authentication service or profile. At least one ID or name or authenticationOption is required in the request. Or could input the 'Always Accept' or 'Local DB'.
	Name *string `json:"name,omitempty"`

	RealmBasedAuth *bool `json:"realmBasedAuth,omitempty"`

	// ThroughController
	// Indicates whether authentication messages were sent through the controller or not
	ThroughController *bool `json:"throughController,omitempty"`
}

func NewWSGWLANAuthentication() *WSGWLANAuthentication {
	m := new(WSGWLANAuthentication)
	return m
}

// WSGWLANBssMinRateMbps
//
// Definition: wlan_wlanBssMinRateMbps
//
// Constraints:
//    - default:'Disable'
//    - oneof:[Disable,1 mbps,2 mbps,5.5 mbps,12 mbps,24 mbps]
type WSGWLANBssMinRateMbps string

func NewWSGWLANBssMinRateMbps() *WSGWLANBssMinRateMbps {
	m := new(WSGWLANBssMinRateMbps)
	return m
}

// WSGWLANConfiguration
//
// Definition: wlan_wlanConfiguration
type WSGWLANConfiguration struct {
	AccessIpsecProfile *WSGCommonGenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *WSGCommonGenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

	AccountingServiceOrProfile *WSGWLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WSGWLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WSGWLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable for CALEA
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable for CALEA
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WSGWLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *WSGCommonGenericRef `json:"defaultUserTrafficProfile,omitempty"`

	// Description
	// Description of the WLAN
	Description *string `json:"description,omitempty"`

	DevicePolicy *WSGCommonGenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *WSGCommonGenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *WSGCommonGenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *WSGDPSKWlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WSGWLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *WSGDPSKWlanExternalDpsk `json:"externalDpsk,omitempty"`

	// FirewallAppPolicyId
	// Firewall Application Policy of WLAN specific
	FirewallAppPolicyId *string `json:"firewallAppPolicyId,omitempty"`

	// FirewallDevicePolicyId
	// Firewall Device Policy of WLAN specific
	FirewallDevicePolicyId *string `json:"firewallDevicePolicyId,omitempty"`

	// FirewallDownlinkRateLimitingMbps
	// Downlink rate limiting, range 0.1 ~ 200 mpbs
	FirewallDownlinkRateLimitingMbps *float64 `json:"firewallDownlinkRateLimitingMbps,omitempty"`

	// FirewallL2AccessControlPolicyId
	// Firewall L2 Access Control Policy of WLAN specific
	FirewallL2AccessControlPolicyId *string `json:"firewallL2AccessControlPolicyId,omitempty"`

	// FirewallL3AccessControlPolicyId
	// Firewall L3 Access Control Policy of WLAN specific
	FirewallL3AccessControlPolicyId *string `json:"firewallL3AccessControlPolicyId,omitempty"`

	// FirewallProfileId
	// Firewall profile of the WLAN
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUplinkRateLimitingMbps
	// Uplink rate limiting, range 0.1 ~ 200 mpbs
	FirewallUplinkRateLimitingMbps *float64 `json:"firewallUplinkRateLimitingMbps,omitempty"`

	// FirewallUrlFilteringPolicyId
	// Firewall Url Filtering Policy of WLAN specific
	FirewallUrlFilteringPolicyId *string `json:"firewallUrlFilteringPolicyId,omitempty"`

	// FirewallWlanSpecificEnabled
	// Firewall WLAN specific enabled
	FirewallWlanSpecificEnabled *bool `json:"firewallWlanSpecificEnabled,omitempty"`

	FlexiVpnProfile *WSGFlexiVPNSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WSGWLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *WSGCommonGenericRef `json:"hotspot20Profile,omitempty"`

	// Id
	// Identifier of the WLAN
	Id *string `json:"id,omitempty"`

	L2ACL *WSGCommonGenericRef `json:"l2ACL,omitempty"`

	MacAuth *WSGWLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	OperatorRealm *WSGCommonRealm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *WSGCommonGenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WSGWLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WSGWLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WSGWLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// SSID of the WLAN
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the WLAN
	// Constraints:
	//    - oneof:[Standard_Open,Standard_8021X,Standard_Mac,Hotspot,Hotspot_MacByPass,Guest,WebAuth,Hotspot20,Hotspot20_Open,Hotspot20_OSEN]
	Type *string `json:"type,omitempty"`

	Vlan *WSGWLANVlan `json:"vlan,omitempty"`

	// ZoneId
	// Identifier of the zone to which the WLAN belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGWLANConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGWLANConfiguration
}

func newWSGWLANConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGWLANConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGWLANConfigurationAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGWLANConfiguration)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGWLANConfiguration() *WSGWLANConfiguration {
	m := new(WSGWLANConfiguration)
	return m
}

// WSGWLANCoreTunnel
//
// Definition: wlan_wlanCoreTunnel
type WSGWLANCoreTunnel struct {
	// Id
	// Identifier of the forwarding profile. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the forwarding profile. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	// Type
	// Tunnel type
	// Constraints:
	//    - required
	//    - oneof:[L2oGRE,Bridge,TTG_PDG]
	Type *string `json:"type"`
}

func NewWSGWLANCoreTunnel() *WSGWLANCoreTunnel {
	m := new(WSGWLANCoreTunnel)
	return m
}

// WSGWLANDSCPSetting
//
// Definition: wlan_wlanDSCPSetting
type WSGWLANDSCPSetting struct {
	// Enable
	// Enabled or disabled
	// Constraints:
	//    - required
	Enable *bool `json:"enable"`

	// Excepts
	// Constraints:
	//    - nullable
	Excepts []int `json:"excepts,omitempty"`

	// High
	// DSCP range - high
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:255
	High *int `json:"high,omitempty"`

	// Low
	// DSCP range - low
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:255
	Low *int `json:"low,omitempty"`

	// Priority
	// Priority
	// Constraints:
	//    - required
	Priority *int `json:"priority"`
}

func NewWSGWLANDSCPSetting() *WSGWLANDSCPSetting {
	m := new(WSGWLANDSCPSetting)
	return m
}

// WSGWLANEncryption
//
// Definition: wlan_wlanEncryption
type WSGWLANEncryption struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	// Constraints:
	//    - oneof:[AES,TKIP_AES,AES_GCMP_256]
	Algorithm *string `json:"algorithm,omitempty"`

	// KeyIndex
	// Key index. This only applies to WEP64 and WEP128.
	KeyIndex *int `json:"keyIndex,omitempty"`

	// KeyInHex
	// Key in hex format. This only applies to WEP64 and WEP128.
	KeyInHex *string `json:"keyInHex,omitempty"`

	// Method
	// Encryption method
	// Constraints:
	//    - required
	//    - default:'None'
	//    - oneof:[WPA2,WPA_Mixed,WEP_64,WEP_128,None,WPA3,WPA23_Mixed,OWE]
	Method *string `json:"method"`

	// Mfp
	// Management frame protection. This only applies to WPA2 + AES or OWE method.
	// Constraints:
	//    - oneof:[disabled,capable,required]
	Mfp *string `json:"mfp,omitempty"`

	// MobilityDomainId
	// mobility Domain Id.
	// Constraints:
	//    - min:1
	//    - max:65535
	MobilityDomainId *int `json:"mobilityDomainId,omitempty"`

	// Passphrase
	// Passphrase. This only applies to WPA2 and WPA mixed mode.
	Passphrase *string `json:"passphrase,omitempty"`

	// SaePassphrase
	// saePassphrase. This only applies to WPA3 and WPA23 mixed mode.
	SaePassphrase *string `json:"saePassphrase,omitempty"`

	// Support80211rEnabled
	// Enable 802.11r Fast BSS Transition, fast Romaing.
	Support80211rEnabled *bool `json:"support80211rEnabled,omitempty"`
}

func NewWSGWLANEncryption() *WSGWLANEncryption {
	m := new(WSGWLANEncryption)
	return m
}

// WSGWLANHESSID
//
// Definition: wlan_wlanHESSID
type WSGWLANHESSID string

func NewWSGWLANHESSID() *WSGWLANHESSID {
	m := new(WSGWLANHESSID)
	return m
}

// WSGWLANList
//
// Definition: wlan_wlanList
type WSGWLANList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGWLANListAPIResponse struct {
	*RawAPIResponse
	Data *WSGWLANList
}

func newWSGWLANListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGWLANListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGWLANListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGWLANList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGWLANList() *WSGWLANList {
	m := new(WSGWLANList)
	return m
}

// WSGWLANMACAuth
//
// Definition: wlan_wlanMACAuth
type WSGWLANMACAuth struct {
	// CustomizedPassword
	// User defined password. When this field is set to an empty string, the MAC address is used as password.
	// Constraints:
	//    - max:64
	CustomizedPassword *string `json:"customizedPassword,omitempty"`

	// MacAuthMacFormat
	// MAC address format. The default format is 0010a42319c0 and the 802.1X format is 00-10-A4-23-19-C0.
	// Constraints:
	//    - oneof:[Default,802.1X,UpperColon,Upper,LowerDash,LowerColon]
	MacAuthMacFormat *string `json:"macAuthMacFormat,omitempty"`
}

func NewWSGWLANMACAuth() *WSGWLANMACAuth {
	m := new(WSGWLANMACAuth)
	return m
}

// WSGWLANMgmtTxRateMbps
//
// Definition: wlan_wlanMgmtTxRateMbps
//
// Constraints:
//    - default:'2 mbps'
//    - oneof:[1 mbps,2 mbps,5.5 mbps,6 mbps,9 mbps,11 mbps,12 mbps,18 mbps,24 mbps,36 mbps,48 mbps,54 mbps]
type WSGWLANMgmtTxRateMbps string

func NewWSGWLANMgmtTxRateMbps() *WSGWLANMgmtTxRateMbps {
	m := new(WSGWLANMgmtTxRateMbps)
	return m
}

// WSGWLANNameSSID
//
// Definition: wlan_wlanNameSSID
//
// Constraints:
//    - max:32
//    - min:1
type WSGWLANNameSSID string

func NewWSGWLANNameSSID() *WSGWLANNameSSID {
	m := new(WSGWLANNameSSID)
	return m
}

// WSGWLANRadius
//
// Definition: wlan_wlanRadius
type WSGWLANRadius struct {
	// CalledStaIdType
	// Called station ID type
	// Constraints:
	//    - default:'WLAN_BSSID'
	//    - oneof:[WLAN_BSSID,AP_MAC,NONE,AP_GROUP]
	CalledStaIdType *string `json:"calledStaIdType,omitempty"`

	// CustomizedNasId
	// User defined NAS ID
	// Constraints:
	//    - max:64
	CustomizedNasId *string `json:"customizedNasId,omitempty"`

	// NasIdType
	// NAS ID type
	// Constraints:
	//    - default:'WLAN_BSSID'
	//    - oneof:[WLAN_BSSID,AP_MAC,Customized]
	NasIdType *string `json:"nasIdType,omitempty"`

	// NasIpType
	// NAS IP type
	// Constraints:
	//    - default:'disabled'
	//    - oneof:[disabled,control,management,userDefined]
	NasIpType *string `json:"nasIpType,omitempty"`

	// NasIpUserDefined
	// User-defined NAS IP
	// Constraints:
	//    - max:45
	NasIpUserDefined *string `json:"nasIpUserDefined,omitempty"`

	// NasMaxRetry
	// NAS request maximum retry
	// Constraints:
	//    - default:2
	//    - min:2
	//    - max:10
	NasMaxRetry *int `json:"nasMaxRetry,omitempty"`

	// NasReconnectPrimaryMin
	// NAS reconnect primary time in minutes
	// Constraints:
	//    - default:5
	//    - min:1
	//    - max:60
	NasReconnectPrimaryMin *int `json:"nasReconnectPrimaryMin,omitempty"`

	// NasRequestTimeoutSec
	// NAS request timeout in seconds
	// Constraints:
	//    - default:3
	//    - min:2
	//    - max:20
	NasRequestTimeoutSec *int `json:"nasRequestTimeoutSec,omitempty"`

	// SingleSessionIdAcctEnabled
	// When Single Accounting Session ID is enabled, APs will maintain one accounting session for client roaming
	SingleSessionIdAcctEnabled *bool `json:"singleSessionIdAcctEnabled,omitempty"`

	// VendorSpecificAttributeProfileId
	// Vendor Specific Attribute Profile ID
	VendorSpecificAttributeProfileId *string `json:"vendorSpecificAttributeProfileId,omitempty"`
}

func NewWSGWLANRadius() *WSGWLANRadius {
	m := new(WSGWLANRadius)
	return m
}

// WSGWLANSchedule
//
// Definition: wlan_wlanSchedule
type WSGWLANSchedule struct {
	// Id
	// Identifier of the schedule profile. At least one ID or name is required in the request.
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the schedule profile. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	// Type
	// Type of WLAN schedule
	// Constraints:
	//    - required
	//    - default:'AlwaysOn'
	//    - oneof:[AlwaysOn,AlwaysOff,Customized]
	Type *string `json:"type"`
}

func NewWSGWLANSchedule() *WSGWLANSchedule {
	m := new(WSGWLANSchedule)
	return m
}

// WSGWLANSummary
//
// Definition: wlan_wlanSummary
type WSGWLANSummary struct {
	// Id
	// Identifier of the WLAN
	Id *string `json:"id,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	// Ssid
	// SSID of the WLAN
	Ssid *string `json:"ssid,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGWLANSummary() *WSGWLANSummary {
	m := new(WSGWLANSummary)
	return m
}

// WSGWLANVlan
//
// Definition: wlan_wlanVlan
type WSGWLANVlan struct {
	// AaaVlanOverride
	// Indicates whether the AAA VLAN settings can be overriden or not
	AaaVlanOverride *bool `json:"aaaVlanOverride,omitempty"`

	// AccessVlan
	// Access VLAN ID
	// Constraints:
	//    - default:1
	//    - min:1
	//    - max:4094
	AccessVlan *int `json:"accessVlan,omitempty"`

	// CoreQinQEnabled
	// Indicates whether Q-in-Q is allowed at the core network or not
	CoreQinQEnabled *bool `json:"coreQinQEnabled,omitempty"`

	// CoreSVlan
	// Core SVLAN ID. This only applies when core Q-in-Q is enabled
	// Constraints:
	//    - min:1
	//    - max:4094
	CoreSVlan *int `json:"coreSVlan,omitempty"`

	VlanPooling *WSGCommonGenericRef `json:"vlanPooling,omitempty"`
}

func NewWSGWLANVlan() *WSGWLANVlan {
	m := new(WSGWLANVlan)
	return m
}

// AddRkszonesWlansByZoneId
//
// Use this API command to create a new standard, open and non-tunneled basic WLAN.
//
// Operation ID: addRkszonesWlansByZoneId
// Operation path: /rkszones/{zoneId}/wlans
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateStandardOpenWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansByZoneId(ctx context.Context, body *WSGWLANCreateStandardOpenWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansGuestByZoneId
//
// Use this API command to create a new guest access WLAN.
//
// Operation ID: addRkszonesWlansGuestByZoneId
// Operation path: /rkszones/{zoneId}/wlans/guest
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateGuestAccessWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansGuestByZoneId(ctx context.Context, body *WSGWLANCreateGuestAccessWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansGuestByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansHotspot20ByZoneId
//
// Use this API command to create a new Hotspot 2.0 access WLAN.
//
// Operation ID: addRkszonesWlansHotspot20ByZoneId
// Operation path: /rkszones/{zoneId}/wlans/hotspot20
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateHotspot20Wlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20ByZoneId(ctx context.Context, body *WSGWLANCreateHotspot20Wlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansHotspot20ByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansHotspot20openByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as 'Open'.
//
// Operation ID: addRkszonesWlansHotspot20openByZoneId
// Operation path: /rkszones/{zoneId}/wlans/hotspot20open
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateHotspot20OpenWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20openByZoneId(ctx context.Context, body *WSGWLANCreateHotspot20OpenWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansHotspot20openByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansHotspot20osenByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as '802.1X'.
//
// Operation ID: addRkszonesWlansHotspot20osenByZoneId
// Operation path: /rkszones/{zoneId}/wlans/hotspot20osen
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateHotspot20OpenWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20osenByZoneId(ctx context.Context, body *WSGWLANCreateHotspot20OpenWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansHotspot20osenByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansQosMapsById
//
// Use this API command to enable Qos Map Set of a WLAN.
//
// Operation ID: addRkszonesWlansQosMapsById
// Operation path: /rkszones/{zoneId}/wlans/{id}/qosMaps
// Success code: 201 (Created)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansQosMapsById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansQosMapsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddRkszonesWlansStandard8021XByZoneId
//
// Use this API command to create a new standard, 802.1X and non-tunneled WLAN.
//
// Operation ID: addRkszonesWlansStandard8021XByZoneId
// Operation path: /rkszones/{zoneId}/wlans/standard8021X
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateStandard80211Wlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandard8021XByZoneId(ctx context.Context, body *WSGWLANCreateStandard80211Wlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansStandard8021XByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansStandard8021XmacByZoneId
//
// Use this API command to create a new standard, 802.1X with MAC address and non-tunneled WLAN.
//
// Operation ID: addRkszonesWlansStandard8021XmacByZoneId
// Operation path: /rkszones/{zoneId}/wlans/standard8021Xmac
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateStandard80211Wlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandard8021XmacByZoneId(ctx context.Context, body *WSGWLANCreateStandard80211Wlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansStandard8021XmacByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansStandardmacByZoneId
//
// Use this API command to create a new standard, MAC auth and non-tunneled WLAN.
//
// Operation ID: addRkszonesWlansStandardmacByZoneId
// Operation path: /rkszones/{zoneId}/wlans/standardmac
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateStandard80211Wlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandardmacByZoneId(ctx context.Context, body *WSGWLANCreateStandard80211Wlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansStandardmacByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansWebauthByZoneId
//
// Use this API command to creates new web authentication WLAN.
//
// Operation ID: addRkszonesWlansWebauthByZoneId
// Operation path: /rkszones/{zoneId}/wlans/webauth
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateWebAuthWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWebauthByZoneId(ctx context.Context, body *WSGWLANCreateWebAuthWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansWebauthByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansWechatByZoneId
//
// Use this API command to create a new wechat WLAN.
//
// Operation ID: addRkszonesWlansWechatByZoneId
// Operation path: /rkszones/{zoneId}/wlans/wechat
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateWechatWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWechatByZoneId(ctx context.Context, body *WSGWLANCreateWechatWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansWechatByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansWispr8021XByZoneId
//
// Use this API command to create a new hotspot (WISPr) with 802.1X WLAN.
//
// Operation ID: addRkszonesWlansWispr8021XByZoneId
// Operation path: /rkszones/{zoneId}/wlans/wispr8021X
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateHotspotWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWispr8021XByZoneId(ctx context.Context, body *WSGWLANCreateHotspotWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansWispr8021XByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansWisprByZoneId
//
// Use this API command to create new hotspot (WISPr) WLAN.
//
// Operation ID: addRkszonesWlansWisprByZoneId
// Operation path: /rkszones/{zoneId}/wlans/wispr
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateHotspotWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWisprByZoneId(ctx context.Context, body *WSGWLANCreateHotspotWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansWisprByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesWlansWisprmacByZoneId
//
// Use this API command to create a new hotspot (WISPr) with MAC bypass WLAN.
//
// Operation ID: addRkszonesWlansWisprmacByZoneId
// Operation path: /rkszones/{zoneId}/wlans/wisprmac
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGWLANCreateHotspotWlan
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWisprmacByZoneId(ctx context.Context, body *WSGWLANCreateHotspotWlan, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesWlansWisprmacByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesWlansAccountingServiceOrProfileById
//
// Use this API command to disable the accounting of a WLAN.
//
// Operation ID: deleteRkszonesWlansAccountingServiceOrProfileById
// Operation path: /rkszones/{zoneId}/wlans/{id}/accountingServiceOrProfile
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansAccountingServiceOrProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlansAccountingServiceOrProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesWlansById
//
// Use this API command to delete a WLAN.
//
// Operation ID: deleteRkszonesWlansById
// Operation path: /rkszones/{zoneId}/wlans/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesWlansDevicePolicyById
//
// Use this API command to disable the device policy of a WLAN.
//
// Operation ID: deleteRkszonesWlansDevicePolicyById
// Operation path: /rkszones/{zoneId}/wlans/{id}/devicePolicy
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDevicePolicyById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlansDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesWlansDiffServProfileById
//
// Use this API command to disable the DiffServ profile of a WLAN.
//
// Operation ID: deleteRkszonesWlansDiffServProfileById
// Operation path: /rkszones/{zoneId}/wlans/{id}/diffServProfile
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDiffServProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlansDiffServProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesWlansDnsServerProfileById
//
// Use this API command to disable DNS server profile of a WLAN.
//
// Operation ID: deleteRkszonesWlansDnsServerProfileById
// Operation path: /rkszones/{zoneId}/wlans/{id}/dnsServerProfile
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDnsServerProfileById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlansDnsServerProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesWlansL2ACLById
//
// Use this API command to disable the layer 2 access control list (ACL) configuration of a WLAN.
//
// Operation ID: deleteRkszonesWlansL2ACLById
// Operation path: /rkszones/{zoneId}/wlans/{id}/l2ACL
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansL2ACLById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlansL2ACLById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesWlansQosMapsById
//
// Use this API command to disable Qos Map Set of a WLAN.
//
// Operation ID: deleteRkszonesWlansQosMapsById
// Operation path: /rkszones/{zoneId}/wlans/{id}/qosMaps
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansQosMapsById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesWlansQosMapsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesWlansById
//
// Use this API command to retrieve a WLAN.
//
// Operation ID: findRkszonesWlansById
// Operation path: /rkszones/{zoneId}/wlans/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) FindRkszonesWlansById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGWLANConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesWlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANConfigurationAPIResponse), err
}

// FindRkszonesWlansByZoneId
//
// Use this API command to retrieve a list of WLANs within a zone.
//
// Operation ID: findRkszonesWlansByZoneId
// Operation path: /rkszones/{zoneId}/wlans
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWLANService) FindRkszonesWlansByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGWLANListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesWlansByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANListAPIResponse), err
}

// FindWlanByQueryCriteria
//
// Query WLANs with specified filters.
//
// Operation ID: findWlanByQueryCriteria
// Operation path: /query/wlan
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWLANService) FindWlanByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGWLANQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGWLANQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGWLANQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindWlanByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGWLANQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGWLANQueryListAPIResponse), err
}

// PartialUpdateRkszonesWlansById
//
// Use this API command to modify the configuration of a WLAN.
//
// Operation ID: partialUpdateRkszonesWlansById
// Operation path: /rkszones/{zoneId}/wlans/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGWLANModifyWlan
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) PartialUpdateRkszonesWlansById(ctx context.Context, body *WSGWLANModifyWlan, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesWlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesWlansById
//
// Use this API command to modify entire information of a WLAN.
//
// Operation ID: updateRkszonesWlansById
// Operation path: /rkszones/{zoneId}/wlans/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGWLANModifyWlan
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWLANService) UpdateRkszonesWlansById(ctx context.Context, body *WSGWLANModifyWlan, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesWlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
