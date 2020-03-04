package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGIdentityGuestPassService struct {
	apiClient *APIClient
}

func NewWSGIdentityGuestPassService(c *APIClient) *WSGIdentityGuestPassService {
	s := new(WSGIdentityGuestPassService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityGuestPassService() *WSGIdentityGuestPassService {
	return NewWSGIdentityGuestPassService(ss.apiClient)
}

// AddIdentityGuestpassGenerate
//
// Use this API command to generate identity guest pass.
//
// Request Body:
//	 - body *WSGIdentityCreateIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassGenerate(ctx context.Context, body *WSGIdentityCreateIdentityGuestPass) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddIdentityGuestpassGenerate), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// AddIdentityGuestpassList
//
// Use this API command to retrieve a list of identity guest pass.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassList(ctx context.Context, body *WSGIdentityQueryCriteria) (*WSGIdentityGuestPassList, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddIdentityGuestpassList), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityGuestPassList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddIdentityGuestpassUpload
//
// Use this API command to upload identity guest pass csv file.
//
// Request Body:
//	 - body []byte
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUpload(ctx context.Context, body []byte) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddIdentityGuestpassUpload), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddIdentityGuestpassUploadCommon
//
// Use this API command to update common identity guest pass settings.
//
// Request Body:
//	 - body *WSGIdentityImportIdentityGuestPass
func (s *WSGIdentityGuestPassService) AddIdentityGuestpassUploadCommon(ctx context.Context, body *WSGIdentityImportIdentityGuestPass) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGAddIdentityGuestpassUploadCommon), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteIdentityGuestpass
//
// Use this API command to delete multiple identity guest passes.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpass(ctx context.Context, body *WSGIdentityDeleteBulk) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGDeleteIdentityGuestpass), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteIdentityGuestpassByUserId
//
// Use this API command to delete identity guest pass.
//
// Required Parameters:
// - userId string
//		- required
func (s *WSGIdentityGuestPassService) DeleteIdentityGuestpassByUserId(ctx context.Context, userId string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, userId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGDeleteIdentityGuestpassByUserId), true)
	req.SetPathParameter("userId", userId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// FindIdentityGuestpass
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
func (s *WSGIdentityGuestPassService) FindIdentityGuestpass(ctx context.Context, optionalParams map[string][]string) (*WSGIdentityGuestPassList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindIdentityGuestpass), true)
	if v, ok := optionalParams["displayName"]; ok {
		req.AddQueryParameter("displayName", v)
	}
	if v, ok := optionalParams["expirationFrom"]; ok {
		req.AddQueryParameter("expirationFrom", v)
	}
	if v, ok := optionalParams["expirationTo"]; ok {
		req.AddQueryParameter("expirationTo", v)
	}
	if v, ok := optionalParams["generatedTimeFrom"]; ok {
		req.AddQueryParameter("generatedTimeFrom", v)
	}
	if v, ok := optionalParams["generatedTimeTo"]; ok {
		req.AddQueryParameter("generatedTimeTo", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["timeZone"]; ok {
		req.AddQueryParameter("timeZone", v)
	}
	if v, ok := optionalParams["wlan"]; ok {
		req.AddQueryParameter("wlan", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGIdentityGuestPassList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
