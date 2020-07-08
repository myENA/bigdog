package bigdog

// API Version: v9_0

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
func (s *SwitchMConfigurationService) AddSwitchconfig(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMConfigurationBackupList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMConfigurationBackupList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchconfig, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMConfigurationBackupList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchconfigBackup
//
// Operation ID: addSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body SwitchMConfigurationBackupSwitchIds
func (s *SwitchMConfigurationService) AddSwitchconfigBackup(ctx context.Context, body SwitchMConfigurationBackupSwitchIds, mutators ...RequestMutator) (*SwitchMConfigurationBackupCreateBackupResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMConfigurationBackupCreateBackupResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchconfigBackup, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMConfigurationBackupCreateBackupResultList()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchconfigBackupDiff
//
// Operation ID: addSwitchconfigBackupDiff
//
// Use this API command to diff between two config back up files for a switch.
//
// Request Body:
//	 - body *SwitchMConfigurationBackupConfigBackupDiffInput
func (s *SwitchMConfigurationService) AddSwitchconfigBackupDiff(ctx context.Context, body *SwitchMConfigurationBackupConfigBackupDiffInput, mutators ...RequestMutator) (*SwitchMConfigurationBackupConfigBackupDiff, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMConfigurationBackupConfigBackupDiff
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchconfigBackupDiff, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMConfigurationBackupConfigBackupDiff()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteSwitchconfig
//
// Operation ID: deleteSwitchconfig
//
// Use this API command to delete config backups by a list of config backup id.
//
// Request Body:
//	 - body SwitchMConfigurationBackupBackupIds
func (s *SwitchMConfigurationService) DeleteSwitchconfig(ctx context.Context, body SwitchMConfigurationBackupBackupIds, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteSwitchconfig, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *SwitchMConfigurationService) DeleteSwitchconfigByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteSwitchconfigByConfigId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("configId", configId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *SwitchMConfigurationService) FindSwitchconfigByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSwitchconfigByConfigId, true)
	req.SetHeader(headerKeyAccept, "plain/text")
	req.SetPathParameter("configId", configId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SwitchMConfigurationService) FindSwitchconfigDownloadByConfigId(ctx context.Context, configId string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSwitchconfigDownloadByConfigId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("configId", configId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
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
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupByGroupId(ctx context.Context, groupId string, groupType string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupByGroupId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("groupId", groupId)
	req.SetPathParameter("groupType", groupType)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId(ctx context.Context, backupId string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupRestoreByBackupId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("backupId", backupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
