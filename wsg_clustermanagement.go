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
// Operation ID: addApPatch
//
// Use this API command to apply AP patch.
func (s *WSGClusterManagementService) AddApPatch(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationApPatchStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchStatusAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationApPatchStatusAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApPatch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApPatchStatusAPIResponse), err
}

// AddApPatchFile
//
// Operation ID: addApPatchFile
//
// Use this API command to upload AP Patch File.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGClusterManagementService) AddApPatchFile(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddApPatchFile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddClusterBackup
//
// Operation ID: addClusterBackup
//
// Backup cluster.
func (s *WSGClusterManagementService) AddClusterBackup(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClusterBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddClusterRestoreById
//
// Operation ID: addClusterRestoreById
//
// Restore cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddClusterRestoreById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddClusterRestoreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddConfigurationBackup
//
// Operation ID: addConfigurationBackup
//
// Backup system configuration.
func (s *WSGClusterManagementService) AddConfigurationBackup(ctx context.Context, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddConfigurationBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddConfigurationRestoreById
//
// Operation ID: addConfigurationRestoreById
//
// Restore system configuration with specified backupUUID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) AddConfigurationRestoreById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddConfigurationRestoreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddConfigurationUpload
//
// Operation ID: addConfigurationUpload
//
// Upload system configuration file.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGClusterManagementService) AddConfigurationUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddConfigurationUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddUpgrade
//
// Operation ID: addUpgrade
//
// Use this API command to do system upgrade.
func (s *WSGClusterManagementService) AddUpgrade(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationUpgradeStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradeStatusAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddUpgrade, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
}

// AddUpgradeUpload
//
// Operation ID: addUpgradeUpload
//
// Use this API command to upload patch file.
//
// Form Data Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddUpgradeUpload, true)
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
// Operation ID: deleteClusterById
//
// Delete cluster backup by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteClusterById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteClusterById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteConfigurationById
//
// Operation ID: deleteConfigurationById
//
// Delete system configuration file.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGClusterManagementService) DeleteConfigurationById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteConfigurationById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindApPatch
//
// Operation ID: findApPatch
//
// Use this API command to retrive uploaded AP patch file info.
func (s *WSGClusterManagementService) FindApPatch(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationApPatchInfoAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchInfoAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationApPatchInfoAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApPatch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApPatchInfoAPIResponse), err
}

// FindApPatchHistory
//
// Operation ID: findApPatchHistory
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
func (s *WSGClusterManagementService) FindApPatchHistory(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationApPatchHistoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchHistoryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationApPatchHistoryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApPatchHistory, true)
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
// Operation ID: findApPatchStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindApPatchStatus(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationApPatchStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApPatchStatusAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationApPatchStatusAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApPatchStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApPatchStatusAPIResponse), err
}

// FindCluster
//
// Operation ID: findCluster
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
func (s *WSGClusterManagementService) FindCluster(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationClusterBackupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationClusterBackupListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationClusterBackupListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCluster, true)
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
// Operation ID: findClusterGeoRedundancy
//
// Get cluster redundancy settings.
func (s *WSGClusterManagementService) FindClusterGeoRedundancy(ctx context.Context, mutators ...RequestMutator) (*WSGClusterRedundancySettingsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGClusterRedundancySettingsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGClusterRedundancySettingsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindClusterGeoRedundancy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGClusterRedundancySettingsAPIResponse), err
}

// FindClusterState
//
// Operation ID: findClusterState
//
// Use this API command to get current cluster, blade, and management service states
func (s *WSGClusterManagementService) FindClusterState(ctx context.Context, mutators ...RequestMutator) (*WSGClusterBladeClusterStateAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGClusterBladeClusterStateAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGClusterBladeClusterStateAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindClusterState, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGClusterBladeClusterStateAPIResponse), err
}

// FindConfiguration
//
// Operation ID: findConfiguration
//
// Retrive system configuration list.
//
// Optional Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationConfigurationBackupListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindConfiguration, true)
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
// Operation ID: findConfigurationDownload
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
func (s *WSGClusterManagementService) FindConfigurationDownload(ctx context.Context, backupUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindConfigurationDownload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("backupUUID", backupUUID)
	if v, ok := optionalParams["timeZone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timeZone", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindConfigurationSettingsAutoExportBackup
//
// Operation ID: findConfigurationSettingsAutoExportBackup
//
// Get Auto Export Backup Settings.
func (s *WSGClusterManagementService) FindConfigurationSettingsAutoExportBackup(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationAutoExportBackupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationAutoExportBackupAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationAutoExportBackupAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindConfigurationSettingsAutoExportBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationAutoExportBackupAPIResponse), err
}

// FindConfigurationSettingsScheduleBackup
//
// Operation ID: findConfigurationSettingsScheduleBackup
//
// Get Schedule Backup Setting.
func (s *WSGClusterManagementService) FindConfigurationSettingsScheduleBackup(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationScheduleBackupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationScheduleBackupAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationScheduleBackupAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindConfigurationSettingsScheduleBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationScheduleBackupAPIResponse), err
}

// FindUpgradeHistory
//
// Operation ID: findUpgradeHistory
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
func (s *WSGClusterManagementService) FindUpgradeHistory(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationUpgradeHistoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradeHistoryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationUpgradeHistoryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindUpgradeHistory, true)
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
// Operation ID: findUpgradePatch
//
// Use this API command to retrive upload file Info.
func (s *WSGClusterManagementService) FindUpgradePatch(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationUpgradePatchInfoAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradePatchInfoAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationUpgradePatchInfoAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindUpgradePatch, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradePatchInfoAPIResponse), err
}

// FindUpgradeStatus
//
// Operation ID: findUpgradeStatus
//
// Use this API command to retrive cluster progress status.
func (s *WSGClusterManagementService) FindUpgradeStatus(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationUpgradeStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationUpgradeStatusAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindUpgradeStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationUpgradeStatusAPIResponse), err
}

// PartialUpdateConfigurationSettingsAutoExportBackup
//
// Operation ID: partialUpdateConfigurationSettingsAutoExportBackup
//
// Modify Auto Export Backup Settings.
//
// Request Body:
//	 - body *WSGAdministrationModifyAutoExportBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsAutoExportBackup(ctx context.Context, body *WSGAdministrationModifyAutoExportBackup, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsAutoExportBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateConfigurationSettingsScheduleBackup
//
// Operation ID: partialUpdateConfigurationSettingsScheduleBackup
//
// Modify Schedule Backup Setting.
//
// Request Body:
//	 - body *WSGAdministrationModifyScheduleBackup
func (s *WSGClusterManagementService) PartialUpdateConfigurationSettingsScheduleBackup(ctx context.Context, body *WSGAdministrationModifyScheduleBackup, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateConfigurationSettingsScheduleBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateClusterGeoRedundancy
//
// Operation ID: updateClusterGeoRedundancy
//
// Update cluster redundancy settings.
//
// Request Body:
//	 - body *WSGClusterRedundancyUpdateClusterRedundancy
func (s *WSGClusterManagementService) UpdateClusterGeoRedundancy(ctx context.Context, body *WSGClusterRedundancyUpdateClusterRedundancy, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateClusterGeoRedundancy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
