package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGL2oGREService struct {
	apiClient *VSZClient
}

func NewWSGL2oGREService(c *VSZClient) *WSGL2oGREService {
	s := new(WSGL2oGREService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2oGREService() *WSGL2oGREService {
	return NewWSGL2oGREService(ss.apiClient)
}

// AddProfilesL2ogre
//
// Use this API command to create L2oGRE profile.
//
// Request Body:
//	 - body *WSGProfileCreateL2oGREProfile
func (s *WSGL2oGREService) AddProfilesL2ogre(ctx context.Context, body *WSGProfileCreateL2oGREProfile, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesL2ogre, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteProfilesL2ogre
//
// Use this API command to delete multiple L2oGRE profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL2oGREService) DeleteProfilesL2ogre(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesL2ogre, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesL2ogreById
//
// Use this API command to delete L2oGRE profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2oGREService) DeleteProfilesL2ogreById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesL2ogreById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesL2ogre
//
// Use this API command to retrieve a list of L2oGRE profile.
func (s *WSGL2oGREService) FindProfilesL2ogre(ctx context.Context, mutators ...RequestMutator) (*WSGProfileList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesL2ogre, true)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesL2ogreById
//
// Use this API command to retrieve L2oGRE profile by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2oGREService) FindProfilesL2ogreById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileL2oGREProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileL2oGREProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesL2ogreById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileL2oGREProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesL2ogreByQueryCriteria
//
// Use this API command to query a list of L2oGRE profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2oGREService) FindProfilesL2ogreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileL2oGREProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileL2oGREProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesL2ogreByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileL2oGREProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesL2ogreById
//
// Use this API command to modify the configuration of L2oGRE profile.
//
// Request Body:
//	 - body *WSGProfileModifyL2oGREProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGL2oGREService) PartialUpdateProfilesL2ogreById(ctx context.Context, body *WSGProfileModifyL2oGREProfile, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesL2ogreById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
