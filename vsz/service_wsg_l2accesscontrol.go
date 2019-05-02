package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGL2AccessControlService struct {
    client *Client
}

func NewWSGL2AccessControlService (client *Client) *WSGL2AccessControlService {
    s := new(WSGL2AccessControlService)
    s.client = client
    return s
}

func (ss *WSGService) WSGL2AccessControlService () *WSGL2AccessControlService {
    serv := new(WSGL2AccessControlService)
    serv.client = ss.client
    return serv
}

