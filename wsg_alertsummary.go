package ruckus

// API Version: v9_0

type WSGAlertsummaryAlarmSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}

func NewWSGAlertsummaryAlarmSummary() *WSGAlertsummaryAlarmSummary {
	m := new(WSGAlertsummaryAlarmSummary)
	return m
}

type WSGAlertsummaryEventSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	DebugCount *int `json:"debugCount,omitempty"`

	InformationalCount *int `json:"informationalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}

func NewWSGAlertsummaryEventSummary() *WSGAlertsummaryEventSummary {
	m := new(WSGAlertsummaryEventSummary)
	return m
}
