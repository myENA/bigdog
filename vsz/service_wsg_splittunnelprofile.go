package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/splittunnel"
)

type WSGSplitTunnelProfileService struct {
    client *Client
}

func NewWSGSplitTunnelProfileService (client *Client) *WSGSplitTunnelProfileService {
    s := new(WSGSplitTunnelProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSplitTunnelProfileService () *WSGSplitTunnelProfileService {
    serv := new(WSGSplitTunnelProfileService)
    serv.client = ss.client
    return serv
}

