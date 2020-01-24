package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGBonjourFencingPolicyService struct {
	apiClient *APIClient
}

func NewWSGBonjourFencingPolicyService(c *APIClient) *WSGBonjourFencingPolicyService {
	s := new(WSGBonjourFencingPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBonjourFencingPolicyService() *WSGBonjourFencingPolicyService {
	return NewWSGBonjourFencingPolicyService(ss.apiClient)
}

// AddRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to create Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGProfileCreateBonjourFencingPolicy
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) AddRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, body *WSGProfileCreateBonjourFencingPolicy, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesBonjourFencingPolicyByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteRkszonesBonjourFencingPolicy
//
// Use this API command to delete Bulk Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicy(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBonjourFencingPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteRkszonesBonjourFencingPolicyById
//
// Use this API command to delete Bonjour Fencing Policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicyById(ctx context.Context, id string) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBonjourFencingPolicyById, true)
}

// FindApsBonjourFencingStatisticByApMac
//
// Use this API command to get Bonjour Fencing Statistic per AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGBonjourFencingPolicyService) FindApsBonjourFencingStatisticByApMac(ctx context.Context, apMac string) (*WSGProfileBonjourFencingStatistic, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBonjourFencingStatistic
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsBonjourFencingStatisticByApMac, true)
}

// FindRkszonesBonjourFencingPolicyById
//
// Use this API command to retrieve Bonjour Fencing Policy.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyById(ctx context.Context, id string, zoneId string) (*WSGProfileBonjourFencingPolicy, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBonjourFencingPolicy
		err  error
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourFencingPolicyById, true)
}

// FindRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, zoneId string) (*WSGProfileBonjourFencingPolicyList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBonjourFencingPolicyList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesBonjourFencingPolicyByZoneId, true)
}

// FindServicesBonjourFencingPolicyByQueryCriteria
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBonjourFencingPolicyService) FindServicesBonjourFencingPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBonjourFencingPolicyList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBonjourFencingPolicyList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesBonjourFencingPolicyByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateRkszonesBonjourFencingPolicyById
//
// Use this API command to modify the basic information of Bonjour Fencing Policy.
//
// Request Body:
//	 - body *WSGProfileModifyBonjourFencingPolicy
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) PartialUpdateRkszonesBonjourFencingPolicyById(ctx context.Context, body *WSGProfileModifyBonjourFencingPolicy, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesBonjourFencingPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
