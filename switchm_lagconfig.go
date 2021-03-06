package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// SwitchMLAGConfigCreate
//
// Definition: lagConfig_create
type SwitchMLAGConfigCreate struct {
	AltoId *string `json:"altoId,omitempty"`

	// GroupId
	// Switch Group Id
	GroupId *string `json:"groupId,omitempty"`

	// Name
	// The Name of LAG Config
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	Ports []string `json:"ports,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// Type
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty"`
}

func NewSwitchMLAGConfigCreate() *SwitchMLAGConfigCreate {
	m := new(SwitchMLAGConfigCreate)
	return m
}

// SwitchMLAGConfigCreateResult
//
// Definition: lagConfig_createResult
type SwitchMLAGConfigCreateResult struct {
	// Id
	// The ID of LAG Config
	Id *string `json:"id,omitempty"`
}

type SwitchMLAGConfigCreateResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMLAGConfigCreateResult
}

func newSwitchMLAGConfigCreateResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMLAGConfigCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMLAGConfigCreateResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMLAGConfigCreateResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMLAGConfigCreateResult() *SwitchMLAGConfigCreateResult {
	m := new(SwitchMLAGConfigCreateResult)
	return m
}

// SwitchMLAGConfig
//
// Definition: lagConfig_lagConfig
type SwitchMLAGConfig struct {
	AltoId *string `json:"altoId,omitempty"`

	// CreatedTime
	// The LAG Types in STATIC or DYNAMIC
	CreatedTime *int `json:"createdTime,omitempty"`

	// GroupId
	// The ID of  Switch Group
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// The ID of LAG Config
	Id *string `json:"id,omitempty"`

	// Name
	// The Name of LAG Config
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	Ports []string `json:"ports,omitempty"`

	// SwitchId
	// The ID of Switch
	SwitchId *string `json:"switchId,omitempty"`

	// Type
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty"`

	// UpdatedTime
	// The LAG Types in STATIC or DYNAMIC
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type SwitchMLAGConfigAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMLAGConfig
}

func newSwitchMLAGConfigAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMLAGConfigAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMLAGConfigAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMLAGConfig)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMLAGConfig() *SwitchMLAGConfig {
	m := new(SwitchMLAGConfig)
	return m
}

// SwitchMLAGConfigList
//
// Definition: lagConfig_list
type SwitchMLAGConfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMLAGConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMLAGConfigListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMLAGConfigList
}

func newSwitchMLAGConfigListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMLAGConfigListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMLAGConfigListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMLAGConfigList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMLAGConfigList() *SwitchMLAGConfigList {
	m := new(SwitchMLAGConfigList)
	return m
}

// SwitchMLAGConfigModify
//
// Definition: lagConfig_modify
type SwitchMLAGConfigModify struct {
	AltoId *string `json:"altoId,omitempty"`

	// Name
	// The Name of LAG Config
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	Ports []string `json:"ports,omitempty"`
}

func NewSwitchMLAGConfigModify() *SwitchMLAGConfigModify {
	m := new(SwitchMLAGConfigModify)
	return m
}
