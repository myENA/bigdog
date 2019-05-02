package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGPrecedenceProfileService struct {
    c *Client
}

func NewWSGPrecedenceProfileService (c *Client) *WSGPrecedenceProfileService {
    s := new(WSGPrecedenceProfileService)
    s.c = c
    return s
}

