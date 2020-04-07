package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMCommonBulkDeleteRequest struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList SwitchMCommonIdList `json:"idList,omitempty"`
}

func NewSwitchMCommonBulkDeleteRequest() *SwitchMCommonBulkDeleteRequest {
	m := new(SwitchMCommonBulkDeleteRequest)
	return m
}

type SwitchMCommonCreateResult struct {
	// Id
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Fields []string `json:"fields,omitempty" validate:"omitempty,dive"`

	// Type
	// Search logic operator
	// Constraints:
	//    - nullable
	//    - oneof:[AND,OR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AND OR"`

	// Value
	// Text or number to search
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	Attributes []string `json:"attributes,omitempty" validate:"omitempty,dive"`

	// Criteria
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	// Constraints:
	//    - nullable
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraFilters []*SwitchMCommonQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty" validate:"omitempty,dive"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraNotFilters []*SwitchMCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty" validate:"omitempty,dive"`

	// ExtraTimeRange
	// Constraints:
	//    - nullable
	ExtraTimeRange *SwitchMCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	// Constraints:
	//    - nullable
	Filters []*SwitchMCommonQueryCriteriaFiltersType `json:"filters,omitempty" validate:"omitempty,dive"`

	// FullTextSearch
	// Constraints:
	//    - nullable
	FullTextSearch *SwitchMCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required information
	// Constraints:
	//    - nullable
	Options *SwitchMCommonQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - nullable
	//    - min:1
	Page *int `json:"page,omitempty" validate:"omitempty,gte=1"`

	// Query
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	// Constraints:
	//    - nullable
	SortInfo *SwitchMCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewSwitchMCommonQueryCriteria() *SwitchMCommonQueryCriteria {
	m := new(SwitchMCommonQueryCriteria)
	return m
}

type SwitchMCommonQueryCriteriaExtraFiltersType struct {
	// Operator
	// Constraints:
	//    - nullable
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attributes
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`

	// Value
	// value to search
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaExtraFiltersType() *SwitchMCommonQueryCriteriaExtraFiltersType {
	m := new(SwitchMCommonQueryCriteriaExtraFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaExtraNotFiltersType() *SwitchMCommonQueryCriteriaExtraNotFiltersType {
	m := new(SwitchMCommonQueryCriteriaExtraNotFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaFiltersType struct {
	// Operator
	// Constraints:
	//    - nullable
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaFiltersType() *SwitchMCommonQueryCriteriaFiltersType {
	m := new(SwitchMCommonQueryCriteriaFiltersType)
	return m
}

// SwitchMCommonQueryCriteriaOptionsType
//
// Specified feature required information
// Constraints:
//    - nullable
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

// SwitchMCommonQueryCriteriaSortInfoType
//
// About sorting
// Constraints:
//    - nullable
type SwitchMCommonQueryCriteriaSortInfoType struct {
	// Dir
	// Constraints:
	//    - nullable
	//    - oneof:[ASC,DESC]
	Dir *string `json:"dir,omitempty" validate:"omitempty,oneof=ASC DESC"`

	// SortColumn
	// Constraints:
	//    - nullable
	SortColumn *string `json:"sortColumn,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSortInfoType() *SwitchMCommonQueryCriteriaSortInfoType {
	m := new(SwitchMCommonQueryCriteriaSortInfoType)
	return m
}

type SwitchMCommonQueryCriteriaSuperSet struct {
	// Attributes
	// Get specific columns only
	// Constraints:
	//    - nullable
	Attributes []string `json:"attributes,omitempty" validate:"omitempty,dive"`

	// Criteria
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	// Constraints:
	//    - nullable
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraFilters []*SwitchMCommonQueryCriteriaSuperSetExtraFiltersType `json:"extraFilters,omitempty" validate:"omitempty,dive"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	// Constraints:
	//    - nullable
	ExtraNotFilters []*SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType `json:"extraNotFilters,omitempty" validate:"omitempty,dive"`

	// ExtraTimeRange
	// Constraints:
	//    - nullable
	ExtraTimeRange *SwitchMCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	// Constraints:
	//    - nullable
	Filters []*SwitchMCommonQueryCriteriaSuperSetFiltersType `json:"filters,omitempty" validate:"omitempty,dive"`

	// FullTextSearch
	// Constraints:
	//    - nullable
	FullTextSearch *SwitchMCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required informaion
	// Constraints:
	//    - nullable
	Options *SwitchMCommonQueryCriteriaSuperSetOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - nullable
	//    - min:1
	Page *int `json:"page,omitempty" validate:"omitempty,gte=1"`

	// Query
	// Add backward compatibility for UI framework
	// Constraints:
	//    - nullable
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	// Constraints:
	//    - nullable
	SortInfo *SwitchMCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSet() *SwitchMCommonQueryCriteriaSuperSet {
	m := new(SwitchMCommonQueryCriteriaSuperSet)
	return m
}

type SwitchMCommonQueryCriteriaSuperSetExtraFiltersType struct {
	// Operator
	// Constraints:
	//    - nullable
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attribute
	// Constraints:
	//    - nullable
	//    - oneof:[CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,ProtocolType,TIMERANGE,RADIOID,WLANID,CATEGORY,CLIENT,CP,DP,CLUSTER,NODE,BLADE,SYNCEDSTATUS,OSTYPE,APP,PORT,STATUS,REGISTRATIONSTATE,GATEWAY,APIPADDRESS,CLIENTIPADDRESS,SEVERITY,ACKNOWLEDGED,MVNOID,USER,USERID,WLANNAME,AUDITIPADDRESS,AUDITUSERUUID,AUDITOBJECT,AUDITACTION,AUDITTENANTUUID,AUDITOBJECTUUID,AUTHTYPE,AUDITTYPE,H20SuppportEnabled,AaaSuppportEnabled,GppSuppportEnabled,Type,RogueMac,SSID,ALARMSTATE,DEVICENAME,SWITCH,SWITCH_GROUP,ZoneAffinityProfileId,FIRMWARE_TYPE,SCHEDULED_TIME,VLAN,FAMILY_ID,SWITCH_ID,transactionId,hasLayerThreeConfig,clientAuthType,clientIpv4Addr,clientIpv6Addr,clientMac,clientUserName,switchName]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=CONTROLBLADE DATABLADE DOMAIN ZONE THIRD_PARTY_ZONE APGROUP WLANGROUP INDOORMAP AP WLAN ProtocolType TIMERANGE RADIOID WLANID CATEGORY CLIENT CP DP CLUSTER NODE BLADE SYNCEDSTATUS OSTYPE APP PORT STATUS REGISTRATIONSTATE GATEWAY APIPADDRESS CLIENTIPADDRESS SEVERITY ACKNOWLEDGED MVNOID USER USERID WLANNAME AUDITIPADDRESS AUDITUSERUUID AUDITOBJECT AUDITACTION AUDITTENANTUUID AUDITOBJECTUUID AUTHTYPE AUDITTYPE H20SuppportEnabled AaaSuppportEnabled GppSuppportEnabled Type RogueMac SSID ALARMSTATE DEVICENAME SWITCH SWITCH_GROUP ZoneAffinityProfileId FIRMWARE_TYPE SCHEDULED_TIME VLAN FAMILY_ID SWITCH_ID transactionId hasLayerThreeConfig clientAuthType clientIpv4Addr clientIpv6Addr clientMac clientUserName switchName"`

	// Value
	// Value to search
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetExtraFiltersType() *SwitchMCommonQueryCriteriaSuperSetExtraFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetExtraFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	// Constraints:
	//    - nullable
	//    - oneof:[CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,ProtocolType,TIMERANGE,RADIOID,WLANID,CATEGORY,CLIENT,CP,DP,CLUSTER,NODE,BLADE,SYNCEDSTATUS,OSTYPE,APP,PORT,STATUS,REGISTRATIONSTATE,GATEWAY,APIPADDRESS,CLIENTIPADDRESS,SEVERITY,ACKNOWLEDGED,MVNOID,USER,USERID,WLANNAME,AUDITIPADDRESS,AUDITUSERUUID,AUDITOBJECT,AUDITACTION,AUDITTENANTUUID,AUDITOBJECTUUID,AUTHTYPE,AUDITTYPE,H20SuppportEnabled,AaaSuppportEnabled,GppSuppportEnabled,Type,RogueMac,SSID,ALARMSTATE,DEVICENAME,SWITCH,ZoneAffinityProfileId,FIRMWARE_TYPE,SCHEDULED_TIME,VLAN,FAMILY_ID,SWITCH_ID,switchStatus.alerts,transactionId,hasLayerThreeConfig]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=CONTROLBLADE DATABLADE DOMAIN ZONE THIRD_PARTY_ZONE APGROUP WLANGROUP INDOORMAP AP WLAN ProtocolType TIMERANGE RADIOID WLANID CATEGORY CLIENT CP DP CLUSTER NODE BLADE SYNCEDSTATUS OSTYPE APP PORT STATUS REGISTRATIONSTATE GATEWAY APIPADDRESS CLIENTIPADDRESS SEVERITY ACKNOWLEDGED MVNOID USER USERID WLANNAME AUDITIPADDRESS AUDITUSERUUID AUDITOBJECT AUDITACTION AUDITTENANTUUID AUDITOBJECTUUID AUTHTYPE AUDITTYPE H20SuppportEnabled AaaSuppportEnabled GppSuppportEnabled Type RogueMac SSID ALARMSTATE DEVICENAME SWITCH ZoneAffinityProfileId FIRMWARE_TYPE SCHEDULED_TIME VLAN FAMILY_ID SWITCH_ID switchStatus.alerts transactionId hasLayerThreeConfig"`

	// Value
	// Value not to search
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType() *SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetExtraNotFiltersType)
	return m
}

type SwitchMCommonQueryCriteriaSuperSetFiltersType struct {
	// Operator
	// Constraints:
	//    - nullable
	Operator *SwitchMCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	// Constraints:
	//    - nullable
	//    - oneof:[SYSTEM,CATEGORY,CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,BLADE,SYNCEDSTATUS,REGISTRATIONSTATE,STATUS,SWITCH_GROUP,PORT]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=SYSTEM CATEGORY CONTROLBLADE DATABLADE DOMAIN ZONE THIRD_PARTY_ZONE APGROUP WLANGROUP INDOORMAP AP WLAN BLADE SYNCEDSTATUS REGISTRATIONSTATE STATUS SWITCH_GROUP PORT"`

	// Value
	// Group ID
	// Constraints:
	//    - nullable
	Value *string `json:"value,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetFiltersType() *SwitchMCommonQueryCriteriaSuperSetFiltersType {
	m := new(SwitchMCommonQueryCriteriaSuperSetFiltersType)
	return m
}

// SwitchMCommonQueryCriteriaSuperSetOptionsType
//
// Specified feature required informaion
// Constraints:
//    - nullable
type SwitchMCommonQueryCriteriaSuperSetOptionsType struct {
	// AcctincludeNa
	// Include Not Available acct service option while returning result
	// Constraints:
	//    - nullable
	AcctincludeNa *bool `json:"acct_includeNa,omitempty"`

	// AccttestableOnly
	// Only get testable service type
	// Constraints:
	//    - nullable
	AccttestableOnly *bool `json:"acct_testableOnly,omitempty"`

	// Accttype
	// Accounting service types to get, use comma to separate, Ex: RADIUS,CGF
	// Constraints:
	//    - nullable
	Accttype *string `json:"acct_type,omitempty"`

	// AuthhostedAaaSupportedEnabled
	// Indicate if Hosted AAA Support is enabled
	// Constraints:
	//    - nullable
	AuthhostedAaaSupportedEnabled *bool `json:"auth_hostedAaaSupportedEnabled,omitempty"`

	// AuthincludeAdGlobal
	// If AD is in list, include only AD with Global Catalog configured
	// Constraints:
	//    - nullable
	AuthincludeAdGlobal *bool `json:"auth_includeAdGlobal,omitempty"`

	// AuthincludeGuest
	// Include Guest auth service while returning result
	// Constraints:
	//    - nullable
	AuthincludeGuest *bool `json:"auth_includeGuest,omitempty"`

	// AuthincludeLocalDb
	// Include LocalDB auth service while returning result
	// Constraints:
	//    - nullable
	AuthincludeLocalDb *bool `json:"auth_includeLocalDb,omitempty"`

	// AuthincludeNa
	// Include Not Available auth service option while returning result
	// Constraints:
	//    - nullable
	AuthincludeNa *bool `json:"auth_includeNa,omitempty"`

	// AuthplmnIdentifierEnabled
	// Indicate if Configure PLMN identifier is enabled
	// Constraints:
	//    - nullable
	AuthplmnIdentifierEnabled *bool `json:"auth_plmnIdentifierEnabled,omitempty"`

	// AuthrealmType
	// To get specific authentication service information for configuring realm based authentication profile
	// Constraints:
	//    - nullable
	//    - oneof:[ALL,RADIUS]
	AuthrealmType *string `json:"auth_realmType,omitempty" validate:"omitempty,oneof=ALL RADIUS"`

	// AuthtestableOnly
	// Only get testable service type
	// Constraints:
	//    - nullable
	AuthtestableOnly *bool `json:"auth_testableOnly,omitempty"`

	// Authtype
	// Authentication service types to get, use comma to separate, Ex: RADIUS,AD
	// Constraints:
	//    - nullable
	Authtype *string `json:"auth_type,omitempty"`

	// Forwardingtype
	// Forwarding service types to get, use comma to separate, Ex: L2oGRE,TTGPDG,Bridge,Advanced
	// Constraints:
	//    - nullable
	Forwardingtype *string `json:"forwarding_type,omitempty"`

	// GlobalFilterId
	// Specify GlobalFilter ID for query
	// Constraints:
	//    - nullable
	GlobalFilterId *string `json:"globalFilterId,omitempty"`

	// IncludeSharedResources
	// Whether to include the resources of parent domain or not
	// Constraints:
	//    - nullable
	IncludeSharedResources *bool `json:"includeSharedResources,omitempty"`

	// IncludeUsers
	// Should also retrieve users or not
	// Constraints:
	//    - nullable
	IncludeUsers *bool `json:"includeUsers,omitempty"`

	// INCLUDERBACMETADATA
	// Whether to include RBAC metadata or not
	// Constraints:
	//    - nullable
	INCLUDERBACMETADATA *bool `json:"INCLUDE_RBAC_METADATA,omitempty"`

	// InMap
	// Specify inMap status for query
	// Constraints:
	//    - nullable
	InMap *bool `json:"inMap,omitempty"`

	// TENANTID
	// Specify Tenant ID for query
	// Constraints:
	//    - nullable
	TENANTID *string `json:"TENANT_ID,omitempty"`
}

func NewSwitchMCommonQueryCriteriaSuperSetOptionsType() *SwitchMCommonQueryCriteriaSuperSetOptionsType {
	m := new(SwitchMCommonQueryCriteriaSuperSetOptionsType)
	return m
}

type SwitchMCommonRbacMetadata struct {
	// RbacMetadata
	// Constraints:
	//    - nullable
	RbacMetadata []string `json:"rbacMetadata,omitempty" validate:"omitempty,dive"`
}

func NewSwitchMCommonRbacMetadata() *SwitchMCommonRbacMetadata {
	m := new(SwitchMCommonRbacMetadata)
	return m
}

type SwitchMCommonTimeRange struct {
	// End
	// end time for collecting data
	// Constraints:
	//    - nullable
	End *float64 `json:"end,omitempty"`

	// Field
	// time field for collecting data
	// Constraints:
	//    - nullable
	//    - oneof:[insertionTime]
	Field *string `json:"field,omitempty" validate:"omitempty,oneof=insertionTime"`

	// Interval
	// time interval in second
	// Constraints:
	//    - nullable
	Interval *float64 `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	// Constraints:
	//    - nullable
	Start *float64 `json:"start,omitempty"`
}

func NewSwitchMCommonTimeRange() *SwitchMCommonTimeRange {
	m := new(SwitchMCommonTimeRange)
	return m
}
