package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGIPSECProfileService struct {
	apiClient *APIClient
}

func NewWSGIPSECProfileService(c *APIClient) *WSGIPSECProfileService {
	s := new(WSGIPSECProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIPSECProfileService() *WSGIPSECProfileService {
	return NewWSGIPSECProfileService(ss.apiClient)
}

// AddProfilesTunnelIpsec
//
// Create a new ipsec.
//
// Request Body:
//	 - body *WSGProfileCreateIpsecProfile
func (s *WSGIPSECProfileService) AddProfilesTunnelIpsec(ctx context.Context, body *WSGProfileCreateIpsecProfile) (*WSGCommonCreateResult, error) {
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddProfilesTunnelIpsec, true)
}

// DeleteProfilesTunnelIpsec
//
// Delete multiple ipsec.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsec(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsec, true)
}

// DeleteProfilesTunnelIpsecById
//
// Delete a ipsec.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsecById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTunnelIpsecById, true)
}

// FindProfilesTunnelIpsec
//
// Retrieve a list of IPSEC.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsec(ctx context.Context, optionalParams map[string]interface{}) (*WSGProfileList, error) {
	var (
		resp *WSGProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelIpsec, true)
}

// FindProfilesTunnelIpsecById
//
// Retrieve a IPSEC.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecById(ctx context.Context, id string) (*WSGProfileIpsecProfile, error) {
	var (
		resp *WSGProfileIpsecProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelIpsecById, true)
}

// FindProfilesTunnelIpsecByQueryCriteria
//
// Query a list of IPSEC.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileIpsecProfileList, error) {
	var (
		resp *WSGProfileIpsecProfileList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTunnelIpsecByQueryCriteria, true)
}

// PartialUpdateProfilesTunnelIpsecById
//
// Modify a specific ipsec basic.
//
// Request Body:
//	 - body *WSGProfileModifyIpsecProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIPSECProfileService) PartialUpdateProfilesTunnelIpsecById(ctx context.Context, body *WSGProfileModifyIpsecProfile, id string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelIpsecById, true)
}
