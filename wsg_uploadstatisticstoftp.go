package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: findGlobalSettingsStatsFtp
//
// Use this API command to retrieve the uploading statistical data to FTP server setting.
func (s *WSGUploadStatisticstoFTPService) FindGlobalSettingsStatsFtp(ctx context.Context, mutators ...RequestMutator) (*WSGSystemFtpGlobalSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpGlobalSettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemFtpGlobalSettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindGlobalSettingsStatsFtp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpGlobalSettingAPIResponse), err
}

// PartialUpdateGlobalSettingsStatsFtp
//
// Operation ID: partialUpdateGlobalSettingsStatsFtp
//
// Use this API command to modify the setting of uploading statistical data to FTP server.
//
// Request Body:
//	 - body *WSGSystemFtpGlobalSetting
func (s *WSGUploadStatisticstoFTPService) PartialUpdateGlobalSettingsStatsFtp(ctx context.Context, body *WSGSystemFtpGlobalSetting, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateGlobalSettingsStatsFtp, true)
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
