package health

// API Version: v8_1

import (
	"encoding/json"
)

type AggMetrics struct {
	// Extra
	// Extra information for Aggregation Metrics
	Extra *AggMetricsExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Aggregation Metrics returned out of the complete ICX Metrics list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more Aggregation Metrics after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Aggs `json:"list,omitempty"`

	// RawDataTotalCount
	// Aggregation Metrics count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of Aggregation Metrics count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAggMetrics() *AggMetrics {
	aggMetricsType := new(AggMetrics)
	return aggMetricsType
}

func NewDefaultAggMetrics() *AggMetrics {
	aggMetricsType := new(AggMetrics)
	return aggMetricsType
}

// AggMetricsExtraType
//
// Extra information for Aggregation Metrics
type AggMetricsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *AggMetricsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = AggMetricsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *AggMetricsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewAggMetricsExtraType() *AggMetricsExtraType {
	aggMetricsExtraTypeType := new(AggMetricsExtraType)
	return aggMetricsExtraTypeType
}

func NewDefaultAggMetricsExtraType() *AggMetricsExtraType {
	aggMetricsExtraTypeType := new(AggMetricsExtraType)
	return aggMetricsExtraTypeType
}

type Aggs struct {
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

func NewAggs() *Aggs {
	aggsType := new(Aggs)
	return aggsType
}

func NewDefaultAggs() *Aggs {
	aggsType := new(Aggs)
	return aggsType
}

type IcxMetrics struct {
	// Extra
	// Extra information for ICX Metrics
	Extra *IcxMetricsExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ICX Metrics returned out of the complete ICX Metrics list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more ICX Metrics after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Metrics `json:"list,omitempty"`

	// RawDataTotalCount
	// ICX Metrics count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of ICX Metrics count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewIcxMetrics() *IcxMetrics {
	icxMetricsType := new(IcxMetrics)
	return icxMetricsType
}

func NewDefaultIcxMetrics() *IcxMetrics {
	icxMetricsType := new(IcxMetrics)
	return icxMetricsType
}

// IcxMetricsExtraType
//
// Extra information for ICX Metrics
type IcxMetricsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *IcxMetricsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = IcxMetricsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *IcxMetricsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewIcxMetricsExtraType() *IcxMetricsExtraType {
	icxMetricsExtraTypeType := new(IcxMetricsExtraType)
	return icxMetricsExtraTypeType
}

func NewDefaultIcxMetricsExtraType() *IcxMetricsExtraType {
	icxMetricsExtraTypeType := new(IcxMetricsExtraType)
	return icxMetricsExtraTypeType
}

type Metrics struct {
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

func NewMetrics() *Metrics {
	metricsType := new(Metrics)
	return metricsType
}

func NewDefaultMetrics() *Metrics {
	metricsType := new(Metrics)
	return metricsType
}

type Status struct {
	// Fan
	// Fan
	Fan []*StatusFanType `json:"fan,omitempty"`

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
	PowerSupply []*StatusPowerSupplyType `json:"powerSupply,omitempty"`

	// Temperature
	// Temperature
	Temperature []*StatusTemperatureType `json:"temperature,omitempty"`
}

func NewStatus() *Status {
	statusType := new(Status)
	return statusType
}

func NewDefaultStatus() *Status {
	statusType := new(Status)
	return statusType
}

type StatusFanType struct {
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

func NewStatusFanType() *StatusFanType {
	statusFanTypeType := new(StatusFanType)
	return statusFanTypeType
}

func NewDefaultStatusFanType() *StatusFanType {
	statusFanTypeType := new(StatusFanType)
	return statusFanTypeType
}

type StatusPowerSupplyType struct {
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

func NewStatusPowerSupplyType() *StatusPowerSupplyType {
	statusPowerSupplyTypeType := new(StatusPowerSupplyType)
	return statusPowerSupplyTypeType
}

func NewDefaultStatusPowerSupplyType() *StatusPowerSupplyType {
	statusPowerSupplyTypeType := new(StatusPowerSupplyType)
	return statusPowerSupplyTypeType
}

type StatusTemperatureType struct {
	// SlotNumber
	// Solt number
	SlotNumber *int `json:"slotNumber,omitempty"`

	// TemperatureValue
	// Slot temperature
	TemperatureValue *float64 `json:"temperatureValue,omitempty"`
}

func NewStatusTemperatureType() *StatusTemperatureType {
	statusTemperatureTypeType := new(StatusTemperatureType)
	return statusTemperatureTypeType
}

func NewDefaultStatusTemperatureType() *StatusTemperatureType {
	statusTemperatureTypeType := new(StatusTemperatureType)
	return statusTemperatureTypeType
}
