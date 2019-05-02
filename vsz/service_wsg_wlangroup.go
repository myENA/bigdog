package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlangroup"
)

type WSGWLANGroupService struct {
    client *Client
}

func NewWSGWLANGroupService (client *Client) *WSGWLANGroupService {
    s := new(WSGWLANGroupService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWLANGroupService () *WSGWLANGroupService {
    serv := new(WSGWLANGroupService)
    serv.client = ss.client
    return serv
}

