package wlan

// API Version: v8_0

type CreateGuestAccessWLAN struct {
	AccessIpsecProfile         *common.GenericRef        `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef        `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                   `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting           `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced             `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication       `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                     `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                     `json:"awsVenueEnable,omitempty"`
	BypassCNA                  *bool                     `json:"bypassCNA,omitempty"`
	CaleaEnabled               *bool                     `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel           `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef        `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                   `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef        `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef        `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef        `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting     `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                     `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                     `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption           `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk    `json:"externalDpsk,omitempty"`
	FlexiVpnProfile            *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`
	Hessid                     *string                   `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef        `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef        `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth              `json:"macAuth,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	OperatorRealm              *string                   `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                   `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef        `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                   `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting        `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius               `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule             `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                   `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                   `json:"ssid,omitempty"`
	Vlan                       *WLANVlan                 `json:"vlan,omitempty"`
}

type CreateHotspot20OpenWLAN struct {
	AccessIpsecProfile         *common.GenericRef     `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef     `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting        `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced          `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication    `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                  `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                  `json:"awsVenueEnable,omitempty"`
	CaleaEnabled               *bool                  `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel        `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef     `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef     `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef     `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef     `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting  `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                  `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                  `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption        `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`
	Hessid                     *string                `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef     `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef     `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth           `json:"macAuth,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	OperatorRealm              *string                `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef     `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting     `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius            `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule          `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                `json:"ssid,omitempty"`
	Vlan                       *WLANVlan              `json:"vlan,omitempty"`
}

type CreateHotspot20OSENWLAN struct {
	*CreateHotspot20OpenWLAN
}

type CreateHotspot20WLAN struct {
	AccessIpsecProfile         *common.GenericRef     `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef     `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting        `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced          `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication    `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                  `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                  `json:"awsVenueEnable,omitempty"`
	CaleaEnabled               *bool                  `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel        `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef     `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef     `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef     `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef     `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting  `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                  `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                  `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption        `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`
	Hessid                     *string                `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef     `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef     `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth           `json:"macAuth,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	OperatorRealm              *string                `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef     `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting     `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius            `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule          `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                `json:"ssid,omitempty"`
	Vlan                       *WLANVlan              `json:"vlan,omitempty"`
}

type CreateHotspot8021XWLAN struct {
	*CreateHotspotWLAN
}

type CreateHotspotMacWLAN struct {
	*CreateHotspotWLAN
}

type CreateHotspotWLAN struct {
	AccessIpsecProfile         *common.GenericRef        `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef        `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                   `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting           `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced             `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication       `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                     `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                     `json:"awsVenueEnable,omitempty"`
	BypassCNA                  *bool                     `json:"bypassCNA,omitempty"`
	CaleaEnabled               *bool                     `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel           `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef        `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                   `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef        `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef        `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef        `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting     `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                     `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                     `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption           `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk    `json:"externalDpsk,omitempty"`
	FlexiVpnProfile            *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`
	Hessid                     *string                   `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef        `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef        `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth              `json:"macAuth,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	OperatorRealm              *string                   `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                   `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef        `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                   `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting        `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius               `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule             `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                   `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                   `json:"ssid,omitempty"`
	Vlan                       *WLANVlan                 `json:"vlan,omitempty"`
}

type CreateStandard80211WLAN struct {
	AccessIpsecProfile         *common.GenericRef        `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef        `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                   `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting           `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced             `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication       `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                     `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                     `json:"awsVenueEnable,omitempty"`
	CaleaEnabled               *bool                     `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel           `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef        `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                   `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef        `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef        `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef        `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting     `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                     `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                     `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption           `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk    `json:"externalDpsk,omitempty"`
	FlexiVpnProfile            *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`
	Hessid                     *string                   `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef        `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef        `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth              `json:"macAuth,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	OperatorRealm              *string                   `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                   `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef        `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                   `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting        `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius               `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule             `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                   `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                   `json:"ssid,omitempty"`
	Vlan                       *WLANVlan                 `json:"vlan,omitempty"`
}

type CreateStandard80211WLANMac struct {
	*CreateStandard80211WLAN
}

type CreateStandardMacWLAN struct {
	*CreateStandard80211WLAN
}

type CreateStandardOpenWLAN struct {
	AccessIpsecProfile         *common.GenericRef        `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef        `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                   `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting           `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced             `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication       `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                     `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                     `json:"awsVenueEnable,omitempty"`
	CaleaEnabled               *bool                     `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel           `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef        `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                   `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef        `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef        `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef        `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting     `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                     `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                     `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption           `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk    `json:"externalDpsk,omitempty"`
	FlexiVpnProfile            *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`
	Hessid                     *string                   `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef        `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef        `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth              `json:"macAuth,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	OperatorRealm              *string                   `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                   `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef        `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                   `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting        `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius               `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule             `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                   `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                   `json:"ssid,omitempty"`
	Vlan                       *WLANVlan                 `json:"vlan,omitempty"`
}

type CreateWebAuthWLAN struct {
	AccessIpsecProfile         *common.GenericRef        `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef        `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                   `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting           `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced             `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication       `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                     `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                     `json:"awsVenueEnable,omitempty"`
	BypassCNA                  *bool                     `json:"bypassCNA,omitempty"`
	CaleaEnabled               *bool                     `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel           `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef        `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                   `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef        `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef        `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef        `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting     `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                     `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                     `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption           `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk    `json:"externalDpsk,omitempty"`
	FlexiVpnProfile            *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`
	Hessid                     *string                   `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef        `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef        `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth              `json:"macAuth,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	OperatorRealm              *string                   `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                   `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef        `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                   `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting        `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius               `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule             `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                   `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                   `json:"ssid,omitempty"`
	Vlan                       *WLANVlan                 `json:"vlan,omitempty"`
}

type CreateWechatWLAN struct {
	AccessIpsecProfile         *common.GenericRef     `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef     `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting        `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced          `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication    `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                  `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                  `json:"awsVenueEnable,omitempty"`
	CaleaEnabled               *bool                  `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel        `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef     `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef     `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef     `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef     `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting  `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                  `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                  `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption        `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk `json:"externalDpsk,omitempty"`
	Hessid                     *string                `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef     `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef     `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth           `json:"macAuth,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	OperatorRealm              *string                `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef     `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting     `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius            `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule          `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                `json:"ssid,omitempty"`
	Vlan                       *WLANVlan              `json:"vlan,omitempty"`
}

type ModifyWLAN struct {
	AccessIpsecProfile         *common.GenericRef        `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef        `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                   `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting           `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced             `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication       `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                     `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                     `json:"awsVenueEnable,omitempty"`
	BypassCNA                  *bool                     `json:"bypassCNA,omitempty"`
	CaleaEnabled               *bool                     `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel           `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef        `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                   `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef        `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef        `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef        `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting     `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                     `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                     `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption           `json:"encryption,omitempty"`
	ExternalDpsk               *dpsk.WLANExternalDpsk    `json:"externalDpsk,omitempty"`
	FlexiVpnProfile            *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`
	Hessid                     *string                   `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef        `json:"hotspot20Profile,omitempty"`
	L2ACL                      *common.GenericRef        `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth              `json:"macAuth,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	OperatorRealm              *string                   `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                   `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef        `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                   `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting        `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius               `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule             `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                   `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                   `json:"ssid,omitempty"`
	Vlan                       *WLANVlan                 `json:"vlan,omitempty"`
}

type WLANAccounting struct {
	AccountingDelayEnabled *bool   `json:"accountingDelayEnabled,omitempty"`
	BackupAccountingID     *string `json:"backupAccountingId,omitempty"`
	BackupAccountingName   *string `json:"backupAccountingName,omitempty"`
	ID                     *string `json:"id,omitempty"`
	InterimUpdateMin       *int    `json:"interimUpdateMin,omitempty"`
	Name                   *string `json:"name,omitempty"`
	RealmBasedAcct         *bool   `json:"realmBasedAcct,omitempty"`
	ThroughController      *bool   `json:"throughController,omitempty"`
}

type WLANAdvanced struct {
	AntiSpoofingEnabled                       *bool              `json:"antiSpoofingEnabled,omitempty"`
	ArpRequestRateLimit                       *int               `json:"arpRequestRateLimit,omitempty"`
	AssocRssiThr                              *int               `json:"assocRssiThr,omitempty"`
	AuthRssiThr                               *int               `json:"authRssiThr,omitempty"`
	AvcEnabled                                *bool              `json:"avcEnabled,omitempty"`
	BandBalancing                             *string            `json:"bandBalancing,omitempty"`
	BssMinRateMbps                            *string            `json:"bssMinRateMbps,omitempty"`
	ClientFingerprintingEnabled               *bool              `json:"clientFingerprintingEnabled,omitempty"`
	ClientIdleTimeoutSec                      *int               `json:"clientIdleTimeoutSec,omitempty"`
	ClientIsolationAutoVrrpEnabled            *bool              `json:"clientIsolationAutoVrrpEnabled,omitempty"`
	ClientIsolationEnabled                    *bool              `json:"clientIsolationEnabled,omitempty"`
	ClientIsolationMulticastEnabled           *bool              `json:"clientIsolationMulticastEnabled,omitempty"`
	ClientIsolationUnicastEnabled             *bool              `json:"clientIsolationUnicastEnabled,omitempty"`
	ClientIsolationWhitelist                  *common.GenericRef `json:"clientIsolationWhitelist,omitempty"`
	ClientLoadBalancingEnabled                *bool              `json:"clientLoadBalancingEnabled,omitempty"`
	DgafEnabled                               *bool              `json:"dgafEnabled,omitempty"`
	DHCP82Format                              *string            `json:"dhcp82Format,omitempty"`
	DHCPOption82Enabled                       *bool              `json:"dhcpOption82Enabled,omitempty"`
	DHCPRequestRateLimit                      *int               `json:"dhcpRequestRateLimit,omitempty"`
	DirectedThreshold                         *int               `json:"directedThreshold,omitempty"`
	DownlinkEnabled                           *bool              `json:"downlinkEnabled,omitempty"`
	DownlinkRate                              *float64           `json:"downlinkRate,omitempty"`
	DtimInterval                              *int               `json:"dtimInterval,omitempty"`
	EnableRadiusBasedDHCPNat                  *bool              `json:"enableRadiusBasedDhcpNat,omitempty"`
	FlowLogEnabled                            *bool              `json:"flowLogEnabled,omitempty"`
	ForceClientDHCPTimeoutSec                 *int               `json:"forceClientDHCPTimeoutSec,omitempty"`
	HdOverheadOptimizeEnable                  *bool              `json:"hdOverheadOptimizeEnable,omitempty"`
	HideSsidEnabled                           *bool              `json:"hideSsidEnabled,omitempty"`
	Hs20Onboarding                            *bool              `json:"hs20Onboarding,omitempty"`
	JoinAcceptTimeout                         *int               `json:"joinAcceptTimeout,omitempty"`
	JoinIgnoreThr                             *int               `json:"joinIgnoreThr,omitempty"`
	JoinIgnoreTimeout                         *int               `json:"joinIgnoreTimeout,omitempty"`
	MaxAllowedRA                              *int               `json:"maxAllowedRA,omitempty"`
	MaxClientsPerRadio                        *int               `json:"maxClientsPerRadio,omitempty"`
	MgmtTxRateMbps                            *string            `json:"mgmtTxRateMbps,omitempty"`
	NdProxyEnabled                            *bool              `json:"ndProxyEnabled,omitempty"`
	OceBroadcastProbeResponseDelay            *int               `json:"oceBroadcastProbeResponseDelay,omitempty"`
	OceEnabled                                *bool              `json:"oceEnabled,omitempty"`
	OceRssiBasedAssociationRejectionThreshold *int               `json:"oceRssiBasedAssociationRejectionThreshold,omitempty"`
	OfdmOnlyEnabled                           *bool              `json:"ofdmOnlyEnabled,omitempty"`
	OkcEnabled                                *bool              `json:"okcEnabled,omitempty"`
	PmkCachingEnabled                         *bool              `json:"pmkCachingEnabled,omitempty"`
	Priority                                  *string            `json:"priority,omitempty"`
	ProbeRssiThr                              *int               `json:"probeRssiThr,omitempty"`
	ProxyARPEnabled                           *bool              `json:"proxyARPEnabled,omitempty"`
	RaInterval                                *int               `json:"raInterval,omitempty"`
	RaProxyEnabled                            *bool              `json:"raProxyEnabled,omitempty"`
	RatePerSTADownlink                        *string            `json:"ratePerSTADownlink,omitempty"`
	RatePerSTAUplink                          *string            `json:"ratePerSTAUplink,omitempty"`
	RaThrottlingEnabled                       *bool              `json:"raThrottlingEnabled,omitempty"`
	RsraGuardEnabled                          *bool              `json:"rsraGuardEnabled,omitempty"`
	Support8021DEnabled                       *bool              `json:"support80211dEnabled,omitempty"`
	Support8021KEnabled                       *bool              `json:"support80211kEnabled,omitempty"`
	SuppressNsEnabled                         *bool              `json:"suppressNsEnabled,omitempty"`
	TransientClientMgmtEnable                 *bool              `json:"transientClientMgmtEnable,omitempty"`
	UnauthClientStatsEnabled                  *bool              `json:"unauthClientStatsEnabled,omitempty"`
	UplinkEnabled                             *bool              `json:"uplinkEnabled,omitempty"`
	UplinkRate                                *float64           `json:"uplinkRate,omitempty"`
	URLFilteringPolicyEnabled                 *bool              `json:"urlFilteringPolicyEnabled,omitempty"`
	URLFilteringPolicyID                      *string            `json:"urlFilteringPolicyId,omitempty"`
	WifiCallingPolicyEnabled                  *bool              `json:"wifiCallingPolicyEnabled,omitempty"`
	WifiCallingPolicyIds                      []string           `json:"wifiCallingPolicyIds,omitempty"`
}

type WLANAuthentication struct {
	AuthenticationOption       *string `json:"authenticationOption,omitempty"`
	BackupAuthenticationID     *string `json:"backupAuthenticationId,omitempty"`
	BackupAuthenticationName   *string `json:"backupAuthenticationName,omitempty"`
	BackupAuthenticationOption *string `json:"backupAuthenticationOption,omitempty"`
	ID                         *string `json:"id,omitempty"`
	LocationDeliveryEnabled    *bool   `json:"locationDeliveryEnabled,omitempty"`
	Name                       *string `json:"name,omitempty"`
	RealmBasedAuth             *bool   `json:"realmBasedAuth,omitempty"`
	ThroughController          *bool   `json:"throughController,omitempty"`
}

type WLANConfiguration struct {
	AccessIpsecProfile         *common.GenericRef        `json:"accessIpsecProfile,omitempty"`
	AccessTunnelProfile        *common.GenericRef        `json:"accessTunnelProfile,omitempty"`
	AccessTunnelType           *string                   `json:"accessTunnelType,omitempty"`
	AccountingServiceOrProfile *WLANAccounting           `json:"accountingServiceOrProfile,omitempty"`
	AdvancedOptions            *WLANAdvanced             `json:"advancedOptions,omitempty"`
	AuthServiceOrProfile       *WLANAuthentication       `json:"authServiceOrProfile,omitempty"`
	AwsExtNasIPEnable          *bool                     `json:"awsExtNasIPEnable,omitempty"`
	AwsVenueEnable             *bool                     `json:"awsVenueEnable,omitempty"`
	BypassCNA                  *bool                     `json:"bypassCNA,omitempty"`
	CaleaEnabled               *bool                     `json:"caleaEnabled,omitempty"`
	CoreTunnelProfile          *WLANCoreTunnel           `json:"coreTunnelProfile,omitempty"`
	DefaultUserTrafficProfile  *common.GenericRef        `json:"defaultUserTrafficProfile,omitempty"`
	Description                *string                   `json:"description,omitempty"`
	DevicePolicy               *common.GenericRef        `json:"devicePolicy,omitempty"`
	DiffServProfile            *common.GenericRef        `json:"diffServProfile,omitempty"`
	DNSServerProfile           *common.GenericRef        `json:"dnsServerProfile,omitempty"`
	Dpsk                       *dpsk.WLANDpskSetting     `json:"dpsk,omitempty"`
	DpTunnelDHCPEnabled        *bool                     `json:"dpTunnelDhcpEnabled,omitempty"`
	DpTunnelNatEnabled         *bool                     `json:"dpTunnelNatEnabled,omitempty"`
	Encryption                 *WLANEncryption           `json:"encryption,omitempty"`
	FlexiVpnProfile            *flexivpn.FlexiVpnSetting `json:"flexiVpnProfile,omitempty"`
	Hessid                     *string                   `json:"hessid,omitempty"`
	Hotspot20Profile           *common.GenericRef        `json:"hotspot20Profile,omitempty"`
	ID                         *string                   `json:"id,omitempty"`
	L2ACL                      *common.GenericRef        `json:"l2ACL,omitempty"`
	MacAuth                    *WLANMACAuth              `json:"macAuth,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	OperatorRealm              *string                   `json:"operatorRealm,omitempty"`
	PortalDetectionProfileID   *string                   `json:"portalDetectionProfileId,omitempty"`
	PortalServiceProfile       *common.GenericRef        `json:"portalServiceProfile,omitempty"`
	PrecedenceProfileID        *string                   `json:"precedenceProfileId,omitempty"`
	QosMaps                    []*WLANDSCPSetting        `json:"qosMaps,omitempty"`
	RadiusOptions              *WLANRadius               `json:"radiusOptions,omitempty"`
	Schedule                   *WLANSchedule             `json:"schedule,omitempty"`
	SplitTunnelProfileID       *string                   `json:"splitTunnelProfileId,omitempty"`
	Ssid                       *string                   `json:"ssid,omitempty"`
	Type                       *string                   `json:"type,omitempty"`
	Vlan                       *WLANVlan                 `json:"vlan,omitempty"`
	ZoneID                     *string                   `json:"zoneId,omitempty"`
}

type WLANCoreTunnel struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

type WLANDSCPSetting struct {
	Enable   *bool `json:"enable,omitempty"`
	Excepts  []int `json:"excepts,omitempty"`
	High     *int  `json:"high,omitempty"`
	Low      *int  `json:"low,omitempty"`
	Priority *int  `json:"priority,omitempty"`
}

type WLANEncryption struct {
	Algorithm           *string `json:"algorithm,omitempty"`
	KeyIndex            *int    `json:"keyIndex,omitempty"`
	KeyInHex            *string `json:"keyInHex,omitempty"`
	Method              *string `json:"method,omitempty"`
	Mfp                 *string `json:"mfp,omitempty"`
	MobilityDomainID    *int    `json:"mobilityDomainId,omitempty"`
	Passphrase          *string `json:"passphrase,omitempty"`
	Support8021REnabled *bool   `json:"support80211rEnabled,omitempty"`
}

type WLANList struct {
	FirstIndex *int           `json:"firstIndex,omitempty"`
	HasMore    *bool          `json:"hasMore,omitempty"`
	List       []*WLANSummary `json:"list,omitempty"`
	TotalCount *int           `json:"totalCount,omitempty"`
}

type WLANMACAuth struct {
	CustomizedPassword *string `json:"customizedPassword,omitempty"`
	MacAuthMacFormat   *string `json:"macAuthMacFormat,omitempty"`
}

type WLANRadius struct {
	CalledStaIDType            *string `json:"calledStaIdType,omitempty"`
	CustomizedNasID            *string `json:"customizedNasId,omitempty"`
	NasIDType                  *string `json:"nasIdType,omitempty"`
	NasIPType                  *string `json:"nasIpType,omitempty"`
	NasIPUserDefined           *string `json:"nasIpUserDefined,omitempty"`
	NasMaxRetry                *int    `json:"nasMaxRetry,omitempty"`
	NasReconnectPrimaryMin     *int    `json:"nasReconnectPrimaryMin,omitempty"`
	NasRequestTimeoutSec       *int    `json:"nasRequestTimeoutSec,omitempty"`
	SingleSessionIDAcctEnabled *bool   `json:"singleSessionIdAcctEnabled,omitempty"`
}

type WLANSchedule struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

type WLANSummary struct {
	ID     *string `json:"id,omitempty"`
	MvnoID *string `json:"mvnoId,omitempty"`
	Name   *string `json:"name,omitempty"`
	Ssid   *string `json:"ssid,omitempty"`
	ZoneID *string `json:"zoneId,omitempty"`
}

type WLANVlan struct {
	AaaVlanOverride *bool              `json:"aaaVlanOverride,omitempty"`
	AccessVlan      *int               `json:"accessVlan,omitempty"`
	CoreQinQEnabled *bool              `json:"coreQinQEnabled,omitempty"`
	CoreSVlan       *int               `json:"coreSVlan,omitempty"`
	VlanPooling     *common.GenericRef `json:"vlanPooling,omitempty"`
}
