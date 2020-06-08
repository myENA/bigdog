package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSMSGatewayService struct {
	apiClient *VSZClient
}

func NewWSGSMSGatewayService(c *VSZClient) *WSGSMSGatewayService {
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
func (s *WSGSMSGatewayService) FindSmsGateway(ctx context.Context, optionalParams map[string][]string) (*WSGSystemSms, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemSms
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSmsGateway, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemSms()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSmsGatewayByQueryCriteria
//
// Query SMS gateway.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSMSGatewayService) FindSmsGatewayByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSystemSmsList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemSmsList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindSmsGatewayByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemSmsList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateSmsGateway
//
// Update SMS gateway.
//
// Request Body:
//	 - body *WSGSystemSms
func (s *WSGSMSGatewayService) PartialUpdateSmsGateway(ctx context.Context, body *WSGSystemSms) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSmsGateway, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
