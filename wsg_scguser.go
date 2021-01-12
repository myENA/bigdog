package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type WSGSCGUserService struct {
	apiClient *VSZClient
}

func NewWSGSCGUserService(c *VSZClient) *WSGSCGUserService {
	s := new(WSGSCGUserService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCGUserService() *WSGSCGUserService {
	return NewWSGSCGUserService(ss.apiClient)
}

// WSGSCGUserCreateScgUser
//
// Definition: scguser_createScgUser
type WSGSCGUserCreateScgUser struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	// Constraints:
	//    - nullable
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
	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	// MinimumPasswordLength
	// The minimum length of the password for the account. (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty"`

	// MinimumPasswordLifetimeEnabled
	// Enable the password should not be changed twice within the 24 hours.
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	// NewPassphrase
	// User login passphrase
	// Constraints:
	//    - required
	NewPassphrase *string `json:"newPassphrase"`

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
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s). (System default admin ONLY)
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty"`

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
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// Title
	// User title
	// Constraints:
	//    - nullable
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	// Constraints:
	//    - required
	UserName *string `json:"userName"`
}

func NewWSGSCGUserCreateScgUser() *WSGSCGUserCreateScgUser {
	m := new(WSGSCGUserCreateScgUser)
	return m
}

// WSGSCGUserGetScgUser
//
// Definition: scguser_getScgUser
type WSGSCGUserGetScgUser struct {
	// AccountLockout
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty"`

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
	// Constraints:
	//    - nullable
	Email *string `json:"email,omitempty"`

	// Enabled
	// User enabled or not
	// Constraints:
	//    - oneof:[0,1]
	Enabled *int `json:"enabled,omitempty"`

	// Id
	// User id
	Id *string `json:"id,omitempty"`

	// Locked
	// User locked or not (0:unlocked/1:locked)
	// Constraints:
	//    - oneof:[0,1,2,3]
	Locked *int `json:"locked,omitempty"`

	// LockoutDuration
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	// MinimumPasswordLength
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty"`

	// MinimumPasswordLifetimeEnabled
	// Enable the password should not be changed twice within the 24 hours.
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
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
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty"`

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
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// TenantUUID
	// Tenant id
	TenantUUID *string `json:"tenantUUID,omitempty"`

	// Title
	// User title
	// Constraints:
	//    - nullable
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}

type WSGSCGUserGetScgUserAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserGetScgUser
}

func newWSGSCGUserGetScgUserAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserGetScgUserAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserGetScgUserAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserGetScgUser)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserGetScgUser() *WSGSCGUserGetScgUser {
	m := new(WSGSCGUserGetScgUser)
	return m
}

// WSGSCGUserModifyScgUser
//
// Definition: scguser_modifyScgUser
type WSGSCGUserModifyScgUser struct {
	// AccountLockout
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty"`

	// DomainId
	// Domain id
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
	Id *string `json:"id"`

	// LockoutDuration
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	// MinimumPasswordLength
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty"`

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
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty"`

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
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// Title
	// User title
	// Constraints:
	//    - nullable
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}

func NewWSGSCGUserModifyScgUser() *WSGSCGUserModifyScgUser {
	m := new(WSGSCGUserModifyScgUser)
	return m
}

// WSGSCGUserPatchScgUserGroup
//
// Definition: scguser_patchScgUserGroup
type WSGSCGUserPatchScgUserGroup struct {
	AccountSecurityProfileId *string `json:"accountSecurityProfileId,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// Permissions
	// Constraints:
	//    - nullable
	Permissions []*WSGSCGUserGroupPermissionWithoutDetailItems `json:"permissions,omitempty"`

	// ResourceGroups
	// Constraints:
	//    - nullable
	ResourceGroups []*WSGSCGUserGroupResourceGroup `json:"resourceGroups,omitempty"`

	// Role
	// Constraints:
	//    - nullable
	//    - oneof:[SUPER_ADMIN,SYSTEM_ADMIN,NETWORK_ADMIN,RO_NETWORK_ADMIN,RO_SYSTEM_ADMIN,AP_ADMIN,GUEST_PASS_ADMIN,MVNO_SUPER_ADMIN,CUSTOM]
	Role *string `json:"role,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Constraints:
	//    - nullable
	Users []*WSGSCGUserGetScgUser `json:"users,omitempty"`
}

func NewWSGSCGUserPatchScgUserGroup() *WSGSCGUserPatchScgUserGroup {
	m := new(WSGSCGUserPatchScgUserGroup)
	return m
}

// WSGSCGUserQueryCriteria
//
// Definition: scguser_queryCriteria
type WSGSCGUserQueryCriteria struct {
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
	ExtraFilters []*WSGCommonQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*WSGCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*WSGSCGUserQueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options interface{} `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGSCGUserQueryCriteria() *WSGSCGUserQueryCriteria {
	m := new(WSGSCGUserQueryCriteria)
	return m
}

// WSGSCGUserQueryCriteriaFiltersType
//
// Definition: scguser_queryCriteriaFiltersType
type WSGSCGUserQueryCriteriaFiltersType struct {
	// Operator
	// operator
	// Constraints:
	//    - oneof:[eq]
	Operator *string `json:"operator,omitempty"`

	// Type
	// Group type
	// Constraints:
	//    - required
	//    - oneof:[DOMAIN]
	Type *string `json:"type"`

	// Value
	// DOMAIN ID
	// Constraints:
	//    - required
	Value *string `json:"value"`
}

func NewWSGSCGUserQueryCriteriaFiltersType() *WSGSCGUserQueryCriteriaFiltersType {
	m := new(WSGSCGUserQueryCriteriaFiltersType)
	return m
}

// WSGSCGUserAuditId
//
// Definition: scguser_scgUserAuditId
type WSGSCGUserAuditId struct {
	// Id
	// the identifier of the SCG user
	Id *string `json:"id,omitempty"`
}

type WSGSCGUserAuditIdAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserAuditId
}

func newWSGSCGUserAuditIdAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserAuditIdAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserAuditIdAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserAuditId)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserAuditId() *WSGSCGUserAuditId {
	m := new(WSGSCGUserAuditId)
	return m
}

// WSGSCGUserGroup
//
// Definition: scguser_scgUserGroup
type WSGSCGUserGroup struct {
	// AccountSecurityProfileId
	// Constraints:
	//    - required
	AccountSecurityProfileId *string `json:"accountSecurityProfileId"`

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
	Name *string `json:"name"`

	// Permissions
	// Permission list
	// Constraints:
	//    - required
	Permissions []*WSGSCGUserGroupPermissionWithoutDetailItems `json:"permissions"`

	// ResourceGroups
	// Resource group id list
	// Constraints:
	//    - required
	ResourceGroups []*WSGSCGUserGroupResourceGroup `json:"resourceGroups"`

	// Role
	// User group role
	// Constraints:
	//    - required
	//    - oneof:[SUPER_ADMIN,SYSTEM_ADMIN,NETWORK_ADMIN,RO_NETWORK_ADMIN,RO_SYSTEM_ADMIN,AP_ADMIN,GUEST_PASS_ADMIN,MVNO_SUPER_ADMIN,CUSTOM]
	Role *string `json:"role"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Users in this user group
	// Constraints:
	//    - nullable
	Users []*WSGSCGUserGetScgUser `json:"users,omitempty"`
}

type WSGSCGUserGroupAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserGroup
}

func newWSGSCGUserGroupAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserGroupAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserGroupAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserGroup)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserGroup() *WSGSCGUserGroup {
	m := new(WSGSCGUserGroup)
	return m
}

// WSGSCGUserGroupAuditId
//
// Definition: scguser_scgUserGroupAuditId
type WSGSCGUserGroupAuditId struct {
	// Id
	// the identifier of the SCG user group
	Id *string `json:"id,omitempty"`
}

type WSGSCGUserGroupAuditIdAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserGroupAuditId
}

func newWSGSCGUserGroupAuditIdAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserGroupAuditIdAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserGroupAuditIdAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserGroupAuditId)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserGroupAuditId() *WSGSCGUserGroupAuditId {
	m := new(WSGSCGUserGroupAuditId)
	return m
}

// WSGSCGUserGroupList
//
// Definition: scguser_scgUserGroupList
type WSGSCGUserGroupList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSCGUserGroupListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserGroupList
}

func newWSGSCGUserGroupListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserGroupListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserGroupListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserGroupList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserGroupList() *WSGSCGUserGroupList {
	m := new(WSGSCGUserGroupList)
	return m
}

// WSGSCGUserGroupPermission
//
// Definition: scguser_scgUserGroupPermission
type WSGSCGUserGroupPermission struct {
	// Access
	// Access level
	// Constraints:
	//    - oneof:[READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty"`

	// Display
	// Resource display name
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	// Constraints:
	//    - nullable
	Ids []string `json:"ids,omitempty"`

	// Items
	// Resource items
	// Constraints:
	//    - nullable
	Items []*WSGSCGUserGroupPermissionItemsType `json:"items,omitempty"`

	// ItemsDescription
	// Descriptions of Resource items
	// Constraints:
	//    - nullable
	ItemsDescription []string `json:"itemsDescription,omitempty"`

	// Resource
	// Resource type
	Resource *string `json:"resource,omitempty"`
}

func NewWSGSCGUserGroupPermission() *WSGSCGUserGroupPermission {
	m := new(WSGSCGUserGroupPermission)
	return m
}

// WSGSCGUserGroupPermissionItemsType
//
// Definition: scguser_scgUserGroupPermissionItemsType
type WSGSCGUserGroupPermissionItemsType struct {
	// Access
	// Constraints:
	//    - oneof:[NA,READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty"`

	Display *string `json:"display,omitempty"`

	Resource *string `json:"resource,omitempty"`
}

func NewWSGSCGUserGroupPermissionItemsType() *WSGSCGUserGroupPermissionItemsType {
	m := new(WSGSCGUserGroupPermissionItemsType)
	return m
}

// WSGSCGUserGroupPermissionList
//
// Definition: scguser_scgUserGroupPermissionList
type WSGSCGUserGroupPermissionList struct {
	// Extra
	// Any additional response data.
	Extra *WSGSCGUserGroupPermissionListExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGroupPermission `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSCGUserGroupPermissionListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserGroupPermissionList
}

func newWSGSCGUserGroupPermissionListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserGroupPermissionListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserGroupPermissionListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserGroupPermissionList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserGroupPermissionList() *WSGSCGUserGroupPermissionList {
	m := new(WSGSCGUserGroupPermissionList)
	return m
}

// WSGSCGUserGroupPermissionListExtraType
//
// Definition: scguser_scgUserGroupPermissionListExtraType
//
// Any additional response data.
type WSGSCGUserGroupPermissionListExtraType struct {
	// IsSuperAdmin
	// whether or not current user is a 'Super Admin' that possesses all 6 permission categories with 'FULL_ACCESS'
	IsSuperAdmin *bool `json:"isSuperAdmin,omitempty"`

	// IsSuperAdminOfDomain
	// whether or not current user is a 'Super Admin of Partner Domain' that possesses all 6 permission categories with 'FULL_ACCESS'
	IsSuperAdminOfDomain *bool `json:"isSuperAdminOfDomain,omitempty"`
}

func NewWSGSCGUserGroupPermissionListExtraType() *WSGSCGUserGroupPermissionListExtraType {
	m := new(WSGSCGUserGroupPermissionListExtraType)
	return m
}

// WSGSCGUserGroupPermissionWithoutDetailItems
//
// Definition: scguser_scgUserGroupPermissionWithoutDetailItems
type WSGSCGUserGroupPermissionWithoutDetailItems struct {
	// Access
	// Access level
	// Constraints:
	//    - oneof:[READ,MODIFY,FULL_ACCESS]
	Access *string `json:"access,omitempty"`

	// Display
	// Resource display name
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	// Constraints:
	//    - nullable
	Ids []string `json:"ids,omitempty"`

	// Resource
	// Resource type
	Resource *string `json:"resource,omitempty"`
}

func NewWSGSCGUserGroupPermissionWithoutDetailItems() *WSGSCGUserGroupPermissionWithoutDetailItems {
	m := new(WSGSCGUserGroupPermissionWithoutDetailItems)
	return m
}

// WSGSCGUserGroupResourceGroup
//
// Definition: scguser_scgUserGroupResourceGroup
type WSGSCGUserGroupResourceGroup struct {
	// Id
	// the identifier of the SCG resource group
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the SCG resource group
	Name *string `json:"name,omitempty"`

	// Type
	// the type of SCG resource group
	// Constraints:
	//    - oneof:[DOMAIN,ZONE,APGROUP]
	Type *string `json:"type,omitempty"`
}

func NewWSGSCGUserGroupResourceGroup() *WSGSCGUserGroupResourceGroup {
	m := new(WSGSCGUserGroupResourceGroup)
	return m
}

// WSGSCGUserGroupRoleLabelValue
//
// Definition: scguser_scgUserGroupRoleLabelValue
type WSGSCGUserGroupRoleLabelValue struct {
	// Label
	// Role display name
	Label *string `json:"label,omitempty"`

	// Value
	// Role value
	Value *string `json:"value,omitempty"`
}

func NewWSGSCGUserGroupRoleLabelValue() *WSGSCGUserGroupRoleLabelValue {
	m := new(WSGSCGUserGroupRoleLabelValue)
	return m
}

// WSGSCGUserGroupRoleLabelValueList
//
// Definition: scguser_scgUserGroupRoleLabelValueList
type WSGSCGUserGroupRoleLabelValueList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGroupRoleLabelValue `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSCGUserGroupRoleLabelValueListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserGroupRoleLabelValueList
}

func newWSGSCGUserGroupRoleLabelValueListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserGroupRoleLabelValueListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserGroupRoleLabelValueListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserGroupRoleLabelValueList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserGroupRoleLabelValueList() *WSGSCGUserGroupRoleLabelValueList {
	m := new(WSGSCGUserGroupRoleLabelValueList)
	return m
}

// WSGSCGUserList
//
// Definition: scguser_scgUserList
type WSGSCGUserList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGetScgUser `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSCGUserListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCGUserList
}

func newWSGSCGUserListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCGUserListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCGUserListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCGUserList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCGUserList() *WSGSCGUserList {
	m := new(WSGSCGUserList)
	return m
}

// AddUsers
//
// Add SCG user.
//
// Operation ID: addUsers
// Operation path: /users
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCGUserCreateScgUser
func (s *WSGSCGUserService) AddUsers(ctx context.Context, body *WSGSCGUserCreateScgUser, mutators ...RequestMutator) (*WSGSCGUserAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserAuditIdAPIResponse), err
}

// DeleteUsers
//
// Delete multiple SCG user.
//
// Operation ID: deleteUsers
// Operation path: /users
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSCGUserService) DeleteUsers(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteUsersByUserId
//
// Delete SCG user.
//
// Operation ID: deleteUsersByUserId
// Operation path: /users/{userId}
// Success code: 204 (No Content)
//
// Required parameters:
// - userId string
//		- required
func (s *WSGSCGUserService) DeleteUsersByUserId(ctx context.Context, userId string, mutators ...RequestMutator) (*WSGSCGUserAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteUsersByUserId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserAuditIdAPIResponse), err
}

// FindUsersByQueryCriteria
//
// Query SCG users.
//
// Operation ID: findUsersByQueryCriteria
// Operation path: /users/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserService) FindUsersByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria, mutators ...RequestMutator) (*WSGSCGUserListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindUsersByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserListAPIResponse), err
}

// FindUsersByUserId
//
// Get SCG user.
//
// Operation ID: findUsersByUserId
// Operation path: /users/{userId}
// Success code: 200 (OK)
//
// Required parameters:
// - userId string
//		- required
func (s *WSGSCGUserService) FindUsersByUserId(ctx context.Context, userId string, mutators ...RequestMutator) (*WSGSCGUserGetScgUserAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserGetScgUserAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUsersByUserId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserGetScgUserAPIResponse), err
}

// PartialUpdateUsersByUserId
//
// Update SCG user.
//
// Operation ID: partialUpdateUsersByUserId
// Operation path: /users/{userId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCGUserModifyScgUser
//
// Required parameters:
// - userId string
//		- required
func (s *WSGSCGUserService) PartialUpdateUsersByUserId(ctx context.Context, body *WSGSCGUserModifyScgUser, userId string, mutators ...RequestMutator) (*WSGSCGUserAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSCGUserAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateUsersByUserId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSCGUserAuditIdAPIResponse), err
}
