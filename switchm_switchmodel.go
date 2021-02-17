package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// SwitchMSwitchModelModelList
//
// Definition: switchModel_modelList
type SwitchMSwitchModelModelList struct {
	// Model
	// Switch Model
	Model *string `json:"model,omitempty"`
}

func NewSwitchMSwitchModelModelList() *SwitchMSwitchModelModelList {
	m := new(SwitchMSwitchModelModelList)
	return m
}

// SwitchMSwitchModelResult
//
// Definition: switchModel_result
type SwitchMSwitchModelResult struct {
	// Extra
	// Extra field
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// The first data index for current reulst
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of remaining data
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchModelModelList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Data Count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Data Count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMSwitchModelResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchModelResult
}

func newSwitchMSwitchModelResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchModelResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchModelResultAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMSwitchModelResult)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMSwitchModelResult() *SwitchMSwitchModelResult {
	m := new(SwitchMSwitchModelResult)
	return m
}
