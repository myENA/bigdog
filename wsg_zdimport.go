package bigdog

// API Version: v9_0

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
// Request Body:
//	 - body *WSGAdministrationZdImport
func (s *WSGZDImportService) AddZdImportConnectZD(ctx context.Context, body *WSGAdministrationZdImport) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddZdImportConnectZD, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// AddZdImportMigrate
//
// Migrate ZD to SCG.
//
// Request Body:
//	 - body *WSGAdministrationZdImport
func (s *WSGZDImportService) AddZdImportMigrate(ctx context.Context, body *WSGAdministrationZdImport) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddZdImportMigrate, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// FindZdImportGetZDAPs
//
// Get ZD AP.
//
// Required Parameters:
// - ip string
//		- required
func (s *WSGZDImportService) FindZdImportGetZDAPs(ctx context.Context, ip string) (*WSGAdministrationZdAPList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindZdImportGetZDAPs, true)
	req.SetQueryParameter("ip", []string{ip})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationZdAPList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindZdImportStatus
//
// Get Migrate Status.
//
// Optional Parameters:
// - details string
//		- nullable
func (s *WSGZDImportService) FindZdImportStatus(ctx context.Context, optionalParams map[string][]string) (*WSGAdministrationZdImportStatus, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindZdImportStatus, true)
	if v, ok := optionalParams["details"]; ok && len(v) > 0 {
		req.SetQueryParameter("details", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdministrationZdImportStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
