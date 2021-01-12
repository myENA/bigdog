package bigdog

// API Version: v9_1

import (
	"context"
	"io"
	"net/http"
)

type WSGClusterManagementService struct {
	apiClient *VSZClient
}

func NewWSGClusterManagementService(c *VSZClient) *WSGClusterManagementService {
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
//
// Operation ID: addApPatch
// Operation path: /apPatch
// Success code: 200 (OK)
func (s *WSGClusterManagementService) AddApPatch(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationApPatchStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApPatch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApPatchStatusAPIResponse), err
}

// AddApPatchFile
//
// Use this API command to upload AP Patch File.
//
// Operation ID: addApPatchFile
// Operation path: /apPatch/file
// Success code: 204 (No Content)
//
// Form data parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGClusterManagementService) AddApPatchFile(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddApPatchFile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddClusterBackup
//
// Backup cluster.
//
// Operation ID: addClusterBackup
// Operation path: /cluster/backup
// Success code: 204 (No Content)
func (s *WSGClusterManagementService) AddClusterBackup(ctx context.Context, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddClusterBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddClusterRestoreById
//
// Restore cluster backup by ID.
//
// Operation ID: addClusterRestoreById
// Operation path: /cluster/restore/{id:.+}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddClusterRestoreById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddClusterRestoreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddConfigurationBackup
//
// Backup system configuration.
//
// Operation ID: addConfigurationBackup
// Operation path: /configuration/backup
// Success code: 201 (Created)
func (s *WSGClusterManagementService) AddConfigurationBackup(ctx context.Context, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddConfigurationBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddConfigurationRestoreById
//
// Restore system configuration with specified backupUUID.
//
// Operation ID: addConfigurationRestoreById
// Operation path: /configuration/restore/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddConfigurationRestoreById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddConfigurationRestoreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddConfigurationUpload
//
// Upload system configuration file.
//
// Operation ID: addConfigurationUpload
// Operation path: /configuration/upload
// Success code: 204 (No Content)
//
// Form data parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGClusterManagementService) AddConfigurationUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddConfigurationUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddUpgrade
//
// Use this API command to do system upgrade.
//
// Operation ID: addUpgrade
// Operation path: /upgrade
// Success code: 200 (OK)
func (s *WSGClusterManagementService) AddUpgrade(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationUpgradeStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradeStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddUpgrade, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
}

// AddUpgradeUpload
//
// Use this API command to upload patch file.
//
// Operation ID: addUpgradeUpload
// Operation path: /upgrade/upload
// Success code: 200 (OK)
//
// Form data parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGClusterManagementService) AddUpgradeUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*WSGAdministrationUpgradeStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradeStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddUpgradeUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
}

// DeleteClusterById
//
// Delete cluster backup by ID.
//
// Operation ID: deleteClusterById
// Operation path: /cluster/{id:.+}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteClusterById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteClusterById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteConfigurationById
//
// Delete system configuration file.
//
// Operation ID: deleteConfigurationById
// Operation path: /configuration/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteConfigurationById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteConfigurationById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindApPatch
//
// Use this API command to retrive uploaded AP patch file info.
//
// Operation ID: findApPatch
// Operation path: /apPatch
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindApPatch(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationApPatchInfoAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchInfoAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApPatch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApPatchInfoAPIResponse), err
}

// FindApPatchHistory
//
// Use this API command to retrive AP patch history.
//
// Operation ID: findApPatchHistory
// Operation path: /apPatch/history
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
// - timezone string
//		- nullable
func (s *WSGClusterManagementService) FindApPatchHistory(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationApPatchHistoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchHistoryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApPatchHistory, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["timezone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timezone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApPatchHistoryListAPIResponse), err
}

// FindApPatchStatus
//
// Use this API command to retrive cluster progress status.
//
// Operation ID: findApPatchStatus
// Operation path: /apPatch/status
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindApPatchStatus(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationApPatchStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApPatchStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApPatchStatusAPIResponse), err
}

// FindCluster
//
// Retrive cluster backup list.
//
// Operation ID: findCluster
// Operation path: /cluster
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
// - timezone string
//		- nullable
func (s *WSGClusterManagementService) FindCluster(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationClusterBackupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationClusterBackupListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindCluster, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["timezone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timezone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationClusterBackupListAPIResponse), err
}

// FindClusterGeoRedundancy
//
// Get cluster redundancy settings.
//
// Operation ID: findClusterGeoRedundancy
// Operation path: /cluster/geoRedundancy
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindClusterGeoRedundancy(ctx context.Context, mutators ...RequestMutator) (*WSGClusterRedundancySettingsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGClusterRedundancySettingsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindClusterGeoRedundancy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGClusterRedundancySettingsAPIResponse), err
}

// FindClusterState
//
// Use this API command to get current cluster, blade, and management service states
//
// Operation ID: findClusterState
// Operation path: /cluster/state
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindClusterState(ctx context.Context, mutators ...RequestMutator) (*WSGClusterBladeClusterStateAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGClusterBladeClusterStateAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindClusterState, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGClusterBladeClusterStateAPIResponse), err
}

// FindConfiguration
//
// Retrive system configuration list.
//
// Operation ID: findConfiguration
// Operation path: /configuration
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGClusterManagementService) FindConfiguration(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationConfigurationBackupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationConfigurationBackupListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindConfiguration, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationConfigurationBackupListAPIResponse), err
}

// FindConfigurationDownload
//
// Download system configuration file.
//
// Operation ID: findConfigurationDownload
// Operation path: /configuration/download
// Success code: 200 (OK)
//
// Required parameters:
// - backupUUID string
//		- required
//
// Optional parameters:
// - timeZone string
//		- nullable
func (s *WSGClusterManagementService) FindConfigurationDownload(ctx context.Context, backupUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindConfigurationDownload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("backupUUID", backupUUID)
	if v, ok := optionalParams["timeZone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timeZone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}

// FindConfigurationSettingsAutoExportBackup
//
// Get Auto Export Backup Settings.
//
// Operation ID: findConfigurationSettingsAutoExportBackup
// Operation path: /configurationSettings/autoExportBackup
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindConfigurationSettingsAutoExportBackup(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationAutoExportBackupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationAutoExportBackupAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindConfigurationSettingsAutoExportBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationAutoExportBackupAPIResponse), err
}

// FindConfigurationSettingsScheduleBackup
//
// Get Schedule Backup Setting.
//
// Operation ID: findConfigurationSettingsScheduleBackup
// Operation path: /configurationSettings/scheduleBackup
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindConfigurationSettingsScheduleBackup(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationScheduleBackupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationScheduleBackupAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindConfigurationSettingsScheduleBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationScheduleBackupAPIResponse), err
}

// FindUpgradeHistory
//
// Use this API command to retrive upgrade history.
//
// Operation ID: findUpgradeHistory
// Operation path: /upgrade/history
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
// - timezone string
//		- nullable
func (s *WSGClusterManagementService) FindUpgradeHistory(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationUpgradeHistoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradeHistoryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUpgradeHistory, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["timezone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timezone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradeHistoryListAPIResponse), err
}

// FindUpgradePatch
//
// Use this API command to retrive upload file Info.
//
// Operation ID: findUpgradePatch
// Operation path: /upgrade/patch
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindUpgradePatch(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationUpgradePatchInfoAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradePatchInfoAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUpgradePatch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradePatchInfoAPIResponse), err
}

// FindUpgradeStatus
//
// Use this API command to retrive cluster progress status.
//
// Operation ID: findUpgradeStatus
// Operation path: /upgrade/status
// Success code: 200 (OK)
func (s *WSGClusterManagementService) FindUpgradeStatus(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationUpgradeStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradeStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindUpgradeStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
}

// PartialUpdateConfigurationSettingsAutoExportBackup
//
// Modify Auto Export Backup Settings.
//
// Operation ID: partialUpdateConfigurationSettingsAutoExportBackup
// Operation path: /configurationSettings/autoExportBackup
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAdministrationModifyAutoExportBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsAutoExportBackup(ctx context.Context, body *WSGAdministrationModifyAutoExportBackup, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsAutoExportBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateConfigurationSettingsScheduleBackup
//
// Modify Schedule Backup Setting.
//
// Operation ID: partialUpdateConfigurationSettingsScheduleBackup
// Operation path: /configurationSettings/scheduleBackup
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAdministrationModifyScheduleBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsScheduleBackup(ctx context.Context, body *WSGAdministrationModifyScheduleBackup, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsScheduleBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateClusterGeoRedundancy
//
// Update cluster redundancy settings.
//
// Operation ID: updateClusterGeoRedundancy
// Operation path: /cluster/geoRedundancy
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGClusterRedundancyUpdateClusterRedundancy
func (s *WSGClusterManagementService) UpdateClusterGeoRedundancy(ctx context.Context, body *WSGClusterRedundancyUpdateClusterRedundancy, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateClusterGeoRedundancy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
