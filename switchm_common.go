package ruckus

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMCommonBulkDeleteRequest struct {
	IdList *SwitchMCommonBulkDeleteRequest `json:"idList,omitempty"`
}

func NewSwitchMCommonBulkDeleteRequest() *SwitchMCommonBulkDeleteRequest {
	m := new(SwitchMCommonBulkDeleteRequest)
	return m
}

type SwitchMCommonCreateResult struct {
	Id *string `json:"id,omitempty"`
}

func NewSwitchMCommonCreateResult() *SwitchMCommonCreateResult {
	m := new(SwitchMCommonCreateResult)
	return m
}

type SwitchMCommonFilterOperator string

func NewSwitchMCommonFilterOperator() *SwitchMCommonFilterOperator {
	m := new(SwitchMCommonFilterOperator)
	return m
}

type SwitchMCommonFullTextSearch struct {
	// Fields
	// Specific fields to search
	Fields []string `json:"fields,omitempty"`

	// Type
	// Search logic operator
	// Constraints:
	//    - oneof:[AND,OR]
	Type *string `json:"type,omitempty"`

	// Value
	// Text or number to search
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonFullTextSearch() *SwitchMCommonFullTextSearch {
	m := new(SwitchMCommonFullTextSearch)
	return m
}

type SwitchMCommonIdList []string

func MakeSwitchMCommonIdList() SwitchMCommonIdList {
	m := make(SwitchMCommonIdList, 0)
	return m
}

type SwitchMCommonQueryCriteria struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []**SwitchMCommonQueryCriteria `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []**SwitchMCommonQueryCriteria `json:"extraNotFilters,omitempty"`

	ExtraTimeRange **SwitchMCommonQueryCriteria `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []**SwitchMCommonQueryCriteria `json:"filters,omitempty"`

	FullTextSearch **SwitchMCommonQueryCriteria `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options **SwitchMCommonQueryCriteria `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo **SwitchMCommonQueryCriteria `json:"sortInfo,omitempty"`
}

func NewSwitchMCommonQueryCriteria() *SwitchMCommonQueryCriteria {
	m := new(SwitchMCommonQueryCriteria)
	return m
}

type SwitchMCommonQueryCriteriaExtraFiltersType struct {
	Operator **SwitchMCommonQueryCriteriaExtraFiltersType `json:"operator,omitempty"`

	// Type
	// Filters for specific attributes
	Type *string `json:"type,omitempty"`

	// Value
	// value to search
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaExtraFiltersType() *SwitchMCommonQueryCriteriaExtraFiltersType {
	m := new(SwitchMCommonQueryCriteriaExtraFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaExtraNotFiltersType() *SwitchMCommonQueryCriteriaExtraNotFiltersType {
	m := new(SwitchMCommonQueryCriteriaExtraNotFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaFiltersType struct {
	Operator **SwitchMCommonQueryCriteriaFiltersType `json:"operator,omitempty"`

	// Type
	// Group type
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaFiltersType() *SwitchMCommonQueryCriteriaFiltersType {
	m := new(SwitchMCommonQueryCriteriaFiltersType)
	return m
}

// SwitchMCommonQueryCriteriaOptionsType
//
// Specified feature required information
type SwitchMCommonQueryCriteriaOptionsType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMCommonQueryCriteriaOptionsType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMCommonQueryCriteriaOptionsType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMCommonQueryCriteriaOptionsType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMCommonQueryCriteriaOptionsType() *SwitchMCommonQueryCriteriaOptionsType {
	m := new(SwitchMCommonQueryCriteriaOptionsType)
	return m
}

type SwitchMCommonQueryCriteriaSortInfoType struct {
	Dir *string `json:"dir,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSortInfoType() *SwitchMCommonQueryCriteriaSortInfoType {
	m := new(SwitchMCommonQueryCriteriaSortInfoType)
	return m
}

type SwitchMCommonQueryCriteriaSuperSet struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []**SwitchMCommonQueryCriteriaSuperSet `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []**SwitchMCommonQueryCriteriaSuperSet `json:"extraNotFilters,omitempty"`

	ExtraTimeRange **SwitchMCommonQueryCriteriaSuperSet `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []**SwitchMCommonQueryCriteriaSuperSet `json:"filters,omitempty"`

	FullTextSearch **SwitchMCommonQueryCriteriaSuperSet `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required informaion
	Options **SwitchMCommonQueryCriteriaSuperSet `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo **SwitchMCommonQueryCriteriaSuperSet `json:"sortInfo,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSet() *SwitchMCommonQueryCriteriaSuperSet {
	m := new(SwitchMCommonQueryCriteriaSuperSet)
	return m
}

type SwitchMCommonQueryCriteriaSuperSetExtraFiltersType struct {
	Type *string `json:"type,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetExtraFiltersType() *SwitchMCommonQueryCriteriaSuperSetExtraFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetExtraFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType struct {
	Type *string `json:"type,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType() *SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaSuperSetFiltersType struct {
	Type *string `json:"type,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetFiltersType() *SwitchMCommonQueryCriteriaSuperSetFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaSuperSetOptionsType struct {
	AuthrealmType *string `json:"auth_realmType,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetOptionsType() *SwitchMCommonQueryCriteriaSuperSetOptionsType {
	m := new(SwitchMCommonQueryCriteriaSuperSetOptionsType)
	return m
}

type SwitchMCommonRbacMetadata struct {
	RbacMetadata interface{} `json:"rbacMetadata,omitempty"`
}

func NewSwitchMCommonRbacMetadata() *SwitchMCommonRbacMetadata {
	m := new(SwitchMCommonRbacMetadata)
	return m
}

type SwitchMCommonTimeRange struct {
	// End
	// end time for collecting data
	End *float64 `json:"end,omitempty"`

	// Field
	// time field for collecting data
	// Constraints:
	//    - oneof:[insertionTime]
	Field *string `json:"field,omitempty"`

	// Interval
	// time interval in second
	Interval *float64 `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	Start *float64 `json:"start,omitempty"`
}

func NewSwitchMCommonTimeRange() *SwitchMCommonTimeRange {
	m := new(SwitchMCommonTimeRange)
	return m
}
