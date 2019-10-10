package service

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ActiveDirectoryService struct {
	AdminDomainName *common.NormalName2to64 `json:"adminDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the active directory authentication service
	Id *string `json:"id,omitempty"`

	Ip *common.IpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=AD"`

	StandbyAdminDomainName *common.NormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *common.IpAddress `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port for standby cluster
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled for standby cluster
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *common.NormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=AD"`

	WindowsDomainName *common.NormalName2to64 `json:"windowsDomainName,omitempty"`
}

type ActiveDirectoryServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ActiveDirectoryService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CommonAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the accounting service
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS CGF"`

	// Type
	// Accounting protocol same as protocol.
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS CGF"`
}

type CommonAccountingServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CommonAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CommonAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS AD LDAP FACEBOOK LINKEDIN GOOGLE GENERICOAUTH SOAP HLR LOCAL_DB GUEST"`

	// Type
	// Authentication protocol same as protocol.
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS AD LDAP FACEBOOK LINKEDIN GOOGLE GENERICOAUTH SOAP HLR LOCAL_DB GUEST"`
}

type CommonAuthenticationServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CommonAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CreateActiveDirectoryAuthentication struct {
	AdminDomainName *common.NormalName2to64 `json:"adminDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled" validate:"required"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *common.IpAddress `json:"ip" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	StandbyAdminDomainName *common.NormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *common.IpAddress `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port for standby cluster
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled for standby cluster
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *common.NormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled" validate:"required"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=AD"`

	WindowsDomainName *common.NormalName2to64 `json:"windowsDomainName" validate:"required"`
}

type CreateHlrAuthentication struct {
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	AuthMapVer *string `json:"authMapVer,omitempty" validate:"oneof=version2 version3"`

	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"oneof=version2 version3"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"oneof=international international_spare national national_spare"`

	LocalPointCode *int `json:"localPointCode,omitempty"`

	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	MncNdcList []*MncNdc `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	PointCode *int `json:"pointCode,omitempty"`

	Protocol *string `json:"protocol,omitempty" validate:"oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*SccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*SctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	Type *string `json:"type,omitempty" validate:"oneof=HLR"`
}

type CreateLDAPAuthentication struct {
	AdminDomainName *common.NormalName2to128 `json:"adminDomainName" validate:"required"`

	BaseDomainName *common.NormalName2to64 `json:"baseDomainName" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *common.IpAddress `json:"ip" validate:"required"`

	KeyAttribute *common.NormalName2to64 `json:"keyAttribute" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	// Password
	// Admin password
	Password *string `json:"password" validate:"required"`

	// Port
	// Port
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	SearchFilter *common.NormalName2to64 `json:"searchFilter" validate:"required"`

	StandbyAdminDomainName *common.NormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *common.NormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *common.IpAddress `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *common.NormalName2to64 `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

	StandbySearchFilter *common.NormalName2to64 `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TlsEnabled
	// LDAP over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled" validate:"required"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=LDAP"`
}

type CreateRadiusAccounting struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	Primary *common.RadiusServer `json:"primary" validate:"required"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

type CreateRadiusAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	Primary *common.RadiusServer `json:"primary" validate:"required"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication server type
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

type DeleteBulkAccountingService struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type DeleteBulkAuthenticationService struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type DnsServer struct {
	// Ip
	// IP of server
	Ip *string `json:"ip,omitempty"`
}

type DnsServerList []*DnsServer

type Ggsn struct {
	// DomainName
	// Domain name of GGSN
	DomainName *string `json:"domainName,omitempty"`

	// GgsnIPAddress
	// IP of GGSN
	GgsnIPAddress *string `json:"ggsnIPAddress,omitempty"`
}

type GgsnConfig struct {
	DnsServerList DnsServerList `json:"dnsServerList,omitempty"`

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
	GroupAttr *string `json:"groupAttr" validate:"required"`

	// Id
	// Group attribute mapping UUID
	Id *string `json:"id,omitempty"`

	// UserRole
	// Identity user role
	UserRole *GroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

// GroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type GroupAttrIdentityUserRoleMappingUserRoleType struct {
	// Id
	// Identity user role UUID
	Id *string `json:"id,omitempty"`

	Name *common.NormalName2to64 `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	UserTrafficProfile *GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

// GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Identity user role
type GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	// Id
	// User traffic profile UUID
	Id *string `json:"id,omitempty"`

	// Name
	// User traffic profile name
	Name *string `json:"name,omitempty"`
}

type GtpSettings struct {
	// DnsNumberOfRetries
	// DNS Number of Retries
	DnsNumberOfRetries *int `json:"dnsNumberOfRetries,omitempty" validate:"gte=1,lte=10"`

	// EchoRequestTimer
	// Echo Request Timerr
	EchoRequestTimer *int `json:"echoRequestTimer,omitempty" validate:"gte=60,lte=300"`

	// NumberOfRetries
	// Number of Retries
	NumberOfRetries *int `json:"numberOfRetries,omitempty" validate:"gte=3,lte=6"`

	// ResponseTimeout
	// DNS Response Timeout
	ResponseTimeout *int `json:"responseTimeout,omitempty" validate:"gte=1,lte=10"`

	// T3ResponseTimer
	// Response Timer
	T3ResponseTimer *int `json:"t3ResponseTimer,omitempty" validate:"gte=2,lte=5"`
}

type HlrService struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	// AuthMapVer
	// - For HLR Authentiaction server
	AuthMapVer *string `json:"authMapVer,omitempty" validate:"oneof=version2 version3"`

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

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultPointCodeFormat
	// - For HLR Authentiaction server
	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DestGtIndicator
	// - For HLR Authentiaction server
	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// DestNatureOfAddressIndicator
	// - For HLR Authentiaction server
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// DestNumberingPlan
	// - For HLR Authentiaction server
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	// DestTransType
	// - For HLR Authentiaction server
	DestTransType *int `json:"destTransType,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// E164Address
	// - For HLR Authentiaction server
	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// - For HLR Authentiaction server
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"oneof=version2 version3"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

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

	// Id
	// Identifier of the HLR authentication service
	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// - For HLR Authentiaction server
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"oneof=international international_spare national national_spare"`

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

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PointCode
	// - For HLR Authentiaction server
	PointCode *int `json:"pointCode,omitempty"`

	Protocol *string `json:"protocol,omitempty" validate:"oneof=HLR"`

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
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// SrcNatureOfAddressIndicator
	// - For HLR Authentiaction server
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// SrcNumberingPlan
	// - For HLR Authentiaction server
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	// SrcTransType
	// - For HLR Authentiaction server
	SrcTransType *int `json:"srcTransType,omitempty"`

	Type *string `json:"type,omitempty" validate:"oneof=HLR"`
}

type HlrServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*HlrService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type LDAPService struct {
	AdminDomainName *common.NormalName2to128 `json:"adminDomainName,omitempty"`

	BaseDomainName *common.NormalName2to64 `json:"baseDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the LDAP authentication service
	Id *string `json:"id,omitempty"`

	Ip *common.IpAddress `json:"ip,omitempty"`

	KeyAttribute *common.NormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	// Protocol
	// Authentication protocol
	Protocol *string `json:"protocol,omitempty" validate:"oneof=LDAP"`

	SearchFilter *common.NormalName2to64 `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *common.NormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *common.NormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *common.IpAddress `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *common.NormalName2to64 `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

	StandbySearchFilter *common.NormalName2to64 `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TlsEnabled
	// LDAP over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=LDAP"`
}

type LDAPServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

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

	// MvnoId
	// MVNO ID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Ndc
	// NDC
	Ndc *string `json:"ndc,omitempty"`

	// ProfileId
	// Profile ID
	ProfileId *string `json:"profileId,omitempty"`
}

type ModifyActiveDirectoryAuthentication struct {
	AdminDomainName *common.NormalName2to64 `json:"adminDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *common.IpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	StandbyAdminDomainName *common.NormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *common.IpAddress `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port standby cluster
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled standby cluster
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *common.NormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=AD"`

	WindowsDomainName *common.NormalName2to64 `json:"windowsDomainName,omitempty"`
}

// ModifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type ModifyGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	GroupAttr *string `json:"groupAttr" validate:"required"`

	// UserRole
	// Identity user role
	UserRole *ModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

// ModifyGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type ModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// Id
	// Identity user role UUID
	Id *string `json:"id,omitempty"`

	Name *common.NormalName2to64 `json:"name,omitempty"`
}

type ModifyHlrAuthentication struct {
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	AuthMapVer *string `json:"authMapVer,omitempty" validate:"oneof=version2 version3"`

	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"oneof=version2 version3"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"oneof=international international_spare national national_spare"`

	LocalPointCode *int `json:"localPointCode,omitempty"`

	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	MncNdcList []*MncNdc `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PointCode *int `json:"pointCode,omitempty"`

	Protocol *string `json:"protocol,omitempty" validate:"oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*SccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*SctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	Type *string `json:"type,omitempty" validate:"oneof=HLR"`
}

type ModifyLDAPAuthentication struct {
	AdminDomainName *common.NormalName2to128 `json:"adminDomainName,omitempty"`

	BaseDomainName *common.NormalName2to64 `json:"baseDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *common.IpAddress `json:"ip,omitempty"`

	KeyAttribute *common.NormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	SearchFilter *common.NormalName2to64 `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *common.NormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *common.NormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *common.IpAddress `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *common.NormalName2to64 `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

	StandbySearchFilter *common.NormalName2to64 `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TlsEnabled
	// LDAP over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=LDAP"`
}

type ModifyLocalDbAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=LOCAL_DB"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=LOCAL_DB"`
}

type ModifyRadiusAccounting struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS accounting service
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

type ModifyRadiusAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication server type
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

type RadiusAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS accounting service
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

type RadiusAccountingServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RadiusAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type RadiusAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	// Protocol
	// Authentication protocol.
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol.
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

type RadiusAuthenticationServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RadiusAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SccpGtt struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	// E164Address
	// - For HLR Authentiaction server
	E164Address *string `json:"e164Address,omitempty"`

	// GtDigits
	// GT digits
	GtDigits *string `json:"gtDigits,omitempty"`

	// GtIndicator
	// - For HLR Authentiaction server
	GtIndicator *string `json:"gtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// HasPointCode
	// Enable Point Code submitted
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSSN
	// Enable SSN- For HLR Authentiaction server
	HasSSN *bool `json:"hasSSN,omitempty"`

	// NatureOfAddressIndicator
	// - For HLR Authentiaction server
	NatureOfAddressIndicator *string `json:"natureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// NumberingPlan
	// - For HLR Authentiaction server
	NumberingPlan *string `json:"numberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

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

	// DestinationIp
	// Destination IP
	DestinationIp *string `json:"destinationIp,omitempty"`

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
	NodeTermination *string `json:"nodeTermination,omitempty" validate:"oneof=active backup both"`

	// SourcePort
	// Source port
	SourcePort *int `json:"sourcePort,omitempty"`
}

type SecondaryRadiusServer struct {
	// AutoFallbackDisable
	// Automatic fallback enabled or disabled
	AutoFallbackDisable *bool `json:"autoFallbackDisable" validate:"required"`

	Ip *common.IpAddress `json:"ip" validate:"required"`

	// Port
	// RADIUS server port
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// SharedSecret
	// RADIUS server shared secrect
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

type TestingConfig struct {
	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

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
