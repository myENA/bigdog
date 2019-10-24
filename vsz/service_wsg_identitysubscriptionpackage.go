package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentitySubscriptionPackageService struct {
	client *Client
}

func NewWSGIdentitySubscriptionPackageService(client *Client) *WSGIdentitySubscriptionPackageService {
	s := new(WSGIdentitySubscriptionPackageService)
	s.client = client
	return s
}

func (ss *WSGService) WSGIdentitySubscriptionPackageService() *WSGIdentitySubscriptionPackageService {
	serv := new(WSGIdentitySubscriptionPackageService)
	serv.client = ss.client
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

// FindIdentityPackages
//
// Use this API command to retrieve a list of subscription package.
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackages(ctx context.Context) (*identity.SubscriptionPackageList, error) {
}

// FindIdentityPackagesById
//
// Use this API command to retrieve subscription package.
func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackagesById(ctx context.Context, pId string) (*identity.SubscriptionPackage, error) {
}
