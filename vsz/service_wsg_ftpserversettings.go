package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGFtpServerSettingsService struct {
    client *Client
}

func NewWSGFtpServerSettingsService (client *Client) *WSGFtpServerSettingsService {
    s := new(WSGFtpServerSettingsService)
    s.client = client
    return s
}

func (ss *WSGService) WSGFtpServerSettingsService () *WSGFtpServerSettingsService {
    serv := new(WSGFtpServerSettingsService)
    serv.client = ss.client
    return serv
}

