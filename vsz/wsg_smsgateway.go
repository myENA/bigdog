package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGSMSGatewayService struct {
	apiClient *APIClient
}

func NewWSGSMSGatewayService(c *APIClient) *WSGSMSGatewayService {
	s := new(WSGSMSGatewayService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSMSGatewayService() *WSGSMSGatewayService {
	return NewWSGSMSGatewayService(ss.apiClient)
}

// FindSmsGateway
//
// Get SMS gateway.
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSMSGatewayService) FindSmsGateway(ctx context.Context, optionalParams map[string]interface{}) (*WSGSystemSms, error) {
	var (
		resp *WSGSystemSms
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindSmsGateway, true)
}

// FindSmsGatewayByQueryCriteria
//
// Query SMS gateway.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSMSGatewayService) FindSmsGatewayByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSystemSmsList, error) {
	var (
		resp *WSGSystemSmsList
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
	req := NewAPIRequest(http.MethodPost, RouteWSGFindSmsGatewayByQueryCriteria, true)
}

// PartialUpdateSmsGateway
//
// Update SMS gateway.
//
// Request Body:
//	 - body *WSGSystemSms
func (s *WSGSMSGatewayService) PartialUpdateSmsGateway(ctx context.Context, body *WSGSystemSms) (interface{}, error) {
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
	req := NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSmsGateway, true)
}
