package group

// API Version: v8_1

import (
	"encoding/json"
)

type AuditId struct {
	// Id
	// Audit Id
	Id *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

func NewAuditId() *AuditId {
	auditIdType := new(AuditId)
	return auditIdType
}

func NewAuditIdWithDefaults() *AuditId {
	auditIdType := new(AuditId)
	return auditIdType
}

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

type CloneConfigByGroup struct {
	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewCloneConfigByGroup() *CloneConfigByGroup {
	cloneConfigByGroupType := new(CloneConfigByGroup)
	return cloneConfigByGroupType
}

func NewCloneConfigByGroupWithDefaults() *CloneConfigByGroup {
	cloneConfigByGroupType := new(CloneConfigByGroup)
	return cloneConfigByGroupType
}

type CloneConfigBySwitch struct {
	// Config
	// Config
	Config *string `json:"config,omitempty"`

	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewCloneConfigBySwitch() *CloneConfigBySwitch {
	cloneConfigBySwitchType := new(CloneConfigBySwitch)
	return cloneConfigBySwitchType
}

func NewCloneConfigBySwitchWithDefaults() *CloneConfigBySwitch {
	cloneConfigBySwitchType := new(CloneConfigBySwitch)
	return cloneConfigBySwitchType
}

type DeleteSwitchGroupResult struct {
	*AuditId
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

type GetConfigBySwitch struct {
	// Id
	// ID
	Id *string `json:"id,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewGetConfigBySwitch() *GetConfigBySwitch {
	getConfigBySwitchType := new(GetConfigBySwitch)
	return getConfigBySwitchType
}

func NewGetConfigBySwitchWithDefaults() *GetConfigBySwitch {
	getConfigBySwitchType := new(GetConfigBySwitch)
	return getConfigBySwitchType
}

type GroupsByIdsQueryResultList struct {
	Data *QueryResultList `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *GroupsByIdsQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Metadata of query result list
	MetaData *GroupsByIdsQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Query result success or not
	Success *bool `json:"success,omitempty"`
}

func NewGroupsByIdsQueryResultList() *GroupsByIdsQueryResultList {
	groupsByIdsQueryResultListType := new(GroupsByIdsQueryResultList)
	return groupsByIdsQueryResultListType
}

func NewGroupsByIdsQueryResultListWithDefaults() *GroupsByIdsQueryResultList {
	groupsByIdsQueryResultListType := new(GroupsByIdsQueryResultList)
	return groupsByIdsQueryResultListType
}

// GroupsByIdsQueryResultListExtraType
//
// Any additional response
type GroupsByIdsQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *GroupsByIdsQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = GroupsByIdsQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *GroupsByIdsQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewGroupsByIdsQueryResultListExtraType() *GroupsByIdsQueryResultListExtraType {
	groupsByIdsQueryResultListExtraTypeType := new(GroupsByIdsQueryResultListExtraType)
	return groupsByIdsQueryResultListExtraTypeType
}

func NewGroupsByIdsQueryResultListExtraTypeWithDefaults() *GroupsByIdsQueryResultListExtraType {
	groupsByIdsQueryResultListExtraTypeType := new(GroupsByIdsQueryResultListExtraType)
	return groupsByIdsQueryResultListExtraTypeType
}

// GroupsByIdsQueryResultListMetaDataType
//
// Metadata of query result list
type GroupsByIdsQueryResultListMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *GroupsByIdsQueryResultListMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = GroupsByIdsQueryResultListMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *GroupsByIdsQueryResultListMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewGroupsByIdsQueryResultListMetaDataType() *GroupsByIdsQueryResultListMetaDataType {
	groupsByIdsQueryResultListMetaDataTypeType := new(GroupsByIdsQueryResultListMetaDataType)
	return groupsByIdsQueryResultListMetaDataTypeType
}

func NewGroupsByIdsQueryResultListMetaDataTypeWithDefaults() *GroupsByIdsQueryResultListMetaDataType {
	groupsByIdsQueryResultListMetaDataTypeType := new(GroupsByIdsQueryResultListMetaDataType)
	return groupsByIdsQueryResultListMetaDataTypeType
}

type QueryResultList struct {
	// Extra
	// Extra information for query result list
	Extra *QueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first query result returned out of the complete query result list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more query result after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ClientObjectID `json:"list,omitempty"`

	// RawDataTotalCount
	// Query result count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total query result count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewQueryResultList() *QueryResultList {
	queryResultListType := new(QueryResultList)
	return queryResultListType
}

func NewQueryResultListWithDefaults() *QueryResultList {
	queryResultListType := new(QueryResultList)
	return queryResultListType
}

// QueryResultListExtraType
//
// Extra information for query result list
type QueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *QueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = QueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *QueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewQueryResultListExtraType() *QueryResultListExtraType {
	queryResultListExtraTypeType := new(QueryResultListExtraType)
	return queryResultListExtraTypeType
}

func NewQueryResultListExtraTypeWithDefaults() *QueryResultListExtraType {
	queryResultListExtraTypeType := new(QueryResultListExtraType)
	return queryResultListExtraTypeType
}

type SwitchGroup struct {
	// CreateDatetime
	// Create datetime of the switch group
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorId
	// Creator Id of the switch group
	CreatorId *string `json:"creatorId,omitempty"`

	// Description
	// Description of the switch group
	Description *string `json:"description,omitempty"`

	// DomainId
	// Identifier of the management domain to which the switch group belong
	DomainId *string `json:"domainId,omitempty"`

	// Firmware
	// Firmware of the switch group
	Firmware *string `json:"firmware,omitempty"`

	// Id
	// Identifier of the switch group
	Id *string `json:"id,omitempty"`

	// LevelOne
	// Level one  of the switch group
	LevelOne *bool `json:"levelOne,omitempty"`

	// LevelTwo
	// Level two of the switch group
	LevelTwo *bool `json:"levelTwo,omitempty"`

	// Name
	// Name of the switch group
	Name *string `json:"name,omitempty"`

	// SampledInstant
	// Sampled instant of the switch group
	SampledInstant *SwitchGroupSampledInstantType `json:"sampledInstant,omitempty"`

	// SwitchGroupLevelOneId
	// Level one Id of the switch group
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Level two Id of the switch group
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// TenantId
	// Tenant Id of the switch group
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSwitchGroup() *SwitchGroup {
	switchGroupType := new(SwitchGroup)
	return switchGroupType
}

func NewSwitchGroupWithDefaults() *SwitchGroup {
	switchGroupType := new(SwitchGroup)
	return switchGroupType
}

type SwitchGroupQueryResult struct {
	*SwitchGroup
}

// SwitchGroupSampledInstantType
//
// Sampled instant of the switch group
type SwitchGroupSampledInstantType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchGroupSampledInstantType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchGroupSampledInstantType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchGroupSampledInstantType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchGroupSampledInstantType() *SwitchGroupSampledInstantType {
	switchGroupSampledInstantTypeType := new(SwitchGroupSampledInstantType)
	return switchGroupSampledInstantTypeType
}

func NewSwitchGroupSampledInstantTypeWithDefaults() *SwitchGroupSampledInstantType {
	switchGroupSampledInstantTypeType := new(SwitchGroupSampledInstantType)
	return switchGroupSampledInstantTypeType
}

type UpdateSwitchGroup struct {
	*SwitchGroup
}

type UpdateSwitchGroupResult struct {
	*AuditId
}
