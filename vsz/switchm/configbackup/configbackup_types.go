package configbackup

// API Version: v8_0

type Cfgbk struct {
	FailureReason    *string `json:"failureReason,omitempty"`
	ID               *string `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	RestoreStatus    *string `json:"restoreStatus,omitempty"`
	RestoreTimestamp *int    `json:"restoreTimestamp,omitempty"`
	Status           *string `json:"status,omitempty"`
	SwitchName       *string `json:"switchName,omitempty"`
	Timestamp        *int    `json:"timestamp,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type ConfigBackupContent struct {
	Config *string `json:"config,omitempty"`
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

type ConfigBackupDiff struct {
	ConfigBackup1 *ConfigBackup1 `json:"configBackup1,omitempty"`
	ConfigBackup2 *ConfigBackup2 `json:"configBackup2,omitempty"`
}

type ConfigBackupDiffInput struct {
	ConfigBackupID1 *string `json:"configBackupId1,omitempty"`
	ConfigBackupID2 *string `json:"configBackupId2,omitempty"`
}

type CreateBackupResultList struct {
	Extra             *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex        *int                 `json:"firstIndex,omitempty"`
	HasMore           *bool                `json:"hasMore,omitempty"`
	List              []*List              `json:"list,omitempty"`
	RawDataTotalCount *int                 `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int                 `json:"totalCount,omitempty"`
}

type CreateBackupResultListType struct {
	ConfigBackupID *string `json:"configBackupId,omitempty"`
	SwitchID       *string `json:"switchId,omitempty"`
}

type List struct {
	Extra             *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex        *int                 `json:"firstIndex,omitempty"`
	HasMore           *bool                `json:"hasMore,omitempty"`
	List              []*List              `json:"list,omitempty"`
	RawDataTotalCount *int                 `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int                 `json:"totalCount,omitempty"`
}

type QueryCriteria struct {
	ConfigBackupQueryCriteria *string `json:"configBackup_queryCriteria,omitempty"`
}
