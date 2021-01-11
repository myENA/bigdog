package bigdog

// API Version: v9_1

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
// Operation ID: addProfilesL2ogre
// Operation path: /profiles/l2ogre
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateL2oGREProfile
func (s *WSGL2oGREService) AddProfilesL2ogre(ctx context.Context, body *WSGProfileCreateL2oGREProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddProfilesL2ogre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteProfilesL2ogre
//
// Use this API command to delete multiple L2oGRE profile.
//
// Operation ID: deleteProfilesL2ogre
// Operation path: /profiles/l2ogre
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL2oGREService) DeleteProfilesL2ogre(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesL2ogre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteProfilesL2ogreById
//
// Use this API command to delete L2oGRE profile.
//
// Operation ID: deleteProfilesL2ogreById
// Operation path: /profiles/l2ogre/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGL2oGREService) DeleteProfilesL2ogreById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteProfilesL2ogreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindProfilesL2ogre
//
// Use this API command to retrieve a list of L2oGRE profile.
//
// Operation ID: findProfilesL2ogre
// Operation path: /profiles/l2ogre
// Success code: 200 (OK)
func (s *WSGL2oGREService) FindProfilesL2ogre(ctx context.Context, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesL2ogre, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileListAPIResponse), err
}

// FindProfilesL2ogreById
//
// Use this API command to retrieve L2oGRE profile by ID.
//
// Operation ID: findProfilesL2ogreById
// Operation path: /profiles/l2ogre/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGL2oGREService) FindProfilesL2ogreById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileL2oGREProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileL2oGREProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileL2oGREProfileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindProfilesL2ogreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileL2oGREProfileAPIResponse), err
}

// FindProfilesL2ogreByQueryCriteria
//
// Use this API command to query a list of L2oGRE profile.
//
// Operation ID: findProfilesL2ogreByQueryCriteria
// Operation path: /profiles/l2ogre/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2oGREService) FindProfilesL2ogreByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileL2oGREProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileL2oGREProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileL2oGREProfileListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindProfilesL2ogreByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileL2oGREProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileL2oGREProfileListAPIResponse), err
}

// PartialUpdateProfilesL2ogreById
//
// Use this API command to modify the configuration of L2oGRE profile.
//
// Operation ID: partialUpdateProfilesL2ogreById
// Operation path: /profiles/l2ogre/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyL2oGREProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGL2oGREService) PartialUpdateProfilesL2ogreById(ctx context.Context, body *WSGProfileModifyL2oGREProfile, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateProfilesL2ogreById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
