package scguser

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateSCGUser struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur. (System default admin ONLY)
	AccountLockout *int `json:"accountLockout,omitempty" validate:"gte=1,lte=100"`

	// DomainID
	// Domain id
	DomainID *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// ID
	// User id
	ID *string `json:"id,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention. (System
	// default admin ONLY)
	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"gte=1,lte=1440"`

	// MinimumPasswordLength
	// The minimum length of the password for the account. (System default admin ONLY)
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"gte=8,lte=64"`

	// NewPassphrase
	// User login passphrase
	NewPassphrase *string `json:"newPassphrase,omitempty" validate:"required"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly. (System default admin
	// ONLY)
	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"gte=1,lte=365"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s). (System default admin ONLY)
	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"gte=1,lte=6"`

	// Phone
	// User phone
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	// SessionIdle
	// A period of idle used to invalid that session. (System default admin ONLY)
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"gte=1,lte=1440"`

	// Title
	// User title
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty" validate:"required"`
}

type GetSCGUser struct {
	AccountLockout *int `json:"accountLockout,omitempty" validate:"gte=1,lte=100"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DomainID
	// Domain id
	DomainID *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// Enabled
	// User enabled or not
	Enabled *int `json:"enabled,omitempty" validate:"oneof=0 1"`

	// ID
	// User id
	ID *string `json:"id,omitempty"`

	// Locked
	// User locked or not (0:unlocked/1:locked)
	Locked *int `json:"locked,omitempty" validate:"oneof=0 1 2"`

	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"gte=1,lte=1440"`

	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"gte=8,lte=64"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"gte=1,lte=365"`

	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"gte=1,lte=6"`

	// Phone
	// User phone
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty" validate:"gte=1,lte=1440"`

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

type ModifySCGUser struct {
	AccountLockout *int `json:"accountLockout,omitempty" validate:"gte=1,lte=100"`

	// DomainID
	// Domain id
	DomainID *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// ID
	// User id
	ID *string `json:"id,omitempty" validate:"required"`

	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"gte=1,lte=1440"`

	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"gte=8,lte=64"`

	// NewPassphrase
	// User new login passphrase
	NewPassphrase *string `json:"newPassphrase,omitempty"`

	// Passphrase
	// User login passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"gte=1,lte=365"`

	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"gte=1,lte=6"`

	// Phone
	// User phone
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty" validate:"gte=1,lte=1440"`

	// Title
	// User title
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}

type PatchSCGUserGroup struct {
	AccountSecurityProfileID *string `json:"accountSecurityProfileId,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorID *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	ID *string `json:"id,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierID *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	Permissions []*SCGUserGroupPermissionWithoutDetailItems `json:"permissions,omitempty"`

	ResourceGroups []*SCGUserGroupResourceGroup `json:"resourceGroups,omitempty"`

	Role *string `json:"role,omitempty"`

	TenantID *string `json:"tenantId,omitempty"`

	Users []*GetSCGUser `json:"users,omitempty"`
}

type QueryCriteria struct{}

type SCGUserAuditID struct {
	// ID
	// the identifier of the SCG user
	ID *string `json:"id,omitempty"`
}

type SCGUserGroup struct {
	AccountSecurityProfileID *string `json:"accountSecurityProfileId,omitempty" validate:"required"`

	AccountSecurityProfileName *string `json:"accountSecurityProfileName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// User group description
	Description *string `json:"description,omitempty"`

	// DomainID
	// Domain Id
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// User group Id
	ID *string `json:"id,omitempty"`

	IsFactoryDefault *bool `json:"isFactoryDefault,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// User group name
	Name *string `json:"name,omitempty" validate:"required"`

	// Permissions
	// Permission list
	Permissions []*SCGUserGroupPermissionWithoutDetailItems `json:"permissions,omitempty" validate:"required"`

	// ResourceGroups
	// Resource group id list
	ResourceGroups []*SCGUserGroupResourceGroup `json:"resourceGroups,omitempty" validate:"required"`

	// Role
	// User group role
	Role *string `json:"role,omitempty" validate:"required"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`

	// Users
	// Users in this user group
	Users []*GetSCGUser `json:"users,omitempty"`
}

type SCGUserGroupAuditID struct {
	// ID
	// the identifier of the SCG user group
	ID *string `json:"id,omitempty"`
}

type SCGUserGroupList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SCGUserGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SCGUserGroupPermission struct {
	// Access
	// Access level
	Access *string `json:"access,omitempty" validate:"oneof=READ MODIFY FULL_ACCESS"`

	// Display
	// Resource display name
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	Ids []string `json:"ids,omitempty"`

	// Items
	// Resource items
	Items []*SCGUserGroupPermissionItemsType `json:"items,omitempty"`

	// ItemsDescription
	// Descriptions of Resource items
	ItemsDescription []string `json:"itemsDescription,omitempty"`

	// Resource
	// Resource type
	Resource *string `json:"resource,omitempty"`
}

type SCGUserGroupPermissionItemsType struct {
	Access *string `json:"access,omitempty" validate:"oneof=NA READ MODIFY FULL_ACCESS"`

	Display *string `json:"display,omitempty"`

	Resource *string `json:"resource,omitempty"`
}

type SCGUserGroupPermissionList struct {
	// Extra
	// Any additional response data.
	Extra *SCGUserGroupPermissionListExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SCGUserGroupPermission `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// SCGUserGroupPermissionListExtraType
//
// Any additional response data.
type SCGUserGroupPermissionListExtraType struct {
	// IsSuperAdmin
	// whether or not current user is a 'Super Admin' that possesses all 6 permission categories with
	// 'FULL_ACCESS'
	IsSuperAdmin *bool `json:"isSuperAdmin,omitempty"`

	// IsSuperAdminOfDomain
	// whether or not current user is a 'Super Admin of Partner Domain' that possesses all 6 permission
	// categories with 'FULL_ACCESS'
	IsSuperAdminOfDomain *bool `json:"isSuperAdminOfDomain,omitempty"`
}

type SCGUserGroupPermissionWithoutDetailItems struct {
	// Access
	// Access level
	Access *string `json:"access,omitempty" validate:"oneof=READ MODIFY FULL_ACCESS"`

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

type SCGUserGroupResourceGroup struct {
	// ID
	// the identifier of the SCG resource group
	ID *string `json:"id,omitempty"`

	// Name
	// the name of the SCG resource group
	Name *string `json:"name,omitempty"`

	// Type
	// the type of SCG resource group
	Type *string `json:"type,omitempty" validate:"oneof=DOMAIN ZONE APGROUP"`
}

type SCGUserGroupRoleLabelValue struct {
	// Label
	// Role display name
	Label *string `json:"label,omitempty"`

	// Value
	// Role value
	Value *string `json:"value,omitempty"`
}

type SCGUserGroupRoleLabelValueList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SCGUserGroupRoleLabelValue `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SCGUserList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*GetSCGUser `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
