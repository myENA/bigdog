package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAuthenticationServiceService struct {
	apiClient *APIClient
}

func NewWSGAuthenticationServiceService(c *APIClient) *WSGAuthenticationServiceService {
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
// Request Body:
//	 - body *WSGServiceCreateActiveDirectoryAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthAd(ctx context.Context, body *WSGServiceCreateActiveDirectoryAuthentication) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthAd, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddServicesAuthHlr
//
// Use this API command to create a new hlr authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateHlrAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthHlr(ctx context.Context, body *WSGServiceCreateHlrAuthentication) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthHlr, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddServicesAuthLdap
//
// Use this API command to create a new LDAP authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateLDAPAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthLdap(ctx context.Context, body *WSGServiceCreateLDAPAuthentication) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthLdap, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddServicesAuthRadius
//
// Use this API command to create a new RADIUS authentication service.
//
// Request Body:
//	 - body *WSGServiceCreateRadiusAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthRadius(ctx context.Context, body *WSGServiceCreateRadiusAuthentication) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthRadius, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
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
func (s *WSGAuthenticationServiceService) AddServicesAuthTestById(ctx context.Context, body *WSGServiceTestingConfig, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServicesAuthTestById, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// DeleteServicesAuth
//
// Use this API command to delete a list of authentication service.
//
// Request Body:
//	 - body *WSGServiceDeleteBulkAuthenticationService
func (s *WSGAuthenticationServiceService) DeleteServicesAuth(ctx context.Context, body *WSGServiceDeleteBulkAuthenticationService) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuth, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteServicesAuthAdById
//
// Use this API command to delete an active directory authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthAdById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthAdById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteServicesAuthById
//
// Use this API command to delete an authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteServicesAuthHlrById
//
// Use this API command to delete a hlr authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthHlrById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthHlrById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteServicesAuthLdapById
//
// Use this API command to delete a LDAP authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthLdapById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthLdapById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteServicesAuthRadiusById
//
// Use this API command to delete a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthRadiusById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteServicesAuthRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusSecondaryById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthRadiusSecondaryById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteServicesAuthRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusStandbyPrimaryById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesAuthRadiusStandbyPrimaryById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindServicesAaaServerAuthByQueryCriteria
//
// Query Non-Proxy Authentication AAAServers with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAaaServerAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAAAServerQueryList, error) {
	var (
		req      *APIRequest
		resp     *WSGAAAServerQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAaaServerAuthByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAAAServerQueryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthAd
//
// Use this API command to retrieve a list of active directory authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthAd(ctx context.Context) (*WSGServiceActiveDirectoryServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceActiveDirectoryServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthAd, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceActiveDirectoryServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthAdById
//
// Use this API command to retrieve an active directory authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthAdById(ctx context.Context, id string) (*WSGServiceActiveDirectoryService, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceActiveDirectoryService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthAdById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceActiveDirectoryService()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthAdByQueryCriteria
//
// Use this API command to retrieve a list of AD Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthAdByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceActiveDirectoryServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceActiveDirectoryServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthAdByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceActiveDirectoryServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthByQueryCriteria
//
// Use this API command to retrieve a list of Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceCommonAuthenticationServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceCommonAuthenticationServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceCommonAuthenticationServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthGuestById
//
// Use this API command to retrieve a Guest authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthGuestById(ctx context.Context, id string) (*WSGServiceCommonAuthenticationService, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceCommonAuthenticationService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthGuestById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceCommonAuthenticationService()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthHlr
//
// Use this API command to retrieve a list of hlr authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthHlr(ctx context.Context) (*WSGServiceHlrServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceHlrServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthHlr, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceHlrServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthHlrById
//
// Use this API command to retrieve a hlr authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrById(ctx context.Context, id string) (*WSGServiceHlrService, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceHlrService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthHlrById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceHlrService()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthHlrByQueryCriteria
//
// Use this API command to retrieve a list of hlr Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceHlrServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceHlrServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthHlrByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceHlrServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthLdap
//
// Use this API command to retrieve a list of LDAP authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthLdap(ctx context.Context) (*WSGServiceLDAPServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceLDAPServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthLdap, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceLDAPServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthLdapById
//
// Use this API command to retrieve a LDAP authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapById(ctx context.Context, id string) (*WSGServiceLDAPService, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceLDAPService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthLdapById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceLDAPService()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthLdapByQueryCriteria
//
// Use this API command to retrieve a list of LDAP Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceLDAPServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceLDAPServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthLdapByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceLDAPServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthLocal_dbById
//
// Use this API command to retrieve a LocalDB authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLocal_dbById(ctx context.Context, id string) (*WSGServiceCommonAuthenticationService, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceCommonAuthenticationService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthLocal_dbById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceCommonAuthenticationService()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthRadius
//
// Use this API command to retrieve a list of RADIUS authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthRadius(ctx context.Context) (*WSGServiceRadiusAuthenticationServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceRadiusAuthenticationServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthRadius, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAuthenticationServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthRadiusById
//
// Use this API command to retrieve a RADIUS authentication service.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusById(ctx context.Context, id string) (*WSGServiceRadiusAuthenticationService, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceRadiusAuthenticationService
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesAuthRadiusById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAuthenticationService()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindServicesAuthRadiusByQueryCriteria
//
// Use this API command to retrieve a list of radius Authentication services by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGServiceRadiusAuthenticationServiceList, error) {
	var (
		req      *APIRequest
		resp     *WSGServiceRadiusAuthenticationServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesAuthRadiusByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceRadiusAuthenticationServiceList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthAdById(ctx context.Context, body *WSGServiceModifyActiveDirectoryAuthentication, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthAdById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthHlrById(ctx context.Context, body *WSGServiceModifyHlrAuthentication, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthHlrById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLdapById(ctx context.Context, body *WSGServiceModifyLDAPAuthentication, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthLdapById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLocal_dbById(ctx context.Context, body *WSGServiceModifyLocalDbAuthentication, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthLocal_dbById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
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
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthRadiusById(ctx context.Context, body *WSGServiceModifyRadiusAuthentication, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesAuthRadiusById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
