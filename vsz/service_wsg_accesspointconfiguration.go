package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apmodel"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/meshnodeinfo"
)

type WSGAccessPointConfigurationService struct {
    client *Client
}

func NewWSGAccessPointConfigurationService (client *Client) *WSGAccessPointConfigurationService {
    s := new(WSGAccessPointConfigurationService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAccessPointConfigurationService () *WSGAccessPointConfigurationService {
    serv := new(WSGAccessPointConfigurationService)
    serv.client = ss.client
    return serv
}

