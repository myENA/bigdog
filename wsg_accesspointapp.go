package bigdog

// API Version: v9_0

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
func (s *WSGAccessPointAppService) FindApsLineman(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAPLinemanSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPLinemanSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsLineman, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["showAlarm"]; ok && len(v) > 0 {
		req.SetQueryParameter("showAlarm", v)
	}
	if v, ok := optionalParams["zoneId"]; ok && len(v) > 0 {
		req.SetQueryParameter("zoneId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAPLinemanSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApsTotalCount
//
// Use this API command to retrieve the total AP count within a zone or a domain.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointAppService) FindApsTotalCount(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsTotalCount, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["zoneId"]; ok && len(v) > 0 {
		req.SetQueryParameter("zoneId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAppService) FindLinemanWorkflow(ctx context.Context, mutators ...RequestMutator) (*FileResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *FileResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLinemanWorkflow, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "application/octet-stream")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(FileResponse)
	rm, err = handleFileResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateLinemanWorkflow
//
// Use this API command to upload a workflow file used by the Ruckus Wireless AP mobile app.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGAccessPointAppService) UpdateLinemanWorkflow(ctx context.Context, uploadFile io.Reader, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateLinemanWorkflow, true)
	req.SetHeader(headerKeyContentType, headerValueMultipartFormData)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = AddRequestMultipartValues(req, map[string]interface{}{"uploadFile": uploadFile}); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
