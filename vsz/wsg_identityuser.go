package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGIdentityUserService struct {
	apiClient *APIClient
}

func NewWSGIdentityUserService(c *APIClient) *WSGIdentityUserService {
	s := new(WSGIdentityUserService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityUserService() *WSGIdentityUserService {
	return NewWSGIdentityUserService(ss.apiClient)
}

// AddIdentityUserList
//
// Use this API command to retrieve a list of identity user.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityUserList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityUserList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUserList, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddIdentityUsers
//
// Use this API command to create identity user.
//
// Request Body:
//	 - body *WSGIdentityCreateUser
func (s *WSGIdentityUserService) AddIdentityUsers(ctx context.Context, body *WSGIdentityCreateUser) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddIdentityUsers, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteIdentityUsers
//
// Use this API command to delete multiple identity users.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserService) DeleteIdentityUsers(ctx context.Context, body *WSGIdentityDeleteBulk) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUsers, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
}

// DeleteIdentityUsersById
//
// Use this API command to delete identity user.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) DeleteIdentityUsersById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteIdentityUsersById, true)
}

// FindIdentityUsers
//
// Use this API command to retrieve a list of identity user.
//
// Optional Parameters:
// - createdOnFrom string
//		- nullable
// - createdOnTo string
//		- nullable
// - displayName string
//		- nullable
// - email string
//		- nullable
// - firstName string
//		- nullable
// - index string
//		- nullable
// - isDisabled string
//		- nullable
// - lastName string
//		- nullable
// - listSize string
//		- nullable
// - phone string
//		- nullable
// - timeZone string
//		- nullable
// - userName string
//		- nullable
// - userSource string
//		- nullable
// - userType string
//		- nullable
func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, optionalParams map[string]interface{}) (*WSGIdentityUserList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityUserList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsers, true)
}

// FindIdentityUsersAaaserver
//
// Use this API command to retrieve a list of aaa server.
func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context) (*WSGIdentityAaaServerList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityAaaServerList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersAaaserver, true)
}

// FindIdentityUsersById
//
// Use this API command to retrieve identity user.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, id string) (*WSGIdentityUserConfiguration, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityUserConfiguration
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersById, true)
}

// FindIdentityUsersCountries
//
// Use this API command to retrieve a list of countries.
func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context) (*WSGIdentityCountryList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityCountryList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersCountries, true)
}

// FindIdentityUsersPackages
//
// Use this API command to retrieve a list of packages.
func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context) (*WSGIdentityPackageList, error) {
	var (
		req  *APIRequest
		resp *WSGIdentityPackageList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindIdentityUsersPackages, true)
}

// PartialUpdateIdentityUsersById
//
// Use this API command to modify the basic information of identity user.
//
// Request Body:
//	 - body *WSGIdentityModifyUser
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) PartialUpdateIdentityUsersById(ctx context.Context, body *WSGIdentityModifyUser, id string) (*WSGCommonEmptyResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateIdentityUsersById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
