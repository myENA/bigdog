package common

// API Version: v8_0

type FilterOperator interface{}

type FullTextSearch struct {
	// Fields
	// Specific fields to search
	Fields []string `json:"fields,omitempty"`

	// Type
	// Search logic operator
	Type *string `json:"type,omitempty"`

	// Value
	// Text or number to search
	Value *string `json:"value,omitempty"`
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
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options map[string]interface{} `json:"options,omitempty"`

	// Page
	// Page number to get
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *QueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

type QueryCriteriaExtraFiltersType struct {
	Operator FilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attributes
	Type *string `json:"type,omitempty"`

	// Value
	// value to search
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaFiltersType struct {
	Operator FilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

// QueryCriteriaSortInfoType
//
// About sorting
type QueryCriteriaSortInfoType struct {
	Dir *string `json:"dir,omitempty"`

	SortColumn *string `json:"sortColumn,omitempty"`
}

type QueryCriteriaSuperSet struct {
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
	ExtraFilters []*QueryCriteriaSuperSetExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*QueryCriteriaSuperSetExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *TimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*QueryCriteriaSuperSetFiltersType `json:"filters,omitempty"`

	FullTextSearch *FullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required informaion
	Options *QueryCriteriaSuperSetOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	Page *int `json:"page,omitempty"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *QueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

type QueryCriteriaSuperSetExtraFiltersType struct {
	Operator FilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// Value to search
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaSuperSetExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// Value not to search
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaSuperSetFiltersType struct {
	Operator FilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

// QueryCriteriaSuperSetOptionsType
//
// Specified feature required informaion
type QueryCriteriaSuperSetOptionsType struct {
	// AcctIncludeNa
	// Include Not Available acct service option while returning result
	AcctIncludeNa *bool `json:"acct_includeNa,omitempty"`

	// AcctTestableOnly
	// Only get testable service type
	AcctTestableOnly *bool `json:"acct_testableOnly,omitempty"`

	// AcctType
	// Accounting service types to get, use comma to separate, Ex: RADIUS,CGF
	AcctType *string `json:"acct_type,omitempty"`

	// AuthHostedAaaSupportedEnabled
	// Indicate if Hosted AAA Support is enabled
	AuthHostedAaaSupportedEnabled *bool `json:"auth_hostedAaaSupportedEnabled,omitempty"`

	// AuthIncludeAdGlobal
	// If AD is in list, include only AD with Global Catalog configured
	AuthIncludeAdGlobal *bool `json:"auth_includeAdGlobal,omitempty"`

	// AuthIncludeGuest
	// Include Guest auth service while returning result
	AuthIncludeGuest *bool `json:"auth_includeGuest,omitempty"`

	// AuthIncludeLocalDb
	// Include LocalDB auth service while returning result
	AuthIncludeLocalDb *bool `json:"auth_includeLocalDb,omitempty"`

	// AuthIncludeNa
	// Include Not Available auth service option while returning result
	AuthIncludeNa *bool `json:"auth_includeNa,omitempty"`

	// AuthPlmnIdentifierEnabled
	// Indicate if Configure PLMN identifier is enabled
	AuthPlmnIdentifierEnabled *bool `json:"auth_plmnIdentifierEnabled,omitempty"`

	// AuthRealmType
	// To get specific authentication service information for configuring realm based authentication profile
	AuthRealmType *string `json:"auth_realmType,omitempty"`

	// AuthTestableOnly
	// Only get testable service type
	AuthTestableOnly *bool `json:"auth_testableOnly,omitempty"`

	// AuthType
	// Authentication service types to get, use comma to separate, Ex: RADIUS,AD
	AuthType *string `json:"auth_type,omitempty"`

	// ForwardingType
	// Forwarding service types to get, use comma to separate, Ex: L2oGRE,TTGPDG,Bridge,Advanced
	ForwardingType *string `json:"forwarding_type,omitempty"`

	// GlobalFilterId
	// Specify GlobalFilter ID for query
	GlobalFilterId *string `json:"globalFilterId,omitempty"`

	// IncludeSharedResources
	// Whether to include the resources of parent domain or not
	IncludeSharedResources *bool `json:"includeSharedResources,omitempty"`

	// IncludeUsers
	// Should also retrieve users or not
	IncludeUsers *bool `json:"includeUsers,omitempty"`

	// INCLUDERBACMETADATA
	// Whether to include RBAC metadata or not
	INCLUDERBACMETADATA *bool `json:"INCLUDE_RBAC_METADATA,omitempty"`

	// InMap
	// Specify inMap status for query
	InMap *bool `json:"inMap,omitempty"`

	// TENANTID
	// Specify Tenant ID for query
	TENANTID *string `json:"TENANT_ID,omitempty"`
}

type RbacMetadata struct {
	RbacMetadata []string `json:"rbacMetadata,omitempty"`
}

type TimeRange struct {
	// End
	// end time for collecting data
	End *float64 `json:"end,omitempty"`

	// Field
	// time field for collecting data
	Field *string `json:"field,omitempty"`

	// Interval
	// time interval in second
	Interval *float64 `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	Start *float64 `json:"start,omitempty"`
}
