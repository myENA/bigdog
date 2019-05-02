package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPProfileService struct {
    client *Client
}

func NewWSGDPDHCPProfileService (client *Client) *WSGDPDHCPProfileService {
    s := new(WSGDPDHCPProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDPDHCPProfileService () *WSGDPDHCPProfileService {
    serv := new(WSGDPDHCPProfileService)
    serv.client = ss.client
    return serv
}

