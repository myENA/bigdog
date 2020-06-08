package ruckus

// API Version: v9_0

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
// Request Body:
//	 - body *WSGServiceCreateRadiusAccounting
func (s *WSGAccountingServiceService) AddServicesAcctRadius(ctx context.Context, body *WSGServiceCreateRadiusAccounting) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAcctRadius, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAcctTestById
//
// Use this API command to test an accounting service.
//
// Request Body:
//	 - body *WSGServiceTestingConfig
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) AddServicesAcctTestById(ctx context.Context, body *WSGServiceTestingConfig, id string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAcctTestById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteServicesAcct
//
// Use this API command to delete a list of accounting service.
//
// Request Body:
//	 - body *WSGServiceDeleteBulkAccountingService
func (s *WSGAccountingServiceService) DeleteServicesAcct(ctx context.Context, body *WSGServiceDeleteBulkAccountingService) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcct, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAcctById
//
// Use this API command to delete an accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAcctRadiusById
//
// Use this API command to delete a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAcctRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusSecondaryById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAcctRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS Accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAcctRadiusStandbyPrimaryById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindServicesAaaServerAcctByQueryCriteria
//
// Query Non-Proxy Accounting AAAServers with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAaaServerAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAAAServerQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAAServerQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAaaServerAcctByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAAAServerQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceCommonAccountingServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceCommonAccountingServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAcctByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceCommonAccountingServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAcctRadius
//
// Use this API command to retrieve a list of RADIUS accounting services.
func (s *WSGAccountingServiceService) FindServicesAcctRadius(ctx context.Context) (*WSGServiceRadiusAccountingServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceRadiusAccountingServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAcctRadius, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAccountingServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAcctRadiusById
//
// Use this API command to retrieve a RADIUS accounting service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) FindServicesAcctRadiusById(ctx context.Context, id string) (*WSGServiceRadiusAccountingService, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceRadiusAccountingService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAcctRadiusById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAccountingService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAcctRadiusByQueryCriteria
//
// Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceRadiusAccountingServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceRadiusAccountingServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAcctRadiusByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAccountingServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateServicesAcctRadiusById
//
// Use this API command to modify the configuration of a RADIUS accounting service.
//
// Request Body:
//	 - body *WSGServiceModifyRadiusAccounting
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountingServiceService) PartialUpdateServicesAcctRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAccounting, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAcctRadiusById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
