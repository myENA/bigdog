package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserGroupService struct {
    c *Client
}

func NewWSGSCGUserGroupService (c *Client) *WSGSCGUserGroupService {
    s := new(WSGSCGUserGroupService)
    s.c = c
    return s
}

