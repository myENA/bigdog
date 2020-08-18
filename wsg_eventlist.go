package bigdog

// API Version: v9_1

// WSGEventListEventQueryResultList
//
// Definition: eventList_eventQueryResultList
type WSGEventListEventQueryResultList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGEventListSingleEvent `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGEventListEventQueryResultList() *WSGEventListEventQueryResultList {
	m := new(WSGEventListEventQueryResultList)
	return m
}

// WSGEventListSingleEvent
//
// Definition: eventList_singleEvent
type WSGEventListSingleEvent struct {
	Activity *string `json:"activity,omitempty"`

	Category *string `json:"category,omitempty"`

	EventCode *int `json:"eventCode,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	Id *string `json:"id,omitempty"`

	InsertionTime *int `json:"insertionTime,omitempty"`

	Severity *string `json:"severity,omitempty"`
}

func NewWSGEventListSingleEvent() *WSGEventListSingleEvent {
	m := new(WSGEventListSingleEvent)
	return m
}
