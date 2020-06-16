package bigdog

// API Version: v9_0

// WSGAlertSummaryAlarmSummary
//
// Definition: alertSummary_alarmSummary
type WSGAlertSummaryAlarmSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
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

func NewWSGAlertSummaryEventSummary() *WSGAlertSummaryEventSummary {
	m := new(WSGAlertSummaryEventSummary)
	return m
}
