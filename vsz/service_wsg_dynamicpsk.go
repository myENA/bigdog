package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
)

type WSGDynamicPSKService struct {
    c *Client
}

func NewWSGDynamicPSKService (c *Client) *WSGDynamicPSKService {
    s := new(WSGDynamicPSKService)
    s.c = c
    return s
}

