package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityGuestPassService struct {
    client *Client
}

func NewWSGIdentityGuestPassService (client *Client) *WSGIdentityGuestPassService {
    s := new(WSGIdentityGuestPassService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIdentityGuestPassService () *WSGIdentityGuestPassService {
    serv := new(WSGIdentityGuestPassService)
    serv.client = ss.client
    return serv
}

