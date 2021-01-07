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
// Operation ID: addZdImportConnectZD
//
// Connect to ZD.
//
// Request Body:
//	 - body *WSGAdministrationConnectZD
func (s *WSGZDImportService) AddZdImportConnectZD(ctx context.Context, body *WSGAdministrationConnectZD, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddZdImportConnectZD, true)
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

// AddZdImportMigrate
//
// Operation ID: addZdImportMigrate
//
// Migrate ZD to SCG.
//
// Request Body:
//	 - body *WSGAdministrationZdImport
func (s *WSGZDImportService) AddZdImportMigrate(ctx context.Context, body *WSGAdministrationZdImport, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddZdImportMigrate, true)
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

// FindZdImportGetZDAPs
//
// Operation ID: findZdImportGetZDAPs
//
// Get ZD AP.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationZdAPListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindZdImportGetZDAPs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("ip", ip)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationZdAPListAPIResponse), err
}

// FindZdImportStatus
//
// Operation ID: findZdImportStatus
//
// Get Migrate Status.
//
// Optional Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationZdImportStatusAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindZdImportStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["details"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("details", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationZdImportStatusAPIResponse), err
}
