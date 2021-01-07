package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
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
// Operation ID: setting_deleteById
//
// Delete a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
func (s *SCISettingService) SettingDeleteById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSCISettingDeleteById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SettingFindById
//
// Operation ID: setting_findById
//
// Find a model instance by id from the data source.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - filter string
//		- nullable
func (s *SCISettingService) SettingFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsSettingAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSCISettingFindById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSetting()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SettingSendTestEmail
//
// Operation ID: setting_sendTestEmail
//
// Form Data Parameters:
// - recipients string
//		- required
func (s *SCISettingService) SettingSendTestEmail(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCISettingSendTestEmail, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SettingUpsert
//
// Operation ID: setting_upsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsSetting
func (s *SCISettingService) SettingUpsert(ctx context.Context, data *SCIModelsSetting, mutators ...RequestMutator) (*SCIModelsSettingAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCISettingUpsert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSetting()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
