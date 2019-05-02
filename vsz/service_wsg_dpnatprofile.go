package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPNATProfileService struct {
    client *Client
}

func NewWSGDPNATProfileService (client *Client) *WSGDPNATProfileService {
    s := new(WSGDPNATProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDPNATProfileService () *WSGDPNATProfileService {
    serv := new(WSGDPNATProfileService)
    serv.client = ss.client
    return serv
}

