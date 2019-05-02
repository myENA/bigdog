package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/urlfiltering"
)

type WSGURLFilteringPolicyService struct {
    client *Client
}

func NewWSGURLFilteringPolicyService (client *Client) *WSGURLFilteringPolicyService {
    s := new(WSGURLFilteringPolicyService)
    s.client = client
    return s
}

func (ss *WSGService) WSGURLFilteringPolicyService () *WSGURLFilteringPolicyService {
    serv := new(WSGURLFilteringPolicyService)
    serv.client = ss.client
    return serv
}

