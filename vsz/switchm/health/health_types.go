package health

// API Version: v8_0

type AggMetrics struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type Aggs struct {
	ID    *string  `json:"id,omitempty"`
	Key   *string  `json:"key,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

type IcxMetrics struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type Metrics struct {
	Avg       *float64 `json:"avg,omitempty"`
	Max       *float64 `json:"max,omitempty"`
	Min       *float64 `json:"min,omitempty"`
	Timestamp *string  `json:"timestamp,omitempty"`
}

type Status struct {
	Fan          []*Fan         `json:"fan,omitempty"`
	FlaggedCount *int           `json:"flaggedCount,omitempty"`
	OfflineCount *int           `json:"offlineCount,omitempty"`
	OnlineCount  *int           `json:"onlineCount,omitempty"`
	PowerSupply  []*PowerSupply `json:"powerSupply,omitempty"`
	Temperature  []*Temperature `json:"temperature,omitempty"`
}

type StatusFanType struct {
	SlotNumber *int    `json:"slotNumber,omitempty"`
	Status     *string `json:"status,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type StatusPowerSupplyType struct {
	SlotNumber *int    `json:"slotNumber,omitempty"`
	Status     *string `json:"status,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type StatusTemperatureType struct {
	SlotNumber       *int     `json:"slotNumber,omitempty"`
	TemperatureValue *float64 `json:"temperatureValue,omitempty"`
}
