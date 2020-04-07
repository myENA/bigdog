package vsz

// API Version: v9_0

type WSGEventListEventQueryResultList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGEventListSingleEvent `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGEventListEventQueryResultList() *WSGEventListEventQueryResultList {
	m := new(WSGEventListEventQueryResultList)
	return m
}

type WSGEventListSingleEvent struct {
	// Activity
	// Constraints:
	//    - nullable
	Activity *string `json:"activity,omitempty"`

	// Category
	// Constraints:
	//    - nullable
	Category *string `json:"category,omitempty"`

	// EventCode
	// Constraints:
	//    - nullable
	EventCode *int `json:"eventCode,omitempty"`

	// EventType
	// Constraints:
	//    - nullable
	EventType *string `json:"eventType,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// InsertionTime
	// Constraints:
	//    - nullable
	InsertionTime *int `json:"insertionTime,omitempty"`

	// Severity
	// Constraints:
	//    - nullable
	Severity *string `json:"severity,omitempty"`
}

func NewWSGEventListSingleEvent() *WSGEventListSingleEvent {
	m := new(WSGEventListSingleEvent)
	return m
}
