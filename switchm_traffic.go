package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type SwitchMTrafficService struct {
	apiClient *VSZClient
}

func NewSwitchMTrafficService(c *VSZClient) *SwitchMTrafficService {
	s := new(SwitchMTrafficService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTrafficService() *SwitchMTrafficService {
	return NewSwitchMTrafficService(ss.apiClient)
}

// SwitchMTrafficBroadcast
//
// Definition: traffic_broadcast
type SwitchMTrafficBroadcast struct {
	BroadcastIn *string `json:"broadcastIn,omitempty"`

	BroadcastOut *string `json:"broadcastOut,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	Total *int `json:"total,omitempty"`
}

func NewSwitchMTrafficBroadcast() *SwitchMTrafficBroadcast {
	m := new(SwitchMTrafficBroadcast)
	return m
}

// SwitchMTrafficBroadcastQueryResultList
//
// Definition: traffic_broadcastQueryResultList
type SwitchMTrafficBroadcastQueryResultList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficBroadcast `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficBroadcastQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficBroadcastQueryResultList
}

func newSwitchMTrafficBroadcastQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficBroadcastQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficBroadcastQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficBroadcastQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficBroadcastQueryResultList() *SwitchMTrafficBroadcastQueryResultList {
	m := new(SwitchMTrafficBroadcastQueryResultList)
	return m
}

// SwitchMTrafficMulticast
//
// Definition: traffic_multicast
type SwitchMTrafficMulticast struct {
	MulticastIn *string `json:"multicastIn,omitempty"`

	MulticastOut *string `json:"multicastOut,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	Total *int `json:"total,omitempty"`
}

func NewSwitchMTrafficMulticast() *SwitchMTrafficMulticast {
	m := new(SwitchMTrafficMulticast)
	return m
}

// SwitchMTrafficMulticastQueryResultList
//
// Definition: traffic_multicastQueryResultList
type SwitchMTrafficMulticastQueryResultList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficMulticast `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficMulticastQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficMulticastQueryResultList
}

func newSwitchMTrafficMulticastQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficMulticastQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficMulticastQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficMulticastQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficMulticastQueryResultList() *SwitchMTrafficMulticastQueryResultList {
	m := new(SwitchMTrafficMulticastQueryResultList)
	return m
}

// SwitchMTrafficPortError
//
// Definition: traffic_portError
type SwitchMTrafficPortError struct {
	CrcErr *string `json:"crcErr,omitempty"`

	InErr *string `json:"inErr,omitempty"`

	OutErr *string `json:"outErr,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	Total *int `json:"total,omitempty"`
}

func NewSwitchMTrafficPortError() *SwitchMTrafficPortError {
	m := new(SwitchMTrafficPortError)
	return m
}

// SwitchMTrafficPortErrorQueryResultList
//
// Definition: traffic_portErrorQueryResultList
type SwitchMTrafficPortErrorQueryResultList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficPortError `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficPortErrorQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficPortErrorQueryResultList
}

func newSwitchMTrafficPortErrorQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficPortErrorQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficPortErrorQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficPortErrorQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficPortErrorQueryResultList() *SwitchMTrafficPortErrorQueryResultList {
	m := new(SwitchMTrafficPortErrorQueryResultList)
	return m
}

// SwitchMTrafficTopPortErrorQueryResultList
//
// Definition: traffic_topPortErrorQueryResultList
type SwitchMTrafficTopPortErrorQueryResultList struct {
	// Extra
	// Extra information for top port error
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port error returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port error after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port error count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port error count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficTopPortErrorQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficTopPortErrorQueryResultList
}

func newSwitchMTrafficTopPortErrorQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficTopPortErrorQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficTopPortErrorQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficTopPortErrorQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficTopPortErrorQueryResultList() *SwitchMTrafficTopPortErrorQueryResultList {
	m := new(SwitchMTrafficTopPortErrorQueryResultList)
	return m
}

// SwitchMTrafficTopPortTrafficUsageQueryResultList
//
// Definition: traffic_topPortTrafficUsageQueryResultList
type SwitchMTrafficTopPortTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top port traffic usage
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficTopPortTrafficUsageQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficTopPortTrafficUsageQueryResultList
}

func newSwitchMTrafficTopPortTrafficUsageQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficTopPortTrafficUsageQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficTopPortTrafficUsageQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficTopPortTrafficUsageQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficTopPortTrafficUsageQueryResultList() *SwitchMTrafficTopPortTrafficUsageQueryResultList {
	m := new(SwitchMTrafficTopPortTrafficUsageQueryResultList)
	return m
}

// SwitchMTrafficTopSwitchPoEUtilizationQueryResultList
//
// Definition: traffic_topSwitchPoEUtilizationQueryResultList
type SwitchMTrafficTopSwitchPoEUtilizationQueryResultList struct {
	// Extra
	// Extra information for top PoE utilization
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top PoE usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top PoE usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// PoE utilization count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total PoE utilization count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficTopSwitchPoEUtilizationQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficTopSwitchPoEUtilizationQueryResultList
}

func newSwitchMTrafficTopSwitchPoEUtilizationQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficTopSwitchPoEUtilizationQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficTopSwitchPoEUtilizationQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficTopSwitchPoEUtilizationQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficTopSwitchPoEUtilizationQueryResultList() *SwitchMTrafficTopSwitchPoEUtilizationQueryResultList {
	m := new(SwitchMTrafficTopSwitchPoEUtilizationQueryResultList)
	return m
}

// SwitchMTrafficTopTrafficUsageQueryResultList
//
// Definition: traffic_topTrafficUsageQueryResultList
type SwitchMTrafficTopTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top traffic usage
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficTopTrafficUsageQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficTopTrafficUsageQueryResultList
}

func newSwitchMTrafficTopTrafficUsageQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficTopTrafficUsageQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficTopTrafficUsageQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficTopTrafficUsageQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficTopTrafficUsageQueryResultList() *SwitchMTrafficTopTrafficUsageQueryResultList {
	m := new(SwitchMTrafficTopTrafficUsageQueryResultList)
	return m
}

// SwitchMTraffic
//
// Definition: traffic_traffic
type SwitchMTraffic struct {
	// Rx
	// RX traffic of the switch
	Rx interface{} `json:"rx,omitempty"`

	// Timestamp
	// Timestamp of the switch traffic
	Timestamp *string `json:"timestamp,omitempty"`

	// Total
	// Total traffic of the switch
	Total interface{} `json:"total,omitempty"`

	// Tx
	// TX traffic of the switch
	Tx interface{} `json:"tx,omitempty"`
}

func NewSwitchMTraffic() *SwitchMTraffic {
	m := new(SwitchMTraffic)
	return m
}

// SwitchMTrafficQueryResultList
//
// Definition: traffic_trafficQueryResultList
type SwitchMTrafficQueryResultList struct {
	// Extra
	// Extra information for traffic list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first traffic list returned out of the complete traffic list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more traffic list after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTraffic `json:"list,omitempty"`

	// RawDataTotalCount
	// Total traffic count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of traffic list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficQueryResultList
}

func newSwitchMTrafficQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficQueryResultList() *SwitchMTrafficQueryResultList {
	m := new(SwitchMTrafficQueryResultList)
	return m
}

// SwitchMTrafficUsage
//
// Definition: traffic_trafficUsage
type SwitchMTrafficUsage struct {
	// Id
	// Identifier of the Traffic Usage
	Id *string `json:"id,omitempty"`

	// Key
	// Interface of the Traffic Usage
	Key *string `json:"key,omitempty"`

	// Value
	// Total Tx and Rx of Traffic Usage
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMTrafficUsage() *SwitchMTrafficUsage {
	m := new(SwitchMTrafficUsage)
	return m
}

// SwitchMTrafficUnicast
//
// Definition: traffic_unicast
type SwitchMTrafficUnicast struct {
	Timestamp *string `json:"timestamp,omitempty"`

	Total *int `json:"total,omitempty"`

	UnicastIn *string `json:"unicastIn,omitempty"`

	UnicastOut *string `json:"unicastOut,omitempty"`
}

func NewSwitchMTrafficUnicast() *SwitchMTrafficUnicast {
	m := new(SwitchMTrafficUnicast)
	return m
}

// SwitchMTrafficUnicastQueryResultList
//
// Definition: traffic_unicastQueryResultList
type SwitchMTrafficUnicastQueryResultList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUnicast `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMTrafficUnicastQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTrafficUnicastQueryResultList
}

func newSwitchMTrafficUnicastQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTrafficUnicastQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTrafficUnicastQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTrafficUnicastQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTrafficUnicastQueryResultList() *SwitchMTrafficUnicastQueryResultList {
	m := new(SwitchMTrafficUnicastQueryResultList)
	return m
}

// AddTrafficTopPoeutilization
//
// Use this API command retrieve the top 10 switches by the PoE utilization.
//
// Operation ID: addTrafficTopPoeutilization
// Operation path: /traffic/top/poeutilization
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPoeutilization(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopSwitchPoEUtilizationQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficTopSwitchPoEUtilizationQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTopPoeutilization, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficTopSwitchPoEUtilizationQueryResultListAPIResponse), err
}

// AddTrafficTopPorterror
//
// Use this API command to get the top 10 switches by the porterror.
//
// Operation ID: addTrafficTopPorterror
// Operation path: /traffic/top/porterror
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPorterror(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopPortErrorQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficTopPortErrorQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTopPorterror, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficTopPortErrorQueryResultListAPIResponse), err
}

// AddTrafficTopPortusage
//
// Use this API command to get the top 10 ports by the traffic.
//
// Operation ID: addTrafficTopPortusage
// Operation path: /traffic/top/portusage
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPortusage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopPortTrafficUsageQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficTopPortTrafficUsageQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTopPortusage, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficTopPortTrafficUsageQueryResultListAPIResponse), err
}

// AddTrafficTopUsage
//
// Use this API command to retrieve Top Swich/Port usage data based on the time duration.
//
// Operation ID: addTrafficTopUsage
// Operation path: /traffic/top/usage
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopUsage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopTrafficUsageQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficTopTrafficUsageQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTopUsage, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficTopTrafficUsageQueryResultListAPIResponse), err
}

// AddTrafficTotalBroadcast
//
// Use this API command to retrieve Switch/Port Broadcast trend data based on the time duration.
//
// Operation ID: addTrafficTotalBroadcast
// Operation path: /traffic/total/broadcast
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTotalBroadcast(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficBroadcastQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficBroadcastQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTotalBroadcast, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficBroadcastQueryResultListAPIResponse), err
}

// AddTrafficTotalMulticast
//
// Use this API command to retrieve Switch/Port Multicast trend data based on the time duration.
//
// Operation ID: addTrafficTotalMulticast
// Operation path: /traffic/total/multicast
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTotalMulticast(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficMulticastQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficMulticastQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTotalMulticast, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficMulticastQueryResultListAPIResponse), err
}

// AddTrafficTotalPorterror
//
// Use this API command to retrieve Switch/Port Port Errors data based on the time duration.
//
// Operation ID: addTrafficTotalPorterror
// Operation path: /traffic/total/porterror
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTotalPorterror(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficPortErrorQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficPortErrorQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTotalPorterror, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficPortErrorQueryResultListAPIResponse), err
}

// AddTrafficTotalTrend
//
// Use this API command to retrieve Swich/Port trend data based on the time duration.
//
// Operation ID: addTrafficTotalTrend
// Operation path: /traffic/total/trend
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTotalTrend(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTotalTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficQueryResultListAPIResponse), err
}

// AddTrafficTotalUnicast
//
// Use this API command to retrieve Switch/Port Unicast trend data based on the time duration.
//
// Operation ID: addTrafficTotalUnicast
// Operation path: /traffic/total/unicast
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTotalUnicast(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficUnicastQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTrafficUnicastQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddTrafficTotalUnicast, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTrafficUnicastQueryResultListAPIResponse), err
}
