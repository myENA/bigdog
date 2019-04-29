package urlfiltering

// API Version: v8_0

type CreateURLFilteringPolicy struct {
	Blacklist                []string `json:"blacklist,omitempty"`
	BlockCategories          []int    `json:"blockCategories,omitempty"`
	Description              *string  `json:"description,omitempty"`
	DomainID                 *string  `json:"domainId,omitempty"`
	FilteringLevel           *string  `json:"filteringLevel,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	SafeSearchBingDNS        *string  `json:"safeSearchBingDns,omitempty"`
	SafeSearchBingEnabled    *bool    `json:"safeSearchBingEnabled,omitempty"`
	SafeSearchGoogleDNS      *string  `json:"safeSearchGoogleDns,omitempty"`
	SafeSearchGoogleEnabled  *bool    `json:"safeSearchGoogleEnabled,omitempty"`
	SafeSearchYouTubeDNS     *string  `json:"safeSearchYouTubeDns,omitempty"`
	SafeSearchYouTubeEnabled *bool    `json:"safeSearchYouTubeEnabled,omitempty"`
	Whitelist                []string `json:"whitelist,omitempty"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyURLFilteringPolicy struct {
	Blacklist                []string `json:"blacklist,omitempty"`
	BlockCategories          []int    `json:"blockCategories,omitempty"`
	Description              *string  `json:"description,omitempty"`
	FilteringLevel           *string  `json:"filteringLevel,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	SafeSearchBingDNS        *string  `json:"safeSearchBingDns,omitempty"`
	SafeSearchBingEnabled    *bool    `json:"safeSearchBingEnabled,omitempty"`
	SafeSearchGoogleDNS      *string  `json:"safeSearchGoogleDns,omitempty"`
	SafeSearchGoogleEnabled  *bool    `json:"safeSearchGoogleEnabled,omitempty"`
	SafeSearchYouTubeDNS     *string  `json:"safeSearchYouTubeDns,omitempty"`
	SafeSearchYouTubeEnabled *bool    `json:"safeSearchYouTubeEnabled,omitempty"`
	Whitelist                []string `json:"whitelist,omitempty"`
}

type URLFilteringBlockCategoriesList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type URLFilteringBlockCategory struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type URLFilteringPolicy struct {
	Blacklist                []string `json:"blacklist,omitempty"`
	BlockCategories          []int    `json:"blockCategories,omitempty"`
	CreateDateTime           *int     `json:"createDateTime,omitempty"`
	CreatorID                *string  `json:"creatorId,omitempty"`
	CreatorUsername          *string  `json:"creatorUsername,omitempty"`
	Description              *string  `json:"description,omitempty"`
	DomainID                 *string  `json:"domainId,omitempty"`
	FilteringLevel           *string  `json:"filteringLevel,omitempty"`
	ID                       *string  `json:"id,omitempty"`
	ModifiedDateTime         *int     `json:"modifiedDateTime,omitempty"`
	ModifierID               *string  `json:"modifierId,omitempty"`
	ModifierUsername         *string  `json:"modifierUsername,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	SafeSearchBingDNS        *string  `json:"safeSearchBingDns,omitempty"`
	SafeSearchBingEnabled    *bool    `json:"safeSearchBingEnabled,omitempty"`
	SafeSearchGoogleDNS      *string  `json:"safeSearchGoogleDns,omitempty"`
	SafeSearchGoogleEnabled  *bool    `json:"safeSearchGoogleEnabled,omitempty"`
	SafeSearchYouTubeDNS     *string  `json:"safeSearchYouTubeDns,omitempty"`
	SafeSearchYouTubeEnabled *bool    `json:"safeSearchYouTubeEnabled,omitempty"`
	TenantID                 *string  `json:"tenantId,omitempty"`
	Whitelist                []string `json:"whitelist,omitempty"`
}

type URLFilteringPolicyList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}
