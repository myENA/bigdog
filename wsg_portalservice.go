package bigdog

// API Version: v9_0

// WSGPortalServiceConnectionCapability
//
// Definition: portalservice_connectionCapability
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

// WSGPortalServiceCreateGuestAccess
//
// Definition: portalservice_createGuestAccess
type WSGPortalServiceCreateGuestAccess struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PortalCustomization
	// Constraints:
	//    - required
	PortalCustomization *WSGCommonPortalCustomization `json:"portalCustomization"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	SmsGateway *WSGCommonGenericRef `json:"smsGateway,omitempty"`

	// UserSession
	// Constraints:
	//    - required
	UserSession *WSGPortalServiceUserSession `json:"userSession"`
}

func NewWSGPortalServiceCreateGuestAccess() *WSGPortalServiceCreateGuestAccess {
	m := new(WSGPortalServiceCreateGuestAccess)
	return m
}

// WSGPortalServiceCreateHotspot20VenueProfile
//
// Definition: portalservice_createHotspot20VenueProfile
type WSGPortalServiceCreateHotspot20VenueProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	// VenueNames
	// Constraints:
	//    - required
	VenueNames []*WSGPortalServiceVenueName `json:"venueNames"`
}

func NewWSGPortalServiceCreateHotspot20VenueProfile() *WSGPortalServiceCreateHotspot20VenueProfile {
	m := new(WSGPortalServiceCreateHotspot20VenueProfile)
	return m
}

// WSGPortalServiceCreateHotspot20WlanProfile
//
// Definition: portalservice_createHotspot20WlanProfile
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
	DefaultIdentityProvider *WSGCommonGenericRef `json:"defaultIdentityProvider"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IdentityProviders
	// Ddentity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*WSGCommonGenericRef `json:"identityProviders,omitempty"`

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
	Name *WSGCommonNormalName `json:"name"`

	// Operator
	// Constraints:
	//    - required
	Operator *WSGCommonGenericRef `json:"operator"`

	SignupSsid *WSGCommonGenericRef `json:"signupSsid,omitempty"`
}

func NewWSGPortalServiceCreateHotspot20WlanProfile() *WSGPortalServiceCreateHotspot20WlanProfile {
	m := new(WSGPortalServiceCreateHotspot20WlanProfile)
	return m
}

// WSGPortalServiceCreateHotspotExternal
//
// Definition: portalservice_createHotspotExternal
type WSGPortalServiceCreateHotspotExternal struct {
	// BackupPortalUrl
	// Backup Portal URL of the Hotspot
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServicePortalLocation `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PortalUrl
	// Constraints:
	//    - required
	PortalUrl *WSGCommonNormalURL `json:"portalUrl"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

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

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceCreateHotspotExternal() *WSGPortalServiceCreateHotspotExternal {
	m := new(WSGPortalServiceCreateHotspotExternal)
	return m
}

// WSGPortalServiceCreateHotspotInternal
//
// Definition: portalservice_createHotspotInternal
type WSGPortalServiceCreateHotspotInternal struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	Location *WSGPortalServicePortalLocation `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - required
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport"`

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceCreateHotspotInternal() *WSGPortalServiceCreateHotspotInternal {
	m := new(WSGPortalServiceCreateHotspotInternal)
	return m
}

// WSGPortalServiceCreateHotspotSmartClientOnly
//
// Definition: portalservice_createHotspotSmartClientOnly
type WSGPortalServiceCreateHotspotSmartClientOnly struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServicePortalLocation `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

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

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceCreateHotspotSmartClientOnly() *WSGPortalServiceCreateHotspotSmartClientOnly {
	m := new(WSGPortalServiceCreateHotspotSmartClientOnly)
	return m
}

// WSGPortalServiceCreateL2ACL
//
// Definition: portalservice_createL2ACL
type WSGPortalServiceCreateL2ACL struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction"`

	RuleMacs []*WSGCommonMac `json:"ruleMacs,omitempty"`
}

func NewWSGPortalServiceCreateL2ACL() *WSGPortalServiceCreateL2ACL {
	m := new(WSGPortalServiceCreateL2ACL)
	return m
}

// WSGPortalServiceCreateWebAuthentication
//
// Definition: portalservice_createWebAuthentication
type WSGPortalServiceCreateWebAuthentication struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	WebAuthenticationPortalCustomization *WSGCommonWebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`
}

func NewWSGPortalServiceCreateWebAuthentication() *WSGPortalServiceCreateWebAuthentication {
	m := new(WSGPortalServiceCreateWebAuthentication)
	return m
}

// WSGPortalServiceCreateWechat
//
// Definition: portalservice_createWechat
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

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	// Constraints:
	//    - required
	DnatDestination *string `json:"dnatDestination"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*WSGPortalServiceDnatPortMapping `json:"dnatPortMapping,omitempty"`

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
	Name *WSGCommonNormalName `json:"name"`

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

// WSGPortalServiceDefaultConnectionCapability
//
// Definition: portalservice_defaultConnectionCapability
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

// WSGPortalServiceDnatPortMapping
//
// Definition: portalservice_dnatPortMapping
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

// WSGPortalServiceGuestAccess
//
// Definition: portalservice_guestAccess
type WSGPortalServiceGuestAccess struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the guest access profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PortalCustomization *WSGCommonPortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	SmsGateway *WSGCommonGenericRef `json:"smsGateway,omitempty"`

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	// ZoneId
	// Identifier of the zone which the guest access profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceGuestAccess() *WSGPortalServiceGuestAccess {
	m := new(WSGPortalServiceGuestAccess)
	return m
}

// WSGPortalServiceHotspot
//
// Definition: portalservice_hotspot
type WSGPortalServiceHotspot struct {
	// BackupPortalUrl
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// Id
	// Identifier of the Hotspot
	Id *string `json:"id,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServicePortalLocation `json:"location,omitempty"`

	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PortalCustomization *WSGCommonPortalCustomization `json:"portalCustomization,omitempty"`

	// PortalType
	// Portal type of the Hotspot
	// Constraints:
	//    - oneof:[Internal,External]
	PortalType *string `json:"portalType,omitempty"`

	PortalUrl *WSGCommonNormalURL `json:"portalUrl,omitempty"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

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

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

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

// WSGPortalServiceHotspot20VeuneProfile
//
// Definition: portalservice_hotspot20VeuneProfile
type WSGPortalServiceHotspot20VeuneProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty"`

	// Id
	// Identifier of the Hotspot 2.0 venue profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*WSGPortalServiceVenueName `json:"venueNames,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 venue profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceHotspot20VeuneProfile() *WSGPortalServiceHotspot20VeuneProfile {
	m := new(WSGPortalServiceHotspot20VeuneProfile)
	return m
}

// WSGPortalServiceHotspot20WlanProfile
//
// Definition: portalservice_hotspot20WlanProfile
type WSGPortalServiceHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*WSGPortalServiceDefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*WSGPortalServiceConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *WSGCommonGenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the Hotspot 2.0 WLAN profile
	Id *string `json:"id,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*WSGCommonGenericRef `json:"identityProviders,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Operator *WSGCommonGenericRef `json:"operator,omitempty"`

	SignupSsid *WSGCommonGenericRef `json:"signupSsid,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 WLAN profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceHotspot20WlanProfile() *WSGPortalServiceHotspot20WlanProfile {
	m := new(WSGPortalServiceHotspot20WlanProfile)
	return m
}

// WSGPortalServiceL2ACL
//
// Definition: portalservice_l2ACL
type WSGPortalServiceL2ACL struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// identifier of the L2 Access Control
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty"`

	RuleMacs []*WSGCommonMac `json:"ruleMacs,omitempty"`

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
// Definition: portalservice_linkSpeedInKbps
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
// Definition: portalservice_macAddressFormatSetting
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

// WSGPortalServiceModifyGuestAccess
//
// Definition: portalservice_modifyGuestAccess
type WSGPortalServiceModifyGuestAccess struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PortalCustomization *WSGCommonPortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	SmsGateway *WSGCommonGenericRef `json:"smsGateway,omitempty"`

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`
}

func NewWSGPortalServiceModifyGuestAccess() *WSGPortalServiceModifyGuestAccess {
	m := new(WSGPortalServiceModifyGuestAccess)
	return m
}

// WSGPortalServiceModifyHotspot
//
// Definition: portalservice_modifyHotspot
type WSGPortalServiceModifyHotspot struct {
	// BackupPortalUrl
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *WSGPortalServicePortalLocation `json:"location,omitempty"`

	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PortalCustomization *WSGCommonPortalCustomization `json:"portalCustomization,omitempty"`

	PortalUrl *WSGCommonNormalURL `json:"portalUrl,omitempty"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

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

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceModifyHotspot() *WSGPortalServiceModifyHotspot {
	m := new(WSGPortalServiceModifyHotspot)
	return m
}

// WSGPortalServiceModifyHotspot20VenueProfile
//
// Definition: portalservice_modifyHotspot20VenueProfile
type WSGPortalServiceModifyHotspot20VenueProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*WSGPortalServiceVenueName `json:"venueNames,omitempty"`
}

func NewWSGPortalServiceModifyHotspot20VenueProfile() *WSGPortalServiceModifyHotspot20VenueProfile {
	m := new(WSGPortalServiceModifyHotspot20VenueProfile)
	return m
}

// WSGPortalServiceModifyHotspot20WlanProfile
//
// Definition: portalservice_modifyHotspot20WlanProfile
type WSGPortalServiceModifyHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*WSGPortalServiceDefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*WSGPortalServiceConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *WSGCommonGenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*WSGCommonGenericRef `json:"identityProviders,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Operator *WSGCommonGenericRef `json:"operator,omitempty"`

	SignupSsid *WSGCommonGenericRef `json:"signupSsid,omitempty"`
}

func NewWSGPortalServiceModifyHotspot20WlanProfile() *WSGPortalServiceModifyHotspot20WlanProfile {
	m := new(WSGPortalServiceModifyHotspot20WlanProfile)
	return m
}

// WSGPortalServiceModifyL2ACL
//
// Definition: portalservice_modifyL2ACL
type WSGPortalServiceModifyL2ACL struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty"`

	RuleMacs []*WSGCommonMac `json:"ruleMacs,omitempty"`
}

func NewWSGPortalServiceModifyL2ACL() *WSGPortalServiceModifyL2ACL {
	m := new(WSGPortalServiceModifyL2ACL)
	return m
}

// WSGPortalServiceModifyWebAuthentication
//
// Definition: portalservice_modifyWebAuthentication
type WSGPortalServiceModifyWebAuthentication struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PortalLanguage *WSGCommonPortalLanguage `json:"portalLanguage,omitempty"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *WSGCommonWebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`
}

func NewWSGPortalServiceModifyWebAuthentication() *WSGPortalServiceModifyWebAuthentication {
	m := new(WSGPortalServiceModifyWebAuthentication)
	return m
}

// WSGPortalServiceModifyWechat
//
// Definition: portalservice_modifyWechat
type WSGPortalServiceModifyWechat struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*WSGPortalServiceDnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

func NewWSGPortalServiceModifyWechat() *WSGPortalServiceModifyWechat {
	m := new(WSGPortalServiceModifyWechat)
	return m
}

// WSGPortalServicePortalLocation
//
// Definition: portalservice_portalLocation
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

// WSGPortalServicePortalRedirect
//
// Definition: portalservice_portalRedirect
type WSGPortalServicePortalRedirect struct {
	Url *WSGCommonNormalURL `json:"url,omitempty"`
}

func NewWSGPortalServicePortalRedirect() *WSGPortalServicePortalRedirect {
	m := new(WSGPortalServicePortalRedirect)
	return m
}

// WSGPortalServiceList
//
// Definition: portalservice_portalServiceList
type WSGPortalServiceList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGPortalServiceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGPortalServiceList() *WSGPortalServiceList {
	m := new(WSGPortalServiceList)
	return m
}

// WSGPortalServiceListType
//
// Definition: portalservice_portalServiceListType
type WSGPortalServiceListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGPortalServiceListType() *WSGPortalServiceListType {
	m := new(WSGPortalServiceListType)
	return m
}

// WSGPortalServiceUserSession
//
// Definition: portalservice_userSession
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

// WSGPortalServiceVenueName
//
// Definition: portalservice_venueName
type WSGPortalServiceVenueName struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonLanguageName `json:"language"`

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

// WSGPortalServiceWebAuthentication
//
// Definition: portalservice_webAuthentication
type WSGPortalServiceWebAuthentication struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the web authentication profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PortalLanguage *WSGCommonPortalLanguage `json:"portalLanguage,omitempty"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *WSGCommonWebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`

	// ZoneId
	// Identifier of the zone which the web authentication profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalServiceWebAuthentication() *WSGPortalServiceWebAuthentication {
	m := new(WSGPortalServiceWebAuthentication)
	return m
}

// WSGPortalServiceWechatConfiguration
//
// Definition: portalservice_wechatConfiguration
type WSGPortalServiceWechatConfiguration struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*WSGPortalServiceDnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

func NewWSGPortalServiceWechatConfiguration() *WSGPortalServiceWechatConfiguration {
	m := new(WSGPortalServiceWechatConfiguration)
	return m
}
