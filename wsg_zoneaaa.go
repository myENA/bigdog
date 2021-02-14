package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGZoneAAAService struct {
	apiClient *VSZClient
}

func NewWSGZoneAAAService(c *VSZClient) *WSGZoneAAAService {
	s := new(WSGZoneAAAService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZoneAAAService() *WSGZoneAAAService {
	return NewWSGZoneAAAService(ss.apiClient)
}

// AddRkszonesAaaAdByZoneId
//
// Use this API command to create a new active directory server of a zone.
//
// Operation ID: addRkszonesAaaAdByZoneId
// Operation path: /rkszones/{zoneId}/aaa/ad
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAAACreateActiveDirectoryServer
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaAdByZoneId(ctx context.Context, body *WSGAAACreateActiveDirectoryServer, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesAaaAdByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesAaaLdapByZoneId
//
// Use this API command to create a new LDAP server of a zone.
//
// Operation ID: addRkszonesAaaLdapByZoneId
// Operation path: /rkszones/{zoneId}/aaa/ldap
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAAACreateLDAPServer
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaLdapByZoneId(ctx context.Context, body *WSGAAACreateLDAPServer, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesAaaLdapByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRkszonesAaaRadiusByZoneId
//
// Use this API command to create a new radius server of a zone.
//
// Operation ID: addRkszonesAaaRadiusByZoneId
// Operation path: /rkszones/{zoneId}/aaa/radius
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGAAACreateAuthenticationServer
//
// Required parameters:
// - zoneId string
//		- required
//
// Optional parameters:
// - forAccounting string
//		- nullable
func (s *WSGZoneAAAService) AddRkszonesAaaRadiusByZoneId(ctx context.Context, body *WSGAAACreateAuthenticationServer, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesAaaRadiusByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["forAccounting"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("forAccounting", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesAaaAdById
//
// Use this API command to delete an active directory server of a zone.
//
// Operation ID: deleteRkszonesAaaAdById
// Operation path: /rkszones/{zoneId}/aaa/ad/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaAdById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesAaaById
//
// Use this API command to delete an AAA server.
//
// Operation ID: deleteRkszonesAaaById
// Operation path: /rkszones/{zoneId}/aaa/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesAaaByZoneId
//
// Use this API command to delete a list of AAA server.
//
// Operation ID: deleteRkszonesAaaByZoneId
// Operation path: /rkszones/{zoneId}/aaa
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAaaByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesAaaLdapById
//
// Use this API command to delete a LDAP server of a zone.
//
// Operation ID: deleteRkszonesAaaLdapById
// Operation path: /rkszones/{zoneId}/aaa/ldap/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaLdapById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesAaaRadiusById
//
// Use this API command to delete a radius server of a zone.
//
// Operation ID: deleteRkszonesAaaRadiusById
// Operation path: /rkszones/{zoneId}/aaa/radius/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesAaaRadiusSecondaryById
//
// Use this API command to disable secondary server on radius server of a zone.
//
// Operation ID: deleteRkszonesAaaRadiusSecondaryById
// Operation path: /rkszones/{zoneId}/aaa/radius/{id}/secondary
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusSecondaryById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusSecondaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesAaaRadiusStandbyPrimaryById
//
// Use this API command to disable primary server on radius server of a zone for Standby Cluster.
//
// Operation ID: deleteRkszonesAaaRadiusStandbyPrimaryById
// Operation path: /rkszones/{zoneId}/aaa/radius/{id}/standbyPrimary
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusStandbyPrimaryById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusStandbyPrimaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesAaaAdById
//
// Use this API command to retrieve an active directory server of a zone.
//
// Operation ID: findRkszonesAaaAdById
// Operation path: /rkszones/{zoneId}/aaa/ad/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGAAAActiveDirectoryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAAAActiveDirectoryAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAAActiveDirectoryAPIResponse), err
}

// FindRkszonesAaaAdByZoneId
//
// Use this API command to retrieve a list of active directory servers of a zone.
//
// Operation ID: findRkszonesAaaAdByZoneId
// Operation path: /rkszones/{zoneId}/aaa/ad
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGAAAActiveDirectoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAAAActiveDirectoryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAaaAdByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAAActiveDirectoryListAPIResponse), err
}

// FindRkszonesAaaLdapById
//
// Use this API command to retrieve a LDAP server of a zone.
//
// Operation ID: findRkszonesAaaLdapById
// Operation path: /rkszones/{zoneId}/aaa/ldap/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGAAALDAPServerAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAAALDAPServerAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAALDAPServerAPIResponse), err
}

// FindRkszonesAaaLdapByZoneId
//
// Use this API command to retrieve a list of LDAP servers of a zone.
//
// Operation ID: findRkszonesAaaLdapByZoneId
// Operation path: /rkszones/{zoneId}/aaa/ldap
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGAAALDAPServerListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAAALDAPServerListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAaaLdapByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAALDAPServerListAPIResponse), err
}

// FindRkszonesAaaRadiusById
//
// Use this API command to retrieve a radius server of a zone.
//
// Operation ID: findRkszonesAaaRadiusById
// Operation path: /rkszones/{zoneId}/aaa/radius/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGAAAAuthenticationServerAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAAAAuthenticationServerAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAAAuthenticationServerAPIResponse), err
}

// FindRkszonesAaaRadiusByZoneId
//
// Use this API command to retrieve a list of radius servers of a zone.
//
// Operation ID: findRkszonesAaaRadiusByZoneId
// Operation path: /rkszones/{zoneId}/aaa/radius
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
//
// Optional parameters:
// - forAccounting string
//		- nullable
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAAAAuthenticationServerListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGAAAAuthenticationServerListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesAaaRadiusByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["forAccounting"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("forAccounting", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAAAAuthenticationServerListAPIResponse), err
}

// PartialUpdateRkszonesAaaAdById
//
// Use this API command to modify the configuration on active directory server of a zone.
//
// Operation ID: partialUpdateRkszonesAaaAdById
// Operation path: /rkszones/{zoneId}/aaa/ad/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateRkszonesAaaLdapById
//
// Use this API command to modify the configuration on LDAP server of a zone.
//
// Operation ID: partialUpdateRkszonesAaaLdapById
// Operation path: /rkszones/{zoneId}/aaa/ldap/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAAAModifyLDAPServer
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// PartialUpdateRkszonesAaaRadiusById
//
// Use this API command to modify the configuration on radius server of a zone.
//
// Operation ID: partialUpdateRkszonesAaaRadiusById
// Operation path: /rkszones/{zoneId}/aaa/radius/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesAaaAdById
//
// Use this API command to modify the configuration on active directory server of a zone by complete attributes.
//
// Operation ID: updateRkszonesAaaAdById
// Operation path: /rkszones/{zoneId}/aaa/ad/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesAaaLdapById
//
// Use this API command to modify the configuration on LDAP server of a zone by complete attributes.
//
// Operation ID: updateRkszonesAaaLdapById
// Operation path: /rkszones/{zoneId}/aaa/ldap/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAAAModifyLDAPServer
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesAaaRadiusById
//
// Use this API command to modify the configuration on radius server of a zone by complete attributes.
//
// Operation ID: updateRkszonesAaaRadiusById
// Operation path: /rkszones/{zoneId}/aaa/radius/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
