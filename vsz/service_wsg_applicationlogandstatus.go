package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
)

type WSGApplicationLogAndStatusService struct {
    c *Client
}

func NewWSGApplicationLogAndStatusService (c *Client) *WSGApplicationLogAndStatusService {
    s := new(WSGApplicationLogAndStatusService)
    s.c = c
    return s
}

