package ruckus

// API Version: v9_0

type WSGPortalServiceConnectionCapability struct {
	// PortNumber
	// Port number of connection capability
	// Constraints:
	//    - required
	//    - min:0.000000
	//    - max:65535.000000
	PortNumber *float64 `json:"portNumber"`

	// ProtocolName
	// Protocol aame of connection capability
	// Constraints:
	//    - required
	ProtocolName *string `json:"protocolName"`

	// ProtocolNumber
	// Protocol number of connection capability
	// Constraints:
	//    - required
	//    - min:0.000000
	//    - max:254.000000
	ProtocolNumber *float64 `json:"protocolNumber"`

	// Status
	// Status of connection capability
	// Constraints:
	//    - required
	//    - oneof:[CLOSED,OPEN,UNKNOWN]
	Status *string `json:"status"`
}

func NewWSGPortalServiceConnectionCapability() *WSGPortalServiceConnectionCapability {
	m := new(WSGPortalServiceConnectionCapability)
	return m
}

type WSGPortalServiceCreateGuestAccess struct {
	Description *WSGPortalServiceCreateGuestAccess `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateGuestAccess `json:"name"`

	// PortalCustomization
	// Constraints:
	//    - required
	PortalCustomization *WSGPortalServiceCreateGuestAccess `json:"portalCustomization"`

	Redirect *WSGPortalServiceCreateGuestAccess `json:"redirect,omitempty"`

	SmsGateway *WSGPortalServiceCreateGuestAccess `json:"smsGateway,omitempty"`

	// UserSession
	// Constraints:
	//    - required
	UserSession *WSGPortalServiceCreateGuestAccess `json:"userSession"`
}

func NewWSGPortalServiceCreateGuestAccess() *WSGPortalServiceCreateGuestAccess {
	m := new(WSGPortalServiceCreateGuestAccess)
	return m
}

type WSGPortalServiceCreateHotspot20VenueProfile struct {
	Description *WSGPortalServiceCreateHotspot20VenueProfile `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceCreateHotspot20VenueProfile `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateHotspot20VenueProfile `json:"name"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *WSGPortalServiceCreateHotspot20VenueProfile `json:"uplinkSpeedInKbps,omitempty"`

	// VenueNames
	// Constraints:
	//    - required
	VenueNames []*WSGPortalServiceCreateHotspot20VenueProfile `json:"venueNames"`
}

func NewWSGPortalServiceCreateHotspot20VenueProfile() *WSGPortalServiceCreateHotspot20VenueProfile {
	m := new(WSGPortalServiceCreateHotspot20VenueProfile)
	return m
}

type WSGPortalServiceCreateHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType"`

	// DefaultIdentityProvider
	// Constraints:
	//    - required
	DefaultIdentityProvider *WSGPortalServiceCreateHotspot20WlanProfile `json:"defaultIdentityProvider"`

	Description *WSGPortalServiceCreateHotspot20WlanProfile `json:"description,omitempty"`

	// IdentityProviders
	// Ddentity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*WSGPortalServiceCreateHotspot20WlanProfile `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	InternetOption *bool `json:"internetOption"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	//    - oneof:[UNAVAILABLE,PUBLIC,PORT_RESTRICTED,SINGLE_NATED_PRIVATE,DOUBLE_NATED_PRIVATE,PORT_RESTRICTED_AND_SINGLE_NATED,PORT_RESTRICTED_AND_DOUBLE_NATED,UNKNOWN]
	Ipv4AddressType *string `json:"ipv4AddressType"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateHotspot20WlanProfile `json:"name"`

	// Operator
	// Constraints:
	//    - required
	Operator *WSGPortalServiceCreateHotspot20WlanProfile `json:"operator"`

	SignupSsid *WSGPortalServiceCreateHotspot20WlanProfile `json:"signupSsid,omitempty"`
}

func NewWSGPortalServiceCreateHotspot20WlanProfile() *WSGPortalServiceCreateHotspot20WlanProfile {
	m := new(WSGPortalServiceCreateHotspot20WlanProfile)
	return m
}

type WSGPortalServiceCreateHotspotExternal struct {
	// BackupPortalUrl
	// Backup Portal URL of the Hotspot
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *WSGPortalServiceCreateHotspotExternal `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServiceCreateHotspotExternal `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *WSGPortalServiceCreateHotspotExternal `json:"macAddressFormat"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateHotspotExternal `json:"name"`

	// PortalUrl
	// Constraints:
	//    - required
	PortalUrl *WSGPortalServiceCreateHotspotExternal `json:"portalUrl"`

	Redirect *WSGPortalServiceCreateHotspotExternal `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - required
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *WSGPortalServiceCreateHotspotExternal `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceCreateHotspotExternal() *WSGPortalServiceCreateHotspotExternal {
	m := new(WSGPortalServiceCreateHotspotExternal)
	return m
}

type WSGPortalServiceCreateHotspotInternal struct {
	Description *WSGPortalServiceCreateHotspotInternal `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	Location *WSGPortalServiceCreateHotspotInternal `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *WSGPortalServiceCreateHotspotInternal `json:"macAddressFormat"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateHotspotInternal `json:"name"`

	Redirect *WSGPortalServiceCreateHotspotInternal `json:"redirect,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - required
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport"`

	UserSession *WSGPortalServiceCreateHotspotInternal `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceCreateHotspotInternal() *WSGPortalServiceCreateHotspotInternal {
	m := new(WSGPortalServiceCreateHotspotInternal)
	return m
}

type WSGPortalServiceCreateHotspotSmartClientOnly struct {
	Description *WSGPortalServiceCreateHotspotSmartClientOnly `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServiceCreateHotspotSmartClientOnly `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *WSGPortalServiceCreateHotspotSmartClientOnly `json:"macAddressFormat"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateHotspotSmartClientOnly `json:"name"`

	Redirect *WSGPortalServiceCreateHotspotSmartClientOnly `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client application.
	// Constraints:
	//    - required
	SmartClientInfo *string `json:"smartClientInfo"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *WSGPortalServiceCreateHotspotSmartClientOnly `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceCreateHotspotSmartClientOnly() *WSGPortalServiceCreateHotspotSmartClientOnly {
	m := new(WSGPortalServiceCreateHotspotSmartClientOnly)
	return m
}

type WSGPortalServiceCreateL2ACL struct {
	Description *WSGPortalServiceCreateL2ACL `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateL2ACL `json:"name"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction"`

	RuleMacs []*WSGPortalServiceCreateL2ACL `json:"ruleMacs,omitempty"`
}

func NewWSGPortalServiceCreateL2ACL() *WSGPortalServiceCreateL2ACL {
	m := new(WSGPortalServiceCreateL2ACL)
	return m
}

type WSGPortalServiceCreateWebAuthentication struct {
	Description *WSGPortalServiceCreateWebAuthentication `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateWebAuthentication `json:"name"`

	Redirect *WSGPortalServiceCreateWebAuthentication `json:"redirect,omitempty"`

	WebAuthenticationPortalCustomization *WSGPortalServiceCreateWebAuthentication `json:"webAuthenticationPortalCustomization,omitempty"`
}

func NewWSGPortalServiceCreateWebAuthentication() *WSGPortalServiceCreateWebAuthentication {
	m := new(WSGPortalServiceCreateWebAuthentication)
	return m
}

type WSGPortalServiceCreateWechat struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	// Constraints:
	//    - required
	AuthUrl *string `json:"authUrl"`

	// BlackList
	// Black list of the wechat profile
	// Constraints:
	//    - required
	BlackList *string `json:"blackList"`

	Description *WSGPortalServiceCreateWechat `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	// Constraints:
	//    - required
	DnatDestination *string `json:"dnatDestination"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*WSGPortalServiceCreateWechat `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGPortalServiceCreateWechat `json:"name"`

	// WhiteList
	// White list of the wechat profile
	// Constraints:
	//    - required
	WhiteList []string `json:"whiteList"`
}

func NewWSGPortalServiceCreateWechat() *WSGPortalServiceCreateWechat {
	m := new(WSGPortalServiceCreateWechat)
	return m
}

type WSGPortalServiceDefaultConnectionCapability struct {
	// PortNumber
	// Port number of connection capability, cannot be modified
	// Constraints:
	//    - required
	//    - min:0.000000
	//    - max:65535.000000
	PortNumber *float64 `json:"portNumber"`

	// ProtocolName
	// Protocol aame of connection capability, cannot be modified
	// Constraints:
	//    - required
	ProtocolName *string `json:"protocolName"`

	// ProtocolNumber
	// Protocol number of connection capability, cannot be modified
	// Constraints:
	//    - required
	//    - min:0.000000
	//    - max:254.000000
	ProtocolNumber *float64 `json:"protocolNumber"`

	// Status
	// Status of connection capability
	// Constraints:
	//    - required
	//    - oneof:[CLOSED,OPEN,UNKNOWN]
	Status *string `json:"status"`
}

func NewWSGPortalServiceDefaultConnectionCapability() *WSGPortalServiceDefaultConnectionCapability {
	m := new(WSGPortalServiceDefaultConnectionCapability)
	return m
}

type WSGPortalServiceDnatPortMapping struct {
	// DestPort
	// Destination port
	// Constraints:
	//    - min:0
	//    - max:65535
	DestPort *int `json:"destPort,omitempty"`

	// SourcePort
	// Source port
	// Constraints:
	//    - min:0
	//    - max:65535
	SourcePort *int `json:"sourcePort,omitempty"`
}

func NewWSGPortalServiceDnatPortMapping() *WSGPortalServiceDnatPortMapping {
	m := new(WSGPortalServiceDnatPortMapping)
	return m
}

type WSGPortalServiceGuestAccess struct {
	Description *WSGPortalServiceGuestAccess `json:"description,omitempty"`

	// Id
	// Identifier of the guest access profile
	Id *string `json:"id,omitempty"`

	Name *WSGPortalServiceGuestAccess `json:"name,omitempty"`

	PortalCustomization *WSGPortalServiceGuestAccess `json:"portalCustomization,omitempty"`

	Redirect *WSGPortalServiceGuestAccess `json:"redirect,omitempty"`

	SmsGateway *WSGPortalServiceGuestAccess `json:"smsGateway,omitempty"`

	UserSession *WSGPortalServiceGuestAccess `json:"userSession,omitempty"`

	// ZoneId
	// Identifier of the zone which the guest access profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceGuestAccess() *WSGPortalServiceGuestAccess {
	m := new(WSGPortalServiceGuestAccess)
	return m
}

type WSGPortalServiceHotspot struct {
	// BackupPortalUrl
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *WSGPortalServiceHotspot `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// Id
	// Identifier of the Hotspot
	Id *string `json:"id,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServiceHotspot `json:"location,omitempty"`

	MacAddressFormat *WSGPortalServiceHotspot `json:"macAddressFormat,omitempty"`

	Name *WSGPortalServiceHotspot `json:"name,omitempty"`

	PortalCustomization *WSGPortalServiceHotspot `json:"portalCustomization,omitempty"`

	// PortalType
	// Portal type of the Hotspot
	// Constraints:
	//    - oneof:[Internal,External]
	PortalType *string `json:"portalType,omitempty"`

	PortalUrl *WSGPortalServiceHotspot `json:"portalUrl,omitempty"`

	Redirect *WSGPortalServiceHotspot `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - oneof:[None,Enabled,SmartClientOnly]
	SmartClientSupport *string `json:"smartClientSupport,omitempty"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *WSGPortalServiceHotspot `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceHotspot() *WSGPortalServiceHotspot {
	m := new(WSGPortalServiceHotspot)
	return m
}

type WSGPortalServiceHotspot20VeuneProfile struct {
	Description *WSGPortalServiceHotspot20VeuneProfile `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceHotspot20VeuneProfile `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty"`

	// Id
	// Identifier of the Hotspot 2.0 venue profile
	Id *string `json:"id,omitempty"`

	Name *WSGPortalServiceHotspot20VeuneProfile `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *WSGPortalServiceHotspot20VeuneProfile `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*WSGPortalServiceHotspot20VeuneProfile `json:"venueNames,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 venue profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceHotspot20VeuneProfile() *WSGPortalServiceHotspot20VeuneProfile {
	m := new(WSGPortalServiceHotspot20VeuneProfile)
	return m
}

type WSGPortalServiceHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*WSGPortalServiceHotspot20WlanProfile `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*WSGPortalServiceHotspot20WlanProfile `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *WSGPortalServiceHotspot20WlanProfile `json:"defaultIdentityProvider,omitempty"`

	Description *WSGPortalServiceHotspot20WlanProfile `json:"description,omitempty"`

	// Id
	// Identifier of the Hotspot 2.0 WLAN profile
	Id *string `json:"id,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*WSGPortalServiceHotspot20WlanProfile `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the v WLAN profile
	// Constraints:
	//    - oneof:[UNAVAILABLE,PUBLIC,PORT_RESTRICTED,SINGLE_NATED_PRIVATE,DOUBLE_NATED_PRIVATE,PORT_RESTRICTED_AND_SINGLE_NATED,PORT_RESTRICTED_AND_DOUBLE_NATED,UNKNOWN]
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty"`

	Name *WSGPortalServiceHotspot20WlanProfile `json:"name,omitempty"`

	Operator *WSGPortalServiceHotspot20WlanProfile `json:"operator,omitempty"`

	SignupSsid *WSGPortalServiceHotspot20WlanProfile `json:"signupSsid,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 WLAN profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceHotspot20WlanProfile() *WSGPortalServiceHotspot20WlanProfile {
	m := new(WSGPortalServiceHotspot20WlanProfile)
	return m
}

type WSGPortalServiceL2ACL struct {
	Description *WSGPortalServiceL2ACL `json:"description,omitempty"`

	// Id
	// identifier of the L2 Access Control
	Id *string `json:"id,omitempty"`

	Name *WSGPortalServiceL2ACL `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty"`

	RuleMacs []*WSGPortalServiceL2ACL `json:"ruleMacs,omitempty"`

	// ZoneId
	// identifier of the zone which the L2 Access Control belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceL2ACL() *WSGPortalServiceL2ACL {
	m := new(WSGPortalServiceL2ACL)
	return m
}

// WSGPortalServiceLinkSpeedInKbps
//
// Link Speed of the Hotspot 2.0 venue profile
// Constraints:
//    - min:0.000000
//    - max:4294967295.000000
type WSGPortalServiceLinkSpeedInKbps float64

func NewWSGPortalServiceLinkSpeedInKbps() *WSGPortalServiceLinkSpeedInKbps {
	m := new(WSGPortalServiceLinkSpeedInKbps)
	return m
}

// WSGPortalServiceMacAddressFormatSetting
//
// mac address format of redirection,the format define: 0(aabbccddeeff), 1(AA-BB-CC-DD-EE-FF), 2(AA:BB:CC:DD:EE:FF), 3(AABBCCDDEEFF), 4(aa-bb-cc-dd-ee-ff), 5(aa:bb:cc:dd:ee:ff)
// Constraints:
//    - default:2
//    - min:0
//    - max:5
type WSGPortalServiceMacAddressFormatSetting int

func NewWSGPortalServiceMacAddressFormatSetting() *WSGPortalServiceMacAddressFormatSetting {
	m := new(WSGPortalServiceMacAddressFormatSetting)
	return m
}

type WSGPortalServiceModifyGuestAccess struct {
	Description *WSGPortalServiceModifyGuestAccess `json:"description,omitempty"`

	Name *WSGPortalServiceModifyGuestAccess `json:"name,omitempty"`

	PortalCustomization *WSGPortalServiceModifyGuestAccess `json:"portalCustomization,omitempty"`

	Redirect *WSGPortalServiceModifyGuestAccess `json:"redirect,omitempty"`

	SmsGateway *WSGPortalServiceModifyGuestAccess `json:"smsGateway,omitempty"`

	UserSession *WSGPortalServiceModifyGuestAccess `json:"userSession,omitempty"`
}

func NewWSGPortalServiceModifyGuestAccess() *WSGPortalServiceModifyGuestAccess {
	m := new(WSGPortalServiceModifyGuestAccess)
	return m
}

type WSGPortalServiceModifyHotspot struct {
	// BackupPortalUrl
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *WSGPortalServiceModifyHotspot `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServiceModifyHotspot `json:"location,omitempty"`

	MacAddressFormat *WSGPortalServiceModifyHotspot `json:"macAddressFormat,omitempty"`

	Name *WSGPortalServiceModifyHotspot `json:"name,omitempty"`

	PortalCustomization *WSGPortalServiceModifyHotspot `json:"portalCustomization,omitempty"`

	PortalUrl *WSGPortalServiceModifyHotspot `json:"portalUrl,omitempty"`

	Redirect *WSGPortalServiceModifyHotspot `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport,omitempty"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *WSGPortalServiceModifyHotspot `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceModifyHotspot() *WSGPortalServiceModifyHotspot {
	m := new(WSGPortalServiceModifyHotspot)
	return m
}

type WSGPortalServiceModifyHotspot20VenueProfile struct {
	Description *WSGPortalServiceModifyHotspot20VenueProfile `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceModifyHotspot20VenueProfile `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty"`

	Name *WSGPortalServiceModifyHotspot20VenueProfile `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *WSGPortalServiceModifyHotspot20VenueProfile `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*WSGPortalServiceModifyHotspot20VenueProfile `json:"venueNames,omitempty"`
}

func NewWSGPortalServiceModifyHotspot20VenueProfile() *WSGPortalServiceModifyHotspot20VenueProfile {
	m := new(WSGPortalServiceModifyHotspot20VenueProfile)
	return m
}

type WSGPortalServiceModifyHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*WSGPortalServiceModifyHotspot20WlanProfile `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*WSGPortalServiceModifyHotspot20WlanProfile `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *WSGPortalServiceModifyHotspot20WlanProfile `json:"defaultIdentityProvider,omitempty"`

	Description *WSGPortalServiceModifyHotspot20WlanProfile `json:"description,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*WSGPortalServiceModifyHotspot20WlanProfile `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 Wlan profile
	// Constraints:
	//    - oneof:[UNAVAILABLE,PUBLIC,PORT_RESTRICTED,SINGLE_NATED_PRIVATE,DOUBLE_NATED_PRIVATE,PORT_RESTRICTED_AND_SINGLE_NATED,PORT_RESTRICTED_AND_DOUBLE_NATED,UNKNOWN]
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 Wlan profile
	// Constraints:
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty"`

	Name *WSGPortalServiceModifyHotspot20WlanProfile `json:"name,omitempty"`

	Operator *WSGPortalServiceModifyHotspot20WlanProfile `json:"operator,omitempty"`

	SignupSsid *WSGPortalServiceModifyHotspot20WlanProfile `json:"signupSsid,omitempty"`
}

func NewWSGPortalServiceModifyHotspot20WlanProfile() *WSGPortalServiceModifyHotspot20WlanProfile {
	m := new(WSGPortalServiceModifyHotspot20WlanProfile)
	return m
}

type WSGPortalServiceModifyL2ACL struct {
	Description *WSGPortalServiceModifyL2ACL `json:"description,omitempty"`

	Name *WSGPortalServiceModifyL2ACL `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty"`

	RuleMacs []*WSGPortalServiceModifyL2ACL `json:"ruleMacs,omitempty"`
}

func NewWSGPortalServiceModifyL2ACL() *WSGPortalServiceModifyL2ACL {
	m := new(WSGPortalServiceModifyL2ACL)
	return m
}

type WSGPortalServiceModifyWebAuthentication struct {
	Description *WSGPortalServiceModifyWebAuthentication `json:"description,omitempty"`

	Name *WSGPortalServiceModifyWebAuthentication `json:"name,omitempty"`

	PortalLanguage *WSGPortalServiceModifyWebAuthentication `json:"portalLanguage,omitempty"`

	Redirect *WSGPortalServiceModifyWebAuthentication `json:"redirect,omitempty"`

	UserSession *WSGPortalServiceModifyWebAuthentication `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *WSGPortalServiceModifyWebAuthentication `json:"webAuthenticationPortalCustomization,omitempty"`
}

func NewWSGPortalServiceModifyWebAuthentication() *WSGPortalServiceModifyWebAuthentication {
	m := new(WSGPortalServiceModifyWebAuthentication)
	return m
}

type WSGPortalServiceModifyWechat struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *WSGPortalServiceModifyWechat `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*WSGPortalServiceModifyWechat `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty"`

	Name *WSGPortalServiceModifyWechat `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

func NewWSGPortalServiceModifyWechat() *WSGPortalServiceModifyWechat {
	m := new(WSGPortalServiceModifyWechat)
	return m
}

type WSGPortalServicePortalLocation struct {
	// Id
	// Portal location id
	Id *string `json:"id,omitempty"`

	// Name
	// Portal location name
	Name *string `json:"name,omitempty"`
}

func NewWSGPortalServicePortalLocation() *WSGPortalServicePortalLocation {
	m := new(WSGPortalServicePortalLocation)
	return m
}

type WSGPortalServicePortalRedirect struct {
	Url *WSGPortalServicePortalRedirect `json:"url,omitempty"`
}

func NewWSGPortalServicePortalRedirect() *WSGPortalServicePortalRedirect {
	m := new(WSGPortalServicePortalRedirect)
	return m
}

type WSGPortalServiceList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGPortalServiceList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGPortalServiceList() *WSGPortalServiceList {
	m := new(WSGPortalServiceList)
	return m
}

type WSGPortalServiceListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGPortalServiceListType `json:"name,omitempty"`
}

func NewWSGPortalServiceListType() *WSGPortalServiceListType {
	m := new(WSGPortalServiceListType)
	return m
}

type WSGPortalServiceUserSession struct {
	// GracePeriodInMin
	// Grace period in minutes
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriodInMin *int `json:"gracePeriodInMin,omitempty"`

	// TimeoutInMin
	// Time out value in minutes
	// Constraints:
	//    - default:1440
	//    - min:2
	//    - max:14400
	TimeoutInMin *int `json:"timeoutInMin,omitempty"`
}

func NewWSGPortalServiceUserSession() *WSGPortalServiceUserSession {
	m := new(WSGPortalServiceUserSession)
	return m
}

type WSGPortalServiceVenueName struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGPortalServiceVenueName `json:"language"`

	// Name
	// Venue name
	// Constraints:
	//    - required
	Name *string `json:"name"`
}

func NewWSGPortalServiceVenueName() *WSGPortalServiceVenueName {
	m := new(WSGPortalServiceVenueName)
	return m
}

type WSGPortalServiceWebAuthentication struct {
	Description *WSGPortalServiceWebAuthentication `json:"description,omitempty"`

	// Id
	// Identifier of the web authentication profile
	Id *string `json:"id,omitempty"`

	Name *WSGPortalServiceWebAuthentication `json:"name,omitempty"`

	PortalLanguage *WSGPortalServiceWebAuthentication `json:"portalLanguage,omitempty"`

	Redirect *WSGPortalServiceWebAuthentication `json:"redirect,omitempty"`

	UserSession *WSGPortalServiceWebAuthentication `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *WSGPortalServiceWebAuthentication `json:"webAuthenticationPortalCustomization,omitempty"`

	// ZoneId
	// Identifier of the zone which the web authentication profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceWebAuthentication() *WSGPortalServiceWebAuthentication {
	m := new(WSGPortalServiceWebAuthentication)
	return m
}

type WSGPortalServiceWechatConfiguration struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *WSGPortalServiceWechatConfiguration `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*WSGPortalServiceWechatConfiguration `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty"`

	Name *WSGPortalServiceWechatConfiguration `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

func NewWSGPortalServiceWechatConfiguration() *WSGPortalServiceWechatConfiguration {
	m := new(WSGPortalServiceWechatConfiguration)
	return m
}
