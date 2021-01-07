package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type SwitchMConfigurationHistoryService struct {
	apiClient *VSZClient
}

func NewSwitchMConfigurationHistoryService(c *VSZClient) *SwitchMConfigurationHistoryService {
	s := new(SwitchMConfigurationHistoryService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMConfigurationHistoryService() *SwitchMConfigurationHistoryService {
	return NewSwitchMConfigurationHistoryService(ss.apiClient)
}

// FindConfigurationHistory
//
// Operation ID: findConfigurationHistory
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistory(ctx context.Context, mutators ...RequestMutator) (*SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMDeployLogConfigurationHistoryQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindConfigurationHistory, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse), err
}

// FindConfigurationHistoryByQueryCriteria
//
// Operation ID: findConfigurationHistoryByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMDeployLogConfigurationHistoryQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMFindConfigurationHistoryByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse), err
}

// FindConfigurationHistoryDetail
//
// Operation ID: findConfigurationHistoryDetail
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetail(ctx context.Context, mutators ...RequestMutator) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindConfigurationHistoryDetail, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse), err
}

// FindConfigurationHistoryDetailByQueryCriteria
//
// Operation ID: findConfigurationHistoryDetailByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMFindConfigurationHistoryDetailByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse), err
}
