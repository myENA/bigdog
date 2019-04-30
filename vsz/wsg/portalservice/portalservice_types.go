package portalservice

// API Version: v8_0

type ConnectionCapability struct {
	// PortNumber
	// Port number of connection capability
	PortNumber *float64 `json:"portNumber,omitempty"`

	// ProtocolName
	// Protocol aame of connection capability
	ProtocolName *string `json:"protocolName,omitempty"`

	// ProtocolNumber
	// Protocol number of connection capability
	ProtocolNumber *float64 `json:"protocolNumber,omitempty"`

	// Status
	// Status of connection capability
	Status *string `json:"status,omitempty"`
}

type CreateGuestAccess struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`
}

type CreateHotspot20VenueProfile struct {
	Description *string `json:"description,omitempty"`

	DownlinkSpeedInKbps *float64 `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	Group *string `json:"group,omitempty"`

	Name *string `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *float64 `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames,omitempty"`
}

type CreateHotspot20WLANProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	AccessNetworkType *string `json:"accessNetworkType,omitempty"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *string `json:"description,omitempty"`

	// IdentityProviders
	// Ddentity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 WLAN profile
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty"`

	Name *string `json:"name,omitempty"`

	Operator *common.GenericRef `json:"operator,omitempty"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`
}

type CreateHotspotExternal struct {
	// BackupPortalURL
	// Backup Portal URL of the Hotspot
	BackupPortalURL *string `json:"backupPortalUrl,omitempty"`

	Description *string `json:"description,omitempty"`

	// HTTPSRedirect
	// HTTPS Redirect is disable or not
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *int `json:"macAddressFormat,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalURL *string `json:"portalUrl,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport,omitempty"`

	// TrafficClassProfileID
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileID *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type CreateHotspotInternal struct {
	Description *string `json:"description,omitempty"`

	// HTTPSRedirect
	// HTTPS Redirect is disable or not
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *int `json:"macAddressFormat,omitempty"`

	Name *string `json:"name,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport,omitempty"`

	// TrafficClassProfileID
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileID *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type CreateHotspotSmartClientOnly struct {
	Description *string `json:"description,omitempty"`

	// HTTPSRedirect
	// HTTPS Redirect is disable or not
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *int `json:"macAddressFormat,omitempty"`

	Name *string `json:"name,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client
	// application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// TrafficClassProfileID
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileID *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type CreateL2ACL struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all
	// stations listed below
	Restriction *string `json:"restriction,omitempty"`

	RuleMacs []string `json:"ruleMacs,omitempty"`
}

type CreateWebAuthentication struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`
}

type CreateWechat struct {
	// AuthURL
	// Authentication URL of the wechat profile
	AuthURL *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *string `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	GracePeriod *int `json:"gracePeriod,omitempty"`

	Name *string `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

type DefaultConnectionCapability struct {
	// PortNumber
	// Port number of connection capability, cannot be modified
	PortNumber *float64 `json:"portNumber,omitempty"`

	// ProtocolName
	// Protocol aame of connection capability, cannot be modified
	ProtocolName *string `json:"protocolName,omitempty"`

	// ProtocolNumber
	// Protocol number of connection capability, cannot be modified
	ProtocolNumber *float64 `json:"protocolNumber,omitempty"`

	// Status
	// Status of connection capability
	Status *string `json:"status,omitempty"`
}

type DnatPortMapping struct {
	// DestPort
	// Destination port
	DestPort *int `json:"destPort,omitempty"`

	// SourcePort
	// Source port
	SourcePort *int `json:"sourcePort,omitempty"`
}

type GuestAccess struct {
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the guest access profile
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// ZoneID
	// Identifier of the zone which the guest access profile belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type Hotspot struct {
	BackupPortalURL *string `json:"backupPortalUrl,omitempty"`

	Description *string `json:"description,omitempty"`

	// HTTPSRedirect
	// HTTPS Redirect is disable or not
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// ID
	// Identifier of the Hotspot
	ID *string `json:"id,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *int `json:"macAddressFormat,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	// PortalType
	// Portal type of the Hotspot
	PortalType *string `json:"portalType,omitempty"`

	PortalURL *string `json:"portalUrl,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client
	// application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport,omitempty"`

	// TrafficClassProfileID
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileID *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`

	// ZoneID
	// Identifier of the zone which the Hotspot belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type Hotspot20VeuneProfile struct {
	Description *string `json:"description,omitempty"`

	DownlinkSpeedInKbps *float64 `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	Group *string `json:"group,omitempty"`

	// ID
	// Identifier of the Hotspot 2.0 venue profile
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *float64 `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames,omitempty"`

	// ZoneID
	// Identifier of the zone which the Hotspot 2.0 venue profile belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type Hotspot20WLANProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	AccessNetworkType *string `json:"accessNetworkType,omitempty"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*ConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the Hotspot 2.0 WLAN profile
	ID *string `json:"id,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the v WLAN profile
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty"`

	Name *string `json:"name,omitempty"`

	Operator *common.GenericRef `json:"operator,omitempty"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`

	// ZoneID
	// Identifier of the zone which the Hotspot 2.0 WLAN profile belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type L2ACL struct {
	Description *string `json:"description,omitempty"`

	// ID
	// identifier of the L2 Access Control
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all
	// stations listed below
	Restriction *string `json:"restriction,omitempty"`

	RuleMacs []string `json:"ruleMacs,omitempty"`

	// ZoneID
	// identifier of the zone which the L2 Access Control belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

// LinkSpeedInKbps
//
// Link Speed of the Hotspot 2.0 venue profile
type LinkSpeedInKbps float64

// MacAddressFormatSetting
//
// mac address format of redirection,the format define: 0(aabbccddeeff), 1(AA-BB-CC-DD-EE-FF), 2(AA:BB:CC:DD:EE:FF),
// 3(AABBCCDDEEFF), 4(aa-bb-cc-dd-ee-ff), 5(aa:bb:cc:dd:ee:ff)
type MacAddressFormatSetting int

type ModifyGuestAccess struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`
}

type ModifyHotspot struct {
	BackupPortalURL *string `json:"backupPortalUrl,omitempty"`

	Description *string `json:"description,omitempty"`

	// HTTPSRedirect
	// HTTPS Redirect is disable or not
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *int `json:"macAddressFormat,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	PortalURL *string `json:"portalUrl,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client
	// application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport,omitempty"`

	// TrafficClassProfileID
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileID *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type ModifyHotspot20VenueProfile struct {
	Description *string `json:"description,omitempty"`

	DownlinkSpeedInKbps *float64 `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	Group *string `json:"group,omitempty"`

	Name *string `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *float64 `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames,omitempty"`
}

type ModifyHotspot20WLANProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	AccessNetworkType *string `json:"accessNetworkType,omitempty"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*ConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *string `json:"description,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 Wlan profile
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 Wlan profile
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty"`

	Name *string `json:"name,omitempty"`

	Operator *common.GenericRef `json:"operator,omitempty"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`
}

type ModifyL2ACL struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all
	// stations listed below
	Restriction *string `json:"restriction,omitempty"`

	RuleMacs []string `json:"ruleMacs,omitempty"`
}

type ModifyWebAuthentication struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalLanguage *string `json:"portalLanguage,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`
}

type ModifyWechat struct {
	// AuthURL
	// Authentication URL of the wechat profile
	AuthURL *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *string `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	GracePeriod *int `json:"gracePeriod,omitempty"`

	Name *string `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

type PortalLocation struct {
	// ID
	// Portal location id
	ID *string `json:"id,omitempty"`

	// Name
	// Portal location name
	Name *string `json:"name,omitempty"`
}

type PortalRedirect struct {
	URL *string `json:"url,omitempty"`
}

type PortalServiceList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortalServiceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type PortalServiceListType struct {
	// ID
	// Identifier of the service
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type UserSession struct {
	// GracePeriodInMin
	// Grace period in minutes
	GracePeriodInMin *int `json:"gracePeriodInMin,omitempty"`

	// TimeoutInMin
	// Time out value in minutes
	TimeoutInMin *int `json:"timeoutInMin,omitempty"`
}

type VenueName struct {
	Language *string `json:"language,omitempty"`

	// Name
	// Venue name
	Name *string `json:"name,omitempty"`
}

type WebAuthentication struct {
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the web authentication profile
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	PortalLanguage *string `json:"portalLanguage,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// ZoneID
	// Identifier of the zone which the web authentication profile belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type WechatConfiguration struct {
	// AuthURL
	// Authentication URL of the wechat profile
	AuthURL *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *string `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	GracePeriod *int `json:"gracePeriod,omitempty"`

	Name *string `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}
