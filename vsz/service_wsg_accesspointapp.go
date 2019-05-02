package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
)

type WSGAccessPointAPPService struct {
    c *Client
}

func NewWSGAccessPointAPPService (c *Client) *WSGAccessPointAPPService {
    s := new(WSGAccessPointAPPService)
    s.c = c
    return s
}

