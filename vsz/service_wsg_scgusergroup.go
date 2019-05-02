package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserGroupService struct {
    client *Client
}

func NewWSGSCGUserGroupService (client *Client) *WSGSCGUserGroupService {
    s := new(WSGSCGUserGroupService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSCGUserGroupService () *WSGSCGUserGroupService {
    serv := new(WSGSCGUserGroupService)
    serv.client = ss.client
    return serv
}

