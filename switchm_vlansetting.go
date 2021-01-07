package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type SwitchMVLANSettingsService struct {
	apiClient *VSZClient
}

func NewSwitchMVLANSettingsService(c *VSZClient) *SwitchMVLANSettingsService {
	s := new(SwitchMVLANSettingsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVLANSettingsService() *SwitchMVLANSettingsService {
	return NewSwitchMVLANSettingsService(ss.apiClient)
}

// AddVlans
//
// Operation ID: addVlans
//
// Use this API command to Create the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVLANConfigCreateVlanConfig
func (s *SwitchMVLANSettingsService) AddVlans(ctx context.Context, body *SwitchMVLANConfigCreateVlanConfig, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddVlans, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteVlans
//
// Operation ID: deleteVlans
//
// Use this API command to Delete the VLAN Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVLANSettingsService) DeleteVlans(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteVlans, true)
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

// DeleteVlansById
//
// Operation ID: deleteVlansById
//
// Use this API command to Delete the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingsService) DeleteVlansById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteVlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindVlans
//
// Operation ID: findVlans
//
// Use this API command to Retrieve the VLAN Config List.
func (s *SwitchMVLANSettingsService) FindVlans(ctx context.Context, mutators ...RequestMutator) (*SwitchMVLANConfigQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMVLANConfigQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMVLANConfigQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindVlans, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMVLANConfigQueryResultAPIResponse), err
}

// FindVlansById
//
// Operation ID: findVlansById
//
// Use this API command to Retrieve the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingsService) FindVlansById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMVLANConfigAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMVLANConfigAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMVLANConfigAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindVlansById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMVLANConfigAPIResponse), err
}

// FindVlansByQueryCriteria
//
// Operation ID: findVlansByQueryCriteria
//
// Use this API command to Retrieve the VLAN Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVLANSettingsService) FindVlansByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMVLANConfigQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMVLANConfigQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMVLANConfigQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMFindVlansByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMVLANConfigQueryResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMVLANConfigQueryResultAPIResponse), err
}

// UpdateVlansById
//
// Operation ID: updateVlansById
//
// Use this API command to Update the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVLANConfigUpdateVlanConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingsService) UpdateVlansById(ctx context.Context, body *SwitchMVLANConfigUpdateVlanConfig, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateVlansById, true)
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
