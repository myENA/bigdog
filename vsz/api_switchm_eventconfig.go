package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMEventconfig struct {
	// Criteria
	// Threshold of each Switch custom event config
	Criteria *int `json:"criteria,omitempty"`

	// Description
	// Description of each Switch custom event config
	Description *string `json:"description,omitempty"`

	// Id
	// Identifier of event config
	Id *int `json:"id,omitempty"`

	// Name
	// Name of each Switch custom event config
	Name *string `json:"name,omitempty"`

	// Severity
	// Severity of each Switch custom event config
	// Constraints:
	//    - oneof:[Warning,Major,Critical]
	Severity *string `json:"severity,omitempty"`

	// TextPattern
	// User defined text pattern of each Switch custom event config
	TextPattern *string `json:"textPattern,omitempty"`

	// TimeWindow
	// Detection time woindow of each Switch custom event config
	// Constraints:
	//    - oneof:[60,120,240,480,720,1440,2880]
	TimeWindow *int `json:"timeWindow,omitempty"`

	// Type
	// Type of each Switch custom event config
	// Constraints:
	//    - oneof:[CPU,Memory,TextPattern]
	Type *string `json:"type,omitempty"`
}

func NewSwitchMEventconfig() *SwitchMEventconfig {
	m := new(SwitchMEventconfig)
	return m
}

type SwitchMEventconfigGetEventConfigList struct {
	// Extra
	// Extra information of responsed Switch custom event config list
	Extra *SwitchMEventconfigGetEventConfigListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// First index of responsed Switch custom event config list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Switch event config
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMEventconfig `json:"list,omitempty"`

	// RawDataTotalCount
	// List of responsed Switch custom event config
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Count of  responsed Switch custom event config
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMEventconfigGetEventConfigList() *SwitchMEventconfigGetEventConfigList {
	m := new(SwitchMEventconfigGetEventConfigList)
	return m
}

// SwitchMEventconfigGetEventConfigListExtraType
//
// Extra information of responsed Switch custom event config list
type SwitchMEventconfigGetEventConfigListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventconfigGetEventConfigListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventconfigGetEventConfigListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventconfigGetEventConfigListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventconfigGetEventConfigListExtraType() *SwitchMEventconfigGetEventConfigListExtraType {
	m := new(SwitchMEventconfigGetEventConfigListExtraType)
	return m
}

type SwitchMEventconfigQueryResponse struct {
	// Data
	// Response data message of Public API
	Data *SwitchMEventconfigQueryResponseDataType `json:"data,omitempty"`

	// Error
	// Response error message of Public API
	Error *SwitchMEventconfigQueryResponseErrorType `json:"error,omitempty"`

	// Extra
	// Extra information of Public API response
	Extra *SwitchMEventconfigQueryResponseExtraType `json:"extra,omitempty"`

	// MetaData
	// Meta-data of Public API response
	MetaData *SwitchMEventconfigQueryResponseMetaDataType `json:"metaData,omitempty"`

	// Success
	// Response success message of Public API
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMEventconfigQueryResponse() *SwitchMEventconfigQueryResponse {
	m := new(SwitchMEventconfigQueryResponse)
	return m
}

// SwitchMEventconfigQueryResponseDataType
//
// Response data message of Public API
type SwitchMEventconfigQueryResponseDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventconfigQueryResponseDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventconfigQueryResponseDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventconfigQueryResponseDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventconfigQueryResponseDataType() *SwitchMEventconfigQueryResponseDataType {
	m := new(SwitchMEventconfigQueryResponseDataType)
	return m
}

// SwitchMEventconfigQueryResponseErrorType
//
// Response error message of Public API
type SwitchMEventconfigQueryResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventconfigQueryResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventconfigQueryResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventconfigQueryResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventconfigQueryResponseErrorType() *SwitchMEventconfigQueryResponseErrorType {
	m := new(SwitchMEventconfigQueryResponseErrorType)
	return m
}

// SwitchMEventconfigQueryResponseExtraType
//
// Extra information of Public API response
type SwitchMEventconfigQueryResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventconfigQueryResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventconfigQueryResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventconfigQueryResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventconfigQueryResponseExtraType() *SwitchMEventconfigQueryResponseExtraType {
	m := new(SwitchMEventconfigQueryResponseExtraType)
	return m
}

// SwitchMEventconfigQueryResponseMetaDataType
//
// Meta-data of Public API response
type SwitchMEventconfigQueryResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventconfigQueryResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventconfigQueryResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventconfigQueryResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventconfigQueryResponseMetaDataType() *SwitchMEventconfigQueryResponseMetaDataType {
	m := new(SwitchMEventconfigQueryResponseMetaDataType)
	return m
}
