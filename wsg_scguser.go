package bigdog

// API Version: v9_0

import (
	"context"
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
	Role *string `json:"role,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Constraints:
	//    - required
	Users []*WSGSCGUserGetScgUser `json:"users"`
}

func NewWSGSCGUserPatchScgUserGroup() *WSGSCGUserPatchScgUserGroup {
	m := new(WSGSCGUserPatchScgUserGroup)
	return m
}

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
	Options *WSGCommonQueryCriteriaOptionsType `json:"options,omitempty"`

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

type WSGSCGUserAuditId struct {
	// Id
	// the identifier of the SCG user
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
	Role *string `json:"role"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Users in this user group
	// Constraints:
	//    - required
	Users []*WSGSCGUserGetScgUser `json:"users"`
}

func NewWSGSCGUserGroup() *WSGSCGUserGroup {
	m := new(WSGSCGUserGroup)
	return m
}

type WSGSCGUserGroupAuditId struct {
	// Id
	// the identifier of the SCG user group
	Id *string `json:"id,omitempty"`
}

func NewWSGSCGUserGroupAuditId() *WSGSCGUserGroupAuditId {
	m := new(WSGSCGUserGroupAuditId)
	return m
}

type WSGSCGUserGroupList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGroup `json:"list,omitempty"`

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

type WSGSCGUserGroupPermissionList struct {
	// Extra
	// Any additional response data.
	Extra *WSGSCGUserGroupPermissionListExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGroupPermission `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCGUserGroupPermissionList() *WSGSCGUserGroupPermissionList {
	m := new(WSGSCGUserGroupPermissionList)
	return m
}

// WSGSCGUserGroupPermissionListExtraType
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

type WSGSCGUserGroupRoleLabelValueList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGroupRoleLabelValue `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCGUserGroupRoleLabelValueList() *WSGSCGUserGroupRoleLabelValueList {
	m := new(WSGSCGUserGroupRoleLabelValueList)
	return m
}

type WSGSCGUserList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCGUserGetScgUser `json:"list,omitempty"`

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
func (s *WSGSCGUserService) AddUsers(ctx context.Context, body *WSGSCGUserCreateScgUser, mutators ...RequestMutator) (*WSGSCGUserAuditId, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddUsers, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *WSGSCGUserService) DeleteUsers(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUsers, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *WSGSCGUserService) DeleteUsersByUserId(ctx context.Context, userId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteUsersByUserId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindUsersByQueryCriteria
//
// Query SCG users.
//
// Request Body:
//	 - body *WSGSCGUserQueryCriteria
func (s *WSGSCGUserService) FindUsersByQueryCriteria(ctx context.Context, body *WSGSCGUserQueryCriteria, mutators ...RequestMutator) (*WSGSCGUserList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindUsersByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *WSGSCGUserService) FindUsersByUserId(ctx context.Context, userId string, mutators ...RequestMutator) (*WSGSCGUserGetScgUser, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUsersByUserId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *WSGSCGUserService) PartialUpdateUsersByUserId(ctx context.Context, body *WSGSCGUserModifyScgUser, userId string, mutators ...RequestMutator) (*WSGSCGUserAuditId, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateUsersByUserId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSCGUserAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
