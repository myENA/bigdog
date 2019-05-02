package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
	"github.com/myENA/ruckus-client/vsz/types/wsg/flexivpn"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlan"
)

type WSGWLANService struct {
    client *Client
}

func NewWSGWLANService (client *Client) *WSGWLANService {
    s := new(WSGWLANService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWLANService () *WSGWLANService {
    serv := new(WSGWLANService)
    serv.client = ss.client
    return serv
}

