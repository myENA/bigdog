package job

// API Version: v8_0

type ErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

type Job struct {
	// Action
	// Action of the job
	Action *string `json:"action,omitempty"`

	// CreatedTimestamp
	// Created timestamp of the job
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	CsvDataMap map[string]string `json:"csvDataMap,omitempty"`

	// DomainID
	// Identifier of the management domain to which the job belong
	DomainID *string `json:"domainId,omitempty"`

	// FailureReason
	// Failure reason of the job
	FailureReason *string `json:"failureReason,omitempty"`

	// ID
	// Identifier of the job
	ID *string `json:"id,omitempty"`

	// ModifiedTimestamp
	// Modified timestamp of the job
	ModifiedTimestamp *int `json:"modifiedTimestamp,omitempty"`

	// ScheduleID
	// Schedule Id of the job
	ScheduleID *string `json:"scheduleId,omitempty"`

	// Status
	// Status of the job
	Status *string `json:"status,omitempty"`

	// StickyNodeID
	// Sticky node Id of the job
	StickyNodeID *string `json:"stickyNodeId,omitempty"`

	// SwitchGroupLevelOneID
	// Switch group level one Id of the job
	SwitchGroupLevelOneID *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoID
	// Switch group level two Id of the job
	SwitchGroupLevelTwoID *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchID
	// Switch Id of the job
	SwitchID *string `json:"switchId,omitempty"`

	// TenantID
	// Tenant Id of the job
	TenantID *string `json:"tenantId,omitempty"`

	// Type
	// Type of the job
	Type *string `json:"type,omitempty"`
}

type JobSchedule struct {
	// CreatedTimestamp
	// Created timestamp of job schedule
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	// JobID
	// Job Id of job schedule
	JobID []string `json:"jobId,omitempty"`

	// TriggerValue
	// Trigger value of job schedule
	TriggerValue *string `json:"triggerValue,omitempty"`
}

type JobScheduleResponse struct {
	Data *JobSchedule `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Extra response of job schedule
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// metaData of job schedule
	MetaData map[string]interface{} `json:"metaData,omitempty"`

	// Success
	// Success response of job schedule
	Success *bool `json:"success,omitempty"`
}

type List struct {
	// Extra
	// Extra information for job list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first job returned out of the complete job list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more jobs after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Job `json:"list,omitempty"`

	// RawDataTotalCount
	// List count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total list count
	TotalCount *int `json:"totalCount,omitempty"`
}
