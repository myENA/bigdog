package eventlist

// API Version: v8_0

type EventQueryResultList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SingleEvent `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SingleEvent struct {
	Activity *string `json:"activity,omitempty"`

	Category *string `json:"category,omitempty"`

	EventCode *int `json:"eventCode,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	ID *string `json:"id,omitempty"`

	InsertionTime *float64 `json:"insertionTime,omitempty"`

	Severity *string `json:"severity,omitempty"`
}
