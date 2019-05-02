package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGSNMPAgentService struct {
    c *Client
}

func NewWSGSNMPAgentService (c *Client) *WSGSNMPAgentService {
    s := new(WSGSNMPAgentService)
    s.c = c
    return s
}

