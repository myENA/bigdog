package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dhcpmsgstats"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dhcppools"
)

type WSGDHCPService struct {
    c *Client
}

func NewWSGDHCPService (c *Client) *WSGDHCPService {
    s := new(WSGDHCPService)
    s.c = c
    return s
}

