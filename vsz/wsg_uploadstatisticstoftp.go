package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGUploadStatisticstoFTPService struct {
	apiClient *APIClient
}

func NewWSGUploadStatisticstoFTPService(c *APIClient) *WSGUploadStatisticstoFTPService {
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
func (s *WSGUploadStatisticstoFTPService) FindGlobalSettingsStatsFtp(ctx context.Context) (*WSGSystemFtpGlobalSetting, *APIResponseMeta, error) {
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateGlobalSettingsStatsFtp
//
// Use this API command to modify the setting of uploading statistical data to FTP server.
//
// Request Body:
//	 - body *WSGSystemFtpGlobalSetting
func (s *WSGUploadStatisticstoFTPService) PartialUpdateGlobalSettingsStatsFtp(ctx context.Context, body *WSGSystemFtpGlobalSetting) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateGlobalSettingsStatsFtp, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
