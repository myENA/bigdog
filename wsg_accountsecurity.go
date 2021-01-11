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
// Use this API command to create the account security proile.
//
// Operation ID: addAccountSecurity
// Operation path: /accountSecurity
// Success code: 200 (OK)
//
// Request body:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddAccountSecurity, true)
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
// Use this API command to selete the account security profile.
//
// Operation ID: deleteAccountSecurity
// Operation path: /accountSecurity
// Success code: 200 (OK)
//
// Request body:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteAccountSecurity, true)
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
// Use this API command to delete the account security profile by id.
//
// Operation ID: deleteAccountSecurityById
// Operation path: /accountSecurity/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAccountSecurityProfileDelete
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteAccountSecurityById, true)
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
// Use this API command to get account security profiles.
//
// Operation ID: findAccountSecurity
// Operation path: /accountSecurity
// Success code: 200 (OK)
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindAccountSecurity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAccountSecurityProfileProfileListResultAPIResponse), err
}

// FindAccountSecurityById
//
// Use this API command to retrieve the specific account security profile.
//
// Operation ID: findAccountSecurityById
// Operation path: /accountSecurity/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAccountSecurityProfileGetById
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindAccountSecurityById, true)
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
// Use this API command to modify the specific account security profile.
//
// Operation ID: partialUpdateAccountSecurityById
// Operation path: /accountSecurity/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateAccountSecurityById, true)
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
// Use this API command to modify the specific account security profile.
//
// Operation ID: updateAccountSecurityById
// Operation path: /accountSecurity/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateAccountSecurityById, true)
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
