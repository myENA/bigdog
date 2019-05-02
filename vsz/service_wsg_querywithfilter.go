package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaaserverquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clientquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
	"github.com/myENA/ruckus-client/vsz/types/wsg/indoormap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/meshneighborinfo"
	"github.com/myENA/ruckus-client/vsz/types/wsg/meshnodeinfo"
	"github.com/myENA/ruckus-client/vsz/types/wsg/rogueinfo"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlanquery"
)

type WSGQueryWithFilterService struct {
    c *Client
}

func NewWSGQueryWithFilterService (c *Client) *WSGQueryWithFilterService {
    s := new(WSGQueryWithFilterService)
    s.c = c
    return s
}

