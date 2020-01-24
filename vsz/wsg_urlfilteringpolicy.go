package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGURLFilteringPolicyService struct {
	apiClient *APIClient
}

func NewWSGURLFilteringPolicyService(c *APIClient) *WSGURLFilteringPolicyService {
	s := new(WSGURLFilteringPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGURLFilteringPolicyService() *WSGURLFilteringPolicyService {
	return NewWSGURLFilteringPolicyService(ss.apiClient)
}

// AddUrlFilteringUrlFilteringPolicy
//
// Use this API command to create a URL Filtering policy.
//
// Request Body:
//	 - body *WSGURLFilteringCreateUrlFilteringPolicy
func (s *WSGURLFilteringPolicyService) AddUrlFilteringUrlFilteringPolicy(ctx context.Context, body *WSGURLFilteringCreateUrlFilteringPolicy) (*WSGCommonCreateResult, error) {
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddUrlFilteringUrlFilteringPolicy, true)
}

// DeleteUrlFilteringUrlFilteringPolicy
//
// Use this API command to delete bulk URL Filtering policies.
//
// Request Body:
//	 - body *WSGURLFilteringDeleteBulk
func (s *WSGURLFilteringPolicyService) DeleteUrlFilteringUrlFilteringPolicy(ctx context.Context, body *WSGURLFilteringDeleteBulk) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteUrlFilteringUrlFilteringPolicy, true)
}

// DeleteUrlFilteringUrlFilteringPolicyById
//
// Use this API command to delete a URL Filtering policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGURLFilteringPolicyService) DeleteUrlFilteringUrlFilteringPolicyById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteUrlFilteringUrlFilteringPolicyById, true)
}

// FindUrlFilteringBlockCategories
//
// Use this API command to retrieve the block categories of URL Filtering.
func (s *WSGURLFilteringPolicyService) FindUrlFilteringBlockCategories(ctx context.Context) (*WSGURLFilteringBlockCategoriesList, error) {
	var (
		resp *WSGURLFilteringBlockCategoriesList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUrlFilteringBlockCategories, true)
}

// FindUrlFilteringByQueryCriteria
//
// Use this API command to retrieve a list of URL Filtering policies by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGURLFilteringPolicyService) FindUrlFilteringByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGURLFilteringPolicyList, error) {
	var (
		resp *WSGURLFilteringPolicyList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindUrlFilteringByQueryCriteria, true)
}

// FindUrlFilteringUrlFilteringPolicy
//
// Use this API command to retrieve list of URL Filtering policies.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicy(ctx context.Context, optionalParams map[string]interface{}) (*WSGURLFilteringPolicyList, error) {
	var (
		resp *WSGURLFilteringPolicyList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUrlFilteringUrlFilteringPolicy, true)
}

// FindUrlFilteringUrlFilteringPolicyById
//
// Use this API command to retrieve an URL Filtering policy.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicyById(ctx context.Context, id string) (*WSGURLFilteringPolicy, error) {
	var (
		resp *WSGURLFilteringPolicy
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindUrlFilteringUrlFilteringPolicyById, true)
}

// PartialUpdateUrlFilteringUrlFilteringPolicyById
//
// Use this API command to patch a URL Filtering policy.
//
// Request Body:
//	 - body *WSGURLFilteringModifyUrlFilteringPolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGURLFilteringPolicyService) PartialUpdateUrlFilteringUrlFilteringPolicyById(ctx context.Context, body *WSGURLFilteringModifyUrlFilteringPolicy, id string) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateUrlFilteringUrlFilteringPolicyById, true)
}
