package administration

// API Version: v8_0

type ApPatchHistory struct {
	ApFwVersion   *string  `json:"apFwVersion,omitempty"`
	ApModelList   []string `json:"apModelList,omitempty"`
	FileName      *string  `json:"fileName,omitempty"`
	StartDateTime *string  `json:"startDateTime,omitempty"`
}

type ApPatchHistoryList struct {
	FirstIndex *int              `json:"firstIndex,omitempty"`
	HasMore    *bool             `json:"hasMore,omitempty"`
	List       []*ApPatchHistory `json:"list,omitempty"`
	TotalCount *int              `json:"totalCount,omitempty"`
}

type ApPatchInfo struct {
	ApModels  []string `json:"apModels,omitempty"`
	ApVersion *string  `json:"apVersion,omitempty"`
	FileName  *string  `json:"fileName,omitempty"`
	FileSize  *int     `json:"fileSize,omitempty"`
}

type ApPatchStatus struct {
	ClusterOperationProgress *clusterblade.ClusterApPatchOperationProgress `json:"clusterOperationProgress,omitempty"`
}

type ApplicationLogAndStatus struct {
	ApplicationName *string  `json:"applicationName,omitempty"`
	HealthStatus    *string  `json:"healthStatus,omitempty"`
	LogFileNames    []string `json:"logFileNames,omitempty"`
	LogLevel        *string  `json:"logLevel,omitempty"`
	NumOfLogs       *int     `json:"numOfLogs,omitempty"`
}

type ApplicationLogAndStatusList struct {
	FirstIndex *int                       `json:"firstIndex,omitempty"`
	HasMore    *bool                      `json:"hasMore,omitempty"`
	List       []*ApplicationLogAndStatus `json:"list,omitempty"`
	TotalCount *int                       `json:"totalCount,omitempty"`
}

type AutoExportBackup struct {
	EnableAutoExportBackup *bool   `json:"enableAutoExportBackup,omitempty"`
	FtpServer              *string `json:"ftpServer,omitempty"`
}

type BackupFile struct {
	BackupElapsed               *float64 `json:"backupElapsed,omitempty"`
	ControlPlaneSoftwareVersion *string  `json:"controlPlaneSoftwareVersion,omitempty"`
	CreatedBy                   *string  `json:"createdBy,omitempty"`
	CreatedOn                   *float64 `json:"createdOn,omitempty"`
	DataPlaneSoftwareVersion    *string  `json:"dataPlaneSoftwareVersion,omitempty"`
	FileSize                    *float64 `json:"fileSize,omitempty"`
	ID                          *string  `json:"id,omitempty"`
	Md5                         *string  `json:"md5,omitempty"`
	ScgVersion                  *string  `json:"scgVersion,omitempty"`
	Type                        *string  `json:"type,omitempty"`
}

type ClusterBackupList struct {
	FirstIndex *int                    `json:"firstIndex,omitempty"`
	HasMore    *bool                   `json:"hasMore,omitempty"`
	List       []*ClusterBackupSummary `json:"list,omitempty"`
	TotalCount *int                    `json:"totalCount,omitempty"`
}

type ClusterBackupSummary struct {
	CreatedOn *string  `json:"createdOn,omitempty"`
	Filesize  *float64 `json:"filesize,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Version   *string  `json:"version,omitempty"`
}

type ConfigurationBackupList struct {
	FirstIndex *int          `json:"firstIndex,omitempty"`
	HasMore    *bool         `json:"hasMore,omitempty"`
	List       []*BackupFile `json:"list,omitempty"`
	TotalCount *int          `json:"totalCount,omitempty"`
}

type Licenses struct {
	Count       *int    `json:"count,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpireDate  *string `json:"expireDate,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type LicenseServer struct {
	IPAddress *string `json:"ipAddress,omitempty"`
	Port      *int    `json:"port,omitempty"`
	UseCloud  *bool   `json:"useCloud,omitempty"`
}

type LicensesList struct {
	FirstIndex *int        `json:"firstIndex,omitempty"`
	HasMore    *bool       `json:"hasMore,omitempty"`
	List       []*Licenses `json:"list,omitempty"`
	TotalCount *int        `json:"totalCount,omitempty"`
}

type LicensesSummary struct {
	CapacityControlLicenseCount *LicensesSummaryCapacityControlLicenseCountType `json:"capacityControlLicenseCount,omitempty"`
	LicenseTypeDescription      *string                                         `json:"licenseTypeDescription,omitempty"`
}

type LicensesSummaryCapacityControlLicenseCountType struct {
	TotalCount *int `json:"totalCount,omitempty"`
	UsedCount  *int `json:"usedCount,omitempty"`
}

type LicensesSummaryList struct {
	FirstIndex *int               `json:"firstIndex,omitempty"`
	HasMore    *bool              `json:"hasMore,omitempty"`
	List       []*LicensesSummary `json:"list,omitempty"`
	TotalCount *int               `json:"totalCount,omitempty"`
}

type LicensesSyncLogs struct {
	CreateDateTime *string `json:"createDateTime,omitempty"`
	SyncResult     *string `json:"syncResult,omitempty"`
}

type LicensesSyncLogsList struct {
	FirstIndex *int                `json:"firstIndex,omitempty"`
	HasMore    *bool               `json:"hasMore,omitempty"`
	List       []*LicensesSyncLogs `json:"list,omitempty"`
	TotalCount *int                `json:"totalCount,omitempty"`
}

type ModfiyLicenseServer struct {
	IPAddress *string `json:"ipAddress,omitempty"`
	Port      *int    `json:"port,omitempty"`
	UseCloud  *bool   `json:"useCloud,omitempty"`
}

type ModifyAutoExportBackup struct {
	EnableAutoExportBackup *bool   `json:"enableAutoExportBackup,omitempty"`
	FtpServer              *string `json:"ftpServer,omitempty"`
}

type ModifyLogLevel struct {
	ApplicationName *string `json:"applicationName,omitempty"`
	LogLevel        *string `json:"logLevel,omitempty"`
}

type ModifyScheduleBackup struct {
	DateOfMonth          *int    `json:"dateOfMonth,omitempty"`
	DayOfWeek            *string `json:"dayOfWeek,omitempty"`
	EnableScheduleBackup *bool   `json:"enableScheduleBackup,omitempty"`
	Hour                 *int    `json:"hour,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	Minute               *int    `json:"minute,omitempty"`
}

type ScheduleBackup struct {
	DateOfMonth          *int    `json:"dateOfMonth,omitempty"`
	DayOfWeek            *string `json:"dayOfWeek,omitempty"`
	EnableScheduleBackup *bool   `json:"enableScheduleBackup,omitempty"`
	Hour                 *int    `json:"hour,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	Minute               *int    `json:"minute,omitempty"`
}

type UpgradeHistoryList struct {
	FirstIndex *int                     `json:"firstIndex,omitempty"`
	HasMore    *bool                    `json:"hasMore,omitempty"`
	List       []*UpgradeHistorySummary `json:"list,omitempty"`
	TotalCount *int                     `json:"totalCount,omitempty"`
}

type UpgradeHistorySummary struct {
	ApFwVersion    *string  `json:"apFwVersion,omitempty"`
	CbVersion      *string  `json:"cbVersion,omitempty"`
	DpVersion      *string  `json:"dpVersion,omitempty"`
	ElapsedSeconds *float64 `json:"elapsedSeconds,omitempty"`
	FileName       *string  `json:"fileName,omitempty"`
	OldApFwVersion *string  `json:"oldApFwVersion,omitempty"`
	OldCbVersion   *string  `json:"oldCbVersion,omitempty"`
	OldDpVersion   *string  `json:"oldDpVersion,omitempty"`
	OldVersion     *string  `json:"oldVersion,omitempty"`
	StartTime      *string  `json:"startTime,omitempty"`
	Version        *string  `json:"version,omitempty"`
}

type UpgradePatchInfo struct {
	ClusterOperationProgress *clusterblade.ClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
	UploadPatchInfo          *clusterblade.UploadPatchInfo          `json:"uploadPatchInfo,omitempty"`
}

type UpgradeStatus struct {
	ClusterOperationProgress *clusterblade.ClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
}

type ZdAP struct {
	Connected *string `json:"connected,omitempty"`
	Mac       *string `json:"mac,omitempty"`
}

type ZdAPList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*ZdAP              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type ZdImport struct {
	ApMacList []string `json:"apMacList,omitempty"`
	IP        *string  `json:"ip,omitempty"`
	Password  *string  `json:"password,omitempty"`
	User      *string  `json:"user,omitempty"`
}

type ZdImportStatus struct {
	Details  *string `json:"details,omitempty"`
	Message  *string `json:"message,omitempty"`
	Progress *int    `json:"progress,omitempty"`
	State    *string `json:"state,omitempty"`
}
