package job

// API Version: v8_0

type ErrorObject struct {
	List    []string `json:"list,omitempty"`
	Message *string  `json:"message,omitempty"`
	MsgKey  *string  `json:"msgKey,omitempty"`
}

type Job struct {
	Action                *string           `json:"action,omitempty"`
	CreatedTimestamp      *int              `json:"createdTimestamp,omitempty"`
	CsvDataMap            map[string]string `json:"csvDataMap,omitempty"`
	DomainID              *string           `json:"domainId,omitempty"`
	FailureReason         *string           `json:"failureReason,omitempty"`
	ID                    *string           `json:"id,omitempty"`
	ModifiedTimestamp     *int              `json:"modifiedTimestamp,omitempty"`
	ScheduleID            *string           `json:"scheduleId,omitempty"`
	Status                *string           `json:"status,omitempty"`
	StickyNodeID          *string           `json:"stickyNodeId,omitempty"`
	SwitchGroupLevelOneID *string           `json:"switchGroupLevelOneId,omitempty"`
	SwitchGroupLevelTwoID *string           `json:"switchGroupLevelTwoId,omitempty"`
	SwitchID              *string           `json:"switchId,omitempty"`
	TenantID              *string           `json:"tenantId,omitempty"`
	Type                  *string           `json:"type,omitempty"`
}

type JobSchedule struct {
	CreatedTimestamp *int     `json:"createdTimestamp,omitempty"`
	JobID            []string `json:"jobId,omitempty"`
	TriggerValue     *string  `json:"triggerValue,omitempty"`
}

type JobScheduleResponse struct {
	Data     *JobSchedule           `json:"data,omitempty"`
	Error    *ErrorObject           `json:"error,omitempty"`
	Extra    interface{}            `json:"extra,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Success  *bool                  `json:"success,omitempty"`
}

type List struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*Job      `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}
