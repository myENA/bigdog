package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGControlPlanesService struct {
    c *Client
}

func NewWSGControlPlanesService (c *Client) *WSGControlPlanesService {
    s := new(WSGControlPlanesService)
    s.c = c
    return s
}

