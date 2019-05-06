package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
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

func (s *WSGSNMPAgentService) FindSystemSnmpAgent (ctx context.Context) (*system.SnmpAgentConfiguration, error) {
}

