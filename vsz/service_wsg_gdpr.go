package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/gdpr"
)

type WSGGDPRService struct {
    c *Client
}

func NewWSGGDPRService (c *Client) *WSGGDPRService {
    s := new(WSGGDPRService)
    s.c = c
    return s
}

