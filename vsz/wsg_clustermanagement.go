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
		req      *APIRequest
		resp     *WSGAdministrationApPatchStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApPatch, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApPatchStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddApPatchFile
//
// Use this API command to upload AP Patch File.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddApPatchFile(ctx context.Context, body []byte) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApPatchFile, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// AddClusterBackup
//
// Backup cluster.
func (s *WSGClusterManagementService) AddClusterBackup(ctx context.Context) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClusterBackup, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// AddClusterRestoreById
//
// Restore cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddClusterRestoreById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClusterRestoreById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// AddConfigurationBackup
//
// Backup system configuration.
func (s *WSGClusterManagementService) AddConfigurationBackup(ctx context.Context) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddConfigurationBackup, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddConfigurationRestoreById
//
// Restore system configuration with specified backupUUID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddConfigurationRestoreById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddConfigurationRestoreById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// AddConfigurationUpload
//
// Upload system configuration file.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddConfigurationUpload(ctx context.Context, body []byte) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddConfigurationUpload, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// AddUpgrade
//
// Use this API command to do system upgrade.
func (s *WSGClusterManagementService) AddUpgrade(ctx context.Context) (*WSGAdministrationUpgradeStatus, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationUpgradeStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddUpgrade, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradeStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddUpgradeUpload
//
// Use this API command to upload patch file.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddUpgradeUpload(ctx context.Context, body []byte) (*WSGAdministrationUpgradeStatus, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationUpgradeStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddUpgradeUpload, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradeStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// DeleteClusterById
//
// Delete cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteClusterById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteClusterById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// DeleteConfigurationById
//
// Delete system configuration file.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteConfigurationById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteConfigurationById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindApPatch
//
// Use this API command to retrive uploaded AP patch file info.
func (s *WSGClusterManagementService) FindApPatch(ctx context.Context) (*WSGAdministrationApPatchInfo, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationApPatchInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApPatch, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApPatchInfo()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGClusterManagementService) FindApPatchHistory(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationApPatchHistoryList, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationApPatchHistoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApPatchHistory, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["timezone"]; ok {
		req.AddQueryParameter("timezone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApPatchHistoryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApPatchStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindApPatchStatus(ctx context.Context) (*WSGAdministrationApPatchStatus, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationApPatchStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApPatchStatus, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApPatchStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGClusterManagementService) FindCluster(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationClusterBackupList, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationClusterBackupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindCluster, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["timezone"]; ok {
		req.AddQueryParameter("timezone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationClusterBackupList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindClusterGeoRedundancy
//
// Get cluster redundancy settings.
func (s *WSGClusterManagementService) FindClusterGeoRedundancy(ctx context.Context) (*WSGClusterRedundancySettings, error) {
	var (
		req      *APIRequest
		resp     *WSGClusterRedundancySettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindClusterGeoRedundancy, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClusterRedundancySettings()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindClusterNodeStatus
//
// Use this API command to get Control node Status.
func (s *WSGClusterManagementService) FindClusterNodeStatus(ctx context.Context) (*WSGClusterBladeControlNodeStatus, error) {
	var (
		req      *APIRequest
		resp     *WSGClusterBladeControlNodeStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindClusterNodeStatus, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClusterBladeControlNodeStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindClusterState
//
// Use this API command to get current cluster, blade, and management service states
func (s *WSGClusterManagementService) FindClusterState(ctx context.Context) (*WSGClusterBladeClusterState, error) {
	var (
		req      *APIRequest
		resp     *WSGClusterBladeClusterState
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindClusterState, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClusterBladeClusterState()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindClusterStatus
//
// Use this API command to get Cluster Status.
func (s *WSGClusterManagementService) FindClusterStatus(ctx context.Context) (*WSGClusterBladeClusterStatus, error) {
	var (
		req      *APIRequest
		resp     *WSGClusterBladeClusterStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindClusterStatus, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClusterBladeClusterStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGClusterManagementService) FindConfiguration(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationConfigurationBackupList, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationConfigurationBackupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindConfiguration, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationConfigurationBackupList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGClusterManagementService) FindConfigurationDownload(ctx context.Context, backupUUID string, optionalParams map[string][]string) ([]byte, error) {
	var (
		req      *APIRequest
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, backupUUID, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindConfigurationDownload, true)
	req.SetQueryParameter("backupUUID", []string{backupUUID})
	if v, ok := optionalParams["timeZone"]; ok {
		req.AddQueryParameter("timeZone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// FindConfigurationSettingsAutoExportBackup
//
// Get Auto Export Backup Settings.
func (s *WSGClusterManagementService) FindConfigurationSettingsAutoExportBackup(ctx context.Context) (*WSGAdministrationAutoExportBackup, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationAutoExportBackup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindConfigurationSettingsAutoExportBackup, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationAutoExportBackup()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindConfigurationSettingsScheduleBackup
//
// Get Schedule Backup Setting.
func (s *WSGClusterManagementService) FindConfigurationSettingsScheduleBackup(ctx context.Context) (*WSGAdministrationScheduleBackup, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationScheduleBackup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindConfigurationSettingsScheduleBackup, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationScheduleBackup()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGClusterManagementService) FindUpgradeHistory(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationUpgradeHistoryList, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationUpgradeHistoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUpgradeHistory, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["timezone"]; ok {
		req.AddQueryParameter("timezone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradeHistoryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindUpgradePatch
//
// Use this API command to retrive upload file Info.
func (s *WSGClusterManagementService) FindUpgradePatch(ctx context.Context) (*WSGAdministrationUpgradePatchInfo, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationUpgradePatchInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUpgradePatch, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradePatchInfo()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindUpgradeStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindUpgradeStatus(ctx context.Context) (*WSGAdministrationUpgradeStatus, error) {
	var (
		req      *APIRequest
		resp     *WSGAdministrationUpgradeStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindUpgradeStatus, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradeStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateConfigurationSettingsAutoExportBackup
//
// Modify Auto Export Backup Settings.
//
// Request Body:
//	 - body *WSGAdministrationModifyAutoExportBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsAutoExportBackup(ctx context.Context, body *WSGAdministrationModifyAutoExportBackup) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsAutoExportBackup, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// PartialUpdateConfigurationSettingsScheduleBackup
//
// Modify Schedule Backup Setting.
//
// Request Body:
//	 - body *WSGAdministrationModifyScheduleBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsScheduleBackup(ctx context.Context, body *WSGAdministrationModifyScheduleBackup) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsScheduleBackup, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// UpdateClusterGeoRedundancy
//
// Update cluster redundancy settings.
//
// Request Body:
//	 - body *WSGClusterRedundancyUpdateClusterRedundancy
func (s *WSGClusterManagementService) UpdateClusterGeoRedundancy(ctx context.Context, body *WSGClusterRedundancyUpdateClusterRedundancy) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateClusterGeoRedundancy, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}
