package vsz

// API Version: v9_0

type WSGPortalServiceConnectionCapability struct {
	// PortNumber
	// Port number of connection capability
	// Constraints:
	//    - required
	//    - min:0.000000
	//    - max:65535.000000
	PortNumber *float64 `json:"portNumber" validate:"required,gte=0.000000,lte=65535.000000"`

	// ProtocolName
	// Protocol aame of connection capability
	// Constraints:
	//    - required
	ProtocolName *string `json:"protocolName" validate:"required"`

	// ProtocolNumber
	// Protocol number of connection capability
	// Constraints:
	//    - required
	//    - min:0.000000
	//    - max:254.000000
	ProtocolNumber *float64 `json:"protocolNumber" validate:"required,gte=0.000000,lte=254.000000"`

	// Status
	// Status of connection capability
	// Constraints:
	//    - required
	//    - oneof:[CLOSED,OPEN,UNKNOWN]
	Status *string `json:"status" validate:"required,oneof=CLOSED OPEN UNKNOWN"`
}

func NewWSGPortalServiceConnectionCapability() *WSGPortalServiceConnectionCapability {
	m := new(WSGPortalServiceConnectionCapability)
	return m
}

type WSGPortalServiceCreateGuestAccess struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// PortalCustomization
	// Constraints:
	//    - required
	PortalCustomization *WSGCommonPortalCustomization `json:"portalCustomization" validate:"required"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	SmsGateway *WSGCommonGenericRef `json:"smsGateway,omitempty"`

	// UserSession
	// Constraints:
	//    - required
	UserSession *WSGPortalServiceUserSession `json:"userSession" validate:"required"`
}

func NewWSGPortalServiceCreateGuestAccess() *WSGPortalServiceCreateGuestAccess {
	m := new(WSGPortalServiceCreateGuestAccess)
	return m
}

type WSGPortalServiceCreateHotspot20VenueProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty" validate:"oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	// VenueNames
	// Constraints:
	//    - required
	VenueNames []*WSGPortalServiceVenueName `json:"venueNames" validate:"required,dive"`
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
	AccessNetworkType *string `json:"accessNetworkType" validate:"required,oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

	// DefaultIdentityProvider
	// Constraints:
	//    - required
	DefaultIdentityProvider *WSGCommonGenericRef `json:"defaultIdentityProvider" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IdentityProviders
	// Ddentity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*WSGCommonGenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	InternetOption *bool `json:"internetOption" validate:"required"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	//    - oneof:[UNAVAILABLE,PUBLIC,PORT_RESTRICTED,SINGLE_NATED_PRIVATE,DOUBLE_NATED_PRIVATE,PORT_RESTRICTED_AND_SINGLE_NATED,PORT_RESTRICTED_AND_DOUBLE_NATED,UNKNOWN]
	Ipv4AddressType *string `json:"ipv4AddressType" validate:"required,oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType" validate:"required,oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Operator
	// Constraints:
	//    - required
	Operator *WSGCommonGenericRef `json:"operator" validate:"required"`

	SignupSsid *WSGCommonGenericRef `json:"signupSsid,omitempty"`
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
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty" validate:"omitempty"`

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
	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat" validate:"required,gte=0,lte=5"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// PortalUrl
	// Constraints:
	//    - required
	PortalUrl *WSGCommonNormalURL `json:"portalUrl" validate:"required"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - required
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport" validate:"required,oneof=None Enabled"`

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

type WSGPortalServiceCreateHotspotInternal struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	Location *WSGPortalServicePortalLocation `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat" validate:"required,gte=0,lte=5"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - required
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport" validate:"required,oneof=None Enabled"`

	UserSession *WSGPortalServiceUserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewWSGPortalServiceCreateHotspotInternal() *WSGPortalServiceCreateHotspotInternal {
	m := new(WSGPortalServiceCreateHotspotInternal)
	return m
}

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
	MacAddressFormat *WSGPortalServiceMacAddressFormatSetting `json:"macAddressFormat" validate:"required,gte=0,lte=5"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client application.
	// Constraints:
	//    - required
	SmartClientInfo *string `json:"smartClientInfo" validate:"required"`

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

type WSGPortalServiceCreateL2ACL struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction" validate:"required,oneof=ALLOW BLOCK"`

	RuleMacs []WSGCommonMac `json:"ruleMacs,omitempty"`
}

func NewWSGPortalServiceCreateL2ACL() *WSGPortalServiceCreateL2ACL {
	m := new(WSGPortalServiceCreateL2ACL)
	return m
}

type WSGPortalServiceCreateWebAuthentication struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	Redirect *WSGPortalServicePortalRedirect `json:"redirect,omitempty"`

	WebAuthenticationPortalCustomization *WSGCommonWebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`
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
	AuthUrl *string `json:"authUrl" validate:"required"`

	// BlackList
	// Black list of the wechat profile
	// Constraints:
	//    - required
	BlackList *string `json:"blackList" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	// Constraints:
	//    - required
	DnatDestination *string `json:"dnatDestination" validate:"required"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*WSGPortalServiceDnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"gte=1,lte=14399"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// WhiteList
	// White list of the wechat profile
	// Constraints:
	//    - required
	WhiteList []string `json:"whiteList" validate:"required,dive"`
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
	PortNumber *float64 `json:"portNumber" validate:"required,gte=0.000000,lte=65535.000000"`

	// ProtocolName
	// Protocol aame of connection capability, cannot be modified
	// Constraints:
	//    - required
	ProtocolName *string `json:"protocolName" validate:"required"`

	// ProtocolNumber
	// Protocol number of connection capability, cannot be modified
	// Constraints:
	//    - required
	//    - min:0.000000
	//    - max:254.000000
	ProtocolNumber *float64 `json:"protocolNumber" validate:"required,gte=0.000000,lte=254.000000"`

	// Status
	// Status of connection capability
	// Constraints:
	//    - required
	//    - oneof:[CLOSED,OPEN,UNKNOWN]
	Status *string `json:"status" validate:"required,oneof=CLOSED OPEN UNKNOWN"`
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
	DestPort *int `json:"destPort,omitempty" validate:"gte=0,lte=65535"`

	// SourcePort
	// Source port
	// Constraints:
	//    - min:0
	//    - max:65535
	SourcePort *int `json:"sourcePort,omitempty" validate:"gte=0,lte=65535"`
}

func NewWSGPortalServiceDnatPortMapping() *WSGPortalServiceDnatPortMapping {
	m := new(WSGPortalServiceDnatPortMapping)
	return m
}

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

type WSGPortalServiceHotspot struct {
	// BackupPortalUrl
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty" validate:"omitempty"`

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
	PortalType *string `json:"portalType,omitempty" validate:"oneof=Internal External"`

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
	SmartClientSupport *string `json:"smartClientSupport,omitempty" validate:"oneof=None Enabled SmartClientOnly"`

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

type WSGPortalServiceHotspot20VeuneProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty" validate:"oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

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

type WSGPortalServiceHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty" validate:"oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

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
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty" validate:"oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty" validate:"oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

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
	Restriction *string `json:"restriction,omitempty" validate:"oneof=ALLOW BLOCK"`

	RuleMacs []WSGCommonMac `json:"ruleMacs,omitempty"`

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

type WSGPortalServiceModifyHotspot struct {
	// BackupPortalUrl
	// Constraints:
	//    - nullable
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty" validate:"omitempty"`

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
	SmartClientSupport *string `json:"smartClientSupport,omitempty" validate:"oneof=None Enabled"`

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

type WSGPortalServiceModifyHotspot20VenueProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DownlinkSpeedInKbps *WSGPortalServiceLinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty" validate:"oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

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

type WSGPortalServiceModifyHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty" validate:"oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

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
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty" validate:"oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 Wlan profile
	// Constraints:
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty" validate:"oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Operator *WSGCommonGenericRef `json:"operator,omitempty"`

	SignupSsid *WSGCommonGenericRef `json:"signupSsid,omitempty"`
}

func NewWSGPortalServiceModifyHotspot20WlanProfile() *WSGPortalServiceModifyHotspot20WlanProfile {
	m := new(WSGPortalServiceModifyHotspot20WlanProfile)
	return m
}

type WSGPortalServiceModifyL2ACL struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty" validate:"oneof=ALLOW BLOCK"`

	RuleMacs []WSGCommonMac `json:"ruleMacs,omitempty"`
}

func NewWSGPortalServiceModifyL2ACL() *WSGPortalServiceModifyL2ACL {
	m := new(WSGPortalServiceModifyL2ACL)
	return m
}

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
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"gte=1,lte=14399"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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
	Url *WSGCommonNormalURL `json:"url,omitempty"`
}

func NewWSGPortalServicePortalRedirect() *WSGPortalServicePortalRedirect {
	m := new(WSGPortalServicePortalRedirect)
	return m
}

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

type WSGPortalServiceUserSession struct {
	// GracePeriodInMin
	// Grace period in minutes
	// Constraints:
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriodInMin *int `json:"gracePeriodInMin,omitempty" validate:"gte=1,lte=14399"`

	// TimeoutInMin
	// Time out value in minutes
	// Constraints:
	//    - default:1440
	//    - min:2
	//    - max:14400
	TimeoutInMin *int `json:"timeoutInMin,omitempty" validate:"gte=2,lte=14400"`
}

func NewWSGPortalServiceUserSession() *WSGPortalServiceUserSession {
	m := new(WSGPortalServiceUserSession)
	return m
}

type WSGPortalServiceVenueName struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonLanguageName `json:"language" validate:"required,oneof=English Chinese Czech Danish Dutch French German Japanese Spanish Korean Swedish Turkish eng chi cze dan dut fre ger jpn kor spa swe tur"`

	// Name
	// Venue name
	// Constraints:
	//    - required
	Name *string `json:"name" validate:"required"`
}

func NewWSGPortalServiceVenueName() *WSGPortalServiceVenueName {
	m := new(WSGPortalServiceVenueName)
	return m
}

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
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"gte=1,lte=14399"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

func NewWSGPortalServiceWechatConfiguration() *WSGPortalServiceWechatConfiguration {
	m := new(WSGPortalServiceWechatConfiguration)
	return m
}
