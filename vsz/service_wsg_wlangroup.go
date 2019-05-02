package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlangroup"
)

type WSGWLANGroupService struct {
    c *Client
}

func NewWSGWLANGroupService (c *Client) *WSGWLANGroupService {
    s := new(WSGWLANGroupService)
    s.c = c
    return s
}

