package wlan

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
	"github.com/myENA/ruckus-client/vsz/types/wsg/flexivpn"
)

type CreateGuestAccessWlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile" validate:"required"`

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

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateGuestAccessWlan() *CreateGuestAccessWlan {
	createGuestAccessWlanType := new(CreateGuestAccessWlan)
	return createGuestAccessWlanType
}

func NewCreateGuestAccessWlanWithDefaults() *CreateGuestAccessWlan {
	createGuestAccessWlanType := new(CreateGuestAccessWlan)
	bypassCNAField := false
	createGuestAccessWlanType.BypassCNA = &bypassCNAField
	return createGuestAccessWlanType
}

type CreateHotspot20OpenWlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateHotspot20OpenWlan() *CreateHotspot20OpenWlan {
	createHotspot20OpenWlanType := new(CreateHotspot20OpenWlan)
	return createHotspot20OpenWlanType
}

func NewCreateHotspot20OpenWlanWithDefaults() *CreateHotspot20OpenWlan {
	createHotspot20OpenWlanType := new(CreateHotspot20OpenWlan)
	return createHotspot20OpenWlanType
}

type CreateHotspot20Wlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	// Hotspot20Profile
	// Constraints:
	//    - required
	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile" validate:"required"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateHotspot20Wlan() *CreateHotspot20Wlan {
	createHotspot20WlanType := new(CreateHotspot20Wlan)
	return createHotspot20WlanType
}

func NewCreateHotspot20WlanWithDefaults() *CreateHotspot20Wlan {
	createHotspot20WlanType := new(CreateHotspot20Wlan)
	return createHotspot20WlanType
}

type CreateHotspotWlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile" validate:"required"`

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

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateHotspotWlan() *CreateHotspotWlan {
	createHotspotWlanType := new(CreateHotspotWlan)
	return createHotspotWlanType
}

func NewCreateHotspotWlanWithDefaults() *CreateHotspotWlan {
	createHotspotWlanType := new(CreateHotspotWlan)
	bypassCNAField := false
	createHotspotWlanType.BypassCNA = &bypassCNAField
	return createHotspotWlanType
}

type CreateStandard80211Wlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile" validate:"required"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateStandard80211Wlan() *CreateStandard80211Wlan {
	createStandard80211WlanType := new(CreateStandard80211Wlan)
	return createStandard80211WlanType
}

func NewCreateStandard80211WlanWithDefaults() *CreateStandard80211Wlan {
	createStandard80211WlanType := new(CreateStandard80211Wlan)
	return createStandard80211WlanType
}

type CreateStandardOpenWlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - default:'APLBO'
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateStandardOpenWlan() *CreateStandardOpenWlan {
	createStandardOpenWlanType := new(CreateStandardOpenWlan)
	return createStandardOpenWlanType
}

func NewCreateStandardOpenWlanWithDefaults() *CreateStandardOpenWlan {
	createStandardOpenWlanType := new(CreateStandardOpenWlan)
	accessTunnelTypeField := `APLBO`
	createStandardOpenWlanType.AccessTunnelType = &accessTunnelTypeField
	return createStandardOpenWlanType
}

type CreateWebAuthWlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	// AuthServiceOrProfile
	// Constraints:
	//    - required
	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile" validate:"required"`

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

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateWebAuthWlan() *CreateWebAuthWlan {
	createWebAuthWlanType := new(CreateWebAuthWlan)
	return createWebAuthWlanType
}

func NewCreateWebAuthWlanWithDefaults() *CreateWebAuthWlan {
	createWebAuthWlanType := new(CreateWebAuthWlan)
	bypassCNAField := false
	createWebAuthWlanType.BypassCNA = &bypassCNAField
	return createWebAuthWlanType
}

type CreateWechatWlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, RuckusGRE means RuckusGRE tunnel to the data plane, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile,omitempty"`

	// AwsExtNasIPEnable
	// Aws ExtNasIP Enable
	AwsExtNasIPEnable *bool `json:"awsExtNasIPEnable,omitempty"`

	// AwsVenueEnable
	// Aws Venue Enable
	AwsVenueEnable *bool `json:"awsVenueEnable,omitempty"`

	// CaleaEnabled
	// DP CALEA Server Enabled
	CaleaEnabled *bool `json:"caleaEnabled,omitempty"`

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WlanNameSSID `json:"name" validate:"required"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	// PortalServiceProfile
	// Constraints:
	//    - required
	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile" validate:"required"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// Constraints:
	//    - required
	Ssid *WlanNameSSID `json:"ssid" validate:"required"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewCreateWechatWlan() *CreateWechatWlan {
	createWechatWlanType := new(CreateWechatWlan)
	return createWechatWlanType
}

func NewCreateWechatWlanWithDefaults() *CreateWechatWlan {
	createWechatWlanType := new(CreateWechatWlan)
	return createWechatWlanType
}

type ModifyWlan struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, and SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile,omitempty"`

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

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	ExternalDpsk *dpsk.WlanExternalDpsk `json:"externalDpsk,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	Name *WlanNameSSID `json:"name,omitempty"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	Ssid *WlanNameSSID `json:"ssid,omitempty"`

	Vlan *WlanVlan `json:"vlan,omitempty"`
}

func NewModifyWlan() *ModifyWlan {
	modifyWlanType := new(ModifyWlan)
	return modifyWlanType
}

func NewModifyWlanWithDefaults() *ModifyWlan {
	modifyWlanType := new(ModifyWlan)
	return modifyWlanType
}

type WlanAccounting struct {
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
	//    - nullable
	//    - min:0
	//    - max:1440
	InterimUpdateMin *int `json:"interimUpdateMin,omitempty" validate:"omitempty,gte=0,lte=1440"`

	// Name
	// Accounting service or profile name. At least one ID or name is required in the request.
	Name *string `json:"name,omitempty"`

	RealmBasedAcct *bool `json:"realmBasedAcct,omitempty"`

	// ThroughController
	// Indicates whether accounting messages were sent through the controller
	ThroughController *bool `json:"throughController,omitempty"`
}

func NewWlanAccounting() *WlanAccounting {
	wlanAccountingType := new(WlanAccounting)
	return wlanAccountingType
}

func NewWlanAccountingWithDefaults() *WlanAccounting {
	wlanAccountingType := new(WlanAccounting)
	return wlanAccountingType
}

type WlanAdvanced struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// AssocRssiThr
	// Assoc RSSI threshold.f
	// Constraints:
	//    - nullable
	//    - min:-90
	//    - max:-60
	AssocRssiThr *int `json:"assocRssiThr,omitempty" validate:"omitempty,gte=-90,lte=-60"`

	// AuthRssiThr
	// Auth RSSI threshold.
	// Constraints:
	//    - nullable
	//    - min:-90
	//    - max:-60
	AuthRssiThr *int `json:"authRssiThr,omitempty" validate:"omitempty,gte=-90,lte=-60"`

	// AvcEnabled
	// Indicator of whether ARC support is enabled or disabled
	AvcEnabled *bool `json:"avcEnabled,omitempty"`

	// BandBalancing
	// Indicates whether band balancing is enabled or disabled
	// Constraints:
	//    - nullable
	//    - default:'UseZoneSetting'
	//    - oneof:[Disabled,UseZoneSetting]
	BandBalancing *string `json:"bandBalancing,omitempty" validate:"omitempty,oneof=Disabled UseZoneSetting"`

	BssMinRateMbps *WlanBssMinRateMbps `json:"bssMinRateMbps,omitempty"`

	// ClientFingerprintingEnabled
	// Indicates whether client fingerprinting is enabled or disabled
	ClientFingerprintingEnabled *bool `json:"clientFingerprintingEnabled,omitempty"`

	// ClientIdleTimeoutSec
	// Client idle timeout in seconds
	// Constraints:
	//    - nullable
	//    - default:120
	//    - min:60
	//    - max:1000
	ClientIdleTimeoutSec *int `json:"clientIdleTimeoutSec,omitempty" validate:"omitempty,gte=60,lte=1000"`

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

	// Dhcp82Format
	// DHCP Option 82 format. This variable no longer supports from v8_1 and only kept for backward compatibility.
	// Constraints:
	//    - nullable
	//    - oneof:[RUCKUS_DEFAULT,RADIUS_DHCP,RADIUS_NAT,RADIUS_DHCP_NAT,RADIUS_NAT_RUCKUS,RADIUS_NAT_SOFTGRE,SOFTGRE_CUSTOMIZED]
	Dhcp82Format *string `json:"dhcp82Format,omitempty" validate:"omitempty,oneof=RUCKUS_DEFAULT RADIUS_DHCP RADIUS_NAT RADIUS_DHCP_NAT RADIUS_NAT_RUCKUS RADIUS_NAT_SOFTGRE SOFTGRE_CUSTOMIZED"`

	// Dhcp82SubOpt1Format
	// Subopt-1 format
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,SUBOPT1_AP_INFO_LOCATION,SUBOPT1_AP_INFO,SUBOPT1_AP_MAC_ESSID_PRIVACYTYPE,SUBOPT1_AP_MAC_hex,SUBOPT1_AP_MAC_hex_ESSID,SUBOPT1_ESSID]
	Dhcp82SubOpt1Format *string `json:"dhcp82SubOpt1Format,omitempty" validate:"omitempty,oneof=NONE SUBOPT1_AP_INFO_LOCATION SUBOPT1_AP_INFO SUBOPT1_AP_MAC_ESSID_PRIVACYTYPE SUBOPT1_AP_MAC_hex SUBOPT1_AP_MAC_hex_ESSID SUBOPT1_ESSID"`

	// Dhcp82SubOpt2Format
	// Subopt-2 format
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,SUBOPT2_CLIENT_MAC,SUBOPT2_CLIENT_MAC_hex,SUBOPT2_CLIENT_MAC_hex_ESSID,SUBOPT2_AP_MAC,SUBOPT2_AP_MAC_hex,SUBOPT2_AP_MAC_hex_ESSID]
	Dhcp82SubOpt2Format *string `json:"dhcp82SubOpt2Format,omitempty" validate:"omitempty,oneof=NONE SUBOPT2_CLIENT_MAC SUBOPT2_CLIENT_MAC_hex SUBOPT2_CLIENT_MAC_hex_ESSID SUBOPT2_AP_MAC SUBOPT2_AP_MAC_hex SUBOPT2_AP_MAC_hex_ESSID"`

	// Dhcp82SubOpt150Format
	// Subopt-150 with VLAN-Id
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,SUBOPT150_VLAN_ID]
	Dhcp82SubOpt150Format *string `json:"dhcp82SubOpt150Format,omitempty" validate:"omitempty,oneof=NONE SUBOPT150_VLAN_ID"`

	// Dhcp82SubOpt151AreaName
	// Subopt-151 Area Name value
	Dhcp82SubOpt151AreaName *string `json:"dhcp82SubOpt151AreaName,omitempty"`

	// Dhcp82SubOpt151Format
	// Subopt-151 format
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,SUBOPT151_AREA_NAME,SUBOPT151_ESSID]
	Dhcp82SubOpt151Format *string `json:"dhcp82SubOpt151Format,omitempty" validate:"omitempty,oneof=NONE SUBOPT151_AREA_NAME SUBOPT151_ESSID"`

	// Dhcp82SubOptRadiusFormat
	// Radius DHCP/NAT option 82 Format
	// Constraints:
	//    - nullable
	//    - oneof:[NONE,RUCKUS_DEFAULT,RADIUS_DHCP,RADIUS_NAT,RADIUS_DHCP_NAT,RADIUS_NAT_RUCKUS,RADIUS_NAT_SOFTGRE,SOFTGRE_CUSTOMIZED]
	Dhcp82SubOptRadiusFormat *string `json:"dhcp82SubOptRadiusFormat,omitempty" validate:"omitempty,oneof=NONE RUCKUS_DEFAULT RADIUS_DHCP RADIUS_NAT RADIUS_DHCP_NAT RADIUS_NAT_RUCKUS RADIUS_NAT_SOFTGRE SOFTGRE_CUSTOMIZED"`

	// DhcpOption82Enabled
	// Indicates whether DCHP Option 82 is enabled or disabled. This variable no longer supports from v8_1 and only kept for backward compatibility.
	DhcpOption82Enabled *bool `json:"dhcpOption82Enabled,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"omitempty,gte=0,lte=100"`

	// DirectedThreshold
	// Directed Threshold Setting, Defines the client count at which an AP will stop converting group addressed data traffic to unicast.
	// Constraints:
	//    - nullable
	//    - default:5
	//    - min:0
	//    - max:128
	DirectedThreshold *int `json:"directedThreshold,omitempty" validate:"omitempty,gte=0,lte=128"`

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
	//    - nullable
	//    - default:1
	//    - min:1
	//    - max:255
	DtimInterval *int `json:"dtimInterval,omitempty" validate:"omitempty,gte=1,lte=255"`

	EnableRadiusBasedDhcpNat *bool `json:"enableRadiusBasedDhcpNat,omitempty"`

	// FlowLogEnabled
	// Flow log enabled.
	FlowLogEnabled *bool `json:"flowLogEnabled,omitempty"`

	// ForceClientDHCPTimeoutSec
	// Force DHCP disconnects the client if the client does not obtain a valid IP address within the timeout peroid. To disable force DHCP, set this value to zero (0).
	// Constraints:
	//    - nullable
	//    - default:0
	//    - oneof:[0,5,6,7,8,9,10,11,12,13,14,15]
	ForceClientDHCPTimeoutSec *int `json:"forceClientDHCPTimeoutSec,omitempty" validate:"omitempty,oneof=0 5 6 7 8 9 10 11 12 13 14 15"`

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
	// Constraints:
	//    - nullable
	//    - default:300
	//    - min:1
	//    - max:300
	JoinAcceptTimeout *int `json:"joinAcceptTimeout,omitempty" validate:"omitempty,gte=1,lte=300"`

	// JoinIgnoreThr
	// Join wait threshold.
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:1
	//    - max:50
	JoinIgnoreThr *int `json:"joinIgnoreThr,omitempty" validate:"omitempty,gte=1,lte=50"`

	// JoinIgnoreTimeout
	// Join wait time.
	// Constraints:
	//    - nullable
	//    - default:30
	//    - min:1
	//    - max:60
	JoinIgnoreTimeout *int `json:"joinIgnoreTimeout,omitempty" validate:"omitempty,gte=1,lte=60"`

	// MaxAllowedRA
	// Max Allowed RAs
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:1
	//    - max:1440
	MaxAllowedRA *int `json:"maxAllowedRA,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// MaxClientsPerRadio
	// Maximum number of clients per radio
	// Constraints:
	//    - nullable
	//    - default:100
	//    - min:1
	//    - max:512
	MaxClientsPerRadio *int `json:"maxClientsPerRadio,omitempty" validate:"omitempty,gte=1,lte=512"`

	MgmtTxRateMbps *WlanMgmtTxRateMbps `json:"mgmtTxRateMbps,omitempty"`

	// NdProxyEnabled
	// Indicates whether ND Proxy is enabled or disabled
	NdProxyEnabled *bool `json:"ndProxyEnabled,omitempty"`

	// OceBroadcastProbeResponseDelay
	// Broadcast probe response delay.
	// Constraints:
	//    - nullable
	//    - default:15
	//    - min:8
	//    - max:120
	OceBroadcastProbeResponseDelay *int `json:"oceBroadcastProbeResponseDelay,omitempty" validate:"omitempty,gte=8,lte=120"`

	// OceEnabled
	// Optimized Connectivity Experience(OCE) enabled.
	OceEnabled *bool `json:"oceEnabled,omitempty"`

	// OceRssiBasedAssociationRejectionThreshold
	// RSSI-based association rejection threshold.
	// Constraints:
	//    - nullable
	//    - default:-75
	//    - min:-90
	//    - max:-60
	OceRssiBasedAssociationRejectionThreshold *int `json:"oceRssiBasedAssociationRejectionThreshold,omitempty" validate:"omitempty,gte=-90,lte=-60"`

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
	//    - nullable
	//    - default:'High'
	//    - oneof:[High,Low]
	Priority *string `json:"priority,omitempty" validate:"omitempty,oneof=High Low"`

	// ProbeRssiThr
	// Join RSSI threshold.
	// Constraints:
	//    - nullable
	//    - default:-85
	//    - min:-90
	//    - max:-60
	ProbeRssiThr *int `json:"probeRssiThr,omitempty" validate:"omitempty,gte=-90,lte=-60"`

	// ProxyARPEnabled
	// Indicates whether proxy ARP is enabled or disabled
	ProxyARPEnabled *bool `json:"proxyARPEnabled,omitempty"`

	// RaInterval
	// A timer that RA proxy runs and once receives unsolicited RA checks against the configured time and allow/drop RA based on next timeout
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:1
	//    - max:256
	RaInterval *int `json:"raInterval,omitempty" validate:"omitempty,gte=1,lte=256"`

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

	// WifiCallingPolicyEnabled
	// Indicator of whether Wi-Fi Calling is enabled or disabled
	WifiCallingPolicyEnabled *bool `json:"wifiCallingPolicyEnabled,omitempty"`

	// WifiCallingPolicyIds
	// The Wi-Fi Calling policy IDs. (Maximum allowed number is 5)
	WifiCallingPolicyIds []string `json:"wifiCallingPolicyIds,omitempty"`
}

func NewWlanAdvanced() *WlanAdvanced {
	wlanAdvancedType := new(WlanAdvanced)
	return wlanAdvancedType
}

func NewWlanAdvancedWithDefaults() *WlanAdvanced {
	wlanAdvancedType := new(WlanAdvanced)
	antiSpoofingEnabledField := false
	wlanAdvancedType.AntiSpoofingEnabled = &antiSpoofingEnabledField
	avcEnabledField := false
	wlanAdvancedType.AvcEnabled = &avcEnabledField
	bandBalancingField := `UseZoneSetting`
	wlanAdvancedType.BandBalancing = &bandBalancingField
	clientFingerprintingEnabledField := false
	wlanAdvancedType.ClientFingerprintingEnabled = &clientFingerprintingEnabledField
	clientIdleTimeoutSecField := 120
	wlanAdvancedType.ClientIdleTimeoutSec = &clientIdleTimeoutSecField
	clientIsolationAutoVrrpEnabledField := false
	wlanAdvancedType.ClientIsolationAutoVrrpEnabled = &clientIsolationAutoVrrpEnabledField
	clientIsolationEnabledField := false
	wlanAdvancedType.ClientIsolationEnabled = &clientIsolationEnabledField
	clientIsolationUnicastEnabledField := false
	wlanAdvancedType.ClientIsolationUnicastEnabled = &clientIsolationUnicastEnabledField
	clientLoadBalancingEnabledField := false
	wlanAdvancedType.ClientLoadBalancingEnabled = &clientLoadBalancingEnabledField
	dhcpOption82EnabledField := false
	wlanAdvancedType.DhcpOption82Enabled = &dhcpOption82EnabledField
	directedThresholdField := 5
	wlanAdvancedType.DirectedThreshold = &directedThresholdField
	downlinkEnabledField := false
	wlanAdvancedType.DownlinkEnabled = &downlinkEnabledField
	downlinkRateField := 0.000000
	wlanAdvancedType.DownlinkRate = &downlinkRateField
	dropRandomProbesEnabledField := false
	wlanAdvancedType.DropRandomProbesEnabled = &dropRandomProbesEnabledField
	dtimIntervalField := 1
	wlanAdvancedType.DtimInterval = &dtimIntervalField
	enableRadiusBasedDhcpNatField := false
	wlanAdvancedType.EnableRadiusBasedDhcpNat = &enableRadiusBasedDhcpNatField
	flowLogEnabledField := false
	wlanAdvancedType.FlowLogEnabled = &flowLogEnabledField
	forceClientDHCPTimeoutSecField := 0
	wlanAdvancedType.ForceClientDHCPTimeoutSec = &forceClientDHCPTimeoutSecField
	hdOverheadOptimizeEnableField := false
	wlanAdvancedType.HdOverheadOptimizeEnable = &hdOverheadOptimizeEnableField
	hideSsidEnabledField := false
	wlanAdvancedType.HideSsidEnabled = &hideSsidEnabledField
	hs20OnboardingField := false
	wlanAdvancedType.Hs20Onboarding = &hs20OnboardingField
	joinAcceptTimeoutField := 300
	wlanAdvancedType.JoinAcceptTimeout = &joinAcceptTimeoutField
	joinIgnoreThrField := 10
	wlanAdvancedType.JoinIgnoreThr = &joinIgnoreThrField
	joinIgnoreTimeoutField := 30
	wlanAdvancedType.JoinIgnoreTimeout = &joinIgnoreTimeoutField
	maxAllowedRAField := 10
	wlanAdvancedType.MaxAllowedRA = &maxAllowedRAField
	maxClientsPerRadioField := 100
	wlanAdvancedType.MaxClientsPerRadio = &maxClientsPerRadioField
	ndProxyEnabledField := false
	wlanAdvancedType.NdProxyEnabled = &ndProxyEnabledField
	oceBroadcastProbeResponseDelayField := 15
	wlanAdvancedType.OceBroadcastProbeResponseDelay = &oceBroadcastProbeResponseDelayField
	oceEnabledField := false
	wlanAdvancedType.OceEnabled = &oceEnabledField
	oceRssiBasedAssociationRejectionThresholdField := -75
	wlanAdvancedType.OceRssiBasedAssociationRejectionThreshold = &oceRssiBasedAssociationRejectionThresholdField
	ofdmOnlyEnabledField := false
	wlanAdvancedType.OfdmOnlyEnabled = &ofdmOnlyEnabledField
	priorityField := `High`
	wlanAdvancedType.Priority = &priorityField
	probeRssiThrField := -85
	wlanAdvancedType.ProbeRssiThr = &probeRssiThrField
	proxyARPEnabledField := false
	wlanAdvancedType.ProxyARPEnabled = &proxyARPEnabledField
	raIntervalField := 10
	wlanAdvancedType.RaInterval = &raIntervalField
	raProxyEnabledField := false
	wlanAdvancedType.RaProxyEnabled = &raProxyEnabledField
	raThrottlingEnabledField := false
	wlanAdvancedType.RaThrottlingEnabled = &raThrottlingEnabledField
	rsraGuardEnabledField := false
	wlanAdvancedType.RsraGuardEnabled = &rsraGuardEnabledField
	support80211dEnabledField := false
	wlanAdvancedType.Support80211dEnabled = &support80211dEnabledField
	support80211kEnabledField := false
	wlanAdvancedType.Support80211kEnabled = &support80211kEnabledField
	suppressNsEnabledField := false
	wlanAdvancedType.SuppressNsEnabled = &suppressNsEnabledField
	transientClientMgmtEnableField := false
	wlanAdvancedType.TransientClientMgmtEnable = &transientClientMgmtEnableField
	unauthClientStatsEnabledField := false
	wlanAdvancedType.UnauthClientStatsEnabled = &unauthClientStatsEnabledField
	uplinkEnabledField := false
	wlanAdvancedType.UplinkEnabled = &uplinkEnabledField
	uplinkRateField := 0.000000
	wlanAdvancedType.UplinkRate = &uplinkRateField
	urlFilteringPolicyEnabledField := false
	wlanAdvancedType.UrlFilteringPolicyEnabled = &urlFilteringPolicyEnabledField
	wifiCallingPolicyEnabledField := false
	wlanAdvancedType.WifiCallingPolicyEnabled = &wifiCallingPolicyEnabledField
	return wlanAdvancedType
}

type WlanAuthentication struct {
	// AuthenticationOption
	// Option of the authentication service or profile, At least one ID or name or authenticationOption is required in the request. This only applies to hotspot and guest WLANs.
	// Constraints:
	//    - nullable
	//    - oneof:[Local DB,Guest,Always Accept]
	AuthenticationOption *string `json:"authenticationOption,omitempty" validate:"omitempty,oneof=Local DB Guest Always Accept"`

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
	BackupAuthenticationOption *string `json:"backupAuthenticationOption,omitempty" validate:"omitempty,oneof=Always Accept"`

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

func NewWlanAuthentication() *WlanAuthentication {
	wlanAuthenticationType := new(WlanAuthentication)
	return wlanAuthenticationType
}

func NewWlanAuthenticationWithDefaults() *WlanAuthentication {
	wlanAuthenticationType := new(WlanAuthentication)
	return wlanAuthenticationType
}

type WlanBssMinRateMbps string

func NewWlanBssMinRateMbps() *WlanBssMinRateMbps {
	wlanBssMinRateMbpsType := new(WlanBssMinRateMbps)
	return wlanBssMinRateMbpsType
}

func NewWlanBssMinRateMbpsWithDefaults() *WlanBssMinRateMbps {
	wlanBssMinRateMbpsType := new(WlanBssMinRateMbps)
	*wlanBssMinRateMbpsType = `Disable`
	return wlanBssMinRateMbpsType
}

type WlanConfiguration struct {
	AccessIpsecProfile *common.GenericRef `json:"accessIpsecProfile,omitempty"`

	AccessTunnelProfile *common.GenericRef `json:"accessTunnelProfile,omitempty"`

	// AccessTunnelType
	// Access tunnel type of the WLAN. APLBO means AP local breakout, SoftGRE means AP direct SoftGRE tunnel
	// Constraints:
	//    - nullable
	//    - oneof:[APLBO,RuckusGRE,SoftGRE]
	AccessTunnelType *string `json:"accessTunnelType,omitempty" validate:"omitempty,oneof=APLBO RuckusGRE SoftGRE"`

	AccountingServiceOrProfile *WlanAccounting `json:"accountingServiceOrProfile,omitempty"`

	AdvancedOptions *WlanAdvanced `json:"advancedOptions,omitempty"`

	AuthServiceOrProfile *WlanAuthentication `json:"authServiceOrProfile,omitempty"`

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

	CoreTunnelProfile *WlanCoreTunnel `json:"coreTunnelProfile,omitempty"`

	DefaultUserTrafficProfile *common.GenericRef `json:"defaultUserTrafficProfile,omitempty"`

	// Description
	// Description of the WLAN
	Description *string `json:"description,omitempty"`

	DevicePolicy *common.GenericRef `json:"devicePolicy,omitempty"`

	DiffServProfile *common.GenericRef `json:"diffServProfile,omitempty"`

	DnsServerProfile *common.GenericRef `json:"dnsServerProfile,omitempty"`

	Dpsk *dpsk.WlanDpskSetting `json:"dpsk,omitempty"`

	// DpTunnelDhcpEnabled
	// DP Tunnel DHCP Enabled
	DpTunnelDhcpEnabled *bool `json:"dpTunnelDhcpEnabled,omitempty"`

	// DpTunnelNatEnabled
	// DP Tunnel NAT Enabled
	DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`

	Encryption *WlanEncryption `json:"encryption,omitempty"`

	FlexiVpnProfile *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`

	Hessid *WlanHESSID `json:"hessid,omitempty"`

	Hotspot20Profile *common.GenericRef `json:"hotspot20Profile,omitempty"`

	// Id
	// Identifier of the WLAN
	Id *string `json:"id,omitempty"`

	L2ACL *common.GenericRef `json:"l2ACL,omitempty"`

	MacAuth *WlanMACAuth `json:"macAuth,omitempty"`

	// Name
	// Name of the WLAN
	Name *string `json:"name,omitempty"`

	OperatorRealm *common.Realm `json:"operatorRealm,omitempty"`

	PortalDetectionProfileId *string `json:"portalDetectionProfileId,omitempty"`

	PortalServiceProfile *common.GenericRef `json:"portalServiceProfile,omitempty"`

	// PrecedenceProfileId
	// Precedence profile of the WLAN
	PrecedenceProfileId *string `json:"precedenceProfileId,omitempty"`

	// QosMaps
	// Qos map set of the WLAN.
	QosMaps []*WlanDSCPSetting `json:"qosMaps,omitempty"`

	RadiusOptions *WlanRadius `json:"radiusOptions,omitempty"`

	Schedule *WlanSchedule `json:"schedule,omitempty"`

	SplitTunnelProfileId *string `json:"splitTunnelProfileId,omitempty"`

	// Ssid
	// SSID of the WLAN
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the WLAN
	// Constraints:
	//    - nullable
	//    - oneof:[Standard_Open,Standard_8021X,Standard_Mac,Hotspot,Hotspot_MacByPass,Guest,WebAuth,Hotspot20,Hotspot20_Open,Hotspot20_OSEN]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=Standard_Open Standard_8021X Standard_Mac Hotspot Hotspot_MacByPass Guest WebAuth Hotspot20 Hotspot20_Open Hotspot20_OSEN"`

	Vlan *WlanVlan `json:"vlan,omitempty"`

	// ZoneId
	// Identifier of the zone to which the WLAN belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWlanConfiguration() *WlanConfiguration {
	wlanConfigurationType := new(WlanConfiguration)
	return wlanConfigurationType
}

func NewWlanConfigurationWithDefaults() *WlanConfiguration {
	wlanConfigurationType := new(WlanConfiguration)
	return wlanConfigurationType
}

type WlanCoreTunnel struct {
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
	Type *string `json:"type" validate:"required,oneof=L2oGRE Bridge TTG_PDG"`
}

func NewWlanCoreTunnel() *WlanCoreTunnel {
	wlanCoreTunnelType := new(WlanCoreTunnel)
	return wlanCoreTunnelType
}

func NewWlanCoreTunnelWithDefaults() *WlanCoreTunnel {
	wlanCoreTunnelType := new(WlanCoreTunnel)
	return wlanCoreTunnelType
}

type WlanDSCPSetting struct {
	// Enable
	// Enabled or disabled
	// Constraints:
	//    - required
	Enable *bool `json:"enable" validate:"required"`

	Excepts []int `json:"excepts,omitempty"`

	// High
	// DSCP range - high
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:255
	High *int `json:"high,omitempty" validate:"omitempty,gte=0,lte=255"`

	// Low
	// DSCP range - low
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:255
	Low *int `json:"low,omitempty" validate:"omitempty,gte=0,lte=255"`

	// Priority
	// Priority
	// Constraints:
	//    - required
	Priority *int `json:"priority" validate:"required"`
}

func NewWlanDSCPSetting() *WlanDSCPSetting {
	wlanDSCPSettingType := new(WlanDSCPSetting)
	return wlanDSCPSettingType
}

func NewWlanDSCPSettingWithDefaults() *WlanDSCPSetting {
	wlanDSCPSettingType := new(WlanDSCPSetting)
	return wlanDSCPSettingType
}

type WlanEncryption struct {
	// Algorithm
	// Encryption algorithm. This only applies to WPA2 and WPA mixed mode.
	// Constraints:
	//    - nullable
	//    - oneof:[AES,TKIP_AES]
	Algorithm *string `json:"algorithm,omitempty" validate:"omitempty,oneof=AES TKIP_AES"`

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
	//    - oneof:[WPA2,WPA_Mixed,WEP_64,WEP_128,None]
	Method *string `json:"method" validate:"required,oneof=WPA2 WPA_Mixed WEP_64 WEP_128 None"`

	// Mfp
	// Management frame protection. This only applies to WPA2 + AES
	// Constraints:
	//    - nullable
	//    - oneof:[disabled,capable,required]
	Mfp *string `json:"mfp,omitempty" validate:"omitempty,oneof=disabled capable required"`

	// MobilityDomainId
	// mobility Domain Id.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:65535
	MobilityDomainId *int `json:"mobilityDomainId,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// Passphrase
	// Passphrase. This only applies to WPA2 and WPA mixed mode.
	Passphrase *string `json:"passphrase,omitempty"`

	// Support80211rEnabled
	// Enable 802.11r Fast BSS Transition, fast Romaing.
	Support80211rEnabled *bool `json:"support80211rEnabled,omitempty"`
}

func NewWlanEncryption() *WlanEncryption {
	wlanEncryptionType := new(WlanEncryption)
	return wlanEncryptionType
}

func NewWlanEncryptionWithDefaults() *WlanEncryption {
	wlanEncryptionType := new(WlanEncryption)
	methodField := `None`
	wlanEncryptionType.Method = &methodField
	return wlanEncryptionType
}

type WlanHESSID string

func NewWlanHESSID() *WlanHESSID {
	wlanHESSIDType := new(WlanHESSID)
	return wlanHESSIDType
}

func NewWlanHESSIDWithDefaults() *WlanHESSID {
	wlanHESSIDType := new(WlanHESSID)
	return wlanHESSIDType
}

type WlanList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WlanSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWlanList() *WlanList {
	wlanListType := new(WlanList)
	return wlanListType
}

func NewWlanListWithDefaults() *WlanList {
	wlanListType := new(WlanList)
	return wlanListType
}

type WlanMACAuth struct {
	// CustomizedPassword
	// User defined password. When this field is set to an empty string, the MAC address is used as password.
	// Constraints:
	//    - nullable
	//    - max:64
	CustomizedPassword *string `json:"customizedPassword,omitempty" validate:"omitempty,max=64"`

	// MacAuthMacFormat
	// MAC address format. The default format is 0010a42319c0 and the 802.1X format is 00-10-A4-23-19-C0.
	// Constraints:
	//    - nullable
	//    - oneof:[Default,802.1X,UpperColon,Upper,LowerDash,LowerColon]
	MacAuthMacFormat *string `json:"macAuthMacFormat,omitempty" validate:"omitempty,oneof=Default 802.1X UpperColon Upper LowerDash LowerColon"`
}

func NewWlanMACAuth() *WlanMACAuth {
	wlanMACAuthType := new(WlanMACAuth)
	return wlanMACAuthType
}

func NewWlanMACAuthWithDefaults() *WlanMACAuth {
	wlanMACAuthType := new(WlanMACAuth)
	return wlanMACAuthType
}

type WlanMgmtTxRateMbps string

func NewWlanMgmtTxRateMbps() *WlanMgmtTxRateMbps {
	wlanMgmtTxRateMbpsType := new(WlanMgmtTxRateMbps)
	return wlanMgmtTxRateMbpsType
}

func NewWlanMgmtTxRateMbpsWithDefaults() *WlanMgmtTxRateMbps {
	wlanMgmtTxRateMbpsType := new(WlanMgmtTxRateMbps)
	*wlanMgmtTxRateMbpsType = `2 mbps`
	return wlanMgmtTxRateMbpsType
}

type WlanNameSSID string

func NewWlanNameSSID() *WlanNameSSID {
	wlanNameSSIDType := new(WlanNameSSID)
	return wlanNameSSIDType
}

func NewWlanNameSSIDWithDefaults() *WlanNameSSID {
	wlanNameSSIDType := new(WlanNameSSID)
	return wlanNameSSIDType
}

type WlanRadius struct {
	// CalledStaIdType
	// Called station ID type
	// Constraints:
	//    - nullable
	//    - default:'WLAN_BSSID'
	//    - oneof:[AP_MAC,NONE,WLAN_BSSID]
	CalledStaIdType *string `json:"calledStaIdType,omitempty" validate:"omitempty,oneof=AP_MAC NONE WLAN_BSSID"`

	// CustomizedNasId
	// User defined NAS ID
	// Constraints:
	//    - nullable
	//    - max:64
	CustomizedNasId *string `json:"customizedNasId,omitempty" validate:"omitempty,max=64"`

	// NasIdType
	// NAS ID type
	// Constraints:
	//    - nullable
	//    - default:'WLAN_BSSID'
	//    - oneof:[AP_MAC,Customized,WLAN_BSSID]
	NasIdType *string `json:"nasIdType,omitempty" validate:"omitempty,oneof=AP_MAC Customized WLAN_BSSID"`

	// NasIpType
	// NAS IP type
	// Constraints:
	//    - nullable
	//    - default:'disabled'
	//    - oneof:[disabled,control,management,userDefined]
	NasIpType *string `json:"nasIpType,omitempty" validate:"omitempty,oneof=disabled control management userDefined"`

	// NasIpUserDefined
	// User-defined NAS IP
	// Constraints:
	//    - nullable
	//    - max:45
	NasIpUserDefined *string `json:"nasIpUserDefined,omitempty" validate:"omitempty,max=45"`

	// NasMaxRetry
	// NAS request maximum retry
	// Constraints:
	//    - nullable
	//    - default:2
	//    - min:2
	//    - max:10
	NasMaxRetry *int `json:"nasMaxRetry,omitempty" validate:"omitempty,gte=2,lte=10"`

	// NasReconnectPrimaryMin
	// NAS reconnect primary time in minutes
	// Constraints:
	//    - nullable
	//    - default:5
	//    - min:1
	//    - max:60
	NasReconnectPrimaryMin *int `json:"nasReconnectPrimaryMin,omitempty" validate:"omitempty,gte=1,lte=60"`

	// NasRequestTimeoutSec
	// NAS request timeout in seconds
	// Constraints:
	//    - nullable
	//    - default:3
	//    - min:2
	//    - max:20
	NasRequestTimeoutSec *int `json:"nasRequestTimeoutSec,omitempty" validate:"omitempty,gte=2,lte=20"`

	// SingleSessionIdAcctEnabled
	// When Single Accounting Session ID is enabled, APs will maintain one accounting session for client roaming
	SingleSessionIdAcctEnabled *bool `json:"singleSessionIdAcctEnabled,omitempty"`

	// VendorSpecificAttributeProfileId
	// Vendor Specific Attribute Profile ID
	VendorSpecificAttributeProfileId *string `json:"vendorSpecificAttributeProfileId,omitempty"`
}

func NewWlanRadius() *WlanRadius {
	wlanRadiusType := new(WlanRadius)
	return wlanRadiusType
}

func NewWlanRadiusWithDefaults() *WlanRadius {
	wlanRadiusType := new(WlanRadius)
	calledStaIdTypeField := `WLAN_BSSID`
	wlanRadiusType.CalledStaIdType = &calledStaIdTypeField
	nasIdTypeField := `WLAN_BSSID`
	wlanRadiusType.NasIdType = &nasIdTypeField
	nasIpTypeField := `disabled`
	wlanRadiusType.NasIpType = &nasIpTypeField
	nasMaxRetryField := 2
	wlanRadiusType.NasMaxRetry = &nasMaxRetryField
	nasReconnectPrimaryMinField := 5
	wlanRadiusType.NasReconnectPrimaryMin = &nasReconnectPrimaryMinField
	nasRequestTimeoutSecField := 3
	wlanRadiusType.NasRequestTimeoutSec = &nasRequestTimeoutSecField
	singleSessionIdAcctEnabledField := false
	wlanRadiusType.SingleSessionIdAcctEnabled = &singleSessionIdAcctEnabledField
	return wlanRadiusType
}

type WlanSchedule struct {
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
	Type *string `json:"type" validate:"required,oneof=AlwaysOn AlwaysOff Customized"`
}

func NewWlanSchedule() *WlanSchedule {
	wlanScheduleType := new(WlanSchedule)
	return wlanScheduleType
}

func NewWlanScheduleWithDefaults() *WlanSchedule {
	wlanScheduleType := new(WlanSchedule)
	typeField := `AlwaysOn`
	wlanScheduleType.Type = &typeField
	return wlanScheduleType
}

type WlanSummary struct {
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

func NewWlanSummary() *WlanSummary {
	wlanSummaryType := new(WlanSummary)
	return wlanSummaryType
}

func NewWlanSummaryWithDefaults() *WlanSummary {
	wlanSummaryType := new(WlanSummary)
	return wlanSummaryType
}

type WlanVlan struct {
	// AaaVlanOverride
	// Indicates whether the AAA VLAN settings can be overriden or not
	AaaVlanOverride *bool `json:"aaaVlanOverride,omitempty"`

	// AccessVlan
	// Access VLAN ID
	// Constraints:
	//    - nullable
	//    - default:1
	//    - min:1
	//    - max:4094
	AccessVlan *int `json:"accessVlan,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// CoreQinQEnabled
	// Indicates whether Q-in-Q is allowed at the core network or not
	CoreQinQEnabled *bool `json:"coreQinQEnabled,omitempty"`

	// CoreSVlan
	// Core SVLAN ID. This only applies when core Q-in-Q is enabled
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	CoreSVlan *int `json:"coreSVlan,omitempty" validate:"omitempty,gte=1,lte=4094"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

func NewWlanVlan() *WlanVlan {
	wlanVlanType := new(WlanVlan)
	return wlanVlanType
}

func NewWlanVlanWithDefaults() *WlanVlan {
	wlanVlanType := new(WlanVlan)
	accessVlanField := 1
	wlanVlanType.AccessVlan = &accessVlanField
	return wlanVlanType
}
