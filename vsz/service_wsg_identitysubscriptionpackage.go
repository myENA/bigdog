package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentitySubscriptionPackageService struct {
    c *Client
}

func NewWSGIdentitySubscriptionPackageService (c *Client) *WSGIdentitySubscriptionPackageService {
    s := new(WSGIdentitySubscriptionPackageService)
    s.c = c
    return s
}

