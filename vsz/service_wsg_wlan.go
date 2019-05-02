package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
	"github.com/myENA/ruckus-client/vsz/types/wsg/flexivpn"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlan"
)

type WSGWLANService struct {
    c *Client
}

func NewWSGWLANService (c *Client) *WSGWLANService {
    s := new(WSGWLANService)
    s.c = c
    return s
}

