package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMPortCapacityService struct {
	apiClient *APIClient
}

func NewSwitchMPortCapacityService(c *APIClient) *SwitchMPortCapacityService {
	s := new(SwitchMPortCapacityService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortCapacityService() *SwitchMPortCapacityService {
	return NewSwitchMPortCapacityService(ss.apiClient)
}

type SwitchMPortCapacityCapacities struct {
	// Capacity
	// Port Speed Capacity
	Capacity *string `json:"capacity,omitempty"`
}

func NewSwitchMPortCapacityCapacities() *SwitchMPortCapacityCapacities {
	m := new(SwitchMPortCapacityCapacities)
	return m
}

type SwitchMPortCapacityResult struct {
	// Extra
	// Extra field
	Extra *SwitchMPortCapacityResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// The first data index for current reulst
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of remaining data
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMPortCapacityCapacities `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Data Count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Data Count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMPortCapacityResult() *SwitchMPortCapacityResult {
	m := new(SwitchMPortCapacityResult)
	return m
}

// SwitchMPortCapacityResultExtraType
//
// Extra field
type SwitchMPortCapacityResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMPortCapacityResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMPortCapacityResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMPortCapacityResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMPortCapacityResultExtraType() *SwitchMPortCapacityResultExtraType {
	m := new(SwitchMPortCapacityResultExtraType)
	return m
}

// FindPortCapacity
//
// Use this API command to Retrieve Switch Port Capacity List.
func (s *SwitchMPortCapacityService) FindPortCapacity(ctx context.Context) (*SwitchMPortCapacityResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortCapacityResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortCapacity, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMPortCapacityResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
