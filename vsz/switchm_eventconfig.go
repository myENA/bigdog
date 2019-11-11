package vsz

// API Version: v8_1

import (
	"encoding/json"
)

type SwitchMEventConfigAddEventConfig struct {
	*SwitchMEventConfig
}

type SwitchMEventConfigAddEventConfigResult struct {
	*SwitchMEventConfigQueryResponse
}

type SwitchMEventConfigDeleteEventConfigResult struct {
	*SwitchMEventConfigQueryResponse
}

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
	//    - nullable
	//    - oneof:[Warning,Major,Critical]
	Severity *string `json:"severity,omitempty" validate:"omitempty,oneof=Warning Major Critical"`

	// TextPattern
	// User defined text pattern of each Switch custom event config
	TextPattern *string `json:"textPattern,omitempty"`

	// TimeWindow
	// Detection time woindow of each Switch custom event config
	// Constraints:
	//    - nullable
	//    - oneof:[60,120,240,480,720,1440,2880]
	TimeWindow *int `json:"timeWindow,omitempty" validate:"omitempty,oneof=60 120 240 480 720 1440 2880"`

	// Type
	// Type of each Switch custom event config
	// Constraints:
	//    - nullable
	//    - oneof:[CPU,Memory,TextPattern]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=CPU Memory TextPattern"`
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

	List []*SwitchMEventConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// List of responsed Switch custom event config
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Count of  responsed Switch custom event config
	TotalCount *int `json:"totalCount,omitempty"`
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

type SwitchMEventConfigUpdateEventConfig struct {
	*SwitchMEventConfig
}

type SwitchMEventConfigUpdateEventConfigResult struct {
	*SwitchMEventConfigQueryResponse
}
