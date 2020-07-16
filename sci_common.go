package bigdog

// API Version: 1.0.0

// SCICommonQueryBody
//
// Definition: common_queryBody
type SCICommonQueryBody struct {
	// End
	// Format: 2016-04-06T16:04:46+00:00
	// Constraints:
	//    - required
	End *string `json:"end"`

	Filter *SCICommonQueryFilter `json:"filter,omitempty"`

	// Granularity
	// Constraints:
	//    - oneof:[all,fifteen_minute,thirty_minute,hour,day]
	Granularity *string `json:"granularity,omitempty"`

	Limit *float64 `json:"limit,omitempty"`

	Metric *string `json:"metric,omitempty"`

	// PagingIdentifiers
	// Query results will return a pagingIdentifiers JSON object that can be reused in the next query for pagination.
	PagingIdentifiers interface{} `json:"pagingIdentifiers,omitempty"`

	// Start
	// Format: 2016-04-06T16:04:46+00:00
	// Constraints:
	//    - required
	Start *string `json:"start"`
}

func NewSCICommonQueryBody() *SCICommonQueryBody {
	m := new(SCICommonQueryBody)
	return m
}

// SCICommonQueryFilter
//
// Definition: common_queryFilter
//
// Data query filter body
type SCICommonQueryFilter struct {
	AlphaNumeric *bool `json:"alphaNumeric,omitempty"`

	Dimension *string `json:"dimension,omitempty"`

	Fields []*SCICommonQueryFilter `json:"fields,omitempty"`

	Lower *string `json:"lower,omitempty"`

	LowerStrict *bool `json:"lowerStrict,omitempty"`

	Pattern *string `json:"pattern,omitempty"`

	// Type
	// Type of filter
	// Constraints:
	//    - required
	Type *string `json:"type"`

	Upper *string `json:"upper,omitempty"`

	UpperStrict *bool `json:"upperStrict,omitempty"`

	Value *string `json:"value,omitempty"`

	Values []string `json:"values,omitempty"`
}

func NewSCICommonQueryFilter() *SCICommonQueryFilter {
	m := new(SCICommonQueryFilter)
	return m
}

// SCICommonReportQueryBody
//
// Definition: common_reportQueryBody
type SCICommonReportQueryBody struct {
	// End
	// Format: 2016-04-06T16:04:46+00:00
	// Constraints:
	//    - required
	End *string `json:"end"`

	Filter *SCICommonQueryFilter `json:"filter,omitempty"`

	// Granularity
	// Constraints:
	//    - oneof:[all,fifteen_minute,thirty_minute,hour,day]
	Granularity *string `json:"granularity,omitempty"`

	Limit *float64 `json:"limit,omitempty"`

	Metric *string `json:"metric,omitempty"`

	// PagingIdentifiers
	// Query results will return a pagingIdentifiers JSON object that can be reused in the next query for pagination.
	PagingIdentifiers interface{} `json:"pagingIdentifiers,omitempty"`

	// Start
	// Format: 2016-04-06T16:04:46+00:00
	// Constraints:
	//    - required
	Start *string `json:"start"`

	SwitchFilter *SCICommonQueryFilter `json:"switchFilter,omitempty"`
}

func NewSCICommonReportQueryBody() *SCICommonReportQueryBody {
	m := new(SCICommonReportQueryBody)
	return m
}
