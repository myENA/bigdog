package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterblade"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterredundancy"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGClusterManagementService struct {
	apiClient *APIClient
}

func NewWSGClusterManagementService(c *APIClient) *WSGClusterManagementService {
	s := new(WSGClusterManagementService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClusterManagementService() *WSGClusterManagementService {
	serv := new(WSGClusterManagementService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddApPatch
//
// Use this API command to apply AP patch.
func (s *WSGClusterManagementService) AddApPatch(ctx context.Context) (*administration.ApPatchStatus, error) {
}

// AddApPatchFile
//
// Use this API command to upload AP Patch File.
func (s *WSGClusterManagementService) AddApPatchFile(ctx context.Context) error {
}

// AddClusterBackup
//
// Backup cluster.
func (s *WSGClusterManagementService) AddClusterBackup(ctx context.Context) error {
}

// AddClusterRestoreById
//
// Restore cluster backup by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGClusterManagementService) AddClusterRestoreById(ctx context.Context, pId string) error {
}

// AddConfigurationBackup
//
// Backup system configuration.
func (s *WSGClusterManagementService) AddConfigurationBackup(ctx context.Context) (*common.EmptyResult, error) {
}

// AddConfigurationRestoreById
//
// Restore system configuration with specified backupUUID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGClusterManagementService) AddConfigurationRestoreById(ctx context.Context, pId string) error {
}

// AddConfigurationUpload
//
// Upload system configuration file.
func (s *WSGClusterManagementService) AddConfigurationUpload(ctx context.Context) error {
}

// AddUpgrade
//
// Use this API command to do system upgrade.
func (s *WSGClusterManagementService) AddUpgrade(ctx context.Context) (*administration.UpgradeStatus, error) {
}

// AddUpgradeUpload
//
// Use this API command to upload patch file.
func (s *WSGClusterManagementService) AddUpgradeUpload(ctx context.Context) (*administration.UpgradeStatus, error) {
}

// DeleteClusterById
//
// Delete cluster backup by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGClusterManagementService) DeleteClusterById(ctx context.Context, pId string) error {
}

// DeleteConfigurationById
//
// Delete system configuration file.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGClusterManagementService) DeleteConfigurationById(ctx context.Context, pId string) error {
}

// FindApPatch
//
// Use this API command to retrive uploaded AP patch file info.
func (s *WSGClusterManagementService) FindApPatch(ctx context.Context) (*administration.ApPatchInfo, error) {
}

// FindApPatchHistory
//
// Use this API command to retrive AP patch history.
//
// Query Parameters:
// - qIndex string
// - qListSize string
// - qTimezone string
func (s *WSGClusterManagementService) FindApPatchHistory(ctx context.Context, qIndex string, qListSize string, qTimezone string) (*administration.ApPatchHistoryList, error) {
}

// FindApPatchStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindApPatchStatus(ctx context.Context) (*administration.ApPatchStatus, error) {
}

// FindCluster
//
// Retrive cluster backup list.
//
// Query Parameters:
// - qIndex string
// - qListSize string
// - qTimezone string
func (s *WSGClusterManagementService) FindCluster(ctx context.Context, qIndex string, qListSize string, qTimezone string) (*administration.ClusterBackupList, error) {
}

// FindClusterGeoRedundancy
//
// Get cluster redundancy settings.
func (s *WSGClusterManagementService) FindClusterGeoRedundancy(ctx context.Context) (*clusterredundancy.ClusterRedundancySettings, error) {
}

// FindClusterNodeStatus
//
// Use this API command to get Control node Status.
func (s *WSGClusterManagementService) FindClusterNodeStatus(ctx context.Context) (*clusterblade.ControlNodeStatus, error) {
}

// FindClusterState
//
// Use this API command to get current cluster, blade, and management service states
func (s *WSGClusterManagementService) FindClusterState(ctx context.Context) (*clusterblade.ClusterState, error) {
}

// FindClusterStatus
//
// Use this API command to get Cluster Status.
func (s *WSGClusterManagementService) FindClusterStatus(ctx context.Context) (*clusterblade.ClusterStatus, error) {
}

// FindConfiguration
//
// Retrive system configuration list.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGClusterManagementService) FindConfiguration(ctx context.Context, qIndex string, qListSize string) (*administration.ConfigurationBackupList, error) {
}

// FindConfigurationDownload
//
// Download system configuration file.
//
// Query Parameters:
// - qBackupUUID string
//		- required
// - qTimeZone string
func (s *WSGClusterManagementService) FindConfigurationDownload(ctx context.Context, qBackupUUID string, qTimeZone string) (json.RawMessage, error) {
}

// FindConfigurationSettingsAutoExportBackup
//
// Get Auto Export Backup Settings.
func (s *WSGClusterManagementService) FindConfigurationSettingsAutoExportBackup(ctx context.Context) (*administration.AutoExportBackup, error) {
}

// FindConfigurationSettingsScheduleBackup
//
// Get Schedule Backup Setting.
func (s *WSGClusterManagementService) FindConfigurationSettingsScheduleBackup(ctx context.Context) (*administration.ScheduleBackup, error) {
}

// FindUpgradeHistory
//
// Use this API command to retrive upgrade history.
//
// Query Parameters:
// - qIndex string
// - qListSize string
// - qTimezone string
func (s *WSGClusterManagementService) FindUpgradeHistory(ctx context.Context, qIndex string, qListSize string, qTimezone string) (*administration.UpgradeHistoryList, error) {
}

// FindUpgradePatch
//
// Use this API command to retrive upload file Info.
func (s *WSGClusterManagementService) FindUpgradePatch(ctx context.Context) (*administration.UpgradePatchInfo, error) {
}

// FindUpgradeStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindUpgradeStatus(ctx context.Context) (*administration.UpgradeStatus, error) {
}

// PartialUpdateConfigurationSettingsAutoExportBackup
//
// Modify Auto Export Backup Settings.
//
// Request Body:
//	 - body *administration.ModifyAutoExportBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsAutoExportBackup(ctx context.Context, body *administration.ModifyAutoExportBackup) (*common.EmptyResult, error) {
}

// PartialUpdateConfigurationSettingsScheduleBackup
//
// Modify Schedule Backup Setting.
//
// Request Body:
//	 - body *administration.ModifyScheduleBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsScheduleBackup(ctx context.Context, body *administration.ModifyScheduleBackup) (*common.EmptyResult, error) {
}

// UpdateClusterGeoRedundancy
//
// Update cluster redundancy settings.
//
// Request Body:
//	 - body *clusterredundancy.UpdateClusterRedundancy
func (s *WSGClusterManagementService) UpdateClusterGeoRedundancy(ctx context.Context, body *clusterredundancy.UpdateClusterRedundancy) error {
}
