package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/meshnodeinfo"
)

type WSGAccessPointConfigurationService struct {
    c *Client
}

func NewWSGAccessPointConfigurationService (c *Client) *WSGAccessPointConfigurationService {
    s := new(WSGAccessPointConfigurationService)
    s.c = c
    return s
}

