package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dhcpmsgstats"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dhcppools"
)

type WSGDHCPService struct {
    client *Client
}

func NewWSGDHCPService (client *Client) *WSGDHCPService {
    s := new(WSGDHCPService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDHCPService () *WSGDHCPService {
    serv := new(WSGDHCPService)
    serv.client = ss.client
    return serv
}

