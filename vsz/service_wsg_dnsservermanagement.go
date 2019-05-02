package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGDNSServerManagementService struct {
    client *Client
}

func NewWSGDNSServerManagementService (client *Client) *WSGDNSServerManagementService {
    s := new(WSGDNSServerManagementService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDNSServerManagementService () *WSGDNSServerManagementService {
    serv := new(WSGDNSServerManagementService)
    serv.client = ss.client
    return serv
}

