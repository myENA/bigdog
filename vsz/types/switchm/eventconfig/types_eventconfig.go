package eventconfig

// API Version: v8_1

import (
	"encoding/json"
)

type AddEventConfig struct {
	*EventConfig
}

type AddEventConfigResult struct {
	*QueryResponse
}

type DeleteEventConfigResult struct {
	*QueryResponse
}

type EventConfig struct {
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
	//    - oneof:[Warning,Major,Critical]
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
	//    - oneof:[60,120,240,480,720,1440,2880]
	//    - oneof:[60,120,240,480,720,1440,2880]
	TimeWindow *int `json:"timeWindow,omitempty" validate:"omitempty,oneof=60 120 240 480 720 1440 2880"`

	// Type
	// Type of each Switch custom event config
	// Constraints:
	//    - nullable
	//    - oneof:[CPU,Memory,TextPattern]
	//    - oneof:[CPU,Memory,TextPattern]
	//    - oneof:[CPU,Memory,TextPattern]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=CPU Memory TextPattern"`
}

type GetEventConfigList struct {
	// Extra
	// Extra information of responsed Switch custom event config list
	Extra *GetEventConfigListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// First index of responsed Switch custom event config list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Switch event config
	HasMore *bool `json:"hasMore,omitempty"`

	List []*EventConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// List of responsed Switch custom event config
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Count of  responsed Switch custom event config
	TotalCount *int `json:"totalCount,omitempty"`
}

// GetEventConfigListExtraType
//
// Extra information of responsed Switch custom event config list
type GetEventConfigListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *GetEventConfigListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = GetEventConfigListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *GetEventConfigListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type QueryResponse struct {
	// Data
	// Response data message of Public API
	Data *QueryResponseDataType `json:"data,omitempty"`

	// Error
	// Response error message of Public API
	Error *QueryResponseErrorType `json:"error,omitempty"`

	// Extra
	// Extra information of Public API response
	Extra *QueryResponseExtraType `json:"extra,omitempty"`

	// MetaData
	// Meta-data of Public API response
	MetaData *QueryResponseMetaDataType `json:"metaData,omitempty"`

	// Success
	// Response success message of Public API
	Success *bool `json:"success,omitempty"`
}

// QueryResponseDataType
//
// Response data message of Public API
type QueryResponseDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *QueryResponseDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = QueryResponseDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *QueryResponseDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

// QueryResponseErrorType
//
// Response error message of Public API
type QueryResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *QueryResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = QueryResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *QueryResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

// QueryResponseExtraType
//
// Extra information of Public API response
type QueryResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *QueryResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = QueryResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *QueryResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

// QueryResponseMetaDataType
//
// Meta-data of Public API response
type QueryResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *QueryResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = QueryResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *QueryResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type UpdateEventConfig struct {
	*EventConfig
}

type UpdateEventConfigResult struct {
	*QueryResponse
}
