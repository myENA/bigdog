package vsz

// API Version: v9_0

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
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - oneof:[AD]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=AD"`

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
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

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
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"oneof=AD"`

	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

func NewWSGServiceActiveDirectoryService() *WSGServiceActiveDirectoryService {
	m := new(WSGServiceActiveDirectoryService)
	return m
}

type WSGServiceActiveDirectoryServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceActiveDirectoryService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceActiveDirectoryServiceList() *WSGServiceActiveDirectoryServiceList {
	m := new(WSGServiceActiveDirectoryServiceList)
	return m
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
	//    - oneof:[RADIUS,CGF]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS CGF"`

	// Type
	// Accounting protocol same as protocol.
	// Constraints:
	//    - oneof:[RADIUS,CGF]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS CGF"`
}

func NewWSGServiceCommonAccountingService() *WSGServiceCommonAccountingService {
	m := new(WSGServiceCommonAccountingService)
	return m
}

type WSGServiceCommonAccountingServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceCommonAccountingServiceList() *WSGServiceCommonAccountingServiceList {
	m := new(WSGServiceCommonAccountingServiceList)
	return m
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
	//    - oneof:[RADIUS,AD,LDAP,FACEBOOK,LINKEDIN,GOOGLE,GENERICOAUTH,SOAP,HLR,LOCAL_DB,GUEST]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS AD LDAP FACEBOOK LINKEDIN GOOGLE GENERICOAUTH SOAP HLR LOCAL_DB GUEST"`

	// Type
	// Authentication protocol same as protocol.
	// Constraints:
	//    - oneof:[RADIUS,AD,LDAP,FACEBOOK,LINKEDIN,GOOGLE,GENERICOAUTH,SOAP,HLR,LOCAL_DB,GUEST]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS AD LDAP FACEBOOK LINKEDIN GOOGLE GENERICOAUTH SOAP HLR LOCAL_DB GUEST"`
}

func NewWSGServiceCommonAuthenticationService() *WSGServiceCommonAuthenticationService {
	m := new(WSGServiceCommonAuthenticationService)
	return m
}

type WSGServiceCommonAuthenticationServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceCommonAuthenticationServiceList() *WSGServiceCommonAuthenticationServiceList {
	m := new(WSGServiceCommonAuthenticationServiceList)
	return m
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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

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
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"oneof=AD"`

	// WindowsDomainName
	// Constraints:
	//    - required
	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName" validate:"required,max=64,min=2"`
}

func NewWSGServiceCreateActiveDirectoryAuthentication() *WSGServiceCreateActiveDirectoryAuthentication {
	m := new(WSGServiceCreateActiveDirectoryAuthentication)
	return m
}

type WSGServiceCreateHlrAuthentication struct {
	// AddressIndicator
	// Constraints:
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	// AuthMapVer
	// Constraints:
	//    - oneof:[version2,version3]
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

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// DestNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// DestNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"oneof=version2 version3"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"oneof=international international_spare national national_spare"`

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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// SrcNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// SrcNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"oneof=HLR"`
}

func NewWSGServiceCreateHlrAuthentication() *WSGServiceCreateHlrAuthentication {
	m := new(WSGServiceCreateHlrAuthentication)
	return m
}

type WSGServiceCreateLDAPAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName" validate:"required,max=128,min=2"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName" validate:"required,max=64,min=2"`

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
	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute" validate:"required,max=64,min=2"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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
	SearchFilter *WSGCommonNormalName2to64 `json:"searchFilter" validate:"required,max=64,min=2"`

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
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

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
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"oneof=LDAP"`
}

func NewWSGServiceCreateLDAPAuthentication() *WSGServiceCreateLDAPAuthentication {
	m := new(WSGServiceCreateLDAPAuthentication)
	return m
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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary" validate:"required"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

func NewWSGServiceCreateRadiusAccounting() *WSGServiceCreateRadiusAccounting {
	m := new(WSGServiceCreateRadiusAccounting)
	return m
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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

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
	// Authentication protocol
	// Constraints:
	//    - default:'RADIUS'
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

func NewWSGServiceCreateRadiusAuthentication() *WSGServiceCreateRadiusAuthentication {
	m := new(WSGServiceCreateRadiusAuthentication)
	return m
}

type WSGServiceDeleteBulkAccountingService struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGServiceDeleteBulkAccountingService() *WSGServiceDeleteBulkAccountingService {
	m := new(WSGServiceDeleteBulkAccountingService)
	return m
}

type WSGServiceDeleteBulkAuthenticationService struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGServiceDeleteBulkAuthenticationService() *WSGServiceDeleteBulkAuthenticationService {
	m := new(WSGServiceDeleteBulkAuthenticationService)
	return m
}

type WSGServiceDnsServer struct {
	// Ip
	// IP of server
	Ip *string `json:"ip,omitempty"`
}

func NewWSGServiceDnsServer() *WSGServiceDnsServer {
	m := new(WSGServiceDnsServer)
	return m
}

type WSGServiceDnsServerList []*WSGServiceDnsServer

func MakeWSGServiceDnsServerList() WSGServiceDnsServerList {
	m := make(WSGServiceDnsServerList, 0)
	return m
}

type WSGServiceGgsn struct {
	// DomainName
	// Domain name of GGSN
	DomainName *string `json:"domainName,omitempty"`

	// GgsnIPAddress
	// IP of GGSN
	GgsnIPAddress *string `json:"ggsnIPAddress,omitempty"`
}

func NewWSGServiceGgsn() *WSGServiceGgsn {
	m := new(WSGServiceGgsn)
	return m
}

type WSGServiceGgsnConfig struct {
	DnsServerList WSGServiceDnsServerList `json:"dnsServerList,omitempty"`

	GgsnList WSGServiceGgsnList `json:"ggsnList,omitempty"`

	GtpSettings *WSGServiceGtpSettings `json:"gtpSettings,omitempty"`
}

func NewWSGServiceGgsnConfig() *WSGServiceGgsnConfig {
	m := new(WSGServiceGgsnConfig)
	return m
}

type WSGServiceGgsnList []*WSGServiceGgsn

func MakeWSGServiceGgsnList() WSGServiceGgsnList {
	m := make(WSGServiceGgsnList, 0)
	return m
}

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

func NewWSGServiceGroupAttrIdentityUserRoleMapping() *WSGServiceGroupAttrIdentityUserRoleMapping {
	m := new(WSGServiceGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType struct {
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// Id
	// Identity user role UUID
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	UserTrafficProfile *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

func NewWSGServiceGroupAttrIdentityUserRoleMappingUserRoleType() *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
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

func NewWSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType() *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType {
	m := new(WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType)
	return m
}

type WSGServiceGtpSettings struct {
	// DnsNumberOfRetries
	// DNS Number of Retries
	// Constraints:
	//    - min:1
	//    - max:10
	DnsNumberOfRetries *int `json:"dnsNumberOfRetries,omitempty" validate:"gte=1,lte=10"`

	// EchoRequestTimer
	// Echo Request Timerr
	// Constraints:
	//    - min:60
	//    - max:300
	EchoRequestTimer *int `json:"echoRequestTimer,omitempty" validate:"gte=60,lte=300"`

	// NumberOfRetries
	// Number of Retries
	// Constraints:
	//    - min:3
	//    - max:6
	NumberOfRetries *int `json:"numberOfRetries,omitempty" validate:"gte=3,lte=6"`

	// ResponseTimeout
	// DNS Response Timeout
	// Constraints:
	//    - min:1
	//    - max:10
	ResponseTimeout *int `json:"responseTimeout,omitempty" validate:"gte=1,lte=10"`

	// T3ResponseTimer
	// Response Timer
	// Constraints:
	//    - min:2
	//    - max:5
	T3ResponseTimer *int `json:"t3ResponseTimer,omitempty" validate:"gte=2,lte=5"`
}

func NewWSGServiceGtpSettings() *WSGServiceGtpSettings {
	m := new(WSGServiceGtpSettings)
	return m
}

type WSGServiceHlrService struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	// AuthMapVer
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[version2,version3]
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

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestGtIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// DestNatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// DestNumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
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
	// Constraints:
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"oneof=version2 version3"`

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
	//    - oneof:[international,international_spare,national,national_spare]
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
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=HLR"`

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
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// SrcNatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// SrcNumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	// SrcTransType
	// - For HLR Authentiaction server
	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"oneof=HLR"`
}

func NewWSGServiceHlrService() *WSGServiceHlrService {
	m := new(WSGServiceHlrService)
	return m
}

type WSGServiceHlrServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceHlrService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceHlrServiceList() *WSGServiceHlrServiceList {
	m := new(WSGServiceHlrServiceList)
	return m
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
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

	// Protocol
	// Authentication protocol
	// Constraints:
	//    - oneof:[LDAP]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=LDAP"`

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
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

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
	// Authentication protocol same as protocol.
	// Constraints:
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"oneof=LDAP"`
}

func NewWSGServiceLDAPService() *WSGServiceLDAPService {
	m := new(WSGServiceLDAPService)
	return m
}

type WSGServiceLDAPServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceLDAPService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceLDAPServiceList() *WSGServiceLDAPServiceList {
	m := new(WSGServiceLDAPServiceList)
	return m
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

func NewWSGServiceMncNdc() *WSGServiceMncNdc {
	m := new(WSGServiceMncNdc)
	return m
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
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

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
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

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
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"oneof=AD"`

	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

func NewWSGServiceModifyActiveDirectoryAuthentication() *WSGServiceModifyActiveDirectoryAuthentication {
	m := new(WSGServiceModifyActiveDirectoryAuthentication)
	return m
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

func NewWSGServiceModifyGroupAttrIdentityUserRoleMapping() *WSGServiceModifyGroupAttrIdentityUserRoleMapping {
	m := new(WSGServiceModifyGroupAttrIdentityUserRoleMapping)
	return m
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

func NewWSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType() *WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
}

type WSGServiceModifyHlrAuthentication struct {
	// AddressIndicator
	// Constraints:
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	// AuthMapVer
	// Constraints:
	//    - oneof:[version2,version3]
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

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DestGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// DestNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// DestNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"oneof=version2 version3"`

	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"oneof=international international_spare national national_spare"`

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
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=HLR"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// SrcNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// SrcNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"oneof=HLR"`
}

func NewWSGServiceModifyHlrAuthentication() *WSGServiceModifyHlrAuthentication {
	m := new(WSGServiceModifyHlrAuthentication)
	return m
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
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"gte=1,lte=65535"`

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
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"gte=1,lte=65535"`

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
	// Authentication protocol same as protocol.
	// Constraints:
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"oneof=LDAP"`
}

func NewWSGServiceModifyLDAPAuthentication() *WSGServiceModifyLDAPAuthentication {
	m := new(WSGServiceModifyLDAPAuthentication)
	return m
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
	//    - oneof:[LOCAL_DB]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=LOCAL_DB"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[LOCAL_DB]
	Type *string `json:"type,omitempty" validate:"oneof=LOCAL_DB"`
}

func NewWSGServiceModifyLocalDbAuthentication() *WSGServiceModifyLocalDbAuthentication {
	m := new(WSGServiceModifyLocalDbAuthentication)
	return m
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
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

func NewWSGServiceModifyRadiusAccounting() *WSGServiceModifyRadiusAccounting {
	m := new(WSGServiceModifyRadiusAccounting)
	return m
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
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

func NewWSGServiceModifyRadiusAuthentication() *WSGServiceModifyRadiusAuthentication {
	m := new(WSGServiceModifyRadiusAuthentication)
	return m
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
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

func NewWSGServiceRadiusAccountingService() *WSGServiceRadiusAccountingService {
	m := new(WSGServiceRadiusAccountingService)
	return m
}

type WSGServiceRadiusAccountingServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceRadiusAccountingServiceList() *WSGServiceRadiusAccountingServiceList {
	m := new(WSGServiceRadiusAccountingServiceList)
	return m
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
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=RADIUS"`

	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS"`
}

func NewWSGServiceRadiusAuthenticationService() *WSGServiceRadiusAuthenticationService {
	m := new(WSGServiceRadiusAuthenticationService)
	return m
}

type WSGServiceRadiusAuthenticationServiceList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceRadiusAuthenticationServiceList() *WSGServiceRadiusAuthenticationServiceList {
	m := new(WSGServiceRadiusAuthenticationServiceList)
	return m
}

type WSGServiceSccpGtt struct {
	// AddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"oneof=route_on_gt route_on_ssn"`

	// E164Address
	// - For HLR Authentiaction server
	E164Address *string `json:"e164Address,omitempty"`

	// GtDigits
	// GT digits
	GtDigits *string `json:"gtDigits,omitempty"`

	// GtIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	GtIndicator *string `json:"gtIndicator,omitempty" validate:"oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// HasPointCode
	// Enable Point Code submitted
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSSN
	// Enable SSN- For HLR Authentiaction server
	HasSSN *bool `json:"hasSSN,omitempty"`

	// NatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	NatureOfAddressIndicator *string `json:"natureOfAddressIndicator,omitempty" validate:"oneof=unknown subscriber_number reserved_for_national_use national_significant_number international_number"`

	// NumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	NumberingPlan *string `json:"numberingPlan,omitempty" validate:"oneof=isdn_telephony_numbering_plan land_mobile_numbering_plan isdn_mobile_numbering_plan"`

	// PointCode
	// - For HLR Authentiaction server
	PointCode *int `json:"pointCode,omitempty"`

	// TransType
	// - For HLR Authentiaction server
	TransType *int `json:"transType,omitempty"`
}

func NewWSGServiceSccpGtt() *WSGServiceSccpGtt {
	m := new(WSGServiceSccpGtt)
	return m
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
	//    - oneof:[active,backup,both]
	NodeTermination *string `json:"nodeTermination,omitempty" validate:"oneof=active backup both"`

	// SourcePort
	// Source port
	SourcePort *int `json:"sourcePort,omitempty"`
}

func NewWSGServiceSctpAssociation() *WSGServiceSctpAssociation {
	m := new(WSGServiceSctpAssociation)
	return m
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

func NewWSGServiceSecondaryRadiusServer() *WSGServiceSecondaryRadiusServer {
	m := new(WSGServiceSecondaryRadiusServer)
	return m
}

type WSGServiceTestingConfig struct {
	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	LoginRequest *WSGServiceTestingConfigLoginRequestType `json:"loginRequest,omitempty"`
}

func NewWSGServiceTestingConfig() *WSGServiceTestingConfig {
	m := new(WSGServiceTestingConfig)
	return m
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

func NewWSGServiceTestingConfigLoginRequestType() *WSGServiceTestingConfigLoginRequestType {
	m := new(WSGServiceTestingConfigLoginRequestType)
	return m
}
