package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addRkszonesAaaAdByZoneId
//
// Use this API command to create a new active directory server of a zone.
//
// Request Body:
//	 - body *WSGAAACreateActiveDirectoryServer
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaAdByZoneId(ctx context.Context, body *WSGAAACreateActiveDirectoryServer, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesAaaAdByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesAaaLdapByZoneId
//
// Operation ID: addRkszonesAaaLdapByZoneId
//
// Use this API command to create a new LDAP server of a zone.
//
// Request Body:
//	 - body *WSGAAACreateLDAPServer
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaLdapByZoneId(ctx context.Context, body *WSGAAACreateLDAPServer, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesAaaLdapByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddRkszonesAaaRadiusByZoneId
//
// Operation ID: addRkszonesAaaRadiusByZoneId
//
// Use this API command to create a new radius server of a zone.
//
// Request Body:
//	 - body *WSGAAACreateAuthenticationServer
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - forAccounting string
//		- nullable
func (s *WSGZoneAAAService) AddRkszonesAaaRadiusByZoneId(ctx context.Context, body *WSGAAACreateAuthenticationServer, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesAaaRadiusByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["forAccounting"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("forAccounting", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesAaaAdById
//
// Operation ID: deleteRkszonesAaaAdById
//
// Use this API command to delete an active directory server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaAdById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesAaaById
//
// Operation ID: deleteRkszonesAaaById
//
// Use this API command to delete an AAA server.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesAaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesAaaByZoneId
//
// Operation ID: deleteRkszonesAaaByZoneId
//
// Use this API command to delete a list of AAA server.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesAaaByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesAaaLdapById
//
// Operation ID: deleteRkszonesAaaLdapById
//
// Use this API command to delete a LDAP server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaLdapById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesAaaRadiusById
//
// Operation ID: deleteRkszonesAaaRadiusById
//
// Use this API command to delete a radius server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesAaaRadiusSecondaryById
//
// Operation ID: deleteRkszonesAaaRadiusSecondaryById
//
// Use this API command to disable secondary server on radius server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusSecondaryById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusSecondaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesAaaRadiusStandbyPrimaryById
//
// Operation ID: deleteRkszonesAaaRadiusStandbyPrimaryById
//
// Use this API command to disable primary server on radius server of a zone for Standby Cluster.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusStandbyPrimaryById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusStandbyPrimaryById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesAaaAdById
//
// Operation ID: findRkszonesAaaAdById
//
// Use this API command to retrieve an active directory server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGAAAActiveDirectoryAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAAActiveDirectory
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAAActiveDirectory()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesAaaAdByZoneId
//
// Operation ID: findRkszonesAaaAdByZoneId
//
// Use this API command to retrieve a list of active directory servers of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGAAAActiveDirectoryListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAAActiveDirectoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesAaaAdByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAAActiveDirectoryList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesAaaLdapById
//
// Operation ID: findRkszonesAaaLdapById
//
// Use this API command to retrieve a LDAP server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGAAALDAPServerAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAALDAPServer
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAALDAPServer()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesAaaLdapByZoneId
//
// Operation ID: findRkszonesAaaLdapByZoneId
//
// Use this API command to retrieve a list of LDAP servers of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGAAALDAPServerListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAALDAPServerList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesAaaLdapByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAALDAPServerList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesAaaRadiusById
//
// Operation ID: findRkszonesAaaRadiusById
//
// Use this API command to retrieve a radius server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGAAAAuthenticationServerAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAAAuthenticationServer
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAAAuthenticationServer()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesAaaRadiusByZoneId
//
// Operation ID: findRkszonesAaaRadiusByZoneId
//
// Use this API command to retrieve a list of radius servers of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - forAccounting string
//		- nullable
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAAAAuthenticationServerListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAAAAuthenticationServerList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesAaaRadiusByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["forAccounting"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("forAccounting", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAAAAuthenticationServerList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesAaaAdById
//
// Operation ID: partialUpdateRkszonesAaaAdById
//
// Use this API command to modify the configuration on active directory server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateRkszonesAaaLdapById
//
// Operation ID: partialUpdateRkszonesAaaLdapById
//
// Use this API command to modify the configuration on LDAP server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyLDAPServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateRkszonesAaaRadiusById
//
// Operation ID: partialUpdateRkszonesAaaRadiusById
//
// Use this API command to modify the configuration on radius server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesAaaAdById
//
// Operation ID: updateRkszonesAaaAdById
//
// Use this API command to modify the configuration on active directory server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesAaaAdById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesAaaLdapById
//
// Operation ID: updateRkszonesAaaLdapById
//
// Use this API command to modify the configuration on LDAP server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyLDAPServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesAaaLdapById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateRkszonesAaaRadiusById
//
// Operation ID: updateRkszonesAaaRadiusById
//
// Use this API command to modify the configuration on radius server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesAaaRadiusById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
