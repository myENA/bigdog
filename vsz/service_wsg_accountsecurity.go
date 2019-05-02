package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/accountsecurityprofile"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAccountSecurityService struct {
    client *Client
}

func NewWSGAccountSecurityService (client *Client) *WSGAccountSecurityService {
    s := new(WSGAccountSecurityService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccountSecurityService () *WSGAccountSecurityService {
    serv := new(WSGAccountSecurityService)
    serv.client = ss.client
    return serv
}

