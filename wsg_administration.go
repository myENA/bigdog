package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGAdministrationService struct {
	apiClient *VSZClient
}

func NewWSGAdministrationService(c *VSZClient) *WSGAdministrationService {
	s := new(WSGAdministrationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAdministrationService() *WSGAdministrationService {
	return NewWSGAdministrationService(ss.apiClient)
}

// WSGAdministrationActiveDirectoryServer
//
// Definition: administration_activeDirectoryServer
type WSGAdministrationActiveDirectoryServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Port
	// Port number of Active Directory Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// WindowsDomainName
	// Windows Domain Name of Active Directory Server object
	// Constraints:
	//    - required
	WindowsDomainName *string `json:"windowsDomainName"`
}

func NewWSGAdministrationActiveDirectoryServer() *WSGAdministrationActiveDirectoryServer {
	m := new(WSGAdministrationActiveDirectoryServer)
	return m
}

// WSGAdministrationApPatchHistory
//
// Definition: administration_apPatchHistory
type WSGAdministrationApPatchHistory struct {
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

func NewWSGAdministrationApPatchHistory() *WSGAdministrationApPatchHistory {
	m := new(WSGAdministrationApPatchHistory)
	return m
}

// WSGAdministrationApPatchHistoryList
//
// Definition: administration_apPatchHistoryList
type WSGAdministrationApPatchHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationApPatchHistory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationApPatchHistoryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApPatchHistoryList
}

func newWSGAdministrationApPatchHistoryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApPatchHistoryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApPatchHistoryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationApPatchHistoryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationApPatchHistoryList() *WSGAdministrationApPatchHistoryList {
	m := new(WSGAdministrationApPatchHistoryList)
	return m
}

// WSGAdministrationApPatchInfo
//
// Definition: administration_apPatchInfo
type WSGAdministrationApPatchInfo struct {
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

type WSGAdministrationApPatchInfoAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApPatchInfo
}

func newWSGAdministrationApPatchInfoAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApPatchInfoAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApPatchInfoAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationApPatchInfo)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationApPatchInfo() *WSGAdministrationApPatchInfo {
	m := new(WSGAdministrationApPatchInfo)
	return m
}

// WSGAdministrationApPatchStatus
//
// Definition: administration_apPatchStatus
type WSGAdministrationApPatchStatus struct {
	ClusterOperationProgress *WSGClusterBladeClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
}

type WSGAdministrationApPatchStatusAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApPatchStatus
}

func newWSGAdministrationApPatchStatusAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApPatchStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApPatchStatusAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationApPatchStatus)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationApPatchStatus() *WSGAdministrationApPatchStatus {
	m := new(WSGAdministrationApPatchStatus)
	return m
}

// WSGAdministrationApplicationLogAndStatus
//
// Definition: administration_applicationLogAndStatus
type WSGAdministrationApplicationLogAndStatus struct {
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

func NewWSGAdministrationApplicationLogAndStatus() *WSGAdministrationApplicationLogAndStatus {
	m := new(WSGAdministrationApplicationLogAndStatus)
	return m
}

// WSGAdministrationApplicationLogAndStatusList
//
// Definition: administration_applicationLogAndStatusList
type WSGAdministrationApplicationLogAndStatusList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationApplicationLogAndStatus `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationApplicationLogAndStatusListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApplicationLogAndStatusList
}

func newWSGAdministrationApplicationLogAndStatusListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApplicationLogAndStatusListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApplicationLogAndStatusListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationApplicationLogAndStatusList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationApplicationLogAndStatusList() *WSGAdministrationApplicationLogAndStatusList {
	m := new(WSGAdministrationApplicationLogAndStatusList)
	return m
}

// WSGAdministrationAutoExportBackup
//
// Definition: administration_autoExportBackup
type WSGAdministrationAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// FTP server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

type WSGAdministrationAutoExportBackupAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationAutoExportBackup
}

func newWSGAdministrationAutoExportBackupAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationAutoExportBackupAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationAutoExportBackupAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationAutoExportBackup)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationAutoExportBackup() *WSGAdministrationAutoExportBackup {
	m := new(WSGAdministrationAutoExportBackup)
	return m
}

// WSGAdministrationBackupFile
//
// Definition: administration_backupFile
type WSGAdministrationBackupFile struct {
	// BackupElapsed
	// backup elapsed of the configuration backup file
	BackupElapsed *int `json:"backupElapsed,omitempty"`

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
	FileSize *int `json:"fileSize,omitempty"`

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

func NewWSGAdministrationBackupFile() *WSGAdministrationBackupFile {
	m := new(WSGAdministrationBackupFile)
	return m
}

// WSGAdministrationClusterBackupList
//
// Definition: administration_clusterBackupList
type WSGAdministrationClusterBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationClusterBackupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationClusterBackupListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationClusterBackupList
}

func newWSGAdministrationClusterBackupListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationClusterBackupListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationClusterBackupListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationClusterBackupList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationClusterBackupList() *WSGAdministrationClusterBackupList {
	m := new(WSGAdministrationClusterBackupList)
	return m
}

// WSGAdministrationClusterBackupSummary
//
// Definition: administration_clusterBackupSummary
type WSGAdministrationClusterBackupSummary struct {
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

func NewWSGAdministrationClusterBackupSummary() *WSGAdministrationClusterBackupSummary {
	m := new(WSGAdministrationClusterBackupSummary)
	return m
}

// WSGAdministrationConfigurationBackupList
//
// Definition: administration_configurationBackupList
type WSGAdministrationConfigurationBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationBackupFile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationConfigurationBackupListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationConfigurationBackupList
}

func newWSGAdministrationConfigurationBackupListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationConfigurationBackupListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationConfigurationBackupListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationConfigurationBackupList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationConfigurationBackupList() *WSGAdministrationConfigurationBackupList {
	m := new(WSGAdministrationConfigurationBackupList)
	return m
}

// WSGAdministrationConnectZD
//
// Definition: administration_connectZD
type WSGAdministrationConnectZD struct {
	Ip *string `json:"ip,omitempty"`

	Password *string `json:"password,omitempty"`

	User *string `json:"user,omitempty"`
}

func NewWSGAdministrationConnectZD() *WSGAdministrationConnectZD {
	m := new(WSGAdministrationConnectZD)
	return m
}

// WSGAdministrationCreateAdminAAAServer
//
// Definition: administration_createAdminAAAServer
type WSGAdministrationCreateAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type"`
}

func NewWSGAdministrationCreateAdminAAAServer() *WSGAdministrationCreateAdminAAAServer {
	m := new(WSGAdministrationCreateAdminAAAServer)
	return m
}

// WSGAdministrationDefaultRoleMapping
//
// Definition: administration_defaultRoleMapping
type WSGAdministrationDefaultRoleMapping struct {
	// DefaultAdmin
	// DefaultAdmin of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultAdmin *string `json:"defaultAdmin"`

	// DefaultUserGroup
	// DefaultUserGroup of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultUserGroup *string `json:"defaultUserGroup"`
}

func NewWSGAdministrationDefaultRoleMapping() *WSGAdministrationDefaultRoleMapping {
	m := new(WSGAdministrationDefaultRoleMapping)
	return m
}

// WSGAdministrationLdapServer
//
// Definition: administration_ldapServer
type WSGAdministrationLdapServer struct {
	// AdminDomainName
	// Admin Domain Name of LDAP Server object
	// Constraints:
	//    - required
	AdminDomainName *string `json:"adminDomainName"`

	// AdminPassword
	// Admin password of LDAP Server object
	// Constraints:
	//    - required
	AdminPassword *string `json:"adminPassword"`

	// BaseDomainName
	// Base Domain Name of LDAP Server object
	// Constraints:
	//    - required
	BaseDomainName *string `json:"baseDomainName"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// KeyAttribute
	// Key attribute of LDAP Server object
	// Constraints:
	//    - required
	KeyAttribute *string `json:"keyAttribute"`

	// Port
	// Port number of LDAP Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// SearchFilter
	// Search filter of LDAP Server object
	// Constraints:
	//    - required
	SearchFilter *string `json:"searchFilter"`
}

func NewWSGAdministrationLdapServer() *WSGAdministrationLdapServer {
	m := new(WSGAdministrationLdapServer)
	return m
}

// WSGAdministrationLicenses
//
// Definition: administration_licenses
type WSGAdministrationLicenses struct {
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

func NewWSGAdministrationLicenses() *WSGAdministrationLicenses {
	m := new(WSGAdministrationLicenses)
	return m
}

// WSGAdministrationLicenseServer
//
// Definition: administration_licenseServer
type WSGAdministrationLicenseServer struct {
	// IpAddress
	// local license server IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// local license server port
	// Constraints:
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// UseCloud
	// use cloud license server
	UseCloud *bool `json:"useCloud,omitempty"`
}

type WSGAdministrationLicenseServerAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicenseServer
}

func newWSGAdministrationLicenseServerAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicenseServerAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicenseServerAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationLicenseServer)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationLicenseServer() *WSGAdministrationLicenseServer {
	m := new(WSGAdministrationLicenseServer)
	return m
}

// WSGAdministrationLicensesList
//
// Definition: administration_licensesList
type WSGAdministrationLicensesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicenses `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicensesList
}

func newWSGAdministrationLicensesListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicensesListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicensesListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationLicensesList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationLicensesList() *WSGAdministrationLicensesList {
	m := new(WSGAdministrationLicensesList)
	return m
}

// WSGAdministrationLicensesSummary
//
// Definition: administration_licensesSummary
type WSGAdministrationLicensesSummary struct {
	CapacityControlLicenseCount *WSGAdministrationLicensesSummaryCapacityControlLicenseCountType `json:"capacityControlLicenseCount,omitempty"`

	// LicenseTypeDescription
	// license type description
	LicenseTypeDescription *string `json:"licenseTypeDescription,omitempty"`
}

func NewWSGAdministrationLicensesSummary() *WSGAdministrationLicensesSummary {
	m := new(WSGAdministrationLicensesSummary)
	return m
}

// WSGAdministrationLicensesSummaryCapacityControlLicenseCountType
//
// Definition: administration_licensesSummaryCapacityControlLicenseCountType
type WSGAdministrationLicensesSummaryCapacityControlLicenseCountType struct {
	// TotalCount
	// total count of licenses
	TotalCount *int `json:"totalCount,omitempty"`

	// UsedCount
	// consumed count of licenses
	UsedCount *int `json:"usedCount,omitempty"`
}

func NewWSGAdministrationLicensesSummaryCapacityControlLicenseCountType() *WSGAdministrationLicensesSummaryCapacityControlLicenseCountType {
	m := new(WSGAdministrationLicensesSummaryCapacityControlLicenseCountType)
	return m
}

// WSGAdministrationLicensesSummaryList
//
// Definition: administration_licensesSummaryList
type WSGAdministrationLicensesSummaryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicensesSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesSummaryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicensesSummaryList
}

func newWSGAdministrationLicensesSummaryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicensesSummaryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicensesSummaryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationLicensesSummaryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationLicensesSummaryList() *WSGAdministrationLicensesSummaryList {
	m := new(WSGAdministrationLicensesSummaryList)
	return m
}

// WSGAdministrationLicensesSyncLogs
//
// Definition: administration_licensesSyncLogs
type WSGAdministrationLicensesSyncLogs struct {
	// CreateDateTime
	// license sync log's create time
	CreateDateTime *string `json:"createDateTime,omitempty"`

	// SyncResult
	// sync license result
	// Constraints:
	//    - oneof:[SUCCESS,FAILURE]
	SyncResult *string `json:"syncResult,omitempty"`
}

func NewWSGAdministrationLicensesSyncLogs() *WSGAdministrationLicensesSyncLogs {
	m := new(WSGAdministrationLicensesSyncLogs)
	return m
}

// WSGAdministrationLicensesSyncLogsList
//
// Definition: administration_licensesSyncLogsList
type WSGAdministrationLicensesSyncLogsList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicensesSyncLogs `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesSyncLogsListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicensesSyncLogsList
}

func newWSGAdministrationLicensesSyncLogsListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicensesSyncLogsListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicensesSyncLogsListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationLicensesSyncLogsList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationLicensesSyncLogsList() *WSGAdministrationLicensesSyncLogsList {
	m := new(WSGAdministrationLicensesSyncLogsList)
	return m
}

// WSGAdministrationModfiyLicenseServer
//
// Definition: administration_modfiyLicenseServer
type WSGAdministrationModfiyLicenseServer struct {
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// Constraints:
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// UseCloud
	// Constraints:
	//    - required
	UseCloud *bool `json:"useCloud"`
}

func NewWSGAdministrationModfiyLicenseServer() *WSGAdministrationModfiyLicenseServer {
	m := new(WSGAdministrationModfiyLicenseServer)
	return m
}

// WSGAdministrationModifyAdminAAAServer
//
// Definition: administration_modifyAdminAAAServer
type WSGAdministrationModifyAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type"`
}

func NewWSGAdministrationModifyAdminAAAServer() *WSGAdministrationModifyAdminAAAServer {
	m := new(WSGAdministrationModifyAdminAAAServer)
	return m
}

// WSGAdministrationModifyAutoExportBackup
//
// Definition: administration_modifyAutoExportBackup
type WSGAdministrationModifyAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// ftp server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

func NewWSGAdministrationModifyAutoExportBackup() *WSGAdministrationModifyAutoExportBackup {
	m := new(WSGAdministrationModifyAutoExportBackup)
	return m
}

// WSGAdministrationModifyLogLevel
//
// Definition: administration_modifyLogLevel
type WSGAdministrationModifyLogLevel struct {
	// ApplicationName
	// Application name.
	ApplicationName *string `json:"applicationName,omitempty"`

	// LogLevel
	// Log level.
	// Constraints:
	//    - oneof:[DEBUG,INFO,WARN,ERROR]
	LogLevel *string `json:"logLevel,omitempty"`
}

func NewWSGAdministrationModifyLogLevel() *WSGAdministrationModifyLogLevel {
	m := new(WSGAdministrationModifyLogLevel)
	return m
}

// WSGAdministrationModifyScheduleBackup
//
// Definition: administration_modifyScheduleBackup
type WSGAdministrationModifyScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

func NewWSGAdministrationModifyScheduleBackup() *WSGAdministrationModifyScheduleBackup {
	m := new(WSGAdministrationModifyScheduleBackup)
	return m
}

// WSGAdministrationRadiusServer
//
// Definition: administration_radiusServer
type WSGAdministrationRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Port
	// Port number of RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Protocol
	// Constraints:
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP,PEAP]
	Protocol *string `json:"protocol,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	SecondaryRadiusServer *WSGAdministrationSecondaryRadiusServer `json:"secondaryRadiusServer,omitempty"`

	// SharedSecret
	// Shared secret of RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGAdministrationRadiusServer() *WSGAdministrationRadiusServer {
	m := new(WSGAdministrationRadiusServer)
	return m
}

// WSGAdministrationRetrieveAdminAAAServer
//
// Definition: administration_retrieveAdminAAAServer
type WSGAdministrationRetrieveAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	// Id
	// ID of this Admin AAA Server
	Id *string `json:"id,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty"`
}

type WSGAdministrationRetrieveAdminAAAServerAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationRetrieveAdminAAAServer
}

func newWSGAdministrationRetrieveAdminAAAServerAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationRetrieveAdminAAAServerAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationRetrieveAdminAAAServerAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationRetrieveAdminAAAServer)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationRetrieveAdminAAAServer() *WSGAdministrationRetrieveAdminAAAServer {
	m := new(WSGAdministrationRetrieveAdminAAAServer)
	return m
}

// WSGAdministrationRetrieveAdminAAAServerList
//
// Definition: administration_retrieveAdminAAAServerList
type WSGAdministrationRetrieveAdminAAAServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationRetrieveAdminAAAServerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationRetrieveAdminAAAServerListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationRetrieveAdminAAAServerList
}

func newWSGAdministrationRetrieveAdminAAAServerListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationRetrieveAdminAAAServerListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationRetrieveAdminAAAServerListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationRetrieveAdminAAAServerList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationRetrieveAdminAAAServerList() *WSGAdministrationRetrieveAdminAAAServerList {
	m := new(WSGAdministrationRetrieveAdminAAAServerList)
	return m
}

// WSGAdministrationRetrieveAdminAAAServerListType
//
// Definition: administration_retrieveAdminAAAServerListType
type WSGAdministrationRetrieveAdminAAAServerListType struct {
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty"`
}

func NewWSGAdministrationRetrieveAdminAAAServerListType() *WSGAdministrationRetrieveAdminAAAServerListType {
	m := new(WSGAdministrationRetrieveAdminAAAServerListType)
	return m
}

// WSGAdministrationScheduleBackup
//
// Definition: administration_scheduleBackup
type WSGAdministrationScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type WSGAdministrationScheduleBackupAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationScheduleBackup
}

func newWSGAdministrationScheduleBackupAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationScheduleBackupAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationScheduleBackupAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationScheduleBackup)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationScheduleBackup() *WSGAdministrationScheduleBackup {
	m := new(WSGAdministrationScheduleBackup)
	return m
}

// WSGAdministrationSecondaryRadiusServer
//
// Definition: administration_secondaryRadiusServer
type WSGAdministrationSecondaryRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// MaxRetries
	// Max number(how many times) of retries for re-connection to primary
	// Constraints:
	//    - required
	//    - default:2
	//    - min:2
	//    - max:10
	MaxRetries *int `json:"maxRetries"`

	// Port
	// Port number of Secondary RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Protocol
	// Constraints:
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP,PEAP]
	Protocol *string `json:"protocol,omitempty"`

	// RequestTimeOut
	// Request timeout(seconds) value of re-connection to primary
	// Constraints:
	//    - required
	//    - default:3
	//    - min:2
	//    - max:20
	RequestTimeOut *int `json:"requestTimeOut"`

	// RetryPriInvl
	// Interval of re-connection to primary(1-60 minute)
	// Constraints:
	//    - required
	//    - default:5
	//    - min:1
	//    - max:60
	RetryPriInvl *int `json:"retryPriInvl"`

	// SharedSecret
	// Shared secret of Secondary RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGAdministrationSecondaryRadiusServer() *WSGAdministrationSecondaryRadiusServer {
	m := new(WSGAdministrationSecondaryRadiusServer)
	return m
}

// WSGAdministrationTacacsServer
//
// Definition: administration_tacacsServer
type WSGAdministrationTacacsServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Port
	// Port number of TACACS+ Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Service
	// Constraints:
	//    - required
	Service *WSGCommonRealm `json:"service"`

	// SharedSecret
	// Shared secret of TACACS+ Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGAdministrationTacacsServer() *WSGAdministrationTacacsServer {
	m := new(WSGAdministrationTacacsServer)
	return m
}

// WSGAdministrationUpgradeHistoryList
//
// Definition: administration_upgradeHistoryList
type WSGAdministrationUpgradeHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationUpgradeHistorySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationUpgradeHistoryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationUpgradeHistoryList
}

func newWSGAdministrationUpgradeHistoryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationUpgradeHistoryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationUpgradeHistoryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationUpgradeHistoryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationUpgradeHistoryList() *WSGAdministrationUpgradeHistoryList {
	m := new(WSGAdministrationUpgradeHistoryList)
	return m
}

// WSGAdministrationUpgradeHistorySummary
//
// Definition: administration_upgradeHistorySummary
type WSGAdministrationUpgradeHistorySummary struct {
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
	ElapsedSeconds *int `json:"elapsedSeconds,omitempty"`

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

func NewWSGAdministrationUpgradeHistorySummary() *WSGAdministrationUpgradeHistorySummary {
	m := new(WSGAdministrationUpgradeHistorySummary)
	return m
}

// WSGAdministrationUpgradePatchInfo
//
// Definition: administration_upgradePatchInfo
type WSGAdministrationUpgradePatchInfo struct {
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`

	UploadPatchInfo *WSGClusterBladeUploadPatchInfo `json:"uploadPatchInfo,omitempty"`
}

type WSGAdministrationUpgradePatchInfoAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationUpgradePatchInfo
}

func newWSGAdministrationUpgradePatchInfoAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationUpgradePatchInfoAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationUpgradePatchInfoAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationUpgradePatchInfo)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationUpgradePatchInfo() *WSGAdministrationUpgradePatchInfo {
	m := new(WSGAdministrationUpgradePatchInfo)
	return m
}

// WSGAdministrationUpgradeStatus
//
// Definition: administration_upgradeStatus
type WSGAdministrationUpgradeStatus struct {
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`
}

type WSGAdministrationUpgradeStatusAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationUpgradeStatus
}

func newWSGAdministrationUpgradeStatusAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationUpgradeStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationUpgradeStatusAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationUpgradeStatus)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationUpgradeStatus() *WSGAdministrationUpgradeStatus {
	m := new(WSGAdministrationUpgradeStatus)
	return m
}

// WSGAdministrationZdAP
//
// Definition: administration_zdAP
type WSGAdministrationZdAP struct {
	// Connected
	// AP Conntected
	Connected *string `json:"connected,omitempty"`

	// Mac
	// AP MAC
	Mac *string `json:"mac,omitempty"`
}

func NewWSGAdministrationZdAP() *WSGAdministrationZdAP {
	m := new(WSGAdministrationZdAP)
	return m
}

// WSGAdministrationZdAPList
//
// Definition: administration_zdAPList
type WSGAdministrationZdAPList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationZdAP `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationZdAPListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationZdAPList
}

func newWSGAdministrationZdAPListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationZdAPListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationZdAPListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationZdAPList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationZdAPList() *WSGAdministrationZdAPList {
	m := new(WSGAdministrationZdAPList)
	return m
}

// WSGAdministrationZdImport
//
// Definition: administration_zdImport
type WSGAdministrationZdImport struct {
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

func NewWSGAdministrationZdImport() *WSGAdministrationZdImport {
	m := new(WSGAdministrationZdImport)
	return m
}

// WSGAdministrationZdImportStatus
//
// Definition: administration_zdImportStatus
type WSGAdministrationZdImportStatus struct {
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

type WSGAdministrationZdImportStatusAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationZdImportStatus
}

func newWSGAdministrationZdImportStatusAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationZdImportStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationZdImportStatusAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAdministrationZdImportStatus)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAdministrationZdImportStatus() *WSGAdministrationZdImportStatus {
	m := new(WSGAdministrationZdImportStatus)
	return m
}

// AddAdminaaa
//
// Use this API command to create a new Admin AAA server
//
// Operation ID: addAdminaaa
// Operation path: /adminaaa
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAdministrationCreateAdminAAAServer
func (s *WSGAdministrationService) AddAdminaaa(ctx context.Context, body *WSGAdministrationCreateAdminAAAServer, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddAdminaaa, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRestart
//
// Use this API command to restart the controller.
//
// Operation ID: addRestart
// Operation path: /restart
// Success code: 204 (No Content)
func (s *WSGAdministrationService) AddRestart(ctx context.Context, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRestart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// AddShutdown
//
// Use this API command to shut down the controller.
//
// Operation ID: addShutdown
// Operation path: /shutdown
// Success code: 204 (No Content)
func (s *WSGAdministrationService) AddShutdown(ctx context.Context, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddShutdown, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteAdminaaaById
//
// Use this API command to delete an existing Admin AAA server
//
// Operation ID: deleteAdminaaaById
// Operation path: /adminaaa/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAdministrationService) DeleteAdminaaaById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteAdminaaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindAdminaaa
//
// Use this API command to retrieve the list of Admin AAA server
//
// Operation ID: findAdminaaa
// Operation path: /adminaaa
// Success code: 200 (OK)
//
// Required parameters:
// - type_ string
//		- required
func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context, type_ string, mutators ...RequestMutator) (*WSGAdministrationRetrieveAdminAAAServerListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationRetrieveAdminAAAServerListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindAdminaaa, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("type", type_)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAdministrationRetrieveAdminAAAServerListAPIResponse), err
}

// FindAdminaaaById
//
// Use this API command to retrieve an existing Admin AAA server
//
// Operation ID: findAdminaaaById
// Operation path: /adminaaa/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAdministrationService) FindAdminaaaById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAdministrationRetrieveAdminAAAServerAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationRetrieveAdminAAAServerAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindAdminaaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAdministrationRetrieveAdminAAAServerAPIResponse), err
}

// FindLicenses
//
// Use this API command to get all licenses currently assign in SCG.
//
// Operation ID: findLicenses
// Operation path: /licenses
// Success code: 200 (OK)
func (s *WSGAdministrationService) FindLicenses(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicensesListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicensesListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindLicenses, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAdministrationLicensesListAPIResponse), err
}

// FindLicenseServer
//
// Use this API command to get license server configuration.
//
// Operation ID: findLicenseServer
// Operation path: /licenseServer
// Success code: 200 (OK)
func (s *WSGAdministrationService) FindLicenseServer(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicenseServerAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicenseServerAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindLicenseServer, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAdministrationLicenseServerAPIResponse), err
}

// FindLicensesSummary
//
// Use this API command to get licenses summary information.
//
// Operation ID: findLicensesSummary
// Operation path: /licensesSummary
// Success code: 200 (OK)
func (s *WSGAdministrationService) FindLicensesSummary(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicensesSummaryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicensesSummaryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindLicensesSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAdministrationLicensesSummaryListAPIResponse), err
}

// FindLicensesSyncLogs
//
// Use this API command to get licenses synchronize logs.
//
// Operation ID: findLicensesSyncLogs
// Operation path: /licensesSyncLogs
// Success code: 200 (OK)
func (s *WSGAdministrationService) FindLicensesSyncLogs(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicensesSyncLogsListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicensesSyncLogsListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindLicensesSyncLogs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAdministrationLicensesSyncLogsListAPIResponse), err
}

// UpdateAdminaaaById
//
// Use this API command to modify an existing Admin AAA server
//
// Operation ID: updateAdminaaaById
// Operation path: /adminaaa/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAdministrationModifyAdminAAAServer
//
// Required parameters:
// - id string
//		- required
func (s *WSGAdministrationService) UpdateAdminaaaById(ctx context.Context, body *WSGAdministrationModifyAdminAAAServer, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateAdminaaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateLicenseServer
//
// Use this API command to update license server configuration.
//
// Operation ID: updateLicenseServer
// Operation path: /licenseServer
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAdministrationModfiyLicenseServer
func (s *WSGAdministrationService) UpdateLicenseServer(ctx context.Context, body *WSGAdministrationModfiyLicenseServer, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateLicenseServer, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateLicensesSync
//
// Use this API command to ask all SCG in cluster to sync licenses from license server.
//
// Operation ID: updateLicensesSync
// Operation path: /licenses/sync
// Success code: 204 (No Content)
func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateLicensesSync, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
