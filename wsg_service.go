package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// WSGServiceActiveDirectoryService
//
// Definition: service_ActiveDirectoryService
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
	Port *int `json:"port,omitempty"`

	// Protocol
	// Authentication protocol.
	// Constraints:
	//    - oneof:[AD]
	Protocol *string `json:"protocol,omitempty"`

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
	StandbyPort *int `json:"standbyPort,omitempty"`

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
	Type *string `json:"type,omitempty"`

	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

type WSGServiceActiveDirectoryServiceAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceActiveDirectoryService
}

func newWSGServiceActiveDirectoryServiceAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceActiveDirectoryServiceAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceActiveDirectoryServiceAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceActiveDirectoryService)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceActiveDirectoryService() *WSGServiceActiveDirectoryService {
	m := new(WSGServiceActiveDirectoryService)
	return m
}

// WSGServiceActiveDirectoryServiceList
//
// Definition: service_ActiveDirectoryServiceList
type WSGServiceActiveDirectoryServiceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceActiveDirectoryService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceActiveDirectoryServiceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceActiveDirectoryServiceList
}

func newWSGServiceActiveDirectoryServiceListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceActiveDirectoryServiceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceActiveDirectoryServiceListAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceActiveDirectoryServiceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceActiveDirectoryServiceList() *WSGServiceActiveDirectoryServiceList {
	m := new(WSGServiceActiveDirectoryServiceList)
	return m
}

// WSGServiceCommonAccountingService
//
// Definition: service_commonAccountingService
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

// WSGServiceCommonAccountingServiceList
//
// Definition: service_commonAccountingServiceList
type WSGServiceCommonAccountingServiceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceCommonAccountingServiceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceCommonAccountingServiceList
}

func newWSGServiceCommonAccountingServiceListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceCommonAccountingServiceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceCommonAccountingServiceListAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceCommonAccountingServiceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceCommonAccountingServiceList() *WSGServiceCommonAccountingServiceList {
	m := new(WSGServiceCommonAccountingServiceList)
	return m
}

// WSGServiceCommonAuthenticationService
//
// Definition: service_commonAuthenticationService
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
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Authentication protocol same as protocol.
	// Constraints:
	//    - oneof:[RADIUS,AD,LDAP,FACEBOOK,LINKEDIN,GOOGLE,GENERICOAUTH,SOAP,HLR,LOCAL_DB,GUEST]
	Type *string `json:"type,omitempty"`
}

type WSGServiceCommonAuthenticationServiceAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceCommonAuthenticationService
}

func newWSGServiceCommonAuthenticationServiceAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceCommonAuthenticationServiceAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceCommonAuthenticationServiceAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceCommonAuthenticationService)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceCommonAuthenticationService() *WSGServiceCommonAuthenticationService {
	m := new(WSGServiceCommonAuthenticationService)
	return m
}

// WSGServiceCommonAuthenticationServiceList
//
// Definition: service_commonAuthenticationServiceList
type WSGServiceCommonAuthenticationServiceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceCommonAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceCommonAuthenticationServiceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceCommonAuthenticationServiceList
}

func newWSGServiceCommonAuthenticationServiceListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceCommonAuthenticationServiceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceCommonAuthenticationServiceListAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceCommonAuthenticationServiceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceCommonAuthenticationServiceList() *WSGServiceCommonAuthenticationServiceList {
	m := new(WSGServiceCommonAuthenticationServiceList)
	return m
}

// WSGServiceCreateActiveDirectoryAuthentication
//
// Definition: service_createActiveDirectoryAuthentication
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
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled"`

	// Id
	// Identifier of the authentication service
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

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
	StandbyPort *int `json:"standbyPort,omitempty"`

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
	TlsEnabled *bool `json:"tlsEnabled"`

	// Type
	// Authentication protocol.
	// Constraints:
	//    - oneof:[AD]
	Type *string `json:"type,omitempty"`

	// WindowsDomainName
	// Constraints:
	//    - required
	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName"`
}

func NewWSGServiceCreateActiveDirectoryAuthentication() *WSGServiceCreateActiveDirectoryAuthentication {
	m := new(WSGServiceCreateActiveDirectoryAuthentication)
	return m
}

// WSGServiceCreateHlrAuthentication
//
// Definition: service_createHlrAuthentication
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

	Description *WSGCommonDescription `json:"description,omitempty"`

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
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty"`

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
	Name *WSGCommonNormalName `json:"name"`

	PointCode *int `json:"pointCode,omitempty"`

	// Protocol
	// Constraints:
	//    - oneof:[HLR]
	Protocol *string `json:"protocol,omitempty"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty"`

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

// WSGServiceCreateLDAPAuthentication
//
// Definition: service_createLDAPAuthentication
type WSGServiceCreateLDAPAuthentication struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName"`

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
	Ip *WSGCommonIpAddress `json:"ip"`

	// KeyAttribute
	// Constraints:
	//    - required
	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGServiceModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

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
	SearchFilter *WSGCommonNormalName2to64 `json:"searchFilter"`

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
	StandbyPort *int `json:"standbyPort,omitempty"`

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

// WSGServiceCreateRadiusAccounting
//
// Definition: service_createRadiusAccounting
type WSGServiceCreateRadiusAccounting struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	HealthCheckPolicy *WSGCommonHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary"`

	// Protocol
	// Accounting protocol.
	// Constraints:
	//    - oneof:[RADIUS]
	Protocol *string `json:"protocol,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCreateRadiusAccounting() *WSGServiceCreateRadiusAccounting {
	m := new(WSGServiceCreateRadiusAccounting)
	return m
}

// WSGServiceCreateRadiusAuthentication
//
// Definition: service_createRadiusAuthentication
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
	Name *WSGCommonNormalName `json:"name"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary"`

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
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceCreateRadiusAuthentication() *WSGServiceCreateRadiusAuthentication {
	m := new(WSGServiceCreateRadiusAuthentication)
	return m
}

// WSGServiceDeleteBulkAccountingService
//
// Definition: service_deleteBulkAccountingService
type WSGServiceDeleteBulkAccountingService struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGServiceDeleteBulkAccountingService() *WSGServiceDeleteBulkAccountingService {
	m := new(WSGServiceDeleteBulkAccountingService)
	return m
}

// WSGServiceDeleteBulkAuthenticationService
//
// Definition: service_deleteBulkAuthenticationService
type WSGServiceDeleteBulkAuthenticationService struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGServiceDeleteBulkAuthenticationService() *WSGServiceDeleteBulkAuthenticationService {
	m := new(WSGServiceDeleteBulkAuthenticationService)
	return m
}

// WSGServiceDnsServer
//
// Definition: service_dnsServer
type WSGServiceDnsServer struct {
	// Ip
	// IP of server
	Ip *string `json:"ip,omitempty"`
}

func NewWSGServiceDnsServer() *WSGServiceDnsServer {
	m := new(WSGServiceDnsServer)
	return m
}

// WSGServiceDnsServerList
//
// Definition: service_dnsServerList
type WSGServiceDnsServerList []*WSGServiceDnsServer

func MakeWSGServiceDnsServerList() WSGServiceDnsServerList {
	m := make(WSGServiceDnsServerList, 0)
	return m
}

// WSGServiceGgsn
//
// Definition: service_ggsn
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

// WSGServiceGgsnConfig
//
// Definition: service_ggsnConfig
type WSGServiceGgsnConfig struct {
	DnsServerList WSGServiceDnsServerList `json:"dnsServerList,omitempty"`

	GgsnList WSGServiceGgsnList `json:"ggsnList,omitempty"`

	GtpSettings *WSGServiceGtpSettings `json:"gtpSettings,omitempty"`
}

type WSGServiceGgsnConfigAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceGgsnConfig
}

func newWSGServiceGgsnConfigAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceGgsnConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceGgsnConfigAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceGgsnConfig)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceGgsnConfig() *WSGServiceGgsnConfig {
	m := new(WSGServiceGgsnConfig)
	return m
}

// WSGServiceGgsnList
//
// Definition: service_ggsnList
type WSGServiceGgsnList []*WSGServiceGgsn

func MakeWSGServiceGgsnList() WSGServiceGgsnList {
	m := make(WSGServiceGgsnList, 0)
	return m
}

// WSGServiceGroupAttrIdentityUserRoleMapping
//
// Definition: service_groupAttrIdentityUserRoleMapping
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
	UserRole *WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole"`
}

func NewWSGServiceGroupAttrIdentityUserRoleMapping() *WSGServiceGroupAttrIdentityUserRoleMapping {
	m := new(WSGServiceGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGServiceGroupAttrIdentityUserRoleMappingUserRoleType
//
// Definition: service_groupAttrIdentityUserRoleMappingUserRoleType
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
// Definition: service_groupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
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

// WSGServiceGtpSettings
//
// Definition: service_gtpSettings
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

// WSGServiceHlrService
//
// Definition: service_hlrService
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

	Description *WSGCommonDescription `json:"description,omitempty"`

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
	Protocol *string `json:"protocol,omitempty"`

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

type WSGServiceHlrServiceAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceHlrService
}

func newWSGServiceHlrServiceAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceHlrServiceAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceHlrServiceAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceHlrService)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceHlrService() *WSGServiceHlrService {
	m := new(WSGServiceHlrService)
	return m
}

// WSGServiceHlrServiceList
//
// Definition: service_hlrServiceList
type WSGServiceHlrServiceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceHlrService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceHlrServiceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceHlrServiceList
}

func newWSGServiceHlrServiceListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceHlrServiceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceHlrServiceListAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceHlrServiceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceHlrServiceList() *WSGServiceHlrServiceList {
	m := new(WSGServiceHlrServiceList)
	return m
}

// WSGServiceLDAPService
//
// Definition: service_LDAPService
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
	Port *int `json:"port,omitempty"`

	// Protocol
	// Authentication protocol
	// Constraints:
	//    - oneof:[LDAP]
	Protocol *string `json:"protocol,omitempty"`

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
	StandbyPort *int `json:"standbyPort,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

type WSGServiceLDAPServiceAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceLDAPService
}

func newWSGServiceLDAPServiceAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceLDAPServiceAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceLDAPServiceAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceLDAPService)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceLDAPService() *WSGServiceLDAPService {
	m := new(WSGServiceLDAPService)
	return m
}

// WSGServiceLDAPServiceList
//
// Definition: service_LDAPServiceList
type WSGServiceLDAPServiceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceLDAPService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceLDAPServiceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceLDAPServiceList
}

func newWSGServiceLDAPServiceListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceLDAPServiceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceLDAPServiceListAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceLDAPServiceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceLDAPServiceList() *WSGServiceLDAPServiceList {
	m := new(WSGServiceLDAPServiceList)
	return m
}

// WSGServiceMncNdc
//
// Definition: service_mncNdc
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

// WSGServiceModifyActiveDirectoryAuthentication
//
// Definition: service_modifyActiveDirectoryAuthentication
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
	Port *int `json:"port,omitempty"`

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
	StandbyPort *int `json:"standbyPort,omitempty"`

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
	Type *string `json:"type,omitempty"`

	WindowsDomainName *WSGCommonNormalName2to64 `json:"windowsDomainName,omitempty"`
}

func NewWSGServiceModifyActiveDirectoryAuthentication() *WSGServiceModifyActiveDirectoryAuthentication {
	m := new(WSGServiceModifyActiveDirectoryAuthentication)
	return m
}

// WSGServiceModifyGroupAttrIdentityUserRoleMapping
//
// Definition: service_modifyGroupAttrIdentityUserRoleMapping
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
	UserRole *WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole"`
}

func NewWSGServiceModifyGroupAttrIdentityUserRoleMapping() *WSGServiceModifyGroupAttrIdentityUserRoleMapping {
	m := new(WSGServiceModifyGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGServiceModifyGroupAttrIdentityUserRoleMappingUserRoleType
//
// Definition: service_modifyGroupAttrIdentityUserRoleMappingUserRoleType
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

// WSGServiceModifyHlrAuthentication
//
// Definition: service_modifyHlrAuthentication
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

	Description *WSGCommonDescription `json:"description,omitempty"`

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
	LocalNetworkIndicator *string `json:"localNetworkIndicator,omitempty"`

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
	Protocol *string `json:"protocol,omitempty"`

	ReuseEnable *bool `json:"reuseEnable,omitempty"`

	RoutingContext *int `json:"routingContext,omitempty"`

	SccpGttList []*WSGServiceSccpGtt `json:"sccpGttList,omitempty"`

	SctpAssociationsList []*WSGServiceSctpAssociation `json:"sctpAssociationsList,omitempty"`

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

// WSGServiceModifyLDAPAuthentication
//
// Definition: service_modifyLDAPAuthentication
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
	Port *int `json:"port,omitempty"`

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
	StandbyPort *int `json:"standbyPort,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceModifyLDAPAuthentication() *WSGServiceModifyLDAPAuthentication {
	m := new(WSGServiceModifyLDAPAuthentication)
	return m
}

// WSGServiceModifyLocalDbAuthentication
//
// Definition: service_modifyLocalDbAuthentication
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

// WSGServiceModifyRadiusAccounting
//
// Definition: service_modifyRadiusAccounting
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
	Protocol *string `json:"protocol,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceModifyRadiusAccounting() *WSGServiceModifyRadiusAccounting {
	m := new(WSGServiceModifyRadiusAccounting)
	return m
}

// WSGServiceModifyRadiusAuthentication
//
// Definition: service_modifyRadiusAuthentication
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
	Type *string `json:"type,omitempty"`
}

func NewWSGServiceModifyRadiusAuthentication() *WSGServiceModifyRadiusAuthentication {
	m := new(WSGServiceModifyRadiusAuthentication)
	return m
}

// WSGServiceRadiusAccountingService
//
// Definition: service_radiusAccountingService
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
	Protocol *string `json:"protocol,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

type WSGServiceRadiusAccountingServiceAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceRadiusAccountingService
}

func newWSGServiceRadiusAccountingServiceAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceRadiusAccountingServiceAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceRadiusAccountingServiceAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceRadiusAccountingService)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceRadiusAccountingService() *WSGServiceRadiusAccountingService {
	m := new(WSGServiceRadiusAccountingService)
	return m
}

// WSGServiceRadiusAccountingServiceList
//
// Definition: service_radiusAccountingServiceList
type WSGServiceRadiusAccountingServiceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAccountingService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceRadiusAccountingServiceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceRadiusAccountingServiceList
}

func newWSGServiceRadiusAccountingServiceListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceRadiusAccountingServiceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceRadiusAccountingServiceListAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceRadiusAccountingServiceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceRadiusAccountingServiceList() *WSGServiceRadiusAccountingServiceList {
	m := new(WSGServiceRadiusAccountingServiceList)
	return m
}

// WSGServiceRadiusAuthenticationService
//
// Definition: service_radiusAuthenticationService
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
	Protocol *string `json:"protocol,omitempty"`

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
	Type *string `json:"type,omitempty"`
}

type WSGServiceRadiusAuthenticationServiceAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceRadiusAuthenticationService
}

func newWSGServiceRadiusAuthenticationServiceAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceRadiusAuthenticationServiceAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceRadiusAuthenticationServiceAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceRadiusAuthenticationService)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceRadiusAuthenticationService() *WSGServiceRadiusAuthenticationService {
	m := new(WSGServiceRadiusAuthenticationService)
	return m
}

// WSGServiceRadiusAuthenticationServiceList
//
// Definition: service_radiusAuthenticationServiceList
type WSGServiceRadiusAuthenticationServiceList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGServiceRadiusAuthenticationService `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGServiceRadiusAuthenticationServiceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceRadiusAuthenticationServiceList
}

func newWSGServiceRadiusAuthenticationServiceListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGServiceRadiusAuthenticationServiceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGServiceRadiusAuthenticationServiceListAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceRadiusAuthenticationServiceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceRadiusAuthenticationServiceList() *WSGServiceRadiusAuthenticationServiceList {
	m := new(WSGServiceRadiusAuthenticationServiceList)
	return m
}

// WSGServiceSccpGtt
//
// Definition: service_sccpGtt
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

// WSGServiceSctpAssociation
//
// Definition: service_sctpAssociation
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

// WSGServiceSecondaryRadiusServer
//
// Definition: service_secondaryRadiusServer
type WSGServiceSecondaryRadiusServer struct {
	// AutoFallbackDisable
	// Automatic fallback enabled or disabled
	// Constraints:
	//    - required
	AutoFallbackDisable *bool `json:"autoFallbackDisable"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

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

// WSGServiceTestingConfig
//
// Definition: service_testingConfig
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

// WSGServiceTestingConfigLoginRequestType
//
// Definition: service_testingConfigLoginRequestType
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
