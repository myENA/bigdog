package administration

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/clusterblade"
	"github.com/myENA/ruckus-client/vsz/wsg/common"
)

type ApPatchHistory struct {
	// ApFwVersion
	// apFwVersion of the AP Patch history
	ApFwVersion *string `json:"apFwVersion,omitempty"`

	// ApModelList
	// AP Models of the AP Patch history
	ApModelList []string `json:"apModelList,omitempty"`

	// FileName
	// file name of the AP Patch history
	FileName *string `json:"fileName,omitempty"`

	// StartDateTime
	// startDateTime of the AP Patch history
	StartDateTime *string `json:"startDateTime,omitempty"`
}

type ApPatchHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApPatchHistory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ApPatchInfo struct {
	// ApModels
	// AP Models of the upload file
	ApModels []string `json:"apModels,omitempty"`

	// ApVersion
	// ApFwVersion of the upload file
	ApVersion *string `json:"apVersion,omitempty"`

	// FileName
	// file name of the upload file
	FileName *string `json:"fileName,omitempty"`

	// FileSize
	// file size(Byte) of the upload file
	FileSize *int `json:"fileSize,omitempty"`
}

type ApPatchStatus struct {
	ClusterOperationProgress *clusterblade.ClusterApPatchOperationProgress `json:"clusterOperationProgress,omitempty"`
}

type ApplicationLogAndStatus struct {
	// ApplicationName
	// Application name
	ApplicationName *string `json:"applicationName,omitempty"`

	// HealthStatus
	// Health status
	HealthStatus *string `json:"healthStatus,omitempty"`

	// LogFileNames
	// List of log file name
	LogFileNames []string `json:"logFileNames,omitempty"`

	// LogLevel
	// Log level
	LogLevel *string `json:"logLevel,omitempty"`

	// NumOfLogs
	// # of Logs
	NumOfLogs *int `json:"numOfLogs,omitempty"`
}

type ApplicationLogAndStatusList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApplicationLogAndStatus `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type AutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// FTP server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

type BackupFile struct {
	// BackupElapsed
	// backup elapsed of the configuration backup file
	BackupElapsed *float64 `json:"backupElapsed,omitempty"`

	// ControlPlaneSoftwareVersion
	// control plane software version of the configuration backup file
	ControlPlaneSoftwareVersion *string `json:"controlPlaneSoftwareVersion,omitempty"`

	// CreatedBy
	// creator of the configuration backup file.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedOn
	// the create time of the configuration backup file.
	CreatedOn *float64 `json:"createdOn,omitempty"`

	// DataPlaneSoftwareVersion
	// data plane software version of the configuration backup file
	DataPlaneSoftwareVersion *string `json:"dataPlaneSoftwareVersion,omitempty"`

	// FileSize
	// file size of the backup file
	FileSize *float64 `json:"fileSize,omitempty"`

	// ID
	// Identifier of system configuration backup file.
	ID *string `json:"id,omitempty"`

	// Md5
	// file md5 of the backup file
	Md5 *string `json:"md5,omitempty"`

	// ScgVersion
	// SCG version of the configuration backup file.
	ScgVersion *string `json:"scgVersion,omitempty"`

	// Type
	// type of the configuration backup file
	Type *string `json:"type,omitempty"`
}

type ClusterBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ClusterBackupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ClusterBackupSummary struct {
	// CreatedOn
	// Created date and time of the cluster backup file
	CreatedOn *string `json:"createdOn,omitempty"`

	// Filesize
	// filesize of the cluster backup file.
	Filesize *float64 `json:"filesize,omitempty"`

	// ID
	// Identifier of cluster backup file.
	ID *string `json:"id,omitempty"`

	// Version
	// the patch version of the cluster backup file.
	Version *string `json:"version,omitempty"`
}

type ConfigurationBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BackupFile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type Licenses struct {
	// Count
	// number of licenses
	Count *int `json:"count,omitempty"`

	// CreateTime
	// license effective date
	CreateTime *string `json:"createTime,omitempty"`

	// Description
	// license description
	Description *string `json:"description,omitempty"`

	// ExpireDate
	// license expiry date
	ExpireDate *string `json:"expireDate,omitempty"`

	// Name
	// license name
	Name *string `json:"name,omitempty"`
}

type LicenseServer struct {
	// IPAddress
	// local license server IP address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Port
	// local license server port
	Port *int `json:"port,omitempty"`

	// UseCloud
	// use cloud license server
	UseCloud *bool `json:"useCloud,omitempty"`
}

type LicensesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Licenses `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type LicensesSummary struct {
	CapacityControlLicenseCount *LicensesSummaryCapacityControlLicenseCountType `json:"capacityControlLicenseCount,omitempty"`

	// LicenseTypeDescription
	// license type description
	LicenseTypeDescription *string `json:"licenseTypeDescription,omitempty"`
}

type LicensesSummaryCapacityControlLicenseCountType struct {
	// TotalCount
	// total count of licenses
	TotalCount *int `json:"totalCount,omitempty"`

	// UsedCount
	// consumed count of licenses
	UsedCount *int `json:"usedCount,omitempty"`
}

type LicensesSummaryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LicensesSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type LicensesSyncLogs struct {
	// CreateDateTime
	// license sync log's create time
	CreateDateTime *string `json:"createDateTime,omitempty"`

	// SyncResult
	// sync license result
	SyncResult *string `json:"syncResult,omitempty"`
}

type LicensesSyncLogsList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LicensesSyncLogs `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ModfiyLicenseServer struct {
	IPAddress *string `json:"ipAddress,omitempty"`

	Port *int `json:"port,omitempty"`

	UseCloud *bool `json:"useCloud,omitempty"`
}

type ModifyAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// ftp server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

type ModifyLogLevel struct {
	// ApplicationName
	// Application name.
	ApplicationName *string `json:"applicationName,omitempty"`

	// LogLevel
	// Log level.
	LogLevel *string `json:"logLevel,omitempty"`
}

type ModifyScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	Interval *string `json:"interval,omitempty"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type ScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	Interval *string `json:"interval,omitempty"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type UpgradeHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UpgradeHistorySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type UpgradeHistorySummary struct {
	// ApFwVersion
	// apFwVersion of the upgrade history
	ApFwVersion *string `json:"apFwVersion,omitempty"`

	// CbVersion
	// cbVersion of the upgrade history
	CbVersion *string `json:"cbVersion,omitempty"`

	// DpVersion
	// dpVersion of the upgrade history
	DpVersion *string `json:"dpVersion,omitempty"`

	// ElapsedSeconds
	// elapsedSeconds of the upgrade history
	ElapsedSeconds *float64 `json:"elapsedSeconds,omitempty"`

	// FileName
	// fileName of the upgrade history
	FileName *string `json:"fileName,omitempty"`

	// OldApFwVersion
	// oldApFwVersion of the upgrade history
	OldApFwVersion *string `json:"oldApFwVersion,omitempty"`

	// OldCbVersion
	// oldCbVersion of the upgrade history
	OldCbVersion *string `json:"oldCbVersion,omitempty"`

	// OldDpVersion
	// oldDpVersion of the upgrade history
	OldDpVersion *string `json:"oldDpVersion,omitempty"`

	// OldVersion
	// oldVersion of the upgrade history
	OldVersion *string `json:"oldVersion,omitempty"`

	// StartTime
	// startTime of the upgrade history
	StartTime *string `json:"startTime,omitempty"`

	// Version
	// version of the upgrade history
	Version *string `json:"version,omitempty"`
}

type UpgradePatchInfo struct {
	ClusterOperationProgress *clusterblade.ClusterOperationProgress `json:"clusterOperationProgress,omitempty"`

	UploadPatchInfo *clusterblade.UploadPatchInfo `json:"uploadPatchInfo,omitempty"`
}

type UpgradeStatus struct {
	ClusterOperationProgress *clusterblade.ClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
}

type ZdAP struct {
	// Connected
	// AP Conntected
	Connected *string `json:"connected,omitempty"`

	// Mac
	// AP MAC
	Mac *string `json:"mac,omitempty"`
}

type ZdAPList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ZdAP `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ZdImport struct {
	// ApMacList
	// List of AP MAC
	ApMacList []string `json:"apMacList,omitempty"`

	// IP
	// ZD IP address
	IP *string `json:"ip,omitempty"`

	// Password
	// ZD password
	Password *string `json:"password,omitempty"`

	// User
	// ZD user name
	User *string `json:"user,omitempty"`
}

type ZdImportStatus struct {
	// Details
	// Details
	Details *string `json:"details,omitempty"`

	// Message
	// Message
	Message *string `json:"message,omitempty"`

	// Progress
	// Progress
	Progress *int `json:"progress,omitempty"`

	// State
	// State
	State *string `json:"state,omitempty"`
}
