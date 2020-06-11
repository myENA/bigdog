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

type SCISettingSendTestEmail200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISettingSendTestEmail200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCISettingSendTestEmail200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCISettingSendTestEmail200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCISettingSendTestEmail200ResponseType() *SCISettingSendTestEmail200ResponseType {
	m := new(SCISettingSendTestEmail200ResponseType)
	return m
}

// SettingFindById
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
func (s *SCISettingService) SettingFindById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*SCIModelsSetting, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSCISettingFindById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["filter"]; ok && len(v) > 0 {
		req.SetQueryParameter("filter", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SettingSendTestEmail
//
// Form Data Parameters:
// - recipients string
//		- required
func (s *SCISettingService) SettingSendTestEmail(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCISettingSendTestEmail200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCISettingSendTestEmail200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCISettingSendTestEmail, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCISettingSendTestEmail200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// SettingUpsert
//
// Update an existing model instance or insert a new one into the data source.
//
// Request Body:
//	 - body *SCIModelsSetting
func (s *SCISettingService) SettingUpsert(ctx context.Context, data *SCIModelsSetting, mutators ...RequestMutator) (*SCIModelsSetting, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSCISettingUpsert, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(data); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
