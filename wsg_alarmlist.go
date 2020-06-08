package ruckus

// API Version: v9_0

type WSGAlarmListAlarmQueryResultList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGAlarmListAlarmQueryResultList `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAlarmListAlarmQueryResultList() *WSGAlarmListAlarmQueryResultList {
	m := new(WSGAlarmListAlarmQueryResultList)
	return m
}

type WSGAlarmListSingleAlarm struct {
	Acknowledged *string `json:"acknowledged,omitempty"`

	AckTime *int `json:"ackTime,omitempty"`

	AckUser *string `json:"ackUser,omitempty"`

	Activity *string `json:"activity,omitempty"`

	AlarmCode *int `json:"alarmCode,omitempty"`

	AlarmState *string `json:"alarmState,omitempty"`

	AlarmType *string `json:"alarmType,omitempty"`

	Category *string `json:"category,omitempty"`

	ClearComment *string `json:"clearComment,omitempty"`

	ClearTime *int `json:"clearTime,omitempty"`

	ClearUser *string `json:"clearUser,omitempty"`

	Id *string `json:"id,omitempty"`

	InsertionTime *int `json:"insertionTime,omitempty"`

	Severity *string `json:"severity,omitempty"`
}

func NewWSGAlarmListSingleAlarm() *WSGAlarmListSingleAlarm {
	m := new(WSGAlarmListSingleAlarm)
	return m
}
