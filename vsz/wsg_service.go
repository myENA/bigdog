package vsz

// API Version: v8_1

type WSGServiceActiveDirectoryService struct {
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the active directory authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

	StandbyAdminDomainName *WSGCommonNormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

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

	StandbyWindowsDomainName *WSGCommonNormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AD"`

	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

type WSGServiceActiveDirectoryServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceActiveDirectoryService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceCommonAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

type WSGServiceCommonAccountingServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceCommonAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

type WSGServiceCommonAuthenticationServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceCreateActiveDirectoryAuthentication struct {
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

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
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

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

	StandbyAdminDomainName *WSGCommonNormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

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

	StandbyWindowsDomainName *WSGCommonNormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

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
	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName" validate:"required"`
}

type WSGServiceCreateHlrAuthentication struct {
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

	Description *WSGCommonDescription `json:"description,omitempty"`

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

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

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

	MncNdcList []*WSGServiceMncNdc `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty"`

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

type WSGServiceCreateLDAPAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName" validate:"required"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName" validate:"required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// KeyAttribute
	// Constraints:
	//    - required
	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

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
	SearchFilter *WSGCommonNormalName2to64 `json:"searchFilter" validate:"required"`

	StandbyAdminDomainName *WSGCommonNormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *WSGCommonNormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *WSGCommonNormalName2to64 `json:"standbyKeyAttribute,omitempty"`

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

	StandbySearchFilter *WSGCommonNormalName2to64 `json:"standbySearchFilter,omitempty"`

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

type WSGServiceCreateRadiusAccounting struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary" validate:"required"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

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

type WSGServiceCreateRadiusAuthentication struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary" validate:"required"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

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

type WSGServiceDeleteBulkAccountingService struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

type WSGServiceDeleteBulkAuthenticationService struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

type WSGServiceDnsServer struct {
	// Ip
	// IP of server
	Ip *string `json:"ip,omitempty"`
}

type WSGServiceDnsServerList []*WSGServiceDnsServer

type WSGServiceGgsn struct {
	// DomainName
	// Domain name of GGSN
	DomainName *string `json:"domainName,omitempty"`

	// GgsnIPAddress
	// IP of GGSN
	GgsnIPAddress *string `json:"ggsnIPAddress,omitempty"`
}

type WSGServiceGgsnConfig struct {
	DnsServerList WSGServiceDnsServerList `json:"dnsServerList,omitempty"`

	GgsnList WSGServiceGgsnList `json:"ggsnList,omitempty"`

	GtpSettings *WSGServiceGtpSettings `json:"gtpSettings,omitempty"`
}

type WSGServiceGgsnList []*WSGServiceGgsn

// WSGServiceGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type WSGServiceGroupAttrIdentityUserRoleMapping struct {
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
	UserRole *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

// WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// Id
	// Identity user role UUID
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	UserTrafficProfile *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

// WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Identity user role
type WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	// Id
	// User traffic profile UUID
	Id *string `json:"id,omitempty"`

	// Name
	// User traffic profile name
	Name *string `json:"name,omitempty"`
}

type WSGServiceGtpSettings struct {
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

type WSGServiceHlrService struct {
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

	Description *WSGCommonDescription `json:"description,omitempty"`

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

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

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
	MncNdcList []*WSGServiceMncNdc `json:"mncNdcList,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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
	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty"`

	// SctpAssociationsList
	// - For HLR Authentiaction server
	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty"`

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

type WSGServiceHlrServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceHlrService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceLDAPService struct {
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName,omitempty"`

	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the LDAP authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

	SearchFilter *WSGCommonNormalName2to64 `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *WSGCommonNormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *WSGCommonNormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *WSGCommonNormalName2to64 `json:"standbyKeyAttribute,omitempty"`

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

	StandbySearchFilter *WSGCommonNormalName2to64 `json:"standbySearchFilter,omitempty"`

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

type WSGServiceLDAPServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceLDAPService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceMncNdc struct {
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

type WSGServiceModifyActiveDirectoryAuthentication struct {
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

	StandbyAdminDomainName *WSGCommonNormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

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

	StandbyWindowsDomainName *WSGCommonNormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AD"`

	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

// WSGServiceModifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type WSGServiceModifyGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	// Constraints:
	//    - required
	GroupAttr *string `json:"groupAttr" validate:"required"`

	// UserRole
	// Identity user role
	// Constraints:
	//    - required
	UserRole *WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

// WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// Id
	// Identity user role UUID
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`
}

type WSGServiceModifyHlrAuthentication struct {
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

	Description *WSGCommonDescription `json:"description,omitempty"`

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

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

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

	MncNdcList []*WSGServiceMncNdc `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty"`

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

type WSGServiceModifyLDAPAuthentication struct {
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName,omitempty"`

	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

	SearchFilter *WSGCommonNormalName2to64 `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *WSGCommonNormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *WSGCommonNormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *WSGCommonNormalName2to64 `json:"standbyKeyAttribute,omitempty"`

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

	StandbySearchFilter *WSGCommonNormalName2to64 `json:"standbySearchFilter,omitempty"`

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

type WSGServiceModifyLocalDbAuthentication struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

type WSGServiceModifyRadiusAccounting struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS accounting service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

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

type WSGServiceModifyRadiusAuthentication struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

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

type WSGServiceRadiusAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

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

type WSGServiceRadiusAccountingServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceRadiusAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

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

type WSGServiceRadiusAuthenticationServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceSccpGtt struct {
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

type WSGServiceSctpAssociation struct {
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

type WSGServiceSecondaryRadiusServer struct {
	// AutoFallbackDisable
	// Automatic fallback enabled or disabled
	// Constraints:
	//    - required
	AutoFallbackDisable *bool `json:"autoFallbackDisable" validate:"required"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

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

type WSGServiceTestingConfig struct {
	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	LoginRequest *WSGServiceTestingConfigLoginRequestType `json:"loginRequest,omitempty"`
}

type WSGServiceTestingConfigLoginRequestType struct {
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
