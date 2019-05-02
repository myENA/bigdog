package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGMarkRogueService struct {
    c *Client
}

func NewWSGMarkRogueService (c *Client) *WSGMarkRogueService {
    s := new(WSGMarkRogueService)
    s.c = c
    return s
}

