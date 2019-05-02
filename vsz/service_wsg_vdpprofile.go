package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGVDPProfileService struct {
    client *Client
}

func NewWSGVDPProfileService (client *Client) *WSGVDPProfileService {
    s := new(WSGVDPProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGVDPProfileService () *WSGVDPProfileService {
    serv := new(WSGVDPProfileService)
    serv.client = ss.client
    return serv
}

