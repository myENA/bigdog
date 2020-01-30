package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGBonjourGatewayPoliciesService struct {
	apiClient *APIClient
}

func NewWSGBonjourGatewayPoliciesService(c *APIClient) *WSGBonjourGatewayPoliciesService {
	s := new(WSGBonjourGatewayPoliciesService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBonjourGatewayPoliciesService() *WSGBonjourGatewayPoliciesService {
	return NewWSGBonjourGatewayPoliciesService(ss.apiClient)
}

// AddRkszonesBonjourGatewayEnableByZoneId
//
// Use this API command to enable/disable bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneModifyBonjourGatewayEnable
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayEnableByZoneId(ctx context.Context, body *WSGZoneModifyBonjourGatewayEnable, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesBonjourGatewayEnableByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to create bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneCreateBonjourGatewayPolicy
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) AddRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, body *WSGZoneCreateBonjourGatewayPolicy, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesBonjourGatewayPoliciesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesBonjourGatewayPoliciesById
//
// Use this API command to delete bonjour gateway policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) DeleteRkszonesBonjourGatewayPoliciesById(ctx context.Context, id string, zoneId string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBonjourGatewayPoliciesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindRkszonesBonjourGatewayPoliciesById
//
// Use this API command to retrieve bonjour gateway policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesById(ctx context.Context, id string, zoneId string) (*WSGZoneBonjourGatewayPolicyConfiguration, error) {
	var (
		req      *APIRequest
		resp     *WSGZoneBonjourGatewayPolicyConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourGatewayPoliciesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGZoneBonjourGatewayPolicyConfiguration()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesBonjourGatewayPoliciesByZoneId
//
// Use this API command to retrieve a list of bonjour gateway policies.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) FindRkszonesBonjourGatewayPoliciesByZoneId(ctx context.Context, zoneId string) (*WSGZoneBonjourGatewayPolicyList, error) {
	var (
		req      *APIRequest
		resp     *WSGZoneBonjourGatewayPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourGatewayPoliciesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGZoneBonjourGatewayPolicyList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesBonjourPolicyByQueryCriteria
//
// Query bonjourPolicy Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBonjourGatewayPoliciesService) FindServicesBonjourPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesBonjourPolicyByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// PartialUpdateRkszonesBonjourGatewayPoliciesById
//
// Use this API command to modify the configuration of bonjour gateway policy.
//
// Request Body:
//	 - body *WSGZoneModifyBonjourGatewayPolicy
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourGatewayPoliciesService) PartialUpdateRkszonesBonjourGatewayPoliciesById(ctx context.Context, body *WSGZoneModifyBonjourGatewayPolicy, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesBonjourGatewayPoliciesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
