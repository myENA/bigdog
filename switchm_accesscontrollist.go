package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type SwitchMAccessControlListService struct {
	apiClient *VSZClient
}

func NewSwitchMAccessControlListService(c *VSZClient) *SwitchMAccessControlListService {
	s := new(SwitchMAccessControlListService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAccessControlListService() *SwitchMAccessControlListService {
	return NewSwitchMAccessControlListService(ss.apiClient)
}

// AddAccessControls
//
// Use this API command to Create the Access Control Config.
//
// Operation ID: addAccessControls
// Operation path: /accessControls
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMACLConfigCreateACLConfig
func (s *SwitchMAccessControlListService) AddAccessControls(ctx context.Context, body *SwitchMACLConfigCreateACLConfig, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddAccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteAccessControls
//
// Use this API command to Delete the Access Control Config by Id list.
//
// Operation ID: deleteAccessControls
// Operation path: /accessControls
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMAccessControlListService) DeleteAccessControls(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteAccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteAccessControlsById
//
// Use this API command to Delete the Access Control Config.
//
// Operation ID: deleteAccessControlsById
// Operation path: /accessControls/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) DeleteAccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteAccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindAccessControlsById
//
// Use this API command to Retrieve the Access Control Config.
//
// Operation ID: findAccessControlsById
// Operation path: /accessControls/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) FindAccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMACLConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMACLConfigAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindAccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMACLConfigAPIResponse), err
}

// FindAccessControlsByQueryCriteria
//
// Use this API command to Retrieve the Access Control Config list.
//
// Operation ID: findAccessControlsByQueryCriteria
// Operation path: /accessControls/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMACLConfigsQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMACLConfigsQueryResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMFindAccessControlsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMACLConfigsQueryResultAPIResponse), err
}

// UpdateAccessControlsById
//
// Use this API command to Update the Access Control Config.
//
// Operation ID: updateAccessControlsById
// Operation path: /accessControls/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMACLConfigUpdateACLConfig
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) UpdateAccessControlsById(ctx context.Context, body *SwitchMACLConfigUpdateACLConfig, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateAccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
