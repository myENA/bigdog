package aaa

// API Version: v8_0

type ActiveDirectory struct {
	AdminDomainName             *string                             `json:"adminDomainName,omitempty"`
	Description                 *string                             `json:"description,omitempty"`
	GlobalCatalogEnabled        *bool                               `json:"globalCatalogEnabled,omitempty"`
	ID                          *string                             `json:"id,omitempty"`
	IP                          *string                             `json:"ip,omitempty"`
	Mappings                    []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	MvnoID                      *string                             `json:"mvnoId,omitempty"`
	Name                        *string                             `json:"name,omitempty"`
	Password                    *string                             `json:"password,omitempty"`
	Port                        *int                                `json:"port,omitempty"`
	StandbyAdminDomainName      *string                             `json:"standbyAdminDomainName,omitempty"`
	StandbyGlobalCatalogEnabled *bool                               `json:"standbyGlobalCatalogEnabled,omitempty"`
	StandbyIP                   *string                             `json:"standbyIp,omitempty"`
	StandbyPassword             *string                             `json:"standbyPassword,omitempty"`
	StandbyPort                 *int                                `json:"standbyPort,omitempty"`
	StandbyServerEnabled        *bool                               `json:"standbyServerEnabled,omitempty"`
	StandbyWindowsDomainName    *string                             `json:"standbyWindowsDomainName,omitempty"`
	WindowsDomainName           *string                             `json:"windowsDomainName,omitempty"`
	ZoneID                      *string                             `json:"zoneId,omitempty"`
}

type ActiveDirectoryList struct {
	FirstIndex *int               `json:"firstIndex,omitempty"`
	HasMore    *bool              `json:"hasMore,omitempty"`
	List       []*ActiveDirectory `json:"list,omitempty"`
	TotalCount *int               `json:"totalCount,omitempty"`
}

type AuthenticationServer struct {
	Description          *string                             `json:"description,omitempty"`
	ID                   *string                             `json:"id,omitempty"`
	Mappings             []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	MvnoID               *string                             `json:"mvnoId,omitempty"`
	Name                 *string                             `json:"name,omitempty"`
	Primary              *common.RadiusServer                `json:"primary,omitempty"`
	Secondary            *common.RadiusServer                `json:"secondary,omitempty"`
	StandbyPrimary       *common.RadiusServer                `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled *bool                               `json:"standbyServerEnabled,omitempty"`
	ZoneID               *string                             `json:"zoneId,omitempty"`
}

type AuthenticationServerList struct {
	FirstIndex *int                    `json:"firstIndex,omitempty"`
	HasMore    *bool                   `json:"hasMore,omitempty"`
	List       []*AuthenticationServer `json:"list,omitempty"`
	TotalCount *int                    `json:"totalCount,omitempty"`
}

type CreateActiveDirectoryServer struct {
	AdminDomainName             *string                                   `json:"adminDomainName,omitempty"`
	Description                 *string                                   `json:"description,omitempty"`
	GlobalCatalogEnabled        *bool                                     `json:"globalCatalogEnabled,omitempty"`
	IP                          *string                                   `json:"ip,omitempty"`
	Mappings                    []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	Name                        *string                                   `json:"name,omitempty"`
	Password                    *string                                   `json:"password,omitempty"`
	Port                        *int                                      `json:"port,omitempty"`
	StandbyAdminDomainName      *string                                   `json:"standbyAdminDomainName,omitempty"`
	StandbyGlobalCatalogEnabled *bool                                     `json:"standbyGlobalCatalogEnabled,omitempty"`
	StandbyIP                   *string                                   `json:"standbyIp,omitempty"`
	StandbyPassword             *string                                   `json:"standbyPassword,omitempty"`
	StandbyPort                 *int                                      `json:"standbyPort,omitempty"`
	StandbyServerEnabled        *bool                                     `json:"standbyServerEnabled,omitempty"`
	StandbyWindowsDomainName    *string                                   `json:"standbyWindowsDomainName,omitempty"`
	WindowsDomainName           *string                                   `json:"windowsDomainName,omitempty"`
}

type CreateAuthenticationServer struct {
	Description          *string                                   `json:"description,omitempty"`
	Mappings             []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	Name                 *string                                   `json:"name,omitempty"`
	Primary              *common.RadiusServer                      `json:"primary,omitempty"`
	Secondary            *common.RadiusServer                      `json:"secondary,omitempty"`
	StandbyPrimary       *common.RadiusServer                      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled *bool                                     `json:"standbyServerEnabled,omitempty"`
}

type CreateLDAPServer struct {
	AdminDomainName        *string                                   `json:"adminDomainName,omitempty"`
	BaseDomainName         *string                                   `json:"baseDomainName,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	IP                     *string                                   `json:"ip,omitempty"`
	KeyAttribute           *string                                   `json:"keyAttribute,omitempty"`
	Mappings               []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	Name                   *string                                   `json:"name,omitempty"`
	Password               *string                                   `json:"password,omitempty"`
	Port                   *int                                      `json:"port,omitempty"`
	SearchFilter           *string                                   `json:"searchFilter,omitempty"`
	StandbyAdminDomainName *string                                   `json:"standbyAdminDomainName,omitempty"`
	StandbyBaseDomainName  *string                                   `json:"standbyBaseDomainName,omitempty"`
	StandbyIP              *string                                   `json:"standbyIp,omitempty"`
	StandbyKeyAttribute    *string                                   `json:"standbyKeyAttribute,omitempty"`
	StandbyPassword        *string                                   `json:"standbyPassword,omitempty"`
	StandbyPort            *int                                      `json:"standbyPort,omitempty"`
	StandbySearchFilter    *string                                   `json:"standbySearchFilter,omitempty"`
	StandbyServerEnabled   *bool                                     `json:"standbyServerEnabled,omitempty"`
}

type DeleteBulkAAAServerList struct {
	AaaDeleteBulkAAAServerList *string `json:"aaa_deleteBulkAAAServerList,omitempty"`
}

type GroupAttrIdentityUserRoleMapping struct {
	GroupAttr *string                                       `json:"groupAttr,omitempty"`
	ID        *string                                       `json:"id,omitempty"`
	UserRole  *GroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole,omitempty"`
}

type GroupAttrIdentityUserRoleMappingUserRoleType struct {
	ID                 *string                                                             `json:"id,omitempty"`
	Name               *string                                                             `json:"name,omitempty"`
	UserTrafficProfile *GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType `json:"userTrafficProfile,omitempty"`
}

type GroupAttrIdentityUserRoleMappingUserRoleTypeUserTrafficProfileType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type LDAPServer struct {
	AdminDomainName        *string                             `json:"adminDomainName,omitempty"`
	BaseDomainName         *string                             `json:"baseDomainName,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	ID                     *string                             `json:"id,omitempty"`
	IP                     *string                             `json:"ip,omitempty"`
	KeyAttribute           *string                             `json:"keyAttribute,omitempty"`
	Mappings               []*GroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	MvnoID                 *string                             `json:"mvnoId,omitempty"`
	Name                   *string                             `json:"name,omitempty"`
	Password               *string                             `json:"password,omitempty"`
	Port                   *int                                `json:"port,omitempty"`
	SearchFilter           *string                             `json:"searchFilter,omitempty"`
	StandbyAdminDomainName *string                             `json:"standbyAdminDomainName,omitempty"`
	StandbyBaseDomainName  *string                             `json:"standbyBaseDomainName,omitempty"`
	StandbyIP              *string                             `json:"standbyIp,omitempty"`
	StandbyKeyAttribute    *string                             `json:"standbyKeyAttribute,omitempty"`
	StandbyPassword        *string                             `json:"standbyPassword,omitempty"`
	StandbyPort            *int                                `json:"standbyPort,omitempty"`
	StandbySearchFilter    *string                             `json:"standbySearchFilter,omitempty"`
	StandbyServerEnabled   *bool                               `json:"standbyServerEnabled,omitempty"`
	ZoneID                 *string                             `json:"zoneId,omitempty"`
}

type LDAPServerList struct {
	FirstIndex *int          `json:"firstIndex,omitempty"`
	HasMore    *bool         `json:"hasMore,omitempty"`
	List       []*LDAPServer `json:"list,omitempty"`
	TotalCount *int          `json:"totalCount,omitempty"`
}

type ModifyActiveDirectoryServer struct {
	AdminDomainName             *string                                   `json:"adminDomainName,omitempty"`
	Description                 *string                                   `json:"description,omitempty"`
	GlobalCatalogEnabled        *bool                                     `json:"globalCatalogEnabled,omitempty"`
	IP                          *string                                   `json:"ip,omitempty"`
	Mappings                    []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	Name                        *string                                   `json:"name,omitempty"`
	Password                    *string                                   `json:"password,omitempty"`
	Port                        *int                                      `json:"port,omitempty"`
	StandbyAdminDomainName      *string                                   `json:"standbyAdminDomainName,omitempty"`
	StandbyGlobalCatalogEnabled *bool                                     `json:"standbyGlobalCatalogEnabled,omitempty"`
	StandbyIP                   *string                                   `json:"standbyIp,omitempty"`
	StandbyPassword             *string                                   `json:"standbyPassword,omitempty"`
	StandbyPort                 *int                                      `json:"standbyPort,omitempty"`
	StandbyServerEnabled        *bool                                     `json:"standbyServerEnabled,omitempty"`
	StandbyWindowsDomainName    *string                                   `json:"standbyWindowsDomainName,omitempty"`
	WindowsDomainName           *string                                   `json:"windowsDomainName,omitempty"`
}

type ModifyAuthenticationServer struct {
	Description          *string                                   `json:"description,omitempty"`
	Mappings             []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	Name                 *string                                   `json:"name,omitempty"`
	Primary              *common.RadiusServer                      `json:"primary,omitempty"`
	Secondary            *common.RadiusServer                      `json:"secondary,omitempty"`
	StandbyPrimary       *common.RadiusServer                      `json:"standbyPrimary,omitempty"`
	StandbyServerEnabled *bool                                     `json:"standbyServerEnabled,omitempty"`
}

type ModifyGroupAttrIdentityUserRoleMapping struct {
	GroupAttr *string                                             `json:"groupAttr,omitempty"`
	UserRole  *ModifyGroupAttrIdentityUserRoleMappingUserRoleType `json:"userRole,omitempty"`
}

type ModifyGroupAttrIdentityUserRoleMappingUserRoleType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ModifyLDAPServer struct {
	AdminDomainName        *string                                   `json:"adminDomainName,omitempty"`
	BaseDomainName         *string                                   `json:"baseDomainName,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	IP                     *string                                   `json:"ip,omitempty"`
	KeyAttribute           *string                                   `json:"keyAttribute,omitempty"`
	Mappings               []*ModifyGroupAttrIdentityUserRoleMapping `json:"mappings,omitempty"`
	Name                   *string                                   `json:"name,omitempty"`
	Password               *string                                   `json:"password,omitempty"`
	Port                   *int                                      `json:"port,omitempty"`
	SearchFilter           *string                                   `json:"searchFilter,omitempty"`
	StandbyAdminDomainName *string                                   `json:"standbyAdminDomainName,omitempty"`
	StandbyBaseDomainName  *string                                   `json:"standbyBaseDomainName,omitempty"`
	StandbyIP              *string                                   `json:"standbyIp,omitempty"`
	StandbyKeyAttribute    *string                                   `json:"standbyKeyAttribute,omitempty"`
	StandbyPassword        *string                                   `json:"standbyPassword,omitempty"`
	StandbyPort            *int                                      `json:"standbyPort,omitempty"`
	StandbySearchFilter    *string                                   `json:"standbySearchFilter,omitempty"`
	StandbyServerEnabled   *bool                                     `json:"standbyServerEnabled,omitempty"`
}

type TestAAAServerResult struct {
	PrimaryServer   *string `json:"primaryServer,omitempty"`
	SecondaryServer *string `json:"secondaryServer,omitempty"`
}

type TestAuthenticationServer struct {
	AaaServer    *common.GenericRef `json:"aaaServer,omitempty"`
	AaaType      *string            `json:"aaaType,omitempty"`
	AuthProtocol *string            `json:"authProtocol,omitempty"`
	Password     *string            `json:"password,omitempty"`
	ServerType   *string            `json:"serverType,omitempty"`
	UserName     *string            `json:"userName,omitempty"`
}
