package aaa

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ActiveDirectory struct {
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
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Enable global catalog support - Standby Cluster settings
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// IP address - Standby Cluster settings
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Windows domain name - Standby Cluster settings
	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	// WindowsDomainName
	// Windows domain name
	WindowsDomainName *string `json:"windowsDomainName,omitempty"`

	// ZoneId
	// Identifier of the zone which the active directory server belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewActiveDirectory() *ActiveDirectory {
	activeDirectoryType := new(ActiveDirectory)
	return activeDirectoryType
}

func NewActiveDirectoryWithDefaults() *ActiveDirectory {
	activeDirectoryType := new(ActiveDirectory)
	return activeDirectoryType
}

type ActiveDirectoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ActiveDirectory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewActiveDirectoryList() *ActiveDirectoryList {
	activeDirectoryListType := new(ActiveDirectoryList)
	return activeDirectoryListType
}

func NewActiveDirectoryListWithDefaults() *ActiveDirectoryList {
	activeDirectoryListType := new(ActiveDirectoryList)
	return activeDirectoryListType
}

type AuthenticationServer struct {
	// Description
	// Description of the RADIUS server
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of the RADIUS server
	Id *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoId
	// Tenant UUID
	MvnoId *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the RADIUS server
	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	Secondary *common.RadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone which the RADIUS server belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewAuthenticationServer() *AuthenticationServer {
	authenticationServerType := new(AuthenticationServer)
	return authenticationServerType
}

func NewAuthenticationServerWithDefaults() *AuthenticationServer {
	authenticationServerType := new(AuthenticationServer)
	return authenticationServerType
}

type AuthenticationServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AuthenticationServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAuthenticationServerList() *AuthenticationServerList {
	authenticationServerListType := new(AuthenticationServerList)
	return authenticationServerListType
}

func NewAuthenticationServerListWithDefaults() *AuthenticationServerList {
	authenticationServerListType := new(AuthenticationServerList)
	return authenticationServerListType
}

type CreateActiveDirectoryServer struct {
	AdminDomainName *common.NormalName2to64 `json:"adminDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	// Constraints:
	//    - required
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled" validate:"required"`

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

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Enable global catalog support - Standby Cluster settings
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// IP address - Standby Cluster settings
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Windows domain name - Standby Cluster settings
	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	WindowsDomainName *common.NormalNameAllowBlank `json:"windowsDomainName,omitempty"`
}

func NewCreateActiveDirectoryServer() *CreateActiveDirectoryServer {
	createActiveDirectoryServerType := new(CreateActiveDirectoryServer)
	return createActiveDirectoryServerType
}

func NewCreateActiveDirectoryServerWithDefaults() *CreateActiveDirectoryServer {
	createActiveDirectoryServerType := new(CreateActiveDirectoryServer)
	portField := 389
	createActiveDirectoryServerType.Port = &portField
	standbyServerEnabledField := false
	createActiveDirectoryServerType.StandbyServerEnabled = &standbyServerEnabledField
	return createActiveDirectoryServerType
}

type CreateAuthenticationServer struct {
	Description *common.Description `json:"description,omitempty"`

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

	Secondary *common.RadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewCreateAuthenticationServer() *CreateAuthenticationServer {
	createAuthenticationServerType := new(CreateAuthenticationServer)
	return createAuthenticationServerType
}

func NewCreateAuthenticationServerWithDefaults() *CreateAuthenticationServer {
	createAuthenticationServerType := new(CreateAuthenticationServer)
	standbyServerEnabledField := false
	createAuthenticationServerType.StandbyServerEnabled = &standbyServerEnabledField
	return createAuthenticationServerType
}

type CreateLDAPServer struct {
	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName *common.NormalName2to128 `json:"adminDomainName" validate:"required"`

	// BaseDomainName
	// Constraints:
	//    - required
	BaseDomainName *common.NormalName2to64 `json:"baseDomainName" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

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

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Base domain name - Standby Cluster settings
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// IP address - Standby Cluster settings
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Key attribute - Standby Cluster settings
	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbySearchFilter
	// Search filter - Standby Cluster settings
	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewCreateLDAPServer() *CreateLDAPServer {
	createLDAPServerType := new(CreateLDAPServer)
	return createLDAPServerType
}

func NewCreateLDAPServerWithDefaults() *CreateLDAPServer {
	createLDAPServerType := new(CreateLDAPServer)
	portField := 389
	createLDAPServerType.Port = &portField
	standbyServerEnabledField := false
	createLDAPServerType.StandbyServerEnabled = &standbyServerEnabledField
	return createLDAPServerType
}

type DeleteBulkAAAServerList struct {
	AaadeleteBulkAAAServerList *string `json:"aaa_deleteBulkAAAServerList,omitempty"`
}

func NewDeleteBulkAAAServerList() *DeleteBulkAAAServerList {
	deleteBulkAAAServerListType := new(DeleteBulkAAAServerList)
	return deleteBulkAAAServerListType
}

func NewDeleteBulkAAAServerListWithDefaults() *DeleteBulkAAAServerList {
	deleteBulkAAAServerListType := new(DeleteBulkAAAServerList)
	return deleteBulkAAAServerListType
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

type LDAPServer struct {
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
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

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

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Base domain name - Standby Cluster settings
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// IP address - Standby Cluster settings
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Key attribute - Standby Cluster settings
	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbySearchFilter
	// Search filter - Standby Cluster settings
	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneId
	// Identifier of the zone which the LDAP server belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewLDAPServer() *LDAPServer {
	lDAPServerType := new(LDAPServer)
	return lDAPServerType
}

func NewLDAPServerWithDefaults() *LDAPServer {
	lDAPServerType := new(LDAPServer)
	return lDAPServerType
}

type LDAPServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LDAPServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewLDAPServerList() *LDAPServerList {
	lDAPServerListType := new(LDAPServerList)
	return lDAPServerListType
}

func NewLDAPServerListWithDefaults() *LDAPServerList {
	lDAPServerListType := new(LDAPServerList)
	return lDAPServerListType
}

type ModifyActiveDirectoryServer struct {
	AdminDomainName *common.NormalName2to64 `json:"adminDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

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

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyGlobalCatalogEnabled
	// Enable global catalog support - Standby Cluster settings
	StandbyGlobalCatalogEnabled *bool `json:"standbyGlobalCatalogEnabled,omitempty"`

	// StandbyIp
	// IP address - Standby Cluster settings
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// StandbyWindowsDomainName
	// Windows domain name - Standby Cluster settings
	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	WindowsDomainName *common.NormalNameAllowBlank `json:"windowsDomainName,omitempty"`
}

func NewModifyActiveDirectoryServer() *ModifyActiveDirectoryServer {
	modifyActiveDirectoryServerType := new(ModifyActiveDirectoryServer)
	return modifyActiveDirectoryServerType
}

func NewModifyActiveDirectoryServerWithDefaults() *ModifyActiveDirectoryServer {
	modifyActiveDirectoryServerType := new(ModifyActiveDirectoryServer)
	portField := 389
	modifyActiveDirectoryServerType.Port = &portField
	return modifyActiveDirectoryServerType
}

type ModifyAuthenticationServer struct {
	Description *common.Description `json:"description,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	Secondary *common.RadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewModifyAuthenticationServer() *ModifyAuthenticationServer {
	modifyAuthenticationServerType := new(ModifyAuthenticationServer)
	return modifyAuthenticationServerType
}

func NewModifyAuthenticationServerWithDefaults() *ModifyAuthenticationServer {
	modifyAuthenticationServerType := new(ModifyAuthenticationServer)
	return modifyAuthenticationServerType
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

type ModifyLDAPServer struct {
	AdminDomainName *common.NormalName2to128 `json:"adminDomainName,omitempty"`

	BaseDomainName *common.NormalName2to64 `json:"baseDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

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

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Base domain name - Standby Cluster settings
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIp
	// IP address - Standby Cluster settings
	StandbyIp *string `json:"standbyIp,omitempty"`

	// StandbyKeyAttribute
	// Key attribute - Standby Cluster settings
	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	// StandbyPassword
	// Admin password - Standby Cluster settings
	StandbyPassword *string `json:"standbyPassword,omitempty"`

	// StandbyPort
	// Port - Standby Cluster settings
	StandbyPort *int `json:"standbyPort,omitempty"`

	// StandbySearchFilter
	// Search filter - Standby Cluster settings
	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`
}

func NewModifyLDAPServer() *ModifyLDAPServer {
	modifyLDAPServerType := new(ModifyLDAPServer)
	return modifyLDAPServerType
}

func NewModifyLDAPServerWithDefaults() *ModifyLDAPServer {
	modifyLDAPServerType := new(ModifyLDAPServer)
	portField := 389
	modifyLDAPServerType.Port = &portField
	return modifyLDAPServerType
}

type TestAAAServerResult struct {
	// PrimaryServer
	// Primary server test result
	PrimaryServer *string `json:"primaryServer,omitempty"`

	// SecondaryServer
	// Secondary server test result
	SecondaryServer *string `json:"secondaryServer,omitempty"`
}

func NewTestAAAServerResult() *TestAAAServerResult {
	testAAAServerResultType := new(TestAAAServerResult)
	return testAAAServerResultType
}

func NewTestAAAServerResultWithDefaults() *TestAAAServerResult {
	testAAAServerResultType := new(TestAAAServerResult)
	return testAAAServerResultType
}

type TestAuthenticationServer struct {
	// AaaServer
	// Constraints:
	//    - required
	AaaServer *common.GenericRef `json:"aaaServer" validate:"required"`

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

func NewTestAuthenticationServer() *TestAuthenticationServer {
	testAuthenticationServerType := new(TestAuthenticationServer)
	return testAuthenticationServerType
}

func NewTestAuthenticationServerWithDefaults() *TestAuthenticationServer {
	testAuthenticationServerType := new(TestAuthenticationServer)
	authProtocolField := `PAP`
	testAuthenticationServerType.AuthProtocol = &authProtocolField
	return testAuthenticationServerType
}
