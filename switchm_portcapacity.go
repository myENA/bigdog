package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchPortCapacityService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchPortCapacityService(c *VSZClient) *SwitchMSwitchPortCapacityService {
	s := new(SwitchMSwitchPortCapacityService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortCapacityService() *SwitchMSwitchPortCapacityService {
	return NewSwitchMSwitchPortCapacityService(ss.apiClient)
}

type SwitchMSwitchPortCapacityCapacities struct {
	// Capacity
	// Port Speed Capacity
	Capacity *string `json:"capacity,omitempty"`
}

func NewSwitchMSwitchPortCapacityCapacities() *SwitchMSwitchPortCapacityCapacities {
	m := new(SwitchMSwitchPortCapacityCapacities)
	return m
}

type SwitchMSwitchPortCapacityResult struct {
	// Extra
	// Extra field
	Extra **SwitchMSwitchPortCapacityResult `json:"extra,omitempty"`

	// FirstIndex
	// The first data index for current reulst
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of remaining data
	HasMore *bool `json:"hasMore,omitempty"`

	List []**SwitchMSwitchPortCapacityResult `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Data Count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Data Count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchPortCapacityResult() *SwitchMSwitchPortCapacityResult {
	m := new(SwitchMSwitchPortCapacityResult)
	return m
}

// SwitchMSwitchPortCapacityResultExtraType
//
// Extra field
type SwitchMSwitchPortCapacityResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchPortCapacityResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchPortCapacityResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchPortCapacityResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchPortCapacityResultExtraType() *SwitchMSwitchPortCapacityResultExtraType {
	m := new(SwitchMSwitchPortCapacityResultExtraType)
	return m
}

// FindPortCapacity
//
// Use this API command to Retrieve Switch Port Capacity List.
func (s *SwitchMSwitchPortCapacityService) FindPortCapacity(ctx context.Context) (*SwitchMSwitchPortCapacityResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchPortCapacityResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortCapacity, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchPortCapacityResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
