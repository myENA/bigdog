package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/vlanpooling"
)

type WSGVlanPoolingService struct {
    c *Client
}

func NewWSGVlanPoolingService (c *Client) *WSGVlanPoolingService {
    s := new(WSGVlanPoolingService)
    s.c = c
    return s
}

