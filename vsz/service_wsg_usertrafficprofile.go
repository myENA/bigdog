package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGUserTrafficProfileService struct {
    client *Client
}

func NewWSGUserTrafficProfileService (client *Client) *WSGUserTrafficProfileService {
    s := new(WSGUserTrafficProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGUserTrafficProfileService () *WSGUserTrafficProfileService {
    serv := new(WSGUserTrafficProfileService)
    serv.client = ss.client
    return serv
}

