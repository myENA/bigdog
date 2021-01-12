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
// Get SMS gateway.
//
// Operation ID: findSmsGateway
// Operation path: /smsGateway
// Success code: 200 (OK)
//
// Optional parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSmsGateway, true)
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
// Query SMS gateway.
//
// Operation ID: findSmsGatewayByQueryCriteria
// Operation path: /smsGateway/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSMSGatewayService) FindSmsGatewayByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSystemSmsListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemSmsListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindSmsGatewayByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemSmsListAPIResponse), err
}

// PartialUpdateSmsGateway
//
// Update SMS gateway.
//
// Operation ID: partialUpdateSmsGateway
// Operation path: /smsGateway
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSystemSms
func (s *WSGSMSGatewayService) PartialUpdateSmsGateway(ctx context.Context, body *WSGSystemSms, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateSmsGateway, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
