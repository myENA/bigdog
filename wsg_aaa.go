package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// WSGAAAActiveDirectory
//
// Definition: aaa_activeDirectory
type WSGAAAActiveDirectory struct {
	// AdminDomainName
	// Admin domain name
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	// Description
	// Description of the active directory server
	Description *string `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the active directory server
	Id *string `json:"id,omitempty"`

	// Ip
	// IP address
	Ip *string `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the active directory server
	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *string `json:"standbyIp,omitempty"`

	StandbyPassword *string `json:"standbyPassword,omitempty"`

	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// WindowsDomainName
	// Windows domain name
	WindowsDomainName *string `json:"windowsDomainName,omitempty"`

	// ZoneId
	// Identifier of the zone which the active directory server belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGAAAActiveDirectoryAPIResponse struct {
	*RawAPIResponse
	Data *WSGAAAActiveDirectory
}

func newWSGAAAActiveDirectoryAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAAAActiveDirectoryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAAAActiveDirectoryAPIResponse) Hydrate() error {
	r.Data = new(WSGAAAActiveDirectory)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAAAActiveDirectory() *WSGAAAActiveDirectory {
	m := new(WSGAAAActiveDirectory)
	return m
}

// WSGAAAActiveDirectoryList
//
// Definition: aaa_activeDirectoryList
type WSGAAAActiveDirectoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAAAActiveDirectory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAAAActiveDirectoryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAAAActiveDirectoryList
}

func newWSGAAAActiveDirectoryListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAAAActiveDirectoryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAAAActiveDirectoryListAPIResponse) Hydrate() error {
	r.Data = new(WSGAAAActiveDirectoryList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAAAActiveDirectoryList() *WSGAAAActiveDirectoryList {
	m := new(WSGAAAActiveDirectoryList)
	return m
}

// WSGAAAAuthenticationServer
//
// Definition: aaa_authenticationServer
type WSGAAAAuthenticationServer struct {
	// Description
	// Description of the RADIUS server
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the RADIUS server
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the RADIUS server
	Name *string `json:"name,omitempty"`

	// PartnerDomainId
	// Identifier of the partner domain which the RADIUS server belongs to
	PartnerDomainId *string `json:"partnerDomainId,omitempty"`

	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	Secondary *WSGCommonRadiusServer `json:"secondary,omitempty"`

	// ServiceType
	// Identify the RADIUS server is belong to Accounting or Authentication
	// Constraints:
	//    - oneof:[Authentication,Accounting]
	ServiceType *string `json:"serviceType,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone which the RADIUS server belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGAAAAuthenticationServerAPIResponse struct {
	*RawAPIResponse
	Data *WSGAAAAuthenticationServer
}

func newWSGAAAAuthenticationServerAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAAAAuthenticationServerAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAAAAuthenticationServerAPIResponse) Hydrate() error {
	r.Data = new(WSGAAAAuthenticationServer)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAAAAuthenticationServer() *WSGAAAAuthenticationServer {
	m := new(WSGAAAAuthenticationServer)
	return m
}

// WSGAAAAuthenticationServerList
//
// Definition: aaa_authenticationServerList
type WSGAAAAuthenticationServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAAAAuthenticationServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAAAAuthenticationServerListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAAAAuthenticationServerList
}

func newWSGAAAAuthenticationServerListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAAAAuthenticationServerListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAAAAuthenticationServerListAPIResponse) Hydrate() error {
	r.Data = new(WSGAAAAuthenticationServerList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAAAAuthenticationServerList() *WSGAAAAuthenticationServerList {
	m := new(WSGAAAAuthenticationServerList)
	return m
}

// WSGAAACreateActiveDirectoryServer
//
// Definition: aaa_createActiveDirectoryServer
type WSGAAACreateActiveDirectoryServer struct {
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	// Constraints:
	//    - required
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *string `json:"standbyIp,omitempty"`

	StandbyPassword *string `json:"standbyPassword,omitempty"`

	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	WindowsDomainName *WSGCommonNormalNameAllowBlank `json:"windowsDomainName,omitempty"`
}

func NewWSGAAACreateActiveDirectoryServer() *WSGAAACreateActiveDirectoryServer {
	m := new(WSGAAACreateActiveDirectoryServer)
	return m
}

// WSGAAACreateAuthenticationServer
//
// Definition: aaa_createAuthenticationServer
type WSGAAACreateAuthenticationServer struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary"`

	Secondary *WSGCommonRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAACreateAuthenticationServer() *WSGAAACreateAuthenticationServer {
	m := new(WSGAAACreateAuthenticationServer)
	return m
}

// WSGAAACreateLDAPServer
//
// Definition: aaa_createLDAPServer
type WSGAAACreateLDAPServer struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName"`

	Description *WSGCommonDescription `json:"description,omitempty"`

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
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *string `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	StandbyPassword *string `json:"standbyPassword,omitempty"`

	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAACreateLDAPServer() *WSGAAACreateLDAPServer {
	m := new(WSGAAACreateLDAPServer)
	return m
}

// WSGAAAGroupAttrIdentityUserRoleMapping
//
// Definition: aaa_groupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type WSGAAAGroupAttrIdentityUserRoleMapping struct {
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
	UserRole *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole"`
}

func NewWSGAAAGroupAttrIdentityUserRoleMapping() *WSGAAAGroupAttrIdentityUserRoleMapping {
	m := new(WSGAAAGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType
//
// Definition: aaa_groupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType struct {
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// Id
	// Identity user role UUID
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	UserTrafficProfile *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

func NewWSGAAAGroupAttrIdentityUserRoleMappingUserRoleType() *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
}

// WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Definition: aaa_groupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Identity user role
type WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	// Id
	// User traffic profile UUID
	Id *string `json:"id,omitempty"`

	// Name
	// User traffic profile name
	Name *string `json:"name,omitempty"`
}

func NewWSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType() *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType {
	m := new(WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType)
	return m
}

// WSGAAALDAPServer
//
// Definition: aaa_LDAPServer
type WSGAAALDAPServer struct {
	// AdminDomainName
	// Admin domain name
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	// BaseDomainName
	// Base domain name
	BaseDomainName *string `json:"baseDomainName,omitempty"`

	// Description
	// Description of the LDAP server
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the LDAP server
	Id *string `json:"id,omitempty"`

	// Ip
	// IP address
	Ip *string `json:"ip,omitempty"`

	// KeyAttribute
	// Key attribute
	KeyAttribute *string `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the LDAP server
	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	// SearchFilter
	// Search filter
	SearchFilter *string `json:"searchFilter,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *string `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	StandbyPassword *string `json:"standbyPassword,omitempty"`

	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone which the LDAP server belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGAAALDAPServerAPIResponse struct {
	*RawAPIResponse
	Data *WSGAAALDAPServer
}

func newWSGAAALDAPServerAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAAALDAPServerAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAAALDAPServerAPIResponse) Hydrate() error {
	r.Data = new(WSGAAALDAPServer)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAAALDAPServer() *WSGAAALDAPServer {
	m := new(WSGAAALDAPServer)
	return m
}

// WSGAAALDAPServerList
//
// Definition: aaa_LDAPServerList
type WSGAAALDAPServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAAALDAPServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAAALDAPServerListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAAALDAPServerList
}

func newWSGAAALDAPServerListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAAALDAPServerListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAAALDAPServerListAPIResponse) Hydrate() error {
	r.Data = new(WSGAAALDAPServerList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAAALDAPServerList() *WSGAAALDAPServerList {
	m := new(WSGAAALDAPServerList)
	return m
}

// WSGAAAModifyActiveDirectoryServer
//
// Definition: aaa_modifyActiveDirectoryServer
type WSGAAAModifyActiveDirectoryServer struct {
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	StandbyIp *string `json:"standbyIp,omitempty"`

	StandbyPassword *string `json:"standbyPassword,omitempty"`

	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	WindowsDomainName *WSGCommonNormalNameAllowBlank `json:"windowsDomainName,omitempty"`
}

func NewWSGAAAModifyActiveDirectoryServer() *WSGAAAModifyActiveDirectoryServer {
	m := new(WSGAAAModifyActiveDirectoryServer)
	return m
}

// WSGAAAModifyAuthenticationServer
//
// Definition: aaa_modifyAuthenticationServer
type WSGAAAModifyAuthenticationServer struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	Secondary *WSGCommonRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAAModifyAuthenticationServer() *WSGAAAModifyAuthenticationServer {
	m := new(WSGAAAModifyAuthenticationServer)
	return m
}

// WSGAAAModifyGroupAttrIdentityUserRoleMapping
//
// Definition: aaa_modifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type WSGAAAModifyGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	// Constraints:
	//    - required
	GroupAttr *string `json:"groupAttr"`

	// UserRole
	// Identity user role
	// Constraints:
	//    - required
	UserRole *WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole"`
}

func NewWSGAAAModifyGroupAttrIdentityUserRoleMapping() *WSGAAAModifyGroupAttrIdentityUserRoleMapping {
	m := new(WSGAAAModifyGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType
//
// Definition: aaa_modifyGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// Id
	// Identity user role UUID
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName2to64 `json:"name,omitempty"`
}

func NewWSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType() *WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
}

// WSGAAAModifyLDAPServer
//
// Definition: aaa_modifyLDAPServer
type WSGAAAModifyLDAPServer struct {
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName,omitempty"`

	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	StandbyIp *string `json:"standbyIp,omitempty"`

	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	StandbyPassword *string `json:"standbyPassword,omitempty"`

	StandbyPort *int `json:"standbyPort,omitempty"`

	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAAModifyLDAPServer() *WSGAAAModifyLDAPServer {
	m := new(WSGAAAModifyLDAPServer)
	return m
}

// WSGAAATestAAAServerResult
//
// Definition: aaa_testAAAServerResult
type WSGAAATestAAAServerResult struct {
	// PrimaryServer
	// Primary server test result
	PrimaryServer *string `json:"primaryServer,omitempty"`

	// SecondaryServer
	// Secondary server test result
	SecondaryServer *string `json:"secondaryServer,omitempty"`
}

type WSGAAATestAAAServerResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGAAATestAAAServerResult
}

func newWSGAAATestAAAServerResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAAATestAAAServerResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAAATestAAAServerResultAPIResponse) Hydrate() error {
	r.Data = new(WSGAAATestAAAServerResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAAATestAAAServerResult() *WSGAAATestAAAServerResult {
	m := new(WSGAAATestAAAServerResult)
	return m
}

// WSGAAATestAuthenticationServer
//
// Definition: aaa_testAuthenticationServer
type WSGAAATestAuthenticationServer struct {
	// AaaServer
	// Constraints:
	//    - required
	AaaServer *WSGCommonGenericRef `json:"aaaServer"`

	// AaaType
	// Authentication/Accounting service protocol. RADIUS for Radius, AD and LDAP. RADIUSAcct for RADIUS Accounting
	// Constraints:
	//    - oneof:[RADIUS,RADIUSAcct]
	AaaType *string `json:"aaaType,omitempty"`

	// AuthProtocol
	// Authentication protocol
	// Constraints:
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP]
	AuthProtocol *string `json:"authProtocol,omitempty"`

	// Password
	// Password
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// ServerType
	// Radius server type.
	// Constraints:
	//    - oneof:[ADMIN,GLOBAL,ZONE]
	ServerType *string `json:"serverType,omitempty"`

	// UserName
	// User name
	// Constraints:
	//    - required
	UserName *string `json:"userName"`
}

func NewWSGAAATestAuthenticationServer() *WSGAAATestAuthenticationServer {
	m := new(WSGAAATestAuthenticationServer)
	return m
}
