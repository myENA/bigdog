package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGApplicationVisibilityControlService struct {
	apiClient *APIClient
}

func NewWSGApplicationVisibilityControlService(c *APIClient) *WSGApplicationVisibilityControlService {
	s := new(WSGApplicationVisibilityControlService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGApplicationVisibilityControlService() *WSGApplicationVisibilityControlService {
	return NewWSGApplicationVisibilityControlService(ss.apiClient)
}

// AddAvcApplicationPolicy
//
// Use this API command to create a new AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCCreateApplicationPolicyProfile
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicy(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcApplicationPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddAvcApplicationPolicyV2
//
// Use this API command to create a new AVC Application Policy profile.
//
// Request Body:
//	 - body *WSGAVCCreateApplicationPolicyProfile
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCCreateApplicationPolicyProfile) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcApplicationPolicyV2, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// AddAvcSignaturePackageUpload
//
// Update AVC Signature Package by upload file (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body []byte
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageUpload(ctx context.Context, body []byte) (*WSGAVCSignaturePackage, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcSignaturePackageUpload, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddAvcSignaturePackageV2Upload
//
// Update AVC Signature Package by upload file.
//
// Request Body:
//	 - body []byte
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageV2Upload(ctx context.Context, body []byte) (*WSGAVCSignaturePackage, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcSignaturePackageV2Upload, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddAvcUserDefined
//
// Use this API command to create a new AVC User Defined profile.
//
// Request Body:
//	 - body *WSGAVCCreateUserDefinedProfile
func (s *WSGApplicationVisibilityControlService) AddAvcUserDefined(ctx context.Context, body *WSGAVCCreateUserDefinedProfile) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAvcUserDefined, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteAvcApplicationPolicy
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicy(ctx context.Context, body *WSGAVCDeleteBulk) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// DeleteAvcApplicationPolicyById
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteAvcApplicationPolicyV2
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2(ctx context.Context, body *WSGAVCDeleteBulk) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyV2, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// DeleteAvcApplicationPolicyV2ById
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2ById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcApplicationPolicyV2ById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteAvcUserDefined
//
// Use this API command to delete a AVC User Defined Profile.
//
// Request Body:
//	 - body *WSGAVCDeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefined(ctx context.Context, body *WSGAVCDeleteBulk) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcUserDefined, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// DeleteAvcUserDefinedById
//
// Use this API command to delete a AVC User Defined Profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefinedById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAvcUserDefinedById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindApplicationPolicyByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAVCApplicationPolicyProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCApplicationPolicyProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApplicationPolicyByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApplicationPolicyV2ByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyV2ByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAVCApplicationPolicyProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCApplicationPolicyProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindApplicationPolicyV2ByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcApplicationPolicyById
//
// Use this API command to retrieve a AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyById(ctx context.Context, id string) (*WSGAVCApplicationPolicyProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCApplicationPolicyProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcApplicationPolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcApplicationPolicyV2ById
//
// Use this API command to retrieve a AVC Application Policy profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyV2ById(ctx context.Context, id string) (*WSGAVCApplicationPolicyProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCApplicationPolicyProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcApplicationPolicyV2ById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationPolicyProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackage
//
// Get current Signature Package info (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackage(ctx context.Context) (*WSGAVCSignaturePackage, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackage, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackageApplicationByApplicationName
//
// Get Application info (catId, appId and name) by application name (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - applicationName string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplicationByApplicationName(ctx context.Context, applicationName string) (*WSGAVCApplication, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCApplication
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, applicationName, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageApplicationByApplicationName, true)
	req.SetPathParameter("applicationName", applicationName)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplication()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackageApplications
//
// Get Application list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplications(ctx context.Context) (*WSGAVCApplicationList, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCApplicationList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageApplications, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackageCategories
//
// Get Application Category list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategories(ctx context.Context) (*WSGAVCAppCategoryList, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCAppCategoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageCategories, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCAppCategoryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackageCategoryByCategoryName
//
// Get Application Category info (catId and name) by category name (for 5.0 and Earlier Firmware Versions).
//
// Required Parameters:
// - categoryName string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategoryByCategoryName(ctx context.Context, categoryName string) (*WSGAVCAppCategory, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCAppCategory
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, categoryName, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageCategoryByCategoryName, true)
	req.SetPathParameter("categoryName", categoryName)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCAppCategory()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackageV2
//
// Get current Signature Package info.
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2(ctx context.Context) (*WSGAVCSignaturePackage, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCSignaturePackage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageV2, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCSignaturePackage()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackageV2Applications
//
// Get Application list from current Signature Package.
//
// Optional Parameters:
// - appName string
//		- nullable
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Applications(ctx context.Context, optionalParams map[string][]string) (*WSGAVCApplicationList, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCApplicationList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageV2Applications, true)
	if v, ok := optionalParams["appName"]; ok {
		req.AddQueryParameter("appName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCApplicationList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcSignaturePackageV2Categories
//
// Get Application Category list from current Signature Package.
//
// Optional Parameters:
// - catName string
//		- nullable
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Categories(ctx context.Context, optionalParams map[string][]string) (*WSGAVCAppCategoryList, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCAppCategoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcSignaturePackageV2Categories, true)
	if v, ok := optionalParams["catName"]; ok {
		req.AddQueryParameter("catName", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCAppCategoryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindAvcUserDefinedById
//
// Use this API command to retrieve a AVC User Defined profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcUserDefinedById(ctx context.Context, id string) (*WSGAVCUserDefinedProfile, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCUserDefinedProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAvcUserDefinedById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCUserDefinedProfile()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindUserDefinedByQueryCriteria
//
// Use this API command to retrieve a list of AVC User Defined profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindUserDefinedByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGAVCUserDefinedProfileList, error) {
	var (
		req      *APIRequest
		resp     *WSGAVCUserDefinedProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindUserDefinedByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAVCUserDefinedProfileList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateAvcApplicationPolicyById
//
// Use this API command to modify the basic information on AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *WSGAVCModifyApplicationPolicyProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAvcApplicationPolicyById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// PartialUpdateAvcApplicationPolicyV2ById
//
// Use this API command to modify the basic information on AVC Application Policy profile.
//
// Request Body:
//	 - body *WSGAVCModifyApplicationPolicyProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyV2ById(ctx context.Context, body *WSGAVCModifyApplicationPolicyProfile, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAvcApplicationPolicyV2ById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// PartialUpdateAvcUserDefinedById
//
// Use this API command to modify the basic information on AVC User Defined profile.
//
// Request Body:
//	 - body *WSGAVCModifyUserDefinedProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcUserDefinedById(ctx context.Context, body *WSGAVCModifyUserDefinedProfile, id string) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAvcUserDefinedById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}
