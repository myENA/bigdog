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
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassGenerate(ctx context.Context, body *WSGIdentityCreateIdentityGuestPass, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassGenerate, true)
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

// AddIdentityGuestpassList
//
// Operation ID: addIdentityGuestpassList
//
// Use this API command to retrieve a list of identity guest pass.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassList(ctx context.Context, body *WSGIdentityQueryCriteria, mutators ...RequestMutator) (*WSGIdentityGuestPassListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityGuestPassListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityGuestPassListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGIdentityGuestPassListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityGuestPassListAPIResponse), err
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
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddIdentityGuestpassUploadCommon
//
// Operation ID: addIdentityGuestpassUploadCommon
//
// Use this API command to update common identity guest pass settings.
//
// Request Body:
//	 - body *WSGIdentityImportIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUploadCommon(ctx context.Context, body *WSGIdentityImportIdentityGuestPass, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityGuestpassUploadCommon, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteIdentityGuestpass
//
// Operation ID: deleteIdentityGuestpass
//
// Use this API command to delete multiple identity guest passes.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpass(ctx context.Context, body *WSGIdentityDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityGuestpass, true)
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

// DeleteIdentityGuestpassByUserId
//
// Operation ID: deleteIdentityGuestpassByUserId
//
// Use this API command to delete identity guest pass.
//
// Required Parameters:
// - userId string
//		- required
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpassByUserId(ctx context.Context, userId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityGuestpassByUserId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
func (s *WSGIdentityGuestPassService) FindIdentityGuestpass(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGIdentityGuestPassListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityGuestPassListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityGuestPassListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityGuestpass, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["displayName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("displayName", v)
	}
	if v, ok := optionalParams["expirationFrom"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("expirationFrom", v)
	}
	if v, ok := optionalParams["expirationTo"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("expirationTo", v)
	}
	if v, ok := optionalParams["generatedTimeFrom"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("generatedTimeFrom", v)
	}
	if v, ok := optionalParams["generatedTimeTo"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("generatedTimeTo", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["timeZone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timeZone", v)
	}
	if v, ok := optionalParams["wlan"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("wlan", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityGuestPassListAPIResponse), err
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
func (s *WSGIdentityGuestPassService) PartialUpdateIdentityGuestpassByUserId(ctx context.Context, body *WSGIdentityModifyGuestPass, userId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateIdentityGuestpassByUserId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
