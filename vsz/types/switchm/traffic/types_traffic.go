package traffic

// API Version: v8_1

import (
	"encoding/json"
)

type TopPortErrorQueryResultList struct {
	// Extra
	// Extra information for top port error
	Extra *TopPortErrorQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port error returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port error after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*TrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port error count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port error count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewTopPortErrorQueryResultList() *TopPortErrorQueryResultList {
	topPortErrorQueryResultListType := new(TopPortErrorQueryResultList)
	return topPortErrorQueryResultListType
}

func NewTopPortErrorQueryResultListWithDefaults() *TopPortErrorQueryResultList {
	topPortErrorQueryResultListType := new(TopPortErrorQueryResultList)
	return topPortErrorQueryResultListType
}

// TopPortErrorQueryResultListExtraType
//
// Extra information for top port error
type TopPortErrorQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *TopPortErrorQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = TopPortErrorQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *TopPortErrorQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewTopPortErrorQueryResultListExtraType() *TopPortErrorQueryResultListExtraType {
	topPortErrorQueryResultListExtraTypeType := new(TopPortErrorQueryResultListExtraType)
	return topPortErrorQueryResultListExtraTypeType
}

func NewTopPortErrorQueryResultListExtraTypeWithDefaults() *TopPortErrorQueryResultListExtraType {
	topPortErrorQueryResultListExtraTypeType := new(TopPortErrorQueryResultListExtraType)
	return topPortErrorQueryResultListExtraTypeType
}

type TopPortTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top port traffic usage
	Extra *TopPortTrafficUsageQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*TrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewTopPortTrafficUsageQueryResultList() *TopPortTrafficUsageQueryResultList {
	topPortTrafficUsageQueryResultListType := new(TopPortTrafficUsageQueryResultList)
	return topPortTrafficUsageQueryResultListType
}

func NewTopPortTrafficUsageQueryResultListWithDefaults() *TopPortTrafficUsageQueryResultList {
	topPortTrafficUsageQueryResultListType := new(TopPortTrafficUsageQueryResultList)
	return topPortTrafficUsageQueryResultListType
}

// TopPortTrafficUsageQueryResultListExtraType
//
// Extra information for top port traffic usage
type TopPortTrafficUsageQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *TopPortTrafficUsageQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = TopPortTrafficUsageQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *TopPortTrafficUsageQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewTopPortTrafficUsageQueryResultListExtraType() *TopPortTrafficUsageQueryResultListExtraType {
	topPortTrafficUsageQueryResultListExtraTypeType := new(TopPortTrafficUsageQueryResultListExtraType)
	return topPortTrafficUsageQueryResultListExtraTypeType
}

func NewTopPortTrafficUsageQueryResultListExtraTypeWithDefaults() *TopPortTrafficUsageQueryResultListExtraType {
	topPortTrafficUsageQueryResultListExtraTypeType := new(TopPortTrafficUsageQueryResultListExtraType)
	return topPortTrafficUsageQueryResultListExtraTypeType
}

type TopSwitchPoEUtilizationQueryResultList struct {
	// Extra
	// Extra information for top PoE utilization
	Extra *TopSwitchPoEUtilizationQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top PoE usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top PoE usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*TrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// PoE utilization count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total PoE utilization count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewTopSwitchPoEUtilizationQueryResultList() *TopSwitchPoEUtilizationQueryResultList {
	topSwitchPoEUtilizationQueryResultListType := new(TopSwitchPoEUtilizationQueryResultList)
	return topSwitchPoEUtilizationQueryResultListType
}

func NewTopSwitchPoEUtilizationQueryResultListWithDefaults() *TopSwitchPoEUtilizationQueryResultList {
	topSwitchPoEUtilizationQueryResultListType := new(TopSwitchPoEUtilizationQueryResultList)
	return topSwitchPoEUtilizationQueryResultListType
}

// TopSwitchPoEUtilizationQueryResultListExtraType
//
// Extra information for top PoE utilization
type TopSwitchPoEUtilizationQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *TopSwitchPoEUtilizationQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = TopSwitchPoEUtilizationQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *TopSwitchPoEUtilizationQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewTopSwitchPoEUtilizationQueryResultListExtraType() *TopSwitchPoEUtilizationQueryResultListExtraType {
	topSwitchPoEUtilizationQueryResultListExtraTypeType := new(TopSwitchPoEUtilizationQueryResultListExtraType)
	return topSwitchPoEUtilizationQueryResultListExtraTypeType
}

func NewTopSwitchPoEUtilizationQueryResultListExtraTypeWithDefaults() *TopSwitchPoEUtilizationQueryResultListExtraType {
	topSwitchPoEUtilizationQueryResultListExtraTypeType := new(TopSwitchPoEUtilizationQueryResultListExtraType)
	return topSwitchPoEUtilizationQueryResultListExtraTypeType
}

type TopTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top traffic usage
	Extra *TopTrafficUsageQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*TrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewTopTrafficUsageQueryResultList() *TopTrafficUsageQueryResultList {
	topTrafficUsageQueryResultListType := new(TopTrafficUsageQueryResultList)
	return topTrafficUsageQueryResultListType
}

func NewTopTrafficUsageQueryResultListWithDefaults() *TopTrafficUsageQueryResultList {
	topTrafficUsageQueryResultListType := new(TopTrafficUsageQueryResultList)
	return topTrafficUsageQueryResultListType
}

// TopTrafficUsageQueryResultListExtraType
//
// Extra information for top traffic usage
type TopTrafficUsageQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *TopTrafficUsageQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = TopTrafficUsageQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *TopTrafficUsageQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewTopTrafficUsageQueryResultListExtraType() *TopTrafficUsageQueryResultListExtraType {
	topTrafficUsageQueryResultListExtraTypeType := new(TopTrafficUsageQueryResultListExtraType)
	return topTrafficUsageQueryResultListExtraTypeType
}

func NewTopTrafficUsageQueryResultListExtraTypeWithDefaults() *TopTrafficUsageQueryResultListExtraType {
	topTrafficUsageQueryResultListExtraTypeType := new(TopTrafficUsageQueryResultListExtraType)
	return topTrafficUsageQueryResultListExtraTypeType
}

type Traffic struct {
	// Rx
	// RX traffic of the switch
	Rx *string `json:"rx,omitempty"`

	// Timestamp
	// Timestamp of the switch traffic
	Timestamp *string `json:"timestamp,omitempty"`

	// Total
	// Total traffic of the switch
	Total *int `json:"total,omitempty"`

	// Tx
	// TX traffic of the switch
	Tx *string `json:"tx,omitempty"`
}

func NewTraffic() *Traffic {
	trafficType := new(Traffic)
	return trafficType
}

func NewTrafficWithDefaults() *Traffic {
	trafficType := new(Traffic)
	return trafficType
}

type TrafficQueryResultList struct {
	// Extra
	// Extra information for traffic list
	Extra *TrafficQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first traffic list returned out of the complete traffic list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more traffic list after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Traffic `json:"list,omitempty"`

	// RawDataTotalCount
	// Total traffic count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of traffic list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewTrafficQueryResultList() *TrafficQueryResultList {
	trafficQueryResultListType := new(TrafficQueryResultList)
	return trafficQueryResultListType
}

func NewTrafficQueryResultListWithDefaults() *TrafficQueryResultList {
	trafficQueryResultListType := new(TrafficQueryResultList)
	return trafficQueryResultListType
}

// TrafficQueryResultListExtraType
//
// Extra information for traffic list
type TrafficQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *TrafficQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = TrafficQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *TrafficQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewTrafficQueryResultListExtraType() *TrafficQueryResultListExtraType {
	trafficQueryResultListExtraTypeType := new(TrafficQueryResultListExtraType)
	return trafficQueryResultListExtraTypeType
}

func NewTrafficQueryResultListExtraTypeWithDefaults() *TrafficQueryResultListExtraType {
	trafficQueryResultListExtraTypeType := new(TrafficQueryResultListExtraType)
	return trafficQueryResultListExtraTypeType
}

type TrafficUsage struct {
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

func NewTrafficUsage() *TrafficUsage {
	trafficUsageType := new(TrafficUsage)
	return trafficUsageType
}

func NewTrafficUsageWithDefaults() *TrafficUsage {
	trafficUsageType := new(TrafficUsage)
	return trafficUsageType
}
