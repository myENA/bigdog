package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// SwitchMJobErrorObject
//
// Definition: job_errorObject
type SwitchMJobErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMJobErrorObject() *SwitchMJobErrorObject {
	m := new(SwitchMJobErrorObject)
	return m
}

// SwitchMJob
//
// Definition: job_job
type SwitchMJob struct {
	// Action
	// Action of the job
	Action *string `json:"action,omitempty"`

	// CreatedTimestamp
	// Created timestamp of the job
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	CsvDataMap interface{} `json:"csvDataMap,omitempty"`

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

type SwitchMJobAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMJob
}

func newSwitchMJobAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMJobAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMJobAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMJob)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMJob() *SwitchMJob {
	m := new(SwitchMJob)
	return m
}

// SwitchMJobSchedule
//
// Definition: job_jobSchedule
type SwitchMJobSchedule struct {
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

func NewSwitchMJobSchedule() *SwitchMJobSchedule {
	m := new(SwitchMJobSchedule)
	return m
}

// SwitchMJobScheduleResponse
//
// Definition: job_jobScheduleResponse
type SwitchMJobScheduleResponse struct {
	Data *SwitchMJobSchedule `json:"data,omitempty"`

	Error *SwitchMJobErrorObject `json:"error,omitempty"`

	// Extra
	// Extra response of job schedule
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// metaData of job schedule
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Success response of job schedule
	Success *bool `json:"success,omitempty"`
}

type SwitchMJobScheduleResponseAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMJobScheduleResponse
}

func newSwitchMJobScheduleResponseAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMJobScheduleResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMJobScheduleResponseAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMJobScheduleResponse)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMJobScheduleResponse() *SwitchMJobScheduleResponse {
	m := new(SwitchMJobScheduleResponse)
	return m
}

// SwitchMJobList
//
// Definition: job_list
type SwitchMJobList struct {
	// Extra
	// Extra information for job list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first job returned out of the complete job list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more jobs after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMJob `json:"list,omitempty"`

	// RawDataTotalCount
	// List count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMJobListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMJobList
}

func newSwitchMJobListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMJobListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMJobListAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMJobList)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMJobList() *SwitchMJobList {
	m := new(SwitchMJobList)
	return m
}
