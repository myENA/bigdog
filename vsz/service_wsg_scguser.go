package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserService struct {
    c *Client
}

func NewWSGSCGUserService (c *Client) *WSGSCGUserService {
    s := new(WSGSCGUserService)
    s.c = c
    return s
}

