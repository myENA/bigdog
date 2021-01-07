package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGHotspot20WiFiOperatorProfileService struct {
	apiClient *VSZClient
}

func NewWSGHotspot20WiFiOperatorProfileService(c *VSZClient) *WSGHotspot20WiFiOperatorProfileService {
	s := new(WSGHotspot20WiFiOperatorProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20WiFiOperatorProfileService() *WSGHotspot20WiFiOperatorProfileService {
	return NewWSGHotspot20WiFiOperatorProfileService(ss.apiClient)
}

// AddProfilesHs20Operators
//
// Operation ID: addProfilesHs20Operators
//
// Use this API command to create a new Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
func (s *WSGHotspot20WiFiOperatorProfileService) AddProfilesHs20Operators(ctx context.Context, body *WSGProfileHs20Operator, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddProfilesHs20Operators, true)
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

// DeleteProfilesHs20Operators
//
// Operation ID: deleteProfilesHs20Operators
//
// Use this API command to delete multiple Hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20Operators(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesHs20Operators, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteProfilesHs20OperatorsById
//
// Operation ID: deleteProfilesHs20OperatorsById
//
// Use this API command to delete a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesHs20OperatorsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteProfilesHs20OperatorsCertificateById
//
// Operation ID: deleteProfilesHs20OperatorsCertificateById
//
// Use this API command to disable certificate of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsCertificateById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteProfilesHs20OperatorsCertificateById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindProfilesHs20Operators
//
// Operation ID: findProfilesHs20Operators
//
// Use this API command to retrieve list of Hotspot 2.0 Wi-Fi Operators.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20Operators(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGProfileHs20OperatorListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileHs20OperatorListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileHs20OperatorListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesHs20Operators, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileHs20OperatorListAPIResponse), err
}

// FindProfilesHs20OperatorsById
//
// Operation ID: findProfilesHs20OperatorsById
//
// Use this API command to retrieve a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGProfileHs20OperatorAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileHs20OperatorAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileHs20OperatorAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindProfilesHs20OperatorsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileHs20OperatorAPIResponse), err
}

// FindProfilesHs20OperatorsByQueryCriteria
//
// Operation ID: findProfilesHs20OperatorsByQueryCriteria
//
// Query hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileHs20OperatorListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGProfileHs20OperatorListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGProfileHs20OperatorListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindProfilesHs20OperatorsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGProfileHs20OperatorListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGProfileHs20OperatorListAPIResponse), err
}

// PartialUpdateProfilesHs20OperatorsById
//
// Operation ID: partialUpdateProfilesHs20OperatorsById
//
// Use this API command to modify the configuration of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileModifyHS20Operator
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) PartialUpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileModifyHS20Operator, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateProfilesHs20OperatorsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateProfilesHs20OperatorsById
//
// Operation ID: updateProfilesHs20OperatorsById
//
// Use this API command to modify entire configuration of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) UpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileHs20Operator, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateProfilesHs20OperatorsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
