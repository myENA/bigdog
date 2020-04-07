package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMJobErrorObject struct {
	// List
	// Constraints:
	//    - nullable
	List []string `json:"list,omitempty" validate:"omitempty,dive"`

	// Message
	// Constraints:
	//    - nullable
	Message *string `json:"message,omitempty"`

	// MsgKey
	// Constraints:
	//    - nullable
	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMJobErrorObject() *SwitchMJobErrorObject {
	m := new(SwitchMJobErrorObject)
	return m
}

type SwitchMJob struct {
	// Action
	// Action of the job
	// Constraints:
	//    - nullable
	Action *string `json:"action,omitempty"`

	// CreatedTimestamp
	// Created timestamp of the job
	// Constraints:
	//    - nullable
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	// CsvDataMap
	// Constraints:
	//    - nullable
	CsvDataMap *SwitchMJobCsvDataMapType `json:"csvDataMap,omitempty"`

	// DomainId
	// Identifier of the management domain to which the job belong
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// FailureReason
	// Failure reason of the job
	// Constraints:
	//    - nullable
	FailureReason *string `json:"failureReason,omitempty"`

	// Id
	// Identifier of the job
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedTimestamp
	// Modified timestamp of the job
	// Constraints:
	//    - nullable
	ModifiedTimestamp *int `json:"modifiedTimestamp,omitempty"`

	// ScheduleId
	// Schedule Id of the job
	// Constraints:
	//    - nullable
	ScheduleId *string `json:"scheduleId,omitempty"`

	// Status
	// Status of the job
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// StickyNodeId
	// Sticky node Id of the job
	// Constraints:
	//    - nullable
	StickyNodeId *string `json:"stickyNodeId,omitempty"`

	// SwitchGroupLevelOneId
	// Switch group level one Id of the job
	// Constraints:
	//    - nullable
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Switch group level two Id of the job
	// Constraints:
	//    - nullable
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// SwitchId
	// Switch Id of the job
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// TenantId
	// Tenant Id of the job
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// Type
	// Type of the job
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewSwitchMJob() *SwitchMJob {
	m := new(SwitchMJob)
	return m
}

type SwitchMJobCsvDataMapType struct {
	XAdditionalProperties map[string]string `json:"-"`
}

func (t *SwitchMJobCsvDataMapType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]string)
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMJobCsvDataMapType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMJobCsvDataMapType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMJobCsvDataMapType() *SwitchMJobCsvDataMapType {
	m := new(SwitchMJobCsvDataMapType)
	return m
}

type SwitchMJobSchedule struct {
	// CreatedTimestamp
	// Created timestamp of job schedule
	// Constraints:
	//    - nullable
	CreatedTimestamp *int `json:"createdTimestamp,omitempty"`

	// JobId
	// Job Id of job schedule
	// Constraints:
	//    - nullable
	JobId []string `json:"jobId,omitempty" validate:"omitempty,dive"`

	// TriggerValue
	// Trigger value of job schedule
	// Constraints:
	//    - nullable
	TriggerValue *string `json:"triggerValue,omitempty"`
}

func NewSwitchMJobSchedule() *SwitchMJobSchedule {
	m := new(SwitchMJobSchedule)
	return m
}

type SwitchMJobScheduleResponse struct {
	// Data
	// Constraints:
	//    - nullable
	Data *SwitchMJobSchedule `json:"data,omitempty"`

	// Error
	// Constraints:
	//    - nullable
	Error *SwitchMJobErrorObject `json:"error,omitempty"`

	// Extra
	// Extra response of job schedule
	// Constraints:
	//    - nullable
	Extra *SwitchMJobScheduleResponseExtraType `json:"extra,omitempty"`

	// MetaData
	// metaData of job schedule
	// Constraints:
	//    - nullable
	MetaData *SwitchMJobScheduleResponseMetaDataType `json:"metaData,omitempty"`

	// Success
	// Success response of job schedule
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMJobScheduleResponse() *SwitchMJobScheduleResponse {
	m := new(SwitchMJobScheduleResponse)
	return m
}

// SwitchMJobScheduleResponseExtraType
//
// Extra response of job schedule
// Constraints:
//    - nullable
type SwitchMJobScheduleResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMJobScheduleResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMJobScheduleResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMJobScheduleResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMJobScheduleResponseExtraType() *SwitchMJobScheduleResponseExtraType {
	m := new(SwitchMJobScheduleResponseExtraType)
	return m
}

// SwitchMJobScheduleResponseMetaDataType
//
// metaData of job schedule
// Constraints:
//    - nullable
type SwitchMJobScheduleResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMJobScheduleResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMJobScheduleResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMJobScheduleResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMJobScheduleResponseMetaDataType() *SwitchMJobScheduleResponseMetaDataType {
	m := new(SwitchMJobScheduleResponseMetaDataType)
	return m
}

type SwitchMJobList struct {
	// Extra
	// Extra information for job list
	// Constraints:
	//    - nullable
	Extra *SwitchMJobListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first job returned out of the complete job list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more jobs after the currently displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMJob `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// List count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total list count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMJobList() *SwitchMJobList {
	m := new(SwitchMJobList)
	return m
}

// SwitchMJobListExtraType
//
// Extra information for job list
// Constraints:
//    - nullable
type SwitchMJobListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMJobListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMJobListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMJobListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMJobListExtraType() *SwitchMJobListExtraType {
	m := new(SwitchMJobListExtraType)
	return m
}
