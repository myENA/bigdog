package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGUploadstatisticstoftpService struct {
	apiClient *VSZClient
}

func NewWSGUploadstatisticstoftpService(c *VSZClient) *WSGUploadstatisticstoftpService {
	s := new(WSGUploadstatisticstoftpService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGUploadstatisticstoftpService() *WSGUploadstatisticstoftpService {
	return NewWSGUploadstatisticstoftpService(ss.apiClient)
}

// FindGlobalSettingsStatsFtp
//
// Use this API command to retrieve the uploading statistical data to FTP server setting.
func (s *WSGUploadstatisticstoftpService) FindGlobalSettingsStatsFtp(ctx context.Context) (*WSGSystemFtpGlobalSetting, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemFtpGlobalSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindGlobalSettingsStatsFtp, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtpGlobalSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateGlobalSettingsStatsFtp
//
// Use this API command to modify the setting of uploading statistical data to FTP server.
//
// Request Body:
//	 - body *WSGSystemFtpGlobalSetting
func (s *WSGUploadstatisticstoftpService) PartialUpdateGlobalSettingsStatsFtp(ctx context.Context, body *WSGSystemFtpGlobalSetting) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateGlobalSettingsStatsFtp, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
