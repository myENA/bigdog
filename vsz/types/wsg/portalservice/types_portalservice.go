package portalservice

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ConnectionCapability struct {
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

func NewConnectionCapability() *ConnectionCapability {
	connectionCapabilityType := new(ConnectionCapability)
	return connectionCapabilityType
}

func NewDefaultConnectionCapability() *ConnectionCapability {
	connectionCapabilityType := new(ConnectionCapability)
	return connectionCapabilityType
}

type CreateGuestAccess struct {
	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PortalCustomization
	// Constraints:
	//    - required
	PortalCustomization *common.PortalCustomization `json:"portalCustomization" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	// UserSession
	// Constraints:
	//    - required
	UserSession *UserSession `json:"userSession" validate:"required"`
}

func NewCreateGuestAccess() *CreateGuestAccess {
	createGuestAccessType := new(CreateGuestAccess)
	return createGuestAccessType
}

func NewDefaultCreateGuestAccess() *CreateGuestAccess {
	createGuestAccessType := new(CreateGuestAccess)
	return createGuestAccessType
}

type CreateHotspot20VenueProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkSpeedInKbps *LinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - nullable
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty" validate:"omitempty,oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *LinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	// VenueNames
	// Constraints:
	//    - required
	VenueNames []*VenueName `json:"venueNames" validate:"required"`
}

func NewCreateHotspot20VenueProfile() *CreateHotspot20VenueProfile {
	createHotspot20VenueProfileType := new(CreateHotspot20VenueProfile)
	return createHotspot20VenueProfileType
}

func NewDefaultCreateHotspot20VenueProfile() *CreateHotspot20VenueProfile {
	createHotspot20VenueProfileType := new(CreateHotspot20VenueProfile)
	return createHotspot20VenueProfileType
}

type CreateHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - required
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType" validate:"required,oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

	// DefaultIdentityProvider
	// Constraints:
	//    - required
	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// IdentityProviders
	// Ddentity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

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
	Name *common.NormalName `json:"name" validate:"required"`

	// Operator
	// Constraints:
	//    - required
	Operator *common.GenericRef `json:"operator" validate:"required"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`
}

func NewCreateHotspot20WlanProfile() *CreateHotspot20WlanProfile {
	createHotspot20WlanProfileType := new(CreateHotspot20WlanProfile)
	return createHotspot20WlanProfileType
}

func NewDefaultCreateHotspot20WlanProfile() *CreateHotspot20WlanProfile {
	createHotspot20WlanProfileType := new(CreateHotspot20WlanProfile)
	return createHotspot20WlanProfileType
}

type CreateHotspotExternal struct {
	// BackupPortalUrl
	// Backup Portal URL of the Hotspot
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PortalUrl
	// Constraints:
	//    - required
	PortalUrl *common.NormalURL `json:"portalUrl" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

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

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewCreateHotspotExternal() *CreateHotspotExternal {
	createHotspotExternalType := new(CreateHotspotExternal)
	return createHotspotExternalType
}

func NewDefaultCreateHotspotExternal() *CreateHotspotExternal {
	createHotspotExternalType := new(CreateHotspotExternal)
	httpsRedirectField := true
	createHotspotExternalType.HttpsRedirect = &httpsRedirectField
	return createHotspotExternalType
}

type CreateHotspotInternal struct {
	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - required
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport" validate:"required,oneof=None Enabled"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewCreateHotspotInternal() *CreateHotspotInternal {
	createHotspotInternalType := new(CreateHotspotInternal)
	return createHotspotInternalType
}

func NewDefaultCreateHotspotInternal() *CreateHotspotInternal {
	createHotspotInternalType := new(CreateHotspotInternal)
	httpsRedirectField := true
	createHotspotInternalType.HttpsRedirect = &httpsRedirectField
	return createHotspotInternalType
}

type CreateHotspotSmartClientOnly struct {
	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	// MacAddressFormat
	// Constraints:
	//    - required
	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

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

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewCreateHotspotSmartClientOnly() *CreateHotspotSmartClientOnly {
	createHotspotSmartClientOnlyType := new(CreateHotspotSmartClientOnly)
	return createHotspotSmartClientOnlyType
}

func NewDefaultCreateHotspotSmartClientOnly() *CreateHotspotSmartClientOnly {
	createHotspotSmartClientOnlyType := new(CreateHotspotSmartClientOnly)
	httpsRedirectField := true
	createHotspotSmartClientOnlyType.HttpsRedirect = &httpsRedirectField
	return createHotspotSmartClientOnlyType
}

type CreateL2ACL struct {
	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction" validate:"required,oneof=ALLOW BLOCK"`

	RuleMacs []common.Mac `json:"ruleMacs,omitempty"`
}

func NewCreateL2ACL() *CreateL2ACL {
	createL2ACLType := new(CreateL2ACL)
	return createL2ACLType
}

func NewDefaultCreateL2ACL() *CreateL2ACL {
	createL2ACLType := new(CreateL2ACL)
	return createL2ACLType
}

type CreateWebAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	WebAuthenticationPortalCustomization *common.WebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`
}

func NewCreateWebAuthentication() *CreateWebAuthentication {
	createWebAuthenticationType := new(CreateWebAuthentication)
	return createWebAuthenticationType
}

func NewDefaultCreateWebAuthentication() *CreateWebAuthentication {
	createWebAuthenticationType := new(CreateWebAuthentication)
	return createWebAuthenticationType
}

type CreateWechat struct {
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

	Description *common.Description `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	// Constraints:
	//    - required
	DnatDestination *string `json:"dnatDestination" validate:"required"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - nullable
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"omitempty,gte=1,lte=14399"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// WhiteList
	// White list of the wechat profile
	// Constraints:
	//    - required
	WhiteList []string `json:"whiteList" validate:"required"`
}

func NewCreateWechat() *CreateWechat {
	createWechatType := new(CreateWechat)
	return createWechatType
}

func NewDefaultCreateWechat() *CreateWechat {
	createWechatType := new(CreateWechat)
	gracePeriodField := 60
	createWechatType.GracePeriod = &gracePeriodField
	return createWechatType
}

type DefaultConnectionCapability struct {
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

func NewDefaultConnectionCapability() *DefaultConnectionCapability {
	defaultConnectionCapabilityType := new(DefaultConnectionCapability)
	return defaultConnectionCapabilityType
}

func NewDefaultDefaultConnectionCapability() *DefaultConnectionCapability {
	defaultConnectionCapabilityType := new(DefaultConnectionCapability)
	return defaultConnectionCapabilityType
}

type DnatPortMapping struct {
	// DestPort
	// Destination port
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:65535
	DestPort *int `json:"destPort,omitempty" validate:"omitempty,gte=0,lte=65535"`

	// SourcePort
	// Source port
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:65535
	SourcePort *int `json:"sourcePort,omitempty" validate:"omitempty,gte=0,lte=65535"`
}

func NewDnatPortMapping() *DnatPortMapping {
	dnatPortMappingType := new(DnatPortMapping)
	return dnatPortMappingType
}

func NewDefaultDnatPortMapping() *DnatPortMapping {
	dnatPortMappingType := new(DnatPortMapping)
	return dnatPortMappingType
}

type GuestAccess struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the guest access profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// ZoneId
	// Identifier of the zone which the guest access profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewGuestAccess() *GuestAccess {
	guestAccessType := new(GuestAccess)
	return guestAccessType
}

func NewDefaultGuestAccess() *GuestAccess {
	guestAccessType := new(GuestAccess)
	return guestAccessType
}

type Hotspot struct {
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// Id
	// Identifier of the Hotspot
	Id *string `json:"id,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	// PortalType
	// Portal type of the Hotspot
	// Constraints:
	//    - nullable
	//    - oneof:[Internal,External]
	PortalType *string `json:"portalType,omitempty" validate:"omitempty,oneof=Internal External"`

	PortalUrl *common.NormalURL `json:"portalUrl,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - nullable
	//    - oneof:[None,Enabled,SmartClientOnly]
	SmartClientSupport *string `json:"smartClientSupport,omitempty" validate:"omitempty,oneof=None Enabled SmartClientOnly"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewHotspot() *Hotspot {
	hotspotType := new(Hotspot)
	return hotspotType
}

func NewDefaultHotspot() *Hotspot {
	hotspotType := new(Hotspot)
	httpsRedirectField := true
	hotspotType.HttpsRedirect = &httpsRedirectField
	return hotspotType
}

type Hotspot20VeuneProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkSpeedInKbps *LinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - nullable
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty" validate:"omitempty,oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

	// Id
	// Identifier of the Hotspot 2.0 venue profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *LinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 venue profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewHotspot20VeuneProfile() *Hotspot20VeuneProfile {
	hotspot20VeuneProfileType := new(Hotspot20VeuneProfile)
	return hotspot20VeuneProfileType
}

func NewDefaultHotspot20VeuneProfile() *Hotspot20VeuneProfile {
	hotspot20VeuneProfileType := new(Hotspot20VeuneProfile)
	return hotspot20VeuneProfileType
}

type Hotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - nullable
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty" validate:"omitempty,oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*ConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the Hotspot 2.0 WLAN profile
	Id *string `json:"id,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the v WLAN profile
	// Constraints:
	//    - nullable
	//    - oneof:[UNAVAILABLE,PUBLIC,PORT_RESTRICTED,SINGLE_NATED_PRIVATE,DOUBLE_NATED_PRIVATE,PORT_RESTRICTED_AND_SINGLE_NATED,PORT_RESTRICTED_AND_DOUBLE_NATED,UNKNOWN]
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty" validate:"omitempty,oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - nullable
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty" validate:"omitempty,oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

	Name *common.NormalName `json:"name,omitempty"`

	Operator *common.GenericRef `json:"operator,omitempty"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 WLAN profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewHotspot20WlanProfile() *Hotspot20WlanProfile {
	hotspot20WlanProfileType := new(Hotspot20WlanProfile)
	return hotspot20WlanProfileType
}

func NewDefaultHotspot20WlanProfile() *Hotspot20WlanProfile {
	hotspot20WlanProfileType := new(Hotspot20WlanProfile)
	return hotspot20WlanProfileType
}

type L2ACL struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// identifier of the L2 Access Control
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - nullable
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty" validate:"omitempty,oneof=ALLOW BLOCK"`

	RuleMacs []common.Mac `json:"ruleMacs,omitempty"`

	// ZoneId
	// identifier of the zone which the L2 Access Control belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewL2ACL() *L2ACL {
	l2ACLType := new(L2ACL)
	return l2ACLType
}

func NewDefaultL2ACL() *L2ACL {
	l2ACLType := new(L2ACL)
	return l2ACLType
}

// LinkSpeedInKbps
//
// Link Speed of the Hotspot 2.0 venue profile
// Constraints:
//    - nullable
//    - min:0.000000
//    - max:4294967295.000000
type LinkSpeedInKbps float64

func NewLinkSpeedInKbps() *LinkSpeedInKbps {
	linkSpeedInKbpsType := new(LinkSpeedInKbps)
	return linkSpeedInKbpsType
}

func NewDefaultLinkSpeedInKbps() *LinkSpeedInKbps {
	linkSpeedInKbpsType := new(LinkSpeedInKbps)
	return linkSpeedInKbpsType
}

// MacAddressFormatSetting
//
// mac address format of redirection,the format define: 0(aabbccddeeff), 1(AA-BB-CC-DD-EE-FF), 2(AA:BB:CC:DD:EE:FF), 3(AABBCCDDEEFF), 4(aa-bb-cc-dd-ee-ff), 5(aa:bb:cc:dd:ee:ff)
// Constraints:
//    - nullable
//    - default:2
//    - min:0
//    - max:5
type MacAddressFormatSetting int

func NewMacAddressFormatSetting() *MacAddressFormatSetting {
	macAddressFormatSettingType := new(MacAddressFormatSetting)
	return macAddressFormatSettingType
}

func NewDefaultMacAddressFormatSetting() *MacAddressFormatSetting {
	macAddressFormatSettingType := new(MacAddressFormatSetting)
	*macAddressFormatSettingType = 2
	return macAddressFormatSettingType
}

type ModifyGuestAccess struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`
}

func NewModifyGuestAccess() *ModifyGuestAccess {
	modifyGuestAccessType := new(ModifyGuestAccess)
	return modifyGuestAccessType
}

func NewDefaultModifyGuestAccess() *ModifyGuestAccess {
	modifyGuestAccessType := new(ModifyGuestAccess)
	return modifyGuestAccessType
}

type ModifyHotspot struct {
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	PortalUrl *common.NormalURL `json:"portalUrl,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	// Constraints:
	//    - nullable
	//    - oneof:[None,Enabled]
	SmartClientSupport *string `json:"smartClientSupport,omitempty" validate:"omitempty,oneof=None Enabled"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g. 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g. www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

func NewModifyHotspot() *ModifyHotspot {
	modifyHotspotType := new(ModifyHotspot)
	return modifyHotspotType
}

func NewDefaultModifyHotspot() *ModifyHotspot {
	modifyHotspotType := new(ModifyHotspot)
	httpsRedirectField := true
	modifyHotspotType.HttpsRedirect = &httpsRedirectField
	return modifyHotspotType
}

type ModifyHotspot20VenueProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkSpeedInKbps *LinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	// Constraints:
	//    - nullable
	//    - oneof:[Unspecified,Assembly,Business,Educational,FactoryAndIndustrial,Institutional,Mercantile,Residential,Storage,UtilityAndMiscellaneous,Vehicular,Outdoor]
	Group *string `json:"group,omitempty" validate:"omitempty,oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

	Name *common.NormalName `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *LinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames,omitempty"`
}

func NewModifyHotspot20VenueProfile() *ModifyHotspot20VenueProfile {
	modifyHotspot20VenueProfileType := new(ModifyHotspot20VenueProfile)
	return modifyHotspot20VenueProfileType
}

func NewDefaultModifyHotspot20VenueProfile() *ModifyHotspot20VenueProfile {
	modifyHotspot20VenueProfileType := new(ModifyHotspot20VenueProfile)
	return modifyHotspot20VenueProfileType
}

type ModifyHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	// Constraints:
	//    - nullable
	//    - oneof:[CHARGEABLE_PUBLIC,FREE_PUBLIC,PERSONAL_DEVICE,PRIVATE,PRIVATE_WITH_GUEST,TEST,WILDCARD]
	AccessNetworkType *string `json:"accessNetworkType,omitempty" validate:"omitempty,oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*ConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 Wlan profile
	// Constraints:
	//    - nullable
	//    - oneof:[UNAVAILABLE,PUBLIC,PORT_RESTRICTED,SINGLE_NATED_PRIVATE,DOUBLE_NATED_PRIVATE,PORT_RESTRICTED_AND_SINGLE_NATED,PORT_RESTRICTED_AND_DOUBLE_NATED,UNKNOWN]
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty" validate:"omitempty,oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 Wlan profile
	// Constraints:
	//    - nullable
	//    - oneof:[UNAVAILABLE,AVAILABLE,UNKNOWN]
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty" validate:"omitempty,oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

	Name *common.NormalName `json:"name,omitempty"`

	Operator *common.GenericRef `json:"operator,omitempty"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`
}

func NewModifyHotspot20WlanProfile() *ModifyHotspot20WlanProfile {
	modifyHotspot20WlanProfileType := new(ModifyHotspot20WlanProfile)
	return modifyHotspot20WlanProfileType
}

func NewDefaultModifyHotspot20WlanProfile() *ModifyHotspot20WlanProfile {
	modifyHotspot20WlanProfileType := new(ModifyHotspot20WlanProfile)
	return modifyHotspot20WlanProfileType
}

type ModifyL2ACL struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - nullable
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty" validate:"omitempty,oneof=ALLOW BLOCK"`

	RuleMacs []common.Mac `json:"ruleMacs,omitempty"`
}

func NewModifyL2ACL() *ModifyL2ACL {
	modifyL2ACLType := new(ModifyL2ACL)
	return modifyL2ACLType
}

func NewDefaultModifyL2ACL() *ModifyL2ACL {
	modifyL2ACLType := new(ModifyL2ACL)
	return modifyL2ACLType
}

type ModifyWebAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalLanguage *common.PortalLanguage `json:"portalLanguage,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *common.WebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`
}

func NewModifyWebAuthentication() *ModifyWebAuthentication {
	modifyWebAuthenticationType := new(ModifyWebAuthentication)
	return modifyWebAuthenticationType
}

func NewDefaultModifyWebAuthentication() *ModifyWebAuthentication {
	modifyWebAuthenticationType := new(ModifyWebAuthentication)
	return modifyWebAuthenticationType
}

type ModifyWechat struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - nullable
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"omitempty,gte=1,lte=14399"`

	Name *common.NormalName `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

func NewModifyWechat() *ModifyWechat {
	modifyWechatType := new(ModifyWechat)
	return modifyWechatType
}

func NewDefaultModifyWechat() *ModifyWechat {
	modifyWechatType := new(ModifyWechat)
	gracePeriodField := 60
	modifyWechatType.GracePeriod = &gracePeriodField
	return modifyWechatType
}

type PortalLocation struct {
	// Id
	// Portal location id
	Id *string `json:"id,omitempty"`

	// Name
	// Portal location name
	Name *string `json:"name,omitempty"`
}

func NewPortalLocation() *PortalLocation {
	portalLocationType := new(PortalLocation)
	return portalLocationType
}

func NewDefaultPortalLocation() *PortalLocation {
	portalLocationType := new(PortalLocation)
	return portalLocationType
}

type PortalRedirect struct {
	Url *common.NormalURL `json:"url,omitempty"`
}

func NewPortalRedirect() *PortalRedirect {
	portalRedirectType := new(PortalRedirect)
	return portalRedirectType
}

func NewDefaultPortalRedirect() *PortalRedirect {
	portalRedirectType := new(PortalRedirect)
	return portalRedirectType
}

type PortalServiceList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortalServiceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewPortalServiceList() *PortalServiceList {
	portalServiceListType := new(PortalServiceList)
	return portalServiceListType
}

func NewDefaultPortalServiceList() *PortalServiceList {
	portalServiceListType := new(PortalServiceList)
	return portalServiceListType
}

type PortalServiceListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

func NewPortalServiceListType() *PortalServiceListType {
	portalServiceListTypeType := new(PortalServiceListType)
	return portalServiceListTypeType
}

func NewDefaultPortalServiceListType() *PortalServiceListType {
	portalServiceListTypeType := new(PortalServiceListType)
	return portalServiceListTypeType
}

type UserSession struct {
	// GracePeriodInMin
	// Grace period in minutes
	// Constraints:
	//    - nullable
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriodInMin *int `json:"gracePeriodInMin,omitempty" validate:"omitempty,gte=1,lte=14399"`

	// TimeoutInMin
	// Time out value in minutes
	// Constraints:
	//    - nullable
	//    - default:1440
	//    - min:2
	//    - max:14400
	TimeoutInMin *int `json:"timeoutInMin,omitempty" validate:"omitempty,gte=2,lte=14400"`
}

func NewUserSession() *UserSession {
	userSessionType := new(UserSession)
	return userSessionType
}

func NewDefaultUserSession() *UserSession {
	userSessionType := new(UserSession)
	gracePeriodInMinField := 60
	userSessionType.GracePeriodInMin = &gracePeriodInMinField
	timeoutInMinField := 1440
	userSessionType.TimeoutInMin = &timeoutInMinField
	return userSessionType
}

type VenueName struct {
	// Language
	// Constraints:
	//    - required
	Language *common.LanguageName `json:"language" validate:"required"`

	// Name
	// Venue name
	// Constraints:
	//    - required
	Name *string `json:"name" validate:"required"`
}

func NewVenueName() *VenueName {
	venueNameType := new(VenueName)
	return venueNameType
}

func NewDefaultVenueName() *VenueName {
	venueNameType := new(VenueName)
	return venueNameType
}

type WebAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the web authentication profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalLanguage *common.PortalLanguage `json:"portalLanguage,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *common.WebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`

	// ZoneId
	// Identifier of the zone which the web authentication profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWebAuthentication() *WebAuthentication {
	webAuthenticationType := new(WebAuthentication)
	return webAuthenticationType
}

func NewDefaultWebAuthentication() *WebAuthentication {
	webAuthenticationType := new(WebAuthentication)
	return webAuthenticationType
}

type WechatConfiguration struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	// Constraints:
	//    - nullable
	//    - default:60
	//    - min:1
	//    - max:14399
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"omitempty,gte=1,lte=14399"`

	Name *common.NormalName `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

func NewWechatConfiguration() *WechatConfiguration {
	wechatConfigurationType := new(WechatConfiguration)
	return wechatConfigurationType
}

func NewDefaultWechatConfiguration() *WechatConfiguration {
	wechatConfigurationType := new(WechatConfiguration)
	gracePeriodField := 60
	wechatConfigurationType.GracePeriod = &gracePeriodField
	return wechatConfigurationType
}
