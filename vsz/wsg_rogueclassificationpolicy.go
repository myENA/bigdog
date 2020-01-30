package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGRogueClassificationPolicyService struct {
	apiClient *APIClient
}

func NewWSGRogueClassificationPolicyService(c *APIClient) *WSGRogueClassificationPolicyService {
	s := new(WSGRogueClassificationPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRogueClassificationPolicyService() *WSGRogueClassificationPolicyService {
	return NewWSGRogueClassificationPolicyService(ss.apiClient)
}

// AddRkszonesRogueApPoliciesByZoneId
//
// Use this API command to create rogue AP policy.
//
// Request Body:
//	 - body *WSGProfileCreateRogueApPolicy
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) AddRkszonesRogueApPoliciesByZoneId(ctx context.Context, body *WSGProfileCreateRogueApPolicy, zoneId string) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesRogueApPoliciesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteRkszonesRogueApPoliciesById
//
// Use this API command to delete rogue AP policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) DeleteRkszonesRogueApPoliciesById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRogueApPoliciesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// DeleteRkszonesRogueApPoliciesByZoneId
//
// Use this API command to delete bulk rogue AP policy.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) DeleteRkszonesRogueApPoliciesByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRogueApPoliciesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindRkszonesRogueApPoliciesById
//
// Use this API command to retrieve rogue AP policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesById(ctx context.Context, id string, zoneId string) (*WSGProfileRogueApPolicy, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileRogueApPolicy
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRogueApPoliciesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRogueApPolicy()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindRkszonesRogueApPoliciesByZoneId
//
// Use this API command to retrieve a list of rogue AP policy.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesByZoneId(ctx context.Context, zoneId string) (*WSGProfileRogueApPolicyList, error) {
	var (
		req      *APIRequest
		resp     *WSGProfileRogueApPolicyList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRogueApPoliciesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRogueApPolicyList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateRkszonesRogueApPoliciesById
//
// Use this API command to modify rogue AP policy.
//
// Request Body:
//	 - body *WSGProfileUpdateRogueApPolicy
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) PartialUpdateRkszonesRogueApPoliciesById(ctx context.Context, body *WSGProfileUpdateRogueApPolicy, id string, zoneId string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesRogueApPoliciesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
