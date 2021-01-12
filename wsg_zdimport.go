package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGZDImportService struct {
	apiClient *VSZClient
}

func NewWSGZDImportService(c *VSZClient) *WSGZDImportService {
	s := new(WSGZDImportService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZDImportService() *WSGZDImportService {
	return NewWSGZDImportService(ss.apiClient)
}

// AddZdImportConnectZD
//
// Connect to ZD.
//
// Operation ID: addZdImportConnectZD
// Operation path: /zdImport/connectZD
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAdministrationConnectZD
func (s *WSGZDImportService) AddZdImportConnectZD(ctx context.Context, body *WSGAdministrationConnectZD, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddZdImportConnectZD, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// AddZdImportMigrate
//
// Migrate ZD to SCG.
//
// Operation ID: addZdImportMigrate
// Operation path: /zdImport/migrate
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAdministrationZdImport
func (s *WSGZDImportService) AddZdImportMigrate(ctx context.Context, body *WSGAdministrationZdImport, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddZdImportMigrate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindZdImportGetZDAPs
//
// Get ZD AP.
//
// Operation ID: findZdImportGetZDAPs
// Operation path: /zdImport/getZDAPs
// Success code: 200 (OK)
//
// Required parameters:
// - ip string
//		- required
func (s *WSGZDImportService) FindZdImportGetZDAPs(ctx context.Context, ip string, mutators ...RequestMutator) (*WSGAdministrationZdAPListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationZdAPListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindZdImportGetZDAPs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("ip", ip)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationZdAPListAPIResponse), err
}

// FindZdImportStatus
//
// Get Migrate Status.
//
// Operation ID: findZdImportStatus
// Operation path: /zdImport/status
// Success code: 200 (OK)
//
// Optional parameters:
// - details string
//		- nullable
func (s *WSGZDImportService) FindZdImportStatus(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationZdImportStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationZdImportStatusAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindZdImportStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["details"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("details", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationZdImportStatusAPIResponse), err
}
