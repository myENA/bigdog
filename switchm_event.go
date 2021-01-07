package bigdog

// API Version: v9_1

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
// Operation ID: addCustomEvent
//
// Use this API command to create a new text pattern event config
//
// Request Body:
//	 - body *SwitchMEventConfig
func (s *SwitchMEventService) AddCustomEvent(ctx context.Context, body *SwitchMEventConfig, mutators ...RequestMutator) (*SwitchMEventConfigQueryResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMEventConfigQueryResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddCustomEvent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
}

// DeleteCustomEventById
//
// Operation ID: deleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) DeleteCustomEventById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMEventConfigQueryResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMEventConfigQueryResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteCustomEventById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
}

// FindCustomEvent
//
// Operation ID: findCustomEvent
//
// Use this API command to retrieve switch event config list
func (s *SwitchMEventService) FindCustomEvent(ctx context.Context, mutators ...RequestMutator) (*SwitchMEventConfigGetEventConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMEventConfigGetEventConfigListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMEventConfigGetEventConfigListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindCustomEvent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigGetEventConfigListAPIResponse), err
}

// FindCustomEventById
//
// Operation ID: findCustomEventById
//
// Use this API command to retrieve one switch event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) FindCustomEventById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMEventConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMEventConfigAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMEventConfigAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindCustomEventById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigAPIResponse), err
}

// UpdateCustomEventById
//
// Operation ID: updateCustomEventById
//
// Use this API command to modify a switch custom event config. The patch variable {id} is same as id attribute in the request payload. For CPU/Memory, only key, type, criteria, and severity attributes are required.
//
// Request Body:
//	 - body *SwitchMEventConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) UpdateCustomEventById(ctx context.Context, body *SwitchMEventConfig, id string, mutators ...RequestMutator) (*SwitchMEventConfigQueryResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMEventConfigQueryResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateCustomEventById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
}
