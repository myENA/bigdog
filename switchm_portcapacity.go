package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type SwitchMPortCapacityService struct {
	apiClient *VSZClient
}

func NewSwitchMPortCapacityService(c *VSZClient) *SwitchMPortCapacityService {
	s := new(SwitchMPortCapacityService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortCapacityService() *SwitchMPortCapacityService {
	return NewSwitchMPortCapacityService(ss.apiClient)
}

// SwitchMPortCapacityCapacities
//
// Definition: portCapacity_capacities
type SwitchMPortCapacityCapacities struct {
	// Capacity
	// Port Speed Capacity
	Capacity *string `json:"capacity,omitempty"`
}

func NewSwitchMPortCapacityCapacities() *SwitchMPortCapacityCapacities {
	m := new(SwitchMPortCapacityCapacities)
	return m
}

// SwitchMPortCapacityResult
//
// Definition: portCapacity_result
type SwitchMPortCapacityResult struct {
	// Extra
	// Extra field
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMPortCapacityResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMPortCapacityResult
}

func newSwitchMPortCapacityResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMPortCapacityResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMPortCapacityResultAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMPortCapacityResult)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMPortCapacityResult() *SwitchMPortCapacityResult {
	m := new(SwitchMPortCapacityResult)
	return m
}

// FindPortCapacity
//
// Use this API command to Retrieve Switch Port Capacity List.
//
// Operation ID: findPortCapacity
// Operation path: /portCapacity
// Success code: 200 (OK)
//
// Required parameters:
// - model string
//		- required
// - portIdentifier string
//		- required
func (s *SwitchMPortCapacityService) FindPortCapacity(ctx context.Context, model string, portIdentifier string, mutators ...RequestMutator) (*SwitchMPortCapacityResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMPortCapacityResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindPortCapacity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("model", model)
	req.QueryParams.Set("portIdentifier", portIdentifier)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMPortCapacityResultAPIResponse), err
}
