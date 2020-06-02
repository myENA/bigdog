package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGApplicationlogandstatusService struct {
	apiClient *APIClient
}

func NewWSGApplicationlogandstatusService(c *APIClient) *WSGApplicationlogandstatusService {
	s := new(WSGApplicationlogandstatusService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGApplicationlogandstatusService() *WSGApplicationlogandstatusService {
	return NewWSGApplicationlogandstatusService(ss.apiClient)
}

// FindApplicationsByBladeUUID
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
func (s *WSGApplicationlogandstatusService) FindApplicationsByBladeUUID(ctx context.Context, bladeUUID string, optionalParams map[string][]string) (*WSGAdministrationApplicationLogAndStatusList, *APIResponseMeta, error) {
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
	req.SetPathParameter("bladeUUID", bladeUUID)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationApplicationLogAndStatusList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApplicationsDownloadByBladeUUID
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
func (s *WSGApplicationlogandstatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, appName string, bladeUUID string, optionalParams map[string][]string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsDownloadByBladeUUID, true)
	req.SetQueryParameter("appName", []string{appName})
	req.SetPathParameter("bladeUUID", bladeUUID)
	if v, ok := optionalParams["logFileName"]; ok && len(v) > 0 {
		req.SetQueryParameter("logFileName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindApplicationsDownloadsnapByBladeUUID
//
// Use this API command to download snapshot logs.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGApplicationlogandstatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, bladeUUID string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsDownloadsnapByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateApplications
//
// Use this API command to modify log level of specified application.
//
// Request Body:
//	 - body *WSGAdministrationModifyLogLevel
func (s *WSGApplicationlogandstatusService) PartialUpdateApplications(ctx context.Context, body *WSGAdministrationModifyLogLevel) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
