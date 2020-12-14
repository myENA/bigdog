package bigdog

// API Version: v9_1

import (
	"context"
	"io"
	"net/http"
)

type WSGIdentityGuestPassService struct {
	apiClient *VSZClient
}

func NewWSGIdentityGuestPassService(c *VSZClient) *WSGIdentityGuestPassService {
	s := new(WSGIdentityGuestPassService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityGuestPassService() *WSGIdentityGuestPassService {
	return NewWSGIdentityGuestPassService(ss.apiClient)
}

// AddIdentityGuestpassGenerate
//
// Operation ID: addIdentityGuestpassGenerate
//
// Use this API command to generate identity guest pass.
//
// Request Body:
//	 - body *WSGIdentityCreateIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassGenerate(ctx context.Context, body *WSGIdentityCreateIdentityGuestPass, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassGenerate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddIdentityGuestpassList
//
// Operation ID: addIdentityGuestpassList
//
// Use this API command to retrieve a list of identity guest pass.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassList(ctx context.Context, body *WSGIdentityQueryCriteria, mutators ...RequestMutator) (*WSGIdentityGuestPassList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityGuestPassList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGIdentityGuestPassList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddIdentityGuestpassUpload
//
// Operation ID: addIdentityGuestpassUpload
//
// Use this API command to upload identity guest pass csv file.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddIdentityGuestpassUploadCommon
//
// Operation ID: addIdentityGuestpassUploadCommon
//
// Use this API command to update common identity guest pass settings.
//
// Request Body:
//	 - body *WSGIdentityImportIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUploadCommon(ctx context.Context, body *WSGIdentityImportIdentityGuestPass, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassUploadCommon, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteIdentityGuestpass
//
// Operation ID: deleteIdentityGuestpass
//
// Use this API command to delete multiple identity guest passes.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpass(ctx context.Context, body *WSGIdentityDeleteBulk, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityGuestpass, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteIdentityGuestpassByUserId
//
// Operation ID: deleteIdentityGuestpassByUserId
//
// Use this API command to delete identity guest pass.
//
// Required Parameters:
// - userId string
//		- required
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpassByUserId(ctx context.Context, userId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityGuestpassByUserId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindIdentityGuestpass
//
// Operation ID: findIdentityGuestpass
//
// Use this API command to retrieve a list of identity guest pass.
//
// Optional Parameters:
// - displayName string
//		- nullable
// - expirationFrom string
//		- nullable
// - expirationTo string
//		- nullable
// - generatedTimeFrom string
//		- nullable
// - generatedTimeTo string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - timeZone string
//		- nullable
// - wlan string
//		- nullable
func (s *WSGIdentityGuestPassService) FindIdentityGuestpass(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGIdentityGuestPassList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGIdentityGuestPassList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityGuestpass, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["displayName"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("displayName", vv)
		}
	}
	if v, ok := optionalParams["expirationFrom"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("expirationFrom", vv)
		}
	}
	if v, ok := optionalParams["expirationTo"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("expirationTo", vv)
		}
	}
	if v, ok := optionalParams["generatedTimeFrom"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("generatedTimeFrom", vv)
		}
	}
	if v, ok := optionalParams["generatedTimeTo"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("generatedTimeTo", vv)
		}
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("index", vv)
		}
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("listSize", vv)
		}
	}
	if v, ok := optionalParams["timeZone"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("timeZone", vv)
		}
	}
	if v, ok := optionalParams["wlan"]; ok && len(v) > 0 {
		for _, vv := range v {
			req.QueryParams.Add("wlan", vv)
		}
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGIdentityGuestPassList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateIdentityGuestpassByUserId
//
// Operation ID: partialUpdateIdentityGuestpassByUserId
//
// Use this API command to modify the configuration of identity guest.
//
// Request Body:
//	 - body *WSGIdentityModifyGuestPass
//
// Required Parameters:
// - userId string
//		- required
func (s *WSGIdentityGuestPassService) PartialUpdateIdentityGuestpassByUserId(ctx context.Context, body *WSGIdentityModifyGuestPass, userId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateIdentityGuestpassByUserId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
