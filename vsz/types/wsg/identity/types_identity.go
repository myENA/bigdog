package identity

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AaaServer struct {
	// Id
	// the identifier of the AAA server
	Id *string `json:"id,omitempty"`

	// Name
	// the identifier of the AAA server
	Name *string `json:"name,omitempty"`
}

type AaaServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AaaServer `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type AuthenticationServerConfig struct {
	// AuthenticationServerID
	// Authentication server id
	AuthenticationServerID *string `json:"AUTHENTICATION_SERVER_ID,omitempty"`

	// AuthenticationServerName
	// Authentication server name
	AuthenticationServerName *string `json:"AUTHENTICATION_SERVER_NAME,omitempty"`

	// AuthenticationServerType
	// Authentication server type
	AuthenticationServerType *string `json:"AUTHENTICATION_SERVER_TYPE,omitempty"`

	// Id
	// server id
	Id *string `json:"id,omitempty"`

	// Local
	// Is local server
	Local *bool `json:"local,omitempty"`

	// Name
	// server name
	Name *string `json:"name,omitempty"`

	// Type
	// server type
	Type *string `json:"type,omitempty"`
}

type CountryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*CountrySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type CountrySummary struct {
	// CountryName
	// Full name of country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Short name of country
	CountryShortName *string `json:"countryShortName,omitempty"`
}

type CreateIdentityGuestPass struct {
	// AutoGeneratedPassword
	// Pass generation
	AutoGeneratedPassword *bool `json:"autoGeneratedPassword,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// GuestName
	// Constraints:
	//    - required
	GuestName *common.NormalName `json:"guestName" validate:"required"`

	// MaxDevices
	// Constraints:
	//    - required
	MaxDevices *MaxDevices `json:"maxDevices" validate:"required"`

	// NumberOfPasses
	// Number of passes
	// Constraints:
	//    - required
	NumberOfPasses *int `json:"numberOfPasses" validate:"required"`

	// PassEffectSince
	// Pass effective since
	// Constraints:
	//    - nullable
	//    - oneof:[CREATION_TIME,FIRST_USE]
	PassEffectSince *string `json:"passEffectSince,omitempty" validate:"omitempty,oneof=CREATION_TIME FIRST_USE"`

	// PassUseDays
	// Expire new guest pass if not used within
	PassUseDays *int `json:"passUseDays,omitempty"`

	// PassValidFor
	// Constraints:
	//    - required
	PassValidFor *PassValidFor `json:"passValidFor" validate:"required"`

	// PassValue
	// Pass value
	PassValue *string `json:"passValue,omitempty"`

	// Remarks
	// Remarks
	Remarks *string `json:"remarks,omitempty"`

	SessionDuration *SessionDuration `json:"sessionDuration,omitempty"`

	// Wlan
	// Constraints:
	//    - required
	Wlan *common.GenericRef `json:"wlan" validate:"required"`

	// Zone
	// Constraints:
	//    - required
	Zone *common.GenericRef `json:"zone" validate:"required"`
}

type CreateIdentityUserRole struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// MaxDevices
	// Constraints:
	//    - required
	MaxDevices *MaxDevices `json:"maxDevices" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName2to64 `json:"name" validate:"required"`

	// UserTrafficProfile
	// Constraints:
	//    - required
	UserTrafficProfile *common.GenericRef `json:"userTrafficProfile" validate:"required"`

	// VlanId
	// vlan id
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4096
	VlanId *int `json:"vlanId,omitempty" validate:"omitempty,gte=1,lte=4096"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

type CreateSubscriptionPackage struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - required
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval" validate:"required,oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	// Constraints:
	//    - required
	ExpirationValue *int `json:"expirationValue" validate:"required"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

type CreateUser struct {
	// Address
	// Address
	Address *string `json:"address,omitempty"`

	// City
	// City
	City *string `json:"city,omitempty"`

	// CountryName
	// Country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Country
	CountryShortName *string `json:"countryShortName,omitempty"`

	// DomainId
	// Domain ID
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// Email
	Email *string `json:"email,omitempty"`

	// FirstName
	// First name
	// Constraints:
	//    - required
	FirstName *string `json:"firstName" validate:"required"`

	// IsDisabled
	// Is Disabled
	// Constraints:
	//    - required
	//    - oneof:[NO,YES]
	IsDisabled *string `json:"isDisabled" validate:"required,oneof=NO YES"`

	// LastName
	// Last Name
	// Constraints:
	//    - required
	LastName *string `json:"lastName" validate:"required"`

	// Password
	// Password
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// Phone
	// Phone
	Phone *string `json:"phone,omitempty"`

	// Remark
	// Remark
	Remark *string `json:"remark,omitempty"`

	// State
	// State
	State *string `json:"state,omitempty"`

	SubscriberPackage *common.GenericRef `json:"subscriberPackage,omitempty"`

	// UserName
	// User Name
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`

	// ZipCode
	// Zip Code
	ZipCode *string `json:"zipCode,omitempty"`
}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type IdentityGuestPassConfiguration struct {
	// AutoGeneratedPassword
	// Pass generation
	AutoGeneratedPassword *bool `json:"autoGeneratedPassword,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationDate
	// Expiration date and time
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// GeneratedOn
	// Generated date and time
	GeneratedOn *string `json:"generatedOn,omitempty"`

	GuestName *common.NormalName `json:"guestName,omitempty"`

	// Id
	// ID of the identity guest pass
	Id *string `json:"id,omitempty"`

	// Key
	// Identifier of the identity guest pass
	Key *string `json:"key,omitempty"`

	MaxDevices *MaxDevices `json:"maxDevices,omitempty"`

	// PassEffectSince
	// Pass effective since
	// Constraints:
	//    - nullable
	//    - oneof:[CREATION_TIME,FIRST_USE]
	PassEffectSince *string `json:"passEffectSince,omitempty" validate:"omitempty,oneof=CREATION_TIME FIRST_USE"`

	// PassUseDays
	// Expire new guest pass if not used within
	PassUseDays *int `json:"passUseDays,omitempty"`

	PassValidFor *PassValidFor `json:"passValidFor,omitempty"`

	// Remarks
	// Remarks
	Remarks *string `json:"remarks,omitempty"`

	SessionDuration *SessionDuration `json:"sessionDuration,omitempty"`

	// Ssid
	// SSID
	Ssid *string `json:"ssid,omitempty"`

	// UserId
	// user ID of the identity guest pass
	UserId *string `json:"userId,omitempty"`

	Wlan *common.GenericRef `json:"wlan,omitempty"`

	// WlanRestrition
	// Wlan description
	WlanRestrition *string `json:"wlanRestrition,omitempty"`

	Zone *common.GenericRef `json:"zone,omitempty"`
}

type IdentityGuestPassList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*IdentityGuestPassConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type IdentityList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*IdentityListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type IdentityListType struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	MaxDevices *MaxDevices `json:"maxDevices,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	UserTrafficProfile *common.GenericRef `json:"userTrafficProfile,omitempty"`

	// VlanId
	// vlan id
	VlanId *int `json:"vlanId,omitempty"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

type IdentityUserRole struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Id
	// the identifier of the object
	Id *string `json:"id,omitempty"`

	MaxDevices *MaxDevices `json:"maxDevices,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName2to64 `json:"name,omitempty"`

	UserTrafficProfile *common.GenericRef `json:"userTrafficProfile,omitempty"`

	// VlanId
	// vlan id
	VlanId *int `json:"vlanId,omitempty"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

type IdentityUserSummary struct {
	// CreatedOn
	// Created on
	CreatedOn *string `json:"createdOn,omitempty"`

	DisplayName *common.NormalName `json:"displayName,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the identity user
	Id *string `json:"id,omitempty"`

	// IsDisabled
	// Is disalbed
	IsDisabled *string `json:"isDisabled,omitempty"`

	// UserName
	// User Name
	UserName *string `json:"userName,omitempty"`

	// UserSource
	// User Source
	UserSource *string `json:"userSource,omitempty"`

	// UserType
	// User Type
	UserType *string `json:"userType,omitempty"`
}

type ImportIdentityGuestPass struct {
	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// MaxDevices
	// Constraints:
	//    - required
	MaxDevices *MaxDevices `json:"maxDevices" validate:"required"`

	// PassEffectSince
	// Pass effective since
	// Constraints:
	//    - nullable
	//    - oneof:[CREATION_TIME,FIRST_USE]
	PassEffectSince *string `json:"passEffectSince,omitempty" validate:"omitempty,oneof=CREATION_TIME FIRST_USE"`

	// PassUseDays
	// Expire new guest pass if not used within
	PassUseDays *int `json:"passUseDays,omitempty"`

	// PassValidFor
	// Constraints:
	//    - required
	PassValidFor *PassValidFor `json:"passValidFor" validate:"required"`

	SessionDuration *SessionDuration `json:"sessionDuration,omitempty"`

	// Wlan
	// Constraints:
	//    - required
	Wlan *common.GenericRef `json:"wlan" validate:"required"`

	// Zone
	// Constraints:
	//    - required
	Zone *common.GenericRef `json:"zone" validate:"required"`
}

type MaxDevices struct {
	// MaxDevicesAllowed
	// Max devices allowed
	// Constraints:
	//    - nullable
	//    - default:'LIMITED'
	//    - oneof:[UNLIMITED,LIMITED]
	MaxDevicesAllowed *string `json:"maxDevicesAllowed,omitempty" validate:"omitempty,oneof=UNLIMITED LIMITED"`

	// MaxDevicesNumber
	// max devices number
	// Constraints:
	//    - nullable
	//    - default:3
	//    - min:1
	//    - max:10
	MaxDevicesNumber *int `json:"maxDevicesNumber,omitempty" validate:"omitempty,gte=1,lte=10"`
}

type ModifyIdentityUserRole struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// ID
	Id *string `json:"id,omitempty"`

	MaxDevices *MaxDevices `json:"maxDevices,omitempty"`

	Name *common.NormalName2to64 `json:"name,omitempty"`

	UserTrafficProfile *common.GenericRef `json:"userTrafficProfile,omitempty"`

	// VlanId
	// vlan id
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4096
	VlanId *int `json:"vlanId,omitempty" validate:"omitempty,gte=1,lte=4096"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

type ModifySubscriptionPackage struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - nullable
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval,omitempty" validate:"omitempty,oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	ExpirationValue *int `json:"expirationValue,omitempty"`

	// Id
	// ID
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type ModifyUser struct {
	// Address
	// Address
	Address *string `json:"address,omitempty"`

	// City
	// City
	City *string `json:"city,omitempty"`

	// CountryName
	// Country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Country
	CountryShortName *string `json:"countryShortName,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// Email
	Email *string `json:"email,omitempty"`

	// FirstName
	// First name
	FirstName *string `json:"firstName,omitempty"`

	// Id
	// ID
	Id *string `json:"id,omitempty"`

	// IsDisabled
	// Is Disabled
	// Constraints:
	//    - nullable
	//    - oneof:[NO,YES]
	IsDisabled *string `json:"isDisabled,omitempty" validate:"omitempty,oneof=NO YES"`

	// LastName
	// Last Name
	LastName *string `json:"lastName,omitempty"`

	// Password
	// Password
	Password *string `json:"password,omitempty"`

	// Phone
	// Phone
	Phone *string `json:"phone,omitempty"`

	// Remark
	// Remark
	Remark *string `json:"remark,omitempty"`

	// State
	// State
	State *string `json:"state,omitempty"`

	SubscriberPackage *common.GenericRef `json:"subscriberPackage,omitempty"`

	// ZipCode
	// Zip Code
	ZipCode *string `json:"zipCode,omitempty"`
}

type PackageConfiguration struct {
	// PackageExpiration
	// Package expiration interval and value
	PackageExpiration *string `json:"packageExpiration,omitempty"`

	SubscriberPackage *common.GenericRef `json:"subscriberPackage,omitempty"`
}

type PackageList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PackageConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type PassValidFor struct {
	// ExpirationUnit
	// Constraints:
	//    - nullable
	//    - oneof:[HOUR,DAY,WEEK]
	ExpirationUnit *string `json:"expirationUnit,omitempty" validate:"omitempty,oneof=HOUR DAY WEEK"`

	ExpirationValue *int `json:"expirationValue,omitempty"`
}

type QueryCriteria struct {}

type SessionDuration struct {
	RequireLoginAgain *bool `json:"requireLoginAgain,omitempty"`

	// SessionUnit
	// Constraints:
	//    - nullable
	//    - oneof:[MIN,HOUR,DAY,WEEK]
	SessionUnit *string `json:"sessionUnit,omitempty" validate:"omitempty,oneof=MIN HOUR DAY WEEK"`

	SessionValue *int `json:"sessionValue,omitempty"`
}

type SubscriptionPackage struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - nullable
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval,omitempty" validate:"omitempty,oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	ExpirationValue *int `json:"expirationValue,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type SubscriptionPackageList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SubscriptionPackageListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SubscriptionPackageListType struct {
	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationInterval
	// Expiration interval
	// Constraints:
	//    - nullable
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	ExpirationInterval *string `json:"expirationInterval,omitempty" validate:"omitempty,oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// ExpirationValue
	// Expiration value
	ExpirationValue *int `json:"expirationValue,omitempty"`

	// Id
	// the identifier of the subscription package
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type UserConfiguration struct {
	// Address
	// Address
	// Constraints:
	//    - nullable
	//    - max:256
	//    - min:2
	Address *string `json:"address,omitempty" validate:"omitempty,max=256,min=2"`

	// City
	// City
	// Constraints:
	//    - nullable
	//    - max:50
	//    - min:2
	City *string `json:"city,omitempty" validate:"omitempty,max=50,min=2"`

	// CountryName
	// Country
	CountryName *string `json:"countryName,omitempty"`

	// CountryShortName
	// Country
	CountryShortName *string `json:"countryShortName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	CredentialsGuestPassDto *UserConfigurationCredentialsGuestPassDtoType `json:"credentialsGuestPassDto,omitempty"`

	// Email
	// Email
	Email *string `json:"email,omitempty"`

	// FirstName
	// First name
	// Constraints:
	//    - nullable
	//    - max:32
	//    - min:2
	FirstName *string `json:"firstName,omitempty" validate:"omitempty,max=32,min=2"`

	// IsDisabled
	// Is Disabled
	// Constraints:
	//    - nullable
	//    - oneof:[NO,YES]
	IsDisabled *string `json:"isDisabled,omitempty" validate:"omitempty,oneof=NO YES"`

	// LastName
	// Last Name
	// Constraints:
	//    - nullable
	//    - max:32
	//    - min:2
	LastName *string `json:"lastName,omitempty" validate:"omitempty,max=32,min=2"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// PackageExpirationDate
	// Package Expiration Date
	PackageExpirationDate *int `json:"packageExpirationDate,omitempty"`

	// PackageExpirationInterval
	// Package Expiration Interval
	// Constraints:
	//    - nullable
	//    - oneof:[HOUR,DAY,WEEK,MONTH,YEAR,NEVER]
	PackageExpirationInterval *string `json:"packageExpirationInterval,omitempty" validate:"omitempty,oneof=HOUR DAY WEEK MONTH YEAR NEVER"`

	// PackageExpirationValue
	// Package Expiration Value
	PackageExpirationValue *int `json:"packageExpirationValue,omitempty"`

	// PackageStatus
	// Package Status
	// Constraints:
	//    - nullable
	//    - oneof:[DEPLETED,AVAILABLE,EXPIRED,TERMINATED,REMOVED]
	PackageStatus *string `json:"packageStatus,omitempty" validate:"omitempty,oneof=DEPLETED AVAILABLE EXPIRED TERMINATED REMOVED"`

	// Phone
	// Phone
	// Constraints:
	//    - nullable
	//    - max:32
	//    - min:2
	Phone *string `json:"phone,omitempty" validate:"omitempty,max=32,min=2"`

	// Remark
	// Remark
	// Constraints:
	//    - nullable
	//    - max:32
	//    - min:2
	Remark *string `json:"remark,omitempty" validate:"omitempty,max=32,min=2"`

	// State
	// State
	// Constraints:
	//    - nullable
	//    - max:32
	//    - min:2
	State *string `json:"state,omitempty" validate:"omitempty,max=32,min=2"`

	SubscriberPackage *common.GenericRef `json:"subscriberPackage,omitempty"`

	// UserName
	// User Name
	// Constraints:
	//    - nullable
	//    - max:64
	//    - min:2
	UserName *string `json:"userName,omitempty" validate:"omitempty,max=64,min=2"`

	UsernamePasswordCredentialsImplDto *UsernamePasswordCredentialsImplDto `json:"usernamePasswordCredentialsImplDto,omitempty"`

	// ZipCode
	// Zip Code
	// Constraints:
	//    - nullable
	//    - max:32
	//    - min:2
	ZipCode *string `json:"zipCode,omitempty" validate:"omitempty,max=32,min=2"`
}

type UserConfigurationCredentialsGuestPassDtoType struct {
	// AuthenticationMethod
	// Authentication method of credential
	// Constraints:
	//    - nullable
	//    - oneof:[GUEST_PASS]
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" validate:"omitempty,oneof=GUEST_PASS"`

	// AutoGeneratePassword
	// Pass generation
	AutoGeneratePassword *bool `json:"autoGeneratePassword,omitempty"`

	Comment *string `json:"comment,omitempty"`

	// CreationDate
	// Creation Date
	CreationDate *int `json:"creationDate,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// DisplayName
	// filter identity user list by display name
	DisplayName *string `json:"displayName,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// ExpirationDate
	// Expiration date and time
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// ExpirationUnit
	// Constraints:
	//    - nullable
	//    - oneof:[HOUR,DAY,WEEK]
	ExpirationUnit *string `json:"expirationUnit,omitempty" validate:"omitempty,oneof=HOUR DAY WEEK"`

	ExpirationValue *int `json:"expirationValue,omitempty"`

	ExpireAfterIfNotUsed *int `json:"expireAfterIfNotUsed,omitempty"`

	// GuestExpiration
	// Expiration time of guest pass
	GuestExpiration *int `json:"guestExpiration,omitempty"`

	// Key
	// Key of this guest pass
	Key *string `json:"key,omitempty"`

	// LoginPassword
	// Login Password
	LoginPassword *string `json:"loginPassword,omitempty"`

	MacAddressList []string `json:"macAddressList,omitempty"`

	// MaxDevices
	// Maximum number of allowed device
	MaxDevices *int `json:"maxDevices,omitempty"`

	// ServiceProviderId
	// Service Provider Id
	ServiceProviderId *string `json:"serviceProviderId,omitempty"`

	// SessionUnit
	// Constraints:
	//    - nullable
	//    - oneof:[MIN,HOUR,DAY,WEEK]
	SessionUnit *string `json:"sessionUnit,omitempty" validate:"omitempty,oneof=MIN HOUR DAY WEEK"`

	SessionValue *int `json:"sessionValue,omitempty"`

	// UserKey
	// user ID of the identity guest pass
	UserKey *string `json:"userKey,omitempty"`

	// UserName
	// Username of this guest pass
	UserName *string `json:"userName,omitempty"`

	// Wlan
	// WLAN Id
	Wlan *string `json:"wlan,omitempty"`

	// WlanName
	// WLAN Name
	WlanName *string `json:"wlanName,omitempty"`
}

type UserList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*IdentityUserSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type UsernamePasswordCredentialsImplDto struct {
	// AuthenticationMethod
	// Authentication Method
	// Constraints:
	//    - nullable
	//    - oneof:[USERNAME_PASSWORD,GUEST_PASS,MAC_WLAN_DPSK,MO,REMOTE,OAUTH2]
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" validate:"omitempty,oneof=USERNAME_PASSWORD GUEST_PASS MAC_WLAN_DPSK MO REMOTE OAUTH2"`

	AuthenticationServerConfig *AuthenticationServerConfig `json:"authenticationServerConfig,omitempty"`

	// CreationDate
	// Creation Date
	CreationDate *int `json:"creationDate,omitempty"`

	// ExpirationDate
	// Expiration Date
	ExpirationDate *int `json:"expirationDate,omitempty"`

	// Key
	// identifier of the UsernamePasswordCredentialsImplDto
	Key *string `json:"key,omitempty"`

	// LoginName
	// Login Name
	LoginName *string `json:"loginName,omitempty"`

	// LoginPassword
	// Login Password
	LoginPassword *string `json:"loginPassword,omitempty"`

	// PasswordCreation
	// Creation Date of Password
	PasswordCreation *string `json:"passwordCreation,omitempty"`

	// PasswordExpiration
	// Expiration Date of Password
	PasswordExpiration *string `json:"passwordExpiration,omitempty"`

	// ServiceProviderId
	// Service Provider Id
	ServiceProviderId *string `json:"serviceProviderId,omitempty"`
}

