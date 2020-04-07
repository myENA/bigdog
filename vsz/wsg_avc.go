package vsz

// API Version: v9_0

type WSGAVCAppCategory struct {
	// Id
	// Identifier of the Application Category
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the Application Category
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGAVCAppCategory() *WSGAVCAppCategory {
	m := new(WSGAVCAppCategory)
	return m
}

type WSGAVCAppCategoryList struct {
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
	List []*WSGAVCAppCategory `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAVCAppCategoryList() *WSGAVCAppCategoryList {
	m := new(WSGAVCAppCategoryList)
	return m
}

type WSGAVCApplication struct {
	// AppId
	// Identifier of the Application
	// Constraints:
	//    - nullable
	AppId *string `json:"appId,omitempty"`

	// CatId
	// Identifier of the Application Category
	// Constraints:
	//    - nullable
	CatId *string `json:"catId,omitempty"`

	// Name
	// Name of the Application
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGAVCApplication() *WSGAVCApplication {
	m := new(WSGAVCApplication)
	return m
}

type WSGAVCApplicationList struct {
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
	List []*WSGAVCApplication `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAVCApplicationList() *WSGAVCApplicationList {
	m := new(WSGAVCApplicationList)
	return m
}

type WSGAVCApplicationPolicyProfile struct {
	// ApplicationRules
	// Constraints:
	//    - nullable
	ApplicationRules []*WSGAVCApplicationRule `json:"applicationRules,omitempty" validate:"omitempty,dive"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	// Constraints:
	//    - nullable
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	// Constraints:
	//    - nullable
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy Profile belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the Application Policy Profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`
}

func NewWSGAVCApplicationPolicyProfile() *WSGAVCApplicationPolicyProfile {
	m := new(WSGAVCApplicationPolicyProfile)
	return m
}

type WSGAVCApplicationPolicyProfileList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

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
	List []*WSGAVCApplicationPolicyProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAVCApplicationPolicyProfileList() *WSGAVCApplicationPolicyProfileList {
	m := new(WSGAVCApplicationPolicyProfileList)
	return m
}

type WSGAVCApplicationRule struct {
	// AppId
	// Identifier of the Application from Signature Package
	// Constraints:
	//    - nullable
	AppId *string `json:"appId,omitempty"`

	// ApplicationType
	// Type of the application when ruleType
	// Constraints:
	//    - nullable
	//    - oneof:[SIGNATURE,USER_DEFINED]
	ApplicationType *string `json:"applicationType,omitempty" validate:"omitempty,oneof=SIGNATURE USER_DEFINED"`

	// AppName
	// Name of the Application from Signature Package
	// Constraints:
	//    - nullable
	AppName *string `json:"appName,omitempty"`

	// CatId
	// Identifier of the Application Category from Signature Package (If applicationType is UserDefind, the catId is 32768)
	// Constraints:
	//    - nullable
	CatId *string `json:"catId,omitempty"`

	// CatName
	// Name of the Application Category from Signature Package
	// Constraints:
	//    - nullable
	CatName *string `json:"catName,omitempty"`

	// ClassificationType
	// QoS downlink classification type
	// Constraints:
	//    - nullable
	//    - oneof:[VOICE,VIDEO,BEST_EFFORT,BACKGROUND]
	ClassificationType *string `json:"classificationType,omitempty" validate:"omitempty,oneof=VOICE VIDEO BEST_EFFORT BACKGROUND"`

	// Downlink
	// Downlink rate limiting (unit: Kbps)
	// Constraints:
	//    - nullable
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

	// Priority
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`

	// RuleType
	// Type of the application rule
	// Constraints:
	//    - nullable
	//    - oneof:[DENY,QOS,RATE_LIMITING]
	RuleType *string `json:"ruleType,omitempty" validate:"omitempty,oneof=DENY QOS RATE_LIMITING"`

	// Uplink
	// Uplink rate limiting (unit: Kbps)
	// Constraints:
	//    - nullable
	Uplink *int `json:"uplink,omitempty"`
}

func NewWSGAVCApplicationRule() *WSGAVCApplicationRule {
	m := new(WSGAVCApplicationRule)
	return m
}

type WSGAVCCreateApplicationPolicyProfile struct {
	// ApplicationRules
	// Constraints:
	//    - required
	ApplicationRules []*WSGAVCApplicationRule `json:"applicationRules" validate:"required,dive"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	// Constraints:
	//    - nullable
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	// Constraints:
	//    - nullable
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Application Policy Profile belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

func NewWSGAVCCreateApplicationPolicyProfile() *WSGAVCCreateApplicationPolicyProfile {
	m := new(WSGAVCCreateApplicationPolicyProfile)
	return m
}

type WSGAVCCreateUserDefinedProfile struct {
	// DestIp
	// Constraints:
	//    - nullable
	DestIp *WSGCommonIpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort" validate:"required,gte=1,lte=65535"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Netmask
	// Constraints:
	//    - nullable
	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

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

func NewWSGAVCCreateUserDefinedProfile() *WSGAVCCreateUserDefinedProfile {
	m := new(WSGAVCCreateUserDefinedProfile)
	return m
}

type WSGAVCDeleteBulk struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGAVCDeleteBulk() *WSGAVCDeleteBulk {
	m := new(WSGAVCDeleteBulk)
	return m
}

type WSGAVCModifyApplicationPolicyProfile struct {
	// ApplicationRules
	// Constraints:
	//    - nullable
	ApplicationRules []*WSGAVCApplicationRule `json:"applicationRules,omitempty" validate:"omitempty,dive"`

	// AvcEventEnable
	// Send ARC logs from AP to SmartZone
	// Constraints:
	//    - nullable
	AvcEventEnable *bool `json:"avcEventEnable,omitempty"`

	// AvcLogEnable
	// Send ARC logs from AP to external syslog server
	// Constraints:
	//    - nullable
	AvcLogEnable *bool `json:"avcLogEnable,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGAVCModifyApplicationPolicyProfile() *WSGAVCModifyApplicationPolicyProfile {
	m := new(WSGAVCModifyApplicationPolicyProfile)
	return m
}

type WSGAVCModifyUserDefinedProfile struct {
	// DestIp
	// Constraints:
	//    - nullable
	DestIp *WSGCommonIpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Netmask
	// Constraints:
	//    - nullable
	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

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

func NewWSGAVCModifyUserDefinedProfile() *WSGAVCModifyUserDefinedProfile {
	m := new(WSGAVCModifyUserDefinedProfile)
	return m
}

type WSGAVCSignaturePackage struct {
	// FileName
	// Name of the Signature Package
	// Constraints:
	//    - nullable
	FileName *string `json:"fileName,omitempty"`

	// Id
	// Identifier of the Signature Package
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Size
	// Size of the Signature Package
	// Constraints:
	//    - nullable
	Size *int `json:"size,omitempty"`

	// Version
	// Version of the Signature Package
	// Constraints:
	//    - nullable
	Version *string `json:"version,omitempty"`
}

func NewWSGAVCSignaturePackage() *WSGAVCSignaturePackage {
	m := new(WSGAVCSignaturePackage)
	return m
}

type WSGAVCUserDefinedProfile struct {
	// AppId
	// AppId for Application Policy's User defined rule type
	// Constraints:
	//    - nullable
	AppId *int `json:"appId,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DestIp
	// Constraints:
	//    - nullable
	DestIp *WSGCommonIpAddress `json:"destIp,omitempty"`

	// DestPort
	// Destination Port of User Defined Profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:65535
	DestPort *int `json:"destPort,omitempty" validate:"omitempty,gte=1,lte=65535"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the User Defined Profile belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the User Defined Profile
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Netmask
	// Constraints:
	//    - nullable
	Netmask *WSGCommonSubNetMask `json:"netmask,omitempty"`

	// Protocol
	// Protocol of User Defined Profile
	// Constraints:
	//    - nullable
	//    - oneof:[TCP,UDP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=TCP UDP"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Type
	// Type of the User Defined Profile
	// Constraints:
	//    - nullable
	//    - oneof:[IP_WITH_PORT,PORT_ONLY]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IP_WITH_PORT PORT_ONLY"`
}

func NewWSGAVCUserDefinedProfile() *WSGAVCUserDefinedProfile {
	m := new(WSGAVCUserDefinedProfile)
	return m
}

type WSGAVCUserDefinedProfileList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

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
	List []*WSGAVCUserDefinedProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAVCUserDefinedProfileList() *WSGAVCUserDefinedProfileList {
	m := new(WSGAVCUserDefinedProfileList)
	return m
}
