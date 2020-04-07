package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMHealthService struct {
	apiClient *APIClient
}

func NewSwitchMHealthService(c *APIClient) *SwitchMHealthService {
	s := new(SwitchMHealthService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMHealthService() *SwitchMHealthService {
	return NewSwitchMHealthService(ss.apiClient)
}

type SwitchMHealthAggMetrics struct {
	// Extra
	// Extra information for Aggregation Metrics
	// Constraints:
	//    - nullable
	Extra *SwitchMHealthAggMetricsExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Aggregation Metrics returned out of the complete ICX Metrics list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more Aggregation Metrics after the currently displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMHealthAggs `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Aggregation Metrics count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of Aggregation Metrics count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMHealthAggMetrics() *SwitchMHealthAggMetrics {
	m := new(SwitchMHealthAggMetrics)
	return m
}

// SwitchMHealthAggMetricsExtraType
//
// Extra information for Aggregation Metrics
// Constraints:
//    - nullable
type SwitchMHealthAggMetricsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMHealthAggMetricsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMHealthAggMetricsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMHealthAggMetricsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMHealthAggMetricsExtraType() *SwitchMHealthAggMetricsExtraType {
	m := new(SwitchMHealthAggMetricsExtraType)
	return m
}

type SwitchMHealthAggs struct {
	// Id
	// Identifier of the aggregation value
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Key
	// Key of the aggregation value
	// Constraints:
	//    - nullable
	Key *string `json:"key,omitempty"`

	// Value
	// Metrics of the aggregation value
	// Constraints:
	//    - nullable
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMHealthAggs() *SwitchMHealthAggs {
	m := new(SwitchMHealthAggs)
	return m
}

type SwitchMHealthIcxMetrics struct {
	// Extra
	// Extra information for ICX Metrics
	// Constraints:
	//    - nullable
	Extra *SwitchMHealthIcxMetricsExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ICX Metrics returned out of the complete ICX Metrics list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more ICX Metrics after the currently displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMHealthMetrics `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// ICX Metrics count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of ICX Metrics count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMHealthIcxMetrics() *SwitchMHealthIcxMetrics {
	m := new(SwitchMHealthIcxMetrics)
	return m
}

// SwitchMHealthIcxMetricsExtraType
//
// Extra information for ICX Metrics
// Constraints:
//    - nullable
type SwitchMHealthIcxMetricsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMHealthIcxMetricsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMHealthIcxMetricsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMHealthIcxMetricsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMHealthIcxMetricsExtraType() *SwitchMHealthIcxMetricsExtraType {
	m := new(SwitchMHealthIcxMetricsExtraType)
	return m
}

type SwitchMHealthMetrics struct {
	// Avg
	// Average metrics
	// Constraints:
	//    - nullable
	Avg *float64 `json:"avg,omitempty"`

	// Max
	// Max metrics
	// Constraints:
	//    - nullable
	Max *float64 `json:"max,omitempty"`

	// Min
	// Min metrics
	// Constraints:
	//    - nullable
	Min *float64 `json:"min,omitempty"`

	// Timestamp
	// Timestamp
	// Constraints:
	//    - nullable
	Timestamp *string `json:"timestamp,omitempty"`
}

func NewSwitchMHealthMetrics() *SwitchMHealthMetrics {
	m := new(SwitchMHealthMetrics)
	return m
}

type SwitchMHealthStatus struct {
	// Fan
	// Fan
	// Constraints:
	//    - nullable
	Fan []*SwitchMHealthStatusFanType `json:"fan,omitempty" validate:"omitempty,dive"`

	// FlaggedCount
	// Flagged status count
	// Constraints:
	//    - nullable
	FlaggedCount *int `json:"flaggedCount,omitempty"`

	// OfflineCount
	// Offline status count
	// Constraints:
	//    - nullable
	OfflineCount *int `json:"offlineCount,omitempty"`

	// OnlineCount
	// Online status count
	// Constraints:
	//    - nullable
	OnlineCount *int `json:"onlineCount,omitempty"`

	// PowerSupply
	// Powersupply
	// Constraints:
	//    - nullable
	PowerSupply []*SwitchMHealthStatusPowerSupplyType `json:"powerSupply,omitempty" validate:"omitempty,dive"`

	// Temperature
	// Temperature
	// Constraints:
	//    - nullable
	Temperature []*SwitchMHealthStatusTemperatureType `json:"temperature,omitempty" validate:"omitempty,dive"`
}

func NewSwitchMHealthStatus() *SwitchMHealthStatus {
	m := new(SwitchMHealthStatus)
	return m
}

type SwitchMHealthStatusFanType struct {
	// SerialNumber
	// Constraints:
	//    - nullable
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Fan slot number
	// Constraints:
	//    - nullable
	SlotNumber *int `json:"slotNumber,omitempty"`

	// Status
	// Fan status
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// Type
	// Fan type
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewSwitchMHealthStatusFanType() *SwitchMHealthStatusFanType {
	m := new(SwitchMHealthStatusFanType)
	return m
}

type SwitchMHealthStatusPowerSupplyType struct {
	// SerialNumber
	// Constraints:
	//    - nullable
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SlotNumber
	// Power supply slot number
	// Constraints:
	//    - nullable
	SlotNumber *int `json:"slotNumber,omitempty"`

	// Status
	// Power supply status
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// Type
	// Power supply type
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewSwitchMHealthStatusPowerSupplyType() *SwitchMHealthStatusPowerSupplyType {
	m := new(SwitchMHealthStatusPowerSupplyType)
	return m
}

type SwitchMHealthStatusTemperatureType struct {
	// SlotNumber
	// Solt number
	// Constraints:
	//    - nullable
	SlotNumber *int `json:"slotNumber,omitempty"`

	// TemperatureValue
	// Slot temperature
	// Constraints:
	//    - nullable
	TemperatureValue *float64 `json:"temperatureValue,omitempty"`
}

func NewSwitchMHealthStatusTemperatureType() *SwitchMHealthStatusTemperatureType {
	m := new(SwitchMHealthStatusTemperatureType)
	return m
}

// AddHealthCpuAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthCpuAgg(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthAggMetrics, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMHealthAggMetrics
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthCpuAgg, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMHealthAggMetrics()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddHealthCpuLine
//
// Use this API command to retrieve CPU trend data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthCpuLine(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthIcxMetrics, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMHealthIcxMetrics
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthCpuLine, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMHealthIcxMetrics()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddHealthMemAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthMemAgg(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthAggMetrics, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMHealthAggMetrics
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthMemAgg, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMHealthAggMetrics()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddHealthMemLine
//
// Use this API command to retrieve switch memory trend data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthMemLine(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthIcxMetrics, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMHealthIcxMetrics
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthMemLine, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMHealthIcxMetrics()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddHealthStatus
//
// Use this API command to retrieve switch health status.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthStatus(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMHealthStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthStatus, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMHealthStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddHealthStatusAll
//
// Use this API command to retrieve fan, temperature and power supply status for the switch managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthStatusAll(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMHealthStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthStatusAll, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMHealthStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
