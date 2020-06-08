package ruckus

// API Version: v9_0

type SwitchMSwitchConfigurationBackupBackupIds []string

func MakeSwitchMSwitchConfigurationBackupBackupIds() SwitchMSwitchConfigurationBackupBackupIds {
	m := make(SwitchMSwitchConfigurationBackupBackupIds, 0)
	return m
}

type SwitchMSwitchConfigurationBackupConfigurationBackup struct {
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

func NewSwitchMSwitchConfigurationBackupConfigurationBackup() *SwitchMSwitchConfigurationBackupConfigurationBackup {
	m := new(SwitchMSwitchConfigurationBackupConfigurationBackup)
	return m
}

type SwitchMSwitchConfigurationBackupConfigBackupContent struct {
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

func NewSwitchMSwitchConfigurationBackupConfigBackupContent() *SwitchMSwitchConfigurationBackupConfigBackupContent {
	m := new(SwitchMSwitchConfigurationBackupConfigBackupContent)
	return m
}

type SwitchMSwitchConfigurationBackupConfigBackupDiff struct {
	ConfigBackup1 *SwitchMSwitchConfigurationBackupConfigBackupDiff `json:"configBackup1,omitempty"`

	ConfigBackup2 *SwitchMSwitchConfigurationBackupConfigBackupDiff `json:"configBackup2,omitempty"`
}

func NewSwitchMSwitchConfigurationBackupConfigBackupDiff() *SwitchMSwitchConfigurationBackupConfigBackupDiff {
	m := new(SwitchMSwitchConfigurationBackupConfigBackupDiff)
	return m
}

type SwitchMSwitchConfigurationBackupConfigBackupDiffInput struct {
	// ConfigBackupId1
	// The first config backup id of diff input
	ConfigBackupId1 *string `json:"configBackupId1,omitempty"`

	// ConfigBackupId2
	// The second config backup id of diff input
	ConfigBackupId2 *string `json:"configBackupId2,omitempty"`
}

func NewSwitchMSwitchConfigurationBackupConfigBackupDiffInput() *SwitchMSwitchConfigurationBackupConfigBackupDiffInput {
	m := new(SwitchMSwitchConfigurationBackupConfigBackupDiffInput)
	return m
}

type SwitchMSwitchConfigurationBackupCreateBackupResultList struct {
	Extra *SwitchMSwitchConfigurationBackupCreateBackupResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchConfigurationBackupCreateBackupResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total ConfigBackupInfo count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current ConfigBackupInfo count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchConfigurationBackupCreateBackupResultList() *SwitchMSwitchConfigurationBackupCreateBackupResultList {
	m := new(SwitchMSwitchConfigurationBackupCreateBackupResultList)
	return m
}

type SwitchMSwitchConfigurationBackupCreateBackupResultListType struct {
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

func NewSwitchMSwitchConfigurationBackupCreateBackupResultListType() *SwitchMSwitchConfigurationBackupCreateBackupResultListType {
	m := new(SwitchMSwitchConfigurationBackupCreateBackupResultListType)
	return m
}

type SwitchMSwitchConfigurationBackupList struct {
	Extra *SwitchMSwitchConfigurationBackupList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ConfigBackup returned out of the complete ConfigBackup list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more ConfigBackup after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchConfigurationBackupList `json:"list,omitempty"`

	// RawDataTotalCount
	// ConfigBackup count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total ConfigBackup count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchConfigurationBackupList() *SwitchMSwitchConfigurationBackupList {
	m := new(SwitchMSwitchConfigurationBackupList)
	return m
}

type SwitchMSwitchConfigurationBackupSwitchIds []string

func MakeSwitchMSwitchConfigurationBackupSwitchIds() SwitchMSwitchConfigurationBackupSwitchIds {
	m := make(SwitchMSwitchConfigurationBackupSwitchIds, 0)
	return m
}
