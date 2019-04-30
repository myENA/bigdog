package aaa

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/common"
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

	// ID
	// Identifier of the active directory server
	ID *string `json:"id,omitempty"`

	// IP
	// IP address
	IP *string `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

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

	// StandbyIP
	// IP address - Standby Cluster settings
	StandbyIP *string `json:"standbyIp,omitempty"`

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

	// ZoneID
	// Identifier of the zone which the active directory server belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type ActiveDirectoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ActiveDirectory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type AuthenticationServer struct {
	// Description
	// Description of the RADIUS server
	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the RADIUS server
	ID *string `json:"id,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

	// Name
	// Name of the RADIUS server
	Name *string `json:"name,omitempty"`

	Primary *common.RadiusServer `json:"primary,omitempty"`

	Secondary *common.RadiusServer `json:"secondary,omitempty"`

	StandbyPrimary *common.RadiusServer `json:"standbyPrimary,omitempty"`

	// StandbyServerEnabled
	// StandbyCluster different AAA Settings Enabled
	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	// ZoneID
	// Identifier of the zone which the RADIUS server belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type AuthenticationServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AuthenticationServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CreateActiveDirectoryServer struct {
	AdminDomainName *common.NormalNameTo64 `json:"adminDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	IP *common.IPAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

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

	// StandbyIP
	// IP address - Standby Cluster settings
	StandbyIP *string `json:"standbyIp,omitempty"`

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

type CreateAuthenticationServer struct {
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

type CreateLDAPServer struct {
	AdminDomainName *common.NormalNameTo128 `json:"adminDomainName,omitempty"`

	BaseDomainName *common.NormalNameTo64 `json:"baseDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	IP *common.IPAddress `json:"ip,omitempty"`

	KeyAttribute *common.NormalNameTo64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	SearchFilter *common.NormalNameTo64 `json:"searchFilter,omitempty"`

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Base domain name - Standby Cluster settings
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIP
	// IP address - Standby Cluster settings
	StandbyIP *string `json:"standbyIp,omitempty"`

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

type DeleteBulkAAAServerList struct {
	AaaDeleteBulkAAAServerList *string `json:"aaa_deleteBulkAAAServerList,omitempty"`
}

// GroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type GroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	GroupAttr *string `json:"groupAttr,omitempty"`

	// ID
	// Group attribute mapping UUID
	ID *string `json:"id,omitempty"`

	// UserRole
	// Identity user role
	UserRole *GroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole,omitempty"`
}

// GroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type GroupAttrIdentityUserRoleMappingUserRoleType struct {
	// ID
	// Identity user role UUID
	ID *string `json:"id,omitempty"`

	Name *common.NormalNameTo64 `json:"name,omitempty"`

	// UserTrafficProfile
	// Identity user role
	UserTrafficProfile *GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

// GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType
//
// Identity user role
type GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	// ID
	// User traffic profile UUID
	ID *string `json:"id,omitempty"`

	// Name
	// User traffic profile name
	Name *string `json:"name,omitempty"`
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

	// ID
	// Identifier of the LDAP server
	ID *string `json:"id,omitempty"`

	// IP
	// IP address
	IP *string `json:"ip,omitempty"`

	// KeyAttribute
	// Key attribute
	KeyAttribute *string `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	// MvnoID
	// Tenant UUID
	MvnoID *string `json:"mvnoId,omitempty"`

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

	// StandbyIP
	// IP address - Standby Cluster settings
	StandbyIP *string `json:"standbyIp,omitempty"`

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

	// ZoneID
	// Identifier of the zone which the LDAP server belongs to
	ZoneID *string `json:"zoneId,omitempty"`
}

type LDAPServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LDAPServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ModifyActiveDirectoryServer struct {
	AdminDomainName *common.NormalNameTo64 `json:"adminDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// GlobalCatalogEnabled
	// Enable global catalog support
	GlobalCatalogEnabled *bool `json:"globalCatalogEnabled,omitempty"`

	IP *common.IPAddress `json:"ip,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

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

	// StandbyIP
	// IP address - Standby Cluster settings
	StandbyIP *string `json:"standbyIp,omitempty"`

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

// ModifyGroupAttrIdentityUserRoleMapping
//
// User traffic profile mapping
type ModifyGroupAttrIdentityUserRoleMapping struct {
	// GroupAttr
	// Group attribute
	GroupAttr *string `json:"groupAttr,omitempty"`

	// UserRole
	// Identity user role
	UserRole *ModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole,omitempty"`
}

// ModifyGroupAttrIdentityUserRoleMappingUserRoleType
//
// Identity user role
type ModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	// ID
	// Identity user role UUID
	ID *string `json:"id,omitempty"`

	Name *common.NormalNameTo64 `json:"name,omitempty"`
}

type ModifyLDAPServer struct {
	AdminDomainName *common.NormalNameTo128 `json:"adminDomainName,omitempty"`

	BaseDomainName *common.NormalNameTo64 `json:"baseDomainName,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	IP *common.IPAddress `json:"ip,omitempty"`

	KeyAttribute *common.NormalNameTo64 `json:"keyAttribute,omitempty"`

	// Mappings
	// Group attribute and user traffic profile mapping
	Mappings []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Password
	// Admin password
	Password *string `json:"password,omitempty"`

	// Port
	// Port
	Port *int `json:"port,omitempty"`

	SearchFilter *common.NormalNameTo64 `json:"searchFilter,omitempty"`

	// StandbyAdminDomainName
	// Admin domain name - Standby Cluster settings
	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	// StandbyBaseDomainName
	// Base domain name - Standby Cluster settings
	StandbyBaseDomainName *string `json:"standbyBaseDomainName,omitempty"`

	// StandbyIP
	// IP address - Standby Cluster settings
	StandbyIP *string `json:"standbyIp,omitempty"`

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

type TestAAAServerResult struct {
	// PrimaryServer
	// Primary server test result
	PrimaryServer *string `json:"primaryServer,omitempty"`

	// SecondaryServer
	// Secondary server test result
	SecondaryServer *string `json:"secondaryServer,omitempty"`
}

type TestAuthenticationServer struct {
	AaaServer *common.GenericRef `json:"aaaServer,omitempty"`

	// AaaType
	// Authentication/Accounting service protocol. RADIUS for Radius, AD and LDAP. RADIUSAcct for RADIUS
	// Accounting
	AaaType *string `json:"aaaType,omitempty"`

	// AuthProtocol
	// Authentication protocol
	AuthProtocol *string `json:"authProtocol,omitempty"`

	// Password
	// Password
	Password *string `json:"password,omitempty"`

	// ServerType
	// Radius server type.
	ServerType *string `json:"serverType,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}
