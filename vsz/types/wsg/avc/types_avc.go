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

func NewAppCategory() *AppCategory {
	appCategoryType := new(AppCategory)
	return appCategoryType
}

func NewAppCategoryWithDefaults() *AppCategory {
	appCategoryType := new(AppCategory)
	return appCategoryType
}

type AppCategoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*AppCategory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAppCategoryList() *AppCategoryList {
	appCategoryListType := new(AppCategoryList)
	return appCategoryListType
}

func NewAppCategoryListWithDefaults() *AppCategoryList {
	appCategoryListType := new(AppCategoryList)
	return appCategoryListType
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

func NewApplication() *Application {
	applicationType := new(Application)
	return applicationType
}

func NewApplicationWithDefaults() *Application {
	applicationType := new(Application)
	return applicationType
}

type ApplicationList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Application `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApplicationList() *ApplicationList {
	applicationListType := new(ApplicationList)
	return applicationListType
}

func NewApplicationListWithDefaults() *ApplicationList {
	applicationListType := new(ApplicationList)
	return applicationListType
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
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy Profile belongs
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

func NewApplicationPolicyProfile() *ApplicationPolicyProfile {
	applicationPolicyProfileType := new(ApplicationPolicyProfile)
	return applicationPolicyProfileType
}

func NewApplicationPolicyProfileWithDefaults() *ApplicationPolicyProfile {
	applicationPolicyProfileType := new(ApplicationPolicyProfile)
	return applicationPolicyProfileType
}

type ApplicationPolicyProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApplicationPolicyProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApplicationPolicyProfileList() *ApplicationPolicyProfileList {
	applicationPolicyProfileListType := new(ApplicationPolicyProfileList)
	return applicationPolicyProfileListType
}

func NewApplicationPolicyProfileListWithDefaults() *ApplicationPolicyProfileList {
	applicationPolicyProfileListType := new(ApplicationPolicyProfileList)
	return applicationPolicyProfileListType
}

type ApplicationRule struct {
	// AppId
	// Identifier of the Application from Signature Package
	AppId *string `json:"appId,omitempty"`

	// ApplicationType
	// Type of the application when ruleType
	// Constraints:
	//    - nullable
	//    - oneof:[SIGNATURE,USER_DEFINED]
	ApplicationType *string `json:"applicationType,omitempty" validate:"omitempty,oneof=SIGNATURE USER_DEFINED"`

	// AppName
	// Name of the Application from Signature Package
	AppName *string `json:"appName,omitempty"`

	// CatId
	// Identifier of the Application Category from Signature Package (If applicationType is UserDefind, the catId is 32768)
	CatId *string `json:"catId,omitempty"`

	// CatName
	// Name of the Application Category from Signature Package
	CatName *string `json:"catName,omitempty"`

	// ClassificationType
	// QoS downlink classification type
	// Constraints:
	//    - nullable
	//    - oneof:[VOICE,VIDEO,BEST_EFFORT,BACKGROUND]
	ClassificationType *string `json:"classificationType,omitempty" validate:"omitempty,oneof=VOICE VIDEO BEST_EFFORT BACKGROUND"`

	// Downlink
	// Downlink rate limiting (unit: Kbps)
	Downlink *int `json:"downlink,omitempty"`

	// MarkingPriority
	// QoS uplink marking priority
	// Constraints:
	//    - nullable
	//    - oneof:[IEEE802_1p,DSCP,BOTH]
	MarkingPriority *string `json:"markingPriority,omitempty" validate:"omitempty,oneof=IEEE802_1p DSCP BOTH"`

	// MarkingType
	// QoS uplink marking type
	// Constraints:
	//    - nullable
	//    - oneof:[VOICE,VIDEO,BEST_EFFORT,BACKGROUND]
	MarkingType *string `json:"markingType,omitempty" validate:"omitempty,oneof=VOICE VIDEO BEST_EFFORT BACKGROUND"`

	Priority *int `json:"priority,omitempty"`

	// RuleType
	// Type of the application rule
	// Constraints:
	//    - nullable
	//    - oneof:[DENY,QOS,RATE_LIMITING]
	RuleType *string `json:"ruleType,omitempty" validate:"omitempty,oneof=DENY QOS RATE_LIMITING"`

	// Uplink
	// Uplink rate limiting (unit: Kbps)
	Uplink *int `json:"uplink,omitempty"`
}

func NewApplicationRule() *ApplicationRule {
	applicationRuleType := new(ApplicationRule)
	return applicationRuleType
}

func NewApplicationRuleWithDefaults() *ApplicationRule {
	applicationRuleType := new(ApplicationRule)
	return applicationRuleType
}

type CreateApplicationPolicyProfile struct {
	// ApplicationRules
	// Constraints:
	//    - required
	ApplicationRules []*ApplicationRule `json:"applicationRules" validate:"required"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`
}

func NewCreateApplicationPolicyProfile() *CreateApplicationPolicyProfile {
	createApplicationPolicyProfileType := new(CreateApplicationPolicyProfile)
	return createApplicationPolicyProfileType
}

func NewCreateApplicationPolicyProfileWithDefaults() *CreateApplicationPolicyProfile {
	createApplicationPolicyProfileType := new(CreateApplicationPolicyProfile)
	return createApplicationPolicyProfileType
}

type CreateUserDefinedProfile struct {
	DestIp *common.IpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort" validate:"required,gte=1,lte=65535"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile belongs
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	// Constraints:
	//    - required
	//    - oneof:[TCP,UDP]
	Protocol *string `json:"protocol" validate:"required,oneof=TCP UDP"`

	// Type
	// Type of the User Defined Profile
	// Constraints:
	//    - required
	//    - oneof:[IP_WITH_PORT,PORT_ONLY]
	Type *string `json:"type" validate:"required,oneof=IP_WITH_PORT PORT_ONLY"`
}

func NewCreateUserDefinedProfile() *CreateUserDefinedProfile {
	createUserDefinedProfileType := new(CreateUserDefinedProfile)
	return createUserDefinedProfileType
}

func NewCreateUserDefinedProfileWithDefaults() *CreateUserDefinedProfile {
	createUserDefinedProfileType := new(CreateUserDefinedProfile)
	return createUserDefinedProfileType
}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulk() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
}

func NewDeleteBulkWithDefaults() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
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

func NewModifyApplicationPolicyProfile() *ModifyApplicationPolicyProfile {
	modifyApplicationPolicyProfileType := new(ModifyApplicationPolicyProfile)
	return modifyApplicationPolicyProfileType
}

func NewModifyApplicationPolicyProfileWithDefaults() *ModifyApplicationPolicyProfile {
	modifyApplicationPolicyProfileType := new(ModifyApplicationPolicyProfile)
	return modifyApplicationPolicyProfileType
}

type ModifyUserDefinedProfile struct {
	DestIp *common.IpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	Name *common.NormalName `json:"name,omitempty"`

	Netmask *common.SubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP"`

	// Type
	// Type of the User Defined Profile
	// Constraints:
	//    - nullable
	//    - oneof:[IP_WITH_PORT,PORT_ONLY]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IP_WITH_PORT PORT_ONLY"`
}

func NewModifyUserDefinedProfile() *ModifyUserDefinedProfile {
	modifyUserDefinedProfileType := new(ModifyUserDefinedProfile)
	return modifyUserDefinedProfileType
}

func NewModifyUserDefinedProfileWithDefaults() *ModifyUserDefinedProfile {
	modifyUserDefinedProfileType := new(ModifyUserDefinedProfile)
	return modifyUserDefinedProfileType
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

func NewSignaturePackage() *SignaturePackage {
	signaturePackageType := new(SignaturePackage)
	return signaturePackageType
}

func NewSignaturePackageWithDefaults() *SignaturePackage {
	signaturePackageType := new(SignaturePackage)
	return signaturePackageType
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
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile belongs
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
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Type
	// Type of the User Defined Profile
	// Constraints:
	//    - nullable
	//    - oneof:[IP_WITH_PORT,PORT_ONLY]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IP_WITH_PORT PORT_ONLY"`
}

func NewUserDefinedProfile() *UserDefinedProfile {
	userDefinedProfileType := new(UserDefinedProfile)
	return userDefinedProfileType
}

func NewUserDefinedProfileWithDefaults() *UserDefinedProfile {
	userDefinedProfileType := new(UserDefinedProfile)
	return userDefinedProfileType
}

type UserDefinedProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UserDefinedProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewUserDefinedProfileList() *UserDefinedProfileList {
	userDefinedProfileListType := new(UserDefinedProfileList)
	return userDefinedProfileListType
}

func NewUserDefinedProfileListWithDefaults() *UserDefinedProfileList {
	userDefinedProfileListType := new(UserDefinedProfileList)
	return userDefinedProfileListType
}
