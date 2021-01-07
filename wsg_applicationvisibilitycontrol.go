package bigdog

// API Version: v9_1

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
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicy(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddAvcApplicationPolicy, true)
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

// AddAvcApplicationPolicyV2
//
// Operation ID: addAvcApplicationPolicyV2
//
// Use this API command to create a new AVC Application Policy profile.
//
// Request Body:
//	 - body *WSGAVCCreateApplicationPolicyProfile
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddAvcApplicationPolicyV2, true)
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

// AddAvcSignaturePackageUpload
//
// Operation ID: addAvcSignaturePackageUpload
//
// Update AVC Signature Package by upload file (for 5.0 and Earlier Firmware Versions).
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*WSGAVCSignaturePackageAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCSignaturePackageAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCSignaturePackageAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddAvcSignaturePackageUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*WSGAVCSignaturePackageAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCSignaturePackageAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageV2Upload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*WSGAVCSignaturePackageAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCSignaturePackageAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCSignaturePackageAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddAvcSignaturePackageV2Upload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*WSGAVCSignaturePackageAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCSignaturePackageAPIResponse), err
}

// AddAvcUserDefined
//
// Operation ID: addAvcUserDefined
//
// Use this API command to create a new AVC User Defined profile.
//
// Request Body:
//	 - body *WSGAVCCreateUserDefinedProfile
func (s *WSGApplicationVisibilityControlService) AddAvcUserDefined(ctx context.Context, body *WSGAVCCreateUserDefinedProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddAvcUserDefined, true)
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

// DeleteAvcApplicationPolicy
//
// Operation ID: deleteAvcApplicationPolicy
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicy(ctx context.Context, body *WSGAVCDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicy, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteAvcApplicationPolicyV2
//
// Operation ID: deleteAvcApplicationPolicyV2
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyV2, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2ById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyV2ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteAvcUserDefined
//
// Operation ID: deleteAvcUserDefined
//
// Use this API command to delete a AVC User Defined Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefined(ctx context.Context, body *WSGAVCDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAvcUserDefined, true)
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

// DeleteAvcUserDefinedById
//
// Operation ID: deleteAvcUserDefinedById
//
// Use this API command to delete a AVC User Defined Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefinedById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAvcUserDefinedById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindApplicationPolicyByQueryCriteria
//
// Operation ID: findApplicationPolicyByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCApplicationPolicyProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCApplicationPolicyProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindApplicationPolicyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAVCApplicationPolicyProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCApplicationPolicyProfileListAPIResponse), err
}

// FindApplicationPolicyV2ByQueryCriteria
//
// Operation ID: findApplicationPolicyV2ByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyV2ByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCApplicationPolicyProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCApplicationPolicyProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindApplicationPolicyV2ByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAVCApplicationPolicyProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCApplicationPolicyProfileListAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCApplicationPolicyProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCApplicationPolicyProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcApplicationPolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCApplicationPolicyProfileAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyV2ById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAVCApplicationPolicyProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCApplicationPolicyProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCApplicationPolicyProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcApplicationPolicyV2ById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCApplicationPolicyProfileAPIResponse), err
}

// FindAvcSignaturePackage
//
// Operation ID: findAvcSignaturePackage
//
// Get current Signature Package info (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackage(ctx context.Context, mutators ...RequestMutator) (*WSGAVCSignaturePackageAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCSignaturePackageAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCSignaturePackageAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackage, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCSignaturePackageAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplicationByApplicationName(ctx context.Context, applicationName string, mutators ...RequestMutator) (*WSGAVCApplicationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCApplicationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCApplicationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackageApplicationByApplicationName, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("applicationName", applicationName)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCApplicationAPIResponse), err
}

// FindAvcSignaturePackageApplications
//
// Operation ID: findAvcSignaturePackageApplications
//
// Get Application list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplications(ctx context.Context, mutators ...RequestMutator) (*WSGAVCApplicationListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCApplicationListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCApplicationListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackageApplications, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCApplicationListAPIResponse), err
}

// FindAvcSignaturePackageCategories
//
// Operation ID: findAvcSignaturePackageCategories
//
// Get Application Category list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategories(ctx context.Context, mutators ...RequestMutator) (*WSGAVCAppCategoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCAppCategoryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCAppCategoryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackageCategories, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCAppCategoryListAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategoryByCategoryName(ctx context.Context, categoryName string, mutators ...RequestMutator) (*WSGAVCAppCategoryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCAppCategoryAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCAppCategoryAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackageCategoryByCategoryName, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("categoryName", categoryName)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCAppCategoryAPIResponse), err
}

// FindAvcSignaturePackageV2
//
// Operation ID: findAvcSignaturePackageV2
//
// Get current Signature Package info.
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2(ctx context.Context, mutators ...RequestMutator) (*WSGAVCSignaturePackageAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCSignaturePackageAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCSignaturePackageAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackageV2, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCSignaturePackageAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Applications(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAVCApplicationListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCApplicationListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCApplicationListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackageV2Applications, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["appName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("appName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCApplicationListAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Categories(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGAVCAppCategoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCAppCategoryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCAppCategoryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcSignaturePackageV2Categories, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["catName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("catName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCAppCategoryListAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) FindAvcUserDefinedById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAVCUserDefinedProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCUserDefinedProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCUserDefinedProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAvcUserDefinedById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCUserDefinedProfileAPIResponse), err
}

// FindUserDefinedByQueryCriteria
//
// Operation ID: findUserDefinedByQueryCriteria
//
// Use this API command to retrieve a list of AVC User Defined profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindUserDefinedByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGAVCUserDefinedProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAVCUserDefinedProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAVCUserDefinedProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindUserDefinedByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGAVCUserDefinedProfileListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAVCUserDefinedProfileListAPIResponse), err
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
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateAvcApplicationPolicyById, true)
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
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyV2ById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateAvcApplicationPolicyV2ById, true)
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
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcUserDefinedById(ctx context.Context, body *WSGAVCModifyUserDefinedProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateAvcUserDefinedById, true)
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
