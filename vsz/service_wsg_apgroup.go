package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/apgroup"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zoneapmodel"
)

type WSGAPGroupService struct {
    c *Client
}

func NewWSGAPGroupService (c *Client) *WSGAPGroupService {
    s := new(WSGAPGroupService)
    s.c = c
    return s
}

