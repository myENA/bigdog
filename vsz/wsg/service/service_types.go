package service

// API Version: v8_0

type ActiveDirectoryService struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// ID
	// Identifier of the active directory authentication service
	ID *string `json:"id,omitempty"`

	IP *string `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIP *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port for standby cluster
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTLSEnabled
	// AD over TLS Enabled for standby cluster
	StandbyTLSEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// TLSEnabled
	// AD over TLS Enabled
	TLSEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`

	WindowsDomainName *string `json:"windowsDomainName,omitempty"`
}

type ActiveDirectoryServiceList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ActiveDirectoryService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CommonAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the accounting service
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Accounting protocol same as protocol.
	Type *string `json:"type,omitempty"`
}

type CommonAccountingServiceList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CommonAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CommonAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Authentication protocol same as protocol.
	Type *string `json:"type,omitempty"`
}

type CommonAuthenticationServiceList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CommonAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CreateActiveDirectoryAuthentication struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	IP *string `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIP *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port for standby cluster
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTLSEnabled
	// AD over TLS Enabled for standby cluster
	StandbyTLSEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// TLSEnabled
	// AD over TLS Enabled
	TLSEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`

	WindowsDomainName *string `json:"windowsDomainName,omitempty"`
}

type CreateHlrAuthentication struct {
	AddressIndicator *string `json:"addressIndicator,omitempty"`

	AuthMapVer *string `json:"authMapVer,omitempty"`

	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorID *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *string `json:"description,omitempty"`

	DestGtIndicator *string `json:"destGtIndicator,omitempty"`

	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty"`

	DestNumberingPlan *string `json:"destNumberingPlan,omitempty"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	EapSimMapVer *string `json:"eapSimMapVer,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	ID *string `json:"id,omitempty"`

	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty"`

	LocalPointCode *int `json:"localPointCode,omitempty"`

	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	MncNdcList []*MncNdc `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierID *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	PointCode *int `json:"pointCode,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*SccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*SctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	SrcGtIndicator *string `json:"srcGtIndicator,omitempty"`

	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty"`

	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	Type *string `json:"type,omitempty"`
}

type CreateLDAPAuthentication struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	BaseDomainName *string `json:"baseDomainName,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	IP *string `json:"ip,omitempty"`

	KeyAttribute *string `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	SearchFilter *string `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	StandbyIP *string `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTLSEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	StandbyTLSEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TLSEnabled
	// LDAP over TLS Enabled
	TLSEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`
}

type CreateRadiusAccounting struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	Type *string `json:"type,omitempty"`
}

type CreateRadiusAuthentication struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication server type
	Type *string `json:"type,omitempty"`
}

type DeleteBulkAccountingService struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DeleteBulkAuthenticationService struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type DNSServer struct {
	// IP
	// IP of server
	IP *string `json:"ip,omitempty"`
}

type DNSServerList []*DNSServer

type Ggsn struct {
	// DomainName
	// Domain name of GGSN
	DomainName *string `json:"domainName,omitempty"`

	// GgsnIPAddress
	// IP of GGSN
	GgsnIPAddress *string `json:"ggsnIPAddress,omitempty"`
}

type GgsnConfig struct {
	DNSServerList DNSServerList `json:"dnsServerList,omitempty"`

	GgsnList GgsnList `json:"ggsnList,omitempty"`

	GtpSettings *GtpSettings `json:"gtpSettings,omitempty"`
}

type GgsnList []*Ggsn

// GroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type GroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	GroupAttr *string `json:"groupAttr,omitempty"`

	// ID
	// Group attribute mapping UUID
	ID *string `json:"id,omitempty"`

	// UserRole
	// Identity user role
	UserRole *GroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole,omitempty"`
}

// GroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type GroupAttrIdentityUserRoleMappingUserRoleType struct {
	// ID
	// Identity user role UUID
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	UserTrafficProfile *GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

// GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Identity user role
type GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	// ID
	// User traffic profile UUID
	ID *string `json:"id,omitempty"`

	// Name
	// User traffic profile name
	Name *string `json:"name,omitempty"`
}

type GtpSettings struct {
	// DNSNumberOfRetries
	// DNS Number of Retries
	DNSNumberOfRetries *int `json:"dnsNumberOfRetries,omitempty"`

	// EchoRequestTimer
	// Echo Request Timerr
	EchoRequestTimer *int `json:"echoRequestTimer,omitempty"`

	// NumberOfRetries
	// Number of Retries
	NumberOfRetries *int `json:"numberOfRetries,omitempty"`

	// ResponseTimeout
	// DNS Response Timeout
	ResponseTimeout *int `json:"responseTimeout,omitempty"`

	// T3ResponseTimer
	// Response Timer
	T3ResponseTimer *int `json:"t3ResponseTimer,omitempty"`
}

type HlrService struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	AddressIndicator *string `json:"addressIndicator,omitempty"`

	// AuthMapVer
	// - For HLR Authentiaction server
	AuthMapVer *string `json:"authMapVer,omitempty"`

	// AuthorizationCachingEnabled
	// - For HLR Authentiaction server
	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	// AvCachingEnabled
	// - For HLR Authentiaction server
	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	// CacheOptionType
	// - For HLR Authentiaction server
	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	// CleanUpTimeHour
	// - For HLR Authentiaction server
	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	// CleanUpTimeMinute
	// - For HLR Authentiaction server
	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultPointCodeFormat
	// - For HLR Authentiaction server
	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *string `json:"description,omitempty"`

	// DestGtIndicator
	// - For HLR Authentiaction server
	DestGtIndicator *string `json:"destGtIndicator,omitempty"`

	// DestNatureOfAddressIndicator
	// - For HLR Authentiaction server
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty"`

	// DestNumberingPlan
	// - For HLR Authentiaction server
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty"`

	// DestTransType
	// - For HLR Authentiaction server
	DestTransType *int `json:"destTransType,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// E164Address
	// - For HLR Authentiaction server
	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// - For HLR Authentiaction server
	EapSimMapVer *string `json:"eapSimMapVer,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// GtPointCode
	// - For HLR Authentiaction server
	GtPointCode *int `json:"gtPointCode,omitempty"`

	// HasPointCode
	// - For HLR Authentiaction server
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSrcPointCode
	// - For HLR Authentiaction server
	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	// HasSSN
	// - For HLR Authentiaction server
	HasSSN *bool `json:"hasSSN,omitempty"`

	// HistoryTime
	// - For HLR Authentiaction server
	HistoryTime *int `json:"historyTime,omitempty"`

	// ID
	// Identifier of the HLR authentication service
	ID *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// - For HLR Authentiaction server
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty"`

	// LocalPointCode
	// - For HLR Authentiaction server
	LocalPointCode *int `json:"localPointCode,omitempty"`

	// LocationDeliveryEnabled
	// - For HLR Authentiaction server
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// MaxReuseTimes
	// - For HLR Authentiaction server
	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	// MncNdcList
	// - For HLR Authentiaction server
	MncNdcList []*MncNdc `json:"mncNdcList,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// PointCode
	// - For HLR Authentiaction server
	PointCode *int `json:"pointCode,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	// ReuseEnable
	// - For HLR Authentiaction server
	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	// RoutingContext
	// - For HLR Authentiaction server
	RoutingContext *int `json:"routingContext,omitempty"`

	// SccpGttList
	// - For HLR Authentiaction server
	SccpGttList []*SccpGtt `json:"sccpGttList,omitempty"`

	// SctpAssociationsList
	// - For HLR Authentiaction server
	SctpAssociationsList []*SctpAssociation `json:"sctpAssociationsList,omitempty"`

	// SgsnIsdnAddress
	// - For HLR Authentiaction server
	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// - For HLR Authentiaction server
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty"`

	// SrcNatureOfAddressIndicator
	// - For HLR Authentiaction server
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty"`

	// SrcNumberingPlan
	// - For HLR Authentiaction server
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty"`

	// SrcTransType
	// - For HLR Authentiaction server
	SrcTransType *int `json:"srcTransType,omitempty"`

	Type *string `json:"type,omitempty"`
}

type HlrServiceList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*HlrService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type LDAPService struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	BaseDomainName *string `json:"baseDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// ID
	// Identifier of the LDAP authentication service
	ID *string `json:"id,omitempty"`

	IP *string `json:"ip,omitempty"`

	KeyAttribute *string `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	// Protocol
	// Authentication protocol
	Protocol *string `json:"protocol,omitempty"`

	SearchFilter *string `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	StandbyIP *string `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTLSEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	StandbyTLSEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TLSEnabled
	// LDAP over TLS Enabled
	TLSEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`
}

type LDAPServiceList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LDAPService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type MncNdc struct {
	// Mcc
	// MCC
	Mcc *string `json:"mcc,omitempty"`

	// Mnc
	// MNC
	Mnc *string `json:"mnc,omitempty"`

	// MvnoID
	// MVNO ID
	MvnoID *string `json:"mvnoId,omitempty"`

	// Ndc
	// NDC
	Ndc *string `json:"ndc,omitempty"`

	// ProfileID
	// Profile ID
	ProfileID *string `json:"profileId,omitempty"`
}

type ModifyActiveDirectoryAuthentication struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	IP *string `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIP *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port standby cluster
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTLSEnabled
	// AD over TLS Enabled standby cluster
	StandbyTLSEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// TLSEnabled
	// AD over TLS Enabled
	TLSEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`

	WindowsDomainName *string `json:"windowsDomainName,omitempty"`
}

// ModifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type ModifyGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	GroupAttr *string `json:"groupAttr,omitempty"`

	// UserRole
	// Identity user role
	UserRole *ModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole,omitempty"`
}

// ModifyGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type ModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// ID
	// Identity user role UUID
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type ModifyHlrAuthentication struct {
	AddressIndicator *string `json:"addressIndicator,omitempty"`

	AuthMapVer *string `json:"authMapVer,omitempty"`

	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorID *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *string `json:"description,omitempty"`

	DestGtIndicator *string `json:"destGtIndicator,omitempty"`

	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty"`

	DestNumberingPlan *string `json:"destNumberingPlan,omitempty"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	EapSimMapVer *string `json:"eapSimMapVer,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	ID *string `json:"id,omitempty"`

	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty"`

	LocalPointCode *int `json:"localPointCode,omitempty"`

	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	MncNdcList []*MncNdc `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierID *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	PointCode *int `json:"pointCode,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*SccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*SctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	SrcGtIndicator *string `json:"srcGtIndicator,omitempty"`

	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty"`

	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	Type *string `json:"type,omitempty"`
}

type ModifyLDAPAuthentication struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	BaseDomainName *string `json:"baseDomainName,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	IP *string `json:"ip,omitempty"`

	KeyAttribute *string `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	SearchFilter *string `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	StandbyIP *string `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTLSEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	StandbyTLSEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TLSEnabled
	// LDAP over TLS Enabled
	TLSEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`
}

type ModifyLocalDbAuthentication struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`
}

type ModifyRadiusAccounting struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// ID
	// Identifier of the RADIUS accounting service
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	Type *string `json:"type,omitempty"`
}

type ModifyRadiusAuthentication struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication server type
	Type *string `json:"type,omitempty"`
}

type RadiusAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// ID
	// Identifier of the RADIUS accounting service
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	Type *string `json:"type,omitempty"`
}

type RadiusAccountingServiceList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RadiusAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type RadiusAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// ID
	// Identifier of the RADIUS authentication service
	ID *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty"`
}

type RadiusAuthenticationServiceList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RadiusAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SccpGtt struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	AddressIndicator *string `json:"addressIndicator,omitempty"`

	// E164Address
	// - For HLR Authentiaction server
	E164Address *string `json:"e164Address,omitempty"`

	// GtDigits
	// GT digits
	GtDigits *string `json:"gtDigits,omitempty"`

	// GtIndicator
	// - For HLR Authentiaction server
	GtIndicator *string `json:"gtIndicator,omitempty"`

	// HasPointCode
	// Enable Point Code submitted
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSSN
	// Enable SSN- For HLR Authentiaction server
	HasSSN *bool `json:"hasSSN,omitempty"`

	// NatureOfAddressIndicator
	// - For HLR Authentiaction server
	NatureOfAddressIndicator *string `json:"natureOfAddressIndicator,omitempty"`

	// NumberingPlan
	// - For HLR Authentiaction server
	NumberingPlan *string `json:"numberingPlan,omitempty"`

	// PointCode
	// - For HLR Authentiaction server
	PointCode *int `json:"pointCode,omitempty"`

	// TransType
	// - For HLR Authentiaction server
	TransType *int `json:"transType,omitempty"`
}

type SctpAssociation struct {
	// AdjPointCode
	// Adj Pointcode
	AdjPointCode *string `json:"adjPointCode,omitempty"`

	// DestinationIP
	// Destination IP
	DestinationIP *string `json:"destinationIp,omitempty"`

	// DestinationPort
	// Destination Port
	DestinationPort *int `json:"destinationPort,omitempty"`

	// MaxInboundsStreams
	// NDC
	MaxInboundsStreams *int `json:"maxInboundsStreams,omitempty"`

	// MaxOutboundsStreams
	// Profile ID
	MaxOutboundsStreams *int `json:"maxOutboundsStreams,omitempty"`

	// NodeTermination
	// Node termination
	NodeTermination *string `json:"nodeTermination,omitempty"`

	// SourcePort
	// Source port
	SourcePort *int `json:"sourcePort,omitempty"`
}

type SecondaryRadiusServer struct {
	// AutoFallbackDisable
	// Automatic fallback enabled or disabled
	AutoFallbackDisable *bool `json:"autoFallbackDisable,omitempty"`

	IP *string `json:"ip,omitempty"`

	// Port
	// RADIUS server port
	Port *int `json:"port,omitempty"`

	// SharedSecret
	// RADIUS server shared secrect
	SharedSecret *string `json:"sharedSecret,omitempty"`
}

type TestingConfig struct {
	// ID
	// Identifier of the authentication service
	ID *string `json:"id,omitempty"`

	LoginRequest *TestingConfigLoginRequestType `json:"loginRequest,omitempty"`
}

type TestingConfigLoginRequestType struct {
	// Password
	// password for test user
	Password *string `json:"password,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	// TimeZoneUtcOffset
	// timezone offset, ex: '+8'
	TimeZoneUtcOffset *string `json:"timeZoneUtcOffset,omitempty"`

	// UserName
	// name for test user
	UserName *string `json:"userName,omitempty"`
}
