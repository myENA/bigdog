package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
)

type SCISettingService struct {
	apiClient *SCIClient
}

func NewSCISettingService(c *SCIClient) *SCISettingService {
	s := new(SCISettingService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCISettingService() *SCISettingService {
	return NewSCISettingService(ss.apiClient)
}

// SettingDeleteById
//
// Delete a model instance by id from the data source.
//
// Operation ID: setting_deleteById
// Operation path: /settings/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SCISettingService) SettingDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSCISettingDeleteById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// SettingFindById
//
// Find a model instance by id from the data source.
//
// Operation ID: setting_findById
// Operation path: /settings/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - filter string
//		- nullable
func (s *SCISettingService) SettingFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsSettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsSettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCISettingFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSettingAPIResponse), err
}

// SettingSendTestEmail
//
// Operation ID: setting_sendTestEmail
// Operation path: /settings/sendTestEmail
// Success code: 200 (OK)
//
// Form data parameters:
// - recipients string
//		- required
func (s *SCISettingService) SettingSendTestEmail(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCISettingSendTestEmail, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// SettingUpsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Operation ID: setting_upsert
// Operation path: /settings
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCIModelsSetting
func (s *SCISettingService) SettingUpsert(ctx context.Context, data *SCIModelsSetting, mutators ...RequestMutator) (*SCIModelsSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsSettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsSettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCISettingUpsert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(data); err != nil {
		return resp.(*SCIModelsSettingAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsSettingAPIResponse), err
}
