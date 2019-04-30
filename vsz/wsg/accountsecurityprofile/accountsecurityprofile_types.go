package accountsecurityprofile

// API Version: v8_0

type Create struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	AccountLockout *int `json:"accountLockout,omitempty"`

	Description *string `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	// DomainID
	// Domain id
	DomainID *string `json:"domainId,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention.
	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	Name *string `json:"name,omitempty"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly.
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s).
	PasswordReuse *int `json:"passwordReuse,omitempty"`

	// SessionIdle
	// A period of idle used to invalid that session.
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// TwoFactorAuthEnabled
	// Enable the two-factor authentication. (This configuration can only be enabled from Web GUI.)
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

type Delete struct {
	// ID
	// Profile id
	ID *string `json:"id,omitempty"`
}

type DeleteList struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type GetByID struct {
	// ID
	// Profile id
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type GetByIDResult struct {
	AccountLockout *int `json:"accountLockout,omitempty"`

	Description *string `json:"description,omitempty"`

	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	ID *string `json:"id,omitempty"`

	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	Name *string `json:"name,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	PasswordReuse *int `json:"passwordReuse,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty"`

	TenantID *string `json:"tenantId,omitempty"`

	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}

type ProfileListResult struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*GetByID `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type Update struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur.
	AccountLockout *int `json:"accountLockout,omitempty"`

	Description *string `json:"description,omitempty"`

	// DisableInactiveAccounts
	// A period in a inactive status used to lockout these accounts.
	DisableInactiveAccounts *int `json:"disableInactiveAccounts,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention.
	LockoutDuration *int `json:"lockoutDuration,omitempty"`

	Name *string `json:"name,omitempty"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly.
	PasswordExpiration *int `json:"passwordExpiration,omitempty"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s).
	PasswordReuse *int `json:"passwordReuse,omitempty"`

	// SessionIdle
	// A period of idle used to invalid that session.
	SessionIdle *int `json:"sessionIdle,omitempty"`

	// TwoFactorAuthEnabled
	// Enable the two-factor authentication. (This configuration can only be enabled from Web GUI.)
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
}
