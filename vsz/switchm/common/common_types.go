package common

// API Version: v8_0

type FullTextSearch struct {
	Fields []string `json:"fields,omitempty"`
	Type   *string  `json:"type,omitempty"`
	Value  *string  `json:"value,omitempty"`
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

type QueryCriteriaExtraFiltersType struct {
	Operator interface{} `json:"operator,omitempty"`
	Type     *string     `json:"type,omitempty"`
	Value    *string     `json:"value,omitempty"`
}

type QueryCriteriaExtraNotFiltersType struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaFiltersType struct {
	Operator interface{} `json:"operator,omitempty"`
	Type     *string     `json:"type,omitempty"`
	Value    *string     `json:"value,omitempty"`
}

type QueryCriteriaSortInfoType struct {
	Dir        *string `json:"dir,omitempty"`
	SortColumn *string `json:"sortColumn,omitempty"`
}

type QueryCriteriaSuperSet struct {
	Attributes      []string           `json:"attributes,omitempty"`
	Criteria        *string            `json:"criteria,omitempty"`
	ExpandDomains   *bool              `json:"expandDomains,omitempty"`
	ExtraFilters    []*ExtraFilters    `json:"extraFilters,omitempty"`
	ExtraNotFilters []*ExtraNotFilters `json:"extraNotFilters,omitempty"`
	ExtraTimeRange  *ExtraTimeRange    `json:"extraTimeRange,omitempty"`
	Filters         []*Filters         `json:"filters,omitempty"`
	FullTextSearch  *FullTextSearch    `json:"fullTextSearch,omitempty"`
	Limit           *int               `json:"limit,omitempty"`
	Options         *Options           `json:"options,omitempty"`
	Page            *int               `json:"page,omitempty"`
	Query           *string            `json:"query,omitempty"`
	SortInfo        *SortInfo          `json:"sortInfo,omitempty"`
}

type QueryCriteriaSuperSetExtraFiltersType struct {
	Operator interface{} `json:"operator,omitempty"`
	Type     *string     `json:"type,omitempty"`
	Value    *string     `json:"value,omitempty"`
}

type QueryCriteriaSuperSetExtraNotFiltersType struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaSuperSetFiltersType struct {
	Operator interface{} `json:"operator,omitempty"`
	Type     *string     `json:"type,omitempty"`
	Value    *string     `json:"value,omitempty"`
}

type QueryCriteriaSuperSetOptionsType struct {
	AcctIncludeNa                 *bool   `json:"acct_includeNa,omitempty"`
	AcctTestableOnly              *bool   `json:"acct_testableOnly,omitempty"`
	AcctType                      *string `json:"acct_type,omitempty"`
	AuthHostedAaaSupportedEnabled *bool   `json:"auth_hostedAaaSupportedEnabled,omitempty"`
	AuthIncludeAdGlobal           *bool   `json:"auth_includeAdGlobal,omitempty"`
	AuthIncludeGuest              *bool   `json:"auth_includeGuest,omitempty"`
	AuthIncludeLocalDb            *bool   `json:"auth_includeLocalDb,omitempty"`
	AuthIncludeNa                 *bool   `json:"auth_includeNa,omitempty"`
	AuthPlmnIdentifierEnabled     *bool   `json:"auth_plmnIdentifierEnabled,omitempty"`
	AuthRealmType                 *string `json:"auth_realmType,omitempty"`
	AuthTestableOnly              *bool   `json:"auth_testableOnly,omitempty"`
	AuthType                      *string `json:"auth_type,omitempty"`
	ForwardingType                *string `json:"forwarding_type,omitempty"`
	GlobalFilterID                *string `json:"globalFilterId,omitempty"`
	IncludeSharedResources        *bool   `json:"includeSharedResources,omitempty"`
	IncludeUsers                  *bool   `json:"includeUsers,omitempty"`
	INCLUDERBACMETADATA           *bool   `json:"INCLUDE_RBAC_METADATA,omitempty"`
	InMap                         *bool   `json:"inMap,omitempty"`
	TENANTID                      *string `json:"TENANT_ID,omitempty"`
}

type RBACMetadata struct {
	RBACMetadata []string `json:"rbacMetadata,omitempty"`
}

type TimeRange struct {
	End      *float64 `json:"end,omitempty"`
	Field    *string  `json:"field,omitempty"`
	Interval *float64 `json:"interval,omitempty"`
	Start    *float64 `json:"start,omitempty"`
}
