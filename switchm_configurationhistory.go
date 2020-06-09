package bigdog

// API Version: v9_0

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
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistory(ctx context.Context) (*SwitchMDeployLogConfigurationHistoryQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeployLogConfigurationHistoryQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindConfigurationHistory, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDeployLogConfigurationHistoryQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindConfigurationHistoryByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeployLogConfigurationHistoryQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeployLogConfigurationHistoryQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindConfigurationHistoryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDeployLogConfigurationHistoryQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindConfigurationHistoryDetail
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetail(ctx context.Context) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindConfigurationHistoryDetail, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDeployLogItemConfigurationHistoryDetailQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindConfigurationHistoryDetailByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindConfigurationHistoryDetailByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDeployLogItemConfigurationHistoryDetailQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
