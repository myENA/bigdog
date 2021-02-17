package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGHistoricalClientConnectionDiagnosticService struct {
	apiClient *VSZClient
}

func NewWSGHistoricalClientConnectionDiagnosticService(c *VSZClient) *WSGHistoricalClientConnectionDiagnosticService {
	s := new(WSGHistoricalClientConnectionDiagnosticService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHistoricalClientConnectionDiagnosticService() *WSGHistoricalClientConnectionDiagnosticService {
	return NewWSGHistoricalClientConnectionDiagnosticService(ss.apiClient)
}

// WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry
//
// Definition: historicalClientConnectionDiagnostic_clientConnectionFailureTypeCountEntry
type WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry struct {
	DisplayName *string `json:"displayName,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry() *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry {
	m := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry)
	return m
}

// WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList
//
// Definition: historicalClientConnectionDiagnostic_clientConnectionFailureTypeCountList
type WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry `json:"list,omitempty"`

	// RawDataTotalCount
	// No idea.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse struct {
	*RawAPIResponse
	Data *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList
}

func newWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList() *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList {
	m := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList)
	return m
}

// HccdCount
//
// Retrieve total count of of specified type of diagnostic entries
//
// Operation ID: hccdCount
// Operation path: /hccd/hccdClientConnection/{type}/count
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - type_ string
//		- required
//		- oneof:[failureType]
func (s *WSGHistoricalClientConnectionDiagnosticService) HccdCount(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGHccdCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("type", type_)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
}

// HccdTypeCount
//
// Operation ID: hccdTypeCount
// Operation path: /hccd/hccdClientConnection/{type}/failedMsgId/count
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - type_ string
//		- required
//		- oneof:[association,authentication,eap,radius,dhcp,userAuthentication]
func (s *WSGHistoricalClientConnectionDiagnosticService) HccdTypeCount(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGHccdTypeCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("type", type_)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
}
