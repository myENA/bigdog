package vsz

// API Version: v8_1

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
		req  *APIRequest
		resp interface{}
		err  error
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
}

// DeleteFtps
//
// Remove FTP servers.
//
// Request Body:
//	 - body *WSGSystemDeleteBulkFtp
func (s *WSGFtpServerSettingsService) DeleteFtps(ctx context.Context, body *WSGSystemDeleteBulkFtp) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
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
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, ftpId, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFtpsByFtpId, true)
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
		req  *APIRequest
		resp *WSGSystemFtp
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, ftpId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsByFtpId, true)
}

// FindFtpsByQueryCriteria
//
// Retrieve information of all FTP server.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFtpServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSystemFtpList, error) {
	var (
		req  *APIRequest
		resp *WSGSystemFtpList
		err  error
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
}

// FindFtpsTest
//
// Test ftp server of specific FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFtpServerSettingsService) FindFtpsTest(ctx context.Context, body *WSGSystemFtp) (*WSGSystemFtpTestResponse, error) {
	var (
		req  *APIRequest
		resp *WSGSystemFtpTestResponse
		err  error
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
		req  *APIRequest
		resp *WSGSystemFtpTestResponse
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, ftpId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsTestByFtpId, true)
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
		req *APIRequest
		err error
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
}
