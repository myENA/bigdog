package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
)

type WSGAdministrationService struct {
    c *Client
}

func NewWSGAdministrationService (c *Client) *WSGAdministrationService {
    s := new(WSGAdministrationService)
    s.c = c
    return s
}

