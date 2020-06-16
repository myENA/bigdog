package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAuthenticationServiceService struct {
	apiClient *VSZClient
}

func NewWSGAuthenticationServiceService(c *VSZClient) *WSGAuthenticationServiceService {
	s := new(WSGAuthenticationServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAuthenticationServiceService() *WSGAuthenticationServiceService {
	return NewWSGAuthenticationServiceService(ss.apiClient)
}

// AddServicesAuthAd
//
// Operation ID: addServicesAuthAd
//
// Use this API command to create a new active directory authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateActiveDirectoryAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthAd(ctx context.Context, body *WSGServiceCreateActiveDirectoryAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthAd, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthHlr
//
// Operation ID: addServicesAuthHlr
//
// Use this API command to create a new hlr authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateHlrAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthHlr(ctx context.Context, body *WSGServiceCreateHlrAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthHlr, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthLdap
//
// Operation ID: addServicesAuthLdap
//
// Use this API command to create a new LDAP authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateLDAPAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthLdap(ctx context.Context, body *WSGServiceCreateLDAPAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthLdap, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthRadius
//
// Operation ID: addServicesAuthRadius
//
// Use this API command to create a new RADIUS authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateRadiusAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthRadius(ctx context.Context, body *WSGServiceCreateRadiusAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthRadius, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthTestById
//
// Operation ID: addServicesAuthTestById
//
// Use this API command to test an authentication service.
//
// Request Body:
//	 - body *WSGServiceTestingConfig
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) AddServicesAuthTestById(ctx context.Context, body *WSGServiceTestingConfig, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthTestById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuth
//
// Operation ID: deleteServicesAuth
//
// Use this API command to delete a list of authentication service.
//
// Request Body:
//	 - body *WSGServiceDeleteBulkAuthenticationService
func (s *WSGAuthenticationServiceService) DeleteServicesAuth(ctx context.Context, body *WSGServiceDeleteBulkAuthenticationService, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuth, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthAdById
//
// Operation ID: deleteServicesAuthAdById
//
// Use this API command to delete an active directory authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthAdById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthAdById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthById
//
// Operation ID: deleteServicesAuthById
//
// Use this API command to delete an authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthHlrById
//
// Operation ID: deleteServicesAuthHlrById
//
// Use this API command to delete a hlr authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthHlrById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthHlrById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthLdapById
//
// Operation ID: deleteServicesAuthLdapById
//
// Use this API command to delete a LDAP authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthLdapById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthLdapById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthRadiusById
//
// Operation ID: deleteServicesAuthRadiusById
//
// Use this API command to delete a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthRadiusById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthRadiusSecondaryById
//
// Operation ID: deleteServicesAuthRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusSecondaryById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthRadiusSecondaryById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthRadiusStandbyPrimaryById
//
// Operation ID: deleteServicesAuthRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusStandbyPrimaryById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthRadiusStandbyPrimaryById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindServicesAaaServerAuthByQueryCriteria
//
// Operation ID: findServicesAaaServerAuthByQueryCriteria
//
// Query Non-Proxy Authentication AAAServers with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAaaServerAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAAAServerQueryList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAaaServerAuthByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAAServerQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthAd
//
// Operation ID: findServicesAuthAd
//
// Use this API command to retrieve a list of active directory authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthAd(ctx context.Context, mutators ...RequestMutator) (*WSGServiceActiveDirectoryServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceActiveDirectoryServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthAd, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceActiveDirectoryServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthAdById
//
// Operation ID: findServicesAuthAdById
//
// Use this API command to retrieve an active directory authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthAdById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceActiveDirectoryService, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceActiveDirectoryService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthAdById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceActiveDirectoryService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthAdByQueryCriteria
//
// Operation ID: findServicesAuthAdByQueryCriteria
//
// Use this API command to retrieve a list of AD Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthAdByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceActiveDirectoryServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceActiveDirectoryServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthAdByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceActiveDirectoryServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthByQueryCriteria
//
// Operation ID: findServicesAuthByQueryCriteria
//
// Use this API command to retrieve a list of Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceCommonAuthenticationServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceCommonAuthenticationServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceCommonAuthenticationServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthGuestById
//
// Operation ID: findServicesAuthGuestById
//
// Use this API command to retrieve a Guest authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthGuestById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceCommonAuthenticationService, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceCommonAuthenticationService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthGuestById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceCommonAuthenticationService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthHlr
//
// Operation ID: findServicesAuthHlr
//
// Use this API command to retrieve a list of hlr authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthHlr(ctx context.Context, mutators ...RequestMutator) (*WSGServiceHlrServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceHlrServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthHlr, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceHlrServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthHlrById
//
// Operation ID: findServicesAuthHlrById
//
// Use this API command to retrieve a hlr authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceHlrService, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceHlrService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthHlrById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceHlrService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthHlrByQueryCriteria
//
// Operation ID: findServicesAuthHlrByQueryCriteria
//
// Use this API command to retrieve a list of hlr Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceHlrServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceHlrServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthHlrByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceHlrServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLdap
//
// Operation ID: findServicesAuthLdap
//
// Use this API command to retrieve a list of LDAP authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthLdap(ctx context.Context, mutators ...RequestMutator) (*WSGServiceLDAPServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceLDAPServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthLdap, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceLDAPServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLdapById
//
// Operation ID: findServicesAuthLdapById
//
// Use this API command to retrieve a LDAP authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceLDAPService, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceLDAPService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthLdapById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceLDAPService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLdapByQueryCriteria
//
// Operation ID: findServicesAuthLdapByQueryCriteria
//
// Use this API command to retrieve a list of LDAP Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceLDAPServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceLDAPServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthLdapByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceLDAPServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLocalDbById
//
// Operation ID: findServicesAuthLocal_dbById
//
// Use this API command to retrieve a LocalDB authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLocalDbById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceCommonAuthenticationService, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceCommonAuthenticationService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthLocalDbById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceCommonAuthenticationService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthRadius
//
// Operation ID: findServicesAuthRadius
//
// Use this API command to retrieve a list of RADIUS authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthRadius(ctx context.Context, mutators ...RequestMutator) (*WSGServiceRadiusAuthenticationServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceRadiusAuthenticationServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthRadius, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceRadiusAuthenticationServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthRadiusById
//
// Operation ID: findServicesAuthRadiusById
//
// Use this API command to retrieve a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceRadiusAuthenticationService, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceRadiusAuthenticationService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthRadiusById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceRadiusAuthenticationService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthRadiusByQueryCriteria
//
// Operation ID: findServicesAuthRadiusByQueryCriteria
//
// Use this API command to retrieve a list of radius Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceRadiusAuthenticationServiceList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceRadiusAuthenticationServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthRadiusByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceRadiusAuthenticationServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateServicesAuthAdById
//
// Operation ID: partialUpdateServicesAuthAdById
//
// Use this API command to modify the configuration of an active directory authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyActiveDirectoryAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthAdById(ctx context.Context, body *WSGServiceModifyActiveDirectoryAuthentication, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthAdById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthHlrById
//
// Operation ID: partialUpdateServicesAuthHlrById
//
// Use this API command to modify the configuration of a hlr authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyHlrAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthHlrById(ctx context.Context, body *WSGServiceModifyHlrAuthentication, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthHlrById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthLdapById
//
// Operation ID: partialUpdateServicesAuthLdapById
//
// Use this API command to modify the configuration of a LDAP authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyLDAPAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLdapById(ctx context.Context, body *WSGServiceModifyLDAPAuthentication, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthLdapById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthLocalDbById
//
// Operation ID: partialUpdateServicesAuthLocal_dbById
//
// Use this API command to update LocalDB authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyLocalDbAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLocalDbById(ctx context.Context, body *WSGServiceModifyLocalDbAuthentication, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthLocalDbById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthRadiusById
//
// Operation ID: partialUpdateServicesAuthRadiusById
//
// Use this API command to modify the configuration of a RADIUS authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyRadiusAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAuthentication, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthRadiusById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
