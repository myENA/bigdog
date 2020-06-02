package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMPortcapacityService struct {
	apiClient *APIClient
}

func NewSwitchMPortcapacityService(c *APIClient) *SwitchMPortcapacityService {
	s := new(SwitchMPortcapacityService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortcapacityService() *SwitchMPortcapacityService {
	return NewSwitchMPortcapacityService(ss.apiClient)
}

type SwitchMPortcapacityCapacities struct {
	// Capacity
	// Port Speed Capacity
	Capacity *string `json:"capacity,omitempty"`
}

func NewSwitchMPortcapacityCapacities() *SwitchMPortcapacityCapacities {
	m := new(SwitchMPortcapacityCapacities)
	return m
}

type SwitchMPortcapacityResult struct {
	// Extra
	// Extra field
	Extra *SwitchMPortcapacityResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// The first data index for current reulst
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of remaining data
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMPortcapacityCapacities `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Data Count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Data Count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMPortcapacityResult() *SwitchMPortcapacityResult {
	m := new(SwitchMPortcapacityResult)
	return m
}

// SwitchMPortcapacityResultExtraType
//
// Extra field
type SwitchMPortcapacityResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMPortcapacityResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMPortcapacityResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMPortcapacityResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMPortcapacityResultExtraType() *SwitchMPortcapacityResultExtraType {
	m := new(SwitchMPortcapacityResultExtraType)
	return m
}

// FindPortCapacity
//
// Use this API command to Retrieve Switch Port Capacity List.
func (s *SwitchMPortcapacityService) FindPortCapacity(ctx context.Context) (*SwitchMPortcapacityResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMPortcapacityResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindPortCapacity, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMPortcapacityResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
