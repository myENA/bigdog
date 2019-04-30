package avc

// API Version: v8_0

type AppCategory struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type AppCategoryList struct {
	FirstIndex *int           `json:"firstIndex,omitempty"`
	HasMore    *bool          `json:"hasMore,omitempty"`
	List       []*AppCategory `json:"list,omitempty"`
	TotalCount *int           `json:"totalCount,omitempty"`
}

type Application struct {
	AppID *string `json:"appId,omitempty"`
	CatID *string `json:"catId,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type ApplicationList struct {
	FirstIndex *int           `json:"firstIndex,omitempty"`
	HasMore    *bool          `json:"hasMore,omitempty"`
	List       []*Application `json:"list,omitempty"`
	TotalCount *int           `json:"totalCount,omitempty"`
}

type ApplicationPolicyProfile struct {
	ApplicationRules []*ApplicationRule `json:"applicationRules,omitempty"`
	AvcEventEnable   *bool              `json:"avcEventEnable,omitempty"`
	AvcLogEnable     *bool              `json:"avcLogEnable,omitempty"`
	CreateDateTime   *int               `json:"createDateTime,omitempty"`
	CreatorID        *string            `json:"creatorId,omitempty"`
	CreatorUsername  *string            `json:"creatorUsername,omitempty"`
	Description      *string            `json:"description,omitempty"`
	DomainID         *string            `json:"domainId,omitempty"`
	ID               *string            `json:"id,omitempty"`
	ModifiedDateTime *int               `json:"modifiedDateTime,omitempty"`
	ModifierID       *string            `json:"modifierId,omitempty"`
	ModifierUsername *string            `json:"modifierUsername,omitempty"`
	Name             *string            `json:"name,omitempty"`
	TenantID         *string            `json:"tenantId,omitempty"`
}

type ApplicationPolicyProfileList struct {
	Extra      *common.RBACMetadata        `json:"extra,omitempty"`
	FirstIndex *int                        `json:"firstIndex,omitempty"`
	HasMore    *bool                       `json:"hasMore,omitempty"`
	List       []*ApplicationPolicyProfile `json:"list,omitempty"`
	TotalCount *int                        `json:"totalCount,omitempty"`
}

type ApplicationRule struct {
	AppID              *string `json:"appId,omitempty"`
	ApplicationType    *string `json:"applicationType,omitempty"`
	AppName            *string `json:"appName,omitempty"`
	CatID              *string `json:"catId,omitempty"`
	CatName            *string `json:"catName,omitempty"`
	ClassificationType *string `json:"classificationType,omitempty"`
	Downlink           *int    `json:"downlink,omitempty"`
	MarkingPriority    *string `json:"markingPriority,omitempty"`
	MarkingType        *string `json:"markingType,omitempty"`
	Priority           *int    `json:"priority,omitempty"`
	RuleType           *string `json:"ruleType,omitempty"`
	Uplink             *int    `json:"uplink,omitempty"`
}

type CreateApplicationPolicyProfile struct {
	ApplicationRules []*ApplicationRule `json:"applicationRules,omitempty"`
	AvcEventEnable   *bool              `json:"avcEventEnable,omitempty"`
	AvcLogEnable     *bool              `json:"avcLogEnable,omitempty"`
	Description      *string            `json:"description,omitempty"`
	DomainID         *string            `json:"domainId,omitempty"`
	Name             *string            `json:"name,omitempty"`
}

type CreateUserDefinedProfile struct {
	DestIP   *string `json:"destIp,omitempty"`
	DestPort *int    `json:"destPort,omitempty"`
	DomainID *string `json:"domainId,omitempty"`
	Name     *string `json:"name,omitempty"`
	Netmask  *string `json:"netmask,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Type     *string `json:"type,omitempty"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyApplicationPolicyProfile struct {
	ApplicationRules []*ApplicationRule `json:"applicationRules,omitempty"`
	AvcEventEnable   *bool              `json:"avcEventEnable,omitempty"`
	AvcLogEnable     *bool              `json:"avcLogEnable,omitempty"`
	Description      *string            `json:"description,omitempty"`
	Name             *string            `json:"name,omitempty"`
}

type ModifyUserDefinedProfile struct {
	DestIP   *string `json:"destIp,omitempty"`
	DestPort *int    `json:"destPort,omitempty"`
	Name     *string `json:"name,omitempty"`
	Netmask  *string `json:"netmask,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Type     *string `json:"type,omitempty"`
}

type SignaturePackage struct {
	FileName *string `json:"fileName,omitempty"`
	ID       *string `json:"id,omitempty"`
	Size     *int    `json:"size,omitempty"`
	Version  *string `json:"version,omitempty"`
}

type UserDefinedProfile struct {
	AppID            *int    `json:"appId,omitempty"`
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	DestIP           *string `json:"destIp,omitempty"`
	DestPort         *int    `json:"destPort,omitempty"`
	DomainID         *string `json:"domainId,omitempty"`
	ID               *string `json:"id,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	Name             *string `json:"name,omitempty"`
	Netmask          *string `json:"netmask,omitempty"`
	Protocol         *string `json:"protocol,omitempty"`
	TenantID         *string `json:"tenantId,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type UserDefinedProfileList struct {
	Extra      *common.RBACMetadata  `json:"extra,omitempty"`
	FirstIndex *int                  `json:"firstIndex,omitempty"`
	HasMore    *bool                 `json:"hasMore,omitempty"`
	List       []*UserDefinedProfile `json:"list,omitempty"`
	TotalCount *int                  `json:"totalCount,omitempty"`
}
