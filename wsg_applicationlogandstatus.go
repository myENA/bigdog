package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Use this API command to retrieve a list of application log and status.
//
// Operation ID: findApplicationsByBladeUUID
// Operation path: /applications/{bladeUUID}
// Success code: 200 (OK)
//
// Required parameters:
// - bladeUUID string
//		- required
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGApplicationLogAndStatusService) FindApplicationsByBladeUUID(ctx context.Context, bladeUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationApplicationLogAndStatusListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApplicationLogAndStatusListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApplicationsByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGAdministrationApplicationLogAndStatusListAPIResponse), err
}

// FindApplicationsDownloadByBladeUUID
//
// Use this API command to download logs of the application.
//
// Operation ID: findApplicationsDownloadByBladeUUID
// Operation path: /applications/download/{bladeUUID}
// Success code: 200 (OK)
//
// Required parameters:
// - appName string
//		- required
// - bladeUUID string
//		- required
//
// Optional parameters:
// - logFileName string
//		- nullable
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, appName string, bladeUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApplicationsDownloadByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("appName", appName)
	req.PathParams.Set("bladeUUID", bladeUUID)
	if v, ok := optionalParams["logFileName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("logFileName", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}

// FindApplicationsDownloadsnapByBladeUUID
//
// Use this API command to download snapshot logs.
//
// Operation ID: findApplicationsDownloadsnapByBladeUUID
// Operation path: /applications/downloadsnap/{bladeUUID}
// Success code: 200 (OK)
//
// Required parameters:
// - bladeUUID string
//		- required
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindApplicationsDownloadsnapByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}

// PartialUpdateApplications
//
// Use this API command to modify log level of specified application.
//
// Operation ID: partialUpdateApplications
// Operation path: /applications
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAdministrationModifyLogLevel
func (s *WSGApplicationLogAndStatusService) PartialUpdateApplications(ctx context.Context, body *WSGAdministrationModifyLogLevel, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateApplications, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
