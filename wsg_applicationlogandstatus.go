package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGApplicationLogAndStatusService struct {
	apiClient *VSZClient
}

func NewWSGApplicationLogAndStatusService(c *VSZClient) *WSGApplicationLogAndStatusService {
	s := new(WSGApplicationLogAndStatusService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGApplicationLogAndStatusService() *WSGApplicationLogAndStatusService {
	return NewWSGApplicationLogAndStatusService(ss.apiClient)
}

// FindApplicationsByBladeUUID
//
// Operation ID: findApplicationsByBladeUUID
//
// Use this API command to retrieve a list of application log and status.
//
// Required Parameters:
// - bladeUUID string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGApplicationLogAndStatusService) FindApplicationsByBladeUUID(ctx context.Context, bladeUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationApplicationLogAndStatusList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationApplicationLogAndStatusList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsByBladeUUID, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("bladeUUID", bladeUUID)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAdministrationApplicationLogAndStatusList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApplicationsDownloadByBladeUUID
//
// Operation ID: findApplicationsDownloadByBladeUUID
//
// Use this API command to download logs of the application.
//
// Required Parameters:
// - appName string
//		- required
// - bladeUUID string
//		- required
//
// Optional Parameters:
// - logFileName string
//		- nullable
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, appName string, bladeUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsDownloadByBladeUUID, true)
	req.SetHeader(headerKeyAccept, "application/octet-stream")
	req.SetQueryParameter("appName", []string{appName})
	req.SetPathParameter("bladeUUID", bladeUUID)
	if v, ok := optionalParams["logFileName"]; ok && len(v) > 0 {
		req.SetQueryParameter("logFileName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApplicationsDownloadsnapByBladeUUID
//
// Operation ID: findApplicationsDownloadsnapByBladeUUID
//
// Use this API command to download snapshot logs.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsDownloadsnapByBladeUUID, true)
	req.SetHeader(headerKeyAccept, "application/octet-stream")
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateApplications
//
// Operation ID: partialUpdateApplications
//
// Use this API command to modify log level of specified application.
//
// Request Body:
//	 - body *WSGAdministrationModifyLogLevel
func (s *WSGApplicationLogAndStatusService) PartialUpdateApplications(ctx context.Context, body *WSGAdministrationModifyLogLevel, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateApplications, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
