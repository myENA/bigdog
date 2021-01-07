package bigdog

// API Version: v9_1

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
func (s *WSGSMSGatewayService) FindSmsGateway(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSystemSmsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemSmsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemSmsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSmsGateway, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemSmsAPIResponse), err
}

// FindSmsGatewayByQueryCriteria
//
// Operation ID: findSmsGatewayByQueryCriteria
//
// Query SMS gateway.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSMSGatewayService) FindSmsGatewayByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSystemSmsListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemSmsListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemSmsListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindSmsGatewayByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSystemSmsListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemSmsListAPIResponse), err
}

// PartialUpdateSmsGateway
//
// Operation ID: partialUpdateSmsGateway
//
// Update SMS gateway.
//
// Request Body:
//	 - body *WSGSystemSms
func (s *WSGSMSGatewayService) PartialUpdateSmsGateway(ctx context.Context, body *WSGSystemSms, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSmsGateway, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
