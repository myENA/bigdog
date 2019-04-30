package common

// API Version: v8_0

type FullTextSearch struct {
	Fields []string `json:"fields,omitempty"`
	Type   *string  `json:"type,omitempty"`
	Value  *string  `json:"value,omitempty"`
}

type QueryCriteria struct {
	Attributes      []string                            `json:"attributes,omitempty"`
	Criteria        *string                             `json:"criteria,omitempty"`
	ExpandDomains   *bool                               `json:"expandDomains,omitempty"`
	ExtraFilters    []*QueryCriteriaExtraFiltersType    `json:"extraFilters,omitempty"`
	ExtraNotFilters []*QueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`
	ExtraTimeRange  *TimeRange                          `json:"extraTimeRange,omitempty"`
	Filters         []*QueryCriteriaFiltersType         `json:"filters,omitempty"`
	FullTextSearch  *FullTextSearch                     `json:"fullTextSearch,omitempty"`
	Limit           *int                                `json:"limit,omitempty"`
	Options         map[string]interface{}              `json:"options,omitempty"`
	Page            *int                                `json:"page,omitempty"`
	Query           *string                             `json:"query,omitempty"`
	SortInfo        *QueryCriteriaSortInfoType          `json:"sortInfo,omitempty"`
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
