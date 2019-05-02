package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/devicepolicy"
)

type WSGDevicePolicyService struct {
    client *Client
}

func NewWSGDevicePolicyService (client *Client) *WSGDevicePolicyService {
    s := new(WSGDevicePolicyService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDevicePolicyService () *WSGDevicePolicyService {
    serv := new(WSGDevicePolicyService)
    serv.client = ss.client
    return serv
}

