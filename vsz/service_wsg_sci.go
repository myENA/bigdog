package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/sci"
)

type WSGSCIService struct {
    c *Client
}

func NewWSGSCIService (c *Client) *WSGSCIService {
    s := new(WSGSCIService)
    s.c = c
    return s
}

