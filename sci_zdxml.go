package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIZoneDirectorXMLService struct {
	apiClient *SCIClient
}

func NewSCIZoneDirectorXMLService(c *SCIClient) *SCIZoneDirectorXMLService {
	s := new(SCIZoneDirectorXMLService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIZoneDirectorXMLService() *SCIZoneDirectorXMLService {
	return NewSCIZoneDirectorXMLService(ss.apiClient)
}

type SCIZoneDirectorXMLGetAjaxRequest200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIZoneDirectorXMLGetAjaxRequest200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIZoneDirectorXMLGetAjaxRequest200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIZoneDirectorXMLGetAjaxRequest200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIZoneDirectorXMLGetAjaxRequest200ResponseType() *SCIZoneDirectorXMLGetAjaxRequest200ResponseType {
	m := new(SCIZoneDirectorXMLGetAjaxRequest200ResponseType)
	return m
}

type SCIZoneDirectorXMLUpload200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIZoneDirectorXMLUpload200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIZoneDirectorXMLUpload200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIZoneDirectorXMLUpload200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIZoneDirectorXMLUpload200ResponseType() *SCIZoneDirectorXMLUpload200ResponseType {
	m := new(SCIZoneDirectorXMLUpload200ResponseType)
	return m
}

// ZdXmlGetAjaxRequest
//
// Required Parameters:
// - systemid string
//		- required
func (s *SCIZoneDirectorXMLService) ZdXmlGetAjaxRequest(ctx context.Context, systemid string, mutators ...RequestMutator) (*SCIZoneDirectorXMLGetAjaxRequest200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIZoneDirectorXMLGetAjaxRequest200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCIZdXmlGetAjaxRequest, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetQueryParameter("systemid", []string{systemid})
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIZoneDirectorXMLGetAjaxRequest200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ZdXmlUpload
//
// Required Parameters:
// - container string
//		- required
func (s *SCIZoneDirectorXMLService) ZdXmlUpload(ctx context.Context, container string, mutators ...RequestMutator) (*SCIZoneDirectorXMLUpload200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIZoneDirectorXMLUpload200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIZdXmlUpload, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("container", container)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIZoneDirectorXMLUpload200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
