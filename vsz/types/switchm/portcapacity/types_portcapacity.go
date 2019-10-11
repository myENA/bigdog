package portcapacity

// API Version: v8_1

import (
	"encoding/json"
)

type Capacities struct {
	// Capacity
	// Port Speed Capacity
	Capacity *string `json:"capacity,omitempty"`
}

func NewCapacities() *Capacities {
	capacitiesType := new(Capacities)
	return capacitiesType
}

func NewCapacitiesWithDefaults() *Capacities {
	capacitiesType := new(Capacities)
	return capacitiesType
}

type Result struct {
	// Extra
	// Extra field
	Extra *ResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// The first data index for current reulst
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of remaining data
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Capacities `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Data Count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Data Count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewResult() *Result {
	resultType := new(Result)
	return resultType
}

func NewResultWithDefaults() *Result {
	resultType := new(Result)
	return resultType
}

// ResultExtraType
//
// Extra field
type ResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *ResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewResultExtraType() *ResultExtraType {
	resultExtraTypeType := new(ResultExtraType)
	return resultExtraTypeType
}

func NewResultExtraTypeWithDefaults() *ResultExtraType {
	resultExtraTypeType := new(ResultExtraType)
	return resultExtraTypeType
}
