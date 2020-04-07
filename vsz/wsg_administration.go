package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAdministrationService struct {
	apiClient *APIClient
}

func NewWSGAdministrationService(c *APIClient) *WSGAdministrationService {
	s := new(WSGAdministrationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAdministrationService() *WSGAdministrationService {
	return NewWSGAdministrationService(ss.apiClient)
}

type WSGAdministrationActiveDirectoryServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

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
	Realm *WSGCommonRealm `json:"realm" validate:"required"`

	// WindowsDomainName
	// Windows Domain Name of Active Directory Server object
	// Constraints:
	//    - required
	WindowsDomainName *string `json:"windowsDomainName" validate:"required"`
}

func NewWSGAdministrationActiveDirectoryServer() *WSGAdministrationActiveDirectoryServer {
	m := new(WSGAdministrationActiveDirectoryServer)
	return m
}

type WSGAdministrationApPatchHistory struct {
	// ApFwVersion
	// apFwVersion of the AP Patch history
	// Constraints:
	//    - nullable
	ApFwVersion *string `json:"apFwVersion,omitempty"`

	// ApModelList
	// AP Models of the AP Patch history
	// Constraints:
	//    - nullable
	ApModelList []string `json:"apModelList,omitempty" validate:"omitempty,dive"`

	// FileName
	// file name of the AP Patch history
	// Constraints:
	//    - nullable
	FileName *string `json:"fileName,omitempty"`

	// StartDateTime
	// startDateTime of the AP Patch history
	// Constraints:
	//    - nullable
	StartDateTime *string `json:"startDateTime,omitempty"`
}

func NewWSGAdministrationApPatchHistory() *WSGAdministrationApPatchHistory {
	m := new(WSGAdministrationApPatchHistory)
	return m
}

type WSGAdministrationApPatchHistoryList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationApPatchHistory `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationApPatchHistoryList() *WSGAdministrationApPatchHistoryList {
	m := new(WSGAdministrationApPatchHistoryList)
	return m
}

type WSGAdministrationApPatchInfo struct {
	// ApModels
	// AP Models of the upload file
	// Constraints:
	//    - nullable
	ApModels []string `json:"apModels,omitempty" validate:"omitempty,dive"`

	// ApVersion
	// ApFwVersion of the upload file
	// Constraints:
	//    - nullable
	ApVersion *string `json:"apVersion,omitempty"`

	// FileName
	// file name of the upload file
	// Constraints:
	//    - nullable
	FileName *string `json:"fileName,omitempty"`

	// FileSize
	// file size(Byte) of the upload file
	// Constraints:
	//    - nullable
	FileSize *int `json:"fileSize,omitempty"`
}

func NewWSGAdministrationApPatchInfo() *WSGAdministrationApPatchInfo {
	m := new(WSGAdministrationApPatchInfo)
	return m
}

type WSGAdministrationApPatchStatus struct {
	// ClusterOperationProgress
	// Constraints:
	//    - nullable
	ClusterOperationProgress *WSGClusterBladeClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
}

func NewWSGAdministrationApPatchStatus() *WSGAdministrationApPatchStatus {
	m := new(WSGAdministrationApPatchStatus)
	return m
}

type WSGAdministrationApplicationLogAndStatus struct {
	// ApplicationName
	// Application name
	// Constraints:
	//    - nullable
	ApplicationName *string `json:"applicationName,omitempty"`

	// HealthStatus
	// Health status
	// Constraints:
	//    - nullable
	HealthStatus *string `json:"healthStatus,omitempty"`

	// LogFileNames
	// List of log file name
	// Constraints:
	//    - nullable
	LogFileNames []string `json:"logFileNames,omitempty" validate:"omitempty,dive"`

	// LogLevel
	// Log level
	// Constraints:
	//    - nullable
	LogLevel *string `json:"logLevel,omitempty"`

	// NumOfLogs
	// # of Logs
	// Constraints:
	//    - nullable
	NumOfLogs *int `json:"numOfLogs,omitempty"`
}

func NewWSGAdministrationApplicationLogAndStatus() *WSGAdministrationApplicationLogAndStatus {
	m := new(WSGAdministrationApplicationLogAndStatus)
	return m
}

type WSGAdministrationApplicationLogAndStatusList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationApplicationLogAndStatus `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationApplicationLogAndStatusList() *WSGAdministrationApplicationLogAndStatusList {
	m := new(WSGAdministrationApplicationLogAndStatusList)
	return m
}

type WSGAdministrationAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	// Constraints:
	//    - nullable
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// FTP server name
	// Constraints:
	//    - nullable
	FtpServer *string `json:"ftpServer,omitempty"`
}

func NewWSGAdministrationAutoExportBackup() *WSGAdministrationAutoExportBackup {
	m := new(WSGAdministrationAutoExportBackup)
	return m
}

type WSGAdministrationBackupFile struct {
	// BackupElapsed
	// backup elapsed of the configuration backup file
	// Constraints:
	//    - nullable
	BackupElapsed *int `json:"backupElapsed,omitempty"`

	// ControlPlaneSoftwareVersion
	// control plane software version of the configuration backup file
	// Constraints:
	//    - nullable
	ControlPlaneSoftwareVersion *string `json:"controlPlaneSoftwareVersion,omitempty"`

	// CreatedBy
	// creator of the configuration backup file.
	// Constraints:
	//    - nullable
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedOn
	// the create time of the configuration backup file.
	// Constraints:
	//    - nullable
	CreatedOn *float64 `json:"createdOn,omitempty"`

	// DataPlaneSoftwareVersion
	// data plane software version of the configuration backup file
	// Constraints:
	//    - nullable
	DataPlaneSoftwareVersion *string `json:"dataPlaneSoftwareVersion,omitempty"`

	// FileSize
	// file size of the backup file
	// Constraints:
	//    - nullable
	FileSize *int `json:"fileSize,omitempty"`

	// Id
	// Identifier of system configuration backup file.
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Md5
	// file md5 of the backup file
	// Constraints:
	//    - nullable
	Md5 *string `json:"md5,omitempty"`

	// ScgVersion
	// SCG version of the configuration backup file.
	// Constraints:
	//    - nullable
	ScgVersion *string `json:"scgVersion,omitempty"`

	// Type
	// type of the configuration backup file
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewWSGAdministrationBackupFile() *WSGAdministrationBackupFile {
	m := new(WSGAdministrationBackupFile)
	return m
}

type WSGAdministrationClusterBackupList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationClusterBackupSummary `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationClusterBackupList() *WSGAdministrationClusterBackupList {
	m := new(WSGAdministrationClusterBackupList)
	return m
}

type WSGAdministrationClusterBackupSummary struct {
	// CreatedOn
	// Created date and time of the cluster backup file
	// Constraints:
	//    - nullable
	CreatedOn *string `json:"createdOn,omitempty"`

	// Filesize
	// filesize of the cluster backup file.
	// Constraints:
	//    - nullable
	Filesize *float64 `json:"filesize,omitempty"`

	// Id
	// Identifier of cluster backup file.
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Version
	// the patch version of the cluster backup file.
	// Constraints:
	//    - nullable
	Version *string `json:"version,omitempty"`
}

func NewWSGAdministrationClusterBackupSummary() *WSGAdministrationClusterBackupSummary {
	m := new(WSGAdministrationClusterBackupSummary)
	return m
}

type WSGAdministrationConfigurationBackupList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationBackupFile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationConfigurationBackupList() *WSGAdministrationConfigurationBackupList {
	m := new(WSGAdministrationConfigurationBackupList)
	return m
}

type WSGAdministrationCreateAdminAAAServer struct {
	// ActiveDirectoryServer
	// Constraints:
	//    - nullable
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	// DefaultRoleMapping
	// Constraints:
	//    - nullable
	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	// LdapServer
	// Constraints:
	//    - nullable
	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// RadiusServer
	// Constraints:
	//    - nullable
	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	// TacacsServer
	// Constraints:
	//    - nullable
	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type" validate:"required,oneof=RADIUS TACACS AD LDAP"`
}

func NewWSGAdministrationCreateAdminAAAServer() *WSGAdministrationCreateAdminAAAServer {
	m := new(WSGAdministrationCreateAdminAAAServer)
	return m
}

type WSGAdministrationDefaultRoleMapping struct {
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

func NewWSGAdministrationDefaultRoleMapping() *WSGAdministrationDefaultRoleMapping {
	m := new(WSGAdministrationDefaultRoleMapping)
	return m
}

type WSGAdministrationLdapServer struct {
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
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

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
	Realm *WSGCommonRealm `json:"realm" validate:"required"`

	// SearchFilter
	// Search filter of LDAP Server object
	// Constraints:
	//    - required
	SearchFilter *string `json:"searchFilter" validate:"required"`
}

func NewWSGAdministrationLdapServer() *WSGAdministrationLdapServer {
	m := new(WSGAdministrationLdapServer)
	return m
}

type WSGAdministrationLicenses struct {
	// Count
	// number of licenses
	// Constraints:
	//    - nullable
	Count *int `json:"count,omitempty"`

	// CreateTime
	// license effective date
	// Constraints:
	//    - nullable
	CreateTime *string `json:"createTime,omitempty"`

	// Description
	// license description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// ExpireDate
	// license expiry date
	// Constraints:
	//    - nullable
	ExpireDate *string `json:"expireDate,omitempty"`

	// Name
	// license name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`
}

func NewWSGAdministrationLicenses() *WSGAdministrationLicenses {
	m := new(WSGAdministrationLicenses)
	return m
}

type WSGAdministrationLicenseServer struct {
	// IpAddress
	// local license server IP address
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	UseCloud *bool `json:"useCloud,omitempty"`
}

func NewWSGAdministrationLicenseServer() *WSGAdministrationLicenseServer {
	m := new(WSGAdministrationLicenseServer)
	return m
}

type WSGAdministrationLicensesList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationLicenses `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationLicensesList() *WSGAdministrationLicensesList {
	m := new(WSGAdministrationLicensesList)
	return m
}

type WSGAdministrationLicensesSummary struct {
	// CapacityControlLicenseCount
	// Constraints:
	//    - nullable
	CapacityControlLicenseCount *WSGAdministrationLicensesSummaryCapacityControlLicenseCountType `json:"capacityControlLicenseCount,omitempty"`

	// LicenseTypeDescription
	// license type description
	// Constraints:
	//    - nullable
	LicenseTypeDescription *string `json:"licenseTypeDescription,omitempty"`
}

func NewWSGAdministrationLicensesSummary() *WSGAdministrationLicensesSummary {
	m := new(WSGAdministrationLicensesSummary)
	return m
}

type WSGAdministrationLicensesSummaryCapacityControlLicenseCountType struct {
	// TotalCount
	// total count of licenses
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`

	// UsedCount
	// consumed count of licenses
	// Constraints:
	//    - nullable
	UsedCount *int `json:"usedCount,omitempty"`
}

func NewWSGAdministrationLicensesSummaryCapacityControlLicenseCountType() *WSGAdministrationLicensesSummaryCapacityControlLicenseCountType {
	m := new(WSGAdministrationLicensesSummaryCapacityControlLicenseCountType)
	return m
}

type WSGAdministrationLicensesSummaryList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationLicensesSummary `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationLicensesSummaryList() *WSGAdministrationLicensesSummaryList {
	m := new(WSGAdministrationLicensesSummaryList)
	return m
}

type WSGAdministrationLicensesSyncLogs struct {
	// CreateDateTime
	// license sync log's create time
	// Constraints:
	//    - nullable
	CreateDateTime *string `json:"createDateTime,omitempty"`

	// SyncResult
	// sync license result
	// Constraints:
	//    - nullable
	//    - oneof:[SUCCESS,FAILURE]
	SyncResult *string `json:"syncResult,omitempty" validate:"omitempty,oneof=SUCCESS FAILURE"`
}

func NewWSGAdministrationLicensesSyncLogs() *WSGAdministrationLicensesSyncLogs {
	m := new(WSGAdministrationLicensesSyncLogs)
	return m
}

type WSGAdministrationLicensesSyncLogsList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationLicensesSyncLogs `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationLicensesSyncLogsList() *WSGAdministrationLicensesSyncLogsList {
	m := new(WSGAdministrationLicensesSyncLogsList)
	return m
}

type WSGAdministrationModfiyLicenseServer struct {
	// IpAddress
	// Constraints:
	//    - nullable
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

func NewWSGAdministrationModfiyLicenseServer() *WSGAdministrationModfiyLicenseServer {
	m := new(WSGAdministrationModfiyLicenseServer)
	return m
}

type WSGAdministrationModifyAdminAAAServer struct {
	// ActiveDirectoryServer
	// Constraints:
	//    - nullable
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	// DefaultRoleMapping
	// Constraints:
	//    - nullable
	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	// LdapServer
	// Constraints:
	//    - nullable
	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// RadiusServer
	// Constraints:
	//    - nullable
	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	// TacacsServer
	// Constraints:
	//    - nullable
	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type" validate:"required,oneof=RADIUS TACACS AD LDAP"`
}

func NewWSGAdministrationModifyAdminAAAServer() *WSGAdministrationModifyAdminAAAServer {
	m := new(WSGAdministrationModifyAdminAAAServer)
	return m
}

type WSGAdministrationModifyAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	// Constraints:
	//    - nullable
	//    - default:false
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// ftp server name
	// Constraints:
	//    - nullable
	FtpServer *string `json:"ftpServer,omitempty"`
}

func NewWSGAdministrationModifyAutoExportBackup() *WSGAdministrationModifyAutoExportBackup {
	m := new(WSGAdministrationModifyAutoExportBackup)
	return m
}

type WSGAdministrationModifyLogLevel struct {
	// ApplicationName
	// Application name.
	// Constraints:
	//    - nullable
	ApplicationName *string `json:"applicationName,omitempty"`

	// LogLevel
	// Log level.
	// Constraints:
	//    - nullable
	//    - oneof:[DEBUG,INFO,WARN,ERROR]
	LogLevel *string `json:"logLevel,omitempty" validate:"omitempty,oneof=DEBUG INFO WARN ERROR"`
}

func NewWSGAdministrationModifyLogLevel() *WSGAdministrationModifyLogLevel {
	m := new(WSGAdministrationModifyLogLevel)
	return m
}

type WSGAdministrationModifyScheduleBackup struct {
	// DateOfMonth
	// date of the month
	// Constraints:
	//    - nullable
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - nullable
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"omitempty,oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// EnableScheduleBackup
	// enable schedule backup
	// Constraints:
	//    - nullable
	//    - default:false
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	// Constraints:
	//    - nullable
	//    - default:0
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - nullable
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty" validate:"omitempty,oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	// Constraints:
	//    - nullable
	//    - default:0
	Minute *int `json:"minute,omitempty"`
}

func NewWSGAdministrationModifyScheduleBackup() *WSGAdministrationModifyScheduleBackup {
	m := new(WSGAdministrationModifyScheduleBackup)
	return m
}

type WSGAdministrationRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Port
	// Port number of RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Protocol
	// Constraints:
	//    - nullable
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP,PEAP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=PAP CHAP PEAP"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required"`

	// SecondaryRadiusServer
	// Constraints:
	//    - nullable
	SecondaryRadiusServer *WSGAdministrationSecondaryRadiusServer `json:"secondaryRadiusServer,omitempty"`

	// SharedSecret
	// Shared secret of RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

func NewWSGAdministrationRadiusServer() *WSGAdministrationRadiusServer {
	m := new(WSGAdministrationRadiusServer)
	return m
}

type WSGAdministrationRetrieveAdminAAAServer struct {
	// ActiveDirectoryServer
	// Constraints:
	//    - nullable
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	// DefaultRoleMapping
	// Constraints:
	//    - nullable
	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	// Id
	// ID of this Admin AAA Server
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// LdapServer
	// Constraints:
	//    - nullable
	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RadiusServer
	// Constraints:
	//    - nullable
	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	// TacacsServer
	// Constraints:
	//    - nullable
	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS TACACS AD LDAP"`
}

func NewWSGAdministrationRetrieveAdminAAAServer() *WSGAdministrationRetrieveAdminAAAServer {
	m := new(WSGAdministrationRetrieveAdminAAAServer)
	return m
}

type WSGAdministrationRetrieveAdminAAAServerList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationRetrieveAdminAAAServerListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationRetrieveAdminAAAServerList() *WSGAdministrationRetrieveAdminAAAServerList {
	m := new(WSGAdministrationRetrieveAdminAAAServerList)
	return m
}

type WSGAdministrationRetrieveAdminAAAServerListType struct {
	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS TACACS AD LDAP"`
}

func NewWSGAdministrationRetrieveAdminAAAServerListType() *WSGAdministrationRetrieveAdminAAAServerListType {
	m := new(WSGAdministrationRetrieveAdminAAAServerListType)
	return m
}

type WSGAdministrationScheduleBackup struct {
	// DateOfMonth
	// date of the month
	// Constraints:
	//    - nullable
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - nullable
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"omitempty,oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// EnableScheduleBackup
	// enable schedule backup
	// Constraints:
	//    - nullable
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	// Constraints:
	//    - nullable
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - nullable
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty" validate:"omitempty,oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	// Constraints:
	//    - nullable
	Minute *int `json:"minute,omitempty"`
}

func NewWSGAdministrationScheduleBackup() *WSGAdministrationScheduleBackup {
	m := new(WSGAdministrationScheduleBackup)
	return m
}

type WSGAdministrationSecondaryRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

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

	// Protocol
	// Constraints:
	//    - nullable
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP,PEAP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=PAP CHAP PEAP"`

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

func NewWSGAdministrationSecondaryRadiusServer() *WSGAdministrationSecondaryRadiusServer {
	m := new(WSGAdministrationSecondaryRadiusServer)
	return m
}

type WSGAdministrationTacacsServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

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
	Service *WSGCommonRealm `json:"service" validate:"required"`

	// SharedSecret
	// Shared secret of TACACS+ Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

func NewWSGAdministrationTacacsServer() *WSGAdministrationTacacsServer {
	m := new(WSGAdministrationTacacsServer)
	return m
}

type WSGAdministrationUpgradeHistoryList struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationUpgradeHistorySummary `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationUpgradeHistoryList() *WSGAdministrationUpgradeHistoryList {
	m := new(WSGAdministrationUpgradeHistoryList)
	return m
}

type WSGAdministrationUpgradeHistorySummary struct {
	// ApFwVersion
	// apFwVersion of the upgrade history
	// Constraints:
	//    - nullable
	ApFwVersion *string `json:"apFwVersion,omitempty"`

	// CbVersion
	// cbVersion of the upgrade history
	// Constraints:
	//    - nullable
	CbVersion *string `json:"cbVersion,omitempty"`

	// DpVersion
	// dpVersion of the upgrade history
	// Constraints:
	//    - nullable
	DpVersion *string `json:"dpVersion,omitempty"`

	// ElapsedSeconds
	// elapsedSeconds of the upgrade history
	// Constraints:
	//    - nullable
	ElapsedSeconds *int `json:"elapsedSeconds,omitempty"`

	// FileName
	// fileName of the upgrade history
	// Constraints:
	//    - nullable
	FileName *string `json:"fileName,omitempty"`

	// OldApFwVersion
	// oldApFwVersion of the upgrade history
	// Constraints:
	//    - nullable
	OldApFwVersion *string `json:"oldApFwVersion,omitempty"`

	// OldCbVersion
	// oldCbVersion of the upgrade history
	// Constraints:
	//    - nullable
	OldCbVersion *string `json:"oldCbVersion,omitempty"`

	// OldDpVersion
	// oldDpVersion of the upgrade history
	// Constraints:
	//    - nullable
	OldDpVersion *string `json:"oldDpVersion,omitempty"`

	// OldVersion
	// oldVersion of the upgrade history
	// Constraints:
	//    - nullable
	OldVersion *string `json:"oldVersion,omitempty"`

	// StartTime
	// startTime of the upgrade history
	// Constraints:
	//    - nullable
	StartTime *string `json:"startTime,omitempty"`

	// Version
	// version of the upgrade history
	// Constraints:
	//    - nullable
	Version *string `json:"version,omitempty"`
}

func NewWSGAdministrationUpgradeHistorySummary() *WSGAdministrationUpgradeHistorySummary {
	m := new(WSGAdministrationUpgradeHistorySummary)
	return m
}

type WSGAdministrationUpgradePatchInfo struct {
	// ClusterOperationProgress
	// Constraints:
	//    - nullable
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`

	// UploadPatchInfo
	// Constraints:
	//    - nullable
	UploadPatchInfo *WSGClusterBladeUploadPatchInfo `json:"uploadPatchInfo,omitempty"`
}

func NewWSGAdministrationUpgradePatchInfo() *WSGAdministrationUpgradePatchInfo {
	m := new(WSGAdministrationUpgradePatchInfo)
	return m
}

type WSGAdministrationUpgradeStatus struct {
	// ClusterOperationProgress
	// Constraints:
	//    - nullable
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`
}

func NewWSGAdministrationUpgradeStatus() *WSGAdministrationUpgradeStatus {
	m := new(WSGAdministrationUpgradeStatus)
	return m
}

type WSGAdministrationZdAP struct {
	// Connected
	// AP Conntected
	// Constraints:
	//    - nullable
	Connected *string `json:"connected,omitempty"`

	// Mac
	// AP MAC
	// Constraints:
	//    - nullable
	Mac *string `json:"mac,omitempty"`
}

func NewWSGAdministrationZdAP() *WSGAdministrationZdAP {
	m := new(WSGAdministrationZdAP)
	return m
}

type WSGAdministrationZdAPList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGAdministrationZdAP `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdministrationZdAPList() *WSGAdministrationZdAPList {
	m := new(WSGAdministrationZdAPList)
	return m
}

type WSGAdministrationZdImport struct {
	// ApMacList
	// List of AP MAC
	// Constraints:
	//    - nullable
	ApMacList []string `json:"apMacList,omitempty" validate:"omitempty,dive"`

	// Ip
	// ZD IP address
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// Password
	// ZD password
	// Constraints:
	//    - nullable
	Password *string `json:"password,omitempty"`

	// User
	// ZD user name
	// Constraints:
	//    - nullable
	User *string `json:"user,omitempty"`
}

func NewWSGAdministrationZdImport() *WSGAdministrationZdImport {
	m := new(WSGAdministrationZdImport)
	return m
}

type WSGAdministrationZdImportStatus struct {
	// Details
	// Details
	// Constraints:
	//    - nullable
	Details *string `json:"details,omitempty"`

	// Message
	// Message
	// Constraints:
	//    - nullable
	Message *string `json:"message,omitempty"`

	// Progress
	// Progress
	// Constraints:
	//    - nullable
	Progress *int `json:"progress,omitempty"`

	// State
	// State
	// Constraints:
	//    - nullable
	State *string `json:"state,omitempty"`
}

func NewWSGAdministrationZdImportStatus() *WSGAdministrationZdImportStatus {
	m := new(WSGAdministrationZdImportStatus)
	return m
}

// AddAdminaaa
//
// Use this API command to create a new Admin AAA server
//
// Request Body:
//	 - body *WSGAdministrationCreateAdminAAAServer
func (s *WSGAdministrationService) AddAdminaaa(ctx context.Context, body *WSGAdministrationCreateAdminAAAServer) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAdminaaa, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRestart
//
// Use this API command to restart the controller.
func (s *WSGAdministrationService) AddRestart(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRestart, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddShutdown
//
// Use this API command to shut down the controller.
func (s *WSGAdministrationService) AddShutdown(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddShutdown, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteAdminaaaById
//
// Use this API command to delete an existing Admin AAA server
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAdministrationService) DeleteAdminaaaById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAdminaaaById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindAdminaaa
//
// Use this API command to retrieve the list of Admin AAA server
//
// Required Parameters:
// - type_ string
//		- required
func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context, type_ string) (*WSGAdministrationRetrieveAdminAAAServerList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationRetrieveAdminAAAServerList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, type_, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAdminaaa, true)
	req.SetQueryParameter("type_", []string{type_})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationRetrieveAdminAAAServerList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAdminaaaById
//
// Use this API command to retrieve an existing Admin AAA server
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAdministrationService) FindAdminaaaById(ctx context.Context, id string) (*WSGAdministrationRetrieveAdminAAAServer, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationRetrieveAdminAAAServer
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAdminaaaById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationRetrieveAdminAAAServer()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLicenses
//
// Use this API command to get all licenses currently assign in SCG.
func (s *WSGAdministrationService) FindLicenses(ctx context.Context) (*WSGAdministrationLicensesList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationLicensesList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLicenses, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationLicensesList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLicenseServer
//
// Use this API command to get license server configuration.
func (s *WSGAdministrationService) FindLicenseServer(ctx context.Context) (*WSGAdministrationLicenseServer, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationLicenseServer
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLicenseServer, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationLicenseServer()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLicensesSummary
//
// Use this API command to get licenses summary information.
func (s *WSGAdministrationService) FindLicensesSummary(ctx context.Context) (*WSGAdministrationLicensesSummaryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationLicensesSummaryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLicensesSummary, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationLicensesSummaryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLicensesSyncLogs
//
// Use this API command to get licenses synchronize logs.
func (s *WSGAdministrationService) FindLicensesSyncLogs(ctx context.Context) (*WSGAdministrationLicensesSyncLogsList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationLicensesSyncLogsList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLicensesSyncLogs, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationLicensesSyncLogsList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateAdminaaaById
//
// Use this API command to modify an existing Admin AAA server
//
// Request Body:
//	 - body *WSGAdministrationModifyAdminAAAServer
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAdministrationService) UpdateAdminaaaById(ctx context.Context, body *WSGAdministrationModifyAdminAAAServer, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAdminaaaById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// UpdateLicenseServer
//
// Use this API command to update license server configuration.
//
// Request Body:
//	 - body *WSGAdministrationModfiyLicenseServer
func (s *WSGAdministrationService) UpdateLicenseServer(ctx context.Context, body *WSGAdministrationModfiyLicenseServer) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateLicenseServer, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateLicensesSync
//
// Use this API command to ask all SCG in cluster to sync licenses from license server.
func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateLicensesSync, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
