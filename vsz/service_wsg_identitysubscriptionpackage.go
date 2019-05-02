package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
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

