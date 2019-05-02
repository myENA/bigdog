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
    client *Client
}

func NewWSGQueryWithFilterService (client *Client) *WSGQueryWithFilterService {
    s := new(WSGQueryWithFilterService)
    s.client = client
    return s
}

func (ss *WSGService) WSGQueryWithFilterService () *WSGQueryWithFilterService {
    serv := new(WSGQueryWithFilterService)
    serv.client = ss.client
    return serv
}

