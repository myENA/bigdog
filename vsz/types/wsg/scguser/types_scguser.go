package scguser

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateScgUser struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// Id
	// User id
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

	// NewPassphrase
	// User login passphrase
	// Constraints:
	//    - required
	NewPassphrase *string `json:"newPassphrase" validate:"required"`

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
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
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
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`
}

func NewCreateScgUser() *CreateScgUser {
	createScgUserType := new(CreateScgUser)
	return createScgUserType
}

func NewCreateScgUserWithDefaults() *CreateScgUser {
	createScgUserType := new(CreateScgUser)
	return createScgUserType
}

type GetScgUser struct {
	// AccountLockout
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[0,1]
	Enabled *int `json:"enabled,omitempty" validate:"omitempty,oneof=0 1"`

	// Id
	// User id
	Id *string `json:"id,omitempty"`

	// Locked
	// User locked or not (0:unlocked/1:locked)
	// Constraints:
	//    - nullable
	//    - oneof:[0,1,2]
	Locked *int `json:"locked,omitempty" validate:"omitempty,oneof=0 1 2"`

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

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

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
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	// SessionIdle
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

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

func NewGetScgUser() *GetScgUser {
	getScgUserType := new(GetScgUser)
	return getScgUserType
}

func NewGetScgUserWithDefaults() *GetScgUser {
	getScgUserType := new(GetScgUser)
	return getScgUserType
}

type ModifyScgUser struct {
	// AccountLockout
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
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

	// NewPassphrase
	// User new login passphrase
	NewPassphrase *string `json:"newPassphrase,omitempty"`

	// Passphrase
	// User login passphrase
	Passphrase *string `json:"passphrase,omitempty"`

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
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	// SessionIdle
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// Title
	// User title
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}

func NewModifyScgUser() *ModifyScgUser {
	modifyScgUserType := new(ModifyScgUser)
	return modifyScgUserType
}

func NewModifyScgUserWithDefaults() *ModifyScgUser {
	modifyScgUserType := new(ModifyScgUser)
	return modifyScgUserType
}

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

	Users []*GetScgUser `json:"users,omitempty"`
}

func NewPatchScgUserGroup() *PatchScgUserGroup {
	patchScgUserGroupType := new(PatchScgUserGroup)
	return patchScgUserGroupType
}

func NewPatchScgUserGroupWithDefaults() *PatchScgUserGroup {
	patchScgUserGroupType := new(PatchScgUserGroup)
	return patchScgUserGroupType
}

type QueryCriteria struct{}

func NewQueryCriteria() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
}

func NewQueryCriteriaWithDefaults() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
}

type ScgUserAuditId struct {
	// Id
	// the identifier of the SCG user
	Id *string `json:"id,omitempty"`
}

func NewScgUserAuditId() *ScgUserAuditId {
	scgUserAuditIdType := new(ScgUserAuditId)
	return scgUserAuditIdType
}

func NewScgUserAuditIdWithDefaults() *ScgUserAuditId {
	scgUserAuditIdType := new(ScgUserAuditId)
	return scgUserAuditIdType
}

type ScgUserGroup struct {
	// AccountSecurityProfileId
	// Constraints:
	//    - required
	AccountSecurityProfileId *string `json:"accountSecurityProfileId" validate:"required"`

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
	// Constraints:
	//    - required
	Name *string `json:"name" validate:"required"`

	// Permissions
	// Permission list
	// Constraints:
	//    - required
	Permissions []*ScgUserGroupPermissionWithoutDetailItems `json:"permissions" validate:"required"`

	// ResourceGroups
	// Resource group id list
	// Constraints:
	//    - required
	ResourceGroups []*ScgUserGroupResourceGroup `json:"resourceGroups" validate:"required"`

	// Role
	// User group role
	// Constraints:
	//    - required
	Role *string `json:"role" validate:"required"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Users in this user group
	Users []*GetScgUser `json:"users,omitempty"`
}

func NewScgUserGroup() *ScgUserGroup {
	scgUserGroupType := new(ScgUserGroup)
	return scgUserGroupType
}

func NewScgUserGroupWithDefaults() *ScgUserGroup {
	scgUserGroupType := new(ScgUserGroup)
	return scgUserGroupType
}

type ScgUserGroupAuditId struct {
	// Id
	// the identifier of the SCG user group
	Id *string `json:"id,omitempty"`
}

func NewScgUserGroupAuditId() *ScgUserGroupAuditId {
	scgUserGroupAuditIdType := new(ScgUserGroupAuditId)
	return scgUserGroupAuditIdType
}

func NewScgUserGroupAuditIdWithDefaults() *ScgUserGroupAuditId {
	scgUserGroupAuditIdType := new(ScgUserGroupAuditId)
	return scgUserGroupAuditIdType
}

type ScgUserGroupList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewScgUserGroupList() *ScgUserGroupList {
	scgUserGroupListType := new(ScgUserGroupList)
	return scgUserGroupListType
}

func NewScgUserGroupListWithDefaults() *ScgUserGroupList {
	scgUserGroupListType := new(ScgUserGroupList)
	return scgUserGroupListType
}

type ScgUserGroupPermission struct {
	// Access
	// Access level
	// Constraints:
	//    - nullable
	//    - oneof:[READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=READ MODIFY FULL_ACCESS"`

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

func NewScgUserGroupPermission() *ScgUserGroupPermission {
	scgUserGroupPermissionType := new(ScgUserGroupPermission)
	return scgUserGroupPermissionType
}

func NewScgUserGroupPermissionWithDefaults() *ScgUserGroupPermission {
	scgUserGroupPermissionType := new(ScgUserGroupPermission)
	return scgUserGroupPermissionType
}

type ScgUserGroupPermissionItemsType struct {
	// Access
	// Constraints:
	//    - nullable
	//    - oneof:[NA,READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=NA READ MODIFY FULL_ACCESS"`

	Display *string `json:"display,omitempty"`

	Resource *string `json:"resource,omitempty"`
}

func NewScgUserGroupPermissionItemsType() *ScgUserGroupPermissionItemsType {
	scgUserGroupPermissionItemsTypeType := new(ScgUserGroupPermissionItemsType)
	return scgUserGroupPermissionItemsTypeType
}

func NewScgUserGroupPermissionItemsTypeWithDefaults() *ScgUserGroupPermissionItemsType {
	scgUserGroupPermissionItemsTypeType := new(ScgUserGroupPermissionItemsType)
	return scgUserGroupPermissionItemsTypeType
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

func NewScgUserGroupPermissionList() *ScgUserGroupPermissionList {
	scgUserGroupPermissionListType := new(ScgUserGroupPermissionList)
	return scgUserGroupPermissionListType
}

func NewScgUserGroupPermissionListWithDefaults() *ScgUserGroupPermissionList {
	scgUserGroupPermissionListType := new(ScgUserGroupPermissionList)
	return scgUserGroupPermissionListType
}

// ScgUserGroupPermissionListExtraType
//
// Any additional response data.
type ScgUserGroupPermissionListExtraType struct {
	// IsSuperAdmin
	// whether or not current user is a 'Super Admin' that possesses all 6 permission categories with 'FULL_ACCESS'
	IsSuperAdmin *bool `json:"isSuperAdmin,omitempty"`

	// IsSuperAdminOfDomain
	// whether or not current user is a 'Super Admin of Partner Domain' that possesses all 6 permission categories with 'FULL_ACCESS'
	IsSuperAdminOfDomain *bool `json:"isSuperAdminOfDomain,omitempty"`
}

func NewScgUserGroupPermissionListExtraType() *ScgUserGroupPermissionListExtraType {
	scgUserGroupPermissionListExtraTypeType := new(ScgUserGroupPermissionListExtraType)
	return scgUserGroupPermissionListExtraTypeType
}

func NewScgUserGroupPermissionListExtraTypeWithDefaults() *ScgUserGroupPermissionListExtraType {
	scgUserGroupPermissionListExtraTypeType := new(ScgUserGroupPermissionListExtraType)
	return scgUserGroupPermissionListExtraTypeType
}

type ScgUserGroupPermissionWithoutDetailItems struct {
	// Access
	// Access level
	// Constraints:
	//    - nullable
	//    - oneof:[READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=READ MODIFY FULL_ACCESS"`

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

func NewScgUserGroupPermissionWithoutDetailItems() *ScgUserGroupPermissionWithoutDetailItems {
	scgUserGroupPermissionWithoutDetailItemsType := new(ScgUserGroupPermissionWithoutDetailItems)
	return scgUserGroupPermissionWithoutDetailItemsType
}

func NewScgUserGroupPermissionWithoutDetailItemsWithDefaults() *ScgUserGroupPermissionWithoutDetailItems {
	scgUserGroupPermissionWithoutDetailItemsType := new(ScgUserGroupPermissionWithoutDetailItems)
	return scgUserGroupPermissionWithoutDetailItemsType
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
	// Constraints:
	//    - nullable
	//    - oneof:[DOMAIN,ZONE,APGROUP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=DOMAIN ZONE APGROUP"`
}

func NewScgUserGroupResourceGroup() *ScgUserGroupResourceGroup {
	scgUserGroupResourceGroupType := new(ScgUserGroupResourceGroup)
	return scgUserGroupResourceGroupType
}

func NewScgUserGroupResourceGroupWithDefaults() *ScgUserGroupResourceGroup {
	scgUserGroupResourceGroupType := new(ScgUserGroupResourceGroup)
	return scgUserGroupResourceGroupType
}

type ScgUserGroupRoleLabelValue struct {
	// Label
	// Role display name
	Label *string `json:"label,omitempty"`

	// Value
	// Role value
	Value *string `json:"value,omitempty"`
}

func NewScgUserGroupRoleLabelValue() *ScgUserGroupRoleLabelValue {
	scgUserGroupRoleLabelValueType := new(ScgUserGroupRoleLabelValue)
	return scgUserGroupRoleLabelValueType
}

func NewScgUserGroupRoleLabelValueWithDefaults() *ScgUserGroupRoleLabelValue {
	scgUserGroupRoleLabelValueType := new(ScgUserGroupRoleLabelValue)
	return scgUserGroupRoleLabelValueType
}

type ScgUserGroupRoleLabelValueList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroupRoleLabelValue `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewScgUserGroupRoleLabelValueList() *ScgUserGroupRoleLabelValueList {
	scgUserGroupRoleLabelValueListType := new(ScgUserGroupRoleLabelValueList)
	return scgUserGroupRoleLabelValueListType
}

func NewScgUserGroupRoleLabelValueListWithDefaults() *ScgUserGroupRoleLabelValueList {
	scgUserGroupRoleLabelValueListType := new(ScgUserGroupRoleLabelValueList)
	return scgUserGroupRoleLabelValueListType
}

type ScgUserList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*GetScgUser `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewScgUserList() *ScgUserList {
	scgUserListType := new(ScgUserList)
	return scgUserListType
}

func NewScgUserListWithDefaults() *ScgUserList {
	scgUserListType := new(ScgUserList)
	return scgUserListType
}
