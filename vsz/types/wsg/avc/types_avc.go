package avc

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AppCategory struct {
	// Id
	// Identifier of the Application Category
	Id *string `json:"id,omitempty"`

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
	// AppId
	// Identifier of the Application
	AppId *string `json:"appId,omitempty"`

	// CatId
	// Identifier of the Application Category
	CatId *string `json:"catId,omitempty"`

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

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy
	// Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the Application Policy Profile
	Id *string `json:"id,omitempty"`

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

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

type ApplicationPolicyProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApplicationPolicyProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApplicationRule struct {
	// AppId
	// Identifier of the Application from Signature Package
	AppId *string `json:"appId,omitempty"`

	// ApplicationType
	// Type of the application when ruleType
	ApplicationType *string `json:"applicationType,omitempty"`

	// AppName
	// Name of the Application from Signature Package
	AppName *string `json:"appName,omitempty"`

	// CatId
	// Identifier of the Application Category from Signature Package (If applicationType is UserDefind, the
	// catId is 32768)
	CatId *string `json:"catId,omitempty"`

	// CatName
	// Name of the Application Category from Signature Package
	CatName *string `json:"catName,omitempty"`

	// ClassificationType
	// QoS downlink classification type
	ClassificationType *string `json:"classificationType,omitempty"`

	// Downlink
	// Downlink rate limiting (unit: Kbps)
	Downlink *int `json:"downlink,omitempty"`

	// MarkingPriority
	// QoS uplink marking priority
	MarkingPriority *string `json:"markingPriority,omitempty"`

	// MarkingType
	// QoS uplink marking type
	MarkingType *string `json:"markingType,omitempty"`

	Priority *int `json:"priority,omitempty"`

	// RuleType
	// Type of the application rule
	RuleType *string `json:"ruleType,omitempty"`

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

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy
	// Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`
}

type CreateUserDefinedProfile struct {
	DestIp *common.IpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	DestPort *int `json:"destPort,omitempty" validate:"required"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile
	// belongs
	DomainId *string `json:"domainId,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	Protocol *string `json:"protocol,omitempty" validate:"required"`

	// Type
	// Type of the User Defined Profile
	Type *string `json:"type,omitempty" validate:"required"`
}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
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
	DestIp *common.IpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	DestPort *int `json:"destPort,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	Protocol *string `json:"protocol,omitempty"`

	// Type
	// Type of the User Defined Profile
	Type *string `json:"type,omitempty"`
}

type SignaturePackage struct {
	// FileName
	// Name of the Signature Package
	FileName *string `json:"fileName,omitempty"`

	// Id
	// Identifier of the Signature Package
	Id *string `json:"id,omitempty"`

	// Size
	// Size of the Signature Package
	Size *int `json:"size,omitempty"`

	// Version
	// Version of the Signature Package
	Version *string `json:"version,omitempty"`
}

type UserDefinedProfile struct {
	// AppId
	// AppId for Application Policy's User defined rule type
	AppId *int `json:"appId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	DestIp *common.IpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	DestPort *int `json:"destPort,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile
	// belongs
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the User Defined Profile
	Id *string `json:"id,omitempty"`

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

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	Protocol *string `json:"protocol,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Type
	// Type of the User Defined Profile
	Type *string `json:"type,omitempty"`
}

type UserDefinedProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UserDefinedProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
