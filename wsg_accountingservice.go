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
// Operation ID: addServicesAcctRadius
//
// Use this API command to create a new RADIUS accounting service.
//
// Request Body:
//	 - body *WSGServiceCreateRadiusAccounting
func (s *WSGAccountingServiceService) AddServicesAcctRadius(ctx context.Context, body *WSGServiceCreateRadiusAccounting, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddServicesAcctRadius, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddServicesAcctTestById
//
// Operation ID: addServicesAcctTestById
//
// Use this API command to test an accounting service.
//
// Request Body:
//	 - body *WSGServiceTestingConfig
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddServicesAcctTestById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteServicesAcct
//
// Operation ID: deleteServicesAcct
//
// Use this API command to delete a list of accounting service.
//
// Request Body:
//	 - body *WSGServiceDeleteBulkAccountingService
func (s *WSGAccountingServiceService) DeleteServicesAcct(ctx context.Context, body *WSGServiceDeleteBulkAccountingService, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServicesAcct, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteServicesAcctById
//
// Operation ID: deleteServicesAcctById
//
// Use this API command to delete an accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServicesAcctById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteServicesAcctRadiusById
//
// Operation ID: deleteServicesAcctRadiusById
//
// Use this API command to delete a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteServicesAcctRadiusSecondaryById
//
// Operation ID: deleteServicesAcctRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusSecondaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteServicesAcctRadiusStandbyPrimaryById
//
// Operation ID: deleteServicesAcctRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS Accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusStandbyPrimaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindServicesAaaServerAcctByQueryCriteria
//
// Operation ID: findServicesAaaServerAcctByQueryCriteria
//
// Query Non-Proxy Accounting AAAServers with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAaaServerAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAAAServerQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAAAServerQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAAAServerQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesAaaServerAcctByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAAAServerQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAAServerQueryListAPIResponse), err
}

// FindServicesAcctByQueryCriteria
//
// Operation ID: findServicesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceCommonAccountingServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceCommonAccountingServiceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGServiceCommonAccountingServiceListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesAcctByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGServiceCommonAccountingServiceListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceCommonAccountingServiceListAPIResponse), err
}

// FindServicesAcctRadius
//
// Operation ID: findServicesAcctRadius
//
// Use this API command to retrieve a list of RADIUS accounting services.
func (s *WSGAccountingServiceService) FindServicesAcctRadius(ctx context.Context, mutators ...RequestMutator) (*WSGServiceRadiusAccountingServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAccountingServiceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGServiceRadiusAccountingServiceListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindServicesAcctRadius, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAccountingServiceListAPIResponse), err
}

// FindServicesAcctRadiusById
//
// Operation ID: findServicesAcctRadiusById
//
// Use this API command to retrieve a RADIUS accounting service.
//
// Required Parameters:
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
	if err = ctx.Err(); err != nil {
		return resp.(*WSGServiceRadiusAccountingServiceAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindServicesAcctRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAccountingServiceAPIResponse), err
}

// FindServicesAcctRadiusByQueryCriteria
//
// Operation ID: findServicesAcctRadiusByQueryCriteria
//
// Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceRadiusAccountingServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAccountingServiceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGServiceRadiusAccountingServiceListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesAcctRadiusByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGServiceRadiusAccountingServiceListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAccountingServiceListAPIResponse), err
}

// PartialUpdateServicesAcctRadiusById
//
// Operation ID: partialUpdateServicesAcctRadiusById
//
// Use this API command to modify the configuration of a RADIUS accounting service.
//
// Request Body:
//	 - body *WSGServiceModifyRadiusAccounting
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) PartialUpdateServicesAcctRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAccounting, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateServicesAcctRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
