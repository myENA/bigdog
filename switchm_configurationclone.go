package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type SwitchMConfigurationCloneService struct {
	apiClient *VSZClient
}

func NewSwitchMConfigurationCloneService(c *VSZClient) *SwitchMConfigurationCloneService {
	s := new(SwitchMConfigurationCloneService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMConfigurationCloneService() *SwitchMConfigurationCloneService {
	return NewSwitchMConfigurationCloneService(ss.apiClient)
}

// AddCloneConfiguration
//
// Operation ID: addCloneConfiguration
//
// Use this API command to Get Switch Config.
//
// Request Body:
//	 - body *SwitchMGroupGetConfigBySwitch
func (s *SwitchMConfigurationCloneService) AddCloneConfiguration(ctx context.Context, body *SwitchMGroupGetConfigBySwitch, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddCloneConfiguration, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddCloneConfigurationByGroup
//
// Operation ID: addCloneConfigurationByGroup
//
// Use this API command to Clone Switch Group Config.
//
// Request Body:
//	 - body *SwitchMGroupCloneConfigByGroup
func (s *SwitchMConfigurationCloneService) AddCloneConfigurationByGroup(ctx context.Context, body *SwitchMGroupCloneConfigByGroup, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddCloneConfigurationByGroup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// UpdateCloneConfiguration
//
// Operation ID: updateCloneConfiguration
//
// Use this API command to Clone Switch Config.
//
// Request Body:
//	 - body *SwitchMGroupCloneConfigBySwitch
func (s *SwitchMConfigurationCloneService) UpdateCloneConfiguration(ctx context.Context, body *SwitchMGroupCloneConfigBySwitch, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateCloneConfiguration, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
