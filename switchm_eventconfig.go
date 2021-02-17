package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// SwitchMEventConfig
//
// Definition: eventConfig_eventConfig
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

type SwitchMEventConfigAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMEventConfig
}

func newSwitchMEventConfigAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMEventConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMEventConfigAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMEventConfig)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMEventConfig() *SwitchMEventConfig {
	m := new(SwitchMEventConfig)
	return m
}

// SwitchMEventConfigGetEventConfigList
//
// Definition: eventConfig_getEventConfigList
type SwitchMEventConfigGetEventConfigList struct {
	// Extra
	// Extra information of responsed Switch custom event config list
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMEventConfigGetEventConfigListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMEventConfigGetEventConfigList
}

func newSwitchMEventConfigGetEventConfigListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMEventConfigGetEventConfigListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMEventConfigGetEventConfigListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMEventConfigGetEventConfigList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMEventConfigGetEventConfigList() *SwitchMEventConfigGetEventConfigList {
	m := new(SwitchMEventConfigGetEventConfigList)
	return m
}

// SwitchMEventConfigQueryResponse
//
// Definition: eventConfig_queryResponse
type SwitchMEventConfigQueryResponse struct {
	// Data
	// Response data message of Public API
	Data interface{} `json:"data,omitempty"`

	// Error
	// Response error message of Public API
	Error interface{} `json:"error,omitempty"`

	// Extra
	// Extra information of Public API response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Meta-data of Public API response
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Response success message of Public API
	Success *bool `json:"success,omitempty"`
}

type SwitchMEventConfigQueryResponseAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMEventConfigQueryResponse
}

func newSwitchMEventConfigQueryResponseAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMEventConfigQueryResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMEventConfigQueryResponseAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMEventConfigQueryResponse)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMEventConfigQueryResponse() *SwitchMEventConfigQueryResponse {
	m := new(SwitchMEventConfigQueryResponse)
	return m
}
