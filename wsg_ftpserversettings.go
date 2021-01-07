package bigdog

// API Version: v9_1

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
// Operation ID: addFtps
//
// Add FTP server.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFTPServerSettingsService) AddFtps(ctx context.Context, body *WSGSystemFtp, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddFtps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteFtps
//
// Operation ID: deleteFtps
//
// Remove FTP servers.
//
// Request Body:
//	 - body *WSGSystemDeleteBulkFtp
func (s *WSGFTPServerSettingsService) DeleteFtps(ctx context.Context, body *WSGSystemDeleteBulkFtp, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteFtps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteFtpsByFtpId
//
// Operation ID: deleteFtpsByFtpId
//
// Remove FTP server.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) DeleteFtpsByFtpId(ctx context.Context, ftpId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteFtpsByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindFtpsByFtpId
//
// Operation ID: findFtpsByFtpId
//
// Retrieve information of specific FTP server.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) FindFtpsByFtpId(ctx context.Context, ftpId string, mutators ...RequestMutator) (*WSGSystemFtpAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemFtpAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindFtpsByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpAPIResponse), err
}

// FindFtpsByQueryCriteria
//
// Operation ID: findFtpsByQueryCriteria
//
// Retrieve information of all FTP server.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFTPServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSystemFtpListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemFtpListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindFtpsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSystemFtpListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpListAPIResponse), err
}

// FindFtpsTest
//
// Operation ID: findFtpsTest
//
// Test ftp server of specific FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
func (s *WSGFTPServerSettingsService) FindFtpsTest(ctx context.Context, body *WSGSystemFtp, mutators ...RequestMutator) (*WSGSystemFtpTestResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpTestResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemFtpTestResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindFtpsTest, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSystemFtpTestResponseAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpTestResponseAPIResponse), err
}

// FindFtpsTestByFtpId
//
// Operation ID: findFtpsTestByFtpId
//
// Test ftp server of specific FTP server id.
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) FindFtpsTestByFtpId(ctx context.Context, ftpId string, mutators ...RequestMutator) (*WSGSystemFtpTestResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpTestResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemFtpTestResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindFtpsTestByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpTestResponseAPIResponse), err
}

// PartialUpdateFtpsByFtpId
//
// Operation ID: partialUpdateFtpsByFtpId
//
// Update FTP server settings.
//
// Request Body:
//	 - body *WSGSystemFtp
//
// Required Parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) PartialUpdateFtpsByFtpId(ctx context.Context, body *WSGSystemFtp, ftpId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateFtpsByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
