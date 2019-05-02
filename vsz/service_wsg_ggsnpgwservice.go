package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGGGSNPGWServiceService struct {
    c *Client
}

func NewWSGGGSNPGWServiceService (c *Client) *WSGGGSNPGWServiceService {
    s := new(WSGGGSNPGWServiceService)
    s.c = c
    return s
}

