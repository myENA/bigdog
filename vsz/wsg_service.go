package vsz

// API Version: v9_0

type WSGServiceActiveDirectoryService struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	// Constraints:
	//    - nullable
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the active directory authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Password
	// Admin password
	// Constraints:
	//    - nullable
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

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *WSGCommonNormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	// Constraints:
	//    - nullable
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled for standby cluster
	// Constraints:
	//    - nullable
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Constraints:
	//    - nullable
	StandbyWindowsDomainName *WSGCommonNormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	// Constraints:
	//    - nullable
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AD"`

	// WindowsDomainName
	// Constraints:
	//    - nullable
	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

func NewWSGServiceActiveDirectoryService() *WSGServiceActiveDirectoryService {
	m := new(WSGServiceActiveDirectoryService)
	return m
}

type WSGServiceActiveDirectoryServiceList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGServiceActiveDirectoryService `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceActiveDirectoryServiceList() *WSGServiceActiveDirectoryServiceList {
	m := new(WSGServiceActiveDirectoryServiceList)
	return m
}

type WSGServiceCommonAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the accounting service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
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

func NewWSGServiceCommonAccountingService() *WSGServiceCommonAccountingService {
	m := new(WSGServiceCommonAccountingService)
	return m
}

type WSGServiceCommonAccountingServiceList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGServiceCommonAccountingService `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceCommonAccountingServiceList() *WSGServiceCommonAccountingServiceList {
	m := new(WSGServiceCommonAccountingServiceList)
	return m
}

type WSGServiceCommonAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
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

func NewWSGServiceCommonAuthenticationService() *WSGServiceCommonAuthenticationService {
	m := new(WSGServiceCommonAuthenticationService)
	return m
}

type WSGServiceCommonAuthenticationServiceList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGServiceCommonAuthenticationService `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceCommonAuthenticationServiceList() *WSGServiceCommonAuthenticationServiceList {
	m := new(WSGServiceCommonAuthenticationServiceList)
	return m
}

type WSGServiceCreateActiveDirectoryAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	// Constraints:
	//    - required
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled" validate:"required"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Password
	// Admin password
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - required
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *WSGCommonNormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled for standby cluster
	// Constraints:
	//    - nullable
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password for standby cluster
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	//    - default:false
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled for standby cluster
	// Constraints:
	//    - nullable
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Constraints:
	//    - nullable
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

func NewWSGServiceCreateActiveDirectoryAuthentication() *WSGServiceCreateActiveDirectoryAuthentication {
	m := new(WSGServiceCreateActiveDirectoryAuthentication)
	return m
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

	// AuthorizationCachingEnabled
	// Constraints:
	//    - nullable
	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	// AvCachingEnabled
	// Constraints:
	//    - nullable
	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	// CacheOptionType
	// Constraints:
	//    - nullable
	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	// CleanUpTimeHour
	// Constraints:
	//    - nullable
	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	// CleanUpTimeMinute
	// Constraints:
	//    - nullable
	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	// CreateDateTime
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultPointCodeFormat
	// Constraints:
	//    - nullable
	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	// Description
	// Constraints:
	//    - nullable
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

	// DestTransType
	// Constraints:
	//    - nullable
	DestTransType *int `json:"destTransType,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// E164Address
	// Constraints:
	//    - nullable
	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GtPointCode
	// Constraints:
	//    - nullable
	GtPointCode *int `json:"gtPointCode,omitempty"`

	// HasPointCode
	// Constraints:
	//    - nullable
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSrcPointCode
	// Constraints:
	//    - nullable
	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	// HasSSN
	// Constraints:
	//    - nullable
	HasSSN *bool `json:"hasSSN,omitempty"`

	// HistoryTime
	// Constraints:
	//    - nullable
	HistoryTime *int `json:"historyTime,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"omitempty,oneof=international international_spare national national_spare"`

	// LocalPointCode
	// Constraints:
	//    - nullable
	LocalPointCode *int `json:"localPointCode,omitempty"`

	// LocationDeliveryEnabled
	// Constraints:
	//    - nullable
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// MaxReuseTimes
	// Constraints:
	//    - nullable
	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	// MncNdcList
	// Constraints:
	//    - nullable
	MncNdcList []*WSGServiceMncNdc `json:"mncNdcList,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PointCode
	// Constraints:
	//    - nullable
	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

	// ReuseEnable
	// Constraints:
	//    - nullable
	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	// RoutingContext
	// Constraints:
	//    - nullable
	RoutingContext *int `json:"routingContext,omitempty"`

	// SccpGttList
	// Constraints:
	//    - nullable
	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty" validate:"omitempty,dive"`

	// SctpAssociationsList
	// Constraints:
	//    - nullable
	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty" validate:"omitempty,dive"`

	// SgsnIsdnAddress
	// Constraints:
	//    - nullable
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

	// SrcTransType
	// Constraints:
	//    - nullable
	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=HLR"`
}

func NewWSGServiceCreateHlrAuthentication() *WSGServiceCreateHlrAuthentication {
	m := new(WSGServiceCreateHlrAuthentication)
	return m
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

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

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

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *WSGCommonNormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Constraints:
	//    - nullable
	StandbyBaseDomainName *WSGCommonNormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Constraints:
	//    - nullable
	StandbyKeyAttribute *WSGCommonNormalName2to64 `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// StandbySearchFilter
	// Constraints:
	//    - nullable
	StandbySearchFilter *WSGCommonNormalName2to64 `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	//    - default:false
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	// Constraints:
	//    - nullable
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

func NewWSGServiceCreateLDAPAuthentication() *WSGServiceCreateLDAPAuthentication {
	m := new(WSGServiceCreateLDAPAuthentication)
	return m
}

type WSGServiceCreateRadiusAccounting struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// HealthCheckPolicy
	// Constraints:
	//    - nullable
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

	// RateLimiting
	// Constraints:
	//    - nullable
	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	//    - default:false
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewWSGServiceCreateRadiusAccounting() *WSGServiceCreateRadiusAccounting {
	m := new(WSGServiceCreateRadiusAccounting)
	return m
}

type WSGServiceCreateRadiusAuthentication struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// HealthCheckPolicy
	// Constraints:
	//    - nullable
	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	// Constraints:
	//    - nullable
	//    - default:false
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary" validate:"required"`

	// RateLimiting
	// Constraints:
	//    - nullable
	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	//    - default:false
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol
	// Constraints:
	//    - nullable
	//    - default:'RADIUS'
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewWSGServiceCreateRadiusAuthentication() *WSGServiceCreateRadiusAuthentication {
	m := new(WSGServiceCreateRadiusAuthentication)
	return m
}

type WSGServiceDeleteBulkAccountingService struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGServiceDeleteBulkAccountingService() *WSGServiceDeleteBulkAccountingService {
	m := new(WSGServiceDeleteBulkAccountingService)
	return m
}

type WSGServiceDeleteBulkAuthenticationService struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGServiceDeleteBulkAuthenticationService() *WSGServiceDeleteBulkAuthenticationService {
	m := new(WSGServiceDeleteBulkAuthenticationService)
	return m
}

type WSGServiceDnsServer struct {
	// Ip
	// IP of server
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	DomainName *string `json:"domainName,omitempty"`

	// GgsnIPAddress
	// IP of GGSN
	// Constraints:
	//    - nullable
	GgsnIPAddress *string `json:"ggsnIPAddress,omitempty"`
}

func NewWSGServiceGgsn() *WSGServiceGgsn {
	m := new(WSGServiceGgsn)
	return m
}

type WSGServiceGgsnConfig struct {
	// DnsServerList
	// Constraints:
	//    - nullable
	DnsServerList WSGServiceDnsServerList `json:"dnsServerList,omitempty"`

	// GgsnList
	// Constraints:
	//    - nullable
	GgsnList WSGServiceGgsnList `json:"ggsnList,omitempty"`

	// GtpSettings
	// Constraints:
	//    - nullable
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
// Constraints:
//    - nullable
type WSGServiceGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	// Constraints:
	//    - required
	GroupAttr *string `json:"groupAttr" validate:"required"`

	// Id
	// Group attribute mapping UUID
	// Constraints:
	//    - nullable
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
// Constraints:
//    - nullable
type WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// FirewallProfileId
	// Constraints:
	//    - nullable
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// Id
	// Identity user role UUID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	// Constraints:
	//    - nullable
	UserTrafficProfile *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

func NewWSGServiceGroupAttrIdentityUserRoleMappingUserRoleType() *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
}

// WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Identity user role
// Constraints:
//    - nullable
type WSGServiceGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	// Id
	// User traffic profile UUID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// User traffic profile name
	// Constraints:
	//    - nullable
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

func NewWSGServiceGtpSettings() *WSGServiceGtpSettings {
	m := new(WSGServiceGtpSettings)
	return m
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
	// Constraints:
	//    - nullable
	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	// AvCachingEnabled
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	// CacheOptionType
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	// CleanUpTimeHour
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	// CleanUpTimeMinute
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultPointCodeFormat
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	// Description
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	DestTransType *int `json:"destTransType,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// E164Address
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GtPointCode
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	GtPointCode *int `json:"gtPointCode,omitempty"`

	// HasPointCode
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSrcPointCode
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	// HasSSN
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	HasSSN *bool `json:"hasSSN,omitempty"`

	// HistoryTime
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	HistoryTime *int `json:"historyTime,omitempty"`

	// Id
	// Identifier of the HLR authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"omitempty,oneof=international international_spare national national_spare"`

	// LocalPointCode
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	LocalPointCode *int `json:"localPointCode,omitempty"`

	// LocationDeliveryEnabled
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// MaxReuseTimes
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	// MncNdcList
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	MncNdcList []*WSGServiceMncNdc `json:"mncNdcList,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PointCode
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

	// ReuseEnable
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	// RoutingContext
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	RoutingContext *int `json:"routingContext,omitempty"`

	// SccpGttList
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty" validate:"omitempty,dive"`

	// SctpAssociationsList
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty" validate:"omitempty,dive"`

	// SgsnIsdnAddress
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=HLR"`
}

func NewWSGServiceHlrService() *WSGServiceHlrService {
	m := new(WSGServiceHlrService)
	return m
}

type WSGServiceHlrServiceList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGServiceHlrService `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceHlrServiceList() *WSGServiceHlrServiceList {
	m := new(WSGServiceHlrServiceList)
	return m
}

type WSGServiceLDAPService struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName,omitempty"`

	// BaseDomainName
	// Constraints:
	//    - nullable
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the LDAP authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// KeyAttribute
	// Constraints:
	//    - nullable
	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Password
	// Admin password
	// Constraints:
	//    - nullable
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

	// SearchFilter
	// Constraints:
	//    - nullable
	SearchFilter *WSGCommonNormalName2to64 `json:"searchFilter,omitempty"`

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *WSGCommonNormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Constraints:
	//    - nullable
	StandbyBaseDomainName *WSGCommonNormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Constraints:
	//    - nullable
	StandbyKeyAttribute *WSGCommonNormalName2to64 `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// StandbySearchFilter
	// Constraints:
	//    - nullable
	StandbySearchFilter *WSGCommonNormalName2to64 `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	// Constraints:
	//    - nullable
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TlsEnabled
	// LDAP over TLS Enabled
	// Constraints:
	//    - nullable
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol same as protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=LDAP"`
}

func NewWSGServiceLDAPService() *WSGServiceLDAPService {
	m := new(WSGServiceLDAPService)
	return m
}

type WSGServiceLDAPServiceList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGServiceLDAPService `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceLDAPServiceList() *WSGServiceLDAPServiceList {
	m := new(WSGServiceLDAPServiceList)
	return m
}

type WSGServiceMncNdc struct {
	// Mcc
	// MCC
	// Constraints:
	//    - nullable
	Mcc *string `json:"mcc,omitempty"`

	// Mnc
	// MNC
	// Constraints:
	//    - nullable
	Mnc *string `json:"mnc,omitempty"`

	// MvnoId
	// MVNO ID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Ndc
	// NDC
	// Constraints:
	//    - nullable
	Ndc *string `json:"ndc,omitempty"`

	// ProfileId
	// Profile ID
	// Constraints:
	//    - nullable
	ProfileId *string `json:"profileId,omitempty"`
}

func NewWSGServiceMncNdc() *WSGServiceMncNdc {
	m := new(WSGServiceMncNdc)
	return m
}

type WSGServiceModifyActiveDirectoryAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GlobalCatalogEnabled
	// Global catalog support enabled or disabled
	// Constraints:
	//    - nullable
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Password
	// Admin password
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *WSGCommonNormalName2to64 `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Global catalog support enabled or disabled standby cluster
	// Constraints:
	//    - nullable
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password standby cluster
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// AD over TLS Enabled standby cluster
	// Constraints:
	//    - nullable
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Constraints:
	//    - nullable
	StandbyWindowsDomainName *WSGCommonNormalName2to64 `json:"standbyWindowsDomainName,omitempty"`

	// TlsEnabled
	// AD over TLS Enabled
	// Constraints:
	//    - nullable
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[AD]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AD"`

	// WindowsDomainName
	// Constraints:
	//    - nullable
	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

func NewWSGServiceModifyActiveDirectoryAuthentication() *WSGServiceModifyActiveDirectoryAuthentication {
	m := new(WSGServiceModifyActiveDirectoryAuthentication)
	return m
}

// WSGServiceModifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
type WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// Id
	// Identity user role UUID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`
}

func NewWSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType() *WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
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

	// AuthorizationCachingEnabled
	// Constraints:
	//    - nullable
	AuthorizationCachingEnabled *bool `json:"authorizationCachingEnabled,omitempty"`

	// AvCachingEnabled
	// Constraints:
	//    - nullable
	AvCachingEnabled *bool `json:"avCachingEnabled,omitempty"`

	// CacheOptionType
	// Constraints:
	//    - nullable
	CacheOptionType *string `json:"cacheOptionType,omitempty"`

	// CleanUpTimeHour
	// Constraints:
	//    - nullable
	CleanUpTimeHour *int `json:"cleanUpTimeHour,omitempty"`

	// CleanUpTimeMinute
	// Constraints:
	//    - nullable
	CleanUpTimeMinute *int `json:"cleanUpTimeMinute,omitempty"`

	// CreateDateTime
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DefaultPointCodeFormat
	// Constraints:
	//    - nullable
	DefaultPointCodeFormat *string `json:"defaultPointCodeFormat,omitempty"`

	// Description
	// Constraints:
	//    - nullable
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

	// DestTransType
	// Constraints:
	//    - nullable
	DestTransType *int `json:"destTransType,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// E164Address
	// Constraints:
	//    - nullable
	E164Address *string `json:"e164Address,omitempty"`

	// EapSimMapVer
	// Constraints:
	//    - nullable
	//    - oneof:[version2,version3]
	EapSimMapVer *string `json:"eapSimMapVer,omitempty" validate:"omitempty,oneof=version2 version3"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// GtPointCode
	// Constraints:
	//    - nullable
	GtPointCode *int `json:"gtPointCode,omitempty"`

	// HasPointCode
	// Constraints:
	//    - nullable
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSrcPointCode
	// Constraints:
	//    - nullable
	HasSrcPointCode *bool `json:"hasSrcPointCode,omitempty"`

	// HasSSN
	// Constraints:
	//    - nullable
	HasSSN *bool `json:"hasSSN,omitempty"`

	// HistoryTime
	// Constraints:
	//    - nullable
	HistoryTime *int `json:"historyTime,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LocalNetworkIndicator
	// Constraints:
	//    - nullable
	//    - oneof:[international,international_spare,national,national_spare]
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty" validate:"omitempty,oneof=international international_spare national national_spare"`

	// LocalPointCode
	// Constraints:
	//    - nullable
	LocalPointCode *int `json:"localPointCode,omitempty"`

	// LocationDeliveryEnabled
	// Constraints:
	//    - nullable
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// MaxReuseTimes
	// Constraints:
	//    - nullable
	MaxReuseTimes *int `json:"maxReuseTimes,omitempty"`

	// MncNdcList
	// Constraints:
	//    - nullable
	MncNdcList []*WSGServiceMncNdc `json:"mncNdcList,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PointCode
	// Constraints:
	//    - nullable
	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=HLR"`

	// ReuseEnable
	// Constraints:
	//    - nullable
	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	// RoutingContext
	// Constraints:
	//    - nullable
	RoutingContext *int `json:"routingContext,omitempty"`

	// SccpGttList
	// Constraints:
	//    - nullable
	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty" validate:"omitempty,dive"`

	// SctpAssociationsList
	// Constraints:
	//    - nullable
	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty" validate:"omitempty,dive"`

	// SgsnIsdnAddress
	// Constraints:
	//    - nullable
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

	// SrcTransType
	// Constraints:
	//    - nullable
	SrcTransType *int `json:"srcTransType,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[HLR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=HLR"`
}

func NewWSGServiceModifyHlrAuthentication() *WSGServiceModifyHlrAuthentication {
	m := new(WSGServiceModifyHlrAuthentication)
	return m
}

type WSGServiceModifyLDAPAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName,omitempty"`

	// BaseDomainName
	// Constraints:
	//    - nullable
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// KeyAttribute
	// Constraints:
	//    - nullable
	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Password
	// Admin password
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// SearchFilter
	// Constraints:
	//    - nullable
	SearchFilter *WSGCommonNormalName2to64 `json:"searchFilter,omitempty"`

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *WSGCommonNormalName2to128 `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Constraints:
	//    - nullable
	StandbyBaseDomainName *WSGCommonNormalName2to64 `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *WSGCommonIpAddress `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Constraints:
	//    - nullable
	StandbyKeyAttribute *WSGCommonNormalName2to64 `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	// Constraints:
	//    - nullable
	//    - default:389
	//    - min:1
	//    - max:65535
	StandbyPort *int `json:"standbyPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// StandbySearchFilter
	// Constraints:
	//    - nullable
	StandbySearchFilter *WSGCommonNormalName2to64 `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyTlsEnabled
	// LDAP over TLS Enabled - Standby Cluster settings
	// Constraints:
	//    - nullable
	StandbyTlsEnabled *bool `json:"standbyTlsEnabled,omitempty"`

	// TlsEnabled
	// LDAP over TLS Enabled
	// Constraints:
	//    - nullable
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	// Type
	// Authentication protocol same as protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=LDAP"`
}

func NewWSGServiceModifyLDAPAuthentication() *WSGServiceModifyLDAPAuthentication {
	m := new(WSGServiceModifyLDAPAuthentication)
	return m
}

type WSGServiceModifyLocalDbAuthentication struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
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

func NewWSGServiceModifyLocalDbAuthentication() *WSGServiceModifyLocalDbAuthentication {
	m := new(WSGServiceModifyLocalDbAuthentication)
	return m
}

type WSGServiceModifyRadiusAccounting struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// HealthCheckPolicy
	// Constraints:
	//    - nullable
	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS accounting service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Primary
	// Constraints:
	//    - nullable
	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	// RateLimiting
	// Constraints:
	//    - nullable
	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewWSGServiceModifyRadiusAccounting() *WSGServiceModifyRadiusAccounting {
	m := new(WSGServiceModifyRadiusAccounting)
	return m
}

type WSGServiceModifyRadiusAuthentication struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// HealthCheckPolicy
	// Constraints:
	//    - nullable
	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	// Constraints:
	//    - nullable
	//    - default:false
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Primary
	// Constraints:
	//    - nullable
	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// RateLimiting
	// Constraints:
	//    - nullable
	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication server type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewWSGServiceModifyRadiusAuthentication() *WSGServiceModifyRadiusAuthentication {
	m := new(WSGServiceModifyRadiusAuthentication)
	return m
}

type WSGServiceRadiusAccountingService struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// HealthCheckPolicy
	// Constraints:
	//    - nullable
	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS accounting service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Primary
	// Constraints:
	//    - nullable
	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	// RateLimiting
	// Constraints:
	//    - nullable
	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Accounting protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewWSGServiceRadiusAccountingService() *WSGServiceRadiusAccountingService {
	m := new(WSGServiceRadiusAccountingService)
	return m
}

type WSGServiceRadiusAccountingServiceList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGServiceRadiusAccountingService `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGServiceRadiusAccountingServiceList() *WSGServiceRadiusAccountingServiceList {
	m := new(WSGServiceRadiusAccountingServiceList)
	return m
}

type WSGServiceRadiusAuthenticationService struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FriendlyName
	// Constraints:
	//    - nullable
	FriendlyName *WSGCommonNormalNameAllowBlank `json:"friendlyName,omitempty"`

	// HealthCheckPolicy
	// Constraints:
	//    - nullable
	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Id
	// Identifier of the RADIUS authentication service
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LocationDeliveryEnabled
	// RFC5580 out of band location delivery support(for Ruckus AP only)
	// Constraints:
	//    - nullable
	//    - default:false
	LocationDeliveryEnabled *bool `json:"locationDeliveryEnabled,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGServiceGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Primary
	// Constraints:
	//    - nullable
	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=RADIUS"`

	// RateLimiting
	// Constraints:
	//    - nullable
	RateLimiting *WSGCommonRateLimiting `json:"rateLimiting,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGServiceSecondaryRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS"`
}

func NewWSGServiceRadiusAuthenticationService() *WSGServiceRadiusAuthenticationService {
	m := new(WSGServiceRadiusAuthenticationService)
	return m
}

type WSGServiceRadiusAuthenticationServiceList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGServiceRadiusAuthenticationService `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
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
	//    - nullable
	//    - oneof:[route_on_gt,route_on_ssn]
	AddressIndicator *string `json:"addressIndicator,omitempty" validate:"omitempty,oneof=route_on_gt route_on_ssn"`

	// E164Address
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	E164Address *string `json:"e164Address,omitempty"`

	// GtDigits
	// GT digits
	// Constraints:
	//    - nullable
	GtDigits *string `json:"gtDigits,omitempty"`

	// GtIndicator
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	//    - oneof:[global_title_includes_translation_type_only,global_title_includes_translation_type_numbering_plan_and_encoding_scheme,global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator]
	GtIndicator *string `json:"gtIndicator,omitempty" validate:"omitempty,oneof=global_title_includes_translation_type_only global_title_includes_translation_type_numbering_plan_and_encoding_scheme global_title_includes_translation_type_numbering_plan_encoding_scheme_and_nature_of_address_indicator"`

	// HasPointCode
	// Enable Point Code submitted
	// Constraints:
	//    - nullable
	HasPointCode *bool `json:"hasPointCode,omitempty"`

	// HasSSN
	// Enable SSN- For HLR Authentiaction server
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	PointCode *int `json:"pointCode,omitempty"`

	// TransType
	// - For HLR Authentiaction server
	// Constraints:
	//    - nullable
	TransType *int `json:"transType,omitempty"`
}

func NewWSGServiceSccpGtt() *WSGServiceSccpGtt {
	m := new(WSGServiceSccpGtt)
	return m
}

type WSGServiceSctpAssociation struct {
	// AdjPointCode
	// Adj Pointcode
	// Constraints:
	//    - nullable
	AdjPointCode *string `json:"adjPointCode,omitempty"`

	// DestinationIp
	// Destination IP
	// Constraints:
	//    - nullable
	DestinationIp *string `json:"destinationIp,omitempty"`

	// DestinationPort
	// Destination Port
	// Constraints:
	//    - nullable
	DestinationPort *int `json:"destinationPort,omitempty"`

	// MaxInboundsStreams
	// NDC
	// Constraints:
	//    - nullable
	MaxInboundsStreams *int `json:"maxInboundsStreams,omitempty"`

	// MaxOutboundsStreams
	// Profile ID
	// Constraints:
	//    - nullable
	MaxOutboundsStreams *int `json:"maxOutboundsStreams,omitempty"`

	// NodeTermination
	// Node termination
	// Constraints:
	//    - nullable
	//    - oneof:[active,backup,both]
	NodeTermination *string `json:"nodeTermination,omitempty" validate:"omitempty,oneof=active backup both"`

	// SourcePort
	// Source port
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LoginRequest
	// Constraints:
	//    - nullable
	LoginRequest *WSGServiceTestingConfigLoginRequestType `json:"loginRequest,omitempty"`
}

func NewWSGServiceTestingConfig() *WSGServiceTestingConfig {
	m := new(WSGServiceTestingConfig)
	return m
}

type WSGServiceTestingConfigLoginRequestType struct {
	// Password
	// password for test user
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// Protocol
	// Constraints:
	//    - nullable
	Protocol *string `json:"protocol,omitempty"`

	// TimeZoneUtcOffset
	// timezone offset, ex: '+8'
	// Constraints:
	//    - nullable
	TimeZoneUtcOffset *string `json:"timeZoneUtcOffset,omitempty"`

	// UserName
	// name for test user
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`
}

func NewWSGServiceTestingConfigLoginRequestType() *WSGServiceTestingConfigLoginRequestType {
	m := new(WSGServiceTestingConfigLoginRequestType)
	return m
}
