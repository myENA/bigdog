package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGAlertSummaryAlarmSummary
//
// Definition: alertSummary_alarmSummary
type WSGAlertSummaryAlarmSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}

type WSGAlertSummaryAlarmSummaryAPIResponse struct {
	*RawAPIResponse
	Data *WSGAlertSummaryAlarmSummary
}

func newWSGAlertSummaryAlarmSummaryAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAlertSummaryAlarmSummaryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAlertSummaryAlarmSummaryAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAlertSummaryAlarmSummary)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAlertSummaryAlarmSummary() *WSGAlertSummaryAlarmSummary {
	m := new(WSGAlertSummaryAlarmSummary)
	return m
}

// WSGAlertSummaryEventSummary
//
// Definition: alertSummary_eventSummary
type WSGAlertSummaryEventSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	DebugCount *int `json:"debugCount,omitempty"`

	InformationalCount *int `json:"informationalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}

type WSGAlertSummaryEventSummaryAPIResponse struct {
	*RawAPIResponse
	Data *WSGAlertSummaryEventSummary
}

func newWSGAlertSummaryEventSummaryAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAlertSummaryEventSummaryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAlertSummaryEventSummaryAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAlertSummaryEventSummary)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAlertSummaryEventSummary() *WSGAlertSummaryEventSummary {
	m := new(WSGAlertSummaryEventSummary)
	return m
}
