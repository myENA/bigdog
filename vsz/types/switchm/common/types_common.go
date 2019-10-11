package common

// API Version: v8_1

import (
	"encoding/json"
)

type BulkDeleteRequest struct {
	IdList IdList `json:"idList,omitempty"`
}

func NewBulkDeleteRequest() *BulkDeleteRequest {
	bulkDeleteRequestType := new(BulkDeleteRequest)
	return bulkDeleteRequestType
}

func NewDefaultBulkDeleteRequest() *BulkDeleteRequest {
	bulkDeleteRequestType := new(BulkDeleteRequest)
	return bulkDeleteRequestType
}

type CreateResult struct {
	Id *string `json:"id,omitempty"`
}

func NewCreateResult() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

func NewDefaultCreateResult() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

type FilterOperator string

func NewFilterOperator() *FilterOperator {
	filterOperatorType := new(FilterOperator)
	return filterOperatorType
}

func NewDefaultFilterOperator() *FilterOperator {
	filterOperatorType := new(FilterOperator)
	return filterOperatorType
}

type FullTextSearch struct {
	// Fields
	// Specific fields to search
	Fields []string `json:"fields,omitempty"`

	// Type
	// Search logic operator
	// Constraints:
	//    - nullable
	//    - oneof:[AND,OR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AND OR"`

	// Value
	// Text or number to search
	Value *string `json:"value,omitempty"`
}

func NewFullTextSearch() *FullTextSearch {
	fullTextSearchType := new(FullTextSearch)
	return fullTextSearchType
}

func NewDefaultFullTextSearch() *FullTextSearch {
	fullTextSearchType := new(FullTextSearch)
	return fullTextSearchType
}

type IdList []string

func NewIdList() *IdList {
	idListType := make(IdList, 0)
	return &idListType
}

func NewDefaultIdList() *IdList {
	idListType := make(IdList, 0)
	return &idListType
}

type QueryCriteria struct {
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
	ExtraFilters []*QueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*QueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *TimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*QueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *FullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required information
	Options *QueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - nullable
	//    - min:1
	Page *int `json:"page,omitempty" validate:"omitempty,gte=1"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *QueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewQueryCriteria() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
}

func NewDefaultQueryCriteria() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
}

type QueryCriteriaExtraFiltersType struct {
	Operator *FilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attributes
	Type *string `json:"type,omitempty"`

	// Value
	// value to search
	Value *string `json:"value,omitempty"`
}

func NewQueryCriteriaExtraFiltersType() *QueryCriteriaExtraFiltersType {
	queryCriteriaExtraFiltersTypeType := new(QueryCriteriaExtraFiltersType)
	return queryCriteriaExtraFiltersTypeType
}

func NewDefaultQueryCriteriaExtraFiltersType() *QueryCriteriaExtraFiltersType {
	queryCriteriaExtraFiltersTypeType := new(QueryCriteriaExtraFiltersType)
	return queryCriteriaExtraFiltersTypeType
}

type QueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
}

func NewQueryCriteriaExtraNotFiltersType() *QueryCriteriaExtraNotFiltersType {
	queryCriteriaExtraNotFiltersTypeType := new(QueryCriteriaExtraNotFiltersType)
	return queryCriteriaExtraNotFiltersTypeType
}

func NewDefaultQueryCriteriaExtraNotFiltersType() *QueryCriteriaExtraNotFiltersType {
	queryCriteriaExtraNotFiltersTypeType := new(QueryCriteriaExtraNotFiltersType)
	return queryCriteriaExtraNotFiltersTypeType
}

type QueryCriteriaFiltersType struct {
	Operator *FilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

func NewQueryCriteriaFiltersType() *QueryCriteriaFiltersType {
	queryCriteriaFiltersTypeType := new(QueryCriteriaFiltersType)
	return queryCriteriaFiltersTypeType
}

func NewDefaultQueryCriteriaFiltersType() *QueryCriteriaFiltersType {
	queryCriteriaFiltersTypeType := new(QueryCriteriaFiltersType)
	return queryCriteriaFiltersTypeType
}

// QueryCriteriaOptionsType
//
// Specified feature required information
type QueryCriteriaOptionsType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *QueryCriteriaOptionsType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = QueryCriteriaOptionsType{XAdditionalProperties: tmp}
	return nil
}

func (t *QueryCriteriaOptionsType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewQueryCriteriaOptionsType() *QueryCriteriaOptionsType {
	queryCriteriaOptionsTypeType := new(QueryCriteriaOptionsType)
	return queryCriteriaOptionsTypeType
}

func NewDefaultQueryCriteriaOptionsType() *QueryCriteriaOptionsType {
	queryCriteriaOptionsTypeType := new(QueryCriteriaOptionsType)
	return queryCriteriaOptionsTypeType
}

// QueryCriteriaSortInfoType
//
// About sorting
type QueryCriteriaSortInfoType struct {
	// Dir
	// Constraints:
	//    - nullable
	//    - oneof:[ASC,DESC]
	Dir *string `json:"dir,omitempty" validate:"omitempty,oneof=ASC DESC"`

	SortColumn *string `json:"sortColumn,omitempty"`
}

func NewQueryCriteriaSortInfoType() *QueryCriteriaSortInfoType {
	queryCriteriaSortInfoTypeType := new(QueryCriteriaSortInfoType)
	return queryCriteriaSortInfoTypeType
}

func NewDefaultQueryCriteriaSortInfoType() *QueryCriteriaSortInfoType {
	queryCriteriaSortInfoTypeType := new(QueryCriteriaSortInfoType)
	return queryCriteriaSortInfoTypeType
}

type QueryCriteriaSuperSet struct{}

func NewQueryCriteriaSuperSet() *QueryCriteriaSuperSet {
	queryCriteriaSuperSetType := new(QueryCriteriaSuperSet)
	return queryCriteriaSuperSetType
}

func NewDefaultQueryCriteriaSuperSet() *QueryCriteriaSuperSet {
	queryCriteriaSuperSetType := new(QueryCriteriaSuperSet)
	return queryCriteriaSuperSetType
}

type RbacMetadata struct {
	RbacMetadata []string `json:"rbacMetadata,omitempty"`
}

func NewRbacMetadata() *RbacMetadata {
	rbacMetadataType := new(RbacMetadata)
	return rbacMetadataType
}

func NewDefaultRbacMetadata() *RbacMetadata {
	rbacMetadataType := new(RbacMetadata)
	return rbacMetadataType
}

type TimeRange struct {
	// End
	// end time for collecting data
	End *float64 `json:"end,omitempty"`

	// Field
	// time field for collecting data
	// Constraints:
	//    - nullable
	//    - oneof:[insertionTime]
	Field *string `json:"field,omitempty" validate:"omitempty,oneof=insertionTime"`

	// Interval
	// time interval in second
	Interval *float64 `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	Start *float64 `json:"start,omitempty"`
}

func NewTimeRange() *TimeRange {
	timeRangeType := new(TimeRange)
	return timeRangeType
}

func NewDefaultTimeRange() *TimeRange {
	timeRangeType := new(TimeRange)
	return timeRangeType
}
