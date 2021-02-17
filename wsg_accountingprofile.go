package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGAccountingProfileService struct {
	apiClient *VSZClient
}

func NewWSGAccountingProfileService(c *VSZClient) *WSGAccountingProfileService {
	s := new(WSGAccountingProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountingProfileService() *WSGAccountingProfileService {
	return NewWSGAccountingProfileService(ss.apiClient)
}

// AddProfilesAcct
//
// Use this API command to create a new accounting profile.
//
// Operation ID: addProfilesAcct
// Operation path: /profiles/acct
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateAccountingProfile
func (s *WSGAccountingProfileService) AddProfilesAcct(ctx context.Context, body *WSGProfileCreateAccountingProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesAcct, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddProfilesAcctCloneById
//
// Use this API command to clone an accounting profile.
//
// Operation ID: addProfilesAcctCloneById
// Operation path: /profiles/acct/clone/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGProfileClone
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) AddProfilesAcctCloneById(ctx context.Context, body *WSGProfileClone, id string, mutators ...RequestMutator) (*WSGProfileCloneAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileCloneAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesAcctCloneById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileCloneAPIResponse), err
}

// DeleteProfilesAcct
//
// Use this API command to delete a list of accounting profile.
//
// Operation ID: deleteProfilesAcct
// Operation path: /profiles/acct
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileDeleteBulkAccountingProfile
func (s *WSGAccountingProfileService) DeleteProfilesAcct(ctx context.Context, body *WSGProfileDeleteBulkAccountingProfile, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesAcct, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesAcctById
//
// Use this API command to delete an accounting profile.
//
// Operation ID: deleteProfilesAcctById
// Operation path: /profiles/acct/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) DeleteProfilesAcctById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesAcctById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesAcct
//
// Use this API command to retrieve a list of accounting profiles.
//
// Operation ID: findProfilesAcct
// Operation path: /profiles/acct
// Success code: 200 (OK)
func (s *WSGAccountingProfileService) FindProfilesAcct(ctx context.Context, mutators ...RequestMutator) (*WSGProfileAccountingProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileAccountingProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesAcct, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileAccountingProfileListAPIResponse), err
}

// FindProfilesAcctById
//
// Use this API command to retrieve an accounting profile.
//
// Operation ID: findProfilesAcctById
// Operation path: /profiles/acct/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) FindProfilesAcctById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileAccountingProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileAccountingProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesAcctById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileAccountingProfileAPIResponse), err
}

// FindProfilesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting profiles by query criteria.
//
// Operation ID: findProfilesAcctByQueryCriteria
// Operation path: /profiles/acct/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGAccountingProfileService) FindProfilesAcctByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileAccountingProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileAccountingProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesAcctByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileAccountingProfileListAPIResponse), err
}

// PartialUpdateProfilesAcctById
//
// Use this API command to modify the configuration of an accounting profile.
//
// Operation ID: partialUpdateProfilesAcctById
// Operation path: /profiles/acct/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyAccountingProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGAccountingProfileService) PartialUpdateProfilesAcctById(ctx context.Context, body *WSGProfileModifyAccountingProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesAcctById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
