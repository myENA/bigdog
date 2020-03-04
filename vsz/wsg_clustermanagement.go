package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
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
func (s *WSGClusterManagementService) AddApPatch(ctx context.Context) (*WSGAdministrationApPatchStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationApPatchStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddApPatch), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApPatchStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddApPatchFile
//
// Use this API command to upload AP Patch File.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddApPatchFile(ctx context.Context, body []byte) (*APIResponseMeta, error) {
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
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddApPatchFile), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddClusterBackup
//
// Backup cluster.
func (s *WSGClusterManagementService) AddClusterBackup(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddClusterBackup), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddClusterRestoreById
//
// Restore cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddClusterRestoreById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddClusterRestoreById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddConfigurationBackup
//
// Backup system configuration.
func (s *WSGClusterManagementService) AddConfigurationBackup(ctx context.Context) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddConfigurationBackup), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// AddConfigurationRestoreById
//
// Restore system configuration with specified backupUUID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddConfigurationRestoreById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddConfigurationRestoreById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddConfigurationUpload
//
// Upload system configuration file.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddConfigurationUpload(ctx context.Context, body []byte) (*APIResponseMeta, error) {
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
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddConfigurationUpload), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddUpgrade
//
// Use this API command to do system upgrade.
func (s *WSGClusterManagementService) AddUpgrade(ctx context.Context) (*WSGAdministrationUpgradeStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationUpgradeStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddUpgrade), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradeStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddUpgradeUpload
//
// Use this API command to upload patch file.
//
// Request Body:
//	 - body []byte
func (s *WSGClusterManagementService) AddUpgradeUpload(ctx context.Context, body []byte) (*WSGAdministrationUpgradeStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationUpgradeStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddUpgradeUpload), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradeStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteClusterById
//
// Delete cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteClusterById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGDeleteClusterById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteConfigurationById
//
// Delete system configuration file.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteConfigurationById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGDeleteConfigurationById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApPatch
//
// Use this API command to retrive uploaded AP patch file info.
func (s *WSGClusterManagementService) FindApPatch(ctx context.Context) (*WSGAdministrationApPatchInfo, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationApPatchInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindApPatch), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApPatchInfo()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGClusterManagementService) FindApPatchHistory(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationApPatchHistoryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationApPatchHistoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindApPatchHistory), true)
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindApPatchStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindApPatchStatus(ctx context.Context) (*WSGAdministrationApPatchStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationApPatchStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindApPatchStatus), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApPatchStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGClusterManagementService) FindCluster(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationClusterBackupList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationClusterBackupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindCluster), true)
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindClusterGeoRedundancy
//
// Get cluster redundancy settings.
func (s *WSGClusterManagementService) FindClusterGeoRedundancy(ctx context.Context) (*WSGClusterRedundancySettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClusterRedundancySettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindClusterGeoRedundancy), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClusterRedundancySettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindClusterState
//
// Use this API command to get current cluster, blade, and management service states
func (s *WSGClusterManagementService) FindClusterState(ctx context.Context) (*WSGClusterBladeClusterState, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClusterBladeClusterState
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindClusterState), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClusterBladeClusterState()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGClusterManagementService) FindConfiguration(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationConfigurationBackupList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationConfigurationBackupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindConfiguration), true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationConfigurationBackupList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGClusterManagementService) FindConfigurationDownload(ctx context.Context, backupUUID string, optionalParams map[string][]string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, backupUUID, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindConfigurationDownload), true)
	req.SetQueryParameter("backupUUID", []string{backupUUID})
	if v, ok := optionalParams["timeZone"]; ok {
		req.AddQueryParameter("timeZone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindConfigurationSettingsAutoExportBackup
//
// Get Auto Export Backup Settings.
func (s *WSGClusterManagementService) FindConfigurationSettingsAutoExportBackup(ctx context.Context) (*WSGAdministrationAutoExportBackup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationAutoExportBackup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindConfigurationSettingsAutoExportBackup), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationAutoExportBackup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindConfigurationSettingsScheduleBackup
//
// Get Schedule Backup Setting.
func (s *WSGClusterManagementService) FindConfigurationSettingsScheduleBackup(ctx context.Context) (*WSGAdministrationScheduleBackup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationScheduleBackup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindConfigurationSettingsScheduleBackup), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationScheduleBackup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *WSGClusterManagementService) FindUpgradeHistory(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationUpgradeHistoryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationUpgradeHistoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindUpgradeHistory), true)
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindUpgradePatch
//
// Use this API command to retrive upload file Info.
func (s *WSGClusterManagementService) FindUpgradePatch(ctx context.Context) (*WSGAdministrationUpgradePatchInfo, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationUpgradePatchInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindUpgradePatch), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradePatchInfo()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindUpgradeStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindUpgradeStatus(ctx context.Context) (*WSGAdministrationUpgradeStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationUpgradeStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindUpgradeStatus), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationUpgradeStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateConfigurationSettingsAutoExportBackup
//
// Modify Auto Export Backup Settings.
//
// Request Body:
//	 - body *WSGAdministrationModifyAutoExportBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsAutoExportBackup(ctx context.Context, body *WSGAdministrationModifyAutoExportBackup) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateConfigurationSettingsAutoExportBackup), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateConfigurationSettingsScheduleBackup
//
// Modify Schedule Backup Setting.
//
// Request Body:
//	 - body *WSGAdministrationModifyScheduleBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsScheduleBackup(ctx context.Context, body *WSGAdministrationModifyScheduleBackup) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateConfigurationSettingsScheduleBackup), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateClusterGeoRedundancy
//
// Update cluster redundancy settings.
//
// Request Body:
//	 - body *WSGClusterRedundancyUpdateClusterRedundancy
func (s *WSGClusterManagementService) UpdateClusterGeoRedundancy(ctx context.Context, body *WSGClusterRedundancyUpdateClusterRedundancy) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
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
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGUpdateClusterGeoRedundancy), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
