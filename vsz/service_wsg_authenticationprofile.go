package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAuthenticationProfileService struct {
    client *Client
}

func NewWSGAuthenticationProfileService (client *Client) *WSGAuthenticationProfileService {
    s := new(WSGAuthenticationProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAuthenticationProfileService () *WSGAuthenticationProfileService {
    serv := new(WSGAuthenticationProfileService)
    serv.client = ss.client
    return serv
}

