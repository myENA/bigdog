package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// SwitchMConfigurationBackupBackupIds
//
// Definition: configBackup_backupIds
type SwitchMConfigurationBackupBackupIds []string

func MakeSwitchMConfigurationBackupBackupIds() SwitchMConfigurationBackupBackupIds {
	m := make(SwitchMConfigurationBackupBackupIds, 0)
	return m
}

// SwitchMConfigurationBackup
//
// Definition: configBackup_cfgbk
type SwitchMConfigurationBackup struct {
	// FailureReason
	// Failure reason of the config backup and config restore
	FailureReason *string `json:"failureReason,omitempty"`

	// Id
	// the identifier of the config backup
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the config backup
	Name *string `json:"name,omitempty"`

	// RestoreStatus
	// Status of config restore
	RestoreStatus *string `json:"restoreStatus,omitempty"`

	// RestoreTimestamp
	// the timestamp of the config restore
	RestoreTimestamp *int `json:"restoreTimestamp,omitempty"`

	// Status
	// the status of the config backup
	Status *string `json:"status,omitempty"`

	// SwitchName
	// Switch Name of the config backup
	SwitchName *string `json:"switchName,omitempty"`

	// Timestamp
	// the timestamp of the config backup
	Timestamp *int `json:"timestamp,omitempty"`

	// Type
	// Scheduled or Manual
	Type *string `json:"type,omitempty"`
}

func NewSwitchMConfigurationBackup() *SwitchMConfigurationBackup {
	m := new(SwitchMConfigurationBackup)
	return m
}

// SwitchMConfigurationBackupConfigBackupContent
//
// Definition: configBackup_configBackupContent
type SwitchMConfigurationBackupConfigBackupContent struct {
	// Config
	// Content of the ConfigBackup
	Config *string `json:"config,omitempty"`

	// Id
	// the identifier of the ConfigBackup Content
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the ConfigBackup Content
	Name *string `json:"name,omitempty"`
}

func NewSwitchMConfigurationBackupConfigBackupContent() *SwitchMConfigurationBackupConfigBackupContent {
	m := new(SwitchMConfigurationBackupConfigBackupContent)
	return m
}

// SwitchMConfigurationBackupConfigBackupDiff
//
// Definition: configBackup_configBackupDiff
type SwitchMConfigurationBackupConfigBackupDiff struct {
	ConfigBackup1 *SwitchMConfigurationBackupConfigBackupContent `json:"configBackup1,omitempty"`

	ConfigBackup2 *SwitchMConfigurationBackupConfigBackupContent `json:"configBackup2,omitempty"`
}

type SwitchMConfigurationBackupConfigBackupDiffAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMConfigurationBackupConfigBackupDiff
}

func newSwitchMConfigurationBackupConfigBackupDiffAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMConfigurationBackupConfigBackupDiffAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMConfigurationBackupConfigBackupDiffAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMConfigurationBackupConfigBackupDiff)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMConfigurationBackupConfigBackupDiff() *SwitchMConfigurationBackupConfigBackupDiff {
	m := new(SwitchMConfigurationBackupConfigBackupDiff)
	return m
}

// SwitchMConfigurationBackupConfigBackupDiffInput
//
// Definition: configBackup_configBackupDiffInput
type SwitchMConfigurationBackupConfigBackupDiffInput struct {
	// ConfigBackupId1
	// The first config backup id of diff input
	ConfigBackupId1 *string `json:"configBackupId1,omitempty"`

	// ConfigBackupId2
	// The second config backup id of diff input
	ConfigBackupId2 *string `json:"configBackupId2,omitempty"`
}

func NewSwitchMConfigurationBackupConfigBackupDiffInput() *SwitchMConfigurationBackupConfigBackupDiffInput {
	m := new(SwitchMConfigurationBackupConfigBackupDiffInput)
	return m
}

// SwitchMConfigurationBackupCreateBackupResultList
//
// Definition: configBackup_createBackupResultList
type SwitchMConfigurationBackupCreateBackupResultList struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMConfigurationBackupCreateBackupResultListType `json:"list,omitempty"`

	// RawDataTotalCount
	// Total ConfigBackupInfo count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current ConfigBackupInfo count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMConfigurationBackupCreateBackupResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMConfigurationBackupCreateBackupResultList
}

func newSwitchMConfigurationBackupCreateBackupResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMConfigurationBackupCreateBackupResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMConfigurationBackupCreateBackupResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMConfigurationBackupCreateBackupResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMConfigurationBackupCreateBackupResultList() *SwitchMConfigurationBackupCreateBackupResultList {
	m := new(SwitchMConfigurationBackupCreateBackupResultList)
	return m
}

// SwitchMConfigurationBackupCreateBackupResultListType
//
// Definition: configBackup_createBackupResultListType
type SwitchMConfigurationBackupCreateBackupResultListType struct {
	// ConfigBackupId
	// Identifier of config backup
	ConfigBackupId *string `json:"configBackupId,omitempty"`

	// ErrorMessage
	// Error message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// SwitchId
	// Identifier of switch
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMConfigurationBackupCreateBackupResultListType() *SwitchMConfigurationBackupCreateBackupResultListType {
	m := new(SwitchMConfigurationBackupCreateBackupResultListType)
	return m
}

// SwitchMConfigurationBackupList
//
// Definition: configBackup_list
type SwitchMConfigurationBackupList struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ConfigBackup returned out of the complete ConfigBackup list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more ConfigBackup after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMConfigurationBackup `json:"list,omitempty"`

	// RawDataTotalCount
	// ConfigBackup count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total ConfigBackup count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMConfigurationBackupListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMConfigurationBackupList
}

func newSwitchMConfigurationBackupListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMConfigurationBackupListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMConfigurationBackupListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMConfigurationBackupList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMConfigurationBackupList() *SwitchMConfigurationBackupList {
	m := new(SwitchMConfigurationBackupList)
	return m
}

// SwitchMConfigurationBackupSwitchIds
//
// Definition: configBackup_switchIds
type SwitchMConfigurationBackupSwitchIds []string

func MakeSwitchMConfigurationBackupSwitchIds() SwitchMConfigurationBackupSwitchIds {
	m := make(SwitchMConfigurationBackupSwitchIds, 0)
	return m
}
