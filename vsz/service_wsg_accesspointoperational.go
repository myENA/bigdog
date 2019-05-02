package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAccessPointOperationalService struct {
    client *Client
}

func NewWSGAccessPointOperationalService (client *Client) *WSGAccessPointOperationalService {
    s := new(WSGAccessPointOperationalService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccessPointOperationalService () *WSGAccessPointOperationalService {
    serv := new(WSGAccessPointOperationalService)
    serv.client = ss.client
    return serv
}

