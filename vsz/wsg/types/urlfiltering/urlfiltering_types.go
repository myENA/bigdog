package urlfiltering

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/types/common"
)

type CreateURLFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	FilteringLevel *string `json:"filteringLevel,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// SafeSearchBingDNS
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDNS *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDNS
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDNS *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDNS
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDNS *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyURLFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	FilteringLevel *string `json:"filteringLevel,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// SafeSearchBingDNS
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDNS *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDNS
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDNS *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDNS
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDNS *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

type URLFilteringBlockCategoriesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*URLFilteringBlockCategory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type URLFilteringBlockCategory struct {
	// ID
	// Identifier of the URL Filtering Category
	ID *int `json:"id,omitempty"`

	// Name
	// name of the URL Filtering Category
	Name *string `json:"name,omitempty"`
}

type URLFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

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
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	FilteringLevel *string `json:"filteringLevel,omitempty"`

	// ID
	// Identifier of the URL Filtering policy
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

	// SafeSearchBingDNS
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDNS *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDNS
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDNS *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDNS
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDNS *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

type URLFilteringPolicyList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*URLFilteringPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
