package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGFTPServerSettingsService struct {
	apiClient *VSZClient
}

func NewWSGFTPServerSettingsService(c *VSZClient) *WSGFTPServerSettingsService {
	s := new(WSGFTPServerSettingsService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFTPServerSettingsService() *WSGFTPServerSettingsService {
	return NewWSGFTPServerSettingsService(ss.apiClient)
}

// AddFtps
//
// Add FTP server.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFTPServerSettingsService) AddFtps(ctx context.Context, body *WSGSystemFtp) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddFtps, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteFtps
//
// Remove FTP servers.
//
// Request Body:
//	 - body *WSGSystemDeleteBulkFtp
func (s *WSGFTPServerSettingsService) DeleteFtps(ctx context.Context, body *WSGSystemDeleteBulkFtp) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFtps, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteFtpsByFtpId
//
// Remove FTP server.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) DeleteFtpsByFtpId(ctx context.Context, ftpId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteFtpsByFtpId, true)
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindFtpsByFtpId
//
// Retrieve information of specific FTP server.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) FindFtpsByFtpId(ctx context.Context, ftpId string) (*WSGSystemFtp, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemFtp
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsByFtpId, true)
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtp()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFtpsByQueryCriteria
//
// Retrieve information of all FTP server.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFTPServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSystemFtpList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemFtpList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindFtpsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtpList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFtpsTest
//
// Test ftp server of specific FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFTPServerSettingsService) FindFtpsTest(ctx context.Context, body *WSGSystemFtp) (*WSGSystemFtpTestResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemFtpTestResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsTest, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtpTestResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFtpsTestByFtpId
//
// Test ftp server of specific FTP server id.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) FindFtpsTestByFtpId(ctx context.Context, ftpId string) (*WSGSystemFtpTestResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemFtpTestResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindFtpsTestByFtpId, true)
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFtpTestResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGFTPServerSettingsService) PartialUpdateFtpsByFtpId(ctx context.Context, body *WSGSystemFtp, ftpId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateFtpsByFtpId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
