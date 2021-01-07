package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGTTGPDGService struct {
	apiClient *VSZClient
}

func NewWSGTTGPDGService(c *VSZClient) *WSGTTGPDGService {
	s := new(WSGTTGPDGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTTGPDGService() *WSGTTGPDGService {
	return NewWSGTTGPDGService(ss.apiClient)
}

// AddProfilesTtgpdg
//
// Operation ID: addProfilesTtgpdg
//
// Use this API command to create TTG+PDG profile.
//
// Request Body:
//	 - body *WSGProfileCreateTtgpdgProfile
func (s *WSGTTGPDGService) AddProfilesTtgpdg(ctx context.Context, body *WSGProfileCreateTtgpdgProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddProfilesTtgpdg, true)
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

// DeleteProfilesTtgpdg
//
// Operation ID: deleteProfilesTtgpdg
//
// Use this API command to delete multiple TTG PDG profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGTTGPDGService) DeleteProfilesTtgpdg(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesTtgpdg, true)
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

// DeleteProfilesTtgpdgApnRealmsById
//
// Operation ID: deleteProfilesTtgpdgApnRealmsById
//
// Use this API command to disable the APN realm of TTG PDG profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgApnRealmsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesTtgpdgApnRealmsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesTtgpdgById
//
// Operation ID: deleteProfilesTtgpdgById
//
// Use this API command to delete TTG PDG profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesTtgpdgById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesTtgpdgDhcpRelayById
//
// Operation ID: deleteProfilesTtgpdgDhcpRelayById
//
// Use this API command to disable the DHCP relay of TTG PDG profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgDhcpRelayById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesTtgpdgDhcpRelayById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesTtgpdg
//
// Operation ID: findProfilesTtgpdg
//
// Use this API command to retrieve a list of TTG+PDG profile.
func (s *WSGTTGPDGService) FindProfilesTtgpdg(ctx context.Context, mutators ...RequestMutator) (*WSGProfileListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesTtgpdg, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesTtgpdgById
//
// Operation ID: findProfilesTtgpdgById
//
// Use this API command to retrieve TTG+PDG profile by ID.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) FindProfilesTtgpdgById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileTtgpdgProfileAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileTtgpdgProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesTtgpdgById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileTtgpdgProfile()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesTtgpdgByQueryCriteria
//
// Operation ID: findProfilesTtgpdgByQueryCriteria
//
// Use this API command to query a list of TTG+PDG profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGTTGPDGService) FindProfilesTtgpdgByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileTtgpdgProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileTtgpdgProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindProfilesTtgpdgByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileTtgpdgProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesTtgpdgById
//
// Operation ID: partialUpdateProfilesTtgpdgById
//
// Use this API command to modify the configuration of TTG+PDG profile.
//
// Request Body:
//	 - body *WSGProfileTtgpdgProfileConfiguration
//
// Required Parameters:
// - id string
//		- required
func (s *WSGTTGPDGService) PartialUpdateProfilesTtgpdgById(ctx context.Context, body *WSGProfileTtgpdgProfileConfiguration, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesTtgpdgById, true)
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
