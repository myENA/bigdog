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
// Use this API command to create a new text pattern event config
//
// Operation ID: addCustomEvent
// Operation path: /customEvent
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMEventConfig
func (s *SwitchMEventService) AddCustomEvent(ctx context.Context, body *SwitchMEventConfig, mutators ...RequestMutator) (*SwitchMEventConfigQueryResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMEventConfigQueryResponseAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddCustomEvent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
}

// DeleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Operation ID: deleteCustomEventById
// Operation path: /customEvent/{id}
// Success code: 200 (OK)
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteCustomEventById, true)
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
// Use this API command to retrieve switch event config list
//
// Operation ID: findCustomEvent
// Operation path: /customEvent
// Success code: 200 (OK)
func (s *SwitchMEventService) FindCustomEvent(ctx context.Context, mutators ...RequestMutator) (*SwitchMEventConfigGetEventConfigListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMEventConfigGetEventConfigListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindCustomEvent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigGetEventConfigListAPIResponse), err
}

// FindCustomEventById
//
// Use this API command to retrieve one switch event config
//
// Operation ID: findCustomEventById
// Operation path: /customEvent/{id}
// Success code: 200 (OK)
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindCustomEventById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigAPIResponse), err
}

// UpdateCustomEventById
//
// Use this API command to modify a switch custom event config. The patch variable {id} is same as id attribute in the request payload. For CPU/Memory, only key, type, criteria, and severity attributes are required.
//
// Operation ID: updateCustomEventById
// Operation path: /customEvent/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMEventConfig
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateCustomEventById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMEventConfigQueryResponseAPIResponse), err
}
