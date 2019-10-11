package job

// API Version: v8_1

import (
	"encoding/json"
)

type ErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewErrorObject() *ErrorObject {
	errorObjectType := new(ErrorObject)
	return errorObjectType
}

func NewErrorObjectWithDefaults() *ErrorObject {
	errorObjectType := new(ErrorObject)
	return errorObjectType
}

type Job struct {
	// Action
	// Action of the job
	Action *string `json:"action,omitempty"`

	// CreatedTimestamp
	// Created timestamp of the job
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	CsvDataMap *JobCsvDataMapType `json:"csvDataMap,omitempty"`

	// DomainId
	// Identifier of the management domain to which the job belong
	DomainId *string `json:"domainId,omitempty"`

	// FailureReason
	// Failure reason of the job
	FailureReason *string `json:"failureReason,omitempty"`

	// Id
	// Identifier of the job
	Id *string `json:"id,omitempty"`

	// ModifiedTimestamp
	// Modified timestamp of the job
	ModifiedTimestamp *int `json:"modifiedTimestamp,omitempty"`

	// ScheduleId
	// Schedule Id of the job
	ScheduleId *string `json:"scheduleId,omitempty"`

	// Status
	// Status of the job
	Status *string `json:"status,omitempty"`

	// StickyNodeId
	// Sticky node Id of the job
	StickyNodeId *string `json:"stickyNodeId,omitempty"`

	// SwitchGroupLevelOneId
	// Switch group level one Id of the job
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Switch group level two Id of the job
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchId
	// Switch Id of the job
	SwitchId *string `json:"switchId,omitempty"`

	// TenantId
	// Tenant Id of the job
	TenantId *string `json:"tenantId,omitempty"`

	// Type
	// Type of the job
	Type *string `json:"type,omitempty"`
}

func NewJob() *Job {
	jobType := new(Job)
	return jobType
}

func NewJobWithDefaults() *Job {
	jobType := new(Job)
	return jobType
}

type JobCsvDataMapType struct {
	XAdditionalProperties map[string]string `json:"-"`
}

func (t *JobCsvDataMapType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]string)
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = JobCsvDataMapType{XAdditionalProperties: tmp}
	return nil
}

func (t *JobCsvDataMapType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewJobCsvDataMapType() *JobCsvDataMapType {
	jobCsvDataMapTypeType := new(JobCsvDataMapType)
	return jobCsvDataMapTypeType
}

func NewJobCsvDataMapTypeWithDefaults() *JobCsvDataMapType {
	jobCsvDataMapTypeType := new(JobCsvDataMapType)
	return jobCsvDataMapTypeType
}

type JobSchedule struct {
	// CreatedTimestamp
	// Created timestamp of job schedule
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	// JobId
	// Job Id of job schedule
	JobId []string `json:"jobId,omitempty"`

	// TriggerValue
	// Trigger value of job schedule
	TriggerValue *string `json:"triggerValue,omitempty"`
}

func NewJobSchedule() *JobSchedule {
	jobScheduleType := new(JobSchedule)
	return jobScheduleType
}

func NewJobScheduleWithDefaults() *JobSchedule {
	jobScheduleType := new(JobSchedule)
	return jobScheduleType
}

type JobScheduleResponse struct {
	Data *JobSchedule `json:"data,omitempty"`

	Error *ErrorObject `json:"error,omitempty"`

	// Extra
	// Extra response of job schedule
	Extra *JobScheduleResponseExtraType `json:"extra,omitempty"`

	// MetaData
	// metaData of job schedule
	MetaData *JobScheduleResponseMetaDataType `json:"metaData,omitempty"`

	// Success
	// Success response of job schedule
	Success *bool `json:"success,omitempty"`
}

func NewJobScheduleResponse() *JobScheduleResponse {
	jobScheduleResponseType := new(JobScheduleResponse)
	return jobScheduleResponseType
}

func NewJobScheduleResponseWithDefaults() *JobScheduleResponse {
	jobScheduleResponseType := new(JobScheduleResponse)
	return jobScheduleResponseType
}

// JobScheduleResponseExtraType
//
// Extra response of job schedule
type JobScheduleResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *JobScheduleResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = JobScheduleResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *JobScheduleResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewJobScheduleResponseExtraType() *JobScheduleResponseExtraType {
	jobScheduleResponseExtraTypeType := new(JobScheduleResponseExtraType)
	return jobScheduleResponseExtraTypeType
}

func NewJobScheduleResponseExtraTypeWithDefaults() *JobScheduleResponseExtraType {
	jobScheduleResponseExtraTypeType := new(JobScheduleResponseExtraType)
	return jobScheduleResponseExtraTypeType
}

// JobScheduleResponseMetaDataType
//
// metaData of job schedule
type JobScheduleResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *JobScheduleResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = JobScheduleResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *JobScheduleResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewJobScheduleResponseMetaDataType() *JobScheduleResponseMetaDataType {
	jobScheduleResponseMetaDataTypeType := new(JobScheduleResponseMetaDataType)
	return jobScheduleResponseMetaDataTypeType
}

func NewJobScheduleResponseMetaDataTypeWithDefaults() *JobScheduleResponseMetaDataType {
	jobScheduleResponseMetaDataTypeType := new(JobScheduleResponseMetaDataType)
	return jobScheduleResponseMetaDataTypeType
}

type List struct {
	// Extra
	// Extra information for job list
	Extra *ListExtraType `json:"extra,omitempty"`

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

func NewList() *List {
	listType := new(List)
	return listType
}

func NewListWithDefaults() *List {
	listType := new(List)
	return listType
}

// ListExtraType
//
// Extra information for job list
type ListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *ListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewListExtraType() *ListExtraType {
	listExtraTypeType := new(ListExtraType)
	return listExtraTypeType
}

func NewListExtraTypeWithDefaults() *ListExtraType {
	listExtraTypeType := new(ListExtraType)
	return listExtraTypeType
}
