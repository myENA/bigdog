package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGDNSServerManagementService struct {
	apiClient *APIClient
}

func NewWSGDNSServerManagementService(c *APIClient) *WSGDNSServerManagementService {
	s := new(WSGDNSServerManagementService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDNSServerManagementService() *WSGDNSServerManagementService {
	return NewWSGDNSServerManagementService(ss.apiClient)
}

// AddProfilesDnsserver
//
// Use this API command to create DNS server profile.
//
// Request Body:
//	 - body *WSGProfileCreateDnsServerProfile
func (s *WSGDNSServerManagementService) AddProfilesDnsserver(ctx context.Context, body *WSGProfileCreateDnsServerProfile) (*WSGCommonCreateResult, error) {
	var (
		resp *WSGCommonCreateResult
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddProfilesDnsserver, true)
}

// AddProfilesDnsserverCloneById
//
// Use this API command to clone an DNS server profile.
//
// Request Body:
//	 - body *WSGProfileCloneRequest
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDNSServerManagementService) AddProfilesDnsserverCloneById(ctx context.Context, body *WSGProfileCloneRequest, id string) (*WSGProfileCloneResponse, error) {
	var (
		resp *WSGProfileCloneResponse
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteWSGAddProfilesDnsserverCloneById, true)
}

// DeleteProfilesDnsserver
//
// Use this API command to delete a list of DNS server profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserver(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesDnsserver, true)
}

// DeleteProfilesDnsserverById
//
// Use this API command to delete DNS server profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDNSServerManagementService) DeleteProfilesDnsserverById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesDnsserverById, true)
}

// FindProfilesDnsserver
//
// Use this API command to retrieve a list of DNS server profile.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGDNSServerManagementService) FindProfilesDnsserver(ctx context.Context, optionalParams map[string]interface{}) (*WSGProfileDnsServerProfileList, error) {
	var (
		resp *WSGProfileDnsServerProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesDnsserver, true)
}

// FindProfilesDnsserverById
//
// Use this API command to retrieve DNS server profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDNSServerManagementService) FindProfilesDnsserverById(ctx context.Context, id string) (*WSGProfileDnsServerProfile, error) {
	var (
		resp *WSGProfileDnsServerProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesDnsserverById, true)
}

// FindProfilesDnsserverByQueryCriteria
//
// Use this API command to retrieve a list of DNS server profile  by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDNSServerManagementService) FindProfilesDnsserverByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileDnsServerProfileList, error) {
	var (
		resp *WSGProfileDnsServerProfileList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindProfilesDnsserverByQueryCriteria, true)
}

// PartialUpdateProfilesDnsserverById
//
// Use this API command to modify the basic information of DNS server profile.
//
// Request Body:
//	 - body *WSGProfileModifyDnsServerProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDNSServerManagementService) PartialUpdateProfilesDnsserverById(ctx context.Context, body *WSGProfileModifyDnsServerProfile, id string) (*WSGCommonEmptyResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesDnsserverById, true)
}
