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
func (s *WSGZDImportService) AddZdImportConnectZD(ctx context.Context, body *WSGAdministrationConnectZD, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddZdImportConnectZD, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// AddZdImportMigrate
//
// Operation ID: addZdImportMigrate
//
// Migrate ZD to SCG.
//
// Request Body:
//	 - body *WSGAdministrationZdImport
func (s *WSGZDImportService) AddZdImportMigrate(ctx context.Context, body *WSGAdministrationZdImport, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddZdImportMigrate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
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
func (s *WSGZDImportService) FindZdImportGetZDAPs(ctx context.Context, ip string, mutators ...RequestMutator) (*WSGAdministrationZdAPList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationZdAPList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindZdImportGetZDAPs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.QueryParams.Set("ip", ip)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAdministrationZdAPList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGZDImportService) FindZdImportStatus(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAdministrationZdImportStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdministrationZdImportStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindZdImportStatus, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["details"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("details", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAdministrationZdImportStatus()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
