package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/avc"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
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
	serv := new(WSGApplicationVisibilityControlService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddAvcApplicationPolicy
//
// Use this API command to create a new AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *avc.CreateApplicationPolicyProfile
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicy(ctx context.Context, body *avc.CreateApplicationPolicyProfile) (*common.CreateResult, error) {
}

// AddAvcApplicationPolicyV2
//
// Use this API command to create a new AVC Application Policy profile.
//
// Request Body:
//	 - body *avc.CreateApplicationPolicyProfile
func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicyV2(ctx context.Context, body *avc.CreateApplicationPolicyProfile) (*common.CreateResult, error) {
}

// AddAvcSignaturePackageUpload
//
// Update AVC Signature Package by upload file (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageUpload(ctx context.Context) (*avc.SignaturePackage, error) {
}

// AddAvcSignaturePackageV2Upload
//
// Update AVC Signature Package by upload file.
func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageV2Upload(ctx context.Context) (*avc.SignaturePackage, error) {
}

// AddAvcUserDefined
//
// Use this API command to create a new AVC User Defined profile.
//
// Request Body:
//	 - body *avc.CreateUserDefinedProfile
func (s *WSGApplicationVisibilityControlService) AddAvcUserDefined(ctx context.Context, body *avc.CreateUserDefinedProfile) (*common.CreateResult, error) {
}

// DeleteAvcApplicationPolicy
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *avc.DeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicy(ctx context.Context, body *avc.DeleteBulk) error {
}

// DeleteAvcApplicationPolicyById
//
// Use this API command to delete a AVC Application Policy Profile (for 5.0 and Earlier Firmware Versions).
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteAvcApplicationPolicyV2
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Request Body:
//	 - body *avc.DeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2(ctx context.Context, body *avc.DeleteBulk) error {
}

// DeleteAvcApplicationPolicyV2ById
//
// Use this API command to delete a AVC Application Policy Profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcApplicationPolicyV2ById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteAvcUserDefined
//
// Use this API command to delete a AVC User Defined Profile.
//
// Request Body:
//	 - body *avc.DeleteBulk
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefined(ctx context.Context, body *avc.DeleteBulk) (*common.EmptyResult, error) {
}

// DeleteAvcUserDefinedById
//
// Use this API command to delete a AVC User Defined Profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) DeleteAvcUserDefinedById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindApplicationPolicyByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*avc.ApplicationPolicyProfileList, error) {
}

// FindApplicationPolicyV2ByQueryCriteria
//
// Use this API command to retrieve a list of AVC Application Policy profiles.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyV2ByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*avc.ApplicationPolicyProfileList, error) {
}

// FindAvcApplicationPolicyById
//
// Use this API command to retrieve a AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyById(ctx context.Context, pId string) (*avc.ApplicationPolicyProfile, error) {
}

// FindAvcApplicationPolicyV2ById
//
// Use this API command to retrieve a AVC Application Policy profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyV2ById(ctx context.Context, pId string) (*avc.ApplicationPolicyProfile, error) {
}

// FindAvcSignaturePackage
//
// Get current Signature Package info (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackage(ctx context.Context) (*avc.SignaturePackage, error) {
}

// FindAvcSignaturePackageApplicationByApplicationName
//
// Get Application info (catId, appId and name) by application name (for 5.0 and Earlier Firmware Versions).
//
// Path Parameters:
// - pApplicationName string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplicationByApplicationName(ctx context.Context, pApplicationName string) (*avc.Application, error) {
}

// FindAvcSignaturePackageApplications
//
// Get Application list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplications(ctx context.Context) (*avc.ApplicationList, error) {
}

// FindAvcSignaturePackageCategories
//
// Get Application Category list from current Signature Package (for 5.0 and Earlier Firmware Versions).
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategories(ctx context.Context) (*avc.AppCategoryList, error) {
}

// FindAvcSignaturePackageCategoryByCategoryName
//
// Get Application Category info (catId and name) by category name (for 5.0 and Earlier Firmware Versions).
//
// Path Parameters:
// - pCategoryName string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategoryByCategoryName(ctx context.Context, pCategoryName string) (*avc.AppCategory, error) {
}

// FindAvcSignaturePackageV2
//
// Get current Signature Package info.
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2(ctx context.Context) (*avc.SignaturePackage, error) {
}

// FindAvcSignaturePackageV2Applications
//
// Get Application list from current Signature Package.
//
// Query Parameters:
// - qAppName string
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Applications(ctx context.Context, qAppName string) (*avc.ApplicationList, error) {
}

// FindAvcSignaturePackageV2Categories
//
// Get Application Category list from current Signature Package.
//
// Query Parameters:
// - qCatName string
func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Categories(ctx context.Context, qCatName string) (*avc.AppCategoryList, error) {
}

// FindAvcUserDefinedById
//
// Use this API command to retrieve a AVC User Defined profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) FindAvcUserDefinedById(ctx context.Context, pId string) (*avc.UserDefinedProfile, error) {
}

// FindUserDefinedByQueryCriteria
//
// Use this API command to retrieve a list of AVC User Defined profiles.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGApplicationVisibilityControlService) FindUserDefinedByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*avc.UserDefinedProfileList, error) {
}

// PartialUpdateAvcApplicationPolicyById
//
// Use this API command to modify the basic information on AVC Application Policy profile (for 5.0 and Earlier Firmware Versions).
//
// Request Body:
//	 - body *avc.ModifyApplicationPolicyProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyById(ctx context.Context, body *avc.ModifyApplicationPolicyProfile, pId string) (*common.EmptyResult, error) {
}

// PartialUpdateAvcApplicationPolicyV2ById
//
// Use this API command to modify the basic information on AVC Application Policy profile.
//
// Request Body:
//	 - body *avc.ModifyApplicationPolicyProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcApplicationPolicyV2ById(ctx context.Context, body *avc.ModifyApplicationPolicyProfile, pId string) (*common.EmptyResult, error) {
}

// PartialUpdateAvcUserDefinedById
//
// Use this API command to modify the basic information on AVC User Defined profile.
//
// Request Body:
//	 - body *avc.ModifyUserDefinedProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGApplicationVisibilityControlService) PartialUpdateAvcUserDefinedById(ctx context.Context, body *avc.ModifyUserDefinedProfile, pId string) (*common.EmptyResult, error) {
}
