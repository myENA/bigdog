package bigdog

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
// Operation ID: findSmsGateway
//
// Get SMS gateway.
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSMSGatewayService) FindSmsGateway(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSystemSms, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSystemSms()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSmsGatewayByQueryCriteria
//
// Operation ID: findSmsGatewayByQueryCriteria
//
// Query SMS gateway.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSMSGatewayService) FindSmsGatewayByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSystemSmsList, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSystemSmsList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateSmsGateway
//
// Operation ID: partialUpdateSmsGateway
//
// Update SMS gateway.
//
// Request Body:
//	 - body *WSGSystemSms
func (s *WSGSMSGatewayService) PartialUpdateSmsGateway(ctx context.Context, body *WSGSystemSms, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSmsGateway, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
