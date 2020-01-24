package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGWiredClientService struct {
	apiClient *APIClient
}

func NewWSGWiredClientService(c *APIClient) *WSGWiredClientService {
	s := new(WSGWiredClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWiredClientService() *WSGWiredClientService {
	return NewWSGWiredClientService(ss.apiClient)
}

// AddWiredClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWiredClientService) AddWiredClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddWiredClientsBulkDeauth, true)
}

// AddWiredClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWiredClientService) AddWiredClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient) (*WSGCommonEmptyResult, error) {
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddWiredClientsDeauth, true)
}

// FindWiredclientByQueryCriteria
//
// Query wired clients with specified filters
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWiredClientService) FindWiredclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGWiredClientQueryClientQueryList, error) {
	var (
		resp *WSGWiredClientQueryClientQueryList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindWiredclientByQueryCriteria, true)
}
