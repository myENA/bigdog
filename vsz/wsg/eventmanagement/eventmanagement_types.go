package eventmanagement

// API Version: v8_0

type EventDataList struct {
	HasMore           *bool                 `json:"hasMore,omitempty"`
	List              []*SingleEventSetting `json:"list,omitempty"`
	RawDataTotalCount *int                  `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int                  `json:"totalCount,omitempty"`
}

type EventDataResponse struct {
	Data    *EventDataList `json:"data,omitempty"`
	Error   *string        `json:"error,omitempty"`
	Extra   *string        `json:"extra,omitempty"`
	Success *bool          `json:"success,omitempty"`
}

type EventEmailSetting struct {
	EmailEnabled *bool   `json:"emailEnabled,omitempty"`
	MailTo       *string `json:"mailTo,omitempty"`
}

type EventSettingList []*SingleEventSetting

type SingleEventSetting struct {
	Category       *string `json:"category,omitempty"`
	ConfigPageDesc *string `json:"configPageDesc,omitempty"`
	DbPersistence  *bool   `json:"dbPersistence,omitempty"`
	EventCode      *int    `json:"eventCode,omitempty"`
	Oid            *string `json:"oid,omitempty"`
	Severity       *string `json:"severity,omitempty"`
	TriggerEmail   *bool   `json:"triggerEmail,omitempty"`
	TriggerTrap    *bool   `json:"triggerTrap,omitempty"`
	Type           *string `json:"type,omitempty"`
	ZoneOverride   *bool   `json:"zoneOverride,omitempty"`
}
