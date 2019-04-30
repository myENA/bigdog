package portalservice

// API Version: v8_0

type ConnectionCapability struct {
	PortNumber     *float64 `json:"portNumber,omitempty"`
	ProtocolName   *string  `json:"protocolName,omitempty"`
	ProtocolNumber *float64 `json:"protocolNumber,omitempty"`
	Status         *string  `json:"status,omitempty"`
}

type CreateGuestAccess struct {
	Description         *string                     `json:"description,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`
	Redirect            *PortalRedirect             `json:"redirect,omitempty"`
	SmsGateway          *common.GenericRef          `json:"smsGateway,omitempty"`
	UserSession         *UserSession                `json:"userSession,omitempty"`
}

type CreateHotspot20VenueProfile struct {
	Description         *string      `json:"description,omitempty"`
	DownlinkSpeedInKbps *float64     `json:"downlinkSpeedInKbps,omitempty"`
	Group               *string      `json:"group,omitempty"`
	Name                *string      `json:"name,omitempty"`
	Type                *string      `json:"type,omitempty"`
	UplinkSpeedInKbps   *float64     `json:"uplinkSpeedInKbps,omitempty"`
	VenueNames          []*VenueName `json:"venueNames,omitempty"`
}

type CreateHotspot20WLANProfile struct {
	AccessNetworkType       *string              `json:"accessNetworkType,omitempty"`
	DefaultIdentityProvider *common.GenericRef   `json:"defaultIdentityProvider,omitempty"`
	Description             *string              `json:"description,omitempty"`
	IdentityProviders       []*common.GenericRef `json:"identityProviders,omitempty"`
	InternetOption          *bool                `json:"internetOption,omitempty"`
	Ipv4AddressType         *string              `json:"ipv4AddressType,omitempty"`
	Ipv6AddressType         *string              `json:"ipv6AddressType,omitempty"`
	Name                    *string              `json:"name,omitempty"`
	Operator                *common.GenericRef   `json:"operator,omitempty"`
	SignupSsid              *common.GenericRef   `json:"signupSsid,omitempty"`
}

type CreateHotspotExternal struct {
	BackupPortalURL       *string         `json:"backupPortalUrl,omitempty"`
	Description           *string         `json:"description,omitempty"`
	HTTPSRedirect         *bool           `json:"httpsRedirect,omitempty"`
	InternalNode          *string         `json:"internalNode,omitempty"`
	Location              *PortalLocation `json:"location,omitempty"`
	MacAddressFormat      *int            `json:"macAddressFormat,omitempty"`
	Name                  *string         `json:"name,omitempty"`
	PortalURL             *string         `json:"portalUrl,omitempty"`
	Redirect              *PortalRedirect `json:"redirect,omitempty"`
	SignatureSigningKey   *string         `json:"signatureSigningKey,omitempty"`
	SmartClientSupport    *string         `json:"smartClientSupport,omitempty"`
	TrafficClassProfileID *string         `json:"trafficClassProfileId,omitempty"`
	UserSession           *UserSession    `json:"userSession,omitempty"`
	WalledGardens         []string        `json:"walledGardens,omitempty"`
}

type CreateHotspotInternal struct {
	Description           *string         `json:"description,omitempty"`
	HTTPSRedirect         *bool           `json:"httpsRedirect,omitempty"`
	InternalNode          *string         `json:"internalNode,omitempty"`
	Location              *PortalLocation `json:"location,omitempty"`
	MacAddressFormat      *int            `json:"macAddressFormat,omitempty"`
	Name                  *string         `json:"name,omitempty"`
	Redirect              *PortalRedirect `json:"redirect,omitempty"`
	SignatureSigningKey   *string         `json:"signatureSigningKey,omitempty"`
	SmartClientSupport    *string         `json:"smartClientSupport,omitempty"`
	TrafficClassProfileID *string         `json:"trafficClassProfileId,omitempty"`
	UserSession           *UserSession    `json:"userSession,omitempty"`
	WalledGardens         []string        `json:"walledGardens,omitempty"`
}

type CreateHotspotSmartClientOnly struct {
	Description           *string         `json:"description,omitempty"`
	HTTPSRedirect         *bool           `json:"httpsRedirect,omitempty"`
	InternalNode          *string         `json:"internalNode,omitempty"`
	Location              *PortalLocation `json:"location,omitempty"`
	MacAddressFormat      *int            `json:"macAddressFormat,omitempty"`
	Name                  *string         `json:"name,omitempty"`
	Redirect              *PortalRedirect `json:"redirect,omitempty"`
	SignatureSigningKey   *string         `json:"signatureSigningKey,omitempty"`
	SmartClientInfo       *string         `json:"smartClientInfo,omitempty"`
	TrafficClassProfileID *string         `json:"trafficClassProfileId,omitempty"`
	UserSession           *UserSession    `json:"userSession,omitempty"`
	WalledGardens         []string        `json:"walledGardens,omitempty"`
}

type CreateL2ACL struct {
	Description *string  `json:"description,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Restriction *string  `json:"restriction,omitempty"`
	RuleMacs    []string `json:"ruleMacs,omitempty"`
}

type CreateWebAuthentication struct {
	Description *string         `json:"description,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Redirect    *PortalRedirect `json:"redirect,omitempty"`
}

type CreateWechat struct {
	AuthURL         *string            `json:"authUrl,omitempty"`
	BlackList       *string            `json:"blackList,omitempty"`
	Description     *string            `json:"description,omitempty"`
	DnatDestination *string            `json:"dnatDestination,omitempty"`
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`
	GracePeriod     *int               `json:"gracePeriod,omitempty"`
	Name            *string            `json:"name,omitempty"`
	WhiteList       []string           `json:"whiteList,omitempty"`
}

type DefaultConnectionCapability struct {
	PortNumber     *float64 `json:"portNumber,omitempty"`
	ProtocolName   *string  `json:"protocolName,omitempty"`
	ProtocolNumber *float64 `json:"protocolNumber,omitempty"`
	Status         *string  `json:"status,omitempty"`
}

type DnatPortMapping struct {
	DestPort   *int `json:"destPort,omitempty"`
	SourcePort *int `json:"sourcePort,omitempty"`
}

type GuestAccess struct {
	Description         *string                     `json:"description,omitempty"`
	ID                  *string                     `json:"id,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`
	Redirect            *PortalRedirect             `json:"redirect,omitempty"`
	SmsGateway          *common.GenericRef          `json:"smsGateway,omitempty"`
	UserSession         *UserSession                `json:"userSession,omitempty"`
	ZoneID              *string                     `json:"zoneId,omitempty"`
}

type Hotspot struct {
	BackupPortalURL       *string                     `json:"backupPortalUrl,omitempty"`
	Description           *string                     `json:"description,omitempty"`
	HTTPSRedirect         *bool                       `json:"httpsRedirect,omitempty"`
	ID                    *string                     `json:"id,omitempty"`
	InternalNode          *string                     `json:"internalNode,omitempty"`
	Location              *PortalLocation             `json:"location,omitempty"`
	MacAddressFormat      *int                        `json:"macAddressFormat,omitempty"`
	Name                  *string                     `json:"name,omitempty"`
	PortalCustomization   *common.PortalCustomization `json:"portalCustomization,omitempty"`
	PortalType            *string                     `json:"portalType,omitempty"`
	PortalURL             *string                     `json:"portalUrl,omitempty"`
	Redirect              *PortalRedirect             `json:"redirect,omitempty"`
	SignatureSigningKey   *string                     `json:"signatureSigningKey,omitempty"`
	SmartClientInfo       *string                     `json:"smartClientInfo,omitempty"`
	SmartClientSupport    *string                     `json:"smartClientSupport,omitempty"`
	TrafficClassProfileID *string                     `json:"trafficClassProfileId,omitempty"`
	UserSession           *UserSession                `json:"userSession,omitempty"`
	WalledGardens         []string                    `json:"walledGardens,omitempty"`
	ZoneID                *string                     `json:"zoneId,omitempty"`
}

type Hotspot20VeuneProfile struct {
	Description         *string      `json:"description,omitempty"`
	DownlinkSpeedInKbps *float64     `json:"downlinkSpeedInKbps,omitempty"`
	Group               *string      `json:"group,omitempty"`
	ID                  *string      `json:"id,omitempty"`
	Name                *string      `json:"name,omitempty"`
	Type                *string      `json:"type,omitempty"`
	UplinkSpeedInKbps   *float64     `json:"uplinkSpeedInKbps,omitempty"`
	VenueNames          []*VenueName `json:"venueNames,omitempty"`
	ZoneID              *string      `json:"zoneId,omitempty"`
}

type Hotspot20WLANProfile struct {
	AccessNetworkType            *string                        `json:"accessNetworkType,omitempty"`
	ConnectionCapabilities       []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`
	CustomConnectionCapabilities []*ConnectionCapability        `json:"customConnectionCapabilities,omitempty"`
	DefaultIdentityProvider      *common.GenericRef             `json:"defaultIdentityProvider,omitempty"`
	Description                  *string                        `json:"description,omitempty"`
	ID                           *string                        `json:"id,omitempty"`
	IdentityProviders            []*common.GenericRef           `json:"identityProviders,omitempty"`
	InternetOption               *bool                          `json:"internetOption,omitempty"`
	Ipv4AddressType              *string                        `json:"ipv4AddressType,omitempty"`
	Ipv6AddressType              *string                        `json:"ipv6AddressType,omitempty"`
	Name                         *string                        `json:"name,omitempty"`
	Operator                     *common.GenericRef             `json:"operator,omitempty"`
	SignupSsid                   *common.GenericRef             `json:"signupSsid,omitempty"`
	ZoneID                       *string                        `json:"zoneId,omitempty"`
}

type L2ACL struct {
	Description *string  `json:"description,omitempty"`
	ID          *string  `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Restriction *string  `json:"restriction,omitempty"`
	RuleMacs    []string `json:"ruleMacs,omitempty"`
	ZoneID      *string  `json:"zoneId,omitempty"`
}

type ModifyGuestAccess struct {
	Description         *string                     `json:"description,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`
	Redirect            *PortalRedirect             `json:"redirect,omitempty"`
	SmsGateway          *common.GenericRef          `json:"smsGateway,omitempty"`
	UserSession         *UserSession                `json:"userSession,omitempty"`
}

type ModifyHotspot struct {
	BackupPortalURL       *string                     `json:"backupPortalUrl,omitempty"`
	Description           *string                     `json:"description,omitempty"`
	HTTPSRedirect         *bool                       `json:"httpsRedirect,omitempty"`
	InternalNode          *string                     `json:"internalNode,omitempty"`
	Location              *PortalLocation             `json:"location,omitempty"`
	MacAddressFormat      *int                        `json:"macAddressFormat,omitempty"`
	Name                  *string                     `json:"name,omitempty"`
	PortalCustomization   *common.PortalCustomization `json:"portalCustomization,omitempty"`
	PortalURL             *string                     `json:"portalUrl,omitempty"`
	Redirect              *PortalRedirect             `json:"redirect,omitempty"`
	SignatureSigningKey   *string                     `json:"signatureSigningKey,omitempty"`
	SmartClientInfo       *string                     `json:"smartClientInfo,omitempty"`
	SmartClientSupport    *string                     `json:"smartClientSupport,omitempty"`
	TrafficClassProfileID *string                     `json:"trafficClassProfileId,omitempty"`
	UserSession           *UserSession                `json:"userSession,omitempty"`
	WalledGardens         []string                    `json:"walledGardens,omitempty"`
}

type ModifyHotspot20VenueProfile struct {
	Description         *string      `json:"description,omitempty"`
	DownlinkSpeedInKbps *float64     `json:"downlinkSpeedInKbps,omitempty"`
	Group               *string      `json:"group,omitempty"`
	Name                *string      `json:"name,omitempty"`
	Type                *string      `json:"type,omitempty"`
	UplinkSpeedInKbps   *float64     `json:"uplinkSpeedInKbps,omitempty"`
	VenueNames          []*VenueName `json:"venueNames,omitempty"`
}

type ModifyHotspot20WLANProfile struct {
	AccessNetworkType            *string                        `json:"accessNetworkType,omitempty"`
	ConnectionCapabilities       []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`
	CustomConnectionCapabilities []*ConnectionCapability        `json:"customConnectionCapabilities,omitempty"`
	DefaultIdentityProvider      *common.GenericRef             `json:"defaultIdentityProvider,omitempty"`
	Description                  *string                        `json:"description,omitempty"`
	IdentityProviders            []*common.GenericRef           `json:"identityProviders,omitempty"`
	InternetOption               *bool                          `json:"internetOption,omitempty"`
	Ipv4AddressType              *string                        `json:"ipv4AddressType,omitempty"`
	Ipv6AddressType              *string                        `json:"ipv6AddressType,omitempty"`
	Name                         *string                        `json:"name,omitempty"`
	Operator                     *common.GenericRef             `json:"operator,omitempty"`
	SignupSsid                   *common.GenericRef             `json:"signupSsid,omitempty"`
}

type ModifyL2ACL struct {
	Description *string  `json:"description,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Restriction *string  `json:"restriction,omitempty"`
	RuleMacs    []string `json:"ruleMacs,omitempty"`
}

type ModifyWebAuthentication struct {
	Description    *string         `json:"description,omitempty"`
	Name           *string         `json:"name,omitempty"`
	PortalLanguage *string         `json:"portalLanguage,omitempty"`
	Redirect       *PortalRedirect `json:"redirect,omitempty"`
	UserSession    *UserSession    `json:"userSession,omitempty"`
}

type ModifyWechat struct {
	AuthURL         *string            `json:"authUrl,omitempty"`
	BlackList       *string            `json:"blackList,omitempty"`
	Description     *string            `json:"description,omitempty"`
	DnatDestination *string            `json:"dnatDestination,omitempty"`
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`
	GracePeriod     *int               `json:"gracePeriod,omitempty"`
	Name            *string            `json:"name,omitempty"`
	WhiteList       []string           `json:"whiteList,omitempty"`
}

type PortalLocation struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PortalRedirect struct {
	URL *string `json:"url,omitempty"`
}

type PortalServiceList struct {
	FirstIndex *int                     `json:"firstIndex,omitempty"`
	HasMore    *bool                    `json:"hasMore,omitempty"`
	List       []*PortalServiceListType `json:"list,omitempty"`
	TotalCount *int                     `json:"totalCount,omitempty"`
}

type PortalServiceListType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UserSession struct {
	GracePeriodInMin *int `json:"gracePeriodInMin,omitempty"`
	TimeoutInMin     *int `json:"timeoutInMin,omitempty"`
}

type VenueName struct {
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type WebAuthentication struct {
	Description    *string         `json:"description,omitempty"`
	ID             *string         `json:"id,omitempty"`
	Name           *string         `json:"name,omitempty"`
	PortalLanguage *string         `json:"portalLanguage,omitempty"`
	Redirect       *PortalRedirect `json:"redirect,omitempty"`
	UserSession    *UserSession    `json:"userSession,omitempty"`
	ZoneID         *string         `json:"zoneId,omitempty"`
}

type WechatConfiguration struct {
	AuthURL         *string            `json:"authUrl,omitempty"`
	BlackList       *string            `json:"blackList,omitempty"`
	Description     *string            `json:"description,omitempty"`
	DnatDestination *string            `json:"dnatDestination,omitempty"`
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`
	GracePeriod     *int               `json:"gracePeriod,omitempty"`
	Name            *string            `json:"name,omitempty"`
	WhiteList       []string           `json:"whiteList,omitempty"`
}
