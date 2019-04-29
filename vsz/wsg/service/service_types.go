package service

// API Version: v8_0

type ActiveDirectoryService struct {
	AdminDomainName             *string     `json:"adminDomainName,omitempty"`
	CreateDateTime              *int        `json:"createDateTime,omitempty"`
	CreatorID                   *string     `json:"creatorId,omitempty"`
	CreatorUsername             *string     `json:"creatorUsername,omitempty"`
	Description                 *string     `json:"description,omitempty"`
	DomainID                    *string     `json:"domainId,omitempty"`
	FriendlyName                *string     `json:"friendlyName,omitempty"`
	GlobalCatalogEnabled        *bool       `json:"globalCatalogEnabled,omitempty"`
	ID                          *string     `json:"id,omitempty"`
	IP                          *string     `json:"ip,omitempty"`
	Mappings                    []*Mappings `json:"mappings,omitempty"`
	ModifiedDateTime            *int        `json:"modifiedDateTime,omitempty"`
	ModifierID                  *string     `json:"modifierId,omitempty"`
	ModifierUsername            *string     `json:"modifierUsername,omitempty"`
	MvnoID                      *string     `json:"mvnoId,omitempty"`
	Name                        *string     `json:"name,omitempty"`
	Password                    *string     `json:"password,omitempty"`
	Port                        *int        `json:"port,omitempty"`
	Protocol                    *string     `json:"protocol,omitempty"`
	StandbyAdminDomainName      *string     `json:"standbyAdminDomainName,omitempty"`
	StandbyGlobalCatalogEnabled *bool       `json:"standbyGlobalCatalogEnabled,omitempty"`
	StandbyIP                   *string     `json:"standbyIp,omitempty"`
	StandbyPassword             *string     `json:"standbyPassword,omitempty"`
	StandbyPort                 *int        `json:"standbyPort,omitempty"`
	StandbyServerEnabled        *bool       `json:"standbyServerEnabled,omitempty"`
	StandbyTLSEnabled           *bool       `json:"standbyTlsEnabled,omitempty"`
	StandbyWindowsDomainName    *string     `json:"standbyWindowsDomainName,omitempty"`
	TLSEnabled                  *bool       `json:"tlsEnabled,omitempty"`
	Type                        *string     `json:"type,omitempty"`
	WindowsDomainName           *string     `json:"windowsDomainName,omitempty"`
}

type ActiveDirectoryServiceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type CommonAccountingService struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	Description      *string `json:"description,omitempty"`
	DomainID         *string `json:"domainId,omitempty"`
	ID               *string `json:"id,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	MvnoID           *string `json:"mvnoId,omitempty"`
	Name             *string `json:"name,omitempty"`
	Protocol         *string `json:"protocol,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type CommonAccountingServiceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type CommonAuthenticationService struct {
	CreateDateTime   *int        `json:"createDateTime,omitempty"`
	CreatorID        *string     `json:"creatorId,omitempty"`
	CreatorUsername  *string     `json:"creatorUsername,omitempty"`
	Description      *string     `json:"description,omitempty"`
	DomainID         *string     `json:"domainId,omitempty"`
	FriendlyName     *string     `json:"friendlyName,omitempty"`
	ID               *string     `json:"id,omitempty"`
	Mappings         []*Mappings `json:"mappings,omitempty"`
	ModifiedDateTime *int        `json:"modifiedDateTime,omitempty"`
	ModifierID       *string     `json:"modifierId,omitempty"`
	ModifierUsername *string     `json:"modifierUsername,omitempty"`
	MvnoID           *string     `json:"mvnoId,omitempty"`
	Name             *string     `json:"name,omitempty"`
	Protocol         *string     `json:"protocol,omitempty"`
	Type             *string     `json:"type,omitempty"`
}

type CommonAuthenticationServiceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type CreateActiveDirectoryAuthentication struct {
	AdminDomainName             *string     `json:"adminDomainName,omitempty"`
	Description                 *string     `json:"description,omitempty"`
	DomainID                    *string     `json:"domainId,omitempty"`
	FriendlyName                *string     `json:"friendlyName,omitempty"`
	GlobalCatalogEnabled        *bool       `json:"globalCatalogEnabled,omitempty"`
	ID                          *string     `json:"id,omitempty"`
	IP                          *string     `json:"ip,omitempty"`
	Mappings                    []*Mappings `json:"mappings,omitempty"`
	Name                        *string     `json:"name,omitempty"`
	Password                    *string     `json:"password,omitempty"`
	Port                        *int        `json:"port,omitempty"`
	StandbyAdminDomainName      *string     `json:"standbyAdminDomainName,omitempty"`
	StandbyGlobalCatalogEnabled *bool       `json:"standbyGlobalCatalogEnabled,omitempty"`
	StandbyIP                   *string     `json:"standbyIp,omitempty"`
	StandbyPassword             *string     `json:"standbyPassword,omitempty"`
	StandbyPort                 *int        `json:"standbyPort,omitempty"`
	StandbyServerEnabled        *bool       `json:"standbyServerEnabled,omitempty"`
	StandbyTLSEnabled           *bool       `json:"standbyTlsEnabled,omitempty"`
	StandbyWindowsDomainName    *string     `json:"standbyWindowsDomainName,omitempty"`
	TLSEnabled                  *bool       `json:"tlsEnabled,omitempty"`
	Type                        *string     `json:"type,omitempty"`
	WindowsDomainName           *string     `json:"windowsDomainName,omitempty"`
}

type CreateHlrAuthentication struct {
	AddressIndicator             *string                 `json:"addressIndicator,omitempty"`
	AuthMapVer                   *string                 `json:"authMapVer,omitempty"`
	AuthorizationCachingEnabled  *bool                   `json:"authorizationCachingEnabled,omitempty"`
	AvCachingEnabled             *bool                   `json:"avCachingEnabled,omitempty"`
	CacheOptionType              *string                 `json:"cacheOptionType,omitempty"`
	CleanUpTimeHour              *int                    `json:"cleanUpTimeHour,omitempty"`
	CleanUpTimeMinute            *int                    `json:"cleanUpTimeMinute,omitempty"`
	CreateDateTime               *int                    `json:"createDateTime,omitempty"`
	CreatorID                    *string                 `json:"creatorId,omitempty"`
	CreatorUsername              *string                 `json:"creatorUsername,omitempty"`
	DefaultPointCodeFormat       *string                 `json:"defaultPointCodeFormat,omitempty"`
	Description                  *string                 `json:"description,omitempty"`
	DestGtIndicator              *string                 `json:"destGtIndicator,omitempty"`
	DestNatureOfAddressIndicator *string                 `json:"destNatureOfAddressIndicator,omitempty"`
	DestNumberingPlan            *string                 `json:"destNumberingPlan,omitempty"`
	DestTransType                *int                    `json:"destTransType,omitempty"`
	DomainID                     *string                 `json:"domainId,omitempty"`
	E164Address                  *string                 `json:"e164Address,omitempty"`
	EapSimMapVer                 *string                 `json:"eapSimMapVer,omitempty"`
	FriendlyName                 *string                 `json:"friendlyName,omitempty"`
	GtPointCode                  *int                    `json:"gtPointCode,omitempty"`
	HasPointCode                 *bool                   `json:"hasPointCode,omitempty"`
	HasSrcPointCode              *bool                   `json:"hasSrcPointCode,omitempty"`
	HasSSN                       *bool                   `json:"hasSSN,omitempty"`
	HistoryTime                  *int                    `json:"historyTime,omitempty"`
	ID                           *string                 `json:"id,omitempty"`
	LocalNetworkIndicator        *string                 `json:"localNetworkIndicator,omitempty"`
	LocalPointCode               *int                    `json:"localPointCode,omitempty"`
	LocationDeliveryEnabled      *bool                   `json:"locationDeliveryEnabled,omitempty"`
	MaxReuseTimes                *int                    `json:"maxReuseTimes,omitempty"`
	MncNdcList                   []*MncNdcList           `json:"mncNdcList,omitempty"`
	ModifiedDateTime             *int                    `json:"modifiedDateTime,omitempty"`
	ModifierID                   *string                 `json:"modifierId,omitempty"`
	ModifierUsername             *string                 `json:"modifierUsername,omitempty"`
	MvnoID                       *string                 `json:"mvnoId,omitempty"`
	Name                         *string                 `json:"name,omitempty"`
	PointCode                    *int                    `json:"pointCode,omitempty"`
	Protocol                     *string                 `json:"protocol,omitempty"`
	ReuseEnable                  *bool                   `json:"reuseEnable,omitempty"`
	RoutingContext               *int                    `json:"routingContext,omitempty"`
	SccpGttList                  []*SccpGttList          `json:"sccpGttList,omitempty"`
	SctpAssociationsList         []*SctpAssociationsList `json:"sctpAssociationsList,omitempty"`
	SgsnIsdnAddress              *string                 `json:"sgsnIsdnAddress,omitempty"`
	SrcGtIndicator               *string                 `json:"srcGtIndicator,omitempty"`
	SrcNatureOfAddressIndicator  *string                 `json:"srcNatureOfAddressIndicator,omitempty"`
	SrcNumberingPlan             *string                 `json:"srcNumberingPlan,omitempty"`
	SrcTransType                 *int                    `json:"srcTransType,omitempty"`
	Type                         *string                 `json:"type,omitempty"`
}

type CreateLDAPAuthentication struct {
	AdminDomainName        *string     `json:"adminDomainName,omitempty"`
	BaseDomainName         *string     `json:"baseDomainName,omitempty"`
	Description            *string     `json:"description,omitempty"`
	DomainID               *string     `json:"domainId,omitempty"`
	FriendlyName           *string     `json:"friendlyName,omitempty"`
	ID                     *string     `json:"id,omitempty"`
	IP                     *string     `json:"ip,omitempty"`
	KeyAttribute           *string     `json:"keyAttribute,omitempty"`
	Mappings               []*Mappings `json:"mappings,omitempty"`
	Name                   *string     `json:"name,omitempty"`
	Password               *string     `json:"password,omitempty"`
	Port                   *int        `json:"port,omitempty"`
	SearchFilter           *string     `json:"searchFilter,omitempty"`
	StandbyAdminDomainName *string     `json:"standbyAdminDomainName,omitempty"`
	StandbyBaseDomainName  *string     `json:"standbyBaseDomainName,omitempty"`
	StandbyIP              *string     `json:"standbyIp,omitempty"`
	StandbyKeyAttribute    *string     `json:"standbyKeyAttribute,omitempty"`
	StandbyPassword        *string     `json:"standbyPassword,omitempty"`
	StandbyPort            *int        `json:"standbyPort,omitempty"`
	StandbySearchFilter    *string     `json:"standbySearchFilter,omitempty"`
	StandbyServerEnabled   *bool       `json:"standbyServerEnabled,omitempty"`
	StandbyTLSEnabled      *bool       `json:"standbyTlsEnabled,omitempty"`
	TLSEnabled             *bool       `json:"tlsEnabled,omitempty"`
	Type                   *string     `json:"type,omitempty"`
}

type CreateRadiusAccounting struct {
	Description          *string                   `json:"description,omitempty"`
	DomainID             *string                   `json:"domainId,omitempty"`
	HealthCheckPolicy    *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
	Primary              *common.RadiusServer      `json:"primary,omitempty"`
	Protocol             *string                   `json:"protocol,omitempty"`
	RateLimiting         *common.RateLimiting      `json:"rateLimiting,omitempty"`
	Secondary            *Secondary                `json:"secondary,omitempty"`
	StandbyPrimary       *common.RadiusServer      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled *bool                     `json:"standbyServerEnabled,omitempty"`
	Type                 *string                   `json:"type,omitempty"`
}

type CreateRadiusAuthentication struct {
	Description             *string                   `json:"description,omitempty"`
	DomainID                *string                   `json:"domainId,omitempty"`
	FriendlyName            *string                   `json:"friendlyName,omitempty"`
	HealthCheckPolicy       *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
	ID                      *string                   `json:"id,omitempty"`
	LocationDeliveryEnabled *bool                     `json:"locationDeliveryEnabled,omitempty"`
	Mappings                []*Mappings               `json:"mappings,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	Primary                 *common.RadiusServer      `json:"primary,omitempty"`
	RateLimiting            *common.RateLimiting      `json:"rateLimiting,omitempty"`
	Secondary               *Secondary                `json:"secondary,omitempty"`
	StandbyPrimary          *common.RadiusServer      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled    *bool                     `json:"standbyServerEnabled,omitempty"`
	Type                    *string                   `json:"type,omitempty"`
}

type DeleteBulkAccountingService struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DeleteBulkAuthenticationService struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DNSServer struct {
	IP *string `json:"ip,omitempty"`
}

type Ggsn struct {
	DomainName    *string `json:"domainName,omitempty"`
	GgsnIPAddress *string `json:"ggsnIPAddress,omitempty"`
}

type GgsnConfig struct {
	DNSServerList DNSServerList `json:"dnsServerList,omitempty"`
	GgsnList      GgsnList      `json:"ggsnList,omitempty"`
	GtpSettings   *GtpSettings  `json:"gtpSettings,omitempty"`
}

type GroupAttrIdentityUserRoleMapping struct {
	GroupAttr *string   `json:"groupAttr,omitempty"`
	ID        *string   `json:"id,omitempty"`
	UserRole  *UserRole `json:"userRole,omitempty"`
}

type GroupAttrIdentityUserRoleMappingUserRoleType struct {
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	UserTrafficProfile *UserTrafficProfile `json:"userTrafficProfile,omitempty"`
}

type GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GtpSettings struct {
	DNSNumberOfRetries *int `json:"dnsNumberOfRetries,omitempty"`
	EchoRequestTimer   *int `json:"echoRequestTimer,omitempty"`
	NumberOfRetries    *int `json:"numberOfRetries,omitempty"`
	ResponseTimeout    *int `json:"responseTimeout,omitempty"`
	T3ResponseTimer    *int `json:"t3ResponseTimer,omitempty"`
}

type HlrService struct {
	AddressIndicator             *string                 `json:"addressIndicator,omitempty"`
	AuthMapVer                   *string                 `json:"authMapVer,omitempty"`
	AuthorizationCachingEnabled  *bool                   `json:"authorizationCachingEnabled,omitempty"`
	AvCachingEnabled             *bool                   `json:"avCachingEnabled,omitempty"`
	CacheOptionType              *string                 `json:"cacheOptionType,omitempty"`
	CleanUpTimeHour              *int                    `json:"cleanUpTimeHour,omitempty"`
	CleanUpTimeMinute            *int                    `json:"cleanUpTimeMinute,omitempty"`
	CreateDateTime               *int                    `json:"createDateTime,omitempty"`
	CreatorID                    *string                 `json:"creatorId,omitempty"`
	CreatorUsername              *string                 `json:"creatorUsername,omitempty"`
	DefaultPointCodeFormat       *string                 `json:"defaultPointCodeFormat,omitempty"`
	Description                  *string                 `json:"description,omitempty"`
	DestGtIndicator              *string                 `json:"destGtIndicator,omitempty"`
	DestNatureOfAddressIndicator *string                 `json:"destNatureOfAddressIndicator,omitempty"`
	DestNumberingPlan            *string                 `json:"destNumberingPlan,omitempty"`
	DestTransType                *int                    `json:"destTransType,omitempty"`
	DomainID                     *string                 `json:"domainId,omitempty"`
	E164Address                  *string                 `json:"e164Address,omitempty"`
	EapSimMapVer                 *string                 `json:"eapSimMapVer,omitempty"`
	FriendlyName                 *string                 `json:"friendlyName,omitempty"`
	GtPointCode                  *int                    `json:"gtPointCode,omitempty"`
	HasPointCode                 *bool                   `json:"hasPointCode,omitempty"`
	HasSrcPointCode              *bool                   `json:"hasSrcPointCode,omitempty"`
	HasSSN                       *bool                   `json:"hasSSN,omitempty"`
	HistoryTime                  *int                    `json:"historyTime,omitempty"`
	ID                           *string                 `json:"id,omitempty"`
	LocalNetworkIndicator        *string                 `json:"localNetworkIndicator,omitempty"`
	LocalPointCode               *int                    `json:"localPointCode,omitempty"`
	LocationDeliveryEnabled      *bool                   `json:"locationDeliveryEnabled,omitempty"`
	MaxReuseTimes                *int                    `json:"maxReuseTimes,omitempty"`
	MncNdcList                   []*MncNdcList           `json:"mncNdcList,omitempty"`
	ModifiedDateTime             *int                    `json:"modifiedDateTime,omitempty"`
	ModifierID                   *string                 `json:"modifierId,omitempty"`
	ModifierUsername             *string                 `json:"modifierUsername,omitempty"`
	MvnoID                       *string                 `json:"mvnoId,omitempty"`
	Name                         *string                 `json:"name,omitempty"`
	PointCode                    *int                    `json:"pointCode,omitempty"`
	Protocol                     *string                 `json:"protocol,omitempty"`
	ReuseEnable                  *bool                   `json:"reuseEnable,omitempty"`
	RoutingContext               *int                    `json:"routingContext,omitempty"`
	SccpGttList                  []*SccpGttList          `json:"sccpGttList,omitempty"`
	SctpAssociationsList         []*SctpAssociationsList `json:"sctpAssociationsList,omitempty"`
	SgsnIsdnAddress              *string                 `json:"sgsnIsdnAddress,omitempty"`
	SrcGtIndicator               *string                 `json:"srcGtIndicator,omitempty"`
	SrcNatureOfAddressIndicator  *string                 `json:"srcNatureOfAddressIndicator,omitempty"`
	SrcNumberingPlan             *string                 `json:"srcNumberingPlan,omitempty"`
	SrcTransType                 *int                    `json:"srcTransType,omitempty"`
	Type                         *string                 `json:"type,omitempty"`
}

type HlrServiceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type LDAPService struct {
	AdminDomainName        *string     `json:"adminDomainName,omitempty"`
	BaseDomainName         *string     `json:"baseDomainName,omitempty"`
	CreateDateTime         *int        `json:"createDateTime,omitempty"`
	CreatorID              *string     `json:"creatorId,omitempty"`
	CreatorUsername        *string     `json:"creatorUsername,omitempty"`
	Description            *string     `json:"description,omitempty"`
	DomainID               *string     `json:"domainId,omitempty"`
	FriendlyName           *string     `json:"friendlyName,omitempty"`
	ID                     *string     `json:"id,omitempty"`
	IP                     *string     `json:"ip,omitempty"`
	KeyAttribute           *string     `json:"keyAttribute,omitempty"`
	Mappings               []*Mappings `json:"mappings,omitempty"`
	ModifiedDateTime       *int        `json:"modifiedDateTime,omitempty"`
	ModifierID             *string     `json:"modifierId,omitempty"`
	ModifierUsername       *string     `json:"modifierUsername,omitempty"`
	MvnoID                 *string     `json:"mvnoId,omitempty"`
	Name                   *string     `json:"name,omitempty"`
	Password               *string     `json:"password,omitempty"`
	Port                   *int        `json:"port,omitempty"`
	Protocol               *string     `json:"protocol,omitempty"`
	SearchFilter           *string     `json:"searchFilter,omitempty"`
	StandbyAdminDomainName *string     `json:"standbyAdminDomainName,omitempty"`
	StandbyBaseDomainName  *string     `json:"standbyBaseDomainName,omitempty"`
	StandbyIP              *string     `json:"standbyIp,omitempty"`
	StandbyKeyAttribute    *string     `json:"standbyKeyAttribute,omitempty"`
	StandbyPassword        *string     `json:"standbyPassword,omitempty"`
	StandbyPort            *int        `json:"standbyPort,omitempty"`
	StandbySearchFilter    *string     `json:"standbySearchFilter,omitempty"`
	StandbyServerEnabled   *bool       `json:"standbyServerEnabled,omitempty"`
	StandbyTLSEnabled      *bool       `json:"standbyTlsEnabled,omitempty"`
	TLSEnabled             *bool       `json:"tlsEnabled,omitempty"`
	Type                   *string     `json:"type,omitempty"`
}

type LDAPServiceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type MncNdc struct {
	Mcc       *string `json:"mcc,omitempty"`
	Mnc       *string `json:"mnc,omitempty"`
	MvnoID    *string `json:"mvnoId,omitempty"`
	Ndc       *string `json:"ndc,omitempty"`
	ProfileID *string `json:"profileId,omitempty"`
}

type ModifyActiveDirectoryAuthentication struct {
	AdminDomainName             *string     `json:"adminDomainName,omitempty"`
	Description                 *string     `json:"description,omitempty"`
	DomainID                    *string     `json:"domainId,omitempty"`
	FriendlyName                *string     `json:"friendlyName,omitempty"`
	GlobalCatalogEnabled        *bool       `json:"globalCatalogEnabled,omitempty"`
	ID                          *string     `json:"id,omitempty"`
	IP                          *string     `json:"ip,omitempty"`
	Mappings                    []*Mappings `json:"mappings,omitempty"`
	Name                        *string     `json:"name,omitempty"`
	Password                    *string     `json:"password,omitempty"`
	Port                        *int        `json:"port,omitempty"`
	StandbyAdminDomainName      *string     `json:"standbyAdminDomainName,omitempty"`
	StandbyGlobalCatalogEnabled *bool       `json:"standbyGlobalCatalogEnabled,omitempty"`
	StandbyIP                   *string     `json:"standbyIp,omitempty"`
	StandbyPassword             *string     `json:"standbyPassword,omitempty"`
	StandbyPort                 *int        `json:"standbyPort,omitempty"`
	StandbyServerEnabled        *bool       `json:"standbyServerEnabled,omitempty"`
	StandbyTLSEnabled           *bool       `json:"standbyTlsEnabled,omitempty"`
	StandbyWindowsDomainName    *string     `json:"standbyWindowsDomainName,omitempty"`
	TLSEnabled                  *bool       `json:"tlsEnabled,omitempty"`
	Type                        *string     `json:"type,omitempty"`
	WindowsDomainName           *string     `json:"windowsDomainName,omitempty"`
}

type ModifyGroupAttrIdentityUserRoleMapping struct {
	GroupAttr *string   `json:"groupAttr,omitempty"`
	UserRole  *UserRole `json:"userRole,omitempty"`
}

type ModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ModifyHlrAuthentication struct {
	AddressIndicator             *string                 `json:"addressIndicator,omitempty"`
	AuthMapVer                   *string                 `json:"authMapVer,omitempty"`
	AuthorizationCachingEnabled  *bool                   `json:"authorizationCachingEnabled,omitempty"`
	AvCachingEnabled             *bool                   `json:"avCachingEnabled,omitempty"`
	CacheOptionType              *string                 `json:"cacheOptionType,omitempty"`
	CleanUpTimeHour              *int                    `json:"cleanUpTimeHour,omitempty"`
	CleanUpTimeMinute            *int                    `json:"cleanUpTimeMinute,omitempty"`
	CreateDateTime               *int                    `json:"createDateTime,omitempty"`
	CreatorID                    *string                 `json:"creatorId,omitempty"`
	CreatorUsername              *string                 `json:"creatorUsername,omitempty"`
	DefaultPointCodeFormat       *string                 `json:"defaultPointCodeFormat,omitempty"`
	Description                  *string                 `json:"description,omitempty"`
	DestGtIndicator              *string                 `json:"destGtIndicator,omitempty"`
	DestNatureOfAddressIndicator *string                 `json:"destNatureOfAddressIndicator,omitempty"`
	DestNumberingPlan            *string                 `json:"destNumberingPlan,omitempty"`
	DestTransType                *int                    `json:"destTransType,omitempty"`
	DomainID                     *string                 `json:"domainId,omitempty"`
	E164Address                  *string                 `json:"e164Address,omitempty"`
	EapSimMapVer                 *string                 `json:"eapSimMapVer,omitempty"`
	FriendlyName                 *string                 `json:"friendlyName,omitempty"`
	GtPointCode                  *int                    `json:"gtPointCode,omitempty"`
	HasPointCode                 *bool                   `json:"hasPointCode,omitempty"`
	HasSrcPointCode              *bool                   `json:"hasSrcPointCode,omitempty"`
	HasSSN                       *bool                   `json:"hasSSN,omitempty"`
	HistoryTime                  *int                    `json:"historyTime,omitempty"`
	ID                           *string                 `json:"id,omitempty"`
	LocalNetworkIndicator        *string                 `json:"localNetworkIndicator,omitempty"`
	LocalPointCode               *int                    `json:"localPointCode,omitempty"`
	LocationDeliveryEnabled      *bool                   `json:"locationDeliveryEnabled,omitempty"`
	MaxReuseTimes                *int                    `json:"maxReuseTimes,omitempty"`
	MncNdcList                   []*MncNdcList           `json:"mncNdcList,omitempty"`
	ModifiedDateTime             *int                    `json:"modifiedDateTime,omitempty"`
	ModifierID                   *string                 `json:"modifierId,omitempty"`
	ModifierUsername             *string                 `json:"modifierUsername,omitempty"`
	MvnoID                       *string                 `json:"mvnoId,omitempty"`
	Name                         *string                 `json:"name,omitempty"`
	PointCode                    *int                    `json:"pointCode,omitempty"`
	Protocol                     *string                 `json:"protocol,omitempty"`
	ReuseEnable                  *bool                   `json:"reuseEnable,omitempty"`
	RoutingContext               *int                    `json:"routingContext,omitempty"`
	SccpGttList                  []*SccpGttList          `json:"sccpGttList,omitempty"`
	SctpAssociationsList         []*SctpAssociationsList `json:"sctpAssociationsList,omitempty"`
	SgsnIsdnAddress              *string                 `json:"sgsnIsdnAddress,omitempty"`
	SrcGtIndicator               *string                 `json:"srcGtIndicator,omitempty"`
	SrcNatureOfAddressIndicator  *string                 `json:"srcNatureOfAddressIndicator,omitempty"`
	SrcNumberingPlan             *string                 `json:"srcNumberingPlan,omitempty"`
	SrcTransType                 *int                    `json:"srcTransType,omitempty"`
	Type                         *string                 `json:"type,omitempty"`
}

type ModifyLDAPAuthentication struct {
	AdminDomainName        *string     `json:"adminDomainName,omitempty"`
	BaseDomainName         *string     `json:"baseDomainName,omitempty"`
	Description            *string     `json:"description,omitempty"`
	DomainID               *string     `json:"domainId,omitempty"`
	FriendlyName           *string     `json:"friendlyName,omitempty"`
	ID                     *string     `json:"id,omitempty"`
	IP                     *string     `json:"ip,omitempty"`
	KeyAttribute           *string     `json:"keyAttribute,omitempty"`
	Mappings               []*Mappings `json:"mappings,omitempty"`
	Name                   *string     `json:"name,omitempty"`
	Password               *string     `json:"password,omitempty"`
	Port                   *int        `json:"port,omitempty"`
	SearchFilter           *string     `json:"searchFilter,omitempty"`
	StandbyAdminDomainName *string     `json:"standbyAdminDomainName,omitempty"`
	StandbyBaseDomainName  *string     `json:"standbyBaseDomainName,omitempty"`
	StandbyIP              *string     `json:"standbyIp,omitempty"`
	StandbyKeyAttribute    *string     `json:"standbyKeyAttribute,omitempty"`
	StandbyPassword        *string     `json:"standbyPassword,omitempty"`
	StandbyPort            *int        `json:"standbyPort,omitempty"`
	StandbySearchFilter    *string     `json:"standbySearchFilter,omitempty"`
	StandbyServerEnabled   *bool       `json:"standbyServerEnabled,omitempty"`
	StandbyTLSEnabled      *bool       `json:"standbyTlsEnabled,omitempty"`
	TLSEnabled             *bool       `json:"tlsEnabled,omitempty"`
	Type                   *string     `json:"type,omitempty"`
}

type ModifyLocalDbAuthentication struct {
	Description  *string     `json:"description,omitempty"`
	DomainID     *string     `json:"domainId,omitempty"`
	FriendlyName *string     `json:"friendlyName,omitempty"`
	ID           *string     `json:"id,omitempty"`
	Mappings     []*Mappings `json:"mappings,omitempty"`
	MvnoID       *string     `json:"mvnoId,omitempty"`
	Name         *string     `json:"name,omitempty"`
	Protocol     *string     `json:"protocol,omitempty"`
	Type         *string     `json:"type,omitempty"`
}

type ModifyRadiusAccounting struct {
	Description          *string                   `json:"description,omitempty"`
	DomainID             *string                   `json:"domainId,omitempty"`
	HealthCheckPolicy    *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
	ID                   *string                   `json:"id,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
	Primary              *common.RadiusServer      `json:"primary,omitempty"`
	Protocol             *string                   `json:"protocol,omitempty"`
	RateLimiting         *common.RateLimiting      `json:"rateLimiting,omitempty"`
	Secondary            *Secondary                `json:"secondary,omitempty"`
	StandbyPrimary       *common.RadiusServer      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled *bool                     `json:"standbyServerEnabled,omitempty"`
	Type                 *string                   `json:"type,omitempty"`
}

type ModifyRadiusAuthentication struct {
	Description             *string                   `json:"description,omitempty"`
	DomainID                *string                   `json:"domainId,omitempty"`
	FriendlyName            *string                   `json:"friendlyName,omitempty"`
	HealthCheckPolicy       *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
	ID                      *string                   `json:"id,omitempty"`
	LocationDeliveryEnabled *bool                     `json:"locationDeliveryEnabled,omitempty"`
	Mappings                []*Mappings               `json:"mappings,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	Primary                 *common.RadiusServer      `json:"primary,omitempty"`
	RateLimiting            *common.RateLimiting      `json:"rateLimiting,omitempty"`
	Secondary               *Secondary                `json:"secondary,omitempty"`
	StandbyPrimary          *common.RadiusServer      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled    *bool                     `json:"standbyServerEnabled,omitempty"`
	Type                    *string                   `json:"type,omitempty"`
}

type RadiusAccountingService struct {
	CreateDateTime       *int                      `json:"createDateTime,omitempty"`
	CreatorID            *string                   `json:"creatorId,omitempty"`
	CreatorUsername      *string                   `json:"creatorUsername,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	DomainID             *string                   `json:"domainId,omitempty"`
	HealthCheckPolicy    *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
	ID                   *string                   `json:"id,omitempty"`
	ModifiedDateTime     *int                      `json:"modifiedDateTime,omitempty"`
	ModifierID           *string                   `json:"modifierId,omitempty"`
	ModifierUsername     *string                   `json:"modifierUsername,omitempty"`
	MvnoID               *string                   `json:"mvnoId,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
	Primary              *common.RadiusServer      `json:"primary,omitempty"`
	Protocol             *string                   `json:"protocol,omitempty"`
	RateLimiting         *common.RateLimiting      `json:"rateLimiting,omitempty"`
	Secondary            *Secondary                `json:"secondary,omitempty"`
	StandbyPrimary       *common.RadiusServer      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled *bool                     `json:"standbyServerEnabled,omitempty"`
	Type                 *string                   `json:"type,omitempty"`
}

type RadiusAccountingServiceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type RadiusAuthenticationService struct {
	CreateDateTime          *int                      `json:"createDateTime,omitempty"`
	CreatorID               *string                   `json:"creatorId,omitempty"`
	CreatorUsername         *string                   `json:"creatorUsername,omitempty"`
	Description             *string                   `json:"description,omitempty"`
	DomainID                *string                   `json:"domainId,omitempty"`
	FriendlyName            *string                   `json:"friendlyName,omitempty"`
	HealthCheckPolicy       *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
	ID                      *string                   `json:"id,omitempty"`
	LocationDeliveryEnabled *bool                     `json:"locationDeliveryEnabled,omitempty"`
	Mappings                []*Mappings               `json:"mappings,omitempty"`
	ModifiedDateTime        *int                      `json:"modifiedDateTime,omitempty"`
	ModifierID              *string                   `json:"modifierId,omitempty"`
	ModifierUsername        *string                   `json:"modifierUsername,omitempty"`
	MvnoID                  *string                   `json:"mvnoId,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	Primary                 *common.RadiusServer      `json:"primary,omitempty"`
	Protocol                *string                   `json:"protocol,omitempty"`
	RateLimiting            *common.RateLimiting      `json:"rateLimiting,omitempty"`
	Secondary               *Secondary                `json:"secondary,omitempty"`
	StandbyPrimary          *common.RadiusServer      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled    *bool                     `json:"standbyServerEnabled,omitempty"`
	Type                    *string                   `json:"type,omitempty"`
}

type RadiusAuthenticationServiceList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type SccpGtt struct {
	AddressIndicator         *string `json:"addressIndicator,omitempty"`
	E164Address              *string `json:"e164Address,omitempty"`
	GtDigits                 *string `json:"gtDigits,omitempty"`
	GtIndicator              *string `json:"gtIndicator,omitempty"`
	HasPointCode             *bool   `json:"hasPointCode,omitempty"`
	HasSSN                   *bool   `json:"hasSSN,omitempty"`
	NatureOfAddressIndicator *string `json:"natureOfAddressIndicator,omitempty"`
	NumberingPlan            *string `json:"numberingPlan,omitempty"`
	PointCode                *int    `json:"pointCode,omitempty"`
	TransType                *int    `json:"transType,omitempty"`
}

type SctpAssociation struct {
	AdjPointCode        *string `json:"adjPointCode,omitempty"`
	DestinationIP       *string `json:"destinationIp,omitempty"`
	DestinationPort     *int    `json:"destinationPort,omitempty"`
	MaxInboundsStreams  *int    `json:"maxInboundsStreams,omitempty"`
	MaxOutboundsStreams *int    `json:"maxOutboundsStreams,omitempty"`
	NodeTermination     *string `json:"nodeTermination,omitempty"`
	SourcePort          *int    `json:"sourcePort,omitempty"`
}

type SecondaryRadiusServer struct {
	AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
	IP                  *string `json:"ip,omitempty"`
	Port                *int    `json:"port,omitempty"`
	SharedSecret        *string `json:"sharedSecret,omitempty"`
}

type TestingConfig struct {
	ID           *string       `json:"id,omitempty"`
	LoginRequest *LoginRequest `json:"loginRequest,omitempty"`
}

type TestingConfigLoginRequestType struct {
	Password          *string `json:"password,omitempty"`
	Protocol          *string `json:"protocol,omitempty"`
	TimeZoneUtcOffset *string `json:"timeZoneUtcOffset,omitempty"`
	UserName          *string `json:"userName,omitempty"`
}
