package configbackup

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
)

type BackupIds []string

type Cfgbk struct {
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

type ConfigBackupContent struct {
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

type ConfigBackupDiff struct {
	ConfigBackup1 *ConfigBackupContent `json:"configBackup1,omitempty"`

	ConfigBackup2 *ConfigBackupContent `json:"configBackup2,omitempty"`
}

type ConfigBackupDiffInput struct {
	// ConfigBackupId1
	// The first config backup id of diff input
	ConfigBackupId1 *string `json:"configBackupId1,omitempty"`

	// ConfigBackupId2
	// The second config backup id of diff input
	ConfigBackupId2 *string `json:"configBackupId2,omitempty"`
}

type CreateBackupResultList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*CreateBackupResultListType `json:"list,omitempty"`

	// RawDataTotalCount
	// Total ConfigBackupInfo count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Current ConfigBackupInfo count
	TotalCount *int `json:"totalCount,omitempty"`
}

type CreateBackupResultListType struct {
	// ConfigBackupId
	// Identifier of config backup
	ConfigBackupId *string `json:"configBackupId,omitempty"`

	// SwitchId
	// Identifier of switch
	SwitchId *string `json:"switchId,omitempty"`
}

type List struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first ConfigBackup returned out of the complete ConfigBackup list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more ConfigBackup after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Cfgbk `json:"list,omitempty"`

	// RawDataTotalCount
	// ConfigBackup count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total ConfigBackup count
	TotalCount *int `json:"totalCount,omitempty"`
}

type QueryCriteria struct {
	ConfigBackupqueryCriteria *string `json:"configBackup_queryCriteria,omitempty"`
}

type SwitchIds []string
