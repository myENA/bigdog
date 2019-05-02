package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/eventmanagement"
)

type WSGEventManagementSettingService struct {
    c *Client
}

func NewWSGEventManagementSettingService (c *Client) *WSGEventManagementSettingService {
    s := new(WSGEventManagementSettingService)
    s.c = c
    return s
}

