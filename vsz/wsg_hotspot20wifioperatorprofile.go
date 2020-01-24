package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGHotspot20WiFiOperatorProfileService struct {
	apiClient *APIClient
}

func NewWSGHotspot20WiFiOperatorProfileService(c *APIClient) *WSGHotspot20WiFiOperatorProfileService {
	s := new(WSGHotspot20WiFiOperatorProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20WiFiOperatorProfileService() *WSGHotspot20WiFiOperatorProfileService {
	return NewWSGHotspot20WiFiOperatorProfileService(ss.apiClient)
}

// AddProfilesHs20Operators
//
// Use this API command to create a new Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
func (s *WSGHotspot20WiFiOperatorProfileService) AddProfilesHs20Operators(ctx context.Context, body *WSGProfileHs20Operator) (*WSGCommonCreateResult, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesHs20Operators, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesHs20Operators
//
// Use this API command to delete multiple Hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20Operators(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20Operators, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteProfilesHs20OperatorsById
//
// Use this API command to delete a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20OperatorsById, true)
}

// DeleteProfilesHs20OperatorsCertificateById
//
// Use this API command to disable certificate of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsCertificateById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20OperatorsCertificateById, true)
}

// FindProfilesHs20Operators
//
// Use this API command to retrieve list of Hotspot 2.0 Wi-Fi Operators.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20Operators(ctx context.Context, optionalParams map[string]interface{}) (*WSGProfileHs20OperatorList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileHs20OperatorList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20Operators, true)
}

// FindProfilesHs20OperatorsById
//
// Use this API command to retrieve a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsById(ctx context.Context, id string) (*WSGProfileHs20Operator, error) {
	var (
		req  *APIRequest
		resp *WSGProfileHs20Operator
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20OperatorsById, true)
}

// FindProfilesHs20OperatorsByQueryCriteria
//
// Query hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileHs20OperatorList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileHs20OperatorList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesHs20OperatorsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateProfilesHs20OperatorsById
//
// Use this API command to modify the basic information of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileModifyHS20Operator
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) PartialUpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileModifyHS20Operator, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesHs20OperatorsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateProfilesHs20OperatorsById
//
// Use this API command to modify entire configuration of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) UpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileHs20Operator, id string) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateProfilesHs20OperatorsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
