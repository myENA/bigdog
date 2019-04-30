package wlan

// API Version: v8_0

type CreateGuestAccessWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateHotspot20OpenWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateHotspot20OSENWLAN struct {
	*CreateHotspot20OpenWLAN
}

type CreateHotspot20WLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateHotspot8021XWLAN struct {
	*CreateHotspotWLAN
}

type CreateHotspotMacWLAN struct {
	*CreateHotspotWLAN
}

type CreateHotspotWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateStandard80211WLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateStandard80211WLANMac struct {
	*CreateStandard80211WLAN
}

type CreateStandardMacWLAN struct {
	*CreateStandard80211WLAN
}

type CreateStandardOpenWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateWebAuthWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type CreateWechatWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the
	// data plane, and SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

	Vlan *WLANVlan `json:"vlan,omitempty"`
}

type ModifyWLAN struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE
	// tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DNSServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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

	Ssid *string `json:"ssid,omitempty"`

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
	InterimUpdateMin *int `json:"interimUpdateMin,omitempty"`

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
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	// AssocRssiThr
	// Assoc RSSI threshold.f
	AssocRssiThr *int `json:"assocRssiThr,omitempty"`

	// AuthRssiThr
	// Auth RSSI threshold.
	AuthRssiThr *int `json:"authRssiThr,omitempty"`

	// AvcEnabled
	// Indicator of whether ARC support is enabled or disabled
	AvcEnabled *bool `json:"avcEnabled,omitempty"`

	// BandBalancing
	// Indicates whether band balancing is enabled or disabled
	BandBalancing *string `json:"bandBalancing,omitempty"`

	BssMinRateMbps *string `json:"bssMinRateMbps,omitempty"`

	// ClientFingerprintingEnabled
	// Indicates whether client fingerprinting is enabled or disabled
	ClientFingerprintingEnabled *bool `json:"clientFingerprintingEnabled,omitempty"`

	// ClientIdleTimeoutSec
	// Client idle timeout in seconds
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

	ClientIsolationWhitelist *common.GenericRef `json:"clientIsolationWhitelist,omitempty"`

	// ClientLoadBalancingEnabled
	// Indicates whether Client Load Balancing is enabled or disabled
	ClientLoadBalancingEnabled *bool `json:"clientLoadBalancingEnabled,omitempty"`

	// DgafEnabled
	// Indicates whether dgaf is enabled or disabled
	DgafEnabled *bool `json:"dgafEnabled,omitempty"`

	// DHCP82Format
	// DHCP Option 82 format
	DHCP82Format *string `json:"dhcp82Format,omitempty"`

	// DHCPOption82Enabled
	// Indicates whether DCHP Option 82 is enabled or disabled
	DHCPOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// DHCPRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	DHCPRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	// DirectedThreshold
	// Directed Threshold Setting, Defines the client count at which an AP will stop converting group
	// addressed data traffic to unicast.
	DirectedThreshold *int `json:"directedThreshold,omitempty"`

	// DownlinkEnabled
	// SSID Rate Limiting downlink enabled.
	DownlinkEnabled *bool `json:"downlinkEnabled,omitempty"`

	// DownlinkRate
	// SSID Rate Limiting downlink.
	DownlinkRate *float64 `json:"downlinkRate,omitempty"`

	// DtimInterval
	// DTIM Interval
	DtimInterval *int `json:"dtimInterval,omitempty"`

	EnableRadiusBasedDHCPNat *bool `json:"enableRadiusBasedDhcpNat,omitempty"`

	// FlowLogEnabled
	// Flow log enabled.
	FlowLogEnabled *bool `json:"flowLogEnabled,omitempty"`

	// ForceClientDHCPTimeoutSec
	// Force DHCP disconnects the client if the client does not obtain a valid IP address within the timeout
	// peroid. To disable force DHCP, set this value to zero (0).
	ForceClientDHCPTimeoutSec *int `json:"forceClientDHCPTimeoutSec,omitempty"`

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
	JoinAcceptTimeout *int `json:"joinAcceptTimeout,omitempty"`

	// JoinIgnoreThr
	// Join wait threshold.
	JoinIgnoreThr *int `json:"joinIgnoreThr,omitempty"`

	// JoinIgnoreTimeout
	// Join wait time.
	JoinIgnoreTimeout *int `json:"joinIgnoreTimeout,omitempty"`

	// MaxAllowedRA
	// Max Allowed RAs
	MaxAllowedRA *int `json:"maxAllowedRA,omitempty"`

	// MaxClientsPerRadio
	// Maximum number of clients per radio
	MaxClientsPerRadio *int `json:"maxClientsPerRadio,omitempty"`

	MgmtTxRateMbps *string `json:"mgmtTxRateMbps,omitempty"`

	// NdProxyEnabled
	// Indicates whether ND Proxy is enabled or disabled
	NdProxyEnabled *bool `json:"ndProxyEnabled,omitempty"`

	// OceBroadcastProbeResponseDelay
	// Broadcast probe response delay.
	OceBroadcastProbeResponseDelay *int `json:"oceBroadcastProbeResponseDelay,omitempty"`

	// OceEnabled
	// Optimized Connectivity Experience(OCE) enabled.
	OceEnabled *bool `json:"oceEnabled,omitempty"`

	// OceRssiBasedAssociationRejectionThreshold
	// RSSI-based association rejection threshold.
	OceRssiBasedAssociationRejectionThreshold *int `json:"oceRssiBasedAssociationRejectionThreshold,omitempty"`

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
	Priority *string `json:"priority,omitempty"`

	// ProbeRssiThr
	// Join RSSI threshold.
	ProbeRssiThr *int `json:"probeRssiThr,omitempty"`

	// ProxyARPEnabled
	// Indicates whether proxy ARP is enabled or disabled
	ProxyARPEnabled *bool `json:"proxyARPEnabled,omitempty"`

	// RaInterval
	// A timer that RA proxy runs and once receives unsolicited RA checks against the configured time and
	// allow/drop RA based on next timeout
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

	// Support8021DEnabled
	// Indicates whether support for 802.11d is enabled or disabled
	Support8021DEnabled *bool `json:"support80211dEnabled,omitempty"`

	// Support8021KEnabled
	// Indicates whether support for 802.11k is enabled or disabled
	Support8021KEnabled *bool `json:"support80211kEnabled,omitempty"`

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
	AuthenticationOption *string `json:"authenticationOption,omitempty"`

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
	BackupAuthenticationOption *string `json:"backupAuthenticationOption,omitempty"`

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

type WLANConfiguration struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, SoftGRE means AP direct SoftGRE tunnel
	AccessTunnelType *string `json:"accessTunnelType,omitempty"`

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

	Dpsk *dpsk.WLANDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDHCPEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDHCPEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WLANEncryption `json:"encryption,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *string `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	// ID
	// Identifier of the WLAN
	ID *string `json:"id,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WLANMACAuth `json:"macAuth,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	OperatorRealm *string `json:"operatorRealm,omitempty"`

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
	Type *string `json:"type,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

type WLANDSCPSetting struct {
	// Enable
	// Enabled or disabled
	Enable *bool `json:"enable,omitempty"`

	Excepts []int `json:"excepts,omitempty"`

	// High
	// DSCP range - high
	High *int `json:"high,omitempty"`

	// Low
	// DSCP range - low
	Low *int `json:"low,omitempty"`

	// Priority
	// Priority
	Priority *int `json:"priority,omitempty"`
}

type WLANEncryption struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	Algorithm *string `json:"algorithm,omitempty"`

	// KeyIndex
	// Key index. This only applies to WEP64 and WEP128.
	KeyIndex *int `json:"keyIndex,omitempty"`

	// KeyInHex
	// Key in hex format. This only applies to WEP64 and WEP128.
	KeyInHex *string `json:"keyInHex,omitempty"`

	// Method
	// Encryption method
	Method *string `json:"method,omitempty"`

	// Mfp
	// Management frame protection. This only applies to WPA2 + AES
	Mfp *string `json:"mfp,omitempty"`

	// MobilityDomainID
	// mobility Domain Id.
	MobilityDomainID *int `json:"mobilityDomainId,omitempty"`

	// Passphrase
	// Passphrase. This only applies to WPA2 and WPA mixed mode.
	Passphrase *string `json:"passphrase,omitempty"`

	// Support8021REnabled
	// Enable 802.11r Fast BSS Transition, fast Romaing.
	Support8021REnabled *bool `json:"support80211rEnabled,omitempty"`
}

type WLANList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WLANSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WLANMACAuth struct {
	// CustomizedPassword
	// User defined password. When this field is set to an empty string, the MAC address is used as password.
	CustomizedPassword *string `json:"customizedPassword,omitempty"`

	// MacAuthMacFormat
	// MAC address format. The default format is 0010a42319c0 and the 802.1X format is 00-10-A4-23-19-C0.
	MacAuthMacFormat *string `json:"macAuthMacFormat,omitempty"`
}

type WLANRadius struct {
	// CalledStaIDType
	// Called station ID type
	CalledStaIDType *string `json:"calledStaIdType,omitempty"`

	// CustomizedNasID
	// User defined NAS ID
	CustomizedNasID *string `json:"customizedNasId,omitempty"`

	// NasIDType
	// NAS ID type
	NasIDType *string `json:"nasIdType,omitempty"`

	// NasIPType
	// NAS IP type
	NasIPType *string `json:"nasIpType,omitempty"`

	// NasIPUserDefined
	// User-defined NAS IP
	NasIPUserDefined *string `json:"nasIpUserDefined,omitempty"`

	// NasMaxRetry
	// NAS request maximum retry
	NasMaxRetry *int `json:"nasMaxRetry,omitempty"`

	// NasReconnectPrimaryMin
	// NAS reconnect primary time in minutes
	NasReconnectPrimaryMin *int `json:"nasReconnectPrimaryMin,omitempty"`

	// NasRequestTimeoutSec
	// NAS request timeout in seconds
	NasRequestTimeoutSec *int `json:"nasRequestTimeoutSec,omitempty"`

	// SingleSessionIDAcctEnabled
	// When Single Accounting Session ID is enabled, APs will maintain one accounting session for client
	// roaming
	SingleSessionIDAcctEnabled *bool `json:"singleSessionIdAcctEnabled,omitempty"`
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
	Type *string `json:"type,omitempty"`
}

type WLANSummary struct {
	// ID
	// Identifier of the WLAN
	ID *string `json:"id,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

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
	AccessVlan *int `json:"accessVlan,omitempty"`

	// CoreQinQEnabled
	// Indicates whether Q-in-Q is allowed at the core network or not
	CoreQinQEnabled *bool `json:"coreQinQEnabled,omitempty"`

	// CoreSVlan
	// Core SVLAN ID. This only applies when core Q-in-Q is enabled
	CoreSVlan *int `json:"coreSVlan,omitempty"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}
