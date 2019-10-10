package group

// API Version: v8_1

import (
	"encoding/json"
)

type AuditID struct {
	// ID
	// Audit Id
	ID *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

type ClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues *ClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

	// ID
	// Identifier of the client
	ID *string `json:"id,omitempty"`

	// Label
	// Label of the client
	Label *string `json:"label,omitempty"`

	// Type
	// Type of the client
	Type *string `json:"type,omitempty"`
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

type CloneConfigByGroup struct {
	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
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

type DeleteSwitchGroupResult struct {
	*AuditID
}

type ErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

type GetConfigBySwitch struct {
	// ID
	// ID
	ID *string `json:"id,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
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

type SwitchGroup struct {
	// CreateDatetime
	// Create datetime of the switch group
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorID
	// Creator Id of the switch group
	CreatorID *string `json:"creatorId,omitempty"`

	// Description
	// Description of the switch group
	Description *string `json:"description,omitempty"`

	// DomainID
	// Identifier of the management domain to which the switch group belong
	DomainID *string `json:"domainId,omitempty"`

	// Firmware
	// Firmware of the switch group
	Firmware *string `json:"firmware,omitempty"`

	// ID
	// Identifier of the switch group
	ID *string `json:"id,omitempty"`

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

	// SwitchGroupLevelOneID
	// Level one Id of the switch group
	SwitchGroupLevelOneID *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoID
	// Level two Id of the switch group
	SwitchGroupLevelTwoID *string `json:"switchGroupLevelTwoId,omitempty"`

	// TenantID
	// Tenant Id of the switch group
	TenantID *string `json:"tenantId,omitempty"`
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

type UpdateSwitchGroup struct {
	*SwitchGroup
}

type UpdateSwitchGroupResult struct {
	*AuditID
}
