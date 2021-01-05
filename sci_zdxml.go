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

// ZdXmlGetAjaxRequest
//
// Operation ID: zdXml_getAjaxRequest
//
// Required Parameters:
// - systemid string
//		- required
func (s *SCIZoneDirectorXMLService) ZdXmlGetAjaxRequest(ctx context.Context, systemid string, mutators ...RequestMutator) (*RawAPIResponse, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSCIZdXmlGetAjaxRequest, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.QueryParams.Set("system_id", systemid)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ZdXmlUpload
//
// Operation ID: zdXml_upload
//
// Required Parameters:
// - container string
//		- required
func (s *SCIZoneDirectorXMLService) ZdXmlUpload(ctx context.Context, container string, mutators ...RequestMutator) (*RawAPIResponse, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIZdXmlUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("container", container)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
