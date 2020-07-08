package bigdog

// API Version: v9_0

import (
	"context"
	"io"
	"net/http"
)

type WSGApplicationVisibilityControlService struct {
	apiClient *VSZClient
}

func NewWSGApplicationVisibilityControlService(c *VSZClient) *WSGApplicationVisibilityControlService {
	s := new(WSGApplicationVisibilityControlService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGApplicationVisibilityControlService() *WSGApplicationVisibilityControlService {
	return NewWSGApplicationVisibilityControlService(ss.apiClient)
}

// AddAvcApplicationPolicy
//
// Operation ID: addAvcApplicationPolicy
//
// Use this API command to create a new AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCCreateApplicationPolicyProfile
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicy(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcApplicationPolicy, true)
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

// AddAvcApplicationPolicyV2
//
// Operation ID: addAvcApplicationPolicyV2
//
// Use this API command to create a new AVC Application Policy profile.
//
// Request Body:
//	 - body *WSGAVCCreateApplicationPolicyProfile
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcApplicationPolicyV2, true)
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

// AddAvcSignaturePackageUpload
//
// Operation ID: addAvcSignaturePackageUpload
//
// Update AVC Signature Package by upload file (for 5.0 and Earlier Firmware Versions).
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcSignaturePackageUpload, true)
	req.SetHeader(headerKeyContentType, headerValueMultipartFormData)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAvcSignaturePackageV2Upload
//
// Operation ID: addAvcSignaturePackageV2Upload
//
// Update AVC Signature Package by upload file.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageV2Upload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcSignaturePackageV2Upload, true)
	req.SetHeader(headerKeyContentType, headerValueMultipartFormData)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAvcUserDefined
//
// Operation ID: addAvcUserDefined
//
// Use this API command to create a new AVC User Defined profile.
//
// Request Body:
//	 - body *WSGAVCCreateUserDefinedProfile
func (s *WSGApplicationVisibilityControlService) AddAvcUserDefined(ctx context.Context, body *WSGAVCCreateUserDefinedProfile, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcUserDefined, true)
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

// DeleteAvcApplicationPolicy
//
// Operation ID: deleteAvcApplicationPolicy
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicy(ctx context.Context, body *WSGAVCDeleteBulk, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicy, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteAvcApplicationPolicyById
//
// Operation ID: deleteAvcApplicationPolicyById
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteAvcApplicationPolicyV2
//
// Operation ID: deleteAvcApplicationPolicyV2
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCDeleteBulk, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyV2, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteAvcApplicationPolicyV2ById
//
// Operation ID: deleteAvcApplicationPolicyV2ById
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2ById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyV2ById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteAvcUserDefined
//
// Operation ID: deleteAvcUserDefined
//
// Use this API command to delete a AVC User Defined Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefined(ctx context.Context, body *WSGAVCDeleteBulk, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcUserDefined, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteAvcUserDefinedById
//
// Operation ID: deleteAvcUserDefinedById
//
// Use this API command to delete a AVC User Defined Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefinedById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcUserDefinedById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApplicationPolicyByQueryCriteria
//
// Operation ID: findApplicationPolicyByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCApplicationPolicyProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApplicationPolicyByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCApplicationPolicyProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApplicationPolicyV2ByQueryCriteria
//
// Operation ID: findApplicationPolicyV2ByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyV2ByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCApplicationPolicyProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApplicationPolicyV2ByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCApplicationPolicyProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcApplicationPolicyById
//
// Operation ID: findAvcApplicationPolicyById
//
// Use this API command to retrieve a AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCApplicationPolicyProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcApplicationPolicyById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCApplicationPolicyProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcApplicationPolicyV2ById
//
// Operation ID: findAvcApplicationPolicyV2ById
//
// Use this API command to retrieve a AVC Application Policy profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyV2ById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCApplicationPolicyProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcApplicationPolicyV2ById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCApplicationPolicyProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackage
//
// Operation ID: findAvcSignaturePackage
//
// Get current Signature Package info (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackage(ctx context.Context, mutators ...RequestMutator) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackage, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageApplicationByApplicationName
//
// Operation ID: findAvcSignaturePackageApplicationByApplicationName
//
// Get Application info (catId, appId and name) by application name (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - applicationName string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplicationByApplicationName(ctx context.Context, applicationName string, mutators ...RequestMutator) (*WSGAVCApplication, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCApplication
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageApplicationByApplicationName, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("applicationName", applicationName)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCApplication()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageApplications
//
// Operation ID: findAvcSignaturePackageApplications
//
// Get Application list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplications(ctx context.Context, mutators ...RequestMutator) (*WSGAVCApplicationList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCApplicationList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageApplications, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCApplicationList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageCategories
//
// Operation ID: findAvcSignaturePackageCategories
//
// Get Application Category list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategories(ctx context.Context, mutators ...RequestMutator) (*WSGAVCAppCategoryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCAppCategoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageCategories, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCAppCategoryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageCategoryByCategoryName
//
// Operation ID: findAvcSignaturePackageCategoryByCategoryName
//
// Get Application Category info (catId and name) by category name (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - categoryName string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategoryByCategoryName(ctx context.Context, categoryName string, mutators ...RequestMutator) (*WSGAVCAppCategory, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCAppCategory
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageCategoryByCategoryName, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("categoryName", categoryName)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCAppCategory()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageV2
//
// Operation ID: findAvcSignaturePackageV2
//
// Get current Signature Package info.
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2(ctx context.Context, mutators ...RequestMutator) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageV2, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageV2Applications
//
// Operation ID: findAvcSignaturePackageV2Applications
//
// Get Application list from current Signature Package.
//
// Optional Parameters:
// - appName string
//		- nullable
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Applications(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAVCApplicationList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCApplicationList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageV2Applications, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["appName"]; ok && len(v) > 0 {
		req.SetQueryParameter("appName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCApplicationList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageV2Categories
//
// Operation ID: findAvcSignaturePackageV2Categories
//
// Get Application Category list from current Signature Package.
//
// Optional Parameters:
// - catName string
//		- nullable
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Categories(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAVCAppCategoryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCAppCategoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageV2Categories, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if v, ok := optionalParams["catName"]; ok && len(v) > 0 {
		req.SetQueryParameter("catName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCAppCategoryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcUserDefinedById
//
// Operation ID: findAvcUserDefinedById
//
// Use this API command to retrieve a AVC User Defined profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcUserDefinedById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAVCUserDefinedProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCUserDefinedProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcUserDefinedById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCUserDefinedProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUserDefinedByQueryCriteria
//
// Operation ID: findUserDefinedByQueryCriteria
//
// Use this API command to retrieve a list of AVC User Defined profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindUserDefinedByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAVCUserDefinedProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAVCUserDefinedProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindUserDefinedByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAVCUserDefinedProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateAvcApplicationPolicyById
//
// Operation ID: partialUpdateAvcApplicationPolicyById
//
// Use this API command to modify the configuration on AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCModifyApplicationPolicyProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAvcApplicationPolicyById, true)
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

// PartialUpdateAvcApplicationPolicyV2ById
//
// Operation ID: partialUpdateAvcApplicationPolicyV2ById
//
// Use this API command to modify the configuration on AVC Application Policy profile.
//
// Request Body:
//	 - body *WSGAVCModifyApplicationPolicyProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyV2ById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAvcApplicationPolicyV2ById, true)
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

// PartialUpdateAvcUserDefinedById
//
// Operation ID: partialUpdateAvcUserDefinedById
//
// Use this API command to modify the configuration on AVC User Defined profile.
//
// Request Body:
//	 - body *WSGAVCModifyUserDefinedProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcUserDefinedById(ctx context.Context, body *WSGAVCModifyUserDefinedProfile, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAvcUserDefinedById, true)
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
