package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type SwitchMConfigurationService struct {
	apiClient *VSZClient
}

func NewSwitchMConfigurationService(c *VSZClient) *SwitchMConfigurationService {
	s := new(SwitchMConfigurationService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMConfigurationService() *SwitchMConfigurationService {
	return NewSwitchMConfigurationService(ss.apiClient)
}

// AddSwitchconfig
//
// Use this API command to retrieve configuration backup list with specified filters.
//
// Operation ID: addSwitchconfig
// Operation path: /switchconfig
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationService) AddSwitchconfig(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMConfigurationBackupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMConfigurationBackupListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchconfig, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMConfigurationBackupListAPIResponse), err
}

// AddSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Operation ID: addSwitchconfigBackup
// Operation path: /switchconfig/backup
// Success code: 201 (Created)
//
// Request body:
//	 - body SwitchMConfigurationBackupSwitchIds
func (s *SwitchMConfigurationService) AddSwitchconfigBackup(ctx context.Context, body SwitchMConfigurationBackupSwitchIds, mutators ...RequestMutator) (*SwitchMConfigurationBackupCreateBackupResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMConfigurationBackupCreateBackupResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchconfigBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMConfigurationBackupCreateBackupResultListAPIResponse), err
}

// AddSwitchconfigBackupDiff
//
// Use this API command to diff between two config back up files for a switch.
//
// Operation ID: addSwitchconfigBackupDiff
// Operation path: /switchconfig/backup/diff
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMConfigurationBackupConfigBackupDiffInput
func (s *SwitchMConfigurationService) AddSwitchconfigBackupDiff(ctx context.Context, body *SwitchMConfigurationBackupConfigBackupDiffInput, mutators ...RequestMutator) (*SwitchMConfigurationBackupConfigBackupDiffAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMConfigurationBackupConfigBackupDiffAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchconfigBackupDiff, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMConfigurationBackupConfigBackupDiffAPIResponse), err
}

// DeleteSwitchconfig
//
// Use this API command to delete config backups by a list of config backup id.
//
// Operation ID: deleteSwitchconfig
// Operation path: /switchconfig
// Success code: 204 (No Content)
//
// Request body:
//	 - body SwitchMConfigurationBackupBackupIds
func (s *SwitchMConfigurationService) DeleteSwitchconfig(ctx context.Context, body SwitchMConfigurationBackupBackupIds, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteSwitchconfig, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteSwitchconfigByConfigId
//
// Use this API command to delete the configuration backup.
//
// Operation ID: deleteSwitchconfigByConfigId
// Operation path: /switchconfig/{configId}
// Success code: 204 (No Content)
//
// Required parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) DeleteSwitchconfigByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteSwitchconfigByConfigId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("configId", configId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindSwitchconfigByConfigId
//
// Use this API command to retrieve configuration backup content.
//
// Operation ID: findSwitchconfigByConfigId
// Operation path: /switchconfig/{configId}
// Success code: 200 (OK)
//
// Required parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) FindSwitchconfigByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSwitchconfigByConfigId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("configId", configId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// FindSwitchconfigDownloadByConfigId
//
// Use this API command to download configuration backup content as plain text.
//
// Operation ID: findSwitchconfigDownloadByConfigId
// Operation path: /switchconfig/download/{configId}
// Success code: 200 (OK)
//
// Required parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) FindSwitchconfigDownloadByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSwitchconfigDownloadByConfigId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("configId", configId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// UpdateSwitchconfigBackupByGroupId
//
// Use this API command to backup configurations for all switches under a group.
//
// Operation ID: updateSwitchconfigBackupByGroupId
// Operation path: /switchconfig/{groupType}/{groupId}/backup
// Success code: 204 (No Content)
//
// Required parameters:
// - groupId string
//		- required
// - groupType string
//		- required
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupByGroupId(ctx context.Context, groupId string, groupType string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("groupType", groupType)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateSwitchconfigBackupRestoreByBackupId
//
// Restore a configuration backup to the switch.
//
// Operation ID: updateSwitchconfigBackupRestoreByBackupId
// Operation path: /switchconfig/backup/{backupId}/restore
// Success code: 200 (OK)
//
// Required parameters:
// - backupId string
//		- required
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId(ctx context.Context, backupId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupRestoreByBackupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("backupId", backupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}
