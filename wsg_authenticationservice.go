package bigdog

// API Version: v9_1

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
// Use this API command to create a new active directory authentication service.
//
// Operation ID: addServicesAuthAd
// Operation path: /services/auth/ad
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGServiceCreateActiveDirectoryAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthAd(ctx context.Context, body *WSGServiceCreateActiveDirectoryAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServicesAuthAd, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddServicesAuthHlr
//
// Use this API command to create a new hlr authentication service.
//
// Operation ID: addServicesAuthHlr
// Operation path: /services/auth/hlr
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGServiceCreateHlrAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthHlr(ctx context.Context, body *WSGServiceCreateHlrAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServicesAuthHlr, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddServicesAuthLdap
//
// Use this API command to create a new LDAP authentication service.
//
// Operation ID: addServicesAuthLdap
// Operation path: /services/auth/ldap
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGServiceCreateLDAPAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthLdap(ctx context.Context, body *WSGServiceCreateLDAPAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServicesAuthLdap, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddServicesAuthRadius
//
// Use this API command to create a new RADIUS authentication service.
//
// Operation ID: addServicesAuthRadius
// Operation path: /services/auth/radius
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGServiceCreateRadiusAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthRadius(ctx context.Context, body *WSGServiceCreateRadiusAuthentication, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServicesAuthRadius, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddServicesAuthTestById
//
// Use this API command to test an authentication service.
//
// Operation ID: addServicesAuthTestById
// Operation path: /services/auth/test/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceTestingConfig
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) AddServicesAuthTestById(ctx context.Context, body *WSGServiceTestingConfig, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServicesAuthTestById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuth
//
// Use this API command to delete a list of authentication service.
//
// Operation ID: deleteServicesAuth
// Operation path: /services/auth
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceDeleteBulkAuthenticationService
func (s *WSGAuthenticationServiceService) DeleteServicesAuth(ctx context.Context, body *WSGServiceDeleteBulkAuthenticationService, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuth, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuthAdById
//
// Use this API command to delete an active directory authentication service.
//
// Operation ID: deleteServicesAuthAdById
// Operation path: /services/auth/ad/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthAdById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuthAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuthById
//
// Use this API command to delete an authentication service.
//
// Operation ID: deleteServicesAuthById
// Operation path: /services/auth/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuthById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuthHlrById
//
// Use this API command to delete a hlr authentication service.
//
// Operation ID: deleteServicesAuthHlrById
// Operation path: /services/auth/hlr/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthHlrById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuthHlrById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuthLdapById
//
// Use this API command to delete a LDAP authentication service.
//
// Operation ID: deleteServicesAuthLdapById
// Operation path: /services/auth/ldap/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthLdapById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuthLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuthRadiusById
//
// Use this API command to delete a RADIUS authentication service.
//
// Operation ID: deleteServicesAuthRadiusById
// Operation path: /services/auth/radius/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuthRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuthRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS authentication service.
//
// Operation ID: deleteServicesAuthRadiusSecondaryById
// Operation path: /services/auth/radius/{id}/secondary
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusSecondaryById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuthRadiusSecondaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteServicesAuthRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS authentication service.
//
// Operation ID: deleteServicesAuthRadiusStandbyPrimaryById
// Operation path: /services/auth/radius/{id}/standbyPrimary
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusStandbyPrimaryById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServicesAuthRadiusStandbyPrimaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindServicesAaaServerAuthByQueryCriteria
//
// Query Non-Proxy Authentication AAAServers with specified filters.
//
// Operation ID: findServicesAaaServerAuthByQueryCriteria
// Operation path: /query/services/aaaServer/auth
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAaaServerAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAAAServerQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAAAServerQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAaaServerAuthByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAAServerQueryListAPIResponse), err
}

// FindServicesAuthAd
//
// Use this API command to retrieve a list of active directory authentication services.
//
// Operation ID: findServicesAuthAd
// Operation path: /services/auth/ad
// Success code: 200 (OK)
func (s *WSGAuthenticationServiceService) FindServicesAuthAd(ctx context.Context, mutators ...RequestMutator) (*WSGServiceActiveDirectoryServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceActiveDirectoryServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthAd, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceActiveDirectoryServiceListAPIResponse), err
}

// FindServicesAuthAdById
//
// Use this API command to retrieve an active directory authentication service.
//
// Operation ID: findServicesAuthAdById
// Operation path: /services/auth/ad/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthAdById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceActiveDirectoryServiceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceActiveDirectoryServiceAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceActiveDirectoryServiceAPIResponse), err
}

// FindServicesAuthAdByQueryCriteria
//
// Use this API command to retrieve a list of AD Authentication services by query criteria.
//
// Operation ID: findServicesAuthAdByQueryCriteria
// Operation path: /services/auth/ad/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthAdByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceActiveDirectoryServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceActiveDirectoryServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAuthAdByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceActiveDirectoryServiceListAPIResponse), err
}

// FindServicesAuthByQueryCriteria
//
// Use this API command to retrieve a list of Authentication services by query criteria.
//
// Operation ID: findServicesAuthByQueryCriteria
// Operation path: /services/auth/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceCommonAuthenticationServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceCommonAuthenticationServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAuthByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceCommonAuthenticationServiceListAPIResponse), err
}

// FindServicesAuthGuestById
//
// Use this API command to retrieve a Guest authentication service.
//
// Operation ID: findServicesAuthGuestById
// Operation path: /services/auth/guest/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthGuestById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceCommonAuthenticationServiceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceCommonAuthenticationServiceAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthGuestById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceCommonAuthenticationServiceAPIResponse), err
}

// FindServicesAuthHlr
//
// Use this API command to retrieve a list of hlr authentication services.
//
// Operation ID: findServicesAuthHlr
// Operation path: /services/auth/hlr
// Success code: 200 (OK)
func (s *WSGAuthenticationServiceService) FindServicesAuthHlr(ctx context.Context, mutators ...RequestMutator) (*WSGServiceHlrServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceHlrServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthHlr, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceHlrServiceListAPIResponse), err
}

// FindServicesAuthHlrById
//
// Use this API command to retrieve a hlr authentication service.
//
// Operation ID: findServicesAuthHlrById
// Operation path: /services/auth/hlr/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceHlrServiceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceHlrServiceAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthHlrById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceHlrServiceAPIResponse), err
}

// FindServicesAuthHlrByQueryCriteria
//
// Use this API command to retrieve a list of hlr Authentication services by query criteria.
//
// Operation ID: findServicesAuthHlrByQueryCriteria
// Operation path: /services/auth/hlr/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceHlrServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceHlrServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAuthHlrByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceHlrServiceListAPIResponse), err
}

// FindServicesAuthLdap
//
// Use this API command to retrieve a list of LDAP authentication services.
//
// Operation ID: findServicesAuthLdap
// Operation path: /services/auth/ldap
// Success code: 200 (OK)
func (s *WSGAuthenticationServiceService) FindServicesAuthLdap(ctx context.Context, mutators ...RequestMutator) (*WSGServiceLDAPServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceLDAPServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthLdap, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceLDAPServiceListAPIResponse), err
}

// FindServicesAuthLdapById
//
// Use this API command to retrieve a LDAP authentication service.
//
// Operation ID: findServicesAuthLdapById
// Operation path: /services/auth/ldap/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceLDAPServiceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceLDAPServiceAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceLDAPServiceAPIResponse), err
}

// FindServicesAuthLdapByQueryCriteria
//
// Use this API command to retrieve a list of LDAP Authentication services by query criteria.
//
// Operation ID: findServicesAuthLdapByQueryCriteria
// Operation path: /services/auth/ldap/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceLDAPServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceLDAPServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAuthLdapByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceLDAPServiceListAPIResponse), err
}

// FindServicesAuthLocalDbById
//
// Use this API command to retrieve a LocalDB authentication service.
//
// Operation ID: findServicesAuthLocal_dbById
// Operation path: /services/auth/local_db/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLocalDbById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceCommonAuthenticationServiceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceCommonAuthenticationServiceAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthLocalDbById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceCommonAuthenticationServiceAPIResponse), err
}

// FindServicesAuthRadius
//
// Use this API command to retrieve a list of RADIUS authentication services.
//
// Operation ID: findServicesAuthRadius
// Operation path: /services/auth/radius
// Success code: 200 (OK)
func (s *WSGAuthenticationServiceService) FindServicesAuthRadius(ctx context.Context, mutators ...RequestMutator) (*WSGServiceRadiusAuthenticationServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAuthenticationServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthRadius, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAuthenticationServiceListAPIResponse), err
}

// FindServicesAuthRadiusById
//
// Use this API command to retrieve a RADIUS authentication service.
//
// Operation ID: findServicesAuthRadiusById
// Operation path: /services/auth/radius/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGServiceRadiusAuthenticationServiceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAuthenticationServiceAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindServicesAuthRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAuthenticationServiceAPIResponse), err
}

// FindServicesAuthRadiusByQueryCriteria
//
// Use this API command to retrieve a list of radius Authentication services by query criteria.
//
// Operation ID: findServicesAuthRadiusByQueryCriteria
// Operation path: /services/auth/radius/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGServiceRadiusAuthenticationServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceRadiusAuthenticationServiceListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesAuthRadiusByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceRadiusAuthenticationServiceListAPIResponse), err
}

// PartialUpdateServicesAuthAdById
//
// Use this API command to modify the configuration of an active directory authentication service.
//
// Operation ID: partialUpdateServicesAuthAdById
// Operation path: /services/auth/ad/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceModifyActiveDirectoryAuthentication
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthAdById(ctx context.Context, body *WSGServiceModifyActiveDirectoryAuthentication, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesAuthAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateServicesAuthHlrById
//
// Use this API command to modify the configuration of a hlr authentication service.
//
// Operation ID: partialUpdateServicesAuthHlrById
// Operation path: /services/auth/hlr/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceModifyHlrAuthentication
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthHlrById(ctx context.Context, body *WSGServiceModifyHlrAuthentication, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesAuthHlrById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateServicesAuthLdapById
//
// Use this API command to modify the configuration of a LDAP authentication service.
//
// Operation ID: partialUpdateServicesAuthLdapById
// Operation path: /services/auth/ldap/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceModifyLDAPAuthentication
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLdapById(ctx context.Context, body *WSGServiceModifyLDAPAuthentication, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesAuthLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateServicesAuthLocalDbById
//
// Use this API command to update LocalDB authentication service.
//
// Operation ID: partialUpdateServicesAuthLocal_dbById
// Operation path: /services/auth/local_db/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceModifyLocalDbAuthentication
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLocalDbById(ctx context.Context, body *WSGServiceModifyLocalDbAuthentication, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesAuthLocalDbById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateServicesAuthRadiusById
//
// Use this API command to modify the configuration of a RADIUS authentication service.
//
// Operation ID: partialUpdateServicesAuthRadiusById
// Operation path: /services/auth/radius/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGServiceModifyRadiusAuthentication
//
// Required parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAuthentication, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateServicesAuthRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
