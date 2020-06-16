package bigdog

// API Version: v9_0

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
func (s *WSGTTGPDGService) AddProfilesTtgpdg(ctx context.Context, body *WSGProfileCreateTtgpdgProfile, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesTtgpdg, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
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
func (s *WSGTTGPDGService) DeleteProfilesTtgpdg(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTtgpdg, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
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
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgApnRealmsById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTtgpdgApnRealmsById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
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
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTtgpdgById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
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
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgDhcpRelayById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesTtgpdgDhcpRelayById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesTtgpdg
//
// Operation ID: findProfilesTtgpdg
//
// Use this API command to retrieve a list of TTG+PDG profile.
func (s *WSGTTGPDGService) FindProfilesTtgpdg(ctx context.Context, mutators ...RequestMutator) (*WSGProfileList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTtgpdg, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *WSGTTGPDGService) FindProfilesTtgpdgById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileTtgpdgProfile, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTtgpdgById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileTtgpdgProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *WSGTTGPDGService) FindProfilesTtgpdgByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileTtgpdgProfileList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTtgpdgByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileTtgpdgProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *WSGTTGPDGService) PartialUpdateProfilesTtgpdgById(ctx context.Context, body *WSGProfileTtgpdgProfileConfiguration, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTtgpdgById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
