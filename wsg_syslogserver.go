package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGSyslogServerService struct {
	apiClient *VSZClient
}

func NewWSGSyslogServerService(c *VSZClient) *WSGSyslogServerService {
	s := new(WSGSyslogServerService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSyslogServerService() *WSGSyslogServerService {
	return NewWSGSyslogServerService(ss.apiClient)
}

// FindSystemSyslog
//
// Retrieve syslog server sertting.
//
// Operation ID: findSystemSyslog
// Operation path: /system/syslog
// Success code: 200 (OK)
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGSyslogServerService) FindSystemSyslog(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSyslogServerSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSyslogServerSettingAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSystemSyslog, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSyslogServerSettingAPIResponse), err
}

// PartialUpdateSystemSyslog
//
// Modify syslog server setting.
//
// Operation ID: partialUpdateSystemSyslog
// Operation path: /system/syslog
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSyslogModifySyslogSettings
func (s *WSGSyslogServerService) PartialUpdateSystemSyslog(ctx context.Context, body *WSGSyslogModifySyslogSettings, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateSystemSyslog, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateSystemSyslogPrimaryServer
//
// Modify Primary Server of syslog.
//
// Operation ID: partialUpdateSystemSyslogPrimaryServer
// Operation path: /system/syslog/primaryServer
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSyslogPrimaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPrimaryServer(ctx context.Context, body *WSGSyslogPrimaryServer, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateSystemSyslogPrimaryServer, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateSystemSyslogPriority
//
// Modify Priority of syslog.
//
// Operation ID: partialUpdateSystemSyslogPriority
// Operation path: /system/syslog/priority
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSyslogPriority
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPriority(ctx context.Context, body *WSGSyslogPriority, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateSystemSyslogPriority, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateSystemSyslogSecondaryServer
//
// Modify Secondary Server of syslog.
//
// Operation ID: partialUpdateSystemSyslogSecondaryServer
// Operation path: /system/syslog/secondaryServer
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSyslogSecondaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogSecondaryServer(ctx context.Context, body *WSGSyslogSecondaryServer, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateSystemSyslogSecondaryServer, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
