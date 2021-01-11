package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// SwitchMCommonBulkDeleteRequest
//
// Definition: common_bulkDeleteRequest
type SwitchMCommonBulkDeleteRequest struct {
	IdList SwitchMCommonIdList `json:"idList,omitempty"`
}

func NewSwitchMCommonBulkDeleteRequest() *SwitchMCommonBulkDeleteRequest {
	m := new(SwitchMCommonBulkDeleteRequest)
	return m
}

// SwitchMCommonCreateResult
//
// Definition: common_createResult
type SwitchMCommonCreateResult struct {
	Id *string `json:"id,omitempty"`
}

type SwitchMCommonCreateResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMCommonCreateResult
}

func newSwitchMCommonCreateResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMCommonCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMCommonCreateResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMCommonCreateResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMCommonCreateResult() *SwitchMCommonCreateResult {
	m := new(SwitchMCommonCreateResult)
	return m
}

// SwitchMCommonFilterOperator
//
// Definition: common_filterOperator
//
// Constraints:
//    - oneof:[eq,gt,lt,gte,lte]
type SwitchMCommonFilterOperator string

func NewSwitchMCommonFilterOperator() *SwitchMCommonFilterOperator {
	m := new(SwitchMCommonFilterOperator)
	return m
}

// SwitchMCommonFullTextSearch
//
// Definition: common_fullTextSearch
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

// SwitchMCommonIdList
//
// Definition: common_idList
type SwitchMCommonIdList []string

func MakeSwitchMCommonIdList() SwitchMCommonIdList {
	m := make(SwitchMCommonIdList, 0)
	return m
}

// SwitchMCommonQueryCriteria
//
// Definition: common_queryCriteria
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
	ExtraFilters []*SwitchMCommonQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*SwitchMCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *SwitchMCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*SwitchMCommonQueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *SwitchMCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options interface{} `json:"options,omitempty"`

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
	SortInfo *SwitchMCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewSwitchMCommonQueryCriteria() *SwitchMCommonQueryCriteria {
	m := new(SwitchMCommonQueryCriteria)
	return m
}

// SwitchMCommonQueryCriteriaExtraFiltersType
//
// Definition: common_queryCriteriaExtraFiltersType
type SwitchMCommonQueryCriteriaExtraFiltersType struct {
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

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

// SwitchMCommonQueryCriteriaExtraNotFiltersType
//
// Definition: common_queryCriteriaExtraNotFiltersType
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

// SwitchMCommonQueryCriteriaFiltersType
//
// Definition: common_queryCriteriaFiltersType
type SwitchMCommonQueryCriteriaFiltersType struct {
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

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

// SwitchMCommonQueryCriteriaSortInfoType
//
// Definition: common_queryCriteriaSortInfoType
//
// About sorting
type SwitchMCommonQueryCriteriaSortInfoType struct {
	// Dir
	// Constraints:
	//    - oneof:[ASC,DESC]
	Dir *string `json:"dir,omitempty"`

	SortColumn *string `json:"sortColumn,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSortInfoType() *SwitchMCommonQueryCriteriaSortInfoType {
	m := new(SwitchMCommonQueryCriteriaSortInfoType)
	return m
}

// SwitchMCommonQueryCriteriaSuperSet
//
// Definition: common_queryCriteriaSuperSet
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
	ExtraFilters []*SwitchMCommonQueryCriteriaSuperSetExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *SwitchMCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*SwitchMCommonQueryCriteriaSuperSetFiltersType `json:"filters,omitempty"`

	FullTextSearch *SwitchMCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required informaion
	Options *SwitchMCommonQueryCriteriaSuperSetOptionsType `json:"options,omitempty"`

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
	SortInfo *SwitchMCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSet() *SwitchMCommonQueryCriteriaSuperSet {
	m := new(SwitchMCommonQueryCriteriaSuperSet)
	return m
}

// SwitchMCommonQueryCriteriaSuperSetExtraFiltersType
//
// Definition: common_queryCriteriaSuperSetExtraFiltersType
type SwitchMCommonQueryCriteriaSuperSetExtraFiltersType struct {
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attribute
	// Constraints:
	//    - oneof:[CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,ProtocolType,TIMERANGE,RADIOID,WLANID,CATEGORY,CLIENT,CP,DP,CLUSTER,NODE,BLADE,SYNCEDSTATUS,OSTYPE,APP,PORT,STATUS,REGISTRATIONSTATE,GATEWAY,APIPADDRESS,CLIENTIPADDRESS,SEVERITY,ACKNOWLEDGED,MVNOID,USER,USERID,WLANNAME,AUDITIPADDRESS,AUDITUSERUUID,AUDITOBJECT,AUDITACTION,AUDITTENANTUUID,AUDITOBJECTUUID,AUTHTYPE,AUDITTYPE,H20SuppportEnabled,AaaSuppportEnabled,GppSuppportEnabled,Type,RogueMac,SSID,ALARMSTATE,DEVICENAME,SWITCH,SWITCH_GROUP,ZoneAffinityProfileId,FIRMWARE_TYPE,SCHEDULED_TIME,VLAN,FAMILY_ID,SWITCH_ID,transactionId,hasLayerThreeConfig,clientAuthType,clientIpv4Addr,clientIpv6Addr,clientMac,clientUserName,switchName,ADMIN_STATUS,PORT_STATUS,POE_ENABLED]
	Type *string `json:"type,omitempty"`

	// Value
	// Value to search
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetExtraFiltersType() *SwitchMCommonQueryCriteriaSuperSetExtraFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetExtraFiltersType)
	return m
}

// SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType
//
// Definition: common_queryCriteriaSuperSetExtraNotFiltersType
type SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	// Constraints:
	//    - oneof:[CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,ProtocolType,TIMERANGE,RADIOID,WLANID,CATEGORY,CLIENT,CP,DP,CLUSTER,NODE,BLADE,SYNCEDSTATUS,OSTYPE,APP,PORT,STATUS,REGISTRATIONSTATE,GATEWAY,APIPADDRESS,CLIENTIPADDRESS,SEVERITY,ACKNOWLEDGED,MVNOID,USER,USERID,WLANNAME,AUDITIPADDRESS,AUDITUSERUUID,AUDITOBJECT,AUDITACTION,AUDITTENANTUUID,AUDITOBJECTUUID,AUTHTYPE,AUDITTYPE,H20SuppportEnabled,AaaSuppportEnabled,GppSuppportEnabled,Type,RogueMac,SSID,ALARMSTATE,DEVICENAME,SWITCH,ZoneAffinityProfileId,FIRMWARE_TYPE,SCHEDULED_TIME,VLAN,FAMILY_ID,SWITCH_ID,switchStatus.alerts,switchStatus.status,transactionId,hasLayerThreeConfig]
	Type *string `json:"type,omitempty"`

	// Value
	// Value not to search
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType() *SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType)
	return m
}

// SwitchMCommonQueryCriteriaSuperSetFiltersType
//
// Definition: common_queryCriteriaSuperSetFiltersType
type SwitchMCommonQueryCriteriaSuperSetFiltersType struct {
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	// Constraints:
	//    - oneof:[SYSTEM,CATEGORY,CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,BLADE,SYNCEDSTATUS,REGISTRATIONSTATE,STATUS,SWITCH_GROUP,FAMILY_ID,PORT,SWITCH_ID_PORT_IDENTIFIER]
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetFiltersType() *SwitchMCommonQueryCriteriaSuperSetFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetFiltersType)
	return m
}

// SwitchMCommonQueryCriteriaSuperSetOptionsType
//
// Definition: common_queryCriteriaSuperSetOptionsType
//
// Specified feature required informaion
type SwitchMCommonQueryCriteriaSuperSetOptionsType struct {
	// AcctincludeNa
	// Include Not Available acct service option while returning result
	AcctincludeNa *bool `json:"acct_includeNa,omitempty"`

	// AccttestableOnly
	// Only get testable service type
	AccttestableOnly *bool `json:"acct_testableOnly,omitempty"`

	// Accttype
	// Accounting service types to get, use comma to separate, Ex: RADIUS,CGF
	Accttype *string `json:"acct_type,omitempty"`

	// AuthhostedAaaSupportedEnabled
	// Indicate if Hosted AAA Support is enabled
	AuthhostedAaaSupportedEnabled *bool `json:"auth_hostedAaaSupportedEnabled,omitempty"`

	// AuthincludeAdGlobal
	// If AD is in list, include only AD with Global Catalog configured
	AuthincludeAdGlobal *bool `json:"auth_includeAdGlobal,omitempty"`

	// AuthincludeGuest
	// Include Guest auth service while returning result
	AuthincludeGuest *bool `json:"auth_includeGuest,omitempty"`

	// AuthincludeLocalDb
	// Include LocalDB auth service while returning result
	AuthincludeLocalDb *bool `json:"auth_includeLocalDb,omitempty"`

	// AuthincludeNa
	// Include Not Available auth service option while returning result
	AuthincludeNa *bool `json:"auth_includeNa,omitempty"`

	// AuthplmnIdentifierEnabled
	// Indicate if Configure PLMN identifier is enabled
	AuthplmnIdentifierEnabled *bool `json:"auth_plmnIdentifierEnabled,omitempty"`

	// AuthrealmType
	// To get specific authentication service information for configuring realm based authentication profile
	// Constraints:
	//    - oneof:[ALL,RADIUS]
	AuthrealmType *string `json:"auth_realmType,omitempty"`

	// AuthtestableOnly
	// Only get testable service type
	AuthtestableOnly *bool `json:"auth_testableOnly,omitempty"`

	// Authtype
	// Authentication service types to get, use comma to separate, Ex: RADIUS,AD
	Authtype *string `json:"auth_type,omitempty"`

	// Forwardingtype
	// Forwarding service types to get, use comma to separate, Ex: L2oGRE,TTGPDG,Bridge,Advanced
	Forwardingtype *string `json:"forwarding_type,omitempty"`

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

func NewSwitchMCommonQueryCriteriaSuperSetOptionsType() *SwitchMCommonQueryCriteriaSuperSetOptionsType {
	m := new(SwitchMCommonQueryCriteriaSuperSetOptionsType)
	return m
}

// SwitchMCommonTimeRange
//
// Definition: common_timeRange
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
