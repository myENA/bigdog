package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGAccountingServiceService struct {
	apiClient *VSZClient
}

func NewWSGAccountingServiceService(c *VSZClient) *WSGAccountingServiceService {
	s := new(WSGAccountingServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountingServiceService() *WSGAccountingServiceService {
	return NewWSGAccountingServiceService(ss.apiClient)
}

// AddServicesAcctRadius
//
// Use this API command to create a new RADIUS accounting service.
//
// Operation ID: addServicesAcctRadius
// Operation path: /services/acct/radius
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGServiceCreateRadiusAccounting
func (s *WSGAccountingServiceService) AddServicesAcctRadius(ctx context.Context, body *WSGServiceCreateRadiusAccounting, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServicesAcctRadius, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddServicesAcctTestById
//
// Use this API command to test an accounting service.
//
// Operation ID: addServicesAcctTestById
// Operation path: /services/acct/test/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGServiceTestingConfig
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) AddServicesAcctTestById(ctx context.Context, body *WSGServiceTestingConfig, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServicesAcctTestById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteServicesAcct
//
// Use this API command to delete a list of accounting service.
//
// Operation ID: deleteServicesAcct
// Operation path: /services/acct
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceDeleteBulkAccountingService
func (s *WSGAccountingServiceService) DeleteServicesAcct(ctx context.Context, body *WSGServiceDeleteBulkAccountingService, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAcct, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAcctById
//
// Use this API command to delete an accounting service.
//
// Operation ID: deleteServicesAcctById
// Operation path: /services/acct/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAcctById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAcctRadiusById
//
// Use this API command to delete a RADIUS accounting service.
//
// Operation ID: deleteServicesAcctRadiusById
// Operation path: /services/acct/radius/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAcctRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAcctRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
//
// Operation ID: deleteServicesAcctRadiusSecondaryById
// Operation path: /services/acct/radius/{id}/secondary
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAcctRadiusSecondaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAcctRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS Accounting service.
//
// Operation ID: deleteServicesAcctRadiusStandbyPrimaryById
// Operation path: /services/acct/radius/{id}/standbyPrimary
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAcctRadiusStandbyPrimaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindServicesAaaServerAcctByQueryCriteria
//
// Query Non-Proxy Accounting AAAServers with specified filters.
//
// Operation ID: findServicesAaaServerAcctByQueryCriteria
// Operation path: /query/services/aaaServer/acct
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAaaServerAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAAAServerQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAAAServerQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAaaServerAcctByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAAServerQueryListAPIResponse), err
}

// FindServicesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting services by query criteria.
//
// Operation ID: findServicesAcctByQueryCriteria
// Operation path: /services/acct/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceCommonAccountingServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceCommonAccountingServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAcctByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceCommonAccountingServiceListAPIResponse), err
}

// FindServicesAcctRadius
//
// Use this API command to retrieve a list of RADIUS accounting services.
//
// Operation ID: findServicesAcctRadius
// Operation path: /services/acct/radius
// Success code: 200 (OK)
func (s *WSGAccountingServiceService) FindServicesAcctRadius(ctx context.Context, mutators ...RequestMutator) (*WSGServiceRadiusAccountingServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAccountingServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAcctRadius, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAccountingServiceListAPIResponse), err
}

// FindServicesAcctRadiusById
//
// Use this API command to retrieve a RADIUS accounting service.
//
// Operation ID: findServicesAcctRadiusById
// Operation path: /services/acct/radius/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) FindServicesAcctRadiusById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceRadiusAccountingServiceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAccountingServiceAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAcctRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAccountingServiceAPIResponse), err
}

// FindServicesAcctRadiusByQueryCriteria
//
// Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Operation ID: findServicesAcctRadiusByQueryCriteria
// Operation path: /services/acct/radius/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceRadiusAccountingServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAccountingServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAcctRadiusByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAccountingServiceListAPIResponse), err
}

// PartialUpdateServicesAcctRadiusById
//
// Use this API command to modify the configuration of a RADIUS accounting service.
//
// Operation ID: partialUpdateServicesAcctRadiusById
// Operation path: /services/acct/radius/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceModifyRadiusAccounting
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) PartialUpdateServicesAcctRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAccounting, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesAcctRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
