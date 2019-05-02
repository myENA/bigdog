package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAccessPointOperationalService struct {
    c *Client
}

func NewWSGAccessPointOperationalService (c *Client) *WSGAccessPointOperationalService {
    s := new(WSGAccessPointOperationalService)
    s.c = c
    return s
}

