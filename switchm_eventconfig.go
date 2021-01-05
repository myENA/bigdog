package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
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

func newSwitchMEventConfigAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMEventConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMEventConfigAPIResponse) Hydrate() error {
	r.Data = new(SwitchMEventConfig)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMEventConfigGetEventConfigListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMEventConfigGetEventConfigListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMEventConfigGetEventConfigListAPIResponse) Hydrate() error {
	r.Data = new(SwitchMEventConfigGetEventConfigList)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMEventConfigQueryResponseAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMEventConfigQueryResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMEventConfigQueryResponseAPIResponse) Hydrate() error {
	r.Data = new(SwitchMEventConfigQueryResponse)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMEventConfigQueryResponse() *SwitchMEventConfigQueryResponse {
	m := new(SwitchMEventConfigQueryResponse)
	return m
}
