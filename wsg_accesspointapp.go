package bigdog

// API Version: v9_1

import (
	"context"
	"io"
	"net/http"
	"time"
)

type WSGAccessPointAppService struct {
	apiClient *VSZClient
}

func NewWSGAccessPointAppService(c *VSZClient) *WSGAccessPointAppService {
	s := new(WSGAccessPointAppService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointAppService() *WSGAccessPointAppService {
	return NewWSGAccessPointAppService(ss.apiClient)
}

// FindApsLineman
//
// Use this API command to retrieve the summary information of an AP. This is used by the Ruckus Wireless AP mobile app.
//
// Operation ID: findApsLineman
// Operation path: /aps/lineman
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - showAlarm string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointAppService) FindApsLineman(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAPLinemanSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAPLinemanSummaryAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsLineman, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["showAlarm"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("showAlarm", v)
	}
	if v, ok := optionalParams["zoneId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("zoneId", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAPLinemanSummaryAPIResponse), err
}

// FindApsTotalCount
//
// Use this API command to retrieve the total AP count within a zone or a domain.
//
// Operation ID: findApsTotalCount
// Operation path: /aps/totalCount
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointAppService) FindApsTotalCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApsTotalCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["zoneId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("zoneId", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// FindLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
//
// Operation ID: findLinemanWorkflow
// Operation path: /lineman/workflow
// Success code: 200 (OK)
func (s *WSGAccessPointAppService) FindLinemanWorkflow(ctx context.Context, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindLinemanWorkflow, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}

// UpdateLinemanWorkflow
//
// Use this API command to upload a workflow file used by the Ruckus Wireless AP mobile app.
//
// Operation ID: updateLinemanWorkflow
// Operation path: /lineman/workflow
// Success code: 204 (No Content)
//
// Form data parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGAccessPointAppService) UpdateLinemanWorkflow(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateLinemanWorkflow, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
