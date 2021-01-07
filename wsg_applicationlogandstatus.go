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
func (s *WSGApplicationLogAndStatusService) FindApplicationsByBladeUUID(ctx context.Context, bladeUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationApplicationLogAndStatusListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationApplicationLogAndStatusListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationApplicationLogAndStatusListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApplicationsByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationApplicationLogAndStatusListAPIResponse), err
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
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, appName string, bladeUUID string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApplicationsDownloadByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("appName", appName)
	req.PathParams.Set("bladeUUID", bladeUUID)
	if v, ok := optionalParams["logFileName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("logFileName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, bladeUUID string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindApplicationsDownloadsnapByBladeUUID, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateApplications
//
// Operation ID: partialUpdateApplications
//
// Use this API command to modify log level of specified application.
//
// Request Body:
//	 - body *WSGAdministrationModifyLogLevel
func (s *WSGApplicationLogAndStatusService) PartialUpdateApplications(ctx context.Context, body *WSGAdministrationModifyLogLevel, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateApplications, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
