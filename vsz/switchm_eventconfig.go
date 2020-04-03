package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMEventConfig struct {
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
	Severity *string `json:"severity,omitempty" validate:"oneof=Warning Major Critical"`

	// TextPattern
	// User defined text pattern of each Switch custom event config
	TextPattern *string `json:"textPattern,omitempty"`

	// TimeWindow
	// Detection time woindow of each Switch custom event config
	// Constraints:
	//    - oneof:[60,120,240,480,720,1440,2880]
	TimeWindow *int `json:"timeWindow,omitempty" validate:"oneof=60 120 240 480 720 1440 2880"`

	// Type
	// Type of each Switch custom event config
	// Constraints:
	//    - oneof:[CPU,Memory,TextPattern]
	Type *string `json:"type,omitempty" validate:"oneof=CPU Memory TextPattern"`
}

func NewSwitchMEventConfig() *SwitchMEventConfig {
	m := new(SwitchMEventConfig)
	return m
}

type SwitchMEventConfigGetEventConfigList struct {
	// Extra
	// Extra information of responsed Switch custom event config list
	Extra *SwitchMEventConfigGetEventConfigListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// First index of responsed Switch custom event config list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Switch event config
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMEventConfig `json:"list"`

	// RawDataTotalCount
	// List of responsed Switch custom event config
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Count of  responsed Switch custom event config
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMEventConfigGetEventConfigList() *SwitchMEventConfigGetEventConfigList {
	m := new(SwitchMEventConfigGetEventConfigList)
	return m
}

// SwitchMEventConfigGetEventConfigListExtraType
//
// Extra information of responsed Switch custom event config list
type SwitchMEventConfigGetEventConfigListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventConfigGetEventConfigListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventConfigGetEventConfigListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventConfigGetEventConfigListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventConfigGetEventConfigListExtraType() *SwitchMEventConfigGetEventConfigListExtraType {
	m := new(SwitchMEventConfigGetEventConfigListExtraType)
	return m
}

type SwitchMEventConfigQueryResponse struct {
	// Data
	// Response data message of Public API
	Data *SwitchMEventConfigQueryResponseDataType `json:"data,omitempty"`

	// Error
	// Response error message of Public API
	Error *SwitchMEventConfigQueryResponseErrorType `json:"error,omitempty"`

	// Extra
	// Extra information of Public API response
	Extra *SwitchMEventConfigQueryResponseExtraType `json:"extra,omitempty"`

	// MetaData
	// Meta-data of Public API response
	MetaData *SwitchMEventConfigQueryResponseMetaDataType `json:"metaData,omitempty"`

	// Success
	// Response success message of Public API
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMEventConfigQueryResponse() *SwitchMEventConfigQueryResponse {
	m := new(SwitchMEventConfigQueryResponse)
	return m
}

// SwitchMEventConfigQueryResponseDataType
//
// Response data message of Public API
type SwitchMEventConfigQueryResponseDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventConfigQueryResponseDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventConfigQueryResponseDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventConfigQueryResponseDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventConfigQueryResponseDataType() *SwitchMEventConfigQueryResponseDataType {
	m := new(SwitchMEventConfigQueryResponseDataType)
	return m
}

// SwitchMEventConfigQueryResponseErrorType
//
// Response error message of Public API
type SwitchMEventConfigQueryResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventConfigQueryResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventConfigQueryResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventConfigQueryResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventConfigQueryResponseErrorType() *SwitchMEventConfigQueryResponseErrorType {
	m := new(SwitchMEventConfigQueryResponseErrorType)
	return m
}

// SwitchMEventConfigQueryResponseExtraType
//
// Extra information of Public API response
type SwitchMEventConfigQueryResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventConfigQueryResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventConfigQueryResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventConfigQueryResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventConfigQueryResponseExtraType() *SwitchMEventConfigQueryResponseExtraType {
	m := new(SwitchMEventConfigQueryResponseExtraType)
	return m
}

// SwitchMEventConfigQueryResponseMetaDataType
//
// Meta-data of Public API response
type SwitchMEventConfigQueryResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMEventConfigQueryResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMEventConfigQueryResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMEventConfigQueryResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMEventConfigQueryResponseMetaDataType() *SwitchMEventConfigQueryResponseMetaDataType {
	m := new(SwitchMEventConfigQueryResponseMetaDataType)
	return m
}
