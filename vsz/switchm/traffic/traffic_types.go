package traffic

// API Version: v8_0

type TopPortErrorQueryResultList struct {
	Extra             interface{}     `json:"extra,omitempty"`
	FirstIndex        *int            `json:"firstIndex,omitempty"`
	HasMore           *bool           `json:"hasMore,omitempty"`
	List              []*TrafficUsage `json:"list,omitempty"`
	RawDataTotalCount *int            `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int            `json:"totalCount,omitempty"`
}

type TopPortTrafficUsageQueryResultList struct {
	Extra             interface{}     `json:"extra,omitempty"`
	FirstIndex        *int            `json:"firstIndex,omitempty"`
	HasMore           *bool           `json:"hasMore,omitempty"`
	List              []*TrafficUsage `json:"list,omitempty"`
	RawDataTotalCount *int            `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int            `json:"totalCount,omitempty"`
}

type TopSwitchPoEUtilizationQueryResultList struct {
	Extra             interface{}     `json:"extra,omitempty"`
	FirstIndex        *int            `json:"firstIndex,omitempty"`
	HasMore           *bool           `json:"hasMore,omitempty"`
	List              []*TrafficUsage `json:"list,omitempty"`
	RawDataTotalCount *int            `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int            `json:"totalCount,omitempty"`
}

type TopTrafficUsageQueryResultList struct {
	Extra             interface{}     `json:"extra,omitempty"`
	FirstIndex        *int            `json:"firstIndex,omitempty"`
	HasMore           *bool           `json:"hasMore,omitempty"`
	List              []*TrafficUsage `json:"list,omitempty"`
	RawDataTotalCount *int            `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int            `json:"totalCount,omitempty"`
}

type Traffic struct {
	Rx        *string `json:"rx,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Total     *int    `json:"total,omitempty"`
	Tx        *string `json:"tx,omitempty"`
}

type TrafficQueryResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*Traffic  `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type TrafficUsage struct {
	ID    *string  `json:"id,omitempty"`
	Key   *string  `json:"key,omitempty"`
	Value *float64 `json:"value,omitempty"`
}
