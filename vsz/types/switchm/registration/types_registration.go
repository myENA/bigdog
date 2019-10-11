package registration

// API Version: v8_1

import (
	"encoding/json"
)

type ClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues *ClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

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

func NewClientObjectID() *ClientObjectID {
	clientObjectIDType := new(ClientObjectID)
	return clientObjectIDType
}

func NewClientObjectIDWithDefaults() *ClientObjectID {
	clientObjectIDType := new(ClientObjectID)
	return clientObjectIDType
}

// ClientObjectIDExtraValuesType
//
// Extra values of the client
type ClientObjectIDExtraValuesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ClientObjectIDExtraValuesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ClientObjectIDExtraValuesType{XAdditionalProperties: tmp}
	return nil
}

func (t *ClientObjectIDExtraValuesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewClientObjectIDExtraValuesType() *ClientObjectIDExtraValuesType {
	clientObjectIDExtraValuesTypeType := new(ClientObjectIDExtraValuesType)
	return clientObjectIDExtraValuesTypeType
}

func NewClientObjectIDExtraValuesTypeWithDefaults() *ClientObjectIDExtraValuesType {
	clientObjectIDExtraValuesTypeType := new(ClientObjectIDExtraValuesType)
	return clientObjectIDExtraValuesTypeType
}

type CreateResult struct {
	Data *ClientObjectID `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *CreateResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule create result
	MetaData *CreateResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Create result success or not
	Success *bool `json:"success,omitempty"`
}

func NewCreateResult() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

func NewCreateResultWithDefaults() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

// CreateResultExtraType
//
// Any additional response
type CreateResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *CreateResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = CreateResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *CreateResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewCreateResultExtraType() *CreateResultExtraType {
	createResultExtraTypeType := new(CreateResultExtraType)
	return createResultExtraTypeType
}

func NewCreateResultExtraTypeWithDefaults() *CreateResultExtraType {
	createResultExtraTypeType := new(CreateResultExtraType)
	return createResultExtraTypeType
}

// CreateResultMetaDataType
//
// Matadata of Rule create result
type CreateResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *CreateResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = CreateResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *CreateResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewCreateResultMetaDataType() *CreateResultMetaDataType {
	createResultMetaDataTypeType := new(CreateResultMetaDataType)
	return createResultMetaDataTypeType
}

func NewCreateResultMetaDataTypeWithDefaults() *CreateResultMetaDataType {
	createResultMetaDataTypeType := new(CreateResultMetaDataType)
	return createResultMetaDataTypeType
}

type DeleteMultipleResult struct {
	Data *List `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *DeleteMultipleResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of delete multiple rules result
	MetaData *DeleteMultipleResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete multiple result success or not
	Success *bool `json:"success,omitempty"`
}

func NewDeleteMultipleResult() *DeleteMultipleResult {
	deleteMultipleResultType := new(DeleteMultipleResult)
	return deleteMultipleResultType
}

func NewDeleteMultipleResultWithDefaults() *DeleteMultipleResult {
	deleteMultipleResultType := new(DeleteMultipleResult)
	return deleteMultipleResultType
}

// DeleteMultipleResultExtraType
//
// Any additional response
type DeleteMultipleResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *DeleteMultipleResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = DeleteMultipleResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *DeleteMultipleResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewDeleteMultipleResultExtraType() *DeleteMultipleResultExtraType {
	deleteMultipleResultExtraTypeType := new(DeleteMultipleResultExtraType)
	return deleteMultipleResultExtraTypeType
}

func NewDeleteMultipleResultExtraTypeWithDefaults() *DeleteMultipleResultExtraType {
	deleteMultipleResultExtraTypeType := new(DeleteMultipleResultExtraType)
	return deleteMultipleResultExtraTypeType
}

// DeleteMultipleResultMetaDataType
//
// Matadata of delete multiple rules result
type DeleteMultipleResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *DeleteMultipleResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = DeleteMultipleResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *DeleteMultipleResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewDeleteMultipleResultMetaDataType() *DeleteMultipleResultMetaDataType {
	deleteMultipleResultMetaDataTypeType := new(DeleteMultipleResultMetaDataType)
	return deleteMultipleResultMetaDataTypeType
}

func NewDeleteMultipleResultMetaDataTypeWithDefaults() *DeleteMultipleResultMetaDataType {
	deleteMultipleResultMetaDataTypeType := new(DeleteMultipleResultMetaDataType)
	return deleteMultipleResultMetaDataTypeType
}

type DeleteResult struct {
	Data *ClientObjectID `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *DeleteResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule delete result
	MetaData *DeleteResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete result success or not
	Success *bool `json:"success,omitempty"`
}

func NewDeleteResult() *DeleteResult {
	deleteResultType := new(DeleteResult)
	return deleteResultType
}

func NewDeleteResultWithDefaults() *DeleteResult {
	deleteResultType := new(DeleteResult)
	return deleteResultType
}

// DeleteResultExtraType
//
// Any additional response
type DeleteResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *DeleteResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = DeleteResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *DeleteResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewDeleteResultExtraType() *DeleteResultExtraType {
	deleteResultExtraTypeType := new(DeleteResultExtraType)
	return deleteResultExtraTypeType
}

func NewDeleteResultExtraTypeWithDefaults() *DeleteResultExtraType {
	deleteResultExtraTypeType := new(DeleteResultExtraType)
	return deleteResultExtraTypeType
}

// DeleteResultMetaDataType
//
// Matadata of Rule delete result
type DeleteResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *DeleteResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = DeleteResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *DeleteResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewDeleteResultMetaDataType() *DeleteResultMetaDataType {
	deleteResultMetaDataTypeType := new(DeleteResultMetaDataType)
	return deleteResultMetaDataTypeType
}

func NewDeleteResultMetaDataTypeWithDefaults() *DeleteResultMetaDataType {
	deleteResultMetaDataTypeType := new(DeleteResultMetaDataType)
	return deleteResultMetaDataTypeType
}

type ErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewErrorObject() *ErrorObject {
	errorObjectType := new(ErrorObject)
	return errorObjectType
}

func NewErrorObjectWithDefaults() *ErrorObject {
	errorObjectType := new(ErrorObject)
	return errorObjectType
}

type List struct {
	// Extra
	// Any additional response data
	Extra *ListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more registration rule after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*RegistrationRule `json:"list,omitempty"`

	// RawDataTotalCount
	// Registration rule list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Registration rule list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewList() *List {
	listType := new(List)
	return listType
}

func NewListWithDefaults() *List {
	listType := new(List)
	return listType
}

// ListExtraType
//
// Any additional response data
type ListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *ListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewListExtraType() *ListExtraType {
	listExtraTypeType := new(ListExtraType)
	return listExtraTypeType
}

func NewListExtraTypeWithDefaults() *ListExtraType {
	listExtraTypeType := new(ListExtraType)
	return listExtraTypeType
}

type ModifyResult struct {
	Data *ClientObjectID `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *ModifyResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of rule modify result
	MetaData *ModifyResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Modify result success or not
	Success *bool `json:"success,omitempty"`
}

func NewModifyResult() *ModifyResult {
	modifyResultType := new(ModifyResult)
	return modifyResultType
}

func NewModifyResultWithDefaults() *ModifyResult {
	modifyResultType := new(ModifyResult)
	return modifyResultType
}

// ModifyResultExtraType
//
// Any additional response
type ModifyResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ModifyResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ModifyResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *ModifyResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewModifyResultExtraType() *ModifyResultExtraType {
	modifyResultExtraTypeType := new(ModifyResultExtraType)
	return modifyResultExtraTypeType
}

func NewModifyResultExtraTypeWithDefaults() *ModifyResultExtraType {
	modifyResultExtraTypeType := new(ModifyResultExtraType)
	return modifyResultExtraTypeType
}

// ModifyResultMetaDataType
//
// Matadata of rule modify result
type ModifyResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ModifyResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ModifyResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *ModifyResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewModifyResultMetaDataType() *ModifyResultMetaDataType {
	modifyResultMetaDataTypeType := new(ModifyResultMetaDataType)
	return modifyResultMetaDataTypeType
}

func NewModifyResultMetaDataTypeWithDefaults() *ModifyResultMetaDataType {
	modifyResultMetaDataTypeType := new(ModifyResultMetaDataType)
	return modifyResultMetaDataTypeType
}

type RegistrationRule struct {
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
	//    - nullable
	//    - oneof:[IP_RANGE,SUBNET,MODEL_NUMBER]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=IP_RANGE SUBNET MODEL_NUMBER"`
}

func NewRegistrationRule() *RegistrationRule {
	registrationRuleType := new(RegistrationRule)
	return registrationRuleType
}

func NewRegistrationRuleWithDefaults() *RegistrationRule {
	registrationRuleType := new(RegistrationRule)
	return registrationRuleType
}

type RuleQueryResultList struct {
	Data *List `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *RuleQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule query result
	MetaData *RuleQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Rule query result success or not
	Success *bool `json:"success,omitempty"`
}

func NewRuleQueryResultList() *RuleQueryResultList {
	ruleQueryResultListType := new(RuleQueryResultList)
	return ruleQueryResultListType
}

func NewRuleQueryResultListWithDefaults() *RuleQueryResultList {
	ruleQueryResultListType := new(RuleQueryResultList)
	return ruleQueryResultListType
}

// RuleQueryResultListExtraType
//
// Any additional response
type RuleQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *RuleQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = RuleQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *RuleQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewRuleQueryResultListExtraType() *RuleQueryResultListExtraType {
	ruleQueryResultListExtraTypeType := new(RuleQueryResultListExtraType)
	return ruleQueryResultListExtraTypeType
}

func NewRuleQueryResultListExtraTypeWithDefaults() *RuleQueryResultListExtraType {
	ruleQueryResultListExtraTypeType := new(RuleQueryResultListExtraType)
	return ruleQueryResultListExtraTypeType
}

// RuleQueryResultListMetaDataType
//
// Matadata of Rule query result
type RuleQueryResultListMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *RuleQueryResultListMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = RuleQueryResultListMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *RuleQueryResultListMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewRuleQueryResultListMetaDataType() *RuleQueryResultListMetaDataType {
	ruleQueryResultListMetaDataTypeType := new(RuleQueryResultListMetaDataType)
	return ruleQueryResultListMetaDataTypeType
}

func NewRuleQueryResultListMetaDataTypeWithDefaults() *RuleQueryResultListMetaDataType {
	ruleQueryResultListMetaDataTypeType := new(RuleQueryResultListMetaDataType)
	return ruleQueryResultListMetaDataTypeType
}

type RuleUUIDs []string

func NewRuleUUIDs() *RuleUUIDs {
	ruleUUIDsType := make(RuleUUIDs, 0)
	return &ruleUUIDsType
}

func NewRuleUUIDsWithDefaults() *RuleUUIDs {
	ruleUUIDsType := make(RuleUUIDs, 0)
	return &ruleUUIDsType
}
