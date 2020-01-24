package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMConfigurationHistoryService struct {
	apiClient *APIClient
}

func NewSwitchMConfigurationHistoryService(c *APIClient) *SwitchMConfigurationHistoryService {
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
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistory(ctx context.Context) (*SwitchMDeployLogConfigurationHistoryQueryResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMDeployLogConfigurationHistoryQueryResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindConfigurationHistory, true)
}

// FindConfigurationHistoryByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeployLogConfigurationHistoryQueryResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMDeployLogConfigurationHistoryQueryResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindConfigurationHistoryByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// FindConfigurationHistoryDetail
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetail(ctx context.Context) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindConfigurationHistoryDetail, true)
}

// FindConfigurationHistoryDetailByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMConfigurationHistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMDeployLogItemConfigurationHistoryDetailQueryResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindConfigurationHistoryDetailByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
