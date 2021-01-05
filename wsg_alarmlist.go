package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// WSGAlarmListAlarmQueryResultList
//
// Definition: alarmList_alarmQueryResultList
type WSGAlarmListAlarmQueryResultList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAlarmListSingleAlarm `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAlarmListAlarmQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAlarmListAlarmQueryResultList
}

func newWSGAlarmListAlarmQueryResultListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGAlarmListAlarmQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGAlarmListAlarmQueryResultListAPIResponse) Hydrate() error {
	r.Data = new(WSGAlarmListAlarmQueryResultList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAlarmListAlarmQueryResultList() *WSGAlarmListAlarmQueryResultList {
	m := new(WSGAlarmListAlarmQueryResultList)
	return m
}

// WSGAlarmListSingleAlarm
//
// Definition: alarmList_singleAlarm
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
