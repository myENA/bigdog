package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGAccountSecurityService struct {
	apiClient *APIClient
}

func NewWSGAccountSecurityService(c *APIClient) *WSGAccountSecurityService {
	s := new(WSGAccountSecurityService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountSecurityService() *WSGAccountSecurityService {
	return NewWSGAccountSecurityService(ss.apiClient)
}

// AddAccountSecurity
//
// Use this API command to create the account security proile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileCreate
func (s *WSGAccountSecurityService) AddAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileCreate) (*WSGCommonCreateResultIdName, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResultIdName
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAccountSecurity, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteAccountSecurity
//
// Use this API command to selete the account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDeleteList
func (s *WSGAccountSecurityService) DeleteAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileDeleteList) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAccountSecurity, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteAccountSecurityById
//
// Use this API command to delete the account security profile by id.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) DeleteAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileDelete, id string) (*WSGCommonCreateResultIdName, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResultIdName
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// FindAccountSecurity
//
// Use this API command to get account security profiles.
func (s *WSGAccountSecurityService) FindAccountSecurity(ctx context.Context) (*WSGAccountSecurityProfileProfileListResult, error) {
	var (
		req  *APIRequest
		resp *WSGAccountSecurityProfileProfileListResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAccountSecurity, true)
}

// FindAccountSecurityById
//
// Use this API command to retrieve the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileGetById
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) FindAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileGetById, id string) (*WSGAccountSecurityProfileGetByIdResult, error) {
	var (
		req  *APIRequest
		resp *WSGAccountSecurityProfileGetByIdResult
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) PartialUpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, id string) (*WSGCommonCreateResultIdName, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResultIdName
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) UpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, id string) (*WSGCommonCreateResultIdName, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResultIdName
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
