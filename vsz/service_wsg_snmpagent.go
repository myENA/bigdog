package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGSNMPAgentService struct {
    client *Client
}

func NewWSGSNMPAgentService (client *Client) *WSGSNMPAgentService {
    s := new(WSGSNMPAgentService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSNMPAgentService () *WSGSNMPAgentService {
    serv := new(WSGSNMPAgentService)
    serv.client = ss.client
    return serv
}

