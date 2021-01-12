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
// Add FTP server.
//
// Operation ID: addFtps
// Operation path: /ftps
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGSystemFtp
func (s *WSGFTPServerSettingsService) AddFtps(ctx context.Context, body *WSGSystemFtp, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddFtps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteFtps
//
// Remove FTP servers.
//
// Operation ID: deleteFtps
// Operation path: /ftps
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSystemDeleteBulkFtp
func (s *WSGFTPServerSettingsService) DeleteFtps(ctx context.Context, body *WSGSystemDeleteBulkFtp, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteFtps, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteFtpsByFtpId
//
// Remove FTP server.
//
// Operation ID: deleteFtpsByFtpId
// Operation path: /ftps/{ftpId}
// Success code: 204 (No Content)
//
// Required parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) DeleteFtpsByFtpId(ctx context.Context, ftpId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteFtpsByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindFtpsByFtpId
//
// Retrieve information of specific FTP server.
//
// Operation ID: findFtpsByFtpId
// Operation path: /ftps/{ftpId}
// Success code: 200 (OK)
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindFtpsByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpAPIResponse), err
}

// FindFtpsByQueryCriteria
//
// Retrieve information of all FTP server.
//
// Operation ID: findFtpsByQueryCriteria
// Operation path: /ftps/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFTPServerSettingsService) FindFtpsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSystemFtpListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindFtpsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpListAPIResponse), err
}

// FindFtpsTest
//
// Test ftp server of specific FTP server settings.
//
// Operation ID: findFtpsTest
// Operation path: /ftps/test
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSystemFtp
func (s *WSGFTPServerSettingsService) FindFtpsTest(ctx context.Context, body *WSGSystemFtp, mutators ...RequestMutator) (*WSGSystemFtpTestResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemFtpTestResponseAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindFtpsTest, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpTestResponseAPIResponse), err
}

// FindFtpsTestByFtpId
//
// Test ftp server of specific FTP server id.
//
// Operation ID: findFtpsTestByFtpId
// Operation path: /ftps/test/{ftpId}
// Success code: 200 (OK)
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindFtpsTestByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemFtpTestResponseAPIResponse), err
}

// PartialUpdateFtpsByFtpId
//
// Update FTP server settings.
//
// Operation ID: partialUpdateFtpsByFtpId
// Operation path: /ftps/{ftpId}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemFtp
//
// Required parameters:
// - ftpId string
//		- required
func (s *WSGFTPServerSettingsService) PartialUpdateFtpsByFtpId(ctx context.Context, body *WSGSystemFtp, ftpId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateFtpsByFtpId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("ftpId", ftpId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
