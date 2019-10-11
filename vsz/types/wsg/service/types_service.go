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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=AD"`

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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AD"`

	WindowsDomainName *common.NormalName2to64 `json:"windowsDomainName,omitempty"`
}

func NewActiveDirectoryService() *ActiveDirectoryService {
	activeDirectoryServiceType := new(ActiveDirectoryService)
	return activeDirectoryServiceType
}

func NewActiveDirectoryServiceWithDefaults() *ActiveDirectoryService {
	activeDirectoryServiceType := new(ActiveDirectoryService)
	portField := 389
	activeDirectoryServiceType.Port = &portField
	standbyPortField := 389
	activeDirectoryServiceType.StandbyPort = &standbyPortField
	return activeDirectoryServiceType
}

type ActiveDirectoryServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ActiveDirectoryService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewActiveDirectoryServiceList() *ActiveDirectoryServiceList {
	activeDirectoryServiceListType := new(ActiveDirectoryServiceList)
	return activeDirectoryServiceListType
}

func NewActiveDirectoryServiceListWithDefaults() *ActiveDirectoryServiceList {
	activeDirectoryServiceListType := new(ActiveDirectoryServiceList)
	return activeDirectoryServiceListType
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
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,CGF]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS CGF"`

	// Type
	// Accounting protocol same as protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,CGF]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS CGF"`
}

func NewCommonAccountingService() *CommonAccountingService {
	commonAccountingServiceType := new(CommonAccountingService)
	return commonAccountingServiceType
}

func NewCommonAccountingServiceWithDefaults() *CommonAccountingService {
	commonAccountingServiceType := new(CommonAccountingService)
	return commonAccountingServiceType
}

type CommonAccountingServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CommonAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewCommonAccountingServiceList() *CommonAccountingServiceList {
	commonAccountingServiceListType := new(CommonAccountingServiceList)
	return commonAccountingServiceListType
}

func NewCommonAccountingServiceListWithDefaults() *CommonAccountingServiceList {
	commonAccountingServiceListType := new(CommonAccountingServiceList)
	return commonAccountingServiceListType
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
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,AD,LDAP,FACEBOOK,LINKEDIN,GOOGLE,GENERICOAUTH,SOAP,HLR,LOCAL_DB,GUEST]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS AD LDAP FACEBOOK LINKEDIN GOOGLE GENERICOAUTH SOAP HLR LOCAL_DB GUEST"`

	// Type
	// Authentication protocol same as protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,AD,LDAP,FACEBOOK,LINKEDIN,GOOGLE,GENERICOAUTH,SOAP,HLR,LOCAL_DB,GUEST]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS AD LDAP FACEBOOK LINKEDIN GOOGLE GENERICOAUTH SOAP HLR LOCAL_DB GUEST"`
}

func NewCommonAuthenticationService() *CommonAuthenticationService {
	commonAuthenticationServiceType := new(CommonAuthenticationService)
	return commonAuthenticationServiceType
}

func NewCommonAuthenticationServiceWithDefaults() *CommonAuthenticationService {
	commonAuthenticationServiceType := new(CommonAuthenticationService)
	return commonAuthenticationServiceType
}

type CommonAuthenticationServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CommonAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewCommonAuthenticationServiceList() *CommonAuthenticationServiceList {
	commonAuthenticationServiceListType := new(CommonAuthenticationServiceList)
	return commonAuthenticationServiceListType
}

func NewCommonAuthenticationServiceListWithDefaults() *CommonAuthenticationServiceList {
	commonAuthenticationServiceListType := new(CommonAuthenticationServiceList)
	return commonAuthenticationServiceListType
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
	// Constraints:
	//    - required
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled" validate:"required"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - required
	//    - default:389
	//    - min:1
	//    - max:65535
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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled for standby cluster
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *common.NormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	// Constraints:
	//    - required
	TlsEnabled *bool `json:"tlsEnabled" validate:"required"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AD"`

	// WindowsDomainName
	// Constraints:
	//    - required
	WindowsDomainName *common.NormalName2to64 `json:"windowsDomainName" validate:"required"`
}

func NewCreateActiveDirectoryAuthentication() *CreateActiveDirectoryAuthentication {
	createActiveDirectoryAuthenticationType := new(CreateActiveDirectoryAuthentication)
	return createActiveDirectoryAuthenticationType
}

func NewCreateActiveDirectoryAuthenticationWithDefaults() *CreateActiveDirectoryAuthentication {
	createActiveDirectoryAuthenticationType := new(CreateActiveDirectoryAuthentication)
	portField := 389
	createActiveDirectoryAuthenticationType.Port = &portField
	standbyPortField := 389
	createActiveDirectoryAuthenticationType.StandbyPort = &standbyPortField
	standbyServerEnabledField := false
	createActiveDirectoryAuthenticationType.StandbyServerEnabled = &standbyServerEnabledField
	return createActiveDirectoryAuthenticationType
}

type CreateHlrAuthentication struct {
	// AddressIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"omitempty,oneof=route_on_gt route_on_ssn"`

	// AuthMapVer
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	AuthMapVer *string `json:"authMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

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

	// DestGtIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// DestNatureOfAddressIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"omitempty,oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// DestNumberingPlan
	// Constraints:
	//    - nullable
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"omitempty,oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"omitempty,oneof=international international_spare national national_spare"`

	LocalPointCode *int `json:"localPointCode,omitempty"`

	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	MncNdcList []*MncNdc `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*SccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*SctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// SrcNatureOfAddressIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"omitempty,oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// SrcNumberingPlan
	// Constraints:
	//    - nullable
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"omitempty,oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=HLR"`
}

func NewCreateHlrAuthentication() *CreateHlrAuthentication {
	createHlrAuthenticationType := new(CreateHlrAuthentication)
	return createHlrAuthenticationType
}

func NewCreateHlrAuthenticationWithDefaults() *CreateHlrAuthentication {
	createHlrAuthenticationType := new(CreateHlrAuthentication)
	return createHlrAuthenticationType
}

type CreateLDAPAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *common.NormalName2to128 `json:"adminDomainName" validate:"required"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *common.NormalName2to64 `json:"baseDomainName" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// KeyAttribute
	// Constraints:
	//    - required
	KeyAttribute *common.NormalName2to64 `json:"keyAttribute" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Password
	// Admin password
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// Port
	// Port
	// Constraints:
	//    - required
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// SearchFilter
	// Constraints:
	//    - required
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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	StandbySearchFilter *common.NormalName2to64 `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TlsEnabled
	// LDAP over TLS Enabled
	// Constraints:
	//    - required
	TlsEnabled *bool `json:"tlsEnabled" validate:"required"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=LDAP"`
}

func NewCreateLDAPAuthentication() *CreateLDAPAuthentication {
	createLDAPAuthenticationType := new(CreateLDAPAuthentication)
	return createLDAPAuthenticationType
}

func NewCreateLDAPAuthenticationWithDefaults() *CreateLDAPAuthentication {
	createLDAPAuthenticationType := new(CreateLDAPAuthentication)
	portField := 389
	createLDAPAuthenticationType.Port = &portField
	standbyPortField := 389
	createLDAPAuthenticationType.StandbyPort = &standbyPortField
	standbyServerEnabledField := false
	createLDAPAuthenticationType.StandbyServerEnabled = &standbyServerEnabledField
	return createLDAPAuthenticationType
}

type CreateRadiusAccounting struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *common.HealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Primary
	// Constraints:
	//    - required
	Primary *common.RadiusServer `json:"primary" validate:"required"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewCreateRadiusAccounting() *CreateRadiusAccounting {
	createRadiusAccountingType := new(CreateRadiusAccounting)
	return createRadiusAccountingType
}

func NewCreateRadiusAccountingWithDefaults() *CreateRadiusAccounting {
	createRadiusAccountingType := new(CreateRadiusAccounting)
	standbyServerEnabledField := false
	createRadiusAccountingType.StandbyServerEnabled = &standbyServerEnabledField
	return createRadiusAccountingType
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

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// Primary
	// Constraints:
	//    - required
	Primary *common.RadiusServer `json:"primary" validate:"required"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication server type
	// Constraints:
	//    - nullable
	//    - default:'RADIUS'
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewCreateRadiusAuthentication() *CreateRadiusAuthentication {
	createRadiusAuthenticationType := new(CreateRadiusAuthentication)
	return createRadiusAuthenticationType
}

func NewCreateRadiusAuthenticationWithDefaults() *CreateRadiusAuthentication {
	createRadiusAuthenticationType := new(CreateRadiusAuthentication)
	locationDeliveryEnabledField := false
	createRadiusAuthenticationType.LocationDeliveryEnabled = &locationDeliveryEnabledField
	standbyServerEnabledField := false
	createRadiusAuthenticationType.StandbyServerEnabled = &standbyServerEnabledField
	typeField := `RADIUS`
	createRadiusAuthenticationType.Type = &typeField
	return createRadiusAuthenticationType
}

type DeleteBulkAccountingService struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulkAccountingService() *DeleteBulkAccountingService {
	deleteBulkAccountingServiceType := new(DeleteBulkAccountingService)
	return deleteBulkAccountingServiceType
}

func NewDeleteBulkAccountingServiceWithDefaults() *DeleteBulkAccountingService {
	deleteBulkAccountingServiceType := new(DeleteBulkAccountingService)
	return deleteBulkAccountingServiceType
}

type DeleteBulkAuthenticationService struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulkAuthenticationService() *DeleteBulkAuthenticationService {
	deleteBulkAuthenticationServiceType := new(DeleteBulkAuthenticationService)
	return deleteBulkAuthenticationServiceType
}

func NewDeleteBulkAuthenticationServiceWithDefaults() *DeleteBulkAuthenticationService {
	deleteBulkAuthenticationServiceType := new(DeleteBulkAuthenticationService)
	return deleteBulkAuthenticationServiceType
}

type DnsServer struct {
	// Ip
	// IP of server
	Ip *string `json:"ip,omitempty"`
}

func NewDnsServer() *DnsServer {
	dnsServerType := new(DnsServer)
	return dnsServerType
}

func NewDnsServerWithDefaults() *DnsServer {
	dnsServerType := new(DnsServer)
	return dnsServerType
}

type DnsServerList []*DnsServer

func NewDnsServerList() *DnsServerList {
	dnsServerListType := make(DnsServerList, 0)
	return &dnsServerListType
}

func NewDnsServerListWithDefaults() *DnsServerList {
	dnsServerListType := make(DnsServerList, 0)
	return &dnsServerListType
}

type Ggsn struct {
	// DomainName
	// Domain name of GGSN
	DomainName *string `json:"domainName,omitempty"`

	// GgsnIPAddress
	// IP of GGSN
	GgsnIPAddress *string `json:"ggsnIPAddress,omitempty"`
}

func NewGgsn() *Ggsn {
	ggsnType := new(Ggsn)
	return ggsnType
}

func NewGgsnWithDefaults() *Ggsn {
	ggsnType := new(Ggsn)
	return ggsnType
}

type GgsnConfig struct {
	DnsServerList DnsServerList `json:"dnsServerList,omitempty"`

	GgsnList GgsnList `json:"ggsnList,omitempty"`

	GtpSettings *GtpSettings `json:"gtpSettings,omitempty"`
}

func NewGgsnConfig() *GgsnConfig {
	ggsnConfigType := new(GgsnConfig)
	return ggsnConfigType
}

func NewGgsnConfigWithDefaults() *GgsnConfig {
	ggsnConfigType := new(GgsnConfig)
	return ggsnConfigType
}

type GgsnList []*Ggsn

func NewGgsnList() *GgsnList {
	ggsnListType := make(GgsnList, 0)
	return &ggsnListType
}

func NewGgsnListWithDefaults() *GgsnList {
	ggsnListType := make(GgsnList, 0)
	return &ggsnListType
}

// GroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type GroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	// Constraints:
	//    - required
	GroupAttr *string `json:"groupAttr" validate:"required"`

	// Id
	// Group attribute mapping UUID
	Id *string `json:"id,omitempty"`

	// UserRole
	// Identity user role
	// Constraints:
	//    - required
	UserRole *GroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

func NewGroupAttrIdentityUserRoleMapping() *GroupAttrIdentityUserRoleMapping {
	groupAttrIdentityUserRoleMappingType := new(GroupAttrIdentityUserRoleMapping)
	return groupAttrIdentityUserRoleMappingType
}

func NewGroupAttrIdentityUserRoleMappingWithDefaults() *GroupAttrIdentityUserRoleMapping {
	groupAttrIdentityUserRoleMappingType := new(GroupAttrIdentityUserRoleMapping)
	return groupAttrIdentityUserRoleMappingType
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

func NewGroupAttrIdentityUserRoleMappingUserRoleType() *GroupAttrIdentityUserRoleMappingUserRoleType {
	groupAttrIdentityUserRoleMappingUserRoleTypeType := new(GroupAttrIdentityUserRoleMappingUserRoleType)
	return groupAttrIdentityUserRoleMappingUserRoleTypeType
}

func NewGroupAttrIdentityUserRoleMappingUserRoleTypeWithDefaults() *GroupAttrIdentityUserRoleMappingUserRoleType {
	groupAttrIdentityUserRoleMappingUserRoleTypeType := new(GroupAttrIdentityUserRoleMappingUserRoleType)
	return groupAttrIdentityUserRoleMappingUserRoleTypeType
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

func NewGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType() *GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType {
	groupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileTypeType := new(GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType)
	return groupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileTypeType
}

func NewGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileTypeWithDefaults() *GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType {
	groupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileTypeType := new(GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType)
	return groupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileTypeType
}

type GtpSettings struct {
	// DnsNumberOfRetries
	// DNS Number of Retries
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:10
	DnsNumberOfRetries *int `json:"dnsNumberOfRetries,omitempty" validate:"omitempty,gte=1,lte=10"`

	// EchoRequestTimer
	// Echo Request Timerr
	// Constraints:
	//    - nullable
	//    - min:60
	//    - max:300
	EchoRequestTimer *int `json:"echoRequestTimer,omitempty" validate:"omitempty,gte=60,lte=300"`

	// NumberOfRetries
	// Number of Retries
	// Constraints:
	//    - nullable
	//    - min:3
	//    - max:6
	NumberOfRetries *int `json:"numberOfRetries,omitempty" validate:"omitempty,gte=3,lte=6"`

	// ResponseTimeout
	// DNS Response Timeout
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:10
	ResponseTimeout *int `json:"responseTimeout,omitempty" validate:"omitempty,gte=1,lte=10"`

	// T3ResponseTimer
	// Response Timer
	// Constraints:
	//    - nullable
	//    - min:2
	//    - max:5
	T3ResponseTimer *int `json:"t3ResponseTimer,omitempty" validate:"omitempty,gte=2,lte=5"`
}

func NewGtpSettings() *GtpSettings {
	gtpSettingsType := new(GtpSettings)
	return gtpSettingsType
}

func NewGtpSettingsWithDefaults() *GtpSettings {
	gtpSettingsType := new(GtpSettings)
	return gtpSettingsType
}

type HlrService struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"omitempty,oneof=route_on_gt route_on_ssn"`

	// AuthMapVer
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	AuthMapVer *string `json:"authMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// DestNatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"omitempty,oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// DestNumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"omitempty,oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"omitempty,oneof=international international_spare national national_spare"`

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

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// SrcNatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"omitempty,oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// SrcNumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"omitempty,oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	// SrcTransType
	// - For HLR Authentiaction server
	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=HLR"`
}

func NewHlrService() *HlrService {
	hlrServiceType := new(HlrService)
	return hlrServiceType
}

func NewHlrServiceWithDefaults() *HlrService {
	hlrServiceType := new(HlrService)
	return hlrServiceType
}

type HlrServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*HlrService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewHlrServiceList() *HlrServiceList {
	hlrServiceListType := new(HlrServiceList)
	return hlrServiceListType
}

func NewHlrServiceListWithDefaults() *HlrServiceList {
	hlrServiceListType := new(HlrServiceList)
	return hlrServiceListType
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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// Protocol
	// Authentication protocol
	// Constraints:
	//    - nullable
	//    - oneof:[LDAP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=LDAP"`

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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=LDAP"`
}

func NewLDAPService() *LDAPService {
	lDAPServiceType := new(LDAPService)
	return lDAPServiceType
}

func NewLDAPServiceWithDefaults() *LDAPService {
	lDAPServiceType := new(LDAPService)
	portField := 389
	lDAPServiceType.Port = &portField
	standbyPortField := 389
	lDAPServiceType.StandbyPort = &standbyPortField
	return lDAPServiceType
}

type LDAPServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LDAPService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewLDAPServiceList() *LDAPServiceList {
	lDAPServiceListType := new(LDAPServiceList)
	return lDAPServiceListType
}

func NewLDAPServiceListWithDefaults() *LDAPServiceList {
	lDAPServiceListType := new(LDAPServiceList)
	return lDAPServiceListType
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

func NewMncNdc() *MncNdc {
	mncNdcType := new(MncNdc)
	return mncNdcType
}

func NewMncNdcWithDefaults() *MncNdc {
	mncNdcType := new(MncNdc)
	return mncNdcType
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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=1,lte=65535"`

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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AD"`

	WindowsDomainName *common.NormalName2to64 `json:"windowsDomainName,omitempty"`
}

func NewModifyActiveDirectoryAuthentication() *ModifyActiveDirectoryAuthentication {
	modifyActiveDirectoryAuthenticationType := new(ModifyActiveDirectoryAuthentication)
	return modifyActiveDirectoryAuthenticationType
}

func NewModifyActiveDirectoryAuthenticationWithDefaults() *ModifyActiveDirectoryAuthentication {
	modifyActiveDirectoryAuthenticationType := new(ModifyActiveDirectoryAuthentication)
	portField := 389
	modifyActiveDirectoryAuthenticationType.Port = &portField
	standbyPortField := 389
	modifyActiveDirectoryAuthenticationType.StandbyPort = &standbyPortField
	return modifyActiveDirectoryAuthenticationType
}

// ModifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type ModifyGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	// Constraints:
	//    - required
	GroupAttr *string `json:"groupAttr" validate:"required"`

	// UserRole
	// Identity user role
	// Constraints:
	//    - required
	UserRole *ModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

func NewModifyGroupAttrIdentityUserRoleMapping() *ModifyGroupAttrIdentityUserRoleMapping {
	modifyGroupAttrIdentityUserRoleMappingType := new(ModifyGroupAttrIdentityUserRoleMapping)
	return modifyGroupAttrIdentityUserRoleMappingType
}

func NewModifyGroupAttrIdentityUserRoleMappingWithDefaults() *ModifyGroupAttrIdentityUserRoleMapping {
	modifyGroupAttrIdentityUserRoleMappingType := new(ModifyGroupAttrIdentityUserRoleMapping)
	return modifyGroupAttrIdentityUserRoleMappingType
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

func NewModifyGroupAttrIdentityUserRoleMappingUserRoleType() *ModifyGroupAttrIdentityUserRoleMappingUserRoleType {
	modifyGroupAttrIdentityUserRoleMappingUserRoleTypeType := new(ModifyGroupAttrIdentityUserRoleMappingUserRoleType)
	return modifyGroupAttrIdentityUserRoleMappingUserRoleTypeType
}

func NewModifyGroupAttrIdentityUserRoleMappingUserRoleTypeWithDefaults() *ModifyGroupAttrIdentityUserRoleMappingUserRoleType {
	modifyGroupAttrIdentityUserRoleMappingUserRoleTypeType := new(ModifyGroupAttrIdentityUserRoleMappingUserRoleType)
	return modifyGroupAttrIdentityUserRoleMappingUserRoleTypeType
}

type ModifyHlrAuthentication struct {
	// AddressIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"omitempty,oneof=route_on_gt route_on_ssn"`

	// AuthMapVer
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	AuthMapVer *string `json:"authMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

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

	// DestGtIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// DestNatureOfAddressIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"omitempty,oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// DestNumberingPlan
	// Constraints:
	//    - nullable
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"omitempty,oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

	FriendlyName *common.NormalNameAllowBlank `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"omitempty,oneof=international international_spare national national_spare"`

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

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*SccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*SctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// SrcNatureOfAddressIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"omitempty,oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// SrcNumberingPlan
	// Constraints:
	//    - nullable
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"omitempty,oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=HLR"`
}

func NewModifyHlrAuthentication() *ModifyHlrAuthentication {
	modifyHlrAuthenticationType := new(ModifyHlrAuthentication)
	return modifyHlrAuthenticationType
}

func NewModifyHlrAuthenticationWithDefaults() *ModifyHlrAuthentication {
	modifyHlrAuthenticationType := new(ModifyHlrAuthentication)
	return modifyHlrAuthenticationType
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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=1,lte=65535"`

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
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=LDAP"`
}

func NewModifyLDAPAuthentication() *ModifyLDAPAuthentication {
	modifyLDAPAuthenticationType := new(ModifyLDAPAuthentication)
	return modifyLDAPAuthenticationType
}

func NewModifyLDAPAuthenticationWithDefaults() *ModifyLDAPAuthentication {
	modifyLDAPAuthenticationType := new(ModifyLDAPAuthentication)
	portField := 389
	modifyLDAPAuthenticationType.Port = &portField
	standbyPortField := 389
	modifyLDAPAuthenticationType.StandbyPort = &standbyPortField
	return modifyLDAPAuthenticationType
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
	// Constraints:
	//    - nullable
	//    - oneof:[LOCAL_DB]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=LOCAL_DB"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[LOCAL_DB]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=LOCAL_DB"`
}

func NewModifyLocalDbAuthentication() *ModifyLocalDbAuthentication {
	modifyLocalDbAuthenticationType := new(ModifyLocalDbAuthentication)
	return modifyLocalDbAuthenticationType
}

func NewModifyLocalDbAuthenticationWithDefaults() *ModifyLocalDbAuthentication {
	modifyLocalDbAuthenticationType := new(ModifyLocalDbAuthentication)
	return modifyLocalDbAuthenticationType
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
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewModifyRadiusAccounting() *ModifyRadiusAccounting {
	modifyRadiusAccountingType := new(ModifyRadiusAccounting)
	return modifyRadiusAccountingType
}

func NewModifyRadiusAccountingWithDefaults() *ModifyRadiusAccounting {
	modifyRadiusAccountingType := new(ModifyRadiusAccounting)
	return modifyRadiusAccountingType
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
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewModifyRadiusAuthentication() *ModifyRadiusAuthentication {
	modifyRadiusAuthenticationType := new(ModifyRadiusAuthentication)
	return modifyRadiusAuthenticationType
}

func NewModifyRadiusAuthenticationWithDefaults() *ModifyRadiusAuthentication {
	modifyRadiusAuthenticationType := new(ModifyRadiusAuthentication)
	locationDeliveryEnabledField := false
	modifyRadiusAuthenticationType.LocationDeliveryEnabled = &locationDeliveryEnabledField
	return modifyRadiusAuthenticationType
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
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewRadiusAccountingService() *RadiusAccountingService {
	radiusAccountingServiceType := new(RadiusAccountingService)
	return radiusAccountingServiceType
}

func NewRadiusAccountingServiceWithDefaults() *RadiusAccountingService {
	radiusAccountingServiceType := new(RadiusAccountingService)
	return radiusAccountingServiceType
}

type RadiusAccountingServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RadiusAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRadiusAccountingServiceList() *RadiusAccountingServiceList {
	radiusAccountingServiceListType := new(RadiusAccountingServiceList)
	return radiusAccountingServiceListType
}

func NewRadiusAccountingServiceListWithDefaults() *RadiusAccountingServiceList {
	radiusAccountingServiceListType := new(RadiusAccountingServiceList)
	return radiusAccountingServiceListType
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
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *common.RateLimiting `json:"rateLimiting,omitempty"`

	Secondary *SecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewRadiusAuthenticationService() *RadiusAuthenticationService {
	radiusAuthenticationServiceType := new(RadiusAuthenticationService)
	return radiusAuthenticationServiceType
}

func NewRadiusAuthenticationServiceWithDefaults() *RadiusAuthenticationService {
	radiusAuthenticationServiceType := new(RadiusAuthenticationService)
	locationDeliveryEnabledField := false
	radiusAuthenticationServiceType.LocationDeliveryEnabled = &locationDeliveryEnabledField
	return radiusAuthenticationServiceType
}

type RadiusAuthenticationServiceList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RadiusAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRadiusAuthenticationServiceList() *RadiusAuthenticationServiceList {
	radiusAuthenticationServiceListType := new(RadiusAuthenticationServiceList)
	return radiusAuthenticationServiceListType
}

func NewRadiusAuthenticationServiceListWithDefaults() *RadiusAuthenticationServiceList {
	radiusAuthenticationServiceListType := new(RadiusAuthenticationServiceList)
	return radiusAuthenticationServiceListType
}

type SccpGtt struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"omitempty,oneof=route_on_gt route_on_ssn"`

	// E164Address
	// - For HLR Authentiaction server
	E164Address *string `json:"e164Address,omitempty"`

	// GtDigits
	// GT digits
	GtDigits *string `json:"gtDigits,omitempty"`

	// GtIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	GtIndicator *string `json:"gtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// HasPointCode
	// Enable Point Code submitted
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSSN
	// Enable SSN- For HLR Authentiaction server
	HasSSN *bool `json:"hasSSN,omitempty"`

	// NatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	NatureOfAddressIndicator *string `json:"natureOfAddressIndicator,omitempty" validate:"omitempty,oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// NumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	NumberingPlan *string `json:"numberingPlan,omitempty" validate:"omitempty,oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	// PointCode
	// - For HLR Authentiaction server
	PointCode *int `json:"pointCode,omitempty"`

	// TransType
	// - For HLR Authentiaction server
	TransType *int `json:"transType,omitempty"`
}

func NewSccpGtt() *SccpGtt {
	sccpGttType := new(SccpGtt)
	return sccpGttType
}

func NewSccpGttWithDefaults() *SccpGtt {
	sccpGttType := new(SccpGtt)
	return sccpGttType
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
	// Constraints:
	//    - nullable
	//    - oneof:[active,backup,both]
	NodeTermination *string `json:"nodeTermination,omitempty" validate:"omitempty,oneof=active backup both"`

	// SourcePort
	// Source port
	SourcePort *int `json:"sourcePort,omitempty"`
}

func NewSctpAssociation() *SctpAssociation {
	sctpAssociationType := new(SctpAssociation)
	return sctpAssociationType
}

func NewSctpAssociationWithDefaults() *SctpAssociation {
	sctpAssociationType := new(SctpAssociation)
	return sctpAssociationType
}

type SecondaryRadiusServer struct {
	// AutoFallbackDisable
	// Automatic fallback enabled or disabled
	// Constraints:
	//    - required
	AutoFallbackDisable *bool `json:"autoFallbackDisable" validate:"required"`

	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// Port
	// RADIUS server port
	// Constraints:
	//    - required
	//    - default:1812
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// SharedSecret
	// RADIUS server shared secrect
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

func NewSecondaryRadiusServer() *SecondaryRadiusServer {
	secondaryRadiusServerType := new(SecondaryRadiusServer)
	return secondaryRadiusServerType
}

func NewSecondaryRadiusServerWithDefaults() *SecondaryRadiusServer {
	secondaryRadiusServerType := new(SecondaryRadiusServer)
	portField := 1812
	secondaryRadiusServerType.Port = &portField
	return secondaryRadiusServerType
}

type TestingConfig struct {
	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	LoginRequest *TestingConfigLoginRequestType `json:"loginRequest,omitempty"`
}

func NewTestingConfig() *TestingConfig {
	testingConfigType := new(TestingConfig)
	return testingConfigType
}

func NewTestingConfigWithDefaults() *TestingConfig {
	testingConfigType := new(TestingConfig)
	return testingConfigType
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

func NewTestingConfigLoginRequestType() *TestingConfigLoginRequestType {
	testingConfigLoginRequestTypeType := new(TestingConfigLoginRequestType)
	return testingConfigLoginRequestTypeType
}

func NewTestingConfigLoginRequestTypeWithDefaults() *TestingConfigLoginRequestType {
	testingConfigLoginRequestTypeType := new(TestingConfigLoginRequestType)
	return testingConfigLoginRequestTypeType
}
