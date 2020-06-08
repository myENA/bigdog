package ruckus

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIZdxmlService struct {
	apiClient *SCIClient
}

func NewSCIZdxmlService(c *SCIClient) *SCIZdxmlService {
	s := new(SCIZdxmlService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIZdxmlService() *SCIZdxmlService {
	return NewSCIZdxmlService(ss.apiClient)
}

type SCIZdxmlGetAjaxRequest200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIZdxmlGetAjaxRequest200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIZdxmlGetAjaxRequest200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIZdxmlGetAjaxRequest200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIZdxmlGetAjaxRequest200ResponseType() *SCIZdxmlGetAjaxRequest200ResponseType {
	m := new(SCIZdxmlGetAjaxRequest200ResponseType)
	return m
}

type SCIZdxmlUpload200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIZdxmlUpload200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIZdxmlUpload200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIZdxmlUpload200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIZdxmlUpload200ResponseType() *SCIZdxmlUpload200ResponseType {
	m := new(SCIZdxmlUpload200ResponseType)
	return m
}

// ZdXmlGetAjaxRequest
//
// Required Parameters:
// - systemid string
//		- required
func (s *SCIZdxmlService) ZdXmlGetAjaxRequest(ctx context.Context, systemid string) (*SCIZdxmlGetAjaxRequest200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIZdxmlGetAjaxRequest200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIZdXmlGetAjaxRequest, false)
	req.SetQueryParameter("systemid", []string{systemid})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIZdxmlGetAjaxRequest200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ZdXmlUpload
//
// Required Parameters:
// - container string
//		- required
func (s *SCIZdxmlService) ZdXmlUpload(ctx context.Context, container string) (*SCIZdxmlUpload200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIZdxmlUpload200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIZdXmlUpload, false)
	req.SetPathParameter("container", container)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCIZdxmlUpload200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
