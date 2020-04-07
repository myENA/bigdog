package vsz

// API Version: v9_0

type SwitchMConfigBackupBackupIds []string

func MakeSwitchMConfigBackupBackupIds() SwitchMConfigBackupBackupIds {
	m := make(SwitchMConfigBackupBackupIds, 0)
	return m
}

type SwitchMConfigBackupCfgbk struct {
	// FailureReason
	// Failure reason of the config backup and config restore
	// Constraints:
	//    - nullable
	FailureReason *string `json:"failureReason,omitempty"`

	// Id
	// the identifier of the config backup
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the config backup
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// RestoreStatus
	// Status of config restore
	// Constraints:
	//    - nullable
	RestoreStatus *string `json:"restoreStatus,omitempty"`

	// RestoreTimestamp
	// the timestamp of the config restore
	// Constraints:
	//    - nullable
	RestoreTimestamp *int `json:"restoreTimestamp,omitempty"`

	// Status
	// the status of the config backup
	// Constraints:
	//    - nullable
	Status *string `json:"status,omitempty"`

	// SwitchName
	// Switch Name of the config backup
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// Timestamp
	// the timestamp of the config backup
	// Constraints:
	//    - nullable
	Timestamp *int `json:"timestamp,omitempty"`

	// Type
	// Scheduled or Manual
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewSwitchMConfigBackupCfgbk() *SwitchMConfigBackupCfgbk {
	m := new(SwitchMConfigBackupCfgbk)
	return m
}

type SwitchMConfigBackupContent struct {
	// Config
	// Content of the ConfigBackup
	// Constraints:
	//    - nullable
	Config *string `json:"config,omitempty"`

	// Id
	// the identifier of the ConfigBackup Content
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Name of the ConfigBackup Content
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewSwitchMConfigBackupContent() *SwitchMConfigBackupContent {
	m := new(SwitchMConfigBackupContent)
	return m
}

type SwitchMConfigBackupDiff struct {
	// ConfigBackup1
	// Constraints:
	//    - nullable
	ConfigBackup1 *SwitchMConfigBackupContent `json:"configBackup1,omitempty"`

	// ConfigBackup2
	// Constraints:
	//    - nullable
	ConfigBackup2 *SwitchMConfigBackupContent `json:"configBackup2,omitempty"`
}

func NewSwitchMConfigBackupDiff() *SwitchMConfigBackupDiff {
	m := new(SwitchMConfigBackupDiff)
	return m
}

type SwitchMConfigBackupDiffInput struct {
	// ConfigBackupId1
	// The first config backup id of diff input
	// Constraints:
	//    - nullable
	ConfigBackupId1 *string `json:"configBackupId1,omitempty"`

	// ConfigBackupId2
	// The second config backup id of diff input
	// Constraints:
	//    - nullable
	ConfigBackupId2 *string `json:"configBackupId2,omitempty"`
}

func NewSwitchMConfigBackupDiffInput() *SwitchMConfigBackupDiffInput {
	m := new(SwitchMConfigBackupDiffInput)
	return m
}

type SwitchMConfigBackupCreateBackupResultList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *SwitchMCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMConfigBackupCreateBackupResultListType `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Total ConfigBackupInfo count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current ConfigBackupInfo count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMConfigBackupCreateBackupResultList() *SwitchMConfigBackupCreateBackupResultList {
	m := new(SwitchMConfigBackupCreateBackupResultList)
	return m
}

type SwitchMConfigBackupCreateBackupResultListType struct {
	// ConfigBackupId
	// Identifier of config backup
	// Constraints:
	//    - nullable
	ConfigBackupId *string `json:"configBackupId,omitempty"`

	// ErrorMessage
	// Error message
	// Constraints:
	//    - nullable
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// SwitchId
	// Identifier of switch
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMConfigBackupCreateBackupResultListType() *SwitchMConfigBackupCreateBackupResultListType {
	m := new(SwitchMConfigBackupCreateBackupResultListType)
	return m
}

type SwitchMConfigBackupList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *SwitchMCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ConfigBackup returned out of the complete ConfigBackup list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more ConfigBackup after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMConfigBackupCfgbk `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// ConfigBackup count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total ConfigBackup count
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMConfigBackupList() *SwitchMConfigBackupList {
	m := new(SwitchMConfigBackupList)
	return m
}

type SwitchMConfigBackupSwitchIds []string

func MakeSwitchMConfigBackupSwitchIds() SwitchMConfigBackupSwitchIds {
	m := make(SwitchMConfigBackupSwitchIds, 0)
	return m
}
