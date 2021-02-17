package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
	"time"
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
// Operation path: /zdXmls/getAjaxRequest
// Success code: 200 (OK)
//
// Required parameters:
// - systemid string
//		- required
func (s *SCIZoneDirectorXMLService) ZdXmlGetAjaxRequest(ctx context.Context, systemid string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodGet, RouteSCIZdXmlGetAjaxRequest, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("system_id", systemid)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// ZdXmlUpload
//
// Operation ID: zdXml_upload
// Operation path: /zdXmls/{container}/upload
// Success code: 200 (OK)
//
// Required parameters:
// - container string
//		- required
func (s *SCIZoneDirectorXMLService) ZdXmlUpload(ctx context.Context, container string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIZdXmlUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("container", container)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}
