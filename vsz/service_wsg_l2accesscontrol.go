package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGL2AccessControlService struct {
    c *Client
}

func NewWSGL2AccessControlService (c *Client) *WSGL2AccessControlService {
    s := new(WSGL2AccessControlService)
    s.c = c
    return s
}

