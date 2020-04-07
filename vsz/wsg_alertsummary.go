package vsz

// API Version: v9_0

type WSGAlertSummaryAlarmSummary struct {
	// CriticalCount
	// Constraints:
	//    - nullable
	CriticalCount *int `json:"criticalCount,omitempty"`

	// MajorCount
	// Constraints:
	//    - nullable
	MajorCount *int `json:"majorCount,omitempty"`

	// MinorCount
	// Constraints:
	//    - nullable
	MinorCount *int `json:"minorCount,omitempty"`

	// WarningCount
	// Constraints:
	//    - nullable
	WarningCount *int `json:"warningCount,omitempty"`
}

func NewWSGAlertSummaryAlarmSummary() *WSGAlertSummaryAlarmSummary {
	m := new(WSGAlertSummaryAlarmSummary)
	return m
}

type WSGAlertSummaryEventSummary struct {
	// CriticalCount
	// Constraints:
	//    - nullable
	CriticalCount *int `json:"criticalCount,omitempty"`

	// DebugCount
	// Constraints:
	//    - nullable
	DebugCount *int `json:"debugCount,omitempty"`

	// InformationalCount
	// Constraints:
	//    - nullable
	InformationalCount *int `json:"informationalCount,omitempty"`

	// MajorCount
	// Constraints:
	//    - nullable
	MajorCount *int `json:"majorCount,omitempty"`

	// MinorCount
	// Constraints:
	//    - nullable
	MinorCount *int `json:"minorCount,omitempty"`

	// WarningCount
	// Constraints:
	//    - nullable
	WarningCount *int `json:"warningCount,omitempty"`
}

func NewWSGAlertSummaryEventSummary() *WSGAlertSummaryEventSummary {
	m := new(WSGAlertSummaryEventSummary)
	return m
}
