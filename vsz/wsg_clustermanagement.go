package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddApPatch, true)
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddApPatchFile, true)
}

// AddClusterBackup
//
// Backup cluster.
func (s *WSGClusterManagementService) AddClusterBackup(ctx context.Context) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodPost, RouteWSGAddClusterBackup, true)
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddClusterRestoreById, true)
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddConfigurationBackup, true)
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddConfigurationRestoreById, true)
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddConfigurationUpload, true)
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddUpgrade, true)
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddUpgradeUpload, true)
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteClusterById, true)
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteConfigurationById, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindApPatch, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindApPatchHistory, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindApPatchStatus, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindCluster, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindClusterGeoRedundancy, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindClusterNodeStatus, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindClusterState, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindClusterStatus, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindConfiguration, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindConfigurationDownload, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindConfigurationSettingsAutoExportBackup, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindConfigurationSettingsScheduleBackup, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUpgradeHistory, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUpgradePatch, true)
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
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUpgradeStatus, true)
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsAutoExportBackup, true)
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsScheduleBackup, true)
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
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateClusterGeoRedundancy, true)
}
