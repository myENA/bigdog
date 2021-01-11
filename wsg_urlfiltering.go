package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGURLFilteringCreateUrlFilteringPolicy
//
// Definition: urlFiltering_createUrlFilteringPolicy
type WSGURLFilteringCreateUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

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
	Name *WSGCommonNormalName `json:"name"`

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

// WSGURLFilteringDeleteBulk
//
// Definition: urlFiltering_deleteBulk
type WSGURLFilteringDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGURLFilteringDeleteBulk() *WSGURLFilteringDeleteBulk {
	m := new(WSGURLFilteringDeleteBulk)
	return m
}

// WSGURLFilteringModifyUrlFilteringPolicy
//
// Definition: urlFiltering_modifyUrlFilteringPolicy
type WSGURLFilteringModifyUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

// WSGURLFilteringBlockCategoriesList
//
// Definition: urlFiltering_urlFilteringBlockCategoriesList
type WSGURLFilteringBlockCategoriesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGURLFilteringBlockCategory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGURLFilteringBlockCategoriesListAPIResponse struct {
	*RawAPIResponse
	Data *WSGURLFilteringBlockCategoriesList
}

func newWSGURLFilteringBlockCategoriesListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGURLFilteringBlockCategoriesListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGURLFilteringBlockCategoriesListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGURLFilteringBlockCategoriesList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGURLFilteringBlockCategoriesList() *WSGURLFilteringBlockCategoriesList {
	m := new(WSGURLFilteringBlockCategoriesList)
	return m
}

// WSGURLFilteringBlockCategory
//
// Definition: urlFiltering_urlFilteringBlockCategory
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

// WSGURLFilteringPolicy
//
// Definition: urlFiltering_urlFilteringPolicy
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

	Description *WSGCommonDescription `json:"description,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

type WSGURLFilteringPolicyAPIResponse struct {
	*RawAPIResponse
	Data *WSGURLFilteringPolicy
}

func newWSGURLFilteringPolicyAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGURLFilteringPolicyAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGURLFilteringPolicyAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGURLFilteringPolicy)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGURLFilteringPolicy() *WSGURLFilteringPolicy {
	m := new(WSGURLFilteringPolicy)
	return m
}

// WSGURLFilteringPolicyList
//
// Definition: urlFiltering_urlFilteringPolicyList
type WSGURLFilteringPolicyList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGURLFilteringPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGURLFilteringPolicyListAPIResponse struct {
	*RawAPIResponse
	Data *WSGURLFilteringPolicyList
}

func newWSGURLFilteringPolicyListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGURLFilteringPolicyListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGURLFilteringPolicyListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGURLFilteringPolicyList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGURLFilteringPolicyList() *WSGURLFilteringPolicyList {
	m := new(WSGURLFilteringPolicyList)
	return m
}
