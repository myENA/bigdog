package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSCGUserService struct {
	apiClient *APIClient
}

func NewWSGSCGUserService(c *APIClient) *WSGSCGUserService {
	s := new(WSGSCGUserService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCGUserService() *WSGSCGUserService {
	return NewWSGSCGUserService(ss.apiClient)
}

type WSGSCGUserCreateScgUser struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// DomainId
	// Domain id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	// Constraints:
	//    - nullable
	Email *string `json:"email,omitempty"`

	// Id
	// User id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// MinimumPasswordLength
	// The minimum length of the password for the account. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	// MinimumPasswordLifetimeEnabled
	// Enable the password should not be changed twice within the 24 hours.
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	// NewPassphrase
	// User login passphrase
	// Constraints:
	//    - required
	NewPassphrase *string `json:"newPassphrase" validate:"required"`

	// PasswordComplexityEnabled
	// Enable the password complexity, should apply the rules as: At least one upper-case character; At least one lower-case character; At least one numeric character:At least one special character; At least 8-chars within the old password should be changed.
	// Constraints:
	//    - nullable
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:365
	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s). (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// Phone
	// User phone
	// Constraints:
	//    - nullable
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	// Constraints:
	//    - nullable
	RealName *string `json:"realName,omitempty"`

	// SessionIdle
	// A period of idle used to invalid that session. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// Title
	// User title
	// Constraints:
	//    - nullable
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`
}

func NewWSGSCGUserCreateScgUser() *WSGSCGUserCreateScgUser {
	m := new(WSGSCGUserCreateScgUser)
	return m
}

type WSGSCGUserGetScgUser struct {
	// AccountLockout
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

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

	// DomainId
	// Domain id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	// Constraints:
	//    - nullable
	Email *string `json:"email,omitempty"`

	// Enabled
	// User enabled or not
	// Constraints:
	//    - nullable
	//    - oneof:[0,1]
	Enabled *int `json:"enabled,omitempty" validate:"omitempty,oneof=0 1"`

	// Id
	// User id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Locked
	// User locked or not (0:unlocked/1:locked)
	// Constraints:
	//    - nullable
	//    - oneof:[0,1,2,3]
	Locked *int `json:"locked,omitempty" validate:"omitempty,oneof=0 1 2 3"`

	// LockoutDuration
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// MinimumPasswordLength
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	// MinimumPasswordLifetimeEnabled
	// Enable the password should not be changed twice within the 24 hours.
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

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

	// PasswordComplexityEnabled
	// Enable the password complexity, should apply the rules as: At least one upper-case character; At least one lower-case character; At least one numeric character:At least one special character; At least 8-chars within the old password should be changed.
	// Constraints:
	//    - nullable
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:365
	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	// PasswordReuse
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// Phone
	// User phone
	// Constraints:
	//    - nullable
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	// Constraints:
	//    - nullable
	RealName *string `json:"realName,omitempty"`

	// SessionIdle
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// TenantUUID
	// Tenant id
	// Constraints:
	//    - nullable
	TenantUUID *string `json:"tenantUUID,omitempty"`

	// Title
	// User title
	// Constraints:
	//    - nullable
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`
}

func NewWSGSCGUserGetScgUser() *WSGSCGUserGetScgUser {
	m := new(WSGSCGUserGetScgUser)
	return m
}

type WSGSCGUserModifyScgUser struct {
	// AccountLockout
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// DomainId
	// Domain id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	// Constraints:
	//    - nullable
	Email *string `json:"email,omitempty"`

	// Id
	// User id
	// Constraints:
	//    - required
	Id *string `json:"id" validate:"required"`

	// LockoutDuration
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// MinimumPasswordLength
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	// MinimumPasswordLifetimeEnabled
	// Enable the password should not be changed twice within the 24 hours.
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	// NewPassphrase
	// User new login passphrase
	// Constraints:
	//    - nullable
	NewPassphrase *string `json:"newPassphrase,omitempty"`

	// Passphrase
	// User login passphrase
	// Constraints:
	//    - nullable
	Passphrase *string `json:"passphrase,omitempty"`

	// PasswordComplexityEnabled
	// Enable the password complexity, should apply the rules as: At least one upper-case character; At least one lower-case character; At least one numeric character:At least one special character; At least 8-chars within the old password should be changed.
	// Constraints:
	//    - nullable
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:365
	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	// PasswordReuse
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// Phone
	// User phone
	// Constraints:
	//    - nullable
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	// Constraints:
	//    - nullable
	RealName *string `json:"realName,omitempty"`

	// SessionIdle
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// Title
	// User title
	// Constraints:
	//    - nullable
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`
}

func NewWSGSCGUserModifyScgUser() *WSGSCGUserModifyScgUser {
	m := new(WSGSCGUserModifyScgUser)
	return m
}

type WSGSCGUserPatchScgUserGroup struct {
	// AccountSecurityProfileId
	// Constraints:
	//    - nullable
	AccountSecurityProfileId *string `json:"accountSecurityProfileId,omitempty"`

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

	// Description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

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

	// Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Permissions
	// Constraints:
	//    - nullable
	Permissions []*WSGSCGUserGroupPermissionWithoutDetailItems `json:"permissions,omitempty" validate:"omitempty,dive"`

	// ResourceGroups
	// Constraints:
	//    - nullable
	ResourceGroups []*WSGSCGUserGroupResourceGroup `json:"resourceGroups,omitempty" validate:"omitempty,dive"`

	// Role
	// Constraints:
	//    - nullable
	Role *string `json:"role,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Constraints:
	//    - required
	Users []*WSGSCGUserGetScgUser `json:"users" validate:"required,dive"`
}

func NewWSGSCGUserPatchScgUserGroup() *WSGSCGUserPatchScgUserGroup {
	m := new(WSGSCGUserPatchScgUserGroup)
	return m
}

type WSGSCGUserQueryCriteria struct {
	// Attributes
	// Get specific columns only
	// Constraints:
	//    - nullable
	Attributes []string `json:"attributes,omitempty" validate:"omitempty,dive"`

	// Criteria
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	// Constraints:
	//    - nullable
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraFilters []*WSGCommonQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty" validate:"omitempty,dive"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraNotFilters []*WSGCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty" validate:"omitempty,dive"`

	// ExtraTimeRange
	// Constraints:
	//    - nullable
	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	// Constraints:
	//    - nullable
	Filters []*WSGSCGUserQueryCriteriaFiltersType `json:"filters,omitempty" validate:"omitempty,dive"`

	// FullTextSearch
	// Constraints:
	//    - nullable
	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required information
	// Constraints:
	//    - nullable
	Options *WSGCommonQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - nullable
	//    - min:1
	Page *int `json:"page,omitempty" validate:"omitempty,gte=1"`

	// Query
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	// Constraints:
	//    - nullable
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGSCGUserQueryCriteria() *WSGSCGUserQueryCriteria {
	m := new(WSGSCGUserQueryCriteria)
	return m
}

type WSGSCGUserQueryCriteriaFiltersType struct {
	// Operator
	// operator
	// Constraints:
	//    - nullable
	//    - oneof:[eq]
	Operator *string `json:"operator,omitempty" validate:"omitempty,oneof=eq"`

	// Type
	// Group type
	// Constraints:
	//    - required
	//    - oneof:[DOMAIN]
	Type *string `json:"type" validate:"required,oneof=DOMAIN"`

	// Value
	// DOMAIN ID
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`
}

func NewWSGSCGUserQueryCriteriaFiltersType() *WSGSCGUserQueryCriteriaFiltersType {
	m := new(WSGSCGUserQueryCriteriaFiltersType)
	return m
}

type WSGSCGUserAuditId struct {
	// Id
	// the identifier of the SCG user
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`
}

func NewWSGSCGUserAuditId() *WSGSCGUserAuditId {
	m := new(WSGSCGUserAuditId)
	return m
}

type WSGSCGUserGroup struct {
	// AccountSecurityProfileId
	// Constraints:
	//    - required
	AccountSecurityProfileId *string `json:"accountSecurityProfileId" validate:"required"`

	// AccountSecurityProfileName
	// Constraints:
	//    - nullable
	AccountSecurityProfileName *string `json:"accountSecurityProfileName,omitempty"`

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
	// User group description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// User group Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IsFactoryDefault
	// Constraints:
	//    - nullable
	IsFactoryDefault *bool `json:"isFactoryDefault,omitempty"`

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

	// Name
	// User group name
	// Constraints:
	//    - required
	Name *string `json:"name" validate:"required"`

	// Permissions
	// Permission list
	// Constraints:
	//    - required
	Permissions []*WSGSCGUserGroupPermissionWithoutDetailItems `json:"permissions" validate:"required,dive"`

	// ResourceGroups
	// Resource group id list
	// Constraints:
	//    - required
	ResourceGroups []*WSGSCGUserGroupResourceGroup `json:"resourceGroups" validate:"required,dive"`

	// Role
	// User group role
	// Constraints:
	//    - required
	Role *string `json:"role" validate:"required"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Users in this user group
	// Constraints:
	//    - required
	Users []*WSGSCGUserGetScgUser `json:"users" validate:"required,dive"`
}

func NewWSGSCGUserGroup() *WSGSCGUserGroup {
	m := new(WSGSCGUserGroup)
	return m
}

type WSGSCGUserGroupAuditId struct {
	// Id
	// the identifier of the SCG user group
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`
}

func NewWSGSCGUserGroupAuditId() *WSGSCGUserGroupAuditId {
	m := new(WSGSCGUserGroupAuditId)
	return m
}

type WSGSCGUserGroupList struct {
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
	List []*WSGSCGUserGroup `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCGUserGroupList() *WSGSCGUserGroupList {
	m := new(WSGSCGUserGroupList)
	return m
}

type WSGSCGUserGroupPermission struct {
	// Access
	// Access level
	// Constraints:
	//    - nullable
	//    - oneof:[READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=READ MODIFY FULL_ACCESS"`

	// Display
	// Resource display name
	// Constraints:
	//    - nullable
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	// Constraints:
	//    - nullable
	Ids []string `json:"ids,omitempty" validate:"omitempty,dive"`

	// Items
	// Resource items
	// Constraints:
	//    - nullable
	Items []*WSGSCGUserGroupPermissionItemsType `json:"items,omitempty" validate:"omitempty,dive"`

	// ItemsDescription
	// Descriptions of Resource items
	// Constraints:
	//    - nullable
	ItemsDescription []string `json:"itemsDescription,omitempty" validate:"omitempty,dive"`

	// Resource
	// Resource type
	// Constraints:
	//    - nullable
	Resource *string `json:"resource,omitempty"`
}

func NewWSGSCGUserGroupPermission() *WSGSCGUserGroupPermission {
	m := new(WSGSCGUserGroupPermission)
	return m
}

type WSGSCGUserGroupPermissionItemsType struct {
	// Access
	// Constraints:
	//    - nullable
	//    - oneof:[NA,READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=NA READ MODIFY FULL_ACCESS"`

	// Display
	// Constraints:
	//    - nullable
	Display *string `json:"display,omitempty"`

	// Resource
	// Constraints:
	//    - nullable
	Resource *string `json:"resource,omitempty"`
}

func NewWSGSCGUserGroupPermissionItemsType() *WSGSCGUserGroupPermissionItemsType {
	m := new(WSGSCGUserGroupPermissionItemsType)
	return m
}

type WSGSCGUserGroupPermissionList struct {
	// Extra
	// Any additional response data.
	// Constraints:
	//    - nullable
	Extra *WSGSCGUserGroupPermissionListExtraType `json:"extra,omitempty"`

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
	List []*WSGSCGUserGroupPermission `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCGUserGroupPermissionList() *WSGSCGUserGroupPermissionList {
	m := new(WSGSCGUserGroupPermissionList)
	return m
}

// WSGSCGUserGroupPermissionListExtraType
//
// Any additional response data.
// Constraints:
//    - nullable
type WSGSCGUserGroupPermissionListExtraType struct {
	// IsSuperAdmin
	// whether or not current user is a 'Super Admin' that possesses all 6 permission categories with 'FULL_ACCESS'
	// Constraints:
	//    - nullable
	IsSuperAdmin *bool `json:"isSuperAdmin,omitempty"`

	// IsSuperAdminOfDomain
	// whether or not current user is a 'Super Admin of Partner Domain' that possesses all 6 permission categories with 'FULL_ACCESS'
	// Constraints:
	//    - nullable
	IsSuperAdminOfDomain *bool `json:"isSuperAdminOfDomain,omitempty"`
}

func NewWSGSCGUserGroupPermissionListExtraType() *WSGSCGUserGroupPermissionListExtraType {
	m := new(WSGSCGUserGroupPermissionListExtraType)
	return m
}

type WSGSCGUserGroupPermissionWithoutDetailItems struct {
	// Access
	// Access level
	// Constraints:
	//    - nullable
	//    - oneof:[READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=READ MODIFY FULL_ACCESS"`

	// Display
	// Resource display name
	// Constraints:
	//    - nullable
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	// Constraints:
	//    - nullable
	Ids []string `json:"ids,omitempty" validate:"omitempty,dive"`

	// Resource
	// Resource type
	// Constraints:
	//    - nullable
	Resource *string `json:"resource,omitempty"`
}

func NewWSGSCGUserGroupPermissionWithoutDetailItems() *WSGSCGUserGroupPermissionWithoutDetailItems {
	m := new(WSGSCGUserGroupPermissionWithoutDetailItems)
	return m
}

type WSGSCGUserGroupResourceGroup struct {
	// Id
	// the identifier of the SCG resource group
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the SCG resource group
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Type
	// the type of SCG resource group
	// Constraints:
	//    - nullable
	//    - oneof:[DOMAIN,ZONE,APGROUP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=DOMAIN ZONE APGROUP"`
}

func NewWSGSCGUserGroupResourceGroup() *WSGSCGUserGroupResourceGroup {
	m := new(WSGSCGUserGroupResourceGroup)
	return m
}

type WSGSCGUserGroupRoleLabelValue struct {
	// Label
	// Role display name
	// Constraints:
	//    - nullable
	Label *string `json:"label,omitempty"`

	// Value
	// Role value
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewWSGSCGUserGroupRoleLabelValue() *WSGSCGUserGroupRoleLabelValue {
	m := new(WSGSCGUserGroupRoleLabelValue)
	return m
}

type WSGSCGUserGroupRoleLabelValueList struct {
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
	List []*WSGSCGUserGroupRoleLabelValue `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCGUserGroupRoleLabelValueList() *WSGSCGUserGroupRoleLabelValueList {
	m := new(WSGSCGUserGroupRoleLabelValueList)
	return m
}

type WSGSCGUserList struct {
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
	List []*WSGSCGUserGetScgUser `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCGUserList() *WSGSCGUserList {
	m := new(WSGSCGUserList)
	return m
}

// AddUsers
//
// Add SCG user.
//
// Request Body:
//	 - body *WSGSCGUserCreateScgUser
func (s *WSGSCGUserService) AddUsers(ctx context.Context, body *WSGSCGUserCreateScgUser) (*WSGSCGUserAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddUsers, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteUsers
//
// Delete multiple SCG user.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserService) DeleteUsers(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUsers, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// DeleteUsersByUserId
//
// Delete SCG user.
//
// Required Parameters:
// - userId string
//		- required
func (s *WSGSCGUserService) DeleteUsersByUserId(ctx context.Context, userId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, userId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUsersByUserId, true)
	req.SetPathParameter("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindUsersByQueryCriteria
//
// Query SCG users.
//
// Request Body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserService) FindUsersByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria) (*WSGSCGUserList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindUsersByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUsersByUserId
//
// Get SCG user.
//
// Required Parameters:
// - userId string
//		- required
func (s *WSGSCGUserService) FindUsersByUserId(ctx context.Context, userId string) (*WSGSCGUserGetScgUser, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserGetScgUser
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, userId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUsersByUserId, true)
	req.SetPathParameter("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserGetScgUser()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateUsersByUserId
//
// Update SCG user.
//
// Request Body:
//	 - body *WSGSCGUserModifyScgUser
//
// Required Parameters:
// - userId string
//		- required
func (s *WSGSCGUserService) PartialUpdateUsersByUserId(ctx context.Context, body *WSGSCGUserModifyScgUser, userId string) (*WSGSCGUserAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCGUserAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, userId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateUsersByUserId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSCGUserAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
