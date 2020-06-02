package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAuthenticationserviceService struct {
	apiClient *APIClient
}

func NewWSGAuthenticationserviceService(c *APIClient) *WSGAuthenticationserviceService {
	s := new(WSGAuthenticationserviceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAuthenticationserviceService() *WSGAuthenticationserviceService {
	return NewWSGAuthenticationserviceService(ss.apiClient)
}

// AddServicesAuthAd
//
// Use this API command to create a new active directory authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateActiveDirectoryAuthentication
func (s *WSGAuthenticationserviceService) AddServicesAuthAd(ctx context.Context, body *WSGServiceCreateActiveDirectoryAuthentication) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthHlr
//
// Use this API command to create a new hlr authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateHlrAuthentication
func (s *WSGAuthenticationserviceService) AddServicesAuthHlr(ctx context.Context, body *WSGServiceCreateHlrAuthentication) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthLdap
//
// Use this API command to create a new LDAP authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateLDAPAuthentication
func (s *WSGAuthenticationserviceService) AddServicesAuthLdap(ctx context.Context, body *WSGServiceCreateLDAPAuthentication) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthRadius
//
// Use this API command to create a new RADIUS authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateRadiusAuthentication
func (s *WSGAuthenticationserviceService) AddServicesAuthRadius(ctx context.Context, body *WSGServiceCreateRadiusAuthentication) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddServicesAuthTestById
//
// Use this API command to test an authentication service.
//
// Request Body:
//	 - body *WSGServiceTestingConfig
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) AddServicesAuthTestById(ctx context.Context, body *WSGServiceTestingConfig, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuth
//
// Use this API command to delete a list of authentication service.
//
// Request Body:
//	 - body *WSGServiceDeleteBulkAuthenticationService
func (s *WSGAuthenticationserviceService) DeleteServicesAuth(ctx context.Context, body *WSGServiceDeleteBulkAuthenticationService) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthAdById
//
// Use this API command to delete an active directory authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) DeleteServicesAuthAdById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthById
//
// Use this API command to delete an authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) DeleteServicesAuthById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthHlrById
//
// Use this API command to delete a hlr authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) DeleteServicesAuthHlrById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthLdapById
//
// Use this API command to delete a LDAP authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) DeleteServicesAuthLdapById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthRadiusById
//
// Use this API command to delete a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) DeleteServicesAuthRadiusById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) DeleteServicesAuthRadiusSecondaryById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesAuthRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) DeleteServicesAuthRadiusStandbyPrimaryById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindServicesAaaServerAuthByQueryCriteria
//
// Query Non-Proxy Authentication AAAServers with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationserviceService) FindServicesAaaServerAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAAAServerQueryList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAAAServerQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthAd
//
// Use this API command to retrieve a list of active directory authentication services.
func (s *WSGAuthenticationserviceService) FindServicesAuthAd(ctx context.Context) (*WSGServiceActiveDirectoryServiceList, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceActiveDirectoryServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthAdById
//
// Use this API command to retrieve an active directory authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) FindServicesAuthAdById(ctx context.Context, id string) (*WSGServiceActiveDirectoryService, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceActiveDirectoryService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthAdByQueryCriteria
//
// Use this API command to retrieve a list of AD Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationserviceService) FindServicesAuthAdByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceActiveDirectoryServiceList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceActiveDirectoryServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthByQueryCriteria
//
// Use this API command to retrieve a list of Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationserviceService) FindServicesAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceCommonAuthenticationServiceList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceCommonAuthenticationServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthGuestById
//
// Use this API command to retrieve a Guest authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) FindServicesAuthGuestById(ctx context.Context, id string) (*WSGServiceCommonAuthenticationService, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceCommonAuthenticationService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthHlr
//
// Use this API command to retrieve a list of hlr authentication services.
func (s *WSGAuthenticationserviceService) FindServicesAuthHlr(ctx context.Context) (*WSGServiceHlrServiceList, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceHlrServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthHlrById
//
// Use this API command to retrieve a hlr authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) FindServicesAuthHlrById(ctx context.Context, id string) (*WSGServiceHlrService, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceHlrService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthHlrByQueryCriteria
//
// Use this API command to retrieve a list of hlr Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationserviceService) FindServicesAuthHlrByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceHlrServiceList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceHlrServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLdap
//
// Use this API command to retrieve a list of LDAP authentication services.
func (s *WSGAuthenticationserviceService) FindServicesAuthLdap(ctx context.Context) (*WSGServiceLDAPServiceList, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceLDAPServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLdapById
//
// Use this API command to retrieve a LDAP authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) FindServicesAuthLdapById(ctx context.Context, id string) (*WSGServiceLDAPService, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceLDAPService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLdapByQueryCriteria
//
// Use this API command to retrieve a list of LDAP Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationserviceService) FindServicesAuthLdapByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceLDAPServiceList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceLDAPServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthLocal_dbById
//
// Use this API command to retrieve a LocalDB authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) FindServicesAuthLocal_dbById(ctx context.Context, id string) (*WSGServiceCommonAuthenticationService, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthLocal_dbById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceCommonAuthenticationService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthRadius
//
// Use this API command to retrieve a list of RADIUS authentication services.
func (s *WSGAuthenticationserviceService) FindServicesAuthRadius(ctx context.Context) (*WSGServiceRadiusAuthenticationServiceList, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAuthenticationServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthRadiusById
//
// Use this API command to retrieve a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) FindServicesAuthRadiusById(ctx context.Context, id string) (*WSGServiceRadiusAuthenticationService, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAuthenticationService()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesAuthRadiusByQueryCriteria
//
// Use this API command to retrieve a list of radius Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationserviceService) FindServicesAuthRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceRadiusAuthenticationServiceList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAuthenticationServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateServicesAuthAdById
//
// Use this API command to modify the configuration of an active directory authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyActiveDirectoryAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) PartialUpdateServicesAuthAdById(ctx context.Context, body *WSGServiceModifyActiveDirectoryAuthentication, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthHlrById
//
// Use this API command to modify the configuration of a hlr authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyHlrAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) PartialUpdateServicesAuthHlrById(ctx context.Context, body *WSGServiceModifyHlrAuthentication, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthLdapById
//
// Use this API command to modify the configuration of a LDAP authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyLDAPAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) PartialUpdateServicesAuthLdapById(ctx context.Context, body *WSGServiceModifyLDAPAuthentication, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthLocal_dbById
//
// Use this API command to update LocalDB authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyLocalDbAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) PartialUpdateServicesAuthLocal_dbById(ctx context.Context, body *WSGServiceModifyLocalDbAuthentication, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthLocal_dbById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesAuthRadiusById
//
// Use this API command to modify the configuration of a RADIUS authentication service.
//
// Request Body:
//	 - body *WSGServiceModifyRadiusAuthentication
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationserviceService) PartialUpdateServicesAuthRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAuthentication, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
