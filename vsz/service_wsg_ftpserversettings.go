package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGFtpServerSettingsService struct {
    c *Client
}

func NewWSGFtpServerSettingsService (c *Client) *WSGFtpServerSettingsService {
    s := new(WSGFtpServerSettingsService)
    s.c = c
    return s
}

