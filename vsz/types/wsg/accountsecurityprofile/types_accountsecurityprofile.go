package accountsecurityprofile

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type Create struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	Description *common.Description `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1000
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty" validate:"omitempty,gte=1,lte=1000"`

	// DomainId
	// Domain id
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

	Name *common.NormalName `json:"name,omitempty"`

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
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

type Delete struct {
	// Id
	// Profile id
	Id *string `json:"id,omitempty"`
}

type DeleteList struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type GetById struct {
	// Id
	// Profile id
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type GetByIdResult struct {
	AccountLockout *int `json:"accountLockout,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	// MinimumPasswordLength
	// Constraints:
	//    - nullable
	//    - min:8
	//    - max:64
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	Name *common.NormalName `json:"name,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	PasswordReuse *int `json:"passwordReuse,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

type ProfileListResult struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*GetById `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type Update struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:100
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	Description *common.Description `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:1000
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty" validate:"omitempty,gte=1,lte=1000"`

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

	Name *common.NormalName `json:"name,omitempty"`

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
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}
