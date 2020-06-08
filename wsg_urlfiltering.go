package ruckus

// API Version: v9_0

type WSGURLFilteringCreateUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *WSGURLFilteringCreateUrlFilteringPolicy `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - required
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel"`

	// Name
	// Constraints:
	//    - required
	Name *WSGURLFilteringCreateUrlFilteringPolicy `json:"name"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

func NewWSGURLFilteringCreateUrlFilteringPolicy() *WSGURLFilteringCreateUrlFilteringPolicy {
	m := new(WSGURLFilteringCreateUrlFilteringPolicy)
	return m
}

type WSGURLFilteringDeleteBulk struct {
	IdList *WSGURLFilteringDeleteBulk `json:"idList,omitempty"`
}

func NewWSGURLFilteringDeleteBulk() *WSGURLFilteringDeleteBulk {
	m := new(WSGURLFilteringDeleteBulk)
	return m
}

type WSGURLFilteringModifyUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *WSGURLFilteringModifyUrlFilteringPolicy `json:"description,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel,omitempty"`

	Name *WSGURLFilteringModifyUrlFilteringPolicy `json:"name,omitempty"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

func NewWSGURLFilteringModifyUrlFilteringPolicy() *WSGURLFilteringModifyUrlFilteringPolicy {
	m := new(WSGURLFilteringModifyUrlFilteringPolicy)
	return m
}

type WSGURLFilteringBlockCategoriesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGURLFilteringBlockCategoriesList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGURLFilteringBlockCategoriesList() *WSGURLFilteringBlockCategoriesList {
	m := new(WSGURLFilteringBlockCategoriesList)
	return m
}

type WSGURLFilteringBlockCategory struct {
	// Id
	// Identifier of the URL Filtering Category
	Id *int `json:"id,omitempty"`

	// Name
	// name of the URL Filtering Category
	Name *string `json:"name,omitempty"`
}

func NewWSGURLFilteringBlockCategory() *WSGURLFilteringBlockCategory {
	m := new(WSGURLFilteringBlockCategory)
	return m
}

type WSGURLFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGURLFilteringPolicy `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel,omitempty"`

	// Id
	// Identifier of the URL Filtering policy
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

	Name *WSGURLFilteringPolicy `json:"name,omitempty"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

func NewWSGURLFilteringPolicy() *WSGURLFilteringPolicy {
	m := new(WSGURLFilteringPolicy)
	return m
}

type WSGURLFilteringPolicyList struct {
	Extra *WSGURLFilteringPolicyList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGURLFilteringPolicyList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGURLFilteringPolicyList() *WSGURLFilteringPolicyList {
	m := new(WSGURLFilteringPolicyList)
	return m
}