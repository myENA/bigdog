package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAccountingProfileService struct {
    client *Client
}

func NewWSGAccountingProfileService (client *Client) *WSGAccountingProfileService {
    s := new(WSGAccountingProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccountingProfileService () *WSGAccountingProfileService {
    serv := new(WSGAccountingProfileService)
    serv.client = ss.client
    return serv
}

