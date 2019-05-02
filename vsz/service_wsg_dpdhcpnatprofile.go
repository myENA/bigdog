package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPNATProfileService struct {
    client *Client
}

func NewWSGDPDHCPNATProfileService (client *Client) *WSGDPDHCPNATProfileService {
    s := new(WSGDPDHCPNATProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDPDHCPNATProfileService () *WSGDPDHCPNATProfileService {
    serv := new(WSGDPDHCPNATProfileService)
    serv.client = ss.client
    return serv
}

