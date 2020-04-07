package vsz

// API Version: v9_0

type WSGAccountSecurityProfileCreate struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// AttemptsLockEnabled
	// Constraints:
	//    - nullable
	AttemptsLockEnabled *bool `json:"attemptsLockEnabled,omitempty"`

	// AttemptsLockoutPeriod
	// A period of attempts times.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	AttemptsLockoutPeriod *int `json:"attemptsLockoutPeriod,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// AttemptsLockoutTimes
	// The attempts times.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AttemptsLockoutTimes *int `json:"attemptsLockoutTimes,omitempty" validate:"omitempty,gte=1,lte=100"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1000
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty" validate:"omitempty,gte=1,lte=1000"`

	// DomainId
	// Domain id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention.
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
	// Control minimum password life time.
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PasswordComplexityEnabled
	// Control password complexity mechanism.
	// Constraints:
	//    - nullable
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:365
	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s).
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// SessionIdle
	// A period of idle used to invalid that session.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// TwoFactorAuthEnabled
	// Enable the two-factor authentication. (This configuration can only be enabled from Web GUI.)
	// Constraints:
	//    - nullable
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

func NewWSGAccountSecurityProfileCreate() *WSGAccountSecurityProfileCreate {
	m := new(WSGAccountSecurityProfileCreate)
	return m
}

type WSGAccountSecurityProfileDelete struct {
	// Id
	// Profile id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`
}

func NewWSGAccountSecurityProfileDelete() *WSGAccountSecurityProfileDelete {
	m := new(WSGAccountSecurityProfileDelete)
	return m
}

type WSGAccountSecurityProfileDeleteList struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGAccountSecurityProfileDeleteList() *WSGAccountSecurityProfileDeleteList {
	m := new(WSGAccountSecurityProfileDeleteList)
	return m
}

type WSGAccountSecurityProfileGetById struct {
	// Id
	// Profile id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGAccountSecurityProfileGetById() *WSGAccountSecurityProfileGetById {
	m := new(WSGAccountSecurityProfileGetById)
	return m
}

type WSGAccountSecurityProfileGetByIdResult struct {
	// AccountLockout
	// Constraints:
	//    - nullable
	AccountLockout *int `json:"accountLockout,omitempty"`

	// AttemptsLockEnabled
	// Constraints:
	//    - nullable
	AttemptsLockEnabled *bool `json:"attemptsLockEnabled,omitempty"`

	// AttemptsLockoutPeriod
	// Constraints:
	//    - nullable
	AttemptsLockoutPeriod *int `json:"attemptsLockoutPeriod,omitempty"`

	// AttemptsLockoutTimes
	// Constraints:
	//    - nullable
	AttemptsLockoutTimes *int `json:"attemptsLockoutTimes,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DisableInactiveAccounts
	// Constraints:
	//    - nullable
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LockoutDuration
	// Constraints:
	//    - nullable
	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	// MinimumPasswordLength
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	// MinimumPasswordLifetimeEnabled
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PasswordComplexityEnabled
	// Constraints:
	//    - nullable
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// Constraints:
	//    - nullable
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// Constraints:
	//    - nullable
	PasswordReuse *int `json:"passwordReuse,omitempty"`

	// SessionIdle
	// Constraints:
	//    - nullable
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// TwoFactorAuthEnabled
	// Constraints:
	//    - nullable
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

func NewWSGAccountSecurityProfileGetByIdResult() *WSGAccountSecurityProfileGetByIdResult {
	m := new(WSGAccountSecurityProfileGetByIdResult)
	return m
}

type WSGAccountSecurityProfileProfileListResult struct {
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
	List []*WSGAccountSecurityProfileGetById `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAccountSecurityProfileProfileListResult() *WSGAccountSecurityProfileProfileListResult {
	m := new(WSGAccountSecurityProfileProfileListResult)
	return m
}

type WSGAccountSecurityProfileUpdate struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// AttemptsLockEnabled
	// Constraints:
	//    - nullable
	AttemptsLockEnabled *bool `json:"attemptsLockEnabled,omitempty"`

	// AttemptsLockoutPeriod
	// A period of attempts times.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	AttemptsLockoutPeriod *int `json:"attemptsLockoutPeriod,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// AttemptsLockoutTimes
	// The attempts times.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AttemptsLockoutTimes *int `json:"attemptsLockoutTimes,omitempty" validate:"omitempty,gte=1,lte=100"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1000
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty" validate:"omitempty,gte=1,lte=1000"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention.
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
	// Control minimum password life time.
	// Constraints:
	//    - nullable
	MinimumPasswordLifetimeEnabled *bool `json:"minimumPasswordLifetimeEnabled,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PasswordComplexityEnabled
	// Control password complexity mechanism.
	// Constraints:
	//    - nullable
	PasswordComplexityEnabled *bool `json:"passwordComplexityEnabled,omitempty"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:365
	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s).
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:6
	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// SessionIdle
	// A period of idle used to invalid that session.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1440
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// TwoFactorAuthEnabled
	// Enable the two-factor authentication. (This configuration can only be enabled from Web GUI.)
	// Constraints:
	//    - nullable
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

func NewWSGAccountSecurityProfileUpdate() *WSGAccountSecurityProfileUpdate {
	m := new(WSGAccountSecurityProfileUpdate)
	return m
}
