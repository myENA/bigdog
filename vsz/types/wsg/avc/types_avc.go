package avc

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AppCategory struct {
	// ID
	// Identifier of the Application Category
	ID *string `json:"id,omitempty"`

	// Name
	// Name of the Application Category
	Name *string `json:"name,omitempty"`
}

type AppCategoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AppCategory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type Application struct {
	// AppID
	// Identifier of the Application
	AppID *string `json:"appId,omitempty"`

	// CatID
	// Identifier of the Application Category
	CatID *string `json:"catId,omitempty"`

	// Name
	// Name of the Application
	Name *string `json:"name,omitempty"`
}

type ApplicationList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Application `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApplicationPolicyProfile struct {
	ApplicationRules []*ApplicationRule `json:"applicationRules,omitempty"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy
	// Profile belongs
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the Application Policy Profile
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`
}

type ApplicationPolicyProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApplicationPolicyProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApplicationRule struct {
	// AppID
	// Identifier of the Application from Signature Package
	AppID *string `json:"appId,omitempty"`

	// ApplicationType
	// Type of the application when ruleType
	ApplicationType *string `json:"applicationType,omitempty" validate:"oneof=SIGNATURE USER_DEFINED"`

	// AppName
	// Name of the Application from Signature Package
	AppName *string `json:"appName,omitempty"`

	// CatID
	// Identifier of the Application Category from Signature Package (If applicationType is UserDefind, the
	// catId is 32768)
	CatID *string `json:"catId,omitempty"`

	// CatName
	// Name of the Application Category from Signature Package
	CatName *string `json:"catName,omitempty"`

	// ClassificationType
	// QoS downlink classification type
	ClassificationType *string `json:"classificationType,omitempty" validate:"oneof=VOICE VIDEO BEST_EFFORT BACKGROUND"`

	// Downlink
	// Downlink rate limiting (unit: Kbps)
	Downlink *int `json:"downlink,omitempty"`

	// MarkingPriority
	// QoS uplink marking priority
	MarkingPriority *string `json:"markingPriority,omitempty" validate:"oneof=IEEE802_1p DSCP BOTH"`

	// MarkingType
	// QoS uplink marking type
	MarkingType *string `json:"markingType,omitempty" validate:"oneof=VOICE VIDEO BEST_EFFORT BACKGROUND"`

	Priority *int `json:"priority,omitempty"`

	// RuleType
	// Type of the application rule
	RuleType *string `json:"ruleType,omitempty" validate:"oneof=DENY QOS RATE_LIMITING"`

	// Uplink
	// Uplink rate limiting (unit: Kbps)
	Uplink *int `json:"uplink,omitempty"`
}

type CreateApplicationPolicyProfile struct {
	ApplicationRules []*ApplicationRule `json:"applicationRules,omitempty" validate:"required"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy
	// Profile belongs
	DomainID *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`
}

type CreateUserDefinedProfile struct {
	DestIP *common.IPAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	DestPort *int `json:"destPort,omitempty" validate:"required,gte=1,lte=65535"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	Protocol *string `json:"protocol,omitempty" validate:"required,oneof=TCP UDP"`

	// Type
	// Type of the User Defined Profile
	Type *string `json:"type,omitempty" validate:"required,oneof=IP_WITH_PORT PORT_ONLY"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyApplicationPolicyProfile struct {
	ApplicationRules []*ApplicationRule `json:"applicationRules,omitempty"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type ModifyUserDefinedProfile struct {
	DestIP *common.IPAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	DestPort *int `json:"destPort,omitempty" validate:"gte=1,lte=65535"`

	Name *common.NormalName `json:"name,omitempty"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	Protocol *string `json:"protocol,omitempty" validate:"oneof=TCP UDP"`

	// Type
	// Type of the User Defined Profile
	Type *string `json:"type,omitempty" validate:"oneof=IP_WITH_PORT PORT_ONLY"`
}

type SignaturePackage struct {
	// FileName
	// Name of the Signature Package
	FileName *string `json:"fileName,omitempty"`

	// ID
	// Identifier of the Signature Package
	ID *string `json:"id,omitempty"`

	// Size
	// Size of the Signature Package
	Size *int `json:"size,omitempty"`

	// Version
	// Version of the Signature Package
	Version *string `json:"version,omitempty"`
}

type UserDefinedProfile struct {
	// AppID
	// AppId for Application Policy's User defined rule type
	AppID *int `json:"appId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DestIP *common.IPAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	DestPort *int `json:"destPort,omitempty" validate:"gte=1,lte=65535"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	// ID
	// Identifier of the User Defined Profile
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	Protocol *string `json:"protocol,omitempty" validate:"oneof=TCP UDP"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`

	// Type
	// Type of the User Defined Profile
	Type *string `json:"type,omitempty" validate:"oneof=IP_WITH_PORT PORT_ONLY"`
}

type UserDefinedProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UserDefinedProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
