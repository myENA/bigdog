package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGTrafficClassProfileService struct {
    client *Client
}

func NewWSGTrafficClassProfileService (client *Client) *WSGTrafficClassProfileService {
    s := new(WSGTrafficClassProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGTrafficClassProfileService () *WSGTrafficClassProfileService {
    serv := new(WSGTrafficClassProfileService)
    serv.client = ss.client
    return serv
}

