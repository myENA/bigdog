package vsz

// API Version: v9_0

type WSGURLFilteringCreateUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	// Constraints:
	//    - nullable
	Blacklist []string `json:"blacklist,omitempty" validate:"omitempty,dive"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	// Constraints:
	//    - nullable
	BlockCategories []int `json:"blockCategories,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - required
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel" validate:"required,oneof=NO_ADULT CLEAN_AND_SAFE CHILD_AND_STUDENT_FRIENDLY STRICT CUSTOM"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	// Constraints:
	//    - nullable
	Whitelist []string `json:"whitelist,omitempty" validate:"omitempty,dive"`
}

func NewWSGURLFilteringCreateUrlFilteringPolicy() *WSGURLFilteringCreateUrlFilteringPolicy {
	m := new(WSGURLFilteringCreateUrlFilteringPolicy)
	return m
}

type WSGURLFilteringDeleteBulk struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGURLFilteringDeleteBulk() *WSGURLFilteringDeleteBulk {
	m := new(WSGURLFilteringDeleteBulk)
	return m
}

type WSGURLFilteringModifyUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	// Constraints:
	//    - nullable
	Blacklist []string `json:"blacklist,omitempty" validate:"omitempty,dive"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	// Constraints:
	//    - nullable
	BlockCategories []int `json:"blockCategories,omitempty" validate:"omitempty,dive"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - nullable
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel,omitempty" validate:"omitempty,oneof=NO_ADULT CLEAN_AND_SAFE CHILD_AND_STUDENT_FRIENDLY STRICT CUSTOM"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	// Constraints:
	//    - nullable
	Whitelist []string `json:"whitelist,omitempty" validate:"omitempty,dive"`
}

func NewWSGURLFilteringModifyUrlFilteringPolicy() *WSGURLFilteringModifyUrlFilteringPolicy {
	m := new(WSGURLFilteringModifyUrlFilteringPolicy)
	return m
}

type WSGURLFilteringBlockCategoriesList struct {
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
	List []*WSGURLFilteringBlockCategory `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGURLFilteringBlockCategoriesList() *WSGURLFilteringBlockCategoriesList {
	m := new(WSGURLFilteringBlockCategoriesList)
	return m
}

type WSGURLFilteringBlockCategory struct {
	// Id
	// Identifier of the URL Filtering Category
	// Constraints:
	//    - nullable
	Id *int `json:"id,omitempty"`

	// Name
	// name of the URL Filtering Category
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGURLFilteringBlockCategory() *WSGURLFilteringBlockCategory {
	m := new(WSGURLFilteringBlockCategory)
	return m
}

type WSGURLFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	// Constraints:
	//    - nullable
	Blacklist []string `json:"blacklist,omitempty" validate:"omitempty,dive"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	// Constraints:
	//    - nullable
	BlockCategories []int `json:"blockCategories,omitempty" validate:"omitempty,dive"`

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
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - nullable
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel,omitempty" validate:"omitempty,oneof=NO_ADULT CLEAN_AND_SAFE CHILD_AND_STUDENT_FRIENDLY STRICT CUSTOM"`

	// Id
	// Identifier of the URL Filtering policy
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

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	// Constraints:
	//    - nullable
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	// Constraints:
	//    - nullable
	Whitelist []string `json:"whitelist,omitempty" validate:"omitempty,dive"`
}

func NewWSGURLFilteringPolicy() *WSGURLFilteringPolicy {
	m := new(WSGURLFilteringPolicy)
	return m
}

type WSGURLFilteringPolicyList struct {
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
	List []*WSGURLFilteringPolicy `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGURLFilteringPolicyList() *WSGURLFilteringPolicyList {
	m := new(WSGURLFilteringPolicyList)
	return m
}
