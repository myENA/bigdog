package ruckus

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMSwitchRegistrationRulesClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues *SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

	// Id
	// Identifier of the client
	Id *string `json:"id,omitempty"`

	// Label
	// Label of the client
	Label *string `json:"label,omitempty"`

	// Type
	// Type of the client
	Type *string `json:"type,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesClientObjectID() *SwitchMSwitchRegistrationRulesClientObjectID {
	m := new(SwitchMSwitchRegistrationRulesClientObjectID)
	return m
}

// SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType
//
// Extra values of the client
type SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType() *SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType {
	m := new(SwitchMSwitchRegistrationRulesClientObjectIDExtraValuesType)
	return m
}

type SwitchMSwitchRegistrationRulesCreateResult struct {
	Data *SwitchMSwitchRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMSwitchRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMSwitchRegistrationRulesCreateResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule create result
	MetaData *SwitchMSwitchRegistrationRulesCreateResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Create result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesCreateResult() *SwitchMSwitchRegistrationRulesCreateResult {
	m := new(SwitchMSwitchRegistrationRulesCreateResult)
	return m
}

// SwitchMSwitchRegistrationRulesCreateResultExtraType
//
// Any additional response
type SwitchMSwitchRegistrationRulesCreateResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesCreateResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesCreateResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesCreateResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesCreateResultExtraType() *SwitchMSwitchRegistrationRulesCreateResultExtraType {
	m := new(SwitchMSwitchRegistrationRulesCreateResultExtraType)
	return m
}

// SwitchMSwitchRegistrationRulesCreateResultMetaDataType
//
// Matadata of Rule create result
type SwitchMSwitchRegistrationRulesCreateResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesCreateResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesCreateResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesCreateResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesCreateResultMetaDataType() *SwitchMSwitchRegistrationRulesCreateResultMetaDataType {
	m := new(SwitchMSwitchRegistrationRulesCreateResultMetaDataType)
	return m
}

type SwitchMSwitchRegistrationRulesDeleteMultipleResult struct {
	Data *SwitchMSwitchRegistrationRulesList `json:"data,omitempty"`

	Error *SwitchMSwitchRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of delete multiple rules result
	MetaData *SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete multiple result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesDeleteMultipleResult() *SwitchMSwitchRegistrationRulesDeleteMultipleResult {
	m := new(SwitchMSwitchRegistrationRulesDeleteMultipleResult)
	return m
}

// SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType
//
// Any additional response
type SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType() *SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType {
	m := new(SwitchMSwitchRegistrationRulesDeleteMultipleResultExtraType)
	return m
}

// SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType
//
// Matadata of delete multiple rules result
type SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType() *SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType {
	m := new(SwitchMSwitchRegistrationRulesDeleteMultipleResultMetaDataType)
	return m
}

type SwitchMSwitchRegistrationRulesDeleteResult struct {
	Data *SwitchMSwitchRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMSwitchRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMSwitchRegistrationRulesDeleteResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule delete result
	MetaData *SwitchMSwitchRegistrationRulesDeleteResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesDeleteResult() *SwitchMSwitchRegistrationRulesDeleteResult {
	m := new(SwitchMSwitchRegistrationRulesDeleteResult)
	return m
}

// SwitchMSwitchRegistrationRulesDeleteResultExtraType
//
// Any additional response
type SwitchMSwitchRegistrationRulesDeleteResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesDeleteResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesDeleteResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesDeleteResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesDeleteResultExtraType() *SwitchMSwitchRegistrationRulesDeleteResultExtraType {
	m := new(SwitchMSwitchRegistrationRulesDeleteResultExtraType)
	return m
}

// SwitchMSwitchRegistrationRulesDeleteResultMetaDataType
//
// Matadata of Rule delete result
type SwitchMSwitchRegistrationRulesDeleteResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesDeleteResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesDeleteResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesDeleteResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesDeleteResultMetaDataType() *SwitchMSwitchRegistrationRulesDeleteResultMetaDataType {
	m := new(SwitchMSwitchRegistrationRulesDeleteResultMetaDataType)
	return m
}

type SwitchMSwitchRegistrationRulesErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesErrorObject() *SwitchMSwitchRegistrationRulesErrorObject {
	m := new(SwitchMSwitchRegistrationRulesErrorObject)
	return m
}

type SwitchMSwitchRegistrationRulesList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchRegistrationRulesListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more registration rule after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchRegistrationRulesRegistrationRule `json:"list,omitempty"`

	// RawDataTotalCount
	// Registration rule list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Registration rule list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesList() *SwitchMSwitchRegistrationRulesList {
	m := new(SwitchMSwitchRegistrationRulesList)
	return m
}

// SwitchMSwitchRegistrationRulesListExtraType
//
// Any additional response data
type SwitchMSwitchRegistrationRulesListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesListExtraType() *SwitchMSwitchRegistrationRulesListExtraType {
	m := new(SwitchMSwitchRegistrationRulesListExtraType)
	return m
}

type SwitchMSwitchRegistrationRulesModifyResult struct {
	Data *SwitchMSwitchRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMSwitchRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMSwitchRegistrationRulesModifyResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of rule modify result
	MetaData *SwitchMSwitchRegistrationRulesModifyResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Modify result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesModifyResult() *SwitchMSwitchRegistrationRulesModifyResult {
	m := new(SwitchMSwitchRegistrationRulesModifyResult)
	return m
}

// SwitchMSwitchRegistrationRulesModifyResultExtraType
//
// Any additional response
type SwitchMSwitchRegistrationRulesModifyResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesModifyResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesModifyResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesModifyResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesModifyResultExtraType() *SwitchMSwitchRegistrationRulesModifyResultExtraType {
	m := new(SwitchMSwitchRegistrationRulesModifyResultExtraType)
	return m
}

// SwitchMSwitchRegistrationRulesModifyResultMetaDataType
//
// Matadata of rule modify result
type SwitchMSwitchRegistrationRulesModifyResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesModifyResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesModifyResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesModifyResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesModifyResultMetaDataType() *SwitchMSwitchRegistrationRulesModifyResultMetaDataType {
	m := new(SwitchMSwitchRegistrationRulesModifyResultMetaDataType)
	return m
}

type SwitchMSwitchRegistrationRulesRegistrationRule struct {
	// CreateDatetime
	// Create datetime of the registration rule
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorId
	// Creator Id of the registration rule
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorName
	// Creator name of the registration rule
	CreatorName *string `json:"creatorName,omitempty"`

	// Description
	// Description of the registration rule
	Description *string `json:"description,omitempty"`

	// GroupName
	// Switch group name of the registration rule
	GroupName *string `json:"groupName,omitempty"`

	// Id
	// Identifier of the registration rule
	Id *string `json:"id,omitempty"`

	// IpFrom
	// Start IP range of the registration rule
	IpFrom *string `json:"ipFrom,omitempty"`

	// IpRange
	// IP range of the registration rule
	IpRange *string `json:"ipRange,omitempty"`

	// IpTo
	// End IP range of the registration rule
	IpTo *string `json:"ipTo,omitempty"`

	// Label
	// Lable of the registration rule
	Label *string `json:"label,omitempty"`

	// ModelNumber
	// Switch Model number of the registration rule
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Network
	// Network of the registration rule
	Network *string `json:"network,omitempty"`

	// Rank
	// Rank of the registration rule
	Rank *int `json:"rank,omitempty"`

	// SubnetMask
	// Subnet mask of the registration rule
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchGroupId
	// Switch group Id of the registration rule
	SwitchGroupId *string `json:"switchGroupId,omitempty"`

	// Type
	// Type of the registration rule
	// Constraints:
	//    - oneof:[IP_RANGE,SUBNET,MODEL_NUMBER]
	Type *string `json:"type,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesRegistrationRule() *SwitchMSwitchRegistrationRulesRegistrationRule {
	m := new(SwitchMSwitchRegistrationRulesRegistrationRule)
	return m
}

type SwitchMSwitchRegistrationRulesRuleQueryResultList struct {
	Data *SwitchMSwitchRegistrationRulesList `json:"data,omitempty"`

	Error *SwitchMSwitchRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule query result
	MetaData *SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Rule query result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchRegistrationRulesRuleQueryResultList() *SwitchMSwitchRegistrationRulesRuleQueryResultList {
	m := new(SwitchMSwitchRegistrationRulesRuleQueryResultList)
	return m
}

// SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType
//
// Any additional response
type SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesRuleQueryResultListExtraType() *SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType {
	m := new(SwitchMSwitchRegistrationRulesRuleQueryResultListExtraType)
	return m
}

// SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType
//
// Matadata of Rule query result
type SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType() *SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType {
	m := new(SwitchMSwitchRegistrationRulesRuleQueryResultListMetaDataType)
	return m
}

type SwitchMSwitchRegistrationRulesRuleUUIDs []string

func MakeSwitchMSwitchRegistrationRulesRuleUUIDs() SwitchMSwitchRegistrationRulesRuleUUIDs {
	m := make(SwitchMSwitchRegistrationRulesRuleUUIDs, 0)
	return m
}
