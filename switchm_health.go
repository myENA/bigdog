package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type SwitchMHealthService struct {
	apiClient *VSZClient
}

func NewSwitchMHealthService(c *VSZClient) *SwitchMHealthService {
	s := new(SwitchMHealthService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMHealthService() *SwitchMHealthService {
	return NewSwitchMHealthService(ss.apiClient)
}

// SwitchMHealthAggMetrics
//
// Definition: health_aggMetrics
type SwitchMHealthAggMetrics struct {
	// Extra
	// Extra information for Aggregation Metrics
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Aggregation Metrics returned out of the complete ICX Metrics list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more Aggregation Metrics after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMHealthAggs `json:"list,omitempty"`

	// RawDataTotalCount
	// Aggregation Metrics count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of Aggregation Metrics count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMHealthAggMetricsAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMHealthAggMetrics
}

func newSwitchMHealthAggMetricsAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMHealthAggMetricsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMHealthAggMetricsAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMHealthAggMetrics)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMHealthAggMetrics() *SwitchMHealthAggMetrics {
	m := new(SwitchMHealthAggMetrics)
	return m
}

// SwitchMHealthAggs
//
// Definition: health_aggs
type SwitchMHealthAggs struct {
	// Id
	// Identifier of the aggregation value
	Id *string `json:"id,omitempty"`

	// Key
	// Key of the aggregation value
	Key *string `json:"key,omitempty"`

	// Value
	// Metrics of the aggregation value
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMHealthAggs() *SwitchMHealthAggs {
	m := new(SwitchMHealthAggs)
	return m
}

// SwitchMHealthFan
//
// Definition: health_fan
type SwitchMHealthFan struct {
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Fan slot number
	SlotNumber *int `json:"slotNumber,omitempty"`

	// Status
	// Fan status
	Status *string `json:"status,omitempty"`

	// Type
	// Fan type
	Type *string `json:"type,omitempty"`
}

func NewSwitchMHealthFan() *SwitchMHealthFan {
	m := new(SwitchMHealthFan)
	return m
}

// SwitchMHealthIcxMetrics
//
// Definition: health_icxMetrics
type SwitchMHealthIcxMetrics struct {
	// Extra
	// Extra information for ICX Metrics
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ICX Metrics returned out of the complete ICX Metrics list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more ICX Metrics after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMHealthMetrics `json:"list,omitempty"`

	// RawDataTotalCount
	// ICX Metrics count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of ICX Metrics count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMHealthIcxMetricsAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMHealthIcxMetrics
}

func newSwitchMHealthIcxMetricsAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMHealthIcxMetricsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMHealthIcxMetricsAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMHealthIcxMetrics)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMHealthIcxMetrics() *SwitchMHealthIcxMetrics {
	m := new(SwitchMHealthIcxMetrics)
	return m
}

// SwitchMHealthMetrics
//
// Definition: health_metrics
type SwitchMHealthMetrics struct {
	// Avg
	// Average metrics.  Value may be either float or "NaN"
	Avg interface{} `json:"avg,omitempty"`

	// Max
	// Max metrics.  Value may be either float or "-Infinity"
	Max interface{} `json:"max,omitempty"`

	// Min
	// Min metrics.  value may be either float or "Infinity"
	Min interface{} `json:"min,omitempty"`

	// Timestamp
	// Timestamp
	Timestamp *string `json:"timestamp,omitempty"`
}

func NewSwitchMHealthMetrics() *SwitchMHealthMetrics {
	m := new(SwitchMHealthMetrics)
	return m
}

// SwitchMHealthPowerSupply
//
// Definition: health_powerSupply
type SwitchMHealthPowerSupply struct {
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Power supply slot number
	SlotNumber *int `json:"slotNumber,omitempty"`

	// Status
	// Power supply status
	Status *string `json:"status,omitempty"`

	// Type
	// Power supply type
	Type *string `json:"type,omitempty"`
}

func NewSwitchMHealthPowerSupply() *SwitchMHealthPowerSupply {
	m := new(SwitchMHealthPowerSupply)
	return m
}

// SwitchMHealthStatus
//
// Definition: health_status
type SwitchMHealthStatus struct {
	// Fan
	// Fan
	Fan []*SwitchMHealthStatusFanType `json:"fan,omitempty"`

	// FlaggedCount
	// Flagged status count
	FlaggedCount *int `json:"flaggedCount,omitempty"`

	// OfflineCount
	// Offline status count
	OfflineCount *int `json:"offlineCount,omitempty"`

	// OnlineCount
	// Online status count
	OnlineCount *int `json:"onlineCount,omitempty"`

	// PowerHealthStatus
	// Health status for ICX/Stack that contains power supply, temperature and fan
	PowerHealthStatus []*SwitchMHealthStatusPowerHealthStatusType `json:"powerHealthStatus,omitempty"`

	// PowerSupply
	// Powersupply
	PowerSupply []*SwitchMHealthStatusPowerSupplyType `json:"powerSupply,omitempty"`

	// Temperature
	// Temperature
	Temperature []*SwitchMHealthStatusTemperatureType `json:"temperature,omitempty"`
}

type SwitchMHealthStatusAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMHealthStatus
}

func newSwitchMHealthStatusAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMHealthStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMHealthStatusAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMHealthStatus)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMHealthStatus() *SwitchMHealthStatus {
	m := new(SwitchMHealthStatus)
	return m
}

// SwitchMHealthStatusFanType
//
// Definition: health_statusFanType
type SwitchMHealthStatusFanType struct {
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Fan slot number
	SlotNumber *int `json:"slotNumber,omitempty"`

	// Status
	// Fan status
	Status *string `json:"status,omitempty"`

	// Type
	// Fan type
	Type *string `json:"type,omitempty"`
}

func NewSwitchMHealthStatusFanType() *SwitchMHealthStatusFanType {
	m := new(SwitchMHealthStatusFanType)
	return m
}

// SwitchMHealthStatusPowerHealthStatusType
//
// Definition: health_statusPowerHealthStatusType
type SwitchMHealthStatusPowerHealthStatusType struct {
	Fan []*SwitchMHealthFan `json:"fan,omitempty"`

	PowerSupply []*SwitchMHealthPowerSupply `json:"powerSupply,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty"`

	Temperature []*SwitchMHealthTemperature `json:"temperature,omitempty"`
}

func NewSwitchMHealthStatusPowerHealthStatusType() *SwitchMHealthStatusPowerHealthStatusType {
	m := new(SwitchMHealthStatusPowerHealthStatusType)
	return m
}

// SwitchMHealthStatusPowerSupplyType
//
// Definition: health_statusPowerSupplyType
type SwitchMHealthStatusPowerSupplyType struct {
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Power supply slot number
	SlotNumber *int `json:"slotNumber,omitempty"`

	// Status
	// Power supply status
	Status *string `json:"status,omitempty"`

	// Type
	// Power supply type
	Type *string `json:"type,omitempty"`
}

func NewSwitchMHealthStatusPowerSupplyType() *SwitchMHealthStatusPowerSupplyType {
	m := new(SwitchMHealthStatusPowerSupplyType)
	return m
}

// SwitchMHealthStatusTemperatureType
//
// Definition: health_statusTemperatureType
type SwitchMHealthStatusTemperatureType struct {
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Solt number
	SlotNumber *int `json:"slotNumber,omitempty"`

	// TemperatureValue
	// Slot temperature
	TemperatureValue *float64 `json:"temperatureValue,omitempty"`
}

func NewSwitchMHealthStatusTemperatureType() *SwitchMHealthStatusTemperatureType {
	m := new(SwitchMHealthStatusTemperatureType)
	return m
}

// SwitchMHealthTemperature
//
// Definition: health_temperature
type SwitchMHealthTemperature struct {
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Solt number
	SlotNumber *int `json:"slotNumber,omitempty"`

	// TemperatureValue
	// Slot temperature
	TemperatureValue *float64 `json:"temperatureValue,omitempty"`
}

func NewSwitchMHealthTemperature() *SwitchMHealthTemperature {
	m := new(SwitchMHealthTemperature)
	return m
}

// AddHealthCpuAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Operation ID: addHealthCpuAgg
// Operation path: /health/cpu/agg
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthCpuAgg(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMHealthAggMetricsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMHealthAggMetricsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddHealthCpuAgg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMHealthAggMetricsAPIResponse), err
}

// AddHealthCpuLine
//
// Use this API command to retrieve CPU trend data based on the time duration.
//
// Operation ID: addHealthCpuLine
// Operation path: /health/cpu/line
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthCpuLine(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMHealthIcxMetricsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMHealthIcxMetricsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddHealthCpuLine, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMHealthIcxMetricsAPIResponse), err
}

// AddHealthMemAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Operation ID: addHealthMemAgg
// Operation path: /health/mem/agg
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthMemAgg(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMHealthAggMetricsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMHealthAggMetricsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddHealthMemAgg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMHealthAggMetricsAPIResponse), err
}

// AddHealthMemLine
//
// Use this API command to retrieve switch memory trend data based on the time duration.
//
// Operation ID: addHealthMemLine
// Operation path: /health/mem/line
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthMemLine(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMHealthIcxMetricsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMHealthIcxMetricsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddHealthMemLine, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMHealthIcxMetricsAPIResponse), err
}

// AddHealthStatus
//
// Use this API command to retrieve switch health status.
//
// Operation ID: addHealthStatus
// Operation path: /health/status
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthStatus(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMHealthStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMHealthStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddHealthStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMHealthStatusAPIResponse), err
}

// AddHealthStatusAll
//
// Use this API command to retrieve fan, temperature and power supply status for the switch managed by SmartZone.
//
// Operation ID: addHealthStatusAll
// Operation path: /health/status/all
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthStatusAll(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMHealthStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMHealthStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddHealthStatusAll, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMHealthStatusAPIResponse), err
}

// AddHealthStatusBySerialNumber
//
// Use this API command to retrieve fan, temperature and power supply status for the switch unit managed by SmartZone.
//
// Operation ID: addHealthStatusBySerialNumber
// Operation path: /health/status/{serialNumber}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
//
// Required parameters:
// - serialNumber string
//		- required
func (s *SwitchMHealthService) AddHealthStatusBySerialNumber(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, serialNumber string, mutators ...RequestMutator) (*SwitchMHealthStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMHealthStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddHealthStatusBySerialNumber, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("serialNumber", serialNumber)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMHealthStatusAPIResponse), err
}
