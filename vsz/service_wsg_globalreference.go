package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGGlobalreferenceService struct {
    c *Client
}

func NewWSGGlobalreferenceService (c *Client) *WSGGlobalreferenceService {
    s := new(WSGGlobalreferenceService)
    s.c = c
    return s
}

