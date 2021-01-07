package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGAccountSecurityService struct {
	apiClient *VSZClient
}

func NewWSGAccountSecurityService(c *VSZClient) *WSGAccountSecurityService {
	s := new(WSGAccountSecurityService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountSecurityService() *WSGAccountSecurityService {
	return NewWSGAccountSecurityService(ss.apiClient)
}

// AddAccountSecurity
//
// Operation ID: addAccountSecurity
//
// Use this API command to create the account security proile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileCreate
func (s *WSGAccountSecurityService) AddAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileCreate, mutators ...RequestMutator) (*WSGCommonCreateResultIdNameAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultIdNameAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddAccountSecurity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
}

// DeleteAccountSecurity
//
// Operation ID: deleteAccountSecurity
//
// Use this API command to selete the account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDeleteList
func (s *WSGAccountSecurityService) DeleteAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileDeleteList, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAccountSecurity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteAccountSecurityById
//
// Operation ID: deleteAccountSecurityById
//
// Use this API command to delete the account security profile by id.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) DeleteAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileDelete, id string, mutators ...RequestMutator) (*WSGCommonCreateResultIdNameAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultIdNameAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAccountSecurityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
}

// FindAccountSecurity
//
// Operation ID: findAccountSecurity
//
// Use this API command to get account security profiles.
func (s *WSGAccountSecurityService) FindAccountSecurity(ctx context.Context, mutators ...RequestMutator) (*WSGAccountSecurityProfileProfileListResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAccountSecurityProfileProfileListResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAccountSecurityProfileProfileListResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAccountSecurity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAccountSecurityProfileProfileListResultAPIResponse), err
}

// FindAccountSecurityById
//
// Operation ID: findAccountSecurityById
//
// Use this API command to retrieve the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileGetById
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) FindAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileGetById, id string, mutators ...RequestMutator) (*WSGAccountSecurityProfileGetByIdResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAccountSecurityProfileGetByIdResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAccountSecurityProfileGetByIdResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAccountSecurityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAccountSecurityProfileGetByIdResultAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAccountSecurityProfileGetByIdResultAPIResponse), err
}

// PartialUpdateAccountSecurityById
//
// Operation ID: partialUpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) PartialUpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, id string, mutators ...RequestMutator) (*WSGCommonCreateResultIdNameAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultIdNameAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateAccountSecurityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
}

// UpdateAccountSecurityById
//
// Operation ID: updateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountSecurityService) UpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, id string, mutators ...RequestMutator) (*WSGCommonCreateResultIdNameAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultIdNameAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateAccountSecurityById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultIdNameAPIResponse), err
}
