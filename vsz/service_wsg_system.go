package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGSystemService struct {
    c *Client
}

func NewWSGSystemService (c *Client) *WSGSystemService {
    s := new(WSGSystemService)
    s.c = c
    return s
}

