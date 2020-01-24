package vsz

// API Version: v8_1

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
	Extra *SwitchMHealthAggMetricsExtraType `json:"extra,omitempty"`

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

func NewSwitchMHealthAggMetrics() *SwitchMHealthAggMetrics {
	m := new(SwitchMHealthAggMetrics)
	return m
}

// SwitchMHealthAggMetricsExtraType
//
// Extra information for Aggregation Metrics
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

type SwitchMHealthIcxMetrics struct {
	// Extra
	// Extra information for ICX Metrics
	Extra *SwitchMHealthIcxMetricsExtraType `json:"extra,omitempty"`

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

func NewSwitchMHealthIcxMetrics() *SwitchMHealthIcxMetrics {
	m := new(SwitchMHealthIcxMetrics)
	return m
}

// SwitchMHealthIcxMetricsExtraType
//
// Extra information for ICX Metrics
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
	Avg *float64 `json:"avg,omitempty"`

	// Max
	// Max metrics
	Max *float64 `json:"max,omitempty"`

	// Min
	// Min metrics
	Min *float64 `json:"min,omitempty"`

	// Timestamp
	// Timestamp
	Timestamp *string `json:"timestamp,omitempty"`
}

func NewSwitchMHealthMetrics() *SwitchMHealthMetrics {
	m := new(SwitchMHealthMetrics)
	return m
}

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

	// PowerSupply
	// Powersupply
	PowerSupply []*SwitchMHealthStatusPowerSupplyType `json:"powerSupply,omitempty"`

	// Temperature
	// Temperature
	Temperature []*SwitchMHealthStatusTemperatureType `json:"temperature,omitempty"`
}

func NewSwitchMHealthStatus() *SwitchMHealthStatus {
	m := new(SwitchMHealthStatus)
	return m
}

type SwitchMHealthStatusFanType struct {
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

type SwitchMHealthStatusPowerSupplyType struct {
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

type SwitchMHealthStatusTemperatureType struct {
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

// AddHealthCpuAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthCpuAgg(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthAggMetrics, error) {
	var (
		resp *SwitchMHealthAggMetrics
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthCpuAgg, true)
}

// AddHealthCpuLine
//
// Use this API command to retrieve CPU trend data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthCpuLine(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthIcxMetrics, error) {
	var (
		resp *SwitchMHealthIcxMetrics
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthCpuLine, true)
}

// AddHealthMemAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthMemAgg(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthAggMetrics, error) {
	var (
		resp *SwitchMHealthAggMetrics
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthMemAgg, true)
}

// AddHealthMemLine
//
// Use this API command to retrieve switch memory trend data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthMemLine(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthIcxMetrics, error) {
	var (
		resp *SwitchMHealthIcxMetrics
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthMemLine, true)
}

// AddHealthStatus
//
// Use this API command to retrieve switch health status.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthStatus(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthStatus, error) {
	var (
		resp *SwitchMHealthStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthStatus, true)
}

// AddHealthStatusAll
//
// Use this API command to retrieve fan, temperature and power supply status for the switch managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMHealthService) AddHealthStatusAll(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMHealthStatus, error) {
	var (
		resp *SwitchMHealthStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddHealthStatusAll, true)
}