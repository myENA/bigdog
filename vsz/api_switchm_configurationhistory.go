package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMConfigurationhistoryService struct {
	apiClient *APIClient
}

func NewSwitchMConfigurationhistoryService(c *APIClient) *SwitchMConfigurationhistoryService {
	s := new(SwitchMConfigurationhistoryService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMConfigurationhistoryService() *SwitchMConfigurationhistoryService {
	return NewSwitchMConfigurationhistoryService(ss.apiClient)
}

// FindConfigurationHistory
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationhistoryService) FindConfigurationHistory(ctx context.Context) (*SwitchMDeploylogConfigurationHistoryQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeploylogConfigurationHistoryQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindConfigurationHistory, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDeploylogConfigurationHistoryQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindConfigurationHistoryByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationhistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeploylogConfigurationHistoryQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeploylogConfigurationHistoryQueryResult
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
	resp = NewSwitchMDeploylogConfigurationHistoryQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindConfigurationHistoryDetail
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationhistoryService) FindConfigurationHistoryDetail(ctx context.Context) (*SwitchMDeploylogitemConfigurationHistoryDetailQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeploylogitemConfigurationHistoryDetailQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindConfigurationHistoryDetail, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMDeploylogitemConfigurationHistoryDetailQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindConfigurationHistoryDetailByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationhistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeploylogitemConfigurationHistoryDetailQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMDeploylogitemConfigurationHistoryDetailQueryResult
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
	resp = NewSwitchMDeploylogitemConfigurationHistoryDetailQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
