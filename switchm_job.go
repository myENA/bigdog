package ruckus

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMSwitchJobErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMSwitchJobErrorObject() *SwitchMSwitchJobErrorObject {
	m := new(SwitchMSwitchJobErrorObject)
	return m
}

type SwitchMSwitchJobJob struct {
	// Action
	// Action of the job
	Action *string `json:"action,omitempty"`

	// CreatedTimestamp
	// Created timestamp of the job
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	CsvDataMap *SwitchMSwitchJobJob `json:"csvDataMap,omitempty"`

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

func NewSwitchMSwitchJobJob() *SwitchMSwitchJobJob {
	m := new(SwitchMSwitchJobJob)
	return m
}

type SwitchMSwitchJobJobCsvDataMapType struct {
	XAdditionalProperties map[string]string `json:"-"`
}

func (t *SwitchMSwitchJobJobCsvDataMapType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]string)
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchJobJobCsvDataMapType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchJobJobCsvDataMapType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchJobJobCsvDataMapType() *SwitchMSwitchJobJobCsvDataMapType {
	m := new(SwitchMSwitchJobJobCsvDataMapType)
	return m
}

type SwitchMSwitchJobJobSchedule struct {
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

func NewSwitchMSwitchJobJobSchedule() *SwitchMSwitchJobJobSchedule {
	m := new(SwitchMSwitchJobJobSchedule)
	return m
}

type SwitchMSwitchJobJobScheduleResponse struct {
	Data *SwitchMSwitchJobJobScheduleResponse `json:"data,omitempty"`

	Error *SwitchMSwitchJobJobScheduleResponse `json:"error,omitempty"`

	// Extra
	// Extra response of job schedule
	Extra *SwitchMSwitchJobJobScheduleResponse `json:"extra,omitempty"`

	// MetaData
	// metaData of job schedule
	MetaData *SwitchMSwitchJobJobScheduleResponse `json:"metaData,omitempty"`

	// Success
	// Success response of job schedule
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchJobJobScheduleResponse() *SwitchMSwitchJobJobScheduleResponse {
	m := new(SwitchMSwitchJobJobScheduleResponse)
	return m
}

// SwitchMSwitchJobJobScheduleResponseExtraType
//
// Extra response of job schedule
type SwitchMSwitchJobJobScheduleResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchJobJobScheduleResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchJobJobScheduleResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchJobJobScheduleResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchJobJobScheduleResponseExtraType() *SwitchMSwitchJobJobScheduleResponseExtraType {
	m := new(SwitchMSwitchJobJobScheduleResponseExtraType)
	return m
}

// SwitchMSwitchJobJobScheduleResponseMetaDataType
//
// metaData of job schedule
type SwitchMSwitchJobJobScheduleResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchJobJobScheduleResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchJobJobScheduleResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchJobJobScheduleResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchJobJobScheduleResponseMetaDataType() *SwitchMSwitchJobJobScheduleResponseMetaDataType {
	m := new(SwitchMSwitchJobJobScheduleResponseMetaDataType)
	return m
}

type SwitchMSwitchJobList struct {
	// Extra
	// Extra information for job list
	Extra *SwitchMSwitchJobList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first job returned out of the complete job list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more jobs after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchJobList `json:"list,omitempty"`

	// RawDataTotalCount
	// List count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchJobList() *SwitchMSwitchJobList {
	m := new(SwitchMSwitchJobList)
	return m
}

// SwitchMSwitchJobListExtraType
//
// Extra information for job list
type SwitchMSwitchJobListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchJobListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchJobListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchJobListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchJobListExtraType() *SwitchMSwitchJobListExtraType {
	m := new(SwitchMSwitchJobListExtraType)
	return m
}
