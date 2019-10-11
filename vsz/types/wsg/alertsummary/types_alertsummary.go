package alertsummary

// API Version: v8_1

type AlarmSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}

func NewAlarmSummary() *AlarmSummary {
	alarmSummaryType := new(AlarmSummary)
	return alarmSummaryType
}

func NewAlarmSummaryWithDefaults() *AlarmSummary {
	alarmSummaryType := new(AlarmSummary)
	return alarmSummaryType
}

type EventSummary struct {
	CriticalCount *int `json:"criticalCount,omitempty"`

	DebugCount *int `json:"debugCount,omitempty"`

	InformationalCount *int `json:"informationalCount,omitempty"`

	MajorCount *int `json:"majorCount,omitempty"`

	MinorCount *int `json:"minorCount,omitempty"`

	WarningCount *int `json:"warningCount,omitempty"`
}

func NewEventSummary() *EventSummary {
	eventSummaryType := new(EventSummary)
	return eventSummaryType
}

func NewEventSummaryWithDefaults() *EventSummary {
	eventSummaryType := new(EventSummary)
	return eventSummaryType
}
