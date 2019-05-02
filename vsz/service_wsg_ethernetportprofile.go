package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ethernetport"
)

type WSGEthernetPortProfileService struct {
    client *Client
}

func NewWSGEthernetPortProfileService (client *Client) *WSGEthernetPortProfileService {
    s := new(WSGEthernetPortProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGEthernetPortProfileService () *WSGEthernetPortProfileService {
    serv := new(WSGEthernetPortProfileService)
    serv.client = ss.client
    return serv
}

