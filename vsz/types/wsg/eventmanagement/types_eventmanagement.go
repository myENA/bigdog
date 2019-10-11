package eventmanagement

// API Version: v8_1

type EventDataList struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SingleEventSetting `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewEventDataList() *EventDataList {
	eventDataListType := new(EventDataList)
	return eventDataListType
}

func NewEventDataListWithDefaults() *EventDataList {
	eventDataListType := new(EventDataList)
	return eventDataListType
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

func NewEventDataResponse() *EventDataResponse {
	eventDataResponseType := new(EventDataResponse)
	return eventDataResponseType
}

func NewEventDataResponseWithDefaults() *EventDataResponse {
	eventDataResponseType := new(EventDataResponse)
	return eventDataResponseType
}

type EventEmailSetting struct {
	// EmailEnabled
	// Enable/Disable Email sending function
	EmailEnabled *bool `json:"emailEnabled,omitempty"`

	// MailTo
	// E-mail recipients
	MailTo *string `json:"mailTo,omitempty"`
}

func NewEventEmailSetting() *EventEmailSetting {
	eventEmailSettingType := new(EventEmailSetting)
	return eventEmailSettingType
}

func NewEventEmailSettingWithDefaults() *EventEmailSetting {
	eventEmailSettingType := new(EventEmailSetting)
	return eventEmailSettingType
}

type EventSettingList []*SingleEventSetting

func NewEventSettingList() *EventSettingList {
	eventSettingListType := make(EventSettingList, 0)
	return &eventSettingListType
}

func NewEventSettingListWithDefaults() *EventSettingList {
	eventSettingListType := make(EventSettingList, 0)
	return &eventSettingListType
}

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

func NewSingleEventSetting() *SingleEventSetting {
	singleEventSettingType := new(SingleEventSetting)
	return singleEventSettingType
}

func NewSingleEventSettingWithDefaults() *SingleEventSetting {
	singleEventSettingType := new(SingleEventSetting)
	return singleEventSettingType
}
