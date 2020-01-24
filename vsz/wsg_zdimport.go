package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGZDImportService struct {
	apiClient *APIClient
}

func NewWSGZDImportService(c *APIClient) *WSGZDImportService {
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
func (s *WSGZDImportService) AddZdImportConnectZD(ctx context.Context, body *WSGAdministrationZdImport) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddZdImportConnectZD, true)
}

// AddZdImportMigrate
//
// Migrate ZD to SCG.
//
// Request Body:
//	 - body *WSGAdministrationZdImport
func (s *WSGZDImportService) AddZdImportMigrate(ctx context.Context, body *WSGAdministrationZdImport) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddZdImportMigrate, true)
}

// FindZdImportGetZDAPs
//
// Get ZD AP.
//
// Required Parameters:
// - ip string
//		- required
func (s *WSGZDImportService) FindZdImportGetZDAPs(ctx context.Context, ip string) (*WSGAdministrationZdAPList, error) {
	var (
		resp *WSGAdministrationZdAPList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, ip, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindZdImportGetZDAPs, true)
}

// FindZdImportStatus
//
// Get Migrate Status.
//
// Optional Parameters:
// - details string
//		- nullable
func (s *WSGZDImportService) FindZdImportStatus(ctx context.Context, optionalParams map[string]interface{}) (*WSGAdministrationZdImportStatus, error) {
	var (
		resp *WSGAdministrationZdImportStatus
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindZdImportStatus, true)
}
