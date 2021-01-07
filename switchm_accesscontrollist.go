package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addAccessControls
//
// Use this API command to Create the Access Control Config.
//
// Request Body:
//	 - body *SwitchMACLConfigCreateACLConfig
func (s *SwitchMAccessControlListService) AddAccessControls(ctx context.Context, body *SwitchMACLConfigCreateACLConfig, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddAccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteAccessControls
//
// Operation ID: deleteAccessControls
//
// Use this API command to Delete the Access Control Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMAccessControlListService) DeleteAccessControls(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteAccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteAccessControlsById
//
// Operation ID: deleteAccessControlsById
//
// Use this API command to Delete the Access Control Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) DeleteAccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteAccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindAccessControlsById
//
// Operation ID: findAccessControlsById
//
// Use this API command to Retrieve the Access Control Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) FindAccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMACLConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMACLConfigAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMACLConfigAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindAccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMACLConfigAPIResponse), err
}

// FindAccessControlsByQueryCriteria
//
// Operation ID: findAccessControlsByQueryCriteria
//
// Use this API command to Retrieve the Access Control Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMACLConfigsQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMACLConfigsQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMACLConfigsQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMFindAccessControlsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMACLConfigsQueryResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMACLConfigsQueryResultAPIResponse), err
}

// UpdateAccessControlsById
//
// Operation ID: updateAccessControlsById
//
// Use this API command to Update the Access Control Config.
//
// Request Body:
//	 - body *SwitchMACLConfigUpdateACLConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMAccessControlListService) UpdateAccessControlsById(ctx context.Context, body *SwitchMACLConfigUpdateACLConfig, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateAccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
