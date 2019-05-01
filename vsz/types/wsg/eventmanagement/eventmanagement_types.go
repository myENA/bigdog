package eventmanagement

// API Version: v8_0

type EventDataList struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SingleEventSetting `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type EventDataResponse struct {
	Data *EventDataList `json:"data,omitempty"`

	// Error
	// The error message of http request
	Error *string `json:"error,omitempty"`

	// Extra
	// Extra information for event management setting
	Extra *string `json:"extra,omitempty"`

	// Success
	// The status of http request
	Success *bool `json:"success,omitempty"`
}

type EventEmailSetting struct {
	// EmailEnabled
	// Enable/Disable Email sending function
	EmailEnabled *bool `json:"emailEnabled,omitempty"`

	// MailTo
	// E-mail recipients
	MailTo *string `json:"mailTo,omitempty"`
}

type EventSettingList []*SingleEventSetting

type SingleEventSetting struct {
	// Category
	// Event category
	Category *string `json:"category,omitempty"`

	// ConfigPageDesc
	// Event description
	ConfigPageDesc *string `json:"configPageDesc,omitempty"`

	// DbPersistence
	// Enable/Disable DB persistence for this event
	DbPersistence *bool `json:"dbPersistence,omitempty"`

	// EventCode
	// Event code
	EventCode *int `json:"eventCode,omitempty"`

	// Oid
	// Event OID
	Oid *string `json:"oid,omitempty"`

	// Severity
	// Event severity
	Severity *string `json:"severity,omitempty"`

	// TriggerEmail
	// Enable/Disable Email sending for this event
	TriggerEmail *bool `json:"triggerEmail,omitempty"`

	// TriggerTrap
	// Enable/Disable SNMP function for this event
	TriggerTrap *bool `json:"triggerTrap,omitempty"`

	// Type
	// Event type
	Type *string `json:"type,omitempty"`

	// ZoneOverride
	// Enable/Disable override event system settings
	ZoneOverride *bool `json:"zoneOverride,omitempty"`
}
