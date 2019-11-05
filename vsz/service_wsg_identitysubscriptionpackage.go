package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentitySubscriptionPackageService struct {
	apiClient *APIClient
}

func NewWSGIdentitySubscriptionPackageService(c *APIClient) *WSGIdentitySubscriptionPackageService {
	s := new(WSGIdentitySubscriptionPackageService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentitySubscriptionPackageService() *WSGIdentitySubscriptionPackageService {
	serv := new(WSGIdentitySubscriptionPackageService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddIdentityPackageList
//
// Use this API command to retrieve a list of subscription package.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackageList(ctx context.Context, body *identity.QueryCriteria) (*identity.SubscriptionPackageList, error) {
}

// AddIdentityPackages
//
// Use this API command to create subscription package.
//
// Request Body:
//	 - body *identity.CreateSubscriptionPackage
func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackages(ctx context.Context, body *identity.CreateSubscriptionPackage) (*common.CreateResult, error) {
}

// DeleteIdentityPackages
//
// Use this API command to delete multiple subscription packages.
//
// Request Body:
//	 - body *identity.DeleteBulk
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackages(ctx context.Context, body *identity.DeleteBulk) error {
}

// DeleteIdentityPackagesById
//
// Use this API command to delete subscription package.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentitySubscriptionPackageService) DeleteIdentityPackagesById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindIdentityPackages
//
// Use this API command to retrieve a list of subscription package.
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackages(ctx context.Context) (*identity.SubscriptionPackageList, error) {
}

// FindIdentityPackagesById
//
// Use this API command to retrieve subscription package.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackagesById(ctx context.Context, pId string) (*identity.SubscriptionPackage, error) {
}

// PartialUpdateIdentityPackagesById
//
// Use this API command to modify the basic information of subscription package.
//
// Request Body:
//	 - body *identity.ModifySubscriptionPackage
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentitySubscriptionPackageService) PartialUpdateIdentityPackagesById(ctx context.Context, body *identity.ModifySubscriptionPackage, pId string) (*common.EmptyResult, error) {
}
