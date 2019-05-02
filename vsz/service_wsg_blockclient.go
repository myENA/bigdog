package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBlockClientService struct {
    c *Client
}

func NewWSGBlockClientService (c *Client) *WSGBlockClientService {
    s := new(WSGBlockClientService)
    s.c = c
    return s
}

