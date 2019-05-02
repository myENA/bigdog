package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ethernetport"
)

type WSGEthernetPortProfileService struct {
    c *Client
}

func NewWSGEthernetPortProfileService (c *Client) *WSGEthernetPortProfileService {
    s := new(WSGEthernetPortProfileService)
    s.c = c
    return s
}

