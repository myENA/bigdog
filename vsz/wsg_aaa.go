package vsz

// API Version: v9_0

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
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings"`

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

func NewWSGAAAActiveDirectory() *WSGAAAActiveDirectory {
	m := new(WSGAAAActiveDirectory)
	return m
}

type WSGAAAActiveDirectoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAAAActiveDirectory `json:"list"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAAActiveDirectoryList() *WSGAAAActiveDirectoryList {
	m := new(WSGAAAActiveDirectoryList)
	return m
}

type WSGAAAAuthenticationServer struct {
	// Description
	// Description of the RADIUS server
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the RADIUS server
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings"`

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
	ServiceType *string `json:"serviceType,omitempty" validate:"oneof=Authentication Accounting"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone which the RADIUS server belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGAAAAuthenticationServer() *WSGAAAAuthenticationServer {
	m := new(WSGAAAAuthenticationServer)
	return m
}

type WSGAAAAuthenticationServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAAAAuthenticationServer `json:"list"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAAAuthenticationServerList() *WSGAAAAuthenticationServerList {
	m := new(WSGAAAAuthenticationServerList)
	return m
}

type WSGAAACreateActiveDirectoryServer struct {
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	// Constraints:
	//    - required
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled" validate:"required"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings"`

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

type WSGAAACreateAuthenticationServer struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary" validate:"required"`

	Secondary *WSGCommonRadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAACreateAuthenticationServer() *WSGAAACreateAuthenticationServer {
	m := new(WSGAAACreateAuthenticationServer)
	return m
}

type WSGAAACreateLDAPServer struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName" validate:"required,max=128,min=2"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName" validate:"required,max=64,min=2"`

	Description *WSGCommonDescription `json:"description,omitempty"`

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
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings"`

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
// User traffic profile mapping
type WSGAAAGroupAttrIdentityUserRoleMapping struct {
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
	UserRole *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

func NewWSGAAAGroupAttrIdentityUserRoleMapping() *WSGAAAGroupAttrIdentityUserRoleMapping {
	m := new(WSGAAAGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType
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
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings"`

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

func NewWSGAAALDAPServer() *WSGAAALDAPServer {
	m := new(WSGAAALDAPServer)
	return m
}

type WSGAAALDAPServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAAALDAPServer `json:"list"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAALDAPServerList() *WSGAAALDAPServerList {
	m := new(WSGAAALDAPServerList)
	return m
}

type WSGAAAModifyActiveDirectoryServer struct {
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings"`

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

type WSGAAAModifyAuthenticationServer struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings"`

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
// User traffic profile mapping
type WSGAAAModifyGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	// Constraints:
	//    - required
	GroupAttr *string `json:"groupAttr" validate:"required"`

	// UserRole
	// Identity user role
	// Constraints:
	//    - required
	UserRole *WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

func NewWSGAAAModifyGroupAttrIdentityUserRoleMapping() *WSGAAAModifyGroupAttrIdentityUserRoleMapping {
	m := new(WSGAAAModifyGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType
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

type WSGAAAModifyLDAPServer struct {
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName,omitempty"`

	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	KeyAttribute *WSGCommonNormalName2to64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings"`

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

type WSGAAATestAAAServerResult struct {
	// PrimaryServer
	// Primary server test result
	PrimaryServer *string `json:"primaryServer,omitempty"`

	// SecondaryServer
	// Secondary server test result
	SecondaryServer *string `json:"secondaryServer,omitempty"`
}

func NewWSGAAATestAAAServerResult() *WSGAAATestAAAServerResult {
	m := new(WSGAAATestAAAServerResult)
	return m
}

type WSGAAATestAuthenticationServer struct {
	// AaaServer
	// Constraints:
	//    - required
	AaaServer *WSGCommonGenericRef `json:"aaaServer" validate:"required"`

	// AaaType
	// Authentication/Accounting service protocol. RADIUS for Radius, AD and LDAP. RADIUSAcct for RADIUS Accounting
	// Constraints:
	//    - oneof:[RADIUS,RADIUSAcct]
	AaaType *string `json:"aaaType,omitempty" validate:"oneof=RADIUS RADIUSAcct"`

	// AuthProtocol
	// Authentication protocol
	// Constraints:
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP]
	AuthProtocol *string `json:"authProtocol,omitempty" validate:"oneof=PAP CHAP"`

	// Password
	// Password
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// ServerType
	// Radius server type.
	// Constraints:
	//    - oneof:[ADMIN,GLOBAL,ZONE]
	ServerType *string `json:"serverType,omitempty" validate:"oneof=ADMIN GLOBAL ZONE"`

	// UserName
	// User name
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`
}

func NewWSGAAATestAuthenticationServer() *WSGAAATestAuthenticationServer {
	m := new(WSGAAATestAuthenticationServer)
	return m
}
