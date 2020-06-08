package ruckus

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMSwitchEventConfigEventConfig struct {
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

func NewSwitchMSwitchEventConfigEventConfig() *SwitchMSwitchEventConfigEventConfig {
	m := new(SwitchMSwitchEventConfigEventConfig)
	return m
}

type SwitchMSwitchEventConfigGetEventConfigList struct {
	// Extra
	// Extra information of responsed Switch custom event config list
	Extra *SwitchMSwitchEventConfigGetEventConfigList `json:"extra,omitempty"`

	// FirstIndex
	// First index of responsed Switch custom event config list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Switch event config
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchEventConfigGetEventConfigList `json:"list,omitempty"`

	// RawDataTotalCount
	// List of responsed Switch custom event config
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Count of  responsed Switch custom event config
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchEventConfigGetEventConfigList() *SwitchMSwitchEventConfigGetEventConfigList {
	m := new(SwitchMSwitchEventConfigGetEventConfigList)
	return m
}

// SwitchMSwitchEventConfigGetEventConfigListExtraType
//
// Extra information of responsed Switch custom event config list
type SwitchMSwitchEventConfigGetEventConfigListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchEventConfigGetEventConfigListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchEventConfigGetEventConfigListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchEventConfigGetEventConfigListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchEventConfigGetEventConfigListExtraType() *SwitchMSwitchEventConfigGetEventConfigListExtraType {
	m := new(SwitchMSwitchEventConfigGetEventConfigListExtraType)
	return m
}

type SwitchMSwitchEventConfigQueryResponse struct {
	// Data
	// Response data message of Public API
	Data *SwitchMSwitchEventConfigQueryResponse `json:"data,omitempty"`

	// Error
	// Response error message of Public API
	Error *SwitchMSwitchEventConfigQueryResponse `json:"error,omitempty"`

	// Extra
	// Extra information of Public API response
	Extra *SwitchMSwitchEventConfigQueryResponse `json:"extra,omitempty"`

	// MetaData
	// Meta-data of Public API response
	MetaData *SwitchMSwitchEventConfigQueryResponse `json:"metaData,omitempty"`

	// Success
	// Response success message of Public API
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchEventConfigQueryResponse() *SwitchMSwitchEventConfigQueryResponse {
	m := new(SwitchMSwitchEventConfigQueryResponse)
	return m
}

// SwitchMSwitchEventConfigQueryResponseDataType
//
// Response data message of Public API
type SwitchMSwitchEventConfigQueryResponseDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchEventConfigQueryResponseDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchEventConfigQueryResponseDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchEventConfigQueryResponseDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchEventConfigQueryResponseDataType() *SwitchMSwitchEventConfigQueryResponseDataType {
	m := new(SwitchMSwitchEventConfigQueryResponseDataType)
	return m
}

// SwitchMSwitchEventConfigQueryResponseErrorType
//
// Response error message of Public API
type SwitchMSwitchEventConfigQueryResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchEventConfigQueryResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchEventConfigQueryResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchEventConfigQueryResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchEventConfigQueryResponseErrorType() *SwitchMSwitchEventConfigQueryResponseErrorType {
	m := new(SwitchMSwitchEventConfigQueryResponseErrorType)
	return m
}

// SwitchMSwitchEventConfigQueryResponseExtraType
//
// Extra information of Public API response
type SwitchMSwitchEventConfigQueryResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchEventConfigQueryResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchEventConfigQueryResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchEventConfigQueryResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchEventConfigQueryResponseExtraType() *SwitchMSwitchEventConfigQueryResponseExtraType {
	m := new(SwitchMSwitchEventConfigQueryResponseExtraType)
	return m
}

// SwitchMSwitchEventConfigQueryResponseMetaDataType
//
// Meta-data of Public API response
type SwitchMSwitchEventConfigQueryResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchEventConfigQueryResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchEventConfigQueryResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchEventConfigQueryResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchEventConfigQueryResponseMetaDataType() *SwitchMSwitchEventConfigQueryResponseMetaDataType {
	m := new(SwitchMSwitchEventConfigQueryResponseMetaDataType)
	return m
}
