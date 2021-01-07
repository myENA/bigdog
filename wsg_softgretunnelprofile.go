package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGSoftGRETunnelProfileService struct {
	apiClient *VSZClient
}

func NewWSGSoftGRETunnelProfileService(c *VSZClient) *WSGSoftGRETunnelProfileService {
	s := new(WSGSoftGRETunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSoftGRETunnelProfileService() *WSGSoftGRETunnelProfileService {
	return NewWSGSoftGRETunnelProfileService(ss.apiClient)
}

// AddProfilesTunnelSoftgre
//
// Operation ID: addProfilesTunnelSoftgre
//
// Use this API command to create SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileCreateSoftGREProfile
func (s *WSGSoftGRETunnelProfileService) AddProfilesTunnelSoftgre(ctx context.Context, body *WSGProfileCreateSoftGREProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddProfilesTunnelSoftgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteProfilesTunnelSoftgre
//
// Operation ID: deleteProfilesTunnelSoftgre
//
// Use this API command to delete multiple SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgre(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesTunnelSoftgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesTunnelSoftgreById
//
// Operation ID: deleteProfilesTunnelSoftgreById
//
// Use this API command to delete SoftGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) DeleteProfilesTunnelSoftgreById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesTunnelSoftgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesTunnelSoftgre
//
// Operation ID: findProfilesTunnelSoftgre
//
// Use this API command to retrieve a list of SoftGRE tunnel profile.
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgre(ctx context.Context, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesTunnelSoftgre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesTunnelSoftgreById
//
// Operation ID: findProfilesTunnelSoftgreById
//
// Use this API command to retrieve SoftGRE tunnel profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileSoftGREProfileAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileSoftGREProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesTunnelSoftgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileSoftGREProfile()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesTunnelSoftgreByQueryCriteria
//
// Operation ID: findProfilesTunnelSoftgreByQueryCriteria
//
// Use this API command to query a list of SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSoftGRETunnelProfileService) FindProfilesTunnelSoftgreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileSoftGREProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileSoftGREProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindProfilesTunnelSoftgreByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileSoftGREProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesTunnelSoftgreById
//
// Operation ID: partialUpdateProfilesTunnelSoftgreById
//
// Use this API command to modify the configuration of SoftGRE tunnel profile.
//
// Request Body:
//	 - body *WSGProfileModifySoftGREProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSoftGRETunnelProfileService) PartialUpdateProfilesTunnelSoftgreById(ctx context.Context, body *WSGProfileModifySoftGREProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelSoftgreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
