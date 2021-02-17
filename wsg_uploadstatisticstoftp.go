package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGUploadStatisticstoFTPService struct {
	apiClient *VSZClient
}

func NewWSGUploadStatisticstoFTPService(c *VSZClient) *WSGUploadStatisticstoFTPService {
	s := new(WSGUploadStatisticstoFTPService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGUploadStatisticstoFTPService() *WSGUploadStatisticstoFTPService {
	return NewWSGUploadStatisticstoFTPService(ss.apiClient)
}

// FindGlobalSettingsStatsFtp
//
// Use this API command to retrieve the uploading statistical data to FTP server setting.
//
// Operation ID: findGlobalSettingsStatsFtp
// Operation path: /globalSettings/statsFtp
// Success code: 200 (OK)
func (s *WSGUploadStatisticstoFTPService) FindGlobalSettingsStatsFtp(ctx context.Context, mutators ...RequestMutator) (*WSGSystemFtpGlobalSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpGlobalSettingAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindGlobalSettingsStatsFtp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSystemFtpGlobalSettingAPIResponse), err
}

// PartialUpdateGlobalSettingsStatsFtp
//
// Use this API command to modify the setting of uploading statistical data to FTP server.
//
// Operation ID: partialUpdateGlobalSettingsStatsFtp
// Operation path: /globalSettings/statsFtp
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemFtpGlobalSetting
func (s *WSGUploadStatisticstoFTPService) PartialUpdateGlobalSettingsStatsFtp(ctx context.Context, body *WSGSystemFtpGlobalSetting, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateGlobalSettingsStatsFtp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
