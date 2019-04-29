package scguser

// API Version: v8_0

type PatchScgUserGroup struct {
	AccountSecurityProfileID *string           `json:"accountSecurityProfileId,omitempty"`
	CreateDateTime           *int              `json:"createDateTime,omitempty"`
	CreatorID                *string           `json:"creatorId,omitempty"`
	CreatorUsername          *string           `json:"creatorUsername,omitempty"`
	Description              *string           `json:"description,omitempty"`
	DomainID                 *string           `json:"domainId,omitempty"`
	ID                       *string           `json:"id,omitempty"`
	ModifiedDateTime         *int              `json:"modifiedDateTime,omitempty"`
	ModifierID               *string           `json:"modifierId,omitempty"`
	ModifierUsername         *string           `json:"modifierUsername,omitempty"`
	Name                     *string           `json:"name,omitempty"`
	Permissions              []*Permissions    `json:"permissions,omitempty"`
	ResourceGroups           []*ResourceGroups `json:"resourceGroups,omitempty"`
	Role                     *string           `json:"role,omitempty"`
	TenantID                 *string           `json:"tenantId,omitempty"`
	Users                    []*Users          `json:"users,omitempty"`
}

type QueryCriteria struct {
	Attributes      []string               `json:"attributes,omitempty"`
	Criteria        *string                `json:"criteria,omitempty"`
	ExpandDomains   *bool                  `json:"expandDomains,omitempty"`
	ExtraFilters    []*ExtraFilters        `json:"extraFilters,omitempty"`
	ExtraNotFilters []*ExtraNotFilters     `json:"extraNotFilters,omitempty"`
	ExtraTimeRange  *ExtraTimeRange        `json:"extraTimeRange,omitempty"`
	Filters         []*Filters             `json:"filters,omitempty"`
	FullTextSearch  *FullTextSearch        `json:"fullTextSearch,omitempty"`
	Limit           *int                   `json:"limit,omitempty"`
	Options         map[string]interface{} `json:"options,omitempty"`
	Page            *int                   `json:"page,omitempty"`
	Query           *string                `json:"query,omitempty"`
	SortInfo        *SortInfo              `json:"sortInfo,omitempty"`
}

type QueryCriteriaFiltersType struct {
	Operator *string `json:"operator,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type ScgUser struct {
	CompanyName        *string `json:"companyName,omitempty"`
	CreateDateTime     *int    `json:"createDateTime,omitempty"`
	CreatorID          *string `json:"creatorId,omitempty"`
	CreatorUsername    *string `json:"creatorUsername,omitempty"`
	DomainID           *string `json:"domainId,omitempty"`
	Email              *string `json:"email,omitempty"`
	Enabled            *int    `json:"enabled,omitempty"`
	ID                 *string `json:"id,omitempty"`
	Locked             *int    `json:"locked,omitempty"`
	ModifiedDateTime   *int    `json:"modifiedDateTime,omitempty"`
	ModifierID         *string `json:"modifierId,omitempty"`
	ModifierUsername   *string `json:"modifierUsername,omitempty"`
	Passphrase         *string `json:"passphrase,omitempty"`
	PasswordExpiration *int    `json:"passwordExpiration,omitempty"`
	PasswordReuse      *int    `json:"passwordReuse,omitempty"`
	Phone              *string `json:"phone,omitempty"`
	RealName           *string `json:"realName,omitempty"`
	SessionIdle        *int    `json:"sessionIdle,omitempty"`
	TenantUUID         *string `json:"tenantUUID,omitempty"`
	Title              *string `json:"title,omitempty"`
	UserName           *string `json:"userName,omitempty"`
}

type ScgUserAuditID struct {
	ID *string `json:"id,omitempty"`
}

type ScgUserGroup struct {
	AccountSecurityProfileID   *string           `json:"accountSecurityProfileId,omitempty"`
	AccountSecurityProfileName *string           `json:"accountSecurityProfileName,omitempty"`
	CreateDateTime             *int              `json:"createDateTime,omitempty"`
	CreatorID                  *string           `json:"creatorId,omitempty"`
	CreatorUsername            *string           `json:"creatorUsername,omitempty"`
	Description                *string           `json:"description,omitempty"`
	DomainID                   *string           `json:"domainId,omitempty"`
	ID                         *string           `json:"id,omitempty"`
	IsFactoryDefault           *bool             `json:"isFactoryDefault,omitempty"`
	ModifiedDateTime           *int              `json:"modifiedDateTime,omitempty"`
	ModifierID                 *string           `json:"modifierId,omitempty"`
	ModifierUsername           *string           `json:"modifierUsername,omitempty"`
	Name                       *string           `json:"name,omitempty"`
	Permissions                []*Permissions    `json:"permissions,omitempty"`
	ResourceGroups             []*ResourceGroups `json:"resourceGroups,omitempty"`
	Role                       *string           `json:"role,omitempty"`
	TenantID                   *string           `json:"tenantId,omitempty"`
	Users                      []*Users          `json:"users,omitempty"`
}

type ScgUserGroupAuditID struct {
	ID *string `json:"id,omitempty"`
}

type ScgUserGroupList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type ScgUserGroupPermission struct {
	Access           *string  `json:"access,omitempty"`
	Display          *string  `json:"display,omitempty"`
	Ids              []string `json:"ids,omitempty"`
	Items            []*Items `json:"items,omitempty"`
	ItemsDescription []string `json:"itemsDescription,omitempty"`
	Resource         *string  `json:"resource,omitempty"`
}

type ScgUserGroupPermissionItemsType struct {
	Access   *string `json:"access,omitempty"`
	Display  *string `json:"display,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type ScgUserGroupPermissionList struct {
	Extra      interface{} `json:"extra,omitempty"`
	FirstIndex *int        `json:"firstIndex,omitempty"`
	HasMore    *bool       `json:"hasMore,omitempty"`
	List       []*List     `json:"list,omitempty"`
	TotalCount *int        `json:"totalCount,omitempty"`
}

type ScgUserGroupPermissionListExtraType struct {
	IsSuperAdmin         *bool `json:"isSuperAdmin,omitempty"`
	IsSuperAdminOfDomain *bool `json:"isSuperAdminOfDomain,omitempty"`
}

type ScgUserGroupPermissionWithoutDetailItems struct {
	Access   *string  `json:"access,omitempty"`
	Display  *string  `json:"display,omitempty"`
	Ids      []string `json:"ids,omitempty"`
	Resource *string  `json:"resource,omitempty"`
}

type ScgUserGroupResourceGroup struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

type ScgUserGroupRoleLabelValue struct {
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ScgUserGroupRoleLabelValueList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type ScgUserList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}
