package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityUserService struct {
    client *Client
}

func NewWSGIdentityUserService (client *Client) *WSGIdentityUserService {
    s := new(WSGIdentityUserService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIdentityUserService () *WSGIdentityUserService {
    serv := new(WSGIdentityUserService)
    serv.client = ss.client
    return serv
}

