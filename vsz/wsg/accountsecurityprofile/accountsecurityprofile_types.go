package accountsecurityprofile

// API Version: v8_0

type Create struct {
	AccountLockout          *int    `json:"accountLockout,omitempty"`
	Description             *string `json:"description,omitempty"`
	DisableInactiveAccounts *int    `json:"disableInactiveAccounts,omitempty"`
	DomainID                *string `json:"domainId,omitempty"`
	LockoutDuration         *int    `json:"lockoutDuration,omitempty"`
	Name                    *string `json:"name,omitempty"`
	PasswordExpiration      *int    `json:"passwordExpiration,omitempty"`
	PasswordReuse           *int    `json:"passwordReuse,omitempty"`
	SessionIdle             *int    `json:"sessionIdle,omitempty"`
	TwoFactorAuthEnabled    *bool   `json:"twoFactorAuthEnabled,omitempty"`
}

type Delete struct {
	ID *string `json:"id,omitempty"`
}

type DeleteList struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type GetByID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetByIDResult struct {
	AccountLockout          *int    `json:"accountLockout,omitempty"`
	Description             *string `json:"description,omitempty"`
	DisableInactiveAccounts *int    `json:"disableInactiveAccounts,omitempty"`
	DomainID                *string `json:"domainId,omitempty"`
	ID                      *string `json:"id,omitempty"`
	LockoutDuration         *int    `json:"lockoutDuration,omitempty"`
	Name                    *string `json:"name,omitempty"`
	PasswordExpiration      *int    `json:"passwordExpiration,omitempty"`
	PasswordReuse           *int    `json:"passwordReuse,omitempty"`
	SessionIdle             *int    `json:"sessionIdle,omitempty"`
	TenantID                *string `json:"tenantId,omitempty"`
	TwoFactorAuthEnabled    *bool   `json:"twoFactorAuthEnabled,omitempty"`
}

type ProfileListResult struct {
	FirstIndex *int       `json:"firstIndex,omitempty"`
	HasMore    *bool      `json:"hasMore,omitempty"`
	List       []*GetByID `json:"list,omitempty"`
	TotalCount *int       `json:"totalCount,omitempty"`
}

type Update struct {
	AccountLockout          *int    `json:"accountLockout,omitempty"`
	Description             *string `json:"description,omitempty"`
	DisableInactiveAccounts *int    `json:"disableInactiveAccounts,omitempty"`
	DomainID                *string `json:"domainId,omitempty"`
	LockoutDuration         *int    `json:"lockoutDuration,omitempty"`
	Name                    *string `json:"name,omitempty"`
	PasswordExpiration      *int    `json:"passwordExpiration,omitempty"`
	PasswordReuse           *int    `json:"passwordReuse,omitempty"`
	SessionIdle             *int    `json:"sessionIdle,omitempty"`
	TwoFactorAuthEnabled    *bool   `json:"twoFactorAuthEnabled,omitempty"`
}
