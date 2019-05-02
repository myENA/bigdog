package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGDiffServService struct {
    c *Client
}

func NewWSGDiffServService (c *Client) *WSGDiffServService {
    s := new(WSGDiffServService)
    s.c = c
    return s
}

