package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMEventService struct {
	apiClient *VSZClient
}

func NewSwitchMEventService(c *VSZClient) *SwitchMEventService {
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
//	 - body *SwitchMSwitchEventConfigEventConfig
func (s *SwitchMEventService) AddCustomEvent(ctx context.Context, body *SwitchMSwitchEventConfigEventConfig) (*SwitchMSwitchEventConfigQueryResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchEventConfigQueryResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddCustomEvent, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchEventConfigQueryResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) DeleteCustomEventById(ctx context.Context, id string) (*SwitchMSwitchEventConfigQueryResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchEventConfigQueryResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteCustomEventById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchEventConfigQueryResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindCustomEvent
//
// Use this API command to retrieve switch event config list
func (s *SwitchMEventService) FindCustomEvent(ctx context.Context) (*SwitchMSwitchEventConfigGetEventConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchEventConfigGetEventConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindCustomEvent, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchEventConfigGetEventConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindCustomEventById
//
// Use this API command to retrieve one switch event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) FindCustomEventById(ctx context.Context, id string) (*SwitchMSwitchEventConfigEventConfig, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchEventConfigEventConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindCustomEventById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchEventConfigEventConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateCustomEventById
//
// Use this API command to modify a switch custom event config. The patch variable {id} is same as id attribute in the request payload. For CPU/Memory, only key, type, criteria, and severity attributes are required.
//
// Request Body:
//	 - body *SwitchMSwitchEventConfigEventConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) UpdateCustomEventById(ctx context.Context, body *SwitchMSwitchEventConfigEventConfig, id string) (*SwitchMSwitchEventConfigQueryResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchEventConfigQueryResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateCustomEventById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchEventConfigQueryResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
