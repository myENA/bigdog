package vsz

// API Version: v9_0

type WSGEventManagementEventDataList struct {
	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGEventManagementSingleEventSetting `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGEventManagementEventDataList() *WSGEventManagementEventDataList {
	m := new(WSGEventManagementEventDataList)
	return m
}

type WSGEventManagementEventDataResponse struct {
	// Data
	// Constraints:
	//    - nullable
	Data *WSGEventManagementEventDataList `json:"data,omitempty"`

	// Error
	// The error message of http request
	// Constraints:
	//    - nullable
	Error *string `json:"error,omitempty"`

	// Extra
	// Extra information for event management setting
	// Constraints:
	//    - nullable
	Extra *string `json:"extra,omitempty"`

	// Success
	// The status of http request
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewWSGEventManagementEventDataResponse() *WSGEventManagementEventDataResponse {
	m := new(WSGEventManagementEventDataResponse)
	return m
}

type WSGEventManagementEventEmailSetting struct {
	// EmailEnabled
	// Enable/Disable Email sending function
	// Constraints:
	//    - nullable
	EmailEnabled *bool `json:"emailEnabled,omitempty"`

	// MailTo
	// E-mail recipients
	// Constraints:
	//    - nullable
	MailTo *string `json:"mailTo,omitempty"`
}

func NewWSGEventManagementEventEmailSetting() *WSGEventManagementEventEmailSetting {
	m := new(WSGEventManagementEventEmailSetting)
	return m
}

type WSGEventManagementEventSettingList []*WSGEventManagementSingleEventSetting

func MakeWSGEventManagementEventSettingList() WSGEventManagementEventSettingList {
	m := make(WSGEventManagementEventSettingList, 0)
	return m
}

type WSGEventManagementSingleEventSetting struct {
	// Category
	// Event category
	// Constraints:
	//    - nullable
	Category *string `json:"category,omitempty"`

	// ConfigPageDesc
	// Event description
	// Constraints:
	//    - nullable
	ConfigPageDesc *string `json:"configPageDesc,omitempty"`

	// DbPersistence
	// Enable/Disable DB persistence for this event
	// Constraints:
	//    - nullable
	DbPersistence *bool `json:"dbPersistence,omitempty"`

	// EventCode
	// Event code
	// Constraints:
	//    - nullable
	EventCode *int `json:"eventCode,omitempty"`

	// Oid
	// Event OID
	// Constraints:
	//    - nullable
	Oid *string `json:"oid,omitempty"`

	// Severity
	// Event severity
	// Constraints:
	//    - nullable
	Severity *string `json:"severity,omitempty"`

	// TriggerEmail
	// Enable/Disable Email sending for this event
	// Constraints:
	//    - nullable
	TriggerEmail *bool `json:"triggerEmail,omitempty"`

	// TriggerTrap
	// Enable/Disable SNMP function for this event
	// Constraints:
	//    - nullable
	TriggerTrap *bool `json:"triggerTrap,omitempty"`

	// Type
	// Event type
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`

	// ZoneOverride
	// Enable/Disable override event system settings
	// Constraints:
	//    - nullable
	ZoneOverride *bool `json:"zoneOverride,omitempty"`
}

func NewWSGEventManagementSingleEventSetting() *WSGEventManagementSingleEventSetting {
	m := new(WSGEventManagementSingleEventSetting)
	return m
}
