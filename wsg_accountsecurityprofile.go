package bigdog

// API Version: v9_1

// WSGAccountSecurityProfileCreate
//
// Definition: accountSecurityProfile_create
type WSGAccountSecurityProfileCreate struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	// Constraints:
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty"`

	AttemptsLockEnabled *bool `json:"attemptsLockEnabled,omitempty"`

	// AttemptsLockoutPeriod
	// A period of attempts times.
	// Constraints:
	//    - min:1
	//    - max:1440
	AttemptsLockoutPeriod *int `json:"attemptsLockoutPeriod,omitempty"`

	// AttemptsLockoutTimes
	// The attempts times.
	// Constraints:
	//    - min:1
	//    - max:100
	AttemptsLockoutTimes *int `json:"attemptsLockoutTimes,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	// Constraints:
	//    - min:1
	//    - max:1000
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention.
	// Constraints:
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
	// Control minimum password life time.
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PasswordComplexityEnabled
	// Control password complexity mechanism.
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly.
	// Constraints:
	//    - min:1
	//    - max:365
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s).
	// Constraints:
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty"`

	// SessionIdle
	// A period of idle used to invalid that session.
	// Constraints:
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// TwoFactorAuthEnabled
	// Enable the two-factor authentication. (This configuration can only be enabled from Web GUI.)
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

func NewWSGAccountSecurityProfileCreate() *WSGAccountSecurityProfileCreate {
	m := new(WSGAccountSecurityProfileCreate)
	return m
}

// WSGAccountSecurityProfileDelete
//
// Definition: accountSecurityProfile_delete
type WSGAccountSecurityProfileDelete struct {
	// Id
	// Profile id
	Id *string `json:"id,omitempty"`
}

func NewWSGAccountSecurityProfileDelete() *WSGAccountSecurityProfileDelete {
	m := new(WSGAccountSecurityProfileDelete)
	return m
}

// WSGAccountSecurityProfileDeleteList
//
// Definition: accountSecurityProfile_deleteList
type WSGAccountSecurityProfileDeleteList struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGAccountSecurityProfileDeleteList() *WSGAccountSecurityProfileDeleteList {
	m := new(WSGAccountSecurityProfileDeleteList)
	return m
}

// WSGAccountSecurityProfileGetById
//
// Definition: accountSecurityProfile_getById
type WSGAccountSecurityProfileGetById struct {
	// Id
	// Profile id
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewWSGAccountSecurityProfileGetById() *WSGAccountSecurityProfileGetById {
	m := new(WSGAccountSecurityProfileGetById)
	return m
}

// WSGAccountSecurityProfileGetByIdResult
//
// Definition: accountSecurityProfile_getByIdResult
type WSGAccountSecurityProfileGetByIdResult struct {
	AccountLockout *int `json:"accountLockout,omitempty"`

	AttemptsLockEnabled *bool `json:"attemptsLockEnabled,omitempty"`

	AttemptsLockoutPeriod *int `json:"attemptsLockoutPeriod,omitempty"`

	AttemptsLockoutTimes *int `json:"attemptsLockoutTimes,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	// MinimumPasswordLength
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty"`

	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	PasswordReuse *int `json:"passwordReuse,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

func NewWSGAccountSecurityProfileGetByIdResult() *WSGAccountSecurityProfileGetByIdResult {
	m := new(WSGAccountSecurityProfileGetByIdResult)
	return m
}

// WSGAccountSecurityProfileProfileListResult
//
// Definition: accountSecurityProfile_profileListResult
type WSGAccountSecurityProfileProfileListResult struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAccountSecurityProfileGetById `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAccountSecurityProfileProfileListResult() *WSGAccountSecurityProfileProfileListResult {
	m := new(WSGAccountSecurityProfileProfileListResult)
	return m
}

// WSGAccountSecurityProfileUpdate
//
// Definition: accountSecurityProfile_update
type WSGAccountSecurityProfileUpdate struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	// Constraints:
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty"`

	AttemptsLockEnabled *bool `json:"attemptsLockEnabled,omitempty"`

	// AttemptsLockoutPeriod
	// A period of attempts times.
	// Constraints:
	//    - min:1
	//    - max:1440
	AttemptsLockoutPeriod *int `json:"attemptsLockoutPeriod,omitempty"`

	// AttemptsLockoutTimes
	// The attempts times.
	// Constraints:
	//    - min:1
	//    - max:100
	AttemptsLockoutTimes *int `json:"attemptsLockoutTimes,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	// Constraints:
	//    - min:1
	//    - max:1000
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention.
	// Constraints:
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
	// Control minimum password life time.
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PasswordComplexityEnabled
	// Control password complexity mechanism.
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly.
	// Constraints:
	//    - min:1
	//    - max:365
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s).
	// Constraints:
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty"`

	// SessionIdle
	// A period of idle used to invalid that session.
	// Constraints:
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// TwoFactorAuthEnabled
	// Enable the two-factor authentication. (This configuration can only be enabled from Web GUI.)
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

func NewWSGAccountSecurityProfileUpdate() *WSGAccountSecurityProfileUpdate {
	m := new(WSGAccountSecurityProfileUpdate)
	return m
}
