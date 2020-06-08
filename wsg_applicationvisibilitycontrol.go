package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGApplicationvisibilitycontrolService struct {
	apiClient *VSZClient
}

func NewWSGApplicationvisibilitycontrolService(c *VSZClient) *WSGApplicationvisibilitycontrolService {
	s := new(WSGApplicationvisibilitycontrolService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGApplicationvisibilitycontrolService() *WSGApplicationvisibilitycontrolService {
	return NewWSGApplicationvisibilitycontrolService(ss.apiClient)
}

// AddAvcApplicationPolicy
//
// Use this API command to create a new AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCCreateApplicationPolicyProfile
func (s *WSGApplicationvisibilitycontrolService) AddAvcApplicationPolicy(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddAvcApplicationPolicyV2
//
// Use this API command to create a new AVC Application Policy profile.
//
// Request Body:
//	 - body *WSGAVCCreateApplicationPolicyProfile
func (s *WSGApplicationvisibilitycontrolService) AddAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddAvcSignaturePackageUpload
//
// Update AVC Signature Package by upload file (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body []byte
func (s *WSGApplicationvisibilitycontrolService) AddAvcSignaturePackageUpload(ctx context.Context, body []byte) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAvcSignaturePackageV2Upload
//
// Update AVC Signature Package by upload file.
//
// Request Body:
//	 - body []byte
func (s *WSGApplicationvisibilitycontrolService) AddAvcSignaturePackageV2Upload(ctx context.Context, body []byte) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddAvcUserDefined
//
// Use this API command to create a new AVC User Defined profile.
//
// Request Body:
//	 - body *WSGAVCCreateUserDefinedProfile
func (s *WSGApplicationvisibilitycontrolService) AddAvcUserDefined(ctx context.Context, body *WSGAVCCreateUserDefinedProfile) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteAvcApplicationPolicy
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationvisibilitycontrolService) DeleteAvcApplicationPolicy(ctx context.Context, body *WSGAVCDeleteBulk) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteAvcApplicationPolicyById
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) DeleteAvcApplicationPolicyById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteAvcApplicationPolicyV2
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationvisibilitycontrolService) DeleteAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCDeleteBulk) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyV2, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteAvcApplicationPolicyV2ById
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) DeleteAvcApplicationPolicyV2ById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteAvcUserDefined
//
// Use this API command to delete a AVC User Defined Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationvisibilitycontrolService) DeleteAvcUserDefined(ctx context.Context, body *WSGAVCDeleteBulk) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteAvcUserDefinedById
//
// Use this API command to delete a AVC User Defined Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) DeleteAvcUserDefinedById(ctx context.Context, id string) (*APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApplicationPolicyByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationvisibilitycontrolService) FindApplicationPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAVCApplicationPolicyProfileList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApplicationPolicyV2ByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationvisibilitycontrolService) FindApplicationPolicyV2ByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAVCApplicationPolicyProfileList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcApplicationPolicyById
//
// Use this API command to retrieve a AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) FindAvcApplicationPolicyById(ctx context.Context, id string) (*WSGAVCApplicationPolicyProfile, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcApplicationPolicyV2ById
//
// Use this API command to retrieve a AVC Application Policy profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) FindAvcApplicationPolicyV2ById(ctx context.Context, id string) (*WSGAVCApplicationPolicyProfile, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackage
//
// Get current Signature Package info (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackage(ctx context.Context) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageApplicationByApplicationName
//
// Get Application info (catId, appId and name) by application name (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - applicationName string
//		- required
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackageApplicationByApplicationName(ctx context.Context, applicationName string) (*WSGAVCApplication, *APIResponseMeta, error) {
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
	req.SetPathParameter("applicationName", applicationName)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplication()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageApplications
//
// Get Application list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackageApplications(ctx context.Context) (*WSGAVCApplicationList, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageCategories
//
// Get Application Category list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackageCategories(ctx context.Context) (*WSGAVCAppCategoryList, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCAppCategoryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageCategoryByCategoryName
//
// Get Application Category info (catId and name) by category name (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - categoryName string
//		- required
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackageCategoryByCategoryName(ctx context.Context, categoryName string) (*WSGAVCAppCategory, *APIResponseMeta, error) {
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
	req.SetPathParameter("categoryName", categoryName)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCAppCategory()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageV2
//
// Get current Signature Package info.
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackageV2(ctx context.Context) (*WSGAVCSignaturePackage, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageV2Applications
//
// Get Application list from current Signature Package.
//
// Optional Parameters:
// - appName string
//		- nullable
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackageV2Applications(ctx context.Context, optionalParams map[string][]string) (*WSGAVCApplicationList, *APIResponseMeta, error) {
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
	if v, ok := optionalParams["appName"]; ok && len(v) > 0 {
		req.SetQueryParameter("appName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcSignaturePackageV2Categories
//
// Get Application Category list from current Signature Package.
//
// Optional Parameters:
// - catName string
//		- nullable
func (s *WSGApplicationvisibilitycontrolService) FindAvcSignaturePackageV2Categories(ctx context.Context, optionalParams map[string][]string) (*WSGAVCAppCategoryList, *APIResponseMeta, error) {
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
	if v, ok := optionalParams["catName"]; ok && len(v) > 0 {
		req.SetQueryParameter("catName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCAppCategoryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAvcUserDefinedById
//
// Use this API command to retrieve a AVC User Defined profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) FindAvcUserDefinedById(ctx context.Context, id string) (*WSGAVCUserDefinedProfile, *APIResponseMeta, error) {
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
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCUserDefinedProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindUserDefinedByQueryCriteria
//
// Use this API command to retrieve a list of AVC User Defined profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationvisibilitycontrolService) FindUserDefinedByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAVCUserDefinedProfileList, *APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCUserDefinedProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateAvcApplicationPolicyById
//
// Use this API command to modify the configuration on AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCModifyApplicationPolicyProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) PartialUpdateAvcApplicationPolicyById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateAvcApplicationPolicyV2ById
//
// Use this API command to modify the configuration on AVC Application Policy profile.
//
// Request Body:
//	 - body *WSGAVCModifyApplicationPolicyProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) PartialUpdateAvcApplicationPolicyV2ById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateAvcUserDefinedById
//
// Use this API command to modify the configuration on AVC User Defined profile.
//
// Request Body:
//	 - body *WSGAVCModifyUserDefinedProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationvisibilitycontrolService) PartialUpdateAvcUserDefinedById(ctx context.Context, body *WSGAVCModifyUserDefinedProfile, id string) (*APIResponseMeta, error) {
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
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
