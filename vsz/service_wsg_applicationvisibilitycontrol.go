package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/avc"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGApplicationVisibilityControlService struct {
    client *Client
}

func NewWSGApplicationVisibilityControlService (client *Client) *WSGApplicationVisibilityControlService {
    s := new(WSGApplicationVisibilityControlService)
    s.client = client
    return s
}

func (ss *WSGService) WSGApplicationVisibilityControlService () *WSGApplicationVisibilityControlService {
    serv := new(WSGApplicationVisibilityControlService)
    serv.client = ss.client
    return serv
}

func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicy (ctx context.Context) (common.CreateResult, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicyV2 (ctx context.Context) (common.CreateResult, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageUpload (ctx context.Context) (avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageV2Upload (ctx context.Context) (avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcUserDefined (ctx context.Context) (common.CreateResult, error) {
}

func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyByQueryCriteria (ctx context.Context) (avc.ApplicationPolicyProfileList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyV2ByQueryCriteria (ctx context.Context) (avc.ApplicationPolicyProfileList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyById (ctx context.Context, id string) (avc.ApplicationPolicyProfile, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyV2ById (ctx context.Context, id string) (avc.ApplicationPolicyProfile, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackage (ctx context.Context) (avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplicationByApplicationName (ctx context.Context, applicationName string) (avc.Application, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplications (ctx context.Context) (avc.ApplicationList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategories (ctx context.Context) (avc.AppCategoryList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategoryByCategoryName (ctx context.Context, categoryName string) (avc.AppCategory, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2 (ctx context.Context) (avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Applications (ctx context.Context) (avc.ApplicationList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Categories (ctx context.Context) (avc.AppCategoryList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcUserDefinedById (ctx context.Context, id string) (avc.UserDefinedProfile, error) {
}

func (s *WSGApplicationVisibilityControlService) FindUserDefinedByQueryCriteria (ctx context.Context) (avc.UserDefinedProfileList, error) {
}

