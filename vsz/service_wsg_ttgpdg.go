package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGTTGPDGService struct {
    c *Client
}

func NewWSGTTGPDGService (c *Client) *WSGTTGPDGService {
    s := new(WSGTTGPDGService)
    s.c = c
    return s
}

