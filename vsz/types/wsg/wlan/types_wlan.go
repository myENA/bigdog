package wlan

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
	"github.com/myENA/ruckus-client/vsz/types/wsg/flexivpn"
)

type CreateGuestAccessWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty" validate:"required"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateHotspot20OpenWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateHotspot20WLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty" validate:"required"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateHotspotWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty" validate:"required"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateStandard80211WLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateStandardOpenWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateWebAuthWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty" validate:"required"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateWechatWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty" validate:"required"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty" validate:"required"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type ModifyWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE
	// tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDPSK *dpsk.WLANExternalDPSK `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *WLANNameSSID `json:"name,omitempty"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WLANNameSSID `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type WLANAccounting struct {
	// AccountingDelayEnabled
	// Indicates whether accounting delay time is enabled
	AccountingDelayEnabled *bool `json:"accountingDelayEnabled,omitempty"`

	// BackupAccountingID
	// Backup accounting service or profile ID. At least one backupAccountingId or backupAccountingName is
	// required in the request when setting backup accounting service.
	BackupAccountingID *string `json:"backupAccountingId,omitempty"`

	// BackupAccountingName
	// Backup accounting service or profile name. At least one backupAccountingId or backupAccountingName is
	// required in the request when setting backup accounting service.
	BackupAccountingName *string `json:"backupAccountingName,omitempty"`

	// ID
	// Accounting service or profile ID. At least one ID or name is required in the request.
	ID *string `json:"id,omitempty"`

	// InterimUpdateMin
	// Interval (in minutes) for sending interim updates
	InterimUpdateMin *int `json:"interimUpdateMin,omitempty" validate:"gte=0,lte=1440"`

	// Name
	// Accounting service or profile name. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	RealmBasedAcct *bool `json:"realmBasedAcct,omitempty"`

	// ThroughController
	// Indicates whether accounting messages were sent through the controller
	ThroughController *bool `json:"throughController,omitempty"`
}

type WLANAdvanced struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	// AssocRssiThr
	// Assoc RSSI threshold.f
	AssocRssiThr *int `json:"assocRssiThr,omitempty" validate:"gte=-90,lte=-60"`

	// AuthRssiThr
	// Auth RSSI threshold.
	AuthRssiThr *int `json:"authRssiThr,omitempty" validate:"gte=-90,lte=-60"`

	// AvcEnabled
	// Indicator of whether ARC support is enabled or disabled
	AvcEnabled *bool `json:"avcEnabled,omitempty"`

	// BandBalancing
	// Indicates whether band balancing is enabled or disabled
	BandBalancing *string `json:"bandBalancing,omitempty" validate:"oneof=Disabled UseZoneSetting"`

	BssMinRateMbps *WLANBssMinRateMbps `json:"bssMinRateMbps,omitempty"`

	// ClientFingerprintingEnabled
	// Indicates whether client fingerprinting is enabled or disabled
	ClientFingerprintingEnabled *bool `json:"clientFingerprintingEnabled,omitempty"`

	// ClientIdleTimeoutSec
	// Client idle timeout in seconds
	ClientIdleTimeoutSec *int `json:"clientIdleTimeoutSec,omitempty" validate:"gte=60,lte=1000"`

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

	ClientIsolationWhitelist *common.GenericRef `json:"clientIsolationWhitelist,omitempty"`

	// ClientLoadBalancingEnabled
	// Indicates whether Client Load Balancing is enabled or disabled
	ClientLoadBalancingEnabled *bool `json:"clientLoadBalancingEnabled,omitempty"`

	// DgafEnabled
	// Indicates whether dgaf is enabled or disabled
	DgafEnabled *bool `json:"dgafEnabled,omitempty"`

	// DHCP82Format
	// DHCP Option 82 format. This variable no longer supports from v8_1 and only kept for backward
	// compatibility.
	DHCP82Format *string `json:"dhcp82Format,omitempty" validate:"oneof=RUCKUS_DEFAULT RADIUS_DHCP RADIUS_NAT RADIUS_DHCP_NAT RADIUS_NAT_RUCKUS RADIUS_NAT_SOFTGRE SOFTGRE_CUSTOMIZED"`

	// DHCP82SubOpt1Format
	// Subopt-1 format
	DHCP82SubOpt1Format *string `json:"dhcp82SubOpt1Format,omitempty" validate:"oneof=NONE SUBOPT1_AP_INFO_LOCATION SUBOPT1_AP_INFO SUBOPT1_AP_MAC_ESSID_PRIVACYTYPE SUBOPT1_AP_MAC_hex SUBOPT1_AP_MAC_hex_ESSID SUBOPT1_ESSID"`

	// DHCP82SubOpt2Format
	// Subopt-2 format
	DHCP82SubOpt2Format *string `json:"dhcp82SubOpt2Format,omitempty" validate:"oneof=NONE SUBOPT2_CLIENT_MAC SUBOPT2_CLIENT_MAC_hex SUBOPT2_CLIENT_MAC_hex_ESSID SUBOPT2_AP_MAC SUBOPT2_AP_MAC_hex SUBOPT2_AP_MAC_hex_ESSID"`

	// DHCP82SubOpt150Format
	// Subopt-150 with VLAN-Id
	DHCP82SubOpt150Format *string `json:"dhcp82SubOpt150Format,omitempty" validate:"oneof=NONE SUBOPT150_VLAN_ID"`

	// DHCP82SubOpt151AreaName
	// Subopt-151 Area Name value
	DHCP82SubOpt151AreaName *string `json:"dhcp82SubOpt151AreaName,omitempty"`

	// DHCP82SubOpt151Format
	// Subopt-151 format
	DHCP82SubOpt151Format *string `json:"dhcp82SubOpt151Format,omitempty" validate:"oneof=NONE SUBOPT151_AREA_NAME SUBOPT151_ESSID"`

	// DHCP82SubOptRadiusFormat
	// Radius DHCP/NAT option 82 Format
	DHCP82SubOptRadiusFormat *string `json:"dhcp82SubOptRadiusFormat,omitempty" validate:"oneof=NONE RUCKUS_DEFAULT RADIUS_DHCP RADIUS_NAT RADIUS_DHCP_NAT RADIUS_NAT_RUCKUS RADIUS_NAT_SOFTGRE SOFTGRE_CUSTOMIZED"`

	// DHCPOption82Enabled
	// Indicates whether DCHP Option 82 is enabled or disabled. This variable no longer supports from v8_1 and
	// only kept for backward compatibility.
	DHCPOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// DHCPRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DHCPRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	// DirectedThreshold
	// Directed Threshold Setting, Defines the client count at which an AP will stop converting group
	// addressed data traffic to unicast.
	DirectedThreshold *int `json:"directedThreshold,omitempty" validate:"gte=0,lte=128"`

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
	DtimInterval *int `json:"dtimInterval,omitempty" validate:"gte=1,lte=255"`

	EnableRadiusBasedDHCPNat *bool `json:"enableRadiusBasedDhcpNat,omitempty"`

	// FlowLogEnabled
	// Flow log enabled.
	FlowLogEnabled *bool `json:"flowLogEnabled,omitempty"`

	// ForceClientDHCPTimeoutSec
	// Force DHCP disconnects the client if the client does not obtain a valid IP address within the timeout
	// peroid. To disable force DHCP, set this value to zero (0).
	ForceClientDHCPTimeoutSec *int `json:"forceClientDHCPTimeoutSec,omitempty" validate:"oneof=0 5 6 7 8 9 10 11 12 13 14 15"`

	// HdOverheadOptimizeEnable
	// Airtime Decongestion enabled.
	HdOverheadOptimizeEnable *bool `json:"hdOverheadOptimizeEnable,omitempty"`

	// HideSsidEnabled
	// Indicates whether the SSID is hidden or broadcast
	HideSsidEnabled *bool `json:"hideSsidEnabled,omitempty"`

	// Hs20Onboarding
	// Allow WISPr WLAN for Hotspot 2.0 Onboarding
	Hs20Onboarding *bool `json:"hs20Onboarding,omitempty"`

	// JoinAcceptTimeout
	// Join expire time.
	JoinAcceptTimeout *int `json:"joinAcceptTimeout,omitempty" validate:"gte=1,lte=300"`

	// JoinIgnoreThr
	// Join wait threshold.
	JoinIgnoreThr *int `json:"joinIgnoreThr,omitempty" validate:"gte=1,lte=50"`

	// JoinIgnoreTimeout
	// Join wait time.
	JoinIgnoreTimeout *int `json:"joinIgnoreTimeout,omitempty" validate:"gte=1,lte=60"`

	// MaxAllowedRA
	// Max Allowed RAs
	MaxAllowedRA *int `json:"maxAllowedRA,omitempty" validate:"gte=1,lte=1440"`

	// MaxClientsPerRadio
	// Maximum number of clients per radio
	MaxClientsPerRadio *int `json:"maxClientsPerRadio,omitempty" validate:"gte=1,lte=512"`

	MgmtTxRateMbps *WLANMgmtTxRateMbps `json:"mgmtTxRateMbps,omitempty"`

	// NdProxyEnabled
	// Indicates whether ND Proxy is enabled or disabled
	NdProxyEnabled *bool `json:"ndProxyEnabled,omitempty"`

	// OceBroadcastProbeResponseDelay
	// Broadcast probe response delay.
	OceBroadcastProbeResponseDelay *int `json:"oceBroadcastProbeResponseDelay,omitempty" validate:"gte=8,lte=120"`

	// OceEnabled
	// Optimized Connectivity Experience(OCE) enabled.
	OceEnabled *bool `json:"oceEnabled,omitempty"`

	// OceRssiBasedAssociationRejectionThreshold
	// RSSI-based association rejection threshold.
	OceRssiBasedAssociationRejectionThreshold *int `json:"oceRssiBasedAssociationRejectionThreshold,omitempty" validate:"gte=-90,lte=-60"`

	// OfdmOnlyEnabled
	// Indicates whether OFDM only is enabled or disabled
	OfdmOnlyEnabled *bool `json:"ofdmOnlyEnabled,omitempty"`

	// OkcEnabled
	// Indicator of whether OKC support is enabled or disabled. The default value is true when the WLAN is
	// WPA+AES non open WLAN.
	OkcEnabled *bool `json:"okcEnabled,omitempty"`

	// PmkCachingEnabled
	// Indicator of whether PKM caching support is enabled or disabled. The default value is true when the
	// WLAN is WPA+AES non open WLAN.
	PmkCachingEnabled *bool `json:"pmkCachingEnabled,omitempty"`

	// Priority
	// Priority of the WLAN
	Priority *string `json:"priority,omitempty" validate:"oneof=High Low"`

	// ProbeRssiThr
	// Join RSSI threshold.
	ProbeRssiThr *int `json:"probeRssiThr,omitempty" validate:"gte=-90,lte=-60"`

	// ProxyARPEnabled
	// Indicates whether proxy ARP is enabled or disabled
	ProxyARPEnabled *bool `json:"proxyARPEnabled,omitempty"`

	// RaInterval
	// A timer that RA proxy runs and once receives unsolicited RA checks against the configured time and
	// allow/drop RA based on next timeout
	RaInterval *int `json:"raInterval,omitempty" validate:"gte=1,lte=256"`

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

	// Support80211DEnabled
	// Indicates whether support for 802.11d is enabled or disabled
	Support80211DEnabled *bool `json:"support80211dEnabled,omitempty"`

	// Support80211KEnabled
	// Indicates whether support for 802.11k is enabled or disabled
	Support80211KEnabled *bool `json:"support80211kEnabled,omitempty"`

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

	// URLFilteringPolicyEnabled
	// Indicator of whether URL Filtering is enabled or disabled
	URLFilteringPolicyEnabled *bool `json:"urlFilteringPolicyEnabled,omitempty"`

	// URLFilteringPolicyID
	// The URL Filtering policy ID.
	URLFilteringPolicyID *string `json:"urlFilteringPolicyId,omitempty"`

	// WifiCallingPolicyEnabled
	// Indicator of whether Wi-Fi Calling is enabled or disabled
	WifiCallingPolicyEnabled *bool `json:"wifiCallingPolicyEnabled,omitempty"`

	// WifiCallingPolicyIds
	// The Wi-Fi Calling policy IDs. (Maximum allowed number is 5)
	WifiCallingPolicyIds []string `json:"wifiCallingPolicyIds,omitempty"`
}

type WLANAuthentication struct {
	// AuthenticationOption
	// Option of the authentication service or profile, At least one ID or name or authenticationOption is
	// required in the request. This only applies to hotspot and guest WLANs.
	AuthenticationOption *string `json:"authenticationOption,omitempty" validate:"oneof=Local DB Guest Always Accept "`

	// BackupAuthenticationID
	// Identifier of the backup authentication service or profile. At least one backupAuthenticationId or
	// backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup
	// authentication service.
	BackupAuthenticationID *string `json:"backupAuthenticationId,omitempty"`

	// BackupAuthenticationName
	// Name of the backup authentication service or profile. At least one backupAuthenticationId or
	// backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup
	// authentication service. Or could input the 'Always Accept'.
	BackupAuthenticationName *string `json:"backupAuthenticationName,omitempty"`

	// BackupAuthenticationOption
	// Option of the backup authentication service or profile, At least one backupAuthenticationId or
	// backupAuthenticationName or backupAuthenticationOption is required in the request when setting backup
	// authentication service. This only applies to hotspot WLANs.
	BackupAuthenticationOption *string `json:"backupAuthenticationOption,omitempty" validate:"oneof=Always Accept "`

	// ID
	// Identifier of the authentication service or profile. At least one ID or name or authenticationOption is
	// required in the request.
	ID *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 location delivery support
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Name
	// Name of the authentication service or profile. At least one ID or name or authenticationOption is
	// required in the request. Or could input the 'Always Accept' or 'Local DB'.
	Name *string `json:"name,omitempty"`

	RealmBasedAuth *bool `json:"realmBasedAuth,omitempty"`

	// ThroughController
	// Indicates whether authentication messages were sent through the controller or not
	ThroughController *bool `json:"throughController,omitempty"`
}

type WLANBssMinRateMbps string

type WLANConfiguration struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WLANAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WLANAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WLANAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// BypassCNA
	// Bypass Capitive Network Assitance
	BypassCNA *bool `json:"bypassCNA,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WLANCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	// Description
	// Description of the WLAN
	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	DPSK *dpsk.WLANDPSKSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WLANHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	// ID
	// Identifier of the WLAN
	ID *string `json:"id,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileID *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileID
	// Precedence profile of the WLAN
	PrecedenceProfileID *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WLANDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WLANRadius `json:"radiusOptions,omitempty"`

	Schedule *WLANSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileID *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// SSID of the WLAN
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the WLAN
	Type *string `json:"type,omitempty" validate:"oneof=Standard_Open Standard_8021X Standard_Mac Hotspot Hotspot_MacByPass Guest WebAuth Hotspot20 Hotspot20_Open Hotspot20_OSEN"`

	Vlan *WLANVlan `json:"vlan,omitempty"`

	// ZoneID
	// Identifier of the zone to which the WLAN belongs
	ZoneID *string `json:"zoneId,omitempty"`
}

type WLANCoreTunnel struct {
	// ID
	// Identifier of the forwarding profile. At least one ID or name is required in the request.
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the forwarding profile. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	// Type
	// Tunnel type
	Type *string `json:"type,omitempty" validate:"required,oneof=L2oGRE Bridge TTG_PDG"`
}

type WLANDSCPSetting struct {
	// Enable
	// Enabled or disabled
	Enable *bool `json:"enable,omitempty" validate:"required"`

	Excepts []int `json:"excepts,omitempty"`

	// High
	// DSCP range - high
	High *int `json:"high,omitempty" validate:"gte=0,lte=255"`

	// Low
	// DSCP range - low
	Low *int `json:"low,omitempty" validate:"gte=0,lte=255"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty" validate:"required"`
}

type WLANEncryption struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	Algorithm *string `json:"algorithm,omitempty" validate:"oneof=AES TKIP_AES"`

	// KeyIndex
	// Key index. This only applies to WEP64 and WEP128.
	KeyIndex *int `json:"keyIndex,omitempty"`

	// KeyInHex
	// Key in hex format. This only applies to WEP64 and WEP128.
	KeyInHex *string `json:"keyInHex,omitempty"`

	// Method
	// Encryption method
	Method *string `json:"method,omitempty" validate:"required,oneof=WPA2 WPA_Mixed WEP_64 WEP_128 None"`

	// Mfp
	// Management frame protection. This only applies to WPA2 + AES
	Mfp *string `json:"mfp,omitempty" validate:"oneof=disabled capable required"`

	// MobilityDomainID
	// mobility Domain Id.
	MobilityDomainID *int `json:"mobilityDomainId,omitempty" validate:"gte=1,lte=65535"`

	// Passphrase
	// Passphrase. This only applies to WPA2 and WPA mixed mode.
	Passphrase *string `json:"passphrase,omitempty"`

	// Support80211REnabled
	// Enable 802.11r Fast BSS Transition, fast Romaing.
	Support80211REnabled *bool `json:"support80211rEnabled,omitempty"`
}

type WLANHESSID string

type WLANList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WLANSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WLANMACAuth struct {
	// CustomizedPassword
	// User defined password. When this field is set to an empty string, the MAC address is used as password.
	CustomizedPassword *string `json:"customizedPassword,omitempty" validate:"max=64"`

	// MacAuthMacFormat
	// MAC address format. The default format is 0010a42319c0 and the 802.1X format is 00-10-A4-23-19-C0.
	MacAuthMacFormat *string `json:"macAuthMacFormat,omitempty" validate:"oneof=Default 802.1X UpperColon Upper LowerDash LowerColon"`
}

type WLANMgmtTxRateMbps string

type WLANNameSSID string

type WLANRadius struct {
	// CalledStaIDType
	// Called station ID type
	CalledStaIDType *string `json:"calledStaIdType,omitempty" validate:"oneof=AP_MAC NONE WLAN_BSSID"`

	// CustomizedNasID
	// User defined NAS ID
	CustomizedNasID *string `json:"customizedNasId,omitempty" validate:"max=64"`

	// NasIDType
	// NAS ID type
	NasIDType *string `json:"nasIdType,omitempty" validate:"oneof=AP_MAC Customized WLAN_BSSID"`

	// NasIPType
	// NAS IP type
	NasIPType *string `json:"nasIpType,omitempty" validate:"oneof=disabled control management userDefined"`

	// NasIPUserDefined
	// User-defined NAS IP
	NasIPUserDefined *string `json:"nasIpUserDefined,omitempty" validate:"max=45"`

	// NasMaxRetry
	// NAS request maximum retry
	NasMaxRetry *int `json:"nasMaxRetry,omitempty" validate:"gte=2,lte=10"`

	// NasReconnectPrimaryMin
	// NAS reconnect primary time in minutes
	NasReconnectPrimaryMin *int `json:"nasReconnectPrimaryMin,omitempty" validate:"gte=1,lte=60"`

	// NasRequestTimeoutSec
	// NAS request timeout in seconds
	NasRequestTimeoutSec *int `json:"nasRequestTimeoutSec,omitempty" validate:"gte=2,lte=20"`

	// SingleSessionIDAcctEnabled
	// When Single Accounting Session ID is enabled, APs will maintain one accounting session for client
	// roaming
	SingleSessionIDAcctEnabled *bool `json:"singleSessionIdAcctEnabled,omitempty"`

	// VendorSpecificAttributeProfileID
	// Vendor Specific Attribute Profile ID
	VendorSpecificAttributeProfileID *string `json:"vendorSpecificAttributeProfileId,omitempty"`
}

type WLANSchedule struct {
	// ID
	// Identifier of the schedule profile. At least one ID or name is required in the request.
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the schedule profile. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	// Type
	// Type of WLAN schedule
	Type *string `json:"type,omitempty" validate:"required,oneof=AlwaysOn AlwaysOff Customized"`
}

type WLANSummary struct {
	// ID
	// Identifier of the WLAN
	ID *string `json:"id,omitempty"`

	// MVNOID
	// Tenant UUID
	MVNOID *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	// Ssid
	// SSID of the WLAN
	Ssid *string `json:"ssid,omitempty"`

	// ZoneID
	// Zone ID
	ZoneID *string `json:"zoneId,omitempty"`
}

type WLANVlan struct {
	// AaaVlanOverride
	// Indicates whether the AAA VLAN settings can be overriden or not
	AaaVlanOverride *bool `json:"aaaVlanOverride,omitempty"`

	// AccessVlan
	// Access VLAN ID
	AccessVlan *int `json:"accessVlan,omitempty" validate:"gte=1,lte=4094"`

	// CoreQinQEnabled
	// Indicates whether Q-in-Q is allowed at the core network or not
	CoreQinQEnabled *bool `json:"coreQinQEnabled,omitempty"`

	// CoreSVlan
	// Core SVLAN ID. This only applies when core Q-in-Q is enabled
	CoreSVlan *int `json:"coreSVlan,omitempty" validate:"gte=1,lte=4094"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}
