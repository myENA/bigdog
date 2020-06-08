package ruckus

// API Version: v9_0

type WSGServiceActiveDirectoryService struct {
	AdminDomainName *WSGServiceActiveDirectoryService `json:"adminDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGServiceActiveDirectoryService `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceActiveDirectoryService `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the active directory authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGServiceActiveDirectoryService `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceActiveDirectoryService `json:"mappings,omitempty"`

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

	Name *WSGServiceActiveDirectoryService `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - oneof:[AD]
	Protocol *string `json:"protocol,omitempty"`

	StandbyAdminDomainName *WSGServiceActiveDirectoryService `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *WSGServiceActiveDirectoryService `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port for standby cluster
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled for standby cluster
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *WSGServiceActiveDirectoryService `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[AD]
	Type *string `json:"type,omitempty"`

	WindowsDomainName *WSGServiceActiveDirectoryService `json:"windowsDomainName,omitempty"`
}

func NewWSGServiceActiveDirectoryService() *WSGServiceActiveDirectoryService {
	m := new(WSGServiceActiveDirectoryService)
	return m
}

type WSGServiceActiveDirectoryServiceList struct {
	Extra *WSGServiceActiveDirectoryServiceList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceActiveDirectoryServiceList `json:"list,omitempty"`

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

	Description *WSGServiceCommonAccountingService `json:"description,omitempty"`

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

	Name *WSGServiceCommonAccountingService `json:"name,omitempty"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS,CGF]
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Accounting protocol same as protocol.
	// Constraints:
	//    - oneof:[RADIUS,CGF]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCommonAccountingService() *WSGServiceCommonAccountingService {
	m := new(WSGServiceCommonAccountingService)
	return m
}

type WSGServiceCommonAccountingServiceList struct {
	Extra *WSGServiceCommonAccountingServiceList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAccountingServiceList `json:"list,omitempty"`

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

	Description *WSGServiceCommonAuthenticationService `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceCommonAuthenticationService `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceCommonAuthenticationService `json:"mappings,omitempty"`

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

	Name *WSGServiceCommonAuthenticationService `json:"name,omitempty"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - oneof:[RADIUS,AD,LDAP,FACEBOOK,LINKEDIN,GOOGLE,GENERICOAUTH,SOAP,HLR,LOCAL_DB,GUEST]
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Authentication protocol same as protocol.
	// Constraints:
	//    - oneof:[RADIUS,AD,LDAP,FACEBOOK,LINKEDIN,GOOGLE,GENERICOAUTH,SOAP,HLR,LOCAL_DB,GUEST]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCommonAuthenticationService() *WSGServiceCommonAuthenticationService {
	m := new(WSGServiceCommonAuthenticationService)
	return m
}

type WSGServiceCommonAuthenticationServiceList struct {
	Extra *WSGServiceCommonAuthenticationServiceList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAuthenticationServiceList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceCommonAuthenticationServiceList() *WSGServiceCommonAuthenticationServiceList {
	m := new(WSGServiceCommonAuthenticationServiceList)
	return m
}

type WSGServiceCreateActiveDirectoryAuthentication struct {
	AdminDomainName *WSGServiceCreateActiveDirectoryAuthentication `json:"adminDomainName,omitempty"`

	Description *WSGServiceCreateActiveDirectoryAuthentication `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceCreateActiveDirectoryAuthentication `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	// Constraints:
	//    - required
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGServiceCreateActiveDirectoryAuthentication `json:"ip"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceCreateActiveDirectoryAuthentication `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGServiceCreateActiveDirectoryAuthentication `json:"name"`

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
	Port *int `json:"port"`

	StandbyAdminDomainName *WSGServiceCreateActiveDirectoryAuthentication `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *WSGServiceCreateActiveDirectoryAuthentication `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port for standby cluster
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled for standby cluster
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *WSGServiceCreateActiveDirectoryAuthentication `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	// Constraints:
	//    - required
	TlsEnabled *bool `json:"tlsEnabled"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[AD]
	Type *string `json:"type,omitempty"`

	// WindowsDomainName
	// Constraints:
	//    - required
	WindowsDomainName *WSGServiceCreateActiveDirectoryAuthentication `json:"windowsDomainName"`
}

func NewWSGServiceCreateActiveDirectoryAuthentication() *WSGServiceCreateActiveDirectoryAuthentication {
	m := new(WSGServiceCreateActiveDirectoryAuthentication)
	return m
}

type WSGServiceCreateHlrAuthentication struct {
	// AddressIndicator
	// Constraints:
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty"`

	// AuthMapVer
	// Constraints:
	//    - oneof:[version2,version3]
	AuthMapVer *string `json:"authMapVer,omitempty"`

	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *WSGServiceCreateHlrAuthentication `json:"description,omitempty"`

	// DestGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty"`

	// DestNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty"`

	// DestNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty"`

	FriendlyName *WSGServiceCreateHlrAuthentication `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty"`

	LocalPointCode *int `json:"localPointCode,omitempty"`

	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	MncNdcList []*WSGServiceCreateHlrAuthentication `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGServiceCreateHlrAuthentication `json:"name"`

	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceCreateHlrAuthentication `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceCreateHlrAuthentication `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty"`

	// SrcNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty"`

	// SrcNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCreateHlrAuthentication() *WSGServiceCreateHlrAuthentication {
	m := new(WSGServiceCreateHlrAuthentication)
	return m
}

type WSGServiceCreateLDAPAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *WSGServiceCreateLDAPAuthentication `json:"adminDomainName"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *WSGServiceCreateLDAPAuthentication `json:"baseDomainName"`

	Description *WSGServiceCreateLDAPAuthentication `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceCreateLDAPAuthentication `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGServiceCreateLDAPAuthentication `json:"ip"`

	// KeyAttribute
	// Constraints:
	//    - required
	KeyAttribute *WSGServiceCreateLDAPAuthentication `json:"keyAttribute"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceCreateLDAPAuthentication `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGServiceCreateLDAPAuthentication `json:"name"`

	// Password
	// Admin password
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// Port
	// Port
	// Constraints:
	//    - required
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// SearchFilter
	// Constraints:
	//    - required
	SearchFilter *WSGServiceCreateLDAPAuthentication `json:"searchFilter"`

	StandbyAdminDomainName *WSGServiceCreateLDAPAuthentication `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *WSGServiceCreateLDAPAuthentication `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *WSGServiceCreateLDAPAuthentication `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *WSGServiceCreateLDAPAuthentication `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *WSGServiceCreateLDAPAuthentication `json:"standbySearchFilter,omitempty"`

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
	TlsEnabled *bool `json:"tlsEnabled"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCreateLDAPAuthentication() *WSGServiceCreateLDAPAuthentication {
	m := new(WSGServiceCreateLDAPAuthentication)
	return m
}

type WSGServiceCreateRadiusAccounting struct {
	Description *WSGServiceCreateRadiusAccounting `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *WSGServiceCreateRadiusAccounting `json:"healthCheckPolicy,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGServiceCreateRadiusAccounting `json:"name"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGServiceCreateRadiusAccounting `json:"primary"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *WSGServiceCreateRadiusAccounting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceCreateRadiusAccounting `json:"secondary,omitempty"`

	StandbyPrimary *WSGServiceCreateRadiusAccounting `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCreateRadiusAccounting() *WSGServiceCreateRadiusAccounting {
	m := new(WSGServiceCreateRadiusAccounting)
	return m
}

type WSGServiceCreateRadiusAuthentication struct {
	Description *WSGServiceCreateRadiusAuthentication `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceCreateRadiusAuthentication `json:"friendlyName,omitempty"`

	HealthCheckPolicy *WSGServiceCreateRadiusAuthentication `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceCreateRadiusAuthentication `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGServiceCreateRadiusAuthentication `json:"name"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGServiceCreateRadiusAuthentication `json:"primary"`

	RateLimiting *WSGServiceCreateRadiusAuthentication `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceCreateRadiusAuthentication `json:"secondary,omitempty"`

	StandbyPrimary *WSGServiceCreateRadiusAuthentication `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol
	// Constraints:
	//    - default:'RADIUS'
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCreateRadiusAuthentication() *WSGServiceCreateRadiusAuthentication {
	m := new(WSGServiceCreateRadiusAuthentication)
	return m
}

type WSGServiceDeleteBulkAccountingService struct {
	IdList *WSGServiceDeleteBulkAccountingService `json:"idList,omitempty"`
}

func NewWSGServiceDeleteBulkAccountingService() *WSGServiceDeleteBulkAccountingService {
	m := new(WSGServiceDeleteBulkAccountingService)
	return m
}

type WSGServiceDeleteBulkAuthenticationService struct {
	IdList *WSGServiceDeleteBulkAuthenticationService `json:"idList,omitempty"`
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
	DnsServerList *WSGServiceGgsnConfig `json:"dnsServerList,omitempty"`

	GgsnList *WSGServiceGgsnConfig `json:"ggsnList,omitempty"`

	GtpSettings *WSGServiceGgsnConfig `json:"gtpSettings,omitempty"`
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
	GroupAttr *string `json:"groupAttr"`

	// Id
	// Group attribute mapping UUID
	Id *string `json:"id,omitempty"`

	// UserRole
	// Identity user role
	// Constraints:
	//    - required
	UserRole *WSGServiceGroupAttrIdentityUserRoleMapping `json:"userRole"`
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

	Name *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	UserTrafficProfile *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType `json:"userTrafficProfile,omitempty"`
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
	DnsNumberOfRetries *int `json:"dnsNumberOfRetries,omitempty"`

	// EchoRequestTimer
	// Echo Request Timerr
	// Constraints:
	//    - min:60
	//    - max:300
	EchoRequestTimer *int `json:"echoRequestTimer,omitempty"`

	// NumberOfRetries
	// Number of Retries
	// Constraints:
	//    - min:3
	//    - max:6
	NumberOfRetries *int `json:"numberOfRetries,omitempty"`

	// ResponseTimeout
	// DNS Response Timeout
	// Constraints:
	//    - min:1
	//    - max:10
	ResponseTimeout *int `json:"responseTimeout,omitempty"`

	// T3ResponseTimer
	// Response Timer
	// Constraints:
	//    - min:2
	//    - max:5
	T3ResponseTimer *int `json:"t3ResponseTimer,omitempty"`
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
	AddressIndicator *string `json:"addressIndicator,omitempty"`

	// AuthMapVer
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[version2,version3]
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

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultPointCodeFormat
	// - For HLR Authentiaction server
	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *WSGServiceHlrService `json:"description,omitempty"`

	// DestGtIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty"`

	// DestNatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty"`

	// DestNumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty"`

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
	EapSimMapVer *string `json:"eapSimMapVer,omitempty"`

	FriendlyName *WSGServiceHlrService `json:"friendlyName,omitempty"`

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
	MncNdcList []*WSGServiceHlrService `json:"mncNdcList,omitempty"`

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

	Name *WSGServiceHlrService `json:"name,omitempty"`

	// PointCode
	// - For HLR Authentiaction server
	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty"`

	// ReuseEnable
	// - For HLR Authentiaction server
	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	// RoutingContext
	// - For HLR Authentiaction server
	RoutingContext *int `json:"routingContext,omitempty"`

	// SccpGttList
	// - For HLR Authentiaction server
	SccpGttList []*WSGServiceHlrService `json:"sccpGttList,omitempty"`

	// SctpAssociationsList
	// - For HLR Authentiaction server
	SctpAssociationsList []*WSGServiceHlrService `json:"sctpAssociationsList,omitempty"`

	// SgsnIsdnAddress
	// - For HLR Authentiaction server
	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty"`

	// SrcNatureOfAddressIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty"`

	// SrcNumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty"`

	// SrcTransType
	// - For HLR Authentiaction server
	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceHlrService() *WSGServiceHlrService {
	m := new(WSGServiceHlrService)
	return m
}

type WSGServiceHlrServiceList struct {
	Extra *WSGServiceHlrServiceList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceHlrServiceList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceHlrServiceList() *WSGServiceHlrServiceList {
	m := new(WSGServiceHlrServiceList)
	return m
}

type WSGServiceLDAPService struct {
	AdminDomainName *WSGServiceLDAPService `json:"adminDomainName,omitempty"`

	BaseDomainName *WSGServiceLDAPService `json:"baseDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGServiceLDAPService `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceLDAPService `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the LDAP authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGServiceLDAPService `json:"ip,omitempty"`

	KeyAttribute *WSGServiceLDAPService `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceLDAPService `json:"mappings,omitempty"`

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

	Name *WSGServiceLDAPService `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// Protocol
	// Authentication protocol
	// Constraints:
	//    - oneof:[LDAP]
	Protocol *string `json:"protocol,omitempty"`

	SearchFilter *WSGServiceLDAPService `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *WSGServiceLDAPService `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *WSGServiceLDAPService `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *WSGServiceLDAPService `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *WSGServiceLDAPService `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *WSGServiceLDAPService `json:"standbySearchFilter,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceLDAPService() *WSGServiceLDAPService {
	m := new(WSGServiceLDAPService)
	return m
}

type WSGServiceLDAPServiceList struct {
	Extra *WSGServiceLDAPServiceList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceLDAPServiceList `json:"list,omitempty"`

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
	AdminDomainName *WSGServiceModifyActiveDirectoryAuthentication `json:"adminDomainName,omitempty"`

	Description *WSGServiceModifyActiveDirectoryAuthentication `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceModifyActiveDirectoryAuthentication `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGServiceModifyActiveDirectoryAuthentication `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyActiveDirectoryAuthentication `json:"mappings,omitempty"`

	Name *WSGServiceModifyActiveDirectoryAuthentication `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty"`

	StandbyAdminDomainName *WSGServiceModifyActiveDirectoryAuthentication `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled standby cluster
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *WSGServiceModifyActiveDirectoryAuthentication `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password standby cluster
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port standby cluster
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled standby cluster
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	StandbyWindowsDomainName *WSGServiceModifyActiveDirectoryAuthentication `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[AD]
	Type *string `json:"type,omitempty"`

	WindowsDomainName *WSGServiceModifyActiveDirectoryAuthentication `json:"windowsDomainName,omitempty"`
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
	GroupAttr *string `json:"groupAttr"`

	// UserRole
	// Identity user role
	// Constraints:
	//    - required
	UserRole *WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"userRole"`
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

	Name *WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"name,omitempty"`
}

func NewWSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType() *WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
}

type WSGServiceModifyHlrAuthentication struct {
	// AddressIndicator
	// Constraints:
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty"`

	// AuthMapVer
	// Constraints:
	//    - oneof:[version2,version3]
	AuthMapVer *string `json:"authMapVer,omitempty"`

	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	Description *WSGServiceModifyHlrAuthentication `json:"description,omitempty"`

	// DestGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	DestGtIndicator *string `json:"destGtIndicator,omitempty"`

	// DestNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	DestNatureOfAddressIndicator *string `json:"destNatureOfAddressIndicator,omitempty"`

	// DestNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	DestNumberingPlan *string `json:"destNumberingPlan,omitempty"`

	DestTransType *int `json:"destTransType,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty"`

	FriendlyName *WSGServiceModifyHlrAuthentication `json:"friendlyName,omitempty"`

	GtPointCode *int `json:"gtPointCode,omitempty"`

	HasPointCode *bool `json:"hasPointCode,omitempty"`

	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	HasSSN *bool `json:"hasSSN,omitempty"`

	HistoryTime *int `json:"historyTime,omitempty"`

	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty"`

	LocalPointCode *int `json:"localPointCode,omitempty"`

	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	MncNdcList []*WSGServiceModifyHlrAuthentication `json:"mncNdcList,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGServiceModifyHlrAuthentication `json:"name,omitempty"`

	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceModifyHlrAuthentication `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceModifyHlrAuthentication `json:"sctpAssociationsList,omitempty"`

	SgsnIsdnAddress *string `json:"sgsnIsdnAddress,omitempty"`

	// SrcGtIndicator
	// Constraints:
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	SrcGtIndicator *string `json:"srcGtIndicator,omitempty"`

	// SrcNatureOfAddressIndicator
	// Constraints:
	//    - oneof:[unknown,subscriber_number,reserved_for_national_use,national_significant_number,international_number]
	SrcNatureOfAddressIndicator *string `json:"srcNatureOfAddressIndicator,omitempty"`

	// SrcNumberingPlan
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	SrcNumberingPlan *string `json:"srcNumberingPlan,omitempty"`

	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceModifyHlrAuthentication() *WSGServiceModifyHlrAuthentication {
	m := new(WSGServiceModifyHlrAuthentication)
	return m
}

type WSGServiceModifyLDAPAuthentication struct {
	AdminDomainName *WSGServiceModifyLDAPAuthentication `json:"adminDomainName,omitempty"`

	BaseDomainName *WSGServiceModifyLDAPAuthentication `json:"baseDomainName,omitempty"`

	Description *WSGServiceModifyLDAPAuthentication `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceModifyLDAPAuthentication `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	Ip *WSGServiceModifyLDAPAuthentication `json:"ip,omitempty"`

	KeyAttribute *WSGServiceModifyLDAPAuthentication `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyLDAPAuthentication `json:"mappings,omitempty"`

	Name *WSGServiceModifyLDAPAuthentication `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty"`

	SearchFilter *WSGServiceModifyLDAPAuthentication `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *WSGServiceModifyLDAPAuthentication `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *WSGServiceModifyLDAPAuthentication `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *WSGServiceModifyLDAPAuthentication `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *WSGServiceModifyLDAPAuthentication `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	// Constraints:
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *WSGServiceModifyLDAPAuthentication `json:"standbySearchFilter,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceModifyLDAPAuthentication() *WSGServiceModifyLDAPAuthentication {
	m := new(WSGServiceModifyLDAPAuthentication)
	return m
}

type WSGServiceModifyLocalDbAuthentication struct {
	Description *WSGServiceModifyLocalDbAuthentication `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceModifyLocalDbAuthentication `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyLocalDbAuthentication `json:"mappings,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	Name *WSGServiceModifyLocalDbAuthentication `json:"name,omitempty"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - oneof:[LOCAL_DB]
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[LOCAL_DB]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceModifyLocalDbAuthentication() *WSGServiceModifyLocalDbAuthentication {
	m := new(WSGServiceModifyLocalDbAuthentication)
	return m
}

type WSGServiceModifyRadiusAccounting struct {
	Description *WSGServiceModifyRadiusAccounting `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *WSGServiceModifyRadiusAccounting `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS accounting service
	Id *string `json:"id,omitempty"`

	Name *WSGServiceModifyRadiusAccounting `json:"name,omitempty"`

	Primary *WSGServiceModifyRadiusAccounting `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *WSGServiceModifyRadiusAccounting `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceModifyRadiusAccounting `json:"secondary,omitempty"`

	StandbyPrimary *WSGServiceModifyRadiusAccounting `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceModifyRadiusAccounting() *WSGServiceModifyRadiusAccounting {
	m := new(WSGServiceModifyRadiusAccounting)
	return m
}

type WSGServiceModifyRadiusAuthentication struct {
	Description *WSGServiceModifyRadiusAuthentication `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceModifyRadiusAuthentication `json:"friendlyName,omitempty"`

	HealthCheckPolicy *WSGServiceModifyRadiusAuthentication `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyRadiusAuthentication `json:"mappings,omitempty"`

	Name *WSGServiceModifyRadiusAuthentication `json:"name,omitempty"`

	Primary *WSGServiceModifyRadiusAuthentication `json:"primary,omitempty"`

	RateLimiting *WSGServiceModifyRadiusAuthentication `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceModifyRadiusAuthentication `json:"secondary,omitempty"`

	StandbyPrimary *WSGServiceModifyRadiusAuthentication `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication server type
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty"`
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

	Description *WSGServiceRadiusAccountingService `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *WSGServiceRadiusAccountingService `json:"healthCheckPolicy,omitempty"`

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

	Name *WSGServiceRadiusAccountingService `json:"name,omitempty"`

	Primary *WSGServiceRadiusAccountingService `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *WSGServiceRadiusAccountingService `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceRadiusAccountingService `json:"secondary,omitempty"`

	StandbyPrimary *WSGServiceRadiusAccountingService `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceRadiusAccountingService() *WSGServiceRadiusAccountingService {
	m := new(WSGServiceRadiusAccountingService)
	return m
}

type WSGServiceRadiusAccountingServiceList struct {
	Extra *WSGServiceRadiusAccountingServiceList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAccountingServiceList `json:"list,omitempty"`

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

	Description *WSGServiceRadiusAuthenticationService `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	FriendlyName *WSGServiceRadiusAuthenticationService `json:"friendlyName,omitempty"`

	HealthCheckPolicy *WSGServiceRadiusAuthenticationService `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS authentication service
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceRadiusAuthenticationService `json:"mappings,omitempty"`

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

	Name *WSGServiceRadiusAuthenticationService `json:"name,omitempty"`

	Primary *WSGServiceRadiusAuthenticationService `json:"primary,omitempty"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty"`

	RateLimiting *WSGServiceRadiusAuthenticationService `json:"rateLimiting,omitempty"`

	Secondary *WSGServiceRadiusAuthenticationService `json:"secondary,omitempty"`

	StandbyPrimary *WSGServiceRadiusAuthenticationService `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceRadiusAuthenticationService() *WSGServiceRadiusAuthenticationService {
	m := new(WSGServiceRadiusAuthenticationService)
	return m
}

type WSGServiceRadiusAuthenticationServiceList struct {
	Extra *WSGServiceRadiusAuthenticationServiceList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAuthenticationServiceList `json:"list,omitempty"`

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
	AddressIndicator *string `json:"addressIndicator,omitempty"`

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
	GtIndicator *string `json:"gtIndicator,omitempty"`

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
	NatureOfAddressIndicator *string `json:"natureOfAddressIndicator,omitempty"`

	// NumberingPlan
	// - For HLR Authentiaction server
	// Constraints:
	//    - oneof:[isdn_telephony_numbering_plan,land_mobile_numbering_plan,isdn_mobile_numbering_plan]
	NumberingPlan *string `json:"numberingPlan,omitempty"`

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
	NodeTermination *string `json:"nodeTermination,omitempty"`

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
	AutoFallbackDisable *bool `json:"autoFallbackDisable"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGServiceSecondaryRadiusServer `json:"ip"`

	// Port
	// RADIUS server port
	// Constraints:
	//    - required
	//    - default:1812
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// SharedSecret
	// RADIUS server shared secrect
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGServiceSecondaryRadiusServer() *WSGServiceSecondaryRadiusServer {
	m := new(WSGServiceSecondaryRadiusServer)
	return m
}

type WSGServiceTestingConfig struct {
	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	LoginRequest *WSGServiceTestingConfig `json:"loginRequest,omitempty"`
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
