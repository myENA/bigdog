package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentitySubscriptionPackageService struct {
    client *Client
}

func NewWSGIdentitySubscriptionPackageService (client *Client) *WSGIdentitySubscriptionPackageService {
    s := new(WSGIdentitySubscriptionPackageService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIdentitySubscriptionPackageService () *WSGIdentitySubscriptionPackageService {
    serv := new(WSGIdentitySubscriptionPackageService)
    serv.client = ss.client
    return serv
}

func (s *WSGIdentitySubscriptionPackageService) AddIdentityPackageList (ctx context.Context) (*identity.SubscriptionPackageList, error) {
}

func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackages (ctx context.Context) (*identity.SubscriptionPackageList, error) {
}

func (s *WSGIdentitySubscriptionPackageService) FindIdentityPackagesById (ctx context.Context, id string) (*identity.SubscriptionPackage, error) {
}

