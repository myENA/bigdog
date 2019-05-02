package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGL2oGREService struct {
    c *Client
}

func NewWSGL2oGREService (c *Client) *WSGL2oGREService {
    s := new(WSGL2oGREService)
    s.c = c
    return s
}

