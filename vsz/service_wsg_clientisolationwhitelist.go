package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGClientIsolationWhitelistService struct {
    client *Client
}

func NewWSGClientIsolationWhitelistService (client *Client) *WSGClientIsolationWhitelistService {
    s := new(WSGClientIsolationWhitelistService)
    s.client = client
    return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService () *WSGClientIsolationWhitelistService {
    serv := new(WSGClientIsolationWhitelistService)
    serv.client = ss.client
    return serv
}

