package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBonjourFencingPolicyService struct {
    client *Client
}

func NewWSGBonjourFencingPolicyService (client *Client) *WSGBonjourFencingPolicyService {
    s := new(WSGBonjourFencingPolicyService)
    s.client = client
    return s
}

func (ss *WSGService) WSGBonjourFencingPolicyService () *WSGBonjourFencingPolicyService {
    serv := new(WSGBonjourFencingPolicyService)
    serv.client = ss.client
    return serv
}

