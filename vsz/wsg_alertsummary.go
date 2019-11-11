package vsz

// API Version: v8_1

type WSGAlertSummaryAlarmSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}

type WSGAlertSummaryEventSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	DebugCount *int `json:"debugCount,omitempty"`

	InformationalCount *int `json:"informationalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}
