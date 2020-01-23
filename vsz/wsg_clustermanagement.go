package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
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
	return NewWSGClusterManagementService(ss.apiClient)
}

// AddApPatch
//
// Use this API command to apply AP patch.
func (s *WSGClusterManagementService) AddApPatch(ctx context.Context) (*WSGAdministrationApPatchStatus, error) {
	var (
		resp *WSGAdministrationApPatchStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// AddApPatchFile
//
// Use this API command to upload AP Patch File.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddApPatchFile(ctx context.Context, body []byte) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	}
}

// AddClusterBackup
//
// Backup cluster.
func (s *WSGClusterManagementService) AddClusterBackup(ctx context.Context) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
}

// AddClusterRestoreById
//
// Restore cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddClusterRestoreById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
}

// AddConfigurationBackup
//
// Backup system configuration.
func (s *WSGClusterManagementService) AddConfigurationBackup(ctx context.Context) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// AddConfigurationRestoreById
//
// Restore system configuration with specified backupUUID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddConfigurationRestoreById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
}

// AddConfigurationUpload
//
// Upload system configuration file.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddConfigurationUpload(ctx context.Context, body []byte) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	}
}

// AddUpgrade
//
// Use this API command to do system upgrade.
func (s *WSGClusterManagementService) AddUpgrade(ctx context.Context) (*WSGAdministrationUpgradeStatus, error) {
	var (
		resp *WSGAdministrationUpgradeStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// AddUpgradeUpload
//
// Use this API command to upload patch file.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddUpgradeUpload(ctx context.Context, body []byte) (*WSGAdministrationUpgradeStatus, error) {
	var (
		resp *WSGAdministrationUpgradeStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
}

// DeleteClusterById
//
// Delete cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteClusterById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
}

// DeleteConfigurationById
//
// Delete system configuration file.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteConfigurationById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
}

// FindApPatch
//
// Use this API command to retrive uploaded AP patch file info.
func (s *WSGClusterManagementService) FindApPatch(ctx context.Context) (*WSGAdministrationApPatchInfo, error) {
	var (
		resp *WSGAdministrationApPatchInfo
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindApPatchHistory
//
// Use this API command to retrive AP patch history.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
// - timezone string
//		- nullable
func (s *WSGClusterManagementService) FindApPatchHistory(ctx context.Context, optionalParams map[string]interface{}) (*WSGAdministrationApPatchHistoryList, error) {
	var (
		resp *WSGAdministrationApPatchHistoryList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindApPatchStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindApPatchStatus(ctx context.Context) (*WSGAdministrationApPatchStatus, error) {
	var (
		resp *WSGAdministrationApPatchStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindCluster
//
// Retrive cluster backup list.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
// - timezone string
//		- nullable
func (s *WSGClusterManagementService) FindCluster(ctx context.Context, optionalParams map[string]interface{}) (*WSGAdministrationClusterBackupList, error) {
	var (
		resp *WSGAdministrationClusterBackupList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindClusterGeoRedundancy
//
// Get cluster redundancy settings.
func (s *WSGClusterManagementService) FindClusterGeoRedundancy(ctx context.Context) (*WSGClusterRedundancySettings, error) {
	var (
		resp *WSGClusterRedundancySettings
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindClusterNodeStatus
//
// Use this API command to get Control node Status.
func (s *WSGClusterManagementService) FindClusterNodeStatus(ctx context.Context) (*WSGClusterBladeControlNodeStatus, error) {
	var (
		resp *WSGClusterBladeControlNodeStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindClusterState
//
// Use this API command to get current cluster, blade, and management service states
func (s *WSGClusterManagementService) FindClusterState(ctx context.Context) (*WSGClusterBladeClusterState, error) {
	var (
		resp *WSGClusterBladeClusterState
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindClusterStatus
//
// Use this API command to get Cluster Status.
func (s *WSGClusterManagementService) FindClusterStatus(ctx context.Context) (*WSGClusterBladeClusterStatus, error) {
	var (
		resp *WSGClusterBladeClusterStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindConfiguration
//
// Retrive system configuration list.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGClusterManagementService) FindConfiguration(ctx context.Context, optionalParams map[string]interface{}) (*WSGAdministrationConfigurationBackupList, error) {
	var (
		resp *WSGAdministrationConfigurationBackupList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindConfigurationDownload
//
// Download system configuration file.
//
// Required Parameters:
// - backupUUID string
//		- required
//
// Optional Parameters:
// - timeZone string
//		- nullable
func (s *WSGClusterManagementService) FindConfigurationDownload(ctx context.Context, backupUUID string, optionalParams map[string]interface{}) ([]byte, error) {
	var (
		resp []byte
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, backupUUID, "required"); err != nil {
		return resp, err
	}
}

// FindConfigurationSettingsAutoExportBackup
//
// Get Auto Export Backup Settings.
func (s *WSGClusterManagementService) FindConfigurationSettingsAutoExportBackup(ctx context.Context) (*WSGAdministrationAutoExportBackup, error) {
	var (
		resp *WSGAdministrationAutoExportBackup
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindConfigurationSettingsScheduleBackup
//
// Get Schedule Backup Setting.
func (s *WSGClusterManagementService) FindConfigurationSettingsScheduleBackup(ctx context.Context) (*WSGAdministrationScheduleBackup, error) {
	var (
		resp *WSGAdministrationScheduleBackup
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindUpgradeHistory
//
// Use this API command to retrive upgrade history.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
// - timezone string
//		- nullable
func (s *WSGClusterManagementService) FindUpgradeHistory(ctx context.Context, optionalParams map[string]interface{}) (*WSGAdministrationUpgradeHistoryList, error) {
	var (
		resp *WSGAdministrationUpgradeHistoryList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindUpgradePatch
//
// Use this API command to retrive upload file Info.
func (s *WSGClusterManagementService) FindUpgradePatch(ctx context.Context) (*WSGAdministrationUpgradePatchInfo, error) {
	var (
		resp *WSGAdministrationUpgradePatchInfo
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindUpgradeStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindUpgradeStatus(ctx context.Context) (*WSGAdministrationUpgradeStatus, error) {
	var (
		resp *WSGAdministrationUpgradeStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// PartialUpdateConfigurationSettingsAutoExportBackup
//
// Modify Auto Export Backup Settings.
//
// Request Body:
//	 - body *WSGAdministrationModifyAutoExportBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsAutoExportBackup(ctx context.Context, body *WSGAdministrationModifyAutoExportBackup) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
}

// PartialUpdateConfigurationSettingsScheduleBackup
//
// Modify Schedule Backup Setting.
//
// Request Body:
//	 - body *WSGAdministrationModifyScheduleBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsScheduleBackup(ctx context.Context, body *WSGAdministrationModifyScheduleBackup) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
}

// UpdateClusterGeoRedundancy
//
// Update cluster redundancy settings.
//
// Request Body:
//	 - body *WSGClusterRedundancyUpdateClusterRedundancy
func (s *WSGClusterManagementService) UpdateClusterGeoRedundancy(ctx context.Context, body *WSGClusterRedundancyUpdateClusterRedundancy) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
}
