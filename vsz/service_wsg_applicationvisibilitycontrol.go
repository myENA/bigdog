package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/avc"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGApplicationVisibilityControlService struct {
	client *Client
}

func NewWSGApplicationVisibilityControlService(client *Client) *WSGApplicationVisibilityControlService {
	s := new(WSGApplicationVisibilityControlService)
	s.client = client
	return s
}

func (ss *WSGService) WSGApplicationVisibilityControlService() *WSGApplicationVisibilityControlService {
	serv := new(WSGApplicationVisibilityControlService)
	serv.client = ss.client
	return serv
}

func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicy(ctx context.Context, body *avc.CreateApplicationPolicyProfile) (*common.CreateResult, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcApplicationPolicyV2(ctx context.Context, body *avc.CreateApplicationPolicyProfile) (*common.CreateResult, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageUpload(ctx context.Context) (*avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcSignaturePackageV2Upload(ctx context.Context) (*avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) AddAvcUserDefined(ctx context.Context, body *avc.CreateUserDefinedProfile) (*common.CreateResult, error) {
}

func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*avc.ApplicationPolicyProfileList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindApplicationPolicyV2ByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*avc.ApplicationPolicyProfileList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyById(ctx context.Context, pId string) (*avc.ApplicationPolicyProfile, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcApplicationPolicyV2ById(ctx context.Context, pId string) (*avc.ApplicationPolicyProfile, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackage(ctx context.Context) (*avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplicationByApplicationName(ctx context.Context, pApplicationName string) (*avc.Application, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageApplications(ctx context.Context) (*avc.ApplicationList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategories(ctx context.Context) (*avc.AppCategoryList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageCategoryByCategoryName(ctx context.Context, pCategoryName string) (*avc.AppCategory, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2(ctx context.Context) (*avc.SignaturePackage, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Applications(ctx context.Context, qAppName string) (*avc.ApplicationList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcSignaturePackageV2Categories(ctx context.Context, qCatName string) (*avc.AppCategoryList, error) {
}

func (s *WSGApplicationVisibilityControlService) FindAvcUserDefinedById(ctx context.Context, pId string) (*avc.UserDefinedProfile, error) {
}

func (s *WSGApplicationVisibilityControlService) FindUserDefinedByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*avc.UserDefinedProfileList, error) {
}
