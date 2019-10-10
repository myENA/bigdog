package administration

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterblade"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ActiveDirectoryServer struct {
	IP *common.IPAddress `json:"ip,omitempty" validate:"required"`

	// Port
	// Port number of Active Directory Server object
	Port *int `json:"port,omitempty" validate:"required,gte=1,lte=65535"`

	Realm *common.Realm `json:"realm,omitempty" validate:"required"`

	// WindowsDomainName
	// Windows Domain Name of Active Directory Server object
	WindowsDomainName *string `json:"windowsDomainName,omitempty" validate:"required"`
}

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
	ClusterOperationProgress *clusterblade.ClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
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

	// SCGVersion
	// SCG version of the configuration backup file.
	SCGVersion *string `json:"scgVersion,omitempty"`

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

type CreateAdminAAAServer struct {
	ActiveDirectoryServer *ActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *DefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *LdapServer `json:"ldapServer,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	RadiusServer *RadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *TacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name
	// [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	Type *string `json:"type,omitempty" validate:"required,oneof=RADIUS TACACS AD LDAP"`
}

type DefaultRoleMapping struct {
	// DefaultAdmin
	// DefaultAdmin of DefaultRoleMapping object
	DefaultAdmin *string `json:"defaultAdmin,omitempty" validate:"required"`

	// DefaultUserGroup
	// DefaultUserGroup of DefaultRoleMapping object
	DefaultUserGroup *string `json:"defaultUserGroup,omitempty" validate:"required"`
}

type LdapServer struct {
	// AdminDomainName
	// Admin Domain Name of LDAP Server object
	AdminDomainName *string `json:"adminDomainName,omitempty" validate:"required"`

	// AdminPassword
	// Admin password of LDAP Server object
	AdminPassword *string `json:"adminPassword,omitempty" validate:"required"`

	// BaseDomainName
	// Base Domain Name of LDAP Server object
	BaseDomainName *string `json:"baseDomainName,omitempty" validate:"required"`

	IP *common.IPAddress `json:"ip,omitempty" validate:"required"`

	// KeyAttribute
	// Key attribute of LDAP Server object
	KeyAttribute *string `json:"keyAttribute,omitempty" validate:"required"`

	// Port
	// Port number of LDAP Server object
	Port *int `json:"port,omitempty" validate:"required,gte=1,lte=65535"`

	Realm *common.Realm `json:"realm,omitempty" validate:"required"`

	// SearchFilter
	// Search filter of LDAP Server object
	SearchFilter *string `json:"searchFilter,omitempty" validate:"required"`
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
	Port *int `json:"port,omitempty" validate:"gte=0,lte=65535"`

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
	SyncResult *string `json:"syncResult,omitempty" validate:"oneof=SUCCESS FAILURE"`
}

type LicensesSyncLogsList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LicensesSyncLogs `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ModfiyLicenseServer struct {
	IPAddress *string `json:"ipAddress,omitempty"`

	Port *int `json:"port,omitempty" validate:"gte=0,lte=65535"`

	UseCloud *bool `json:"useCloud,omitempty" validate:"required"`
}

type ModifyAdminAAAServer struct {
	ActiveDirectoryServer *ActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *DefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *LdapServer `json:"ldapServer,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	RadiusServer *RadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *TacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name
	// [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	Type *string `json:"type,omitempty" validate:"required,oneof=RADIUS TACACS AD LDAP"`
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
	LogLevel *string `json:"logLevel,omitempty" validate:"oneof=DEBUG INFO WARN ERROR"`
}

type ModifyScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	Interval *string `json:"interval,omitempty" validate:"oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type RadiusServer struct {
	IP *common.IPAddress `json:"ip,omitempty" validate:"required"`

	// Port
	// Port number of RADIUS Server object
	Port *int `json:"port,omitempty" validate:"required,gte=1,lte=65535"`

	Realm *common.Realm `json:"realm,omitempty" validate:"required"`

	SecondaryRadiusServer *SecondaryRadiusServer `json:"secondaryRadiusServer,omitempty"`

	// SharedSecret
	// Shared secret of RADIUS Server object
	SharedSecret *string `json:"sharedSecret,omitempty" validate:"required"`
}

type RetrieveAdminAAAServer struct {
	ActiveDirectoryServer *ActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *DefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	// ID
	// ID of this Admin AAA Server
	ID *string `json:"id,omitempty"`

	LdapServer *LdapServer `json:"ldapServer,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	RadiusServer *RadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *TacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is
	// for TACACS+ (Terminal Access Controller Access-Control System Plus)
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS TACACS AD LDAP"`
}

type RetrieveAdminAAAServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RetrieveAdminAAAServerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type RetrieveAdminAAAServerListType struct {
	ID *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	Type *string `json:"type,omitempty" validate:"oneof=RADIUS TACACS AD LDAP"`
}

type ScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	Interval *string `json:"interval,omitempty" validate:"oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type SecondaryRadiusServer struct {
	IP *common.IPAddress `json:"ip,omitempty" validate:"required"`

	// MaxRetries
	// Max number(how many times) of retries for re-connection to primary
	MaxRetries *int `json:"maxRetries,omitempty" validate:"required,gte=2,lte=10"`

	// Port
	// Port number of Secondary RADIUS Server object
	Port *int `json:"port,omitempty" validate:"required,gte=1,lte=65535"`

	// RequestTimeOut
	// Request timeout(seconds) value of re-connection to primary
	RequestTimeOut *int `json:"requestTimeOut,omitempty" validate:"required,gte=2,lte=20"`

	// RetryPriInvl
	// Interval of re-connection to primary(1-60 minute)
	RetryPriInvl *int `json:"retryPriInvl,omitempty" validate:"required,gte=1,lte=60"`

	// SharedSecret
	// Shared secret of Secondary RADIUS Server object
	SharedSecret *string `json:"sharedSecret,omitempty" validate:"required"`
}

type TacacsServer struct {
	IP *common.IPAddress `json:"ip,omitempty" validate:"required"`

	// Port
	// Port number of TACACS+ Server object
	Port *int `json:"port,omitempty" validate:"required,gte=1,lte=65535"`

	Service *common.Realm `json:"service,omitempty" validate:"required"`

	// SharedSecret
	// Shared secret of TACACS+ Server object
	SharedSecret *string `json:"sharedSecret,omitempty" validate:"required"`
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
	ClusterOperationProgress *clusterblade.ClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`

	UploadPatchInfo *clusterblade.UploadPatchInfo `json:"uploadPatchInfo,omitempty"`
}

type UpgradeStatus struct {
	ClusterOperationProgress *clusterblade.ClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`
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
