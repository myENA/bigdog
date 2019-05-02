package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGSoftGRETunnelProfileService struct {
    client *Client
}

func NewWSGSoftGRETunnelProfileService (client *Client) *WSGSoftGRETunnelProfileService {
    s := new(WSGSoftGRETunnelProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSoftGRETunnelProfileService () *WSGSoftGRETunnelProfileService {
    serv := new(WSGSoftGRETunnelProfileService)
    serv.client = ss.client
    return serv
}

