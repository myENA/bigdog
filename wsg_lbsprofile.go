package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGLBSprofileService struct {
	apiClient *VSZClient
}

func NewWSGLBSprofileService(c *VSZClient) *WSGLBSprofileService {
	s := new(WSGLBSprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGLBSprofileService() *WSGLBSprofileService {
	return NewWSGLBSprofileService(ss.apiClient)
}

// AddProfilesLbs
//
// Operation ID: addProfilesLbs
//
// Create LBS profile.
//
// Request Body:
//	 - body *WSGProfileLbsProfile
func (s *WSGLBSprofileService) AddProfilesLbs(ctx context.Context, body *WSGProfileLbsProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddProfilesLbs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesLbs
//
// Operation ID: deleteProfilesLbs
//
// Delete multiple LBS profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGLBSprofileService) DeleteProfilesLbs(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesLbs, true)
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

// DeleteProfilesLbsById
//
// Operation ID: deleteProfilesLbsById
//
// Delete LBS profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGLBSprofileService) DeleteProfilesLbsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesLbsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindProfilesLbsById
//
// Operation ID: findProfilesLbsById
//
// Retrieve LBS profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGLBSprofileService) FindProfilesLbsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileLbsProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileLbsProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileLbsProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesLbsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileLbsProfileAPIResponse), err
}

// FindProfilesLbsByQueryCriteria
//
// Operation ID: findProfilesLbsByQueryCriteria
//
// Query LBS profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGLBSprofileService) FindProfilesLbsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileLbsProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileLbsProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileLbsProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindProfilesLbsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileLbsProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileLbsProfileListAPIResponse), err
}

// PartialUpdateProfilesLbsById
//
// Operation ID: partialUpdateProfilesLbsById
//
// Update LBS profile.
//
// Request Body:
//	 - body *WSGProfileLbsProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGLBSprofileService) PartialUpdateProfilesLbsById(ctx context.Context, body *WSGProfileLbsProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesLbsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
