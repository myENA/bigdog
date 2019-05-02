package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGTestAAAServerService struct {
    c *Client
}

func NewWSGTestAAAServerService (c *Client) *WSGTestAAAServerService {
    s := new(WSGTestAAAServerService)
    s.c = c
    return s
}

