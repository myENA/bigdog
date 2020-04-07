package vsz

// API Version: v9_0

type WSGAlarmListAlarmQueryResultList struct {
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
	List []*WSGAlarmListSingleAlarm `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAlarmListAlarmQueryResultList() *WSGAlarmListAlarmQueryResultList {
	m := new(WSGAlarmListAlarmQueryResultList)
	return m
}

type WSGAlarmListSingleAlarm struct {
	// Acknowledged
	// Constraints:
	//    - nullable
	Acknowledged *string `json:"acknowledged,omitempty"`

	// AckTime
	// Constraints:
	//    - nullable
	AckTime *int `json:"ackTime,omitempty"`

	// AckUser
	// Constraints:
	//    - nullable
	AckUser *string `json:"ackUser,omitempty"`

	// Activity
	// Constraints:
	//    - nullable
	Activity *string `json:"activity,omitempty"`

	// AlarmCode
	// Constraints:
	//    - nullable
	AlarmCode *int `json:"alarmCode,omitempty"`

	// AlarmState
	// Constraints:
	//    - nullable
	AlarmState *string `json:"alarmState,omitempty"`

	// AlarmType
	// Constraints:
	//    - nullable
	AlarmType *string `json:"alarmType,omitempty"`

	// Category
	// Constraints:
	//    - nullable
	Category *string `json:"category,omitempty"`

	// ClearComment
	// Constraints:
	//    - nullable
	ClearComment *string `json:"clearComment,omitempty"`

	// ClearTime
	// Constraints:
	//    - nullable
	ClearTime *int `json:"clearTime,omitempty"`

	// ClearUser
	// Constraints:
	//    - nullable
	ClearUser *string `json:"clearUser,omitempty"`

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

func NewWSGAlarmListSingleAlarm() *WSGAlarmListSingleAlarm {
	m := new(WSGAlarmListSingleAlarm)
	return m
}
