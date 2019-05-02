package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGAccountingProfileService struct {
    c *Client
}

func NewWSGAccountingProfileService (c *Client) *WSGAccountingProfileService {
    s := new(WSGAccountingProfileService)
    s.c = c
    return s
}

