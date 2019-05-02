package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGL3RoamingService struct {
    c *Client
}

func NewWSGL3RoamingService (c *Client) *WSGL3RoamingService {
    s := new(WSGL3RoamingService)
    s.c = c
    return s
}

