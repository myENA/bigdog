package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type SwitchMEventService struct {
	apiClient *APIClient
}

func NewSwitchMEventService(c *APIClient) *SwitchMEventService {
	s := new(SwitchMEventService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMEventService() *SwitchMEventService {
	return NewSwitchMEventService(ss.apiClient)
}

// AddCustomEvent
//
// Use this API command to create a new text pattern event config
//
// Request Body:
//	 - body *SwitchMEventConfig
func (s *SwitchMEventService) AddCustomEvent(ctx context.Context, body *SwitchMEventConfig) (*SwitchMEventConfigQueryResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMEventConfigQueryResponse
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMAddCustomEvent), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMEventConfigQueryResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) DeleteCustomEventById(ctx context.Context, id string) (*SwitchMEventConfigQueryResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMEventConfigQueryResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMDeleteCustomEventById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMEventConfigQueryResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindCustomEvent
//
// Use this API command to retrieve switch event config list
func (s *SwitchMEventService) FindCustomEvent(ctx context.Context) (*SwitchMEventConfigGetEventConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMEventConfigGetEventConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMFindCustomEvent), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMEventConfigGetEventConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindCustomEventById
//
// Use this API command to retrieve one switch event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) FindCustomEventById(ctx context.Context, id string) (*SwitchMEventConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMEventConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMFindCustomEventById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMEventConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateCustomEventById
//
// Use this API command to modify a switch custom event config. The patch variable {id} is same as id attribute in the request payload. For CPU/Memory, only key, type, criteria, and severity attributes are required.
//
// Request Body:
//	 - body *SwitchMEventConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) UpdateCustomEventById(ctx context.Context, body *SwitchMEventConfig, id string) (*SwitchMEventConfigQueryResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMEventConfigQueryResponse
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMUpdateCustomEventById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMEventConfigQueryResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
