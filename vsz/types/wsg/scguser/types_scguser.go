package scguser

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type PatchScgUserGroup struct {
	AccountSecurityProfileId *string `json:"accountSecurityProfileId,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	Permissions []*ScgUserGroupPermissionWithoutDetailItems `json:"permissions,omitempty"`

	ResourceGroups []*ScgUserGroupResourceGroup `json:"resourceGroups,omitempty"`

	Role *string `json:"role,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	Users []*ScgUser `json:"users,omitempty"`
}

type QueryCriteria struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []*common.QueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*common.QueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *common.TimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*QueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *common.FullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options *common.QueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *common.QueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

type QueryCriteriaFiltersType struct {
	// Operator
	// operator
	Operator *string `json:"operator,omitempty"`

	// Type
	// Group type
	Type *string `json:"type,omitempty"`

	// Value
	// DOMAIN ID
	Value *string `json:"value,omitempty"`
}

type ScgUser struct {
	// CompanyName
	// User company name
	CompanyName *string `json:"companyName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// Enabled
	// User enabled or not
	Enabled *int `json:"enabled,omitempty"`

	// Id
	// User id
	Id *string `json:"id,omitempty"`

	// Locked
	// User locked or not (0:unlocked/1:locked)
	Locked *int `json:"locked,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Passphrase
	// User login passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	PasswordReuse *int `json:"passwordReuse,omitempty"`

	// Phone
	// User phone
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty"`

	// TenantUUID
	// Tenant id
	TenantUUID *string `json:"tenantUUID,omitempty"`

	// Title
	// User title
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}

type ScgUserAuditId struct {
	// Id
	// the identifier of the SCG user
	Id *string `json:"id,omitempty"`
}

type ScgUserGroup struct {
	AccountSecurityProfileId *string `json:"accountSecurityProfileId,omitempty"`

	AccountSecurityProfileName *string `json:"accountSecurityProfileName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// User group description
	Description *string `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// User group Id
	Id *string `json:"id,omitempty"`

	IsFactoryDefault *bool `json:"isFactoryDefault,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// User group name
	Name *string `json:"name,omitempty"`

	// Permissions
	// Permission list
	Permissions []*ScgUserGroupPermissionWithoutDetailItems `json:"permissions,omitempty"`

	// ResourceGroups
	// Resource group id list
	ResourceGroups []*ScgUserGroupResourceGroup `json:"resourceGroups,omitempty"`

	// Role
	// User group role
	Role *string `json:"role,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Users in this user group
	Users []*ScgUser `json:"users,omitempty"`
}

type ScgUserGroupAuditId struct {
	// Id
	// the identifier of the SCG user group
	Id *string `json:"id,omitempty"`
}

type ScgUserGroupList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ScgUserGroupPermission struct {
	// Access
	// Access level
	Access *string `json:"access,omitempty"`

	// Display
	// Resource display name
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	Ids []string `json:"ids,omitempty"`

	// Items
	// Resource items
	Items []*ScgUserGroupPermissionItemsType `json:"items,omitempty"`

	// ItemsDescription
	// Descriptions of Resource items
	ItemsDescription []string `json:"itemsDescription,omitempty"`

	// Resource
	// Resource type
	Resource *string `json:"resource,omitempty"`
}

type ScgUserGroupPermissionItemsType struct {
	Access *string `json:"access,omitempty"`

	Display *string `json:"display,omitempty"`

	Resource *string `json:"resource,omitempty"`
}

type ScgUserGroupPermissionList struct {
	// Extra
	// Any additional response data.
	Extra *ScgUserGroupPermissionListExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroupPermission `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// ScgUserGroupPermissionListExtraType
//
// Any additional response data.
type ScgUserGroupPermissionListExtraType struct {
	// IsSuperAdmin
	// whether or not current user is a 'Super Admin' that possesses all 6 permission categories with
	// 'FULL_ACCESS'
	IsSuperAdmin *bool `json:"isSuperAdmin,omitempty"`

	// IsSuperAdminOfDomain
	// whether or not current user is a 'Super Admin of Partner Domain' that possesses all 6 permission
	// categories with 'FULL_ACCESS'
	IsSuperAdminOfDomain *bool `json:"isSuperAdminOfDomain,omitempty"`
}

type ScgUserGroupPermissionWithoutDetailItems struct {
	// Access
	// Access level
	Access *string `json:"access,omitempty"`

	// Display
	// Resource display name
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	Ids []string `json:"ids,omitempty"`

	// Resource
	// Resource type
	Resource *string `json:"resource,omitempty"`
}

type ScgUserGroupResourceGroup struct {
	// Id
	// the identifier of the SCG resource group
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the SCG resource group
	Name *string `json:"name,omitempty"`

	// Type
	// the type of SCG resource group
	Type *string `json:"type,omitempty"`
}

type ScgUserGroupRoleLabelValue struct {
	// Label
	// Role display name
	Label *string `json:"label,omitempty"`

	// Value
	// Role value
	Value *string `json:"value,omitempty"`
}

type ScgUserGroupRoleLabelValueList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroupRoleLabelValue `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ScgUserList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUser `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
