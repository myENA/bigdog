package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addSwitchconfig
//
// Use this API command to retrieve configuration backup list with specified filters.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationService) AddSwitchconfig(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMConfigurationBackupListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMConfigurationBackupListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMConfigurationBackupListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchconfig, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMConfigurationBackupListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMConfigurationBackupListAPIResponse), err
}

// AddSwitchconfigBackup
//
// Operation ID: addSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body SwitchMConfigurationBackupSwitchIds
func (s *SwitchMConfigurationService) AddSwitchconfigBackup(ctx context.Context, body SwitchMConfigurationBackupSwitchIds, mutators ...RequestMutator) (*SwitchMConfigurationBackupCreateBackupResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMConfigurationBackupCreateBackupResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMConfigurationBackupCreateBackupResultListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchconfigBackup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMConfigurationBackupCreateBackupResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMConfigurationBackupCreateBackupResultListAPIResponse), err
}

// AddSwitchconfigBackupDiff
//
// Operation ID: addSwitchconfigBackupDiff
//
// Use this API command to diff between two config back up files for a switch.
//
// Request Body:
//	 - body *SwitchMConfigurationBackupConfigBackupDiffInput
func (s *SwitchMConfigurationService) AddSwitchconfigBackupDiff(ctx context.Context, body *SwitchMConfigurationBackupConfigBackupDiffInput, mutators ...RequestMutator) (*SwitchMConfigurationBackupConfigBackupDiffAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMConfigurationBackupConfigBackupDiffAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMConfigurationBackupConfigBackupDiffAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchconfigBackupDiff, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMConfigurationBackupConfigBackupDiffAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMConfigurationBackupConfigBackupDiffAPIResponse), err
}

// DeleteSwitchconfig
//
// Operation ID: deleteSwitchconfig
//
// Use this API command to delete config backups by a list of config backup id.
//
// Request Body:
//	 - body SwitchMConfigurationBackupBackupIds
func (s *SwitchMConfigurationService) DeleteSwitchconfig(ctx context.Context, body SwitchMConfigurationBackupBackupIds, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteSwitchconfig, true)
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

// DeleteSwitchconfigByConfigId
//
// Operation ID: deleteSwitchconfigByConfigId
//
// Use this API command to delete the configuration backup.
//
// Required Parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) DeleteSwitchconfigByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteSwitchconfigByConfigId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("configId", configId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindSwitchconfigByConfigId
//
// Operation ID: findSwitchconfigByConfigId
//
// Use this API command to retrieve configuration backup content.
//
// Required Parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) FindSwitchconfigByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindSwitchconfigByConfigId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("configId", configId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindSwitchconfigDownloadByConfigId
//
// Operation ID: findSwitchconfigDownloadByConfigId
//
// Use this API command to download configuration backup content as plain text.
//
// Required Parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) FindSwitchconfigDownloadByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindSwitchconfigDownloadByConfigId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("configId", configId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateSwitchconfigBackupByGroupId
//
// Operation ID: updateSwitchconfigBackupByGroupId
//
// Use this API command to backup configurations for all switches under a group.
//
// Required Parameters:
// - groupId string
//		- required
// - groupType string
//		- required
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupByGroupId(ctx context.Context, groupId string, groupType string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("groupType", groupType)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateSwitchconfigBackupRestoreByBackupId
//
// Operation ID: updateSwitchconfigBackupRestoreByBackupId
//
// Restore a configuration backup to the switch.
//
// Required Parameters:
// - backupId string
//		- required
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId(ctx context.Context, backupId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupRestoreByBackupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("backupId", backupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
