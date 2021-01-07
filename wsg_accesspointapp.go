package bigdog

// API Version: v9_1

import (
	"context"
	"io"
	"net/http"
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
// Operation ID: findApsLineman
//
// Use this API command to retrieve the summary information of an AP. This is used by the Ruckus Wireless AP mobile app.
//
// Optional Parameters:
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
		resp     APIResponse
		err      error

		respFn = newWSGAPLinemanSummaryAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPLinemanSummaryAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsLineman, true)
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
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPLinemanSummaryAPIResponse), err
}

// FindApsTotalCount
//
// Operation ID: findApsTotalCount
//
// Use this API command to retrieve the total AP count within a zone or a domain.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointAppService) FindApsTotalCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApsTotalCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["zoneId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("zoneId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindLinemanWorkflow
//
// Operation ID: findLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAppService) FindLinemanWorkflow(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindLinemanWorkflow, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateLinemanWorkflow
//
// Operation ID: updateLinemanWorkflow
//
// Use this API command to upload a workflow file used by the Ruckus Wireless AP mobile app.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGAccessPointAppService) UpdateLinemanWorkflow(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateLinemanWorkflow, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
