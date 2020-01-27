package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGZoneAAAService struct {
	apiClient *APIClient
}

func NewWSGZoneAAAService(c *APIClient) *WSGZoneAAAService {
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
// Request Body:
//	 - body *WSGAAACreateActiveDirectoryServer
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaAdByZoneId(ctx context.Context, body *WSGAAACreateActiveDirectoryServer, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesAaaAdByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
}

// AddRkszonesAaaLdapByZoneId
//
// Use this API command to create a new LDAP server of a zone.
//
// Request Body:
//	 - body *WSGAAACreateLDAPServer
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaLdapByZoneId(ctx context.Context, body *WSGAAACreateLDAPServer, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesAaaLdapByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
}

// AddRkszonesAaaRadiusByZoneId
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
func (s *WSGZoneAAAService) AddRkszonesAaaRadiusByZoneId(ctx context.Context, body *WSGAAACreateAuthenticationServer, zoneId string, optionalParams map[string]interface{}) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesAaaRadiusByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["forAccounting"]; ok {
		req.AddQueryParameter("forAccounting", v)
	}
}

// DeleteRkszonesAaaAdById
//
// Use this API command to delete an active directory server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaAdById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAaaAdById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// DeleteRkszonesAaaById
//
// Use this API command to delete an AAA server.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAaaById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// DeleteRkszonesAaaByZoneId
//
// Use this API command to delete a list of AAA server.
//
// Request Body:
//	 - body *WSGAAADeleteBulkAAAServerList
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaByZoneId(ctx context.Context, body *WSGAAADeleteBulkAAAServerList, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAaaByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("zoneId", zoneId)
}

// DeleteRkszonesAaaLdapById
//
// Use this API command to delete a LDAP server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaLdapById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAaaLdapById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// DeleteRkszonesAaaRadiusById
//
// Use this API command to delete a radius server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// DeleteRkszonesAaaRadiusSecondaryById
//
// Use this API command to disable secondary server on radius server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusSecondaryById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusSecondaryById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// DeleteRkszonesAaaRadiusStandbyPrimaryById
//
// Use this API command to disable primary server on radius server of a zone for Standby Cluster.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusStandbyPrimaryById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAaaRadiusStandbyPrimaryById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// FindRkszonesAaaAdById
//
// Use this API command to retrieve an active directory server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdById(ctx context.Context, id string, zoneId string) (*WSGAAAActiveDirectory, error) {
	var (
		req  *APIRequest
		resp *WSGAAAActiveDirectory
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAaaAdById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// FindRkszonesAaaAdByZoneId
//
// Use this API command to retrieve a list of active directory servers of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdByZoneId(ctx context.Context, zoneId string) (*WSGAAAActiveDirectoryList, error) {
	var (
		req  *APIRequest
		resp *WSGAAAActiveDirectoryList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAaaAdByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
}

// FindRkszonesAaaLdapById
//
// Use this API command to retrieve a LDAP server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapById(ctx context.Context, id string, zoneId string) (*WSGAAALDAPServer, error) {
	var (
		req  *APIRequest
		resp *WSGAAALDAPServer
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAaaLdapById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// FindRkszonesAaaLdapByZoneId
//
// Use this API command to retrieve a list of LDAP servers of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapByZoneId(ctx context.Context, zoneId string) (*WSGAAALDAPServerList, error) {
	var (
		req  *APIRequest
		resp *WSGAAALDAPServerList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAaaLdapByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
}

// FindRkszonesAaaRadiusById
//
// Use this API command to retrieve a radius server of a zone.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusById(ctx context.Context, id string, zoneId string) (*WSGAAAAuthenticationServer, error) {
	var (
		req  *APIRequest
		resp *WSGAAAAuthenticationServer
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAaaRadiusById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// FindRkszonesAaaRadiusByZoneId
//
// Use this API command to retrieve a list of radius servers of a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusByZoneId(ctx context.Context, zoneId string) (*WSGAAAAuthenticationServerList, error) {
	var (
		req  *APIRequest
		resp *WSGAAAAuthenticationServerList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAaaRadiusByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
}

// PartialUpdateRkszonesAaaAdById
//
// Use this API command to modify the basic information on active directory server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaAdById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// PartialUpdateRkszonesAaaLdapById
//
// Use this API command to modify the basic information on LDAP server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyLDAPServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaLdapById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// PartialUpdateRkszonesAaaRadiusById
//
// Use this API command to modify the basic information on radius server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesAaaRadiusById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// UpdateRkszonesAaaAdById
//
// Use this API command to modify the basic information on active directory server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesAaaAdById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// UpdateRkszonesAaaLdapById
//
// Use this API command to modify the basic information on LDAP server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyLDAPServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesAaaLdapById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}

// UpdateRkszonesAaaRadiusById
//
// Use this API command to modify the basic information on radius server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesAaaRadiusById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
}
