package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMRegistrationClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	// Constraints:
	//    - nullable
	ExtraValues *SwitchMRegistrationClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

	// Id
	// Identifier of the client
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Label
	// Label of the client
	// Constraints:
	//    - nullable
	Label *string `json:"label,omitempty"`

	// Type
	// Type of the client
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewSwitchMRegistrationClientObjectID() *SwitchMRegistrationClientObjectID {
	m := new(SwitchMRegistrationClientObjectID)
	return m
}

// SwitchMRegistrationClientObjectIDExtraValuesType
//
// Extra values of the client
// Constraints:
//    - nullable
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
	// Data
	// Constraints:
	//    - nullable
	Data *SwitchMRegistrationClientObjectID `json:"data,omitempty"`

	// Error
	// Constraints:
	//    - nullable
	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	// Constraints:
	//    - nullable
	Extra *SwitchMRegistrationCreateResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule create result
	// Constraints:
	//    - nullable
	MetaData *SwitchMRegistrationCreateResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Create result success or not
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationCreateResult() *SwitchMRegistrationCreateResult {
	m := new(SwitchMRegistrationCreateResult)
	return m
}

// SwitchMRegistrationCreateResultExtraType
//
// Any additional response
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
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
	// Data
	// Constraints:
	//    - nullable
	Data *SwitchMRegistrationList `json:"data,omitempty"`

	// Error
	// Constraints:
	//    - nullable
	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	// Constraints:
	//    - nullable
	Extra *SwitchMRegistrationDeleteMultipleResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of delete multiple rules result
	// Constraints:
	//    - nullable
	MetaData *SwitchMRegistrationDeleteMultipleResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete multiple result success or not
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationDeleteMultipleResult() *SwitchMRegistrationDeleteMultipleResult {
	m := new(SwitchMRegistrationDeleteMultipleResult)
	return m
}

// SwitchMRegistrationDeleteMultipleResultExtraType
//
// Any additional response
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
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
	// Data
	// Constraints:
	//    - nullable
	Data *SwitchMRegistrationClientObjectID `json:"data,omitempty"`

	// Error
	// Constraints:
	//    - nullable
	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	// Constraints:
	//    - nullable
	Extra *SwitchMRegistrationDeleteResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule delete result
	// Constraints:
	//    - nullable
	MetaData *SwitchMRegistrationDeleteResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete result success or not
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationDeleteResult() *SwitchMRegistrationDeleteResult {
	m := new(SwitchMRegistrationDeleteResult)
	return m
}

// SwitchMRegistrationDeleteResultExtraType
//
// Any additional response
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
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
	// List
	// Constraints:
	//    - nullable
	List []string `json:"list,omitempty" validate:"omitempty,dive"`

	// Message
	// Constraints:
	//    - nullable
	Message *string `json:"message,omitempty"`

	// MsgKey
	// Constraints:
	//    - nullable
	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMRegistrationErrorObject() *SwitchMRegistrationErrorObject {
	m := new(SwitchMRegistrationErrorObject)
	return m
}

type SwitchMRegistrationList struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMRegistrationListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more registration rule after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMRegistrationRule `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Registration rule list count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Registration rule list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMRegistrationList() *SwitchMRegistrationList {
	m := new(SwitchMRegistrationList)
	return m
}

// SwitchMRegistrationListExtraType
//
// Any additional response data
// Constraints:
//    - nullable
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
	// Data
	// Constraints:
	//    - nullable
	Data *SwitchMRegistrationClientObjectID `json:"data,omitempty"`

	// Error
	// Constraints:
	//    - nullable
	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	// Constraints:
	//    - nullable
	Extra *SwitchMRegistrationModifyResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of rule modify result
	// Constraints:
	//    - nullable
	MetaData *SwitchMRegistrationModifyResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Modify result success or not
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationModifyResult() *SwitchMRegistrationModifyResult {
	m := new(SwitchMRegistrationModifyResult)
	return m
}

// SwitchMRegistrationModifyResultExtraType
//
// Any additional response
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
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
	// Constraints:
	//    - nullable
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorId
	// Creator Id of the registration rule
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorName
	// Creator name of the registration rule
	// Constraints:
	//    - nullable
	CreatorName *string `json:"creatorName,omitempty"`

	// Description
	// Description of the registration rule
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// GroupName
	// Switch group name of the registration rule
	// Constraints:
	//    - nullable
	GroupName *string `json:"groupName,omitempty"`

	// Id
	// Identifier of the registration rule
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IpFrom
	// Start IP range of the registration rule
	// Constraints:
	//    - nullable
	IpFrom *string `json:"ipFrom,omitempty"`

	// IpRange
	// IP range of the registration rule
	// Constraints:
	//    - nullable
	IpRange *string `json:"ipRange,omitempty"`

	// IpTo
	// End IP range of the registration rule
	// Constraints:
	//    - nullable
	IpTo *string `json:"ipTo,omitempty"`

	// Label
	// Lable of the registration rule
	// Constraints:
	//    - nullable
	Label *string `json:"label,omitempty"`

	// ModelNumber
	// Switch Model number of the registration rule
	// Constraints:
	//    - nullable
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Network
	// Network of the registration rule
	// Constraints:
	//    - nullable
	Network *string `json:"network,omitempty"`

	// Rank
	// Rank of the registration rule
	// Constraints:
	//    - nullable
	Rank *int `json:"rank,omitempty"`

	// SubnetMask
	// Subnet mask of the registration rule
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchGroupId
	// Switch group Id of the registration rule
	// Constraints:
	//    - nullable
	SwitchGroupId *string `json:"switchGroupId,omitempty"`

	// Type
	// Type of the registration rule
	// Constraints:
	//    - nullable
	//    - oneof:[IP_RANGE,SUBNET,MODEL_NUMBER]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IP_RANGE SUBNET MODEL_NUMBER"`
}

func NewSwitchMRegistrationRule() *SwitchMRegistrationRule {
	m := new(SwitchMRegistrationRule)
	return m
}

type SwitchMRegistrationRuleQueryResultList struct {
	// Data
	// Constraints:
	//    - nullable
	Data *SwitchMRegistrationList `json:"data,omitempty"`

	// Error
	// Constraints:
	//    - nullable
	Error *SwitchMRegistrationErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	// Constraints:
	//    - nullable
	Extra *SwitchMRegistrationRuleQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule query result
	// Constraints:
	//    - nullable
	MetaData *SwitchMRegistrationRuleQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Rule query result success or not
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationRuleQueryResultList() *SwitchMRegistrationRuleQueryResultList {
	m := new(SwitchMRegistrationRuleQueryResultList)
	return m
}

// SwitchMRegistrationRuleQueryResultListExtraType
//
// Any additional response
// Constraints:
//    - nullable
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
// Constraints:
//    - nullable
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
