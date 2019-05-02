package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGHotspot20IdentityProviderProfileService struct {
    client *Client
}

func NewWSGHotspot20IdentityProviderProfileService (client *Client) *WSGHotspot20IdentityProviderProfileService {
    s := new(WSGHotspot20IdentityProviderProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGHotspot20IdentityProviderProfileService () *WSGHotspot20IdentityProviderProfileService {
    serv := new(WSGHotspot20IdentityProviderProfileService)
    serv.client = ss.client
    return serv
}

