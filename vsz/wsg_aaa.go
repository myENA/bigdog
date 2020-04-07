package vsz

// API Version: v9_0

type WSGAAAActiveDirectory struct {
	// AdminDomainName
	// Admin domain name
	// Constraints:
	//    - nullable
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	// Description
	// Description of the active directory server
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	// Constraints:
	//    - nullable
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Id
	// Identifier of the active directory server
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// IP address
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the active directory server
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - nullable
	Port *int `json:"port,omitempty"`

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Constraints:
	//    - nullable
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Constraints:
	//    - nullable
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Constraints:
	//    - nullable
	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// WindowsDomainName
	// Windows domain name
	// Constraints:
	//    - nullable
	WindowsDomainName *string `json:"windowsDomainName,omitempty"`

	// ZoneId
	// Identifier of the zone which the active directory server belongs to
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGAAAActiveDirectory() *WSGAAAActiveDirectory {
	m := new(WSGAAAActiveDirectory)
	return m
}

type WSGAAAActiveDirectoryList struct {
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
	List []*WSGAAAActiveDirectory `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAAActiveDirectoryList() *WSGAAAActiveDirectoryList {
	m := new(WSGAAAActiveDirectoryList)
	return m
}

type WSGAAAAuthenticationServer struct {
	// Description
	// Description of the RADIUS server
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the RADIUS server
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the RADIUS server
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// PartnerDomainId
	// Identifier of the partner domain which the RADIUS server belongs to
	// Constraints:
	//    - nullable
	PartnerDomainId *string `json:"partnerDomainId,omitempty"`

	// Primary
	// Constraints:
	//    - nullable
	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGCommonRadiusServer `json:"secondary,omitempty"`

	// ServiceType
	// Identify the RADIUS server is belong to Accounting or Authentication
	// Constraints:
	//    - nullable
	//    - oneof:[Authentication,Accounting]
	ServiceType *string `json:"serviceType,omitempty" validate:"omitempty,oneof=Authentication Accounting"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone which the RADIUS server belongs to
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGAAAAuthenticationServer() *WSGAAAAuthenticationServer {
	m := new(WSGAAAAuthenticationServer)
	return m
}

type WSGAAAAuthenticationServerList struct {
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
	List []*WSGAAAAuthenticationServer `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAAAuthenticationServerList() *WSGAAAAuthenticationServerList {
	m := new(WSGAAAAuthenticationServerList)
	return m
}

type WSGAAACreateActiveDirectoryServer struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	// Description
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

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
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Constraints:
	//    - nullable
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Constraints:
	//    - nullable
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	//    - default:false
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Constraints:
	//    - nullable
	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// WindowsDomainName
	// Constraints:
	//    - nullable
	WindowsDomainName *WSGCommonNormalNameAllowBlank `json:"windowsDomainName,omitempty"`
}

func NewWSGAAACreateActiveDirectoryServer() *WSGAAACreateActiveDirectoryServer {
	m := new(WSGAAACreateActiveDirectoryServer)
	return m
}

type WSGAAACreateAuthenticationServer struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Primary
	// Constraints:
	//    - required
	Primary *WSGCommonRadiusServer `json:"primary" validate:"required"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGCommonRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	//    - default:false
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
	AdminDomainName *WSGCommonNormalName2to128 `json:"adminDomainName" validate:"required"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *WSGCommonNormalName2to64 `json:"baseDomainName" validate:"required"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

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
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

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
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Constraints:
	//    - nullable
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Constraints:
	//    - nullable
	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Constraints:
	//    - nullable
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbySearchFilter
	// Constraints:
	//    - nullable
	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	//    - default:false
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAACreateLDAPServer() *WSGAAACreateLDAPServer {
	m := new(WSGAAACreateLDAPServer)
	return m
}

// WSGAAAGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
// Constraints:
//    - nullable
type WSGAAAGroupAttrIdentityUserRoleMapping struct {
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
	UserRole *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole" validate:"required"`
}

func NewWSGAAAGroupAttrIdentityUserRoleMapping() *WSGAAAGroupAttrIdentityUserRoleMapping {
	m := new(WSGAAAGroupAttrIdentityUserRoleMapping)
	return m
}

// WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
// Constraints:
//    - nullable
type WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType struct {
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
	UserTrafficProfile *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

func NewWSGAAAGroupAttrIdentityUserRoleMappingUserRoleType() *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGAAAGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
}

// WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Identity user role
// Constraints:
//    - nullable
type WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
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

func NewWSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType() *WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType {
	m := new(WSGAAAGroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType)
	return m
}

type WSGAAALDAPServer struct {
	// AdminDomainName
	// Admin domain name
	// Constraints:
	//    - nullable
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	// BaseDomainName
	// Base domain name
	// Constraints:
	//    - nullable
	BaseDomainName *string `json:"baseDomainName,omitempty"`

	// Description
	// Description of the LDAP server
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the LDAP server
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// IP address
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// KeyAttribute
	// Key attribute
	// Constraints:
	//    - nullable
	KeyAttribute *string `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGAAAGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// MvnoId
	// Tenant UUID
	// Constraints:
	//    - nullable
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the LDAP server
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Password
	// Admin password
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	// Constraints:
	//    - nullable
	Port *int `json:"port,omitempty"`

	// SearchFilter
	// Search filter
	// Constraints:
	//    - nullable
	SearchFilter *string `json:"searchFilter,omitempty"`

	// StandbyAdminDomainName
	// Constraints:
	//    - nullable
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Constraints:
	//    - nullable
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Constraints:
	//    - nullable
	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Constraints:
	//    - nullable
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbySearchFilter
	// Constraints:
	//    - nullable
	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone which the LDAP server belongs to
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGAAALDAPServer() *WSGAAALDAPServer {
	m := new(WSGAAALDAPServer)
	return m
}

type WSGAAALDAPServerList struct {
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
	List []*WSGAAALDAPServer `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAALDAPServerList() *WSGAAALDAPServerList {
	m := new(WSGAAALDAPServerList)
	return m
}

type WSGAAAModifyActiveDirectoryServer struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *WSGCommonNormalName2to64 `json:"adminDomainName,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	// Constraints:
	//    - nullable
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *WSGCommonIpAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

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
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Constraints:
	//    - nullable
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Constraints:
	//    - nullable
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Constraints:
	//    - nullable
	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// WindowsDomainName
	// Constraints:
	//    - nullable
	WindowsDomainName *WSGCommonNormalNameAllowBlank `json:"windowsDomainName,omitempty"`
}

func NewWSGAAAModifyActiveDirectoryServer() *WSGAAAModifyActiveDirectoryServer {
	m := new(WSGAAAModifyActiveDirectoryServer)
	return m
}

type WSGAAAModifyAuthenticationServer struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	// Constraints:
	//    - nullable
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Primary
	// Constraints:
	//    - nullable
	Primary *WSGCommonRadiusServer `json:"primary,omitempty"`

	// Secondary
	// Constraints:
	//    - nullable
	Secondary *WSGCommonRadiusServer `json:"secondary,omitempty"`

	// StandbyPrimary
	// Constraints:
	//    - nullable
	StandbyPrimary *WSGCommonRadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAAModifyAuthenticationServer() *WSGAAAModifyAuthenticationServer {
	m := new(WSGAAAModifyAuthenticationServer)
	return m
}

// WSGAAAModifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
type WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
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

func NewWSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType() *WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType {
	m := new(WSGAAAModifyGroupAttrIdentityUserRoleMappingUserRoleType)
	return m
}

type WSGAAAModifyLDAPServer struct {
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
	Mappings []*WSGAAAModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty" validate:"omitempty,dive"`

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
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Constraints:
	//    - nullable
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// Constraints:
	//    - nullable
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Constraints:
	//    - nullable
	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Constraints:
	//    - nullable
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Constraints:
	//    - nullable
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbySearchFilter
	// Constraints:
	//    - nullable
	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// Constraints:
	//    - nullable
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewWSGAAAModifyLDAPServer() *WSGAAAModifyLDAPServer {
	m := new(WSGAAAModifyLDAPServer)
	return m
}

type WSGAAATestAAAServerResult struct {
	// PrimaryServer
	// Primary server test result
	// Constraints:
	//    - nullable
	PrimaryServer *string `json:"primaryServer,omitempty"`

	// SecondaryServer
	// Secondary server test result
	// Constraints:
	//    - nullable
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
	//    - nullable
	//    - oneof:[RADIUS,RADIUSAcct]
	AaaType *string `json:"aaaType,omitempty" validate:"omitempty,oneof=RADIUS RADIUSAcct"`

	// AuthProtocol
	// Authentication protocol
	// Constraints:
	//    - nullable
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP]
	AuthProtocol *string `json:"authProtocol,omitempty" validate:"omitempty,oneof=PAP CHAP"`

	// Password
	// Password
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// ServerType
	// Radius server type.
	// Constraints:
	//    - nullable
	//    - oneof:[ADMIN,GLOBAL,ZONE]
	ServerType *string `json:"serverType,omitempty" validate:"omitempty,oneof=ADMIN GLOBAL ZONE"`

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
