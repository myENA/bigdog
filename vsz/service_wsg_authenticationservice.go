package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAuthenticationServiceService struct {
    client *Client
}

func NewWSGAuthenticationServiceService (client *Client) *WSGAuthenticationServiceService {
    s := new(WSGAuthenticationServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAuthenticationServiceService () *WSGAuthenticationServiceService {
    serv := new(WSGAuthenticationServiceService)
    serv.client = ss.client
    return serv
}

