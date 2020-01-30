package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGFtpServerSettingsService struct {
	apiClient *APIClient
}

func NewWSGFtpServerSettingsService(c *APIClient) *WSGFtpServerSettingsService {
	s := new(WSGFtpServerSettingsService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFtpServerSettingsService() *WSGFtpServerSettingsService {
	return NewWSGFtpServerSettingsService(ss.apiClient)
}

// AddFtps
//
// Add FTP server.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFtpServerSettingsService) AddFtps(ctx context.Context, body *WSGSystemFtp) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddFtps, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusCreated, httpResp, resp, err)
}

// DeleteFtps
//
// Remove FTP servers.
//
// Request Body:
//	 - body *WSGSystemDeleteBulkFtp
func (s *WSGFtpServerSettingsService) DeleteFtps(ctx context.Context, body *WSGSystemDeleteBulkFtp) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFtps, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// DeleteFtpsByFtpId
//
// Remove FTP server.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFtpServerSettingsService) DeleteFtpsByFtpId(ctx context.Context, ftpId string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, ftpId, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFtpsByFtpId, true)
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindFtpsByFtpId
//
// Retrieve information of specific FTP server.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFtpServerSettingsService) FindFtpsByFtpId(ctx context.Context, ftpId string) (*WSGSystemFtp, error) {
	var (
		req      *APIRequest
		resp     *WSGSystemFtp
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, ftpId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsByFtpId, true)
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtp()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindFtpsByQueryCriteria
//
// Retrieve information of all FTP server.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFtpServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSystemFtpList, error) {
	var (
		req      *APIRequest
		resp     *WSGSystemFtpList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindFtpsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtpList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindFtpsTest
//
// Test ftp server of specific FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFtpServerSettingsService) FindFtpsTest(ctx context.Context, body *WSGSystemFtp) (*WSGSystemFtpTestResponse, error) {
	var (
		req      *APIRequest
		resp     *WSGSystemFtpTestResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsTest, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtpTestResponse()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindFtpsTestByFtpId
//
// Test ftp server of specific FTP server id.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFtpServerSettingsService) FindFtpsTestByFtpId(ctx context.Context, ftpId string) (*WSGSystemFtpTestResponse, error) {
	var (
		req      *APIRequest
		resp     *WSGSystemFtpTestResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, ftpId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsTestByFtpId, true)
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtpTestResponse()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateFtpsByFtpId
//
// Update FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFtpServerSettingsService) PartialUpdateFtpsByFtpId(ctx context.Context, body *WSGSystemFtp, ftpId string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, ftpId, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateFtpsByFtpId, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}
