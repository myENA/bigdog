package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGMarkRogueService struct {
    client *Client
}

func NewWSGMarkRogueService (client *Client) *WSGMarkRogueService {
    s := new(WSGMarkRogueService)
    s.client = client
    return s
}

func (ss *WSGService) WSGMarkRogueService () *WSGMarkRogueService {
    serv := new(WSGMarkRogueService)
    serv.client = ss.client
    return serv
}

