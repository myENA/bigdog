package administration

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterblade"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ActiveDirectoryServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// Port
	// Port number of Active Directory Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`

	// WindowsDomainName
	// Windows Domain Name of Active Directory Server object
	// Constraints:
	//    - required
	WindowsDomainName *string `json:"windowsDomainName" validate:"required"`
}

func NewActiveDirectoryServer() *ActiveDirectoryServer {
	activeDirectoryServerType := new(ActiveDirectoryServer)
	return activeDirectoryServerType
}

func NewDefaultActiveDirectoryServer() *ActiveDirectoryServer {
	activeDirectoryServerType := new(ActiveDirectoryServer)
	return activeDirectoryServerType
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

func NewApPatchHistory() *ApPatchHistory {
	apPatchHistoryType := new(ApPatchHistory)
	return apPatchHistoryType
}

func NewDefaultApPatchHistory() *ApPatchHistory {
	apPatchHistoryType := new(ApPatchHistory)
	return apPatchHistoryType
}

type ApPatchHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApPatchHistory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApPatchHistoryList() *ApPatchHistoryList {
	apPatchHistoryListType := new(ApPatchHistoryList)
	return apPatchHistoryListType
}

func NewDefaultApPatchHistoryList() *ApPatchHistoryList {
	apPatchHistoryListType := new(ApPatchHistoryList)
	return apPatchHistoryListType
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

func NewApPatchInfo() *ApPatchInfo {
	apPatchInfoType := new(ApPatchInfo)
	return apPatchInfoType
}

func NewDefaultApPatchInfo() *ApPatchInfo {
	apPatchInfoType := new(ApPatchInfo)
	return apPatchInfoType
}

type ApPatchStatus struct {
	ClusterOperationProgress *clusterblade.ClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
}

func NewApPatchStatus() *ApPatchStatus {
	apPatchStatusType := new(ApPatchStatus)
	return apPatchStatusType
}

func NewDefaultApPatchStatus() *ApPatchStatus {
	apPatchStatusType := new(ApPatchStatus)
	return apPatchStatusType
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

func NewApplicationLogAndStatus() *ApplicationLogAndStatus {
	applicationLogAndStatusType := new(ApplicationLogAndStatus)
	return applicationLogAndStatusType
}

func NewDefaultApplicationLogAndStatus() *ApplicationLogAndStatus {
	applicationLogAndStatusType := new(ApplicationLogAndStatus)
	return applicationLogAndStatusType
}

type ApplicationLogAndStatusList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ApplicationLogAndStatus `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewApplicationLogAndStatusList() *ApplicationLogAndStatusList {
	applicationLogAndStatusListType := new(ApplicationLogAndStatusList)
	return applicationLogAndStatusListType
}

func NewDefaultApplicationLogAndStatusList() *ApplicationLogAndStatusList {
	applicationLogAndStatusListType := new(ApplicationLogAndStatusList)
	return applicationLogAndStatusListType
}

type AutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// FTP server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

func NewAutoExportBackup() *AutoExportBackup {
	autoExportBackupType := new(AutoExportBackup)
	return autoExportBackupType
}

func NewDefaultAutoExportBackup() *AutoExportBackup {
	autoExportBackupType := new(AutoExportBackup)
	return autoExportBackupType
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

	// Id
	// Identifier of system configuration backup file.
	Id *string `json:"id,omitempty"`

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

func NewBackupFile() *BackupFile {
	backupFileType := new(BackupFile)
	return backupFileType
}

func NewDefaultBackupFile() *BackupFile {
	backupFileType := new(BackupFile)
	return backupFileType
}

type ClusterBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ClusterBackupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewClusterBackupList() *ClusterBackupList {
	clusterBackupListType := new(ClusterBackupList)
	return clusterBackupListType
}

func NewDefaultClusterBackupList() *ClusterBackupList {
	clusterBackupListType := new(ClusterBackupList)
	return clusterBackupListType
}

type ClusterBackupSummary struct {
	// CreatedOn
	// Created date and time of the cluster backup file
	CreatedOn *string `json:"createdOn,omitempty"`

	// Filesize
	// filesize of the cluster backup file.
	Filesize *float64 `json:"filesize,omitempty"`

	// Id
	// Identifier of cluster backup file.
	Id *string `json:"id,omitempty"`

	// Version
	// the patch version of the cluster backup file.
	Version *string `json:"version,omitempty"`
}

func NewClusterBackupSummary() *ClusterBackupSummary {
	clusterBackupSummaryType := new(ClusterBackupSummary)
	return clusterBackupSummaryType
}

func NewDefaultClusterBackupSummary() *ClusterBackupSummary {
	clusterBackupSummaryType := new(ClusterBackupSummary)
	return clusterBackupSummaryType
}

type ConfigurationBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*BackupFile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewConfigurationBackupList() *ConfigurationBackupList {
	configurationBackupListType := new(ConfigurationBackupList)
	return configurationBackupListType
}

func NewDefaultConfigurationBackupList() *ConfigurationBackupList {
	configurationBackupListType := new(ConfigurationBackupList)
	return configurationBackupListType
}

type CreateAdminAAAServer struct {
	ActiveDirectoryServer *ActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *DefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *LdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	RadiusServer *RadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *TacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type" validate:"required,oneof=RADIUS TACACS AD LDAP"`
}

func NewCreateAdminAAAServer() *CreateAdminAAAServer {
	createAdminAAAServerType := new(CreateAdminAAAServer)
	return createAdminAAAServerType
}

func NewDefaultCreateAdminAAAServer() *CreateAdminAAAServer {
	createAdminAAAServerType := new(CreateAdminAAAServer)
	return createAdminAAAServerType
}

type DefaultRoleMapping struct {
	// DefaultAdmin
	// DefaultAdmin of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultAdmin *string `json:"defaultAdmin" validate:"required"`

	// DefaultUserGroup
	// DefaultUserGroup of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultUserGroup *string `json:"defaultUserGroup" validate:"required"`
}

func NewDefaultRoleMapping() *DefaultRoleMapping {
	defaultRoleMappingType := new(DefaultRoleMapping)
	return defaultRoleMappingType
}

func NewDefaultDefaultRoleMapping() *DefaultRoleMapping {
	defaultRoleMappingType := new(DefaultRoleMapping)
	return defaultRoleMappingType
}

type LdapServer struct {
	// AdminDomainName
	// Admin Domain Name of LDAP Server object
	// Constraints:
	//    - required
	AdminDomainName *string `json:"adminDomainName" validate:"required"`

	// AdminPassword
	// Admin password of LDAP Server object
	// Constraints:
	//    - required
	AdminPassword *string `json:"adminPassword" validate:"required"`

	// BaseDomainName
	// Base Domain Name of LDAP Server object
	// Constraints:
	//    - required
	BaseDomainName *string `json:"baseDomainName" validate:"required"`

	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// KeyAttribute
	// Key attribute of LDAP Server object
	// Constraints:
	//    - required
	KeyAttribute *string `json:"keyAttribute" validate:"required"`

	// Port
	// Port number of LDAP Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`

	// SearchFilter
	// Search filter of LDAP Server object
	// Constraints:
	//    - required
	SearchFilter *string `json:"searchFilter" validate:"required"`
}

func NewLdapServer() *LdapServer {
	ldapServerType := new(LdapServer)
	return ldapServerType
}

func NewDefaultLdapServer() *LdapServer {
	ldapServerType := new(LdapServer)
	return ldapServerType
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

func NewLicenses() *Licenses {
	licensesType := new(Licenses)
	return licensesType
}

func NewDefaultLicenses() *Licenses {
	licensesType := new(Licenses)
	return licensesType
}

type LicenseServer struct {
	// IpAddress
	// local license server IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// local license server port
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=0,lte=65535"`

	// UseCloud
	// use cloud license server
	UseCloud *bool `json:"useCloud,omitempty"`
}

func NewLicenseServer() *LicenseServer {
	licenseServerType := new(LicenseServer)
	return licenseServerType
}

func NewDefaultLicenseServer() *LicenseServer {
	licenseServerType := new(LicenseServer)
	return licenseServerType
}

type LicensesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*Licenses `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewLicensesList() *LicensesList {
	licensesListType := new(LicensesList)
	return licensesListType
}

func NewDefaultLicensesList() *LicensesList {
	licensesListType := new(LicensesList)
	return licensesListType
}

type LicensesSummary struct {
	CapacityControlLicenseCount *LicensesSummaryCapacityControlLicenseCountType `json:"capacityControlLicenseCount,omitempty"`

	// LicenseTypeDescription
	// license type description
	LicenseTypeDescription *string `json:"licenseTypeDescription,omitempty"`
}

func NewLicensesSummary() *LicensesSummary {
	licensesSummaryType := new(LicensesSummary)
	return licensesSummaryType
}

func NewDefaultLicensesSummary() *LicensesSummary {
	licensesSummaryType := new(LicensesSummary)
	return licensesSummaryType
}

type LicensesSummaryCapacityControlLicenseCountType struct {
	// TotalCount
	// total count of licenses
	TotalCount *int `json:"totalCount,omitempty"`

	// UsedCount
	// consumed count of licenses
	UsedCount *int `json:"usedCount,omitempty"`
}

func NewLicensesSummaryCapacityControlLicenseCountType() *LicensesSummaryCapacityControlLicenseCountType {
	licensesSummaryCapacityControlLicenseCountTypeType := new(LicensesSummaryCapacityControlLicenseCountType)
	return licensesSummaryCapacityControlLicenseCountTypeType
}

func NewDefaultLicensesSummaryCapacityControlLicenseCountType() *LicensesSummaryCapacityControlLicenseCountType {
	licensesSummaryCapacityControlLicenseCountTypeType := new(LicensesSummaryCapacityControlLicenseCountType)
	return licensesSummaryCapacityControlLicenseCountTypeType
}

type LicensesSummaryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LicensesSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewLicensesSummaryList() *LicensesSummaryList {
	licensesSummaryListType := new(LicensesSummaryList)
	return licensesSummaryListType
}

func NewDefaultLicensesSummaryList() *LicensesSummaryList {
	licensesSummaryListType := new(LicensesSummaryList)
	return licensesSummaryListType
}

type LicensesSyncLogs struct {
	// CreateDateTime
	// license sync log's create time
	CreateDateTime *string `json:"createDateTime,omitempty"`

	// SyncResult
	// sync license result
	// Constraints:
	//    - nullable
	//    - oneof:[SUCCESS,FAILURE]
	SyncResult *string `json:"syncResult,omitempty" validate:"omitempty,oneof=SUCCESS FAILURE"`
}

func NewLicensesSyncLogs() *LicensesSyncLogs {
	licensesSyncLogsType := new(LicensesSyncLogs)
	return licensesSyncLogsType
}

func NewDefaultLicensesSyncLogs() *LicensesSyncLogs {
	licensesSyncLogsType := new(LicensesSyncLogs)
	return licensesSyncLogsType
}

type LicensesSyncLogsList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*LicensesSyncLogs `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewLicensesSyncLogsList() *LicensesSyncLogsList {
	licensesSyncLogsListType := new(LicensesSyncLogsList)
	return licensesSyncLogsListType
}

func NewDefaultLicensesSyncLogsList() *LicensesSyncLogsList {
	licensesSyncLogsListType := new(LicensesSyncLogsList)
	return licensesSyncLogsListType
}

type ModfiyLicenseServer struct {
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"omitempty,gte=0,lte=65535"`

	// UseCloud
	// Constraints:
	//    - required
	UseCloud *bool `json:"useCloud" validate:"required"`
}

func NewModfiyLicenseServer() *ModfiyLicenseServer {
	modfiyLicenseServerType := new(ModfiyLicenseServer)
	return modfiyLicenseServerType
}

func NewDefaultModfiyLicenseServer() *ModfiyLicenseServer {
	modfiyLicenseServerType := new(ModfiyLicenseServer)
	return modfiyLicenseServerType
}

type ModifyAdminAAAServer struct {
	ActiveDirectoryServer *ActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *DefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *LdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	RadiusServer *RadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *TacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type" validate:"required,oneof=RADIUS TACACS AD LDAP"`
}

func NewModifyAdminAAAServer() *ModifyAdminAAAServer {
	modifyAdminAAAServerType := new(ModifyAdminAAAServer)
	return modifyAdminAAAServerType
}

func NewDefaultModifyAdminAAAServer() *ModifyAdminAAAServer {
	modifyAdminAAAServerType := new(ModifyAdminAAAServer)
	return modifyAdminAAAServerType
}

type ModifyAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// ftp server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

func NewModifyAutoExportBackup() *ModifyAutoExportBackup {
	modifyAutoExportBackupType := new(ModifyAutoExportBackup)
	return modifyAutoExportBackupType
}

func NewDefaultModifyAutoExportBackup() *ModifyAutoExportBackup {
	modifyAutoExportBackupType := new(ModifyAutoExportBackup)
	enableAutoExportBackupField := false
	modifyAutoExportBackupType.EnableAutoExportBackup = &enableAutoExportBackupField
	return modifyAutoExportBackupType
}

type ModifyLogLevel struct {
	// ApplicationName
	// Application name.
	ApplicationName *string `json:"applicationName,omitempty"`

	// LogLevel
	// Log level.
	// Constraints:
	//    - nullable
	//    - oneof:[DEBUG,INFO,WARN,ERROR]
	LogLevel *string `json:"logLevel,omitempty" validate:"omitempty,oneof=DEBUG INFO WARN ERROR"`
}

func NewModifyLogLevel() *ModifyLogLevel {
	modifyLogLevelType := new(ModifyLogLevel)
	return modifyLogLevelType
}

func NewDefaultModifyLogLevel() *ModifyLogLevel {
	modifyLogLevelType := new(ModifyLogLevel)
	return modifyLogLevelType
}

type ModifyScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - nullable
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"omitempty,oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - nullable
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty" validate:"omitempty,oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

func NewModifyScheduleBackup() *ModifyScheduleBackup {
	modifyScheduleBackupType := new(ModifyScheduleBackup)
	return modifyScheduleBackupType
}

func NewDefaultModifyScheduleBackup() *ModifyScheduleBackup {
	modifyScheduleBackupType := new(ModifyScheduleBackup)
	enableScheduleBackupField := false
	modifyScheduleBackupType.EnableScheduleBackup = &enableScheduleBackupField
	hourField := 0
	modifyScheduleBackupType.Hour = &hourField
	minuteField := 0
	modifyScheduleBackupType.Minute = &minuteField
	return modifyScheduleBackupType
}

type RadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// Port
	// Port number of RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Realm
	// Constraints:
	//    - required
	Realm *common.Realm `json:"realm" validate:"required"`

	SecondaryRadiusServer *SecondaryRadiusServer `json:"secondaryRadiusServer,omitempty"`

	// SharedSecret
	// Shared secret of RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

func NewRadiusServer() *RadiusServer {
	radiusServerType := new(RadiusServer)
	return radiusServerType
}

func NewDefaultRadiusServer() *RadiusServer {
	radiusServerType := new(RadiusServer)
	return radiusServerType
}

type RetrieveAdminAAAServer struct {
	ActiveDirectoryServer *ActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *DefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	// Id
	// ID of this Admin AAA Server
	Id *string `json:"id,omitempty"`

	LdapServer *LdapServer `json:"ldapServer,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	RadiusServer *RadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *TacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS TACACS AD LDAP"`
}

func NewRetrieveAdminAAAServer() *RetrieveAdminAAAServer {
	retrieveAdminAAAServerType := new(RetrieveAdminAAAServer)
	return retrieveAdminAAAServerType
}

func NewDefaultRetrieveAdminAAAServer() *RetrieveAdminAAAServer {
	retrieveAdminAAAServerType := new(RetrieveAdminAAAServer)
	return retrieveAdminAAAServerType
}

type RetrieveAdminAAAServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RetrieveAdminAAAServerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRetrieveAdminAAAServerList() *RetrieveAdminAAAServerList {
	retrieveAdminAAAServerListType := new(RetrieveAdminAAAServerList)
	return retrieveAdminAAAServerListType
}

func NewDefaultRetrieveAdminAAAServerList() *RetrieveAdminAAAServerList {
	retrieveAdminAAAServerListType := new(RetrieveAdminAAAServerList)
	return retrieveAdminAAAServerListType
}

type RetrieveAdminAAAServerListType struct {
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS TACACS AD LDAP"`
}

func NewRetrieveAdminAAAServerListType() *RetrieveAdminAAAServerListType {
	retrieveAdminAAAServerListTypeType := new(RetrieveAdminAAAServerListType)
	return retrieveAdminAAAServerListTypeType
}

func NewDefaultRetrieveAdminAAAServerListType() *RetrieveAdminAAAServerListType {
	retrieveAdminAAAServerListTypeType := new(RetrieveAdminAAAServerListType)
	return retrieveAdminAAAServerListTypeType
}

type ScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - nullable
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"omitempty,oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - nullable
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty" validate:"omitempty,oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

func NewScheduleBackup() *ScheduleBackup {
	scheduleBackupType := new(ScheduleBackup)
	return scheduleBackupType
}

func NewDefaultScheduleBackup() *ScheduleBackup {
	scheduleBackupType := new(ScheduleBackup)
	return scheduleBackupType
}

type SecondaryRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// MaxRetries
	// Max number(how many times) of retries for re-connection to primary
	// Constraints:
	//    - required
	//    - default:2
	//    - min:2
	//    - max:10
	MaxRetries *int `json:"maxRetries" validate:"required,gte=2,lte=10"`

	// Port
	// Port number of Secondary RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// RequestTimeOut
	// Request timeout(seconds) value of re-connection to primary
	// Constraints:
	//    - required
	//    - default:3
	//    - min:2
	//    - max:20
	RequestTimeOut *int `json:"requestTimeOut" validate:"required,gte=2,lte=20"`

	// RetryPriInvl
	// Interval of re-connection to primary(1-60 minute)
	// Constraints:
	//    - required
	//    - default:5
	//    - min:1
	//    - max:60
	RetryPriInvl *int `json:"retryPriInvl" validate:"required,gte=1,lte=60"`

	// SharedSecret
	// Shared secret of Secondary RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

func NewSecondaryRadiusServer() *SecondaryRadiusServer {
	secondaryRadiusServerType := new(SecondaryRadiusServer)
	return secondaryRadiusServerType
}

func NewDefaultSecondaryRadiusServer() *SecondaryRadiusServer {
	secondaryRadiusServerType := new(SecondaryRadiusServer)
	maxRetriesField := 2
	secondaryRadiusServerType.MaxRetries = &maxRetriesField
	requestTimeOutField := 3
	secondaryRadiusServerType.RequestTimeOut = &requestTimeOutField
	retryPriInvlField := 5
	secondaryRadiusServerType.RetryPriInvl = &retryPriInvlField
	return secondaryRadiusServerType
}

type TacacsServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *common.IpAddress `json:"ip" validate:"required"`

	// Port
	// Port number of TACACS+ Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Service
	// Constraints:
	//    - required
	Service *common.Realm `json:"service" validate:"required"`

	// SharedSecret
	// Shared secret of TACACS+ Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

func NewTacacsServer() *TacacsServer {
	tacacsServerType := new(TacacsServer)
	return tacacsServerType
}

func NewDefaultTacacsServer() *TacacsServer {
	tacacsServerType := new(TacacsServer)
	return tacacsServerType
}

type UpgradeHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UpgradeHistorySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewUpgradeHistoryList() *UpgradeHistoryList {
	upgradeHistoryListType := new(UpgradeHistoryList)
	return upgradeHistoryListType
}

func NewDefaultUpgradeHistoryList() *UpgradeHistoryList {
	upgradeHistoryListType := new(UpgradeHistoryList)
	return upgradeHistoryListType
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

func NewUpgradeHistorySummary() *UpgradeHistorySummary {
	upgradeHistorySummaryType := new(UpgradeHistorySummary)
	return upgradeHistorySummaryType
}

func NewDefaultUpgradeHistorySummary() *UpgradeHistorySummary {
	upgradeHistorySummaryType := new(UpgradeHistorySummary)
	return upgradeHistorySummaryType
}

type UpgradePatchInfo struct {
	ClusterOperationProgress *clusterblade.ClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`

	UploadPatchInfo *clusterblade.UploadPatchInfo `json:"uploadPatchInfo,omitempty"`
}

func NewUpgradePatchInfo() *UpgradePatchInfo {
	upgradePatchInfoType := new(UpgradePatchInfo)
	return upgradePatchInfoType
}

func NewDefaultUpgradePatchInfo() *UpgradePatchInfo {
	upgradePatchInfoType := new(UpgradePatchInfo)
	return upgradePatchInfoType
}

type UpgradeStatus struct {
	ClusterOperationProgress *clusterblade.ClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`
}

func NewUpgradeStatus() *UpgradeStatus {
	upgradeStatusType := new(UpgradeStatus)
	return upgradeStatusType
}

func NewDefaultUpgradeStatus() *UpgradeStatus {
	upgradeStatusType := new(UpgradeStatus)
	return upgradeStatusType
}

type ZdAP struct {
	// Connected
	// AP Conntected
	Connected *string `json:"connected,omitempty"`

	// Mac
	// AP MAC
	Mac *string `json:"mac,omitempty"`
}

func NewZdAP() *ZdAP {
	zdAPType := new(ZdAP)
	return zdAPType
}

func NewDefaultZdAP() *ZdAP {
	zdAPType := new(ZdAP)
	return zdAPType
}

type ZdAPList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ZdAP `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewZdAPList() *ZdAPList {
	zdAPListType := new(ZdAPList)
	return zdAPListType
}

func NewDefaultZdAPList() *ZdAPList {
	zdAPListType := new(ZdAPList)
	return zdAPListType
}

type ZdImport struct {
	// ApMacList
	// List of AP MAC
	ApMacList []string `json:"apMacList,omitempty"`

	// Ip
	// ZD IP address
	Ip *string `json:"ip,omitempty"`

	// Password
	// ZD password
	Password *string `json:"password,omitempty"`

	// User
	// ZD user name
	User *string `json:"user,omitempty"`
}

func NewZdImport() *ZdImport {
	zdImportType := new(ZdImport)
	return zdImportType
}

func NewDefaultZdImport() *ZdImport {
	zdImportType := new(ZdImport)
	return zdImportType
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

func NewZdImportStatus() *ZdImportStatus {
	zdImportStatusType := new(ZdImportStatus)
	return zdImportStatusType
}

func NewDefaultZdImportStatus() *ZdImportStatus {
	zdImportStatusType := new(ZdImportStatus)
	return zdImportStatusType
}
