package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGGuestAccessService struct {
    client *Client
}

func NewWSGGuestAccessService (client *Client) *WSGGuestAccessService {
    s := new(WSGGuestAccessService)
    s.client = client
    return s
}

func (ss *WSGService) WSGGuestAccessService () *WSGGuestAccessService {
    serv := new(WSGGuestAccessService)
    serv.client = ss.client
    return serv
}

