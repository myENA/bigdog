package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMRegistrationClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues *SwitchMRegistrationClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

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

func NewSwitchMRegistrationClientObjectID() *SwitchMRegistrationClientObjectID {
	m := new(SwitchMRegistrationClientObjectID)
	return m
}

// SwitchMRegistrationClientObjectIDExtraValuesType
//
// Extra values of the client
type SwitchMRegistrationClientObjectIDExtraValuesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationClientObjectIDExtraValuesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationClientObjectIDExtraValuesType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationClientObjectIDExtraValuesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationClientObjectIDExtraValuesType() *SwitchMRegistrationClientObjectIDExtraValuesType {
	m := new(SwitchMRegistrationClientObjectIDExtraValuesType)
	return m
}

type SwitchMRegistrationCreateResult struct {
	Data *SwitchMRegistrationClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationCreateResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule create result
	MetaData *SwitchMRegistrationCreateResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Create result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationCreateResult() *SwitchMRegistrationCreateResult {
	m := new(SwitchMRegistrationCreateResult)
	return m
}

// SwitchMRegistrationCreateResultExtraType
//
// Any additional response
type SwitchMRegistrationCreateResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationCreateResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationCreateResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationCreateResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationCreateResultExtraType() *SwitchMRegistrationCreateResultExtraType {
	m := new(SwitchMRegistrationCreateResultExtraType)
	return m
}

// SwitchMRegistrationCreateResultMetaDataType
//
// Matadata of Rule create result
type SwitchMRegistrationCreateResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationCreateResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationCreateResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationCreateResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationCreateResultMetaDataType() *SwitchMRegistrationCreateResultMetaDataType {
	m := new(SwitchMRegistrationCreateResultMetaDataType)
	return m
}

type SwitchMRegistrationDeleteMultipleResult struct {
	Data *SwitchMRegistrationList `json:"data,omitempty"`

	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationDeleteMultipleResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of delete multiple rules result
	MetaData *SwitchMRegistrationDeleteMultipleResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete multiple result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationDeleteMultipleResult() *SwitchMRegistrationDeleteMultipleResult {
	m := new(SwitchMRegistrationDeleteMultipleResult)
	return m
}

// SwitchMRegistrationDeleteMultipleResultExtraType
//
// Any additional response
type SwitchMRegistrationDeleteMultipleResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationDeleteMultipleResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationDeleteMultipleResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationDeleteMultipleResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationDeleteMultipleResultExtraType() *SwitchMRegistrationDeleteMultipleResultExtraType {
	m := new(SwitchMRegistrationDeleteMultipleResultExtraType)
	return m
}

// SwitchMRegistrationDeleteMultipleResultMetaDataType
//
// Matadata of delete multiple rules result
type SwitchMRegistrationDeleteMultipleResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationDeleteMultipleResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationDeleteMultipleResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationDeleteMultipleResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationDeleteMultipleResultMetaDataType() *SwitchMRegistrationDeleteMultipleResultMetaDataType {
	m := new(SwitchMRegistrationDeleteMultipleResultMetaDataType)
	return m
}

type SwitchMRegistrationDeleteResult struct {
	Data *SwitchMRegistrationClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationDeleteResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule delete result
	MetaData *SwitchMRegistrationDeleteResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationDeleteResult() *SwitchMRegistrationDeleteResult {
	m := new(SwitchMRegistrationDeleteResult)
	return m
}

// SwitchMRegistrationDeleteResultExtraType
//
// Any additional response
type SwitchMRegistrationDeleteResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationDeleteResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationDeleteResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationDeleteResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationDeleteResultExtraType() *SwitchMRegistrationDeleteResultExtraType {
	m := new(SwitchMRegistrationDeleteResultExtraType)
	return m
}

// SwitchMRegistrationDeleteResultMetaDataType
//
// Matadata of Rule delete result
type SwitchMRegistrationDeleteResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationDeleteResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationDeleteResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationDeleteResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationDeleteResultMetaDataType() *SwitchMRegistrationDeleteResultMetaDataType {
	m := new(SwitchMRegistrationDeleteResultMetaDataType)
	return m
}

type SwitchMRegistrationErrorObject struct {
	List []string `json:"list"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMRegistrationErrorObject() *SwitchMRegistrationErrorObject {
	m := new(SwitchMRegistrationErrorObject)
	return m
}

type SwitchMRegistrationList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMRegistrationListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more registration rule after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMRegistrationRule `json:"list"`

	// RawDataTotalCount
	// Registration rule list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Registration rule list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMRegistrationList() *SwitchMRegistrationList {
	m := new(SwitchMRegistrationList)
	return m
}

// SwitchMRegistrationListExtraType
//
// Any additional response data
type SwitchMRegistrationListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationListExtraType() *SwitchMRegistrationListExtraType {
	m := new(SwitchMRegistrationListExtraType)
	return m
}

type SwitchMRegistrationModifyResult struct {
	Data *SwitchMRegistrationClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationModifyResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of rule modify result
	MetaData *SwitchMRegistrationModifyResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Modify result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationModifyResult() *SwitchMRegistrationModifyResult {
	m := new(SwitchMRegistrationModifyResult)
	return m
}

// SwitchMRegistrationModifyResultExtraType
//
// Any additional response
type SwitchMRegistrationModifyResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationModifyResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationModifyResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationModifyResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationModifyResultExtraType() *SwitchMRegistrationModifyResultExtraType {
	m := new(SwitchMRegistrationModifyResultExtraType)
	return m
}

// SwitchMRegistrationModifyResultMetaDataType
//
// Matadata of rule modify result
type SwitchMRegistrationModifyResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationModifyResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationModifyResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationModifyResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationModifyResultMetaDataType() *SwitchMRegistrationModifyResultMetaDataType {
	m := new(SwitchMRegistrationModifyResultMetaDataType)
	return m
}

type SwitchMRegistrationRule struct {
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
	Type *string `json:"type,omitempty" validate:"oneof=IP_RANGE SUBNET MODEL_NUMBER"`
}

func NewSwitchMRegistrationRule() *SwitchMRegistrationRule {
	m := new(SwitchMRegistrationRule)
	return m
}

type SwitchMRegistrationRuleQueryResultList struct {
	Data *SwitchMRegistrationList `json:"data,omitempty"`

	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationRuleQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule query result
	MetaData *SwitchMRegistrationRuleQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Rule query result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationRuleQueryResultList() *SwitchMRegistrationRuleQueryResultList {
	m := new(SwitchMRegistrationRuleQueryResultList)
	return m
}

// SwitchMRegistrationRuleQueryResultListExtraType
//
// Any additional response
type SwitchMRegistrationRuleQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRuleQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRuleQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRuleQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRuleQueryResultListExtraType() *SwitchMRegistrationRuleQueryResultListExtraType {
	m := new(SwitchMRegistrationRuleQueryResultListExtraType)
	return m
}

// SwitchMRegistrationRuleQueryResultListMetaDataType
//
// Matadata of Rule query result
type SwitchMRegistrationRuleQueryResultListMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRuleQueryResultListMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRuleQueryResultListMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRuleQueryResultListMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRuleQueryResultListMetaDataType() *SwitchMRegistrationRuleQueryResultListMetaDataType {
	m := new(SwitchMRegistrationRuleQueryResultListMetaDataType)
	return m
}

type SwitchMRegistrationRuleUUIDs []string

func MakeSwitchMRegistrationRuleUUIDs() SwitchMRegistrationRuleUUIDs {
	m := make(SwitchMRegistrationRuleUUIDs, 0)
	return m
}
