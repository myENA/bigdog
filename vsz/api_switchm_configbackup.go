package vsz

// API Version: v9_0

type SwitchMConfigbackupBackupIds []string

func MakeSwitchMConfigbackupBackupIds() SwitchMConfigbackupBackupIds {
	m := make(SwitchMConfigbackupBackupIds, 0)
	return m
}

type SwitchMConfigbackupCfgbk struct {
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

func NewSwitchMConfigbackupCfgbk() *SwitchMConfigbackupCfgbk {
	m := new(SwitchMConfigbackupCfgbk)
	return m
}

type SwitchMConfigbackupContent struct {
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

func NewSwitchMConfigbackupContent() *SwitchMConfigbackupContent {
	m := new(SwitchMConfigbackupContent)
	return m
}

type SwitchMConfigbackupDiff struct {
	ConfigBackup1 *SwitchMConfigbackupContent `json:"configBackup1,omitempty"`

	ConfigBackup2 *SwitchMConfigbackupContent `json:"configBackup2,omitempty"`
}

func NewSwitchMConfigbackupDiff() *SwitchMConfigbackupDiff {
	m := new(SwitchMConfigbackupDiff)
	return m
}

type SwitchMConfigbackupDiffInput struct {
	// ConfigBackupId1
	// The first config backup id of diff input
	ConfigBackupId1 *string `json:"configBackupId1,omitempty"`

	// ConfigBackupId2
	// The second config backup id of diff input
	ConfigBackupId2 *string `json:"configBackupId2,omitempty"`
}

func NewSwitchMConfigbackupDiffInput() *SwitchMConfigbackupDiffInput {
	m := new(SwitchMConfigbackupDiffInput)
	return m
}

type SwitchMConfigbackupCreateBackupResultList struct {
	Extra *SwitchMCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMConfigbackupCreateBackupResultListType `json:"list,omitempty"`

	// RawDataTotalCount
	// Total ConfigBackupInfo count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current ConfigBackupInfo count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMConfigbackupCreateBackupResultList() *SwitchMConfigbackupCreateBackupResultList {
	m := new(SwitchMConfigbackupCreateBackupResultList)
	return m
}

type SwitchMConfigbackupCreateBackupResultListType struct {
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

func NewSwitchMConfigbackupCreateBackupResultListType() *SwitchMConfigbackupCreateBackupResultListType {
	m := new(SwitchMConfigbackupCreateBackupResultListType)
	return m
}

type SwitchMConfigbackupList struct {
	Extra *SwitchMCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ConfigBackup returned out of the complete ConfigBackup list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more ConfigBackup after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMConfigbackupCfgbk `json:"list,omitempty"`

	// RawDataTotalCount
	// ConfigBackup count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total ConfigBackup count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMConfigbackupList() *SwitchMConfigbackupList {
	m := new(SwitchMConfigbackupList)
	return m
}

type SwitchMConfigbackupSwitchIds []string

func MakeSwitchMConfigbackupSwitchIds() SwitchMConfigbackupSwitchIds {
	m := make(SwitchMConfigbackupSwitchIds, 0)
	return m
}
