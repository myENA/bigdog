package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGWebAuthenticationService struct {
    client *Client
}

func NewWSGWebAuthenticationService (client *Client) *WSGWebAuthenticationService {
    s := new(WSGWebAuthenticationService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWebAuthenticationService () *WSGWebAuthenticationService {
    serv := new(WSGWebAuthenticationService)
    serv.client = ss.client
    return serv
}

